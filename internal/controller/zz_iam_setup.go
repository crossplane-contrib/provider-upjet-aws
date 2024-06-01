// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accesskey "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/accesskey"
	accountalias "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/accountalias"
	accountpasswordpolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/accountpasswordpolicy"
	group "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/group"
	groupmembership "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/groupmembership"
	grouppolicyattachment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/grouppolicyattachment"
	instanceprofile "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/openidconnectprovider"
	policy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/policy"
	role "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/role"
	rolepolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/rolepolicy"
	rolepolicyattachment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/rolepolicyattachment"
	samlprovider "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/samlprovider"
	servercertificate "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/servercertificate"
	servicelinkedrole "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/servicelinkedrole"
	servicespecificcredential "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/servicespecificcredential"
	signingcertificate "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/signingcertificate"
	user "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/user"
	usergroupmembership "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/usergroupmembership"
	userloginprofile "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/userloginprofile"
	userpolicyattachment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/userpolicyattachment"
	usersshkey "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/usersshkey"
	virtualmfadevice "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/iam/virtualmfadevice"
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
