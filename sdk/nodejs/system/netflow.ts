// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure NetFlow.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Netflow("trname", {
 *     activeFlowTimeout: 30,
 *     collectorIp: "0.0.0.0",
 *     collectorPort: 2055,
 *     inactiveFlowTimeout: 15,
 *     sourceIp: "0.0.0.0",
 *     templateTxCounter: 20,
 *     templateTxTimeout: 30,
 * });
 * ```
 *
 * ## Import
 *
 * System Netflow can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/netflow:Netflow labelname SystemNetflow
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/netflow:Netflow labelname SystemNetflow
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Netflow extends pulumi.CustomResource {
    /**
     * Get an existing Netflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetflowState, opts?: pulumi.CustomResourceOptions): Netflow {
        return new Netflow(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/netflow:Netflow';

    /**
     * Returns true if the given object is an instance of Netflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Netflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Netflow.__pulumiType;
    }

    /**
     * Timeout to report active flows. On FortiOS versions 6.2.0-7.0.0: 1 - 60 min, default = 30. On FortiOS versions >= 7.0.1: 60 - 3600 sec, default = 1800.
     */
    public readonly activeFlowTimeout!: pulumi.Output<number>;
    /**
     * Collector IP.
     */
    public readonly collectorIp!: pulumi.Output<string>;
    /**
     * NetFlow collector port number.
     */
    public readonly collectorPort!: pulumi.Output<number>;
    /**
     * Netflow collectors. The structure of `collectors` block is documented below.
     */
    public readonly collectors!: pulumi.Output<outputs.system.NetflowCollector[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
     */
    public readonly inactiveFlowTimeout!: pulumi.Output<number>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    /**
     * Source IP address for communication with the NetFlow agent.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Counter of flowset records before resending a template flowset record.
     */
    public readonly templateTxCounter!: pulumi.Output<number>;
    /**
     * Timeout for periodic template flowset transmission. On FortiOS versions 6.2.0-7.0.0: 1 - 1440 min, default = 30. On FortiOS versions >= 7.0.1: 60 - 86400 sec, default = 1800.
     */
    public readonly templateTxTimeout!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Netflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetflowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetflowArgs | NetflowState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetflowState | undefined;
            resourceInputs["activeFlowTimeout"] = state ? state.activeFlowTimeout : undefined;
            resourceInputs["collectorIp"] = state ? state.collectorIp : undefined;
            resourceInputs["collectorPort"] = state ? state.collectorPort : undefined;
            resourceInputs["collectors"] = state ? state.collectors : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["inactiveFlowTimeout"] = state ? state.inactiveFlowTimeout : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["templateTxCounter"] = state ? state.templateTxCounter : undefined;
            resourceInputs["templateTxTimeout"] = state ? state.templateTxTimeout : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as NetflowArgs | undefined;
            resourceInputs["activeFlowTimeout"] = args ? args.activeFlowTimeout : undefined;
            resourceInputs["collectorIp"] = args ? args.collectorIp : undefined;
            resourceInputs["collectorPort"] = args ? args.collectorPort : undefined;
            resourceInputs["collectors"] = args ? args.collectors : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["inactiveFlowTimeout"] = args ? args.inactiveFlowTimeout : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["templateTxCounter"] = args ? args.templateTxCounter : undefined;
            resourceInputs["templateTxTimeout"] = args ? args.templateTxTimeout : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Netflow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Netflow resources.
 */
export interface NetflowState {
    /**
     * Timeout to report active flows. On FortiOS versions 6.2.0-7.0.0: 1 - 60 min, default = 30. On FortiOS versions >= 7.0.1: 60 - 3600 sec, default = 1800.
     */
    activeFlowTimeout?: pulumi.Input<number>;
    /**
     * Collector IP.
     */
    collectorIp?: pulumi.Input<string>;
    /**
     * NetFlow collector port number.
     */
    collectorPort?: pulumi.Input<number>;
    /**
     * Netflow collectors. The structure of `collectors` block is documented below.
     */
    collectors?: pulumi.Input<pulumi.Input<inputs.system.NetflowCollector>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
     */
    inactiveFlowTimeout?: pulumi.Input<number>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Source IP address for communication with the NetFlow agent.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Counter of flowset records before resending a template flowset record.
     */
    templateTxCounter?: pulumi.Input<number>;
    /**
     * Timeout for periodic template flowset transmission. On FortiOS versions 6.2.0-7.0.0: 1 - 1440 min, default = 30. On FortiOS versions >= 7.0.1: 60 - 86400 sec, default = 1800.
     */
    templateTxTimeout?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Netflow resource.
 */
export interface NetflowArgs {
    /**
     * Timeout to report active flows. On FortiOS versions 6.2.0-7.0.0: 1 - 60 min, default = 30. On FortiOS versions >= 7.0.1: 60 - 3600 sec, default = 1800.
     */
    activeFlowTimeout?: pulumi.Input<number>;
    /**
     * Collector IP.
     */
    collectorIp?: pulumi.Input<string>;
    /**
     * NetFlow collector port number.
     */
    collectorPort?: pulumi.Input<number>;
    /**
     * Netflow collectors. The structure of `collectors` block is documented below.
     */
    collectors?: pulumi.Input<pulumi.Input<inputs.system.NetflowCollector>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
     */
    inactiveFlowTimeout?: pulumi.Input<number>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Source IP address for communication with the NetFlow agent.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Counter of flowset records before resending a template flowset record.
     */
    templateTxCounter?: pulumi.Input<number>;
    /**
     * Timeout for periodic template flowset transmission. On FortiOS versions 6.2.0-7.0.0: 1 - 1440 min, default = 30. On FortiOS versions >= 7.0.1: 60 - 86400 sec, default = 1800.
     */
    templateTxTimeout?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
