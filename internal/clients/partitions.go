package clients

import "regexp"

type awsPartition struct {
	id                      string
	name                    string
	regionRegex             *regexp.Regexp
	dnsSuffix               string
	serviceToDefaultRegions map[string]string
	regions                 map[string]awsRegion
}

type awsRegion struct {
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
