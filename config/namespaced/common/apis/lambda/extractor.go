// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package lambda

import (
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/upbound/provider-aws/apis/namespaced/lambda/v1beta1"
)

// FunctionInvokeARN returns the invoke ARN value of the lambda function.
func FunctionInvokeARN() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		return func(mg metav1.Object) string {
			f, ok := mg.(*v1beta1.Function)
			if !ok {
				return ""
			}
			if !ok || f.Status.AtProvider.InvokeArn == nil {
				return ""
			}
			return *f.Status.AtProvider.InvokeArn
		}(mg)
	}
}
