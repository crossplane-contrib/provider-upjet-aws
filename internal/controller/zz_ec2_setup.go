// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	ami "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/ami"
	amicopy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/amicopy"
	amilaunchpermission "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/amilaunchpermission"
	availabilityzonegroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/availabilityzonegroup"
	capacityreservation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/capacityreservation"
	carriergateway "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/carriergateway"
	customergateway "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/customergateway"
	defaultnetworkacl "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/defaultnetworkacl"
	defaultroutetable "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/defaultroutetable"
	defaultsecuritygroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/defaultsecuritygroup"
	defaultsubnet "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/defaultsubnet"
	defaultvpc "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/defaultvpc"
	defaultvpcdhcpoptions "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/defaultvpcdhcpoptions"
	ebsdefaultkmskey "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/ebsdefaultkmskey"
	ebsencryptionbydefault "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/ebsencryptionbydefault"
	ebssnapshot "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/ebssnapshot"
	ebssnapshotcopy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/ebssnapshotcopy"
	ebssnapshotimport "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/ebssnapshotimport"
	ebsvolume "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/ebsvolume"
	egressonlyinternetgateway "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/egressonlyinternetgateway"
	eip "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/eip"
	eipassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/eipassociation"
	flowlog "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/flowlog"
	host "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/host"
	instance "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/instance"
	instancestate "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/instancestate"
	internetgateway "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/internetgateway"
	keypair "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/keypair"
	launchtemplate "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/launchtemplate"
	mainroutetableassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/mainroutetableassociation"
	managedprefixlist "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/managedprefixlist"
	managedprefixlistentry "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/managedprefixlistentry"
	natgateway "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/natgateway"
	networkacl "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/networkacl"
	networkaclrule "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/networkaclrule"
	networkinsightsanalysis "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/networkinsightsanalysis"
	networkinsightspath "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/networkinsightspath"
	networkinterface "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/networkinterface"
	networkinterfaceattachment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/networkinterfaceattachment"
	networkinterfacesgattachment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/networkinterfacesgattachment"
	placementgroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/placementgroup"
	route "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/route"
	routetable "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/routetableassociation"
	securitygroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/securitygroup"
	securitygroupegressrule "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/securitygroupegressrule"
	securitygroupingressrule "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/securitygroupingressrule"
	securitygrouprule "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/securitygrouprule"
	serialconsoleaccess "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/serialconsoleaccess"
	snapshotcreatevolumepermission "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/snapshotcreatevolumepermission"
	spotdatafeedsubscription "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/spotdatafeedsubscription"
	spotfleetrequest "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/spotfleetrequest"
	spotinstancerequest "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/spotinstancerequest"
	subnet "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/subnet"
	subnetcidrreservation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/subnetcidrreservation"
	tag "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/tag"
	trafficmirrorfilter "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/trafficmirrorfilter"
	trafficmirrorfilterrule "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/trafficmirrorfilterrule"
	transitgateway "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgateway"
	transitgatewayconnect "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewayconnect"
	transitgatewayconnectpeer "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewayconnectpeer"
	transitgatewaymulticastdomain "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewaymulticastdomain"
	transitgatewaymulticastdomainassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewaymulticastdomainassociation"
	transitgatewaymulticastgroupmember "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewaymulticastgroupmember"
	transitgatewaymulticastgroupsource "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewaymulticastgroupsource"
	transitgatewaypeeringattachment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewaypeeringattachment"
	transitgatewaypeeringattachmentaccepter "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewaypeeringattachmentaccepter"
	transitgatewaypolicytable "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewaypolicytable"
	transitgatewayprefixlistreference "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewayprefixlistreference"
	transitgatewayroute "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/transitgatewayvpcattachmentaccepter"
	volumeattachment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/volumeattachment"
	vpc "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpc"
	vpcdhcpoptions "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcdhcpoptions"
	vpcdhcpoptionsassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcdhcpoptionsassociation"
	vpcendpoint "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcendpoint"
	vpcendpointconnectionnotification "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcendpointconnectionnotification"
	vpcendpointroutetableassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcendpointroutetableassociation"
	vpcendpointsecuritygroupassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcendpointsecuritygroupassociation"
	vpcendpointservice "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcendpointservice"
	vpcendpointserviceallowedprincipal "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcendpointserviceallowedprincipal"
	vpcendpointsubnetassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcendpointsubnetassociation"
	vpcipam "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcipam"
	vpcipampool "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcipampool"
	vpcipampoolcidr "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcipampoolcidr"
	vpcipampoolcidrallocation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcipampoolcidrallocation"
	vpcipamscope "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcipamscope"
	vpcipv4cidrblockassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcpeeringconnection"
	vpcpeeringconnectionaccepter "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcpeeringconnectionaccepter"
	vpcpeeringconnectionoptions "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpcpeeringconnectionoptions"
	vpnconnection "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpnconnection"
	vpnconnectionroute "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpnconnectionroute"
	vpngateway "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpngateway"
	vpngatewayattachment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpngatewayattachment"
	vpngatewayroutepropagation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ec2/vpngatewayroutepropagation"
)

