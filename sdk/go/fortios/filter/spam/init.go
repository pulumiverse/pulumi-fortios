// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spam

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
	case "fortios:filter/spam/bwl:Bwl":
		r = &Bwl{}
	case "fortios:filter/spam/bword:Bword":
		r = &Bword{}
	case "fortios:filter/spam/dnsbl:Dnsbl":
		r = &Dnsbl{}
	case "fortios:filter/spam/fortishield:Fortishield":
		r = &Fortishield{}
	case "fortios:filter/spam/iptrust:Iptrust":
		r = &Iptrust{}
	case "fortios:filter/spam/mheader:Mheader":
		r = &Mheader{}
	case "fortios:filter/spam/options:Options":
		r = &Options{}
	case "fortios:filter/spam/profile:Profile":
		r = &Profile{}
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
		"filter/spam/bwl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"filter/spam/bword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"filter/spam/dnsbl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"filter/spam/fortishield",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"filter/spam/iptrust",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"filter/spam/mheader",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"filter/spam/options",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"filter/spam/profile",
		&module{version},
	)
}
