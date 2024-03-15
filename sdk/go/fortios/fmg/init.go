// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fmg

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
	case "fortios:fmg/devicemanagerDevice:DevicemanagerDevice":
		r = &DevicemanagerDevice{}
	case "fortios:fmg/devicemanagerInstallDevice:DevicemanagerInstallDevice":
		r = &DevicemanagerInstallDevice{}
	case "fortios:fmg/devicemanagerInstallPolicypackage:DevicemanagerInstallPolicypackage":
		r = &DevicemanagerInstallPolicypackage{}
	case "fortios:fmg/devicemanagerScript:DevicemanagerScript":
		r = &DevicemanagerScript{}
	case "fortios:fmg/devicemanagerScriptExecute:DevicemanagerScriptExecute":
		r = &DevicemanagerScriptExecute{}
	case "fortios:fmg/firewallObjectAddress:FirewallObjectAddress":
		r = &FirewallObjectAddress{}
	case "fortios:fmg/firewallObjectIppool:FirewallObjectIppool":
		r = &FirewallObjectIppool{}
	case "fortios:fmg/firewallObjectService:FirewallObjectService":
		r = &FirewallObjectService{}
	case "fortios:fmg/firewallObjectVip:FirewallObjectVip":
		r = &FirewallObjectVip{}
	case "fortios:fmg/firewallSecurityPolicy:FirewallSecurityPolicy":
		r = &FirewallSecurityPolicy{}
	case "fortios:fmg/firewallSecurityPolicypackage:FirewallSecurityPolicypackage":
		r = &FirewallSecurityPolicypackage{}
	case "fortios:fmg/jsonrpcRequest:JsonrpcRequest":
		r = &JsonrpcRequest{}
	case "fortios:fmg/objectAdomRevision:ObjectAdomRevision":
		r = &ObjectAdomRevision{}
	case "fortios:fmg/systemAdmin:SystemAdmin":
		r = &SystemAdmin{}
	case "fortios:fmg/systemAdminProfiles:SystemAdminProfiles":
		r = &SystemAdminProfiles{}
	case "fortios:fmg/systemAdminUser:SystemAdminUser":
		r = &SystemAdminUser{}
	case "fortios:fmg/systemAdom:SystemAdom":
		r = &SystemAdom{}
	case "fortios:fmg/systemDns:SystemDns":
		r = &SystemDns{}
	case "fortios:fmg/systemGlobal:SystemGlobal":
		r = &SystemGlobal{}
	case "fortios:fmg/systemLicenseForticare:SystemLicenseForticare":
		r = &SystemLicenseForticare{}
	case "fortios:fmg/systemLicenseVm:SystemLicenseVm":
		r = &SystemLicenseVm{}
	case "fortios:fmg/systemNetworkInterface:SystemNetworkInterface":
		r = &SystemNetworkInterface{}
	case "fortios:fmg/systemNetworkRoute:SystemNetworkRoute":
		r = &SystemNetworkRoute{}
	case "fortios:fmg/systemNtp:SystemNtp":
		r = &SystemNtp{}
	case "fortios:fmg/systemSyslogserver:SystemSyslogserver":
		r = &SystemSyslogserver{}
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
		"fmg/devicemanagerDevice",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/devicemanagerInstallDevice",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/devicemanagerInstallPolicypackage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/devicemanagerScript",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/devicemanagerScriptExecute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/firewallObjectAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/firewallObjectIppool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/firewallObjectService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/firewallObjectVip",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/firewallSecurityPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/firewallSecurityPolicypackage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/jsonrpcRequest",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/objectAdomRevision",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/systemAdmin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/systemAdminProfiles",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/systemAdminUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/systemAdom",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/systemDns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/systemGlobal",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/systemLicenseForticare",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/systemLicenseVm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/systemNetworkInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/systemNetworkRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/systemNtp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"fmg/systemSyslogserver",
		&module{version},
	)
}