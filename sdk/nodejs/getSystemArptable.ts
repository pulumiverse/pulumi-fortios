// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios system arptable
 */
export function getSystemArptable(args: GetSystemArptableArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemArptableResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemArptable:getSystemArptable", {
        "fosid": args.fosid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemArptable.
 */
export interface GetSystemArptableArgs {
    /**
     * Specify the fosid of the desired system arptable.
     */
    fosid: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSystemArptable.
 */
export interface GetSystemArptableResult {
    /**
     * Unique integer ID of the entry.
     */
    readonly fosid: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Interface name.
     */
    readonly interface: string;
    /**
     * IP address.
     */
    readonly ip: string;
    /**
     * MAC address.
     */
    readonly mac: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios system arptable
 */
export function getSystemArptableOutput(args: GetSystemArptableOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemArptableResult> {
    return pulumi.output(args).apply((a: any) => getSystemArptable(a, opts))
}

/**
 * A collection of arguments for invoking getSystemArptable.
 */
export interface GetSystemArptableOutputArgs {
    /**
     * Specify the fosid of the desired system arptable.
     */
    fosid: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}