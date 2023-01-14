/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	eip "github.com/upbound/provider-aws/internal/controller/ec2/eip"
	instance "github.com/upbound/provider-aws/internal/controller/ec2/instance"
	internetgateway "github.com/upbound/provider-aws/internal/controller/ec2/internetgateway"
	managedprefixlist "github.com/upbound/provider-aws/internal/controller/ec2/managedprefixlist"
	natgateway "github.com/upbound/provider-aws/internal/controller/ec2/natgateway"
	networkinterface "github.com/upbound/provider-aws/internal/controller/ec2/networkinterface"
	route "github.com/upbound/provider-aws/internal/controller/ec2/route"
	routetable "github.com/upbound/provider-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/routetableassociation"
	securitygroup "github.com/upbound/provider-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/upbound/provider-aws/internal/controller/ec2/securitygrouprule"
	subnet "github.com/upbound/provider-aws/internal/controller/ec2/subnet"
	transitgateway "github.com/upbound/provider-aws/internal/controller/ec2/transitgateway"
	transitgatewaypeeringattachment "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaypeeringattachment"
	transitgatewaypeeringattachmentaccepter "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaypeeringattachmentaccepter"
	transitgatewayroutetable "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayvpcattachment"
	vpc "github.com/upbound/provider-aws/internal/controller/ec2/vpc"
	vpcendpoint "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpoint"
	vpcipv4cidrblockassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/upbound/provider-aws/internal/controller/ec2/vpcpeeringconnection"
)

// Setup_ec2 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ec2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		eip.Setup,
		instance.Setup,
		internetgateway.Setup,
		managedprefixlist.Setup,
		natgateway.Setup,
		networkinterface.Setup,
		route.Setup,
		routetable.Setup,
		routetableassociation.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		transitgateway.Setup,
		transitgatewaypeeringattachment.Setup,
		transitgatewaypeeringattachmentaccepter.Setup,
		transitgatewayroutetable.Setup,
		transitgatewayroutetableassociation.Setup,
		transitgatewayroutetablepropagation.Setup,
		transitgatewayvpcattachment.Setup,
		vpc.Setup,
		vpcendpoint.Setup,
		vpcipv4cidrblockassociation.Setup,
		vpcpeeringconnection.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
