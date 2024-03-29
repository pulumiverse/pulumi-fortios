// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on fortios system fssopolling
 */
export function getFssopolling(args?: GetFssopollingArgs, opts?: pulumi.InvokeOptions): Promise<GetFssopollingResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:system/getFssopolling:getFssopolling", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getFssopolling.
 */
export interface GetFssopollingArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getFssopolling.
 */
export interface GetFssopollingResult {
    /**
     * Password to connect to FSSO Agent.
     */
    readonly authPassword: string;
    /**
     * Enable/disable FSSO Agent Authentication.
     */
    readonly authentication: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Listening port to accept clients (1 - 65535).
     */
    readonly listeningPort: number;
    /**
     * Enable/disable FSSO Polling Mode.
     */
    readonly status: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on fortios system fssopolling
 */
export function getFssopollingOutput(args?: GetFssopollingOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFssopollingResult> {
    return pulumi.output(args).apply((a: any) => getFssopolling(a, opts))
}

/**
 * A collection of arguments for invoking getFssopolling.
 */
export interface GetFssopollingOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
