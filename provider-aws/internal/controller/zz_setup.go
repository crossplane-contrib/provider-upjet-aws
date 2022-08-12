/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	certificate "github.com/upbound/official-providers/provider-aws/internal/controller/acm/certificate"
	certificatevalidation "github.com/upbound/official-providers/provider-aws/internal/controller/acm/certificatevalidation"
	certificateacmpca "github.com/upbound/official-providers/provider-aws/internal/controller/acmpca/certificate"
	certificateauthority "github.com/upbound/official-providers/provider-aws/internal/controller/acmpca/certificateauthority"
	certificateauthoritycertificate "github.com/upbound/official-providers/provider-aws/internal/controller/acmpca/certificateauthoritycertificate"
	api "github.com/upbound/official-providers/provider-aws/internal/controller/apigatewayv2/api"
	apimapping "github.com/upbound/official-providers/provider-aws/internal/controller/apigatewayv2/apimapping"
	authorizer "github.com/upbound/official-providers/provider-aws/internal/controller/apigatewayv2/authorizer"
	deployment "github.com/upbound/official-providers/provider-aws/internal/controller/apigatewayv2/deployment"
	domainname "github.com/upbound/official-providers/provider-aws/internal/controller/apigatewayv2/domainname"
	integration "github.com/upbound/official-providers/provider-aws/internal/controller/apigatewayv2/integration"
	integrationresponse "github.com/upbound/official-providers/provider-aws/internal/controller/apigatewayv2/integrationresponse"
	model "github.com/upbound/official-providers/provider-aws/internal/controller/apigatewayv2/model"
	route "github.com/upbound/official-providers/provider-aws/internal/controller/apigatewayv2/route"
	routeresponse "github.com/upbound/official-providers/provider-aws/internal/controller/apigatewayv2/routeresponse"
	stage "github.com/upbound/official-providers/provider-aws/internal/controller/apigatewayv2/stage"
	vpclink "github.com/upbound/official-providers/provider-aws/internal/controller/apigatewayv2/vpclink"
	workgroup "github.com/upbound/official-providers/provider-aws/internal/controller/athena/workgroup"
	attachment "github.com/upbound/official-providers/provider-aws/internal/controller/autoscaling/attachment"
	autoscalinggroup "github.com/upbound/official-providers/provider-aws/internal/controller/autoscaling/autoscalinggroup"
	framework "github.com/upbound/official-providers/provider-aws/internal/controller/backup/framework"
	globalsettings "github.com/upbound/official-providers/provider-aws/internal/controller/backup/globalsettings"
	plan "github.com/upbound/official-providers/provider-aws/internal/controller/backup/plan"
	regionsettings "github.com/upbound/official-providers/provider-aws/internal/controller/backup/regionsettings"
	reportplan "github.com/upbound/official-providers/provider-aws/internal/controller/backup/reportplan"
	selection "github.com/upbound/official-providers/provider-aws/internal/controller/backup/selection"
	vault "github.com/upbound/official-providers/provider-aws/internal/controller/backup/vault"
	vaultlockconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/backup/vaultlockconfiguration"
	vaultnotifications "github.com/upbound/official-providers/provider-aws/internal/controller/backup/vaultnotifications"
	vaultpolicy "github.com/upbound/official-providers/provider-aws/internal/controller/backup/vaultpolicy"
	cachepolicy "github.com/upbound/official-providers/provider-aws/internal/controller/cloudfront/cachepolicy"
	distribution "github.com/upbound/official-providers/provider-aws/internal/controller/cloudfront/distribution"
	fieldlevelencryptionconfig "github.com/upbound/official-providers/provider-aws/internal/controller/cloudfront/fieldlevelencryptionconfig"
	fieldlevelencryptionprofile "github.com/upbound/official-providers/provider-aws/internal/controller/cloudfront/fieldlevelencryptionprofile"
	function "github.com/upbound/official-providers/provider-aws/internal/controller/cloudfront/function"
	keygroup "github.com/upbound/official-providers/provider-aws/internal/controller/cloudfront/keygroup"
	monitoringsubscription "github.com/upbound/official-providers/provider-aws/internal/controller/cloudfront/monitoringsubscription"
	originaccessidentity "github.com/upbound/official-providers/provider-aws/internal/controller/cloudfront/originaccessidentity"
	originrequestpolicy "github.com/upbound/official-providers/provider-aws/internal/controller/cloudfront/originrequestpolicy"
	publickey "github.com/upbound/official-providers/provider-aws/internal/controller/cloudfront/publickey"
	realtimelogconfig "github.com/upbound/official-providers/provider-aws/internal/controller/cloudfront/realtimelogconfig"
	responseheaderspolicy "github.com/upbound/official-providers/provider-aws/internal/controller/cloudfront/responseheaderspolicy"
	group "github.com/upbound/official-providers/provider-aws/internal/controller/cloudwatchlogs/group"
	pool "github.com/upbound/official-providers/provider-aws/internal/controller/cognitoidentity/pool"
	identityprovider "github.com/upbound/official-providers/provider-aws/internal/controller/cognitoidp/identityprovider"
	resourceserver "github.com/upbound/official-providers/provider-aws/internal/controller/cognitoidp/resourceserver"
	usergroup "github.com/upbound/official-providers/provider-aws/internal/controller/cognitoidp/usergroup"
	userpool "github.com/upbound/official-providers/provider-aws/internal/controller/cognitoidp/userpool"
	userpoolclient "github.com/upbound/official-providers/provider-aws/internal/controller/cognitoidp/userpoolclient"
	userpooldomain "github.com/upbound/official-providers/provider-aws/internal/controller/cognitoidp/userpooldomain"
	userpooluicustomization "github.com/upbound/official-providers/provider-aws/internal/controller/cognitoidp/userpooluicustomization"
	cluster "github.com/upbound/official-providers/provider-aws/internal/controller/dax/cluster"
	parametergroup "github.com/upbound/official-providers/provider-aws/internal/controller/dax/parametergroup"
	subnetgroup "github.com/upbound/official-providers/provider-aws/internal/controller/dax/subnetgroup"
	clusterdocdb "github.com/upbound/official-providers/provider-aws/internal/controller/docdb/cluster"
	clusterinstance "github.com/upbound/official-providers/provider-aws/internal/controller/docdb/clusterinstance"
	globalcluster "github.com/upbound/official-providers/provider-aws/internal/controller/docdb/globalcluster"
	subnetgroupdocdb "github.com/upbound/official-providers/provider-aws/internal/controller/docdb/subnetgroup"
	contributorinsights "github.com/upbound/official-providers/provider-aws/internal/controller/dynamodb/contributorinsights"
	globaltable "github.com/upbound/official-providers/provider-aws/internal/controller/dynamodb/globaltable"
	kinesisstreamingdestination "github.com/upbound/official-providers/provider-aws/internal/controller/dynamodb/kinesisstreamingdestination"
	table "github.com/upbound/official-providers/provider-aws/internal/controller/dynamodb/table"
	tableitem "github.com/upbound/official-providers/provider-aws/internal/controller/dynamodb/tableitem"
	ebssnapshot "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/ebssnapshot"
	ebsvolume "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/ebsvolume"
	egressonlyinternetgateway "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/egressonlyinternetgateway"
	eip "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/eip"
	eipassociation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/eipassociation"
	flowlog "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/flowlog"
	instance "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/instance"
	internetgateway "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/internetgateway"
	keypair "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/keypair"
	launchtemplate "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/launchtemplate"
	mainroutetableassociation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/mainroutetableassociation"
	managedprefixlist "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/managedprefixlist"
	managedprefixlistentry "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/managedprefixlistentry"
	natgateway "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/natgateway"
	networkacl "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/networkacl"
	networkaclrule "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/networkaclrule"
	networkinterface "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/networkinterface"
	networkinterfaceattachment "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/networkinterfaceattachment"
	networkinterfacesgattachment "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/networkinterfacesgattachment"
	placementgroup "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/placementgroup"
	routeec2 "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/route"
	routetable "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/routetableassociation"
	securitygroup "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/securitygrouprule"
	spotdatafeedsubscription "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/spotdatafeedsubscription"
	spotinstancerequest "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/spotinstancerequest"
	subnet "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/subnet"
	transitgateway "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgateway"
	transitgatewaymulticastdomain "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewaymulticastdomain"
	transitgatewaymulticastdomainassociation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewaymulticastdomainassociation"
	transitgatewaymulticastgroupmember "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewaymulticastgroupmember"
	transitgatewaymulticastgroupsource "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewaymulticastgroupsource"
	transitgatewaypeeringattachment "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewaypeeringattachment"
	transitgatewayprefixlistreference "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayprefixlistreference"
	transitgatewayroute "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/transitgatewayvpcattachmentaccepter"
	volumeattachment "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/volumeattachment"
	vpc "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpc"
	vpcdhcpoptions "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcdhcpoptions"
	vpcdhcpoptionsassociation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcdhcpoptionsassociation"
	vpcendpoint "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcendpoint"
	vpcendpointconnectionnotification "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcendpointconnectionnotification"
	vpcendpointroutetableassociation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcendpointroutetableassociation"
	vpcendpointservice "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcendpointservice"
	vpcendpointserviceallowedprincipal "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcendpointserviceallowedprincipal"
	vpcendpointsubnetassociation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcendpointsubnetassociation"
	vpcipv4cidrblockassociation "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/upbound/official-providers/provider-aws/internal/controller/ec2/vpcpeeringconnection"
	lifecyclepolicy "github.com/upbound/official-providers/provider-aws/internal/controller/ecr/lifecyclepolicy"
	pullthroughcacherule "github.com/upbound/official-providers/provider-aws/internal/controller/ecr/pullthroughcacherule"
	registrypolicy "github.com/upbound/official-providers/provider-aws/internal/controller/ecr/registrypolicy"
	registryscanningconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/ecr/registryscanningconfiguration"
	replicationconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/ecr/replicationconfiguration"
	repository "github.com/upbound/official-providers/provider-aws/internal/controller/ecr/repository"
	repositorypolicy "github.com/upbound/official-providers/provider-aws/internal/controller/ecr/repositorypolicy"
	repositoryecrpublic "github.com/upbound/official-providers/provider-aws/internal/controller/ecrpublic/repository"
	repositorypolicyecrpublic "github.com/upbound/official-providers/provider-aws/internal/controller/ecrpublic/repositorypolicy"
	accountsettingdefault "github.com/upbound/official-providers/provider-aws/internal/controller/ecs/accountsettingdefault"
	capacityprovider "github.com/upbound/official-providers/provider-aws/internal/controller/ecs/capacityprovider"
	clusterecs "github.com/upbound/official-providers/provider-aws/internal/controller/ecs/cluster"
	clustercapacityproviders "github.com/upbound/official-providers/provider-aws/internal/controller/ecs/clustercapacityproviders"
	service "github.com/upbound/official-providers/provider-aws/internal/controller/ecs/service"
	taskdefinition "github.com/upbound/official-providers/provider-aws/internal/controller/ecs/taskdefinition"
	accesspoint "github.com/upbound/official-providers/provider-aws/internal/controller/efs/accesspoint"
	backuppolicy "github.com/upbound/official-providers/provider-aws/internal/controller/efs/backuppolicy"
	filesystem "github.com/upbound/official-providers/provider-aws/internal/controller/efs/filesystem"
	filesystempolicy "github.com/upbound/official-providers/provider-aws/internal/controller/efs/filesystempolicy"
	mounttarget "github.com/upbound/official-providers/provider-aws/internal/controller/efs/mounttarget"
	addon "github.com/upbound/official-providers/provider-aws/internal/controller/eks/addon"
	clustereks "github.com/upbound/official-providers/provider-aws/internal/controller/eks/cluster"
	fargateprofile "github.com/upbound/official-providers/provider-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/upbound/official-providers/provider-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/upbound/official-providers/provider-aws/internal/controller/eks/nodegroup"
	clusterelasticache "github.com/upbound/official-providers/provider-aws/internal/controller/elasticache/cluster"
	parametergroupelasticache "github.com/upbound/official-providers/provider-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/upbound/official-providers/provider-aws/internal/controller/elasticache/replicationgroup"
	subnetgroupelasticache "github.com/upbound/official-providers/provider-aws/internal/controller/elasticache/subnetgroup"
	user "github.com/upbound/official-providers/provider-aws/internal/controller/elasticache/user"
	usergroupelasticache "github.com/upbound/official-providers/provider-aws/internal/controller/elasticache/usergroup"
	attachmentelb "github.com/upbound/official-providers/provider-aws/internal/controller/elb/attachment"
	elb "github.com/upbound/official-providers/provider-aws/internal/controller/elb/elb"
	lb "github.com/upbound/official-providers/provider-aws/internal/controller/elbv2/lb"
	lblistener "github.com/upbound/official-providers/provider-aws/internal/controller/elbv2/lblistener"
	lbtargetgroup "github.com/upbound/official-providers/provider-aws/internal/controller/elbv2/lbtargetgroup"
	lbtargetgroupattachment "github.com/upbound/official-providers/provider-aws/internal/controller/elbv2/lbtargetgroupattachment"
	deliverystream "github.com/upbound/official-providers/provider-aws/internal/controller/firehose/deliverystream"
	alias "github.com/upbound/official-providers/provider-aws/internal/controller/gamelift/alias"
	build "github.com/upbound/official-providers/provider-aws/internal/controller/gamelift/build"
	fleet "github.com/upbound/official-providers/provider-aws/internal/controller/gamelift/fleet"
	gamesessionqueue "github.com/upbound/official-providers/provider-aws/internal/controller/gamelift/gamesessionqueue"
	script "github.com/upbound/official-providers/provider-aws/internal/controller/gamelift/script"
	accelerator "github.com/upbound/official-providers/provider-aws/internal/controller/globalaccelerator/accelerator"
	endpointgroup "github.com/upbound/official-providers/provider-aws/internal/controller/globalaccelerator/endpointgroup"
	listener "github.com/upbound/official-providers/provider-aws/internal/controller/globalaccelerator/listener"
	catalogdatabase "github.com/upbound/official-providers/provider-aws/internal/controller/glue/catalogdatabase"
	catalogtable "github.com/upbound/official-providers/provider-aws/internal/controller/glue/catalogtable"
	classifier "github.com/upbound/official-providers/provider-aws/internal/controller/glue/classifier"
	datacatalogencryptionsettings "github.com/upbound/official-providers/provider-aws/internal/controller/glue/datacatalogencryptionsettings"
	job "github.com/upbound/official-providers/provider-aws/internal/controller/glue/job"
	registry "github.com/upbound/official-providers/provider-aws/internal/controller/glue/registry"
	resourcepolicy "github.com/upbound/official-providers/provider-aws/internal/controller/glue/resourcepolicy"
	trigger "github.com/upbound/official-providers/provider-aws/internal/controller/glue/trigger"
	userdefinedfunction "github.com/upbound/official-providers/provider-aws/internal/controller/glue/userdefinedfunction"
	workflow "github.com/upbound/official-providers/provider-aws/internal/controller/glue/workflow"
	roleassociation "github.com/upbound/official-providers/provider-aws/internal/controller/grafana/roleassociation"
	workspace "github.com/upbound/official-providers/provider-aws/internal/controller/grafana/workspace"
	workspacesamlconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/grafana/workspacesamlconfiguration"
	accesskey "github.com/upbound/official-providers/provider-aws/internal/controller/iam/accesskey"
	accountalias "github.com/upbound/official-providers/provider-aws/internal/controller/iam/accountalias"
	accountpasswordpolicy "github.com/upbound/official-providers/provider-aws/internal/controller/iam/accountpasswordpolicy"
	groupiam "github.com/upbound/official-providers/provider-aws/internal/controller/iam/group"
	groupmembership "github.com/upbound/official-providers/provider-aws/internal/controller/iam/groupmembership"
	grouppolicyattachment "github.com/upbound/official-providers/provider-aws/internal/controller/iam/grouppolicyattachment"
	instanceprofile "github.com/upbound/official-providers/provider-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/upbound/official-providers/provider-aws/internal/controller/iam/openidconnectprovider"
	policy "github.com/upbound/official-providers/provider-aws/internal/controller/iam/policy"
	role "github.com/upbound/official-providers/provider-aws/internal/controller/iam/role"
	rolepolicyattachment "github.com/upbound/official-providers/provider-aws/internal/controller/iam/rolepolicyattachment"
	samlprovider "github.com/upbound/official-providers/provider-aws/internal/controller/iam/samlprovider"
	servercertificate "github.com/upbound/official-providers/provider-aws/internal/controller/iam/servercertificate"
	servicelinkedrole "github.com/upbound/official-providers/provider-aws/internal/controller/iam/servicelinkedrole"
	servicespecificcredential "github.com/upbound/official-providers/provider-aws/internal/controller/iam/servicespecificcredential"
	signingcertificate "github.com/upbound/official-providers/provider-aws/internal/controller/iam/signingcertificate"
	useriam "github.com/upbound/official-providers/provider-aws/internal/controller/iam/user"
	usergroupmembership "github.com/upbound/official-providers/provider-aws/internal/controller/iam/usergroupmembership"
	userloginprofile "github.com/upbound/official-providers/provider-aws/internal/controller/iam/userloginprofile"
	userpolicyattachment "github.com/upbound/official-providers/provider-aws/internal/controller/iam/userpolicyattachment"
	usersshkey "github.com/upbound/official-providers/provider-aws/internal/controller/iam/usersshkey"
	virtualmfadevice "github.com/upbound/official-providers/provider-aws/internal/controller/iam/virtualmfadevice"
	policyiot "github.com/upbound/official-providers/provider-aws/internal/controller/iot/policy"
	thing "github.com/upbound/official-providers/provider-aws/internal/controller/iot/thing"
	clusterkafka "github.com/upbound/official-providers/provider-aws/internal/controller/kafka/cluster"
	configuration "github.com/upbound/official-providers/provider-aws/internal/controller/kafka/configuration"
	stream "github.com/upbound/official-providers/provider-aws/internal/controller/kinesis/stream"
	streamconsumer "github.com/upbound/official-providers/provider-aws/internal/controller/kinesis/streamconsumer"
	application "github.com/upbound/official-providers/provider-aws/internal/controller/kinesisanalytics/application"
	applicationkinesisanalyticsv2 "github.com/upbound/official-providers/provider-aws/internal/controller/kinesisanalyticsv2/application"
	applicationsnapshot "github.com/upbound/official-providers/provider-aws/internal/controller/kinesisanalyticsv2/applicationsnapshot"
	streamkinesisvideo "github.com/upbound/official-providers/provider-aws/internal/controller/kinesisvideo/stream"
	aliaskms "github.com/upbound/official-providers/provider-aws/internal/controller/kms/alias"
	ciphertext "github.com/upbound/official-providers/provider-aws/internal/controller/kms/ciphertext"
	externalkey "github.com/upbound/official-providers/provider-aws/internal/controller/kms/externalkey"
	grant "github.com/upbound/official-providers/provider-aws/internal/controller/kms/grant"
	key "github.com/upbound/official-providers/provider-aws/internal/controller/kms/key"
	replicaexternalkey "github.com/upbound/official-providers/provider-aws/internal/controller/kms/replicaexternalkey"
	replicakey "github.com/upbound/official-providers/provider-aws/internal/controller/kms/replicakey"
	datalakesettings "github.com/upbound/official-providers/provider-aws/internal/controller/lakeformation/datalakesettings"
	permissions "github.com/upbound/official-providers/provider-aws/internal/controller/lakeformation/permissions"
	resource "github.com/upbound/official-providers/provider-aws/internal/controller/lakeformation/resource"
	aliaslambda "github.com/upbound/official-providers/provider-aws/internal/controller/lambda/alias"
	codesigningconfig "github.com/upbound/official-providers/provider-aws/internal/controller/lambda/codesigningconfig"
	eventsourcemapping "github.com/upbound/official-providers/provider-aws/internal/controller/lambda/eventsourcemapping"
	functionlambda "github.com/upbound/official-providers/provider-aws/internal/controller/lambda/function"
	functioneventinvokeconfig "github.com/upbound/official-providers/provider-aws/internal/controller/lambda/functioneventinvokeconfig"
	functionurl "github.com/upbound/official-providers/provider-aws/internal/controller/lambda/functionurl"
	invocation "github.com/upbound/official-providers/provider-aws/internal/controller/lambda/invocation"
	layerversion "github.com/upbound/official-providers/provider-aws/internal/controller/lambda/layerversion"
	layerversionpermission "github.com/upbound/official-providers/provider-aws/internal/controller/lambda/layerversionpermission"
	permission "github.com/upbound/official-providers/provider-aws/internal/controller/lambda/permission"
	provisionedconcurrencyconfig "github.com/upbound/official-providers/provider-aws/internal/controller/lambda/provisionedconcurrencyconfig"
	bot "github.com/upbound/official-providers/provider-aws/internal/controller/lexmodels/bot"
	botalias "github.com/upbound/official-providers/provider-aws/internal/controller/lexmodels/botalias"
	intent "github.com/upbound/official-providers/provider-aws/internal/controller/lexmodels/intent"
	slottype "github.com/upbound/official-providers/provider-aws/internal/controller/lexmodels/slottype"
	association "github.com/upbound/official-providers/provider-aws/internal/controller/licensemanager/association"
	licenseconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/licensemanager/licenseconfiguration"
	broker "github.com/upbound/official-providers/provider-aws/internal/controller/mq/broker"
	configurationmq "github.com/upbound/official-providers/provider-aws/internal/controller/mq/configuration"
	clusterneptune "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/cluster"
	clusterendpoint "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/clusterendpoint"
	clusterinstanceneptune "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/clusterinstance"
	clusterparametergroup "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/clusterparametergroup"
	clustersnapshot "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/clustersnapshot"
	eventsubscription "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/eventsubscription"
	parametergroupneptune "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/parametergroup"
	subnetgroupneptune "github.com/upbound/official-providers/provider-aws/internal/controller/neptune/subnetgroup"
	providerconfig "github.com/upbound/official-providers/provider-aws/internal/controller/providerconfig"
	resourceshare "github.com/upbound/official-providers/provider-aws/internal/controller/ram/resourceshare"
	clusterrds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/cluster"
	clusteractivitystream "github.com/upbound/official-providers/provider-aws/internal/controller/rds/clusteractivitystream"
	clusterendpointrds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/clusterendpoint"
	clusterinstancerds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/clusterinstance"
	clusterparametergrouprds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/clusterparametergroup"
	clusterroleassociation "github.com/upbound/official-providers/provider-aws/internal/controller/rds/clusterroleassociation"
	globalclusterrds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/globalcluster"
	instancerds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/instance"
	instanceroleassociation "github.com/upbound/official-providers/provider-aws/internal/controller/rds/instanceroleassociation"
	optiongroup "github.com/upbound/official-providers/provider-aws/internal/controller/rds/optiongroup"
	parametergrouprds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/parametergroup"
	proxy "github.com/upbound/official-providers/provider-aws/internal/controller/rds/proxy"
	proxydefaulttargetgroup "github.com/upbound/official-providers/provider-aws/internal/controller/rds/proxydefaulttargetgroup"
	proxyendpoint "github.com/upbound/official-providers/provider-aws/internal/controller/rds/proxyendpoint"
	proxytarget "github.com/upbound/official-providers/provider-aws/internal/controller/rds/proxytarget"
	securitygrouprds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/securitygroup"
	snapshot "github.com/upbound/official-providers/provider-aws/internal/controller/rds/snapshot"
	subnetgrouprds "github.com/upbound/official-providers/provider-aws/internal/controller/rds/subnetgroup"
	clusterredshift "github.com/upbound/official-providers/provider-aws/internal/controller/redshift/cluster"
	groupresourcegroups "github.com/upbound/official-providers/provider-aws/internal/controller/resourcegroups/group"
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
	bucketaccelerateconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketaccelerateconfiguration"
	bucketacl "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketacl"
	bucketanalyticsconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketanalyticsconfiguration"
	bucketcorsconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketcorsconfiguration"
	bucketintelligenttieringconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketintelligenttieringconfiguration"
	bucketinventory "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketinventory"
	bucketlifecycleconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketlifecycleconfiguration"
	bucketlogging "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketlogging"
	bucketmetric "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketmetric"
	bucketnotification "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketnotification"
	bucketobject "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketobject"
	bucketobjectlockconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketobjectlockconfiguration"
	bucketownershipcontrols "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketownershipcontrols"
	bucketpolicy "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketpolicy"
	bucketpublicaccessblock "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketpublicaccessblock"
	bucketreplicationconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketreplicationconfiguration"
	bucketrequestpaymentconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketrequestpaymentconfiguration"
	bucketserversideencryptionconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketserversideencryptionconfiguration"
	bucketversioning "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketversioning"
	bucketwebsiteconfiguration "github.com/upbound/official-providers/provider-aws/internal/controller/s3/bucketwebsiteconfiguration"
	object "github.com/upbound/official-providers/provider-aws/internal/controller/s3/object"
	secret "github.com/upbound/official-providers/provider-aws/internal/controller/secretsmanager/secret"
	httpnamespace "github.com/upbound/official-providers/provider-aws/internal/controller/servicediscovery/httpnamespace"
	privatednsnamespace "github.com/upbound/official-providers/provider-aws/internal/controller/servicediscovery/privatednsnamespace"
	publicdnsnamespace "github.com/upbound/official-providers/provider-aws/internal/controller/servicediscovery/publicdnsnamespace"
	activity "github.com/upbound/official-providers/provider-aws/internal/controller/sfn/activity"
	statemachine "github.com/upbound/official-providers/provider-aws/internal/controller/sfn/statemachine"
	signingprofile "github.com/upbound/official-providers/provider-aws/internal/controller/signer/signingprofile"
	topic "github.com/upbound/official-providers/provider-aws/internal/controller/sns/topic"
	topicsubscription "github.com/upbound/official-providers/provider-aws/internal/controller/sns/topicsubscription"
	queue "github.com/upbound/official-providers/provider-aws/internal/controller/sqs/queue"
	server "github.com/upbound/official-providers/provider-aws/internal/controller/transfer/server"
	usertransfer "github.com/upbound/official-providers/provider-aws/internal/controller/transfer/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		certificate.Setup,
		certificatevalidation.Setup,
		certificateacmpca.Setup,
		certificateauthority.Setup,
		certificateauthoritycertificate.Setup,
		api.Setup,
		apimapping.Setup,
		authorizer.Setup,
		deployment.Setup,
		domainname.Setup,
		integration.Setup,
		integrationresponse.Setup,
		model.Setup,
		route.Setup,
		routeresponse.Setup,
		stage.Setup,
		vpclink.Setup,
		workgroup.Setup,
		attachment.Setup,
		autoscalinggroup.Setup,
		framework.Setup,
		globalsettings.Setup,
		plan.Setup,
		regionsettings.Setup,
		reportplan.Setup,
		selection.Setup,
		vault.Setup,
		vaultlockconfiguration.Setup,
		vaultnotifications.Setup,
		vaultpolicy.Setup,
		cachepolicy.Setup,
		distribution.Setup,
		fieldlevelencryptionconfig.Setup,
		fieldlevelencryptionprofile.Setup,
		function.Setup,
		keygroup.Setup,
		monitoringsubscription.Setup,
		originaccessidentity.Setup,
		originrequestpolicy.Setup,
		publickey.Setup,
		realtimelogconfig.Setup,
		responseheaderspolicy.Setup,
		group.Setup,
		pool.Setup,
		identityprovider.Setup,
		resourceserver.Setup,
		usergroup.Setup,
		userpool.Setup,
		userpoolclient.Setup,
		userpooldomain.Setup,
		userpooluicustomization.Setup,
		cluster.Setup,
		parametergroup.Setup,
		subnetgroup.Setup,
		clusterdocdb.Setup,
		clusterinstance.Setup,
		globalcluster.Setup,
		subnetgroupdocdb.Setup,
		contributorinsights.Setup,
		globaltable.Setup,
		kinesisstreamingdestination.Setup,
		table.Setup,
		tableitem.Setup,
		ebssnapshot.Setup,
		ebsvolume.Setup,
		egressonlyinternetgateway.Setup,
		eip.Setup,
		eipassociation.Setup,
		flowlog.Setup,
		instance.Setup,
		internetgateway.Setup,
		keypair.Setup,
		launchtemplate.Setup,
		mainroutetableassociation.Setup,
		managedprefixlist.Setup,
		managedprefixlistentry.Setup,
		natgateway.Setup,
		networkacl.Setup,
		networkaclrule.Setup,
		networkinterface.Setup,
		networkinterfaceattachment.Setup,
		networkinterfacesgattachment.Setup,
		placementgroup.Setup,
		routeec2.Setup,
		routetable.Setup,
		routetableassociation.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		spotdatafeedsubscription.Setup,
		spotinstancerequest.Setup,
		subnet.Setup,
		transitgateway.Setup,
		transitgatewaymulticastdomain.Setup,
		transitgatewaymulticastdomainassociation.Setup,
		transitgatewaymulticastgroupmember.Setup,
		transitgatewaymulticastgroupsource.Setup,
		transitgatewaypeeringattachment.Setup,
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
		vpcendpointservice.Setup,
		vpcendpointserviceallowedprincipal.Setup,
		vpcendpointsubnetassociation.Setup,
		vpcipv4cidrblockassociation.Setup,
		vpcpeeringconnection.Setup,
		lifecyclepolicy.Setup,
		pullthroughcacherule.Setup,
		registrypolicy.Setup,
		registryscanningconfiguration.Setup,
		replicationconfiguration.Setup,
		repository.Setup,
		repositorypolicy.Setup,
		repositoryecrpublic.Setup,
		repositorypolicyecrpublic.Setup,
		accountsettingdefault.Setup,
		capacityprovider.Setup,
		clusterecs.Setup,
		clustercapacityproviders.Setup,
		service.Setup,
		taskdefinition.Setup,
		accesspoint.Setup,
		backuppolicy.Setup,
		filesystem.Setup,
		filesystempolicy.Setup,
		mounttarget.Setup,
		addon.Setup,
		clustereks.Setup,
		fargateprofile.Setup,
		identityproviderconfig.Setup,
		nodegroup.Setup,
		clusterelasticache.Setup,
		parametergroupelasticache.Setup,
		replicationgroup.Setup,
		subnetgroupelasticache.Setup,
		user.Setup,
		usergroupelasticache.Setup,
		attachmentelb.Setup,
		elb.Setup,
		lb.Setup,
		lblistener.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		deliverystream.Setup,
		alias.Setup,
		build.Setup,
		fleet.Setup,
		gamesessionqueue.Setup,
		script.Setup,
		accelerator.Setup,
		endpointgroup.Setup,
		listener.Setup,
		catalogdatabase.Setup,
		catalogtable.Setup,
		classifier.Setup,
		datacatalogencryptionsettings.Setup,
		job.Setup,
		registry.Setup,
		resourcepolicy.Setup,
		trigger.Setup,
		userdefinedfunction.Setup,
		workflow.Setup,
		roleassociation.Setup,
		workspace.Setup,
		workspacesamlconfiguration.Setup,
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		groupiam.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		instanceprofile.Setup,
		openidconnectprovider.Setup,
		policy.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		servercertificate.Setup,
		servicelinkedrole.Setup,
		servicespecificcredential.Setup,
		signingcertificate.Setup,
		useriam.Setup,
		usergroupmembership.Setup,
		userloginprofile.Setup,
		userpolicyattachment.Setup,
		usersshkey.Setup,
		virtualmfadevice.Setup,
		policyiot.Setup,
		thing.Setup,
		clusterkafka.Setup,
		configuration.Setup,
		stream.Setup,
		streamconsumer.Setup,
		application.Setup,
		applicationkinesisanalyticsv2.Setup,
		applicationsnapshot.Setup,
		streamkinesisvideo.Setup,
		aliaskms.Setup,
		ciphertext.Setup,
		externalkey.Setup,
		grant.Setup,
		key.Setup,
		replicaexternalkey.Setup,
		replicakey.Setup,
		datalakesettings.Setup,
		permissions.Setup,
		resource.Setup,
		aliaslambda.Setup,
		codesigningconfig.Setup,
		eventsourcemapping.Setup,
		functionlambda.Setup,
		functioneventinvokeconfig.Setup,
		functionurl.Setup,
		invocation.Setup,
		layerversion.Setup,
		layerversionpermission.Setup,
		permission.Setup,
		provisionedconcurrencyconfig.Setup,
		bot.Setup,
		botalias.Setup,
		intent.Setup,
		slottype.Setup,
		association.Setup,
		licenseconfiguration.Setup,
		broker.Setup,
		configurationmq.Setup,
		clusterneptune.Setup,
		clusterendpoint.Setup,
		clusterinstanceneptune.Setup,
		clusterparametergroup.Setup,
		clustersnapshot.Setup,
		eventsubscription.Setup,
		parametergroupneptune.Setup,
		subnetgroupneptune.Setup,
		providerconfig.Setup,
		resourceshare.Setup,
		clusterrds.Setup,
		clusteractivitystream.Setup,
		clusterendpointrds.Setup,
		clusterinstancerds.Setup,
		clusterparametergrouprds.Setup,
		clusterroleassociation.Setup,
		globalclusterrds.Setup,
		instancerds.Setup,
		instanceroleassociation.Setup,
		optiongroup.Setup,
		parametergrouprds.Setup,
		proxy.Setup,
		proxydefaulttargetgroup.Setup,
		proxyendpoint.Setup,
		proxytarget.Setup,
		securitygrouprds.Setup,
		snapshot.Setup,
		subnetgrouprds.Setup,
		clusterredshift.Setup,
		groupresourcegroups.Setup,
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
		bucketaccelerateconfiguration.Setup,
		bucketacl.Setup,
		bucketanalyticsconfiguration.Setup,
		bucketcorsconfiguration.Setup,
		bucketintelligenttieringconfiguration.Setup,
		bucketinventory.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketlogging.Setup,
		bucketmetric.Setup,
		bucketnotification.Setup,
		bucketobject.Setup,
		bucketobjectlockconfiguration.Setup,
		bucketownershipcontrols.Setup,
		bucketpolicy.Setup,
		bucketpublicaccessblock.Setup,
		bucketreplicationconfiguration.Setup,
		bucketrequestpaymentconfiguration.Setup,
		bucketserversideencryptionconfiguration.Setup,
		bucketversioning.Setup,
		bucketwebsiteconfiguration.Setup,
		object.Setup,
		secret.Setup,
		httpnamespace.Setup,
		privatednsnamespace.Setup,
		publicdnsnamespace.Setup,
		activity.Setup,
		statemachine.Setup,
		signingprofile.Setup,
		topic.Setup,
		topicsubscription.Setup,
		queue.Setup,
		server.Setup,
		usertransfer.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
