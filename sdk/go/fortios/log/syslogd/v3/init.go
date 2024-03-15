// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

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
	case "fortios:log/syslogd/v3/filter:Filter":
		r = &Filter{}
	case "fortios:log/syslogd/v3/overridefilter:Overridefilter":
		r = &Overridefilter{}
	case "fortios:log/syslogd/v3/overridesetting:Overridesetting":
		r = &Overridesetting{}
	case "fortios:log/syslogd/v3/setting:Setting":
		r = &Setting{}
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
		"log/syslogd/v3/filter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"log/syslogd/v3/overridefilter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"log/syslogd/v3/overridesetting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"log/syslogd/v3/setting",
		&module{version},
	)
}
