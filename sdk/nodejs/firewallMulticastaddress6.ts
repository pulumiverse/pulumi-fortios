// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure IPv6 multicast address.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.FirewallMulticastaddress6("trname", {
 *     color: 0,
 *     ip6: "ff02::1:ff0e:8c6c/128",
 *     visibility: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall MulticastAddress6 can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallMulticastaddress6:FirewallMulticastaddress6 labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallMulticastaddress6:FirewallMulticastaddress6 labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallMulticastaddress6 extends pulumi.CustomResource {
    /**
     * Get an existing FirewallMulticastaddress6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallMulticastaddress6State, opts?: pulumi.CustomResourceOptions): FirewallMulticastaddress6 {
        return new FirewallMulticastaddress6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallMulticastaddress6:FirewallMulticastaddress6';

    /**
     * Returns true if the given object is an instance of FirewallMulticastaddress6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallMulticastaddress6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallMulticastaddress6.__pulumiType;
    }

    /**
     * Color of icon on the GUI.
     */
    public readonly color!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
     */
    public readonly ip6!: pulumi.Output<string>;
    /**
     * IPv6 multicast address name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    public readonly taggings!: pulumi.Output<outputs.FirewallMulticastaddress6Tagging[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
     */
    public readonly visibility!: pulumi.Output<string>;

    /**
     * Create a FirewallMulticastaddress6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallMulticastaddress6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallMulticastaddress6Args | FirewallMulticastaddress6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallMulticastaddress6State | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["ip6"] = state ? state.ip6 : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["taggings"] = state ? state.taggings : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as FirewallMulticastaddress6Args | undefined;
            if ((!args || args.ip6 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip6'");
            }
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["ip6"] = args ? args.ip6 : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["taggings"] = args ? args.taggings : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallMulticastaddress6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallMulticastaddress6 resources.
 */
export interface FirewallMulticastaddress6State {
    /**
     * Color of icon on the GUI.
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
     */
    ip6?: pulumi.Input<string>;
    /**
     * IPv6 multicast address name.
     */
    name?: pulumi.Input<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.FirewallMulticastaddress6Tagging>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
     */
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallMulticastaddress6 resource.
 */
export interface FirewallMulticastaddress6Args {
    /**
     * Color of icon on the GUI.
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
     */
    ip6: pulumi.Input<string>;
    /**
     * IPv6 multicast address name.
     */
    name?: pulumi.Input<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.FirewallMulticastaddress6Tagging>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
     */
    visibility?: pulumi.Input<string>;
}