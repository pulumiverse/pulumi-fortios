// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Configure YouTube channel filter. Applies to FortiOS Version `>= 7.0.1`.
 *
 * ## Import
 *
 * Videofilter YoutubeChannelFilter can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:filter/video/youtubechannelfilter:Youtubechannelfilter labelname {{fosid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:filter/video/youtubechannelfilter:Youtubechannelfilter labelname {{fosid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Youtubechannelfilter extends pulumi.CustomResource {
    /**
     * Get an existing Youtubechannelfilter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: YoutubechannelfilterState, opts?: pulumi.CustomResourceOptions): Youtubechannelfilter {
        return new Youtubechannelfilter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:filter/video/youtubechannelfilter:Youtubechannelfilter';

    /**
     * Returns true if the given object is an instance of Youtubechannelfilter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Youtubechannelfilter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Youtubechannelfilter.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * YouTube channel filter default action. Valid values: `allow`, `monitor`, `block`.
     */
    public readonly defaultAction!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * YouTube filter entries. The structure of `entries` block is documented below.
     */
    public readonly entries!: pulumi.Output<outputs.filter.video.YoutubechannelfilterEntry[] | undefined>;
    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Eanble/disable logging. Valid values: `enable`, `disable`.
     */
    public readonly log!: pulumi.Output<string>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable overriding category filtering result. Valid values: `enable`, `disable`.
     */
    public readonly overrideCategory!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Youtubechannelfilter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: YoutubechannelfilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: YoutubechannelfilterArgs | YoutubechannelfilterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as YoutubechannelfilterState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["defaultAction"] = state ? state.defaultAction : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = state ? state.entries : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["log"] = state ? state.log : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["overrideCategory"] = state ? state.overrideCategory : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as YoutubechannelfilterArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["defaultAction"] = args ? args.defaultAction : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = args ? args.entries : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["log"] = args ? args.log : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overrideCategory"] = args ? args.overrideCategory : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Youtubechannelfilter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Youtubechannelfilter resources.
 */
export interface YoutubechannelfilterState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * YouTube channel filter default action. Valid values: `allow`, `monitor`, `block`.
     */
    defaultAction?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * YouTube filter entries. The structure of `entries` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.filter.video.YoutubechannelfilterEntry>[]>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Eanble/disable logging. Valid values: `enable`, `disable`.
     */
    log?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable overriding category filtering result. Valid values: `enable`, `disable`.
     */
    overrideCategory?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Youtubechannelfilter resource.
 */
export interface YoutubechannelfilterArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * YouTube channel filter default action. Valid values: `allow`, `monitor`, `block`.
     */
    defaultAction?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * YouTube filter entries. The structure of `entries` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.filter.video.YoutubechannelfilterEntry>[]>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Eanble/disable logging. Valid values: `enable`, `disable`.
     */
    log?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable overriding category filtering result. Valid values: `enable`, `disable`.
     */
    overrideCategory?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
