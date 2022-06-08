/*
Copyright 2021 Upbound Inc.
*/

package v1alpha2

import (
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

// KMSKeyARN returns an extractor that returns ARN of Key.
func KMSKeyARN() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		return func(mg metav1.Object) string {
			key, ok := mg.(*Key)
			if !ok {
				return ""
			}
			return pointer.StringDeref(key.Status.AtProvider.Arn, "")
		}(mg)
	}
}
