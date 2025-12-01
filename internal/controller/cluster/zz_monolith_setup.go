// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	analyzer "github.com/upbound/provider-aws/internal/controller/cluster/accessanalyzer/analyzer"
	archiverule "github.com/upbound/provider-aws/internal/controller/cluster/accessanalyzer/archiverule"
	alternatecontact "github.com/upbound/provider-aws/internal/controller/cluster/account/alternatecontact"
	region "github.com/upbound/provider-aws/internal/controller/cluster/account/region"
	certificate "github.com/upbound/provider-aws/internal/controller/cluster/acm/certificate"
	certificatevalidation "github.com/upbound/provider-aws/internal/controller/cluster/acm/certificatevalidation"
	certificateacmpca "github.com/upbound/provider-aws/internal/controller/cluster/acmpca/certificate"
	certificateauthority "github.com/upbound/provider-aws/internal/controller/cluster/acmpca/certificateauthority"
	certificateauthoritycertificate "github.com/upbound/provider-aws/internal/controller/cluster/acmpca/certificateauthoritycertificate"
	permission "github.com/upbound/provider-aws/internal/controller/cluster/acmpca/permission"
	policy "github.com/upbound/provider-aws/internal/controller/cluster/acmpca/policy"
	alertmanagerdefinition "github.com/upbound/provider-aws/internal/controller/cluster/amp/alertmanagerdefinition"
	rulegroupnamespace "github.com/upbound/provider-aws/internal/controller/cluster/amp/rulegroupnamespace"
	scraper "github.com/upbound/provider-aws/internal/controller/cluster/amp/scraper"
	workspace "github.com/upbound/provider-aws/internal/controller/cluster/amp/workspace"
	app "github.com/upbound/provider-aws/internal/controller/cluster/amplify/app"
	backendenvironment "github.com/upbound/provider-aws/internal/controller/cluster/amplify/backendenvironment"
	branch "github.com/upbound/provider-aws/internal/controller/cluster/amplify/branch"
	webhook "github.com/upbound/provider-aws/internal/controller/cluster/amplify/webhook"
	account "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/account"
	apikey "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/apikey"
	authorizer "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/authorizer"
	basepathmapping "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/basepathmapping"
	clientcertificate "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/clientcertificate"
	deployment "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/deployment"
	documentationpart "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/documentationpart"
	documentationversion "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/documentationversion"
	domainname "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/domainname"
	gatewayresponse "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/gatewayresponse"
	integration "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/integration"
	integrationresponse "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/integrationresponse"
	method "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/method"
	methodresponse "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/methodresponse"
	methodsettings "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/methodsettings"
	model "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/model"
	requestvalidator "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/requestvalidator"
	resource "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/resource"
	restapi "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/restapi"
	restapipolicy "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/restapipolicy"
	stage "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/stage"
	usageplan "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/usageplan"
	usageplankey "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/usageplankey"
	vpclink "github.com/upbound/provider-aws/internal/controller/cluster/apigateway/vpclink"
	api "github.com/upbound/provider-aws/internal/controller/cluster/apigatewayv2/api"
	apimapping "github.com/upbound/provider-aws/internal/controller/cluster/apigatewayv2/apimapping"
	authorizerapigatewayv2 "github.com/upbound/provider-aws/internal/controller/cluster/apigatewayv2/authorizer"
	deploymentapigatewayv2 "github.com/upbound/provider-aws/internal/controller/cluster/apigatewayv2/deployment"
	domainnameapigatewayv2 "github.com/upbound/provider-aws/internal/controller/cluster/apigatewayv2/domainname"
	integrationapigatewayv2 "github.com/upbound/provider-aws/internal/controller/cluster/apigatewayv2/integration"
	integrationresponseapigatewayv2 "github.com/upbound/provider-aws/internal/controller/cluster/apigatewayv2/integrationresponse"
	modelapigatewayv2 "github.com/upbound/provider-aws/internal/controller/cluster/apigatewayv2/model"
	route "github.com/upbound/provider-aws/internal/controller/cluster/apigatewayv2/route"
	routeresponse "github.com/upbound/provider-aws/internal/controller/cluster/apigatewayv2/routeresponse"
	stageapigatewayv2 "github.com/upbound/provider-aws/internal/controller/cluster/apigatewayv2/stage"
	vpclinkapigatewayv2 "github.com/upbound/provider-aws/internal/controller/cluster/apigatewayv2/vpclink"
	policyappautoscaling "github.com/upbound/provider-aws/internal/controller/cluster/appautoscaling/policy"
	scheduledaction "github.com/upbound/provider-aws/internal/controller/cluster/appautoscaling/scheduledaction"
	target "github.com/upbound/provider-aws/internal/controller/cluster/appautoscaling/target"
	application "github.com/upbound/provider-aws/internal/controller/cluster/appconfig/application"
	configurationprofile "github.com/upbound/provider-aws/internal/controller/cluster/appconfig/configurationprofile"
	deploymentappconfig "github.com/upbound/provider-aws/internal/controller/cluster/appconfig/deployment"
	deploymentstrategy "github.com/upbound/provider-aws/internal/controller/cluster/appconfig/deploymentstrategy"
	environment "github.com/upbound/provider-aws/internal/controller/cluster/appconfig/environment"
	extension "github.com/upbound/provider-aws/internal/controller/cluster/appconfig/extension"
	extensionassociation "github.com/upbound/provider-aws/internal/controller/cluster/appconfig/extensionassociation"
	hostedconfigurationversion "github.com/upbound/provider-aws/internal/controller/cluster/appconfig/hostedconfigurationversion"
	flow "github.com/upbound/provider-aws/internal/controller/cluster/appflow/flow"
	eventintegration "github.com/upbound/provider-aws/internal/controller/cluster/appintegrations/eventintegration"
	applicationapplicationinsights "github.com/upbound/provider-aws/internal/controller/cluster/applicationinsights/application"
	gatewayroute "github.com/upbound/provider-aws/internal/controller/cluster/appmesh/gatewayroute"
	mesh "github.com/upbound/provider-aws/internal/controller/cluster/appmesh/mesh"
	routeappmesh "github.com/upbound/provider-aws/internal/controller/cluster/appmesh/route"
	virtualgateway "github.com/upbound/provider-aws/internal/controller/cluster/appmesh/virtualgateway"
	virtualnode "github.com/upbound/provider-aws/internal/controller/cluster/appmesh/virtualnode"
	virtualrouter "github.com/upbound/provider-aws/internal/controller/cluster/appmesh/virtualrouter"
	virtualservice "github.com/upbound/provider-aws/internal/controller/cluster/appmesh/virtualservice"
	autoscalingconfigurationversion "github.com/upbound/provider-aws/internal/controller/cluster/apprunner/autoscalingconfigurationversion"
	connection "github.com/upbound/provider-aws/internal/controller/cluster/apprunner/connection"
	observabilityconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/apprunner/observabilityconfiguration"
	service "github.com/upbound/provider-aws/internal/controller/cluster/apprunner/service"
	vpcconnector "github.com/upbound/provider-aws/internal/controller/cluster/apprunner/vpcconnector"
	directoryconfig "github.com/upbound/provider-aws/internal/controller/cluster/appstream/directoryconfig"
	fleet "github.com/upbound/provider-aws/internal/controller/cluster/appstream/fleet"
	fleetstackassociation "github.com/upbound/provider-aws/internal/controller/cluster/appstream/fleetstackassociation"
	imagebuilder "github.com/upbound/provider-aws/internal/controller/cluster/appstream/imagebuilder"
	stack "github.com/upbound/provider-aws/internal/controller/cluster/appstream/stack"
	user "github.com/upbound/provider-aws/internal/controller/cluster/appstream/user"
	userstackassociation "github.com/upbound/provider-aws/internal/controller/cluster/appstream/userstackassociation"
	apicache "github.com/upbound/provider-aws/internal/controller/cluster/appsync/apicache"
	apikeyappsync "github.com/upbound/provider-aws/internal/controller/cluster/appsync/apikey"
	datasource "github.com/upbound/provider-aws/internal/controller/cluster/appsync/datasource"
	function "github.com/upbound/provider-aws/internal/controller/cluster/appsync/function"
	graphqlapi "github.com/upbound/provider-aws/internal/controller/cluster/appsync/graphqlapi"
	resolver "github.com/upbound/provider-aws/internal/controller/cluster/appsync/resolver"
	database "github.com/upbound/provider-aws/internal/controller/cluster/athena/database"
	datacatalog "github.com/upbound/provider-aws/internal/controller/cluster/athena/datacatalog"
	namedquery "github.com/upbound/provider-aws/internal/controller/cluster/athena/namedquery"
	workgroup "github.com/upbound/provider-aws/internal/controller/cluster/athena/workgroup"
	attachment "github.com/upbound/provider-aws/internal/controller/cluster/autoscaling/attachment"
	autoscalinggroup "github.com/upbound/provider-aws/internal/controller/cluster/autoscaling/autoscalinggroup"
	grouptag "github.com/upbound/provider-aws/internal/controller/cluster/autoscaling/grouptag"
	launchconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/autoscaling/launchconfiguration"
	lifecyclehook "github.com/upbound/provider-aws/internal/controller/cluster/autoscaling/lifecyclehook"
	notification "github.com/upbound/provider-aws/internal/controller/cluster/autoscaling/notification"
	policyautoscaling "github.com/upbound/provider-aws/internal/controller/cluster/autoscaling/policy"
	schedule "github.com/upbound/provider-aws/internal/controller/cluster/autoscaling/schedule"
	scalingplan "github.com/upbound/provider-aws/internal/controller/cluster/autoscalingplans/scalingplan"
	framework "github.com/upbound/provider-aws/internal/controller/cluster/backup/framework"
	globalsettings "github.com/upbound/provider-aws/internal/controller/cluster/backup/globalsettings"
	plan "github.com/upbound/provider-aws/internal/controller/cluster/backup/plan"
	regionsettings "github.com/upbound/provider-aws/internal/controller/cluster/backup/regionsettings"
	reportplan "github.com/upbound/provider-aws/internal/controller/cluster/backup/reportplan"
	selection "github.com/upbound/provider-aws/internal/controller/cluster/backup/selection"
	vault "github.com/upbound/provider-aws/internal/controller/cluster/backup/vault"
	vaultlockconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/backup/vaultlockconfiguration"
	vaultnotifications "github.com/upbound/provider-aws/internal/controller/cluster/backup/vaultnotifications"
	vaultpolicy "github.com/upbound/provider-aws/internal/controller/cluster/backup/vaultpolicy"
	computeenvironment "github.com/upbound/provider-aws/internal/controller/cluster/batch/computeenvironment"
	jobdefinition "github.com/upbound/provider-aws/internal/controller/cluster/batch/jobdefinition"
	jobqueue "github.com/upbound/provider-aws/internal/controller/cluster/batch/jobqueue"
	schedulingpolicy "github.com/upbound/provider-aws/internal/controller/cluster/batch/schedulingpolicy"
	inferenceprofile "github.com/upbound/provider-aws/internal/controller/cluster/bedrock/inferenceprofile"
	agent "github.com/upbound/provider-aws/internal/controller/cluster/bedrockagent/agent"
	budget "github.com/upbound/provider-aws/internal/controller/cluster/budgets/budget"
	budgetaction "github.com/upbound/provider-aws/internal/controller/cluster/budgets/budgetaction"
	anomalymonitor "github.com/upbound/provider-aws/internal/controller/cluster/ce/anomalymonitor"
	voiceconnector "github.com/upbound/provider-aws/internal/controller/cluster/chime/voiceconnector"
	voiceconnectorgroup "github.com/upbound/provider-aws/internal/controller/cluster/chime/voiceconnectorgroup"
	voiceconnectorlogging "github.com/upbound/provider-aws/internal/controller/cluster/chime/voiceconnectorlogging"
	voiceconnectororigination "github.com/upbound/provider-aws/internal/controller/cluster/chime/voiceconnectororigination"
	voiceconnectorstreaming "github.com/upbound/provider-aws/internal/controller/cluster/chime/voiceconnectorstreaming"
	voiceconnectortermination "github.com/upbound/provider-aws/internal/controller/cluster/chime/voiceconnectortermination"
	voiceconnectorterminationcredentials "github.com/upbound/provider-aws/internal/controller/cluster/chime/voiceconnectorterminationcredentials"
	environmentec2 "github.com/upbound/provider-aws/internal/controller/cluster/cloud9/environmentec2"
	environmentmembership "github.com/upbound/provider-aws/internal/controller/cluster/cloud9/environmentmembership"
	resourcecloudcontrol "github.com/upbound/provider-aws/internal/controller/cluster/cloudcontrol/resource"
	stackcloudformation "github.com/upbound/provider-aws/internal/controller/cluster/cloudformation/stack"
	stackset "github.com/upbound/provider-aws/internal/controller/cluster/cloudformation/stackset"
	stacksetinstance "github.com/upbound/provider-aws/internal/controller/cluster/cloudformation/stacksetinstance"
	cachepolicy "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/cachepolicy"
	distribution "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/distribution"
	fieldlevelencryptionconfig "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/fieldlevelencryptionconfig"
	fieldlevelencryptionprofile "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/fieldlevelencryptionprofile"
	functioncloudfront "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/function"
	keygroup "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/keygroup"
	monitoringsubscription "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/monitoringsubscription"
	originaccesscontrol "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/originaccesscontrol"
	originaccessidentity "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/originaccessidentity"
	originrequestpolicy "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/originrequestpolicy"
	publickey "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/publickey"
	realtimelogconfig "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/realtimelogconfig"
	responseheaderspolicy "github.com/upbound/provider-aws/internal/controller/cluster/cloudfront/responseheaderspolicy"
	domain "github.com/upbound/provider-aws/internal/controller/cluster/cloudsearch/domain"
	domainserviceaccesspolicy "github.com/upbound/provider-aws/internal/controller/cluster/cloudsearch/domainserviceaccesspolicy"
	eventdatastore "github.com/upbound/provider-aws/internal/controller/cluster/cloudtrail/eventdatastore"
	trail "github.com/upbound/provider-aws/internal/controller/cluster/cloudtrail/trail"
	compositealarm "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatch/compositealarm"
	dashboard "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatch/dashboard"
	metricalarm "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatch/metricalarm"
	metricstream "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatch/metricstream"
	apidestination "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/apidestination"
	archive "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/archive"
	bus "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/bus"
	buspolicy "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/buspolicy"
	connectioncloudwatchevents "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/connection"
	permissioncloudwatchevents "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/permission"
	rule "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/rule"
	targetcloudwatchevents "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/target"
	definition "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchlogs/definition"
	destination "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchlogs/destination"
	destinationpolicy "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchlogs/destinationpolicy"
	group "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchlogs/group"
	metricfilter "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchlogs/metricfilter"
	resourcepolicy "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchlogs/resourcepolicy"
	stream "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchlogs/stream"
	subscriptionfilter "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchlogs/subscriptionfilter"
	domaincodeartifact "github.com/upbound/provider-aws/internal/controller/cluster/codeartifact/domain"
	domainpermissionspolicy "github.com/upbound/provider-aws/internal/controller/cluster/codeartifact/domainpermissionspolicy"
	repository "github.com/upbound/provider-aws/internal/controller/cluster/codeartifact/repository"
	repositorypermissionspolicy "github.com/upbound/provider-aws/internal/controller/cluster/codeartifact/repositorypermissionspolicy"
	approvalruletemplate "github.com/upbound/provider-aws/internal/controller/cluster/codecommit/approvalruletemplate"
	approvalruletemplateassociation "github.com/upbound/provider-aws/internal/controller/cluster/codecommit/approvalruletemplateassociation"
	repositorycodecommit "github.com/upbound/provider-aws/internal/controller/cluster/codecommit/repository"
	trigger "github.com/upbound/provider-aws/internal/controller/cluster/codecommit/trigger"
	profilinggroup "github.com/upbound/provider-aws/internal/controller/cluster/codeguruprofiler/profilinggroup"
	codepipeline "github.com/upbound/provider-aws/internal/controller/cluster/codepipeline/codepipeline"
	customactiontype "github.com/upbound/provider-aws/internal/controller/cluster/codepipeline/customactiontype"
	webhookcodepipeline "github.com/upbound/provider-aws/internal/controller/cluster/codepipeline/webhook"
	connectioncodestarconnections "github.com/upbound/provider-aws/internal/controller/cluster/codestarconnections/connection"
	host "github.com/upbound/provider-aws/internal/controller/cluster/codestarconnections/host"
	notificationrule "github.com/upbound/provider-aws/internal/controller/cluster/codestarnotifications/notificationrule"
	cognitoidentitypoolproviderprincipaltag "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidentity/cognitoidentitypoolproviderprincipaltag"
	pool "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidentity/pool"
	poolrolesattachment "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidentity/poolrolesattachment"
	identityprovider "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/identityprovider"
	resourceserver "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/resourceserver"
	riskconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/riskconfiguration"
	usercognitoidp "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/user"
	usergroup "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/usergroup"
	useringroup "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/useringroup"
	userpool "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/userpool"
	userpoolclient "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/userpoolclient"
	userpooldomain "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/userpooldomain"
	userpooluicustomization "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/userpooluicustomization"
	awsconfigurationrecorderstatus "github.com/upbound/provider-aws/internal/controller/cluster/configservice/awsconfigurationrecorderstatus"
	configrule "github.com/upbound/provider-aws/internal/controller/cluster/configservice/configrule"
	configurationaggregator "github.com/upbound/provider-aws/internal/controller/cluster/configservice/configurationaggregator"
	configurationrecorder "github.com/upbound/provider-aws/internal/controller/cluster/configservice/configurationrecorder"
	conformancepack "github.com/upbound/provider-aws/internal/controller/cluster/configservice/conformancepack"
	deliverychannel "github.com/upbound/provider-aws/internal/controller/cluster/configservice/deliverychannel"
	remediationconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/configservice/remediationconfiguration"
	botassociation "github.com/upbound/provider-aws/internal/controller/cluster/connect/botassociation"
	contactflow "github.com/upbound/provider-aws/internal/controller/cluster/connect/contactflow"
	contactflowmodule "github.com/upbound/provider-aws/internal/controller/cluster/connect/contactflowmodule"
	hoursofoperation "github.com/upbound/provider-aws/internal/controller/cluster/connect/hoursofoperation"
	instance "github.com/upbound/provider-aws/internal/controller/cluster/connect/instance"
	instancestorageconfig "github.com/upbound/provider-aws/internal/controller/cluster/connect/instancestorageconfig"
	lambdafunctionassociation "github.com/upbound/provider-aws/internal/controller/cluster/connect/lambdafunctionassociation"
	phonenumber "github.com/upbound/provider-aws/internal/controller/cluster/connect/phonenumber"
	queue "github.com/upbound/provider-aws/internal/controller/cluster/connect/queue"
	quickconnect "github.com/upbound/provider-aws/internal/controller/cluster/connect/quickconnect"
	routingprofile "github.com/upbound/provider-aws/internal/controller/cluster/connect/routingprofile"
	securityprofile "github.com/upbound/provider-aws/internal/controller/cluster/connect/securityprofile"
	userconnect "github.com/upbound/provider-aws/internal/controller/cluster/connect/user"
	userhierarchystructure "github.com/upbound/provider-aws/internal/controller/cluster/connect/userhierarchystructure"
	vocabulary "github.com/upbound/provider-aws/internal/controller/cluster/connect/vocabulary"
	reportdefinition "github.com/upbound/provider-aws/internal/controller/cluster/cur/reportdefinition"
	dataset "github.com/upbound/provider-aws/internal/controller/cluster/dataexchange/dataset"
	revision "github.com/upbound/provider-aws/internal/controller/cluster/dataexchange/revision"
	pipeline "github.com/upbound/provider-aws/internal/controller/cluster/datapipeline/pipeline"
	locations3 "github.com/upbound/provider-aws/internal/controller/cluster/datasync/locations3"
	task "github.com/upbound/provider-aws/internal/controller/cluster/datasync/task"
	cluster "github.com/upbound/provider-aws/internal/controller/cluster/dax/cluster"
	parametergroup "github.com/upbound/provider-aws/internal/controller/cluster/dax/parametergroup"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/cluster/dax/subnetgroup"
	appdeploy "github.com/upbound/provider-aws/internal/controller/cluster/deploy/app"
	deploymentconfig "github.com/upbound/provider-aws/internal/controller/cluster/deploy/deploymentconfig"
	deploymentgroup "github.com/upbound/provider-aws/internal/controller/cluster/deploy/deploymentgroup"
	graph "github.com/upbound/provider-aws/internal/controller/cluster/detective/graph"
	invitationaccepter "github.com/upbound/provider-aws/internal/controller/cluster/detective/invitationaccepter"
	member "github.com/upbound/provider-aws/internal/controller/cluster/detective/member"
	devicepool "github.com/upbound/provider-aws/internal/controller/cluster/devicefarm/devicepool"
	instanceprofile "github.com/upbound/provider-aws/internal/controller/cluster/devicefarm/instanceprofile"
	networkprofile "github.com/upbound/provider-aws/internal/controller/cluster/devicefarm/networkprofile"
	project "github.com/upbound/provider-aws/internal/controller/cluster/devicefarm/project"
	testgridproject "github.com/upbound/provider-aws/internal/controller/cluster/devicefarm/testgridproject"
	upload "github.com/upbound/provider-aws/internal/controller/cluster/devicefarm/upload"
	bgppeer "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/bgppeer"
	connectiondirectconnect "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/connection"
	connectionassociation "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/connectionassociation"
	gateway "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/gateway"
	gatewayassociation "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/gatewayassociation"
	gatewayassociationproposal "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/gatewayassociationproposal"
	hostedprivatevirtualinterface "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/hostedprivatevirtualinterface"
	hostedprivatevirtualinterfaceaccepter "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/hostedprivatevirtualinterfaceaccepter"
	hostedpublicvirtualinterface "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/hostedpublicvirtualinterface"
	hostedpublicvirtualinterfaceaccepter "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/hostedpublicvirtualinterfaceaccepter"
	hostedtransitvirtualinterface "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/hostedtransitvirtualinterface"
	hostedtransitvirtualinterfaceaccepter "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/hostedtransitvirtualinterfaceaccepter"
	lag "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/lag"
	privatevirtualinterface "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/privatevirtualinterface"
	publicvirtualinterface "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/publicvirtualinterface"
	transitvirtualinterface "github.com/upbound/provider-aws/internal/controller/cluster/directconnect/transitvirtualinterface"
	lifecyclepolicy "github.com/upbound/provider-aws/internal/controller/cluster/dlm/lifecyclepolicy"
	certificatedms "github.com/upbound/provider-aws/internal/controller/cluster/dms/certificate"
	endpoint "github.com/upbound/provider-aws/internal/controller/cluster/dms/endpoint"
	eventsubscription "github.com/upbound/provider-aws/internal/controller/cluster/dms/eventsubscription"
	replicationinstance "github.com/upbound/provider-aws/internal/controller/cluster/dms/replicationinstance"
	replicationsubnetgroup "github.com/upbound/provider-aws/internal/controller/cluster/dms/replicationsubnetgroup"
	replicationtask "github.com/upbound/provider-aws/internal/controller/cluster/dms/replicationtask"
	s3endpoint "github.com/upbound/provider-aws/internal/controller/cluster/dms/s3endpoint"
	clusterdocdb "github.com/upbound/provider-aws/internal/controller/cluster/docdb/cluster"
	clusterinstance "github.com/upbound/provider-aws/internal/controller/cluster/docdb/clusterinstance"
	clusterparametergroup "github.com/upbound/provider-aws/internal/controller/cluster/docdb/clusterparametergroup"
	clustersnapshot "github.com/upbound/provider-aws/internal/controller/cluster/docdb/clustersnapshot"
	eventsubscriptiondocdb "github.com/upbound/provider-aws/internal/controller/cluster/docdb/eventsubscription"
	globalcluster "github.com/upbound/provider-aws/internal/controller/cluster/docdb/globalcluster"
	subnetgroupdocdb "github.com/upbound/provider-aws/internal/controller/cluster/docdb/subnetgroup"
	conditionalforwarder "github.com/upbound/provider-aws/internal/controller/cluster/ds/conditionalforwarder"
	directory "github.com/upbound/provider-aws/internal/controller/cluster/ds/directory"
	shareddirectory "github.com/upbound/provider-aws/internal/controller/cluster/ds/shareddirectory"
	clusterdsql "github.com/upbound/provider-aws/internal/controller/cluster/dsql/cluster"
	clusterpeering "github.com/upbound/provider-aws/internal/controller/cluster/dsql/clusterpeering"
	contributorinsights "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/contributorinsights"
	globaltable "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/globaltable"
	kinesisstreamingdestination "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/kinesisstreamingdestination"
	resourcepolicydynamodb "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/resourcepolicy"
	table "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/table"
	tableitem "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/tableitem"
	tablereplica "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/tablereplica"
	tag "github.com/upbound/provider-aws/internal/controller/cluster/dynamodb/tag"
	ami "github.com/upbound/provider-aws/internal/controller/cluster/ec2/ami"
	amicopy "github.com/upbound/provider-aws/internal/controller/cluster/ec2/amicopy"
	amilaunchpermission "github.com/upbound/provider-aws/internal/controller/cluster/ec2/amilaunchpermission"
	availabilityzonegroup "github.com/upbound/provider-aws/internal/controller/cluster/ec2/availabilityzonegroup"
	capacityreservation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/capacityreservation"
	carriergateway "github.com/upbound/provider-aws/internal/controller/cluster/ec2/carriergateway"
	customergateway "github.com/upbound/provider-aws/internal/controller/cluster/ec2/customergateway"
	defaultnetworkacl "github.com/upbound/provider-aws/internal/controller/cluster/ec2/defaultnetworkacl"
	defaultroutetable "github.com/upbound/provider-aws/internal/controller/cluster/ec2/defaultroutetable"
	defaultsecuritygroup "github.com/upbound/provider-aws/internal/controller/cluster/ec2/defaultsecuritygroup"
	defaultsubnet "github.com/upbound/provider-aws/internal/controller/cluster/ec2/defaultsubnet"
	defaultvpc "github.com/upbound/provider-aws/internal/controller/cluster/ec2/defaultvpc"
	defaultvpcdhcpoptions "github.com/upbound/provider-aws/internal/controller/cluster/ec2/defaultvpcdhcpoptions"
	ebsdefaultkmskey "github.com/upbound/provider-aws/internal/controller/cluster/ec2/ebsdefaultkmskey"
	ebsencryptionbydefault "github.com/upbound/provider-aws/internal/controller/cluster/ec2/ebsencryptionbydefault"
	ebssnapshot "github.com/upbound/provider-aws/internal/controller/cluster/ec2/ebssnapshot"
	ebssnapshotcopy "github.com/upbound/provider-aws/internal/controller/cluster/ec2/ebssnapshotcopy"
	ebssnapshotimport "github.com/upbound/provider-aws/internal/controller/cluster/ec2/ebssnapshotimport"
	ebsvolume "github.com/upbound/provider-aws/internal/controller/cluster/ec2/ebsvolume"
	egressonlyinternetgateway "github.com/upbound/provider-aws/internal/controller/cluster/ec2/egressonlyinternetgateway"
	eip "github.com/upbound/provider-aws/internal/controller/cluster/ec2/eip"
	eipassociation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/eipassociation"
	fleetec2 "github.com/upbound/provider-aws/internal/controller/cluster/ec2/fleet"
	flowlog "github.com/upbound/provider-aws/internal/controller/cluster/ec2/flowlog"
	hostec2 "github.com/upbound/provider-aws/internal/controller/cluster/ec2/host"
	instanceec2 "github.com/upbound/provider-aws/internal/controller/cluster/ec2/instance"
	instancestate "github.com/upbound/provider-aws/internal/controller/cluster/ec2/instancestate"
	internetgateway "github.com/upbound/provider-aws/internal/controller/cluster/ec2/internetgateway"
	keypair "github.com/upbound/provider-aws/internal/controller/cluster/ec2/keypair"
	launchtemplate "github.com/upbound/provider-aws/internal/controller/cluster/ec2/launchtemplate"
	mainroutetableassociation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/mainroutetableassociation"
	managedprefixlist "github.com/upbound/provider-aws/internal/controller/cluster/ec2/managedprefixlist"
	managedprefixlistentry "github.com/upbound/provider-aws/internal/controller/cluster/ec2/managedprefixlistentry"
	natgateway "github.com/upbound/provider-aws/internal/controller/cluster/ec2/natgateway"
	networkacl "github.com/upbound/provider-aws/internal/controller/cluster/ec2/networkacl"
	networkaclrule "github.com/upbound/provider-aws/internal/controller/cluster/ec2/networkaclrule"
	networkinsightsanalysis "github.com/upbound/provider-aws/internal/controller/cluster/ec2/networkinsightsanalysis"
	networkinsightspath "github.com/upbound/provider-aws/internal/controller/cluster/ec2/networkinsightspath"
	networkinterface "github.com/upbound/provider-aws/internal/controller/cluster/ec2/networkinterface"
	networkinterfaceattachment "github.com/upbound/provider-aws/internal/controller/cluster/ec2/networkinterfaceattachment"
	networkinterfacesgattachment "github.com/upbound/provider-aws/internal/controller/cluster/ec2/networkinterfacesgattachment"
	placementgroup "github.com/upbound/provider-aws/internal/controller/cluster/ec2/placementgroup"
	routeec2 "github.com/upbound/provider-aws/internal/controller/cluster/ec2/route"
	routetable "github.com/upbound/provider-aws/internal/controller/cluster/ec2/routetable"
	routetableassociation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/routetableassociation"
	securitygroup "github.com/upbound/provider-aws/internal/controller/cluster/ec2/securitygroup"
	securitygroupegressrule "github.com/upbound/provider-aws/internal/controller/cluster/ec2/securitygroupegressrule"
	securitygroupingressrule "github.com/upbound/provider-aws/internal/controller/cluster/ec2/securitygroupingressrule"
	securitygrouprule "github.com/upbound/provider-aws/internal/controller/cluster/ec2/securitygrouprule"
	serialconsoleaccess "github.com/upbound/provider-aws/internal/controller/cluster/ec2/serialconsoleaccess"
	snapshotcreatevolumepermission "github.com/upbound/provider-aws/internal/controller/cluster/ec2/snapshotcreatevolumepermission"
	spotdatafeedsubscription "github.com/upbound/provider-aws/internal/controller/cluster/ec2/spotdatafeedsubscription"
	spotfleetrequest "github.com/upbound/provider-aws/internal/controller/cluster/ec2/spotfleetrequest"
	spotinstancerequest "github.com/upbound/provider-aws/internal/controller/cluster/ec2/spotinstancerequest"
	subnet "github.com/upbound/provider-aws/internal/controller/cluster/ec2/subnet"
	subnetcidrreservation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/subnetcidrreservation"
	tagec2 "github.com/upbound/provider-aws/internal/controller/cluster/ec2/tag"
	trafficmirrorfilter "github.com/upbound/provider-aws/internal/controller/cluster/ec2/trafficmirrorfilter"
	trafficmirrorfilterrule "github.com/upbound/provider-aws/internal/controller/cluster/ec2/trafficmirrorfilterrule"
	transitgateway "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgateway"
	transitgatewayconnect "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewayconnect"
	transitgatewayconnectpeer "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewayconnectpeer"
	transitgatewaymulticastdomain "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewaymulticastdomain"
	transitgatewaymulticastdomainassociation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewaymulticastdomainassociation"
	transitgatewaymulticastgroupmember "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewaymulticastgroupmember"
	transitgatewaymulticastgroupsource "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewaymulticastgroupsource"
	transitgatewaypeeringattachment "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewaypeeringattachment"
	transitgatewaypeeringattachmentaccepter "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewaypeeringattachmentaccepter"
	transitgatewaypolicytable "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewaypolicytable"
	transitgatewayprefixlistreference "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewayprefixlistreference"
	transitgatewayroute "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/upbound/provider-aws/internal/controller/cluster/ec2/transitgatewayvpcattachmentaccepter"
	volumeattachment "github.com/upbound/provider-aws/internal/controller/cluster/ec2/volumeattachment"
	vpc "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpc"
	vpcdhcpoptions "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcdhcpoptions"
	vpcdhcpoptionsassociation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcdhcpoptionsassociation"
	vpcendpoint "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcendpoint"
	vpcendpointconnectionnotification "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcendpointconnectionnotification"
	vpcendpointroutetableassociation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcendpointroutetableassociation"
	vpcendpointsecuritygroupassociation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcendpointsecuritygroupassociation"
	vpcendpointservice "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcendpointservice"
	vpcendpointserviceallowedprincipal "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcendpointserviceallowedprincipal"
	vpcendpointsubnetassociation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcendpointsubnetassociation"
	vpcipam "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcipam"
	vpcipampool "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcipampool"
	vpcipampoolcidr "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcipampoolcidr"
	vpcipampoolcidrallocation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcipampoolcidrallocation"
	vpcipamscope "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcipamscope"
	vpcipv4cidrblockassociation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcpeeringconnection"
	vpcpeeringconnectionaccepter "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcpeeringconnectionaccepter"
	vpcpeeringconnectionoptions "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpcpeeringconnectionoptions"
	vpnconnection "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpnconnection"
	vpnconnectionroute "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpnconnectionroute"
	vpngateway "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpngateway"
	vpngatewayattachment "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpngatewayattachment"
	vpngatewayroutepropagation "github.com/upbound/provider-aws/internal/controller/cluster/ec2/vpngatewayroutepropagation"
	lifecyclepolicyecr "github.com/upbound/provider-aws/internal/controller/cluster/ecr/lifecyclepolicy"
	pullthroughcacherule "github.com/upbound/provider-aws/internal/controller/cluster/ecr/pullthroughcacherule"
	registrypolicy "github.com/upbound/provider-aws/internal/controller/cluster/ecr/registrypolicy"
	registryscanningconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/ecr/registryscanningconfiguration"
	replicationconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/ecr/replicationconfiguration"
	repositoryecr "github.com/upbound/provider-aws/internal/controller/cluster/ecr/repository"
	repositorypolicy "github.com/upbound/provider-aws/internal/controller/cluster/ecr/repositorypolicy"
	repositoryecrpublic "github.com/upbound/provider-aws/internal/controller/cluster/ecrpublic/repository"
	repositorypolicyecrpublic "github.com/upbound/provider-aws/internal/controller/cluster/ecrpublic/repositorypolicy"
	accountsettingdefault "github.com/upbound/provider-aws/internal/controller/cluster/ecs/accountsettingdefault"
	capacityprovider "github.com/upbound/provider-aws/internal/controller/cluster/ecs/capacityprovider"
	clusterecs "github.com/upbound/provider-aws/internal/controller/cluster/ecs/cluster"
	clustercapacityproviders "github.com/upbound/provider-aws/internal/controller/cluster/ecs/clustercapacityproviders"
	serviceecs "github.com/upbound/provider-aws/internal/controller/cluster/ecs/service"
	taskdefinition "github.com/upbound/provider-aws/internal/controller/cluster/ecs/taskdefinition"
	accesspoint "github.com/upbound/provider-aws/internal/controller/cluster/efs/accesspoint"
	backuppolicy "github.com/upbound/provider-aws/internal/controller/cluster/efs/backuppolicy"
	filesystem "github.com/upbound/provider-aws/internal/controller/cluster/efs/filesystem"
	filesystempolicy "github.com/upbound/provider-aws/internal/controller/cluster/efs/filesystempolicy"
	mounttarget "github.com/upbound/provider-aws/internal/controller/cluster/efs/mounttarget"
	replicationconfigurationefs "github.com/upbound/provider-aws/internal/controller/cluster/efs/replicationconfiguration"
	accessentry "github.com/upbound/provider-aws/internal/controller/cluster/eks/accessentry"
	accesspolicyassociation "github.com/upbound/provider-aws/internal/controller/cluster/eks/accesspolicyassociation"
	addon "github.com/upbound/provider-aws/internal/controller/cluster/eks/addon"
	clustereks "github.com/upbound/provider-aws/internal/controller/cluster/eks/cluster"
	clusterauth "github.com/upbound/provider-aws/internal/controller/cluster/eks/clusterauth"
	fargateprofile "github.com/upbound/provider-aws/internal/controller/cluster/eks/fargateprofile"
	identityproviderconfig "github.com/upbound/provider-aws/internal/controller/cluster/eks/identityproviderconfig"
	nodegroup "github.com/upbound/provider-aws/internal/controller/cluster/eks/nodegroup"
	podidentityassociation "github.com/upbound/provider-aws/internal/controller/cluster/eks/podidentityassociation"
	clusterelasticache "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/cluster"
	globalreplicationgroup "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/globalreplicationgroup"
	parametergroupelasticache "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/parametergroup"
	replicationgroup "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/replicationgroup"
	serverlesscache "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/serverlesscache"
	subnetgroupelasticache "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/subnetgroup"
	userelasticache "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/user"
	usergroupelasticache "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/usergroup"
	applicationelasticbeanstalk "github.com/upbound/provider-aws/internal/controller/cluster/elasticbeanstalk/application"
	applicationversion "github.com/upbound/provider-aws/internal/controller/cluster/elasticbeanstalk/applicationversion"
	configurationtemplate "github.com/upbound/provider-aws/internal/controller/cluster/elasticbeanstalk/configurationtemplate"
	domainelasticsearch "github.com/upbound/provider-aws/internal/controller/cluster/elasticsearch/domain"
	domainpolicy "github.com/upbound/provider-aws/internal/controller/cluster/elasticsearch/domainpolicy"
	domainsamloptions "github.com/upbound/provider-aws/internal/controller/cluster/elasticsearch/domainsamloptions"
	pipelineelastictranscoder "github.com/upbound/provider-aws/internal/controller/cluster/elastictranscoder/pipeline"
	preset "github.com/upbound/provider-aws/internal/controller/cluster/elastictranscoder/preset"
	appcookiestickinesspolicy "github.com/upbound/provider-aws/internal/controller/cluster/elb/appcookiestickinesspolicy"
	attachmentelb "github.com/upbound/provider-aws/internal/controller/cluster/elb/attachment"
	backendserverpolicy "github.com/upbound/provider-aws/internal/controller/cluster/elb/backendserverpolicy"
	elb "github.com/upbound/provider-aws/internal/controller/cluster/elb/elb"
	lbcookiestickinesspolicy "github.com/upbound/provider-aws/internal/controller/cluster/elb/lbcookiestickinesspolicy"
	lbsslnegotiationpolicy "github.com/upbound/provider-aws/internal/controller/cluster/elb/lbsslnegotiationpolicy"
	listenerpolicy "github.com/upbound/provider-aws/internal/controller/cluster/elb/listenerpolicy"
	policyelb "github.com/upbound/provider-aws/internal/controller/cluster/elb/policy"
	proxyprotocolpolicy "github.com/upbound/provider-aws/internal/controller/cluster/elb/proxyprotocolpolicy"
	lb "github.com/upbound/provider-aws/internal/controller/cluster/elbv2/lb"
	lblistener "github.com/upbound/provider-aws/internal/controller/cluster/elbv2/lblistener"
	lblistenercertificate "github.com/upbound/provider-aws/internal/controller/cluster/elbv2/lblistenercertificate"
	lblistenerrule "github.com/upbound/provider-aws/internal/controller/cluster/elbv2/lblistenerrule"
	lbtargetgroup "github.com/upbound/provider-aws/internal/controller/cluster/elbv2/lbtargetgroup"
	lbtargetgroupattachment "github.com/upbound/provider-aws/internal/controller/cluster/elbv2/lbtargetgroupattachment"
	lbtruststore "github.com/upbound/provider-aws/internal/controller/cluster/elbv2/lbtruststore"
	securityconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/emr/securityconfiguration"
	applicationemrserverless "github.com/upbound/provider-aws/internal/controller/cluster/emrserverless/application"
	feature "github.com/upbound/provider-aws/internal/controller/cluster/evidently/feature"
	projectevidently "github.com/upbound/provider-aws/internal/controller/cluster/evidently/project"
	segment "github.com/upbound/provider-aws/internal/controller/cluster/evidently/segment"
	deliverystream "github.com/upbound/provider-aws/internal/controller/cluster/firehose/deliverystream"
	experimenttemplate "github.com/upbound/provider-aws/internal/controller/cluster/fis/experimenttemplate"
	backup "github.com/upbound/provider-aws/internal/controller/cluster/fsx/backup"
	datarepositoryassociation "github.com/upbound/provider-aws/internal/controller/cluster/fsx/datarepositoryassociation"
	lustrefilesystem "github.com/upbound/provider-aws/internal/controller/cluster/fsx/lustrefilesystem"
	ontapfilesystem "github.com/upbound/provider-aws/internal/controller/cluster/fsx/ontapfilesystem"
	ontapstoragevirtualmachine "github.com/upbound/provider-aws/internal/controller/cluster/fsx/ontapstoragevirtualmachine"
	windowsfilesystem "github.com/upbound/provider-aws/internal/controller/cluster/fsx/windowsfilesystem"
	alias "github.com/upbound/provider-aws/internal/controller/cluster/gamelift/alias"
	build "github.com/upbound/provider-aws/internal/controller/cluster/gamelift/build"
	fleetgamelift "github.com/upbound/provider-aws/internal/controller/cluster/gamelift/fleet"
	gamesessionqueue "github.com/upbound/provider-aws/internal/controller/cluster/gamelift/gamesessionqueue"
	script "github.com/upbound/provider-aws/internal/controller/cluster/gamelift/script"
	vaultglacier "github.com/upbound/provider-aws/internal/controller/cluster/glacier/vault"
	vaultlock "github.com/upbound/provider-aws/internal/controller/cluster/glacier/vaultlock"
	accelerator "github.com/upbound/provider-aws/internal/controller/cluster/globalaccelerator/accelerator"
	endpointgroup "github.com/upbound/provider-aws/internal/controller/cluster/globalaccelerator/endpointgroup"
	listener "github.com/upbound/provider-aws/internal/controller/cluster/globalaccelerator/listener"
	catalogdatabase "github.com/upbound/provider-aws/internal/controller/cluster/glue/catalogdatabase"
	catalogtable "github.com/upbound/provider-aws/internal/controller/cluster/glue/catalogtable"
	catalogtableoptimizer "github.com/upbound/provider-aws/internal/controller/cluster/glue/catalogtableoptimizer"
	classifier "github.com/upbound/provider-aws/internal/controller/cluster/glue/classifier"
	connectionglue "github.com/upbound/provider-aws/internal/controller/cluster/glue/connection"
	crawler "github.com/upbound/provider-aws/internal/controller/cluster/glue/crawler"
	datacatalogencryptionsettings "github.com/upbound/provider-aws/internal/controller/cluster/glue/datacatalogencryptionsettings"
	job "github.com/upbound/provider-aws/internal/controller/cluster/glue/job"
	registry "github.com/upbound/provider-aws/internal/controller/cluster/glue/registry"
	resourcepolicyglue "github.com/upbound/provider-aws/internal/controller/cluster/glue/resourcepolicy"
	schema "github.com/upbound/provider-aws/internal/controller/cluster/glue/schema"
	securityconfigurationglue "github.com/upbound/provider-aws/internal/controller/cluster/glue/securityconfiguration"
	triggerglue "github.com/upbound/provider-aws/internal/controller/cluster/glue/trigger"
	userdefinedfunction "github.com/upbound/provider-aws/internal/controller/cluster/glue/userdefinedfunction"
	workflow "github.com/upbound/provider-aws/internal/controller/cluster/glue/workflow"
	licenseassociation "github.com/upbound/provider-aws/internal/controller/cluster/grafana/licenseassociation"
	roleassociation "github.com/upbound/provider-aws/internal/controller/cluster/grafana/roleassociation"
	workspacegrafana "github.com/upbound/provider-aws/internal/controller/cluster/grafana/workspace"
	workspaceapikey "github.com/upbound/provider-aws/internal/controller/cluster/grafana/workspaceapikey"
	workspacesamlconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/grafana/workspacesamlconfiguration"
	detector "github.com/upbound/provider-aws/internal/controller/cluster/guardduty/detector"
	filter "github.com/upbound/provider-aws/internal/controller/cluster/guardduty/filter"
	memberguardduty "github.com/upbound/provider-aws/internal/controller/cluster/guardduty/member"
	accesskey "github.com/upbound/provider-aws/internal/controller/cluster/iam/accesskey"
	accountalias "github.com/upbound/provider-aws/internal/controller/cluster/iam/accountalias"
	accountpasswordpolicy "github.com/upbound/provider-aws/internal/controller/cluster/iam/accountpasswordpolicy"
	groupiam "github.com/upbound/provider-aws/internal/controller/cluster/iam/group"
	groupmembership "github.com/upbound/provider-aws/internal/controller/cluster/iam/groupmembership"
	grouppolicyattachment "github.com/upbound/provider-aws/internal/controller/cluster/iam/grouppolicyattachment"
	instanceprofileiam "github.com/upbound/provider-aws/internal/controller/cluster/iam/instanceprofile"
	openidconnectprovider "github.com/upbound/provider-aws/internal/controller/cluster/iam/openidconnectprovider"
	policyiam "github.com/upbound/provider-aws/internal/controller/cluster/iam/policy"
	role "github.com/upbound/provider-aws/internal/controller/cluster/iam/role"
	rolepolicy "github.com/upbound/provider-aws/internal/controller/cluster/iam/rolepolicy"
	rolepolicyattachment "github.com/upbound/provider-aws/internal/controller/cluster/iam/rolepolicyattachment"
	samlprovider "github.com/upbound/provider-aws/internal/controller/cluster/iam/samlprovider"
	servercertificate "github.com/upbound/provider-aws/internal/controller/cluster/iam/servercertificate"
	servicelinkedrole "github.com/upbound/provider-aws/internal/controller/cluster/iam/servicelinkedrole"
	servicespecificcredential "github.com/upbound/provider-aws/internal/controller/cluster/iam/servicespecificcredential"
	signingcertificate "github.com/upbound/provider-aws/internal/controller/cluster/iam/signingcertificate"
	useriam "github.com/upbound/provider-aws/internal/controller/cluster/iam/user"
	usergroupmembership "github.com/upbound/provider-aws/internal/controller/cluster/iam/usergroupmembership"
	userloginprofile "github.com/upbound/provider-aws/internal/controller/cluster/iam/userloginprofile"
	userpolicyattachment "github.com/upbound/provider-aws/internal/controller/cluster/iam/userpolicyattachment"
	usersshkey "github.com/upbound/provider-aws/internal/controller/cluster/iam/usersshkey"
	virtualmfadevice "github.com/upbound/provider-aws/internal/controller/cluster/iam/virtualmfadevice"
	groupidentitystore "github.com/upbound/provider-aws/internal/controller/cluster/identitystore/group"
	groupmembershipidentitystore "github.com/upbound/provider-aws/internal/controller/cluster/identitystore/groupmembership"
	useridentitystore "github.com/upbound/provider-aws/internal/controller/cluster/identitystore/user"
	component "github.com/upbound/provider-aws/internal/controller/cluster/imagebuilder/component"
	containerrecipe "github.com/upbound/provider-aws/internal/controller/cluster/imagebuilder/containerrecipe"
	distributionconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/imagebuilder/distributionconfiguration"
	image "github.com/upbound/provider-aws/internal/controller/cluster/imagebuilder/image"
	imagepipeline "github.com/upbound/provider-aws/internal/controller/cluster/imagebuilder/imagepipeline"
	imagerecipe "github.com/upbound/provider-aws/internal/controller/cluster/imagebuilder/imagerecipe"
	infrastructureconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/imagebuilder/infrastructureconfiguration"
	assessmenttarget "github.com/upbound/provider-aws/internal/controller/cluster/inspector/assessmenttarget"
	assessmenttemplate "github.com/upbound/provider-aws/internal/controller/cluster/inspector/assessmenttemplate"
	resourcegroup "github.com/upbound/provider-aws/internal/controller/cluster/inspector/resourcegroup"
	enabler "github.com/upbound/provider-aws/internal/controller/cluster/inspector2/enabler"
	authorizeriot "github.com/upbound/provider-aws/internal/controller/cluster/iot/authorizer"
	certificateiot "github.com/upbound/provider-aws/internal/controller/cluster/iot/certificate"
	domainconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/iot/domainconfiguration"
	indexingconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/iot/indexingconfiguration"
	loggingoptions "github.com/upbound/provider-aws/internal/controller/cluster/iot/loggingoptions"
	policyiot "github.com/upbound/provider-aws/internal/controller/cluster/iot/policy"
	policyattachment "github.com/upbound/provider-aws/internal/controller/cluster/iot/policyattachment"
	provisioningtemplate "github.com/upbound/provider-aws/internal/controller/cluster/iot/provisioningtemplate"
	rolealias "github.com/upbound/provider-aws/internal/controller/cluster/iot/rolealias"
	thing "github.com/upbound/provider-aws/internal/controller/cluster/iot/thing"
	thinggroup "github.com/upbound/provider-aws/internal/controller/cluster/iot/thinggroup"
	thinggroupmembership "github.com/upbound/provider-aws/internal/controller/cluster/iot/thinggroupmembership"
	thingprincipalattachment "github.com/upbound/provider-aws/internal/controller/cluster/iot/thingprincipalattachment"
	thingtype "github.com/upbound/provider-aws/internal/controller/cluster/iot/thingtype"
	topicrule "github.com/upbound/provider-aws/internal/controller/cluster/iot/topicrule"
	topicruledestination "github.com/upbound/provider-aws/internal/controller/cluster/iot/topicruledestination"
	channel "github.com/upbound/provider-aws/internal/controller/cluster/ivs/channel"
	recordingconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/ivs/recordingconfiguration"
	clusterkafka "github.com/upbound/provider-aws/internal/controller/cluster/kafka/cluster"
	clusterpolicy "github.com/upbound/provider-aws/internal/controller/cluster/kafka/clusterpolicy"
	configuration "github.com/upbound/provider-aws/internal/controller/cluster/kafka/configuration"
	replicator "github.com/upbound/provider-aws/internal/controller/cluster/kafka/replicator"
	scramsecretassociation "github.com/upbound/provider-aws/internal/controller/cluster/kafka/scramsecretassociation"
	serverlesscluster "github.com/upbound/provider-aws/internal/controller/cluster/kafka/serverlesscluster"
	singlescramsecretassociation "github.com/upbound/provider-aws/internal/controller/cluster/kafka/singlescramsecretassociation"
	connector "github.com/upbound/provider-aws/internal/controller/cluster/kafkaconnect/connector"
	customplugin "github.com/upbound/provider-aws/internal/controller/cluster/kafkaconnect/customplugin"
	workerconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/kafkaconnect/workerconfiguration"
	datasourcekendra "github.com/upbound/provider-aws/internal/controller/cluster/kendra/datasource"
	experience "github.com/upbound/provider-aws/internal/controller/cluster/kendra/experience"
	index "github.com/upbound/provider-aws/internal/controller/cluster/kendra/index"
	querysuggestionsblocklist "github.com/upbound/provider-aws/internal/controller/cluster/kendra/querysuggestionsblocklist"
	thesaurus "github.com/upbound/provider-aws/internal/controller/cluster/kendra/thesaurus"
	keyspace "github.com/upbound/provider-aws/internal/controller/cluster/keyspaces/keyspace"
	tablekeyspaces "github.com/upbound/provider-aws/internal/controller/cluster/keyspaces/table"
	streamkinesis "github.com/upbound/provider-aws/internal/controller/cluster/kinesis/stream"
	streamconsumer "github.com/upbound/provider-aws/internal/controller/cluster/kinesis/streamconsumer"
	applicationkinesisanalytics "github.com/upbound/provider-aws/internal/controller/cluster/kinesisanalytics/application"
	applicationkinesisanalyticsv2 "github.com/upbound/provider-aws/internal/controller/cluster/kinesisanalyticsv2/application"
	applicationsnapshot "github.com/upbound/provider-aws/internal/controller/cluster/kinesisanalyticsv2/applicationsnapshot"
	streamkinesisvideo "github.com/upbound/provider-aws/internal/controller/cluster/kinesisvideo/stream"
	aliaskms "github.com/upbound/provider-aws/internal/controller/cluster/kms/alias"
	ciphertext "github.com/upbound/provider-aws/internal/controller/cluster/kms/ciphertext"
	externalkey "github.com/upbound/provider-aws/internal/controller/cluster/kms/externalkey"
	grant "github.com/upbound/provider-aws/internal/controller/cluster/kms/grant"
	key "github.com/upbound/provider-aws/internal/controller/cluster/kms/key"
	replicaexternalkey "github.com/upbound/provider-aws/internal/controller/cluster/kms/replicaexternalkey"
	replicakey "github.com/upbound/provider-aws/internal/controller/cluster/kms/replicakey"
	datalakesettings "github.com/upbound/provider-aws/internal/controller/cluster/lakeformation/datalakesettings"
	permissions "github.com/upbound/provider-aws/internal/controller/cluster/lakeformation/permissions"
	resourcelakeformation "github.com/upbound/provider-aws/internal/controller/cluster/lakeformation/resource"
	aliaslambda "github.com/upbound/provider-aws/internal/controller/cluster/lambda/alias"
	codesigningconfig "github.com/upbound/provider-aws/internal/controller/cluster/lambda/codesigningconfig"
	eventsourcemapping "github.com/upbound/provider-aws/internal/controller/cluster/lambda/eventsourcemapping"
	functionlambda "github.com/upbound/provider-aws/internal/controller/cluster/lambda/function"
	functioneventinvokeconfig "github.com/upbound/provider-aws/internal/controller/cluster/lambda/functioneventinvokeconfig"
	functionurl "github.com/upbound/provider-aws/internal/controller/cluster/lambda/functionurl"
	invocation "github.com/upbound/provider-aws/internal/controller/cluster/lambda/invocation"
	layerversion "github.com/upbound/provider-aws/internal/controller/cluster/lambda/layerversion"
	layerversionpermission "github.com/upbound/provider-aws/internal/controller/cluster/lambda/layerversionpermission"
	permissionlambda "github.com/upbound/provider-aws/internal/controller/cluster/lambda/permission"
	provisionedconcurrencyconfig "github.com/upbound/provider-aws/internal/controller/cluster/lambda/provisionedconcurrencyconfig"
	bot "github.com/upbound/provider-aws/internal/controller/cluster/lexmodels/bot"
	botalias "github.com/upbound/provider-aws/internal/controller/cluster/lexmodels/botalias"
	intent "github.com/upbound/provider-aws/internal/controller/cluster/lexmodels/intent"
	slottype "github.com/upbound/provider-aws/internal/controller/cluster/lexmodels/slottype"
	association "github.com/upbound/provider-aws/internal/controller/cluster/licensemanager/association"
	licenseconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/licensemanager/licenseconfiguration"
	bucket "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/bucket"
	certificatelightsail "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/certificate"
	containerservice "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/containerservice"
	disk "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/disk"
	diskattachment "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/diskattachment"
	domainlightsail "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/domain"
	domainentry "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/domainentry"
	instancelightsail "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/instance"
	instancepublicports "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/instancepublicports"
	keypairlightsail "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/keypair"
	lblightsail "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/lb"
	lbattachment "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/lbattachment"
	lbcertificate "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/lbcertificate"
	lbstickinesspolicy "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/lbstickinesspolicy"
	staticip "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/staticip"
	staticipattachment "github.com/upbound/provider-aws/internal/controller/cluster/lightsail/staticipattachment"
	geofencecollection "github.com/upbound/provider-aws/internal/controller/cluster/location/geofencecollection"
	placeindex "github.com/upbound/provider-aws/internal/controller/cluster/location/placeindex"
	routecalculator "github.com/upbound/provider-aws/internal/controller/cluster/location/routecalculator"
	tracker "github.com/upbound/provider-aws/internal/controller/cluster/location/tracker"
	trackerassociation "github.com/upbound/provider-aws/internal/controller/cluster/location/trackerassociation"
	accountmacie2 "github.com/upbound/provider-aws/internal/controller/cluster/macie2/account"
	classificationjob "github.com/upbound/provider-aws/internal/controller/cluster/macie2/classificationjob"
	customdataidentifier "github.com/upbound/provider-aws/internal/controller/cluster/macie2/customdataidentifier"
	findingsfilter "github.com/upbound/provider-aws/internal/controller/cluster/macie2/findingsfilter"
	invitationacceptermacie2 "github.com/upbound/provider-aws/internal/controller/cluster/macie2/invitationaccepter"
	membermacie2 "github.com/upbound/provider-aws/internal/controller/cluster/macie2/member"
	queuemediaconvert "github.com/upbound/provider-aws/internal/controller/cluster/mediaconvert/queue"
	channelmedialive "github.com/upbound/provider-aws/internal/controller/cluster/medialive/channel"
	input "github.com/upbound/provider-aws/internal/controller/cluster/medialive/input"
	inputsecuritygroup "github.com/upbound/provider-aws/internal/controller/cluster/medialive/inputsecuritygroup"
	multiplex "github.com/upbound/provider-aws/internal/controller/cluster/medialive/multiplex"
	channelmediapackage "github.com/upbound/provider-aws/internal/controller/cluster/mediapackage/channel"
	container "github.com/upbound/provider-aws/internal/controller/cluster/mediastore/container"
	containerpolicy "github.com/upbound/provider-aws/internal/controller/cluster/mediastore/containerpolicy"
	acl "github.com/upbound/provider-aws/internal/controller/cluster/memorydb/acl"
	clustermemorydb "github.com/upbound/provider-aws/internal/controller/cluster/memorydb/cluster"
	parametergroupmemorydb "github.com/upbound/provider-aws/internal/controller/cluster/memorydb/parametergroup"
	snapshot "github.com/upbound/provider-aws/internal/controller/cluster/memorydb/snapshot"
	subnetgroupmemorydb "github.com/upbound/provider-aws/internal/controller/cluster/memorydb/subnetgroup"
	usermemorydb "github.com/upbound/provider-aws/internal/controller/cluster/memorydb/user"
	broker "github.com/upbound/provider-aws/internal/controller/cluster/mq/broker"
	configurationmq "github.com/upbound/provider-aws/internal/controller/cluster/mq/configuration"
	usermq "github.com/upbound/provider-aws/internal/controller/cluster/mq/user"
	environmentmwaa "github.com/upbound/provider-aws/internal/controller/cluster/mwaa/environment"
	clusterneptune "github.com/upbound/provider-aws/internal/controller/cluster/neptune/cluster"
	clusterendpoint "github.com/upbound/provider-aws/internal/controller/cluster/neptune/clusterendpoint"
	clusterinstanceneptune "github.com/upbound/provider-aws/internal/controller/cluster/neptune/clusterinstance"
	clusterparametergroupneptune "github.com/upbound/provider-aws/internal/controller/cluster/neptune/clusterparametergroup"
	clustersnapshotneptune "github.com/upbound/provider-aws/internal/controller/cluster/neptune/clustersnapshot"
	eventsubscriptionneptune "github.com/upbound/provider-aws/internal/controller/cluster/neptune/eventsubscription"
	globalclusterneptune "github.com/upbound/provider-aws/internal/controller/cluster/neptune/globalcluster"
	parametergroupneptune "github.com/upbound/provider-aws/internal/controller/cluster/neptune/parametergroup"
	subnetgroupneptune "github.com/upbound/provider-aws/internal/controller/cluster/neptune/subnetgroup"
	firewall "github.com/upbound/provider-aws/internal/controller/cluster/networkfirewall/firewall"
	firewallpolicy "github.com/upbound/provider-aws/internal/controller/cluster/networkfirewall/firewallpolicy"
	loggingconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/networkfirewall/loggingconfiguration"
	rulegroup "github.com/upbound/provider-aws/internal/controller/cluster/networkfirewall/rulegroup"
	attachmentaccepter "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/attachmentaccepter"
	connectattachment "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/connectattachment"
	connectionnetworkmanager "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/connection"
	corenetwork "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/corenetwork"
	customergatewayassociation "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/customergatewayassociation"
	device "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/device"
	globalnetwork "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/globalnetwork"
	link "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/link"
	linkassociation "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/linkassociation"
	site "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/site"
	transitgatewayconnectpeerassociation "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/transitgatewayconnectpeerassociation"
	transitgatewayregistration "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/transitgatewayregistration"
	vpcattachment "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/vpcattachment"
	domainopensearch "github.com/upbound/provider-aws/internal/controller/cluster/opensearch/domain"
	domainpolicyopensearch "github.com/upbound/provider-aws/internal/controller/cluster/opensearch/domainpolicy"
	domainsamloptionsopensearch "github.com/upbound/provider-aws/internal/controller/cluster/opensearch/domainsamloptions"
	accesspolicy "github.com/upbound/provider-aws/internal/controller/cluster/opensearchserverless/accesspolicy"
	collection "github.com/upbound/provider-aws/internal/controller/cluster/opensearchserverless/collection"
	lifecyclepolicyopensearchserverless "github.com/upbound/provider-aws/internal/controller/cluster/opensearchserverless/lifecyclepolicy"
	securityconfig "github.com/upbound/provider-aws/internal/controller/cluster/opensearchserverless/securityconfig"
	securitypolicy "github.com/upbound/provider-aws/internal/controller/cluster/opensearchserverless/securitypolicy"
	vpcendpointopensearchserverless "github.com/upbound/provider-aws/internal/controller/cluster/opensearchserverless/vpcendpoint"
	accountorganizations "github.com/upbound/provider-aws/internal/controller/cluster/organizations/account"
	delegatedadministrator "github.com/upbound/provider-aws/internal/controller/cluster/organizations/delegatedadministrator"
	organization "github.com/upbound/provider-aws/internal/controller/cluster/organizations/organization"
	organizationalunit "github.com/upbound/provider-aws/internal/controller/cluster/organizations/organizationalunit"
	policyorganizations "github.com/upbound/provider-aws/internal/controller/cluster/organizations/policy"
	policyattachmentorganizations "github.com/upbound/provider-aws/internal/controller/cluster/organizations/policyattachment"
	pipelineosis "github.com/upbound/provider-aws/internal/controller/cluster/osis/pipeline"
	apppinpoint "github.com/upbound/provider-aws/internal/controller/cluster/pinpoint/app"
	smschannel "github.com/upbound/provider-aws/internal/controller/cluster/pinpoint/smschannel"
	pipe "github.com/upbound/provider-aws/internal/controller/cluster/pipes/pipe"
	providerconfig "github.com/upbound/provider-aws/internal/controller/cluster/providerconfig"
	ledger "github.com/upbound/provider-aws/internal/controller/cluster/qldb/ledger"
	streamqldb "github.com/upbound/provider-aws/internal/controller/cluster/qldb/stream"
	groupquicksight "github.com/upbound/provider-aws/internal/controller/cluster/quicksight/group"
	userquicksight "github.com/upbound/provider-aws/internal/controller/cluster/quicksight/user"
	principalassociation "github.com/upbound/provider-aws/internal/controller/cluster/ram/principalassociation"
	resourceassociation "github.com/upbound/provider-aws/internal/controller/cluster/ram/resourceassociation"
	resourceshare "github.com/upbound/provider-aws/internal/controller/cluster/ram/resourceshare"
	resourceshareaccepter "github.com/upbound/provider-aws/internal/controller/cluster/ram/resourceshareaccepter"
	clusterrds "github.com/upbound/provider-aws/internal/controller/cluster/rds/cluster"
	clusteractivitystream "github.com/upbound/provider-aws/internal/controller/cluster/rds/clusteractivitystream"
	clusterendpointrds "github.com/upbound/provider-aws/internal/controller/cluster/rds/clusterendpoint"
	clusterinstancerds "github.com/upbound/provider-aws/internal/controller/cluster/rds/clusterinstance"
	clusterparametergrouprds "github.com/upbound/provider-aws/internal/controller/cluster/rds/clusterparametergroup"
	clusterroleassociation "github.com/upbound/provider-aws/internal/controller/cluster/rds/clusterroleassociation"
	clustersnapshotrds "github.com/upbound/provider-aws/internal/controller/cluster/rds/clustersnapshot"
	dbinstanceautomatedbackupsreplication "github.com/upbound/provider-aws/internal/controller/cluster/rds/dbinstanceautomatedbackupsreplication"
	dbsnapshotcopy "github.com/upbound/provider-aws/internal/controller/cluster/rds/dbsnapshotcopy"
	eventsubscriptionrds "github.com/upbound/provider-aws/internal/controller/cluster/rds/eventsubscription"
	globalclusterrds "github.com/upbound/provider-aws/internal/controller/cluster/rds/globalcluster"
	instancerds "github.com/upbound/provider-aws/internal/controller/cluster/rds/instance"
	instanceroleassociation "github.com/upbound/provider-aws/internal/controller/cluster/rds/instanceroleassociation"
	instancestaterds "github.com/upbound/provider-aws/internal/controller/cluster/rds/instancestate"
	optiongroup "github.com/upbound/provider-aws/internal/controller/cluster/rds/optiongroup"
	parametergrouprds "github.com/upbound/provider-aws/internal/controller/cluster/rds/parametergroup"
	proxy "github.com/upbound/provider-aws/internal/controller/cluster/rds/proxy"
	proxydefaulttargetgroup "github.com/upbound/provider-aws/internal/controller/cluster/rds/proxydefaulttargetgroup"
	proxyendpoint "github.com/upbound/provider-aws/internal/controller/cluster/rds/proxyendpoint"
	proxytarget "github.com/upbound/provider-aws/internal/controller/cluster/rds/proxytarget"
	snapshotrds "github.com/upbound/provider-aws/internal/controller/cluster/rds/snapshot"
	subnetgrouprds "github.com/upbound/provider-aws/internal/controller/cluster/rds/subnetgroup"
	authenticationprofile "github.com/upbound/provider-aws/internal/controller/cluster/redshift/authenticationprofile"
	clusterredshift "github.com/upbound/provider-aws/internal/controller/cluster/redshift/cluster"
	endpointaccess "github.com/upbound/provider-aws/internal/controller/cluster/redshift/endpointaccess"
	eventsubscriptionredshift "github.com/upbound/provider-aws/internal/controller/cluster/redshift/eventsubscription"
	hsmclientcertificate "github.com/upbound/provider-aws/internal/controller/cluster/redshift/hsmclientcertificate"
	hsmconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/redshift/hsmconfiguration"
	parametergroupredshift "github.com/upbound/provider-aws/internal/controller/cluster/redshift/parametergroup"
	scheduledactionredshift "github.com/upbound/provider-aws/internal/controller/cluster/redshift/scheduledaction"
	snapshotcopygrant "github.com/upbound/provider-aws/internal/controller/cluster/redshift/snapshotcopygrant"
	snapshotschedule "github.com/upbound/provider-aws/internal/controller/cluster/redshift/snapshotschedule"
	snapshotscheduleassociation "github.com/upbound/provider-aws/internal/controller/cluster/redshift/snapshotscheduleassociation"
	subnetgroupredshift "github.com/upbound/provider-aws/internal/controller/cluster/redshift/subnetgroup"
	usagelimit "github.com/upbound/provider-aws/internal/controller/cluster/redshift/usagelimit"
	endpointaccessredshiftserverless "github.com/upbound/provider-aws/internal/controller/cluster/redshiftserverless/endpointaccess"
	redshiftserverlessnamespace "github.com/upbound/provider-aws/internal/controller/cluster/redshiftserverless/redshiftserverlessnamespace"
	resourcepolicyredshiftserverless "github.com/upbound/provider-aws/internal/controller/cluster/redshiftserverless/resourcepolicy"
	snapshotredshiftserverless "github.com/upbound/provider-aws/internal/controller/cluster/redshiftserverless/snapshot"
	usagelimitredshiftserverless "github.com/upbound/provider-aws/internal/controller/cluster/redshiftserverless/usagelimit"
	workgroupredshiftserverless "github.com/upbound/provider-aws/internal/controller/cluster/redshiftserverless/workgroup"
	groupresourcegroups "github.com/upbound/provider-aws/internal/controller/cluster/resourcegroups/group"
	profile "github.com/upbound/provider-aws/internal/controller/cluster/rolesanywhere/profile"
	delegationset "github.com/upbound/provider-aws/internal/controller/cluster/route53/delegationset"
	healthcheck "github.com/upbound/provider-aws/internal/controller/cluster/route53/healthcheck"
	hostedzonednssec "github.com/upbound/provider-aws/internal/controller/cluster/route53/hostedzonednssec"
	querylog "github.com/upbound/provider-aws/internal/controller/cluster/route53/querylog"
	record "github.com/upbound/provider-aws/internal/controller/cluster/route53/record"
	resolverconfig "github.com/upbound/provider-aws/internal/controller/cluster/route53/resolverconfig"
	trafficpolicy "github.com/upbound/provider-aws/internal/controller/cluster/route53/trafficpolicy"
	trafficpolicyinstance "github.com/upbound/provider-aws/internal/controller/cluster/route53/trafficpolicyinstance"
	vpcassociationauthorization "github.com/upbound/provider-aws/internal/controller/cluster/route53/vpcassociationauthorization"
	zone "github.com/upbound/provider-aws/internal/controller/cluster/route53/zone"
	zoneassociation "github.com/upbound/provider-aws/internal/controller/cluster/route53/zoneassociation"
	clusterroute53recoverycontrolconfig "github.com/upbound/provider-aws/internal/controller/cluster/route53recoverycontrolconfig/cluster"
	controlpanel "github.com/upbound/provider-aws/internal/controller/cluster/route53recoverycontrolconfig/controlpanel"
	routingcontrol "github.com/upbound/provider-aws/internal/controller/cluster/route53recoverycontrolconfig/routingcontrol"
	safetyrule "github.com/upbound/provider-aws/internal/controller/cluster/route53recoverycontrolconfig/safetyrule"
	cell "github.com/upbound/provider-aws/internal/controller/cluster/route53recoveryreadiness/cell"
	readinesscheck "github.com/upbound/provider-aws/internal/controller/cluster/route53recoveryreadiness/readinesscheck"
	recoverygroup "github.com/upbound/provider-aws/internal/controller/cluster/route53recoveryreadiness/recoverygroup"
	resourceset "github.com/upbound/provider-aws/internal/controller/cluster/route53recoveryreadiness/resourceset"
	endpointroute53resolver "github.com/upbound/provider-aws/internal/controller/cluster/route53resolver/endpoint"
	ruleroute53resolver "github.com/upbound/provider-aws/internal/controller/cluster/route53resolver/rule"
	ruleassociation "github.com/upbound/provider-aws/internal/controller/cluster/route53resolver/ruleassociation"
	appmonitor "github.com/upbound/provider-aws/internal/controller/cluster/rum/appmonitor"
	metricsdestination "github.com/upbound/provider-aws/internal/controller/cluster/rum/metricsdestination"
	buckets3 "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucket"
	bucketaccelerateconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketaccelerateconfiguration"
	bucketacl "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketacl"
	bucketanalyticsconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketanalyticsconfiguration"
	bucketcorsconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketcorsconfiguration"
	bucketintelligenttieringconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketintelligenttieringconfiguration"
	bucketinventory "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketinventory"
	bucketlifecycleconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketlifecycleconfiguration"
	bucketlogging "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketlogging"
	bucketmetric "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketmetric"
	bucketnotification "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketnotification"
	bucketobject "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketobject"
	bucketobjectlockconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketobjectlockconfiguration"
	bucketownershipcontrols "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketownershipcontrols"
	bucketpolicy "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketpolicy"
	bucketpublicaccessblock "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketpublicaccessblock"
	bucketreplicationconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketreplicationconfiguration"
	bucketrequestpaymentconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketrequestpaymentconfiguration"
	bucketserversideencryptionconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketserversideencryptionconfiguration"
	bucketversioning "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketversioning"
	bucketwebsiteconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/s3/bucketwebsiteconfiguration"
	directorybucket "github.com/upbound/provider-aws/internal/controller/cluster/s3/directorybucket"
	object "github.com/upbound/provider-aws/internal/controller/cluster/s3/object"
	objectcopy "github.com/upbound/provider-aws/internal/controller/cluster/s3/objectcopy"
	accesspoints3control "github.com/upbound/provider-aws/internal/controller/cluster/s3control/accesspoint"
	accesspointpolicy "github.com/upbound/provider-aws/internal/controller/cluster/s3control/accesspointpolicy"
	accountpublicaccessblock "github.com/upbound/provider-aws/internal/controller/cluster/s3control/accountpublicaccessblock"
	multiregionaccesspoint "github.com/upbound/provider-aws/internal/controller/cluster/s3control/multiregionaccesspoint"
	multiregionaccesspointpolicy "github.com/upbound/provider-aws/internal/controller/cluster/s3control/multiregionaccesspointpolicy"
	objectlambdaaccesspoint "github.com/upbound/provider-aws/internal/controller/cluster/s3control/objectlambdaaccesspoint"
	objectlambdaaccesspointpolicy "github.com/upbound/provider-aws/internal/controller/cluster/s3control/objectlambdaaccesspointpolicy"
	storagelensconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/s3control/storagelensconfiguration"
	appsagemaker "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/app"
	appimageconfig "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/appimageconfig"
	coderepository "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/coderepository"
	devicesagemaker "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/device"
	devicefleet "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/devicefleet"
	domainsagemaker "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/domain"
	endpointsagemaker "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/endpoint"
	endpointconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/endpointconfiguration"
	featuregroup "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/featuregroup"
	imagesagemaker "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/image"
	imageversion "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/imageversion"
	mlflowtrackingserver "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/mlflowtrackingserver"
	modelsagemaker "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/model"
	modelpackagegroup "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/modelpackagegroup"
	modelpackagegrouppolicy "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/modelpackagegrouppolicy"
	notebookinstance "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/notebookinstance"
	notebookinstancelifecycleconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/notebookinstancelifecycleconfiguration"
	servicecatalogportfoliostatus "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/servicecatalogportfoliostatus"
	space "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/space"
	studiolifecycleconfig "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/studiolifecycleconfig"
	userprofile "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/userprofile"
	workforce "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/workforce"
	workteam "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/workteam"
	schedulescheduler "github.com/upbound/provider-aws/internal/controller/cluster/scheduler/schedule"
	schedulegroup "github.com/upbound/provider-aws/internal/controller/cluster/scheduler/schedulegroup"
	discoverer "github.com/upbound/provider-aws/internal/controller/cluster/schemas/discoverer"
	registryschemas "github.com/upbound/provider-aws/internal/controller/cluster/schemas/registry"
	schemaschemas "github.com/upbound/provider-aws/internal/controller/cluster/schemas/schema"
	secret "github.com/upbound/provider-aws/internal/controller/cluster/secretsmanager/secret"
	secretpolicy "github.com/upbound/provider-aws/internal/controller/cluster/secretsmanager/secretpolicy"
	secretrotation "github.com/upbound/provider-aws/internal/controller/cluster/secretsmanager/secretrotation"
	secretversion "github.com/upbound/provider-aws/internal/controller/cluster/secretsmanager/secretversion"
	accountsecurityhub "github.com/upbound/provider-aws/internal/controller/cluster/securityhub/account"
	actiontarget "github.com/upbound/provider-aws/internal/controller/cluster/securityhub/actiontarget"
	findingaggregator "github.com/upbound/provider-aws/internal/controller/cluster/securityhub/findingaggregator"
	insight "github.com/upbound/provider-aws/internal/controller/cluster/securityhub/insight"
	inviteaccepter "github.com/upbound/provider-aws/internal/controller/cluster/securityhub/inviteaccepter"
	membersecurityhub "github.com/upbound/provider-aws/internal/controller/cluster/securityhub/member"
	productsubscription "github.com/upbound/provider-aws/internal/controller/cluster/securityhub/productsubscription"
	standardssubscription "github.com/upbound/provider-aws/internal/controller/cluster/securityhub/standardssubscription"
	cloudformationstack "github.com/upbound/provider-aws/internal/controller/cluster/serverlessrepo/cloudformationstack"
	budgetresourceassociation "github.com/upbound/provider-aws/internal/controller/cluster/servicecatalog/budgetresourceassociation"
	constraint "github.com/upbound/provider-aws/internal/controller/cluster/servicecatalog/constraint"
	portfolio "github.com/upbound/provider-aws/internal/controller/cluster/servicecatalog/portfolio"
	portfolioshare "github.com/upbound/provider-aws/internal/controller/cluster/servicecatalog/portfolioshare"
	principalportfolioassociation "github.com/upbound/provider-aws/internal/controller/cluster/servicecatalog/principalportfolioassociation"
	product "github.com/upbound/provider-aws/internal/controller/cluster/servicecatalog/product"
	productportfolioassociation "github.com/upbound/provider-aws/internal/controller/cluster/servicecatalog/productportfolioassociation"
	provisioningartifact "github.com/upbound/provider-aws/internal/controller/cluster/servicecatalog/provisioningartifact"
	serviceaction "github.com/upbound/provider-aws/internal/controller/cluster/servicecatalog/serviceaction"
	tagoption "github.com/upbound/provider-aws/internal/controller/cluster/servicecatalog/tagoption"
	tagoptionresourceassociation "github.com/upbound/provider-aws/internal/controller/cluster/servicecatalog/tagoptionresourceassociation"
	httpnamespace "github.com/upbound/provider-aws/internal/controller/cluster/servicediscovery/httpnamespace"
	privatednsnamespace "github.com/upbound/provider-aws/internal/controller/cluster/servicediscovery/privatednsnamespace"
	publicdnsnamespace "github.com/upbound/provider-aws/internal/controller/cluster/servicediscovery/publicdnsnamespace"
	serviceservicediscovery "github.com/upbound/provider-aws/internal/controller/cluster/servicediscovery/service"
	servicequota "github.com/upbound/provider-aws/internal/controller/cluster/servicequotas/servicequota"
	activereceiptruleset "github.com/upbound/provider-aws/internal/controller/cluster/ses/activereceiptruleset"
	configurationset "github.com/upbound/provider-aws/internal/controller/cluster/ses/configurationset"
	domaindkim "github.com/upbound/provider-aws/internal/controller/cluster/ses/domaindkim"
	domainidentity "github.com/upbound/provider-aws/internal/controller/cluster/ses/domainidentity"
	domainmailfrom "github.com/upbound/provider-aws/internal/controller/cluster/ses/domainmailfrom"
	emailidentity "github.com/upbound/provider-aws/internal/controller/cluster/ses/emailidentity"
	eventdestination "github.com/upbound/provider-aws/internal/controller/cluster/ses/eventdestination"
	identitynotificationtopic "github.com/upbound/provider-aws/internal/controller/cluster/ses/identitynotificationtopic"
	identitypolicy "github.com/upbound/provider-aws/internal/controller/cluster/ses/identitypolicy"
	receiptfilter "github.com/upbound/provider-aws/internal/controller/cluster/ses/receiptfilter"
	receiptrule "github.com/upbound/provider-aws/internal/controller/cluster/ses/receiptrule"
	receiptruleset "github.com/upbound/provider-aws/internal/controller/cluster/ses/receiptruleset"
	template "github.com/upbound/provider-aws/internal/controller/cluster/ses/template"
	configurationsetsesv2 "github.com/upbound/provider-aws/internal/controller/cluster/sesv2/configurationset"
	configurationseteventdestination "github.com/upbound/provider-aws/internal/controller/cluster/sesv2/configurationseteventdestination"
	dedicatedippool "github.com/upbound/provider-aws/internal/controller/cluster/sesv2/dedicatedippool"
	emailidentitysesv2 "github.com/upbound/provider-aws/internal/controller/cluster/sesv2/emailidentity"
	emailidentityfeedbackattributes "github.com/upbound/provider-aws/internal/controller/cluster/sesv2/emailidentityfeedbackattributes"
	emailidentitymailfromattributes "github.com/upbound/provider-aws/internal/controller/cluster/sesv2/emailidentitymailfromattributes"
	activity "github.com/upbound/provider-aws/internal/controller/cluster/sfn/activity"
	statemachine "github.com/upbound/provider-aws/internal/controller/cluster/sfn/statemachine"
	signingjob "github.com/upbound/provider-aws/internal/controller/cluster/signer/signingjob"
	signingprofile "github.com/upbound/provider-aws/internal/controller/cluster/signer/signingprofile"
	signingprofilepermission "github.com/upbound/provider-aws/internal/controller/cluster/signer/signingprofilepermission"
	platformapplication "github.com/upbound/provider-aws/internal/controller/cluster/sns/platformapplication"
	smspreferences "github.com/upbound/provider-aws/internal/controller/cluster/sns/smspreferences"
	topic "github.com/upbound/provider-aws/internal/controller/cluster/sns/topic"
	topicpolicy "github.com/upbound/provider-aws/internal/controller/cluster/sns/topicpolicy"
	topicsubscription "github.com/upbound/provider-aws/internal/controller/cluster/sns/topicsubscription"
	queuesqs "github.com/upbound/provider-aws/internal/controller/cluster/sqs/queue"
	queuepolicy "github.com/upbound/provider-aws/internal/controller/cluster/sqs/queuepolicy"
	queueredriveallowpolicy "github.com/upbound/provider-aws/internal/controller/cluster/sqs/queueredriveallowpolicy"
	queueredrivepolicy "github.com/upbound/provider-aws/internal/controller/cluster/sqs/queueredrivepolicy"
	activation "github.com/upbound/provider-aws/internal/controller/cluster/ssm/activation"
	associationssm "github.com/upbound/provider-aws/internal/controller/cluster/ssm/association"
	defaultpatchbaseline "github.com/upbound/provider-aws/internal/controller/cluster/ssm/defaultpatchbaseline"
	document "github.com/upbound/provider-aws/internal/controller/cluster/ssm/document"
	maintenancewindow "github.com/upbound/provider-aws/internal/controller/cluster/ssm/maintenancewindow"
	maintenancewindowtarget "github.com/upbound/provider-aws/internal/controller/cluster/ssm/maintenancewindowtarget"
	maintenancewindowtask "github.com/upbound/provider-aws/internal/controller/cluster/ssm/maintenancewindowtask"
	parameter "github.com/upbound/provider-aws/internal/controller/cluster/ssm/parameter"
	patchbaseline "github.com/upbound/provider-aws/internal/controller/cluster/ssm/patchbaseline"
	patchgroup "github.com/upbound/provider-aws/internal/controller/cluster/ssm/patchgroup"
	resourcedatasync "github.com/upbound/provider-aws/internal/controller/cluster/ssm/resourcedatasync"
	servicesetting "github.com/upbound/provider-aws/internal/controller/cluster/ssm/servicesetting"
	accountassignment "github.com/upbound/provider-aws/internal/controller/cluster/ssoadmin/accountassignment"
	customermanagedpolicyattachment "github.com/upbound/provider-aws/internal/controller/cluster/ssoadmin/customermanagedpolicyattachment"
	instanceaccesscontrolattributes "github.com/upbound/provider-aws/internal/controller/cluster/ssoadmin/instanceaccesscontrolattributes"
	managedpolicyattachment "github.com/upbound/provider-aws/internal/controller/cluster/ssoadmin/managedpolicyattachment"
	permissionsboundaryattachment "github.com/upbound/provider-aws/internal/controller/cluster/ssoadmin/permissionsboundaryattachment"
	permissionset "github.com/upbound/provider-aws/internal/controller/cluster/ssoadmin/permissionset"
	permissionsetinlinepolicy "github.com/upbound/provider-aws/internal/controller/cluster/ssoadmin/permissionsetinlinepolicy"
	domainswf "github.com/upbound/provider-aws/internal/controller/cluster/swf/domain"
	databasetimestreamwrite "github.com/upbound/provider-aws/internal/controller/cluster/timestreamwrite/database"
	tabletimestreamwrite "github.com/upbound/provider-aws/internal/controller/cluster/timestreamwrite/table"
	languagemodel "github.com/upbound/provider-aws/internal/controller/cluster/transcribe/languagemodel"
	vocabularytranscribe "github.com/upbound/provider-aws/internal/controller/cluster/transcribe/vocabulary"
	vocabularyfilter "github.com/upbound/provider-aws/internal/controller/cluster/transcribe/vocabularyfilter"
	connectortransfer "github.com/upbound/provider-aws/internal/controller/cluster/transfer/connector"
	server "github.com/upbound/provider-aws/internal/controller/cluster/transfer/server"
	sshkey "github.com/upbound/provider-aws/internal/controller/cluster/transfer/sshkey"
	tagtransfer "github.com/upbound/provider-aws/internal/controller/cluster/transfer/tag"
	usertransfer "github.com/upbound/provider-aws/internal/controller/cluster/transfer/user"
	workflowtransfer "github.com/upbound/provider-aws/internal/controller/cluster/transfer/workflow"
	endpointverifiedaccess "github.com/upbound/provider-aws/internal/controller/cluster/verifiedaccess/endpoint"
	groupverifiedaccess "github.com/upbound/provider-aws/internal/controller/cluster/verifiedaccess/group"
	instanceverifiedaccess "github.com/upbound/provider-aws/internal/controller/cluster/verifiedaccess/instance"
	instanceloggingconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/verifiedaccess/instanceloggingconfiguration"
	instancetrustproviderattachment "github.com/upbound/provider-aws/internal/controller/cluster/verifiedaccess/instancetrustproviderattachment"
	trustprovider "github.com/upbound/provider-aws/internal/controller/cluster/verifiedaccess/trustprovider"
	networkperformancemetricsubscription "github.com/upbound/provider-aws/internal/controller/cluster/vpc/networkperformancemetricsubscription"
	accesslogsubscription "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/accesslogsubscription"
	authpolicy "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/authpolicy"
	listenervpclattice "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/listener"
	listenerrule "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/listenerrule"
	resourceconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/resourceconfiguration"
	resourcegateway "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/resourcegateway"
	resourcepolicyvpclattice "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/resourcepolicy"
	servicevpclattice "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/service"
	servicenetwork "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/servicenetwork"
	servicenetworkresourceassociation "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/servicenetworkresourceassociation"
	servicenetworkserviceassociation "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/servicenetworkserviceassociation"
	servicenetworkvpcassociation "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/servicenetworkvpcassociation"
	targetgroup "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/targetgroup"
	targetgroupattachment "github.com/upbound/provider-aws/internal/controller/cluster/vpclattice/targetgroupattachment"
	bytematchset "github.com/upbound/provider-aws/internal/controller/cluster/waf/bytematchset"
	geomatchset "github.com/upbound/provider-aws/internal/controller/cluster/waf/geomatchset"
	ipset "github.com/upbound/provider-aws/internal/controller/cluster/waf/ipset"
	ratebasedrule "github.com/upbound/provider-aws/internal/controller/cluster/waf/ratebasedrule"
	regexmatchset "github.com/upbound/provider-aws/internal/controller/cluster/waf/regexmatchset"
	regexpatternset "github.com/upbound/provider-aws/internal/controller/cluster/waf/regexpatternset"
	rulewaf "github.com/upbound/provider-aws/internal/controller/cluster/waf/rule"
	sizeconstraintset "github.com/upbound/provider-aws/internal/controller/cluster/waf/sizeconstraintset"
	sqlinjectionmatchset "github.com/upbound/provider-aws/internal/controller/cluster/waf/sqlinjectionmatchset"
	webacl "github.com/upbound/provider-aws/internal/controller/cluster/waf/webacl"
	xssmatchset "github.com/upbound/provider-aws/internal/controller/cluster/waf/xssmatchset"
	bytematchsetwafregional "github.com/upbound/provider-aws/internal/controller/cluster/wafregional/bytematchset"
	geomatchsetwafregional "github.com/upbound/provider-aws/internal/controller/cluster/wafregional/geomatchset"
	ipsetwafregional "github.com/upbound/provider-aws/internal/controller/cluster/wafregional/ipset"
	ratebasedrulewafregional "github.com/upbound/provider-aws/internal/controller/cluster/wafregional/ratebasedrule"
	regexmatchsetwafregional "github.com/upbound/provider-aws/internal/controller/cluster/wafregional/regexmatchset"
	regexpatternsetwafregional "github.com/upbound/provider-aws/internal/controller/cluster/wafregional/regexpatternset"
	rulewafregional "github.com/upbound/provider-aws/internal/controller/cluster/wafregional/rule"
	sizeconstraintsetwafregional "github.com/upbound/provider-aws/internal/controller/cluster/wafregional/sizeconstraintset"
	sqlinjectionmatchsetwafregional "github.com/upbound/provider-aws/internal/controller/cluster/wafregional/sqlinjectionmatchset"
	webaclwafregional "github.com/upbound/provider-aws/internal/controller/cluster/wafregional/webacl"
	xssmatchsetwafregional "github.com/upbound/provider-aws/internal/controller/cluster/wafregional/xssmatchset"
	ipsetwafv2 "github.com/upbound/provider-aws/internal/controller/cluster/wafv2/ipset"
	regexpatternsetwafv2 "github.com/upbound/provider-aws/internal/controller/cluster/wafv2/regexpatternset"
	rulegroupwafv2 "github.com/upbound/provider-aws/internal/controller/cluster/wafv2/rulegroup"
	webaclwafv2 "github.com/upbound/provider-aws/internal/controller/cluster/wafv2/webacl"
	webaclassociation "github.com/upbound/provider-aws/internal/controller/cluster/wafv2/webaclassociation"
	webaclloggingconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/wafv2/webaclloggingconfiguration"
	directoryworkspaces "github.com/upbound/provider-aws/internal/controller/cluster/workspaces/directory"
	ipgroup "github.com/upbound/provider-aws/internal/controller/cluster/workspaces/ipgroup"
	encryptionconfig "github.com/upbound/provider-aws/internal/controller/cluster/xray/encryptionconfig"
	groupxray "github.com/upbound/provider-aws/internal/controller/cluster/xray/group"
	samplingrule "github.com/upbound/provider-aws/internal/controller/cluster/xray/samplingrule"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyzer.Setup,
		archiverule.Setup,
		alternatecontact.Setup,
		region.Setup,
		certificate.Setup,
		certificatevalidation.Setup,
		certificateacmpca.Setup,
		certificateauthority.Setup,
		certificateauthoritycertificate.Setup,
		permission.Setup,
		policy.Setup,
		alertmanagerdefinition.Setup,
		rulegroupnamespace.Setup,
		scraper.Setup,
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
		computeenvironment.Setup,
		jobdefinition.Setup,
		jobqueue.Setup,
		schedulingpolicy.Setup,
		inferenceprofile.Setup,
		agent.Setup,
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
		domaincodeartifact.Setup,
		domainpermissionspolicy.Setup,
		repository.Setup,
		repositorypermissionspolicy.Setup,
		approvalruletemplate.Setup,
		approvalruletemplateassociation.Setup,
		repositorycodecommit.Setup,
		trigger.Setup,
		profilinggroup.Setup,
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
		clusterdsql.Setup,
		clusterpeering.Setup,
		contributorinsights.Setup,
		globaltable.Setup,
		kinesisstreamingdestination.Setup,
		resourcepolicydynamodb.Setup,
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
		fleetec2.Setup,
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
		accessentry.Setup,
		accesspolicyassociation.Setup,
		addon.Setup,
		clustereks.Setup,
		clusterauth.Setup,
		fargateprofile.Setup,
		identityproviderconfig.Setup,
		nodegroup.Setup,
		podidentityassociation.Setup,
		clusterelasticache.Setup,
		globalreplicationgroup.Setup,
		parametergroupelasticache.Setup,
		replicationgroup.Setup,
		serverlesscache.Setup,
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
		lbtruststore.Setup,
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
		catalogtableoptimizer.Setup,
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
		authorizeriot.Setup,
		certificateiot.Setup,
		domainconfiguration.Setup,
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
		clusterpolicy.Setup,
		configuration.Setup,
		replicator.Setup,
		scramsecretassociation.Setup,
		serverlesscluster.Setup,
		singlescramsecretassociation.Setup,
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
		usermq.Setup,
		environmentmwaa.Setup,
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
		accountorganizations.Setup,
		delegatedadministrator.Setup,
		organization.Setup,
		organizationalunit.Setup,
		policyorganizations.Setup,
		policyattachmentorganizations.Setup,
		pipelineosis.Setup,
		apppinpoint.Setup,
		smschannel.Setup,
		pipe.Setup,
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
		instancestaterds.Setup,
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
		endpointaccess.Setup,
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
		endpointaccessredshiftserverless.Setup,
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
		querylog.Setup,
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
		directorybucket.Setup,
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
		mlflowtrackingserver.Setup,
		modelsagemaker.Setup,
		modelpackagegroup.Setup,
		modelpackagegrouppolicy.Setup,
		notebookinstance.Setup,
		notebookinstancelifecycleconfiguration.Setup,
		servicecatalogportfoliostatus.Setup,
		space.Setup,
		studiolifecycleconfig.Setup,
		userprofile.Setup,
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
		connectortransfer.Setup,
		server.Setup,
		sshkey.Setup,
		tagtransfer.Setup,
		usertransfer.Setup,
		workflowtransfer.Setup,
		endpointverifiedaccess.Setup,
		groupverifiedaccess.Setup,
		instanceverifiedaccess.Setup,
		instanceloggingconfiguration.Setup,
		instancetrustproviderattachment.Setup,
		trustprovider.Setup,
		networkperformancemetricsubscription.Setup,
		accesslogsubscription.Setup,
		authpolicy.Setup,
		listenervpclattice.Setup,
		listenerrule.Setup,
		resourceconfiguration.Setup,
		resourcegateway.Setup,
		resourcepolicyvpclattice.Setup,
		servicevpclattice.Setup,
		servicenetwork.Setup,
		servicenetworkresourceassociation.Setup,
		servicenetworkserviceassociation.Setup,
		servicenetworkvpcassociation.Setup,
		targetgroup.Setup,
		targetgroupattachment.Setup,
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
		rulegroupwafv2.Setup,
		webaclwafv2.Setup,
		webaclassociation.Setup,
		webaclloggingconfiguration.Setup,
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

