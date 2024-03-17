// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package antivirus

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
	case "fortios:antivirus/heuristic:Heuristic":
		r = &Heuristic{}
	case "fortios:antivirus/profile:Profile":
		r = &Profile{}
	case "fortios:antivirus/quarantine:Quarantine":
		r = &Quarantine{}
	case "fortios:antivirus/settings:Settings":
		r = &Settings{}
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
		"antivirus/heuristic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"antivirus/profile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"antivirus/quarantine",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"antivirus/settings",
		&module{version},
	)
}
