// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * configure ips view-map Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * Ips ViewMap can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:ips/viewmap:Viewmap labelname {{fosid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:ips/viewmap:Viewmap labelname {{fosid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Viewmap extends pulumi.CustomResource {
    /**
     * Get an existing Viewmap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ViewmapState, opts?: pulumi.CustomResourceOptions): Viewmap {
        return new Viewmap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:ips/viewmap:Viewmap';

    /**
     * Returns true if the given object is an instance of Viewmap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Viewmap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Viewmap.__pulumiType;
    }

    /**
     * View ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * ID-based policy ID.
     */
    public readonly idPolicyId!: pulumi.Output<number>;
    /**
     * Policy ID.
     */
    public readonly policyId!: pulumi.Output<number>;
    /**
     * VDOM ID.
     */
    public readonly vdomId!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;
    /**
     * Policy.
     */
    public readonly which!: pulumi.Output<string>;

    /**
     * Create a Viewmap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ViewmapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ViewmapArgs | ViewmapState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ViewmapState | undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["idPolicyId"] = state ? state.idPolicyId : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
            resourceInputs["vdomId"] = state ? state.vdomId : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["which"] = state ? state.which : undefined;
        } else {
            const args = argsOrState as ViewmapArgs | undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["idPolicyId"] = args ? args.idPolicyId : undefined;
            resourceInputs["policyId"] = args ? args.policyId : undefined;
            resourceInputs["vdomId"] = args ? args.vdomId : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["which"] = args ? args.which : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Viewmap.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Viewmap resources.
 */
export interface ViewmapState {
    /**
     * View ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * ID-based policy ID.
     */
    idPolicyId?: pulumi.Input<number>;
    /**
     * Policy ID.
     */
    policyId?: pulumi.Input<number>;
    /**
     * VDOM ID.
     */
    vdomId?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Policy.
     */
    which?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Viewmap resource.
 */
export interface ViewmapArgs {
    /**
     * View ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * ID-based policy ID.
     */
    idPolicyId?: pulumi.Input<number>;
    /**
     * Policy ID.
     */
    policyId?: pulumi.Input<number>;
    /**
     * VDOM ID.
     */
    vdomId?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Policy.
     */
    which?: pulumi.Input<string>;
}
