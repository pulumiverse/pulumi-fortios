// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure global Web cache settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.wanopt.Webcache("trname", {
 *     alwaysRevalidate: "disable",
 *     cacheByDefault: "disable",
 *     cacheCookie: "disable",
 *     cacheExpired: "disable",
 *     defaultTtl: 1440,
 *     external: "disable",
 *     freshFactor: 100,
 *     hostValidate: "disable",
 *     ignoreConditional: "disable",
 *     ignoreIeReload: "enable",
 *     ignoreIms: "disable",
 *     ignorePnc: "disable",
 *     maxObjectSize: 512000,
 *     maxTtl: 7200,
 *     minTtl: 5,
 *     negRespTime: 0,
 *     revalPnc: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * Wanopt Webcache can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:wanopt/webcache:Webcache labelname WanoptWebcache
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:wanopt/webcache:Webcache labelname WanoptWebcache
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Webcache extends pulumi.CustomResource {
    /**
     * Get an existing Webcache resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebcacheState, opts?: pulumi.CustomResourceOptions): Webcache {
        return new Webcache(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:wanopt/webcache:Webcache';

    /**
     * Returns true if the given object is an instance of Webcache.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Webcache {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Webcache.__pulumiType;
    }

    /**
     * Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
     */
    public readonly alwaysRevalidate!: pulumi.Output<string>;
    /**
     * Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
     */
    public readonly cacheByDefault!: pulumi.Output<string>;
    /**
     * Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
     */
    public readonly cacheCookie!: pulumi.Output<string>;
    /**
     * Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
     */
    public readonly cacheExpired!: pulumi.Output<string>;
    /**
     * Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
     */
    public readonly defaultTtl!: pulumi.Output<number>;
    /**
     * Enable/disable external Web caching. Valid values: `enable`, `disable`.
     */
    public readonly external!: pulumi.Output<string>;
    /**
     * Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
     */
    public readonly freshFactor!: pulumi.Output<number>;
    /**
     * Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
     */
    public readonly hostValidate!: pulumi.Output<string>;
    /**
     * Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
     */
    public readonly ignoreConditional!: pulumi.Output<string>;
    /**
     * Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
     */
    public readonly ignoreIeReload!: pulumi.Output<string>;
    /**
     * Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
     */
    public readonly ignoreIms!: pulumi.Output<string>;
    /**
     * Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
     */
    public readonly ignorePnc!: pulumi.Output<string>;
    /**
     * Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
     */
    public readonly maxObjectSize!: pulumi.Output<number>;
    /**
     * Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
     */
    public readonly maxTtl!: pulumi.Output<number>;
    /**
     * Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
     */
    public readonly minTtl!: pulumi.Output<number>;
    /**
     * Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
     */
    public readonly negRespTime!: pulumi.Output<number>;
    /**
     * Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
     */
    public readonly revalPnc!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Webcache resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WebcacheArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebcacheArgs | WebcacheState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebcacheState | undefined;
            resourceInputs["alwaysRevalidate"] = state ? state.alwaysRevalidate : undefined;
            resourceInputs["cacheByDefault"] = state ? state.cacheByDefault : undefined;
            resourceInputs["cacheCookie"] = state ? state.cacheCookie : undefined;
            resourceInputs["cacheExpired"] = state ? state.cacheExpired : undefined;
            resourceInputs["defaultTtl"] = state ? state.defaultTtl : undefined;
            resourceInputs["external"] = state ? state.external : undefined;
            resourceInputs["freshFactor"] = state ? state.freshFactor : undefined;
            resourceInputs["hostValidate"] = state ? state.hostValidate : undefined;
            resourceInputs["ignoreConditional"] = state ? state.ignoreConditional : undefined;
            resourceInputs["ignoreIeReload"] = state ? state.ignoreIeReload : undefined;
            resourceInputs["ignoreIms"] = state ? state.ignoreIms : undefined;
            resourceInputs["ignorePnc"] = state ? state.ignorePnc : undefined;
            resourceInputs["maxObjectSize"] = state ? state.maxObjectSize : undefined;
            resourceInputs["maxTtl"] = state ? state.maxTtl : undefined;
            resourceInputs["minTtl"] = state ? state.minTtl : undefined;
            resourceInputs["negRespTime"] = state ? state.negRespTime : undefined;
            resourceInputs["revalPnc"] = state ? state.revalPnc : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WebcacheArgs | undefined;
            resourceInputs["alwaysRevalidate"] = args ? args.alwaysRevalidate : undefined;
            resourceInputs["cacheByDefault"] = args ? args.cacheByDefault : undefined;
            resourceInputs["cacheCookie"] = args ? args.cacheCookie : undefined;
            resourceInputs["cacheExpired"] = args ? args.cacheExpired : undefined;
            resourceInputs["defaultTtl"] = args ? args.defaultTtl : undefined;
            resourceInputs["external"] = args ? args.external : undefined;
            resourceInputs["freshFactor"] = args ? args.freshFactor : undefined;
            resourceInputs["hostValidate"] = args ? args.hostValidate : undefined;
            resourceInputs["ignoreConditional"] = args ? args.ignoreConditional : undefined;
            resourceInputs["ignoreIeReload"] = args ? args.ignoreIeReload : undefined;
            resourceInputs["ignoreIms"] = args ? args.ignoreIms : undefined;
            resourceInputs["ignorePnc"] = args ? args.ignorePnc : undefined;
            resourceInputs["maxObjectSize"] = args ? args.maxObjectSize : undefined;
            resourceInputs["maxTtl"] = args ? args.maxTtl : undefined;
            resourceInputs["minTtl"] = args ? args.minTtl : undefined;
            resourceInputs["negRespTime"] = args ? args.negRespTime : undefined;
            resourceInputs["revalPnc"] = args ? args.revalPnc : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Webcache.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Webcache resources.
 */
export interface WebcacheState {
    /**
     * Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
     */
    alwaysRevalidate?: pulumi.Input<string>;
    /**
     * Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
     */
    cacheByDefault?: pulumi.Input<string>;
    /**
     * Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
     */
    cacheCookie?: pulumi.Input<string>;
    /**
     * Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
     */
    cacheExpired?: pulumi.Input<string>;
    /**
     * Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
     */
    defaultTtl?: pulumi.Input<number>;
    /**
     * Enable/disable external Web caching. Valid values: `enable`, `disable`.
     */
    external?: pulumi.Input<string>;
    /**
     * Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
     */
    freshFactor?: pulumi.Input<number>;
    /**
     * Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
     */
    hostValidate?: pulumi.Input<string>;
    /**
     * Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
     */
    ignoreConditional?: pulumi.Input<string>;
    /**
     * Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
     */
    ignoreIeReload?: pulumi.Input<string>;
    /**
     * Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
     */
    ignoreIms?: pulumi.Input<string>;
    /**
     * Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
     */
    ignorePnc?: pulumi.Input<string>;
    /**
     * Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
     */
    maxObjectSize?: pulumi.Input<number>;
    /**
     * Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
     */
    maxTtl?: pulumi.Input<number>;
    /**
     * Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
     */
    minTtl?: pulumi.Input<number>;
    /**
     * Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
     */
    negRespTime?: pulumi.Input<number>;
    /**
     * Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
     */
    revalPnc?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Webcache resource.
 */
export interface WebcacheArgs {
    /**
     * Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
     */
    alwaysRevalidate?: pulumi.Input<string>;
    /**
     * Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
     */
    cacheByDefault?: pulumi.Input<string>;
    /**
     * Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
     */
    cacheCookie?: pulumi.Input<string>;
    /**
     * Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
     */
    cacheExpired?: pulumi.Input<string>;
    /**
     * Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
     */
    defaultTtl?: pulumi.Input<number>;
    /**
     * Enable/disable external Web caching. Valid values: `enable`, `disable`.
     */
    external?: pulumi.Input<string>;
    /**
     * Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
     */
    freshFactor?: pulumi.Input<number>;
    /**
     * Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
     */
    hostValidate?: pulumi.Input<string>;
    /**
     * Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
     */
    ignoreConditional?: pulumi.Input<string>;
    /**
     * Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
     */
    ignoreIeReload?: pulumi.Input<string>;
    /**
     * Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
     */
    ignoreIms?: pulumi.Input<string>;
    /**
     * Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
     */
    ignorePnc?: pulumi.Input<string>;
    /**
     * Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
     */
    maxObjectSize?: pulumi.Input<number>;
    /**
     * Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
     */
    maxTtl?: pulumi.Input<number>;
    /**
     * Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
     */
    minTtl?: pulumi.Input<number>;
    /**
     * Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
     */
    negRespTime?: pulumi.Input<number>;
    /**
     * Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
     */
    revalPnc?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
