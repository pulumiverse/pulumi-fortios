// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure network visibility settings.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Networkvisibility("trname", {
 *     destinationHostnameVisibility: "enable",
 *     destinationLocation: "enable",
 *     destinationVisibility: "enable",
 *     hostnameLimit: 5000,
 *     hostnameTtl: 86400,
 *     sourceLocation: "enable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * System NetworkVisibility can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/networkvisibility:Networkvisibility labelname SystemNetworkVisibility
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/networkvisibility:Networkvisibility labelname SystemNetworkVisibility
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Networkvisibility extends pulumi.CustomResource {
    /**
     * Get an existing Networkvisibility resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkvisibilityState, opts?: pulumi.CustomResourceOptions): Networkvisibility {
        return new Networkvisibility(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/networkvisibility:Networkvisibility';

    /**
     * Returns true if the given object is an instance of Networkvisibility.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Networkvisibility {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Networkvisibility.__pulumiType;
    }

    /**
     * Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
     */
    public readonly destinationHostnameVisibility!: pulumi.Output<string>;
    /**
     * Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
     */
    public readonly destinationLocation!: pulumi.Output<string>;
    /**
     * Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
     */
    public readonly destinationVisibility!: pulumi.Output<string>;
    /**
     * Limit of the number of hostname table entries (0 - 50000).
     */
    public readonly hostnameLimit!: pulumi.Output<number>;
    /**
     * TTL of hostname table entries (60 - 86400).
     */
    public readonly hostnameTtl!: pulumi.Output<number>;
    /**
     * Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
     */
    public readonly sourceLocation!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Networkvisibility resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetworkvisibilityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkvisibilityArgs | NetworkvisibilityState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkvisibilityState | undefined;
            resourceInputs["destinationHostnameVisibility"] = state ? state.destinationHostnameVisibility : undefined;
            resourceInputs["destinationLocation"] = state ? state.destinationLocation : undefined;
            resourceInputs["destinationVisibility"] = state ? state.destinationVisibility : undefined;
            resourceInputs["hostnameLimit"] = state ? state.hostnameLimit : undefined;
            resourceInputs["hostnameTtl"] = state ? state.hostnameTtl : undefined;
            resourceInputs["sourceLocation"] = state ? state.sourceLocation : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as NetworkvisibilityArgs | undefined;
            resourceInputs["destinationHostnameVisibility"] = args ? args.destinationHostnameVisibility : undefined;
            resourceInputs["destinationLocation"] = args ? args.destinationLocation : undefined;
            resourceInputs["destinationVisibility"] = args ? args.destinationVisibility : undefined;
            resourceInputs["hostnameLimit"] = args ? args.hostnameLimit : undefined;
            resourceInputs["hostnameTtl"] = args ? args.hostnameTtl : undefined;
            resourceInputs["sourceLocation"] = args ? args.sourceLocation : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Networkvisibility.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Networkvisibility resources.
 */
export interface NetworkvisibilityState {
    /**
     * Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
     */
    destinationHostnameVisibility?: pulumi.Input<string>;
    /**
     * Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
     */
    destinationLocation?: pulumi.Input<string>;
    /**
     * Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
     */
    destinationVisibility?: pulumi.Input<string>;
    /**
     * Limit of the number of hostname table entries (0 - 50000).
     */
    hostnameLimit?: pulumi.Input<number>;
    /**
     * TTL of hostname table entries (60 - 86400).
     */
    hostnameTtl?: pulumi.Input<number>;
    /**
     * Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
     */
    sourceLocation?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Networkvisibility resource.
 */
export interface NetworkvisibilityArgs {
    /**
     * Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
     */
    destinationHostnameVisibility?: pulumi.Input<string>;
    /**
     * Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
     */
    destinationLocation?: pulumi.Input<string>;
    /**
     * Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
     */
    destinationVisibility?: pulumi.Input<string>;
    /**
     * Limit of the number of hostname table entries (0 - 50000).
     */
    hostnameLimit?: pulumi.Input<number>;
    /**
     * TTL of hostname table entries (60 - 86400).
     */
    hostnameTtl?: pulumi.Input<number>;
    /**
     * Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
     */
    sourceLocation?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
