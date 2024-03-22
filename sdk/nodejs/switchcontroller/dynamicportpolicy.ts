// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device. Applies to FortiOS Version `>= 7.0.1`.
 *
 * ## Import
 *
 * SwitchController DynamicPortPolicy can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/dynamicportpolicy:Dynamicportpolicy labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/dynamicportpolicy:Dynamicportpolicy labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Dynamicportpolicy extends pulumi.CustomResource {
    /**
     * Get an existing Dynamicportpolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DynamicportpolicyState, opts?: pulumi.CustomResourceOptions): Dynamicportpolicy {
        return new Dynamicportpolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/dynamicportpolicy:Dynamicportpolicy';

    /**
     * Returns true if the given object is an instance of Dynamicportpolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dynamicportpolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dynamicportpolicy.__pulumiType;
    }

    /**
     * Description for the Dynamic port policy.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * FortiLink interface for which this Dynamic port policy belongs to.
     */
    public readonly fortilink!: pulumi.Output<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Dynamic port policy name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Port policies with matching criteria and actions. The structure of `policy` block is documented below.
     */
    public readonly policies!: pulumi.Output<outputs.switchcontroller.DynamicportpolicyPolicy[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Dynamicportpolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DynamicportpolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DynamicportpolicyArgs | DynamicportpolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DynamicportpolicyState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fortilink"] = state ? state.fortilink : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as DynamicportpolicyArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fortilink"] = args ? args.fortilink : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Dynamicportpolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dynamicportpolicy resources.
 */
export interface DynamicportpolicyState {
    /**
     * Description for the Dynamic port policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * FortiLink interface for which this Dynamic port policy belongs to.
     */
    fortilink?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Dynamic port policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Port policies with matching criteria and actions. The structure of `policy` block is documented below.
     */
    policies?: pulumi.Input<pulumi.Input<inputs.switchcontroller.DynamicportpolicyPolicy>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Dynamicportpolicy resource.
 */
export interface DynamicportpolicyArgs {
    /**
     * Description for the Dynamic port policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * FortiLink interface for which this Dynamic port policy belongs to.
     */
    fortilink?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Dynamic port policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Port policies with matching criteria and actions. The structure of `policy` block is documented below.
     */
    policies?: pulumi.Input<pulumi.Input<inputs.switchcontroller.DynamicportpolicyPolicy>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
