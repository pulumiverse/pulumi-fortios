// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package application

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "fortios:application/custom:Custom":
		r = &Custom{}
	case "fortios:application/group:Group":
		r = &Group{}
	case "fortios:application/list:List":
		r = &List{}
	case "fortios:application/name:Name":
		r = &Name{}
	case "fortios:application/rulesettings:Rulesettings":
		r = &Rulesettings{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"fortios",
		"application/custom",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"application/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"application/list",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"application/name",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"application/rulesettings",
		&module{version},
	)
}
