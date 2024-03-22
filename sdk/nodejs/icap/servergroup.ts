// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure an ICAP server group consisting of multiple forward servers. Supports failover and load balancing. Applies to FortiOS Version `>= 7.2.0`.
 *
 * ## Import
 *
 * Icap ServerGroup can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:icap/servergroup:Servergroup labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:icap/servergroup:Servergroup labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Servergroup extends pulumi.CustomResource {
    /**
     * Get an existing Servergroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServergroupState, opts?: pulumi.CustomResourceOptions): Servergroup {
        return new Servergroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:icap/servergroup:Servergroup';

    /**
     * Returns true if the given object is an instance of Servergroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Servergroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Servergroup.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
     */
    public readonly ldbMethod!: pulumi.Output<string>;
    /**
     * Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
     */
    public readonly serverLists!: pulumi.Output<outputs.icap.ServergroupServerList[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Servergroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServergroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServergroupArgs | ServergroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServergroupState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["ldbMethod"] = state ? state.ldbMethod : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["serverLists"] = state ? state.serverLists : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ServergroupArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["ldbMethod"] = args ? args.ldbMethod : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serverLists"] = args ? args.serverLists : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Servergroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Servergroup resources.
 */
export interface ServergroupState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
     */
    ldbMethod?: pulumi.Input<string>;
    /**
     * Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
     */
    name?: pulumi.Input<string>;
    /**
     * Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
     */
    serverLists?: pulumi.Input<pulumi.Input<inputs.icap.ServergroupServerList>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Servergroup resource.
 */
export interface ServergroupArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
     */
    ldbMethod?: pulumi.Input<string>;
    /**
     * Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
     */
    name?: pulumi.Input<string>;
    /**
     * Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
     */
    serverLists?: pulumi.Input<pulumi.Input<inputs.icap.ServergroupServerList>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
