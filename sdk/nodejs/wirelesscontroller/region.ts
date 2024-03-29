// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure FortiAP regions (for floor plans and maps).
 *
 * ## Import
 *
 * WirelessController Region can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:wirelesscontroller/region:Region labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:wirelesscontroller/region:Region labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Region extends pulumi.CustomResource {
    /**
     * Get an existing Region resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionState, opts?: pulumi.CustomResourceOptions): Region {
        return new Region(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:wirelesscontroller/region:Region';

    /**
     * Returns true if the given object is an instance of Region.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Region {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Region.__pulumiType;
    }

    /**
     * Comments.
     */
    public readonly comments!: pulumi.Output<string>;
    /**
     * Region image grayscale. Valid values: `enable`, `disable`.
     */
    public readonly grayscale!: pulumi.Output<string>;
    /**
     * FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
     */
    public readonly imageType!: pulumi.Output<string>;
    /**
     * FortiAP region name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Region image opacity (0 - 100).
     */
    public readonly opacity!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Region resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RegionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionArgs | RegionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegionState | undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["grayscale"] = state ? state.grayscale : undefined;
            resourceInputs["imageType"] = state ? state.imageType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["opacity"] = state ? state.opacity : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RegionArgs | undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["grayscale"] = args ? args.grayscale : undefined;
            resourceInputs["imageType"] = args ? args.imageType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["opacity"] = args ? args.opacity : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Region.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Region resources.
 */
export interface RegionState {
    /**
     * Comments.
     */
    comments?: pulumi.Input<string>;
    /**
     * Region image grayscale. Valid values: `enable`, `disable`.
     */
    grayscale?: pulumi.Input<string>;
    /**
     * FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
     */
    imageType?: pulumi.Input<string>;
    /**
     * FortiAP region name.
     */
    name?: pulumi.Input<string>;
    /**
     * Region image opacity (0 - 100).
     */
    opacity?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Region resource.
 */
export interface RegionArgs {
    /**
     * Comments.
     */
    comments?: pulumi.Input<string>;
    /**
     * Region image grayscale. Valid values: `enable`, `disable`.
     */
    grayscale?: pulumi.Input<string>;
    /**
     * FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
     */
    imageType?: pulumi.Input<string>;
    /**
     * FortiAP region name.
     */
    name?: pulumi.Input<string>;
    /**
     * Region image opacity (0 - 100).
     */
    opacity?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
