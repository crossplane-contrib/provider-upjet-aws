#!/usr/bin/env python3

import yaml
import os
import sys


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


exceptions = {
    "provider-aws": {
        'ProviderConfigUsage.aws.upbound.io/v1beta1', 
    },

    "provider-azure": {
        'ProviderConfigUsage.azure.upbound.io/v1beta1'
    },

    "provider-gcp": {
        'ManagedSSLCertificate.compute.gcp.upbound.io/v1beta1', 
        'StoreConfig.gcp.upbound.io/v1alpha1', 
        'ProviderConfigUsage.gcp.upbound.io/v1beta1'
    }
}


# Example usage: check-examples.py <provider name> <CRD dir> <example manifests dir>
if __name__ == "__main__":
    if len(sys.argv) != 4:
        print("Example usage: check-examples.py <provider name> <CRD dir> <example manifests dir>")
        sys.exit(1)
    try:
        exception_set = exceptions[sys.argv[1]]
    except KeyError:
        exception_set = set()
    known_crd_types = load_gvks(sys.argv[2], load_crd_type)
    example_types = load_gvks(sys.argv[3], lambda t: [] if t is None or not {"kind", "apiVersion"}.issubset(t.keys())
        else [f'{t["kind"]}.{t["apiVersion"]}'])
    diff = known_crd_types.difference(example_types.union(exception_set))
    if len(diff) == 0:
        print("All CRDs have at least one example...")
        print(f'Exceptions allowed for: {exception_set}')
        sys.exit(0)
    print(f'Please add example manifests for the following types: {diff}')
    sys.exit(2)
