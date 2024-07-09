// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure threat weight settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.log.Threatweight("trname", {
 *     applications: [
 *         {
 *             category: 2,
 *             id: 1,
 *             level: "low",
 *         },
 *         {
 *             category: 6,
 *             id: 2,
 *             level: "medium",
 *         },
 *     ],
 *     blockedConnection: "high",
 *     failedConnection: "low",
 *     ips: {
 *         criticalSeverity: "critical",
 *         highSeverity: "high",
 *         infoSeverity: "disable",
 *         lowSeverity: "low",
 *         mediumSeverity: "medium",
 *     },
 *     level: {
 *         critical: 50,
 *         high: 30,
 *         low: 5,
 *         medium: 10,
 *     },
 *     malware: {
 *         botnetConnection: "critical",
 *         commandBlocked: "disable",
 *         contentDisarm: "medium",
 *         fileBlocked: "low",
 *         mimefragmented: "disable",
 *         oversized: "disable",
 *         switchProto: "disable",
 *         virusFileTypeExecutable: "medium",
 *         virusInfected: "critical",
 *         virusOutbreakPrevention: "critical",
 *         virusScanError: "high",
 *     },
 *     status: "enable",
 *     urlBlockDetected: "high",
 *     webs: [
 *         {
 *             category: 26,
 *             id: 1,
 *             level: "high",
 *         },
 *         {
 *             category: 61,
 *             id: 2,
 *             level: "high",
 *         },
 *         {
 *             category: 86,
 *             id: 3,
 *             level: "high",
 *         },
 *         {
 *             category: 1,
 *             id: 4,
 *             level: "medium",
 *         },
 *         {
 *             category: 3,
 *             id: 5,
 *             level: "medium",
 *         },
 *         {
 *             category: 4,
 *             id: 6,
 *             level: "medium",
 *         },
 *         {
 *             category: 5,
 *             id: 7,
 *             level: "medium",
 *         },
 *         {
 *             category: 6,
 *             id: 8,
 *             level: "medium",
 *         },
 *         {
 *             category: 12,
 *             id: 9,
 *             level: "medium",
 *         },
 *         {
 *             category: 59,
 *             id: 10,
 *             level: "medium",
 *         },
 *         {
 *             category: 62,
 *             id: 11,
 *             level: "medium",
 *         },
 *         {
 *             category: 83,
 *             id: 12,
 *             level: "medium",
 *         },
 *         {
 *             category: 72,
 *             id: 13,
 *             level: "low",
 *         },
 *         {
 *             category: 14,
 *             id: 14,
 *             level: "low",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Log ThreatWeight can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:log/threatweight:Threatweight labelname LogThreatWeight
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:log/threatweight:Threatweight labelname LogThreatWeight
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Threatweight extends pulumi.CustomResource {
    /**
     * Get an existing Threatweight resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ThreatweightState, opts?: pulumi.CustomResourceOptions): Threatweight {
        return new Threatweight(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:log/threatweight:Threatweight';

    /**
     * Returns true if the given object is an instance of Threatweight.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Threatweight {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Threatweight.__pulumiType;
    }

    /**
     * Application-control threat weight settings. The structure of `application` block is documented below.
     */
    public readonly applications!: pulumi.Output<outputs.log.ThreatweightApplication[] | undefined>;
    /**
     * Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    public readonly blockedConnection!: pulumi.Output<string>;
    /**
     * Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    public readonly botnetConnectionDetected!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    public readonly failedConnection!: pulumi.Output<string>;
    /**
     * Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.
     */
    public readonly geolocations!: pulumi.Output<outputs.log.ThreatweightGeolocation[] | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * IPS threat weight settings. The structure of `ips` block is documented below.
     */
    public readonly ips!: pulumi.Output<outputs.log.ThreatweightIps>;
    /**
     * Score mapping for threat weight levels. The structure of `level` block is documented below.
     */
    public readonly level!: pulumi.Output<outputs.log.ThreatweightLevel>;
    /**
     * Anti-virus malware threat weight settings. The structure of `malware` block is documented below.
     */
    public readonly malware!: pulumi.Output<outputs.log.ThreatweightMalware>;
    /**
     * Enable/disable the threat weight feature. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    public readonly urlBlockDetected!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;
    /**
     * Web filtering threat weight settings. The structure of `web` block is documented below.
     */
    public readonly webs!: pulumi.Output<outputs.log.ThreatweightWeb[] | undefined>;

    /**
     * Create a Threatweight resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ThreatweightArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ThreatweightArgs | ThreatweightState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ThreatweightState | undefined;
            resourceInputs["applications"] = state ? state.applications : undefined;
            resourceInputs["blockedConnection"] = state ? state.blockedConnection : undefined;
            resourceInputs["botnetConnectionDetected"] = state ? state.botnetConnectionDetected : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["failedConnection"] = state ? state.failedConnection : undefined;
            resourceInputs["geolocations"] = state ? state.geolocations : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["ips"] = state ? state.ips : undefined;
            resourceInputs["level"] = state ? state.level : undefined;
            resourceInputs["malware"] = state ? state.malware : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["urlBlockDetected"] = state ? state.urlBlockDetected : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["webs"] = state ? state.webs : undefined;
        } else {
            const args = argsOrState as ThreatweightArgs | undefined;
            resourceInputs["applications"] = args ? args.applications : undefined;
            resourceInputs["blockedConnection"] = args ? args.blockedConnection : undefined;
            resourceInputs["botnetConnectionDetected"] = args ? args.botnetConnectionDetected : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["failedConnection"] = args ? args.failedConnection : undefined;
            resourceInputs["geolocations"] = args ? args.geolocations : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["ips"] = args ? args.ips : undefined;
            resourceInputs["level"] = args ? args.level : undefined;
            resourceInputs["malware"] = args ? args.malware : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["urlBlockDetected"] = args ? args.urlBlockDetected : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["webs"] = args ? args.webs : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Threatweight.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Threatweight resources.
 */
export interface ThreatweightState {
    /**
     * Application-control threat weight settings. The structure of `application` block is documented below.
     */
    applications?: pulumi.Input<pulumi.Input<inputs.log.ThreatweightApplication>[]>;
    /**
     * Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    blockedConnection?: pulumi.Input<string>;
    /**
     * Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    botnetConnectionDetected?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    failedConnection?: pulumi.Input<string>;
    /**
     * Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.
     */
    geolocations?: pulumi.Input<pulumi.Input<inputs.log.ThreatweightGeolocation>[]>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * IPS threat weight settings. The structure of `ips` block is documented below.
     */
    ips?: pulumi.Input<inputs.log.ThreatweightIps>;
    /**
     * Score mapping for threat weight levels. The structure of `level` block is documented below.
     */
    level?: pulumi.Input<inputs.log.ThreatweightLevel>;
    /**
     * Anti-virus malware threat weight settings. The structure of `malware` block is documented below.
     */
    malware?: pulumi.Input<inputs.log.ThreatweightMalware>;
    /**
     * Enable/disable the threat weight feature. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    urlBlockDetected?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Web filtering threat weight settings. The structure of `web` block is documented below.
     */
    webs?: pulumi.Input<pulumi.Input<inputs.log.ThreatweightWeb>[]>;
}

/**
 * The set of arguments for constructing a Threatweight resource.
 */
export interface ThreatweightArgs {
    /**
     * Application-control threat weight settings. The structure of `application` block is documented below.
     */
    applications?: pulumi.Input<pulumi.Input<inputs.log.ThreatweightApplication>[]>;
    /**
     * Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    blockedConnection?: pulumi.Input<string>;
    /**
     * Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    botnetConnectionDetected?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    failedConnection?: pulumi.Input<string>;
    /**
     * Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.
     */
    geolocations?: pulumi.Input<pulumi.Input<inputs.log.ThreatweightGeolocation>[]>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * IPS threat weight settings. The structure of `ips` block is documented below.
     */
    ips?: pulumi.Input<inputs.log.ThreatweightIps>;
    /**
     * Score mapping for threat weight levels. The structure of `level` block is documented below.
     */
    level?: pulumi.Input<inputs.log.ThreatweightLevel>;
    /**
     * Anti-virus malware threat weight settings. The structure of `malware` block is documented below.
     */
    malware?: pulumi.Input<inputs.log.ThreatweightMalware>;
    /**
     * Enable/disable the threat weight feature. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    urlBlockDetected?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Web filtering threat weight settings. The structure of `web` block is documented below.
     */
    webs?: pulumi.Input<pulumi.Input<inputs.log.ThreatweightWeb>[]>;
}
