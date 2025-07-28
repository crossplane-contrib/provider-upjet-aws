package clients

import "testing"

func TestPartitions(t *testing.T) {
	cases := map[string]struct {
		partition        string
		servicesToRegion map[string]string
		wantOk           bool
	}{
		"AWSDefaultPartition": {
			partition: "aws",
			servicesToRegion: map[string]string{
				"iam": "us-east-1",
			},
			wantOk: true,
		},
		"AWSChina": {
			partition: "aws-cn",
			servicesToRegion: map[string]string{
				"iam": "cn-north-1",
			},
			wantOk: true,
		},
		"AWSISO": {
			partition: "aws-iso",
			servicesToRegion: map[string]string{
				"iam": "us-iso-east-1",
			},
			wantOk: true},
		"AWSISOB": {
			partition:        "aws-iso-b",
			servicesToRegion: map[string]string{},
			wantOk:           true,
		},
		"AWSISOE": {
			partition:        "aws-iso-e",
			servicesToRegion: map[string]string{},
			wantOk:           true,
		},
		"AWSISOF": {
			partition:        "aws-iso-f",
			servicesToRegion: map[string]string{},
			wantOk:           true,
		},
		"AWSUSGov": {
			partition: "aws-us-gov",
			servicesToRegion: map[string]string{
				"iam": "us-gov-west-1",
			},
			wantOk: true,
		},
		"NonExistentPartition": {
			partition:        "aws-foo",
			servicesToRegion: map[string]string{},
			wantOk:           false,
		},
	}
	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			partition, ok := partitions[tc.partition]
			if ok != tc.wantOk {
				t.Errorf("expected partition existence: got %v, want %v", ok, tc.wantOk)
			}
			for wantSvc, wantRegion := range tc.servicesToRegion {
				reg, okSvc := partition.serviceToDefaultRegions[wantSvc]
				if !okSvc {
					t.Errorf("expected service %q to exist in partition %q", wantSvc, tc.partition)
				}
				if reg != wantRegion {
					t.Errorf("expected default region %q for service %q in partition %q, got region %q", wantRegion, wantSvc, tc.partition, reg)
				}
			}
		})
	}
	var defaultRegions = map[string]string{}
	for name, partition := range partitions {
		reg, ok := partition.serviceToDefaultRegions["iam"]
		if !ok {
			continue
		}
		defaultRegions[name] = reg
	}
}
