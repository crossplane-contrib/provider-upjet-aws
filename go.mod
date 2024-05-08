// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

module github.com/upbound/provider-aws

go 1.21

require (
	dario.cat/mergo v1.0.0
	github.com/aws/aws-sdk-go v1.49.2
	github.com/aws/aws-sdk-go-v2 v1.24.1
	github.com/aws/aws-sdk-go-v2/config v1.26.1
	github.com/aws/aws-sdk-go-v2/credentials v1.16.12
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.10
	github.com/aws/aws-sdk-go-v2/service/eks v1.35.5
	github.com/aws/aws-sdk-go-v2/service/sts v1.26.5
	github.com/aws/smithy-go v1.19.0
	github.com/crossplane/crossplane-runtime v1.16.0-rc.1.0.20240424114634-8641eb2ba384
	github.com/crossplane/crossplane-tools v0.0.0-20230925130601-628280f8bf79
	github.com/crossplane/upjet v1.3.0
	github.com/go-ini/ini v1.46.0
	github.com/google/go-cmp v0.6.0
	github.com/hashicorp/terraform-json v0.18.0
	github.com/hashicorp/terraform-plugin-framework v1.4.2
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.30.0
	github.com/hashicorp/terraform-provider-aws v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/api v0.29.4
	k8s.io/apimachinery v0.29.4
	k8s.io/client-go v0.29.4
	k8s.io/utils v0.0.0-20240502163921-fe8a2dddb1d0
	sigs.k8s.io/controller-runtime v0.17.3
	sigs.k8s.io/controller-tools v0.14.0
)

require (
	github.com/ProtonMail/go-crypto v0.0.0-20230923063757-afb1ddc0824c // indirect
	github.com/YakDriver/regexache v0.23.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/antchfx/htmlquery v1.2.4 // indirect
	github.com/antchfx/xpath v1.2.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.4 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.15.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.2.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.26.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/account v1.14.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/acm v1.22.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.26.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/appfabric v1.5.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/appflow v1.39.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.25.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.37.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.30.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/bedrock v1.5.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines v1.13.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/chimesdkvoice v1.12.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/cleanrooms v1.8.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.15.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.30.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/codecatalyst v1.10.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.22.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codeguruprofiler v1.18.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/codestarconnections v1.21.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/codestarnotifications v1.20.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.29.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.31.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/connectcases v1.12.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/controltower v1.10.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/customerprofiles v1.34.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.22.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/docdbelastic v1.6.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.26.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.141.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecr v1.24.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/emr v1.35.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/emrserverless v1.14.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/evidently v1.16.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/finspace v1.20.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/fis v1.21.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/glacier v1.19.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/healthlake v1.20.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/iam v1.28.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.21.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.20.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.2.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.8.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.16.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internetmonitor v1.10.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ivschat v1.10.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafka v1.28.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/kendra v1.47.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.7.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/lambda v1.49.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/lexmodelsv2 v1.38.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.32.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/lookoutmetrics v1.25.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediaconnect v1.24.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/medialive v1.43.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.28.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackagev2 v1.7.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/oam v1.7.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.9.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/osis v1.6.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/pipes v1.9.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/polly v1.36.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/pricing v1.24.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/qldb v1.19.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/rbin v1.14.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/rds v1.64.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshiftdata v1.23.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.8.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.19.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.19.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/rolesanywhere v1.6.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.20.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.47.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3control v1.41.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.6.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.44.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.10.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.19.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.24.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/signer v1.19.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/sns v1.26.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.29.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssm v1.44.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmcontacts v1.20.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmincidents v1.27.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.18.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.23.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.21.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/support v1.19.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/swf v1.20.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.23.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/transcribe v1.34.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/verifiedpermissions v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/vpclattice v1.5.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.35.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/xray v1.23.5 // indirect
	github.com/beevik/etree v1.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/dave/jennifer v1.4.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.8.0 // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-logr/zapr v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/gobuffalo/flect v1.0.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go v0.21.0 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.45 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2 v2.0.0-beta.46 // indirect
	github.com/hashicorp/awspolicyequivalence v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-cty v1.4.1-0.20200723130312-85980079f637 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.5.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hc-install v0.6.1 // indirect
	github.com/hashicorp/hcl/v2 v2.19.1 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.19.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-timeouts v0.4.1 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.12.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.19.1 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/hashicorp/terraform-plugin-mux v0.12.0 // indirect
	github.com/hashicorp/terraform-plugin-testing v1.6.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.2.3 // indirect
	github.com/hashicorp/terraform-svchost v0.1.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20230413205102-771768614e91 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/muvaf/typewriter v0.0.0-20210910160850-80e49fe1eb32 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/smartystreets/goconvey v1.8.1 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tmccombs/hcl2json v0.3.3 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yuin/goldmark v1.4.13 // indirect
	github.com/zclconf/go-cty v1.14.1 // indirect
	github.com/zclconf/go-cty-yaml v1.0.3 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.46.1 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.21.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/exp v0.0.0-20240112132812-db7319d0e0e3 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/oauth2 v0.15.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/term v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231120223509-83a465c0220f // indirect
	google.golang.org/grpc v1.61.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apiextensions-apiserver v0.29.2 // indirect
	k8s.io/component-base v0.29.2 // indirect
	k8s.io/klog/v2 v2.110.1 // indirect
	k8s.io/kube-openapi v0.0.0-20231010175941-2dd684a91f00 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

replace golang.org/x/exp => golang.org/x/exp v0.0.0-20231006140011-7918f672742d

replace github.com/hashicorp/terraform-provider-aws => github.com/upbound/terraform-provider-aws v0.0.0-20240328111213-f2f0fdd63866

replace github.com/hashicorp/terraform-plugin-log => github.com/gdavison/terraform-plugin-log v0.0.0-20230928191232-6c653d8ef8fb

// pin versions for https://github.com/crossplane-contrib/provider-upjet-aws/issues/1248
replace (
	github.com/aws/aws-sdk-go-v2 v1.24.1 => github.com/aws/aws-sdk-go-v2 v1.24.0
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.10 => github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.9
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.10 => github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.9
)
