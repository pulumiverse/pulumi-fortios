// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios firewallschedule onetime
 */
export function getOnetime(args: GetOnetimeArgs, opts?: pulumi.InvokeOptions): Promise<GetOnetimeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:firewallschedule/getOnetime:getOnetime", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getOnetime.
 */
export interface GetOnetimeArgs {
    /**
     * Specify the name of the desired firewallschedule onetime.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getOnetime.
 */
export interface GetOnetimeResult {
    /**
     * Color of icon on the GUI.
     */
    readonly color: number;
    /**
     * Schedule end date and time, format hh:mm yyyy/mm/dd.
     */
    readonly end: string;
    /**
     * Write an event log message this many days before the schedule expires.
     */
    readonly expirationDays: number;
    /**
     * Security Fabric global object setting.
     */
    readonly fabricObject: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Onetime schedule name.
     */
    readonly name: string;
    /**
     * Schedule start date and time, format hh:mm yyyy/mm/dd.
     */
    readonly start: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios firewallschedule onetime
 */
export function getOnetimeOutput(args: GetOnetimeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOnetimeResult> {
    return pulumi.output(args).apply((a: any) => getOnetime(a, opts))
}

/**
 * A collection of arguments for invoking getOnetime.
 */
export interface GetOnetimeOutputArgs {
    /**
     * Specify the name of the desired firewallschedule onetime.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
