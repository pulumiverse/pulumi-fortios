// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


// Export sub-modules:
import * as alertemail from "./alertemail";
import * as antivirus from "./antivirus";
import * as application from "./application";
import * as authentication from "./authentication";
import * as automation from "./automation";
import * as certificate from "./certificate";
import * as cifs from "./cifs";
import * as config from "./config";
import * as credentialstore from "./credentialstore";
import * as dlp from "./dlp";
import * as dnsfilter from "./dnsfilter";
import * as dpdk from "./dpdk";
import * as emailfilter from "./emailfilter";
import * as endpointcontrol from "./endpointcontrol";
import * as extendercontroller from "./extendercontroller";
import * as extensioncontroller from "./extensioncontroller";
import * as filefilter from "./filefilter";
import * as firewall from "./firewall";
import * as firewallconsolidated from "./firewallconsolidated";
import * as firewallipmacbinding from "./firewallipmacbinding";
import * as firewallschedule from "./firewallschedule";
import * as firewallservice from "./firewallservice";
import * as firewallshaper from "./firewallshaper";
import * as firewallssh from "./firewallssh";
import * as firewallssl from "./firewallssl";
import * as firewallwildcardfqdn from "./firewallwildcardfqdn";
import * as fmg from "./fmg";
import * as ftpproxy from "./ftpproxy";
import * as icap from "./icap";
import * as ipmask from "./ipmask";
import * as ips from "./ips";
import * as json from "./json";
import * as log from "./log";
import * as logdisk from "./logdisk";
import * as logfortianalyzer from "./logfortianalyzer";
import * as logfortianalyzer2 from "./logfortianalyzer2";
import * as logfortianalyzer3 from "./logfortianalyzer3";
import * as logfortianalyzercloud from "./logfortianalyzercloud";
import * as logfortiguard from "./logfortiguard";
import * as logmemory from "./logmemory";
import * as lognulldevice from "./lognulldevice";
import * as logsyslogd from "./logsyslogd";
import * as logsyslogd2 from "./logsyslogd2";
import * as logsyslogd3 from "./logsyslogd3";
import * as logsyslogd4 from "./logsyslogd4";
import * as logtacacsaccounting from "./logtacacsaccounting";
import * as logtacacsaccounting2 from "./logtacacsaccounting2";
import * as logtacacsaccounting3 from "./logtacacsaccounting3";
import * as logwebtrends from "./logwebtrends";
import * as networking from "./networking";
import * as nsxt from "./nsxt";
import * as report from "./report";
import * as router from "./router";
import * as routerbgp from "./routerbgp";
import * as routerospf from "./routerospf";
import * as routerospf6 from "./routerospf6";
import * as sctpfilter from "./sctpfilter";
import * as spamfilter from "./spamfilter";
import * as sshfilter from "./sshfilter";
import * as switchcontroller from "./switchcontroller";
import * as switchcontrollerautoconfig from "./switchcontrollerautoconfig";
import * as switchcontrollerinitialconfig from "./switchcontrollerinitialconfig";
import * as switchcontrollerptp from "./switchcontrollerptp";
import * as switchcontrollerqos from "./switchcontrollerqos";
import * as switchcontrollersecuritypolicy from "./switchcontrollersecuritypolicy";
import * as sys from "./sys";
import * as system3gmodem from "./system3gmodem";
import * as systemautoupdate from "./systemautoupdate";
import * as systemdhcp from "./systemdhcp";
import * as systemdhcp6 from "./systemdhcp6";
import * as systemlldp from "./systemlldp";
import * as systemreplacemsg from "./systemreplacemsg";
import * as systemsnmp from "./systemsnmp";
import * as types from "./types";
import * as user from "./user";
import * as videofilter from "./videofilter";
import * as voip from "./voip";
import * as vpn from "./vpn";
import * as vpncertificate from "./vpncertificate";
import * as vpnipsec from "./vpnipsec";
import * as vpnssl from "./vpnssl";
import * as vpnsslweb from "./vpnsslweb";
import * as waf from "./waf";
import * as wanopt from "./wanopt";
import * as webfilter from "./webfilter";
import * as webproxy from "./webproxy";
import * as wirelesscontroller from "./wirelesscontroller";
import * as wirelesscontrollerhotspot20 from "./wirelesscontrollerhotspot20";

export {
    alertemail,
    antivirus,
    application,
    authentication,
    automation,
    certificate,
    cifs,
    config,
    credentialstore,
    dlp,
    dnsfilter,
    dpdk,
    emailfilter,
    endpointcontrol,
    extendercontroller,
    extensioncontroller,
    filefilter,
    firewall,
    firewallconsolidated,
    firewallipmacbinding,
    firewallschedule,
    firewallservice,
    firewallshaper,
    firewallssh,
    firewallssl,
    firewallwildcardfqdn,
    fmg,
    ftpproxy,
    icap,
    ipmask,
    ips,
    json,
    log,
    logdisk,
    logfortianalyzer,
    logfortianalyzer2,
    logfortianalyzer3,
    logfortianalyzercloud,
    logfortiguard,
    logmemory,
    lognulldevice,
    logsyslogd,
    logsyslogd2,
    logsyslogd3,
    logsyslogd4,
    logtacacsaccounting,
    logtacacsaccounting2,
    logtacacsaccounting3,
    logwebtrends,
    networking,
    nsxt,
    report,
    router,
    routerbgp,
    routerospf,
    routerospf6,
    sctpfilter,
    spamfilter,
    sshfilter,
    switchcontroller,
    switchcontrollerautoconfig,
    switchcontrollerinitialconfig,
    switchcontrollerptp,
    switchcontrollerqos,
    switchcontrollersecuritypolicy,
    sys,
    system3gmodem,
    systemautoupdate,
    systemdhcp,
    systemdhcp6,
    systemlldp,
    systemreplacemsg,
    systemsnmp,
    types,
    user,
    videofilter,
    voip,
    vpn,
    vpncertificate,
    vpnipsec,
    vpnssl,
    vpnsslweb,
    waf,
    wanopt,
    webfilter,
    webproxy,
    wirelesscontroller,
    wirelesscontrollerhotspot20,
};
pulumi.runtime.registerResourcePackage("fortios", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:fortios") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
