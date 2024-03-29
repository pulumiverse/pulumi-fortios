// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure application signatures.
 *
 * ## Import
 *
 * Application Name can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:application/name:Name labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:application/name:Name labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Name extends pulumi.CustomResource {
    /**
     * Get an existing Name resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NameState, opts?: pulumi.CustomResourceOptions): Name {
        return new Name(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:application/name:Name';

    /**
     * Returns true if the given object is an instance of Name.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Name {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Name.__pulumiType;
    }

    /**
     * Application behavior.
     */
    public readonly behavior!: pulumi.Output<string>;
    /**
     * Application category ID.
     */
    public readonly category!: pulumi.Output<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Application ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Meta data. The structure of `metadata` block is documented below.
     */
    public readonly metadatas!: pulumi.Output<outputs.application.NameMetadata[] | undefined>;
    /**
     * Application name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Application parameter name.
     */
    public readonly parameter!: pulumi.Output<string>;
    /**
     * Application parameters. The structure of `parameters` block is documented below.
     */
    public readonly parameters!: pulumi.Output<outputs.application.NameParameter[] | undefined>;
    /**
     * Application popularity.
     */
    public readonly popularity!: pulumi.Output<number>;
    /**
     * Application protocol.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Application risk.
     */
    public readonly risk!: pulumi.Output<number>;
    /**
     * Application sub-category ID.
     */
    public readonly subCategory!: pulumi.Output<number>;
    /**
     * Application technology.
     */
    public readonly technology!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Application vendor.
     */
    public readonly vendor!: pulumi.Output<string>;
    /**
     * Application weight.
     */
    public readonly weight!: pulumi.Output<number>;

    /**
     * Create a Name resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NameArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NameArgs | NameState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NameState | undefined;
            resourceInputs["behavior"] = state ? state.behavior : undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["metadatas"] = state ? state.metadatas : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parameter"] = state ? state.parameter : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["popularity"] = state ? state.popularity : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["risk"] = state ? state.risk : undefined;
            resourceInputs["subCategory"] = state ? state.subCategory : undefined;
            resourceInputs["technology"] = state ? state.technology : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vendor"] = state ? state.vendor : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as NameArgs | undefined;
            if ((!args || args.category === undefined) && !opts.urn) {
                throw new Error("Missing required property 'category'");
            }
            resourceInputs["behavior"] = args ? args.behavior : undefined;
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["metadatas"] = args ? args.metadatas : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameter"] = args ? args.parameter : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["popularity"] = args ? args.popularity : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["risk"] = args ? args.risk : undefined;
            resourceInputs["subCategory"] = args ? args.subCategory : undefined;
            resourceInputs["technology"] = args ? args.technology : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vendor"] = args ? args.vendor : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Name.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Name resources.
 */
export interface NameState {
    /**
     * Application behavior.
     */
    behavior?: pulumi.Input<string>;
    /**
     * Application category ID.
     */
    category?: pulumi.Input<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Application ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Meta data. The structure of `metadata` block is documented below.
     */
    metadatas?: pulumi.Input<pulumi.Input<inputs.application.NameMetadata>[]>;
    /**
     * Application name.
     */
    name?: pulumi.Input<string>;
    /**
     * Application parameter name.
     */
    parameter?: pulumi.Input<string>;
    /**
     * Application parameters. The structure of `parameters` block is documented below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.application.NameParameter>[]>;
    /**
     * Application popularity.
     */
    popularity?: pulumi.Input<number>;
    /**
     * Application protocol.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Application risk.
     */
    risk?: pulumi.Input<number>;
    /**
     * Application sub-category ID.
     */
    subCategory?: pulumi.Input<number>;
    /**
     * Application technology.
     */
    technology?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Application vendor.
     */
    vendor?: pulumi.Input<string>;
    /**
     * Application weight.
     */
    weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Name resource.
 */
export interface NameArgs {
    /**
     * Application behavior.
     */
    behavior?: pulumi.Input<string>;
    /**
     * Application category ID.
     */
    category: pulumi.Input<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Application ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Meta data. The structure of `metadata` block is documented below.
     */
    metadatas?: pulumi.Input<pulumi.Input<inputs.application.NameMetadata>[]>;
    /**
     * Application name.
     */
    name?: pulumi.Input<string>;
    /**
     * Application parameter name.
     */
    parameter?: pulumi.Input<string>;
    /**
     * Application parameters. The structure of `parameters` block is documented below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.application.NameParameter>[]>;
    /**
     * Application popularity.
     */
    popularity?: pulumi.Input<number>;
    /**
     * Application protocol.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Application risk.
     */
    risk?: pulumi.Input<number>;
    /**
     * Application sub-category ID.
     */
    subCategory?: pulumi.Input<number>;
    /**
     * Application technology.
     */
    technology?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Application vendor.
     */
    vendor?: pulumi.Input<string>;
    /**
     * Application weight.
     */
    weight?: pulumi.Input<number>;
}
