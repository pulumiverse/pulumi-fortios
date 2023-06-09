// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.SystemIpv6tunnel`.
 */
export function getSystemIpv6tunnellist(args?: GetSystemIpv6tunnellistArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemIpv6tunnellistResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemIpv6tunnellist:getSystemIpv6tunnellist", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemIpv6tunnellist.
 */
export interface GetSystemIpv6tunnellistArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSystemIpv6tunnellist.
 */
export interface GetSystemIpv6tunnellistResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.SystemIpv6tunnel`.
     */
    readonly namelists: string[];
    readonly vdomparam?: string;
}
/**
 * Provides a list of `fortios.SystemIpv6tunnel`.
 */
export function getSystemIpv6tunnellistOutput(args?: GetSystemIpv6tunnellistOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemIpv6tunnellistResult> {
    return pulumi.output(args).apply((a: any) => getSystemIpv6tunnellist(a, opts))
}

/**
 * A collection of arguments for invoking getSystemIpv6tunnellist.
 */
export interface GetSystemIpv6tunnellistOutputArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
