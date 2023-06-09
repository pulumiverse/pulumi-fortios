// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure IPS URL filter settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.WebfilterIpsurlfiltersetting("trname", {
 *     distance: 1,
 *     gateway: "0.0.0.0",
 * });
 * ```
 *
 * ## Import
 *
 * Webfilter IpsUrlfilterSetting can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/webfilterIpsurlfiltersetting:WebfilterIpsurlfiltersetting labelname WebfilterIpsUrlfilterSetting
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/webfilterIpsurlfiltersetting:WebfilterIpsurlfiltersetting labelname WebfilterIpsUrlfilterSetting
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WebfilterIpsurlfiltersetting extends pulumi.CustomResource {
    /**
     * Get an existing WebfilterIpsurlfiltersetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebfilterIpsurlfiltersettingState, opts?: pulumi.CustomResourceOptions): WebfilterIpsurlfiltersetting {
        return new WebfilterIpsurlfiltersetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/webfilterIpsurlfiltersetting:WebfilterIpsurlfiltersetting';

    /**
     * Returns true if the given object is an instance of WebfilterIpsurlfiltersetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebfilterIpsurlfiltersetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebfilterIpsurlfiltersetting.__pulumiType;
    }

    /**
     * Interface for this route.
     */
    public readonly device!: pulumi.Output<string>;
    /**
     * Administrative distance (1 - 255) for this route.
     */
    public readonly distance!: pulumi.Output<number>;
    /**
     * Gateway IP address for this route.
     */
    public readonly gateway!: pulumi.Output<string>;
    /**
     * Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
     */
    public readonly geoFilter!: pulumi.Output<string | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WebfilterIpsurlfiltersetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WebfilterIpsurlfiltersettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebfilterIpsurlfiltersettingArgs | WebfilterIpsurlfiltersettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebfilterIpsurlfiltersettingState | undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["distance"] = state ? state.distance : undefined;
            resourceInputs["gateway"] = state ? state.gateway : undefined;
            resourceInputs["geoFilter"] = state ? state.geoFilter : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WebfilterIpsurlfiltersettingArgs | undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["distance"] = args ? args.distance : undefined;
            resourceInputs["gateway"] = args ? args.gateway : undefined;
            resourceInputs["geoFilter"] = args ? args.geoFilter : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WebfilterIpsurlfiltersetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebfilterIpsurlfiltersetting resources.
 */
export interface WebfilterIpsurlfiltersettingState {
    /**
     * Interface for this route.
     */
    device?: pulumi.Input<string>;
    /**
     * Administrative distance (1 - 255) for this route.
     */
    distance?: pulumi.Input<number>;
    /**
     * Gateway IP address for this route.
     */
    gateway?: pulumi.Input<string>;
    /**
     * Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
     */
    geoFilter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebfilterIpsurlfiltersetting resource.
 */
export interface WebfilterIpsurlfiltersettingArgs {
    /**
     * Interface for this route.
     */
    device?: pulumi.Input<string>;
    /**
     * Administrative distance (1 - 255) for this route.
     */
    distance?: pulumi.Input<number>;
    /**
     * Gateway IP address for this route.
     */
    gateway?: pulumi.Input<string>;
    /**
     * Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
     */
    geoFilter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
