// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios system ipiptunnel
 */
export function getIpiptunnel(args: GetIpiptunnelArgs, opts?: pulumi.InvokeOptions): Promise<GetIpiptunnelResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:system/getIpiptunnel:getIpiptunnel", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpiptunnel.
 */
export interface GetIpiptunnelArgs {
    /**
     * Specify the name of the desired system ipiptunnel.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getIpiptunnel.
 */
export interface GetIpiptunnelResult {
    /**
     * Enable/disable tunnel ASIC offloading.
     */
    readonly autoAsicOffload: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Interface name that is associated with the incoming traffic from available options.
     */
    readonly interface: string;
    /**
     * IPv4 address for the local gateway.
     */
    readonly localGw: string;
    /**
     * IPIP Tunnel name.
     */
    readonly name: string;
    /**
     * IPv4 address for the remote gateway.
     */
    readonly remoteGw: string;
    /**
     * Enable/disable use of SD-WAN to reach remote gateway.
     */
    readonly useSdwan: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios system ipiptunnel
 */
export function getIpiptunnelOutput(args: GetIpiptunnelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIpiptunnelResult> {
    return pulumi.output(args).apply((a: any) => getIpiptunnel(a, opts))
}

/**
 * A collection of arguments for invoking getIpiptunnel.
 */
export interface GetIpiptunnelOutputArgs {
    /**
     * Specify the name of the desired system ipiptunnel.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
