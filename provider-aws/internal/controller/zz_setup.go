/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	attachment "github.com/upbound/official-providers/provider-aws/internal/controller/autoscaling/attachment"
	autoscalinggroup "github.com/upbound/official-providers/provider-aws/internal/controller/autoscaling/autoscalinggroup"
	ebsvolume "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/ebsvolume"
	eip "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/eip"
	instance "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/instance"
	internetgateway "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/internetgateway"
	launchtemplate "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/launchtemplate"
	networkinterface "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/networkinterface"
	route "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/route"
	routetable "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/routetableassociation"
	securitygroup "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/securitygrouprule"
	subnet "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/subnet"
	transitgateway "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgateway"
	transitgatewayroute "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayvpcattachmentaccepter"
	vpc "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpc"
	vpcendpoint "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcendpoint"
	vpcipv4cidrblockassociation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcpeeringconnection"
	repository "github.com/upbound/official-providers/provider-aws/internal/controller/ecr/repository"
	repositoryecrpublic "github.com/upbound/official-providers/provider-aws/internal/controller/ecrpublic/repository"
	capacityprovider "github.com/upbound/official-providers/provider-aws/internal/controller/ecs/capacityprovider"
	cluster "github.com/upbound/official-providers/provider-aws/internal/controller/ecs/cluster"
	service "github.com/upbound/official-providers/provider-aws/internal/controller/ecs/service"
	taskdefinition "github.com/upbound/official-providers/provider-aws/internal/controller/ecs/taskdefinition"
	addon "github.com/upbound/official-providers/provider-aws/internal/controller/eks/addon"
	clustereks "github.com/upbound/official-providers/provider-aws/internal/controller/eks/cluster"
	fargateprofile "github.com/upbound/official-providers/provider-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/upbound/official-providers/provider-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/upbound/official-providers/provider-aws/internal/controller/eks/nodegroup"
	clusterelasticache "github.com/upbound/official-providers/provider-aws/internal/controller/elasticache/cluster"
	parametergroup "github.com/upbound/official-providers/provider-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/upbound/official-providers/provider-aws/internal/controller/elasticache/replicationgroup"
	subnetgroup "github.com/upbound/official-providers/provider-aws/internal/controller/elasticache/subnetgroup"
	user "github.com/upbound/official-providers/provider-aws/internal/controller/elasticache/user"
	usergroup "github.com/upbound/official-providers/provider-aws/internal/controller/elasticache/usergroup"
	lb "github.com/upbound/official-providers/provider-aws/internal/controller/elbv2/lb"
	lblistener "github.com/upbound/official-providers/provider-aws/internal/controller/elbv2/lblistener"
	lbtargetgroup "github.com/upbound/official-providers/provider-aws/internal/controller/elbv2/lbtargetgroup"
	lbtargetgroupattachment "github.com/upbound/official-providers/provider-aws/internal/controller/elbv2/lbtargetgroupattachment"
	accelerator "github.com/upbound/official-providers/provider-aws/internal/controller/globalaccelerator/accelerator"
	endpointgroup "github.com/upbound/official-providers/provider-aws/internal/controller/globalaccelerator/endpointgroup"
	listener "github.com/upbound/official-providers/provider-aws/internal/controller/globalaccelerator/listener"
	accesskey "github.com/upbound/official-providers/provider-aws/internal/controller/iam/accesskey"
	group "github.com/upbound/official-providers/provider-aws/internal/controller/iam/group"
	grouppolicyattachment "github.com/upbound/official-providers/provider-aws/internal/controller/iam/grouppolicyattachment"
	instanceprofile "github.com/upbound/official-providers/provider-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/upbound/official-providers/provider-aws/internal/controller/iam/openidconnectprovider"
	policy "github.com/upbound/official-providers/provider-aws/internal/controller/iam/policy"
	role "github.com/upbound/official-providers/provider-aws/internal/controller/iam/role"
	rolepolicyattachment "github.com/upbound/official-providers/provider-aws/internal/controller/iam/rolepolicyattachment"
	useriam "github.com/upbound/official-providers/provider-aws/internal/controller/iam/user"
	usergroupmembership "github.com/upbound/official-providers/provider-aws/internal/controller/iam/usergroupmembership"
	userpolicyattachment "github.com/upbound/official-providers/provider-aws/internal/controller/iam/userpolicyattachment"
	key "github.com/upbound/official-providers/provider-aws/internal/controller/kms/key"
	broker "github.com/upbound/official-providers/provider-aws/internal/controller/mq/broker"
	configuration "github.com/upbound/official-providers/provider-aws/internal/controller/mq/configuration"
	clusterneptune "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/cluster"
	clusterendpoint "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/clusterendpoint"
	clusterinstance "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/clusterinstance"
	clusterparametergroup "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/clusterparametergroup"
	clustersnapshot "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/clustersnapshot"
	eventsubscription "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/eventsubscription"
	parametergroupneptune "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/parametergroup"
	subnetgroupneptune "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/subnetgroup"
	providerconfig "github.com/upbound/official-providers/provider-aws/internal/controller/providerconfig"
	clusterrds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/cluster"
	instancerds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/instance"
	parametergrouprds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/parametergroup"
	subnetgrouprds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/subnetgroup"
	delegationset "github.com/upbound/official-providers/provider-aws/internal/controller/route53/delegationset"
	healthcheck "github.com/upbound/official-providers/provider-aws/internal/controller/route53/healthcheck"
	hostedzonednssec "github.com/upbound/official-providers/provider-aws/internal/controller/route53/hostedzonednssec"
	keysigningkey "github.com/upbound/official-providers/provider-aws/internal/controller/route53/keysigningkey"
	querylog "github.com/upbound/official-providers/provider-aws/internal/controller/route53/querylog"
	record "github.com/upbound/official-providers/provider-aws/internal/controller/route53/record"
	trafficpolicy "github.com/upbound/official-providers/provider-aws/internal/controller/route53/trafficpolicy"
	trafficpolicyinstance "github.com/upbound/official-providers/provider-aws/internal/controller/route53/trafficpolicyinstance"
	vpcassociationauthorization "github.com/upbound/official-providers/provider-aws/internal/controller/route53/vpcassociationauthorization"
	zone "github.com/upbound/official-providers/provider-aws/internal/controller/route53/zone"
	zoneassociation "github.com/upbound/official-providers/provider-aws/internal/controller/route53/zoneassociation"
	dnssecconfig "github.com/upbound/official-providers/provider-aws/internal/controller/route53resolver/dnssecconfig"
	endpoint "github.com/upbound/official-providers/provider-aws/internal/controller/route53resolver/endpoint"
	firewallconfig "github.com/upbound/official-providers/provider-aws/internal/controller/route53resolver/firewallconfig"
	firewalldomainlist "github.com/upbound/official-providers/provider-aws/internal/controller/route53resolver/firewalldomainlist"
	firewallrule "github.com/upbound/official-providers/provider-aws/internal/controller/route53resolver/firewallrule"
	firewallrulegroup "github.com/upbound/official-providers/provider-aws/internal/controller/route53resolver/firewallrulegroup"
	firewallrulegroupassociation "github.com/upbound/official-providers/provider-aws/internal/controller/route53resolver/firewallrulegroupassociation"
	querylogconfig "github.com/upbound/official-providers/provider-aws/internal/controller/route53resolver/querylogconfig"
	querylogconfigassociation "github.com/upbound/official-providers/provider-aws/internal/controller/route53resolver/querylogconfigassociation"
	rule "github.com/upbound/official-providers/provider-aws/internal/controller/route53resolver/rule"
	ruleassociation "github.com/upbound/official-providers/provider-aws/internal/controller/route53resolver/ruleassociation"
	bucket "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucket"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attachment.Setup,
		autoscalinggroup.Setup,
		ebsvolume.Setup,
		eip.Setup,
		instance.Setup,
		internetgateway.Setup,
		launchtemplate.Setup,
		networkinterface.Setup,
		route.Setup,
		routetable.Setup,
		routetableassociation.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		transitgateway.Setup,
		transitgatewayroute.Setup,
		transitgatewayroutetable.Setup,
		transitgatewayroutetableassociation.Setup,
		transitgatewayroutetablepropagation.Setup,
		transitgatewayvpcattachment.Setup,
		transitgatewayvpcattachmentaccepter.Setup,
		vpc.Setup,
		vpcendpoint.Setup,
		vpcipv4cidrblockassociation.Setup,
		vpcpeeringconnection.Setup,
		repository.Setup,
		repositoryecrpublic.Setup,
		capacityprovider.Setup,
		cluster.Setup,
		service.Setup,
		taskdefinition.Setup,
		addon.Setup,
		clustereks.Setup,
		fargateprofile.Setup,
		identityproviderconfig.Setup,
		nodegroup.Setup,
		clusterelasticache.Setup,
		parametergroup.Setup,
		replicationgroup.Setup,
		subnetgroup.Setup,
		user.Setup,
		usergroup.Setup,
		lb.Setup,
		lblistener.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		accelerator.Setup,
		endpointgroup.Setup,
		listener.Setup,
		accesskey.Setup,
		group.Setup,
		grouppolicyattachment.Setup,
		instanceprofile.Setup,
		openidconnectprovider.Setup,
		policy.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		useriam.Setup,
		usergroupmembership.Setup,
		userpolicyattachment.Setup,
		key.Setup,
		broker.Setup,
		configuration.Setup,
		clusterneptune.Setup,
		clusterendpoint.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clustersnapshot.Setup,
		eventsubscription.Setup,
		parametergroupneptune.Setup,
		subnetgroupneptune.Setup,
		providerconfig.Setup,
		clusterrds.Setup,
		instancerds.Setup,
		parametergrouprds.Setup,
		subnetgrouprds.Setup,
		delegationset.Setup,
		healthcheck.Setup,
		hostedzonednssec.Setup,
		keysigningkey.Setup,
		querylog.Setup,
		record.Setup,
		trafficpolicy.Setup,
		trafficpolicyinstance.Setup,
		vpcassociationauthorization.Setup,
		zone.Setup,
		zoneassociation.Setup,
		dnssecconfig.Setup,
		endpoint.Setup,
		firewallconfig.Setup,
		firewalldomainlist.Setup,
		firewallrule.Setup,
		firewallrulegroup.Setup,
		firewallrulegroupassociation.Setup,
		querylogconfig.Setup,
		querylogconfigassociation.Setup,
		rule.Setup,
		ruleassociation.Setup,
		bucket.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
