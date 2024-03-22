// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Filters for FortiAnalyzer.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.log.fortianalyzer.Filter("trname", {
 *     anomaly: "enable",
 *     dlpArchive: "enable",
 *     dns: "enable",
 *     filterType: "include",
 *     forwardTraffic: "enable",
 *     gtp: "enable",
 *     localTraffic: "enable",
 *     multicastTraffic: "enable",
 *     severity: "information",
 *     snifferTraffic: "enable",
 *     ssh: "enable",
 *     voip: "enable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * LogFortianalyzer Filter can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:log/fortianalyzer/filter:Filter labelname LogFortianalyzerFilter
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:log/fortianalyzer/filter:Filter labelname LogFortianalyzerFilter
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Filter extends pulumi.CustomResource {
    /**
     * Get an existing Filter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FilterState, opts?: pulumi.CustomResourceOptions): Filter {
        return new Filter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:log/fortianalyzer/filter:Filter';

    /**
     * Returns true if the given object is an instance of Filter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Filter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Filter.__pulumiType;
    }

    /**
     * Enable/disable anomaly logging. Valid values: `enable`, `disable`.
     */
    public readonly anomaly!: pulumi.Output<string>;
    /**
     * Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
     */
    public readonly dlpArchive!: pulumi.Output<string>;
    /**
     * Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
     */
    public readonly dns!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * FortiAnalyzer log filter.
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
     */
    public readonly filterType!: pulumi.Output<string>;
    /**
     * Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
     */
    public readonly fortiSwitch!: pulumi.Output<string>;
    /**
     * Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly forwardTraffic!: pulumi.Output<string>;
    /**
     * Free Style Filters The structure of `freeStyle` block is documented below.
     */
    public readonly freeStyles!: pulumi.Output<outputs.log.fortianalyzer.FilterFreeStyle[] | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
     */
    public readonly gtp!: pulumi.Output<string>;
    /**
     * Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly localTraffic!: pulumi.Output<string>;
    /**
     * Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly multicastTraffic!: pulumi.Output<string>;
    /**
     * Enable/disable netscan discovery event logging.
     */
    public readonly netscanDiscovery!: pulumi.Output<string>;
    /**
     * Enable/disable netscan vulnerability event logging.
     */
    public readonly netscanVulnerability!: pulumi.Output<string>;
    /**
     * Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly snifferTraffic!: pulumi.Output<string>;
    /**
     * Enable/disable SSH logging. Valid values: `enable`, `disable`.
     */
    public readonly ssh!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable VoIP logging. Valid values: `enable`, `disable`.
     */
    public readonly voip!: pulumi.Output<string>;
    /**
     * Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly ztnaTraffic!: pulumi.Output<string>;

    /**
     * Create a Filter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FilterArgs | FilterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FilterState | undefined;
            resourceInputs["anomaly"] = state ? state.anomaly : undefined;
            resourceInputs["dlpArchive"] = state ? state.dlpArchive : undefined;
            resourceInputs["dns"] = state ? state.dns : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["filterType"] = state ? state.filterType : undefined;
            resourceInputs["fortiSwitch"] = state ? state.fortiSwitch : undefined;
            resourceInputs["forwardTraffic"] = state ? state.forwardTraffic : undefined;
            resourceInputs["freeStyles"] = state ? state.freeStyles : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["gtp"] = state ? state.gtp : undefined;
            resourceInputs["localTraffic"] = state ? state.localTraffic : undefined;
            resourceInputs["multicastTraffic"] = state ? state.multicastTraffic : undefined;
            resourceInputs["netscanDiscovery"] = state ? state.netscanDiscovery : undefined;
            resourceInputs["netscanVulnerability"] = state ? state.netscanVulnerability : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["snifferTraffic"] = state ? state.snifferTraffic : undefined;
            resourceInputs["ssh"] = state ? state.ssh : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["voip"] = state ? state.voip : undefined;
            resourceInputs["ztnaTraffic"] = state ? state.ztnaTraffic : undefined;
        } else {
            const args = argsOrState as FilterArgs | undefined;
            resourceInputs["anomaly"] = args ? args.anomaly : undefined;
            resourceInputs["dlpArchive"] = args ? args.dlpArchive : undefined;
            resourceInputs["dns"] = args ? args.dns : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["filterType"] = args ? args.filterType : undefined;
            resourceInputs["fortiSwitch"] = args ? args.fortiSwitch : undefined;
            resourceInputs["forwardTraffic"] = args ? args.forwardTraffic : undefined;
            resourceInputs["freeStyles"] = args ? args.freeStyles : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["gtp"] = args ? args.gtp : undefined;
            resourceInputs["localTraffic"] = args ? args.localTraffic : undefined;
            resourceInputs["multicastTraffic"] = args ? args.multicastTraffic : undefined;
            resourceInputs["netscanDiscovery"] = args ? args.netscanDiscovery : undefined;
            resourceInputs["netscanVulnerability"] = args ? args.netscanVulnerability : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["snifferTraffic"] = args ? args.snifferTraffic : undefined;
            resourceInputs["ssh"] = args ? args.ssh : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["voip"] = args ? args.voip : undefined;
            resourceInputs["ztnaTraffic"] = args ? args.ztnaTraffic : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Filter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Filter resources.
 */
export interface FilterState {
    /**
     * Enable/disable anomaly logging. Valid values: `enable`, `disable`.
     */
    anomaly?: pulumi.Input<string>;
    /**
     * Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
     */
    dlpArchive?: pulumi.Input<string>;
    /**
     * Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
     */
    dns?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * FortiAnalyzer log filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
     */
    filterType?: pulumi.Input<string>;
    /**
     * Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
     */
    fortiSwitch?: pulumi.Input<string>;
    /**
     * Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
     */
    forwardTraffic?: pulumi.Input<string>;
    /**
     * Free Style Filters The structure of `freeStyle` block is documented below.
     */
    freeStyles?: pulumi.Input<pulumi.Input<inputs.log.fortianalyzer.FilterFreeStyle>[]>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
     */
    gtp?: pulumi.Input<string>;
    /**
     * Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
     */
    localTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
     */
    multicastTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable netscan discovery event logging.
     */
    netscanDiscovery?: pulumi.Input<string>;
    /**
     * Enable/disable netscan vulnerability event logging.
     */
    netscanVulnerability?: pulumi.Input<string>;
    /**
     * Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
     */
    severity?: pulumi.Input<string>;
    /**
     * Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
     */
    snifferTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable SSH logging. Valid values: `enable`, `disable`.
     */
    ssh?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable VoIP logging. Valid values: `enable`, `disable`.
     */
    voip?: pulumi.Input<string>;
    /**
     * Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
     */
    ztnaTraffic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Filter resource.
 */
export interface FilterArgs {
    /**
     * Enable/disable anomaly logging. Valid values: `enable`, `disable`.
     */
    anomaly?: pulumi.Input<string>;
    /**
     * Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
     */
    dlpArchive?: pulumi.Input<string>;
    /**
     * Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
     */
    dns?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * FortiAnalyzer log filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
     */
    filterType?: pulumi.Input<string>;
    /**
     * Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
     */
    fortiSwitch?: pulumi.Input<string>;
    /**
     * Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
     */
    forwardTraffic?: pulumi.Input<string>;
    /**
     * Free Style Filters The structure of `freeStyle` block is documented below.
     */
    freeStyles?: pulumi.Input<pulumi.Input<inputs.log.fortianalyzer.FilterFreeStyle>[]>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
     */
    gtp?: pulumi.Input<string>;
    /**
     * Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
     */
    localTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
     */
    multicastTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable netscan discovery event logging.
     */
    netscanDiscovery?: pulumi.Input<string>;
    /**
     * Enable/disable netscan vulnerability event logging.
     */
    netscanVulnerability?: pulumi.Input<string>;
    /**
     * Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
     */
    severity?: pulumi.Input<string>;
    /**
     * Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
     */
    snifferTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable SSH logging. Valid values: `enable`, `disable`.
     */
    ssh?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable VoIP logging. Valid values: `enable`, `disable`.
     */
    voip?: pulumi.Input<string>;
    /**
     * Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
     */
    ztnaTraffic?: pulumi.Input<string>;
}
