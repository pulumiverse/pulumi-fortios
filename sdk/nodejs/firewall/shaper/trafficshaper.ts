// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Configure shared traffic shaper.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.shaper.Trafficshaper("trname", {
 *     bandwidthUnit: "kbps",
 *     diffserv: "disable",
 *     diffservcode: "000000",
 *     guaranteedBandwidth: 0,
 *     maximumBandwidth: 1024,
 *     perPolicy: "disable",
 *     priority: "low",
 * });
 * ```
 *
 * ## Import
 *
 * FirewallShaper TrafficShaper can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/shaper/trafficshaper:Trafficshaper labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/shaper/trafficshaper:Trafficshaper labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Trafficshaper extends pulumi.CustomResource {
    /**
     * Get an existing Trafficshaper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TrafficshaperState, opts?: pulumi.CustomResourceOptions): Trafficshaper {
        return new Trafficshaper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/shaper/trafficshaper:Trafficshaper';

    /**
     * Returns true if the given object is an instance of Trafficshaper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Trafficshaper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Trafficshaper.__pulumiType;
    }

    /**
     * Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
     */
    public readonly bandwidthUnit!: pulumi.Output<string>;
    /**
     * VLAN CoS mark.
     */
    public readonly cos!: pulumi.Output<string>;
    /**
     * Enable/disable VLAN CoS marking. Valid values: `enable`, `disable`.
     */
    public readonly cosMarking!: pulumi.Output<string>;
    /**
     * Select VLAN CoS marking method. Valid values: `multi-stage`, `static`.
     */
    public readonly cosMarkingMethod!: pulumi.Output<string>;
    /**
     * Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
     */
    public readonly diffserv!: pulumi.Output<string>;
    /**
     * DiffServ setting to be applied to traffic accepted by this shaper.
     */
    public readonly diffservcode!: pulumi.Output<string>;
    /**
     * Select DSCP marking method. Valid values: `multi-stage`, `static`.
     */
    public readonly dscpMarkingMethod!: pulumi.Output<string>;
    /**
     * Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
     */
    public readonly exceedBandwidth!: pulumi.Output<number>;
    /**
     * Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].
     */
    public readonly exceedClassId!: pulumi.Output<number>;
    /**
     * VLAN CoS mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
     */
    public readonly exceedCos!: pulumi.Output<string>;
    /**
     * DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
     */
    public readonly exceedDscp!: pulumi.Output<string>;
    /**
     * Amount of bandwidth guaranteed for this shaper. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
     */
    public readonly guaranteedBandwidth!: pulumi.Output<number>;
    /**
     * Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
     */
    public readonly maximumBandwidth!: pulumi.Output<number>;
    /**
     * VLAN CoS mark for traffic in [exceed-bandwidth, maximum-bandwidth].
     */
    public readonly maximumCos!: pulumi.Output<string>;
    /**
     * DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
     */
    public readonly maximumDscp!: pulumi.Output<string>;
    /**
     * Traffic shaper name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Per-packet size overhead used in rate computations.
     */
    public readonly overhead!: pulumi.Output<number>;
    /**
     * Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy. Valid values: `disable`, `enable`.
     */
    public readonly perPolicy!: pulumi.Output<string>;
    /**
     * Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth. Valid values: `low`, `medium`, `high`.
     */
    public readonly priority!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Trafficshaper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TrafficshaperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TrafficshaperArgs | TrafficshaperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TrafficshaperState | undefined;
            resourceInputs["bandwidthUnit"] = state ? state.bandwidthUnit : undefined;
            resourceInputs["cos"] = state ? state.cos : undefined;
            resourceInputs["cosMarking"] = state ? state.cosMarking : undefined;
            resourceInputs["cosMarkingMethod"] = state ? state.cosMarkingMethod : undefined;
            resourceInputs["diffserv"] = state ? state.diffserv : undefined;
            resourceInputs["diffservcode"] = state ? state.diffservcode : undefined;
            resourceInputs["dscpMarkingMethod"] = state ? state.dscpMarkingMethod : undefined;
            resourceInputs["exceedBandwidth"] = state ? state.exceedBandwidth : undefined;
            resourceInputs["exceedClassId"] = state ? state.exceedClassId : undefined;
            resourceInputs["exceedCos"] = state ? state.exceedCos : undefined;
            resourceInputs["exceedDscp"] = state ? state.exceedDscp : undefined;
            resourceInputs["guaranteedBandwidth"] = state ? state.guaranteedBandwidth : undefined;
            resourceInputs["maximumBandwidth"] = state ? state.maximumBandwidth : undefined;
            resourceInputs["maximumCos"] = state ? state.maximumCos : undefined;
            resourceInputs["maximumDscp"] = state ? state.maximumDscp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["overhead"] = state ? state.overhead : undefined;
            resourceInputs["perPolicy"] = state ? state.perPolicy : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as TrafficshaperArgs | undefined;
            resourceInputs["bandwidthUnit"] = args ? args.bandwidthUnit : undefined;
            resourceInputs["cos"] = args ? args.cos : undefined;
            resourceInputs["cosMarking"] = args ? args.cosMarking : undefined;
            resourceInputs["cosMarkingMethod"] = args ? args.cosMarkingMethod : undefined;
            resourceInputs["diffserv"] = args ? args.diffserv : undefined;
            resourceInputs["diffservcode"] = args ? args.diffservcode : undefined;
            resourceInputs["dscpMarkingMethod"] = args ? args.dscpMarkingMethod : undefined;
            resourceInputs["exceedBandwidth"] = args ? args.exceedBandwidth : undefined;
            resourceInputs["exceedClassId"] = args ? args.exceedClassId : undefined;
            resourceInputs["exceedCos"] = args ? args.exceedCos : undefined;
            resourceInputs["exceedDscp"] = args ? args.exceedDscp : undefined;
            resourceInputs["guaranteedBandwidth"] = args ? args.guaranteedBandwidth : undefined;
            resourceInputs["maximumBandwidth"] = args ? args.maximumBandwidth : undefined;
            resourceInputs["maximumCos"] = args ? args.maximumCos : undefined;
            resourceInputs["maximumDscp"] = args ? args.maximumDscp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overhead"] = args ? args.overhead : undefined;
            resourceInputs["perPolicy"] = args ? args.perPolicy : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Trafficshaper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trafficshaper resources.
 */
export interface TrafficshaperState {
    /**
     * Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
     */
    bandwidthUnit?: pulumi.Input<string>;
    /**
     * VLAN CoS mark.
     */
    cos?: pulumi.Input<string>;
    /**
     * Enable/disable VLAN CoS marking. Valid values: `enable`, `disable`.
     */
    cosMarking?: pulumi.Input<string>;
    /**
     * Select VLAN CoS marking method. Valid values: `multi-stage`, `static`.
     */
    cosMarkingMethod?: pulumi.Input<string>;
    /**
     * Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
     */
    diffserv?: pulumi.Input<string>;
    /**
     * DiffServ setting to be applied to traffic accepted by this shaper.
     */
    diffservcode?: pulumi.Input<string>;
    /**
     * Select DSCP marking method. Valid values: `multi-stage`, `static`.
     */
    dscpMarkingMethod?: pulumi.Input<string>;
    /**
     * Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
     */
    exceedBandwidth?: pulumi.Input<number>;
    /**
     * Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].
     */
    exceedClassId?: pulumi.Input<number>;
    /**
     * VLAN CoS mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
     */
    exceedCos?: pulumi.Input<string>;
    /**
     * DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
     */
    exceedDscp?: pulumi.Input<string>;
    /**
     * Amount of bandwidth guaranteed for this shaper. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
     */
    guaranteedBandwidth?: pulumi.Input<number>;
    /**
     * Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
     */
    maximumBandwidth?: pulumi.Input<number>;
    /**
     * VLAN CoS mark for traffic in [exceed-bandwidth, maximum-bandwidth].
     */
    maximumCos?: pulumi.Input<string>;
    /**
     * DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
     */
    maximumDscp?: pulumi.Input<string>;
    /**
     * Traffic shaper name.
     */
    name?: pulumi.Input<string>;
    /**
     * Per-packet size overhead used in rate computations.
     */
    overhead?: pulumi.Input<number>;
    /**
     * Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy. Valid values: `disable`, `enable`.
     */
    perPolicy?: pulumi.Input<string>;
    /**
     * Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth. Valid values: `low`, `medium`, `high`.
     */
    priority?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Trafficshaper resource.
 */
export interface TrafficshaperArgs {
    /**
     * Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
     */
    bandwidthUnit?: pulumi.Input<string>;
    /**
     * VLAN CoS mark.
     */
    cos?: pulumi.Input<string>;
    /**
     * Enable/disable VLAN CoS marking. Valid values: `enable`, `disable`.
     */
    cosMarking?: pulumi.Input<string>;
    /**
     * Select VLAN CoS marking method. Valid values: `multi-stage`, `static`.
     */
    cosMarkingMethod?: pulumi.Input<string>;
    /**
     * Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
     */
    diffserv?: pulumi.Input<string>;
    /**
     * DiffServ setting to be applied to traffic accepted by this shaper.
     */
    diffservcode?: pulumi.Input<string>;
    /**
     * Select DSCP marking method. Valid values: `multi-stage`, `static`.
     */
    dscpMarkingMethod?: pulumi.Input<string>;
    /**
     * Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
     */
    exceedBandwidth?: pulumi.Input<number>;
    /**
     * Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].
     */
    exceedClassId?: pulumi.Input<number>;
    /**
     * VLAN CoS mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
     */
    exceedCos?: pulumi.Input<string>;
    /**
     * DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
     */
    exceedDscp?: pulumi.Input<string>;
    /**
     * Amount of bandwidth guaranteed for this shaper. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
     */
    guaranteedBandwidth?: pulumi.Input<number>;
    /**
     * Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
     */
    maximumBandwidth?: pulumi.Input<number>;
    /**
     * VLAN CoS mark for traffic in [exceed-bandwidth, maximum-bandwidth].
     */
    maximumCos?: pulumi.Input<string>;
    /**
     * DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
     */
    maximumDscp?: pulumi.Input<string>;
    /**
     * Traffic shaper name.
     */
    name?: pulumi.Input<string>;
    /**
     * Per-packet size overhead used in rate computations.
     */
    overhead?: pulumi.Input<number>;
    /**
     * Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy. Valid values: `disable`, `enable`.
     */
    perPolicy?: pulumi.Input<string>;
    /**
     * Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth. Valid values: `low`, `medium`, `high`.
     */
    priority?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
