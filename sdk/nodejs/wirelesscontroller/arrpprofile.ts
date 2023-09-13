// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles. Applies to FortiOS Version `>= 6.4.2`.
 *
 * ## Import
 *
 * WirelessController ArrpProfile can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:wirelesscontroller/arrpprofile:Arrpprofile labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:wirelesscontroller/arrpprofile:Arrpprofile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Arrpprofile extends pulumi.CustomResource {
    /**
     * Get an existing Arrpprofile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ArrpprofileState, opts?: pulumi.CustomResourceOptions): Arrpprofile {
        return new Arrpprofile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:wirelesscontroller/arrpprofile:Arrpprofile';

    /**
     * Returns true if the given object is an instance of Arrpprofile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Arrpprofile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Arrpprofile.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
     */
    public readonly darrpOptimize!: pulumi.Output<number>;
    /**
     * Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
     */
    public readonly darrpOptimizeSchedules!: pulumi.Output<outputs.wirelesscontroller.ArrpprofileDarrpOptimizeSchedule[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
     */
    public readonly includeDfsChannel!: pulumi.Output<string>;
    /**
     * Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
     */
    public readonly includeWeatherChannel!: pulumi.Output<string>;
    /**
     * Period in seconds to measure average transmit retries and receive errors (default = 300).
     */
    public readonly monitorPeriod!: pulumi.Output<number>;
    /**
     * WiFi ARRP profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable to override setting darrp-optimize and darrp-optimize-schedules (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly overrideDarrpOptimize!: pulumi.Output<string>;
    /**
     * Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
     */
    public readonly selectionPeriod!: pulumi.Output<number>;
    /**
     * Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
     */
    public readonly thresholdAp!: pulumi.Output<number>;
    /**
     * Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
     */
    public readonly thresholdChannelLoad!: pulumi.Output<number>;
    /**
     * Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
     */
    public readonly thresholdNoiseFloor!: pulumi.Output<string>;
    /**
     * Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
     */
    public readonly thresholdRxErrors!: pulumi.Output<number>;
    /**
     * Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
     */
    public readonly thresholdSpectralRssi!: pulumi.Output<string>;
    /**
     * Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
     */
    public readonly thresholdTxRetries!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
     */
    public readonly weightChannelLoad!: pulumi.Output<number>;
    /**
     * Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
     */
    public readonly weightDfsChannel!: pulumi.Output<number>;
    /**
     * Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
     */
    public readonly weightManagedAp!: pulumi.Output<number>;
    /**
     * Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
     */
    public readonly weightNoiseFloor!: pulumi.Output<number>;
    /**
     * Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
     */
    public readonly weightRogueAp!: pulumi.Output<number>;
    /**
     * Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
     */
    public readonly weightSpectralRssi!: pulumi.Output<number>;
    /**
     * Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
     */
    public readonly weightWeatherChannel!: pulumi.Output<number>;

    /**
     * Create a Arrpprofile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ArrpprofileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ArrpprofileArgs | ArrpprofileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ArrpprofileState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["darrpOptimize"] = state ? state.darrpOptimize : undefined;
            resourceInputs["darrpOptimizeSchedules"] = state ? state.darrpOptimizeSchedules : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["includeDfsChannel"] = state ? state.includeDfsChannel : undefined;
            resourceInputs["includeWeatherChannel"] = state ? state.includeWeatherChannel : undefined;
            resourceInputs["monitorPeriod"] = state ? state.monitorPeriod : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["overrideDarrpOptimize"] = state ? state.overrideDarrpOptimize : undefined;
            resourceInputs["selectionPeriod"] = state ? state.selectionPeriod : undefined;
            resourceInputs["thresholdAp"] = state ? state.thresholdAp : undefined;
            resourceInputs["thresholdChannelLoad"] = state ? state.thresholdChannelLoad : undefined;
            resourceInputs["thresholdNoiseFloor"] = state ? state.thresholdNoiseFloor : undefined;
            resourceInputs["thresholdRxErrors"] = state ? state.thresholdRxErrors : undefined;
            resourceInputs["thresholdSpectralRssi"] = state ? state.thresholdSpectralRssi : undefined;
            resourceInputs["thresholdTxRetries"] = state ? state.thresholdTxRetries : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["weightChannelLoad"] = state ? state.weightChannelLoad : undefined;
            resourceInputs["weightDfsChannel"] = state ? state.weightDfsChannel : undefined;
            resourceInputs["weightManagedAp"] = state ? state.weightManagedAp : undefined;
            resourceInputs["weightNoiseFloor"] = state ? state.weightNoiseFloor : undefined;
            resourceInputs["weightRogueAp"] = state ? state.weightRogueAp : undefined;
            resourceInputs["weightSpectralRssi"] = state ? state.weightSpectralRssi : undefined;
            resourceInputs["weightWeatherChannel"] = state ? state.weightWeatherChannel : undefined;
        } else {
            const args = argsOrState as ArrpprofileArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["darrpOptimize"] = args ? args.darrpOptimize : undefined;
            resourceInputs["darrpOptimizeSchedules"] = args ? args.darrpOptimizeSchedules : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["includeDfsChannel"] = args ? args.includeDfsChannel : undefined;
            resourceInputs["includeWeatherChannel"] = args ? args.includeWeatherChannel : undefined;
            resourceInputs["monitorPeriod"] = args ? args.monitorPeriod : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overrideDarrpOptimize"] = args ? args.overrideDarrpOptimize : undefined;
            resourceInputs["selectionPeriod"] = args ? args.selectionPeriod : undefined;
            resourceInputs["thresholdAp"] = args ? args.thresholdAp : undefined;
            resourceInputs["thresholdChannelLoad"] = args ? args.thresholdChannelLoad : undefined;
            resourceInputs["thresholdNoiseFloor"] = args ? args.thresholdNoiseFloor : undefined;
            resourceInputs["thresholdRxErrors"] = args ? args.thresholdRxErrors : undefined;
            resourceInputs["thresholdSpectralRssi"] = args ? args.thresholdSpectralRssi : undefined;
            resourceInputs["thresholdTxRetries"] = args ? args.thresholdTxRetries : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["weightChannelLoad"] = args ? args.weightChannelLoad : undefined;
            resourceInputs["weightDfsChannel"] = args ? args.weightDfsChannel : undefined;
            resourceInputs["weightManagedAp"] = args ? args.weightManagedAp : undefined;
            resourceInputs["weightNoiseFloor"] = args ? args.weightNoiseFloor : undefined;
            resourceInputs["weightRogueAp"] = args ? args.weightRogueAp : undefined;
            resourceInputs["weightSpectralRssi"] = args ? args.weightSpectralRssi : undefined;
            resourceInputs["weightWeatherChannel"] = args ? args.weightWeatherChannel : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Arrpprofile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Arrpprofile resources.
 */
export interface ArrpprofileState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
     */
    darrpOptimize?: pulumi.Input<number>;
    /**
     * Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
     */
    darrpOptimizeSchedules?: pulumi.Input<pulumi.Input<inputs.wirelesscontroller.ArrpprofileDarrpOptimizeSchedule>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
     */
    includeDfsChannel?: pulumi.Input<string>;
    /**
     * Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
     */
    includeWeatherChannel?: pulumi.Input<string>;
    /**
     * Period in seconds to measure average transmit retries and receive errors (default = 300).
     */
    monitorPeriod?: pulumi.Input<number>;
    /**
     * WiFi ARRP profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable to override setting darrp-optimize and darrp-optimize-schedules (default = disable). Valid values: `enable`, `disable`.
     */
    overrideDarrpOptimize?: pulumi.Input<string>;
    /**
     * Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
     */
    selectionPeriod?: pulumi.Input<number>;
    /**
     * Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
     */
    thresholdAp?: pulumi.Input<number>;
    /**
     * Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
     */
    thresholdChannelLoad?: pulumi.Input<number>;
    /**
     * Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
     */
    thresholdNoiseFloor?: pulumi.Input<string>;
    /**
     * Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
     */
    thresholdRxErrors?: pulumi.Input<number>;
    /**
     * Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
     */
    thresholdSpectralRssi?: pulumi.Input<string>;
    /**
     * Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
     */
    thresholdTxRetries?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
     */
    weightChannelLoad?: pulumi.Input<number>;
    /**
     * Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
     */
    weightDfsChannel?: pulumi.Input<number>;
    /**
     * Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
     */
    weightManagedAp?: pulumi.Input<number>;
    /**
     * Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
     */
    weightNoiseFloor?: pulumi.Input<number>;
    /**
     * Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
     */
    weightRogueAp?: pulumi.Input<number>;
    /**
     * Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
     */
    weightSpectralRssi?: pulumi.Input<number>;
    /**
     * Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
     */
    weightWeatherChannel?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Arrpprofile resource.
 */
export interface ArrpprofileArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
     */
    darrpOptimize?: pulumi.Input<number>;
    /**
     * Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrpOptimizeSchedules` block is documented below.
     */
    darrpOptimizeSchedules?: pulumi.Input<pulumi.Input<inputs.wirelesscontroller.ArrpprofileDarrpOptimizeSchedule>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
     */
    includeDfsChannel?: pulumi.Input<string>;
    /**
     * Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
     */
    includeWeatherChannel?: pulumi.Input<string>;
    /**
     * Period in seconds to measure average transmit retries and receive errors (default = 300).
     */
    monitorPeriod?: pulumi.Input<number>;
    /**
     * WiFi ARRP profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable to override setting darrp-optimize and darrp-optimize-schedules (default = disable). Valid values: `enable`, `disable`.
     */
    overrideDarrpOptimize?: pulumi.Input<string>;
    /**
     * Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
     */
    selectionPeriod?: pulumi.Input<number>;
    /**
     * Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
     */
    thresholdAp?: pulumi.Input<number>;
    /**
     * Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
     */
    thresholdChannelLoad?: pulumi.Input<number>;
    /**
     * Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
     */
    thresholdNoiseFloor?: pulumi.Input<string>;
    /**
     * Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
     */
    thresholdRxErrors?: pulumi.Input<number>;
    /**
     * Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
     */
    thresholdSpectralRssi?: pulumi.Input<string>;
    /**
     * Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
     */
    thresholdTxRetries?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
     */
    weightChannelLoad?: pulumi.Input<number>;
    /**
     * Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
     */
    weightDfsChannel?: pulumi.Input<number>;
    /**
     * Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
     */
    weightManagedAp?: pulumi.Input<number>;
    /**
     * Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
     */
    weightNoiseFloor?: pulumi.Input<number>;
    /**
     * Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
     */
    weightRogueAp?: pulumi.Input<number>;
    /**
     * Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
     */
    weightSpectralRssi?: pulumi.Input<number>;
    /**
     * Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
     */
    weightWeatherChannel?: pulumi.Input<number>;
}