// SetupGated_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyzer.SetupGated,
		archiverule.SetupGated,
		alternatecontact.SetupGated,
		region.SetupGated,
		certificate.SetupGated,
		certificatevalidation.SetupGated,
		certificateacmpca.SetupGated,
		certificateauthority.SetupGated,
		certificateauthoritycertificate.SetupGated,
		permission.SetupGated,
		policy.SetupGated,
		alertmanagerdefinition.SetupGated,
		rulegroupnamespace.SetupGated,
		scraper.SetupGated,
		workspace.SetupGated,
		app.SetupGated,
		backendenvironment.SetupGated,
		branch.SetupGated,
		webhook.SetupGated,
		account.SetupGated,
		apikey.SetupGated,
		authorizer.SetupGated,
		basepathmapping.SetupGated,
		clientcertificate.SetupGated,
		deployment.SetupGated,
		documentationpart.SetupGated,
		documentationversion.SetupGated,
		domainname.SetupGated,
		gatewayresponse.SetupGated,
		integration.SetupGated,
		integrationresponse.SetupGated,
		method.SetupGated,
		methodresponse.SetupGated,
		methodsettings.SetupGated,
		model.SetupGated,
		requestvalidator.SetupGated,
		resource.SetupGated,
		restapi.SetupGated,
		restapipolicy.SetupGated,
		stage.SetupGated,
		usageplan.SetupGated,
		usageplankey.SetupGated,
		vpclink.SetupGated,
		api.SetupGated,
		apimapping.SetupGated,
		authorizerapigatewayv2.SetupGated,
		deploymentapigatewayv2.SetupGated,
		domainnameapigatewayv2.SetupGated,
		integrationapigatewayv2.SetupGated,
		integrationresponseapigatewayv2.SetupGated,
		modelapigatewayv2.SetupGated,
		route.SetupGated,
		routeresponse.SetupGated,
		stageapigatewayv2.SetupGated,
		vpclinkapigatewayv2.SetupGated,
		policyappautoscaling.SetupGated,
		scheduledaction.SetupGated,
		target.SetupGated,
		application.SetupGated,
		configurationprofile.SetupGated,
		deploymentappconfig.SetupGated,
		deploymentstrategy.SetupGated,
		environment.SetupGated,
		extension.SetupGated,
		extensionassociation.SetupGated,
		hostedconfigurationversion.SetupGated,
		flow.SetupGated,
		eventintegration.SetupGated,
		applicationapplicationinsights.SetupGated,
		gatewayroute.SetupGated,
		mesh.SetupGated,
		routeappmesh.SetupGated,
		virtualgateway.SetupGated,
		virtualnode.SetupGated,
		virtualrouter.SetupGated,
		virtualservice.SetupGated,
		autoscalingconfigurationversion.SetupGated,
		connection.SetupGated,
		observabilityconfiguration.SetupGated,
		service.SetupGated,
		vpcconnector.SetupGated,
		directoryconfig.SetupGated,
		fleet.SetupGated,
		fleetstackassociation.SetupGated,
		imagebuilder.SetupGated,
		stack.SetupGated,
		user.SetupGated,
		userstackassociation.SetupGated,
		apicache.SetupGated,
		apikeyappsync.SetupGated,
		datasource.SetupGated,
		function.SetupGated,
		graphqlapi.SetupGated,
		resolver.SetupGated,
		database.SetupGated,
		datacatalog.SetupGated,
		namedquery.SetupGated,
		workgroup.SetupGated,
		attachment.SetupGated,
		autoscalinggroup.SetupGated,
		grouptag.SetupGated,
		launchconfiguration.SetupGated,
		lifecyclehook.SetupGated,
		notification.SetupGated,
		policyautoscaling.SetupGated,
		schedule.SetupGated,
		scalingplan.SetupGated,
		framework.SetupGated,
		globalsettings.SetupGated,
		plan.SetupGated,
		regionsettings.SetupGated,
		reportplan.SetupGated,
		selection.SetupGated,
		vault.SetupGated,
		vaultlockconfiguration.SetupGated,
		vaultnotifications.SetupGated,
		vaultpolicy.SetupGated,
		computeenvironment.SetupGated,
		jobdefinition.SetupGated,
		jobqueue.SetupGated,
		schedulingpolicy.SetupGated,
		inferenceprofile.SetupGated,
		agent.SetupGated,
		budget.SetupGated,
		budgetaction.SetupGated,
		anomalymonitor.SetupGated,
		voiceconnector.SetupGated,
		voiceconnectorgroup.SetupGated,
		voiceconnectorlogging.SetupGated,
		voiceconnectororigination.SetupGated,
		voiceconnectorstreaming.SetupGated,
		voiceconnectortermination.SetupGated,
		voiceconnectorterminationcredentials.SetupGated,
		environmentec2.SetupGated,
		environmentmembership.SetupGated,
		resourcecloudcontrol.SetupGated,
		stackcloudformation.SetupGated,
		stackset.SetupGated,
		stacksetinstance.SetupGated,
		cachepolicy.SetupGated,
		distribution.SetupGated,
		fieldlevelencryptionconfig.SetupGated,
		fieldlevelencryptionprofile.SetupGated,
		functioncloudfront.SetupGated,
		keygroup.SetupGated,
		monitoringsubscription.SetupGated,
		originaccesscontrol.SetupGated,
		originaccessidentity.SetupGated,
		originrequestpolicy.SetupGated,
		publickey.SetupGated,
		realtimelogconfig.SetupGated,
		responseheaderspolicy.SetupGated,
		domain.SetupGated,
		domainserviceaccesspolicy.SetupGated,
		eventdatastore.SetupGated,
		trail.SetupGated,
		compositealarm.SetupGated,
		dashboard.SetupGated,
		metricalarm.SetupGated,
		metricstream.SetupGated,
		apidestination.SetupGated,
		archive.SetupGated,
		bus.SetupGated,
		buspolicy.SetupGated,
		connectioncloudwatchevents.SetupGated,
		permissioncloudwatchevents.SetupGated,
		rule.SetupGated,
		targetcloudwatchevents.SetupGated,
		definition.SetupGated,
		destination.SetupGated,
		destinationpolicy.SetupGated,
		group.SetupGated,
		metricfilter.SetupGated,
		resourcepolicy.SetupGated,
		stream.SetupGated,
		subscriptionfilter.SetupGated,
		domaincodeartifact.SetupGated,
		domainpermissionspolicy.SetupGated,
		repository.SetupGated,
		repositorypermissionspolicy.SetupGated,
		approvalruletemplate.SetupGated,
		approvalruletemplateassociation.SetupGated,
		repositorycodecommit.SetupGated,
		trigger.SetupGated,
		profilinggroup.SetupGated,
		codepipeline.SetupGated,
		customactiontype.SetupGated,
		webhookcodepipeline.SetupGated,
		connectioncodestarconnections.SetupGated,
		host.SetupGated,
		notificationrule.SetupGated,
		cognitoidentitypoolproviderprincipaltag.SetupGated,
		pool.SetupGated,
		poolrolesattachment.SetupGated,
		identityprovider.SetupGated,
		resourceserver.SetupGated,
		riskconfiguration.SetupGated,
		usercognitoidp.SetupGated,
		usergroup.SetupGated,
		useringroup.SetupGated,
		userpool.SetupGated,
		userpoolclient.SetupGated,
		userpooldomain.SetupGated,
		userpooluicustomization.SetupGated,
		awsconfigurationrecorderstatus.SetupGated,
		configrule.SetupGated,
		configurationaggregator.SetupGated,
		configurationrecorder.SetupGated,
		conformancepack.SetupGated,
		deliverychannel.SetupGated,
		remediationconfiguration.SetupGated,
		botassociation.SetupGated,
		contactflow.SetupGated,
		contactflowmodule.SetupGated,
		hoursofoperation.SetupGated,
		instance.SetupGated,
		instancestorageconfig.SetupGated,
		lambdafunctionassociation.SetupGated,
		phonenumber.SetupGated,
		queue.SetupGated,
		quickconnect.SetupGated,
		routingprofile.SetupGated,
		securityprofile.SetupGated,
		userconnect.SetupGated,
		userhierarchystructure.SetupGated,
		vocabulary.SetupGated,
		reportdefinition.SetupGated,
		dataset.SetupGated,
		revision.SetupGated,
		pipeline.SetupGated,
		locations3.SetupGated,
		task.SetupGated,
		cluster.SetupGated,
		parametergroup.SetupGated,
		subnetgroup.SetupGated,
		appdeploy.SetupGated,
		deploymentconfig.SetupGated,
		deploymentgroup.SetupGated,
		graph.SetupGated,
		invitationaccepter.SetupGated,
		member.SetupGated,
		devicepool.SetupGated,
		instanceprofile.SetupGated,
		networkprofile.SetupGated,
		project.SetupGated,
		testgridproject.SetupGated,
		upload.SetupGated,
		bgppeer.SetupGated,
		connectiondirectconnect.SetupGated,
		connectionassociation.SetupGated,
		gateway.SetupGated,
		gatewayassociation.SetupGated,
		gatewayassociationproposal.SetupGated,
		hostedprivatevirtualinterface.SetupGated,
		hostedprivatevirtualinterfaceaccepter.SetupGated,
		hostedpublicvirtualinterface.SetupGated,
		hostedpublicvirtualinterfaceaccepter.SetupGated,
		hostedtransitvirtualinterface.SetupGated,
		hostedtransitvirtualinterfaceaccepter.SetupGated,
		lag.SetupGated,
		privatevirtualinterface.SetupGated,
		publicvirtualinterface.SetupGated,
		transitvirtualinterface.SetupGated,
		lifecyclepolicy.SetupGated,
		certificatedms.SetupGated,
		endpoint.SetupGated,
		eventsubscription.SetupGated,
		replicationinstance.SetupGated,
		replicationsubnetgroup.SetupGated,
		replicationtask.SetupGated,
		s3endpoint.SetupGated,
		clusterdocdb.SetupGated,
		clusterinstance.SetupGated,
		clusterparametergroup.SetupGated,
		clustersnapshot.SetupGated,
		eventsubscriptiondocdb.SetupGated,
		globalcluster.SetupGated,
		subnetgroupdocdb.SetupGated,
		conditionalforwarder.SetupGated,
		directory.SetupGated,
		shareddirectory.SetupGated,
		clusterdsql.SetupGated,
		clusterpeering.SetupGated,
		contributorinsights.SetupGated,
		globaltable.SetupGated,
		kinesisstreamingdestination.SetupGated,
		resourcepolicydynamodb.SetupGated,
		table.SetupGated,
		tableitem.SetupGated,
		tablereplica.SetupGated,
		tag.SetupGated,
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
		fleetec2.SetupGated,
		flowlog.SetupGated,
		hostec2.SetupGated,
		instanceec2.SetupGated,
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
		routeec2.SetupGated,
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
		tagec2.SetupGated,
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
		lifecyclepolicyecr.SetupGated,
		pullthroughcacherule.SetupGated,
		registrypolicy.SetupGated,
		registryscanningconfiguration.SetupGated,
		replicationconfiguration.SetupGated,
		repositoryecr.SetupGated,
		repositorypolicy.SetupGated,
		repositoryecrpublic.SetupGated,
		repositorypolicyecrpublic.SetupGated,
		accountsettingdefault.SetupGated,
		capacityprovider.SetupGated,
		clusterecs.SetupGated,
		clustercapacityproviders.SetupGated,
		serviceecs.SetupGated,
		taskdefinition.SetupGated,
		accesspoint.SetupGated,
		backuppolicy.SetupGated,
		filesystem.SetupGated,
		filesystempolicy.SetupGated,
		mounttarget.SetupGated,
		replicationconfigurationefs.SetupGated,
		accessentry.SetupGated,
		accesspolicyassociation.SetupGated,
		addon.SetupGated,
		clustereks.SetupGated,
		clusterauth.SetupGated,
		fargateprofile.SetupGated,
		identityproviderconfig.SetupGated,
		nodegroup.SetupGated,
		podidentityassociation.SetupGated,
		clusterelasticache.SetupGated,
		globalreplicationgroup.SetupGated,
		parametergroupelasticache.SetupGated,
		replicationgroup.SetupGated,
		serverlesscache.SetupGated,
		subnetgroupelasticache.SetupGated,
		userelasticache.SetupGated,
		usergroupelasticache.SetupGated,
		applicationelasticbeanstalk.SetupGated,
		applicationversion.SetupGated,
		configurationtemplate.SetupGated,
		domainelasticsearch.SetupGated,
		domainpolicy.SetupGated,
		domainsamloptions.SetupGated,
		pipelineelastictranscoder.SetupGated,
		preset.SetupGated,
		appcookiestickinesspolicy.SetupGated,
		attachmentelb.SetupGated,
		backendserverpolicy.SetupGated,
		elb.SetupGated,
		lbcookiestickinesspolicy.SetupGated,
		lbsslnegotiationpolicy.SetupGated,
		listenerpolicy.SetupGated,
		policyelb.SetupGated,
		proxyprotocolpolicy.SetupGated,
		lb.SetupGated,
		lblistener.SetupGated,
		lblistenercertificate.SetupGated,
		lblistenerrule.SetupGated,
		lbtargetgroup.SetupGated,
		lbtargetgroupattachment.SetupGated,
		lbtruststore.SetupGated,
		securityconfiguration.SetupGated,
		applicationemrserverless.SetupGated,
		feature.SetupGated,
		projectevidently.SetupGated,
		segment.SetupGated,
		deliverystream.SetupGated,
		experimenttemplate.SetupGated,
		backup.SetupGated,
		datarepositoryassociation.SetupGated,
		lustrefilesystem.SetupGated,
		ontapfilesystem.SetupGated,
		ontapstoragevirtualmachine.SetupGated,
		windowsfilesystem.SetupGated,
		alias.SetupGated,
		build.SetupGated,
		fleetgamelift.SetupGated,
		gamesessionqueue.SetupGated,
		script.SetupGated,
		vaultglacier.SetupGated,
		vaultlock.SetupGated,
		accelerator.SetupGated,
		endpointgroup.SetupGated,
		listener.SetupGated,
		catalogdatabase.SetupGated,
		catalogtable.SetupGated,
		catalogtableoptimizer.SetupGated,
		classifier.SetupGated,
		connectionglue.SetupGated,
		crawler.SetupGated,
		datacatalogencryptionsettings.SetupGated,
		job.SetupGated,
		registry.SetupGated,
		resourcepolicyglue.SetupGated,
		schema.SetupGated,
		securityconfigurationglue.SetupGated,
		triggerglue.SetupGated,
		userdefinedfunction.SetupGated,
		workflow.SetupGated,
		licenseassociation.SetupGated,
		roleassociation.SetupGated,
		workspacegrafana.SetupGated,
		workspaceapikey.SetupGated,
		workspacesamlconfiguration.SetupGated,
		detector.SetupGated,
		filter.SetupGated,
		memberguardduty.SetupGated,
		accesskey.SetupGated,
		accountalias.SetupGated,
		accountpasswordpolicy.SetupGated,
		groupiam.SetupGated,
		groupmembership.SetupGated,
		grouppolicyattachment.SetupGated,
		instanceprofileiam.SetupGated,
		openidconnectprovider.SetupGated,
		policyiam.SetupGated,
		role.SetupGated,
		rolepolicy.SetupGated,
		rolepolicyattachment.SetupGated,
		samlprovider.SetupGated,
		servercertificate.SetupGated,
		servicelinkedrole.SetupGated,
		servicespecificcredential.SetupGated,
		signingcertificate.SetupGated,
		useriam.SetupGated,
		usergroupmembership.SetupGated,
		userloginprofile.SetupGated,
		userpolicyattachment.SetupGated,
		usersshkey.SetupGated,
		virtualmfadevice.SetupGated,
		groupidentitystore.SetupGated,
		groupmembershipidentitystore.SetupGated,
		useridentitystore.SetupGated,
		component.SetupGated,
		containerrecipe.SetupGated,
		distributionconfiguration.SetupGated,
		image.SetupGated,
		imagepipeline.SetupGated,
		imagerecipe.SetupGated,
		infrastructureconfiguration.SetupGated,
		assessmenttarget.SetupGated,
		assessmenttemplate.SetupGated,
		resourcegroup.SetupGated,
		enabler.SetupGated,
		authorizeriot.SetupGated,
		certificateiot.SetupGated,
		domainconfiguration.SetupGated,
		indexingconfiguration.SetupGated,
		loggingoptions.SetupGated,
		policyiot.SetupGated,
		policyattachment.SetupGated,
		provisioningtemplate.SetupGated,
		rolealias.SetupGated,
		thing.SetupGated,
		thinggroup.SetupGated,
		thinggroupmembership.SetupGated,
		thingprincipalattachment.SetupGated,
		thingtype.SetupGated,
		topicrule.SetupGated,
		topicruledestination.SetupGated,
		channel.SetupGated,
		recordingconfiguration.SetupGated,
		clusterkafka.SetupGated,
		clusterpolicy.SetupGated,
		configuration.SetupGated,
		replicator.SetupGated,
		scramsecretassociation.SetupGated,
		serverlesscluster.SetupGated,
		singlescramsecretassociation.SetupGated,
		connector.SetupGated,
		customplugin.SetupGated,
		workerconfiguration.SetupGated,
		datasourcekendra.SetupGated,
		experience.SetupGated,
		index.SetupGated,
		querysuggestionsblocklist.SetupGated,
		thesaurus.SetupGated,
		keyspace.SetupGated,
		tablekeyspaces.SetupGated,
		streamkinesis.SetupGated,
		streamconsumer.SetupGated,
		applicationkinesisanalytics.SetupGated,
		applicationkinesisanalyticsv2.SetupGated,
		applicationsnapshot.SetupGated,
		streamkinesisvideo.SetupGated,
		aliaskms.SetupGated,
		ciphertext.SetupGated,
		externalkey.SetupGated,
		grant.SetupGated,
		key.SetupGated,
		replicaexternalkey.SetupGated,
		replicakey.SetupGated,
		datalakesettings.SetupGated,
		permissions.SetupGated,
		resourcelakeformation.SetupGated,
		aliaslambda.SetupGated,
		codesigningconfig.SetupGated,
		eventsourcemapping.SetupGated,
		functionlambda.SetupGated,
		functioneventinvokeconfig.SetupGated,
		functionurl.SetupGated,
		invocation.SetupGated,
		layerversion.SetupGated,
		layerversionpermission.SetupGated,
		permissionlambda.SetupGated,
		provisionedconcurrencyconfig.SetupGated,
		bot.SetupGated,
		botalias.SetupGated,
		intent.SetupGated,
		slottype.SetupGated,
		association.SetupGated,
		licenseconfiguration.SetupGated,
		bucket.SetupGated,
		certificatelightsail.SetupGated,
		containerservice.SetupGated,
		disk.SetupGated,
		diskattachment.SetupGated,
		domainlightsail.SetupGated,
		domainentry.SetupGated,
		instancelightsail.SetupGated,
		instancepublicports.SetupGated,
		keypairlightsail.SetupGated,
		lblightsail.SetupGated,
		lbattachment.SetupGated,
		lbcertificate.SetupGated,
		lbstickinesspolicy.SetupGated,
		staticip.SetupGated,
		staticipattachment.SetupGated,
		geofencecollection.SetupGated,
		placeindex.SetupGated,
		routecalculator.SetupGated,
		tracker.SetupGated,
		trackerassociation.SetupGated,
		accountmacie2.SetupGated,
		classificationjob.SetupGated,
		customdataidentifier.SetupGated,
		findingsfilter.SetupGated,
		invitationacceptermacie2.SetupGated,
		membermacie2.SetupGated,
		queuemediaconvert.SetupGated,
		channelmedialive.SetupGated,
		input.SetupGated,
		inputsecuritygroup.SetupGated,
		multiplex.SetupGated,
		channelmediapackage.SetupGated,
		container.SetupGated,
		containerpolicy.SetupGated,
		acl.SetupGated,
		clustermemorydb.SetupGated,
		parametergroupmemorydb.SetupGated,
		snapshot.SetupGated,
		subnetgroupmemorydb.SetupGated,
		usermemorydb.SetupGated,
		broker.SetupGated,
		configurationmq.SetupGated,
		usermq.SetupGated,
		environmentmwaa.SetupGated,
		clusterneptune.SetupGated,
		clusterendpoint.SetupGated,
		clusterinstanceneptune.SetupGated,
		clusterparametergroupneptune.SetupGated,
		clustersnapshotneptune.SetupGated,
		eventsubscriptionneptune.SetupGated,
		globalclusterneptune.SetupGated,
		parametergroupneptune.SetupGated,
		subnetgroupneptune.SetupGated,
		firewall.SetupGated,
		firewallpolicy.SetupGated,
		loggingconfiguration.SetupGated,
		rulegroup.SetupGated,
		attachmentaccepter.SetupGated,
		connectattachment.SetupGated,
		connectionnetworkmanager.SetupGated,
		corenetwork.SetupGated,
		customergatewayassociation.SetupGated,
		device.SetupGated,
		globalnetwork.SetupGated,
		link.SetupGated,
		linkassociation.SetupGated,
		site.SetupGated,
		transitgatewayconnectpeerassociation.SetupGated,
		transitgatewayregistration.SetupGated,
		vpcattachment.SetupGated,
		domainopensearch.SetupGated,
		domainpolicyopensearch.SetupGated,
		domainsamloptionsopensearch.SetupGated,
		accesspolicy.SetupGated,
		collection.SetupGated,
		lifecyclepolicyopensearchserverless.SetupGated,
		securityconfig.SetupGated,
		securitypolicy.SetupGated,
		vpcendpointopensearchserverless.SetupGated,
		accountorganizations.SetupGated,
		delegatedadministrator.SetupGated,
		organization.SetupGated,
		organizationalunit.SetupGated,
		policyorganizations.SetupGated,
		policyattachmentorganizations.SetupGated,
		pipelineosis.SetupGated,
		apppinpoint.SetupGated,
		smschannel.SetupGated,
		pipe.SetupGated,
		providerconfig.SetupGated,
		ledger.SetupGated,
		streamqldb.SetupGated,
		groupquicksight.SetupGated,
		userquicksight.SetupGated,
		principalassociation.SetupGated,
		resourceassociation.SetupGated,
		resourceshare.SetupGated,
		resourceshareaccepter.SetupGated,
		clusterrds.SetupGated,
		clusteractivitystream.SetupGated,
		clusterendpointrds.SetupGated,
		clusterinstancerds.SetupGated,
		clusterparametergrouprds.SetupGated,
		clusterroleassociation.SetupGated,
		clustersnapshotrds.SetupGated,
		dbinstanceautomatedbackupsreplication.SetupGated,
		dbsnapshotcopy.SetupGated,
		eventsubscriptionrds.SetupGated,
		globalclusterrds.SetupGated,
		instancerds.SetupGated,
		instanceroleassociation.SetupGated,
		instancestaterds.SetupGated,
		optiongroup.SetupGated,
		parametergrouprds.SetupGated,
		proxy.SetupGated,
		proxydefaulttargetgroup.SetupGated,
		proxyendpoint.SetupGated,
		proxytarget.SetupGated,
		snapshotrds.SetupGated,
		subnetgrouprds.SetupGated,
		authenticationprofile.SetupGated,
		clusterredshift.SetupGated,
		endpointaccess.SetupGated,
		eventsubscriptionredshift.SetupGated,
		hsmclientcertificate.SetupGated,
		hsmconfiguration.SetupGated,
		parametergroupredshift.SetupGated,
		scheduledactionredshift.SetupGated,
		snapshotcopygrant.SetupGated,
		snapshotschedule.SetupGated,
		snapshotscheduleassociation.SetupGated,
		subnetgroupredshift.SetupGated,
		usagelimit.SetupGated,
		endpointaccessredshiftserverless.SetupGated,
		redshiftserverlessnamespace.SetupGated,
		resourcepolicyredshiftserverless.SetupGated,
		snapshotredshiftserverless.SetupGated,
		usagelimitredshiftserverless.SetupGated,
		workgroupredshiftserverless.SetupGated,
		groupresourcegroups.SetupGated,
		profile.SetupGated,
		delegationset.SetupGated,
		healthcheck.SetupGated,
		hostedzonednssec.SetupGated,
		querylog.SetupGated,
		record.SetupGated,
		resolverconfig.SetupGated,
		trafficpolicy.SetupGated,
		trafficpolicyinstance.SetupGated,
		vpcassociationauthorization.SetupGated,
		zone.SetupGated,
		zoneassociation.SetupGated,
		clusterroute53recoverycontrolconfig.SetupGated,
		controlpanel.SetupGated,
		routingcontrol.SetupGated,
		safetyrule.SetupGated,
		cell.SetupGated,
		readinesscheck.SetupGated,
		recoverygroup.SetupGated,
		resourceset.SetupGated,
		endpointroute53resolver.SetupGated,
		ruleroute53resolver.SetupGated,
		ruleassociation.SetupGated,
		appmonitor.SetupGated,
		metricsdestination.SetupGated,
		buckets3.SetupGated,
		bucketaccelerateconfiguration.SetupGated,
		bucketacl.SetupGated,
		bucketanalyticsconfiguration.SetupGated,
		bucketcorsconfiguration.SetupGated,
		bucketintelligenttieringconfiguration.SetupGated,
		bucketinventory.SetupGated,
		bucketlifecycleconfiguration.SetupGated,
		bucketlogging.SetupGated,
		bucketmetric.SetupGated,
		bucketnotification.SetupGated,
		bucketobject.SetupGated,
		bucketobjectlockconfiguration.SetupGated,
		bucketownershipcontrols.SetupGated,
		bucketpolicy.SetupGated,
		bucketpublicaccessblock.SetupGated,
		bucketreplicationconfiguration.SetupGated,
		bucketrequestpaymentconfiguration.SetupGated,
		bucketserversideencryptionconfiguration.SetupGated,
		bucketversioning.SetupGated,
		bucketwebsiteconfiguration.SetupGated,
		directorybucket.SetupGated,
		object.SetupGated,
		objectcopy.SetupGated,
		accesspoints3control.SetupGated,
		accesspointpolicy.SetupGated,
		accountpublicaccessblock.SetupGated,
		multiregionaccesspoint.SetupGated,
		multiregionaccesspointpolicy.SetupGated,
		objectlambdaaccesspoint.SetupGated,
		objectlambdaaccesspointpolicy.SetupGated,
		storagelensconfiguration.SetupGated,
		appsagemaker.SetupGated,
		appimageconfig.SetupGated,
		coderepository.SetupGated,
		devicesagemaker.SetupGated,
		devicefleet.SetupGated,
		domainsagemaker.SetupGated,
		endpointsagemaker.SetupGated,
		endpointconfiguration.SetupGated,
		featuregroup.SetupGated,
		imagesagemaker.SetupGated,
		imageversion.SetupGated,
		mlflowtrackingserver.SetupGated,
		modelsagemaker.SetupGated,
		modelpackagegroup.SetupGated,
		modelpackagegrouppolicy.SetupGated,
		notebookinstance.SetupGated,
		notebookinstancelifecycleconfiguration.SetupGated,
		servicecatalogportfoliostatus.SetupGated,
		space.SetupGated,
		studiolifecycleconfig.SetupGated,
		userprofile.SetupGated,
		workforce.SetupGated,
		workteam.SetupGated,
		schedulescheduler.SetupGated,
		schedulegroup.SetupGated,
		discoverer.SetupGated,
		registryschemas.SetupGated,
		schemaschemas.SetupGated,
		secret.SetupGated,
		secretpolicy.SetupGated,
		secretrotation.SetupGated,
		secretversion.SetupGated,
		accountsecurityhub.SetupGated,
		actiontarget.SetupGated,
		findingaggregator.SetupGated,
		insight.SetupGated,
		inviteaccepter.SetupGated,
		membersecurityhub.SetupGated,
		productsubscription.SetupGated,
		standardssubscription.SetupGated,
		cloudformationstack.SetupGated,
		budgetresourceassociation.SetupGated,
		constraint.SetupGated,
		portfolio.SetupGated,
		portfolioshare.SetupGated,
		principalportfolioassociation.SetupGated,
		product.SetupGated,
		productportfolioassociation.SetupGated,
		provisioningartifact.SetupGated,
		serviceaction.SetupGated,
		tagoption.SetupGated,
		tagoptionresourceassociation.SetupGated,
		httpnamespace.SetupGated,
		privatednsnamespace.SetupGated,
		publicdnsnamespace.SetupGated,
		serviceservicediscovery.SetupGated,
		servicequota.SetupGated,
		activereceiptruleset.SetupGated,
		configurationset.SetupGated,
		domaindkim.SetupGated,
		domainidentity.SetupGated,
		domainmailfrom.SetupGated,
		emailidentity.SetupGated,
		eventdestination.SetupGated,
		identitynotificationtopic.SetupGated,
		identitypolicy.SetupGated,
		receiptfilter.SetupGated,
		receiptrule.SetupGated,
		receiptruleset.SetupGated,
		template.SetupGated,
		configurationsetsesv2.SetupGated,
		configurationseteventdestination.SetupGated,
		dedicatedippool.SetupGated,
		emailidentitysesv2.SetupGated,
		emailidentityfeedbackattributes.SetupGated,
		emailidentitymailfromattributes.SetupGated,
		activity.SetupGated,
		statemachine.SetupGated,
		signingjob.SetupGated,
		signingprofile.SetupGated,
		signingprofilepermission.SetupGated,
		platformapplication.SetupGated,
		smspreferences.SetupGated,
		topic.SetupGated,
		topicpolicy.SetupGated,
		topicsubscription.SetupGated,
		queuesqs.SetupGated,
		queuepolicy.SetupGated,
		queueredriveallowpolicy.SetupGated,
		queueredrivepolicy.SetupGated,
		activation.SetupGated,
		associationssm.SetupGated,
		defaultpatchbaseline.SetupGated,
		document.SetupGated,
		maintenancewindow.SetupGated,
		maintenancewindowtarget.SetupGated,
		maintenancewindowtask.SetupGated,
		parameter.SetupGated,
		patchbaseline.SetupGated,
		patchgroup.SetupGated,
		resourcedatasync.SetupGated,
		servicesetting.SetupGated,
		accountassignment.SetupGated,
		customermanagedpolicyattachment.SetupGated,
		instanceaccesscontrolattributes.SetupGated,
		managedpolicyattachment.SetupGated,
		permissionsboundaryattachment.SetupGated,
		permissionset.SetupGated,
		permissionsetinlinepolicy.SetupGated,
		domainswf.SetupGated,
		databasetimestreamwrite.SetupGated,
		tabletimestreamwrite.SetupGated,
		languagemodel.SetupGated,
		vocabularytranscribe.SetupGated,
		vocabularyfilter.SetupGated,
		connectortransfer.SetupGated,
		server.SetupGated,
		sshkey.SetupGated,
		tagtransfer.SetupGated,
		usertransfer.SetupGated,
		workflowtransfer.SetupGated,
		endpointverifiedaccess.SetupGated,
		groupverifiedaccess.SetupGated,
		instanceverifiedaccess.SetupGated,
		instanceloggingconfiguration.SetupGated,
		instancetrustproviderattachment.SetupGated,
		trustprovider.SetupGated,
		networkperformancemetricsubscription.SetupGated,
		accesslogsubscription.SetupGated,
		authpolicy.SetupGated,
		listenervpclattice.SetupGated,
		listenerrule.SetupGated,
		resourceconfiguration.SetupGated,
		resourcegateway.SetupGated,
		resourcepolicyvpclattice.SetupGated,
		servicevpclattice.SetupGated,
		servicenetwork.SetupGated,
		servicenetworkresourceassociation.SetupGated,
		servicenetworkserviceassociation.SetupGated,
		servicenetworkvpcassociation.SetupGated,
		targetgroup.SetupGated,
		targetgroupattachment.SetupGated,
		bytematchset.SetupGated,
		geomatchset.SetupGated,
		ipset.SetupGated,
		ratebasedrule.SetupGated,
		regexmatchset.SetupGated,
		regexpatternset.SetupGated,
		rulewaf.SetupGated,
		sizeconstraintset.SetupGated,
		sqlinjectionmatchset.SetupGated,
		webacl.SetupGated,
		xssmatchset.SetupGated,
		bytematchsetwafregional.SetupGated,
		geomatchsetwafregional.SetupGated,
		ipsetwafregional.SetupGated,
		ratebasedrulewafregional.SetupGated,
		regexmatchsetwafregional.SetupGated,
		regexpatternsetwafregional.SetupGated,
		rulewafregional.SetupGated,
		sizeconstraintsetwafregional.SetupGated,
		sqlinjectionmatchsetwafregional.SetupGated,
		webaclwafregional.SetupGated,
		xssmatchsetwafregional.SetupGated,
		ipsetwafv2.SetupGated,
		regexpatternsetwafv2.SetupGated,
		rulegroupwafv2.SetupGated,
		webaclwafv2.SetupGated,
		webaclassociation.SetupGated,
		webaclloggingconfiguration.SetupGated,
		directoryworkspaces.SetupGated,
		ipgroup.SetupGated,
		encryptionconfig.SetupGated,
		groupxray.SetupGated,
		samplingrule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
