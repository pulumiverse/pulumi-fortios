// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure IPv6 multicast.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.RouterMulticast6("trname", {
 *     multicastPmtu: "disable",
 *     multicastRouting: "disable",
 *     pimSmGlobal: {
 *         registerRateLimit: 0,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Router Multicast6 can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/routerMulticast6:RouterMulticast6 labelname RouterMulticast6
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/routerMulticast6:RouterMulticast6 labelname RouterMulticast6
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class RouterMulticast6 extends pulumi.CustomResource {
    /**
     * Get an existing RouterMulticast6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterMulticast6State, opts?: pulumi.CustomResourceOptions): RouterMulticast6 {
        return new RouterMulticast6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerMulticast6:RouterMulticast6';

    /**
     * Returns true if the given object is an instance of RouterMulticast6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterMulticast6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterMulticast6.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
     */
    public readonly interfaces!: pulumi.Output<outputs.RouterMulticast6Interface[] | undefined>;
    /**
     * Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
     */
    public readonly multicastPmtu!: pulumi.Output<string>;
    /**
     * Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
     */
    public readonly multicastRouting!: pulumi.Output<string>;
    /**
     * PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
     */
    public readonly pimSmGlobal!: pulumi.Output<outputs.RouterMulticast6PimSmGlobal>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterMulticast6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RouterMulticast6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterMulticast6Args | RouterMulticast6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterMulticast6State | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["multicastPmtu"] = state ? state.multicastPmtu : undefined;
            resourceInputs["multicastRouting"] = state ? state.multicastRouting : undefined;
            resourceInputs["pimSmGlobal"] = state ? state.pimSmGlobal : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RouterMulticast6Args | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["multicastPmtu"] = args ? args.multicastPmtu : undefined;
            resourceInputs["multicastRouting"] = args ? args.multicastRouting : undefined;
            resourceInputs["pimSmGlobal"] = args ? args.pimSmGlobal : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterMulticast6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterMulticast6 resources.
 */
export interface RouterMulticast6State {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.RouterMulticast6Interface>[]>;
    /**
     * Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
     */
    multicastPmtu?: pulumi.Input<string>;
    /**
     * Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
     */
    multicastRouting?: pulumi.Input<string>;
    /**
     * PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
     */
    pimSmGlobal?: pulumi.Input<inputs.RouterMulticast6PimSmGlobal>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterMulticast6 resource.
 */
export interface RouterMulticast6Args {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.RouterMulticast6Interface>[]>;
    /**
     * Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
     */
    multicastPmtu?: pulumi.Input<string>;
    /**
     * Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
     */
    multicastRouting?: pulumi.Input<string>;
    /**
     * PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
     */
    pimSmGlobal?: pulumi.Input<inputs.RouterMulticast6PimSmGlobal>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}