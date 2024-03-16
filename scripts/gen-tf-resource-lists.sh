#!/bin/bash

XP_PROVIDER_PATH=`pwd`

cd $TF_PROVIDER_PATH/internal/service
git grep -l '@FrameworkResource'  | xargs grep '\.TypeName = "aws_' | sed -n 's|\([^/]*\)/\(\w*\)\.go:.*"\(aws_[a-z0-9_]*\)".*|\1 \2 \3|p' > $XP_PROVIDER_PATH/hack/framework-resources.lst
git grep '// @SDKResource("aws' | sed -n 's|\([^/]*\)/\(\w*\)\.go:// @SDKResource(.\([a-z0-9_]*\).*|\1 \2 \3|p' > $XP_PROVIDER_PATH/hack/sdk-resources.lst