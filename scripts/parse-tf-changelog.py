import re
import json
from enum import Enum
import pdb

resources = json.load(open('config/generated.lst'))

print(len(resources))

version_regex = re.compile(r"## (\d+\.\d+\.\d+) \((.*)\)")
new_resource_regex = re.compile(r"\*\*New Resource:\*\* `(aws_\w+)`")
resource_regex = re.compile(r"resource/(aws_\w+):")
section_regex = re.compile(r"(BREAKING CHANGES|NOTES|FEATURES|ENHANCEMENTS|BUG FIXES)")


class Section(Enum):
    BREAKING_CHANGES = "BREAKING CHANGES"
    NOTES = "NOTES"
    FEATURES = "FEATURES"
    ENHANCEMENTS = "ENHANCEMENTS"
    BUG_FIXES = "BUG FIXES"
    NEW_RESOURCE = "NEW RESOURCE"

class TerraformChange:
    def __init__(
            self, 
            version: str, 
            section: Section, 
            resource: str | None, 
            line: str,
            release_date: str):
        self.version = version
        self.section = section
        self.resource = resource
        self.line = line
        self.release_date = release_date

    def __repr__(self) -> str:
        return f"{self.version} {self.section.value} {self.line}"
    
    def __lt__(self, other):
        if self.resource == other.resource:
            return self.version < other.version
        if self.resource is None:
            return True
        if other.resource is None:
            return False
        return self.resource < other.resource

def parse_changelog(old_version):
    section = None
    version = None
    release_date = None
    out = []
    with (open("../terraform-provider-aws/CHANGELOG.md")) as changelog:
        while line := changelog.readline():
            if line.isspace():
                continue
            if m := version_regex.search(line):
                version = m.group(1)
                release_date = m.group(2)
                continue
            if version == old_version:
                break
            if m := section_regex.search(line):
                section = Section(m.group(1))
                continue
            if line.startswith("* data-source/") or line.startswith("* **New Data Source:**"):
                continue
            if m := new_resource_regex.search(line):
                change = TerraformChange(version, Section.NEW_RESOURCE, m.group(1), line.strip(), release_date)
                out.append(change)
                continue
            if m := resource_regex.search(line):
                change = TerraformChange(version, section, m.group(1), line.strip(), release_date)
                out.append(change)
                continue
            other = TerraformChange(version, section, None, line.strip(), release_date)
            out.append(other)
    return out


changes = parse_changelog("5.31.0")
relevant_changes = [c for c in changes if not c.resource or c.resource in resources]

relevant_changes.sort()


for c in sorted(relevant_changes):
    print(c)
print(len(set(c.resource for c in relevant_changes if c.resource is not None)))