// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Web application firewall configuration.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.waf.Profile("trname", {
 *     extendedLog: "disable",
 *     external: "disable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Waf Profile can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:waf/profile:Profile labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:waf/profile:Profile labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Profile extends pulumi.CustomResource {
    /**
     * Get an existing Profile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileState, opts?: pulumi.CustomResourceOptions): Profile {
        return new Profile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:waf/profile:Profile';

    /**
     * Returns true if the given object is an instance of Profile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Profile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Profile.__pulumiType;
    }

    /**
     * Black address list and white address list. The structure of `addressList` block is documented below.
     */
    public readonly addressList!: pulumi.Output<outputs.waf.ProfileAddressList>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * WAF HTTP protocol restrictions. The structure of `constraint` block is documented below.
     */
    public readonly constraint!: pulumi.Output<outputs.waf.ProfileConstraint>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable extended logging. Valid values: `enable`, `disable`.
     */
    public readonly extendedLog!: pulumi.Output<string>;
    /**
     * Disable/Enable external HTTP Inspection. Valid values: `disable`, `enable`.
     */
    public readonly external!: pulumi.Output<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Method restriction. The structure of `method` block is documented below.
     */
    public readonly method!: pulumi.Output<outputs.waf.ProfileMethod>;
    /**
     * WAF Profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * WAF signatures. The structure of `signature` block is documented below.
     */
    public readonly signature!: pulumi.Output<outputs.waf.ProfileSignature>;
    /**
     * URL access list The structure of `urlAccess` block is documented below.
     */
    public readonly urlAccesses!: pulumi.Output<outputs.waf.ProfileUrlAccess[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Profile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileArgs | ProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProfileState | undefined;
            resourceInputs["addressList"] = state ? state.addressList : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["constraint"] = state ? state.constraint : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["extendedLog"] = state ? state.extendedLog : undefined;
            resourceInputs["external"] = state ? state.external : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["method"] = state ? state.method : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["signature"] = state ? state.signature : undefined;
            resourceInputs["urlAccesses"] = state ? state.urlAccesses : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ProfileArgs | undefined;
            resourceInputs["addressList"] = args ? args.addressList : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["constraint"] = args ? args.constraint : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["extendedLog"] = args ? args.extendedLog : undefined;
            resourceInputs["external"] = args ? args.external : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["method"] = args ? args.method : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["signature"] = args ? args.signature : undefined;
            resourceInputs["urlAccesses"] = args ? args.urlAccesses : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Profile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Profile resources.
 */
export interface ProfileState {
    /**
     * Black address list and white address list. The structure of `addressList` block is documented below.
     */
    addressList?: pulumi.Input<inputs.waf.ProfileAddressList>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * WAF HTTP protocol restrictions. The structure of `constraint` block is documented below.
     */
    constraint?: pulumi.Input<inputs.waf.ProfileConstraint>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable extended logging. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * Disable/Enable external HTTP Inspection. Valid values: `disable`, `enable`.
     */
    external?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Method restriction. The structure of `method` block is documented below.
     */
    method?: pulumi.Input<inputs.waf.ProfileMethod>;
    /**
     * WAF Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * WAF signatures. The structure of `signature` block is documented below.
     */
    signature?: pulumi.Input<inputs.waf.ProfileSignature>;
    /**
     * URL access list The structure of `urlAccess` block is documented below.
     */
    urlAccesses?: pulumi.Input<pulumi.Input<inputs.waf.ProfileUrlAccess>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Profile resource.
 */
export interface ProfileArgs {
    /**
     * Black address list and white address list. The structure of `addressList` block is documented below.
     */
    addressList?: pulumi.Input<inputs.waf.ProfileAddressList>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * WAF HTTP protocol restrictions. The structure of `constraint` block is documented below.
     */
    constraint?: pulumi.Input<inputs.waf.ProfileConstraint>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable extended logging. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * Disable/Enable external HTTP Inspection. Valid values: `disable`, `enable`.
     */
    external?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Method restriction. The structure of `method` block is documented below.
     */
    method?: pulumi.Input<inputs.waf.ProfileMethod>;
    /**
     * WAF Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * WAF signatures. The structure of `signature` block is documented below.
     */
    signature?: pulumi.Input<inputs.waf.ProfileSignature>;
    /**
     * URL access list The structure of `urlAccess` block is documented below.
     */
    urlAccesses?: pulumi.Input<pulumi.Input<inputs.waf.ProfileUrlAccess>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
