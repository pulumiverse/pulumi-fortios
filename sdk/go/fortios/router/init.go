// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

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
	case "fortios:router/accesslist6:Accesslist6":
		r = &Accesslist6{}
	case "fortios:router/accesslist:Accesslist":
		r = &Accesslist{}
	case "fortios:router/aspathlist:Aspathlist":
		r = &Aspathlist{}
	case "fortios:router/authpath:Authpath":
		r = &Authpath{}
	case "fortios:router/bfd6:Bfd6":
		r = &Bfd6{}
	case "fortios:router/bfd:Bfd":
		r = &Bfd{}
	case "fortios:router/bgp:Bgp":
		r = &Bgp{}
	case "fortios:router/communitylist:Communitylist":
		r = &Communitylist{}
	case "fortios:router/isis:Isis":
		r = &Isis{}
	case "fortios:router/keychain:Keychain":
		r = &Keychain{}
	case "fortios:router/multicast6:Multicast6":
		r = &Multicast6{}
	case "fortios:router/multicast:Multicast":
		r = &Multicast{}
	case "fortios:router/multicastflow:Multicastflow":
		r = &Multicastflow{}
	case "fortios:router/ospf6:Ospf6":
		r = &Ospf6{}
	case "fortios:router/ospf:Ospf":
		r = &Ospf{}
	case "fortios:router/policy6:Policy6":
		r = &Policy6{}
	case "fortios:router/policy:Policy":
		r = &Policy{}
	case "fortios:router/prefixlist6:Prefixlist6":
		r = &Prefixlist6{}
	case "fortios:router/prefixlist:Prefixlist":
		r = &Prefixlist{}
	case "fortios:router/rip:Rip":
		r = &Rip{}
	case "fortios:router/ripng:Ripng":
		r = &Ripng{}
	case "fortios:router/routemap:Routemap":
		r = &Routemap{}
	case "fortios:router/setting:Setting":
		r = &Setting{}
	case "fortios:router/static6:Static6":
		r = &Static6{}
	case "fortios:router/static:Static":
		r = &Static{}
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
		"router/accesslist",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/accesslist6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/aspathlist",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/authpath",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/bfd",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/bfd6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/bgp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/communitylist",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/isis",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/keychain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/multicast",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/multicast6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/multicastflow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/ospf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/ospf6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/policy6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/prefixlist",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/prefixlist6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/rip",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/ripng",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/routemap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/setting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/static",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"router/static6",
		&module{version},
	)
}