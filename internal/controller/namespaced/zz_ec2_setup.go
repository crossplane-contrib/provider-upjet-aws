// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	ami "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/ami"
	amicopy "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/amicopy"
	amilaunchpermission "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/amilaunchpermission"
	availabilityzonegroup "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/availabilityzonegroup"
	capacityreservation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/capacityreservation"
	carriergateway "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/carriergateway"
	customergateway "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/customergateway"
	defaultnetworkacl "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/defaultnetworkacl"
	defaultroutetable "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/defaultroutetable"
	defaultsecuritygroup "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/defaultsecuritygroup"
	defaultsubnet "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/defaultsubnet"
	defaultvpc "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/defaultvpc"
	defaultvpcdhcpoptions "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/defaultvpcdhcpoptions"
	ebsdefaultkmskey "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/ebsdefaultkmskey"
	ebsencryptionbydefault "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/ebsencryptionbydefault"
	ebssnapshot "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/ebssnapshot"
	ebssnapshotcopy "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/ebssnapshotcopy"
	ebssnapshotimport "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/ebssnapshotimport"
	ebsvolume "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/ebsvolume"
	egressonlyinternetgateway "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/egressonlyinternetgateway"
	eip "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/eip"
	eipassociation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/eipassociation"
	fleet "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/fleet"
	flowlog "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/flowlog"
	host "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/host"
	instance "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/instance"
	instancestate "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/instancestate"
	internetgateway "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/internetgateway"
	keypair "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/keypair"
	launchtemplate "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/launchtemplate"
	mainroutetableassociation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/mainroutetableassociation"
	managedprefixlist "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/managedprefixlist"
	managedprefixlistentry "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/managedprefixlistentry"
	natgateway "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/natgateway"
	networkacl "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/networkacl"
	networkaclrule "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/networkaclrule"
	networkinsightsanalysis "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/networkinsightsanalysis"
	networkinsightspath "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/networkinsightspath"
	networkinterface "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/networkinterface"
	networkinterfaceattachment "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/networkinterfaceattachment"
	networkinterfacesgattachment "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/networkinterfacesgattachment"
	placementgroup "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/placementgroup"
	route "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/route"
	routetable "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/routetable"
	routetableassociation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/routetableassociation"
	securitygroup "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/securitygroup"
	securitygroupegressrule "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/securitygroupegressrule"
	securitygroupingressrule "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/securitygroupingressrule"
	securitygrouprule "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/securitygrouprule"
	serialconsoleaccess "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/serialconsoleaccess"
	snapshotcreatevolumepermission "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/snapshotcreatevolumepermission"
	spotdatafeedsubscription "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/spotdatafeedsubscription"
	spotfleetrequest "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/spotfleetrequest"
	spotinstancerequest "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/spotinstancerequest"
	subnet "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/subnet"
	subnetcidrreservation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/subnetcidrreservation"
	tag "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/tag"
	trafficmirrorfilter "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/trafficmirrorfilter"
	trafficmirrorfilterrule "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/trafficmirrorfilterrule"
	transitgateway "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgateway"
	transitgatewayconnect "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewayconnect"
	transitgatewayconnectpeer "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewayconnectpeer"
	transitgatewaymulticastdomain "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewaymulticastdomain"
	transitgatewaymulticastdomainassociation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewaymulticastdomainassociation"
	transitgatewaymulticastgroupmember "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewaymulticastgroupmember"
	transitgatewaymulticastgroupsource "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewaymulticastgroupsource"
	transitgatewaypeeringattachment "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewaypeeringattachment"
	transitgatewaypeeringattachmentaccepter "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewaypeeringattachmentaccepter"
	transitgatewaypolicytable "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewaypolicytable"
	transitgatewayprefixlistreference "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewayprefixlistreference"
	transitgatewayroute "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/transitgatewayvpcattachmentaccepter"
	volumeattachment "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/volumeattachment"
	vpc "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpc"
	vpcdhcpoptions "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcdhcpoptions"
	vpcdhcpoptionsassociation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcdhcpoptionsassociation"
	vpcendpoint "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcendpoint"
	vpcendpointconnectionnotification "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcendpointconnectionnotification"
	vpcendpointroutetableassociation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcendpointroutetableassociation"
	vpcendpointsecuritygroupassociation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcendpointsecuritygroupassociation"
	vpcendpointservice "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcendpointservice"
	vpcendpointserviceallowedprincipal "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcendpointserviceallowedprincipal"
	vpcendpointsubnetassociation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcendpointsubnetassociation"
	vpcipam "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcipam"
	vpcipampool "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcipampool"
	vpcipampoolcidr "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcipampoolcidr"
	vpcipampoolcidrallocation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcipampoolcidrallocation"
	vpcipamscope "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcipamscope"
	vpcipv4cidrblockassociation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcpeeringconnection"
	vpcpeeringconnectionaccepter "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcpeeringconnectionaccepter"
	vpcpeeringconnectionoptions "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpcpeeringconnectionoptions"
	vpnconnection "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpnconnection"
	vpnconnectionroute "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpnconnectionroute"
	vpngateway "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpngateway"
	vpngatewayattachment "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpngatewayattachment"
	vpngatewayroutepropagation "github.com/upbound/provider-aws/internal/controller/namespaced/ec2/vpngatewayroutepropagation"
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
		fleet.Setup,
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

