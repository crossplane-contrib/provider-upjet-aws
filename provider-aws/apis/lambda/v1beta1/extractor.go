/*
Copyright 2022 Upbound Inc.
*/

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LambdaFunctionInvokeARN returns the invoke ARN value of the lambda function.
func LambdaFunctionInvokeARN() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		return func(mg metav1.Object) string {
			f, ok := mg.(*Function)
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
