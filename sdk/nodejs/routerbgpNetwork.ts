// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * BGP network table.
 *
 * > The provider supports the definition of Network in Router Bgp `fortios.RouterBgp`, and also allows the definition of separate Network resources `fortios.RouterbgpNetwork`, but do not use a `fortios.RouterBgp` with in-line Network in conjunction with any `fortios.RouterbgpNetwork` resources, otherwise conflicts and overwrite will occur.
 *
 * ## Import
 *
 * Routerbgp Network can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/routerbgpNetwork:RouterbgpNetwork labelname {{fosid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/routerbgpNetwork:RouterbgpNetwork labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class RouterbgpNetwork extends pulumi.CustomResource {
    /**
     * Get an existing RouterbgpNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterbgpNetworkState, opts?: pulumi.CustomResourceOptions): RouterbgpNetwork {
        return new RouterbgpNetwork(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerbgpNetwork:RouterbgpNetwork';

    /**
     * Returns true if the given object is an instance of RouterbgpNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterbgpNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterbgpNetwork.__pulumiType;
    }

    /**
     * Enable/disable route as backdoor. Valid values: `enable`, `disable`.
     */
    public readonly backdoor!: pulumi.Output<string>;
    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
     */
    public readonly networkImportCheck!: pulumi.Output<string>;
    /**
     * Network prefix.
     */
    public readonly prefix!: pulumi.Output<string>;
    /**
     * Route map to modify generated route.
     */
    public readonly routeMap!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterbgpNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterbgpNetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterbgpNetworkArgs | RouterbgpNetworkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterbgpNetworkState | undefined;
            resourceInputs["backdoor"] = state ? state.backdoor : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["networkImportCheck"] = state ? state.networkImportCheck : undefined;
            resourceInputs["prefix"] = state ? state.prefix : undefined;
            resourceInputs["routeMap"] = state ? state.routeMap : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RouterbgpNetworkArgs | undefined;
            if ((!args || args.prefix === undefined) && !opts.urn) {
                throw new Error("Missing required property 'prefix'");
            }
            resourceInputs["backdoor"] = args ? args.backdoor : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["networkImportCheck"] = args ? args.networkImportCheck : undefined;
            resourceInputs["prefix"] = args ? args.prefix : undefined;
            resourceInputs["routeMap"] = args ? args.routeMap : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterbgpNetwork.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterbgpNetwork resources.
 */
export interface RouterbgpNetworkState {
    /**
     * Enable/disable route as backdoor. Valid values: `enable`, `disable`.
     */
    backdoor?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
     */
    networkImportCheck?: pulumi.Input<string>;
    /**
     * Network prefix.
     */
    prefix?: pulumi.Input<string>;
    /**
     * Route map to modify generated route.
     */
    routeMap?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterbgpNetwork resource.
 */
export interface RouterbgpNetworkArgs {
    /**
     * Enable/disable route as backdoor. Valid values: `enable`, `disable`.
     */
    backdoor?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
     */
    networkImportCheck?: pulumi.Input<string>;
    /**
     * Network prefix.
     */
    prefix: pulumi.Input<string>;
    /**
     * Route map to modify generated route.
     */
    routeMap?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}