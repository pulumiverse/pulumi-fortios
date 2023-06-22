// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure URL filter lists.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.WebfilterUrlfilter("trname", {
 *     fosid: 1,
 *     ipAddrBlock: "enable",
 *     oneArmIpsUrlfilter: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * Webfilter Urlfilter can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/webfilterUrlfilter:WebfilterUrlfilter labelname {{fosid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/webfilterUrlfilter:WebfilterUrlfilter labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WebfilterUrlfilter extends pulumi.CustomResource {
    /**
     * Get an existing WebfilterUrlfilter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebfilterUrlfilterState, opts?: pulumi.CustomResourceOptions): WebfilterUrlfilter {
        return new WebfilterUrlfilter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/webfilterUrlfilter:WebfilterUrlfilter';

    /**
     * Returns true if the given object is an instance of WebfilterUrlfilter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebfilterUrlfilter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebfilterUrlfilter.__pulumiType;
    }

    /**
     * Optional comments.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * URL filter entries. The structure of `entries` block is documented below.
     */
    public readonly entries!: pulumi.Output<outputs.WebfilterUrlfilterEntry[] | undefined>;
    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
     */
    public readonly ipAddrBlock!: pulumi.Output<string>;
    /**
     * Name of URL filter list.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
     */
    public readonly oneArmIpsUrlfilter!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WebfilterUrlfilter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebfilterUrlfilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebfilterUrlfilterArgs | WebfilterUrlfilterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebfilterUrlfilterState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = state ? state.entries : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["ipAddrBlock"] = state ? state.ipAddrBlock : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["oneArmIpsUrlfilter"] = state ? state.oneArmIpsUrlfilter : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WebfilterUrlfilterArgs | undefined;
            if ((!args || args.fosid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fosid'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = args ? args.entries : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["ipAddrBlock"] = args ? args.ipAddrBlock : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oneArmIpsUrlfilter"] = args ? args.oneArmIpsUrlfilter : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WebfilterUrlfilter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebfilterUrlfilter resources.
 */
export interface WebfilterUrlfilterState {
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * URL filter entries. The structure of `entries` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.WebfilterUrlfilterEntry>[]>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
     */
    ipAddrBlock?: pulumi.Input<string>;
    /**
     * Name of URL filter list.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
     */
    oneArmIpsUrlfilter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebfilterUrlfilter resource.
 */
export interface WebfilterUrlfilterArgs {
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * URL filter entries. The structure of `entries` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.WebfilterUrlfilterEntry>[]>;
    /**
     * ID.
     */
    fosid: pulumi.Input<number>;
    /**
     * Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
     */
    ipAddrBlock?: pulumi.Input<string>;
    /**
     * Name of URL filter list.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
     */
    oneArmIpsUrlfilter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
