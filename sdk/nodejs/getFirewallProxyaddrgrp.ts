// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios firewall proxyaddrgrp
 */
export function getFirewallProxyaddrgrp(args: GetFirewallProxyaddrgrpArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallProxyaddrgrpResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallProxyaddrgrp:getFirewallProxyaddrgrp", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewallProxyaddrgrp.
 */
export interface GetFirewallProxyaddrgrpArgs {
    /**
     * Specify the name of the desired firewall proxyaddrgrp.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getFirewallProxyaddrgrp.
 */
export interface GetFirewallProxyaddrgrpResult {
    /**
     * Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
     */
    readonly color: number;
    /**
     * Optional comments.
     */
    readonly comment: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Members of address group. The structure of `member` block is documented below.
     */
    readonly members: outputs.GetFirewallProxyaddrgrpMember[];
    /**
     * Tag name.
     */
    readonly name: string;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    readonly taggings: outputs.GetFirewallProxyaddrgrpTagging[];
    /**
     * Source or destination address group type.
     */
    readonly type: string;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    readonly uuid: string;
    readonly vdomparam?: string;
    /**
     * Enable/disable visibility of the object in the GUI.
     */
    readonly visibility: string;
}
/**
 * Use this data source to get information on an fortios firewall proxyaddrgrp
 */
export function getFirewallProxyaddrgrpOutput(args: GetFirewallProxyaddrgrpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallProxyaddrgrpResult> {
    return pulumi.output(args).apply((a: any) => getFirewallProxyaddrgrp(a, opts))
}

/**
 * A collection of arguments for invoking getFirewallProxyaddrgrp.
 */
export interface GetFirewallProxyaddrgrpOutputArgs {
    /**
     * Specify the name of the desired firewall proxyaddrgrp.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}