// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure replacement message images.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Replacemsgimage("trname", {
 *     imageBase64: "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAEWAAABFgAVshLGQAAAAMSURBVBhXY/j//z8ABf4C/qc1gYQAAAAASUVORK5CYII=",
 *     imageType: "png",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * System ReplacemsgImage can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/replacemsgimage:Replacemsgimage labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/replacemsgimage:Replacemsgimage labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Replacemsgimage extends pulumi.CustomResource {
    /**
     * Get an existing Replacemsgimage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplacemsgimageState, opts?: pulumi.CustomResourceOptions): Replacemsgimage {
        return new Replacemsgimage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/replacemsgimage:Replacemsgimage';

    /**
     * Returns true if the given object is an instance of Replacemsgimage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Replacemsgimage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Replacemsgimage.__pulumiType;
    }

    /**
     * Image data.
     */
    public readonly imageBase64!: pulumi.Output<string | undefined>;
    /**
     * Image type. Valid values: `gif`, `jpg`, `tiff`, `png`.
     */
    public readonly imageType!: pulumi.Output<string>;
    /**
     * Image name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Replacemsgimage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ReplacemsgimageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplacemsgimageArgs | ReplacemsgimageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReplacemsgimageState | undefined;
            resourceInputs["imageBase64"] = state ? state.imageBase64 : undefined;
            resourceInputs["imageType"] = state ? state.imageType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ReplacemsgimageArgs | undefined;
            resourceInputs["imageBase64"] = args ? args.imageBase64 : undefined;
            resourceInputs["imageType"] = args ? args.imageType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Replacemsgimage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Replacemsgimage resources.
 */
export interface ReplacemsgimageState {
    /**
     * Image data.
     */
    imageBase64?: pulumi.Input<string>;
    /**
     * Image type. Valid values: `gif`, `jpg`, `tiff`, `png`.
     */
    imageType?: pulumi.Input<string>;
    /**
     * Image name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Replacemsgimage resource.
 */
export interface ReplacemsgimageArgs {
    /**
     * Image data.
     */
    imageBase64?: pulumi.Input<string>;
    /**
     * Image type. Valid values: `gif`, `jpg`, `tiff`, `png`.
     */
    imageType?: pulumi.Input<string>;
    /**
     * Image name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
