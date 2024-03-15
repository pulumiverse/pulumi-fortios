// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure IPv6 BFD.
 *
 * ## Import
 *
 * Router Bfd6 can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:router/bfd6:Bfd6 labelname RouterBfd6
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:router/bfd6:Bfd6 labelname RouterBfd6
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Bfd6 extends pulumi.CustomResource {
    /**
     * Get an existing Bfd6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Bfd6State, opts?: pulumi.CustomResourceOptions): Bfd6 {
        return new Bfd6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:router/bfd6:Bfd6';

    /**
     * Returns true if the given object is an instance of Bfd6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bfd6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bfd6.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * BFD IPv6 multi-hop template table. The structure of `multihopTemplate` block is documented below.
     */
    public readonly multihopTemplates!: pulumi.Output<outputs.router.Bfd6MultihopTemplate[] | undefined>;
    /**
     * Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
     */
    public readonly neighbors!: pulumi.Output<outputs.router.Bfd6Neighbor[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Bfd6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: Bfd6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Bfd6Args | Bfd6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Bfd6State | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["multihopTemplates"] = state ? state.multihopTemplates : undefined;
            resourceInputs["neighbors"] = state ? state.neighbors : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Bfd6Args | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["multihopTemplates"] = args ? args.multihopTemplates : undefined;
            resourceInputs["neighbors"] = args ? args.neighbors : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Bfd6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Bfd6 resources.
 */
export interface Bfd6State {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * BFD IPv6 multi-hop template table. The structure of `multihopTemplate` block is documented below.
     */
    multihopTemplates?: pulumi.Input<pulumi.Input<inputs.router.Bfd6MultihopTemplate>[]>;
    /**
     * Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
     */
    neighbors?: pulumi.Input<pulumi.Input<inputs.router.Bfd6Neighbor>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Bfd6 resource.
 */
export interface Bfd6Args {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * BFD IPv6 multi-hop template table. The structure of `multihopTemplate` block is documented below.
     */
    multihopTemplates?: pulumi.Input<pulumi.Input<inputs.router.Bfd6MultihopTemplate>[]>;
    /**
     * Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
     */
    neighbors?: pulumi.Input<pulumi.Input<inputs.router.Bfd6Neighbor>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
