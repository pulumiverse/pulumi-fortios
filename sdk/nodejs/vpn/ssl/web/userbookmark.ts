// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../../types/input";
import * as outputs from "../../../types/output";
import * as utilities from "../../../utilities";

/**
 * Configure SSL VPN user bookmark.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.vpn.ssl.web.Userbookmark("trname", {customLang: "big5"});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * VpnSslWeb UserBookmark can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:vpn/ssl/web/userbookmark:Userbookmark labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:vpn/ssl/web/userbookmark:Userbookmark labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Userbookmark extends pulumi.CustomResource {
    /**
     * Get an existing Userbookmark resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserbookmarkState, opts?: pulumi.CustomResourceOptions): Userbookmark {
        return new Userbookmark(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:vpn/ssl/web/userbookmark:Userbookmark';

    /**
     * Returns true if the given object is an instance of Userbookmark.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Userbookmark {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Userbookmark.__pulumiType;
    }

    /**
     * Bookmark table. The structure of `bookmarks` block is documented below.
     */
    public readonly bookmarks!: pulumi.Output<outputs.vpn.ssl.web.UserbookmarkBookmark[] | undefined>;
    /**
     * Personal language.
     */
    public readonly customLang!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * User and group name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Userbookmark resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserbookmarkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserbookmarkArgs | UserbookmarkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserbookmarkState | undefined;
            resourceInputs["bookmarks"] = state ? state.bookmarks : undefined;
            resourceInputs["customLang"] = state ? state.customLang : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as UserbookmarkArgs | undefined;
            resourceInputs["bookmarks"] = args ? args.bookmarks : undefined;
            resourceInputs["customLang"] = args ? args.customLang : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Userbookmark.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Userbookmark resources.
 */
export interface UserbookmarkState {
    /**
     * Bookmark table. The structure of `bookmarks` block is documented below.
     */
    bookmarks?: pulumi.Input<pulumi.Input<inputs.vpn.ssl.web.UserbookmarkBookmark>[]>;
    /**
     * Personal language.
     */
    customLang?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * User and group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Userbookmark resource.
 */
export interface UserbookmarkArgs {
    /**
     * Bookmark table. The structure of `bookmarks` block is documented below.
     */
    bookmarks?: pulumi.Input<pulumi.Input<inputs.vpn.ssl.web.UserbookmarkBookmark>[]>;
    /**
     * Personal language.
     */
    customLang?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * User and group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
