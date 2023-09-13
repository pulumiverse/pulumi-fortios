// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on fortios system hamonitor
 */
export function getHamonitor(args?: GetHamonitorArgs, opts?: pulumi.InvokeOptions): Promise<GetHamonitorResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:sys/getHamonitor:getHamonitor", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getHamonitor.
 */
export interface GetHamonitorArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getHamonitor.
 */
export interface GetHamonitorResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Enable/disable monitor VLAN interfaces.
     */
    readonly monitorVlan: string;
    readonly vdomparam?: string;
    /**
     * Configure heartbeat interval (seconds).
     */
    readonly vlanHbInterval: number;
    /**
     * VLAN lost heartbeat threshold (1 - 60).
     */
    readonly vlanHbLostThreshold: number;
}
/**
 * Use this data source to get information on fortios system hamonitor
 */
export function getHamonitorOutput(args?: GetHamonitorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHamonitorResult> {
    return pulumi.output(args).apply((a: any) => getHamonitor(a, opts))
}

/**
 * A collection of arguments for invoking getHamonitor.
 */
export interface GetHamonitorOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
