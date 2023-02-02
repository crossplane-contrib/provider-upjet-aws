#!/usr/bin/env python3

import json
import sys

# usage: version_diff.py <generated resource list> <base JSON schema path> <bumped JSON schema path>
# example usage: version_diff.py config/generated.lst .work/schema.json.3.38.0 config/schema.json
if __name__ == "__main__":
    base_path = sys.argv[2]
    bumped_path = sys.argv[3]
    print(f'Reporting schema changes between "{base_path}" as base version and "{bumped_path}" as bumped version')
    with open(sys.argv[1]) as f:
        resources = json.load(f)
    with open(base_path) as f:
        base = json.load(f)
    with open(bumped_path) as f:
        bump = json.load(f)

    provider_name = None
    for k in base["provider_schemas"]:
        # the first key is the provider name
        provider_name = k
        break
    if provider_name is None:
        print(f"Cannot extract the provider name from the base schema: {base_path}")
        sys.exit(-1)
    base_schemas = base["provider_schemas"][provider_name]["resource_schemas"]
    bumped_schemas = bump["provider_schemas"][provider_name]["resource_schemas"]

    for name in resources:
        try:
            if base_schemas[name]["version"] != bumped_schemas[name]["version"]:
                print(f'{name}:{base_schemas[name]["version"]}-{bumped_schemas[name]["version"]}')
        except KeyError as ke:
            print(f'{name} is not found in schema: {ke}')
            continue
