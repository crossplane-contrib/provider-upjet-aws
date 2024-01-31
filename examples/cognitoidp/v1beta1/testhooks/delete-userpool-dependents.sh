#!/usr/bin/env bash
set -aeuo pipefail

# TODO:  we had better seperate these two commands if hooks support
# specifying multiple scripts for a single resource
# Delete the user group resource before the user pool
${KUBECTL} delete usergroup.cognitoidp.aws.upbound.io -l testing.upbound.io/example-name=example
# Delete the useringroup resource before the user pool
${KUBECTL} delete useringroup.cognitoidp.aws.upbound.io -l testing.upbound.io/example-name=example
