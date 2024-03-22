// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package endpointcontrol

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
	case "fortios:endpointcontrol/client:Client":
		r = &Client{}
	case "fortios:endpointcontrol/fctems:Fctems":
		r = &Fctems{}
	case "fortios:endpointcontrol/fctemsoverride:Fctemsoverride":
		r = &Fctemsoverride{}
	case "fortios:endpointcontrol/forticlientems:Forticlientems":
		r = &Forticlientems{}
	case "fortios:endpointcontrol/forticlientregistrationsync:Forticlientregistrationsync":
		r = &Forticlientregistrationsync{}
	case "fortios:endpointcontrol/profile:Profile":
		r = &Profile{}
	case "fortios:endpointcontrol/registeredforticlient:Registeredforticlient":
		r = &Registeredforticlient{}
	case "fortios:endpointcontrol/settings:Settings":
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
		"endpointcontrol/client",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"endpointcontrol/fctems",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"endpointcontrol/fctemsoverride",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"endpointcontrol/forticlientems",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"endpointcontrol/forticlientregistrationsync",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"endpointcontrol/profile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"endpointcontrol/registeredforticlient",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"endpointcontrol/settings",
		&module{version},
	)
}
