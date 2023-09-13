// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on fortios router bfd
 */
export function getBfd(args?: GetBfdArgs, opts?: pulumi.InvokeOptions): Promise<GetBfdResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:router/getBfd:getBfd", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getBfd.
 */
export interface GetBfdArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getBfd.
 */
export interface GetBfdResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * BFD multi-hop template table. The structure of `multihopTemplate` block is documented below.
     */
    readonly multihopTemplates: outputs.router.GetBfdMultihopTemplate[];
    /**
     * neighbor The structure of `neighbor` block is documented below.
     */
    readonly neighbors: outputs.router.GetBfdNeighbor[];
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on fortios router bfd
 */
export function getBfdOutput(args?: GetBfdOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBfdResult> {
    return pulumi.output(args).apply((a: any) => getBfd(a, opts))
}

/**
 * A collection of arguments for invoking getBfd.
 */
export interface GetBfdOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
