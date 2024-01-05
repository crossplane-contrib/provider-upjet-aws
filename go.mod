module github.com/upbound/provider-aws

go 1.21

require (
	dario.cat/mergo v1.0.0
	github.com/aws/aws-sdk-go-v2 v1.18.0
	github.com/aws/aws-sdk-go-v2/config v1.18.12
	github.com/aws/aws-sdk-go-v2/credentials v1.13.12
	github.com/aws/aws-sdk-go-v2/service/eks v1.22.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.18.11
	github.com/aws/smithy-go v1.13.5
	github.com/crossplane/crossplane-runtime v1.15.0-rc.0.0.20231215091746-d23a82b3a2f5
	github.com/crossplane/crossplane-tools v0.0.0-20230925130601-628280f8bf79
	github.com/crossplane/upjet v1.1.0-rc.0.0.20231227120826-4cb45f9104ac
	github.com/go-ini/ini v1.46.0
	github.com/google/go-cmp v0.6.0
	github.com/hashicorp/terraform-json v0.17.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.26.1
	github.com/hashicorp/terraform-provider-aws v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/api v0.29.0
	k8s.io/apimachinery v0.29.0
	k8s.io/client-go v0.29.0
	k8s.io/utils v0.0.0-20240102154912-e7106e64919e
	sigs.k8s.io/controller-runtime v0.16.3
	sigs.k8s.io/controller-tools v0.13.0
)

require (
	github.com/ProtonMail/go-crypto v0.0.0-20230201104953-d1d05f4e2bfb // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/antchfx/htmlquery v1.2.4 // indirect
	github.com/antchfx/xpath v1.2.0 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/aws/aws-sdk-go v1.44.261 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.33 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.27 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.29 // indirect
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.19.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/account v1.10.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/acm v1.17.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.24.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/cleanrooms v1.1.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.11.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.20.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.24.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.17.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/docdbelastic v1.1.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.97.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/fis v1.14.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/healthlake v1.15.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/iam v1.19.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.16.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.13.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.27 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.14.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ivschat v1.4.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/kendra v1.40.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/lambda v1.34.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/medialive v1.31.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/oam v1.1.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.2.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/pipes v1.2.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/rbin v1.8.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/rds v1.44.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.2.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/rolesanywhere v1.1.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.14.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3control v1.31.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.1.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.3.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.17.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssm v1.36.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmcontacts v1.15.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmincidents v1.21.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/transcribe v1.26.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/vpclattice v1.0.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/xray v1.16.11 // indirect
	github.com/beevik/etree v1.1.4 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cloudflare/circl v1.3.3 // indirect
	github.com/dave/jennifer v1.4.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/zapr v1.2.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/gobuffalo/flect v1.0.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go v0.20.0 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.25 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2 v2.0.0-beta.26 // indirect
	github.com/hashicorp/awspolicyequivalence v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.4.10 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hc-install v0.5.0 // indirect
	github.com/hashicorp/hcl/v2 v2.17.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.18.1 // indirect
	github.com/hashicorp/terraform-plugin-framework v1.2.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-timeouts v0.3.1 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.10.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.16.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/hashicorp/terraform-plugin-mux v0.10.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.2.1 // indirect
	github.com/hashicorp/terraform-svchost v0.1.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20200820163806-098863c1fc24 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
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
	github.com/prometheus/client_golang v1.16.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/smartystreets/goconvey v1.7.2 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tmccombs/hcl2json v0.3.3 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yuin/goldmark v1.4.13 // indirect
	github.com/zclconf/go-cty v1.13.2 // indirect
	github.com/zclconf/go-cty-yaml v1.0.3 // indirect
	go.opentelemetry.io/otel v1.20.0 // indirect
	go.opentelemetry.io/otel/trace v1.20.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/oauth2 v0.15.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/term v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.15.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231120223509-83a465c0220f // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.1-0.20231128094519-2087447a6b4a // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.28.4 // indirect
	k8s.io/component-base v0.28.4 // indirect
	k8s.io/klog/v2 v2.110.1 // indirect
	k8s.io/kube-openapi v0.0.0-20231010175941-2dd684a91f00 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

replace golang.org/x/exp => golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1

replace github.com/hashicorp/terraform-provider-aws => github.com/upbound/terraform-provider-aws v0.0.0-20231026091456-f2d38ee240d7
