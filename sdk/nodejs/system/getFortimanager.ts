// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on fortios system fortimanager
 */
export function getFortimanager(args?: GetFortimanagerArgs, opts?: pulumi.InvokeOptions): Promise<GetFortimanagerResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:system/getFortimanager:getFortimanager", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getFortimanager.
 */
export interface GetFortimanagerArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getFortimanager.
 */
export interface GetFortimanagerResult {
    /**
     * Enable/disable FortiManager central management.
     */
    readonly centralManagement: string;
    /**
     * Enable/disable central management auto backup.
     */
    readonly centralMgmtAutoBackup: string;
    /**
     * Enable/disable central management schedule config restore.
     */
    readonly centralMgmtScheduleConfigRestore: string;
    /**
     * Enable/disable central management schedule script restore.
     */
    readonly centralMgmtScheduleScriptRestore: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * IP address.
     */
    readonly ip: string;
    /**
     * Enable/disable FortiManager IPsec tunnel.
     */
    readonly ipsec: string;
    /**
     * Virtual domain name.
     */
    readonly vdom: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on fortios system fortimanager
 */
export function getFortimanagerOutput(args?: GetFortimanagerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFortimanagerResult> {
    return pulumi.output(args).apply((a: any) => getFortimanager(a, opts))
}

/**
 * A collection of arguments for invoking getFortimanager.
 */
export interface GetFortimanagerOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
