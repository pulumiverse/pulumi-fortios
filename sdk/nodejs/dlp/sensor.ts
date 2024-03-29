// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure DLP sensors.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.dlp.Sensor("trname", {
 *     dlpLog: "enable",
 *     extendedLog: "disable",
 *     flowBased: "enable",
 *     nacQuarLog: "disable",
 *     summaryProto: "smtp pop3",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Dlp Sensor can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:dlp/sensor:Sensor labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:dlp/sensor:Sensor labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Sensor extends pulumi.CustomResource {
    /**
     * Get an existing Sensor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SensorState, opts?: pulumi.CustomResourceOptions): Sensor {
        return new Sensor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:dlp/sensor:Sensor';

    /**
     * Returns true if the given object is an instance of Sensor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Sensor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Sensor.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable DLP logging. Valid values: `enable`, `disable`.
     */
    public readonly dlpLog!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * DLP sensor entries. The structure of `entries` block is documented below.
     */
    public readonly entries!: pulumi.Output<outputs.dlp.SensorEntry[] | undefined>;
    /**
     * Expression to evaluate.
     */
    public readonly eval!: pulumi.Output<string>;
    /**
     * Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
     */
    public readonly extendedLog!: pulumi.Output<string>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    public readonly featureSet!: pulumi.Output<string>;
    /**
     * Set up DLP filters for this sensor. The structure of `filter` block is documented below.
     */
    public readonly filters!: pulumi.Output<outputs.dlp.SensorFilter[] | undefined>;
    /**
     * Enable/disable flow-based DLP. Valid values: `enable`, `disable`.
     */
    public readonly flowBased!: pulumi.Output<string>;
    /**
     * Protocols to always content archive.
     */
    public readonly fullArchiveProto!: pulumi.Output<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`, `match-eval`.
     */
    public readonly matchType!: pulumi.Output<string>;
    /**
     * Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
     */
    public readonly nacQuarLog!: pulumi.Output<string>;
    /**
     * Name of the DLP sensor.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configure DLP options.
     */
    public readonly options!: pulumi.Output<string>;
    /**
     * Replacement message group used by this DLP sensor.
     */
    public readonly replacemsgGroup!: pulumi.Output<string>;
    /**
     * Protocols to always log summary.
     */
    public readonly summaryProto!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Sensor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SensorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SensorArgs | SensorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SensorState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dlpLog"] = state ? state.dlpLog : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = state ? state.entries : undefined;
            resourceInputs["eval"] = state ? state.eval : undefined;
            resourceInputs["extendedLog"] = state ? state.extendedLog : undefined;
            resourceInputs["featureSet"] = state ? state.featureSet : undefined;
            resourceInputs["filters"] = state ? state.filters : undefined;
            resourceInputs["flowBased"] = state ? state.flowBased : undefined;
            resourceInputs["fullArchiveProto"] = state ? state.fullArchiveProto : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["matchType"] = state ? state.matchType : undefined;
            resourceInputs["nacQuarLog"] = state ? state.nacQuarLog : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["summaryProto"] = state ? state.summaryProto : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SensorArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dlpLog"] = args ? args.dlpLog : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = args ? args.entries : undefined;
            resourceInputs["eval"] = args ? args.eval : undefined;
            resourceInputs["extendedLog"] = args ? args.extendedLog : undefined;
            resourceInputs["featureSet"] = args ? args.featureSet : undefined;
            resourceInputs["filters"] = args ? args.filters : undefined;
            resourceInputs["flowBased"] = args ? args.flowBased : undefined;
            resourceInputs["fullArchiveProto"] = args ? args.fullArchiveProto : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["matchType"] = args ? args.matchType : undefined;
            resourceInputs["nacQuarLog"] = args ? args.nacQuarLog : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["summaryProto"] = args ? args.summaryProto : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Sensor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Sensor resources.
 */
export interface SensorState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable/disable DLP logging. Valid values: `enable`, `disable`.
     */
    dlpLog?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * DLP sensor entries. The structure of `entries` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.dlp.SensorEntry>[]>;
    /**
     * Expression to evaluate.
     */
    eval?: pulumi.Input<string>;
    /**
     * Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    featureSet?: pulumi.Input<string>;
    /**
     * Set up DLP filters for this sensor. The structure of `filter` block is documented below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.dlp.SensorFilter>[]>;
    /**
     * Enable/disable flow-based DLP. Valid values: `enable`, `disable`.
     */
    flowBased?: pulumi.Input<string>;
    /**
     * Protocols to always content archive.
     */
    fullArchiveProto?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`, `match-eval`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
     */
    nacQuarLog?: pulumi.Input<string>;
    /**
     * Name of the DLP sensor.
     */
    name?: pulumi.Input<string>;
    /**
     * Configure DLP options.
     */
    options?: pulumi.Input<string>;
    /**
     * Replacement message group used by this DLP sensor.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Protocols to always log summary.
     */
    summaryProto?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Sensor resource.
 */
export interface SensorArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable/disable DLP logging. Valid values: `enable`, `disable`.
     */
    dlpLog?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * DLP sensor entries. The structure of `entries` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.dlp.SensorEntry>[]>;
    /**
     * Expression to evaluate.
     */
    eval?: pulumi.Input<string>;
    /**
     * Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    featureSet?: pulumi.Input<string>;
    /**
     * Set up DLP filters for this sensor. The structure of `filter` block is documented below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.dlp.SensorFilter>[]>;
    /**
     * Enable/disable flow-based DLP. Valid values: `enable`, `disable`.
     */
    flowBased?: pulumi.Input<string>;
    /**
     * Protocols to always content archive.
     */
    fullArchiveProto?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`, `match-eval`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
     */
    nacQuarLog?: pulumi.Input<string>;
    /**
     * Name of the DLP sensor.
     */
    name?: pulumi.Input<string>;
    /**
     * Configure DLP options.
     */
    options?: pulumi.Input<string>;
    /**
     * Replacement message group used by this DLP sensor.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Protocols to always log summary.
     */
    summaryProto?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
