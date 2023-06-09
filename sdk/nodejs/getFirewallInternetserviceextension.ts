// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios firewall internetserviceextension
 */
export function getFirewallInternetserviceextension(args: GetFirewallInternetserviceextensionArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallInternetserviceextensionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallInternetserviceextension:getFirewallInternetserviceextension", {
        "fosid": args.fosid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewallInternetserviceextension.
 */
export interface GetFirewallInternetserviceextensionArgs {
    /**
     * Specify the fosid of the desired firewall internetserviceextension.
     */
    fosid: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getFirewallInternetserviceextension.
 */
export interface GetFirewallInternetserviceextensionResult {
    /**
     * Comment.
     */
    readonly comment: string;
    /**
     * Disable entries in the Internet Service database. The structure of `disableEntry` block is documented below.
     */
    readonly disableEntries: outputs.GetFirewallInternetserviceextensionDisableEntry[];
    /**
     * Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
     */
    readonly entries: outputs.GetFirewallInternetserviceextensionEntry[];
    /**
     * Internet Service ID in the Internet Service database.
     */
    readonly fosid: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios firewall internetserviceextension
 */
export function getFirewallInternetserviceextensionOutput(args: GetFirewallInternetserviceextensionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallInternetserviceextensionResult> {
    return pulumi.output(args).apply((a: any) => getFirewallInternetserviceextension(a, opts))
}

/**
 * A collection of arguments for invoking getFirewallInternetserviceextension.
 */
export interface GetFirewallInternetserviceextensionOutputArgs {
    /**
     * Specify the fosid of the desired firewall internetserviceextension.
     */
    fosid: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
