// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure firewall application groups.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.application.Group("trname", {
 *     categories: [{
 *         id: 2,
 *     }],
 *     comment: "group1 test",
 *     type: "category",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Application Group can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:application/group:Group labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:application/group:Group labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:application/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * Application ID list. The structure of `application` block is documented below.
     */
    public readonly applications!: pulumi.Output<outputs.application.GroupApplication[] | undefined>;
    /**
     * Application behavior filter.
     */
    public readonly behavior!: pulumi.Output<string>;
    /**
     * Application category ID list. The structure of `category` block is documented below.
     */
    public readonly categories!: pulumi.Output<outputs.application.GroupCategory[] | undefined>;
    /**
     * Comment
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Application group name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
     */
    public readonly popularity!: pulumi.Output<string>;
    /**
     * Application protocol filter.
     */
    public readonly protocols!: pulumi.Output<string>;
    /**
     * Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
     */
    public readonly risks!: pulumi.Output<outputs.application.GroupRisk[] | undefined>;
    /**
     * Application technology filter.
     */
    public readonly technology!: pulumi.Output<string>;
    /**
     * Application group type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Application vendor filter.
     */
    public readonly vendor!: pulumi.Output<string>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupState | undefined;
            resourceInputs["applications"] = state ? state.applications : undefined;
            resourceInputs["behavior"] = state ? state.behavior : undefined;
            resourceInputs["categories"] = state ? state.categories : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["popularity"] = state ? state.popularity : undefined;
            resourceInputs["protocols"] = state ? state.protocols : undefined;
            resourceInputs["risks"] = state ? state.risks : undefined;
            resourceInputs["technology"] = state ? state.technology : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vendor"] = state ? state.vendor : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            resourceInputs["applications"] = args ? args.applications : undefined;
            resourceInputs["behavior"] = args ? args.behavior : undefined;
            resourceInputs["categories"] = args ? args.categories : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["popularity"] = args ? args.popularity : undefined;
            resourceInputs["protocols"] = args ? args.protocols : undefined;
            resourceInputs["risks"] = args ? args.risks : undefined;
            resourceInputs["technology"] = args ? args.technology : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vendor"] = args ? args.vendor : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * Application ID list. The structure of `application` block is documented below.
     */
    applications?: pulumi.Input<pulumi.Input<inputs.application.GroupApplication>[]>;
    /**
     * Application behavior filter.
     */
    behavior?: pulumi.Input<string>;
    /**
     * Application category ID list. The structure of `category` block is documented below.
     */
    categories?: pulumi.Input<pulumi.Input<inputs.application.GroupCategory>[]>;
    /**
     * Comment
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Application group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
     */
    popularity?: pulumi.Input<string>;
    /**
     * Application protocol filter.
     */
    protocols?: pulumi.Input<string>;
    /**
     * Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
     */
    risks?: pulumi.Input<pulumi.Input<inputs.application.GroupRisk>[]>;
    /**
     * Application technology filter.
     */
    technology?: pulumi.Input<string>;
    /**
     * Application group type.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Application vendor filter.
     */
    vendor?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * Application ID list. The structure of `application` block is documented below.
     */
    applications?: pulumi.Input<pulumi.Input<inputs.application.GroupApplication>[]>;
    /**
     * Application behavior filter.
     */
    behavior?: pulumi.Input<string>;
    /**
     * Application category ID list. The structure of `category` block is documented below.
     */
    categories?: pulumi.Input<pulumi.Input<inputs.application.GroupCategory>[]>;
    /**
     * Comment
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Application group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
     */
    popularity?: pulumi.Input<string>;
    /**
     * Application protocol filter.
     */
    protocols?: pulumi.Input<string>;
    /**
     * Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
     */
    risks?: pulumi.Input<pulumi.Input<inputs.application.GroupRisk>[]>;
    /**
     * Application technology filter.
     */
    technology?: pulumi.Input<string>;
    /**
     * Application group type.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Application vendor filter.
     */
    vendor?: pulumi.Input<string>;
}
