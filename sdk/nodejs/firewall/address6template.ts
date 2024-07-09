// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure IPv6 address templates.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Address6template("trname", {
 *     ip6: "2001:db8:0:b::/64",
 *     subnetSegments: [
 *         {
 *             bits: 4,
 *             exclusive: "disable",
 *             id: 1,
 *             name: "country",
 *         },
 *         {
 *             bits: 4,
 *             exclusive: "disable",
 *             id: 2,
 *             name: "state",
 *         },
 *     ],
 *     subnetSegmentCount: 2,
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Address6Template can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/address6template:Address6template labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/address6template:Address6template labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Address6template extends pulumi.CustomResource {
    /**
     * Get an existing Address6template resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Address6templateState, opts?: pulumi.CustomResourceOptions): Address6template {
        return new Address6template(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/address6template:Address6template';

    /**
     * Returns true if the given object is an instance of Address6template.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Address6template {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Address6template.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Security Fabric global object setting. Valid values: `enable`, `disable`.
     */
    public readonly fabricObject!: pulumi.Output<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * IPv6 address prefix.
     */
    public readonly ip6!: pulumi.Output<string>;
    /**
     * IPv6 address template name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of IPv6 subnet segments.
     */
    public readonly subnetSegmentCount!: pulumi.Output<number>;
    /**
     * IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
     */
    public readonly subnetSegments!: pulumi.Output<outputs.firewall.Address6templateSubnetSegment[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Address6template resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Address6templateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Address6templateArgs | Address6templateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Address6templateState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fabricObject"] = state ? state.fabricObject : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["ip6"] = state ? state.ip6 : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["subnetSegmentCount"] = state ? state.subnetSegmentCount : undefined;
            resourceInputs["subnetSegments"] = state ? state.subnetSegments : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Address6templateArgs | undefined;
            if ((!args || args.ip6 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip6'");
            }
            if ((!args || args.subnetSegmentCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetSegmentCount'");
            }
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fabricObject"] = args ? args.fabricObject : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["ip6"] = args ? args.ip6 : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["subnetSegmentCount"] = args ? args.subnetSegmentCount : undefined;
            resourceInputs["subnetSegments"] = args ? args.subnetSegments : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Address6template.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Address6template resources.
 */
export interface Address6templateState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Security Fabric global object setting. Valid values: `enable`, `disable`.
     */
    fabricObject?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * IPv6 address prefix.
     */
    ip6?: pulumi.Input<string>;
    /**
     * IPv6 address template name.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of IPv6 subnet segments.
     */
    subnetSegmentCount?: pulumi.Input<number>;
    /**
     * IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
     */
    subnetSegments?: pulumi.Input<pulumi.Input<inputs.firewall.Address6templateSubnetSegment>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Address6template resource.
 */
export interface Address6templateArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Security Fabric global object setting. Valid values: `enable`, `disable`.
     */
    fabricObject?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * IPv6 address prefix.
     */
    ip6: pulumi.Input<string>;
    /**
     * IPv6 address template name.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of IPv6 subnet segments.
     */
    subnetSegmentCount: pulumi.Input<number>;
    /**
     * IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
     */
    subnetSegments?: pulumi.Input<pulumi.Input<inputs.firewall.Address6templateSubnetSegment>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
