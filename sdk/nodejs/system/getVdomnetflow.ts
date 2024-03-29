// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on fortios system vdomnetflow
 */
export function getVdomnetflow(args?: GetVdomnetflowArgs, opts?: pulumi.InvokeOptions): Promise<GetVdomnetflowResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:system/getVdomnetflow:getVdomnetflow", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getVdomnetflow.
 */
export interface GetVdomnetflowArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getVdomnetflow.
 */
export interface GetVdomnetflowResult {
    /**
     * Collector IP.
     */
    readonly collectorIp: string;
    /**
     * NetFlow collector port number.
     */
    readonly collectorPort: number;
    /**
     * Netflow collectors. The structure of `collectors` block is documented below.
     */
    readonly collectors: outputs.system.GetVdomnetflowCollector[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Specify outgoing interface to reach server.
     */
    readonly interface: string;
    /**
     * Specify how to select outgoing interface to reach server.
     */
    readonly interfaceSelectMethod: string;
    /**
     * Source IP address for communication with the NetFlow agent.
     */
    readonly sourceIp: string;
    /**
     * Enable/disable NetFlow per VDOM.
     */
    readonly vdomNetflow: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on fortios system vdomnetflow
 */
export function getVdomnetflowOutput(args?: GetVdomnetflowOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVdomnetflowResult> {
    return pulumi.output(args).apply((a: any) => getVdomnetflow(a, opts))
}

/**
 * A collection of arguments for invoking getVdomnetflow.
 */
export interface GetVdomnetflowOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
