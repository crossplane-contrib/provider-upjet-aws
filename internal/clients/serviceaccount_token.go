// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"os"
	"time"

	authv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/pkg/errors"

	"github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

const (
	// tokenExpirationSeconds is the requested duration of the token.
	// AWS recommends tokens with at least 15 minutes (900 seconds) validity.
	// We request 1 hour to provide a buffer.
	tokenExpirationSeconds = 3600

	// defaultTokenAudience is the default intended audience for the token.
	// For AWS IRSA, this should be "sts.amazonaws.com"
	defaultTokenAudience = "sts.amazonaws.com"

	// tokenAudienceEnvVar is the environment variable name for customizing the token audience.
	// This is useful when running the provider in non-AWS environments (GCP, Azure, on-prem)
	// where the OIDC token audience needs to be different.
	tokenAudienceEnvVar = "SERVICE_ACCOUNT_TOKEN_AUDIENCE"
)

// ServiceAccountTokenRetriever retrieves tokens for ServiceAccounts
type ServiceAccountTokenRetriever struct {
	kube client.Client
}

// NewServiceAccountTokenRetriever creates a new ServiceAccountTokenRetriever
func NewServiceAccountTokenRetriever(kube client.Client) *ServiceAccountTokenRetriever {
	return &ServiceAccountTokenRetriever{
		kube: kube,
	}
}

// getTokenAudience returns the token audience, checking the environment variable first,
// then falling back to the default AWS audience.
func getTokenAudience() string {
	if audience := os.Getenv(tokenAudienceEnvVar); audience != "" {
		return audience
	}
	return defaultTokenAudience
}

// GetToken creates a TokenRequest for the specified ServiceAccount and returns the token
func (r *ServiceAccountTokenRetriever) GetToken(ctx context.Context, namespace, name string) (string, error) {
	if namespace == "" {
		return "", errors.New("namespace cannot be empty")
	}
	if name == "" {
		return "", errors.New("service account name cannot be empty")
	}

	expirationSeconds := int64(tokenExpirationSeconds)
	tokenRequest := &authv1.TokenRequest{
		Spec: authv1.TokenRequestSpec{
			Audiences:         []string{getTokenAudience()},
			ExpirationSeconds: &expirationSeconds,
		},
	}

	// Use SubResource to create the token request
	sa := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	if err := r.kube.SubResource("token").Create(ctx, sa, tokenRequest); err != nil {
		return "", errors.Wrapf(err, "failed to create token request for ServiceAccount %s/%s", namespace, name)
	}

	if tokenRequest.Status.Token == "" {
		return "", errors.Errorf("received empty token for ServiceAccount %s/%s", namespace, name)
	}

	return tokenRequest.Status.Token, nil
}

// serviceAccountTokenIdentityRetriever implements stscreds.IdentityTokenRetriever
// for retrieving ServiceAccount tokens
type serviceAccountTokenIdentityRetriever struct {
	retriever *ServiceAccountTokenRetriever
	namespace string
	name      string
}

// GetIdentityToken retrieves the ServiceAccount token.
// This method creates a fresh context with timeout for each token request to ensure
// it's not tied to the lifecycle of any particular reconciliation. This is critical
// because AWS SDK calls this method to refresh credentials long after the original
// reconciliation context may have been canceled.
func (r *serviceAccountTokenIdentityRetriever) GetIdentityToken() ([]byte, error) {
	// Create a bounded context for this specific token request.
	// Token requests are typically fast (<1s) but we allow 30s for API server delays.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token, err := r.retriever.GetToken(ctx, r.namespace, r.name)
	if err != nil {
		return nil, err
	}
	return []byte(token), nil
}

// NewServiceAccountTokenIdentityRetriever creates an identity token retriever
// for the specified ServiceAccount.
func NewServiceAccountTokenIdentityRetriever(ctx context.Context, retriever *ServiceAccountTokenRetriever, namespace, name string) *serviceAccountTokenIdentityRetriever {
	return &serviceAccountTokenIdentityRetriever{
		retriever: retriever,
		namespace: namespace,
		name:      name,
	}
}

// ResolveServiceAccountNamespace resolves the namespace for a ServiceAccount reference.
// For namespace-scoped ProviderConfigs: ALWAYS uses the ProviderConfig's own namespace
// (ignores any namespace field in ServiceAccountReference for security).
// For ClusterProviderConfigs: requires the namespace field in ServiceAccountReference.
func ResolveServiceAccountNamespace(pc *v1beta1.ClusterProviderConfig, saRef *v1beta1.ServiceAccountReference) (string, error) {
	if saRef == nil {
		return "", errors.New("serviceAccountRef cannot be nil")
	}

	// For namespace-scoped ProviderConfig: ALWAYS use the ProviderConfig's namespace
	// This prevents cross-namespace ServiceAccount access even if someone specifies
	// a different namespace in the ServiceAccountReference (security enforcement).
	// Note: pc.Namespace is set in pc_resolver.go:150 when converting ProviderConfig to ClusterProviderConfig
	if pc.Namespace != "" {
		return pc.Namespace, nil
	}

	// For ClusterProviderConfig (pc.Namespace == ""), require explicit namespace in ServiceAccountReference
	if saRef.Namespace == "" {
		return "", errors.Errorf("ClusterProviderConfig %s: serviceAccountRef.namespace is required", pc.Name)
	}
	return saRef.Namespace, nil
}
