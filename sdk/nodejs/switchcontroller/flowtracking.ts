// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure FortiSwitch flow tracking and export via ipfix/netflow. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * SwitchController FlowTracking can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/flowtracking:Flowtracking labelname SwitchControllerFlowTracking
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/flowtracking:Flowtracking labelname SwitchControllerFlowTracking
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Flowtracking extends pulumi.CustomResource {
    /**
     * Get an existing Flowtracking resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlowtrackingState, opts?: pulumi.CustomResourceOptions): Flowtracking {
        return new Flowtracking(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/flowtracking:Flowtracking';

    /**
     * Returns true if the given object is an instance of Flowtracking.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Flowtracking {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Flowtracking.__pulumiType;
    }

    /**
     * Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
     */
    public readonly aggregates!: pulumi.Output<outputs.switchcontroller.FlowtrackingAggregate[] | undefined>;
    /**
     * Configure collector ip address.
     */
    public readonly collectorIp!: pulumi.Output<string>;
    /**
     * Configure collector port number(0-65535, default=0).
     */
    public readonly collectorPort!: pulumi.Output<number>;
    /**
     * Configure collectors for the flow. The structure of `collectors` block is documented below.
     */
    public readonly collectors!: pulumi.Output<outputs.switchcontroller.FlowtrackingCollector[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
     */
    public readonly format!: pulumi.Output<string>;
    /**
     * Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
     */
    public readonly level!: pulumi.Output<string>;
    /**
     * Configure flow max export packet size (512-9216, default=512 bytes).
     */
    public readonly maxExportPktSize!: pulumi.Output<number>;
    /**
     * Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
     */
    public readonly sampleMode!: pulumi.Output<string>;
    /**
     * Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
     */
    public readonly sampleRate!: pulumi.Output<number>;
    /**
     * Configure template export period (1-60, default=5 minutes).
     */
    public readonly templateExportPeriod!: pulumi.Output<number>;
    /**
     * Configure flow session general timeout (60-604800, default=3600 seconds).
     */
    public readonly timeoutGeneral!: pulumi.Output<number>;
    /**
     * Configure flow session ICMP timeout (60-604800, default=300 seconds).
     */
    public readonly timeoutIcmp!: pulumi.Output<number>;
    /**
     * Configure flow session max timeout (60-604800, default=604800 seconds).
     */
    public readonly timeoutMax!: pulumi.Output<number>;
    /**
     * Configure flow session TCP timeout (60-604800, default=3600 seconds).
     */
    public readonly timeoutTcp!: pulumi.Output<number>;
    /**
     * Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
     */
    public readonly timeoutTcpFin!: pulumi.Output<number>;
    /**
     * Configure flow session TCP RST timeout (60-604800, default=120 seconds).
     */
    public readonly timeoutTcpRst!: pulumi.Output<number>;
    /**
     * Configure flow session UDP timeout (60-604800, default=300 seconds).
     */
    public readonly timeoutUdp!: pulumi.Output<number>;
    /**
     * Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
     */
    public readonly transport!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Flowtracking resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FlowtrackingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FlowtrackingArgs | FlowtrackingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FlowtrackingState | undefined;
            resourceInputs["aggregates"] = state ? state.aggregates : undefined;
            resourceInputs["collectorIp"] = state ? state.collectorIp : undefined;
            resourceInputs["collectorPort"] = state ? state.collectorPort : undefined;
            resourceInputs["collectors"] = state ? state.collectors : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["level"] = state ? state.level : undefined;
            resourceInputs["maxExportPktSize"] = state ? state.maxExportPktSize : undefined;
            resourceInputs["sampleMode"] = state ? state.sampleMode : undefined;
            resourceInputs["sampleRate"] = state ? state.sampleRate : undefined;
            resourceInputs["templateExportPeriod"] = state ? state.templateExportPeriod : undefined;
            resourceInputs["timeoutGeneral"] = state ? state.timeoutGeneral : undefined;
            resourceInputs["timeoutIcmp"] = state ? state.timeoutIcmp : undefined;
            resourceInputs["timeoutMax"] = state ? state.timeoutMax : undefined;
            resourceInputs["timeoutTcp"] = state ? state.timeoutTcp : undefined;
            resourceInputs["timeoutTcpFin"] = state ? state.timeoutTcpFin : undefined;
            resourceInputs["timeoutTcpRst"] = state ? state.timeoutTcpRst : undefined;
            resourceInputs["timeoutUdp"] = state ? state.timeoutUdp : undefined;
            resourceInputs["transport"] = state ? state.transport : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FlowtrackingArgs | undefined;
            resourceInputs["aggregates"] = args ? args.aggregates : undefined;
            resourceInputs["collectorIp"] = args ? args.collectorIp : undefined;
            resourceInputs["collectorPort"] = args ? args.collectorPort : undefined;
            resourceInputs["collectors"] = args ? args.collectors : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["level"] = args ? args.level : undefined;
            resourceInputs["maxExportPktSize"] = args ? args.maxExportPktSize : undefined;
            resourceInputs["sampleMode"] = args ? args.sampleMode : undefined;
            resourceInputs["sampleRate"] = args ? args.sampleRate : undefined;
            resourceInputs["templateExportPeriod"] = args ? args.templateExportPeriod : undefined;
            resourceInputs["timeoutGeneral"] = args ? args.timeoutGeneral : undefined;
            resourceInputs["timeoutIcmp"] = args ? args.timeoutIcmp : undefined;
            resourceInputs["timeoutMax"] = args ? args.timeoutMax : undefined;
            resourceInputs["timeoutTcp"] = args ? args.timeoutTcp : undefined;
            resourceInputs["timeoutTcpFin"] = args ? args.timeoutTcpFin : undefined;
            resourceInputs["timeoutTcpRst"] = args ? args.timeoutTcpRst : undefined;
            resourceInputs["timeoutUdp"] = args ? args.timeoutUdp : undefined;
            resourceInputs["transport"] = args ? args.transport : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Flowtracking.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Flowtracking resources.
 */
export interface FlowtrackingState {
    /**
     * Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
     */
    aggregates?: pulumi.Input<pulumi.Input<inputs.switchcontroller.FlowtrackingAggregate>[]>;
    /**
     * Configure collector ip address.
     */
    collectorIp?: pulumi.Input<string>;
    /**
     * Configure collector port number(0-65535, default=0).
     */
    collectorPort?: pulumi.Input<number>;
    /**
     * Configure collectors for the flow. The structure of `collectors` block is documented below.
     */
    collectors?: pulumi.Input<pulumi.Input<inputs.switchcontroller.FlowtrackingCollector>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
     */
    format?: pulumi.Input<string>;
    /**
     * Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
     */
    level?: pulumi.Input<string>;
    /**
     * Configure flow max export packet size (512-9216, default=512 bytes).
     */
    maxExportPktSize?: pulumi.Input<number>;
    /**
     * Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
     */
    sampleMode?: pulumi.Input<string>;
    /**
     * Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
     */
    sampleRate?: pulumi.Input<number>;
    /**
     * Configure template export period (1-60, default=5 minutes).
     */
    templateExportPeriod?: pulumi.Input<number>;
    /**
     * Configure flow session general timeout (60-604800, default=3600 seconds).
     */
    timeoutGeneral?: pulumi.Input<number>;
    /**
     * Configure flow session ICMP timeout (60-604800, default=300 seconds).
     */
    timeoutIcmp?: pulumi.Input<number>;
    /**
     * Configure flow session max timeout (60-604800, default=604800 seconds).
     */
    timeoutMax?: pulumi.Input<number>;
    /**
     * Configure flow session TCP timeout (60-604800, default=3600 seconds).
     */
    timeoutTcp?: pulumi.Input<number>;
    /**
     * Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
     */
    timeoutTcpFin?: pulumi.Input<number>;
    /**
     * Configure flow session TCP RST timeout (60-604800, default=120 seconds).
     */
    timeoutTcpRst?: pulumi.Input<number>;
    /**
     * Configure flow session UDP timeout (60-604800, default=300 seconds).
     */
    timeoutUdp?: pulumi.Input<number>;
    /**
     * Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
     */
    transport?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Flowtracking resource.
 */
export interface FlowtrackingArgs {
    /**
     * Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
     */
    aggregates?: pulumi.Input<pulumi.Input<inputs.switchcontroller.FlowtrackingAggregate>[]>;
    /**
     * Configure collector ip address.
     */
    collectorIp?: pulumi.Input<string>;
    /**
     * Configure collector port number(0-65535, default=0).
     */
    collectorPort?: pulumi.Input<number>;
    /**
     * Configure collectors for the flow. The structure of `collectors` block is documented below.
     */
    collectors?: pulumi.Input<pulumi.Input<inputs.switchcontroller.FlowtrackingCollector>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
     */
    format?: pulumi.Input<string>;
    /**
     * Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
     */
    level?: pulumi.Input<string>;
    /**
     * Configure flow max export packet size (512-9216, default=512 bytes).
     */
    maxExportPktSize?: pulumi.Input<number>;
    /**
     * Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
     */
    sampleMode?: pulumi.Input<string>;
    /**
     * Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
     */
    sampleRate?: pulumi.Input<number>;
    /**
     * Configure template export period (1-60, default=5 minutes).
     */
    templateExportPeriod?: pulumi.Input<number>;
    /**
     * Configure flow session general timeout (60-604800, default=3600 seconds).
     */
    timeoutGeneral?: pulumi.Input<number>;
    /**
     * Configure flow session ICMP timeout (60-604800, default=300 seconds).
     */
    timeoutIcmp?: pulumi.Input<number>;
    /**
     * Configure flow session max timeout (60-604800, default=604800 seconds).
     */
    timeoutMax?: pulumi.Input<number>;
    /**
     * Configure flow session TCP timeout (60-604800, default=3600 seconds).
     */
    timeoutTcp?: pulumi.Input<number>;
    /**
     * Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
     */
    timeoutTcpFin?: pulumi.Input<number>;
    /**
     * Configure flow session TCP RST timeout (60-604800, default=120 seconds).
     */
    timeoutTcpRst?: pulumi.Input<number>;
    /**
     * Configure flow session UDP timeout (60-604800, default=300 seconds).
     */
    timeoutUdp?: pulumi.Input<number>;
    /**
     * Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
     */
    transport?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
