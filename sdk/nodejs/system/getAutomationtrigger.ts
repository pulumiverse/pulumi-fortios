// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios system automationtrigger
 */
export function getAutomationtrigger(args: GetAutomationtriggerArgs, opts?: pulumi.InvokeOptions): Promise<GetAutomationtriggerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:system/getAutomationtrigger:getAutomationtrigger", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getAutomationtrigger.
 */
export interface GetAutomationtriggerArgs {
    /**
     * Specify the name of the desired system automationtrigger.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getAutomationtrigger.
 */
export interface GetAutomationtriggerResult {
    /**
     * Description.
     */
    readonly description: string;
    /**
     * Event type.
     */
    readonly eventType: string;
    /**
     * Fabric connector event handler name.
     */
    readonly fabricEventName: string;
    /**
     * Fabric connector event severity.
     */
    readonly fabricEventSeverity: string;
    /**
     * FortiAnalyzer event handler name.
     */
    readonly fazEventName: string;
    /**
     * FortiAnalyzer event severity.
     */
    readonly fazEventSeverity: string;
    /**
     * FortiAnalyzer event tags.
     */
    readonly fazEventTags: string;
    /**
     * Customized trigger field settings. The structure of `fields` block is documented below.
     */
    readonly fields: outputs.system.GetAutomationtriggerField[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * IOC threat level.
     */
    readonly iocLevel: string;
    /**
     * License type.
     */
    readonly licenseType: string;
    /**
     * Log ID to trigger event.
     */
    readonly logid: number;
    /**
     * Log IDs to trigger event. The structure of `logidBlock` block is documented below.
     */
    readonly logidBlocks: outputs.system.GetAutomationtriggerLogidBlock[];
    /**
     * Name.
     */
    readonly name: string;
    /**
     * Security Rating report.
     */
    readonly reportType: string;
    /**
     * Fabric connector serial number.
     */
    readonly serial: string;
    /**
     * Trigger date and time (YYYY-MM-DD HH:MM:SS).
     */
    readonly triggerDatetime: string;
    /**
     * Day within a month to trigger.
     */
    readonly triggerDay: number;
    /**
     * Scheduled trigger frequency (default = daily).
     */
    readonly triggerFrequency: string;
    /**
     * Hour of the day on which to trigger (0 - 23, default = 1).
     */
    readonly triggerHour: number;
    /**
     * Minute of the hour on which to trigger (0 - 59, 60 to randomize).
     */
    readonly triggerMinute: number;
    /**
     * Trigger type.
     */
    readonly triggerType: string;
    /**
     * Day of week for trigger.
     */
    readonly triggerWeekday: string;
    readonly vdomparam?: string;
    /**
     * Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
     */
    readonly vdoms: outputs.system.GetAutomationtriggerVdom[];
}
/**
 * Use this data source to get information on an fortios system automationtrigger
 */
export function getAutomationtriggerOutput(args: GetAutomationtriggerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAutomationtriggerResult> {
    return pulumi.output(args).apply((a: any) => getAutomationtrigger(a, opts))
}

/**
 * A collection of arguments for invoking getAutomationtrigger.
 */
export interface GetAutomationtriggerOutputArgs {
    /**
     * Specify the name of the desired system automationtrigger.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
