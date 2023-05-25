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

	"github.com/aspyrmedia/pulumi-fortios/provider/pkg/version"
	"github.com/aspyrmedia/terraform-provider-fortios/fortios"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "fortios"
	// modules:
	mainMod                           = "index" // the fortios module
	antivirusMod                      = "antivirus"
	applicationMod                    = "application"
	authenticationMod                 = "authentication"
	automationMod                     = "automation"
	certificateMod                    = "certificate"
	cifsMod                           = "cifs"
	credentialstoreMod                = "credentialstore"
	dlpMod                            = "dlp"
	dnsfilterMod                      = "dnsfilter"
	dpdkMod                           = "dpdk"
	emailfilterMod                    = "emailfilter"
	endpointcontrolMod                = "endpointcontrol"
	extendercontrollerMod             = "extendercontroller"
	extensioncontrollerMod            = "extensioncontroller"
	filefilterMod                     = "filefilter"
	firewallMod                       = "firewall"
	firewallconsolidatedMod           = "firewallconsolidated"
	firewallipmacbindingMod           = "firewallipmacbinding"
	firewallscheduleMod               = "firewallschedule"
	firewallserviceMod                = "firewallservice"
	firewallshaperMod                 = "firewallshaper"
	firewallsshMod                    = "firewallssh"
	firewallsslMod                    = "firewallssl"
	firewallwildcardfqdnMod           = "firewallwildcardfqdn"
	fmgMod                            = "fmg"
	ftpproxyMod                       = "ftpproxy"
	icapMod                           = "icap"
	ipsMod                            = "ips"
	jsonMod                           = "json"
	logMod                            = "log"
	logdiskMod                        = "logdisk"
	logfortianalyzer2Mod              = "logfortianalyzer2"
	logfortianalyzer3Mod              = "logfortianalyzer3"
	logfortianalyzerMod               = "logfortianalyzer"
	logfortianalyzercloudMod          = "logfortianalyzercloud"
	logfortiguardMod                  = "logfortiguard"
	logmemoryMod                      = "logmemory"
	lognulldeviceMod                  = "lognulldevice"
	logsyslogd2Mod                    = "logsyslogd2"
	logsyslogd3Mod                    = "logsyslogd3"
	logsyslogd4Mod                    = "logsyslogd4"
	logsyslogdMod                     = "logsyslogd"
	logtacacsaccounting2Mod           = "logtacacsaccounting2"
	logtacacsaccounting3Mod           = "logtacacsaccounting3"
	logtacacsaccountingMod            = "logtacacsaccounting"
	logwebtrendsMod                   = "logwebtrends"
	networkingMod                     = "networking"
	nsxtMod                           = "nsxt"
	reportMod                         = "report"
	routerMod                         = "router"
	routerbgpMod                      = "routerbgp"
	routerospf6Mod                    = "routerospf6"
	routerospfMod                     = "routerospf"
	sctpfilterMod                     = "sctpfilter"
	spamfilterMod                     = "spamfilter"
	sshfilterMod                      = "sshfilter"
	switchcontrollerMod               = "switchcontroller"
	switchcontrollerautoconfigMod     = "switchcontrollerautoconfig"
	switchcontrollerinitialconfigMod  = "switchcontrollerinitialconfig"
	switchcontrollerptpMod            = "switchcontrollerptp"
	switchcontrollerqosMod            = "switchcontrollerqos"
	switchcontrollersecuritypolicyMod = "switchcontrollersecuritypolicy"
	system3gmodemMod                  = "system3gmodem"
	systemMod                         = "system"
	systemautoupdateMod               = "systemautoupdate"
	systemdhcp6Mod                    = "systemdhcp6"
	systemdhcpMod                     = "systemdhcp"
	systemlldpMod                     = "systemlldp"
	systemreplacemsgMod               = "systemreplacemsg"
	systemsnmpMod                     = "systemsnmp"
	userMod                           = "user"
	videofilterMod                    = "videofilter"
	voipMod                           = "voip"
	vpnMod                            = "vpn"
	vpncertificateMod                 = "vpncertificate"
	vpnipsecMod                       = "vpnipsec"
	vpnsslMod                         = "vpnssl"
	vpnsslwebMod                      = "vpnsslweb"
	wafMod                            = "waf"
	wanoptMod                         = "wanopt"
	webfilterMod                      = "webfilter"
	webproxyMod                       = "webproxy"
	wirelesscontrollerMod             = "wirelesscontroller"
	wirelesscontrollerhotspot20Mod    = "wirelesscontrollerhotspot20"
	ipmaskMod                         = "ipmask"
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
// func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
// 	return nil
// }

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
		DisplayName: "Fortinet FortiOS",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Aspyrmedia",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "https://github.com/aspyrmedia/pulumi-fortios/releases/download/v${VERSION}",
		Description:       "A Pulumi package for creating and managing fortios cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "fortios", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/aspyrmedia/pulumi-fortios",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "fortinetdev",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		// PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
			"fortios_alertemail_setting": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "alertEmailSetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# FortiOS Alert Email Setting Resource"),
				},
			},
			"fortios_antivirus_heuristic": {
				Tok: tfbridge.MakeResource(mainPkg, antivirusMod, "antivirusHeuristic"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# antivirus - heuristic"),
				},
			},
			"fortios_antivirus_profile": {
				Tok: tfbridge.MakeResource(mainPkg, antivirusMod, "antivirusProfile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# antivirus - profile"),
				},
			},
			"fortios_antivirus_quarantine": {
				Tok: tfbridge.MakeResource(mainPkg, antivirusMod, "antivirusQuarantine"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# antivirus - quarantine"),
				},
			},
			"fortios_antivirus_settings": {
				Tok: tfbridge.MakeResource(mainPkg, antivirusMod, "antivirusSettings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# antivirus - settings"),
				},
			},
			"fortios_application_custom": {
				Tok: tfbridge.MakeResource(mainPkg, applicationMod, "custom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# application - custom"),
				},
			},
			"fortios_application_group": {
				Tok: tfbridge.MakeResource(mainPkg, applicationMod, "group"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# application - group"),
				},
			},
			"fortios_application_list": {
				Tok: tfbridge.MakeResource(mainPkg, applicationMod, "list"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# application - list"),
				},
			},
			"fortios_application_name": {
				Tok: tfbridge.MakeResource(mainPkg, applicationMod, "name"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# application - name"),
				},
			},
			"fortios_application_rulesettings": {
				Tok: tfbridge.MakeResource(mainPkg, applicationMod, "rulesettings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# application - rulesettings"),
				},
			},
			"fortios_authentication_rule": {
				Tok: tfbridge.MakeResource(mainPkg, authenticationMod, "rule"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# authentication - rule"),
				},
			},
			"fortios_authentication_scheme": {
				Tok: tfbridge.MakeResource(mainPkg, authenticationMod, "scheme"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# authentication - scheme"),
				},
			},
			"fortios_authentication_setting": {
				Tok: tfbridge.MakeResource(mainPkg, authenticationMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# authentication - setting"),
				},
			},
			"fortios_automation_setting": {
				Tok: tfbridge.MakeResource(mainPkg, automationMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# automation - setting"),
				},
			},
			"fortios_certificate_ca": {
				Tok: tfbridge.MakeResource(mainPkg, certificateMod, "ca"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# certificate - ca"),
				},
			},
			"fortios_certificate_crl": {
				Tok: tfbridge.MakeResource(mainPkg, certificateMod, "crl"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# certificate - crl"),
				},
			},
			"fortios_certificate_local": {
				Tok: tfbridge.MakeResource(mainPkg, certificateMod, "local"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# certificate - local"),
				},
			},
			"fortios_certificate_remote": {
				Tok: tfbridge.MakeResource(mainPkg, certificateMod, "remote"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# certificate - remote"),
				},
			},
			"fortios_cifs_domaincontroller": {
				Tok: tfbridge.MakeResource(mainPkg, cifsMod, "domaincontroller"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# cifs - domaincontroller"),
				},
			},
			"fortios_cifs_profile": {
				Tok: tfbridge.MakeResource(mainPkg, cifsMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# cifs - profile"),
				},
			},
			"fortios_credentialstore_domaincontroller": {
				Tok: tfbridge.MakeResource(mainPkg, credentialstoreMod, "domaincontroller"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# credentialstore - domaincontroller"),
				},
			},
			"fortios_dlp_datatype": {
				Tok: tfbridge.MakeResource(mainPkg, dlpMod, "datatype"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dlp - datatype"),
				},
			},
			"fortios_dlp_dictionary": {
				Tok: tfbridge.MakeResource(mainPkg, dlpMod, "dictionary"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dlp - dictionary"),
				},
			},
			"fortios_dlp_filepattern": {
				Tok: tfbridge.MakeResource(mainPkg, dlpMod, "filepattern"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dlp - filepattern"),
				},
			},
			"fortios_dlp_fpdocsource": {
				Tok: tfbridge.MakeResource(mainPkg, dlpMod, "fpdocsource"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dlp - fpdocsource"),
				},
			},
			"fortios_dlp_fpsensitivity": {
				Tok: tfbridge.MakeResource(mainPkg, dlpMod, "fpsensitivity"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dlp - fpsensitivity"),
				},
			},
			"fortios_dlp_profile": {
				Tok: tfbridge.MakeResource(mainPkg, dlpMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dlp - profile"),
				},
			},
			"fortios_dlp_sensitivity": {
				Tok: tfbridge.MakeResource(mainPkg, dlpMod, "sensitivity"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dlp - sensitivity"),
				},
			},
			"fortios_dlp_sensor": {
				Tok: tfbridge.MakeResource(mainPkg, dlpMod, "sensor"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dlp - sensor"),
				},
			},
			"fortios_dlp_settings": {
				Tok: tfbridge.MakeResource(mainPkg, dlpMod, "settings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dlp - settings"),
				},
			},
			"fortios_dnsfilter_domainfilter": {
				Tok: tfbridge.MakeResource(mainPkg, dnsfilterMod, "domainfilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dnsfilter - domainfilter"),
				},
			},
			"fortios_dnsfilter_profile": {
				Tok: tfbridge.MakeResource(mainPkg, dnsfilterMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dnsfilter - profile"),
				},
			},
			"fortios_dpdk_cpus": {
				Tok: tfbridge.MakeResource(mainPkg, dpdkMod, "cpus"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dpdk - cpus"),
				},
			},
			"fortios_dpdk_global": {
				Tok: tfbridge.MakeResource(mainPkg, dpdkMod, "global"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# dpdk - global"),
				},
			},
			"fortios_emailfilter_blockallowlist": {
				Tok: tfbridge.MakeResource(mainPkg, emailfilterMod, "blockallowlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# emailfilter - blockallowlist"),
				},
			},
			"fortios_emailfilter_bwl": {
				Tok: tfbridge.MakeResource(mainPkg, emailfilterMod, "bwl"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# emailfilter - bwl"),
				},
			},
			"fortios_emailfilter_bword": {
				Tok: tfbridge.MakeResource(mainPkg, emailfilterMod, "bword"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# emailfilter - bword"),
				},
			},
			"fortios_emailfilter_dnsbl": {
				Tok: tfbridge.MakeResource(mainPkg, emailfilterMod, "dnsbl"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# emailfilter - dnsbl"),
				},
			},
			"fortios_emailfilter_fortishield": {
				Tok: tfbridge.MakeResource(mainPkg, emailfilterMod, "fortishield"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# emailfilter - fortishield"),
				},
			},
			"fortios_emailfilter_iptrust": {
				Tok: tfbridge.MakeResource(mainPkg, emailfilterMod, "iptrust"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# emailfilter - iptrust"),
				},
			},
			"fortios_emailfilter_mheader": {
				Tok: tfbridge.MakeResource(mainPkg, emailfilterMod, "mheader"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# emailfilter - mheader"),
				},
			},
			"fortios_emailfilter_options": {
				Tok: tfbridge.MakeResource(mainPkg, emailfilterMod, "options"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# emailfilter - options"),
				},
			},
			"fortios_emailfilter_profile": {
				Tok: tfbridge.MakeResource(mainPkg, emailfilterMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# emailfilter - profile"),
				},
			},
			"fortios_endpointcontrol_client": {
				Tok: tfbridge.MakeResource(mainPkg, endpointcontrolMod, "client"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# endpointcontrol - client"),
				},
			},
			"fortios_endpointcontrol_fctems": {
				Tok: tfbridge.MakeResource(mainPkg, endpointcontrolMod, "fctems"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# endpointcontrol - fctems"),
				},
			},
			"fortios_endpointcontrol_forticlientems": {
				Tok: tfbridge.MakeResource(mainPkg, endpointcontrolMod, "forticlientems"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# endpointcontrol - forticlientems"),
				},
			},
			"fortios_endpointcontrol_forticlientregistrationsync": {
				Tok: tfbridge.MakeResource(mainPkg, endpointcontrolMod, "forticlientregistrationsync"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# endpointcontrol - forticlientregistrationsync"),
				},
			},
			"fortios_endpointcontrol_profile": {
				Tok: tfbridge.MakeResource(mainPkg, endpointcontrolMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# endpointcontrol - profile"),
				},
			},
			"fortios_endpointcontrol_registeredforticlient": {
				Tok: tfbridge.MakeResource(mainPkg, endpointcontrolMod, "registeredforticlient"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# endpointcontrol - registeredforticlient"),
				},
			},
			"fortios_endpointcontrol_settings": {
				Tok: tfbridge.MakeResource(mainPkg, endpointcontrolMod, "settings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# endpointcontrol - settings"),
				},
			},
			"fortios_extendercontroller_dataplan": {
				Tok: tfbridge.MakeResource(mainPkg, extendercontrollerMod, "dataplan"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# extendercontroller - dataplan"),
				},
			},
			"fortios_extendercontroller_extender": {
				Tok: tfbridge.MakeResource(mainPkg, extendercontrollerMod, "extender"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# extendercontroller - extender"),
				},
			},
			"fortios_extendercontroller_extender1": {
				Tok: tfbridge.MakeResource(mainPkg, extendercontrollerMod, "extender1"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# extendercontroller - extender1"),
				},
			},
			"fortios_extendercontroller_extenderprofile": {
				Tok: tfbridge.MakeResource(mainPkg, extendercontrollerMod, "extenderprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# extendercontroller - extenderprofile"),
				},
			},
			"fortios_extensioncontroller_dataplan": {
				Tok: tfbridge.MakeResource(mainPkg, extensioncontrollerMod, "dataplan"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# extensioncontroller - dataplan"),
				},
			},
			"fortios_extensioncontroller_extender": {
				Tok: tfbridge.MakeResource(mainPkg, extensioncontrollerMod, "extender"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# extensioncontroller - extender"),
				},
			},
			"fortios_extensioncontroller_extenderprofile": {
				Tok: tfbridge.MakeResource(mainPkg, extensioncontrollerMod, "extenderprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# extensioncontroller - extenderprofile"),
				},
			},
			"fortios_extensioncontroller_fortigate": {
				Tok: tfbridge.MakeResource(mainPkg, extensioncontrollerMod, "fortigate"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# extensioncontroller - fortigate"),
				},
			},
			"fortios_extensioncontroller_fortigateprofile": {
				Tok: tfbridge.MakeResource(mainPkg, extensioncontrollerMod, "fortigateprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# extensioncontroller - fortigateprofile"),
				},
			},
			"fortios_filefilter_profile": {
				Tok: tfbridge.MakeResource(mainPkg, filefilterMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# filefilter - profile"),
				},
			},
			"fortios_firewall_DoSpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "DoSpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - DoSpolicy"),
				},
			},
			"fortios_firewall_DoSpolicy6": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "DoSpolicy6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - DoSpolicy6"),
				},
			},
			"fortios_firewall_accessproxy": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "accessproxy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - accessproxy"),
				},
			},
			"fortios_firewall_accessproxy6": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "accessproxy6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - accessproxy6"),
				},
			},
			"fortios_firewall_accessproxysshclientcert": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "accessproxysshclientcert"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - accessproxysshclientcert"),
				},
			},
			"fortios_firewall_accessproxyvirtualhost": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "accessproxyvirtualhost"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - accessproxyvirtualhost"),
				},
			},
			"fortios_firewall_address": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "address"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - address"),
				},
			},
			"fortios_firewall_address6": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "address6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - address6"),
				},
			},
			"fortios_firewall_address6template": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "address6template"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - address6template"),
				},
			},
			"fortios_firewall_addrgrp": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "addrgrp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - addrgrp"),
				},
			},
			"fortios_firewall_addrgrp6": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "addrgrp6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - addrgrp6"),
				},
			},
			"fortios_firewall_authportal": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "authportal"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - authportal"),
				},
			},
			"fortios_firewall_centralsnatmap": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "centralsnatmap"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - centralsnatmap"),
				},
			},
			"fortios_firewall_centralsnatmap_move": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "centralsnatmap_move"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - centralsnatmap_move"),
				},
			},
			"fortios_firewall_centralsnatmap_sort": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "centralsnatmap_sort"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - centralsnatmap_sort"),
				},
			},
			"fortios_firewall_city": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "city"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - city"),
				},
			},
			"fortios_firewall_country": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "country"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - country"),
				},
			},
			"fortios_firewall_decryptedtrafficmirror": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "decryptedtrafficmirror"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - decryptedtrafficmirror"),
				},
			},
			"fortios_firewall_dnstranslation": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "dnstranslation"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - dnstranslation"),
				},
			},
			"fortios_firewall_global": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "global"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - global"),
				},
			},
			"fortios_firewall_identitybasedroute": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "identitybasedroute"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - identitybasedroute"),
				},
			},
			"fortios_firewall_interfacepolicy": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "interfacepolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - interfacepolicy"),
				},
			},
			"fortios_firewall_interfacepolicy6": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "interfacepolicy6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - interfacepolicy6"),
				},
			},
			"fortios_firewall_internetservice": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetservice"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservice"),
				},
			},
			"fortios_firewall_internetserviceaddition": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetserviceaddition"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetserviceaddition"),
				},
			},
			"fortios_firewall_internetserviceappend": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetserviceappend"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetserviceappend"),
				},
			},
			"fortios_firewall_internetservicebotnet": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetservicebotnet"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicebotnet"),
				},
			},
			"fortios_firewall_internetservicecustom": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetservicecustom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicecustom"),
				},
			},
			"fortios_firewall_internetservicecustomgroup": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetservicecustomgroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicecustomgroup"),
				},
			},
			"fortios_firewall_internetservicedefinition": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetservicedefinition"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicedefinition"),
				},
			},
			"fortios_firewall_internetserviceextension": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetserviceextension"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetserviceextension"),
				},
			},
			"fortios_firewall_internetservicegroup": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetservicegroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicegroup"),
				},
			},
			"fortios_firewall_internetserviceipblreason": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetserviceipblreason"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetserviceipblreason"),
				},
			},
			"fortios_firewall_internetserviceipblvendor": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetserviceipblvendor"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetserviceipblvendor"),
				},
			},
			"fortios_firewall_internetservicelist": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetservicelist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicelist"),
				},
			},
			"fortios_firewall_internetservicename": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetservicename"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicename"),
				},
			},
			"fortios_firewall_internetserviceowner": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetserviceowner"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetserviceowner"),
				},
			},
			"fortios_firewall_internetservicereputation": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "internetservicereputation"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicereputation"),
				},
			},
			"fortios_firewall_ippool": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "ippool"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - ippool"),
				},
			},
			"fortios_firewall_ippool6": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "ippool6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - ippool6"),
				},
			},
			"fortios_firewall_iptranslation": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "iptranslation"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - iptranslation"),
				},
			},
			"fortios_firewall_ipv6ehfilter": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "ipv6ehfilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - ipv6ehfilter"),
				},
			},
			"fortios_firewall_ldbmonitor": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "ldbmonitor"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - ldbmonitor"),
				},
			},
			"fortios_firewall_localinpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "localinpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - localinpolicy"),
				},
			},
			"fortios_firewall_localinpolicy6": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "localinpolicy6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - localinpolicy6"),
				},
			},
			"fortios_firewall_multicastaddress": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "multicastaddress"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - multicastaddress"),
				},
			},
			"fortios_firewall_multicastaddress6": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "multicastaddress6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - multicastaddress6"),
				},
			},
			"fortios_firewall_multicastpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "multicastpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - multicastpolicy"),
				},
			},
			"fortios_firewall_multicastpolicy6": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "multicastpolicy6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - multicastpolicy6"),
				},
			},
			"fortios_firewall_networkservicedynamic": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "networkservicedynamic"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - networkservicedynamic"),
				},
			},
			"fortios_firewall_object_address": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "object_address"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - object_address"),
				},
			},
			"fortios_firewall_object_addressgroup": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "object_addressgroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - object_addressgroup"),
				},
			},
			"fortios_firewall_object_ippool": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "object_ippool"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - object_ippool"),
				},
			},
			"fortios_firewall_object_service": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "object_service"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - object_service"),
				},
			},
			"fortios_firewall_object_servicecategory": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "object_servicecategory"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - object_servicecategory"),
				},
			},
			"fortios_firewall_object_servicegroup": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "object_servicegroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - object_servicegroup"),
				},
			},
			"fortios_firewall_object_vip": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "object_vip"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - object_vip"),
				},
			},
			"fortios_firewall_object_vipgroup": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "object_vipgroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - object_vipgroup"),
				},
			},
			"fortios_firewall_policy": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "policy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - policy"),
				},
			},
			"fortios_firewall_policy46": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "policy46"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - policy46"),
				},
			},
			"fortios_firewall_policy6": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "policy6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - policy6"),
				},
			},
			"fortios_firewall_policy64": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "policy64"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - policy64"),
				},
			},
			"fortios_firewall_profilegroup": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "profilegroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - profilegroup"),
				},
			},
			"fortios_firewall_profileprotocoloptions": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "profileprotocoloptions"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - profileprotocoloptions"),
				},
			},
			"fortios_firewall_proxyaddress": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "proxyaddress"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - proxyaddress"),
				},
			},
			"fortios_firewall_proxyaddrgrp": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "proxyaddrgrp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - proxyaddrgrp"),
				},
			},
			"fortios_firewall_proxypolicy": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "proxypolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - proxypolicy"),
				},
			},
			"fortios_firewall_proxypolicy_move": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "proxypolicy_move"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - proxypolicy_move"),
				},
			},
			"fortios_firewall_proxypolicy_sort": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "proxypolicy_sort"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - proxypolicy_sort"),
				},
			},
			"fortios_firewall_region": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "region"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - region"),
				},
			},
			"fortios_firewall_security_policy": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "security_policy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - security_policy"),
				},
			},
			"fortios_firewall_security_policyseq": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "security_policyseq"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - security_policyseq"),
				},
			},
			"fortios_firewall_security_policysort": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "security_policysort"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - security_policysort"),
				},
			},
			"fortios_firewall_securitypolicy": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "securitypolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - securitypolicy"),
				},
			},
			"fortios_firewall_shapingpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "shapingpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - shapingpolicy"),
				},
			},
			"fortios_firewall_shapingprofile": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "shapingprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - shapingprofile"),
				},
			},
			"fortios_firewall_sniffer": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "sniffer"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - sniffer"),
				},
			},
			"fortios_firewall_sslserver": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "sslserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - sslserver"),
				},
			},
			"fortios_firewall_sslsshprofile": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "sslsshprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - sslsshprofile"),
				},
			},
			"fortios_firewall_trafficclass": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "trafficclass"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - trafficclass"),
				},
			},
			"fortios_firewall_ttlpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "ttlpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - ttlpolicy"),
				},
			},
			"fortios_firewall_vendormac": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "vendormac"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - vendormac"),
				},
			},
			"fortios_firewall_vip": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "vip"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - vip"),
				},
			},
			"fortios_firewall_vip46": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "vip46"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - vip46"),
				},
			},
			"fortios_firewall_vip6": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "vip6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - vip6"),
				},
			},
			"fortios_firewall_vip64": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "vip64"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - vip64"),
				},
			},
			"fortios_firewall_vipgrp": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "vipgrp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - vipgrp"),
				},
			},
			"fortios_firewall_vipgrp46": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "vipgrp46"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - vipgrp46"),
				},
			},
			"fortios_firewall_vipgrp6": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "vipgrp6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - vipgrp6"),
				},
			},
			"fortios_firewall_vipgrp64": {
				Tok: tfbridge.MakeResource(mainPkg, firewallMod, "vipgrp64"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - vipgrp64"),
				},
			},
			"fortios_firewallconsolidated_policy": {
				Tok: tfbridge.MakeResource(mainPkg, firewallconsolidatedMod, "policy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallconsolidated - policy"),
				},
			},
			"fortios_firewallipmacbinding_setting": {
				Tok: tfbridge.MakeResource(mainPkg, firewallipmacbindingMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallipmacbinding - setting"),
				},
			},
			"fortios_firewallipmacbinding_table": {
				Tok: tfbridge.MakeResource(mainPkg, firewallipmacbindingMod, "table"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallipmacbinding - table"),
				},
			},
			"fortios_firewallschedule_group": {
				Tok: tfbridge.MakeResource(mainPkg, firewallscheduleMod, "group"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallschedule - group"),
				},
			},
			"fortios_firewallschedule_onetime": {
				Tok: tfbridge.MakeResource(mainPkg, firewallscheduleMod, "onetime"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallschedule - onetime"),
				},
			},
			"fortios_firewallschedule_recurring": {
				Tok: tfbridge.MakeResource(mainPkg, firewallscheduleMod, "recurring"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallschedule - recurring"),
				},
			},
			"fortios_firewallservice_category": {
				Tok: tfbridge.MakeResource(mainPkg, firewallserviceMod, "category"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallservice - category"),
				},
			},
			"fortios_firewallservice_custom": {
				Tok: tfbridge.MakeResource(mainPkg, firewallserviceMod, "custom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallservice - custom"),
				},
			},
			"fortios_firewallservice_group": {
				Tok: tfbridge.MakeResource(mainPkg, firewallserviceMod, "group"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallservice - group"),
				},
			},
			"fortios_firewallshaper_peripshaper": {
				Tok: tfbridge.MakeResource(mainPkg, firewallshaperMod, "peripshaper"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallshaper - peripshaper"),
				},
			},
			"fortios_firewallshaper_trafficshaper": {
				Tok: tfbridge.MakeResource(mainPkg, firewallshaperMod, "trafficshaper"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallshaper - trafficshaper"),
				},
			},
			"fortios_firewallssh_hostkey": {
				Tok: tfbridge.MakeResource(mainPkg, firewallsshMod, "hostkey"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallssh - hostkey"),
				},
			},
			"fortios_firewallssh_localca": {
				Tok: tfbridge.MakeResource(mainPkg, firewallsshMod, "localca"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallssh - localca"),
				},
			},
			"fortios_firewallssh_localkey": {
				Tok: tfbridge.MakeResource(mainPkg, firewallsshMod, "localkey"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallssh - localkey"),
				},
			},
			"fortios_firewallssh_setting": {
				Tok: tfbridge.MakeResource(mainPkg, firewallsshMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallssh - setting"),
				},
			},
			"fortios_firewallssl_setting": {
				Tok: tfbridge.MakeResource(mainPkg, firewallsslMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallssl - setting"),
				},
			},
			"fortios_firewallwildcardfqdn_custom": {
				Tok: tfbridge.MakeResource(mainPkg, firewallwildcardfqdnMod, "custom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallwildcardfqdn - custom"),
				},
			},
			"fortios_firewallwildcardfqdn_group": {
				Tok: tfbridge.MakeResource(mainPkg, firewallwildcardfqdnMod, "group"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallwildcardfqdn - group"),
				},
			},
			"fortios_fmg_devicemanager_device": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "devicemanager_device"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - devicemanager_device"),
				},
			},
			"fortios_fmg_devicemanager_install_device": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "devicemanager_install_device"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - devicemanager_install_device"),
				},
			},
			"fortios_fmg_devicemanager_install_policypackage": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "devicemanager_install_policypackage"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - devicemanager_install_policypackage"),
				},
			},
			"fortios_fmg_devicemanager_script": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "devicemanager_script"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - devicemanager_script"),
				},
			},
			"fortios_fmg_devicemanager_script_execute": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "devicemanager_script_execute"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - devicemanager_script_execute"),
				},
			},
			"fortios_fmg_firewall_object_address": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "firewall_object_address"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - firewall_object_address"),
				},
			},
			"fortios_fmg_firewall_object_ippool": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "firewall_object_ippool"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - firewall_object_ippool"),
				},
			},
			"fortios_fmg_firewall_object_service": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "firewall_object_service"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - firewall_object_service"),
				},
			},
			"fortios_fmg_firewall_object_vip": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "firewall_object_vip"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - firewall_object_vip"),
				},
			},
			"fortios_fmg_firewall_security_policy": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "firewall_security_policy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - firewall_security_policy"),
				},
			},
			"fortios_fmg_firewall_security_policypackage": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "firewall_security_policypackage"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - firewall_security_policypackage"),
				},
			},
			"fortios_fmg_jsonrpc_request": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "jsonrpc_request"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - jsonrpc_request"),
				},
			},
			"fortios_fmg_object_adom_revision": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "object_adom_revision"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - object_adom_revision"),
				},
			},
			"fortios_fmg_system_admin": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "system_admin"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - system_admin"),
				},
			},
			"fortios_fmg_system_admin_profiles": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "system_admin_profiles"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - system_admin_profiles"),
				},
			},
			"fortios_fmg_system_admin_user": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "system_admin_user"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - system_admin_user"),
				},
			},
			"fortios_fmg_system_adom": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "system_adom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - system_adom"),
				},
			},
			"fortios_fmg_system_dns": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "system_dns"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - system_dns"),
				},
			},
			"fortios_fmg_system_global": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "system_global"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - system_global"),
				},
			},
			"fortios_fmg_system_license_forticare": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "system_license_forticare"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - system_license_forticare"),
				},
			},
			"fortios_fmg_system_license_vm": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "system_license_vm"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - system_license_vm"),
				},
			},
			"fortios_fmg_system_network_interface": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "system_network_interface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - system_network_interface"),
				},
			},
			"fortios_fmg_system_network_route": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "system_network_route"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - system_network_route"),
				},
			},
			"fortios_fmg_system_ntp": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "system_ntp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - system_ntp"),
				},
			},
			"fortios_fmg_system_syslogserver": {
				Tok: tfbridge.MakeResource(mainPkg, fmgMod, "system_syslogserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# fmg - system_syslogserver"),
				},
			},
			"fortios_ftpproxy_explicit": {
				Tok: tfbridge.MakeResource(mainPkg, ftpproxyMod, "explicit"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# ftpproxy - explicit"),
				},
			},
			"fortios_icap_profile": {
				Tok: tfbridge.MakeResource(mainPkg, icapMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# icap - profile"),
				},
			},
			"fortios_icap_server": {
				Tok: tfbridge.MakeResource(mainPkg, icapMod, "server"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# icap - server"),
				},
			},
			"fortios_icap_servergroup": {
				Tok: tfbridge.MakeResource(mainPkg, icapMod, "servergroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# icap - servergroup"),
				},
			},
			"fortios_ips_custom": {
				Tok: tfbridge.MakeResource(mainPkg, ipsMod, "custom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# ips - custom"),
				},
			},
			"fortios_ips_decoder": {
				Tok: tfbridge.MakeResource(mainPkg, ipsMod, "decoder"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# ips - decoder"),
				},
			},
			"fortios_ips_global": {
				Tok: tfbridge.MakeResource(mainPkg, ipsMod, "global"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# ips - global"),
				},
			},
			"fortios_ips_rule": {
				Tok: tfbridge.MakeResource(mainPkg, ipsMod, "rule"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# ips - rule"),
				},
			},
			"fortios_ips_rulesettings": {
				Tok: tfbridge.MakeResource(mainPkg, ipsMod, "rulesettings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# ips - rulesettings"),
				},
			},
			"fortios_ips_sensor": {
				Tok: tfbridge.MakeResource(mainPkg, ipsMod, "sensor"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# ips - sensor"),
				},
			},
			"fortios_ips_settings": {
				Tok: tfbridge.MakeResource(mainPkg, ipsMod, "settings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# ips - settings"),
				},
			},
			"fortios_ips_viewmap": {
				Tok: tfbridge.MakeResource(mainPkg, ipsMod, "viewmap"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# ips - viewmap"),
				},
			},
			"fortios_json_generic_api": {
				Tok: tfbridge.MakeResource(mainPkg, jsonMod, "generic_api"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# json - generic_api"),
				},
			},
			"fortios_log_customfield": {
				Tok: tfbridge.MakeResource(mainPkg, logMod, "customfield"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# log - customfield"),
				},
			},
			"fortios_log_eventfilter": {
				Tok: tfbridge.MakeResource(mainPkg, logMod, "eventfilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# log - eventfilter"),
				},
			},
			"fortios_log_fortianalyzer_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logMod, "fortianalyzer_setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# log - fortianalyzer_setting"),
				},
			},
			"fortios_log_guidisplay": {
				Tok: tfbridge.MakeResource(mainPkg, logMod, "guidisplay"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# log - guidisplay"),
				},
			},
			"fortios_log_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# log - setting"),
				},
			},
			"fortios_log_syslog_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logMod, "syslog_setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# log - syslog_setting"),
				},
			},
			"fortios_log_threatweight": {
				Tok: tfbridge.MakeResource(mainPkg, logMod, "threatweight"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# log - threatweight"),
				},
			},
			"fortios_logdisk_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logdiskMod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logdisk - filter"),
				},
			},
			"fortios_logdisk_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logdiskMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logdisk - setting"),
				},
			},
			"fortios_logfortianalyzer2_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzer2Mod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzer2 - filter"),
				},
			},
			"fortios_logfortianalyzer2_overridefilter": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzer2Mod, "overridefilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzer2 - overridefilter"),
				},
			},
			"fortios_logfortianalyzer2_overridesetting": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzer2Mod, "overridesetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzer2 - overridesetting"),
				},
			},
			"fortios_logfortianalyzer2_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzer2Mod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzer2 - setting"),
				},
			},
			"fortios_logfortianalyzer3_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzer3Mod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzer3 - filter"),
				},
			},
			"fortios_logfortianalyzer3_overridefilter": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzer3Mod, "overridefilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzer3 - overridefilter"),
				},
			},
			"fortios_logfortianalyzer3_overridesetting": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzer3Mod, "overridesetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzer3 - overridesetting"),
				},
			},
			"fortios_logfortianalyzer3_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzer3Mod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzer3 - setting"),
				},
			},
			"fortios_logfortianalyzer_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzerMod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzer - filter"),
				},
			},
			"fortios_logfortianalyzer_overridefilter": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzerMod, "overridefilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzer - overridefilter"),
				},
			},
			"fortios_logfortianalyzer_overridesetting": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzerMod, "overridesetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzer - overridesetting"),
				},
			},
			"fortios_logfortianalyzer_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzerMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzer - setting"),
				},
			},
			"fortios_logfortianalyzercloud_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzercloudMod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzercloud - filter"),
				},
			},
			"fortios_logfortianalyzercloud_overridefilter": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzercloudMod, "overridefilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzercloud - overridefilter"),
				},
			},
			"fortios_logfortianalyzercloud_overridesetting": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzercloudMod, "overridesetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzercloud - overridesetting"),
				},
			},
			"fortios_logfortianalyzercloud_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logfortianalyzercloudMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortianalyzercloud - setting"),
				},
			},
			"fortios_logfortiguard_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logfortiguardMod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortiguard - filter"),
				},
			},
			"fortios_logfortiguard_overridefilter": {
				Tok: tfbridge.MakeResource(mainPkg, logfortiguardMod, "overridefilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortiguard - overridefilter"),
				},
			},
			"fortios_logfortiguard_overridesetting": {
				Tok: tfbridge.MakeResource(mainPkg, logfortiguardMod, "overridesetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortiguard - overridesetting"),
				},
			},
			"fortios_logfortiguard_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logfortiguardMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logfortiguard - setting"),
				},
			},
			"fortios_logmemory_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logmemoryMod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logmemory - filter"),
				},
			},
			"fortios_logmemory_globalsetting": {
				Tok: tfbridge.MakeResource(mainPkg, logmemoryMod, "globalsetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logmemory - globalsetting"),
				},
			},
			"fortios_logmemory_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logmemoryMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logmemory - setting"),
				},
			},
			"fortios_lognulldevice_filter": {
				Tok: tfbridge.MakeResource(mainPkg, lognulldeviceMod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# lognulldevice - filter"),
				},
			},
			"fortios_lognulldevice_setting": {
				Tok: tfbridge.MakeResource(mainPkg, lognulldeviceMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# lognulldevice - setting"),
				},
			},
			"fortios_logsyslogd2_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogd2Mod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd2 - filter"),
				},
			},
			"fortios_logsyslogd2_overridefilter": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogd2Mod, "overridefilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd2 - overridefilter"),
				},
			},
			"fortios_logsyslogd2_overridesetting": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogd2Mod, "overridesetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd2 - overridesetting"),
				},
			},
			"fortios_logsyslogd2_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogd2Mod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd2 - setting"),
				},
			},
			"fortios_logsyslogd3_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogd3Mod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd3 - filter"),
				},
			},
			"fortios_logsyslogd3_overridefilter": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogd3Mod, "overridefilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd3 - overridefilter"),
				},
			},
			"fortios_logsyslogd3_overridesetting": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogd3Mod, "overridesetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd3 - overridesetting"),
				},
			},
			"fortios_logsyslogd3_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogd3Mod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd3 - setting"),
				},
			},
			"fortios_logsyslogd4_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogd4Mod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd4 - filter"),
				},
			},
			"fortios_logsyslogd4_overridefilter": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogd4Mod, "overridefilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd4 - overridefilter"),
				},
			},
			"fortios_logsyslogd4_overridesetting": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogd4Mod, "overridesetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd4 - overridesetting"),
				},
			},
			"fortios_logsyslogd4_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogd4Mod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd4 - setting"),
				},
			},
			"fortios_logsyslogd_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogdMod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd - filter"),
				},
			},
			"fortios_logsyslogd_overridefilter": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogdMod, "overridefilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd - overridefilter"),
				},
			},
			"fortios_logsyslogd_overridesetting": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogdMod, "overridesetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd - overridesetting"),
				},
			},
			"fortios_logsyslogd_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logsyslogdMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logsyslogd - setting"),
				},
			},
			"fortios_logtacacsaccounting2_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logtacacsaccounting2Mod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logtacacsaccounting2 - filter"),
				},
			},
			"fortios_logtacacsaccounting2_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logtacacsaccounting2Mod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logtacacsaccounting2 - setting"),
				},
			},
			"fortios_logtacacsaccounting3_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logtacacsaccounting3Mod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logtacacsaccounting3 - filter"),
				},
			},
			"fortios_logtacacsaccounting3_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logtacacsaccounting3Mod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logtacacsaccounting3 - setting"),
				},
			},
			"fortios_logtacacsaccounting_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logtacacsaccountingMod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logtacacsaccounting - filter"),
				},
			},
			"fortios_logtacacsaccounting_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logtacacsaccountingMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logtacacsaccounting - setting"),
				},
			},
			"fortios_logwebtrends_filter": {
				Tok: tfbridge.MakeResource(mainPkg, logwebtrendsMod, "filter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logwebtrends - filter"),
				},
			},
			"fortios_logwebtrends_setting": {
				Tok: tfbridge.MakeResource(mainPkg, logwebtrendsMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# logwebtrends - setting"),
				},
			},
			"fortios_networking_interface_port": {
				Tok: tfbridge.MakeResource(mainPkg, networkingMod, "interface_port"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# networking - interface_port"),
				},
			},
			"fortios_networking_route_static": {
				Tok: tfbridge.MakeResource(mainPkg, networkingMod, "route_static"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# networking - route_static"),
				},
			},
			"fortios_nsxt_servicechain": {
				Tok: tfbridge.MakeResource(mainPkg, nsxtMod, "servicechain"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# nsxt - servicechain"),
				},
			},
			"fortios_nsxt_setting": {
				Tok: tfbridge.MakeResource(mainPkg, nsxtMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# nsxt - setting"),
				},
			},
			"fortios_report_chart": {
				Tok: tfbridge.MakeResource(mainPkg, reportMod, "chart"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# report - chart"),
				},
			},
			"fortios_report_dataset": {
				Tok: tfbridge.MakeResource(mainPkg, reportMod, "dataset"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# report - dataset"),
				},
			},
			"fortios_report_layout": {
				Tok: tfbridge.MakeResource(mainPkg, reportMod, "layout"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# report - layout"),
				},
			},
			"fortios_report_setting": {
				Tok: tfbridge.MakeResource(mainPkg, reportMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# report - setting"),
				},
			},
			"fortios_report_style": {
				Tok: tfbridge.MakeResource(mainPkg, reportMod, "style"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# report - style"),
				},
			},
			"fortios_report_theme": {
				Tok: tfbridge.MakeResource(mainPkg, reportMod, "theme"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# report - theme"),
				},
			},
			"fortios_router_accesslist": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "accesslist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - accesslist"),
				},
			},
			"fortios_router_accesslist6": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "accesslist6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - accesslist6"),
				},
			},
			"fortios_router_aspathlist": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "aspathlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - aspathlist"),
				},
			},
			"fortios_router_authpath": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "authpath"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - authpath"),
				},
			},
			"fortios_router_bfd": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "bfd"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - bfd"),
				},
			},
			"fortios_router_bfd6": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "bfd6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - bfd6"),
				},
			},
			"fortios_router_bgp": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "bgp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - bgp"),
				},
			},
			"fortios_router_communitylist": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "communitylist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - communitylist"),
				},
			},
			"fortios_router_isis": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "isis"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - isis"),
				},
			},
			"fortios_router_keychain": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "keychain"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - keychain"),
				},
			},
			"fortios_router_multicast": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "multicast"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - multicast"),
				},
			},
			"fortios_router_multicast6": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "multicast6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - multicast6"),
				},
			},
			"fortios_router_multicastflow": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "multicastflow"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - multicastflow"),
				},
			},
			"fortios_router_ospf": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "ospf"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - ospf"),
				},
			},
			"fortios_router_ospf6": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "ospf6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - ospf6"),
				},
			},
			"fortios_router_policy": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "policy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - policy"),
				},
			},
			"fortios_router_policy6": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "policy6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - policy6"),
				},
			},
			"fortios_router_prefixlist": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "prefixlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - prefixlist"),
				},
			},
			"fortios_router_prefixlist6": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "prefixlist6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - prefixlist6"),
				},
			},
			"fortios_router_rip": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "rip"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - rip"),
				},
			},
			"fortios_router_ripng": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "ripng"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - ripng"),
				},
			},
			"fortios_router_routemap": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "routemap"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - routemap"),
				},
			},
			"fortios_router_setting": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - setting"),
				},
			},
			"fortios_router_static": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "static"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - static"),
				},
			},
			"fortios_router_static6": {
				Tok: tfbridge.MakeResource(mainPkg, routerMod, "static6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - static6"),
				},
			},
			"fortios_routerbgp_neighbor": {
				Tok: tfbridge.MakeResource(mainPkg, routerbgpMod, "neighbor"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# routerbgp - neighbor"),
				},
			},
			"fortios_routerbgp_network": {
				Tok: tfbridge.MakeResource(mainPkg, routerbgpMod, "network"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# routerbgp - network"),
				},
			},
			"fortios_routerbgp_network6": {
				Tok: tfbridge.MakeResource(mainPkg, routerbgpMod, "network6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# routerbgp - network6"),
				},
			},
			"fortios_routerospf6_ospf6interface": {
				Tok: tfbridge.MakeResource(mainPkg, routerospf6Mod, "ospf6interface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# routerospf6 - ospf6interface"),
				},
			},
			"fortios_routerospf_neighbor": {
				Tok: tfbridge.MakeResource(mainPkg, routerospfMod, "neighbor"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# routerospf - neighbor"),
				},
			},
			"fortios_routerospf_network": {
				Tok: tfbridge.MakeResource(mainPkg, routerospfMod, "network"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# routerospf - network"),
				},
			},
			"fortios_routerospf_ospfinterface": {
				Tok: tfbridge.MakeResource(mainPkg, routerospfMod, "ospfinterface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# routerospf - ospfinterface"),
				},
			},
			"fortios_sctpfilter_profile": {
				Tok: tfbridge.MakeResource(mainPkg, sctpfilterMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# sctpfilter - profile"),
				},
			},
			"fortios_spamfilter_bwl": {
				Tok: tfbridge.MakeResource(mainPkg, spamfilterMod, "bwl"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# spamfilter - bwl"),
				},
			},
			"fortios_spamfilter_bword": {
				Tok: tfbridge.MakeResource(mainPkg, spamfilterMod, "bword"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# spamfilter - bword"),
				},
			},
			"fortios_spamfilter_dnsbl": {
				Tok: tfbridge.MakeResource(mainPkg, spamfilterMod, "dnsbl"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# spamfilter - dnsbl"),
				},
			},
			"fortios_spamfilter_fortishield": {
				Tok: tfbridge.MakeResource(mainPkg, spamfilterMod, "fortishield"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# spamfilter - fortishield"),
				},
			},
			"fortios_spamfilter_iptrust": {
				Tok: tfbridge.MakeResource(mainPkg, spamfilterMod, "iptrust"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# spamfilter - iptrust"),
				},
			},
			"fortios_spamfilter_mheader": {
				Tok: tfbridge.MakeResource(mainPkg, spamfilterMod, "mheader"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# spamfilter - mheader"),
				},
			},
			"fortios_spamfilter_options": {
				Tok: tfbridge.MakeResource(mainPkg, spamfilterMod, "options"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# spamfilter - options"),
				},
			},
			"fortios_spamfilter_profile": {
				Tok: tfbridge.MakeResource(mainPkg, spamfilterMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# spamfilter - profile"),
				},
			},
			"fortios_sshfilter_profile": {
				Tok: tfbridge.MakeResource(mainPkg, sshfilterMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# sshfilter - profile"),
				},
			},
			"fortios_switchcontroller_8021Xsettings": {
				// Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "8021Xsettings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - 8021Xsettings"),
				},
			},
			"fortios_switchcontroller_customcommand": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "customcommand"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - customcommand"),
				},
			},
			"fortios_switchcontroller_dynamicportpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "dynamicportpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - dynamicportpolicy"),
				},
			},
			"fortios_switchcontroller_flowtracking": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "flowtracking"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - flowtracking"),
				},
			},
			"fortios_switchcontroller_fortilinksettings": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "fortilinksettings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - fortilinksettings"),
				},
			},
			"fortios_switchcontroller_global": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "global"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - global"),
				},
			},
			"fortios_switchcontroller_igmpsnooping": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "igmpsnooping"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - igmpsnooping"),
				},
			},
			"fortios_switchcontroller_lldpprofile": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "lldpprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - lldpprofile"),
				},
			},
			"fortios_switchcontroller_lldpsettings": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "lldpsettings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - lldpsettings"),
				},
			},
			"fortios_switchcontroller_location": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "location"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - location"),
				},
			},
			"fortios_switchcontroller_macsyncsettings": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "macsyncsettings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - macsyncsettings"),
				},
			},
			"fortios_switchcontroller_managedswitch": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "managedswitch"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - managedswitch"),
				},
			},
			"fortios_switchcontroller_nacdevice": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "nacdevice"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - nacdevice"),
				},
			},
			"fortios_switchcontroller_nacsettings": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "nacsettings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - nacsettings"),
				},
			},
			"fortios_switchcontroller_networkmonitorsettings": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "networkmonitorsettings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - networkmonitorsettings"),
				},
			},
			"fortios_switchcontroller_portpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "portpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - portpolicy"),
				},
			},
			"fortios_switchcontroller_quarantine": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "quarantine"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - quarantine"),
				},
			},
			"fortios_switchcontroller_remotelog": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "remotelog"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - remotelog"),
				},
			},
			"fortios_switchcontroller_sflow": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "sflow"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - sflow"),
				},
			},
			"fortios_switchcontroller_snmpcommunity": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "snmpcommunity"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - snmpcommunity"),
				},
			},
			"fortios_switchcontroller_snmpsysinfo": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "snmpsysinfo"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - snmpsysinfo"),
				},
			},
			"fortios_switchcontroller_snmptrapthreshold": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "snmptrapthreshold"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - snmptrapthreshold"),
				},
			},
			"fortios_switchcontroller_snmpuser": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "snmpuser"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - snmpuser"),
				},
			},
			"fortios_switchcontroller_stormcontrol": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "stormcontrol"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - stormcontrol"),
				},
			},
			"fortios_switchcontroller_stormcontrolpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "stormcontrolpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - stormcontrolpolicy"),
				},
			},
			"fortios_switchcontroller_stpinstance": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "stpinstance"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - stpinstance"),
				},
			},
			"fortios_switchcontroller_stpsettings": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "stpsettings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - stpsettings"),
				},
			},
			"fortios_switchcontroller_switchgroup": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "switchgroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - switchgroup"),
				},
			},
			"fortios_switchcontroller_switchinterfacetag": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "switchinterfacetag"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - switchinterfacetag"),
				},
			},
			"fortios_switchcontroller_switchlog": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "switchlog"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - switchlog"),
				},
			},
			"fortios_switchcontroller_switchprofile": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "switchprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - switchprofile"),
				},
			},
			"fortios_switchcontroller_system": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "system"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - system"),
				},
			},
			"fortios_switchcontroller_trafficpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "trafficpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - trafficpolicy"),
				},
			},
			"fortios_switchcontroller_trafficsniffer": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "trafficsniffer"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - trafficsniffer"),
				},
			},
			"fortios_switchcontroller_virtualportpool": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "virtualportpool"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - virtualportpool"),
				},
			},
			"fortios_switchcontroller_vlan": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "vlan"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - vlan"),
				},
			},
			"fortios_switchcontroller_vlanpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerMod, "vlanpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontroller - vlanpolicy"),
				},
			},
			"fortios_switchcontrollerautoconfig_custom": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerautoconfigMod, "custom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollerautoconfig - custom"),
				},
			},
			"fortios_switchcontrollerautoconfig_default": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerautoconfigMod, "default"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollerautoconfig - default"),
				},
			},
			"fortios_switchcontrollerautoconfig_policy": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerautoconfigMod, "policy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollerautoconfig - policy"),
				},
			},
			"fortios_switchcontrollerinitialconfig_template": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerinitialconfigMod, "template"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollerinitialconfig - template"),
				},
			},
			"fortios_switchcontrollerinitialconfig_vlans": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerinitialconfigMod, "vlans"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollerinitialconfig - vlans"),
				},
			},
			"fortios_switchcontrollerptp_policy": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerptpMod, "policy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollerptp - policy"),
				},
			},
			"fortios_switchcontrollerptp_settings": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerptpMod, "settings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollerptp - settings"),
				},
			},
			"fortios_switchcontrollerqos_dot1pmap": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerqosMod, "dot1pmap"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollerqos - dot1pmap"),
				},
			},
			"fortios_switchcontrollerqos_ipdscpmap": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerqosMod, "ipdscpmap"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollerqos - ipdscpmap"),
				},
			},
			"fortios_switchcontrollerqos_qospolicy": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerqosMod, "qospolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollerqos - qospolicy"),
				},
			},
			"fortios_switchcontrollerqos_queuepolicy": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollerqosMod, "queuepolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollerqos - queuepolicy"),
				},
			},
			"fortios_switchcontrollersecuritypolicy_8021X": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollersecuritypolicyMod, "8021X"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollersecuritypolicy - 8021X"),
				},
			},
			"fortios_switchcontrollersecuritypolicy_captiveportal": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollersecuritypolicyMod, "captiveportal"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollersecuritypolicy - captiveportal"),
				},
			},
			"fortios_switchcontrollersecuritypolicy_localaccess": {
				Tok: tfbridge.MakeResource(mainPkg, switchcontrollersecuritypolicyMod, "localaccess"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# switchcontrollersecuritypolicy - localaccess"),
				},
			},
			"fortios_system3gmodem_custom": {
				Tok: tfbridge.MakeResource(mainPkg, system3gmodemMod, "custom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system3gmodem - custom"),
				},
			},
			"fortios_system_accprofile": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "accprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - accprofile"),
				},
			},
			"fortios_system_acme": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "acme"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - acme"),
				},
			},
			"fortios_system_admin": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "admin"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - admin"),
				},
			},
			"fortios_system_admin_administrator": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "admin_administrator"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - admin_administrator"),
				},
			},
			"fortios_system_admin_profiles": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "admin_profiles"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - admin_profiles"),
				},
			},
			"fortios_system_affinityinterrupt": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "affinityinterrupt"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - affinityinterrupt"),
				},
			},
			"fortios_system_affinitypacketredistribution": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "affinitypacketredistribution"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - affinitypacketredistribution"),
				},
			},
			"fortios_system_alarm": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "alarm"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - alarm"),
				},
			},
			"fortios_system_alias": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "alias"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - alias"),
				},
			},
			"fortios_system_apiuser": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "apiuser"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - apiuser"),
				},
			},
			"fortios_system_apiuser_setting": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "apiuser_setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - apiuser_setting"),
				},
			},
			"fortios_system_arptable": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "arptable"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - arptable"),
				},
			},
			"fortios_system_autoinstall": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "autoinstall"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - autoinstall"),
				},
			},
			"fortios_system_automationaction": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "automationaction"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - automationaction"),
				},
			},
			"fortios_system_automationdestination": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "automationdestination"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - automationdestination"),
				},
			},
			"fortios_system_automationstitch": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "automationstitch"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - automationstitch"),
				},
			},
			"fortios_system_automationtrigger": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "automationtrigger"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - automationtrigger"),
				},
			},
			"fortios_system_autoscript": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "autoscript"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - autoscript"),
				},
			},
			"fortios_system_centralmanagement": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "centralmanagement"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - centralmanagement"),
				},
			},
			"fortios_system_clustersync": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "clustersync"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - clustersync"),
				},
			},
			"fortios_system_console": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "console"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - console"),
				},
			},
			"fortios_system_csf": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "csf"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - csf"),
				},
			},
			"fortios_system_customlanguage": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "customlanguage"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - customlanguage"),
				},
			},
			"fortios_system_ddns": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ddns"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ddns"),
				},
			},
			"fortios_system_dedicatedmgmt": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "dedicatedmgmt"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dedicatedmgmt"),
				},
			},
			"fortios_system_dns": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "dns"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dns"),
				},
			},
			"fortios_system_dns64": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "dns64"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dns64"),
				},
			},
			"fortios_system_dnsdatabase": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "dnsdatabase"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dnsdatabase"),
				},
			},
			"fortios_system_dnsserver": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "dnsserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dnsserver"),
				},
			},
			"fortios_system_dscpbasedpriority": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "dscpbasedpriority"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dscpbasedpriority"),
				},
			},
			"fortios_system_emailserver": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "emailserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - emailserver"),
				},
			},
			"fortios_system_externalresource": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "externalresource"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - externalresource"),
				},
			},
			"fortios_system_federatedupgrade": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "federatedupgrade"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - federatedupgrade"),
				},
			},
			"fortios_system_fipscc": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "fipscc"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fipscc"),
				},
			},
			"fortios_system_fm": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "fm"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fm"),
				},
			},
			"fortios_system_fortiai": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "fortiai"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fortiai"),
				},
			},
			"fortios_system_fortiguard": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "fortiguard"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fortiguard"),
				},
			},
			"fortios_system_fortimanager": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "fortimanager"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fortimanager"),
				},
			},
			"fortios_system_fortindr": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "fortindr"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fortindr"),
				},
			},
			"fortios_system_fortisandbox": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "fortisandbox"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fortisandbox"),
				},
			},
			"fortios_system_fssopolling": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "fssopolling"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fssopolling"),
				},
			},
			"fortios_system_ftmpush": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ftmpush"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ftmpush"),
				},
			},
			"fortios_system_geneve": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "geneve"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - geneve"),
				},
			},
			"fortios_system_geoipcountry": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "geoipcountry"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - geoipcountry"),
				},
			},
			"fortios_system_geoipoverride": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "geoipoverride"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - geoipoverride"),
				},
			},
			"fortios_system_global": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "global"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - global"),
				},
			},
			"fortios_system_gretunnel": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "gretunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - gretunnel"),
				},
			},
			"fortios_system_ha": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ha"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ha"),
				},
			},
			"fortios_system_hamonitor": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "hamonitor"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - hamonitor"),
				},
			},
			"fortios_system_ike": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ike"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ike"),
				},
			},
			"fortios_system_interface": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "interface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - interface"),
				},
			},
			"fortios_system_ipam": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ipam"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipam"),
				},
			},
			"fortios_system_ipiptunnel": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ipiptunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipiptunnel"),
				},
			},
			"fortios_system_ips": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ips"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ips"),
				},
			},
			"fortios_system_ipsecaggregate": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ipsecaggregate"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipsecaggregate"),
				},
			},
			"fortios_system_ipsurlfilterdns": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ipsurlfilterdns"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipsurlfilterdns"),
				},
			},
			"fortios_system_ipsurlfilterdns6": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ipsurlfilterdns6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipsurlfilterdns6"),
				},
			},
			"fortios_system_ipv6neighborcache": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ipv6neighborcache"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipv6neighborcache"),
				},
			},
			"fortios_system_ipv6tunnel": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ipv6tunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipv6tunnel"),
				},
			},
			"fortios_system_license_forticare": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "license_forticare"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - license_forticare"),
				},
			},
			"fortios_system_license_vdom": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "license_vdom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - license_vdom"),
				},
			},
			"fortios_system_license_vm": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "license_vm"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - license_vm"),
				},
			},
			"fortios_system_linkmonitor": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "linkmonitor"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - linkmonitor"),
				},
			},
			"fortios_system_ltemodem": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ltemodem"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ltemodem"),
				},
			},
			"fortios_system_macaddresstable": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "macaddresstable"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - macaddresstable"),
				},
			},
			"fortios_system_managementtunnel": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "managementtunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - managementtunnel"),
				},
			},
			"fortios_system_mobiletunnel": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "mobiletunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - mobiletunnel"),
				},
			},
			"fortios_system_modem": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "modem"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - modem"),
				},
			},
			"fortios_system_nat64": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "nat64"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - nat64"),
				},
			},
			"fortios_system_ndproxy": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ndproxy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ndproxy"),
				},
			},
			"fortios_system_netflow": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "netflow"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - netflow"),
				},
			},
			"fortios_system_networkvisibility": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "networkvisibility"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - networkvisibility"),
				},
			},
			"fortios_system_npu": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "npu"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - npu"),
				},
			},
			"fortios_system_ntp": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ntp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ntp"),
				},
			},
			"fortios_system_objecttagging": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "objecttagging"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - objecttagging"),
				},
			},
			"fortios_system_passwordpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "passwordpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - passwordpolicy"),
				},
			},
			"fortios_system_passwordpolicyguestadmin": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "passwordpolicyguestadmin"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - passwordpolicyguestadmin"),
				},
			},
			"fortios_system_physicalswitch": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "physicalswitch"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - physicalswitch"),
				},
			},
			"fortios_system_pppoeinterface": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "pppoeinterface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - pppoeinterface"),
				},
			},
			"fortios_system_proberesponse": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "proberesponse"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - proberesponse"),
				},
			},
			"fortios_system_proxyarp": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "proxyarp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - proxyarp"),
				},
			},
			"fortios_system_ptp": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ptp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ptp"),
				},
			},
			"fortios_system_replacemsggroup": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "replacemsggroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - replacemsggroup"),
				},
			},
			"fortios_system_replacemsgimage": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "replacemsgimage"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - replacemsgimage"),
				},
			},
			"fortios_system_resourcelimits": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "resourcelimits"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - resourcelimits"),
				},
			},
			"fortios_system_saml": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "saml"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - saml"),
				},
			},
			"fortios_system_sdnconnector": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "sdnconnector"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sdnconnector"),
				},
			},
			"fortios_system_sdwan": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "sdwan"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sdwan"),
				},
			},
			"fortios_system_sessionhelper": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "sessionhelper"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sessionhelper"),
				},
			},
			"fortios_system_sessionttl": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "sessionttl"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sessionttl"),
				},
			},
			"fortios_system_setting_dns": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "setting_dns"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - setting_dns"),
				},
			},
			"fortios_system_setting_global": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "setting_global"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - setting_global"),
				},
			},
			"fortios_system_setting_ntp": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "setting_ntp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - setting_ntp"),
				},
			},
			"fortios_system_settings": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "settings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - settings"),
				},
			},
			"fortios_system_sflow": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "sflow"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sflow"),
				},
			},
			"fortios_system_sittunnel": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "sittunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sittunnel"),
				},
			},
			"fortios_system_smsserver": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "smsserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - smsserver"),
				},
			},
			"fortios_system_speedtestschedule": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "speedtestschedule"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - speedtestschedule"),
				},
			},
			"fortios_system_speedtestserver": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "speedtestserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - speedtestserver"),
				},
			},
			"fortios_system_ssoadmin": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ssoadmin"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ssoadmin"),
				},
			},
			"fortios_system_ssoforticloudadmin": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "ssoforticloudadmin"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ssoforticloudadmin"),
				},
			},
			"fortios_system_standalonecluster": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "standalonecluster"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - standalonecluster"),
				},
			},
			"fortios_system_storage": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "storage"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - storage"),
				},
			},
			"fortios_system_stp": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "stp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - stp"),
				},
			},
			"fortios_system_switchinterface": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "switchinterface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - switchinterface"),
				},
			},
			"fortios_system_tosbasedpriority": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "tosbasedpriority"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - tosbasedpriority"),
				},
			},
			"fortios_system_vdom": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "vdom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdom"),
				},
			},
			"fortios_system_vdom_setting": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "vdom_setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdom_setting"),
				},
			},
			"fortios_system_vdomdns": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "vdomdns"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdomdns"),
				},
			},
			"fortios_system_vdomexception": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "vdomexception"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdomexception"),
				},
			},
			"fortios_system_vdomlink": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "vdomlink"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdomlink"),
				},
			},
			"fortios_system_vdomnetflow": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "vdomnetflow"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdomnetflow"),
				},
			},
			"fortios_system_vdomproperty": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "vdomproperty"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdomproperty"),
				},
			},
			"fortios_system_vdomradiusserver": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "vdomradiusserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdomradiusserver"),
				},
			},
			"fortios_system_vdomsflow": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "vdomsflow"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdomsflow"),
				},
			},
			"fortios_system_virtualswitch": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "virtualswitch"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - virtualswitch"),
				},
			},
			"fortios_system_virtualwanlink": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "virtualwanlink"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - virtualwanlink"),
				},
			},
			"fortios_system_virtualwirepair": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "virtualwirepair"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - virtualwirepair"),
				},
			},
			"fortios_system_vnetunnel": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "vnetunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vnetunnel"),
				},
			},
			"fortios_system_vxlan": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "vxlan"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vxlan"),
				},
			},
			"fortios_system_wccp": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "wccp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - wccp"),
				},
			},
			"fortios_system_zone": {
				Tok: tfbridge.MakeResource(mainPkg, systemMod, "zone"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - zone"),
				},
			},
			"fortios_systemautoupdate_pushupdate": {
				Tok: tfbridge.MakeResource(mainPkg, systemautoupdateMod, "pushupdate"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemautoupdate - pushupdate"),
				},
			},
			"fortios_systemautoupdate_schedule": {
				Tok: tfbridge.MakeResource(mainPkg, systemautoupdateMod, "schedule"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemautoupdate - schedule"),
				},
			},
			"fortios_systemautoupdate_tunneling": {
				Tok: tfbridge.MakeResource(mainPkg, systemautoupdateMod, "tunneling"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemautoupdate - tunneling"),
				},
			},
			"fortios_systemdhcp6_server": {
				Tok: tfbridge.MakeResource(mainPkg, systemdhcp6Mod, "server"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemdhcp6 - server"),
				},
			},
			"fortios_systemdhcp_server": {
				Tok: tfbridge.MakeResource(mainPkg, systemdhcpMod, "server"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemdhcp - server"),
				},
			},
			"fortios_systemlldp_networkpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, systemlldpMod, "networkpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemlldp - networkpolicy"),
				},
			},
			"fortios_systemreplacemsg_admin": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "admin"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - admin"),
				},
			},
			"fortios_systemreplacemsg_alertmail": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "alertmail"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - alertmail"),
				},
			},
			"fortios_systemreplacemsg_auth": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "auth"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - auth"),
				},
			},
			"fortios_systemreplacemsg_automation": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "automation"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - automation"),
				},
			},
			"fortios_systemreplacemsg_devicedetectionportal": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "devicedetectionportal"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - devicedetectionportal"),
				},
			},
			"fortios_systemreplacemsg_ec": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "ec"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - ec"),
				},
			},
			"fortios_systemreplacemsg_fortiguardwf": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "fortiguardwf"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - fortiguardwf"),
				},
			},
			"fortios_systemreplacemsg_ftp": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "ftp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - ftp"),
				},
			},
			"fortios_systemreplacemsg_http": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "http"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - http"),
				},
			},
			"fortios_systemreplacemsg_icap": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "icap"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - icap"),
				},
			},
			"fortios_systemreplacemsg_mail": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "mail"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - mail"),
				},
			},
			"fortios_systemreplacemsg_nacquar": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "nacquar"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - nacquar"),
				},
			},
			"fortios_systemreplacemsg_nntp": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "nntp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - nntp"),
				},
			},
			"fortios_systemreplacemsg_spam": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "spam"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - spam"),
				},
			},
			"fortios_systemreplacemsg_sslvpn": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "sslvpn"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - sslvpn"),
				},
			},
			"fortios_systemreplacemsg_trafficquota": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "trafficquota"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - trafficquota"),
				},
			},
			"fortios_systemreplacemsg_utm": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "utm"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - utm"),
				},
			},
			"fortios_systemreplacemsg_webproxy": {
				Tok: tfbridge.MakeResource(mainPkg, systemreplacemsgMod, "webproxy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemreplacemsg - webproxy"),
				},
			},
			"fortios_systemsnmp_community": {
				Tok: tfbridge.MakeResource(mainPkg, systemsnmpMod, "community"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemsnmp - community"),
				},
			},
			"fortios_systemsnmp_mibview": {
				Tok: tfbridge.MakeResource(mainPkg, systemsnmpMod, "mibview"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemsnmp - mibview"),
				},
			},
			"fortios_systemsnmp_sysinfo": {
				Tok: tfbridge.MakeResource(mainPkg, systemsnmpMod, "sysinfo"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemsnmp - sysinfo"),
				},
			},
			"fortios_systemsnmp_user": {
				Tok: tfbridge.MakeResource(mainPkg, systemsnmpMod, "user"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemsnmp - user"),
				},
			},
			"fortios_user_adgrp": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "adgrp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - adgrp"),
				},
			},
			"fortios_user_certificate": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "certificate"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - certificate"),
				},
			},
			"fortios_user_device": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "device"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - device"),
				},
			},
			"fortios_user_deviceaccesslist": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "deviceaccesslist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - deviceaccesslist"),
				},
			},
			"fortios_user_devicecategory": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "devicecategory"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - devicecategory"),
				},
			},
			"fortios_user_devicegroup": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "devicegroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - devicegroup"),
				},
			},
			"fortios_user_domaincontroller": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "domaincontroller"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - domaincontroller"),
				},
			},
			"fortios_user_exchange": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "exchange"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - exchange"),
				},
			},
			"fortios_user_fortitoken": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "fortitoken"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - fortitoken"),
				},
			},
			"fortios_user_fsso": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "fsso"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - fsso"),
				},
			},
			"fortios_user_fssopolling": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "fssopolling"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - fssopolling"),
				},
			},
			"fortios_user_group": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "group"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - group"),
				},
			},
			"fortios_user_krbkeytab": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "krbkeytab"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - krbkeytab"),
				},
			},
			"fortios_user_ldap": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "ldap"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - ldap"),
				},
			},
			"fortios_user_local": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "local"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - local"),
				},
			},
			"fortios_user_nacpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "nacpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - nacpolicy"),
				},
			},
			"fortios_user_passwordpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "passwordpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - passwordpolicy"),
				},
			},
			"fortios_user_peer": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "peer"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - peer"),
				},
			},
			"fortios_user_peergrp": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "peergrp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - peergrp"),
				},
			},
			"fortios_user_pop3": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "pop3"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - pop3"),
				},
			},
			"fortios_user_quarantine": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "quarantine"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - quarantine"),
				},
			},
			"fortios_user_radius": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "radius"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - radius"),
				},
			},
			"fortios_user_saml": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "saml"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - saml"),
				},
			},
			"fortios_user_securityexemptlist": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "securityexemptlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - securityexemptlist"),
				},
			},
			"fortios_user_setting": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - setting"),
				},
			},
			"fortios_user_tacacs": {
				Tok: tfbridge.MakeResource(mainPkg, userMod, "tacacs"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - tacacs"),
				},
			},
			"fortios_videofilter_profile": {
				Tok: tfbridge.MakeResource(mainPkg, videofilterMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# videofilter - profile"),
				},
			},
			"fortios_videofilter_youtubechannelfilter": {
				Tok: tfbridge.MakeResource(mainPkg, videofilterMod, "youtubechannelfilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# videofilter - youtubechannelfilter"),
				},
			},
			"fortios_videofilter_youtubekey": {
				Tok: tfbridge.MakeResource(mainPkg, videofilterMod, "youtubekey"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# videofilter - youtubekey"),
				},
			},
			"fortios_voip_profile": {
				Tok: tfbridge.MakeResource(mainPkg, voipMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# voip - profile"),
				},
			},
			"fortios_vpn_ipsec_phase1interface": {
				Tok: tfbridge.MakeResource(mainPkg, vpnMod, "ipsec_phase1interface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpn - ipsec_phase1interface"),
				},
			},
			"fortios_vpn_ipsec_phase2interface": {
				Tok: tfbridge.MakeResource(mainPkg, vpnMod, "ipsec_phase2interface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpn - ipsec_phase2interface"),
				},
			},
			"fortios_vpn_l2tp": {
				Tok: tfbridge.MakeResource(mainPkg, vpnMod, "l2tp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpn - l2tp"),
				},
			},
			"fortios_vpn_ocvpn": {
				Tok: tfbridge.MakeResource(mainPkg, vpnMod, "ocvpn"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpn - ocvpn"),
				},
			},
			"fortios_vpn_pptp": {
				Tok: tfbridge.MakeResource(mainPkg, vpnMod, "pptp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpn - pptp"),
				},
			},
			"fortios_vpncertificate_ca": {
				Tok: tfbridge.MakeResource(mainPkg, vpncertificateMod, "ca"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpncertificate - ca"),
				},
			},
			"fortios_vpncertificate_crl": {
				Tok: tfbridge.MakeResource(mainPkg, vpncertificateMod, "crl"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpncertificate - crl"),
				},
			},
			"fortios_vpncertificate_local": {
				Tok: tfbridge.MakeResource(mainPkg, vpncertificateMod, "local"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpncertificate - local"),
				},
			},
			"fortios_vpncertificate_ocspserver": {
				Tok: tfbridge.MakeResource(mainPkg, vpncertificateMod, "ocspserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpncertificate - ocspserver"),
				},
			},
			"fortios_vpncertificate_remote": {
				Tok: tfbridge.MakeResource(mainPkg, vpncertificateMod, "remote"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpncertificate - remote"),
				},
			},
			"fortios_vpncertificate_setting": {
				Tok: tfbridge.MakeResource(mainPkg, vpncertificateMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpncertificate - setting"),
				},
			},
			"fortios_vpnipsec_concentrator": {
				Tok: tfbridge.MakeResource(mainPkg, vpnipsecMod, "concentrator"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnipsec - concentrator"),
				},
			},
			"fortios_vpnipsec_fec": {
				Tok: tfbridge.MakeResource(mainPkg, vpnipsecMod, "fec"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnipsec - fec"),
				},
			},
			"fortios_vpnipsec_forticlient": {
				Tok: tfbridge.MakeResource(mainPkg, vpnipsecMod, "forticlient"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnipsec - forticlient"),
				},
			},
			"fortios_vpnipsec_manualkey": {
				Tok: tfbridge.MakeResource(mainPkg, vpnipsecMod, "manualkey"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnipsec - manualkey"),
				},
			},
			"fortios_vpnipsec_manualkeyinterface": {
				Tok: tfbridge.MakeResource(mainPkg, vpnipsecMod, "manualkeyinterface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnipsec - manualkeyinterface"),
				},
			},
			"fortios_vpnipsec_phase1": {
				Tok: tfbridge.MakeResource(mainPkg, vpnipsecMod, "phase1"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnipsec - phase1"),
				},
			},
			"fortios_vpnipsec_phase1interface": {
				Tok: tfbridge.MakeResource(mainPkg, vpnipsecMod, "phase1interface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnipsec - phase1interface"),
				},
			},
			"fortios_vpnipsec_phase2": {
				Tok: tfbridge.MakeResource(mainPkg, vpnipsecMod, "phase2"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnipsec - phase2"),
				},
			},
			"fortios_vpnipsec_phase2interface": {
				Tok: tfbridge.MakeResource(mainPkg, vpnipsecMod, "phase2interface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnipsec - phase2interface"),
				},
			},
			"fortios_vpnssl_client": {
				Tok: tfbridge.MakeResource(mainPkg, vpnsslMod, "client"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnssl - client"),
				},
			},
			"fortios_vpnssl_settings": {
				Tok: tfbridge.MakeResource(mainPkg, vpnsslMod, "settings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnssl - settings"),
				},
			},
			"fortios_vpnsslweb_hostchecksoftware": {
				Tok: tfbridge.MakeResource(mainPkg, vpnsslwebMod, "hostchecksoftware"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnsslweb - hostchecksoftware"),
				},
			},
			"fortios_vpnsslweb_portal": {
				Tok: tfbridge.MakeResource(mainPkg, vpnsslwebMod, "portal"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnsslweb - portal"),
				},
			},
			"fortios_vpnsslweb_realm": {
				Tok: tfbridge.MakeResource(mainPkg, vpnsslwebMod, "realm"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnsslweb - realm"),
				},
			},
			"fortios_vpnsslweb_userbookmark": {
				Tok: tfbridge.MakeResource(mainPkg, vpnsslwebMod, "userbookmark"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnsslweb - userbookmark"),
				},
			},
			"fortios_vpnsslweb_usergroupbookmark": {
				Tok: tfbridge.MakeResource(mainPkg, vpnsslwebMod, "usergroupbookmark"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# vpnsslweb - usergroupbookmark"),
				},
			},
			"fortios_waf_mainclass": {
				Tok: tfbridge.MakeResource(mainPkg, wafMod, "mainclass"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# waf - mainclass"),
				},
			},
			"fortios_waf_profile": {
				Tok: tfbridge.MakeResource(mainPkg, wafMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# waf - profile"),
				},
			},
			"fortios_waf_signature": {
				Tok: tfbridge.MakeResource(mainPkg, wafMod, "signature"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# waf - signature"),
				},
			},
			"fortios_waf_subclass": {
				Tok: tfbridge.MakeResource(mainPkg, wafMod, "subclass"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# waf - subclass"),
				},
			},
			"fortios_wanopt_authgroup": {
				Tok: tfbridge.MakeResource(mainPkg, wanoptMod, "authgroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wanopt - authgroup"),
				},
			},
			"fortios_wanopt_cacheservice": {
				Tok: tfbridge.MakeResource(mainPkg, wanoptMod, "cacheservice"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wanopt - cacheservice"),
				},
			},
			"fortios_wanopt_contentdeliverynetworkrule": {
				Tok: tfbridge.MakeResource(mainPkg, wanoptMod, "contentdeliverynetworkrule"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wanopt - contentdeliverynetworkrule"),
				},
			},
			"fortios_wanopt_peer": {
				Tok: tfbridge.MakeResource(mainPkg, wanoptMod, "peer"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wanopt - peer"),
				},
			},
			"fortios_wanopt_profile": {
				Tok: tfbridge.MakeResource(mainPkg, wanoptMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wanopt - profile"),
				},
			},
			"fortios_wanopt_remotestorage": {
				Tok: tfbridge.MakeResource(mainPkg, wanoptMod, "remotestorage"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wanopt - remotestorage"),
				},
			},
			"fortios_wanopt_settings": {
				Tok: tfbridge.MakeResource(mainPkg, wanoptMod, "settings"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wanopt - settings"),
				},
			},
			"fortios_wanopt_webcache": {
				Tok: tfbridge.MakeResource(mainPkg, wanoptMod, "webcache"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wanopt - webcache"),
				},
			},
			"fortios_webfilter_content": {
				Tok: tfbridge.MakeResource(mainPkg, webfilterMod, "content"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webfilter - content"),
				},
			},
			"fortios_webfilter_contentheader": {
				Tok: tfbridge.MakeResource(mainPkg, webfilterMod, "contentheader"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webfilter - contentheader"),
				},
			},
			"fortios_webfilter_fortiguard": {
				Tok: tfbridge.MakeResource(mainPkg, webfilterMod, "fortiguard"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webfilter - fortiguard"),
				},
			},
			"fortios_webfilter_ftgdlocalcat": {
				Tok: tfbridge.MakeResource(mainPkg, webfilterMod, "ftgdlocalcat"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webfilter - ftgdlocalcat"),
				},
			},
			"fortios_webfilter_ftgdlocalrating": {
				Tok: tfbridge.MakeResource(mainPkg, webfilterMod, "ftgdlocalrating"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webfilter - ftgdlocalrating"),
				},
			},
			"fortios_webfilter_ipsurlfiltercachesetting": {
				Tok: tfbridge.MakeResource(mainPkg, webfilterMod, "ipsurlfiltercachesetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webfilter - ipsurlfiltercachesetting"),
				},
			},
			"fortios_webfilter_ipsurlfiltersetting": {
				Tok: tfbridge.MakeResource(mainPkg, webfilterMod, "ipsurlfiltersetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webfilter - ipsurlfiltersetting"),
				},
			},
			"fortios_webfilter_ipsurlfiltersetting6": {
				Tok: tfbridge.MakeResource(mainPkg, webfilterMod, "ipsurlfiltersetting6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webfilter - ipsurlfiltersetting6"),
				},
			},
			"fortios_webfilter_override": {
				Tok: tfbridge.MakeResource(mainPkg, webfilterMod, "override"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webfilter - override"),
				},
			},
			"fortios_webfilter_profile": {
				Tok: tfbridge.MakeResource(mainPkg, webfilterMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webfilter - profile"),
				},
			},
			"fortios_webfilter_searchengine": {
				Tok: tfbridge.MakeResource(mainPkg, webfilterMod, "searchengine"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webfilter - searchengine"),
				},
			},
			"fortios_webfilter_urlfilter": {
				Tok: tfbridge.MakeResource(mainPkg, webfilterMod, "urlfilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webfilter - urlfilter"),
				},
			},
			"fortios_webproxy_debugurl": {
				Tok: tfbridge.MakeResource(mainPkg, webproxyMod, "debugurl"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webproxy - debugurl"),
				},
			},
			"fortios_webproxy_explicit": {
				Tok: tfbridge.MakeResource(mainPkg, webproxyMod, "explicit"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webproxy - explicit"),
				},
			},
			"fortios_webproxy_forwardserver": {
				Tok: tfbridge.MakeResource(mainPkg, webproxyMod, "forwardserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webproxy - forwardserver"),
				},
			},
			"fortios_webproxy_forwardservergroup": {
				Tok: tfbridge.MakeResource(mainPkg, webproxyMod, "forwardservergroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webproxy - forwardservergroup"),
				},
			},
			"fortios_webproxy_global": {
				Tok: tfbridge.MakeResource(mainPkg, webproxyMod, "global"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webproxy - global"),
				},
			},
			"fortios_webproxy_profile": {
				Tok: tfbridge.MakeResource(mainPkg, webproxyMod, "profile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webproxy - profile"),
				},
			},
			"fortios_webproxy_urlmatch": {
				Tok: tfbridge.MakeResource(mainPkg, webproxyMod, "urlmatch"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webproxy - urlmatch"),
				},
			},
			"fortios_webproxy_wisp": {
				Tok: tfbridge.MakeResource(mainPkg, webproxyMod, "wisp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# webproxy - wisp"),
				},
			},
			"fortios_wirelesscontroller_accesscontrollist": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "accesscontrollist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - accesscontrollist"),
				},
			},
			"fortios_wirelesscontroller_address": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "address"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - address"),
				},
			},
			"fortios_wirelesscontroller_addrgrp": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "addrgrp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - addrgrp"),
				},
			},
			"fortios_wirelesscontroller_apcfgprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "apcfgprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - apcfgprofile"),
				},
			},
			"fortios_wirelesscontroller_apstatus": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "apstatus"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - apstatus"),
				},
			},
			"fortios_wirelesscontroller_arrpprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "arrpprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - arrpprofile"),
				},
			},
			"fortios_wirelesscontroller_bleprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "bleprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - bleprofile"),
				},
			},
			"fortios_wirelesscontroller_bonjourprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "bonjourprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - bonjourprofile"),
				},
			},
			"fortios_wirelesscontroller_global": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "global"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - global"),
				},
			},
			"fortios_wirelesscontroller_intercontroller": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "intercontroller"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - intercontroller"),
				},
			},
			"fortios_wirelesscontroller_log": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "log"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - log"),
				},
			},
			"fortios_wirelesscontroller_mpskprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "mpskprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - mpskprofile"),
				},
			},
			"fortios_wirelesscontroller_nacprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "nacprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - nacprofile"),
				},
			},
			"fortios_wirelesscontroller_qosprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "qosprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - qosprofile"),
				},
			},
			"fortios_wirelesscontroller_region": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "region"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - region"),
				},
			},
			"fortios_wirelesscontroller_setting": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "setting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - setting"),
				},
			},
			"fortios_wirelesscontroller_snmp": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "snmp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - snmp"),
				},
			},
			"fortios_wirelesscontroller_ssidpolicy": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "ssidpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - ssidpolicy"),
				},
			},
			"fortios_wirelesscontroller_syslogprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "syslogprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - syslogprofile"),
				},
			},
			"fortios_wirelesscontroller_timers": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "timers"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - timers"),
				},
			},
			"fortios_wirelesscontroller_utmprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "utmprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - utmprofile"),
				},
			},
			"fortios_wirelesscontroller_vap": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "vap"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - vap"),
				},
			},
			"fortios_wirelesscontroller_vapgroup": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "vapgroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - vapgroup"),
				},
			},
			"fortios_wirelesscontroller_wagprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "wagprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - wagprofile"),
				},
			},
			"fortios_wirelesscontroller_widsprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "widsprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - widsprofile"),
				},
			},
			"fortios_wirelesscontroller_wtp": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "wtp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - wtp"),
				},
			},
			"fortios_wirelesscontroller_wtpgroup": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "wtpgroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - wtpgroup"),
				},
			},
			"fortios_wirelesscontroller_wtpprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerMod, "wtpprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontroller - wtpprofile"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_anqp3gppcellular": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "anqp3gppcellular"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - anqp3gppcellular"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_anqpipaddresstype": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "anqpipaddresstype"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - anqpipaddresstype"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_anqpnairealm": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "anqpnairealm"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - anqpnairealm"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "anqpnetworkauthtype"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - anqpnetworkauthtype"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_anqproamingconsortium": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "anqproamingconsortium"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - anqproamingconsortium"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_anqpvenuename": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "anqpvenuename"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - anqpvenuename"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_anqpvenueurl": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "anqpvenueurl"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - anqpvenueurl"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "h2qpadviceofcharge"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - h2qpadviceofcharge"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_h2qpconncapability": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "h2qpconncapability"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - h2qpconncapability"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_h2qpoperatorname": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "h2qpoperatorname"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - h2qpoperatorname"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_h2qposuprovider": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "h2qposuprovider"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - h2qposuprovider"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_h2qposuprovidernai": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "h2qposuprovidernai"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - h2qposuprovidernai"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_h2qptermsandconditions": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "h2qptermsandconditions"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - h2qptermsandconditions"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_h2qpwanmetric": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "h2qpwanmetric"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - h2qpwanmetric"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_hsprofile": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "hsprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - hsprofile"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_icon": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "icon"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - icon"),
				},
			},
			"fortios_wirelesscontrollerhotspot20_qosmap": {
				Tok: tfbridge.MakeResource(mainPkg, wirelesscontrollerhotspot20Mod, "qosmap"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# wirelesscontrollerhotspot20 - qosmap"),
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getgetAmi")},
			"fortios_firewall_DoSpolicy": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFirewallDoSpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - DoSpolicy"),
				},
			},
			"fortios_firewall_DoSpolicy6": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getDoSpolicy6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - DoSpolicy6"),
				},
			},
			"fortios_firewall_DoSpolicy6list": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getDoSpolicy6list"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - DoSpolicy6list"),
				},
			},
			"fortios_firewall_DoSpolicylist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getDoSpolicylist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - DoSpolicylist"),
				},
			},
			"fortios_firewall_address": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getaddress"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - address"),
				},
				// Fields: map[string]*tfbridge.SchemaInfo{
				// 	addressFssoGroup: {
			},
			"fortios_firewall_address6": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getaddress6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - address6"),
				},
			},
			"fortios_firewall_address6list": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getaddress6list"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - address6list"),
				},
			},
			"fortios_firewall_address6template": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getaddress6template"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - address6template"),
				},
			},
			"fortios_firewall_address6templatelist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getaddress6templatelist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - address6templatelist"),
				},
			},
			"fortios_firewall_addresslist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getaddresslist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - addresslist"),
				},
			},
			"fortios_firewall_addrgrp": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getaddrgrp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - addrgrp"),
				},
			},
			"fortios_firewall_addrgrp6": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getaddrgrp6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - addrgrp6"),
				},
			},
			"fortios_firewall_addrgrp6list": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getaddrgrp6list"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - addrgrp6list"),
				},
			},
			"fortios_firewall_addrgrplist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getaddrgrplist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - addrgrplist"),
				},
			},
			"fortios_firewall_centralsnatmap": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getcentralsnatmap"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - centralsnatmap"),
				},
			},
			"fortios_firewall_centralsnatmaplist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getcentralsnatmaplist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - centralsnatmaplist"),
				},
			},
			"fortios_firewall_internetservice": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getinternetservice"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservice"),
				},
			},
			"fortios_firewall_internetservicecustom": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getinternetservicecustom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicecustom"),
				},
			},
			"fortios_firewall_internetservicecustomgroup": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getinternetservicecustomgroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicecustomgroup"),
				},
			},
			"fortios_firewall_internetservicecustomgrouplist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getinternetservicecustomgrouplist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicecustomgrouplist"),
				},
			},
			"fortios_firewall_internetservicecustomlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getinternetservicecustomlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicecustomlist"),
				},
			},
			"fortios_firewall_internetservicedefinition": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getinternetservicedefinition"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicedefinition"),
				},
			},
			"fortios_firewall_internetservicedefinitionlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getinternetservicedefinitionlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicedefinitionlist"),
				},
			},
			"fortios_firewall_internetserviceextension": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getinternetserviceextension"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetserviceextension"),
				},
			},
			"fortios_firewall_internetserviceextensionlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getinternetserviceextensionlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetserviceextensionlist"),
				},
			},
			"fortios_firewall_internetservicegroup": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getinternetservicegroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicegroup"),
				},
			},
			"fortios_firewall_internetservicegrouplist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getinternetservicegrouplist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicegrouplist"),
				},
			},
			"fortios_firewall_internetservicelist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getinternetservicelist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - internetservicelist"),
				},
			},
			"fortios_firewall_ipv6ehfilter": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getipv6ehfilter"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - ipv6ehfilter"),
				},
			},
			"fortios_firewall_multicastaddress": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getmulticastaddress"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - multicastaddress"),
				},
			},
			"fortios_firewall_multicastaddress6": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getmulticastaddress6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - multicastaddress6"),
				},
			},
			"fortios_firewall_multicastaddress6list": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getmulticastaddress6list"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - multicastaddress6list"),
				},
			},
			"fortios_firewall_multicastaddresslist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getmulticastaddresslist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - multicastaddresslist"),
				},
			},
			"fortios_firewall_policy": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - policy"),
				},
			},
			"fortios_firewall_policy46": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getpolicy46"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - policy46"),
				},
			},
			"fortios_firewall_policy46list": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getpolicy46list"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - policy46list"),
				},
			},
			"fortios_firewall_policy6": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getpolicy6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - policy6"),
				},
			},
			"fortios_firewall_policy64": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getpolicy64"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - policy64"),
				},
			},
			"fortios_firewall_policy64list": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getpolicy64list"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - policy64list"),
				},
			},
			"fortios_firewall_policy6list": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getpolicy6list"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - policy6list"),
				},
			},
			"fortios_firewall_policylist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getpolicylist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - policylist"),
				},
			},
			"fortios_firewall_profileprotocoloptions": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getprofileprotocoloptions"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - profileprotocoloptions"),
				},
			},
			"fortios_firewall_profileprotocoloptionslist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getprofileprotocoloptionslist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - profileprotocoloptionslist"),
				},
			},
			"fortios_firewall_proxyaddress": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getproxyaddress"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - proxyaddress"),
				},
			},
			"fortios_firewall_proxyaddresslist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getproxyaddresslist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - proxyaddresslist"),
				},
			},
			"fortios_firewall_proxyaddrgrp": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getproxyaddrgrp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - proxyaddrgrp"),
				},
			},
			"fortios_firewall_proxyaddrgrplist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getproxyaddrgrplist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - proxyaddrgrplist"),
				},
			},
			"fortios_firewall_proxypolicy": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getproxypolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - proxypolicy"),
				},
			},
			"fortios_firewall_proxypolicylist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallMod, "getproxypolicylist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewall - proxypolicylist"),
				},
			},
			"fortios_firewallconsolidated_policy": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallconsolidatedMod, "getpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallconsolidated - policy"),
				},
			},
			"fortios_firewallconsolidated_policylist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallconsolidatedMod, "getpolicylist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallconsolidated - policylist"),
				},
			},
			"fortios_firewallschedule_group": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallscheduleMod, "getgroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallschedule - group"),
				},
			},
			"fortios_firewallschedule_grouplist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallscheduleMod, "getgrouplist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallschedule - grouplist"),
				},
			},
			"fortios_firewallschedule_onetime": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallscheduleMod, "getonetime"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallschedule - onetime"),
				},
			},
			"fortios_firewallschedule_onetimelist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallscheduleMod, "getonetimelist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallschedule - onetimelist"),
				},
			},
			"fortios_firewallschedule_recurring": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallscheduleMod, "getrecurring"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallschedule - recurring"),
				},
			},
			"fortios_firewallschedule_recurringlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallscheduleMod, "getrecurringlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallschedule - recurringlist"),
				},
			},
			"fortios_firewallservice_category": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallserviceMod, "getcategory"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallservice - category"),
				},
			},
			"fortios_firewallservice_categorylist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallserviceMod, "getcategorylist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallservice - categorylist"),
				},
			},
			"fortios_firewallservice_custom": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallserviceMod, "getcustom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallservice - custom"),
				},
			},
			"fortios_firewallservice_customlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallserviceMod, "getcustomlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallservice - customlist"),
				},
			},
			"fortios_firewallservice_group": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallserviceMod, "getgroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallservice - group"),
				},
			},
			"fortios_firewallservice_grouplist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallserviceMod, "getgrouplist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallservice - grouplist"),
				},
			},
			"fortios_firewallshaper_peripshaper": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallshaperMod, "getperipshaper"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallshaper - peripshaper"),
				},
			},
			"fortios_firewallshaper_peripshaperlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallshaperMod, "getperipshaperlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallshaper - peripshaperlist"),
				},
			},
			"fortios_firewallshaper_trafficshaper": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallshaperMod, "gettrafficshaper"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallshaper - trafficshaper"),
				},
			},
			"fortios_firewallshaper_trafficshaperlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallshaperMod, "gettrafficshaperlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallshaper - trafficshaperlist"),
				},
			},
			"fortios_firewallwildcardfqdn_custom": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallwildcardfqdnMod, "getcustom"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallwildcardfqdn - custom"),
				},
			},
			"fortios_firewallwildcardfqdn_customlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallwildcardfqdnMod, "getcustomlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallwildcardfqdn - customlist"),
				},
			},
			"fortios_firewallwildcardfqdn_group": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallwildcardfqdnMod, "getgroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallwildcardfqdn - group"),
				},
			},
			"fortios_firewallwildcardfqdn_grouplist": {
				Tok: tfbridge.MakeDataSource(mainPkg, firewallwildcardfqdnMod, "getgrouplist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# firewallwildcardfqdn - grouplist"),
				},
			},
			"fortios_ipmask_cidr": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipmaskMod, "getcidr"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# ipmask - cidr"),
				},
			},
			"fortios_json_generic_api": {
				Tok: tfbridge.MakeDataSource(mainPkg, jsonMod, "getgeneric_api"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# json - generic_api"),
				},
			},
			"fortios_router_accesslist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getaccesslist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - accesslist"),
				},
			},
			"fortios_router_accesslist6": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getaccesslist6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - accesslist6"),
				},
			},
			"fortios_router_accesslist6list": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getaccesslist6list"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - accesslist6list"),
				},
			},
			"fortios_router_accesslistlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getaccesslistlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - accesslistlist"),
				},
			},
			"fortios_router_aspathlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getaspathlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - aspathlist"),
				},
			},
			"fortios_router_aspathlistlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getaspathlistlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - aspathlistlist"),
				},
			},
			"fortios_router_authpath": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getauthpath"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - authpath"),
				},
			},
			"fortios_router_authpathlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getauthpathlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - authpathlist"),
				},
			},
			"fortios_router_bfd": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getbfd"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - bfd"),
				},
			},
			"fortios_router_bfd6": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getbfd6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - bfd6"),
				},
			},
			"fortios_router_bgp": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getbgp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - bgp"),
				},
			},
			"fortios_router_communitylist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getcommunitylist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - communitylist"),
				},
			},
			"fortios_router_communitylistlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getcommunitylistlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - communitylistlist"),
				},
			},
			"fortios_router_isis": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getisis"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - isis"),
				},
			},
			"fortios_router_keychain": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getkeychain"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - keychain"),
				},
			},
			"fortios_router_keychainlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getkeychainlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - keychainlist"),
				},
			},
			"fortios_router_multicast": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getmulticast"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - multicast"),
				},
			},
			"fortios_router_multicast6": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getmulticast6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - multicast6"),
				},
			},
			"fortios_router_multicastflow": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getmulticastflow"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - multicastflow"),
				},
			},
			"fortios_router_multicastflowlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getmulticastflowlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - multicastflowlist"),
				},
			},
			"fortios_router_ospf": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getospf"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - ospf"),
				},
			},
			"fortios_router_ospf6": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getospf6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - ospf6"),
				},
			},
			"fortios_router_policy": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - policy"),
				},
			},
			"fortios_router_policy6": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getpolicy6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - policy6"),
				},
			},
			"fortios_router_policy6list": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getpolicy6list"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - policy6list"),
				},
			},
			"fortios_router_policylist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getpolicylist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - policylist"),
				},
			},
			"fortios_router_prefixlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getprefixlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - prefixlist"),
				},
			},
			"fortios_router_prefixlist6": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getprefixlist6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - prefixlist6"),
				},
			},
			"fortios_router_prefixlist6list": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getprefixlist6list"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - prefixlist6list"),
				},
			},
			"fortios_router_prefixlistlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getprefixlistlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - prefixlistlist"),
				},
			},
			"fortios_router_rip": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getrip"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - rip"),
				},
			},
			"fortios_router_ripng": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getripng"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - ripng"),
				},
			},
			"fortios_router_routemap": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getroutemap"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - routemap"),
				},
			},
			"fortios_router_routemaplist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getroutemaplist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - routemaplist"),
				},
			},
			"fortios_router_setting": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getsetting"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - setting"),
				},
			},
			"fortios_router_static": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getstatic"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - static"),
				},
			},
			"fortios_router_static6": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getstatic6"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - static6"),
				},
			},
			"fortios_router_static6list": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getstatic6list"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - static6list"),
				},
			},
			"fortios_router_staticlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerMod, "getstaticlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# router - staticlist"),
				},
			},
			"fortios_routerbgp_neighbor": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerbgpMod, "getneighbor"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# routerbgp - neighbor"),
				},
			},
			"fortios_routerbgp_neighborlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, routerbgpMod, "getneighborlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# routerbgp - neighborlist"),
				},
			},
			"fortios_system_accprofile": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getaccprofile"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - accprofile"),
				},
			},
			"fortios_system_accprofilelist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getaccprofilelist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - accprofilelist"),
				},
			},
			"fortios_system_admin": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getadmin"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - admin"),
				},
			},
			"fortios_system_adminlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getadminlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - adminlist"),
				},
			},
			"fortios_system_alias": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getalias"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - alias"),
				},
			},
			"fortios_system_aliaslist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getaliaslist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - aliaslist"),
				},
			},
			"fortios_system_apiuser": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getapiuser"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - apiuser"),
				},
			},
			"fortios_system_apiuserlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getapiuserlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - apiuserlist"),
				},
			},
			"fortios_system_arptable": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getarptable"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - arptable"),
				},
			},
			"fortios_system_arptablelist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getarptablelist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - arptablelist"),
				},
			},
			"fortios_system_autoinstall": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getautoinstall"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - autoinstall"),
				},
			},
			"fortios_system_automationaction": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getautomationaction"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - automationaction"),
				},
			},
			"fortios_system_automationactionlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getautomationactionlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - automationactionlist"),
				},
			},
			"fortios_system_automationdestination": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getautomationdestination"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - automationdestination"),
				},
			},
			"fortios_system_automationdestinationlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getautomationdestinationlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - automationdestinationlist"),
				},
			},
			"fortios_system_automationtrigger": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getautomationtrigger"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - automationtrigger"),
				},
			},
			"fortios_system_automationtriggerlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getautomationtriggerlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - automationtriggerlist"),
				},
			},
			"fortios_system_autoscript": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getautoscript"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - autoscript"),
				},
			},
			"fortios_system_autoscriptlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getautoscriptlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - autoscriptlist"),
				},
			},
			"fortios_system_centralmanagement": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getcentralmanagement"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - centralmanagement"),
				},
			},
			"fortios_system_clustersync": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getclustersync"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - clustersync"),
				},
			},
			"fortios_system_clustersynclist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getclustersynclist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - clustersynclist"),
				},
			},
			"fortios_system_console": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getconsole"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - console"),
				},
			},
			"fortios_system_csf": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getcsf"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - csf"),
				},
			},
			"fortios_system_ddns": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getddns"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ddns"),
				},
			},
			"fortios_system_ddnslist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getddnslist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ddnslist"),
				},
			},
			"fortios_system_dns": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getdns"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dns"),
				},
			},
			"fortios_system_dnsdatabase": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getdnsdatabase"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dnsdatabase"),
				},
			},
			"fortios_system_dnsdatabaselist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getdnsdatabaselist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dnsdatabaselist"),
				},
			},
			"fortios_system_dnsserver": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getdnsserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dnsserver"),
				},
			},
			"fortios_system_dnsserverlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getdnsserverlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dnsserverlist"),
				},
			},
			"fortios_system_dscpbasedpriority": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getdscpbasedpriority"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dscpbasedpriority"),
				},
			},
			"fortios_system_dscpbasedprioritylist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getdscpbasedprioritylist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - dscpbasedprioritylist"),
				},
			},
			"fortios_system_emailserver": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getemailserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - emailserver"),
				},
			},
			"fortios_system_externalresource": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getexternalresource"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - externalresource"),
				},
			},
			"fortios_system_externalresourcelist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getexternalresourcelist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - externalresourcelist"),
				},
			},
			"fortios_system_fipscc": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getfipscc"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fipscc"),
				},
			},
			"fortios_system_fm": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getfm"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fm"),
				},
			},
			"fortios_system_fortiguard": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getfortiguard"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fortiguard"),
				},
			},
			"fortios_system_fortimanager": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getfortimanager"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fortimanager"),
				},
			},
			"fortios_system_fortisandbox": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getfortisandbox"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fortisandbox"),
				},
			},
			"fortios_system_fssopolling": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getfssopolling"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - fssopolling"),
				},
			},
			"fortios_system_ftmpush": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getftmpush"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ftmpush"),
				},
			},
			"fortios_system_global": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getglobal"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - global"),
				},
			},
			"fortios_system_gretunnel": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getgretunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - gretunnel"),
				},
			},
			"fortios_system_gretunnellist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getgretunnellist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - gretunnellist"),
				},
			},
			"fortios_system_ha": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getha"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ha"),
				},
			},
			"fortios_system_hamonitor": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "gethamonitor"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - hamonitor"),
				},
			},
			"fortios_system_interface": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getinterface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - interface"),
				},
			},
			"fortios_system_interfacelist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getinterfacelist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - interfacelist"),
				},
			},
			"fortios_system_ipiptunnel": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getipiptunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipiptunnel"),
				},
			},
			"fortios_system_ipiptunnellist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getipiptunnellist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipiptunnellist"),
				},
			},
			"fortios_system_ipv6neighborcache": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getipv6neighborcache"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipv6neighborcache"),
				},
			},
			"fortios_system_ipv6neighborcachelist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getipv6neighborcachelist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipv6neighborcachelist"),
				},
			},
			"fortios_system_ipv6tunnel": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getipv6tunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipv6tunnel"),
				},
			},
			"fortios_system_ipv6tunnellist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getipv6tunnellist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ipv6tunnellist"),
				},
			},
			"fortios_system_linkmonitor": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getlinkmonitor"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - linkmonitor"),
				},
			},
			"fortios_system_linkmonitorlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getlinkmonitorlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - linkmonitorlist"),
				},
			},
			"fortios_system_managementtunnel": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getmanagementtunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - managementtunnel"),
				},
			},
			"fortios_system_mobiletunnel": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getmobiletunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - mobiletunnel"),
				},
			},
			"fortios_system_mobiletunnellist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getmobiletunnellist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - mobiletunnellist"),
				},
			},
			"fortios_system_nat64": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getnat64"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - nat64"),
				},
			},
			"fortios_system_ndproxy": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getndproxy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ndproxy"),
				},
			},
			"fortios_system_netflow": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getnetflow"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - netflow"),
				},
			},
			"fortios_system_networkvisibility": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getnetworkvisibility"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - networkvisibility"),
				},
			},
			"fortios_system_ntp": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getntp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - ntp"),
				},
			},
			"fortios_system_objecttagging": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getobjecttagging"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - objecttagging"),
				},
			},
			"fortios_system_objecttagginglist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getobjecttagginglist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - objecttagginglist"),
				},
			},
			"fortios_system_passwordpolicy": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getpasswordpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - passwordpolicy"),
				},
			},
			"fortios_system_passwordpolicyguestadmin": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getpasswordpolicyguestadmin"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - passwordpolicyguestadmin"),
				},
			},
			"fortios_system_pppoeinterface": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getpppoeinterface"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - pppoeinterface"),
				},
			},
			"fortios_system_pppoeinterfacelist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getpppoeinterfacelist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - pppoeinterfacelist"),
				},
			},
			"fortios_system_proberesponse": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getproberesponse"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - proberesponse"),
				},
			},
			"fortios_system_proxyarp": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getproxyarp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - proxyarp"),
				},
			},
			"fortios_system_proxyarplist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getproxyarplist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - proxyarplist"),
				},
			},
			"fortios_system_replacemsggroup": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getreplacemsggroup"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - replacemsggroup"),
				},
			},
			"fortios_system_replacemsggrouplist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getreplacemsggrouplist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - replacemsggrouplist"),
				},
			},
			"fortios_system_replacemsgimage": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getreplacemsgimage"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - replacemsgimage"),
				},
			},
			"fortios_system_replacemsgimagelist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getreplacemsgimagelist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - replacemsgimagelist"),
				},
			},
			"fortios_system_resourcelimits": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getresourcelimits"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - resourcelimits"),
				},
			},
			"fortios_system_sdnconnector": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getsdnconnector"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sdnconnector"),
				},
			},
			"fortios_system_sdnconnectorlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getsdnconnectorlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sdnconnectorlist"),
				},
			},
			"fortios_system_sessionhelper": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getsessionhelper"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sessionhelper"),
				},
			},
			"fortios_system_sessionhelperlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getsessionhelperlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sessionhelperlist"),
				},
			},
			"fortios_system_sessionttl": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getsessionttl"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sessionttl"),
				},
			},
			"fortios_system_sflow": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getsflow"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sflow"),
				},
			},
			"fortios_system_sittunnel": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getsittunnel"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sittunnel"),
				},
			},
			"fortios_system_sittunnellist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getsittunnellist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - sittunnellist"),
				},
			},
			"fortios_system_smsserver": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getsmsserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - smsserver"),
				},
			},
			"fortios_system_smsserverlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getsmsserverlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - smsserverlist"),
				},
			},
			"fortios_system_tosbasedpriority": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "gettosbasedpriority"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - tosbasedpriority"),
				},
			},
			"fortios_system_tosbasedprioritylist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "gettosbasedprioritylist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - tosbasedprioritylist"),
				},
			},
			"fortios_system_vdomexception": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getvdomexception"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdomexception"),
				},
			},
			"fortios_system_vdomexceptionlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getvdomexceptionlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdomexceptionlist"),
				},
			},
			"fortios_system_vdomnetflow": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getvdomnetflow"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdomnetflow"),
				},
			},
			"fortios_system_vdomsflow": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getvdomsflow"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vdomsflow"),
				},
			},
			"fortios_system_virtualwanlink": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getvirtualwanlink"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - virtualwanlink"),
				},
			},
			"fortios_system_vxlan": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getvxlan"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vxlan"),
				},
			},
			"fortios_system_vxlanlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getvxlanlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - vxlanlist"),
				},
			},
			"fortios_system_wccp": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getwccp"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - wccp"),
				},
			},
			"fortios_system_wccplist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getwccplist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - wccplist"),
				},
			},
			"fortios_system_zone": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getzone"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - zone"),
				},
			},
			"fortios_system_zonelist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemMod, "getzonelist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# system - zonelist"),
				},
			},
			"fortios_systemautoupdate_pushupdate": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemautoupdateMod, "getpushupdate"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemautoupdate - pushupdate"),
				},
			},
			"fortios_systemautoupdate_schedule": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemautoupdateMod, "getschedule"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemautoupdate - schedule"),
				},
			},
			"fortios_systemautoupdate_tunneling": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemautoupdateMod, "gettunneling"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemautoupdate - tunneling"),
				},
			},
			"fortios_systemdhcp_server": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemdhcpMod, "getserver"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemdhcp - server"),
				},
			},
			"fortios_systemdhcp_serverlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemdhcpMod, "getserverlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemdhcp - serverlist"),
				},
			},
			"fortios_systemlldp_networkpolicy": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemlldpMod, "getnetworkpolicy"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemlldp - networkpolicy"),
				},
			},
			"fortios_systemlldp_networkpolicylist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemlldpMod, "getnetworkpolicylist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemlldp - networkpolicylist"),
				},
			},
			"fortios_systemsnmp_community": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemsnmpMod, "getcommunity"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemsnmp - community"),
				},
			},
			"fortios_systemsnmp_communitylist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemsnmpMod, "getcommunitylist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemsnmp - communitylist"),
				},
			},
			"fortios_systemsnmp_sysinfo": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemsnmpMod, "getsysinfo"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemsnmp - sysinfo"),
				},
			},
			"fortios_systemsnmp_user": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemsnmpMod, "getuser"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemsnmp - user"),
				},
			},
			"fortios_systemsnmp_userlist": {
				Tok: tfbridge.MakeDataSource(mainPkg, systemsnmpMod, "getuserlist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# systemsnmp - userlist"),
				},
			},
			"fortios_user_saml": {
				Tok: tfbridge.MakeDataSource(mainPkg, userMod, "getsaml"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - saml"),
				},
			},
			"fortios_user_samllist": {
				Tok: tfbridge.MakeDataSource(mainPkg, userMod, "getsamllist"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte("# user - samllist"),
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@aspyrmedia/pulumi-fortios",
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
			PackageName: "aspyrmedia_fortios",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/aspyrmedia/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Aspyrmedia",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
