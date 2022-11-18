#!/usr/bin/env bash
set -aeuo pipefail

# Delete the VoiceConnectorTerminationCredentials resource before deleting the VoiceConnector istself
${KUBECTL} delete voiceconnectorterminationcredentials.chime.aws.upbound.io --all