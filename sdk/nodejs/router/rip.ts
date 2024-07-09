// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure RIP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.router.Rip("trname", {
 *     defaultInformationOriginate: "disable",
 *     defaultMetric: 1,
 *     garbageTimer: 120,
 *     maxOutMetric: 0,
 *     recvBufferSize: 655360,
 *     redistributes: [
 *         {
 *             metric: 10,
 *             name: "connected",
 *             status: "disable",
 *         },
 *         {
 *             metric: 10,
 *             name: "static",
 *             status: "disable",
 *         },
 *         {
 *             metric: 10,
 *             name: "ospf",
 *             status: "disable",
 *         },
 *         {
 *             metric: 10,
 *             name: "bgp",
 *             status: "disable",
 *         },
 *         {
 *             metric: 10,
 *             name: "isis",
 *             status: "disable",
 *         },
 *     ],
 *     timeoutTimer: 180,
 *     updateTimer: 30,
 *     version: "2",
 * });
 * ```
 *
 * ## Import
 *
 * Router Rip can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:router/rip:Rip labelname RouterRip
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:router/rip:Rip labelname RouterRip
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Rip extends pulumi.CustomResource {
    /**
     * Get an existing Rip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RipState, opts?: pulumi.CustomResourceOptions): Rip {
        return new Rip(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:router/rip:Rip';

    /**
     * Returns true if the given object is an instance of Rip.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rip {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rip.__pulumiType;
    }

    /**
     * Enable/disable generation of default route. Valid values: `enable`, `disable`.
     */
    public readonly defaultInformationOriginate!: pulumi.Output<string>;
    /**
     * Default metric.
     */
    public readonly defaultMetric!: pulumi.Output<number>;
    /**
     * distance The structure of `distance` block is documented below.
     */
    public readonly distances!: pulumi.Output<outputs.router.RipDistance[] | undefined>;
    /**
     * Distribute list. The structure of `distributeList` block is documented below.
     */
    public readonly distributeLists!: pulumi.Output<outputs.router.RipDistributeList[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Garbage timer in seconds.
     */
    public readonly garbageTimer!: pulumi.Output<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * RIP interface configuration. The structure of `interface` block is documented below.
     */
    public readonly interfaces!: pulumi.Output<outputs.router.RipInterface[] | undefined>;
    /**
     * Maximum metric allowed to output(0 means 'not set').
     */
    public readonly maxOutMetric!: pulumi.Output<number>;
    /**
     * neighbor The structure of `neighbor` block is documented below.
     */
    public readonly neighbors!: pulumi.Output<outputs.router.RipNeighbor[] | undefined>;
    /**
     * network The structure of `network` block is documented below.
     */
    public readonly networks!: pulumi.Output<outputs.router.RipNetwork[] | undefined>;
    /**
     * Offset list. The structure of `offsetList` block is documented below.
     */
    public readonly offsetLists!: pulumi.Output<outputs.router.RipOffsetList[] | undefined>;
    /**
     * Passive interface configuration. The structure of `passiveInterface` block is documented below.
     */
    public readonly passiveInterfaces!: pulumi.Output<outputs.router.RipPassiveInterface[] | undefined>;
    /**
     * Receiving buffer size.
     */
    public readonly recvBufferSize!: pulumi.Output<number>;
    /**
     * Redistribute configuration. The structure of `redistribute` block is documented below.
     */
    public readonly redistributes!: pulumi.Output<outputs.router.RipRedistribute[] | undefined>;
    /**
     * Timeout timer in seconds.
     */
    public readonly timeoutTimer!: pulumi.Output<number>;
    /**
     * Update timer in seconds.
     */
    public readonly updateTimer!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;
    /**
     * RIP version. Valid values: `1`, `2`.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Rip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RipArgs | RipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RipState | undefined;
            resourceInputs["defaultInformationOriginate"] = state ? state.defaultInformationOriginate : undefined;
            resourceInputs["defaultMetric"] = state ? state.defaultMetric : undefined;
            resourceInputs["distances"] = state ? state.distances : undefined;
            resourceInputs["distributeLists"] = state ? state.distributeLists : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["garbageTimer"] = state ? state.garbageTimer : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["maxOutMetric"] = state ? state.maxOutMetric : undefined;
            resourceInputs["neighbors"] = state ? state.neighbors : undefined;
            resourceInputs["networks"] = state ? state.networks : undefined;
            resourceInputs["offsetLists"] = state ? state.offsetLists : undefined;
            resourceInputs["passiveInterfaces"] = state ? state.passiveInterfaces : undefined;
            resourceInputs["recvBufferSize"] = state ? state.recvBufferSize : undefined;
            resourceInputs["redistributes"] = state ? state.redistributes : undefined;
            resourceInputs["timeoutTimer"] = state ? state.timeoutTimer : undefined;
            resourceInputs["updateTimer"] = state ? state.updateTimer : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as RipArgs | undefined;
            resourceInputs["defaultInformationOriginate"] = args ? args.defaultInformationOriginate : undefined;
            resourceInputs["defaultMetric"] = args ? args.defaultMetric : undefined;
            resourceInputs["distances"] = args ? args.distances : undefined;
            resourceInputs["distributeLists"] = args ? args.distributeLists : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["garbageTimer"] = args ? args.garbageTimer : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["maxOutMetric"] = args ? args.maxOutMetric : undefined;
            resourceInputs["neighbors"] = args ? args.neighbors : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["offsetLists"] = args ? args.offsetLists : undefined;
            resourceInputs["passiveInterfaces"] = args ? args.passiveInterfaces : undefined;
            resourceInputs["recvBufferSize"] = args ? args.recvBufferSize : undefined;
            resourceInputs["redistributes"] = args ? args.redistributes : undefined;
            resourceInputs["timeoutTimer"] = args ? args.timeoutTimer : undefined;
            resourceInputs["updateTimer"] = args ? args.updateTimer : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Rip.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rip resources.
 */
export interface RipState {
    /**
     * Enable/disable generation of default route. Valid values: `enable`, `disable`.
     */
    defaultInformationOriginate?: pulumi.Input<string>;
    /**
     * Default metric.
     */
    defaultMetric?: pulumi.Input<number>;
    /**
     * distance The structure of `distance` block is documented below.
     */
    distances?: pulumi.Input<pulumi.Input<inputs.router.RipDistance>[]>;
    /**
     * Distribute list. The structure of `distributeList` block is documented below.
     */
    distributeLists?: pulumi.Input<pulumi.Input<inputs.router.RipDistributeList>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Garbage timer in seconds.
     */
    garbageTimer?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * RIP interface configuration. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.router.RipInterface>[]>;
    /**
     * Maximum metric allowed to output(0 means 'not set').
     */
    maxOutMetric?: pulumi.Input<number>;
    /**
     * neighbor The structure of `neighbor` block is documented below.
     */
    neighbors?: pulumi.Input<pulumi.Input<inputs.router.RipNeighbor>[]>;
    /**
     * network The structure of `network` block is documented below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.router.RipNetwork>[]>;
    /**
     * Offset list. The structure of `offsetList` block is documented below.
     */
    offsetLists?: pulumi.Input<pulumi.Input<inputs.router.RipOffsetList>[]>;
    /**
     * Passive interface configuration. The structure of `passiveInterface` block is documented below.
     */
    passiveInterfaces?: pulumi.Input<pulumi.Input<inputs.router.RipPassiveInterface>[]>;
    /**
     * Receiving buffer size.
     */
    recvBufferSize?: pulumi.Input<number>;
    /**
     * Redistribute configuration. The structure of `redistribute` block is documented below.
     */
    redistributes?: pulumi.Input<pulumi.Input<inputs.router.RipRedistribute>[]>;
    /**
     * Timeout timer in seconds.
     */
    timeoutTimer?: pulumi.Input<number>;
    /**
     * Update timer in seconds.
     */
    updateTimer?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * RIP version. Valid values: `1`, `2`.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Rip resource.
 */
export interface RipArgs {
    /**
     * Enable/disable generation of default route. Valid values: `enable`, `disable`.
     */
    defaultInformationOriginate?: pulumi.Input<string>;
    /**
     * Default metric.
     */
    defaultMetric?: pulumi.Input<number>;
    /**
     * distance The structure of `distance` block is documented below.
     */
    distances?: pulumi.Input<pulumi.Input<inputs.router.RipDistance>[]>;
    /**
     * Distribute list. The structure of `distributeList` block is documented below.
     */
    distributeLists?: pulumi.Input<pulumi.Input<inputs.router.RipDistributeList>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Garbage timer in seconds.
     */
    garbageTimer?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * RIP interface configuration. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.router.RipInterface>[]>;
    /**
     * Maximum metric allowed to output(0 means 'not set').
     */
    maxOutMetric?: pulumi.Input<number>;
    /**
     * neighbor The structure of `neighbor` block is documented below.
     */
    neighbors?: pulumi.Input<pulumi.Input<inputs.router.RipNeighbor>[]>;
    /**
     * network The structure of `network` block is documented below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.router.RipNetwork>[]>;
    /**
     * Offset list. The structure of `offsetList` block is documented below.
     */
    offsetLists?: pulumi.Input<pulumi.Input<inputs.router.RipOffsetList>[]>;
    /**
     * Passive interface configuration. The structure of `passiveInterface` block is documented below.
     */
    passiveInterfaces?: pulumi.Input<pulumi.Input<inputs.router.RipPassiveInterface>[]>;
    /**
     * Receiving buffer size.
     */
    recvBufferSize?: pulumi.Input<number>;
    /**
     * Redistribute configuration. The structure of `redistribute` block is documented below.
     */
    redistributes?: pulumi.Input<pulumi.Input<inputs.router.RipRedistribute>[]>;
    /**
     * Timeout timer in seconds.
     */
    timeoutTimer?: pulumi.Input<number>;
    /**
     * Update timer in seconds.
     */
    updateTimer?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * RIP version. Valid values: `1`, `2`.
     */
    version?: pulumi.Input<string>;
}
