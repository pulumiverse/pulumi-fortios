// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to configure static route of FortiOS.
 *
 * !> **Warning:** The resource will be deprecated and replaced by new resource `fortios.router.Static`, we recommend that you use the new resource.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const subnet = new fortios.networking.RouteStatic("subnet", {
 *     blackhole: "disable",
 *     comment: "Terraform test",
 *     device: "port2",
 *     distance: "22",
 *     dst: "110.2.2.122/32",
 *     gateway: "2.2.2.2",
 *     priority: "3",
 *     status: "enable",
 *     weight: "3",
 * });
 * const internetService = new fortios.networking.RouteStatic("internetService", {
 *     blackhole: "disable",
 *     comment: "Terraform Test",
 *     device: "port2",
 *     distance: "22",
 *     gateway: "2.2.2.2",
 *     internetService: 5242881,
 *     priority: "3",
 *     status: "enable",
 *     weight: "3",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class RouteStatic extends pulumi.CustomResource {
    /**
     * Get an existing RouteStatic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteStaticState, opts?: pulumi.CustomResourceOptions): RouteStatic {
        return new RouteStatic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:networking/routeStatic:RouteStatic';

    /**
     * Returns true if the given object is an instance of RouteStatic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteStatic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteStatic.__pulumiType;
    }

    /**
     * Enable/disable black hole.
     */
    public readonly blackhole!: pulumi.Output<string>;
    /**
     * Optional comments.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Gateway out interface or tunnel.
     */
    public readonly device!: pulumi.Output<string>;
    /**
     * Administrative distance.
     */
    public readonly distance!: pulumi.Output<string>;
    /**
     * Destination IP and mask for this route.
     */
    public readonly dst!: pulumi.Output<string>;
    /**
     * Gateway IP for this route.
     */
    public readonly gateway!: pulumi.Output<string>;
    /**
     * Application ID in the Internet service database.
     */
    public readonly internetService!: pulumi.Output<number>;
    /**
     * Administrative priority.
     */
    public readonly priority!: pulumi.Output<string>;
    /**
     * Enable/disable this static route. default is "enable".
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Administrative weight.
     */
    public readonly weight!: pulumi.Output<string>;

    /**
     * Create a RouteStatic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteStaticArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteStaticArgs | RouteStaticState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouteStaticState | undefined;
            resourceInputs["blackhole"] = state ? state.blackhole : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["distance"] = state ? state.distance : undefined;
            resourceInputs["dst"] = state ? state.dst : undefined;
            resourceInputs["gateway"] = state ? state.gateway : undefined;
            resourceInputs["internetService"] = state ? state.internetService : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as RouteStaticArgs | undefined;
            if ((!args || args.device === undefined) && !opts.urn) {
                throw new Error("Missing required property 'device'");
            }
            if ((!args || args.gateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gateway'");
            }
            resourceInputs["blackhole"] = args ? args.blackhole : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["distance"] = args ? args.distance : undefined;
            resourceInputs["dst"] = args ? args.dst : undefined;
            resourceInputs["gateway"] = args ? args.gateway : undefined;
            resourceInputs["internetService"] = args ? args.internetService : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouteStatic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteStatic resources.
 */
export interface RouteStaticState {
    /**
     * Enable/disable black hole.
     */
    blackhole?: pulumi.Input<string>;
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Gateway out interface or tunnel.
     */
    device?: pulumi.Input<string>;
    /**
     * Administrative distance.
     */
    distance?: pulumi.Input<string>;
    /**
     * Destination IP and mask for this route.
     */
    dst?: pulumi.Input<string>;
    /**
     * Gateway IP for this route.
     */
    gateway?: pulumi.Input<string>;
    /**
     * Application ID in the Internet service database.
     */
    internetService?: pulumi.Input<number>;
    /**
     * Administrative priority.
     */
    priority?: pulumi.Input<string>;
    /**
     * Enable/disable this static route. default is "enable".
     */
    status?: pulumi.Input<string>;
    /**
     * Administrative weight.
     */
    weight?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteStatic resource.
 */
export interface RouteStaticArgs {
    /**
     * Enable/disable black hole.
     */
    blackhole?: pulumi.Input<string>;
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Gateway out interface or tunnel.
     */
    device: pulumi.Input<string>;
    /**
     * Administrative distance.
     */
    distance?: pulumi.Input<string>;
    /**
     * Destination IP and mask for this route.
     */
    dst?: pulumi.Input<string>;
    /**
     * Gateway IP for this route.
     */
    gateway: pulumi.Input<string>;
    /**
     * Application ID in the Internet service database.
     */
    internetService?: pulumi.Input<number>;
    /**
     * Administrative priority.
     */
    priority?: pulumi.Input<string>;
    /**
     * Enable/disable this static route. default is "enable".
     */
    status?: pulumi.Input<string>;
    /**
     * Administrative weight.
     */
    weight?: pulumi.Input<string>;
}
