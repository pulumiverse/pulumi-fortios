// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on fortios system fortiguard
 */
export function getSystemFortiguard(args?: GetSystemFortiguardArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemFortiguardResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemFortiguard:getSystemFortiguard", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemFortiguard.
 */
export interface GetSystemFortiguardArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSystemFortiguard.
 */
export interface GetSystemFortiguardResult {
    /**
     * Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance.
     */
    readonly antispamCache: string;
    /**
     * Maximum percent of FortiGate memory the antispam cache is allowed to use (1 - 15%).
     */
    readonly antispamCacheMpercent: number;
    /**
     * Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.
     */
    readonly antispamCacheTtl: number;
    /**
     * Expiration date of the FortiGuard antispam contract.
     */
    readonly antispamExpiration: number;
    /**
     * Enable/disable turning off the FortiGuard antispam service.
     */
    readonly antispamForceOff: string;
    /**
     * Interval of time between license checks for the FortiGuard antispam contract.
     */
    readonly antispamLicense: number;
    /**
     * Antispam query time out (1 - 30 sec, default = 7).
     */
    readonly antispamTimeout: number;
    /**
     * IP address of the FortiGuard anycast DNS rating server.
     */
    readonly anycastSdnsServerIp: string;
    /**
     * Port to connect to on the FortiGuard anycast DNS rating server.
     */
    readonly anycastSdnsServerPort: number;
    /**
     * Enable/disable automatic patch-level firmware upgrade from FortiGuard. The FortiGate unit searches for new patches only in the same major and minor version.
     */
    readonly autoFirmwareUpgrade: string;
    /**
     * Allowed day(s) of the week to start automatic patch-level firmware upgrade from FortiGuard.
     */
    readonly autoFirmwareUpgradeDay: string;
    /**
     * End time in the designated time window for automatic patch-level firmware upgrade from FortiGuard in 24 hour time (0 ~ 23, default = 4). When the end time is smaller than the start time, the end time is interpreted as the next day. The actual upgrade time is selected randomly within the time window.
     */
    readonly autoFirmwareUpgradeEndHour: number;
    /**
     * Start time in the designated time window for automatic patch-level firmware upgrade from FortiGuard in 24 hour time (0 ~ 23, default = 2). The actual upgrade time is selected randomly within the time window.
     */
    readonly autoFirmwareUpgradeStartHour: number;
    /**
     * Automatically connect to and login to FortiCloud.
     */
    readonly autoJoinForticloud: string;
    /**
     * IP address of the FortiDDNS server.
     */
    readonly ddnsServerIp: string;
    /**
     * IPv6 address of the FortiDDNS server.
     */
    readonly ddnsServerIp6: string;
    /**
     * Port used to communicate with FortiDDNS servers.
     */
    readonly ddnsServerPort: number;
    /**
     * Enable/disable use of FortiGuard's anycast network.
     */
    readonly fortiguardAnycast: string;
    /**
     * Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet.
     */
    readonly fortiguardAnycastSource: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Specify outgoing interface to reach server.
     */
    readonly interface: string;
    /**
     * Specify how to select outgoing interface to reach server.
     */
    readonly interfaceSelectMethod: string;
    /**
     * Number of servers to alternate between as first FortiGuard option.
     */
    readonly loadBalanceServers: number;
    /**
     * Enable/disable FortiGuard Virus Outbreak Prevention cache.
     */
    readonly outbreakPreventionCache: string;
    /**
     * Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).
     */
    readonly outbreakPreventionCacheMpercent: number;
    /**
     * Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).
     */
    readonly outbreakPreventionCacheTtl: number;
    /**
     * Expiration date of FortiGuard Virus Outbreak Prevention contract.
     */
    readonly outbreakPreventionExpiration: number;
    /**
     * Turn off FortiGuard Virus Outbreak Prevention service.
     */
    readonly outbreakPreventionForceOff: string;
    /**
     * Interval of time between license checks for FortiGuard Virus Outbreak Prevention contract.
     */
    readonly outbreakPreventionLicense: number;
    /**
     * FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).
     */
    readonly outbreakPreventionTimeout: number;
    /**
     * Enable/disable use of persistent connection to receive update notification from FortiGuard.
     */
    readonly persistentConnection: string;
    /**
     * Port used to communicate with the FortiGuard servers.
     */
    readonly port: string;
    /**
     * Protocol used to communicate with the FortiGuard servers.
     */
    readonly protocol: string;
    /**
     * Proxy user password.
     */
    readonly proxyPassword: string;
    /**
     * IP address of the proxy server.
     */
    readonly proxyServerIp: string;
    /**
     * Port used to communicate with the proxy server.
     */
    readonly proxyServerPort: number;
    /**
     * Proxy user name.
     */
    readonly proxyUsername: string;
    /**
     * Enable/disable FortiCloud Sandbox inline-scan.
     */
    readonly sandboxInlineScan: string;
    /**
     * Cloud sandbox region.
     */
    readonly sandboxRegion: string;
    /**
     * Customization options for the FortiGuard DNS service.
     */
    readonly sdnsOptions: string;
    /**
     * IP address of the FortiDNS server.
     */
    readonly sdnsServerIp: string;
    /**
     * Port used to communicate with FortiDNS servers.
     */
    readonly sdnsServerPort: number;
    /**
     * Service account ID.
     */
    readonly serviceAccountId: string;
    /**
     * Source IPv4 address used to communicate with FortiGuard.
     */
    readonly sourceIp: string;
    /**
     * Source IPv6 address used to communicate with FortiGuard.
     */
    readonly sourceIp6: string;
    /**
     * Enable/disable proxy dictionary rebuild.
     */
    readonly updateBuildProxy: string;
    /**
     * Enable/disable external resource update.
     */
    readonly updateExtdb: string;
    /**
     * Enable/disable Internet Service Database update.
     */
    readonly updateFfdb: string;
    /**
     * Signature update server location.
     */
    readonly updateServerLocation: string;
    /**
     * Enable/disable allowlist update.
     */
    readonly updateUwdb: string;
    /**
     * FortiGuard Service virtual domain name.
     */
    readonly vdom: string;
    readonly vdomparam?: string;
    /**
     * Expiration date of the FortiGuard video filter contract.
     */
    readonly videofilterExpiration: number;
    /**
     * Interval of time between license checks for the FortiGuard video filter contract.
     */
    readonly videofilterLicense: number;
    /**
     * Enable/disable FortiGuard web filter caching.
     */
    readonly webfilterCache: string;
    /**
     * Time-to-live for web filter cache entries in seconds (300 - 86400).
     */
    readonly webfilterCacheTtl: number;
    /**
     * Expiration date of the FortiGuard web filter contract.
     */
    readonly webfilterExpiration: number;
    /**
     * Enable/disable turning off the FortiGuard web filtering service.
     */
    readonly webfilterForceOff: string;
    /**
     * Interval of time between license checks for the FortiGuard web filter contract.
     */
    readonly webfilterLicense: number;
    /**
     * Web filter query time out (1 - 30 sec, default = 7).
     */
    readonly webfilterTimeout: number;
}
/**
 * Use this data source to get information on fortios system fortiguard
 */
export function getSystemFortiguardOutput(args?: GetSystemFortiguardOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemFortiguardResult> {
    return pulumi.output(args).apply((a: any) => getSystemFortiguard(a, opts))
}

/**
 * A collection of arguments for invoking getSystemFortiguard.
 */
export interface GetSystemFortiguardOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
