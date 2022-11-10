/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/upbound/upjet/pkg/pipeline"

	"github.com/upbound/provider-aws/config"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}
	p := config.GetProvider()
	pipeline.Run(p, absRootDir)
	if fp := os.Getenv("SKIPPED_RESOURCES_CSV"); len(fp) != 0 {
		if err := os.WriteFile(fp, []byte(strings.Join(p.GetSkippedResourceNames(), ";")), 0o600); err != nil {
			panic(fmt.Sprintf("cannot write skipped resources CSV to file %s: %s", fp, err.Error()))
		}
	}
}