// SetupGated_ec2 creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ec2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		ami.SetupGated,
		amicopy.SetupGated,
		amilaunchpermission.SetupGated,
		availabilityzonegroup.SetupGated,
		capacityreservation.SetupGated,
		carriergateway.SetupGated,
		customergateway.SetupGated,
		defaultnetworkacl.SetupGated,
		defaultroutetable.SetupGated,
		defaultsecuritygroup.SetupGated,
		defaultsubnet.SetupGated,
		defaultvpc.SetupGated,
		defaultvpcdhcpoptions.SetupGated,
		ebsdefaultkmskey.SetupGated,
		ebsencryptionbydefault.SetupGated,
		ebssnapshot.SetupGated,
		ebssnapshotcopy.SetupGated,
		ebssnapshotimport.SetupGated,
		ebsvolume.SetupGated,
		egressonlyinternetgateway.SetupGated,
		eip.SetupGated,
		eipassociation.SetupGated,
		fleet.SetupGated,
		flowlog.SetupGated,
		host.SetupGated,
		instance.SetupGated,
		instancestate.SetupGated,
		internetgateway.SetupGated,
		keypair.SetupGated,
		launchtemplate.SetupGated,
		mainroutetableassociation.SetupGated,
		managedprefixlist.SetupGated,
		managedprefixlistentry.SetupGated,
		natgateway.SetupGated,
		networkacl.SetupGated,
		networkaclrule.SetupGated,
		networkinsightsanalysis.SetupGated,
		networkinsightspath.SetupGated,
		networkinterface.SetupGated,
		networkinterfaceattachment.SetupGated,
		networkinterfacesgattachment.SetupGated,
		placementgroup.SetupGated,
		route.SetupGated,
		routetable.SetupGated,
		routetableassociation.SetupGated,
		securitygroup.SetupGated,
		securitygroupegressrule.SetupGated,
		securitygroupingressrule.SetupGated,
		securitygrouprule.SetupGated,
		serialconsoleaccess.SetupGated,
		snapshotcreatevolumepermission.SetupGated,
		spotdatafeedsubscription.SetupGated,
		spotfleetrequest.SetupGated,
		spotinstancerequest.SetupGated,
		subnet.SetupGated,
		subnetcidrreservation.SetupGated,
		tag.SetupGated,
		trafficmirrorfilter.SetupGated,
		trafficmirrorfilterrule.SetupGated,
		transitgateway.SetupGated,
		transitgatewayconnect.SetupGated,
		transitgatewayconnectpeer.SetupGated,
		transitgatewaymulticastdomain.SetupGated,
		transitgatewaymulticastdomainassociation.SetupGated,
		transitgatewaymulticastgroupmember.SetupGated,
		transitgatewaymulticastgroupsource.SetupGated,
		transitgatewaypeeringattachment.SetupGated,
		transitgatewaypeeringattachmentaccepter.SetupGated,
		transitgatewaypolicytable.SetupGated,
		transitgatewayprefixlistreference.SetupGated,
		transitgatewayroute.SetupGated,
		transitgatewayroutetable.SetupGated,
		transitgatewayroutetableassociation.SetupGated,
		transitgatewayroutetablepropagation.SetupGated,
		transitgatewayvpcattachment.SetupGated,
		transitgatewayvpcattachmentaccepter.SetupGated,
		volumeattachment.SetupGated,
		vpc.SetupGated,
		vpcdhcpoptions.SetupGated,
		vpcdhcpoptionsassociation.SetupGated,
		vpcendpoint.SetupGated,
		vpcendpointconnectionnotification.SetupGated,
		vpcendpointroutetableassociation.SetupGated,
		vpcendpointsecuritygroupassociation.SetupGated,
		vpcendpointservice.SetupGated,
		vpcendpointserviceallowedprincipal.SetupGated,
		vpcendpointsubnetassociation.SetupGated,
		vpcipam.SetupGated,
		vpcipampool.SetupGated,
		vpcipampoolcidr.SetupGated,
		vpcipampoolcidrallocation.SetupGated,
		vpcipamscope.SetupGated,
		vpcipv4cidrblockassociation.SetupGated,
		vpcpeeringconnection.SetupGated,
		vpcpeeringconnectionaccepter.SetupGated,
		vpcpeeringconnectionoptions.SetupGated,
		vpnconnection.SetupGated,
		vpnconnectionroute.SetupGated,
		vpngateway.SetupGated,
		vpngatewayattachment.SetupGated,
		vpngatewayroutepropagation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
