/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	database "github.com/upbound/provider-aws/internal/controller/athena/database"
	datacatalog "github.com/upbound/provider-aws/internal/controller/athena/datacatalog"
	namedquery "github.com/upbound/provider-aws/internal/controller/athena/namedquery"
	workgroup "github.com/upbound/provider-aws/internal/controller/athena/workgroup"
	table "github.com/upbound/provider-aws/internal/controller/dynamodb/table"
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
	repository "github.com/upbound/provider-aws/internal/controller/ecr/repository"
	repositorypolicy "github.com/upbound/provider-aws/internal/controller/ecr/repositorypolicy"
	addon "github.com/upbound/provider-aws/internal/controller/eks/addon"
	cluster "github.com/upbound/provider-aws/internal/controller/eks/cluster"
	clusterauth "github.com/upbound/provider-aws/internal/controller/eks/clusterauth"
	nodegroup "github.com/upbound/provider-aws/internal/controller/eks/nodegroup"
	catalogdatabase "github.com/upbound/provider-aws/internal/controller/glue/catalogdatabase"
	crawler "github.com/upbound/provider-aws/internal/controller/glue/crawler"
	accesskey "github.com/upbound/provider-aws/internal/controller/iam/accesskey"
	openidconnectprovider "github.com/upbound/provider-aws/internal/controller/iam/openidconnectprovider"
	policy "github.com/upbound/provider-aws/internal/controller/iam/policy"
	role "github.com/upbound/provider-aws/internal/controller/iam/role"
	rolepolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/rolepolicyattachment"
	user "github.com/upbound/provider-aws/internal/controller/iam/user"
	userpolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/userpolicyattachment"
	key "github.com/upbound/provider-aws/internal/controller/kms/key"
	alias "github.com/upbound/provider-aws/internal/controller/lambda/alias"
	function "github.com/upbound/provider-aws/internal/controller/lambda/function"
	functioneventinvokeconfig "github.com/upbound/provider-aws/internal/controller/lambda/functioneventinvokeconfig"
	permission "github.com/upbound/provider-aws/internal/controller/lambda/permission"
	providerconfig "github.com/upbound/provider-aws/internal/controller/providerconfig"
	instancerds "github.com/upbound/provider-aws/internal/controller/rds/instance"
	optiongroup "github.com/upbound/provider-aws/internal/controller/rds/optiongroup"
	parametergroup "github.com/upbound/provider-aws/internal/controller/rds/parametergroup"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/rds/subnetgroup"
	delegationset "github.com/upbound/provider-aws/internal/controller/route53/delegationset"
	healthcheck "github.com/upbound/provider-aws/internal/controller/route53/healthcheck"
	record "github.com/upbound/provider-aws/internal/controller/route53/record"
	zone "github.com/upbound/provider-aws/internal/controller/route53/zone"
	bucket "github.com/upbound/provider-aws/internal/controller/s3/bucket"
	bucketlifecycleconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketlifecycleconfiguration"
	bucketnotification "github.com/upbound/provider-aws/internal/controller/s3/bucketnotification"
	topic "github.com/upbound/provider-aws/internal/controller/sns/topic"
	queue "github.com/upbound/provider-aws/internal/controller/sqs/queue"
	server "github.com/upbound/provider-aws/internal/controller/transfer/server"
	sshkey "github.com/upbound/provider-aws/internal/controller/transfer/sshkey"
	usertransfer "github.com/upbound/provider-aws/internal/controller/transfer/user"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		datacatalog.Setup,
		namedquery.Setup,
		workgroup.Setup,
		table.Setup,
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
		repository.Setup,
		repositorypolicy.Setup,
		addon.Setup,
		cluster.Setup,
		clusterauth.Setup,
		nodegroup.Setup,
		catalogdatabase.Setup,
		crawler.Setup,
		accesskey.Setup,
		openidconnectprovider.Setup,
		policy.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		user.Setup,
		userpolicyattachment.Setup,
		key.Setup,
		alias.Setup,
		function.Setup,
		functioneventinvokeconfig.Setup,
		permission.Setup,
		providerconfig.Setup,
		instancerds.Setup,
		optiongroup.Setup,
		parametergroup.Setup,
		subnetgroup.Setup,
		delegationset.Setup,
		healthcheck.Setup,
		record.Setup,
		zone.Setup,
		bucket.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketnotification.Setup,
		topic.Setup,
		queue.Setup,
		server.Setup,
		sshkey.Setup,
		usertransfer.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
