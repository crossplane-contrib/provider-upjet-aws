// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package apis

import (
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var s = runtime.NewScheme()

func GetManagedResource(group, version, kind, listKind string) (xpresource.Managed, xpresource.ManagedList, error) {
	gv := schema.GroupVersion{
		Group:   group,
		Version: version,
	}
	kingGVK := gv.WithKind(kind)
	m, err := s.New(kingGVK)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to get a new API object of GVK %q from the runtime scheme", kingGVK)
	}

	listGVK := gv.WithKind(listKind)
	l, err := s.New(listGVK)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to get a new API object list of GVK %q from the runtime scheme", listGVK)
	}
	return m.(xpresource.Managed), l.(xpresource.ManagedList), nil
}

func BuildScheme(sb runtime.SchemeBuilder) error {
	return errors.Wrap(sb.AddToScheme(s), "failed to register the GVKs with the runtime scheme")
}
