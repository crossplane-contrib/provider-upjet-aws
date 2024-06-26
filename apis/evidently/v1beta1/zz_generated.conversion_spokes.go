// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	ujconversion "github.com/crossplane/upjet/pkg/controller/conversion"
	"github.com/crossplane/upjet/pkg/resource"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this Feature to the hub type.
func (tr *Feature) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the Feature type.
func (tr *Feature) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this Project to the hub type.
func (tr *Project) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the Project type.
func (tr *Project) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}
