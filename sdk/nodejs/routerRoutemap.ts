// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure route maps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.RouterRoutemap("trname", {rules: [{
 *     action: "deny",
 *     matchCommunityExact: "disable",
 *     matchFlags: 0,
 *     matchMetric: 0,
 *     matchOrigin: "none",
 *     matchRouteType: "No type specified",
 *     matchTag: 0,
 *     setAggregatorAs: 0,
 *     setAggregatorIp: "0.0.0.0",
 *     setAspathAction: "prepend",
 *     setAtomicAggregate: "disable",
 *     setCommunityAdditive: "disable",
 *     setDampeningMaxSuppress: 0,
 *     setDampeningReachabilityHalfLife: 0,
 *     setDampeningReuse: 0,
 *     setDampeningSuppress: 0,
 *     setDampeningUnreachabilityHalfLife: 0,
 *     setFlags: 128,
 *     setIp6Nexthop: "::",
 *     setIp6NexthopLocal: "::",
 *     setIpNexthop: "0.0.0.0",
 *     setLocalPreference: 0,
 *     setMetric: 0,
 *     setMetricType: "No type specified",
 *     setOrigin: "none",
 *     setOriginatorId: "0.0.0.0",
 *     setRouteTag: 0,
 *     setTag: 0,
 *     setWeight: 21,
 * }]});
 * ```
 *
 * ## Import
 *
 * Router RouteMap can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/routerRoutemap:RouterRoutemap labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/routerRoutemap:RouterRoutemap labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class RouterRoutemap extends pulumi.CustomResource {
    /**
     * Get an existing RouterRoutemap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterRoutemapState, opts?: pulumi.CustomResourceOptions): RouterRoutemap {
        return new RouterRoutemap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerRoutemap:RouterRoutemap';

    /**
     * Returns true if the given object is an instance of RouterRoutemap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterRoutemap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterRoutemap.__pulumiType;
    }

    /**
     * Optional comments.
     */
    public readonly comments!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Rule. The structure of `rule` block is documented below.
     */
    public readonly rules!: pulumi.Output<outputs.RouterRoutemapRule[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterRoutemap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RouterRoutemapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterRoutemapArgs | RouterRoutemapState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterRoutemapState | undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RouterRoutemapArgs | undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterRoutemap.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterRoutemap resources.
 */
export interface RouterRoutemapState {
    /**
     * Optional comments.
     */
    comments?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Rule. The structure of `rule` block is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.RouterRoutemapRule>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterRoutemap resource.
 */
export interface RouterRoutemapArgs {
    /**
     * Optional comments.
     */
    comments?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Rule. The structure of `rule` block is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.RouterRoutemapRule>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
