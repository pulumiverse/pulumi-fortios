// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.SystemsnmpCommunity`.
 */
export function getSystemsnmpCommunitylist(args?: GetSystemsnmpCommunitylistArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemsnmpCommunitylistResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemsnmpCommunitylist:getSystemsnmpCommunitylist", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemsnmpCommunitylist.
 */
export interface GetSystemsnmpCommunitylistArgs {
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
 * A collection of values returned by getSystemsnmpCommunitylist.
 */
export interface GetSystemsnmpCommunitylistResult {
    readonly filter?: string;
    /**
     * A list of the `fortios.SystemsnmpCommunity`.
     */
    readonly fosidlists: number[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly vdomparam?: string;
}
/**
 * Provides a list of `fortios.SystemsnmpCommunity`.
 */
export function getSystemsnmpCommunitylistOutput(args?: GetSystemsnmpCommunitylistOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemsnmpCommunitylistResult> {
    return pulumi.output(args).apply((a: any) => getSystemsnmpCommunitylist(a, opts))
}

/**
 * A collection of arguments for invoking getSystemsnmpCommunitylist.
 */
export interface GetSystemsnmpCommunitylistOutputArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
