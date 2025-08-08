// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
)

//go:embed partitions_gen.go.tmpl
var templateBody string

const documentVersion = 3

type PartitionDatum struct {
	ID                          string
	Name                        string
	DNSSuffix                   string
	RegionRegex                 string
	CredentialScopeRegion       string
	IAMRegions                  map[string]string
	GlobalServiceSigningRegions map[string]string
	Regions                     []RegionDatum
}

type RegionDatum struct {
	ID          string
	Description string
}

type TemplateData struct {
	Partitions []PartitionDatum
}

type EndpointsDocument struct {
	Partitions []PartitionModel `json:"partitions"`
	Version    uint64           `json:"version"`
}

type PartitionModel struct {
	Defaults      DefaultsModel           `json:"defaults"`
	DnsSuffix     string                  `json:"dnsSuffix"`
	Partition     string                  `json:"partition"`
	PartitionName string                  `json:"partitionName"`
	RegionRegex   string                  `json:"regionRegex"`
	Regions       map[string]RegionModel  `json:"regions"`
	Services      map[string]ServiceModel `json:"services"`
}

type DefaultsModel struct {
	Hostname          string         `json:"hostname"`
	Protocols         []string       `json:"protocols"`
	SignatureVersions []string       `json:"signatureVersions"`
	Variants          []VariantModel `json:"variants"`
}

type VariantModel struct {
	DnsSuffix string   `json:"dnsSuffix"`
	Hostname  string   `json:"hostname"`
	Tags      []string `json:"tags"`
}

type RegionModel struct {
	Description string `json:"description"`
}

type ServiceModel struct {
	Endpoints         map[string]EndpointModel `json:"endpoints"`
	IsRegionalized    bool                     `json:"isRegionalized,omitempty"`
	PartitionEndpoint string                   `json:"partitionEndpoint,omitempty"`
	Defaults          *DefaultsModel           `json:"defaults,omitempty"`
}

type EndpointModel struct {
	CredentialScope *CredentialScopeModel `json:"credentialScope,omitempty"`
	Hostname        string                `json:"hostname"`
	Protocols       []string              `json:"protocols,omitempty"`
	Variants        []VariantModel        `json:"variants,omitempty"`
	Deprecated      bool                  `json:"deprecated,omitempty"`
}

type CredentialScopeModel struct {
	Region  string `json:"region,omitempty"`
	Service string `json:"service,omitempty"`
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go <aws-sdk-go-v2-endpoints-json-url>\n\n")
}

func main() { //nolint:gocyclo
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(2)
	}

	inputURL := args[0]
	targetFilename := `zz_partitions_gen.go`
	var endpointDocu EndpointsDocument

	log.Println("Generating AWS partition definitions file", targetFilename)

	if err := readEndpointsDocumentFromURL(inputURL, &endpointDocu); err != nil {
		log.Fatalf("error reading JSON from %s: %s", inputURL, err)
	}

	templateData := TemplateData{}

	if endpointDocu.Version != documentVersion {
		log.Fatalf("unsupported endpoints document version: %d, expected version: %d", endpointDocu.Version, documentVersion)
	}

	for _, partition := range endpointDocu.Partitions {
		partitionDatum := PartitionDatum{
			GlobalServiceSigningRegions: make(map[string]string),
		}
		partitionDatum.ID = partition.Partition
		partitionDatum.Name = partition.PartitionName
		partitionDatum.DNSSuffix = partition.DnsSuffix
		partitionDatum.RegionRegex = partition.RegionRegex
		for id, region := range partition.Regions {
			regionDatum := RegionDatum{
				ID:          id,
				Description: region.Description,
			}
			partitionDatum.Regions = append(partitionDatum.Regions, regionDatum)
		}

		for svcName, svc := range partition.Services {
			if svc.PartitionEndpoint != "" {
				defaultEndpoint, ok := svc.Endpoints[svc.PartitionEndpoint]
				if !ok {
					log.Fatalf("partition endpoint %q not found for service %q in partition %q", svc.PartitionEndpoint, svcName, partition.Partition)
				}
				if defaultEndpoint.CredentialScope == nil {
					continue
				}
				partitionDatum.GlobalServiceSigningRegions[svcName] = defaultEndpoint.CredentialScope.Region
			}
		}
		templateData.Partitions = append(templateData.Partitions, partitionDatum)

	}

	sort.SliceStable(templateData.Partitions, func(i, j int) bool {
		return templateData.Partitions[i].ID < templateData.Partitions[j].ID
	})

	for i := 0; i < len(templateData.Partitions); i++ {
		sort.SliceStable(templateData.Partitions[i].Regions, func(j, k int) bool {
			return templateData.Partitions[i].Regions[j].ID < templateData.Partitions[i].Regions[k].ID
		})
	}

	tmpl, err := template.New("partitions").Parse(templateBody)

	if err != nil {
		log.Fatalf("parsing function template: %v", err)
	}

	targetFile, err := os.OpenFile(targetFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("failed to open file %q for write: %v", targetFilename, err)
	}
	defer func() {
		if err := targetFile.Close(); err != nil {
			log.Fatalf("Failed to close the file %q: %s", targetFilename, err.Error())
		}
	}()

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, templateData); err != nil {
		log.Fatalf("cannot execute template: %v", err) //nolint:gocritic
	}
	gofmtFormattedBytes, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("cannot gofmt generated partitions file: %v", err)
	}
	if _, err := targetFile.Write(gofmtFormattedBytes); err != nil {
		log.Fatalf("cannot write generated file: %v", err)
	}

	log.Println("Successfully generated AWS partition definitions file", targetFilename)
}

func readEndpointsDocumentFromURL(url string, to *EndpointsDocument) error {
	r, err := http.Get(url) //nolint only for endpoint generation, with a fixed AWS url
	if err != nil {
		return errors.Wrap(err, "cannot fetch remote endpoints document")
	}
	if r.StatusCode < 200 || r.StatusCode > 299 {
		return errors.Errorf("fetching endpoints document returned non-2xx HTTP status code: %s", r.Status)
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

	epDocumentRaw, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "cannot read HTTP body")
	}
	return errors.Wrap(json.Unmarshal(epDocumentRaw, to), "cannot unmarshal endpoints document")
}