// Setup_ec2 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ec2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		ami.Setup,
		amicopy.Setup,
		amilaunchpermission.Setup,
		availabilityzonegroup.Setup,
		capacityreservation.Setup,
		carriergateway.Setup,
		customergateway.Setup,
		defaultnetworkacl.Setup,
		defaultroutetable.Setup,
		defaultsecuritygroup.Setup,
		defaultsubnet.Setup,
		defaultvpc.Setup,
		defaultvpcdhcpoptions.Setup,
		ebsdefaultkmskey.Setup,
		ebsencryptionbydefault.Setup,
		ebssnapshot.Setup,
		ebssnapshotcopy.Setup,
		ebssnapshotimport.Setup,
		ebsvolume.Setup,
		egressonlyinternetgateway.Setup,
		eip.Setup,
		eipassociation.Setup,
		flowlog.Setup,
		host.Setup,
		instance.Setup,
		instancestate.Setup,
		internetgateway.Setup,
		keypair.Setup,
		launchtemplate.Setup,
		mainroutetableassociation.Setup,
		managedprefixlist.Setup,
		managedprefixlistentry.Setup,
		natgateway.Setup,
		networkacl.Setup,
		networkaclrule.Setup,
		networkinsightsanalysis.Setup,
		networkinsightspath.Setup,
		networkinterface.Setup,
		networkinterfaceattachment.Setup,
		networkinterfacesgattachment.Setup,
		placementgroup.Setup,
		route.Setup,
		routetable.Setup,
		routetableassociation.Setup,
		securitygroup.Setup,
		securitygroupegressrule.Setup,
		securitygroupingressrule.Setup,
		securitygrouprule.Setup,
		serialconsoleaccess.Setup,
		snapshotcreatevolumepermission.Setup,
		spotdatafeedsubscription.Setup,
		spotfleetrequest.Setup,
		spotinstancerequest.Setup,
		subnet.Setup,
		subnetcidrreservation.Setup,
		tag.Setup,
		trafficmirrorfilter.Setup,
		trafficmirrorfilterrule.Setup,
		transitgateway.Setup,
		transitgatewayconnect.Setup,
		transitgatewayconnectpeer.Setup,
		transitgatewaymulticastdomain.Setup,
		transitgatewaymulticastdomainassociation.Setup,
		transitgatewaymulticastgroupmember.Setup,
		transitgatewaymulticastgroupsource.Setup,
		transitgatewaypeeringattachment.Setup,
		transitgatewaypeeringattachmentaccepter.Setup,
		transitgatewaypolicytable.Setup,
		transitgatewayprefixlistreference.Setup,
		transitgatewayroute.Setup,
		transitgatewayroutetable.Setup,
		transitgatewayroutetableassociation.Setup,
		transitgatewayroutetablepropagation.Setup,
		transitgatewayvpcattachment.Setup,
		transitgatewayvpcattachmentaccepter.Setup,
		volumeattachment.Setup,
		vpc.Setup,
		vpcdhcpoptions.Setup,
		vpcdhcpoptionsassociation.Setup,
		vpcendpoint.Setup,
		vpcendpointconnectionnotification.Setup,
		vpcendpointroutetableassociation.Setup,
		vpcendpointsecuritygroupassociation.Setup,
		vpcendpointservice.Setup,
		vpcendpointserviceallowedprincipal.Setup,
		vpcendpointsubnetassociation.Setup,
		vpcipam.Setup,
		vpcipampool.Setup,
		vpcipampoolcidr.Setup,
		vpcipampoolcidrallocation.Setup,
		vpcipamscope.Setup,
		vpcipv4cidrblockassociation.Setup,
		vpcpeeringconnection.Setup,
		vpcpeeringconnectionaccepter.Setup,
		vpcpeeringconnectionoptions.Setup,
		vpnconnection.Setup,
		vpnconnectionroute.Setup,
		vpngateway.Setup,
		vpngatewayattachment.Setup,
		vpngatewayroutepropagation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
