#!/bin/bash -e

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

check_duplicate_types() {
    knownVersions=("v1beta2" "v1beta1")

    for group in $(ls); do
        if [ -d "$group" ]; then
            versions=($(ls "$group"))
            length=${#versions[@]}
            if [ $length == 1 ]; then
                continue
            fi
            for (( i=0; i<${#knownVersions[@]}-1; i++ )); do
                if [ -d "$group/${knownVersions[i]}" ]; then
                    for file in $(ls "$group/${knownVersions[i]}" | grep "_types.go"); do
                        kind=$(echo "$file" | cut -d'_' -f2 | cut -d'.' -f1)
                        if [ "$kind" != "" ]; then
                              for (( j=i+1; j<${#knownVersions[@]}; j++ )); do
                                  if [ -f "$group/${knownVersions[j]}/zz_${kind}_types.go" ]; then
                                      echo "$group/${knownVersions[j]}/zz_${kind}_types.go will be checked out to ${CHECKOUT_RELEASE_VERSION}"
                                      git checkout ${CHECKOUT_RELEASE_VERSION} -- "$group/${knownVersions[j]}/zz_${kind}_types.go"
                                  fi
                              done
                        fi
                    done
                    break
                fi
            done
        fi
    done
}

cd apis
check_duplicate_types
