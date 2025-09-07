#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Delete the VoiceConnectorTerminationCredentials resource before deleting the VoiceConnector istself
${KUBECTL} delete voiceconnectorterminationcredentials.chime.aws.upbound.io --all