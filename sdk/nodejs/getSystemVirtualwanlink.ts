// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on fortios system virtualwanlink
 */
export function getSystemVirtualwanlink(args?: GetSystemVirtualwanlinkArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemVirtualwanlinkResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemVirtualwanlink:getSystemVirtualwanlink", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemVirtualwanlink.
 */
export interface GetSystemVirtualwanlinkArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSystemVirtualwanlink.
 */
export interface GetSystemVirtualwanlinkResult {
    /**
     * Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
     */
    readonly failAlertInterfaces: outputs.GetSystemVirtualwanlinkFailAlertInterface[];
    /**
     * Enable/disable SD-WAN Internet connection status checking (failure detection).
     */
    readonly failDetect: string;
    /**
     * Virtual WAN Link health-check.
     */
    readonly healthChecks: outputs.GetSystemVirtualwanlinkHealthCheck[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Algorithm or mode to use for load balancing Internet traffic to SD-WAN members.
     */
    readonly loadBalanceMode: string;
    /**
     * Member sequence number list. The structure of `members` block is documented below.
     */
    readonly members: outputs.GetSystemVirtualwanlinkMember[];
    /**
     * Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
     */
    readonly neighborHoldBootTime: number;
    /**
     * Enable/disable hold switching from the secondary neighbor to the primary neighbor.
     */
    readonly neighborHoldDown: string;
    /**
     * Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
     */
    readonly neighborHoldDownTime: number;
    /**
     * Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
     */
    readonly neighbors: outputs.GetSystemVirtualwanlinkNeighbor[];
    /**
     * Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
     */
    readonly services: outputs.GetSystemVirtualwanlinkService[];
    /**
     * Enable/disable SD-WAN service.
     */
    readonly status: string;
    readonly vdomparam?: string;
    /**
     * Configure SD-WAN zones. The structure of `zone` block is documented below.
     */
    readonly zones: outputs.GetSystemVirtualwanlinkZone[];
}
/**
 * Use this data source to get information on fortios system virtualwanlink
 */
export function getSystemVirtualwanlinkOutput(args?: GetSystemVirtualwanlinkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemVirtualwanlinkResult> {
    return pulumi.output(args).apply((a: any) => getSystemVirtualwanlink(a, opts))
}

/**
 * A collection of arguments for invoking getSystemVirtualwanlink.
 */
export interface GetSystemVirtualwanlinkOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
