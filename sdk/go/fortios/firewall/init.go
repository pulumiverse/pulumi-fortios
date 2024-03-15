// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

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
	case "fortios:firewall/accessproxy6:Accessproxy6":
		r = &Accessproxy6{}
	case "fortios:firewall/accessproxy:Accessproxy":
		r = &Accessproxy{}
	case "fortios:firewall/accessproxysshclientcert:Accessproxysshclientcert":
		r = &Accessproxysshclientcert{}
	case "fortios:firewall/accessproxyvirtualhost:Accessproxyvirtualhost":
		r = &Accessproxyvirtualhost{}
	case "fortios:firewall/address6:Address6":
		r = &Address6{}
	case "fortios:firewall/address6template:Address6template":
		r = &Address6template{}
	case "fortios:firewall/address:Address":
		r = &Address{}
	case "fortios:firewall/addrgrp6:Addrgrp6":
		r = &Addrgrp6{}
	case "fortios:firewall/addrgrp:Addrgrp":
		r = &Addrgrp{}
	case "fortios:firewall/authportal:Authportal":
		r = &Authportal{}
	case "fortios:firewall/centralsnatmap:Centralsnatmap":
		r = &Centralsnatmap{}
	case "fortios:firewall/centralsnatmapMove:CentralsnatmapMove":
		r = &CentralsnatmapMove{}
	case "fortios:firewall/centralsnatmapSort:CentralsnatmapSort":
		r = &CentralsnatmapSort{}
	case "fortios:firewall/city:City":
		r = &City{}
	case "fortios:firewall/country:Country":
		r = &Country{}
	case "fortios:firewall/decryptedtrafficmirror:Decryptedtrafficmirror":
		r = &Decryptedtrafficmirror{}
	case "fortios:firewall/dnstranslation:Dnstranslation":
		r = &Dnstranslation{}
	case "fortios:firewall/doSpolicy6:DoSpolicy6":
		r = &DoSpolicy6{}
	case "fortios:firewall/doSpolicy:DoSpolicy":
		r = &DoSpolicy{}
	case "fortios:firewall/global:Global":
		r = &Global{}
	case "fortios:firewall/identitybasedroute:Identitybasedroute":
		r = &Identitybasedroute{}
	case "fortios:firewall/interfacepolicy6:Interfacepolicy6":
		r = &Interfacepolicy6{}
	case "fortios:firewall/interfacepolicy:Interfacepolicy":
		r = &Interfacepolicy{}
	case "fortios:firewall/internetservice:Internetservice":
		r = &Internetservice{}
	case "fortios:firewall/internetserviceaddition:Internetserviceaddition":
		r = &Internetserviceaddition{}
	case "fortios:firewall/internetserviceappend:Internetserviceappend":
		r = &Internetserviceappend{}
	case "fortios:firewall/internetservicebotnet:Internetservicebotnet":
		r = &Internetservicebotnet{}
	case "fortios:firewall/internetservicecustom:Internetservicecustom":
		r = &Internetservicecustom{}
	case "fortios:firewall/internetservicecustomgroup:Internetservicecustomgroup":
		r = &Internetservicecustomgroup{}
	case "fortios:firewall/internetservicedefinition:Internetservicedefinition":
		r = &Internetservicedefinition{}
	case "fortios:firewall/internetserviceextension:Internetserviceextension":
		r = &Internetserviceextension{}
	case "fortios:firewall/internetservicegroup:Internetservicegroup":
		r = &Internetservicegroup{}
	case "fortios:firewall/internetserviceipblreason:Internetserviceipblreason":
		r = &Internetserviceipblreason{}
	case "fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor":
		r = &Internetserviceipblvendor{}
	case "fortios:firewall/internetservicelist:Internetservicelist":
		r = &Internetservicelist{}
	case "fortios:firewall/internetservicename:Internetservicename":
		r = &Internetservicename{}
	case "fortios:firewall/internetserviceowner:Internetserviceowner":
		r = &Internetserviceowner{}
	case "fortios:firewall/internetservicereputation:Internetservicereputation":
		r = &Internetservicereputation{}
	case "fortios:firewall/ippool6:Ippool6":
		r = &Ippool6{}
	case "fortios:firewall/ippool:Ippool":
		r = &Ippool{}
	case "fortios:firewall/iptranslation:Iptranslation":
		r = &Iptranslation{}
	case "fortios:firewall/ipv6ehfilter:Ipv6ehfilter":
		r = &Ipv6ehfilter{}
	case "fortios:firewall/ldbmonitor:Ldbmonitor":
		r = &Ldbmonitor{}
	case "fortios:firewall/localinpolicy6:Localinpolicy6":
		r = &Localinpolicy6{}
	case "fortios:firewall/localinpolicy:Localinpolicy":
		r = &Localinpolicy{}
	case "fortios:firewall/multicastaddress6:Multicastaddress6":
		r = &Multicastaddress6{}
	case "fortios:firewall/multicastaddress:Multicastaddress":
		r = &Multicastaddress{}
	case "fortios:firewall/multicastpolicy6:Multicastpolicy6":
		r = &Multicastpolicy6{}
	case "fortios:firewall/multicastpolicy:Multicastpolicy":
		r = &Multicastpolicy{}
	case "fortios:firewall/networkservicedynamic:Networkservicedynamic":
		r = &Networkservicedynamic{}
	case "fortios:firewall/objectAddress:ObjectAddress":
		r = &ObjectAddress{}
	case "fortios:firewall/objectAddressgroup:ObjectAddressgroup":
		r = &ObjectAddressgroup{}
	case "fortios:firewall/objectIppool:ObjectIppool":
		r = &ObjectIppool{}
	case "fortios:firewall/objectService:ObjectService":
		r = &ObjectService{}
	case "fortios:firewall/objectServicecategory:ObjectServicecategory":
		r = &ObjectServicecategory{}
	case "fortios:firewall/objectServicegroup:ObjectServicegroup":
		r = &ObjectServicegroup{}
	case "fortios:firewall/objectVip:ObjectVip":
		r = &ObjectVip{}
	case "fortios:firewall/objectVipgroup:ObjectVipgroup":
		r = &ObjectVipgroup{}
	case "fortios:firewall/policy46:Policy46":
		r = &Policy46{}
	case "fortios:firewall/policy64:Policy64":
		r = &Policy64{}
	case "fortios:firewall/policy6:Policy6":
		r = &Policy6{}
	case "fortios:firewall/policy:Policy":
		r = &Policy{}
	case "fortios:firewall/profilegroup:Profilegroup":
		r = &Profilegroup{}
	case "fortios:firewall/profileprotocoloptions:Profileprotocoloptions":
		r = &Profileprotocoloptions{}
	case "fortios:firewall/proxyaddress:Proxyaddress":
		r = &Proxyaddress{}
	case "fortios:firewall/proxyaddrgrp:Proxyaddrgrp":
		r = &Proxyaddrgrp{}
	case "fortios:firewall/proxypolicy:Proxypolicy":
		r = &Proxypolicy{}
	case "fortios:firewall/proxypolicyMove:ProxypolicyMove":
		r = &ProxypolicyMove{}
	case "fortios:firewall/proxypolicySort:ProxypolicySort":
		r = &ProxypolicySort{}
	case "fortios:firewall/region:Region":
		r = &Region{}
	case "fortios:firewall/securityPolicyseq:SecurityPolicyseq":
		r = &SecurityPolicyseq{}
	case "fortios:firewall/securityPolicysort:SecurityPolicysort":
		r = &SecurityPolicysort{}
	case "fortios:firewall/securitypolicy:Securitypolicy":
		r = &Securitypolicy{}
	case "fortios:firewall/shapingpolicy:Shapingpolicy":
		r = &Shapingpolicy{}
	case "fortios:firewall/shapingprofile:Shapingprofile":
		r = &Shapingprofile{}
	case "fortios:firewall/sniffer:Sniffer":
		r = &Sniffer{}
	case "fortios:firewall/sslserver:Sslserver":
		r = &Sslserver{}
	case "fortios:firewall/sslsshprofile:Sslsshprofile":
		r = &Sslsshprofile{}
	case "fortios:firewall/trafficclass:Trafficclass":
		r = &Trafficclass{}
	case "fortios:firewall/ttlpolicy:Ttlpolicy":
		r = &Ttlpolicy{}
	case "fortios:firewall/vendormac:Vendormac":
		r = &Vendormac{}
	case "fortios:firewall/vip46:Vip46":
		r = &Vip46{}
	case "fortios:firewall/vip64:Vip64":
		r = &Vip64{}
	case "fortios:firewall/vip6:Vip6":
		r = &Vip6{}
	case "fortios:firewall/vip:Vip":
		r = &Vip{}
	case "fortios:firewall/vipgrp46:Vipgrp46":
		r = &Vipgrp46{}
	case "fortios:firewall/vipgrp64:Vipgrp64":
		r = &Vipgrp64{}
	case "fortios:firewall/vipgrp6:Vipgrp6":
		r = &Vipgrp6{}
	case "fortios:firewall/vipgrp:Vipgrp":
		r = &Vipgrp{}
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
		"firewall/accessproxy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/accessproxy6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/accessproxysshclientcert",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/accessproxyvirtualhost",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/address",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/address6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/address6template",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/addrgrp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/addrgrp6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/authportal",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/centralsnatmap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/centralsnatmapMove",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/centralsnatmapSort",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/city",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/country",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/decryptedtrafficmirror",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/dnstranslation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/doSpolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/doSpolicy6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/global",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/identitybasedroute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/interfacepolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/interfacepolicy6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetservice",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetserviceaddition",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetserviceappend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetservicebotnet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetservicecustom",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetservicecustomgroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetservicedefinition",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetserviceextension",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetservicegroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetserviceipblreason",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetserviceipblvendor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetservicelist",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetservicename",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetserviceowner",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/internetservicereputation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/ippool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/ippool6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/iptranslation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/ipv6ehfilter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/ldbmonitor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/localinpolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/localinpolicy6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/multicastaddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/multicastaddress6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/multicastpolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/multicastpolicy6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/networkservicedynamic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/objectAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/objectAddressgroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/objectIppool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/objectService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/objectServicecategory",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/objectServicegroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/objectVip",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/objectVipgroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/policy46",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/policy6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/policy64",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/profilegroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/profileprotocoloptions",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/proxyaddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/proxyaddrgrp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/proxypolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/proxypolicyMove",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/proxypolicySort",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/region",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/securityPolicyseq",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/securityPolicysort",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/securitypolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/shapingpolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/shapingprofile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/sniffer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/sslserver",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/sslsshprofile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/trafficclass",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/ttlpolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/vendormac",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/vip",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/vip46",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/vip6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/vip64",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/vipgrp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/vipgrp46",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/vipgrp6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewall/vipgrp64",
		&module{version},
	)
}
