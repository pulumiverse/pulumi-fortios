// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios firewall internetservicedefinition
 */
export function getInternetservicedefinition(args: GetInternetservicedefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetInternetservicedefinitionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:firewall/getInternetservicedefinition:getInternetservicedefinition", {
        "fosid": args.fosid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getInternetservicedefinition.
 */
export interface GetInternetservicedefinitionArgs {
    /**
     * Specify the fosid of the desired firewall internetservicedefinition.
     */
    fosid: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getInternetservicedefinition.
 */
export interface GetInternetservicedefinitionResult {
    /**
     * Protocol and port information in an Internet Service entry. The structure of `entry` block is documented below.
     */
    readonly entries: outputs.firewall.GetInternetservicedefinitionEntry[];
    /**
     * Internet Service application list ID.
     */
    readonly fosid: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios firewall internetservicedefinition
 */
export function getInternetservicedefinitionOutput(args: GetInternetservicedefinitionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInternetservicedefinitionResult> {
    return pulumi.output(args).apply((a: any) => getInternetservicedefinition(a, opts))
}

/**
 * A collection of arguments for invoking getInternetservicedefinition.
 */
export interface GetInternetservicedefinitionOutputArgs {
    /**
     * Specify the fosid of the desired firewall internetservicedefinition.
     */
    fosid: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
