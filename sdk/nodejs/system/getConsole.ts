// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on fortios system console
 */
export function getConsole(args?: GetConsoleArgs, opts?: pulumi.InvokeOptions): Promise<GetConsoleResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:system/getConsole:getConsole", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getConsole.
 */
export interface GetConsoleArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getConsole.
 */
export interface GetConsoleResult {
    /**
     * Console baud rate.
     */
    readonly baudrate: string;
    /**
     * Enable/disable access for FortiExplorer.
     */
    readonly fortiexplorer: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Enable/disable serial console and FortiExplorer.
     */
    readonly login: string;
    /**
     * Console mode.
     */
    readonly mode: string;
    /**
     * Console output mode.
     */
    readonly output: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on fortios system console
 */
export function getConsoleOutput(args?: GetConsoleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConsoleResult> {
    return pulumi.output(args).apply((a: any) => getConsole(a, opts))
}

/**
 * A collection of arguments for invoking getConsole.
 */
export interface GetConsoleOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
