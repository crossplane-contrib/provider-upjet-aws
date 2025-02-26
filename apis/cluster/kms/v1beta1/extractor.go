// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

// KMSKeyARN returns an extractor that returns ARN of Key.
func KMSKeyARN() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		return func(mg metav1.Object) string {
			key, ok := mg.(*Key)
			if !ok {
				return ""
			}
			return ptr.Deref(key.Status.AtProvider.Arn, "")
		}(mg)
	}
}
