// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios system ipv6tunnel
 */
export function getSystemIpv6tunnel(args: GetSystemIpv6tunnelArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemIpv6tunnelResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemIpv6tunnel:getSystemIpv6tunnel", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemIpv6tunnel.
 */
export interface GetSystemIpv6tunnelArgs {
    /**
     * Specify the name of the desired system ipv6tunnel.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSystemIpv6tunnel.
 */
export interface GetSystemIpv6tunnelResult {
    /**
     * Enable/disable tunnel ASIC offloading.
     */
    readonly autoAsicOffload: string;
    /**
     * Remote IPv6 address of the tunnel.
     */
    readonly destination: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Interface name.
     */
    readonly interface: string;
    /**
     * IPv6 tunnel name.
     */
    readonly name: string;
    /**
     * Local IPv6 address of the tunnel.
     */
    readonly source: string;
    /**
     * Enable/disable use of SD-WAN to reach remote gateway.
     */
    readonly useSdwan: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios system ipv6tunnel
 */
export function getSystemIpv6tunnelOutput(args: GetSystemIpv6tunnelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemIpv6tunnelResult> {
    return pulumi.output(args).apply((a: any) => getSystemIpv6tunnel(a, opts))
}

/**
 * A collection of arguments for invoking getSystemIpv6tunnel.
 */
export interface GetSystemIpv6tunnelOutputArgs {
    /**
     * Specify the name of the desired system ipv6tunnel.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
