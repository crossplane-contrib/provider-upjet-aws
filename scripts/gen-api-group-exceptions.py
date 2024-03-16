#!/usr/bin/env python3
from functools import cmp_to_key


# First run gen-tf-resource-list.sh to generate the list of sdk and framework resources
# from the source code of terraform-provider-aws. Then run this script to generate the
# contents of the GroupMap in config/groups.go.
def find_exceptions(filename: str):
    needs_change = []
    with open(filename) as f:
        resources = f.readlines()
    for r in resources:
        line = r.strip().split(" ")
        if len(line) != 3:
            raise Exception(f"Invalid line: {r}")
        [aws_group, tf_filename, tf_resource_name] = line

        # Some customizations for places where terraform-provider-aws is inconsistent, or where this provider
        # was inconsistent in the past and we need to keep that inconsistency to avoid a breaking change.

        # This resource is in the route53resolver SDK group, but it's in the route53 crossplane provider api group,
        # because of a bug in the previous manually-maintained mapping.
        if tf_resource_name == "aws_route53_resolver_config":
            continue

        # This resource is in the ec2 SDK group, but it's in the vpc crossplane provider api group (and it's the only
        # resource in that group)
        if tf_resource_name == "aws_vpc_network_performance_metric_subscription":
            continue

        # terraform-provider-aws doesn't use the same group names as the AWS SDKs for cloudwatch logs and events
        if aws_group in ["logs", "events"]:
            aws_group = "cloudwatch" + aws_group

        # The vpc ipam resources don't include "vpc_" in their filenames, but it is in the terraform resource name
        if tf_filename.startswith("ipam_"):
            tf_filename = "vpc_" + tf_filename

        # Some of the ec2 resources have filenames like "ec2_<resource>.go" and others are just "<resource>.go"
        if aws_group == "ec2" and tf_filename.startswith("ec2_"):
            tf_filename = tf_filename.removeprefix("ec2_")
        # terraform-provider-aws uses tag_gen.go as a filename for generated resources that implement tag-like behavior,
        # and are named "tag"
        if tf_filename == "tag_gen":
            tf_filename = "tag"

        name = tf_resource_name.removeprefix("aws_")
        if name.endswith(tf_filename):
            calculated_group = name.removesuffix(tf_filename).removesuffix("_")
        elif aws_group == "ec2" and name.startswith("ec2_"):
            # Several ec2 resources have filenames like transitgateway_policy_table.go, and a resource name like
            # transit_gateway_policy_table.
            calculated_group = "ec2"
        else:
            # These are mostly the really old resources like "aws_subnet" that predate the concept of api groups
            calculated_group = ""
        words_to_drop = len([w for w in calculated_group.split("_") if w])
        if aws_group == "elbv2" and tf_resource_name.startswith("aws_lb"):
            # The elbv2 resources currently all have a kind that starts with LB, so we need to not drop that word.
            words_to_drop = 0
        if aws_group != calculated_group:
            needs_change.append([tf_resource_name, aws_group, words_to_drop])
    return needs_change


# python and golang differ in their sorting behavior when one string is a prefix of another. Implement the golang order.
# hacked together to only sort lists by their first element, which is a string.
@cmp_to_key
def golang_cmp(a, b):
    if a[0] == b[0]:
        return 0
    if a[0].startswith(b[0]):
        return -1
    if b[0].startswith(a[0]):
        return 1
    return 1 if a > b else -1


if __name__ == "__main__":
    exceptions = find_exceptions("../hack/sdk-resources.lst")
    exceptions.extend(find_exceptions("../hack/framework-resources.lst"))
    exceptions.sort(key=golang_cmp)
    for l in exceptions:
        [tf_resource_name, aws_group, drop] = l
        print(f'"{tf_resource_name}": ReplaceGroupWords("{aws_group}", {drop}),')




