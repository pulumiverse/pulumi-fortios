// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios system gretunnel
 */
export function getGretunnel(args: GetGretunnelArgs, opts?: pulumi.InvokeOptions): Promise<GetGretunnelResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:system/getGretunnel:getGretunnel", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getGretunnel.
 */
export interface GetGretunnelArgs {
    /**
     * Specify the name of the desired system gretunnel.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getGretunnel.
 */
export interface GetGretunnelResult {
    /**
     * Enable/disable validating checksums in received GRE packets.
     */
    readonly checksumReception: string;
    /**
     * Enable/disable including checksums in transmitted GRE packets.
     */
    readonly checksumTransmission: string;
    /**
     * DiffServ setting to be applied to GRE tunnel outer IP header.
     */
    readonly diffservcode: string;
    /**
     * Enable/disable DSCP copying.
     */
    readonly dscpCopying: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Interface name.
     */
    readonly interface: string;
    /**
     * IP version to use for VPN interface.
     */
    readonly ipVersion: string;
    /**
     * Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
     */
    readonly keepaliveFailtimes: number;
    /**
     * Keepalive message interval (0 - 32767, 0 = disabled).
     */
    readonly keepaliveInterval: number;
    /**
     * Require received GRE packets contain this key (0 - 4294967295).
     */
    readonly keyInbound: number;
    /**
     * Include this key in transmitted GRE packets (0 - 4294967295).
     */
    readonly keyOutbound: number;
    /**
     * IP address of the local gateway.
     */
    readonly localGw: string;
    /**
     * IPv6 address of the local gateway.
     */
    readonly localGw6: string;
    /**
     * Tunnel name.
     */
    readonly name: string;
    /**
     * IP address of the remote gateway.
     */
    readonly remoteGw: string;
    /**
     * IPv6 address of the remote gateway.
     */
    readonly remoteGw6: string;
    /**
     * Enable/disable validating sequence numbers in received GRE packets.
     */
    readonly sequenceNumberReception: string;
    /**
     * Enable/disable including of sequence numbers in transmitted GRE packets.
     */
    readonly sequenceNumberTransmission: string;
    /**
     * Enable/disable use of SD-WAN to reach remote gateway.
     */
    readonly useSdwan: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios system gretunnel
 */
export function getGretunnelOutput(args: GetGretunnelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGretunnelResult> {
    return pulumi.output(args).apply((a: any) => getGretunnel(a, opts))
}

/**
 * A collection of arguments for invoking getGretunnel.
 */
export interface GetGretunnelOutputArgs {
    /**
     * Specify the name of the desired system gretunnel.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
