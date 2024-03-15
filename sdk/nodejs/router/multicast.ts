// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure router multicast.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.router.Multicast("trname", {
 *     multicastRouting: "disable",
 *     pimSmGlobal: {
 *         bsrAllowQuickRefresh: "disable",
 *         bsrCandidate: "disable",
 *         bsrHash: 10,
 *         bsrPriority: 0,
 *         ciscoCrpPrefix: "disable",
 *         ciscoIgnoreRpSetPriority: "disable",
 *         ciscoRegisterChecksum: "disable",
 *         joinPruneHoldtime: 210,
 *         messageInterval: 60,
 *         nullRegisterRetries: 1,
 *         registerRateLimit: 0,
 *         registerRpReachability: "enable",
 *         registerSource: "disable",
 *         registerSourceIp: "0.0.0.0",
 *         registerSupression: 60,
 *         rpRegisterKeepalive: 185,
 *         sptThreshold: "enable",
 *         ssm: "disable",
 *     },
 *     routeLimit: 2147483647,
 *     routeThreshold: 2147483647,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Router Multicast can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:router/multicast:Multicast labelname RouterMulticast
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:router/multicast:Multicast labelname RouterMulticast
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Multicast extends pulumi.CustomResource {
    /**
     * Get an existing Multicast resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MulticastState, opts?: pulumi.CustomResourceOptions): Multicast {
        return new Multicast(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:router/multicast:Multicast';

    /**
     * Returns true if the given object is an instance of Multicast.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Multicast {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Multicast.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * PIM interfaces. The structure of `interface` block is documented below.
     */
    public readonly interfaces!: pulumi.Output<outputs.router.MulticastInterface[] | undefined>;
    /**
     * Enable/disable IP multicast routing. Valid values: `enable`, `disable`.
     */
    public readonly multicastRouting!: pulumi.Output<string>;
    /**
     * PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
     */
    public readonly pimSmGlobal!: pulumi.Output<outputs.router.MulticastPimSmGlobal>;
    /**
     * Maximum number of multicast routes.
     */
    public readonly routeLimit!: pulumi.Output<number>;
    /**
     * Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
     */
    public readonly routeThreshold!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Multicast resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MulticastArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MulticastArgs | MulticastState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MulticastState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["multicastRouting"] = state ? state.multicastRouting : undefined;
            resourceInputs["pimSmGlobal"] = state ? state.pimSmGlobal : undefined;
            resourceInputs["routeLimit"] = state ? state.routeLimit : undefined;
            resourceInputs["routeThreshold"] = state ? state.routeThreshold : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as MulticastArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["multicastRouting"] = args ? args.multicastRouting : undefined;
            resourceInputs["pimSmGlobal"] = args ? args.pimSmGlobal : undefined;
            resourceInputs["routeLimit"] = args ? args.routeLimit : undefined;
            resourceInputs["routeThreshold"] = args ? args.routeThreshold : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Multicast.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Multicast resources.
 */
export interface MulticastState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * PIM interfaces. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.router.MulticastInterface>[]>;
    /**
     * Enable/disable IP multicast routing. Valid values: `enable`, `disable`.
     */
    multicastRouting?: pulumi.Input<string>;
    /**
     * PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
     */
    pimSmGlobal?: pulumi.Input<inputs.router.MulticastPimSmGlobal>;
    /**
     * Maximum number of multicast routes.
     */
    routeLimit?: pulumi.Input<number>;
    /**
     * Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
     */
    routeThreshold?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Multicast resource.
 */
export interface MulticastArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * PIM interfaces. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.router.MulticastInterface>[]>;
    /**
     * Enable/disable IP multicast routing. Valid values: `enable`, `disable`.
     */
    multicastRouting?: pulumi.Input<string>;
    /**
     * PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
     */
    pimSmGlobal?: pulumi.Input<inputs.router.MulticastPimSmGlobal>;
    /**
     * Maximum number of multicast routes.
     */
    routeLimit?: pulumi.Input<number>;
    /**
     * Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
     */
    routeThreshold?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
