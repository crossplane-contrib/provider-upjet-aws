#!/usr/bin/env python3

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

import yaml
import os
import sys


def load_gvks(path, loader):
    types = set()
    empty_docs = []
    for root, _, files in os.walk(path):
        for f in files:
            if f.endswith(".yaml"):
                filepath = os.path.join(root, f)
                with open(filepath) as s:
                    docs = list(yaml.safe_load_all(s))
                if None in docs:
                    empty_docs.append(filepath)
                for t in docs:
                    for gvk in loader(t):
                        types.add(gvk)
    return types, empty_docs


def load_crd_type(t):
    kind = t["spec"]["names"]["kind"]
    group = t["spec"]["group"]
    for v in t["spec"]["versions"]:
        yield f'{kind}.{group}/{v["name"]}'


exceptions = {
    "provider-aws": {
        'ProviderConfigUsage.aws.upbound.io/v1beta1',
        'ProviderConfigUsage.aws.m.upbound.io/v1beta1',
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
    known_crd_types, _ = load_gvks(sys.argv[1], load_crd_type)
    example_types, empty_docs = load_gvks(sys.argv[2], lambda t: [] if t is None or not {"kind", "apiVersion"}.issubset(t.keys())
        else [f'{t["kind"]}.{t["apiVersion"]}'])

    exit_code = 0
    if empty_docs:
        print("The following example manifests contain an empty YAML "
              "document, usually caused by a trailing '---' separator with "
              "no content after it:")
        for f in empty_docs:
            print(f'  {f}')
        exit_code = 2

    diff = known_crd_types.difference(example_types.union(exception_set))
    if len(diff) == 0:
        print("All CRDs have at least one example...")
        print(f'Exceptions allowed for: {exception_set}')
    else:
        print(f'Please add example manifests for the following types: {diff}')
        exit_code = 2

    sys.exit(exit_code)
