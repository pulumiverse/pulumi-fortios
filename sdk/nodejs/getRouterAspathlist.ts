// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios router aspathlist
 */
export function getRouterAspathlist(args: GetRouterAspathlistArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterAspathlistResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getRouterAspathlist:getRouterAspathlist", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouterAspathlist.
 */
export interface GetRouterAspathlistArgs {
    /**
     * Specify the name of the desired router aspathlist.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getRouterAspathlist.
 */
export interface GetRouterAspathlistResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * AS path list name.
     */
    readonly name: string;
    /**
     * AS path list rule. The structure of `rule` block is documented below.
     */
    readonly rules: outputs.GetRouterAspathlistRule[];
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios router aspathlist
 */
export function getRouterAspathlistOutput(args: GetRouterAspathlistOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterAspathlistResult> {
    return pulumi.output(args).apply((a: any) => getRouterAspathlist(a, opts))
}

/**
 * A collection of arguments for invoking getRouterAspathlist.
 */
export interface GetRouterAspathlistOutputArgs {
    /**
     * Specify the name of the desired router aspathlist.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}