#!/usr/bin/env python3

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

import yaml
import os
import sys
import re


def load_gvks(path, loader):
    types = set()
    for root, _, files in os.walk(path):
        for f in files:
            if f.endswith(".yaml"):
                with open(os.path.join(root, f)) as s:
                    for t in yaml.safe_load_all(s):
                        for gvk in loader(t):
                            types.add(gvk)
    return types


def load_crd_type(t):
    kind = t["spec"]["names"]["kind"]
    group = t["spec"]["group"]
    for v in t["spec"]["versions"]:
        yield f'{kind}.{group}/{v["name"]}'


def is_exception(type_name, exception_patterns):
    """
    Check if a type matches any of the given exception patterns.
    Both exact matches and wildcard patterns are supported.
    """
    for pattern in exception_patterns:
        if '*' in pattern:
            # For wildcard patterns, convert * to .* for regex
            regex_pattern = pattern.replace('.', '\\.').replace('*', '.*')
            if re.match(f"^{regex_pattern}$", type_name):
                return True
        else:
            # For exact matches, simple string comparison
            if pattern == type_name:
                return True
    return False


exceptions = {
    "provider-aws": {
        # Exact match exceptions
        'ProviderConfigUsage.aws.upbound.io/v1beta1',
        # Wildcard pattern exceptions
        '*.m.upbound.io/*',
    },
}

# NOTE(muvaf): Please consider tackling https://github.com/upbound/squad-control-planes/issues/806
# before adding new functionality here.

# Example usage: check-examples.py <CRD dir> <example manifests dir>
if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Example usage: check-examples.py <CRD dir> <example manifests "
              "dir>")
        sys.exit(1)
    try:
        exception_set = exceptions["provider-aws"]
    except KeyError:
        exception_set = set()
    known_crd_types = load_gvks(sys.argv[1], load_crd_type)
    example_types = load_gvks(sys.argv[2], lambda t: [] if t is None or not {"kind", "apiVersion"}.issubset(t.keys())
        else [f'{t["kind"]}.{t["apiVersion"]}'])
    # Find all missing types that are not in the examples and are not an allowed exception
    missing_types = {t for t in known_crd_types
                     if t not in example_types and not is_exception(t, exception_set)}
    if len(missing_types) == 0:
        print("All CRDs have at least one example or are excluded...")
        print(f'Exceptions allowed for: {exception_set}')
        sys.exit(0)
    print(f'Please add example manifests for the following types: {missing_types}')
    sys.exit(2)
