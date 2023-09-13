// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fortios

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/ettle/strcase"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumiverse/pulumi-fortios/provider/pkg/version"
	"github.com/terraform-providers/terraform-provider-fortios/fortios"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	mainMod = "index" // the fortios module
)

var module_overrides = map[string]string{
	"system": "sys",
}

func convertName(tfname string) (module string, name string) {
	tfNameItems := strings.Split(tfname, "_")
	contract.Assertf(len(tfNameItems) >= 2, "Invalid snake case name %s", tfname)
	contract.Assertf(tfNameItems[0] == "fortios", "Invalid snake case name %s. Does not start with fortios", tfname)
	if len(tfNameItems) == 2 {
		module = mainMod
		name = tfNameItems[1]
	} else {
		module = tfNameItems[1]
		name = strings.Join(tfNameItems[2:], "_")
	}
	if v, ok := module_overrides[module]; ok {
		module = v
	}
	contract.Assertf(!unicode.IsDigit(rune(name[0])), "Pulumi name must not start with a digit: %s", name)
	name = strcase.ToPascal(name)
	return
}

func makeDataSource(_ string, ds string) tokens.ModuleMember {
	mod, name := convertName(ds)
	return tfbridge.MakeDataSource("fortios", mod, "get"+name)
}

func makeResource(_ string, res string) tokens.Type {
	mod, name := convertName(res)
	return tfbridge.MakeResource("fortios", mod, name)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(fortios.Provider())
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "fortios",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Fortios",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "pulumiverse",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://raw.githubusercontent.com/pulumiverse/pulumi-fortios/main/docs/fortios.png",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/pulumiverse/pulumi-fortios",
		Description:       "A Pulumi package for creating and managing Fortios resources",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords: []string{
			"pulumi",
			"fortios",
			"category/cloud",
		},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/pulumiverse/pulumi-fortios",
		Repository: "https://github.com/pulumiverse/pulumi-fortios",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		Version:   version.Version,
		GitHubOrg: "fortinetdev",
		Config: map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
			"hostname": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_ACCESS_HOSTNAME"},
				},
			},
			"token": {
				MarkAsOptional: tfbridge.True(),
				Secret:         tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_ACCESS_TOKEN"},
				},
			},
			"insecure": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_INSECURE"},
				},
			},
			"cabundlefile": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_CA_CABUNDLE"},
				},
			},
			"cabundlecontent": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_CA_CABUNDLECONTENT"},
				},
			},
			"http_proxy": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_HTTP_PROXY"},
				},
			},
			"peerauth": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_CA_PEERAUTH"},
				},
			},
			"cacert": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_CA_CACERT"},
				},
			},
			"clientcert": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_CA_CLIENTCERT"},
				},
			},
			"clientkey": {
				MarkAsOptional: tfbridge.True(),
				Secret:         tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_CA_CLIENTKEY"},
				},
			},
			"vdom": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_VDOM"},
				},
			},
			"fmg_hostname": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_FMG_HOSTNAME"},
				},
			},
			"fmg_username": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_FMG_USERNAME"},
				},
			},
			"fmg_passwd": {
				MarkAsOptional: tfbridge.True(),
				Secret:         tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_FMG_PASSWORD"},
				},
			},
			"fmg_insecure": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_FMG_INSECURE"},
				},
			},
			"fmg_cabundlefile": {
				MarkAsOptional: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FORTIOS_FMG_CABUNDLE"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: makeResource(mainMod(mainMod, "aws_iam_role")}
			//
			// "aws_acm_certificate": {
			// 	Tok: Tok: makeResource(mainMod(mainMod, "aws_acm_certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType("fortios", "Tags")},
			// 	},
			// },
			"fortios_alertemail_setting": {
				Tok: makeResource(mainMod, "fortios_alertemail_setting"),
			},
			"fortios_antivirus_heuristic": {
				Tok: makeResource(mainMod, "fortios_antivirus_heuristic"),
			},
			"fortios_antivirus_profile": {
				Tok: makeResource(mainMod, "fortios_antivirus_profile"),
			},
			"fortios_antivirus_quarantine": {
				Tok: makeResource(mainMod, "fortios_antivirus_quarantine"),
			},
			"fortios_antivirus_settings": {
				Tok: makeResource(mainMod, "fortios_antivirus_settings"),
			},
			"fortios_application_custom": {
				Tok: makeResource(mainMod, "fortios_application_custom"),
			},
			"fortios_application_group": {
				Tok: makeResource(mainMod, "fortios_application_group"),
			},
			"fortios_application_list": {
				Tok: makeResource(mainMod, "fortios_application_list"),
			},
			"fortios_application_name": {
				Tok: makeResource(mainMod, "fortios_application_name"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"name": {
						CSharpName: "ApplicationName",
					},
				},
			},
			"fortios_application_rulesettings": {
				Tok: makeResource(mainMod, "fortios_application_rulesettings"),
			},
			"fortios_authentication_rule": {
				Tok: makeResource(mainMod, "fortios_authentication_rule"),
			},
			"fortios_authentication_scheme": {
				Tok: makeResource(mainMod, "fortios_authentication_scheme"),
			},
			"fortios_authentication_setting": {
				Tok: makeResource(mainMod, "fortios_authentication_setting"),
			},
			"fortios_automation_setting": {
				Tok: makeResource(mainMod, "fortios_automation_setting"),
			},
			"fortios_certificate_ca": {
				Tok: makeResource(mainMod, "fortios_certificate_ca"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"ca": {
						CSharpName: "Certificate",
					},
				},
			},
			"fortios_certificate_crl": {
				Tok: makeResource(mainMod, "fortios_certificate_crl"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"crl": {
						CSharpName: "Data",
					},
				},
			},
			"fortios_certificate_local": {
				Tok: makeResource(mainMod, "fortios_certificate_local"),
			},
			"fortios_certificate_remote": {
				Tok: makeResource(mainMod, "fortios_certificate_remote"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"remote": {
						CSharpName: "Certificate",
					},
				},
			},
			"fortios_cifs_domaincontroller": {
				Tok: makeResource(mainMod, "fortios_cifs_domaincontroller"),
			},
			"fortios_cifs_profile": {
				Tok: makeResource(mainMod, "fortios_cifs_profile"),
			},
			"fortios_credentialstore_domaincontroller": {
				Tok: makeResource(mainMod, "fortios_credentialstore_domaincontroller"),
			},
			"fortios_dlp_datatype": {
				Tok: makeResource(mainMod, "fortios_dlp_datatype"),
			},
			"fortios_dlp_dictionary": {
				Tok: makeResource(mainMod, "fortios_dlp_dictionary"),
			},
			"fortios_dlp_filepattern": {
				Tok: makeResource(mainMod, "fortios_dlp_filepattern"),
			},
			"fortios_dlp_fpdocsource": {
				Tok: makeResource(mainMod, "fortios_dlp_fpdocsource"),
			},
			"fortios_dlp_fpsensitivity": {
				Tok: makeResource(mainMod, "fortios_dlp_fpsensitivity"),
			},
			"fortios_dlp_profile": {
				Tok: makeResource(mainMod, "fortios_dlp_profile"),
			},
			"fortios_dlp_sensitivity": {
				Tok: makeResource(mainMod, "fortios_dlp_sensitivity"),
			},
			"fortios_dlp_sensor": {
				Tok: makeResource(mainMod, "fortios_dlp_sensor"),
			},
			"fortios_dlp_settings": {
				Tok: makeResource(mainMod, "fortios_dlp_settings"),
			},
			"fortios_dnsfilter_domainfilter": {
				Tok: makeResource(mainMod, "fortios_dnsfilter_domainfilter"),
			},
			"fortios_dnsfilter_profile": {
				Tok: makeResource(mainMod, "fortios_dnsfilter_profile"),
			},
			"fortios_dpdk_cpus": {
				Tok: makeResource(mainMod, "fortios_dpdk_cpus"),
			},
			"fortios_dpdk_global": {
				Tok: makeResource(mainMod, "fortios_dpdk_global"),
			},
			"fortios_emailfilter_blockallowlist": {
				Tok: makeResource(mainMod, "fortios_emailfilter_blockallowlist"),
			},
			"fortios_emailfilter_bwl": {
				Tok: makeResource(mainMod, "fortios_emailfilter_bwl"),
			},
			"fortios_emailfilter_bword": {
				Tok: makeResource(mainMod, "fortios_emailfilter_bword"),
			},
			"fortios_emailfilter_dnsbl": {
				Tok: makeResource(mainMod, "fortios_emailfilter_dnsbl"),
			},
			"fortios_emailfilter_fortishield": {
				Tok: makeResource(mainMod, "fortios_emailfilter_fortishield"),
			},
			"fortios_emailfilter_iptrust": {
				Tok: makeResource(mainMod, "fortios_emailfilter_iptrust"),
			},
			"fortios_emailfilter_mheader": {
				Tok: makeResource(mainMod, "fortios_emailfilter_mheader"),
			},
			"fortios_emailfilter_options": {
				Tok: makeResource(mainMod, "fortios_emailfilter_options"),
			},
			"fortios_emailfilter_profile": {
				Tok: makeResource(mainMod, "fortios_emailfilter_profile"),
			},
			"fortios_endpointcontrol_client": {
				Tok: makeResource(mainMod, "fortios_endpointcontrol_client"),
			},
			"fortios_endpointcontrol_fctems": {
				Tok: makeResource(mainMod, "fortios_endpointcontrol_fctems"),
			},
			"fortios_endpointcontrol_forticlientems": {
				Tok: makeResource(mainMod, "fortios_endpointcontrol_forticlientems"),
			},
			"fortios_endpointcontrol_forticlientregistrationsync": {
				Tok: makeResource(mainMod, "fortios_endpointcontrol_forticlientregistrationsync"),
			},
			"fortios_endpointcontrol_profile": {
				Tok: makeResource(mainMod, "fortios_endpointcontrol_profile"),
			},
			"fortios_endpointcontrol_registeredforticlient": {
				Tok: makeResource(mainMod, "fortios_endpointcontrol_registeredforticlient"),
			},
			"fortios_endpointcontrol_settings": {
				Tok: makeResource(mainMod, "fortios_endpointcontrol_settings"),
			},
			"fortios_extendercontroller_dataplan": {
				Tok: makeResource(mainMod, "fortios_extendercontroller_dataplan"),
			},
			"fortios_extendercontroller_extender": {
				Tok: makeResource(mainMod, "fortios_extendercontroller_extender"),
			},
			"fortios_extendercontroller_extender1": {
				Tok: makeResource(mainMod, "fortios_extendercontroller_extender1"),
			},
			"fortios_extendercontroller_extenderprofile": {
				Tok: makeResource(mainMod, "fortios_extendercontroller_extenderprofile"),
			},
			"fortios_extensioncontroller_dataplan": {
				Tok: makeResource(mainMod, "fortios_extensioncontroller_dataplan"),
			},
			"fortios_extensioncontroller_extender": {
				Tok: makeResource(mainMod, "fortios_extensioncontroller_extender"),
			},
			"fortios_extensioncontroller_extenderprofile": {
				Tok: makeResource(mainMod, "fortios_extensioncontroller_extenderprofile"),
			},
			"fortios_extensioncontroller_fortigate": {
				Tok: makeResource(mainMod, "fortios_extensioncontroller_fortigate"),
			},
			"fortios_extensioncontroller_fortigateprofile": {
				Tok: makeResource(mainMod, "fortios_extensioncontroller_fortigateprofile"),
			},
			"fortios_filefilter_profile": {
				Tok: makeResource(mainMod, "fortios_filefilter_profile"),
			},
			"fortios_firewall_DoSpolicy": {
				Tok: makeResource(mainMod, "fortios_firewall_DoSpolicy"),
			},
			"fortios_firewall_DoSpolicy6": {
				Tok: makeResource(mainMod, "fortios_firewall_DoSpolicy6"),
			},
			"fortios_firewall_accessproxy": {
				Tok: makeResource(mainMod, "fortios_firewall_accessproxy"),
			},
			"fortios_firewall_accessproxy6": {
				Tok: makeResource(mainMod, "fortios_firewall_accessproxy6"),
			},
			"fortios_firewall_accessproxysshclientcert": {
				Tok: makeResource(mainMod, "fortios_firewall_accessproxysshclientcert"),
			},
			"fortios_firewall_accessproxyvirtualhost": {
				Tok: makeResource(mainMod, "fortios_firewall_accessproxyvirtualhost"),
			},
			"fortios_firewall_address": {
				Tok: makeResource(mainMod, "fortios_firewall_address"),
			},
			"fortios_firewall_address6": {
				Tok: makeResource(mainMod, "fortios_firewall_address6"),
			},
			"fortios_firewall_address6template": {
				Tok: makeResource(mainMod, "fortios_firewall_address6template"),
			},
			"fortios_firewall_addrgrp": {
				Tok: makeResource(mainMod, "fortios_firewall_addrgrp"),
			},
			"fortios_firewall_addrgrp6": {
				Tok: makeResource(mainMod, "fortios_firewall_addrgrp6"),
			},
			"fortios_firewall_authportal": {
				Tok: makeResource(mainMod, "fortios_firewall_authportal"),
			},
			"fortios_firewall_centralsnatmap": {
				Tok: makeResource(mainMod, "fortios_firewall_centralsnatmap"),
			},
			"fortios_firewall_centralsnatmap_move": {
				Tok: makeResource(mainMod, "fortios_firewall_centralsnatmap_move"),
			},
			"fortios_firewall_centralsnatmap_sort": {
				Tok: makeResource(mainMod, "fortios_firewall_centralsnatmap_sort"),
			},
			"fortios_firewall_city": {
				Tok: makeResource(mainMod, "fortios_firewall_city"),
			},
			"fortios_firewall_country": {
				Tok: makeResource(mainMod, "fortios_firewall_country"),
			},
			"fortios_firewall_decryptedtrafficmirror": {
				Tok: makeResource(mainMod, "fortios_firewall_decryptedtrafficmirror"),
			},
			"fortios_firewall_dnstranslation": {
				Tok: makeResource(mainMod, "fortios_firewall_dnstranslation"),
			},
			"fortios_firewall_global": {
				Tok: makeResource(mainMod, "fortios_firewall_global"),
			},
			"fortios_firewall_identitybasedroute": {
				Tok: makeResource(mainMod, "fortios_firewall_identitybasedroute"),
			},
			"fortios_firewall_interfacepolicy": {
				Tok: makeResource(mainMod, "fortios_firewall_interfacepolicy"),
			},
			"fortios_firewall_interfacepolicy6": {
				Tok: makeResource(mainMod, "fortios_firewall_interfacepolicy6"),
			},
			"fortios_firewall_internetservice": {
				Tok: makeResource(mainMod, "fortios_firewall_internetservice"),
			},
			"fortios_firewall_internetserviceaddition": {
				Tok: makeResource(mainMod, "fortios_firewall_internetserviceaddition"),
			},
			"fortios_firewall_internetserviceappend": {
				Tok: makeResource(mainMod, "fortios_firewall_internetserviceappend"),
			},
			"fortios_firewall_internetservicebotnet": {
				Tok: makeResource(mainMod, "fortios_firewall_internetservicebotnet"),
			},
			"fortios_firewall_internetservicecustom": {
				Tok: makeResource(mainMod, "fortios_firewall_internetservicecustom"),
			},
			"fortios_firewall_internetservicecustomgroup": {
				Tok: makeResource(mainMod, "fortios_firewall_internetservicecustomgroup"),
			},
			"fortios_firewall_internetservicedefinition": {
				Tok: makeResource(mainMod, "fortios_firewall_internetservicedefinition"),
			},
			"fortios_firewall_internetserviceextension": {
				Tok: makeResource(mainMod, "fortios_firewall_internetserviceextension"),
			},
			"fortios_firewall_internetservicegroup": {
				Tok: makeResource(mainMod, "fortios_firewall_internetservicegroup"),
			},
			"fortios_firewall_internetserviceipblreason": {
				Tok: makeResource(mainMod, "fortios_firewall_internetserviceipblreason"),
			},
			"fortios_firewall_internetserviceipblvendor": {
				Tok: makeResource(mainMod, "fortios_firewall_internetserviceipblvendor"),
			},
			"fortios_firewall_internetservicelist": {
				Tok: makeResource(mainMod, "fortios_firewall_internetservicelist"),
			},
			"fortios_firewall_internetservicename": {
				Tok: makeResource(mainMod, "fortios_firewall_internetservicename"),
			},
			"fortios_firewall_internetserviceowner": {
				Tok: makeResource(mainMod, "fortios_firewall_internetserviceowner"),
			},
			"fortios_firewall_internetservicereputation": {
				Tok: makeResource(mainMod, "fortios_firewall_internetservicereputation"),
			},
			"fortios_firewall_ippool": {
				Tok: makeResource(mainMod, "fortios_firewall_ippool"),
			},
			"fortios_firewall_ippool6": {
				Tok: makeResource(mainMod, "fortios_firewall_ippool6"),
			},
			"fortios_firewall_iptranslation": {
				Tok: makeResource(mainMod, "fortios_firewall_iptranslation"),
			},
			"fortios_firewall_ipv6ehfilter": {
				Tok: makeResource(mainMod, "fortios_firewall_ipv6ehfilter"),
			},
			"fortios_firewall_ldbmonitor": {
				Tok: makeResource(mainMod, "fortios_firewall_ldbmonitor"),
			},
			"fortios_firewall_localinpolicy": {
				Tok: makeResource(mainMod, "fortios_firewall_localinpolicy"),
			},
			"fortios_firewall_localinpolicy6": {
				Tok: makeResource(mainMod, "fortios_firewall_localinpolicy6"),
			},
			"fortios_firewall_multicastaddress": {
				Tok: makeResource(mainMod, "fortios_firewall_multicastaddress"),
			},
			"fortios_firewall_multicastaddress6": {
				Tok: makeResource(mainMod, "fortios_firewall_multicastaddress6"),
			},
			"fortios_firewall_multicastpolicy": {
				Tok: makeResource(mainMod, "fortios_firewall_multicastpolicy"),
			},
			"fortios_firewall_multicastpolicy6": {
				Tok: makeResource(mainMod, "fortios_firewall_multicastpolicy6"),
			},
			"fortios_firewall_networkservicedynamic": {
				Tok: makeResource(mainMod, "fortios_firewall_networkservicedynamic"),
			},
			"fortios_firewall_object_address": {
				Tok: makeResource(mainMod, "fortios_firewall_object_address"),
			},
			"fortios_firewall_object_addressgroup": {
				Tok: makeResource(mainMod, "fortios_firewall_object_addressgroup"),
			},
			"fortios_firewall_object_ippool": {
				Tok: makeResource(mainMod, "fortios_firewall_object_ippool"),
			},
			"fortios_firewall_object_service": {
				Tok: makeResource(mainMod, "fortios_firewall_object_service"),
			},
			"fortios_firewall_object_servicecategory": {
				Tok: makeResource(mainMod, "fortios_firewall_object_servicecategory"),
			},
			"fortios_firewall_object_servicegroup": {
				Tok: makeResource(mainMod, "fortios_firewall_object_servicegroup"),
			},
			"fortios_firewall_object_vip": {
				Tok: makeResource(mainMod, "fortios_firewall_object_vip"),
			},
			"fortios_firewall_object_vipgroup": {
				Tok: makeResource(mainMod, "fortios_firewall_object_vipgroup"),
			},
			"fortios_firewall_policy": {
				Tok: makeResource(mainMod, "fortios_firewall_policy"),
			},
			"fortios_firewall_policy46": {
				Tok: makeResource(mainMod, "fortios_firewall_policy46"),
			},
			"fortios_firewall_policy6": {
				Tok: makeResource(mainMod, "fortios_firewall_policy6"),
			},
			"fortios_firewall_policy64": {
				Tok: makeResource(mainMod, "fortios_firewall_policy64"),
			},
			"fortios_firewall_profilegroup": {
				Tok: makeResource(mainMod, "fortios_firewall_profilegroup"),
			},
			"fortios_firewall_profileprotocoloptions": {
				Tok: makeResource(mainMod, "fortios_firewall_profileprotocoloptions"),
			},
			"fortios_firewall_proxyaddress": {
				Tok: makeResource(mainMod, "fortios_firewall_proxyaddress"),
			},
			"fortios_firewall_proxyaddrgrp": {
				Tok: makeResource(mainMod, "fortios_firewall_proxyaddrgrp"),
			},
			"fortios_firewall_proxypolicy": {
				Tok: makeResource(mainMod, "fortios_firewall_proxypolicy"),
			},
			"fortios_firewall_proxypolicy_move": {
				Tok: makeResource(mainMod, "fortios_firewall_proxypolicy_move"),
			},
			"fortios_firewall_proxypolicy_sort": {
				Tok: makeResource(mainMod, "fortios_firewall_proxypolicy_sort"),
			},
			"fortios_firewall_region": {
				Tok: makeResource(mainMod, "fortios_firewall_region"),
			},
			"fortios_firewall_security_policyseq": {
				Tok: makeResource(mainMod, "fortios_firewall_security_policyseq"),
			},
			"fortios_firewall_security_policysort": {
				Tok: makeResource(mainMod, "fortios_firewall_security_policysort"),
			},
			"fortios_firewall_securitypolicy": {
				Tok: makeResource(mainMod, "fortios_firewall_securitypolicy"),
			},
			"fortios_firewall_shapingpolicy": {
				Tok: makeResource(mainMod, "fortios_firewall_shapingpolicy"),
			},
			"fortios_firewall_shapingprofile": {
				Tok: makeResource(mainMod, "fortios_firewall_shapingprofile"),
			},
			"fortios_firewall_sniffer": {
				Tok: makeResource(mainMod, "fortios_firewall_sniffer"),
			},
			"fortios_firewall_sslserver": {
				Tok: makeResource(mainMod, "fortios_firewall_sslserver"),
			},
			"fortios_firewall_sslsshprofile": {
				Tok: makeResource(mainMod, "fortios_firewall_sslsshprofile"),
			},
			"fortios_firewall_trafficclass": {
				Tok: makeResource(mainMod, "fortios_firewall_trafficclass"),
			},
			"fortios_firewall_ttlpolicy": {
				Tok: makeResource(mainMod, "fortios_firewall_ttlpolicy"),
			},
			"fortios_firewall_vendormac": {
				Tok: makeResource(mainMod, "fortios_firewall_vendormac"),
			},
			"fortios_firewall_vip": {
				Tok: makeResource(mainMod, "fortios_firewall_vip"),
			},
			"fortios_firewall_vip46": {
				Tok: makeResource(mainMod, "fortios_firewall_vip46"),
			},
			"fortios_firewall_vip6": {
				Tok: makeResource(mainMod, "fortios_firewall_vip6"),
			},
			"fortios_firewall_vip64": {
				Tok: makeResource(mainMod, "fortios_firewall_vip64"),
			},
			"fortios_firewall_vipgrp": {
				Tok: makeResource(mainMod, "fortios_firewall_vipgrp"),
			},
			"fortios_firewall_vipgrp46": {
				Tok: makeResource(mainMod, "fortios_firewall_vipgrp46"),
			},
			"fortios_firewall_vipgrp6": {
				Tok: makeResource(mainMod, "fortios_firewall_vipgrp6"),
			},
			"fortios_firewall_vipgrp64": {
				Tok: makeResource(mainMod, "fortios_firewall_vipgrp64"),
			},
			"fortios_firewallconsolidated_policy": {
				Tok: makeResource(mainMod, "fortios_firewallconsolidated_policy"),
			},
			"fortios_firewallipmacbinding_setting": {
				Tok: makeResource(mainMod, "fortios_firewallipmacbinding_setting"),
			},
			"fortios_firewallipmacbinding_table": {
				Tok: makeResource(mainMod, "fortios_firewallipmacbinding_table"),
			},
			"fortios_firewallschedule_group": {
				Tok: makeResource(mainMod, "fortios_firewallschedule_group"),
			},
			"fortios_firewallschedule_onetime": {
				Tok: makeResource(mainMod, "fortios_firewallschedule_onetime"),
			},
			"fortios_firewallschedule_recurring": {
				Tok: makeResource(mainMod, "fortios_firewallschedule_recurring"),
			},
			"fortios_firewallservice_category": {
				Tok: makeResource(mainMod, "fortios_firewallservice_category"),
			},
			"fortios_firewallservice_custom": {
				Tok: makeResource(mainMod, "fortios_firewallservice_custom"),
			},
			"fortios_firewallservice_group": {
				Tok: makeResource(mainMod, "fortios_firewallservice_group"),
			},
			"fortios_firewallshaper_peripshaper": {
				Tok: makeResource(mainMod, "fortios_firewallshaper_peripshaper"),
			},
			"fortios_firewallshaper_trafficshaper": {
				Tok: makeResource(mainMod, "fortios_firewallshaper_trafficshaper"),
			},
			"fortios_firewallssh_hostkey": {
				Tok: makeResource(mainMod, "fortios_firewallssh_hostkey"),
			},
			"fortios_firewallssh_localca": {
				Tok: makeResource(mainMod, "fortios_firewallssh_localca"),
			},
			"fortios_firewallssh_localkey": {
				Tok: makeResource(mainMod, "fortios_firewallssh_localkey"),
			},
			"fortios_firewallssh_setting": {
				Tok: makeResource(mainMod, "fortios_firewallssh_setting"),
			},
			"fortios_firewallssl_setting": {
				Tok: makeResource(mainMod, "fortios_firewallssl_setting"),
			},
			"fortios_firewallwildcardfqdn_custom": {
				Tok: makeResource(mainMod, "fortios_firewallwildcardfqdn_custom"),
			},
			"fortios_firewallwildcardfqdn_group": {
				Tok: makeResource(mainMod, "fortios_firewallwildcardfqdn_group"),
			},
			"fortios_fmg_devicemanager_device": {
				Tok: makeResource(mainMod, "fortios_fmg_devicemanager_device"),
			},
			"fortios_fmg_devicemanager_install_device": {
				Tok: makeResource(mainMod, "fortios_fmg_devicemanager_install_device"),
			},
			"fortios_fmg_devicemanager_install_policypackage": {
				Tok: makeResource(mainMod, "fortios_fmg_devicemanager_install_policypackage"),
			},
			"fortios_fmg_devicemanager_script": {
				Tok: makeResource(mainMod, "fortios_fmg_devicemanager_script"),
			},
			"fortios_fmg_devicemanager_script_execute": {
				Tok: makeResource(mainMod, "fortios_fmg_devicemanager_script_execute"),
			},
			"fortios_fmg_firewall_object_address": {
				Tok: makeResource(mainMod, "fortios_fmg_firewall_object_address"),
			},
			"fortios_fmg_firewall_object_ippool": {
				Tok: makeResource(mainMod, "fortios_fmg_firewall_object_ippool"),
			},
			"fortios_fmg_firewall_object_service": {
				Tok: makeResource(mainMod, "fortios_fmg_firewall_object_service"),
			},
			"fortios_fmg_firewall_object_vip": {
				Tok: makeResource(mainMod, "fortios_fmg_firewall_object_vip"),
			},
			"fortios_fmg_firewall_security_policy": {
				Tok: makeResource(mainMod, "fortios_fmg_firewall_security_policy"),
			},
			"fortios_fmg_firewall_security_policypackage": {
				Tok: makeResource(mainMod, "fortios_fmg_firewall_security_policypackage"),
			},
			"fortios_fmg_jsonrpc_request": {
				Tok: makeResource(mainMod, "fortios_fmg_jsonrpc_request"),
			},
			"fortios_fmg_object_adom_revision": {
				Tok: makeResource(mainMod, "fortios_fmg_object_adom_revision"),
			},
			"fortios_fmg_system_admin": {
				Tok: makeResource(mainMod, "fortios_fmg_system_admin"),
			},
			"fortios_fmg_system_admin_profiles": {
				Tok: makeResource(mainMod, "fortios_fmg_system_admin_profiles"),
			},
			"fortios_fmg_system_admin_user": {
				Tok: makeResource(mainMod, "fortios_fmg_system_admin_user"),
			},
			"fortios_fmg_system_adom": {
				Tok: makeResource(mainMod, "fortios_fmg_system_adom"),
			},
			"fortios_fmg_system_dns": {
				Tok: makeResource(mainMod, "fortios_fmg_system_dns"),
			},
			"fortios_fmg_system_global": {
				Tok: makeResource(mainMod, "fortios_fmg_system_global"),
			},
			"fortios_fmg_system_license_forticare": {
				Tok: makeResource(mainMod, "fortios_fmg_system_license_forticare"),
			},
			"fortios_fmg_system_license_vm": {
				Tok: makeResource(mainMod, "fortios_fmg_system_license_vm"),
			},
			"fortios_fmg_system_network_interface": {
				Tok: makeResource(mainMod, "fortios_fmg_system_network_interface"),
			},
			"fortios_fmg_system_network_route": {
				Tok: makeResource(mainMod, "fortios_fmg_system_network_route"),
			},
			"fortios_fmg_system_ntp": {
				Tok: makeResource(mainMod, "fortios_fmg_system_ntp"),
			},
			"fortios_fmg_system_syslogserver": {
				Tok: makeResource(mainMod, "fortios_fmg_system_syslogserver"),
			},
			"fortios_ftpproxy_explicit": {
				Tok: makeResource(mainMod, "fortios_ftpproxy_explicit"),
			},
			"fortios_icap_profile": {
				Tok: makeResource(mainMod, "fortios_icap_profile"),
			},
			"fortios_icap_server": {
				Tok: makeResource(mainMod, "fortios_icap_server"),
			},
			"fortios_icap_servergroup": {
				Tok: makeResource(mainMod, "fortios_icap_servergroup"),
			},
			"fortios_ips_custom": {
				Tok: makeResource(mainMod, "fortios_ips_custom"),
			},
			"fortios_ips_decoder": {
				Tok: makeResource(mainMod, "fortios_ips_decoder"),
			},
			"fortios_ips_global": {
				Tok: makeResource(mainMod, "fortios_ips_global"),
			},
			"fortios_ips_rule": {
				Tok: makeResource(mainMod, "fortios_ips_rule"),
			},
			"fortios_ips_rulesettings": {
				Tok: makeResource(mainMod, "fortios_ips_rulesettings"),
			},
			"fortios_ips_sensor": {
				Tok: makeResource(mainMod, "fortios_ips_sensor"),
			},
			"fortios_ips_settings": {
				Tok: makeResource(mainMod, "fortios_ips_settings"),
			},
			"fortios_ips_viewmap": {
				Tok: makeResource(mainMod, "fortios_ips_viewmap"),
			},
			"fortios_json_generic_api": {
				Tok: makeResource(mainMod, "fortios_json_generic_api"),
			},
			"fortios_log_customfield": {
				Tok: makeResource(mainMod, "fortios_log_customfield"),
			},
			"fortios_log_eventfilter": {
				Tok: makeResource(mainMod, "fortios_log_eventfilter"),
			},
			"fortios_log_guidisplay": {
				Tok: makeResource(mainMod, "fortios_log_guidisplay"),
			},
			"fortios_log_setting": {
				Tok: makeResource(mainMod, "fortios_log_setting"),
			},
			"fortios_log_syslog_setting": {
				Tok: makeResource(mainMod, "fortios_log_syslog_setting"),
			},
			"fortios_log_threatweight": {
				Tok: makeResource(mainMod, "fortios_log_threatweight"),
			},
			"fortios_logdisk_filter": {
				Tok: makeResource(mainMod, "fortios_logdisk_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logdisk_setting": {
				Tok: makeResource(mainMod, "fortios_logdisk_setting"),
			},
			"fortios_logfortianalyzer2_filter": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzer2_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logfortianalyzer2_overridefilter": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzer2_overridefilter"),
			},
			"fortios_logfortianalyzer2_overridesetting": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzer2_overridesetting"),
			},
			"fortios_logfortianalyzer2_setting": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzer2_setting"),
			},
			"fortios_logfortianalyzer3_filter": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzer3_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logfortianalyzer3_overridefilter": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzer3_overridefilter"),
			},
			"fortios_logfortianalyzer3_overridesetting": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzer3_overridesetting"),
			},
			"fortios_logfortianalyzer3_setting": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzer3_setting"),
			},
			"fortios_logfortianalyzer_filter": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzer_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logfortianalyzer_overridefilter": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzer_overridefilter"),
			},
			"fortios_logfortianalyzer_overridesetting": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzer_overridesetting"),
			},
			"fortios_logfortianalyzer_setting": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzer_setting"),
			},
			"fortios_logfortianalyzercloud_filter": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzercloud_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logfortianalyzercloud_overridefilter": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzercloud_overridefilter"),
			},
			"fortios_logfortianalyzercloud_overridesetting": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzercloud_overridesetting"),
			},
			"fortios_logfortianalyzercloud_setting": {
				Tok: makeResource(mainMod, "fortios_logfortianalyzercloud_setting"),
			},
			"fortios_logfortiguard_filter": {
				Tok: makeResource(mainMod, "fortios_logfortiguard_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logfortiguard_overridefilter": {
				Tok: makeResource(mainMod, "fortios_logfortiguard_overridefilter"),
			},
			"fortios_logfortiguard_overridesetting": {
				Tok: makeResource(mainMod, "fortios_logfortiguard_overridesetting"),
			},
			"fortios_logfortiguard_setting": {
				Tok: makeResource(mainMod, "fortios_logfortiguard_setting"),
			},
			"fortios_logmemory_filter": {
				Tok: makeResource(mainMod, "fortios_logmemory_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logmemory_globalsetting": {
				Tok: makeResource(mainMod, "fortios_logmemory_globalsetting"),
			},
			"fortios_logmemory_setting": {
				Tok: makeResource(mainMod, "fortios_logmemory_setting"),
			},
			"fortios_lognulldevice_filter": {
				Tok: makeResource(mainMod, "fortios_lognulldevice_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_lognulldevice_setting": {
				Tok: makeResource(mainMod, "fortios_lognulldevice_setting"),
			},
			"fortios_logsyslogd2_filter": {
				Tok: makeResource(mainMod, "fortios_logsyslogd2_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logsyslogd2_overridefilter": {
				Tok: makeResource(mainMod, "fortios_logsyslogd2_overridefilter"),
			},
			"fortios_logsyslogd2_overridesetting": {
				Tok: makeResource(mainMod, "fortios_logsyslogd2_overridesetting"),
			},
			"fortios_logsyslogd2_setting": {
				Tok: makeResource(mainMod, "fortios_logsyslogd2_setting"),
			},
			"fortios_logsyslogd3_filter": {
				Tok: makeResource(mainMod, "fortios_logsyslogd3_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logsyslogd3_overridefilter": {
				Tok: makeResource(mainMod, "fortios_logsyslogd3_overridefilter"),
			},
			"fortios_logsyslogd3_overridesetting": {
				Tok: makeResource(mainMod, "fortios_logsyslogd3_overridesetting"),
			},
			"fortios_logsyslogd3_setting": {
				Tok: makeResource(mainMod, "fortios_logsyslogd3_setting"),
			},
			"fortios_logsyslogd4_filter": {
				Tok: makeResource(mainMod, "fortios_logsyslogd4_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logsyslogd4_overridefilter": {
				Tok: makeResource(mainMod, "fortios_logsyslogd4_overridefilter"),
			},
			"fortios_logsyslogd4_overridesetting": {
				Tok: makeResource(mainMod, "fortios_logsyslogd4_overridesetting"),
			},
			"fortios_logsyslogd4_setting": {
				Tok: makeResource(mainMod, "fortios_logsyslogd4_setting"),
			},
			"fortios_logsyslogd_filter": {
				Tok: makeResource(mainMod, "fortios_logsyslogd_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logsyslogd_overridefilter": {
				Tok: makeResource(mainMod, "fortios_logsyslogd_overridefilter"),
			},
			"fortios_logsyslogd_overridesetting": {
				Tok: makeResource(mainMod, "fortios_logsyslogd_overridesetting"),
			},
			"fortios_logsyslogd_setting": {
				Tok: makeResource(mainMod, "fortios_logsyslogd_setting"),
			},
			"fortios_logtacacsaccounting2_filter": {
				Tok: makeResource(mainMod, "fortios_logtacacsaccounting2_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logtacacsaccounting2_setting": {
				Tok: makeResource(mainMod, "fortios_logtacacsaccounting2_setting"),
			},
			"fortios_logtacacsaccounting3_filter": {
				Tok: makeResource(mainMod, "fortios_logtacacsaccounting3_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logtacacsaccounting3_setting": {
				Tok: makeResource(mainMod, "fortios_logtacacsaccounting3_setting"),
			},
			"fortios_logtacacsaccounting_filter": {
				Tok: makeResource(mainMod, "fortios_logtacacsaccounting_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logtacacsaccounting_setting": {
				Tok: makeResource(mainMod, "fortios_logtacacsaccounting_setting"),
			},
			"fortios_logwebtrends_filter": {
				Tok: makeResource(mainMod, "fortios_logwebtrends_filter"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"filter": {
						CSharpName: "FilterDefinition",
					},
				},
			},
			"fortios_logwebtrends_setting": {
				Tok: makeResource(mainMod, "fortios_logwebtrends_setting"),
			},
			"fortios_networking_interface_port": {
				Tok: makeResource(mainMod, "fortios_networking_interface_port"),
			},
			"fortios_networking_route_static": {
				Tok: makeResource(mainMod, "fortios_networking_route_static"),
			},
			"fortios_nsxt_servicechain": {
				Tok: makeResource(mainMod, "fortios_nsxt_servicechain"),
			},
			"fortios_nsxt_setting": {
				Tok: makeResource(mainMod, "fortios_nsxt_setting"),
			},
			"fortios_report_chart": {
				Tok: makeResource(mainMod, "fortios_report_chart"),
			},
			"fortios_report_dataset": {
				Tok: makeResource(mainMod, "fortios_report_dataset"),
			},
			"fortios_report_layout": {
				Tok: makeResource(mainMod, "fortios_report_layout"),
			},
			"fortios_report_setting": {
				Tok: makeResource(mainMod, "fortios_report_setting"),
			},
			"fortios_report_style": {
				Tok: makeResource(mainMod, "fortios_report_style"),
			},
			"fortios_report_theme": {
				Tok: makeResource(mainMod, "fortios_report_theme"),
			},
			"fortios_router_accesslist": {
				Tok: makeResource(mainMod, "fortios_router_accesslist"),
			},
			"fortios_router_accesslist6": {
				Tok: makeResource(mainMod, "fortios_router_accesslist6"),
			},
			"fortios_router_aspathlist": {
				Tok: makeResource(mainMod, "fortios_router_aspathlist"),
			},
			"fortios_router_authpath": {
				Tok: makeResource(mainMod, "fortios_router_authpath"),
			},
			"fortios_router_bfd": {
				Tok: makeResource(mainMod, "fortios_router_bfd"),
			},
			"fortios_router_bfd6": {
				Tok: makeResource(mainMod, "fortios_router_bfd6"),
			},
			"fortios_router_bgp": {
				Tok: makeResource(mainMod, "fortios_router_bgp"),
			},
			"fortios_router_communitylist": {
				Tok: makeResource(mainMod, "fortios_router_communitylist"),
			},
			"fortios_router_isis": {
				Tok: makeResource(mainMod, "fortios_router_isis"),
			},
			"fortios_router_keychain": {
				Tok: makeResource(mainMod, "fortios_router_keychain"),
			},
			"fortios_router_multicast": {
				Tok: makeResource(mainMod, "fortios_router_multicast"),
			},
			"fortios_router_multicast6": {
				Tok: makeResource(mainMod, "fortios_router_multicast6"),
			},
			"fortios_router_multicastflow": {
				Tok: makeResource(mainMod, "fortios_router_multicastflow"),
			},
			"fortios_router_ospf": {
				Tok: makeResource(mainMod, "fortios_router_ospf"),
			},
			"fortios_router_ospf6": {
				Tok: makeResource(mainMod, "fortios_router_ospf6"),
			},
			"fortios_router_policy": {
				Tok: makeResource(mainMod, "fortios_router_policy"),
			},
			"fortios_router_policy6": {
				Tok: makeResource(mainMod, "fortios_router_policy6"),
			},
			"fortios_router_prefixlist": {
				Tok: makeResource(mainMod, "fortios_router_prefixlist"),
			},
			"fortios_router_prefixlist6": {
				Tok: makeResource(mainMod, "fortios_router_prefixlist6"),
			},
			"fortios_router_rip": {
				Tok: makeResource(mainMod, "fortios_router_rip"),
			},
			"fortios_router_ripng": {
				Tok: makeResource(mainMod, "fortios_router_ripng"),
			},
			"fortios_router_routemap": {
				Tok: makeResource(mainMod, "fortios_router_routemap"),
			},
			"fortios_router_setting": {
				Tok: makeResource(mainMod, "fortios_router_setting"),
			},
			"fortios_router_static": {
				Tok: makeResource(mainMod, "fortios_router_static"),
			},
			"fortios_router_static6": {
				Tok: makeResource(mainMod, "fortios_router_static6"),
			},
			"fortios_routerbgp_neighbor": {
				Tok: makeResource(mainMod, "fortios_routerbgp_neighbor"),
			},
			"fortios_routerbgp_network": {
				Tok: makeResource(mainMod, "fortios_routerbgp_network"),
			},
			"fortios_routerbgp_network6": {
				Tok: makeResource(mainMod, "fortios_routerbgp_network6"),
			},
			"fortios_routerospf6_ospf6interface": {
				Tok: makeResource(mainMod, "fortios_routerospf6_ospf6interface"),
			},
			"fortios_routerospf_neighbor": {
				Tok: makeResource(mainMod, "fortios_routerospf_neighbor"),
			},
			"fortios_routerospf_network": {
				Tok: makeResource(mainMod, "fortios_routerospf_network"),
			},
			"fortios_routerospf_ospfinterface": {
				Tok: makeResource(mainMod, "fortios_routerospf_ospfinterface"),
			},
			"fortios_sctpfilter_profile": {
				Tok: makeResource(mainMod, "fortios_sctpfilter_profile"),
			},
			"fortios_spamfilter_bwl": {
				Tok: makeResource(mainMod, "fortios_spamfilter_bwl"),
			},
			"fortios_spamfilter_bword": {
				Tok: makeResource(mainMod, "fortios_spamfilter_bword"),
			},
			"fortios_spamfilter_dnsbl": {
				Tok: makeResource(mainMod, "fortios_spamfilter_dnsbl"),
			},
			"fortios_spamfilter_fortishield": {
				Tok: makeResource(mainMod, "fortios_spamfilter_fortishield"),
			},
			"fortios_spamfilter_iptrust": {
				Tok: makeResource(mainMod, "fortios_spamfilter_iptrust"),
			},
			"fortios_spamfilter_mheader": {
				Tok: makeResource(mainMod, "fortios_spamfilter_mheader"),
			},
			"fortios_spamfilter_options": {
				Tok: makeResource(mainMod, "fortios_spamfilter_options"),
			},
			"fortios_spamfilter_profile": {
				Tok: makeResource(mainMod, "fortios_spamfilter_profile"),
			},
			"fortios_sshfilter_profile": {
				Tok: makeResource(mainMod, "fortios_sshfilter_profile"),
			},
			"fortios_switchcontroller_8021Xsettings": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_settings8021X"),
			},
			"fortios_switchcontroller_customcommand": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_customcommand"),
			},
			"fortios_switchcontroller_dynamicportpolicy": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_dynamicportpolicy"),
			},
			"fortios_switchcontroller_flowtracking": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_flowtracking"),
			},
			"fortios_switchcontroller_fortilinksettings": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_fortilinksettings"),
			},
			"fortios_switchcontroller_global": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_global"),
			},
			"fortios_switchcontroller_igmpsnooping": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_igmpsnooping"),
			},
			"fortios_switchcontroller_lldpprofile": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_lldpprofile"),
			},
			"fortios_switchcontroller_lldpsettings": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_lldpsettings"),
			},
			"fortios_switchcontroller_location": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_location"),
			},
			"fortios_switchcontroller_macsyncsettings": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_macsyncsettings"),
			},
			"fortios_switchcontroller_managedswitch": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_managedswitch"),
			},
			"fortios_switchcontroller_nacdevice": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_nacdevice"),
			},
			"fortios_switchcontroller_nacsettings": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_nacsettings"),
			},
			"fortios_switchcontroller_networkmonitorsettings": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_networkmonitorsettings"),
			},
			"fortios_switchcontroller_portpolicy": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_portpolicy"),
			},
			"fortios_switchcontroller_quarantine": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_quarantine"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"quarantine": {
						CSharpName: "Data",
					},
				},
			},
			"fortios_switchcontroller_remotelog": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_remotelog"),
			},
			"fortios_switchcontroller_sflow": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_sflow"),
			},
			"fortios_switchcontroller_snmpcommunity": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_snmpcommunity"),
			},
			"fortios_switchcontroller_snmpsysinfo": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_snmpsysinfo"),
			},
			"fortios_switchcontroller_snmptrapthreshold": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_snmptrapthreshold"),
			},
			"fortios_switchcontroller_snmpuser": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_snmpuser"),
			},
			"fortios_switchcontroller_stormcontrol": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_stormcontrol"),
			},
			"fortios_switchcontroller_stormcontrolpolicy": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_stormcontrolpolicy"),
			},
			"fortios_switchcontroller_stpinstance": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_stpinstance"),
			},
			"fortios_switchcontroller_stpsettings": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_stpsettings"),
			},
			"fortios_switchcontroller_switchgroup": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_switchgroup"),
			},
			"fortios_switchcontroller_switchinterfacetag": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_switchinterfacetag"),
			},
			"fortios_switchcontroller_switchlog": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_switchlog"),
			},
			"fortios_switchcontroller_switchprofile": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_switchprofile"),
			},
			"fortios_switchcontroller_system": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_system"),
			},
			"fortios_switchcontroller_trafficpolicy": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_trafficpolicy"),
			},
			"fortios_switchcontroller_trafficsniffer": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_trafficsniffer"),
			},
			"fortios_switchcontroller_virtualportpool": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_virtualportpool"),
			},
			"fortios_switchcontroller_vlan": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_vlan"),
			},
			"fortios_switchcontroller_vlanpolicy": {
				Tok: makeResource(mainMod, "fortios_switchcontroller_vlanpolicy"),
			},
			"fortios_switchcontrollerautoconfig_custom": {
				Tok: makeResource(mainMod, "fortios_switchcontrollerautoconfig_custom"),
			},
			"fortios_switchcontrollerautoconfig_default": {
				Tok: makeResource(mainMod, "fortios_switchcontrollerautoconfig_default"),
			},
			"fortios_switchcontrollerautoconfig_policy": {
				Tok: makeResource(mainMod, "fortios_switchcontrollerautoconfig_policy"),
			},
			"fortios_switchcontrollerinitialconfig_template": {
				Tok: makeResource(mainMod, "fortios_switchcontrollerinitialconfig_template"),
			},
			"fortios_switchcontrollerinitialconfig_vlans": {
				Tok: makeResource(mainMod, "fortios_switchcontrollerinitialconfig_vlans"),
			},
			"fortios_switchcontrollerptp_policy": {
				Tok: makeResource(mainMod, "fortios_switchcontrollerptp_policy"),
			},
			"fortios_switchcontrollerptp_settings": {
				Tok: makeResource(mainMod, "fortios_switchcontrollerptp_settings"),
			},
			"fortios_switchcontrollerqos_dot1pmap": {
				Tok: makeResource(mainMod, "fortios_switchcontrollerqos_dot1pmap"),
			},
			"fortios_switchcontrollerqos_ipdscpmap": {
				Tok: makeResource(mainMod, "fortios_switchcontrollerqos_ipdscpmap"),
			},
			"fortios_switchcontrollerqos_qospolicy": {
				Tok: makeResource(mainMod, "fortios_switchcontrollerqos_qospolicy"),
			},
			"fortios_switchcontrollerqos_queuepolicy": {
				Tok: makeResource(mainMod, "fortios_switchcontrollerqos_queuepolicy"),
			},
			"fortios_switchcontrollersecuritypolicy_8021X": {
				Tok: makeResource(mainMod, "fortios_switchcontrollersecuritypolicy_policy8021X"),
			},
			"fortios_switchcontrollersecuritypolicy_captiveportal": {
				Tok: makeResource(mainMod, "fortios_switchcontrollersecuritypolicy_captiveportal"),
			},
			"fortios_switchcontrollersecuritypolicy_localaccess": {
				Tok: makeResource(mainMod, "fortios_switchcontrollersecuritypolicy_localaccess"),
			},
			"fortios_system3gmodem_custom": {
				Tok: makeResource(mainMod, "fortios_system3gmodem_custom"),
			},
			"fortios_system_accprofile": {
				Tok: makeResource(mainMod, "fortios_system_accprofile"),
			},
			"fortios_system_acme": {
				Tok: makeResource(mainMod, "fortios_system_acme"),
			},
			"fortios_system_admin": {
				Tok: makeResource(mainMod, "fortios_system_admin"),
			},
			"fortios_system_admin_administrator": {
				Tok: makeResource(mainMod, "fortios_system_admin_administrator"),
			},
			"fortios_system_admin_profiles": {
				Tok: makeResource(mainMod, "fortios_system_admin_profiles"),
			},
			"fortios_system_affinityinterrupt": {
				Tok: makeResource(mainMod, "fortios_system_affinityinterrupt"),
			},
			"fortios_system_affinitypacketredistribution": {
				Tok: makeResource(mainMod, "fortios_system_affinitypacketredistribution"),
			},
			"fortios_system_alarm": {
				Tok: makeResource(mainMod, "fortios_system_alarm"),
			},
			"fortios_system_alias": {
				Tok: makeResource(mainMod, "fortios_system_alias"),
			},
			"fortios_system_apiuser": {
				Tok: makeResource(mainMod, "fortios_system_apiuser"),
			},
			"fortios_system_apiuser_setting": {
				Tok: makeResource(mainMod, "fortios_system_apiuser_setting"),
			},
			"fortios_system_arptable": {
				Tok: makeResource(mainMod, "fortios_system_arptable"),
			},
			"fortios_system_autoinstall": {
				Tok: makeResource(mainMod, "fortios_system_autoinstall"),
			},
			"fortios_system_automationaction": {
				Tok: makeResource(mainMod, "fortios_system_automationaction"),
			},
			"fortios_system_automationdestination": {
				Tok: makeResource(mainMod, "fortios_system_automationdestination"),
			},
			"fortios_system_automationstitch": {
				Tok: makeResource(mainMod, "fortios_system_automationstitch"),
			},
			"fortios_system_automationtrigger": {
				Tok: makeResource(mainMod, "fortios_system_automationtrigger"),
			},
			"fortios_system_autoscript": {
				Tok: makeResource(mainMod, "fortios_system_autoscript"),
			},
			"fortios_system_centralmanagement": {
				Tok: makeResource(mainMod, "fortios_system_centralmanagement"),
			},
			"fortios_system_clustersync": {
				Tok: makeResource(mainMod, "fortios_system_clustersync"),
			},
			"fortios_system_console": {
				Tok: makeResource(mainMod, "fortios_system_console"),
			},
			"fortios_system_csf": {
				Tok: makeResource(mainMod, "fortios_system_csf"),
			},
			"fortios_system_customlanguage": {
				Tok: makeResource(mainMod, "fortios_system_customlanguage"),
			},
			"fortios_system_ddns": {
				Tok: makeResource(mainMod, "fortios_system_ddns"),
			},
			"fortios_system_dedicatedmgmt": {
				Tok: makeResource(mainMod, "fortios_system_dedicatedmgmt"),
			},
			"fortios_system_dns": {
				Tok: makeResource(mainMod, "fortios_system_dns"),
			},
			"fortios_system_dns64": {
				Tok: makeResource(mainMod, "fortios_system_dns64"),
			},
			"fortios_system_dnsdatabase": {
				Tok: makeResource(mainMod, "fortios_system_dnsdatabase"),
			},
			"fortios_system_dnsserver": {
				Tok: makeResource(mainMod, "fortios_system_dnsserver"),
			},
			"fortios_system_dscpbasedpriority": {
				Tok: makeResource(mainMod, "fortios_system_dscpbasedpriority"),
			},
			"fortios_system_emailserver": {
				Tok: makeResource(mainMod, "fortios_system_emailserver"),
			},
			"fortios_system_externalresource": {
				Tok: makeResource(mainMod, "fortios_system_externalresource"),
			},
			"fortios_system_federatedupgrade": {
				Tok: makeResource(mainMod, "fortios_system_federatedupgrade"),
			},
			"fortios_system_fipscc": {
				Tok: makeResource(mainMod, "fortios_system_fipscc"),
			},
			"fortios_system_fm": {
				Tok: makeResource(mainMod, "fortios_system_fm"),
			},
			"fortios_system_fortiai": {
				Tok: makeResource(mainMod, "fortios_system_fortiai"),
			},
			"fortios_system_fortiguard": {
				Tok: makeResource(mainMod, "fortios_system_fortiguard"),
			},
			"fortios_system_fortimanager": {
				Tok: makeResource(mainMod, "fortios_system_fortimanager"),
			},
			"fortios_system_fortindr": {
				Tok: makeResource(mainMod, "fortios_system_fortindr"),
			},
			"fortios_system_fortisandbox": {
				Tok: makeResource(mainMod, "fortios_system_fortisandbox"),
			},
			"fortios_system_fssopolling": {
				Tok: makeResource(mainMod, "fortios_system_fssopolling"),
			},
			"fortios_system_ftmpush": {
				Tok: makeResource(mainMod, "fortios_system_ftmpush"),
			},
			"fortios_system_geneve": {
				Tok: makeResource(mainMod, "fortios_system_geneve"),
			},
			"fortios_system_geoipcountry": {
				Tok: makeResource(mainMod, "fortios_system_geoipcountry"),
			},
			"fortios_system_geoipoverride": {
				Tok: makeResource(mainMod, "fortios_system_geoipoverride"),
			},
			"fortios_system_global": {
				Tok: makeResource(mainMod, "fortios_system_global"),
			},
			"fortios_system_gretunnel": {
				Tok: makeResource(mainMod, "fortios_system_gretunnel"),
			},
			"fortios_system_ha": {
				Tok: makeResource(mainMod, "fortios_system_ha"),
			},
			"fortios_system_hamonitor": {
				Tok: makeResource(mainMod, "fortios_system_hamonitor"),
			},
			"fortios_system_ike": {
				Tok: makeResource(mainMod, "fortios_system_ike"),
			},
			"fortios_system_interface": {
				Tok: makeResource(mainMod, "fortios_system_interface"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"interface": {
						CSharpName: "Data",
					},
				},
			},
			"fortios_system_ipam": {
				Tok: makeResource(mainMod, "fortios_system_ipam"),
			},
			"fortios_system_ipiptunnel": {
				Tok: makeResource(mainMod, "fortios_system_ipiptunnel"),
			},
			"fortios_system_ips": {
				Tok: makeResource(mainMod, "fortios_system_ips"),
			},
			"fortios_system_ipsecaggregate": {
				Tok: makeResource(mainMod, "fortios_system_ipsecaggregate"),
			},
			"fortios_system_ipsurlfilterdns": {
				Tok: makeResource(mainMod, "fortios_system_ipsurlfilterdns"),
			},
			"fortios_system_ipsurlfilterdns6": {
				Tok: makeResource(mainMod, "fortios_system_ipsurlfilterdns6"),
			},
			"fortios_system_ipv6neighborcache": {
				Tok: makeResource(mainMod, "fortios_system_ipv6neighborcache"),
			},
			"fortios_system_ipv6tunnel": {
				Tok: makeResource(mainMod, "fortios_system_ipv6tunnel"),
			},
			"fortios_system_license_forticare": {
				Tok: makeResource(mainMod, "fortios_system_license_forticare"),
			},
			"fortios_system_license_vdom": {
				Tok: makeResource(mainMod, "fortios_system_license_vdom"),
			},
			"fortios_system_license_vm": {
				Tok: makeResource(mainMod, "fortios_system_license_vm"),
			},
			"fortios_system_linkmonitor": {
				Tok: makeResource(mainMod, "fortios_system_linkmonitor"),
			},
			"fortios_system_ltemodem": {
				Tok: makeResource(mainMod, "fortios_system_ltemodem"),
			},
			"fortios_system_macaddresstable": {
				Tok: makeResource(mainMod, "fortios_system_macaddresstable"),
			},
			"fortios_system_managementtunnel": {
				Tok: makeResource(mainMod, "fortios_system_managementtunnel"),
			},
			"fortios_system_mobiletunnel": {
				Tok: makeResource(mainMod, "fortios_system_mobiletunnel"),
			},
			"fortios_system_modem": {
				Tok: makeResource(mainMod, "fortios_system_modem"),
			},
			"fortios_system_nat64": {
				Tok: makeResource(mainMod, "fortios_system_nat64"),
			},
			"fortios_system_ndproxy": {
				Tok: makeResource(mainMod, "fortios_system_ndproxy"),
			},
			"fortios_system_netflow": {
				Tok: makeResource(mainMod, "fortios_system_netflow"),
			},
			"fortios_system_networkvisibility": {
				Tok: makeResource(mainMod, "fortios_system_networkvisibility"),
			},
			"fortios_system_npu": {
				Tok: makeResource(mainMod, "fortios_system_npu"),
			},
			"fortios_system_ntp": {
				Tok: makeResource(mainMod, "fortios_system_ntp"),
			},
			"fortios_system_objecttagging": {
				Tok: makeResource(mainMod, "fortios_system_objecttagging"),
			},
			"fortios_system_passwordpolicy": {
				Tok: makeResource(mainMod, "fortios_system_passwordpolicy"),
			},
			"fortios_system_passwordpolicyguestadmin": {
				Tok: makeResource(mainMod, "fortios_system_passwordpolicyguestadmin"),
			},
			"fortios_system_physicalswitch": {
				Tok: makeResource(mainMod, "fortios_system_physicalswitch"),
			},
			"fortios_system_pppoeinterface": {
				Tok: makeResource(mainMod, "fortios_system_pppoeinterface"),
			},
			"fortios_system_proberesponse": {
				Tok: makeResource(mainMod, "fortios_system_proberesponse"),
			},
			"fortios_system_proxyarp": {
				Tok: makeResource(mainMod, "fortios_system_proxyarp"),
			},
			"fortios_system_ptp": {
				Tok: makeResource(mainMod, "fortios_system_ptp"),
			},
			"fortios_system_replacemsggroup": {
				Tok: makeResource(mainMod, "fortios_system_replacemsggroup"),
			},
			"fortios_system_replacemsgimage": {
				Tok: makeResource(mainMod, "fortios_system_replacemsgimage"),
			},
			"fortios_system_resourcelimits": {
				Tok: makeResource(mainMod, "fortios_system_resourcelimits"),
			},
			"fortios_system_saml": {
				Tok: makeResource(mainMod, "fortios_system_saml"),
			},
			"fortios_system_sdnconnector": {
				Tok: makeResource(mainMod, "fortios_system_sdnconnector"),
			},
			"fortios_system_sdwan": {
				Tok: makeResource(mainMod, "fortios_system_sdwan"),
			},
			"fortios_system_sessionhelper": {
				Tok: makeResource(mainMod, "fortios_system_sessionhelper"),
			},
			"fortios_system_sessionttl": {
				Tok: makeResource(mainMod, "fortios_system_sessionttl"),
			},
			"fortios_system_setting_dns": {
				Tok: makeResource(mainMod, "fortios_system_setting_dns"),
			},
			"fortios_system_setting_global": {
				Tok: makeResource(mainMod, "fortios_system_setting_global"),
			},
			"fortios_system_setting_ntp": {
				Tok: makeResource(mainMod, "fortios_system_setting_ntp"),
			},
			"fortios_system_settings": {
				Tok: makeResource(mainMod, "fortios_system_settings"),
			},
			"fortios_system_sflow": {
				Tok: makeResource(mainMod, "fortios_system_sflow"),
			},
			"fortios_system_sittunnel": {
				Tok: makeResource(mainMod, "fortios_system_sittunnel"),
			},
			"fortios_system_smsserver": {
				Tok: makeResource(mainMod, "fortios_system_smsserver"),
			},
			"fortios_system_speedtestschedule": {
				Tok: makeResource(mainMod, "fortios_system_speedtestschedule"),
			},
			"fortios_system_speedtestserver": {
				Tok: makeResource(mainMod, "fortios_system_speedtestserver"),
			},
			"fortios_system_ssoadmin": {
				Tok: makeResource(mainMod, "fortios_system_ssoadmin"),
			},
			"fortios_system_ssoforticloudadmin": {
				Tok: makeResource(mainMod, "fortios_system_ssoforticloudadmin"),
			},
			"fortios_system_standalonecluster": {
				Tok: makeResource(mainMod, "fortios_system_standalonecluster"),
			},
			"fortios_system_storage": {
				Tok: makeResource(mainMod, "fortios_system_storage"),
			},
			"fortios_system_stp": {
				Tok: makeResource(mainMod, "fortios_system_stp"),
			},
			"fortios_system_switchinterface": {
				Tok: makeResource(mainMod, "fortios_system_switchinterface"),
			},
			"fortios_system_tosbasedpriority": {
				Tok: makeResource(mainMod, "fortios_system_tosbasedpriority"),
			},
			"fortios_system_vdom": {
				Tok: makeResource(mainMod, "fortios_system_vdom"),
			},
			"fortios_system_vdom_setting": {
				Tok: makeResource(mainMod, "fortios_system_vdom_setting"),
			},
			"fortios_system_vdomdns": {
				Tok: makeResource(mainMod, "fortios_system_vdomdns"),
			},
			"fortios_system_vdomexception": {
				Tok: makeResource(mainMod, "fortios_system_vdomexception"),
			},
			"fortios_system_vdomlink": {
				Tok: makeResource(mainMod, "fortios_system_vdomlink"),
			},
			"fortios_system_vdomnetflow": {
				Tok: makeResource(mainMod, "fortios_system_vdomnetflow"),
			},
			"fortios_system_vdomproperty": {
				Tok: makeResource(mainMod, "fortios_system_vdomproperty"),
			},
			"fortios_system_vdomradiusserver": {
				Tok: makeResource(mainMod, "fortios_system_vdomradiusserver"),
			},
			"fortios_system_vdomsflow": {
				Tok: makeResource(mainMod, "fortios_system_vdomsflow"),
			},
			"fortios_system_virtualswitch": {
				Tok: makeResource(mainMod, "fortios_system_virtualswitch"),
			},
			"fortios_system_virtualwanlink": {
				Tok: makeResource(mainMod, "fortios_system_virtualwanlink"),
			},
			"fortios_system_virtualwirepair": {
				Tok: makeResource(mainMod, "fortios_system_virtualwirepair"),
			},
			"fortios_system_vnetunnel": {
				Tok: makeResource(mainMod, "fortios_system_vnetunnel"),
			},
			"fortios_system_vxlan": {
				Tok: makeResource(mainMod, "fortios_system_vxlan"),
			},
			"fortios_system_wccp": {
				Tok: makeResource(mainMod, "fortios_system_wccp"),
			},
			"fortios_system_zone": {
				Tok: makeResource(mainMod, "fortios_system_zone"),
			},
			"fortios_systemautoupdate_pushupdate": {
				Tok: makeResource(mainMod, "fortios_systemautoupdate_pushupdate"),
			},
			"fortios_systemautoupdate_schedule": {
				Tok: makeResource(mainMod, "fortios_systemautoupdate_schedule"),
			},
			"fortios_systemautoupdate_tunneling": {
				Tok: makeResource(mainMod, "fortios_systemautoupdate_tunneling"),
			},
			"fortios_systemdhcp6_server": {
				Tok: makeResource(mainMod, "fortios_systemdhcp6_server"),
			},
			"fortios_systemdhcp_server": {
				Tok: makeResource(mainMod, "fortios_systemdhcp_server"),
			},
			"fortios_systemlldp_networkpolicy": {
				Tok: makeResource(mainMod, "fortios_systemlldp_networkpolicy"),
			},
			"fortios_systemreplacemsg_admin": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_admin"),
			},
			"fortios_systemreplacemsg_alertmail": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_alertmail"),
			},
			"fortios_systemreplacemsg_auth": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_auth"),
			},
			"fortios_systemreplacemsg_automation": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_automation"),
			},
			"fortios_systemreplacemsg_devicedetectionportal": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_devicedetectionportal"),
			},
			"fortios_systemreplacemsg_ec": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_ec"),
			},
			"fortios_systemreplacemsg_fortiguardwf": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_fortiguardwf"),
			},
			"fortios_systemreplacemsg_ftp": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_ftp"),
			},
			"fortios_systemreplacemsg_http": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_http"),
			},
			"fortios_systemreplacemsg_icap": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_icap"),
			},
			"fortios_systemreplacemsg_mail": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_mail"),
			},
			"fortios_systemreplacemsg_nacquar": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_nacquar"),
			},
			"fortios_systemreplacemsg_nntp": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_nntp"),
			},
			"fortios_systemreplacemsg_spam": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_spam"),
			},
			"fortios_systemreplacemsg_sslvpn": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_sslvpn"),
			},
			"fortios_systemreplacemsg_trafficquota": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_trafficquota"),
			},
			"fortios_systemreplacemsg_utm": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_utm"),
			},
			"fortios_systemreplacemsg_webproxy": {
				Tok: makeResource(mainMod, "fortios_systemreplacemsg_webproxy"),
			},
			"fortios_systemsnmp_community": {
				Tok: makeResource(mainMod, "fortios_systemsnmp_community"),
			},
			"fortios_systemsnmp_mibview": {
				Tok: makeResource(mainMod, "fortios_systemsnmp_mibview"),
			},
			"fortios_systemsnmp_sysinfo": {
				Tok: makeResource(mainMod, "fortios_systemsnmp_sysinfo"),
			},
			"fortios_systemsnmp_user": {
				Tok: makeResource(mainMod, "fortios_systemsnmp_user"),
			},
			"fortios_user_adgrp": {
				Tok: makeResource(mainMod, "fortios_user_adgrp"),
			},
			"fortios_user_certificate": {
				Tok: makeResource(mainMod, "fortios_user_certificate"),
			},
			"fortios_user_device": {
				Tok: makeResource(mainMod, "fortios_user_device"),
			},
			"fortios_user_deviceaccesslist": {
				Tok: makeResource(mainMod, "fortios_user_deviceaccesslist"),
			},
			"fortios_user_devicecategory": {
				Tok: makeResource(mainMod, "fortios_user_devicecategory"),
			},
			"fortios_user_devicegroup": {
				Tok: makeResource(mainMod, "fortios_user_devicegroup"),
			},
			"fortios_user_domaincontroller": {
				Tok: makeResource(mainMod, "fortios_user_domaincontroller"),
			},
			"fortios_user_exchange": {
				Tok: makeResource(mainMod, "fortios_user_exchange"),
			},
			"fortios_user_fortitoken": {
				Tok: makeResource(mainMod, "fortios_user_fortitoken"),
			},
			"fortios_user_fsso": {
				Tok: makeResource(mainMod, "fortios_user_fsso"),
			},
			"fortios_user_fssopolling": {
				Tok: makeResource(mainMod, "fortios_user_fssopolling"),
			},
			"fortios_user_group": {
				Tok: makeResource(mainMod, "fortios_user_group"),
			},
			"fortios_user_krbkeytab": {
				Tok: makeResource(mainMod, "fortios_user_krbkeytab"),
			},
			"fortios_user_ldap": {
				Tok: makeResource(mainMod, "fortios_user_ldap"),
			},
			"fortios_user_local": {
				Tok: makeResource(mainMod, "fortios_user_local"),
			},
			"fortios_user_nacpolicy": {
				Tok: makeResource(mainMod, "fortios_user_nacpolicy"),
			},
			"fortios_user_passwordpolicy": {
				Tok: makeResource(mainMod, "fortios_user_passwordpolicy"),
			},
			"fortios_user_peer": {
				Tok: makeResource(mainMod, "fortios_user_peer"),
			},
			"fortios_user_peergrp": {
				Tok: makeResource(mainMod, "fortios_user_peergrp"),
			},
			"fortios_user_pop3": {
				Tok: makeResource(mainMod, "fortios_user_pop3"),
			},
			"fortios_user_quarantine": {
				Tok: makeResource(mainMod, "fortios_user_quarantine"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"quarantine": {
						CSharpName: "Data",
					},
				},
			},
			"fortios_user_radius": {
				Tok: makeResource(mainMod, "fortios_user_radius"),
			},
			"fortios_user_saml": {
				Tok: makeResource(mainMod, "fortios_user_saml"),
			},
			"fortios_user_securityexemptlist": {
				Tok: makeResource(mainMod, "fortios_user_securityexemptlist"),
			},
			"fortios_user_setting": {
				Tok: makeResource(mainMod, "fortios_user_setting"),
			},
			"fortios_user_tacacs": {
				Tok: makeResource(mainMod, "fortios_user_tacacs"),
			},
			"fortios_videofilter_profile": {
				Tok: makeResource(mainMod, "fortios_videofilter_profile"),
			},
			"fortios_videofilter_youtubechannelfilter": {
				Tok: makeResource(mainMod, "fortios_videofilter_youtubechannelfilter"),
			},
			"fortios_videofilter_youtubekey": {
				Tok: makeResource(mainMod, "fortios_videofilter_youtubekey"),
			},
			"fortios_voip_profile": {
				Tok: makeResource(mainMod, "fortios_voip_profile"),
			},
			"fortios_vpn_l2tp": {
				Tok: makeResource(mainMod, "fortios_vpn_l2tp"),
			},
			"fortios_vpn_ocvpn": {
				Tok: makeResource(mainMod, "fortios_vpn_ocvpn"),
			},
			"fortios_vpn_pptp": {
				Tok: makeResource(mainMod, "fortios_vpn_pptp"),
			},
			"fortios_vpncertificate_ca": {
				Tok: makeResource(mainMod, "fortios_vpncertificate_ca"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"ca": {
						CSharpName: "Certificate",
					},
				},
			},
			"fortios_vpncertificate_crl": {
				Tok: makeResource(mainMod, "fortios_vpncertificate_crl"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"crl": {
						CSharpName: "Data",
					},
				},
			},
			"fortios_vpncertificate_local": {
				Tok: makeResource(mainMod, "fortios_vpncertificate_local"),
			},
			"fortios_vpncertificate_ocspserver": {
				Tok: makeResource(mainMod, "fortios_vpncertificate_ocspserver"),
			},
			"fortios_vpncertificate_remote": {
				Tok: makeResource(mainMod, "fortios_vpncertificate_remote"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"remote": {
						CSharpName: "Data",
					},
				},
			},
			"fortios_vpncertificate_setting": {
				Tok: makeResource(mainMod, "fortios_vpncertificate_setting"),
			},
			"fortios_vpnipsec_concentrator": {
				Tok: makeResource(mainMod, "fortios_vpnipsec_concentrator"),
			},
			"fortios_vpnipsec_fec": {
				Tok: makeResource(mainMod, "fortios_vpnipsec_fec"),
			},
			"fortios_vpnipsec_forticlient": {
				Tok: makeResource(mainMod, "fortios_vpnipsec_forticlient"),
			},
			"fortios_vpnipsec_manualkey": {
				Tok: makeResource(mainMod, "fortios_vpnipsec_manualkey"),
			},
			"fortios_vpnipsec_manualkeyinterface": {
				Tok: makeResource(mainMod, "fortios_vpnipsec_manualkeyinterface"),
			},
			"fortios_vpnipsec_phase1": {
				Tok: makeResource(mainMod, "fortios_vpnipsec_phase1"),
			},
			"fortios_vpnipsec_phase1interface": {
				Tok: makeResource(mainMod, "fortios_vpnipsec_phase1interface"),
			},
			"fortios_vpnipsec_phase2": {
				Tok: makeResource(mainMod, "fortios_vpnipsec_phase2"),
			},
			"fortios_vpnipsec_phase2interface": {
				Tok: makeResource(mainMod, "fortios_vpnipsec_phase2interface"),
			},
			"fortios_vpnssl_client": {
				Tok: makeResource(mainMod, "fortios_vpnssl_client"),
			},
			"fortios_vpnssl_settings": {
				Tok: makeResource(mainMod, "fortios_vpnssl_settings"),
			},
			"fortios_vpnsslweb_hostchecksoftware": {
				Tok: makeResource(mainMod, "fortios_vpnsslweb_hostchecksoftware"),
			},
			"fortios_vpnsslweb_portal": {
				Tok: makeResource(mainMod, "fortios_vpnsslweb_portal"),
			},
			"fortios_vpnsslweb_realm": {
				Tok: makeResource(mainMod, "fortios_vpnsslweb_realm"),
			},
			"fortios_vpnsslweb_userbookmark": {
				Tok: makeResource(mainMod, "fortios_vpnsslweb_userbookmark"),
			},
			"fortios_vpnsslweb_usergroupbookmark": {
				Tok: makeResource(mainMod, "fortios_vpnsslweb_usergroupbookmark"),
			},
			"fortios_waf_mainclass": {
				Tok: makeResource(mainMod, "fortios_waf_mainclass"),
			},
			"fortios_waf_profile": {
				Tok: makeResource(mainMod, "fortios_waf_profile"),
			},
			"fortios_waf_signature": {
				Tok: makeResource(mainMod, "fortios_waf_signature"),
			},
			"fortios_waf_subclass": {
				Tok: makeResource(mainMod, "fortios_waf_subclass"),
			},
			"fortios_wanopt_authgroup": {
				Tok: makeResource(mainMod, "fortios_wanopt_authgroup"),
			},
			"fortios_wanopt_cacheservice": {
				Tok: makeResource(mainMod, "fortios_wanopt_cacheservice"),
			},
			"fortios_wanopt_contentdeliverynetworkrule": {
				Tok: makeResource(mainMod, "fortios_wanopt_contentdeliverynetworkrule"),
			},
			"fortios_wanopt_peer": {
				Tok: makeResource(mainMod, "fortios_wanopt_peer"),
			},
			"fortios_wanopt_profile": {
				Tok: makeResource(mainMod, "fortios_wanopt_profile"),
			},
			"fortios_wanopt_remotestorage": {
				Tok: makeResource(mainMod, "fortios_wanopt_remotestorage"),
			},
			"fortios_wanopt_settings": {
				Tok: makeResource(mainMod, "fortios_wanopt_settings"),
			},
			"fortios_wanopt_webcache": {
				Tok: makeResource(mainMod, "fortios_wanopt_webcache"),
			},
			"fortios_webfilter_content": {
				Tok: makeResource(mainMod, "fortios_webfilter_content"),
			},
			"fortios_webfilter_contentheader": {
				Tok: makeResource(mainMod, "fortios_webfilter_contentheader"),
			},
			"fortios_webfilter_fortiguard": {
				Tok: makeResource(mainMod, "fortios_webfilter_fortiguard"),
			},
			"fortios_webfilter_ftgdlocalcat": {
				Tok: makeResource(mainMod, "fortios_webfilter_ftgdlocalcat"),
			},
			"fortios_webfilter_ftgdlocalrating": {
				Tok: makeResource(mainMod, "fortios_webfilter_ftgdlocalrating"),
			},
			"fortios_webfilter_ipsurlfiltercachesetting": {
				Tok: makeResource(mainMod, "fortios_webfilter_ipsurlfiltercachesetting"),
			},
			"fortios_webfilter_ipsurlfiltersetting": {
				Tok: makeResource(mainMod, "fortios_webfilter_ipsurlfiltersetting"),
			},
			"fortios_webfilter_ipsurlfiltersetting6": {
				Tok: makeResource(mainMod, "fortios_webfilter_ipsurlfiltersetting6"),
			},
			"fortios_webfilter_override": {
				Tok: makeResource(mainMod, "fortios_webfilter_override"),
			},
			"fortios_webfilter_profile": {
				Tok: makeResource(mainMod, "fortios_webfilter_profile"),
			},
			"fortios_webfilter_searchengine": {
				Tok: makeResource(mainMod, "fortios_webfilter_searchengine"),
			},
			"fortios_webfilter_urlfilter": {
				Tok: makeResource(mainMod, "fortios_webfilter_urlfilter"),
			},
			"fortios_webproxy_debugurl": {
				Tok: makeResource(mainMod, "fortios_webproxy_debugurl"),
			},
			"fortios_webproxy_explicit": {
				Tok: makeResource(mainMod, "fortios_webproxy_explicit"),
			},
			"fortios_webproxy_forwardserver": {
				Tok: makeResource(mainMod, "fortios_webproxy_forwardserver"),
			},
			"fortios_webproxy_forwardservergroup": {
				Tok: makeResource(mainMod, "fortios_webproxy_forwardservergroup"),
			},
			"fortios_webproxy_global": {
				Tok: makeResource(mainMod, "fortios_webproxy_global"),
			},
			"fortios_webproxy_profile": {
				Tok: makeResource(mainMod, "fortios_webproxy_profile"),
			},
			"fortios_webproxy_urlmatch": {
				Tok: makeResource(mainMod, "fortios_webproxy_urlmatch"),
			},
			"fortios_webproxy_wisp": {
				Tok: makeResource(mainMod, "fortios_webproxy_wisp"),
			},
			"fortios_wirelesscontroller_accesscontrollist": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_accesscontrollist"),
			},
			"fortios_wirelesscontroller_address": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_address"),
			},
			"fortios_wirelesscontroller_addrgrp": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_addrgrp"),
			},
			"fortios_wirelesscontroller_apcfgprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_apcfgprofile"),
			},
			"fortios_wirelesscontroller_apstatus": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_apstatus"),
			},
			"fortios_wirelesscontroller_arrpprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_arrpprofile"),
			},
			"fortios_wirelesscontroller_bleprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_bleprofile"),
			},
			"fortios_wirelesscontroller_bonjourprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_bonjourprofile"),
			},
			"fortios_wirelesscontroller_global": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_global"),
			},
			"fortios_wirelesscontroller_intercontroller": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_intercontroller"),
			},
			"fortios_wirelesscontroller_log": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_log"),
			},
			"fortios_wirelesscontroller_mpskprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_mpskprofile"),
			},
			"fortios_wirelesscontroller_nacprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_nacprofile"),
			},
			"fortios_wirelesscontroller_qosprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_qosprofile"),
			},
			"fortios_wirelesscontroller_region": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_region"),
			},
			"fortios_wirelesscontroller_setting": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_setting"),
			},
			"fortios_wirelesscontroller_snmp": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_snmp"),
			},
			"fortios_wirelesscontroller_ssidpolicy": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_ssidpolicy"),
			},
			"fortios_wirelesscontroller_syslogprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_syslogprofile"),
			},
			"fortios_wirelesscontroller_timers": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_timers"),
			},
			"fortios_wirelesscontroller_utmprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_utmprofile"),
			},
			"fortios_wirelesscontroller_vap": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_vap"),
			},
			"fortios_wirelesscontroller_vapgroup": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_vapgroup"),
			},
			"fortios_wirelesscontroller_wagprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_wagprofile"),
			},
			"fortios_wirelesscontroller_widsprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_widsprofile"),
			},
			"fortios_wirelesscontroller_wtp": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_wtp"),
			},
			"fortios_wirelesscontroller_wtpgroup": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_wtpgroup"),
			},
			"fortios_wirelesscontroller_wtpprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontroller_wtpprofile"),
			},
			"fortios_wirelesscontrollerhotspot20_anqp3gppcellular": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_anqp3gppcellular"),
			},
			"fortios_wirelesscontrollerhotspot20_anqpipaddresstype": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_anqpipaddresstype"),
			},
			"fortios_wirelesscontrollerhotspot20_anqpnairealm": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_anqpnairealm"),
			},
			"fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype"),
			},
			"fortios_wirelesscontrollerhotspot20_anqproamingconsortium": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_anqproamingconsortium"),
			},
			"fortios_wirelesscontrollerhotspot20_anqpvenuename": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_anqpvenuename"),
			},
			"fortios_wirelesscontrollerhotspot20_anqpvenueurl": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_anqpvenueurl"),
			},
			"fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge"),
			},
			"fortios_wirelesscontrollerhotspot20_h2qpconncapability": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_h2qpconncapability"),
			},
			"fortios_wirelesscontrollerhotspot20_h2qpoperatorname": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_h2qpoperatorname"),
			},
			"fortios_wirelesscontrollerhotspot20_h2qposuprovider": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_h2qposuprovider"),
			},
			"fortios_wirelesscontrollerhotspot20_h2qposuprovidernai": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_h2qposuprovidernai"),
			},
			"fortios_wirelesscontrollerhotspot20_h2qptermsandconditions": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_h2qptermsandconditions"),
			},
			"fortios_wirelesscontrollerhotspot20_h2qpwanmetric": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_h2qpwanmetric"),
			},
			"fortios_wirelesscontrollerhotspot20_hsprofile": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_hsprofile"),
			},
			"fortios_wirelesscontrollerhotspot20_icon": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_icon"),
			},
			"fortios_wirelesscontrollerhotspot20_qosmap": {
				Tok: makeResource(mainMod, "fortios_wirelesscontrollerhotspot20_qosmap"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: makeDataSource(mainMod, "aws_ami")},
			"fortios_firewall_DoSpolicy": {
				Tok: makeDataSource(mainMod, "fortios_firewall_DoSpolicy"),
			},
			"fortios_firewall_DoSpolicy6": {
				Tok: makeDataSource(mainMod, "fortios_firewall_DoSpolicy6"),
			},
			"fortios_firewall_DoSpolicy6list": {
				Tok: makeDataSource(mainMod, "fortios_firewall_DoSpolicy6list"),
			},
			"fortios_firewall_DoSpolicylist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_DoSpolicylist"),
			},
			"fortios_firewall_address": {
				Tok: makeDataSource(mainMod, "fortios_firewall_address"),
			},
			"fortios_firewall_address6": {
				Tok: makeDataSource(mainMod, "fortios_firewall_address6"),
			},
			"fortios_firewall_address6list": {
				Tok: makeDataSource(mainMod, "fortios_firewall_address6list"),
			},
			"fortios_firewall_address6template": {
				Tok: makeDataSource(mainMod, "fortios_firewall_address6template"),
			},
			"fortios_firewall_address6templatelist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_address6templatelist"),
			},
			"fortios_firewall_addresslist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_addresslist"),
			},
			"fortios_firewall_addrgrp": {
				Tok: makeDataSource(mainMod, "fortios_firewall_addrgrp"),
			},
			"fortios_firewall_addrgrp6": {
				Tok: makeDataSource(mainMod, "fortios_firewall_addrgrp6"),
			},
			"fortios_firewall_addrgrp6list": {
				Tok: makeDataSource(mainMod, "fortios_firewall_addrgrp6list"),
			},
			"fortios_firewall_addrgrplist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_addrgrplist"),
			},
			"fortios_firewall_centralsnatmap": {
				Tok: makeDataSource(mainMod, "fortios_firewall_centralsnatmap"),
			},
			"fortios_firewall_centralsnatmaplist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_centralsnatmaplist"),
			},
			"fortios_firewall_internetservice": {
				Tok: makeDataSource(mainMod, "fortios_firewall_internetservice"),
			},
			"fortios_firewall_internetservicecustom": {
				Tok: makeDataSource(mainMod, "fortios_firewall_internetservicecustom"),
			},
			"fortios_firewall_internetservicecustomgroup": {
				Tok: makeDataSource(mainMod, "fortios_firewall_internetservicecustomgroup"),
			},
			"fortios_firewall_internetservicecustomgrouplist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_internetservicecustomgrouplist"),
			},
			"fortios_firewall_internetservicecustomlist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_internetservicecustomlist"),
			},
			"fortios_firewall_internetservicedefinition": {
				Tok: makeDataSource(mainMod, "fortios_firewall_internetservicedefinition"),
			},
			"fortios_firewall_internetservicedefinitionlist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_internetservicedefinitionlist"),
			},
			"fortios_firewall_internetserviceextension": {
				Tok: makeDataSource(mainMod, "fortios_firewall_internetserviceextension"),
			},
			"fortios_firewall_internetserviceextensionlist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_internetserviceextensionlist"),
			},
			"fortios_firewall_internetservicegroup": {
				Tok: makeDataSource(mainMod, "fortios_firewall_internetservicegroup"),
			},
			"fortios_firewall_internetservicegrouplist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_internetservicegrouplist"),
			},
			"fortios_firewall_internetservicelist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_internetservicelist"),
			},
			"fortios_firewall_ipv6ehfilter": {
				Tok: makeDataSource(mainMod, "fortios_firewall_ipv6ehfilter"),
			},
			"fortios_firewall_multicastaddress": {
				Tok: makeDataSource(mainMod, "fortios_firewall_multicastaddress"),
			},
			"fortios_firewall_multicastaddress6": {
				Tok: makeDataSource(mainMod, "fortios_firewall_multicastaddress6"),
			},
			"fortios_firewall_multicastaddress6list": {
				Tok: makeDataSource(mainMod, "fortios_firewall_multicastaddress6list"),
			},
			"fortios_firewall_multicastaddresslist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_multicastaddresslist"),
			},
			"fortios_firewall_policy": {
				Tok: makeDataSource(mainMod, "fortios_firewall_policy"),
			},
			"fortios_firewall_policy46": {
				Tok: makeDataSource(mainMod, "fortios_firewall_policy46"),
			},
			"fortios_firewall_policy46list": {
				Tok: makeDataSource(mainMod, "fortios_firewall_policy46list"),
			},
			"fortios_firewall_policy6": {
				Tok: makeDataSource(mainMod, "fortios_firewall_policy6"),
			},
			"fortios_firewall_policy64": {
				Tok: makeDataSource(mainMod, "fortios_firewall_policy64"),
			},
			"fortios_firewall_policy64list": {
				Tok: makeDataSource(mainMod, "fortios_firewall_policy64list"),
			},
			"fortios_firewall_policy6list": {
				Tok: makeDataSource(mainMod, "fortios_firewall_policy6list"),
			},
			"fortios_firewall_policylist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_policylist"),
			},
			"fortios_firewall_profileprotocoloptions": {
				Tok: makeDataSource(mainMod, "fortios_firewall_profileprotocoloptions"),
			},
			"fortios_firewall_profileprotocoloptionslist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_profileprotocoloptionslist"),
			},
			"fortios_firewall_proxyaddress": {
				Tok: makeDataSource(mainMod, "fortios_firewall_proxyaddress"),
			},
			"fortios_firewall_proxyaddresslist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_proxyaddresslist"),
			},
			"fortios_firewall_proxyaddrgrp": {
				Tok: makeDataSource(mainMod, "fortios_firewall_proxyaddrgrp"),
			},
			"fortios_firewall_proxyaddrgrplist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_proxyaddrgrplist"),
			},
			"fortios_firewall_proxypolicy": {
				Tok: makeDataSource(mainMod, "fortios_firewall_proxypolicy"),
			},
			"fortios_firewall_proxypolicylist": {
				Tok: makeDataSource(mainMod, "fortios_firewall_proxypolicylist"),
			},
			"fortios_firewallconsolidated_policy": {
				Tok: makeDataSource(mainMod, "fortios_firewallconsolidated_policy"),
			},
			"fortios_firewallconsolidated_policylist": {
				Tok: makeDataSource(mainMod, "fortios_firewallconsolidated_policylist"),
			},
			"fortios_firewallschedule_group": {
				Tok: makeDataSource(mainMod, "fortios_firewallschedule_group"),
			},
			"fortios_firewallschedule_grouplist": {
				Tok: makeDataSource(mainMod, "fortios_firewallschedule_grouplist"),
			},
			"fortios_firewallschedule_onetime": {
				Tok: makeDataSource(mainMod, "fortios_firewallschedule_onetime"),
			},
			"fortios_firewallschedule_onetimelist": {
				Tok: makeDataSource(mainMod, "fortios_firewallschedule_onetimelist"),
			},
			"fortios_firewallschedule_recurring": {
				Tok: makeDataSource(mainMod, "fortios_firewallschedule_recurring"),
			},
			"fortios_firewallschedule_recurringlist": {
				Tok: makeDataSource(mainMod, "fortios_firewallschedule_recurringlist"),
			},
			"fortios_firewallservice_category": {
				Tok: makeDataSource(mainMod, "fortios_firewallservice_category"),
			},
			"fortios_firewallservice_categorylist": {
				Tok: makeDataSource(mainMod, "fortios_firewallservice_categorylist"),
			},
			"fortios_firewallservice_custom": {
				Tok: makeDataSource(mainMod, "fortios_firewallservice_custom"),
			},
			"fortios_firewallservice_customlist": {
				Tok: makeDataSource(mainMod, "fortios_firewallservice_customlist"),
			},
			"fortios_firewallservice_group": {
				Tok: makeDataSource(mainMod, "fortios_firewallservice_group"),
			},
			"fortios_firewallservice_grouplist": {
				Tok: makeDataSource(mainMod, "fortios_firewallservice_grouplist"),
			},
			"fortios_firewallshaper_peripshaper": {
				Tok: makeDataSource(mainMod, "fortios_firewallshaper_peripshaper"),
			},
			"fortios_firewallshaper_peripshaperlist": {
				Tok: makeDataSource(mainMod, "fortios_firewallshaper_peripshaperlist"),
			},
			"fortios_firewallshaper_trafficshaper": {
				Tok: makeDataSource(mainMod, "fortios_firewallshaper_trafficshaper"),
			},
			"fortios_firewallshaper_trafficshaperlist": {
				Tok: makeDataSource(mainMod, "fortios_firewallshaper_trafficshaperlist"),
			},
			"fortios_firewallwildcardfqdn_custom": {
				Tok: makeDataSource(mainMod, "fortios_firewallwildcardfqdn_custom"),
			},
			"fortios_firewallwildcardfqdn_customlist": {
				Tok: makeDataSource(mainMod, "fortios_firewallwildcardfqdn_customlist"),
			},
			"fortios_firewallwildcardfqdn_group": {
				Tok: makeDataSource(mainMod, "fortios_firewallwildcardfqdn_group"),
			},
			"fortios_firewallwildcardfqdn_grouplist": {
				Tok: makeDataSource(mainMod, "fortios_firewallwildcardfqdn_grouplist"),
			},
			"fortios_ipmask_cidr": {
				Tok: makeDataSource(mainMod, "fortios_ipmask_cidr"),
			},
			"fortios_json_generic_api": {
				Tok: makeDataSource(mainMod, "fortios_json_generic_api"),
			},
			"fortios_router_accesslist": {
				Tok: makeDataSource(mainMod, "fortios_router_accesslist"),
			},
			"fortios_router_accesslist6": {
				Tok: makeDataSource(mainMod, "fortios_router_accesslist6"),
			},
			"fortios_router_accesslist6list": {
				Tok: makeDataSource(mainMod, "fortios_router_accesslist6list"),
			},
			"fortios_router_accesslistlist": {
				Tok: makeDataSource(mainMod, "fortios_router_accesslistlist"),
			},
			"fortios_router_aspathlist": {
				Tok: makeDataSource(mainMod, "fortios_router_aspathlist"),
			},
			"fortios_router_aspathlistlist": {
				Tok: makeDataSource(mainMod, "fortios_router_aspathlistlist"),
			},
			"fortios_router_authpath": {
				Tok: makeDataSource(mainMod, "fortios_router_authpath"),
			},
			"fortios_router_authpathlist": {
				Tok: makeDataSource(mainMod, "fortios_router_authpathlist"),
			},
			"fortios_router_bfd": {
				Tok: makeDataSource(mainMod, "fortios_router_bfd"),
			},
			"fortios_router_bfd6": {
				Tok: makeDataSource(mainMod, "fortios_router_bfd6"),
			},
			"fortios_router_bgp": {
				Tok: makeDataSource(mainMod, "fortios_router_bgp"),
			},
			"fortios_router_communitylist": {
				Tok: makeDataSource(mainMod, "fortios_router_communitylist"),
			},
			"fortios_router_communitylistlist": {
				Tok: makeDataSource(mainMod, "fortios_router_communitylistlist"),
			},
			"fortios_router_isis": {
				Tok: makeDataSource(mainMod, "fortios_router_isis"),
			},
			"fortios_router_keychain": {
				Tok: makeDataSource(mainMod, "fortios_router_keychain"),
			},
			"fortios_router_keychainlist": {
				Tok: makeDataSource(mainMod, "fortios_router_keychainlist"),
			},
			"fortios_router_multicast": {
				Tok: makeDataSource(mainMod, "fortios_router_multicast"),
			},
			"fortios_router_multicast6": {
				Tok: makeDataSource(mainMod, "fortios_router_multicast6"),
			},
			"fortios_router_multicastflow": {
				Tok: makeDataSource(mainMod, "fortios_router_multicastflow"),
			},
			"fortios_router_multicastflowlist": {
				Tok: makeDataSource(mainMod, "fortios_router_multicastflowlist"),
			},
			"fortios_router_ospf": {
				Tok: makeDataSource(mainMod, "fortios_router_ospf"),
			},
			"fortios_router_ospf6": {
				Tok: makeDataSource(mainMod, "fortios_router_ospf6"),
			},
			"fortios_router_policy": {
				Tok: makeDataSource(mainMod, "fortios_router_policy"),
			},
			"fortios_router_policy6": {
				Tok: makeDataSource(mainMod, "fortios_router_policy6"),
			},
			"fortios_router_policy6list": {
				Tok: makeDataSource(mainMod, "fortios_router_policy6list"),
			},
			"fortios_router_policylist": {
				Tok: makeDataSource(mainMod, "fortios_router_policylist"),
			},
			"fortios_router_prefixlist": {
				Tok: makeDataSource(mainMod, "fortios_router_prefixlist"),
			},
			"fortios_router_prefixlist6": {
				Tok: makeDataSource(mainMod, "fortios_router_prefixlist6"),
			},
			"fortios_router_prefixlist6list": {
				Tok: makeDataSource(mainMod, "fortios_router_prefixlist6list"),
			},
			"fortios_router_prefixlistlist": {
				Tok: makeDataSource(mainMod, "fortios_router_prefixlistlist"),
			},
			"fortios_router_rip": {
				Tok: makeDataSource(mainMod, "fortios_router_rip"),
			},
			"fortios_router_ripng": {
				Tok: makeDataSource(mainMod, "fortios_router_ripng"),
			},
			"fortios_router_routemap": {
				Tok: makeDataSource(mainMod, "fortios_router_routemap"),
			},
			"fortios_router_routemaplist": {
				Tok: makeDataSource(mainMod, "fortios_router_routemaplist"),
			},
			"fortios_router_setting": {
				Tok: makeDataSource(mainMod, "fortios_router_setting"),
			},
			"fortios_router_static": {
				Tok: makeDataSource(mainMod, "fortios_router_static"),
			},
			"fortios_router_static6": {
				Tok: makeDataSource(mainMod, "fortios_router_static6"),
			},
			"fortios_router_static6list": {
				Tok: makeDataSource(mainMod, "fortios_router_static6list"),
			},
			"fortios_router_staticlist": {
				Tok: makeDataSource(mainMod, "fortios_router_staticlist"),
			},
			"fortios_routerbgp_neighbor": {
				Tok: makeDataSource(mainMod, "fortios_routerbgp_neighbor"),
			},
			"fortios_routerbgp_neighborlist": {
				Tok: makeDataSource(mainMod, "fortios_routerbgp_neighborlist"),
			},
			"fortios_system_accprofile": {
				Tok: makeDataSource(mainMod, "fortios_system_accprofile"),
			},
			"fortios_system_accprofilelist": {
				Tok: makeDataSource(mainMod, "fortios_system_accprofilelist"),
			},
			"fortios_system_admin": {
				Tok: makeDataSource(mainMod, "fortios_system_admin"),
			},
			"fortios_system_adminlist": {
				Tok: makeDataSource(mainMod, "fortios_system_adminlist"),
			},
			"fortios_system_alias": {
				Tok: makeDataSource(mainMod, "fortios_system_alias"),
			},
			"fortios_system_aliaslist": {
				Tok: makeDataSource(mainMod, "fortios_system_aliaslist"),
			},
			"fortios_system_apiuser": {
				Tok: makeDataSource(mainMod, "fortios_system_apiuser"),
			},
			"fortios_system_apiuserlist": {
				Tok: makeDataSource(mainMod, "fortios_system_apiuserlist"),
			},
			"fortios_system_arptable": {
				Tok: makeDataSource(mainMod, "fortios_system_arptable"),
			},
			"fortios_system_arptablelist": {
				Tok: makeDataSource(mainMod, "fortios_system_arptablelist"),
			},
			"fortios_system_autoinstall": {
				Tok: makeDataSource(mainMod, "fortios_system_autoinstall"),
			},
			"fortios_system_automationaction": {
				Tok: makeDataSource(mainMod, "fortios_system_automationaction"),
			},
			"fortios_system_automationactionlist": {
				Tok: makeDataSource(mainMod, "fortios_system_automationactionlist"),
			},
			"fortios_system_automationdestination": {
				Tok: makeDataSource(mainMod, "fortios_system_automationdestination"),
			},
			"fortios_system_automationdestinationlist": {
				Tok: makeDataSource(mainMod, "fortios_system_automationdestinationlist"),
			},
			"fortios_system_automationtrigger": {
				Tok: makeDataSource(mainMod, "fortios_system_automationtrigger"),
			},
			"fortios_system_automationtriggerlist": {
				Tok: makeDataSource(mainMod, "fortios_system_automationtriggerlist"),
			},
			"fortios_system_autoscript": {
				Tok: makeDataSource(mainMod, "fortios_system_autoscript"),
			},
			"fortios_system_autoscriptlist": {
				Tok: makeDataSource(mainMod, "fortios_system_autoscriptlist"),
			},
			"fortios_system_centralmanagement": {
				Tok: makeDataSource(mainMod, "fortios_system_centralmanagement"),
			},
			"fortios_system_clustersync": {
				Tok: makeDataSource(mainMod, "fortios_system_clustersync"),
			},
			"fortios_system_clustersynclist": {
				Tok: makeDataSource(mainMod, "fortios_system_clustersynclist"),
			},
			"fortios_system_console": {
				Tok: makeDataSource(mainMod, "fortios_system_console"),
			},
			"fortios_system_csf": {
				Tok: makeDataSource(mainMod, "fortios_system_csf"),
			},
			"fortios_system_ddns": {
				Tok: makeDataSource(mainMod, "fortios_system_ddns"),
			},
			"fortios_system_ddnslist": {
				Tok: makeDataSource(mainMod, "fortios_system_ddnslist"),
			},
			"fortios_system_dns": {
				Tok: makeDataSource(mainMod, "fortios_system_dns"),
			},
			"fortios_system_dnsdatabase": {
				Tok: makeDataSource(mainMod, "fortios_system_dnsdatabase"),
			},
			"fortios_system_dnsdatabaselist": {
				Tok: makeDataSource(mainMod, "fortios_system_dnsdatabaselist"),
			},
			"fortios_system_dnsserver": {
				Tok: makeDataSource(mainMod, "fortios_system_dnsserver"),
			},
			"fortios_system_dnsserverlist": {
				Tok: makeDataSource(mainMod, "fortios_system_dnsserverlist"),
			},
			"fortios_system_dscpbasedpriority": {
				Tok: makeDataSource(mainMod, "fortios_system_dscpbasedpriority"),
			},
			"fortios_system_dscpbasedprioritylist": {
				Tok: makeDataSource(mainMod, "fortios_system_dscpbasedprioritylist"),
			},
			"fortios_system_emailserver": {
				Tok: makeDataSource(mainMod, "fortios_system_emailserver"),
			},
			"fortios_system_externalresource": {
				Tok: makeDataSource(mainMod, "fortios_system_externalresource"),
			},
			"fortios_system_externalresourcelist": {
				Tok: makeDataSource(mainMod, "fortios_system_externalresourcelist"),
			},
			"fortios_system_fipscc": {
				Tok: makeDataSource(mainMod, "fortios_system_fipscc"),
			},
			"fortios_system_fm": {
				Tok: makeDataSource(mainMod, "fortios_system_fm"),
			},
			"fortios_system_fortiguard": {
				Tok: makeDataSource(mainMod, "fortios_system_fortiguard"),
			},
			"fortios_system_fortimanager": {
				Tok: makeDataSource(mainMod, "fortios_system_fortimanager"),
			},
			"fortios_system_fortisandbox": {
				Tok: makeDataSource(mainMod, "fortios_system_fortisandbox"),
			},
			"fortios_system_fssopolling": {
				Tok: makeDataSource(mainMod, "fortios_system_fssopolling"),
			},
			"fortios_system_ftmpush": {
				Tok: makeDataSource(mainMod, "fortios_system_ftmpush"),
			},
			"fortios_system_global": {
				Tok: makeDataSource(mainMod, "fortios_system_global"),
			},
			"fortios_system_gretunnel": {
				Tok: makeDataSource(mainMod, "fortios_system_gretunnel"),
			},
			"fortios_system_gretunnellist": {
				Tok: makeDataSource(mainMod, "fortios_system_gretunnellist"),
			},
			"fortios_system_ha": {
				Tok: makeDataSource(mainMod, "fortios_system_ha"),
			},
			"fortios_system_hamonitor": {
				Tok: makeDataSource(mainMod, "fortios_system_hamonitor"),
			},
			"fortios_system_interface": {
				Tok: makeDataSource(mainMod, "fortios_system_interface"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"interface": {
						CSharpName: "Data",
					},
				},
			},
			"fortios_system_interfacelist": {
				Tok: makeDataSource(mainMod, "fortios_system_interfacelist"),
			},
			"fortios_system_ipiptunnel": {
				Tok: makeDataSource(mainMod, "fortios_system_ipiptunnel"),
			},
			"fortios_system_ipiptunnellist": {
				Tok: makeDataSource(mainMod, "fortios_system_ipiptunnellist"),
			},
			"fortios_system_ipv6neighborcache": {
				Tok: makeDataSource(mainMod, "fortios_system_ipv6neighborcache"),
			},
			"fortios_system_ipv6neighborcachelist": {
				Tok: makeDataSource(mainMod, "fortios_system_ipv6neighborcachelist"),
			},
			"fortios_system_ipv6tunnel": {
				Tok: makeDataSource(mainMod, "fortios_system_ipv6tunnel"),
			},
			"fortios_system_ipv6tunnellist": {
				Tok: makeDataSource(mainMod, "fortios_system_ipv6tunnellist"),
			},
			"fortios_system_linkmonitor": {
				Tok: makeDataSource(mainMod, "fortios_system_linkmonitor"),
			},
			"fortios_system_linkmonitorlist": {
				Tok: makeDataSource(mainMod, "fortios_system_linkmonitorlist"),
			},
			"fortios_system_managementtunnel": {
				Tok: makeDataSource(mainMod, "fortios_system_managementtunnel"),
			},
			"fortios_system_mobiletunnel": {
				Tok: makeDataSource(mainMod, "fortios_system_mobiletunnel"),
			},
			"fortios_system_mobiletunnellist": {
				Tok: makeDataSource(mainMod, "fortios_system_mobiletunnellist"),
			},
			"fortios_system_nat64": {
				Tok: makeDataSource(mainMod, "fortios_system_nat64"),
			},
			"fortios_system_ndproxy": {
				Tok: makeDataSource(mainMod, "fortios_system_ndproxy"),
			},
			"fortios_system_netflow": {
				Tok: makeDataSource(mainMod, "fortios_system_netflow"),
			},
			"fortios_system_networkvisibility": {
				Tok: makeDataSource(mainMod, "fortios_system_networkvisibility"),
			},
			"fortios_system_ntp": {
				Tok: makeDataSource(mainMod, "fortios_system_ntp"),
			},
			"fortios_system_objecttagging": {
				Tok: makeDataSource(mainMod, "fortios_system_objecttagging"),
			},
			"fortios_system_objecttagginglist": {
				Tok: makeDataSource(mainMod, "fortios_system_objecttagginglist"),
			},
			"fortios_system_passwordpolicy": {
				Tok: makeDataSource(mainMod, "fortios_system_passwordpolicy"),
			},
			"fortios_system_passwordpolicyguestadmin": {
				Tok: makeDataSource(mainMod, "fortios_system_passwordpolicyguestadmin"),
			},
			"fortios_system_pppoeinterface": {
				Tok: makeDataSource(mainMod, "fortios_system_pppoeinterface"),
			},
			"fortios_system_pppoeinterfacelist": {
				Tok: makeDataSource(mainMod, "fortios_system_pppoeinterfacelist"),
			},
			"fortios_system_proberesponse": {
				Tok: makeDataSource(mainMod, "fortios_system_proberesponse"),
			},
			"fortios_system_proxyarp": {
				Tok: makeDataSource(mainMod, "fortios_system_proxyarp"),
			},
			"fortios_system_proxyarplist": {
				Tok: makeDataSource(mainMod, "fortios_system_proxyarplist"),
			},
			"fortios_system_replacemsggroup": {
				Tok: makeDataSource(mainMod, "fortios_system_replacemsggroup"),
			},
			"fortios_system_replacemsggrouplist": {
				Tok: makeDataSource(mainMod, "fortios_system_replacemsggrouplist"),
			},
			"fortios_system_replacemsgimage": {
				Tok: makeDataSource(mainMod, "fortios_system_replacemsgimage"),
			},
			"fortios_system_replacemsgimagelist": {
				Tok: makeDataSource(mainMod, "fortios_system_replacemsgimagelist"),
			},
			"fortios_system_resourcelimits": {
				Tok: makeDataSource(mainMod, "fortios_system_resourcelimits"),
			},
			"fortios_system_sdnconnector": {
				Tok: makeDataSource(mainMod, "fortios_system_sdnconnector"),
			},
			"fortios_system_sdnconnectorlist": {
				Tok: makeDataSource(mainMod, "fortios_system_sdnconnectorlist"),
			},
			"fortios_system_sessionhelper": {
				Tok: makeDataSource(mainMod, "fortios_system_sessionhelper"),
			},
			"fortios_system_sessionhelperlist": {
				Tok: makeDataSource(mainMod, "fortios_system_sessionhelperlist"),
			},
			"fortios_system_sessionttl": {
				Tok: makeDataSource(mainMod, "fortios_system_sessionttl"),
			},
			"fortios_system_sflow": {
				Tok: makeDataSource(mainMod, "fortios_system_sflow"),
			},
			"fortios_system_sittunnel": {
				Tok: makeDataSource(mainMod, "fortios_system_sittunnel"),
			},
			"fortios_system_sittunnellist": {
				Tok: makeDataSource(mainMod, "fortios_system_sittunnellist"),
			},
			"fortios_system_smsserver": {
				Tok: makeDataSource(mainMod, "fortios_system_smsserver"),
			},
			"fortios_system_smsserverlist": {
				Tok: makeDataSource(mainMod, "fortios_system_smsserverlist"),
			},
			"fortios_system_tosbasedpriority": {
				Tok: makeDataSource(mainMod, "fortios_system_tosbasedpriority"),
			},
			"fortios_system_tosbasedprioritylist": {
				Tok: makeDataSource(mainMod, "fortios_system_tosbasedprioritylist"),
			},
			"fortios_system_vdomexception": {
				Tok: makeDataSource(mainMod, "fortios_system_vdomexception"),
			},
			"fortios_system_vdomexceptionlist": {
				Tok: makeDataSource(mainMod, "fortios_system_vdomexceptionlist"),
			},
			"fortios_system_vdomnetflow": {
				Tok: makeDataSource(mainMod, "fortios_system_vdomnetflow"),
			},
			"fortios_system_vdomsflow": {
				Tok: makeDataSource(mainMod, "fortios_system_vdomsflow"),
			},
			"fortios_system_virtualwanlink": {
				Tok: makeDataSource(mainMod, "fortios_system_virtualwanlink"),
			},
			"fortios_system_vxlan": {
				Tok: makeDataSource(mainMod, "fortios_system_vxlan"),
			},
			"fortios_system_vxlanlist": {
				Tok: makeDataSource(mainMod, "fortios_system_vxlanlist"),
			},
			"fortios_system_wccp": {
				Tok: makeDataSource(mainMod, "fortios_system_wccp"),
			},
			"fortios_system_wccplist": {
				Tok: makeDataSource(mainMod, "fortios_system_wccplist"),
			},
			"fortios_system_zone": {
				Tok: makeDataSource(mainMod, "fortios_system_zone"),
			},
			"fortios_system_zonelist": {
				Tok: makeDataSource(mainMod, "fortios_system_zonelist"),
			},
			"fortios_systemautoupdate_pushupdate": {
				Tok: makeDataSource(mainMod, "fortios_systemautoupdate_pushupdate"),
			},
			"fortios_systemautoupdate_schedule": {
				Tok: makeDataSource(mainMod, "fortios_systemautoupdate_schedule"),
			},
			"fortios_systemautoupdate_tunneling": {
				Tok: makeDataSource(mainMod, "fortios_systemautoupdate_tunneling"),
			},
			"fortios_systemdhcp_server": {
				Tok: makeDataSource(mainMod, "fortios_systemdhcp_server"),
			},
			"fortios_systemdhcp_serverlist": {
				Tok: makeDataSource(mainMod, "fortios_systemdhcp_serverlist"),
			},
			"fortios_systemlldp_networkpolicy": {
				Tok: makeDataSource(mainMod, "fortios_systemlldp_networkpolicy"),
			},
			"fortios_systemlldp_networkpolicylist": {
				Tok: makeDataSource(mainMod, "fortios_systemlldp_networkpolicylist"),
			},
			"fortios_systemsnmp_community": {
				Tok: makeDataSource(mainMod, "fortios_systemsnmp_community"),
			},
			"fortios_systemsnmp_communitylist": {
				Tok: makeDataSource(mainMod, "fortios_systemsnmp_communitylist"),
			},
			"fortios_systemsnmp_sysinfo": {
				Tok: makeDataSource(mainMod, "fortios_systemsnmp_sysinfo"),
			},
			"fortios_systemsnmp_user": {
				Tok: makeDataSource(mainMod, "fortios_systemsnmp_user"),
			},
			"fortios_systemsnmp_userlist": {
				Tok: makeDataSource(mainMod, "fortios_systemsnmp_userlist"),
			},
			"fortios_user_saml": {
				Tok: makeDataSource(mainMod, "fortios_user_saml"),
			},
			"fortios_user_samllist": {
				Tok: makeDataSource(mainMod, "fortios_user_samllist"),
			},
		},
		IgnoreMappings: []string{
			"fortios_vpn_ipsec_phase1interface", // deprecated
			"fortios_vpn_ipsec_phase2interface", // deprecated
			"fortios_log_fortianalyzer_setting", // deprecated
			"fortios_firewall_security_policy",  // deprecated
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/fortios",

			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumiverse_fortios",

			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", "fortios"),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				"fortios",
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumiverse",

			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.pulumiverse",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
