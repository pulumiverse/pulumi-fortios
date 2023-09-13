// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on fortios router ospf6
 */
export function getOspf6(args?: GetOspf6Args, opts?: pulumi.InvokeOptions): Promise<GetOspf6Result> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:router/getOspf6:getOspf6", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getOspf6.
 */
export interface GetOspf6Args {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getOspf6.
 */
export interface GetOspf6Result {
    /**
     * Area border router type.
     */
    readonly abrType: string;
    /**
     * OSPF6 area configuration. The structure of `area` block is documented below.
     */
    readonly areas: outputs.router.GetOspf6Area[];
    /**
     * Reference bandwidth in terms of megabits per second.
     */
    readonly autoCostRefBandwidth: number;
    /**
     * Enable/disable Bidirectional Forwarding Detection (BFD).
     */
    readonly bfd: string;
    /**
     * Default information metric.
     */
    readonly defaultInformationMetric: number;
    /**
     * Default information metric type.
     */
    readonly defaultInformationMetricType: string;
    /**
     * Enable/disable generation of default route.
     */
    readonly defaultInformationOriginate: string;
    /**
     * Default information route map.
     */
    readonly defaultInformationRouteMap: string;
    /**
     * Default metric of redistribute routes.
     */
    readonly defaultMetric: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Enable logging of OSPFv3 neighbour's changes
     */
    readonly logNeighbourChanges: string;
    /**
     * OSPF6 interface configuration. The structure of `ospf6Interface` block is documented below.
     */
    readonly ospf6Interfaces: outputs.router.GetOspf6Ospf6Interface[];
    /**
     * Passive interface configuration. The structure of `passiveInterface` block is documented below.
     */
    readonly passiveInterfaces: outputs.router.GetOspf6PassiveInterface[];
    /**
     * Redistribute configuration. The structure of `redistribute` block is documented below.
     */
    readonly redistributes: outputs.router.GetOspf6Redistribute[];
    /**
     * OSPFv3 restart mode (graceful or none).
     */
    readonly restartMode: string;
    /**
     * Enable/disable continuing graceful restart upon topology change.
     */
    readonly restartOnTopologyChange: string;
    /**
     * Graceful restart period in seconds.
     */
    readonly restartPeriod: number;
    /**
     * A.B.C.D, in IPv4 address format.
     */
    readonly routerId: string;
    /**
     * SPF calculation frequency.
     */
    readonly spfTimers: string;
    /**
     * IPv6 address summary configuration. The structure of `summaryAddress` block is documented below.
     */
    readonly summaryAddresses: outputs.router.GetOspf6SummaryAddress[];
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on fortios router ospf6
 */
export function getOspf6Output(args?: GetOspf6OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOspf6Result> {
    return pulumi.output(args).apply((a: any) => getOspf6(a, opts))
}

/**
 * A collection of arguments for invoking getOspf6.
 */
export interface GetOspf6OutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
