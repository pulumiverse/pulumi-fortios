// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure Internet Services Extension.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Internetserviceextension("trname", {
 *     comment: "EIWE",
 *     fosid: 65536,
 * });
 * ```
 *
 * ## Import
 *
 * Firewall InternetServiceExtension can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/internetserviceextension:Internetserviceextension labelname {{fosid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/internetserviceextension:Internetserviceextension labelname {{fosid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Internetserviceextension extends pulumi.CustomResource {
    /**
     * Get an existing Internetserviceextension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InternetserviceextensionState, opts?: pulumi.CustomResourceOptions): Internetserviceextension {
        return new Internetserviceextension(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/internetserviceextension:Internetserviceextension';

    /**
     * Returns true if the given object is an instance of Internetserviceextension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Internetserviceextension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Internetserviceextension.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Disable entries in the Internet Service database. The structure of `disableEntry` block is documented below.
     */
    public readonly disableEntries!: pulumi.Output<outputs.firewall.InternetserviceextensionDisableEntry[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
     */
    public readonly entries!: pulumi.Output<outputs.firewall.InternetserviceextensionEntry[] | undefined>;
    /**
     * Internet Service ID in the Internet Service database.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Internetserviceextension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InternetserviceextensionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InternetserviceextensionArgs | InternetserviceextensionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InternetserviceextensionState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["disableEntries"] = state ? state.disableEntries : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = state ? state.entries : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as InternetserviceextensionArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["disableEntries"] = args ? args.disableEntries : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = args ? args.entries : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Internetserviceextension.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Internetserviceextension resources.
 */
export interface InternetserviceextensionState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Disable entries in the Internet Service database. The structure of `disableEntry` block is documented below.
     */
    disableEntries?: pulumi.Input<pulumi.Input<inputs.firewall.InternetserviceextensionDisableEntry>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.firewall.InternetserviceextensionEntry>[]>;
    /**
     * Internet Service ID in the Internet Service database.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Internetserviceextension resource.
 */
export interface InternetserviceextensionArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Disable entries in the Internet Service database. The structure of `disableEntry` block is documented below.
     */
    disableEntries?: pulumi.Input<pulumi.Input<inputs.firewall.InternetserviceextensionDisableEntry>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.firewall.InternetserviceextensionEntry>[]>;
    /**
     * Internet Service ID in the Internet Service database.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
