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
