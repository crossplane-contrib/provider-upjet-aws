package clusterauth

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/pkg/errors"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	clusterIDHeader = "x-k8s-aws-id"
	expireHeader    = "X-Amz-Expires"
	v1Prefix        = "k8s-aws-v1."

	errGetPresignGetCallerIdentity = "cannot get caller identity for presign"
	errDecodeCA                    = "cannot decode certificate authority data"
	errProduceKubeconfig           = "cannot produce kubeconfig"
)

func newPresignClient(cfg aws.Config, optFns ...func(*sts.Options)) *sts.PresignClient {
	cl := sts.NewFromConfig(cfg, optFns...)
	return sts.NewPresignClient(cl)
}

// GetConnectionDetails extracts managed.ConnectionDetails out of ekstypes.Cluster.
func GetConnectionDetails(ctx context.Context, stsClient *sts.PresignClient, cluster *ekstypes.Cluster, expiration time.Duration) (managed.ConnectionDetails, error) {
	getCallerIdentity, err := stsClient.PresignGetCallerIdentity(ctx, &sts.GetCallerIdentityInput{},
		func(po *sts.PresignOptions) {
			po.ClientOptions = []func(*sts.Options){
				sts.WithAPIOptions(
					smithyhttp.AddHeaderValue(clusterIDHeader, *cluster.Name),
					// Required to provide.
					// See https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html
					smithyhttp.AddHeaderValue(expireHeader, fmt.Sprintf("%d", int(expiration.Seconds()))),
				),
			}
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, errGetPresignGetCallerIdentity)
	}

	// More information: https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html
	token := v1Prefix + base64.RawURLEncoding.EncodeToString([]byte(getCallerIdentity.URL))

	// NOTE(hasheddan): We must decode the CA data before constructing our
	// Kubeconfig, as the raw Kubeconfig will be base64 encoded again when
	// written as a Secret.
	caData, err := base64.StdEncoding.DecodeString(aws.ToString(cluster.CertificateAuthority.Data))
	if err != nil {
		return managed.ConnectionDetails{}, errors.Wrap(err, errDecodeCA)
	}
	kc := clientcmdapi.Config{
		Clusters: map[string]*clientcmdapi.Cluster{
			*cluster.Name: {
				Server:                   *cluster.Endpoint,
				CertificateAuthorityData: caData,
			},
		},
		Contexts: map[string]*clientcmdapi.Context{
			*cluster.Name: {
				Cluster:  *cluster.Name,
				AuthInfo: *cluster.Name,
			},
		},
		AuthInfos: map[string]*clientcmdapi.AuthInfo{
			*cluster.Name: {
				Token: token,
			},
		},
		CurrentContext: *cluster.Name,
	}

	rawConfig, err := clientcmd.Write(kc)
	if err != nil {
		return managed.ConnectionDetails{}, errors.Wrap(err, errProduceKubeconfig)
	}
	return managed.ConnectionDetails{
		xpv1.ResourceCredentialsSecretEndpointKey:   []byte(*cluster.Endpoint),
		xpv1.ResourceCredentialsSecretKubeconfigKey: rawConfig,
		xpv1.ResourceCredentialsSecretCAKey:         caData,
	}, nil
}
