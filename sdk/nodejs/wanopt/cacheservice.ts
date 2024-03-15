// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Designate cache-service for wan-optimization and webcache.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.wanopt.Cacheservice("trname", {
 *     acceptableConnections: "any",
 *     collaboration: "disable",
 *     deviceId: "default_dev_id",
 *     preferScenario: "balance",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Wanopt CacheService can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:wanopt/cacheservice:Cacheservice labelname WanoptCacheService
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:wanopt/cacheservice:Cacheservice labelname WanoptCacheService
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Cacheservice extends pulumi.CustomResource {
    /**
     * Get an existing Cacheservice resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CacheserviceState, opts?: pulumi.CustomResourceOptions): Cacheservice {
        return new Cacheservice(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:wanopt/cacheservice:Cacheservice';

    /**
     * Returns true if the given object is an instance of Cacheservice.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cacheservice {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cacheservice.__pulumiType;
    }

    /**
     * Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
     */
    public readonly acceptableConnections!: pulumi.Output<string>;
    /**
     * Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
     */
    public readonly collaboration!: pulumi.Output<string>;
    /**
     * Set identifier for this cache device.
     */
    public readonly deviceId!: pulumi.Output<string>;
    /**
     * Modify cache-service destination peer list. The structure of `dstPeer` block is documented below.
     */
    public readonly dstPeers!: pulumi.Output<outputs.wanopt.CacheserviceDstPeer[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
     */
    public readonly preferScenario!: pulumi.Output<string>;
    /**
     * Modify cache-service source peer list. The structure of `srcPeer` block is documented below.
     */
    public readonly srcPeers!: pulumi.Output<outputs.wanopt.CacheserviceSrcPeer[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Cacheservice resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CacheserviceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CacheserviceArgs | CacheserviceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CacheserviceState | undefined;
            resourceInputs["acceptableConnections"] = state ? state.acceptableConnections : undefined;
            resourceInputs["collaboration"] = state ? state.collaboration : undefined;
            resourceInputs["deviceId"] = state ? state.deviceId : undefined;
            resourceInputs["dstPeers"] = state ? state.dstPeers : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["preferScenario"] = state ? state.preferScenario : undefined;
            resourceInputs["srcPeers"] = state ? state.srcPeers : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as CacheserviceArgs | undefined;
            resourceInputs["acceptableConnections"] = args ? args.acceptableConnections : undefined;
            resourceInputs["collaboration"] = args ? args.collaboration : undefined;
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["dstPeers"] = args ? args.dstPeers : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["preferScenario"] = args ? args.preferScenario : undefined;
            resourceInputs["srcPeers"] = args ? args.srcPeers : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cacheservice.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cacheservice resources.
 */
export interface CacheserviceState {
    /**
     * Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
     */
    acceptableConnections?: pulumi.Input<string>;
    /**
     * Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
     */
    collaboration?: pulumi.Input<string>;
    /**
     * Set identifier for this cache device.
     */
    deviceId?: pulumi.Input<string>;
    /**
     * Modify cache-service destination peer list. The structure of `dstPeer` block is documented below.
     */
    dstPeers?: pulumi.Input<pulumi.Input<inputs.wanopt.CacheserviceDstPeer>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
     */
    preferScenario?: pulumi.Input<string>;
    /**
     * Modify cache-service source peer list. The structure of `srcPeer` block is documented below.
     */
    srcPeers?: pulumi.Input<pulumi.Input<inputs.wanopt.CacheserviceSrcPeer>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cacheservice resource.
 */
export interface CacheserviceArgs {
    /**
     * Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
     */
    acceptableConnections?: pulumi.Input<string>;
    /**
     * Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
     */
    collaboration?: pulumi.Input<string>;
    /**
     * Set identifier for this cache device.
     */
    deviceId?: pulumi.Input<string>;
    /**
     * Modify cache-service destination peer list. The structure of `dstPeer` block is documented below.
     */
    dstPeers?: pulumi.Input<pulumi.Input<inputs.wanopt.CacheserviceDstPeer>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
     */
    preferScenario?: pulumi.Input<string>;
    /**
     * Modify cache-service source peer list. The structure of `srcPeer` block is documented below.
     */
    srcPeers?: pulumi.Input<pulumi.Input<inputs.wanopt.CacheserviceSrcPeer>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
