package project

import (
	"github.com/giantswarm/versionbundle"
)

func NewVersionBundle() versionbundle.Bundle {
	return versionbundle.Bundle{
		Changelogs: []versionbundle.Changelog{
			{
				Component:   "update-operator",
				Description: "TODO",
				Kind:        versionbundle.KindChanged,
			},
		},
		Components: []versionbundle.Component{},
		Name:       "update-operator",
		Version:    BundleVersion(),
	}
}
