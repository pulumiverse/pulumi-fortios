// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios system sessionhelper
 */
export function getSessionhelper(args: GetSessionhelperArgs, opts?: pulumi.InvokeOptions): Promise<GetSessionhelperResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:system/getSessionhelper:getSessionhelper", {
        "fosid": args.fosid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSessionhelper.
 */
export interface GetSessionhelperArgs {
    /**
     * Specify the fosid of the desired system sessionhelper.
     */
    fosid: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSessionhelper.
 */
export interface GetSessionhelperResult {
    /**
     * Session helper ID.
     */
    readonly fosid: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Helper name.
     */
    readonly name: string;
    /**
     * Protocol port.
     */
    readonly port: number;
    /**
     * Protocol number.
     */
    readonly protocol: number;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios system sessionhelper
 */
export function getSessionhelperOutput(args: GetSessionhelperOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSessionhelperResult> {
    return pulumi.output(args).apply((a: any) => getSessionhelper(a, opts))
}

/**
 * A collection of arguments for invoking getSessionhelper.
 */
export interface GetSessionhelperOutputArgs {
    /**
     * Specify the fosid of the desired system sessionhelper.
     */
    fosid: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
