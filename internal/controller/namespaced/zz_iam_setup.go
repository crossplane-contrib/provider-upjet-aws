// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accesskey "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/accesskey"
	accountalias "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/accountalias"
	accountpasswordpolicy "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/accountpasswordpolicy"
	group "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/group"
	groupmembership "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/groupmembership"
	grouppolicyattachment "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/grouppolicyattachment"
	instanceprofile "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/instanceprofile"
	openidconnectprovider "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/openidconnectprovider"
	policy "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/policy"
	role "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/role"
	rolepolicy "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/rolepolicy"
	rolepolicyattachment "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/rolepolicyattachment"
	samlprovider "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/samlprovider"
	servercertificate "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/servercertificate"
	servicelinkedrole "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/servicelinkedrole"
	servicespecificcredential "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/servicespecificcredential"
	signingcertificate "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/signingcertificate"
	user "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/user"
	usergroupmembership "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/usergroupmembership"
	userloginprofile "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/userloginprofile"
	userpolicyattachment "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/userpolicyattachment"
	usersshkey "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/usersshkey"
	virtualmfadevice "github.com/upbound/provider-aws/v2/internal/controller/namespaced/iam/virtualmfadevice"
)

// Setup_iam creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_iam(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		group.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		instanceprofile.Setup,
		openidconnectprovider.Setup,
		policy.Setup,
		role.Setup,
		rolepolicy.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		servercertificate.Setup,
		servicelinkedrole.Setup,
		servicespecificcredential.Setup,
		signingcertificate.Setup,
		user.Setup,
		usergroupmembership.Setup,
		userloginprofile.Setup,
		userpolicyattachment.Setup,
		usersshkey.Setup,
		virtualmfadevice.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_iam creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_iam(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesskey.SetupGated,
		accountalias.SetupGated,
		accountpasswordpolicy.SetupGated,
		group.SetupGated,
		groupmembership.SetupGated,
		grouppolicyattachment.SetupGated,
		instanceprofile.SetupGated,
		openidconnectprovider.SetupGated,
		policy.SetupGated,
		role.SetupGated,
		rolepolicy.SetupGated,
		rolepolicyattachment.SetupGated,
		samlprovider.SetupGated,
		servercertificate.SetupGated,
		servicelinkedrole.SetupGated,
		servicespecificcredential.SetupGated,
		signingcertificate.SetupGated,
		user.SetupGated,
		usergroupmembership.SetupGated,
		userloginprofile.SetupGated,
		userpolicyattachment.SetupGated,
		usersshkey.SetupGated,
		virtualmfadevice.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
