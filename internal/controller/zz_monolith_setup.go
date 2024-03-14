// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	analyzer "github.com/upbound/provider-aws/internal/controller/accessanalyzer/analyzer"
	archiverule "github.com/upbound/provider-aws/internal/controller/accessanalyzer/archiverule"
	alternatecontact "github.com/upbound/provider-aws/internal/controller/account/alternatecontact"
	certificate "github.com/upbound/provider-aws/internal/controller/acm/certificate"
	certificatevalidation "github.com/upbound/provider-aws/internal/controller/acm/certificatevalidation"
	certificateacmpca "github.com/upbound/provider-aws/internal/controller/acmpca/certificate"
	certificateauthority "github.com/upbound/provider-aws/internal/controller/acmpca/certificateauthority"
	certificateauthoritycertificate "github.com/upbound/provider-aws/internal/controller/acmpca/certificateauthoritycertificate"
	permission "github.com/upbound/provider-aws/internal/controller/acmpca/permission"
	policy "github.com/upbound/provider-aws/internal/controller/acmpca/policy"
	alertmanagerdefinition "github.com/upbound/provider-aws/internal/controller/amp/alertmanagerdefinition"
	rulegroupnamespace "github.com/upbound/provider-aws/internal/controller/amp/rulegroupnamespace"
	workspace "github.com/upbound/provider-aws/internal/controller/amp/workspace"
	app "github.com/upbound/provider-aws/internal/controller/amplify/app"
	backendenvironment "github.com/upbound/provider-aws/internal/controller/amplify/backendenvironment"
	branch "github.com/upbound/provider-aws/internal/controller/amplify/branch"
	webhook "github.com/upbound/provider-aws/internal/controller/amplify/webhook"
	account "github.com/upbound/provider-aws/internal/controller/apigateway/account"
	apikey "github.com/upbound/provider-aws/internal/controller/apigateway/apikey"
	authorizer "github.com/upbound/provider-aws/internal/controller/apigateway/authorizer"
	basepathmapping "github.com/upbound/provider-aws/internal/controller/apigateway/basepathmapping"
	clientcertificate "github.com/upbound/provider-aws/internal/controller/apigateway/clientcertificate"
	deployment "github.com/upbound/provider-aws/internal/controller/apigateway/deployment"
	documentationpart "github.com/upbound/provider-aws/internal/controller/apigateway/documentationpart"
	documentationversion "github.com/upbound/provider-aws/internal/controller/apigateway/documentationversion"
	domainname "github.com/upbound/provider-aws/internal/controller/apigateway/domainname"
	gatewayresponse "github.com/upbound/provider-aws/internal/controller/apigateway/gatewayresponse"
	integration "github.com/upbound/provider-aws/internal/controller/apigateway/integration"
	integrationresponse "github.com/upbound/provider-aws/internal/controller/apigateway/integrationresponse"
	method "github.com/upbound/provider-aws/internal/controller/apigateway/method"
	methodresponse "github.com/upbound/provider-aws/internal/controller/apigateway/methodresponse"
	methodsettings "github.com/upbound/provider-aws/internal/controller/apigateway/methodsettings"
	model "github.com/upbound/provider-aws/internal/controller/apigateway/model"
	requestvalidator "github.com/upbound/provider-aws/internal/controller/apigateway/requestvalidator"
	resource "github.com/upbound/provider-aws/internal/controller/apigateway/resource"
	restapi "github.com/upbound/provider-aws/internal/controller/apigateway/restapi"
	restapipolicy "github.com/upbound/provider-aws/internal/controller/apigateway/restapipolicy"
	stage "github.com/upbound/provider-aws/internal/controller/apigateway/stage"
	usageplan "github.com/upbound/provider-aws/internal/controller/apigateway/usageplan"
	usageplankey "github.com/upbound/provider-aws/internal/controller/apigateway/usageplankey"
	vpclink "github.com/upbound/provider-aws/internal/controller/apigateway/vpclink"
	api "github.com/upbound/provider-aws/internal/controller/apigatewayv2/api"
	apimapping "github.com/upbound/provider-aws/internal/controller/apigatewayv2/apimapping"
	authorizerapigatewayv2 "github.com/upbound/provider-aws/internal/controller/apigatewayv2/authorizer"
	deploymentapigatewayv2 "github.com/upbound/provider-aws/internal/controller/apigatewayv2/deployment"
	domainnameapigatewayv2 "github.com/upbound/provider-aws/internal/controller/apigatewayv2/domainname"
	integrationapigatewayv2 "github.com/upbound/provider-aws/internal/controller/apigatewayv2/integration"
	integrationresponseapigatewayv2 "github.com/upbound/provider-aws/internal/controller/apigatewayv2/integrationresponse"
	modelapigatewayv2 "github.com/upbound/provider-aws/internal/controller/apigatewayv2/model"
	route "github.com/upbound/provider-aws/internal/controller/apigatewayv2/route"
	routeresponse "github.com/upbound/provider-aws/internal/controller/apigatewayv2/routeresponse"
	stageapigatewayv2 "github.com/upbound/provider-aws/internal/controller/apigatewayv2/stage"
	vpclinkapigatewayv2 "github.com/upbound/provider-aws/internal/controller/apigatewayv2/vpclink"
	policyappautoscaling "github.com/upbound/provider-aws/internal/controller/appautoscaling/policy"
	scheduledaction "github.com/upbound/provider-aws/internal/controller/appautoscaling/scheduledaction"
	target "github.com/upbound/provider-aws/internal/controller/appautoscaling/target"
	application "github.com/upbound/provider-aws/internal/controller/appconfig/application"
	configurationprofile "github.com/upbound/provider-aws/internal/controller/appconfig/configurationprofile"
	deploymentappconfig "github.com/upbound/provider-aws/internal/controller/appconfig/deployment"
	deploymentstrategy "github.com/upbound/provider-aws/internal/controller/appconfig/deploymentstrategy"
	environment "github.com/upbound/provider-aws/internal/controller/appconfig/environment"
	extension "github.com/upbound/provider-aws/internal/controller/appconfig/extension"
	extensionassociation "github.com/upbound/provider-aws/internal/controller/appconfig/extensionassociation"
	hostedconfigurationversion "github.com/upbound/provider-aws/internal/controller/appconfig/hostedconfigurationversion"
	flow "github.com/upbound/provider-aws/internal/controller/appflow/flow"
	eventintegration "github.com/upbound/provider-aws/internal/controller/appintegrations/eventintegration"
	applicationapplicationinsights "github.com/upbound/provider-aws/internal/controller/applicationinsights/application"
	gatewayroute "github.com/upbound/provider-aws/internal/controller/appmesh/gatewayroute"
	mesh "github.com/upbound/provider-aws/internal/controller/appmesh/mesh"
	routeappmesh "github.com/upbound/provider-aws/internal/controller/appmesh/route"
	virtualgateway "github.com/upbound/provider-aws/internal/controller/appmesh/virtualgateway"
	virtualnode "github.com/upbound/provider-aws/internal/controller/appmesh/virtualnode"
	virtualrouter "github.com/upbound/provider-aws/internal/controller/appmesh/virtualrouter"
	virtualservice "github.com/upbound/provider-aws/internal/controller/appmesh/virtualservice"
	autoscalingconfigurationversion "github.com/upbound/provider-aws/internal/controller/apprunner/autoscalingconfigurationversion"
	connection "github.com/upbound/provider-aws/internal/controller/apprunner/connection"
	observabilityconfiguration "github.com/upbound/provider-aws/internal/controller/apprunner/observabilityconfiguration"
	service "github.com/upbound/provider-aws/internal/controller/apprunner/service"
	vpcconnector "github.com/upbound/provider-aws/internal/controller/apprunner/vpcconnector"
	directoryconfig "github.com/upbound/provider-aws/internal/controller/appstream/directoryconfig"
	fleet "github.com/upbound/provider-aws/internal/controller/appstream/fleet"
	fleetstackassociation "github.com/upbound/provider-aws/internal/controller/appstream/fleetstackassociation"
	imagebuilder "github.com/upbound/provider-aws/internal/controller/appstream/imagebuilder"
	stack "github.com/upbound/provider-aws/internal/controller/appstream/stack"
	user "github.com/upbound/provider-aws/internal/controller/appstream/user"
	userstackassociation "github.com/upbound/provider-aws/internal/controller/appstream/userstackassociation"
	apicache "github.com/upbound/provider-aws/internal/controller/appsync/apicache"
	apikeyappsync "github.com/upbound/provider-aws/internal/controller/appsync/apikey"
	datasource "github.com/upbound/provider-aws/internal/controller/appsync/datasource"
	function "github.com/upbound/provider-aws/internal/controller/appsync/function"
	graphqlapi "github.com/upbound/provider-aws/internal/controller/appsync/graphqlapi"
	resolver "github.com/upbound/provider-aws/internal/controller/appsync/resolver"
	database "github.com/upbound/provider-aws/internal/controller/athena/database"
	datacatalog "github.com/upbound/provider-aws/internal/controller/athena/datacatalog"
	namedquery "github.com/upbound/provider-aws/internal/controller/athena/namedquery"
	workgroup "github.com/upbound/provider-aws/internal/controller/athena/workgroup"
	attachment "github.com/upbound/provider-aws/internal/controller/autoscaling/attachment"
	autoscalinggroup "github.com/upbound/provider-aws/internal/controller/autoscaling/autoscalinggroup"
	grouptag "github.com/upbound/provider-aws/internal/controller/autoscaling/grouptag"
	launchconfiguration "github.com/upbound/provider-aws/internal/controller/autoscaling/launchconfiguration"
	lifecyclehook "github.com/upbound/provider-aws/internal/controller/autoscaling/lifecyclehook"
	notification "github.com/upbound/provider-aws/internal/controller/autoscaling/notification"
	policyautoscaling "github.com/upbound/provider-aws/internal/controller/autoscaling/policy"
	schedule "github.com/upbound/provider-aws/internal/controller/autoscaling/schedule"
	scalingplan "github.com/upbound/provider-aws/internal/controller/autoscalingplans/scalingplan"
	framework "github.com/upbound/provider-aws/internal/controller/backup/framework"
	globalsettings "github.com/upbound/provider-aws/internal/controller/backup/globalsettings"
	plan "github.com/upbound/provider-aws/internal/controller/backup/plan"
	regionsettings "github.com/upbound/provider-aws/internal/controller/backup/regionsettings"
	reportplan "github.com/upbound/provider-aws/internal/controller/backup/reportplan"
	selection "github.com/upbound/provider-aws/internal/controller/backup/selection"
	vault "github.com/upbound/provider-aws/internal/controller/backup/vault"
	vaultlockconfiguration "github.com/upbound/provider-aws/internal/controller/backup/vaultlockconfiguration"
	vaultnotifications "github.com/upbound/provider-aws/internal/controller/backup/vaultnotifications"
	vaultpolicy "github.com/upbound/provider-aws/internal/controller/backup/vaultpolicy"
	jobdefinition "github.com/upbound/provider-aws/internal/controller/batch/jobdefinition"
	schedulingpolicy "github.com/upbound/provider-aws/internal/controller/batch/schedulingpolicy"
	budget "github.com/upbound/provider-aws/internal/controller/budgets/budget"
	budgetaction "github.com/upbound/provider-aws/internal/controller/budgets/budgetaction"
	anomalymonitor "github.com/upbound/provider-aws/internal/controller/ce/anomalymonitor"
	voiceconnector "github.com/upbound/provider-aws/internal/controller/chime/voiceconnector"
	voiceconnectorgroup "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectorgroup"
	voiceconnectorlogging "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectorlogging"
	voiceconnectororigination "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectororigination"
	voiceconnectorstreaming "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectorstreaming"
	voiceconnectortermination "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectortermination"
	voiceconnectorterminationcredentials "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectorterminationcredentials"
	environmentec2 "github.com/upbound/provider-aws/internal/controller/cloud9/environmentec2"
	environmentmembership "github.com/upbound/provider-aws/internal/controller/cloud9/environmentmembership"
	resourcecloudcontrol "github.com/upbound/provider-aws/internal/controller/cloudcontrol/resource"
	stackcloudformation "github.com/upbound/provider-aws/internal/controller/cloudformation/stack"
	stackset "github.com/upbound/provider-aws/internal/controller/cloudformation/stackset"
	stacksetinstance "github.com/upbound/provider-aws/internal/controller/cloudformation/stacksetinstance"
	cachepolicy "github.com/upbound/provider-aws/internal/controller/cloudfront/cachepolicy"
	distribution "github.com/upbound/provider-aws/internal/controller/cloudfront/distribution"
	fieldlevelencryptionconfig "github.com/upbound/provider-aws/internal/controller/cloudfront/fieldlevelencryptionconfig"
	fieldlevelencryptionprofile "github.com/upbound/provider-aws/internal/controller/cloudfront/fieldlevelencryptionprofile"
	functioncloudfront "github.com/upbound/provider-aws/internal/controller/cloudfront/function"
	keygroup "github.com/upbound/provider-aws/internal/controller/cloudfront/keygroup"
	monitoringsubscription "github.com/upbound/provider-aws/internal/controller/cloudfront/monitoringsubscription"
	originaccesscontrol "github.com/upbound/provider-aws/internal/controller/cloudfront/originaccesscontrol"
	originaccessidentity "github.com/upbound/provider-aws/internal/controller/cloudfront/originaccessidentity"
	originrequestpolicy "github.com/upbound/provider-aws/internal/controller/cloudfront/originrequestpolicy"
	publickey "github.com/upbound/provider-aws/internal/controller/cloudfront/publickey"
	realtimelogconfig "github.com/upbound/provider-aws/internal/controller/cloudfront/realtimelogconfig"
	responseheaderspolicy "github.com/upbound/provider-aws/internal/controller/cloudfront/responseheaderspolicy"
	domain "github.com/upbound/provider-aws/internal/controller/cloudsearch/domain"
	domainserviceaccesspolicy "github.com/upbound/provider-aws/internal/controller/cloudsearch/domainserviceaccesspolicy"
	eventdatastore "github.com/upbound/provider-aws/internal/controller/cloudtrail/eventdatastore"
	trail "github.com/upbound/provider-aws/internal/controller/cloudtrail/trail"
	compositealarm "github.com/upbound/provider-aws/internal/controller/cloudwatch/compositealarm"
	dashboard "github.com/upbound/provider-aws/internal/controller/cloudwatch/dashboard"
	metricalarm "github.com/upbound/provider-aws/internal/controller/cloudwatch/metricalarm"
	metricstream "github.com/upbound/provider-aws/internal/controller/cloudwatch/metricstream"
	apidestination "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/apidestination"
	archive "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/archive"
	bus "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/bus"
	buspolicy "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/buspolicy"
	connectioncloudwatchevents "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/connection"
	permissioncloudwatchevents "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/permission"
	rule "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/rule"
	targetcloudwatchevents "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/target"
	definition "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/definition"
	destination "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/destination"
	destinationpolicy "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/destinationpolicy"
	group "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/group"
	metricfilter "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/metricfilter"
	resourcepolicy "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/resourcepolicy"
	stream "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/stream"
	subscriptionfilter "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/subscriptionfilter"
	approvalruletemplate "github.com/upbound/provider-aws/internal/controller/codecommit/approvalruletemplate"
	approvalruletemplateassociation "github.com/upbound/provider-aws/internal/controller/codecommit/approvalruletemplateassociation"
	repository "github.com/upbound/provider-aws/internal/controller/codecommit/repository"
	trigger "github.com/upbound/provider-aws/internal/controller/codecommit/trigger"
	codepipeline "github.com/upbound/provider-aws/internal/controller/codepipeline/codepipeline"
	customactiontype "github.com/upbound/provider-aws/internal/controller/codepipeline/customactiontype"
	webhookcodepipeline "github.com/upbound/provider-aws/internal/controller/codepipeline/webhook"
	connectioncodestarconnections "github.com/upbound/provider-aws/internal/controller/codestarconnections/connection"
	host "github.com/upbound/provider-aws/internal/controller/codestarconnections/host"
	notificationrule "github.com/upbound/provider-aws/internal/controller/codestarnotifications/notificationrule"
	cognitoidentitypoolproviderprincipaltag "github.com/upbound/provider-aws/internal/controller/cognitoidentity/cognitoidentitypoolproviderprincipaltag"
	pool "github.com/upbound/provider-aws/internal/controller/cognitoidentity/pool"
	poolrolesattachment "github.com/upbound/provider-aws/internal/controller/cognitoidentity/poolrolesattachment"
	identityprovider "github.com/upbound/provider-aws/internal/controller/cognitoidp/identityprovider"
	resourceserver "github.com/upbound/provider-aws/internal/controller/cognitoidp/resourceserver"
	riskconfiguration "github.com/upbound/provider-aws/internal/controller/cognitoidp/riskconfiguration"
	usercognitoidp "github.com/upbound/provider-aws/internal/controller/cognitoidp/user"
	usergroup "github.com/upbound/provider-aws/internal/controller/cognitoidp/usergroup"
	useringroup "github.com/upbound/provider-aws/internal/controller/cognitoidp/useringroup"
	userpool "github.com/upbound/provider-aws/internal/controller/cognitoidp/userpool"
	userpoolclient "github.com/upbound/provider-aws/internal/controller/cognitoidp/userpoolclient"
	userpooldomain "github.com/upbound/provider-aws/internal/controller/cognitoidp/userpooldomain"
	userpooluicustomization "github.com/upbound/provider-aws/internal/controller/cognitoidp/userpooluicustomization"
	awsconfigurationrecorderstatus "github.com/upbound/provider-aws/internal/controller/configservice/awsconfigurationrecorderstatus"
	configrule "github.com/upbound/provider-aws/internal/controller/configservice/configrule"
	configurationaggregator "github.com/upbound/provider-aws/internal/controller/configservice/configurationaggregator"
	configurationrecorder "github.com/upbound/provider-aws/internal/controller/configservice/configurationrecorder"
	conformancepack "github.com/upbound/provider-aws/internal/controller/configservice/conformancepack"
	deliverychannel "github.com/upbound/provider-aws/internal/controller/configservice/deliverychannel"
	remediationconfiguration "github.com/upbound/provider-aws/internal/controller/configservice/remediationconfiguration"
	botassociation "github.com/upbound/provider-aws/internal/controller/connect/botassociation"
	contactflow "github.com/upbound/provider-aws/internal/controller/connect/contactflow"
	contactflowmodule "github.com/upbound/provider-aws/internal/controller/connect/contactflowmodule"
	hoursofoperation "github.com/upbound/provider-aws/internal/controller/connect/hoursofoperation"
	instance "github.com/upbound/provider-aws/internal/controller/connect/instance"
	instancestorageconfig "github.com/upbound/provider-aws/internal/controller/connect/instancestorageconfig"
	lambdafunctionassociation "github.com/upbound/provider-aws/internal/controller/connect/lambdafunctionassociation"
	phonenumber "github.com/upbound/provider-aws/internal/controller/connect/phonenumber"
	queue "github.com/upbound/provider-aws/internal/controller/connect/queue"
	quickconnect "github.com/upbound/provider-aws/internal/controller/connect/quickconnect"
	routingprofile "github.com/upbound/provider-aws/internal/controller/connect/routingprofile"
	securityprofile "github.com/upbound/provider-aws/internal/controller/connect/securityprofile"
	userconnect "github.com/upbound/provider-aws/internal/controller/connect/user"
	userhierarchystructure "github.com/upbound/provider-aws/internal/controller/connect/userhierarchystructure"
	vocabulary "github.com/upbound/provider-aws/internal/controller/connect/vocabulary"
	reportdefinition "github.com/upbound/provider-aws/internal/controller/cur/reportdefinition"
	dataset "github.com/upbound/provider-aws/internal/controller/dataexchange/dataset"
	revision "github.com/upbound/provider-aws/internal/controller/dataexchange/revision"
	pipeline "github.com/upbound/provider-aws/internal/controller/datapipeline/pipeline"
	locations3 "github.com/upbound/provider-aws/internal/controller/datasync/locations3"
	task "github.com/upbound/provider-aws/internal/controller/datasync/task"
	cluster "github.com/upbound/provider-aws/internal/controller/dax/cluster"
	parametergroup "github.com/upbound/provider-aws/internal/controller/dax/parametergroup"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/dax/subnetgroup"
	appdeploy "github.com/upbound/provider-aws/internal/controller/deploy/app"
	deploymentconfig "github.com/upbound/provider-aws/internal/controller/deploy/deploymentconfig"
	deploymentgroup "github.com/upbound/provider-aws/internal/controller/deploy/deploymentgroup"
	graph "github.com/upbound/provider-aws/internal/controller/detective/graph"
	invitationaccepter "github.com/upbound/provider-aws/internal/controller/detective/invitationaccepter"
	member "github.com/upbound/provider-aws/internal/controller/detective/member"
	devicepool "github.com/upbound/provider-aws/internal/controller/devicefarm/devicepool"
	instanceprofile "github.com/upbound/provider-aws/internal/controller/devicefarm/instanceprofile"
	networkprofile "github.com/upbound/provider-aws/internal/controller/devicefarm/networkprofile"
	project "github.com/upbound/provider-aws/internal/controller/devicefarm/project"
	testgridproject "github.com/upbound/provider-aws/internal/controller/devicefarm/testgridproject"
	upload "github.com/upbound/provider-aws/internal/controller/devicefarm/upload"
	bgppeer "github.com/upbound/provider-aws/internal/controller/directconnect/bgppeer"
	connectiondirectconnect "github.com/upbound/provider-aws/internal/controller/directconnect/connection"
	connectionassociation "github.com/upbound/provider-aws/internal/controller/directconnect/connectionassociation"
	gateway "github.com/upbound/provider-aws/internal/controller/directconnect/gateway"
	gatewayassociation "github.com/upbound/provider-aws/internal/controller/directconnect/gatewayassociation"
	gatewayassociationproposal "github.com/upbound/provider-aws/internal/controller/directconnect/gatewayassociationproposal"
	hostedprivatevirtualinterface "github.com/upbound/provider-aws/internal/controller/directconnect/hostedprivatevirtualinterface"
	hostedprivatevirtualinterfaceaccepter "github.com/upbound/provider-aws/internal/controller/directconnect/hostedprivatevirtualinterfaceaccepter"
	hostedpublicvirtualinterface "github.com/upbound/provider-aws/internal/controller/directconnect/hostedpublicvirtualinterface"
	hostedpublicvirtualinterfaceaccepter "github.com/upbound/provider-aws/internal/controller/directconnect/hostedpublicvirtualinterfaceaccepter"
	hostedtransitvirtualinterface "github.com/upbound/provider-aws/internal/controller/directconnect/hostedtransitvirtualinterface"
	hostedtransitvirtualinterfaceaccepter "github.com/upbound/provider-aws/internal/controller/directconnect/hostedtransitvirtualinterfaceaccepter"
	lag "github.com/upbound/provider-aws/internal/controller/directconnect/lag"
	privatevirtualinterface "github.com/upbound/provider-aws/internal/controller/directconnect/privatevirtualinterface"
	publicvirtualinterface "github.com/upbound/provider-aws/internal/controller/directconnect/publicvirtualinterface"
	transitvirtualinterface "github.com/upbound/provider-aws/internal/controller/directconnect/transitvirtualinterface"
	lifecyclepolicy "github.com/upbound/provider-aws/internal/controller/dlm/lifecyclepolicy"
	certificatedms "github.com/upbound/provider-aws/internal/controller/dms/certificate"
	endpoint "github.com/upbound/provider-aws/internal/controller/dms/endpoint"
	eventsubscription "github.com/upbound/provider-aws/internal/controller/dms/eventsubscription"
	replicationinstance "github.com/upbound/provider-aws/internal/controller/dms/replicationinstance"
	replicationsubnetgroup "github.com/upbound/provider-aws/internal/controller/dms/replicationsubnetgroup"
	replicationtask "github.com/upbound/provider-aws/internal/controller/dms/replicationtask"
	s3endpoint "github.com/upbound/provider-aws/internal/controller/dms/s3endpoint"
	clusterdocdb "github.com/upbound/provider-aws/internal/controller/docdb/cluster"
	clusterinstance "github.com/upbound/provider-aws/internal/controller/docdb/clusterinstance"
	clusterparametergroup "github.com/upbound/provider-aws/internal/controller/docdb/clusterparametergroup"
	clustersnapshot "github.com/upbound/provider-aws/internal/controller/docdb/clustersnapshot"
	eventsubscriptiondocdb "github.com/upbound/provider-aws/internal/controller/docdb/eventsubscription"
	globalcluster "github.com/upbound/provider-aws/internal/controller/docdb/globalcluster"
	subnetgroupdocdb "github.com/upbound/provider-aws/internal/controller/docdb/subnetgroup"
	conditionalforwarder "github.com/upbound/provider-aws/internal/controller/ds/conditionalforwarder"
	directory "github.com/upbound/provider-aws/internal/controller/ds/directory"
	shareddirectory "github.com/upbound/provider-aws/internal/controller/ds/shareddirectory"
	contributorinsights "github.com/upbound/provider-aws/internal/controller/dynamodb/contributorinsights"
	globaltable "github.com/upbound/provider-aws/internal/controller/dynamodb/globaltable"
	kinesisstreamingdestination "github.com/upbound/provider-aws/internal/controller/dynamodb/kinesisstreamingdestination"
	table "github.com/upbound/provider-aws/internal/controller/dynamodb/table"
	tableitem "github.com/upbound/provider-aws/internal/controller/dynamodb/tableitem"
	tablereplica "github.com/upbound/provider-aws/internal/controller/dynamodb/tablereplica"
	tag "github.com/upbound/provider-aws/internal/controller/dynamodb/tag"
	ami "github.com/upbound/provider-aws/internal/controller/ec2/ami"
	amicopy "github.com/upbound/provider-aws/internal/controller/ec2/amicopy"
	amilaunchpermission "github.com/upbound/provider-aws/internal/controller/ec2/amilaunchpermission"
	availabilityzonegroup "github.com/upbound/provider-aws/internal/controller/ec2/availabilityzonegroup"
	capacityreservation "github.com/upbound/provider-aws/internal/controller/ec2/capacityreservation"
	carriergateway "github.com/upbound/provider-aws/internal/controller/ec2/carriergateway"
	customergateway "github.com/upbound/provider-aws/internal/controller/ec2/customergateway"
	defaultnetworkacl "github.com/upbound/provider-aws/internal/controller/ec2/defaultnetworkacl"
	defaultroutetable "github.com/upbound/provider-aws/internal/controller/ec2/defaultroutetable"
	defaultsecuritygroup "github.com/upbound/provider-aws/internal/controller/ec2/defaultsecuritygroup"
	defaultsubnet "github.com/upbound/provider-aws/internal/controller/ec2/defaultsubnet"
	defaultvpc "github.com/upbound/provider-aws/internal/controller/ec2/defaultvpc"
	defaultvpcdhcpoptions "github.com/upbound/provider-aws/internal/controller/ec2/defaultvpcdhcpoptions"
	ebsdefaultkmskey "github.com/upbound/provider-aws/internal/controller/ec2/ebsdefaultkmskey"
	ebsencryptionbydefault "github.com/upbound/provider-aws/internal/controller/ec2/ebsencryptionbydefault"
	ebssnapshot "github.com/upbound/provider-aws/internal/controller/ec2/ebssnapshot"
	ebssnapshotcopy "github.com/upbound/provider-aws/internal/controller/ec2/ebssnapshotcopy"
	ebssnapshotimport "github.com/upbound/provider-aws/internal/controller/ec2/ebssnapshotimport"
	ebsvolume "github.com/upbound/provider-aws/internal/controller/ec2/ebsvolume"
	egressonlyinternetgateway "github.com/upbound/provider-aws/internal/controller/ec2/egressonlyinternetgateway"
	eip "github.com/upbound/provider-aws/internal/controller/ec2/eip"
	eipassociation "github.com/upbound/provider-aws/internal/controller/ec2/eipassociation"
	flowlog "github.com/upbound/provider-aws/internal/controller/ec2/flowlog"
	hostec2 "github.com/upbound/provider-aws/internal/controller/ec2/host"
	instanceec2 "github.com/upbound/provider-aws/internal/controller/ec2/instance"
	instancestate "github.com/upbound/provider-aws/internal/controller/ec2/instancestate"
	internetgateway "github.com/upbound/provider-aws/internal/controller/ec2/internetgateway"
	keypair "github.com/upbound/provider-aws/internal/controller/ec2/keypair"
	launchtemplate "github.com/upbound/provider-aws/internal/controller/ec2/launchtemplate"
	mainroutetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/mainroutetableassociation"
	managedprefixlist "github.com/upbound/provider-aws/internal/controller/ec2/managedprefixlist"
	managedprefixlistentry "github.com/upbound/provider-aws/internal/controller/ec2/managedprefixlistentry"
	natgateway "github.com/upbound/provider-aws/internal/controller/ec2/natgateway"
	networkacl "github.com/upbound/provider-aws/internal/controller/ec2/networkacl"
	networkaclrule "github.com/upbound/provider-aws/internal/controller/ec2/networkaclrule"
	networkinsightsanalysis "github.com/upbound/provider-aws/internal/controller/ec2/networkinsightsanalysis"
	networkinsightspath "github.com/upbound/provider-aws/internal/controller/ec2/networkinsightspath"
	networkinterface "github.com/upbound/provider-aws/internal/controller/ec2/networkinterface"
	networkinterfaceattachment "github.com/upbound/provider-aws/internal/controller/ec2/networkinterfaceattachment"
	networkinterfacesgattachment "github.com/upbound/provider-aws/internal/controller/ec2/networkinterfacesgattachment"
	placementgroup "github.com/upbound/provider-aws/internal/controller/ec2/placementgroup"
	routeec2 "github.com/upbound/provider-aws/internal/controller/ec2/route"
	routetable "github.com/upbound/provider-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/routetableassociation"
	securitygroup "github.com/upbound/provider-aws/internal/controller/ec2/securitygroup"
	securitygroupegressrule "github.com/upbound/provider-aws/internal/controller/ec2/securitygroupegressrule"
	securitygroupingressrule "github.com/upbound/provider-aws/internal/controller/ec2/securitygroupingressrule"
	securitygrouprule "github.com/upbound/provider-aws/internal/controller/ec2/securitygrouprule"
	serialconsoleaccess "github.com/upbound/provider-aws/internal/controller/ec2/serialconsoleaccess"
	snapshotcreatevolumepermission "github.com/upbound/provider-aws/internal/controller/ec2/snapshotcreatevolumepermission"
	spotdatafeedsubscription "github.com/upbound/provider-aws/internal/controller/ec2/spotdatafeedsubscription"
	spotfleetrequest "github.com/upbound/provider-aws/internal/controller/ec2/spotfleetrequest"
	spotinstancerequest "github.com/upbound/provider-aws/internal/controller/ec2/spotinstancerequest"
	subnet "github.com/upbound/provider-aws/internal/controller/ec2/subnet"
	subnetcidrreservation "github.com/upbound/provider-aws/internal/controller/ec2/subnetcidrreservation"
	tagec2 "github.com/upbound/provider-aws/internal/controller/ec2/tag"
	trafficmirrorfilter "github.com/upbound/provider-aws/internal/controller/ec2/trafficmirrorfilter"
	trafficmirrorfilterrule "github.com/upbound/provider-aws/internal/controller/ec2/trafficmirrorfilterrule"
	transitgateway "github.com/upbound/provider-aws/internal/controller/ec2/transitgateway"
	transitgatewayconnect "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayconnect"
	transitgatewayconnectpeer "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayconnectpeer"
	transitgatewaymulticastdomain "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaymulticastdomain"
	transitgatewaymulticastdomainassociation "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaymulticastdomainassociation"
	transitgatewaymulticastgroupmember "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaymulticastgroupmember"
	transitgatewaymulticastgroupsource "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaymulticastgroupsource"
	transitgatewaypeeringattachment "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaypeeringattachment"
	transitgatewaypeeringattachmentaccepter "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaypeeringattachmentaccepter"
	transitgatewaypolicytable "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaypolicytable"
	transitgatewayprefixlistreference "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayprefixlistreference"
	transitgatewayroute "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayvpcattachmentaccepter"
	volumeattachment "github.com/upbound/provider-aws/internal/controller/ec2/volumeattachment"
	vpc "github.com/upbound/provider-aws/internal/controller/ec2/vpc"
	vpcdhcpoptions "github.com/upbound/provider-aws/internal/controller/ec2/vpcdhcpoptions"
	vpcdhcpoptionsassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcdhcpoptionsassociation"
	vpcendpoint "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpoint"
	vpcendpointconnectionnotification "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointconnectionnotification"
	vpcendpointroutetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointroutetableassociation"
	vpcendpointsecuritygroupassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointsecuritygroupassociation"
	vpcendpointservice "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointservice"
	vpcendpointserviceallowedprincipal "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointserviceallowedprincipal"
	vpcendpointsubnetassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointsubnetassociation"
	vpcipam "github.com/upbound/provider-aws/internal/controller/ec2/vpcipam"
	vpcipampool "github.com/upbound/provider-aws/internal/controller/ec2/vpcipampool"
	vpcipampoolcidr "github.com/upbound/provider-aws/internal/controller/ec2/vpcipampoolcidr"
	vpcipampoolcidrallocation "github.com/upbound/provider-aws/internal/controller/ec2/vpcipampoolcidrallocation"
	vpcipamscope "github.com/upbound/provider-aws/internal/controller/ec2/vpcipamscope"
	vpcipv4cidrblockassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/upbound/provider-aws/internal/controller/ec2/vpcpeeringconnection"
	vpcpeeringconnectionaccepter "github.com/upbound/provider-aws/internal/controller/ec2/vpcpeeringconnectionaccepter"
	vpcpeeringconnectionoptions "github.com/upbound/provider-aws/internal/controller/ec2/vpcpeeringconnectionoptions"
	vpnconnection "github.com/upbound/provider-aws/internal/controller/ec2/vpnconnection"
	vpnconnectionroute "github.com/upbound/provider-aws/internal/controller/ec2/vpnconnectionroute"
	vpngateway "github.com/upbound/provider-aws/internal/controller/ec2/vpngateway"
	vpngatewayattachment "github.com/upbound/provider-aws/internal/controller/ec2/vpngatewayattachment"
	vpngatewayroutepropagation "github.com/upbound/provider-aws/internal/controller/ec2/vpngatewayroutepropagation"
	lifecyclepolicyecr "github.com/upbound/provider-aws/internal/controller/ecr/lifecyclepolicy"
	pullthroughcacherule "github.com/upbound/provider-aws/internal/controller/ecr/pullthroughcacherule"
	registrypolicy "github.com/upbound/provider-aws/internal/controller/ecr/registrypolicy"
	registryscanningconfiguration "github.com/upbound/provider-aws/internal/controller/ecr/registryscanningconfiguration"
	replicationconfiguration "github.com/upbound/provider-aws/internal/controller/ecr/replicationconfiguration"
	repositoryecr "github.com/upbound/provider-aws/internal/controller/ecr/repository"
	repositorypolicy "github.com/upbound/provider-aws/internal/controller/ecr/repositorypolicy"
	repositoryecrpublic "github.com/upbound/provider-aws/internal/controller/ecrpublic/repository"
	repositorypolicyecrpublic "github.com/upbound/provider-aws/internal/controller/ecrpublic/repositorypolicy"
	accountsettingdefault "github.com/upbound/provider-aws/internal/controller/ecs/accountsettingdefault"
	capacityprovider "github.com/upbound/provider-aws/internal/controller/ecs/capacityprovider"
	clusterecs "github.com/upbound/provider-aws/internal/controller/ecs/cluster"
	clustercapacityproviders "github.com/upbound/provider-aws/internal/controller/ecs/clustercapacityproviders"
	serviceecs "github.com/upbound/provider-aws/internal/controller/ecs/service"
	taskdefinition "github.com/upbound/provider-aws/internal/controller/ecs/taskdefinition"
	accesspoint "github.com/upbound/provider-aws/internal/controller/efs/accesspoint"
	backuppolicy "github.com/upbound/provider-aws/internal/controller/efs/backuppolicy"
	filesystem "github.com/upbound/provider-aws/internal/controller/efs/filesystem"
	filesystempolicy "github.com/upbound/provider-aws/internal/controller/efs/filesystempolicy"
	mounttarget "github.com/upbound/provider-aws/internal/controller/efs/mounttarget"
	replicationconfigurationefs "github.com/upbound/provider-aws/internal/controller/efs/replicationconfiguration"
	addon "github.com/upbound/provider-aws/internal/controller/eks/addon"
	clustereks "github.com/upbound/provider-aws/internal/controller/eks/cluster"
	clusterauth "github.com/upbound/provider-aws/internal/controller/eks/clusterauth"
	fargateprofile "github.com/upbound/provider-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/upbound/provider-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/upbound/provider-aws/internal/controller/eks/nodegroup"
	podidentityassociation "github.com/upbound/provider-aws/internal/controller/eks/podidentityassociation"
	clusterelasticache "github.com/upbound/provider-aws/internal/controller/elasticache/cluster"
	parametergroupelasticache "github.com/upbound/provider-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/upbound/provider-aws/internal/controller/elasticache/replicationgroup"
	subnetgroupelasticache "github.com/upbound/provider-aws/internal/controller/elasticache/subnetgroup"
	userelasticache "github.com/upbound/provider-aws/internal/controller/elasticache/user"
	usergroupelasticache "github.com/upbound/provider-aws/internal/controller/elasticache/usergroup"
	applicationelasticbeanstalk "github.com/upbound/provider-aws/internal/controller/elasticbeanstalk/application"
	applicationversion "github.com/upbound/provider-aws/internal/controller/elasticbeanstalk/applicationversion"
	configurationtemplate "github.com/upbound/provider-aws/internal/controller/elasticbeanstalk/configurationtemplate"
	domainelasticsearch "github.com/upbound/provider-aws/internal/controller/elasticsearch/domain"
	domainpolicy "github.com/upbound/provider-aws/internal/controller/elasticsearch/domainpolicy"
	domainsamloptions "github.com/upbound/provider-aws/internal/controller/elasticsearch/domainsamloptions"
	pipelineelastictranscoder "github.com/upbound/provider-aws/internal/controller/elastictranscoder/pipeline"
	preset "github.com/upbound/provider-aws/internal/controller/elastictranscoder/preset"
	appcookiestickinesspolicy "github.com/upbound/provider-aws/internal/controller/elb/appcookiestickinesspolicy"
	attachmentelb "github.com/upbound/provider-aws/internal/controller/elb/attachment"
	backendserverpolicy "github.com/upbound/provider-aws/internal/controller/elb/backendserverpolicy"
	elb "github.com/upbound/provider-aws/internal/controller/elb/elb"
	lbcookiestickinesspolicy "github.com/upbound/provider-aws/internal/controller/elb/lbcookiestickinesspolicy"
	lbsslnegotiationpolicy "github.com/upbound/provider-aws/internal/controller/elb/lbsslnegotiationpolicy"
	listenerpolicy "github.com/upbound/provider-aws/internal/controller/elb/listenerpolicy"
	policyelb "github.com/upbound/provider-aws/internal/controller/elb/policy"
	proxyprotocolpolicy "github.com/upbound/provider-aws/internal/controller/elb/proxyprotocolpolicy"
	lb "github.com/upbound/provider-aws/internal/controller/elbv2/lb"
	lblistener "github.com/upbound/provider-aws/internal/controller/elbv2/lblistener"
	lblistenercertificate "github.com/upbound/provider-aws/internal/controller/elbv2/lblistenercertificate"
	lblistenerrule "github.com/upbound/provider-aws/internal/controller/elbv2/lblistenerrule"
	lbtargetgroup "github.com/upbound/provider-aws/internal/controller/elbv2/lbtargetgroup"
	lbtargetgroupattachment "github.com/upbound/provider-aws/internal/controller/elbv2/lbtargetgroupattachment"
	securityconfiguration "github.com/upbound/provider-aws/internal/controller/emr/securityconfiguration"
	applicationemrserverless "github.com/upbound/provider-aws/internal/controller/emrserverless/application"
	feature "github.com/upbound/provider-aws/internal/controller/evidently/feature"
	projectevidently "github.com/upbound/provider-aws/internal/controller/evidently/project"
	segment "github.com/upbound/provider-aws/internal/controller/evidently/segment"
	deliverystream "github.com/upbound/provider-aws/internal/controller/firehose/deliverystream"
	experimenttemplate "github.com/upbound/provider-aws/internal/controller/fis/experimenttemplate"
	backup "github.com/upbound/provider-aws/internal/controller/fsx/backup"
	datarepositoryassociation "github.com/upbound/provider-aws/internal/controller/fsx/datarepositoryassociation"
	lustrefilesystem "github.com/upbound/provider-aws/internal/controller/fsx/lustrefilesystem"
	ontapfilesystem "github.com/upbound/provider-aws/internal/controller/fsx/ontapfilesystem"
	ontapstoragevirtualmachine "github.com/upbound/provider-aws/internal/controller/fsx/ontapstoragevirtualmachine"
	windowsfilesystem "github.com/upbound/provider-aws/internal/controller/fsx/windowsfilesystem"
	alias "github.com/upbound/provider-aws/internal/controller/gamelift/alias"
	build "github.com/upbound/provider-aws/internal/controller/gamelift/build"
	fleetgamelift "github.com/upbound/provider-aws/internal/controller/gamelift/fleet"
	gamesessionqueue "github.com/upbound/provider-aws/internal/controller/gamelift/gamesessionqueue"
	script "github.com/upbound/provider-aws/internal/controller/gamelift/script"
	vaultglacier "github.com/upbound/provider-aws/internal/controller/glacier/vault"
	vaultlock "github.com/upbound/provider-aws/internal/controller/glacier/vaultlock"
	accelerator "github.com/upbound/provider-aws/internal/controller/globalaccelerator/accelerator"
	endpointgroup "github.com/upbound/provider-aws/internal/controller/globalaccelerator/endpointgroup"
	listener "github.com/upbound/provider-aws/internal/controller/globalaccelerator/listener"
	catalogdatabase "github.com/upbound/provider-aws/internal/controller/glue/catalogdatabase"
	catalogtable "github.com/upbound/provider-aws/internal/controller/glue/catalogtable"
	classifier "github.com/upbound/provider-aws/internal/controller/glue/classifier"
	connectionglue "github.com/upbound/provider-aws/internal/controller/glue/connection"
	crawler "github.com/upbound/provider-aws/internal/controller/glue/crawler"
	datacatalogencryptionsettings "github.com/upbound/provider-aws/internal/controller/glue/datacatalogencryptionsettings"
	job "github.com/upbound/provider-aws/internal/controller/glue/job"
	registry "github.com/upbound/provider-aws/internal/controller/glue/registry"
	resourcepolicyglue "github.com/upbound/provider-aws/internal/controller/glue/resourcepolicy"
	schema "github.com/upbound/provider-aws/internal/controller/glue/schema"
	securityconfigurationglue "github.com/upbound/provider-aws/internal/controller/glue/securityconfiguration"
	triggerglue "github.com/upbound/provider-aws/internal/controller/glue/trigger"
	userdefinedfunction "github.com/upbound/provider-aws/internal/controller/glue/userdefinedfunction"
	workflow "github.com/upbound/provider-aws/internal/controller/glue/workflow"
	licenseassociation "github.com/upbound/provider-aws/internal/controller/grafana/licenseassociation"
	roleassociation "github.com/upbound/provider-aws/internal/controller/grafana/roleassociation"
	workspacegrafana "github.com/upbound/provider-aws/internal/controller/grafana/workspace"
	workspaceapikey "github.com/upbound/provider-aws/internal/controller/grafana/workspaceapikey"
	workspacesamlconfiguration "github.com/upbound/provider-aws/internal/controller/grafana/workspacesamlconfiguration"
	detector "github.com/upbound/provider-aws/internal/controller/guardduty/detector"
	filter "github.com/upbound/provider-aws/internal/controller/guardduty/filter"
	memberguardduty "github.com/upbound/provider-aws/internal/controller/guardduty/member"
	accesskey "github.com/upbound/provider-aws/internal/controller/iam/accesskey"
	accountalias "github.com/upbound/provider-aws/internal/controller/iam/accountalias"
	accountpasswordpolicy "github.com/upbound/provider-aws/internal/controller/iam/accountpasswordpolicy"
	groupiam "github.com/upbound/provider-aws/internal/controller/iam/group"
	groupmembership "github.com/upbound/provider-aws/internal/controller/iam/groupmembership"
	grouppolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/grouppolicyattachment"
	instanceprofileiam "github.com/upbound/provider-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/upbound/provider-aws/internal/controller/iam/openidconnectprovider"
	policyiam "github.com/upbound/provider-aws/internal/controller/iam/policy"
	role "github.com/upbound/provider-aws/internal/controller/iam/role"
	rolepolicy "github.com/upbound/provider-aws/internal/controller/iam/rolepolicy"
	rolepolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/rolepolicyattachment"
	samlprovider "github.com/upbound/provider-aws/internal/controller/iam/samlprovider"
	servercertificate "github.com/upbound/provider-aws/internal/controller/iam/servercertificate"
	servicelinkedrole "github.com/upbound/provider-aws/internal/controller/iam/servicelinkedrole"
	servicespecificcredential "github.com/upbound/provider-aws/internal/controller/iam/servicespecificcredential"
	signingcertificate "github.com/upbound/provider-aws/internal/controller/iam/signingcertificate"
	useriam "github.com/upbound/provider-aws/internal/controller/iam/user"
	usergroupmembership "github.com/upbound/provider-aws/internal/controller/iam/usergroupmembership"
	userloginprofile "github.com/upbound/provider-aws/internal/controller/iam/userloginprofile"
	userpolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/userpolicyattachment"
	usersshkey "github.com/upbound/provider-aws/internal/controller/iam/usersshkey"
	virtualmfadevice "github.com/upbound/provider-aws/internal/controller/iam/virtualmfadevice"
	groupidentitystore "github.com/upbound/provider-aws/internal/controller/identitystore/group"
	groupmembershipidentitystore "github.com/upbound/provider-aws/internal/controller/identitystore/groupmembership"
	useridentitystore "github.com/upbound/provider-aws/internal/controller/identitystore/user"
	component "github.com/upbound/provider-aws/internal/controller/imagebuilder/component"
	containerrecipe "github.com/upbound/provider-aws/internal/controller/imagebuilder/containerrecipe"
	distributionconfiguration "github.com/upbound/provider-aws/internal/controller/imagebuilder/distributionconfiguration"
	image "github.com/upbound/provider-aws/internal/controller/imagebuilder/image"
	imagepipeline "github.com/upbound/provider-aws/internal/controller/imagebuilder/imagepipeline"
	imagerecipe "github.com/upbound/provider-aws/internal/controller/imagebuilder/imagerecipe"
	infrastructureconfiguration "github.com/upbound/provider-aws/internal/controller/imagebuilder/infrastructureconfiguration"
	assessmenttarget "github.com/upbound/provider-aws/internal/controller/inspector/assessmenttarget"
	assessmenttemplate "github.com/upbound/provider-aws/internal/controller/inspector/assessmenttemplate"
	resourcegroup "github.com/upbound/provider-aws/internal/controller/inspector/resourcegroup"
	enabler "github.com/upbound/provider-aws/internal/controller/inspector2/enabler"
	certificateiot "github.com/upbound/provider-aws/internal/controller/iot/certificate"
	indexingconfiguration "github.com/upbound/provider-aws/internal/controller/iot/indexingconfiguration"
	loggingoptions "github.com/upbound/provider-aws/internal/controller/iot/loggingoptions"
	policyiot "github.com/upbound/provider-aws/internal/controller/iot/policy"
	policyattachment "github.com/upbound/provider-aws/internal/controller/iot/policyattachment"
	provisioningtemplate "github.com/upbound/provider-aws/internal/controller/iot/provisioningtemplate"
	rolealias "github.com/upbound/provider-aws/internal/controller/iot/rolealias"
	thing "github.com/upbound/provider-aws/internal/controller/iot/thing"
	thinggroup "github.com/upbound/provider-aws/internal/controller/iot/thinggroup"
	thinggroupmembership "github.com/upbound/provider-aws/internal/controller/iot/thinggroupmembership"
	thingprincipalattachment "github.com/upbound/provider-aws/internal/controller/iot/thingprincipalattachment"
	thingtype "github.com/upbound/provider-aws/internal/controller/iot/thingtype"
	topicrule "github.com/upbound/provider-aws/internal/controller/iot/topicrule"
	topicruledestination "github.com/upbound/provider-aws/internal/controller/iot/topicruledestination"
	channel "github.com/upbound/provider-aws/internal/controller/ivs/channel"
	recordingconfiguration "github.com/upbound/provider-aws/internal/controller/ivs/recordingconfiguration"
	clusterkafka "github.com/upbound/provider-aws/internal/controller/kafka/cluster"
	configuration "github.com/upbound/provider-aws/internal/controller/kafka/configuration"
	scramsecretassociation "github.com/upbound/provider-aws/internal/controller/kafka/scramsecretassociation"
	serverlesscluster "github.com/upbound/provider-aws/internal/controller/kafka/serverlesscluster"
	connector "github.com/upbound/provider-aws/internal/controller/kafkaconnect/connector"
	customplugin "github.com/upbound/provider-aws/internal/controller/kafkaconnect/customplugin"
	workerconfiguration "github.com/upbound/provider-aws/internal/controller/kafkaconnect/workerconfiguration"
	datasourcekendra "github.com/upbound/provider-aws/internal/controller/kendra/datasource"
	experience "github.com/upbound/provider-aws/internal/controller/kendra/experience"
	index "github.com/upbound/provider-aws/internal/controller/kendra/index"
	querysuggestionsblocklist "github.com/upbound/provider-aws/internal/controller/kendra/querysuggestionsblocklist"
	thesaurus "github.com/upbound/provider-aws/internal/controller/kendra/thesaurus"
	keyspace "github.com/upbound/provider-aws/internal/controller/keyspaces/keyspace"
	tablekeyspaces "github.com/upbound/provider-aws/internal/controller/keyspaces/table"
	streamkinesis "github.com/upbound/provider-aws/internal/controller/kinesis/stream"
	streamconsumer "github.com/upbound/provider-aws/internal/controller/kinesis/streamconsumer"
	applicationkinesisanalytics "github.com/upbound/provider-aws/internal/controller/kinesisanalytics/application"
	applicationkinesisanalyticsv2 "github.com/upbound/provider-aws/internal/controller/kinesisanalyticsv2/application"
	applicationsnapshot "github.com/upbound/provider-aws/internal/controller/kinesisanalyticsv2/applicationsnapshot"
	streamkinesisvideo "github.com/upbound/provider-aws/internal/controller/kinesisvideo/stream"
	aliaskms "github.com/upbound/provider-aws/internal/controller/kms/alias"
	ciphertext "github.com/upbound/provider-aws/internal/controller/kms/ciphertext"
	externalkey "github.com/upbound/provider-aws/internal/controller/kms/externalkey"
	grant "github.com/upbound/provider-aws/internal/controller/kms/grant"
	key "github.com/upbound/provider-aws/internal/controller/kms/key"
	replicaexternalkey "github.com/upbound/provider-aws/internal/controller/kms/replicaexternalkey"
	replicakey "github.com/upbound/provider-aws/internal/controller/kms/replicakey"
	datalakesettings "github.com/upbound/provider-aws/internal/controller/lakeformation/datalakesettings"
	permissions "github.com/upbound/provider-aws/internal/controller/lakeformation/permissions"
	resourcelakeformation "github.com/upbound/provider-aws/internal/controller/lakeformation/resource"
	aliaslambda "github.com/upbound/provider-aws/internal/controller/lambda/alias"
	codesigningconfig "github.com/upbound/provider-aws/internal/controller/lambda/codesigningconfig"
	eventsourcemapping "github.com/upbound/provider-aws/internal/controller/lambda/eventsourcemapping"
	functionlambda "github.com/upbound/provider-aws/internal/controller/lambda/function"
	functioneventinvokeconfig "github.com/upbound/provider-aws/internal/controller/lambda/functioneventinvokeconfig"
	functionurl "github.com/upbound/provider-aws/internal/controller/lambda/functionurl"
	invocation "github.com/upbound/provider-aws/internal/controller/lambda/invocation"
	layerversion "github.com/upbound/provider-aws/internal/controller/lambda/layerversion"
	layerversionpermission "github.com/upbound/provider-aws/internal/controller/lambda/layerversionpermission"
	permissionlambda "github.com/upbound/provider-aws/internal/controller/lambda/permission"
	provisionedconcurrencyconfig "github.com/upbound/provider-aws/internal/controller/lambda/provisionedconcurrencyconfig"
	bot "github.com/upbound/provider-aws/internal/controller/lexmodels/bot"
	botalias "github.com/upbound/provider-aws/internal/controller/lexmodels/botalias"
	intent "github.com/upbound/provider-aws/internal/controller/lexmodels/intent"
	slottype "github.com/upbound/provider-aws/internal/controller/lexmodels/slottype"
	association "github.com/upbound/provider-aws/internal/controller/licensemanager/association"
	licenseconfiguration "github.com/upbound/provider-aws/internal/controller/licensemanager/licenseconfiguration"
	bucket "github.com/upbound/provider-aws/internal/controller/lightsail/bucket"
	certificatelightsail "github.com/upbound/provider-aws/internal/controller/lightsail/certificate"
	containerservice "github.com/upbound/provider-aws/internal/controller/lightsail/containerservice"
	disk "github.com/upbound/provider-aws/internal/controller/lightsail/disk"
	diskattachment "github.com/upbound/provider-aws/internal/controller/lightsail/diskattachment"
	domainlightsail "github.com/upbound/provider-aws/internal/controller/lightsail/domain"
	domainentry "github.com/upbound/provider-aws/internal/controller/lightsail/domainentry"
	instancelightsail "github.com/upbound/provider-aws/internal/controller/lightsail/instance"
	instancepublicports "github.com/upbound/provider-aws/internal/controller/lightsail/instancepublicports"
	keypairlightsail "github.com/upbound/provider-aws/internal/controller/lightsail/keypair"
	lblightsail "github.com/upbound/provider-aws/internal/controller/lightsail/lb"
	lbattachment "github.com/upbound/provider-aws/internal/controller/lightsail/lbattachment"
	lbcertificate "github.com/upbound/provider-aws/internal/controller/lightsail/lbcertificate"
	lbstickinesspolicy "github.com/upbound/provider-aws/internal/controller/lightsail/lbstickinesspolicy"
	staticip "github.com/upbound/provider-aws/internal/controller/lightsail/staticip"
	staticipattachment "github.com/upbound/provider-aws/internal/controller/lightsail/staticipattachment"
	geofencecollection "github.com/upbound/provider-aws/internal/controller/location/geofencecollection"
	placeindex "github.com/upbound/provider-aws/internal/controller/location/placeindex"
	routecalculator "github.com/upbound/provider-aws/internal/controller/location/routecalculator"
	tracker "github.com/upbound/provider-aws/internal/controller/location/tracker"
	trackerassociation "github.com/upbound/provider-aws/internal/controller/location/trackerassociation"
	accountmacie2 "github.com/upbound/provider-aws/internal/controller/macie2/account"
	classificationjob "github.com/upbound/provider-aws/internal/controller/macie2/classificationjob"
	customdataidentifier "github.com/upbound/provider-aws/internal/controller/macie2/customdataidentifier"
	findingsfilter "github.com/upbound/provider-aws/internal/controller/macie2/findingsfilter"
	invitationacceptermacie2 "github.com/upbound/provider-aws/internal/controller/macie2/invitationaccepter"
	membermacie2 "github.com/upbound/provider-aws/internal/controller/macie2/member"
	queuemediaconvert "github.com/upbound/provider-aws/internal/controller/mediaconvert/queue"
	channelmedialive "github.com/upbound/provider-aws/internal/controller/medialive/channel"
	input "github.com/upbound/provider-aws/internal/controller/medialive/input"
	inputsecuritygroup "github.com/upbound/provider-aws/internal/controller/medialive/inputsecuritygroup"
	multiplex "github.com/upbound/provider-aws/internal/controller/medialive/multiplex"
	channelmediapackage "github.com/upbound/provider-aws/internal/controller/mediapackage/channel"
	container "github.com/upbound/provider-aws/internal/controller/mediastore/container"
	containerpolicy "github.com/upbound/provider-aws/internal/controller/mediastore/containerpolicy"
	acl "github.com/upbound/provider-aws/internal/controller/memorydb/acl"
	clustermemorydb "github.com/upbound/provider-aws/internal/controller/memorydb/cluster"
	parametergroupmemorydb "github.com/upbound/provider-aws/internal/controller/memorydb/parametergroup"
	snapshot "github.com/upbound/provider-aws/internal/controller/memorydb/snapshot"
	subnetgroupmemorydb "github.com/upbound/provider-aws/internal/controller/memorydb/subnetgroup"
	usermemorydb "github.com/upbound/provider-aws/internal/controller/memorydb/user"
	broker "github.com/upbound/provider-aws/internal/controller/mq/broker"
	configurationmq "github.com/upbound/provider-aws/internal/controller/mq/configuration"
	clusterneptune "github.com/upbound/provider-aws/internal/controller/neptune/cluster"
	clusterendpoint "github.com/upbound/provider-aws/internal/controller/neptune/clusterendpoint"
	clusterinstanceneptune "github.com/upbound/provider-aws/internal/controller/neptune/clusterinstance"
	clusterparametergroupneptune "github.com/upbound/provider-aws/internal/controller/neptune/clusterparametergroup"
	clustersnapshotneptune "github.com/upbound/provider-aws/internal/controller/neptune/clustersnapshot"
	eventsubscriptionneptune "github.com/upbound/provider-aws/internal/controller/neptune/eventsubscription"
	globalclusterneptune "github.com/upbound/provider-aws/internal/controller/neptune/globalcluster"
	parametergroupneptune "github.com/upbound/provider-aws/internal/controller/neptune/parametergroup"
	subnetgroupneptune "github.com/upbound/provider-aws/internal/controller/neptune/subnetgroup"
	firewall "github.com/upbound/provider-aws/internal/controller/networkfirewall/firewall"
	firewallpolicy "github.com/upbound/provider-aws/internal/controller/networkfirewall/firewallpolicy"
	loggingconfiguration "github.com/upbound/provider-aws/internal/controller/networkfirewall/loggingconfiguration"
	rulegroup "github.com/upbound/provider-aws/internal/controller/networkfirewall/rulegroup"
	attachmentaccepter "github.com/upbound/provider-aws/internal/controller/networkmanager/attachmentaccepter"
	connectattachment "github.com/upbound/provider-aws/internal/controller/networkmanager/connectattachment"
	connectionnetworkmanager "github.com/upbound/provider-aws/internal/controller/networkmanager/connection"
	corenetwork "github.com/upbound/provider-aws/internal/controller/networkmanager/corenetwork"
	customergatewayassociation "github.com/upbound/provider-aws/internal/controller/networkmanager/customergatewayassociation"
	device "github.com/upbound/provider-aws/internal/controller/networkmanager/device"
	globalnetwork "github.com/upbound/provider-aws/internal/controller/networkmanager/globalnetwork"
	link "github.com/upbound/provider-aws/internal/controller/networkmanager/link"
	linkassociation "github.com/upbound/provider-aws/internal/controller/networkmanager/linkassociation"
	site "github.com/upbound/provider-aws/internal/controller/networkmanager/site"
	transitgatewayconnectpeerassociation "github.com/upbound/provider-aws/internal/controller/networkmanager/transitgatewayconnectpeerassociation"
	transitgatewayregistration "github.com/upbound/provider-aws/internal/controller/networkmanager/transitgatewayregistration"
	vpcattachment "github.com/upbound/provider-aws/internal/controller/networkmanager/vpcattachment"
	domainopensearch "github.com/upbound/provider-aws/internal/controller/opensearch/domain"
	domainpolicyopensearch "github.com/upbound/provider-aws/internal/controller/opensearch/domainpolicy"
	domainsamloptionsopensearch "github.com/upbound/provider-aws/internal/controller/opensearch/domainsamloptions"
	accesspolicy "github.com/upbound/provider-aws/internal/controller/opensearchserverless/accesspolicy"
	collection "github.com/upbound/provider-aws/internal/controller/opensearchserverless/collection"
	lifecyclepolicyopensearchserverless "github.com/upbound/provider-aws/internal/controller/opensearchserverless/lifecyclepolicy"
	securityconfig "github.com/upbound/provider-aws/internal/controller/opensearchserverless/securityconfig"
	securitypolicy "github.com/upbound/provider-aws/internal/controller/opensearchserverless/securitypolicy"
	vpcendpointopensearchserverless "github.com/upbound/provider-aws/internal/controller/opensearchserverless/vpcendpoint"
	applicationopsworks "github.com/upbound/provider-aws/internal/controller/opsworks/application"
	customlayer "github.com/upbound/provider-aws/internal/controller/opsworks/customlayer"
	ecsclusterlayer "github.com/upbound/provider-aws/internal/controller/opsworks/ecsclusterlayer"
	ganglialayer "github.com/upbound/provider-aws/internal/controller/opsworks/ganglialayer"
	haproxylayer "github.com/upbound/provider-aws/internal/controller/opsworks/haproxylayer"
	instanceopsworks "github.com/upbound/provider-aws/internal/controller/opsworks/instance"
	javaapplayer "github.com/upbound/provider-aws/internal/controller/opsworks/javaapplayer"
	memcachedlayer "github.com/upbound/provider-aws/internal/controller/opsworks/memcachedlayer"
	mysqllayer "github.com/upbound/provider-aws/internal/controller/opsworks/mysqllayer"
	nodejsapplayer "github.com/upbound/provider-aws/internal/controller/opsworks/nodejsapplayer"
	permissionopsworks "github.com/upbound/provider-aws/internal/controller/opsworks/permission"
	phpapplayer "github.com/upbound/provider-aws/internal/controller/opsworks/phpapplayer"
	railsapplayer "github.com/upbound/provider-aws/internal/controller/opsworks/railsapplayer"
	rdsdbinstance "github.com/upbound/provider-aws/internal/controller/opsworks/rdsdbinstance"
	stackopsworks "github.com/upbound/provider-aws/internal/controller/opsworks/stack"
	staticweblayer "github.com/upbound/provider-aws/internal/controller/opsworks/staticweblayer"
	userprofile "github.com/upbound/provider-aws/internal/controller/opsworks/userprofile"
	accountorganizations "github.com/upbound/provider-aws/internal/controller/organizations/account"
	delegatedadministrator "github.com/upbound/provider-aws/internal/controller/organizations/delegatedadministrator"
	organization "github.com/upbound/provider-aws/internal/controller/organizations/organization"
	organizationalunit "github.com/upbound/provider-aws/internal/controller/organizations/organizationalunit"
	policyorganizations "github.com/upbound/provider-aws/internal/controller/organizations/policy"
	policyattachmentorganizations "github.com/upbound/provider-aws/internal/controller/organizations/policyattachment"
	apppinpoint "github.com/upbound/provider-aws/internal/controller/pinpoint/app"
	smschannel "github.com/upbound/provider-aws/internal/controller/pinpoint/smschannel"
	providerconfig "github.com/upbound/provider-aws/internal/controller/providerconfig"
	ledger "github.com/upbound/provider-aws/internal/controller/qldb/ledger"
	streamqldb "github.com/upbound/provider-aws/internal/controller/qldb/stream"
	groupquicksight "github.com/upbound/provider-aws/internal/controller/quicksight/group"
	userquicksight "github.com/upbound/provider-aws/internal/controller/quicksight/user"
	principalassociation "github.com/upbound/provider-aws/internal/controller/ram/principalassociation"
	resourceassociation "github.com/upbound/provider-aws/internal/controller/ram/resourceassociation"
	resourceshare "github.com/upbound/provider-aws/internal/controller/ram/resourceshare"
	resourceshareaccepter "github.com/upbound/provider-aws/internal/controller/ram/resourceshareaccepter"
	clusterrds "github.com/upbound/provider-aws/internal/controller/rds/cluster"
	clusteractivitystream "github.com/upbound/provider-aws/internal/controller/rds/clusteractivitystream"
	clusterendpointrds "github.com/upbound/provider-aws/internal/controller/rds/clusterendpoint"
	clusterinstancerds "github.com/upbound/provider-aws/internal/controller/rds/clusterinstance"
	clusterparametergrouprds "github.com/upbound/provider-aws/internal/controller/rds/clusterparametergroup"
	clusterroleassociation "github.com/upbound/provider-aws/internal/controller/rds/clusterroleassociation"
	clustersnapshotrds "github.com/upbound/provider-aws/internal/controller/rds/clustersnapshot"
	dbinstanceautomatedbackupsreplication "github.com/upbound/provider-aws/internal/controller/rds/dbinstanceautomatedbackupsreplication"
	dbsnapshotcopy "github.com/upbound/provider-aws/internal/controller/rds/dbsnapshotcopy"
	eventsubscriptionrds "github.com/upbound/provider-aws/internal/controller/rds/eventsubscription"
	globalclusterrds "github.com/upbound/provider-aws/internal/controller/rds/globalcluster"
	instancerds "github.com/upbound/provider-aws/internal/controller/rds/instance"
	instanceroleassociation "github.com/upbound/provider-aws/internal/controller/rds/instanceroleassociation"
	optiongroup "github.com/upbound/provider-aws/internal/controller/rds/optiongroup"
	parametergrouprds "github.com/upbound/provider-aws/internal/controller/rds/parametergroup"
	proxy "github.com/upbound/provider-aws/internal/controller/rds/proxy"
	proxydefaulttargetgroup "github.com/upbound/provider-aws/internal/controller/rds/proxydefaulttargetgroup"
	proxyendpoint "github.com/upbound/provider-aws/internal/controller/rds/proxyendpoint"
	proxytarget "github.com/upbound/provider-aws/internal/controller/rds/proxytarget"
	snapshotrds "github.com/upbound/provider-aws/internal/controller/rds/snapshot"
	subnetgrouprds "github.com/upbound/provider-aws/internal/controller/rds/subnetgroup"
	authenticationprofile "github.com/upbound/provider-aws/internal/controller/redshift/authenticationprofile"
	clusterredshift "github.com/upbound/provider-aws/internal/controller/redshift/cluster"
	eventsubscriptionredshift "github.com/upbound/provider-aws/internal/controller/redshift/eventsubscription"
	hsmclientcertificate "github.com/upbound/provider-aws/internal/controller/redshift/hsmclientcertificate"
	hsmconfiguration "github.com/upbound/provider-aws/internal/controller/redshift/hsmconfiguration"
	parametergroupredshift "github.com/upbound/provider-aws/internal/controller/redshift/parametergroup"
	scheduledactionredshift "github.com/upbound/provider-aws/internal/controller/redshift/scheduledaction"
	snapshotcopygrant "github.com/upbound/provider-aws/internal/controller/redshift/snapshotcopygrant"
	snapshotschedule "github.com/upbound/provider-aws/internal/controller/redshift/snapshotschedule"
	snapshotscheduleassociation "github.com/upbound/provider-aws/internal/controller/redshift/snapshotscheduleassociation"
	subnetgroupredshift "github.com/upbound/provider-aws/internal/controller/redshift/subnetgroup"
	usagelimit "github.com/upbound/provider-aws/internal/controller/redshift/usagelimit"
	endpointaccess "github.com/upbound/provider-aws/internal/controller/redshiftserverless/endpointaccess"
	redshiftserverlessnamespace "github.com/upbound/provider-aws/internal/controller/redshiftserverless/redshiftserverlessnamespace"
	resourcepolicyredshiftserverless "github.com/upbound/provider-aws/internal/controller/redshiftserverless/resourcepolicy"
	snapshotredshiftserverless "github.com/upbound/provider-aws/internal/controller/redshiftserverless/snapshot"
	usagelimitredshiftserverless "github.com/upbound/provider-aws/internal/controller/redshiftserverless/usagelimit"
	workgroupredshiftserverless "github.com/upbound/provider-aws/internal/controller/redshiftserverless/workgroup"
	groupresourcegroups "github.com/upbound/provider-aws/internal/controller/resourcegroups/group"
	profile "github.com/upbound/provider-aws/internal/controller/rolesanywhere/profile"
	delegationset "github.com/upbound/provider-aws/internal/controller/route53/delegationset"
	healthcheck "github.com/upbound/provider-aws/internal/controller/route53/healthcheck"
	hostedzonednssec "github.com/upbound/provider-aws/internal/controller/route53/hostedzonednssec"
	record "github.com/upbound/provider-aws/internal/controller/route53/record"
	resolverconfig "github.com/upbound/provider-aws/internal/controller/route53/resolverconfig"
	trafficpolicy "github.com/upbound/provider-aws/internal/controller/route53/trafficpolicy"
	trafficpolicyinstance "github.com/upbound/provider-aws/internal/controller/route53/trafficpolicyinstance"
	vpcassociationauthorization "github.com/upbound/provider-aws/internal/controller/route53/vpcassociationauthorization"
	zone "github.com/upbound/provider-aws/internal/controller/route53/zone"
	zoneassociation "github.com/upbound/provider-aws/internal/controller/route53/zoneassociation"
	clusterroute53recoverycontrolconfig "github.com/upbound/provider-aws/internal/controller/route53recoverycontrolconfig/cluster"
	controlpanel "github.com/upbound/provider-aws/internal/controller/route53recoverycontrolconfig/controlpanel"
	routingcontrol "github.com/upbound/provider-aws/internal/controller/route53recoverycontrolconfig/routingcontrol"
	safetyrule "github.com/upbound/provider-aws/internal/controller/route53recoverycontrolconfig/safetyrule"
	cell "github.com/upbound/provider-aws/internal/controller/route53recoveryreadiness/cell"
	readinesscheck "github.com/upbound/provider-aws/internal/controller/route53recoveryreadiness/readinesscheck"
	recoverygroup "github.com/upbound/provider-aws/internal/controller/route53recoveryreadiness/recoverygroup"
	resourceset "github.com/upbound/provider-aws/internal/controller/route53recoveryreadiness/resourceset"
	endpointroute53resolver "github.com/upbound/provider-aws/internal/controller/route53resolver/endpoint"
	ruleroute53resolver "github.com/upbound/provider-aws/internal/controller/route53resolver/rule"
	ruleassociation "github.com/upbound/provider-aws/internal/controller/route53resolver/ruleassociation"
	appmonitor "github.com/upbound/provider-aws/internal/controller/rum/appmonitor"
	metricsdestination "github.com/upbound/provider-aws/internal/controller/rum/metricsdestination"
	buckets3 "github.com/upbound/provider-aws/internal/controller/s3/bucket"
	bucketaccelerateconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketaccelerateconfiguration"
	bucketacl "github.com/upbound/provider-aws/internal/controller/s3/bucketacl"
	bucketanalyticsconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketanalyticsconfiguration"
	bucketcorsconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketcorsconfiguration"
	bucketintelligenttieringconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketintelligenttieringconfiguration"
	bucketinventory "github.com/upbound/provider-aws/internal/controller/s3/bucketinventory"
	bucketlifecycleconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketlifecycleconfiguration"
	bucketlogging "github.com/upbound/provider-aws/internal/controller/s3/bucketlogging"
	bucketmetric "github.com/upbound/provider-aws/internal/controller/s3/bucketmetric"
	bucketnotification "github.com/upbound/provider-aws/internal/controller/s3/bucketnotification"
	bucketobject "github.com/upbound/provider-aws/internal/controller/s3/bucketobject"
	bucketobjectlockconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketobjectlockconfiguration"
	bucketownershipcontrols "github.com/upbound/provider-aws/internal/controller/s3/bucketownershipcontrols"
	bucketpolicy "github.com/upbound/provider-aws/internal/controller/s3/bucketpolicy"
	bucketpublicaccessblock "github.com/upbound/provider-aws/internal/controller/s3/bucketpublicaccessblock"
	bucketreplicationconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketreplicationconfiguration"
	bucketrequestpaymentconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketrequestpaymentconfiguration"
	bucketserversideencryptionconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketserversideencryptionconfiguration"
	bucketversioning "github.com/upbound/provider-aws/internal/controller/s3/bucketversioning"
	bucketwebsiteconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketwebsiteconfiguration"
	object "github.com/upbound/provider-aws/internal/controller/s3/object"
	objectcopy "github.com/upbound/provider-aws/internal/controller/s3/objectcopy"
	accesspoints3control "github.com/upbound/provider-aws/internal/controller/s3control/accesspoint"
	accesspointpolicy "github.com/upbound/provider-aws/internal/controller/s3control/accesspointpolicy"
	accountpublicaccessblock "github.com/upbound/provider-aws/internal/controller/s3control/accountpublicaccessblock"
	multiregionaccesspoint "github.com/upbound/provider-aws/internal/controller/s3control/multiregionaccesspoint"
	multiregionaccesspointpolicy "github.com/upbound/provider-aws/internal/controller/s3control/multiregionaccesspointpolicy"
	objectlambdaaccesspoint "github.com/upbound/provider-aws/internal/controller/s3control/objectlambdaaccesspoint"
	objectlambdaaccesspointpolicy "github.com/upbound/provider-aws/internal/controller/s3control/objectlambdaaccesspointpolicy"
	storagelensconfiguration "github.com/upbound/provider-aws/internal/controller/s3control/storagelensconfiguration"
	appsagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/app"
	appimageconfig "github.com/upbound/provider-aws/internal/controller/sagemaker/appimageconfig"
	coderepository "github.com/upbound/provider-aws/internal/controller/sagemaker/coderepository"
	devicesagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/device"
	devicefleet "github.com/upbound/provider-aws/internal/controller/sagemaker/devicefleet"
	domainsagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/domain"
	endpointsagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/endpoint"
	endpointconfiguration "github.com/upbound/provider-aws/internal/controller/sagemaker/endpointconfiguration"
	featuregroup "github.com/upbound/provider-aws/internal/controller/sagemaker/featuregroup"
	imagesagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/image"
	imageversion "github.com/upbound/provider-aws/internal/controller/sagemaker/imageversion"
	modelsagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/model"
	modelpackagegroup "github.com/upbound/provider-aws/internal/controller/sagemaker/modelpackagegroup"
	modelpackagegrouppolicy "github.com/upbound/provider-aws/internal/controller/sagemaker/modelpackagegrouppolicy"
	notebookinstance "github.com/upbound/provider-aws/internal/controller/sagemaker/notebookinstance"
	notebookinstancelifecycleconfiguration "github.com/upbound/provider-aws/internal/controller/sagemaker/notebookinstancelifecycleconfiguration"
	servicecatalogportfoliostatus "github.com/upbound/provider-aws/internal/controller/sagemaker/servicecatalogportfoliostatus"
	space "github.com/upbound/provider-aws/internal/controller/sagemaker/space"
	studiolifecycleconfig "github.com/upbound/provider-aws/internal/controller/sagemaker/studiolifecycleconfig"
	userprofilesagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/userprofile"
	workforce "github.com/upbound/provider-aws/internal/controller/sagemaker/workforce"
	workteam "github.com/upbound/provider-aws/internal/controller/sagemaker/workteam"
	schedulescheduler "github.com/upbound/provider-aws/internal/controller/scheduler/schedule"
	schedulegroup "github.com/upbound/provider-aws/internal/controller/scheduler/schedulegroup"
	discoverer "github.com/upbound/provider-aws/internal/controller/schemas/discoverer"
	registryschemas "github.com/upbound/provider-aws/internal/controller/schemas/registry"
	schemaschemas "github.com/upbound/provider-aws/internal/controller/schemas/schema"
	secret "github.com/upbound/provider-aws/internal/controller/secretsmanager/secret"
	secretpolicy "github.com/upbound/provider-aws/internal/controller/secretsmanager/secretpolicy"
	secretrotation "github.com/upbound/provider-aws/internal/controller/secretsmanager/secretrotation"
	secretversion "github.com/upbound/provider-aws/internal/controller/secretsmanager/secretversion"
	accountsecurityhub "github.com/upbound/provider-aws/internal/controller/securityhub/account"
	actiontarget "github.com/upbound/provider-aws/internal/controller/securityhub/actiontarget"
	findingaggregator "github.com/upbound/provider-aws/internal/controller/securityhub/findingaggregator"
	insight "github.com/upbound/provider-aws/internal/controller/securityhub/insight"
	inviteaccepter "github.com/upbound/provider-aws/internal/controller/securityhub/inviteaccepter"
	membersecurityhub "github.com/upbound/provider-aws/internal/controller/securityhub/member"
	productsubscription "github.com/upbound/provider-aws/internal/controller/securityhub/productsubscription"
	standardssubscription "github.com/upbound/provider-aws/internal/controller/securityhub/standardssubscription"
	cloudformationstack "github.com/upbound/provider-aws/internal/controller/serverlessrepo/cloudformationstack"
	budgetresourceassociation "github.com/upbound/provider-aws/internal/controller/servicecatalog/budgetresourceassociation"
	constraint "github.com/upbound/provider-aws/internal/controller/servicecatalog/constraint"
	portfolio "github.com/upbound/provider-aws/internal/controller/servicecatalog/portfolio"
	portfolioshare "github.com/upbound/provider-aws/internal/controller/servicecatalog/portfolioshare"
	principalportfolioassociation "github.com/upbound/provider-aws/internal/controller/servicecatalog/principalportfolioassociation"
	product "github.com/upbound/provider-aws/internal/controller/servicecatalog/product"
	productportfolioassociation "github.com/upbound/provider-aws/internal/controller/servicecatalog/productportfolioassociation"
	provisioningartifact "github.com/upbound/provider-aws/internal/controller/servicecatalog/provisioningartifact"
	serviceaction "github.com/upbound/provider-aws/internal/controller/servicecatalog/serviceaction"
	tagoption "github.com/upbound/provider-aws/internal/controller/servicecatalog/tagoption"
	tagoptionresourceassociation "github.com/upbound/provider-aws/internal/controller/servicecatalog/tagoptionresourceassociation"
	httpnamespace "github.com/upbound/provider-aws/internal/controller/servicediscovery/httpnamespace"
	privatednsnamespace "github.com/upbound/provider-aws/internal/controller/servicediscovery/privatednsnamespace"
	publicdnsnamespace "github.com/upbound/provider-aws/internal/controller/servicediscovery/publicdnsnamespace"
	serviceservicediscovery "github.com/upbound/provider-aws/internal/controller/servicediscovery/service"
	servicequota "github.com/upbound/provider-aws/internal/controller/servicequotas/servicequota"
	activereceiptruleset "github.com/upbound/provider-aws/internal/controller/ses/activereceiptruleset"
	configurationset "github.com/upbound/provider-aws/internal/controller/ses/configurationset"
	domaindkim "github.com/upbound/provider-aws/internal/controller/ses/domaindkim"
	domainidentity "github.com/upbound/provider-aws/internal/controller/ses/domainidentity"
	domainmailfrom "github.com/upbound/provider-aws/internal/controller/ses/domainmailfrom"
	emailidentity "github.com/upbound/provider-aws/internal/controller/ses/emailidentity"
	eventdestination "github.com/upbound/provider-aws/internal/controller/ses/eventdestination"
	identitynotificationtopic "github.com/upbound/provider-aws/internal/controller/ses/identitynotificationtopic"
	identitypolicy "github.com/upbound/provider-aws/internal/controller/ses/identitypolicy"
	receiptfilter "github.com/upbound/provider-aws/internal/controller/ses/receiptfilter"
	receiptrule "github.com/upbound/provider-aws/internal/controller/ses/receiptrule"
	receiptruleset "github.com/upbound/provider-aws/internal/controller/ses/receiptruleset"
	template "github.com/upbound/provider-aws/internal/controller/ses/template"
	configurationsetsesv2 "github.com/upbound/provider-aws/internal/controller/sesv2/configurationset"
	configurationseteventdestination "github.com/upbound/provider-aws/internal/controller/sesv2/configurationseteventdestination"
	dedicatedippool "github.com/upbound/provider-aws/internal/controller/sesv2/dedicatedippool"
	emailidentitysesv2 "github.com/upbound/provider-aws/internal/controller/sesv2/emailidentity"
	emailidentityfeedbackattributes "github.com/upbound/provider-aws/internal/controller/sesv2/emailidentityfeedbackattributes"
	emailidentitymailfromattributes "github.com/upbound/provider-aws/internal/controller/sesv2/emailidentitymailfromattributes"
	activity "github.com/upbound/provider-aws/internal/controller/sfn/activity"
	statemachine "github.com/upbound/provider-aws/internal/controller/sfn/statemachine"
	signingjob "github.com/upbound/provider-aws/internal/controller/signer/signingjob"
	signingprofile "github.com/upbound/provider-aws/internal/controller/signer/signingprofile"
	signingprofilepermission "github.com/upbound/provider-aws/internal/controller/signer/signingprofilepermission"
	domainsimpledb "github.com/upbound/provider-aws/internal/controller/simpledb/domain"
	platformapplication "github.com/upbound/provider-aws/internal/controller/sns/platformapplication"
	smspreferences "github.com/upbound/provider-aws/internal/controller/sns/smspreferences"
	topic "github.com/upbound/provider-aws/internal/controller/sns/topic"
	topicpolicy "github.com/upbound/provider-aws/internal/controller/sns/topicpolicy"
	topicsubscription "github.com/upbound/provider-aws/internal/controller/sns/topicsubscription"
	queuesqs "github.com/upbound/provider-aws/internal/controller/sqs/queue"
	queuepolicy "github.com/upbound/provider-aws/internal/controller/sqs/queuepolicy"
	queueredriveallowpolicy "github.com/upbound/provider-aws/internal/controller/sqs/queueredriveallowpolicy"
	queueredrivepolicy "github.com/upbound/provider-aws/internal/controller/sqs/queueredrivepolicy"
	activation "github.com/upbound/provider-aws/internal/controller/ssm/activation"
	associationssm "github.com/upbound/provider-aws/internal/controller/ssm/association"
	defaultpatchbaseline "github.com/upbound/provider-aws/internal/controller/ssm/defaultpatchbaseline"
	document "github.com/upbound/provider-aws/internal/controller/ssm/document"
	maintenancewindow "github.com/upbound/provider-aws/internal/controller/ssm/maintenancewindow"
	maintenancewindowtarget "github.com/upbound/provider-aws/internal/controller/ssm/maintenancewindowtarget"
	maintenancewindowtask "github.com/upbound/provider-aws/internal/controller/ssm/maintenancewindowtask"
	parameter "github.com/upbound/provider-aws/internal/controller/ssm/parameter"
	patchbaseline "github.com/upbound/provider-aws/internal/controller/ssm/patchbaseline"
	patchgroup "github.com/upbound/provider-aws/internal/controller/ssm/patchgroup"
	resourcedatasync "github.com/upbound/provider-aws/internal/controller/ssm/resourcedatasync"
	servicesetting "github.com/upbound/provider-aws/internal/controller/ssm/servicesetting"
	accountassignment "github.com/upbound/provider-aws/internal/controller/ssoadmin/accountassignment"
	customermanagedpolicyattachment "github.com/upbound/provider-aws/internal/controller/ssoadmin/customermanagedpolicyattachment"
	instanceaccesscontrolattributes "github.com/upbound/provider-aws/internal/controller/ssoadmin/instanceaccesscontrolattributes"
	managedpolicyattachment "github.com/upbound/provider-aws/internal/controller/ssoadmin/managedpolicyattachment"
	permissionsboundaryattachment "github.com/upbound/provider-aws/internal/controller/ssoadmin/permissionsboundaryattachment"
	permissionset "github.com/upbound/provider-aws/internal/controller/ssoadmin/permissionset"
	permissionsetinlinepolicy "github.com/upbound/provider-aws/internal/controller/ssoadmin/permissionsetinlinepolicy"
	domainswf "github.com/upbound/provider-aws/internal/controller/swf/domain"
	databasetimestreamwrite "github.com/upbound/provider-aws/internal/controller/timestreamwrite/database"
	tabletimestreamwrite "github.com/upbound/provider-aws/internal/controller/timestreamwrite/table"
	languagemodel "github.com/upbound/provider-aws/internal/controller/transcribe/languagemodel"
	vocabularytranscribe "github.com/upbound/provider-aws/internal/controller/transcribe/vocabulary"
	vocabularyfilter "github.com/upbound/provider-aws/internal/controller/transcribe/vocabularyfilter"
	server "github.com/upbound/provider-aws/internal/controller/transfer/server"
	sshkey "github.com/upbound/provider-aws/internal/controller/transfer/sshkey"
	tagtransfer "github.com/upbound/provider-aws/internal/controller/transfer/tag"
	usertransfer "github.com/upbound/provider-aws/internal/controller/transfer/user"
	workflowtransfer "github.com/upbound/provider-aws/internal/controller/transfer/workflow"
	networkperformancemetricsubscription "github.com/upbound/provider-aws/internal/controller/vpc/networkperformancemetricsubscription"
	bytematchset "github.com/upbound/provider-aws/internal/controller/waf/bytematchset"
	geomatchset "github.com/upbound/provider-aws/internal/controller/waf/geomatchset"
	ipset "github.com/upbound/provider-aws/internal/controller/waf/ipset"
	ratebasedrule "github.com/upbound/provider-aws/internal/controller/waf/ratebasedrule"
	regexmatchset "github.com/upbound/provider-aws/internal/controller/waf/regexmatchset"
	regexpatternset "github.com/upbound/provider-aws/internal/controller/waf/regexpatternset"
	rulewaf "github.com/upbound/provider-aws/internal/controller/waf/rule"
	sizeconstraintset "github.com/upbound/provider-aws/internal/controller/waf/sizeconstraintset"
	sqlinjectionmatchset "github.com/upbound/provider-aws/internal/controller/waf/sqlinjectionmatchset"
	webacl "github.com/upbound/provider-aws/internal/controller/waf/webacl"
	xssmatchset "github.com/upbound/provider-aws/internal/controller/waf/xssmatchset"
	bytematchsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/bytematchset"
	geomatchsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/geomatchset"
	ipsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/ipset"
	ratebasedrulewafregional "github.com/upbound/provider-aws/internal/controller/wafregional/ratebasedrule"
	regexmatchsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/regexmatchset"
	regexpatternsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/regexpatternset"
	rulewafregional "github.com/upbound/provider-aws/internal/controller/wafregional/rule"
	sizeconstraintsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/sizeconstraintset"
	sqlinjectionmatchsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/sqlinjectionmatchset"
	webaclwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/webacl"
	xssmatchsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/xssmatchset"
	ipsetwafv2 "github.com/upbound/provider-aws/internal/controller/wafv2/ipset"
	regexpatternsetwafv2 "github.com/upbound/provider-aws/internal/controller/wafv2/regexpatternset"
	directoryworkspaces "github.com/upbound/provider-aws/internal/controller/workspaces/directory"
	ipgroup "github.com/upbound/provider-aws/internal/controller/workspaces/ipgroup"
	encryptionconfig "github.com/upbound/provider-aws/internal/controller/xray/encryptionconfig"
	groupxray "github.com/upbound/provider-aws/internal/controller/xray/group"
	samplingrule "github.com/upbound/provider-aws/internal/controller/xray/samplingrule"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyzer.Setup,
		archiverule.Setup,
		alternatecontact.Setup,
		certificate.Setup,
		certificatevalidation.Setup,
		certificateacmpca.Setup,
		certificateauthority.Setup,
		certificateauthoritycertificate.Setup,
		permission.Setup,
		policy.Setup,
		alertmanagerdefinition.Setup,
		rulegroupnamespace.Setup,
		workspace.Setup,
		app.Setup,
		backendenvironment.Setup,
		branch.Setup,
		webhook.Setup,
		account.Setup,
		apikey.Setup,
		authorizer.Setup,
		basepathmapping.Setup,
		clientcertificate.Setup,
		deployment.Setup,
		documentationpart.Setup,
		documentationversion.Setup,
		domainname.Setup,
		gatewayresponse.Setup,
		integration.Setup,
		integrationresponse.Setup,
		method.Setup,
		methodresponse.Setup,
		methodsettings.Setup,
		model.Setup,
		requestvalidator.Setup,
		resource.Setup,
		restapi.Setup,
		restapipolicy.Setup,
		stage.Setup,
		usageplan.Setup,
		usageplankey.Setup,
		vpclink.Setup,
		api.Setup,
		apimapping.Setup,
		authorizerapigatewayv2.Setup,
		deploymentapigatewayv2.Setup,
		domainnameapigatewayv2.Setup,
		integrationapigatewayv2.Setup,
		integrationresponseapigatewayv2.Setup,
		modelapigatewayv2.Setup,
		route.Setup,
		routeresponse.Setup,
		stageapigatewayv2.Setup,
		vpclinkapigatewayv2.Setup,
		policyappautoscaling.Setup,
		scheduledaction.Setup,
		target.Setup,
		application.Setup,
		configurationprofile.Setup,
		deploymentappconfig.Setup,
		deploymentstrategy.Setup,
		environment.Setup,
		extension.Setup,
		extensionassociation.Setup,
		hostedconfigurationversion.Setup,
		flow.Setup,
		eventintegration.Setup,
		applicationapplicationinsights.Setup,
		gatewayroute.Setup,
		mesh.Setup,
		routeappmesh.Setup,
		virtualgateway.Setup,
		virtualnode.Setup,
		virtualrouter.Setup,
		virtualservice.Setup,
		autoscalingconfigurationversion.Setup,
		connection.Setup,
		observabilityconfiguration.Setup,
		service.Setup,
		vpcconnector.Setup,
		directoryconfig.Setup,
		fleet.Setup,
		fleetstackassociation.Setup,
		imagebuilder.Setup,
		stack.Setup,
		user.Setup,
		userstackassociation.Setup,
		apicache.Setup,
		apikeyappsync.Setup,
		datasource.Setup,
		function.Setup,
		graphqlapi.Setup,
		resolver.Setup,
		database.Setup,
		datacatalog.Setup,
		namedquery.Setup,
		workgroup.Setup,
		attachment.Setup,
		autoscalinggroup.Setup,
		grouptag.Setup,
		launchconfiguration.Setup,
		lifecyclehook.Setup,
		notification.Setup,
		policyautoscaling.Setup,
		schedule.Setup,
		scalingplan.Setup,
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
		jobdefinition.Setup,
		schedulingpolicy.Setup,
		budget.Setup,
		budgetaction.Setup,
		anomalymonitor.Setup,
		voiceconnector.Setup,
		voiceconnectorgroup.Setup,
		voiceconnectorlogging.Setup,
		voiceconnectororigination.Setup,
		voiceconnectorstreaming.Setup,
		voiceconnectortermination.Setup,
		voiceconnectorterminationcredentials.Setup,
		environmentec2.Setup,
		environmentmembership.Setup,
		resourcecloudcontrol.Setup,
		stackcloudformation.Setup,
		stackset.Setup,
		stacksetinstance.Setup,
		cachepolicy.Setup,
		distribution.Setup,
		fieldlevelencryptionconfig.Setup,
		fieldlevelencryptionprofile.Setup,
		functioncloudfront.Setup,
		keygroup.Setup,
		monitoringsubscription.Setup,
		originaccesscontrol.Setup,
		originaccessidentity.Setup,
		originrequestpolicy.Setup,
		publickey.Setup,
		realtimelogconfig.Setup,
		responseheaderspolicy.Setup,
		domain.Setup,
		domainserviceaccesspolicy.Setup,
		eventdatastore.Setup,
		trail.Setup,
		compositealarm.Setup,
		dashboard.Setup,
		metricalarm.Setup,
		metricstream.Setup,
		apidestination.Setup,
		archive.Setup,
		bus.Setup,
		buspolicy.Setup,
		connectioncloudwatchevents.Setup,
		permissioncloudwatchevents.Setup,
		rule.Setup,
		targetcloudwatchevents.Setup,
		definition.Setup,
		destination.Setup,
		destinationpolicy.Setup,
		group.Setup,
		metricfilter.Setup,
		resourcepolicy.Setup,
		stream.Setup,
		subscriptionfilter.Setup,
		approvalruletemplate.Setup,
		approvalruletemplateassociation.Setup,
		repository.Setup,
		trigger.Setup,
		codepipeline.Setup,
		customactiontype.Setup,
		webhookcodepipeline.Setup,
		connectioncodestarconnections.Setup,
		host.Setup,
		notificationrule.Setup,
		cognitoidentitypoolproviderprincipaltag.Setup,
		pool.Setup,
		poolrolesattachment.Setup,
		identityprovider.Setup,
		resourceserver.Setup,
		riskconfiguration.Setup,
		usercognitoidp.Setup,
		usergroup.Setup,
		useringroup.Setup,
		userpool.Setup,
		userpoolclient.Setup,
		userpooldomain.Setup,
		userpooluicustomization.Setup,
		awsconfigurationrecorderstatus.Setup,
		configrule.Setup,
		configurationaggregator.Setup,
		configurationrecorder.Setup,
		conformancepack.Setup,
		deliverychannel.Setup,
		remediationconfiguration.Setup,
		botassociation.Setup,
		contactflow.Setup,
		contactflowmodule.Setup,
		hoursofoperation.Setup,
		instance.Setup,
		instancestorageconfig.Setup,
		lambdafunctionassociation.Setup,
		phonenumber.Setup,
		queue.Setup,
		quickconnect.Setup,
		routingprofile.Setup,
		securityprofile.Setup,
		userconnect.Setup,
		userhierarchystructure.Setup,
		vocabulary.Setup,
		reportdefinition.Setup,
		dataset.Setup,
		revision.Setup,
		pipeline.Setup,
		locations3.Setup,
		task.Setup,
		cluster.Setup,
		parametergroup.Setup,
		subnetgroup.Setup,
		appdeploy.Setup,
		deploymentconfig.Setup,
		deploymentgroup.Setup,
		graph.Setup,
		invitationaccepter.Setup,
		member.Setup,
		devicepool.Setup,
		instanceprofile.Setup,
		networkprofile.Setup,
		project.Setup,
		testgridproject.Setup,
		upload.Setup,
		bgppeer.Setup,
		connectiondirectconnect.Setup,
		connectionassociation.Setup,
		gateway.Setup,
		gatewayassociation.Setup,
		gatewayassociationproposal.Setup,
		hostedprivatevirtualinterface.Setup,
		hostedprivatevirtualinterfaceaccepter.Setup,
		hostedpublicvirtualinterface.Setup,
		hostedpublicvirtualinterfaceaccepter.Setup,
		hostedtransitvirtualinterface.Setup,
		hostedtransitvirtualinterfaceaccepter.Setup,
		lag.Setup,
		privatevirtualinterface.Setup,
		publicvirtualinterface.Setup,
		transitvirtualinterface.Setup,
		lifecyclepolicy.Setup,
		certificatedms.Setup,
		endpoint.Setup,
		eventsubscription.Setup,
		replicationinstance.Setup,
		replicationsubnetgroup.Setup,
		replicationtask.Setup,
		s3endpoint.Setup,
		clusterdocdb.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clustersnapshot.Setup,
		eventsubscriptiondocdb.Setup,
		globalcluster.Setup,
		subnetgroupdocdb.Setup,
		conditionalforwarder.Setup,
		directory.Setup,
		shareddirectory.Setup,
		contributorinsights.Setup,
		globaltable.Setup,
		kinesisstreamingdestination.Setup,
		table.Setup,
		tableitem.Setup,
		tablereplica.Setup,
		tag.Setup,
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
		hostec2.Setup,
		instanceec2.Setup,
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
		routeec2.Setup,
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
		tagec2.Setup,
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
		lifecyclepolicyecr.Setup,
		pullthroughcacherule.Setup,
		registrypolicy.Setup,
		registryscanningconfiguration.Setup,
		replicationconfiguration.Setup,
		repositoryecr.Setup,
		repositorypolicy.Setup,
		repositoryecrpublic.Setup,
		repositorypolicyecrpublic.Setup,
		accountsettingdefault.Setup,
		capacityprovider.Setup,
		clusterecs.Setup,
		clustercapacityproviders.Setup,
		serviceecs.Setup,
		taskdefinition.Setup,
		accesspoint.Setup,
		backuppolicy.Setup,
		filesystem.Setup,
		filesystempolicy.Setup,
		mounttarget.Setup,
		replicationconfigurationefs.Setup,
		addon.Setup,
		clustereks.Setup,
		clusterauth.Setup,
		fargateprofile.Setup,
		identityproviderconfig.Setup,
		nodegroup.Setup,
		podidentityassociation.Setup,
		clusterelasticache.Setup,
		parametergroupelasticache.Setup,
		replicationgroup.Setup,
		subnetgroupelasticache.Setup,
		userelasticache.Setup,
		usergroupelasticache.Setup,
		applicationelasticbeanstalk.Setup,
		applicationversion.Setup,
		configurationtemplate.Setup,
		domainelasticsearch.Setup,
		domainpolicy.Setup,
		domainsamloptions.Setup,
		pipelineelastictranscoder.Setup,
		preset.Setup,
		appcookiestickinesspolicy.Setup,
		attachmentelb.Setup,
		backendserverpolicy.Setup,
		elb.Setup,
		lbcookiestickinesspolicy.Setup,
		lbsslnegotiationpolicy.Setup,
		listenerpolicy.Setup,
		policyelb.Setup,
		proxyprotocolpolicy.Setup,
		lb.Setup,
		lblistener.Setup,
		lblistenercertificate.Setup,
		lblistenerrule.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		securityconfiguration.Setup,
		applicationemrserverless.Setup,
		feature.Setup,
		projectevidently.Setup,
		segment.Setup,
		deliverystream.Setup,
		experimenttemplate.Setup,
		backup.Setup,
		datarepositoryassociation.Setup,
		lustrefilesystem.Setup,
		ontapfilesystem.Setup,
		ontapstoragevirtualmachine.Setup,
		windowsfilesystem.Setup,
		alias.Setup,
		build.Setup,
		fleetgamelift.Setup,
		gamesessionqueue.Setup,
		script.Setup,
		vaultglacier.Setup,
		vaultlock.Setup,
		accelerator.Setup,
		endpointgroup.Setup,
		listener.Setup,
		catalogdatabase.Setup,
		catalogtable.Setup,
		classifier.Setup,
		connectionglue.Setup,
		crawler.Setup,
		datacatalogencryptionsettings.Setup,
		job.Setup,
		registry.Setup,
		resourcepolicyglue.Setup,
		schema.Setup,
		securityconfigurationglue.Setup,
		triggerglue.Setup,
		userdefinedfunction.Setup,
		workflow.Setup,
		licenseassociation.Setup,
		roleassociation.Setup,
		workspacegrafana.Setup,
		workspaceapikey.Setup,
		workspacesamlconfiguration.Setup,
		detector.Setup,
		filter.Setup,
		memberguardduty.Setup,
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		groupiam.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		instanceprofileiam.Setup,
		openidconnectprovider.Setup,
		policyiam.Setup,
		role.Setup,
		rolepolicy.Setup,
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
		groupidentitystore.Setup,
		groupmembershipidentitystore.Setup,
		useridentitystore.Setup,
		component.Setup,
		containerrecipe.Setup,
		distributionconfiguration.Setup,
		image.Setup,
		imagepipeline.Setup,
		imagerecipe.Setup,
		infrastructureconfiguration.Setup,
		assessmenttarget.Setup,
		assessmenttemplate.Setup,
		resourcegroup.Setup,
		enabler.Setup,
		certificateiot.Setup,
		indexingconfiguration.Setup,
		loggingoptions.Setup,
		policyiot.Setup,
		policyattachment.Setup,
		provisioningtemplate.Setup,
		rolealias.Setup,
		thing.Setup,
		thinggroup.Setup,
		thinggroupmembership.Setup,
		thingprincipalattachment.Setup,
		thingtype.Setup,
		topicrule.Setup,
		topicruledestination.Setup,
		channel.Setup,
		recordingconfiguration.Setup,
		clusterkafka.Setup,
		configuration.Setup,
		scramsecretassociation.Setup,
		serverlesscluster.Setup,
		connector.Setup,
		customplugin.Setup,
		workerconfiguration.Setup,
		datasourcekendra.Setup,
		experience.Setup,
		index.Setup,
		querysuggestionsblocklist.Setup,
		thesaurus.Setup,
		keyspace.Setup,
		tablekeyspaces.Setup,
		streamkinesis.Setup,
		streamconsumer.Setup,
		applicationkinesisanalytics.Setup,
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
		resourcelakeformation.Setup,
		aliaslambda.Setup,
		codesigningconfig.Setup,
		eventsourcemapping.Setup,
		functionlambda.Setup,
		functioneventinvokeconfig.Setup,
		functionurl.Setup,
		invocation.Setup,
		layerversion.Setup,
		layerversionpermission.Setup,
		permissionlambda.Setup,
		provisionedconcurrencyconfig.Setup,
		bot.Setup,
		botalias.Setup,
		intent.Setup,
		slottype.Setup,
		association.Setup,
		licenseconfiguration.Setup,
		bucket.Setup,
		certificatelightsail.Setup,
		containerservice.Setup,
		disk.Setup,
		diskattachment.Setup,
		domainlightsail.Setup,
		domainentry.Setup,
		instancelightsail.Setup,
		instancepublicports.Setup,
		keypairlightsail.Setup,
		lblightsail.Setup,
		lbattachment.Setup,
		lbcertificate.Setup,
		lbstickinesspolicy.Setup,
		staticip.Setup,
		staticipattachment.Setup,
		geofencecollection.Setup,
		placeindex.Setup,
		routecalculator.Setup,
		tracker.Setup,
		trackerassociation.Setup,
		accountmacie2.Setup,
		classificationjob.Setup,
		customdataidentifier.Setup,
		findingsfilter.Setup,
		invitationacceptermacie2.Setup,
		membermacie2.Setup,
		queuemediaconvert.Setup,
		channelmedialive.Setup,
		input.Setup,
		inputsecuritygroup.Setup,
		multiplex.Setup,
		channelmediapackage.Setup,
		container.Setup,
		containerpolicy.Setup,
		acl.Setup,
		clustermemorydb.Setup,
		parametergroupmemorydb.Setup,
		snapshot.Setup,
		subnetgroupmemorydb.Setup,
		usermemorydb.Setup,
		broker.Setup,
		configurationmq.Setup,
		clusterneptune.Setup,
		clusterendpoint.Setup,
		clusterinstanceneptune.Setup,
		clusterparametergroupneptune.Setup,
		clustersnapshotneptune.Setup,
		eventsubscriptionneptune.Setup,
		globalclusterneptune.Setup,
		parametergroupneptune.Setup,
		subnetgroupneptune.Setup,
		firewall.Setup,
		firewallpolicy.Setup,
		loggingconfiguration.Setup,
		rulegroup.Setup,
		attachmentaccepter.Setup,
		connectattachment.Setup,
		connectionnetworkmanager.Setup,
		corenetwork.Setup,
		customergatewayassociation.Setup,
		device.Setup,
		globalnetwork.Setup,
		link.Setup,
		linkassociation.Setup,
		site.Setup,
		transitgatewayconnectpeerassociation.Setup,
		transitgatewayregistration.Setup,
		vpcattachment.Setup,
		domainopensearch.Setup,
		domainpolicyopensearch.Setup,
		domainsamloptionsopensearch.Setup,
		accesspolicy.Setup,
		collection.Setup,
		lifecyclepolicyopensearchserverless.Setup,
		securityconfig.Setup,
		securitypolicy.Setup,
		vpcendpointopensearchserverless.Setup,
		applicationopsworks.Setup,
		customlayer.Setup,
		ecsclusterlayer.Setup,
		ganglialayer.Setup,
		haproxylayer.Setup,
		instanceopsworks.Setup,
		javaapplayer.Setup,
		memcachedlayer.Setup,
		mysqllayer.Setup,
		nodejsapplayer.Setup,
		permissionopsworks.Setup,
		phpapplayer.Setup,
		railsapplayer.Setup,
		rdsdbinstance.Setup,
		stackopsworks.Setup,
		staticweblayer.Setup,
		userprofile.Setup,
		accountorganizations.Setup,
		delegatedadministrator.Setup,
		organization.Setup,
		organizationalunit.Setup,
		policyorganizations.Setup,
		policyattachmentorganizations.Setup,
		apppinpoint.Setup,
		smschannel.Setup,
		providerconfig.Setup,
		ledger.Setup,
		streamqldb.Setup,
		groupquicksight.Setup,
		userquicksight.Setup,
		principalassociation.Setup,
		resourceassociation.Setup,
		resourceshare.Setup,
		resourceshareaccepter.Setup,
		clusterrds.Setup,
		clusteractivitystream.Setup,
		clusterendpointrds.Setup,
		clusterinstancerds.Setup,
		clusterparametergrouprds.Setup,
		clusterroleassociation.Setup,
		clustersnapshotrds.Setup,
		dbinstanceautomatedbackupsreplication.Setup,
		dbsnapshotcopy.Setup,
		eventsubscriptionrds.Setup,
		globalclusterrds.Setup,
		instancerds.Setup,
		instanceroleassociation.Setup,
		optiongroup.Setup,
		parametergrouprds.Setup,
		proxy.Setup,
		proxydefaulttargetgroup.Setup,
		proxyendpoint.Setup,
		proxytarget.Setup,
		snapshotrds.Setup,
		subnetgrouprds.Setup,
		authenticationprofile.Setup,
		clusterredshift.Setup,
		eventsubscriptionredshift.Setup,
		hsmclientcertificate.Setup,
		hsmconfiguration.Setup,
		parametergroupredshift.Setup,
		scheduledactionredshift.Setup,
		snapshotcopygrant.Setup,
		snapshotschedule.Setup,
		snapshotscheduleassociation.Setup,
		subnetgroupredshift.Setup,
		usagelimit.Setup,
		endpointaccess.Setup,
		redshiftserverlessnamespace.Setup,
		resourcepolicyredshiftserverless.Setup,
		snapshotredshiftserverless.Setup,
		usagelimitredshiftserverless.Setup,
		workgroupredshiftserverless.Setup,
		groupresourcegroups.Setup,
		profile.Setup,
		delegationset.Setup,
		healthcheck.Setup,
		hostedzonednssec.Setup,
		record.Setup,
		resolverconfig.Setup,
		trafficpolicy.Setup,
		trafficpolicyinstance.Setup,
		vpcassociationauthorization.Setup,
		zone.Setup,
		zoneassociation.Setup,
		clusterroute53recoverycontrolconfig.Setup,
		controlpanel.Setup,
		routingcontrol.Setup,
		safetyrule.Setup,
		cell.Setup,
		readinesscheck.Setup,
		recoverygroup.Setup,
		resourceset.Setup,
		endpointroute53resolver.Setup,
		ruleroute53resolver.Setup,
		ruleassociation.Setup,
		appmonitor.Setup,
		metricsdestination.Setup,
		buckets3.Setup,
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
		objectcopy.Setup,
		accesspoints3control.Setup,
		accesspointpolicy.Setup,
		accountpublicaccessblock.Setup,
		multiregionaccesspoint.Setup,
		multiregionaccesspointpolicy.Setup,
		objectlambdaaccesspoint.Setup,
		objectlambdaaccesspointpolicy.Setup,
		storagelensconfiguration.Setup,
		appsagemaker.Setup,
		appimageconfig.Setup,
		coderepository.Setup,
		devicesagemaker.Setup,
		devicefleet.Setup,
		domainsagemaker.Setup,
		endpointsagemaker.Setup,
		endpointconfiguration.Setup,
		featuregroup.Setup,
		imagesagemaker.Setup,
		imageversion.Setup,
		modelsagemaker.Setup,
		modelpackagegroup.Setup,
		modelpackagegrouppolicy.Setup,
		notebookinstance.Setup,
		notebookinstancelifecycleconfiguration.Setup,
		servicecatalogportfoliostatus.Setup,
		space.Setup,
		studiolifecycleconfig.Setup,
		userprofilesagemaker.Setup,
		workforce.Setup,
		workteam.Setup,
		schedulescheduler.Setup,
		schedulegroup.Setup,
		discoverer.Setup,
		registryschemas.Setup,
		schemaschemas.Setup,
		secret.Setup,
		secretpolicy.Setup,
		secretrotation.Setup,
		secretversion.Setup,
		accountsecurityhub.Setup,
		actiontarget.Setup,
		findingaggregator.Setup,
		insight.Setup,
		inviteaccepter.Setup,
		membersecurityhub.Setup,
		productsubscription.Setup,
		standardssubscription.Setup,
		cloudformationstack.Setup,
		budgetresourceassociation.Setup,
		constraint.Setup,
		portfolio.Setup,
		portfolioshare.Setup,
		principalportfolioassociation.Setup,
		product.Setup,
		productportfolioassociation.Setup,
		provisioningartifact.Setup,
		serviceaction.Setup,
		tagoption.Setup,
		tagoptionresourceassociation.Setup,
		httpnamespace.Setup,
		privatednsnamespace.Setup,
		publicdnsnamespace.Setup,
		serviceservicediscovery.Setup,
		servicequota.Setup,
		activereceiptruleset.Setup,
		configurationset.Setup,
		domaindkim.Setup,
		domainidentity.Setup,
		domainmailfrom.Setup,
		emailidentity.Setup,
		eventdestination.Setup,
		identitynotificationtopic.Setup,
		identitypolicy.Setup,
		receiptfilter.Setup,
		receiptrule.Setup,
		receiptruleset.Setup,
		template.Setup,
		configurationsetsesv2.Setup,
		configurationseteventdestination.Setup,
		dedicatedippool.Setup,
		emailidentitysesv2.Setup,
		emailidentityfeedbackattributes.Setup,
		emailidentitymailfromattributes.Setup,
		activity.Setup,
		statemachine.Setup,
		signingjob.Setup,
		signingprofile.Setup,
		signingprofilepermission.Setup,
		domainsimpledb.Setup,
		platformapplication.Setup,
		smspreferences.Setup,
		topic.Setup,
		topicpolicy.Setup,
		topicsubscription.Setup,
		queuesqs.Setup,
		queuepolicy.Setup,
		queueredriveallowpolicy.Setup,
		queueredrivepolicy.Setup,
		activation.Setup,
		associationssm.Setup,
		defaultpatchbaseline.Setup,
		document.Setup,
		maintenancewindow.Setup,
		maintenancewindowtarget.Setup,
		maintenancewindowtask.Setup,
		parameter.Setup,
		patchbaseline.Setup,
		patchgroup.Setup,
		resourcedatasync.Setup,
		servicesetting.Setup,
		accountassignment.Setup,
		customermanagedpolicyattachment.Setup,
		instanceaccesscontrolattributes.Setup,
		managedpolicyattachment.Setup,
		permissionsboundaryattachment.Setup,
		permissionset.Setup,
		permissionsetinlinepolicy.Setup,
		domainswf.Setup,
		databasetimestreamwrite.Setup,
		tabletimestreamwrite.Setup,
		languagemodel.Setup,
		vocabularytranscribe.Setup,
		vocabularyfilter.Setup,
		server.Setup,
		sshkey.Setup,
		tagtransfer.Setup,
		usertransfer.Setup,
		workflowtransfer.Setup,
		networkperformancemetricsubscription.Setup,
		bytematchset.Setup,
		geomatchset.Setup,
		ipset.Setup,
		ratebasedrule.Setup,
		regexmatchset.Setup,
		regexpatternset.Setup,
		rulewaf.Setup,
		sizeconstraintset.Setup,
		sqlinjectionmatchset.Setup,
		webacl.Setup,
		xssmatchset.Setup,
		bytematchsetwafregional.Setup,
		geomatchsetwafregional.Setup,
		ipsetwafregional.Setup,
		ratebasedrulewafregional.Setup,
		regexmatchsetwafregional.Setup,
		regexpatternsetwafregional.Setup,
		rulewafregional.Setup,
		sizeconstraintsetwafregional.Setup,
		sqlinjectionmatchsetwafregional.Setup,
		webaclwafregional.Setup,
		xssmatchsetwafregional.Setup,
		ipsetwafv2.Setup,
		regexpatternsetwafv2.Setup,
		directoryworkspaces.Setup,
		ipgroup.Setup,
		encryptionconfig.Setup,
		groupxray.Setup,
		samplingrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
