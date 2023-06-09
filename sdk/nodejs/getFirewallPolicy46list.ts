// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.FirewallPolicy46`.
 */
export function getFirewallPolicy46list(args?: GetFirewallPolicy46listArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallPolicy46listResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallPolicy46list:getFirewallPolicy46list", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewallPolicy46list.
 */
export interface GetFirewallPolicy46listArgs {
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
 * A collection of values returned by getFirewallPolicy46list.
 */
export interface GetFirewallPolicy46listResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.FirewallPolicy46`.
     */
    readonly policyidlists: number[];
    readonly vdomparam?: string;
}
/**
 * Provides a list of `fortios.FirewallPolicy46`.
 */
export function getFirewallPolicy46listOutput(args?: GetFirewallPolicy46listOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallPolicy46listResult> {
    return pulumi.output(args).apply((a: any) => getFirewallPolicy46list(a, opts))
}

/**
 * A collection of arguments for invoking getFirewallPolicy46list.
 */
export interface GetFirewallPolicy46listOutputArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
