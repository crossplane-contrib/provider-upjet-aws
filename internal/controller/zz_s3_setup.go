/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	bucket "github.com/upbound/provider-aws/internal/controller/s3/bucket"
	bucketlifecycleconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketlifecycleconfiguration"
	bucketnotification "github.com/upbound/provider-aws/internal/controller/s3/bucketnotification"
)

// Setup_s3 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_s3(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketnotification.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
