// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.FirewallserviceCustom`.
 */
export function getFirewallserviceCustomlist(args?: GetFirewallserviceCustomlistArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallserviceCustomlistResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallserviceCustomlist:getFirewallserviceCustomlist", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewallserviceCustomlist.
 */
export interface GetFirewallserviceCustomlistArgs {
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
 * A collection of values returned by getFirewallserviceCustomlist.
 */
export interface GetFirewallserviceCustomlistResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.FirewallserviceCustom`.
     */
    readonly namelists: string[];
    readonly vdomparam?: string;
}
/**
 * Provides a list of `fortios.FirewallserviceCustom`.
 */
export function getFirewallserviceCustomlistOutput(args?: GetFirewallserviceCustomlistOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallserviceCustomlistResult> {
    return pulumi.output(args).apply((a: any) => getFirewallserviceCustomlist(a, opts))
}

/**
 * A collection of arguments for invoking getFirewallserviceCustomlist.
 */
export interface GetFirewallserviceCustomlistOutputArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
