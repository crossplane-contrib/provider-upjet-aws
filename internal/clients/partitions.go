package clients

import "regexp"

type awsPartition struct {
	id                      string
	name                    string
	regionRegex             *regexp.Regexp //nolint:unused
	dnsSuffix               string
	serviceToDefaultRegions map[string]string
	regions                 map[string]awsRegion //nolint:unused
}

type awsRegion struct { //nolint:unused
	id          string
	description string
}

func getIAMDefaultSigningRegions() map[string]string {
	var partitionToDefaultRegion = map[string]string{}
	for name, partition := range partitions {
		reg, ok := partition.serviceToDefaultRegions["iam"]
		if !ok {
			continue
		}
		partitionToDefaultRegion[name] = reg
	}
	return partitionToDefaultRegion
}
