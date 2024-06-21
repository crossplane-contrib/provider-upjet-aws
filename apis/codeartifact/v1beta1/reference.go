// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ExtractDomainName extracts the name from a Domain
func ExtractDomainName() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		return func(mg metav1.Object) string {
			f, ok := mg.(*Domain)
			if !ok {
				return ""
			}
			if !ok || f.Status.AtProvider.Domain == nil {
				return ""
			}
			return *f.Status.AtProvider.Domain
		}(mg)
	}
}

// ExtractRepositoryName extracts the name from a Domain
func ExtractRepositoryName() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		return func(mg metav1.Object) string {
			f, ok := mg.(*Repository)
			if !ok {
				return ""
			}
			if !ok || f.Status.AtProvider.Repository == nil {
				return ""
			}
			return *f.Status.AtProvider.Repository
		}(mg)
	}
}
