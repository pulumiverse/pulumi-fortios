// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on fortios system fipscc
 */
export function getSystemFipscc(args?: GetSystemFipsccArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemFipsccResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemFipscc:getSystemFipscc", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemFipscc.
 */
export interface GetSystemFipsccArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSystemFipscc.
 */
export interface GetSystemFipsccResult {
    /**
     * Enable/disable/dynamic entropy token.
     */
    readonly entropyToken: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Enable/disable self tests after key generation.
     */
    readonly keyGenerationSelfTest: string;
    /**
     * Self test period.
     */
    readonly selfTestPeriod: number;
    /**
     * Enable/disable FIPS-CC mode.
     */
    readonly status: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on fortios system fipscc
 */
export function getSystemFipsccOutput(args?: GetSystemFipsccOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemFipsccResult> {
    return pulumi.output(args).apply((a: any) => getSystemFipscc(a, opts))
}

/**
 * A collection of arguments for invoking getSystemFipscc.
 */
export interface GetSystemFipsccOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
