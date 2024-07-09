// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure IPv4 virtual IP groups.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname1 = new fortios.firewall.Vip("trname1", {
 *     extintf: "any",
 *     extport: "0-65535",
 *     extip: "2.0.0.1-2.0.0.4",
 *     mappedips: [{
 *         range: "3.0.0.0-3.0.0.3",
 *     }],
 * });
 * const trname = new fortios.firewall.Vipgrp("trname", {
 *     color: 0,
 *     "interface": "any",
 *     members: [{
 *         name: trname1.name,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Vipgrp can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/vipgrp:Vipgrp labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/vipgrp:Vipgrp labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Vipgrp extends pulumi.CustomResource {
    /**
     * Get an existing Vipgrp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VipgrpState, opts?: pulumi.CustomResourceOptions): Vipgrp {
        return new Vipgrp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/vipgrp:Vipgrp';

    /**
     * Returns true if the given object is an instance of Vipgrp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vipgrp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vipgrp.__pulumiType;
    }

    /**
     * Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
     */
    public readonly color!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * interface
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
     */
    public readonly members!: pulumi.Output<outputs.firewall.VipgrpMember[]>;
    /**
     * VIP group name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    public readonly uuid!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Vipgrp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VipgrpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VipgrpArgs | VipgrpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VipgrpState | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as VipgrpArgs | undefined;
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vipgrp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vipgrp resources.
 */
export interface VipgrpState {
    /**
     * Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * interface
     */
    interface?: pulumi.Input<string>;
    /**
     * Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.firewall.VipgrpMember>[]>;
    /**
     * VIP group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Vipgrp resource.
 */
export interface VipgrpArgs {
    /**
     * Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * interface
     */
    interface: pulumi.Input<string>;
    /**
     * Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
     */
    members: pulumi.Input<pulumi.Input<inputs.firewall.VipgrpMember>[]>;
    /**
     * VIP group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
