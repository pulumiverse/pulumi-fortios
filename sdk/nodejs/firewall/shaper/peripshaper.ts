// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Configure per-IP traffic shaper.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.shaper.Peripshaper("trname", {
 *     bandwidthUnit: "kbps",
 *     diffservForward: "disable",
 *     diffservReverse: "disable",
 *     diffservcodeForward: "000000",
 *     diffservcodeRev: "000000",
 *     maxBandwidth: 1024,
 *     maxConcurrentSession: 33,
 * });
 * ```
 *
 * ## Import
 *
 * FirewallShaper PerIpShaper can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/shaper/peripshaper:Peripshaper labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/shaper/peripshaper:Peripshaper labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Peripshaper extends pulumi.CustomResource {
    /**
     * Get an existing Peripshaper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PeripshaperState, opts?: pulumi.CustomResourceOptions): Peripshaper {
        return new Peripshaper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/shaper/peripshaper:Peripshaper';

    /**
     * Returns true if the given object is an instance of Peripshaper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Peripshaper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Peripshaper.__pulumiType;
    }

    /**
     * Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
     */
    public readonly bandwidthUnit!: pulumi.Output<string>;
    /**
     * Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
     */
    public readonly diffservForward!: pulumi.Output<string>;
    /**
     * Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
     */
    public readonly diffservReverse!: pulumi.Output<string>;
    /**
     * Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
     */
    public readonly diffservcodeForward!: pulumi.Output<string>;
    /**
     * Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
     */
    public readonly diffservcodeRev!: pulumi.Output<string>;
    /**
     * Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.14, 7.0.6-7.0.13, >= 7.2.1: 0 - 80000000.
     */
    public readonly maxBandwidth!: pulumi.Output<number>;
    /**
     * Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
     */
    public readonly maxConcurrentSession!: pulumi.Output<number>;
    /**
     * Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
     */
    public readonly maxConcurrentTcpSession!: pulumi.Output<number>;
    /**
     * Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
     */
    public readonly maxConcurrentUdpSession!: pulumi.Output<number>;
    /**
     * Traffic shaper name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Peripshaper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PeripshaperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PeripshaperArgs | PeripshaperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PeripshaperState | undefined;
            resourceInputs["bandwidthUnit"] = state ? state.bandwidthUnit : undefined;
            resourceInputs["diffservForward"] = state ? state.diffservForward : undefined;
            resourceInputs["diffservReverse"] = state ? state.diffservReverse : undefined;
            resourceInputs["diffservcodeForward"] = state ? state.diffservcodeForward : undefined;
            resourceInputs["diffservcodeRev"] = state ? state.diffservcodeRev : undefined;
            resourceInputs["maxBandwidth"] = state ? state.maxBandwidth : undefined;
            resourceInputs["maxConcurrentSession"] = state ? state.maxConcurrentSession : undefined;
            resourceInputs["maxConcurrentTcpSession"] = state ? state.maxConcurrentTcpSession : undefined;
            resourceInputs["maxConcurrentUdpSession"] = state ? state.maxConcurrentUdpSession : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as PeripshaperArgs | undefined;
            resourceInputs["bandwidthUnit"] = args ? args.bandwidthUnit : undefined;
            resourceInputs["diffservForward"] = args ? args.diffservForward : undefined;
            resourceInputs["diffservReverse"] = args ? args.diffservReverse : undefined;
            resourceInputs["diffservcodeForward"] = args ? args.diffservcodeForward : undefined;
            resourceInputs["diffservcodeRev"] = args ? args.diffservcodeRev : undefined;
            resourceInputs["maxBandwidth"] = args ? args.maxBandwidth : undefined;
            resourceInputs["maxConcurrentSession"] = args ? args.maxConcurrentSession : undefined;
            resourceInputs["maxConcurrentTcpSession"] = args ? args.maxConcurrentTcpSession : undefined;
            resourceInputs["maxConcurrentUdpSession"] = args ? args.maxConcurrentUdpSession : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Peripshaper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Peripshaper resources.
 */
export interface PeripshaperState {
    /**
     * Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
     */
    bandwidthUnit?: pulumi.Input<string>;
    /**
     * Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
     */
    diffservForward?: pulumi.Input<string>;
    /**
     * Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
     */
    diffservReverse?: pulumi.Input<string>;
    /**
     * Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
     */
    diffservcodeForward?: pulumi.Input<string>;
    /**
     * Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
     */
    diffservcodeRev?: pulumi.Input<string>;
    /**
     * Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.14, 7.0.6-7.0.13, >= 7.2.1: 0 - 80000000.
     */
    maxBandwidth?: pulumi.Input<number>;
    /**
     * Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
     */
    maxConcurrentSession?: pulumi.Input<number>;
    /**
     * Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
     */
    maxConcurrentTcpSession?: pulumi.Input<number>;
    /**
     * Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
     */
    maxConcurrentUdpSession?: pulumi.Input<number>;
    /**
     * Traffic shaper name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Peripshaper resource.
 */
export interface PeripshaperArgs {
    /**
     * Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
     */
    bandwidthUnit?: pulumi.Input<string>;
    /**
     * Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
     */
    diffservForward?: pulumi.Input<string>;
    /**
     * Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
     */
    diffservReverse?: pulumi.Input<string>;
    /**
     * Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
     */
    diffservcodeForward?: pulumi.Input<string>;
    /**
     * Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
     */
    diffservcodeRev?: pulumi.Input<string>;
    /**
     * Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.14, 7.0.6-7.0.13, >= 7.2.1: 0 - 80000000.
     */
    maxBandwidth?: pulumi.Input<number>;
    /**
     * Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
     */
    maxConcurrentSession?: pulumi.Input<number>;
    /**
     * Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
     */
    maxConcurrentTcpSession?: pulumi.Input<number>;
    /**
     * Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
     */
    maxConcurrentUdpSession?: pulumi.Input<number>;
    /**
     * Traffic shaper name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
