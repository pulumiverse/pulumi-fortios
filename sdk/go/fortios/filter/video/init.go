// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package video

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
	case "fortios:filter/video/keyword:Keyword":
		r = &Keyword{}
	case "fortios:filter/video/profile:Profile":
		r = &Profile{}
	case "fortios:filter/video/youtubechannelfilter:Youtubechannelfilter":
		r = &Youtubechannelfilter{}
	case "fortios:filter/video/youtubekey:Youtubekey":
		r = &Youtubekey{}
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
		"filter/video/keyword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"filter/video/profile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"filter/video/youtubechannelfilter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"filter/video/youtubekey",
		&module{version},
	)
}
