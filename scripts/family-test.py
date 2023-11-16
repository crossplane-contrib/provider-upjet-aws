#!/usr/bin/env python3

import yaml
import os
import sys
import logging
from collections import namedtuple


def get_annotation(r, key, default=None):
    try:
        return r["metadata"]["annotations"][key]
    except KeyError:
        return default


# Example usage: family-test.py <example manifests dir>
if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Example usage: family-test.py <example manifests dir> <default uptest timeout>")
        sys.exit(1)

    examples_dir = sys.argv[1]
    default_timeout = int(sys.argv[2])
    manifests = dict()
    info = namedtuple('Info', ['path', 'count', 'timeout'])

    for root, _, files in os.walk(examples_dir):
        for f in files:
            if f.endswith(".yaml"):
                api_folder = os.path.basename(root)
                if api_folder in {"providerconfig", "examples"}:
                    break
                m = os.path.join(root, f)
                with (open(m) as s):
                    resources = [r for r in yaml.safe_load_all(s)]
                    # check if any resource manifest in the file has
                    # the manual-intervention annotation
                    if len([r for r in resources if get_annotation(r, "upjet.upbound.io/manual-intervention")]) > 0:
                        logging.warning("Skipping %s as at least one resource manifest"
                                        "has the manual-intervention annotation in it...", m)
                        continue

                    i = info(
                        path=m,
                        count=len(resources),
                        timeout=max([int(get_annotation(r, "uptest.upbound.io/timeout", default_timeout))
                                     for r in resources]))
                    if api_folder not in manifests or manifests[api_folder].timeout > i.timeout or \
                            (manifests[api_folder].timeout == i.timeout and manifests[api_folder].count > i.count):
                        manifests[api_folder] = i

    for i in manifests.values():
        print(f'{i.path}')
