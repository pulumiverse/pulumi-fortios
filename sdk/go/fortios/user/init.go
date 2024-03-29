// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

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
	case "fortios:user/adgrp:Adgrp":
		r = &Adgrp{}
	case "fortios:user/certificate:Certificate":
		r = &Certificate{}
	case "fortios:user/device:Device":
		r = &Device{}
	case "fortios:user/deviceaccesslist:Deviceaccesslist":
		r = &Deviceaccesslist{}
	case "fortios:user/devicecategory:Devicecategory":
		r = &Devicecategory{}
	case "fortios:user/devicegroup:Devicegroup":
		r = &Devicegroup{}
	case "fortios:user/domaincontroller:Domaincontroller":
		r = &Domaincontroller{}
	case "fortios:user/exchange:Exchange":
		r = &Exchange{}
	case "fortios:user/externalidentityprovider:Externalidentityprovider":
		r = &Externalidentityprovider{}
	case "fortios:user/fortitoken:Fortitoken":
		r = &Fortitoken{}
	case "fortios:user/fsso:Fsso":
		r = &Fsso{}
	case "fortios:user/fssopolling:Fssopolling":
		r = &Fssopolling{}
	case "fortios:user/group:Group":
		r = &Group{}
	case "fortios:user/krbkeytab:Krbkeytab":
		r = &Krbkeytab{}
	case "fortios:user/ldap:Ldap":
		r = &Ldap{}
	case "fortios:user/local:Local":
		r = &Local{}
	case "fortios:user/nacpolicy:Nacpolicy":
		r = &Nacpolicy{}
	case "fortios:user/passwordpolicy:Passwordpolicy":
		r = &Passwordpolicy{}
	case "fortios:user/peer:Peer":
		r = &Peer{}
	case "fortios:user/peergrp:Peergrp":
		r = &Peergrp{}
	case "fortios:user/pop3:Pop3":
		r = &Pop3{}
	case "fortios:user/quarantine:Quarantine":
		r = &Quarantine{}
	case "fortios:user/radius:Radius":
		r = &Radius{}
	case "fortios:user/saml:Saml":
		r = &Saml{}
	case "fortios:user/securityexemptlist:Securityexemptlist":
		r = &Securityexemptlist{}
	case "fortios:user/setting:Setting":
		r = &Setting{}
	case "fortios:user/tacacs:Tacacs":
		r = &Tacacs{}
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
		"user/adgrp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/certificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/device",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/deviceaccesslist",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/devicecategory",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/devicegroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/domaincontroller",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/exchange",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/externalidentityprovider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/fortitoken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/fsso",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/fssopolling",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/krbkeytab",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/ldap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/local",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/nacpolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/passwordpolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/peer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/peergrp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/pop3",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/quarantine",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/radius",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/saml",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/securityexemptlist",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/setting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"user/tacacs",
		&module{version},
	)
}
