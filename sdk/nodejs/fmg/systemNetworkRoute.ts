// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource supports updating system network route for FortiManager.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.fmg.SystemNetworkRoute("test1", {
 *     destination: "192.168.2.0 255.255.255.0",
 *     device: "port4",
 *     gateway: "192.168.2.1",
 *     routeId: 5,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class SystemNetworkRoute extends pulumi.CustomResource {
    /**
     * Get an existing SystemNetworkRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemNetworkRouteState, opts?: pulumi.CustomResourceOptions): SystemNetworkRoute {
        return new SystemNetworkRoute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:fmg/systemNetworkRoute:SystemNetworkRoute';

    /**
     * Returns true if the given object is an instance of SystemNetworkRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemNetworkRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemNetworkRoute.__pulumiType;
    }

    /**
     * Destination Ip/Mask.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * Gateway out interface.
     */
    public readonly device!: pulumi.Output<string>;
    /**
     * Gateway Ip.
     */
    public readonly gateway!: pulumi.Output<string>;
    /**
     * Route id.
     */
    public readonly routeId!: pulumi.Output<number>;

    /**
     * Create a SystemNetworkRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemNetworkRouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemNetworkRouteArgs | SystemNetworkRouteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemNetworkRouteState | undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["gateway"] = state ? state.gateway : undefined;
            resourceInputs["routeId"] = state ? state.routeId : undefined;
        } else {
            const args = argsOrState as SystemNetworkRouteArgs | undefined;
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.device === undefined) && !opts.urn) {
                throw new Error("Missing required property 'device'");
            }
            if ((!args || args.gateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gateway'");
            }
            if ((!args || args.routeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeId'");
            }
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["gateway"] = args ? args.gateway : undefined;
            resourceInputs["routeId"] = args ? args.routeId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemNetworkRoute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemNetworkRoute resources.
 */
export interface SystemNetworkRouteState {
    /**
     * Destination Ip/Mask.
     */
    destination?: pulumi.Input<string>;
    /**
     * Gateway out interface.
     */
    device?: pulumi.Input<string>;
    /**
     * Gateway Ip.
     */
    gateway?: pulumi.Input<string>;
    /**
     * Route id.
     */
    routeId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SystemNetworkRoute resource.
 */
export interface SystemNetworkRouteArgs {
    /**
     * Destination Ip/Mask.
     */
    destination: pulumi.Input<string>;
    /**
     * Gateway out interface.
     */
    device: pulumi.Input<string>;
    /**
     * Gateway Ip.
     */
    gateway: pulumi.Input<string>;
    /**
     * Route id.
     */
    routeId: pulumi.Input<number>;
}
