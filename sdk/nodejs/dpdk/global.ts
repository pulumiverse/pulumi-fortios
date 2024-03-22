// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure global DPDK options. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * Dpdk Global can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:dpdk/global:Global labelname DpdkGlobal
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:dpdk/global:Global labelname DpdkGlobal
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Global extends pulumi.CustomResource {
    /**
     * Get an existing Global resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalState, opts?: pulumi.CustomResourceOptions): Global {
        return new Global(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:dpdk/global:Global';

    /**
     * Returns true if the given object is an instance of Global.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Global {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Global.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
     */
    public readonly elasticbuffer!: pulumi.Output<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Percentage of main memory allocated to hugepages, which are available for DPDK operation.
     */
    public readonly hugepagePercentage!: pulumi.Output<number>;
    /**
     * Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
     */
    public readonly interfaces!: pulumi.Output<outputs.dpdk.GlobalInterface[] | undefined>;
    /**
     * Enable/disable DPDK IPsec phase 2 offloading. Valid values: `disable`, `enable`.
     */
    public readonly ipsecOffload!: pulumi.Output<string>;
    /**
     * Percentage of main memory allocated to DPDK packet buffer.
     */
    public readonly mbufpoolPercentage!: pulumi.Output<number>;
    /**
     * Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
     */
    public readonly multiqueue!: pulumi.Output<string>;
    /**
     * Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
     */
    public readonly perSessionAccounting!: pulumi.Output<string>;
    /**
     * Special arguments for device
     */
    public readonly protects!: pulumi.Output<string>;
    /**
     * Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
     */
    public readonly sleepOnIdle!: pulumi.Output<string>;
    /**
     * Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Global resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GlobalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GlobalArgs | GlobalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GlobalState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["elasticbuffer"] = state ? state.elasticbuffer : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["hugepagePercentage"] = state ? state.hugepagePercentage : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["ipsecOffload"] = state ? state.ipsecOffload : undefined;
            resourceInputs["mbufpoolPercentage"] = state ? state.mbufpoolPercentage : undefined;
            resourceInputs["multiqueue"] = state ? state.multiqueue : undefined;
            resourceInputs["perSessionAccounting"] = state ? state.perSessionAccounting : undefined;
            resourceInputs["protects"] = state ? state.protects : undefined;
            resourceInputs["sleepOnIdle"] = state ? state.sleepOnIdle : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as GlobalArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["elasticbuffer"] = args ? args.elasticbuffer : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["hugepagePercentage"] = args ? args.hugepagePercentage : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["ipsecOffload"] = args ? args.ipsecOffload : undefined;
            resourceInputs["mbufpoolPercentage"] = args ? args.mbufpoolPercentage : undefined;
            resourceInputs["multiqueue"] = args ? args.multiqueue : undefined;
            resourceInputs["perSessionAccounting"] = args ? args.perSessionAccounting : undefined;
            resourceInputs["protects"] = args ? args.protects : undefined;
            resourceInputs["sleepOnIdle"] = args ? args.sleepOnIdle : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Global.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Global resources.
 */
export interface GlobalState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
     */
    elasticbuffer?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Percentage of main memory allocated to hugepages, which are available for DPDK operation.
     */
    hugepagePercentage?: pulumi.Input<number>;
    /**
     * Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.dpdk.GlobalInterface>[]>;
    /**
     * Enable/disable DPDK IPsec phase 2 offloading. Valid values: `disable`, `enable`.
     */
    ipsecOffload?: pulumi.Input<string>;
    /**
     * Percentage of main memory allocated to DPDK packet buffer.
     */
    mbufpoolPercentage?: pulumi.Input<number>;
    /**
     * Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
     */
    multiqueue?: pulumi.Input<string>;
    /**
     * Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
     */
    perSessionAccounting?: pulumi.Input<string>;
    /**
     * Special arguments for device
     */
    protects?: pulumi.Input<string>;
    /**
     * Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
     */
    sleepOnIdle?: pulumi.Input<string>;
    /**
     * Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Global resource.
 */
export interface GlobalArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
     */
    elasticbuffer?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Percentage of main memory allocated to hugepages, which are available for DPDK operation.
     */
    hugepagePercentage?: pulumi.Input<number>;
    /**
     * Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.dpdk.GlobalInterface>[]>;
    /**
     * Enable/disable DPDK IPsec phase 2 offloading. Valid values: `disable`, `enable`.
     */
    ipsecOffload?: pulumi.Input<string>;
    /**
     * Percentage of main memory allocated to DPDK packet buffer.
     */
    mbufpoolPercentage?: pulumi.Input<number>;
    /**
     * Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
     */
    multiqueue?: pulumi.Input<string>;
    /**
     * Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
     */
    perSessionAccounting?: pulumi.Input<string>;
    /**
     * Special arguments for device
     */
    protects?: pulumi.Input<string>;
    /**
     * Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
     */
    sleepOnIdle?: pulumi.Input<string>;
    /**
     * Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
