// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Use this data source to get information on an fortios firewallservice custom
 */
export function getCustom(args: GetCustomArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:firewall/service/getCustom:getCustom", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getCustom.
 */
export interface GetCustomArgs {
    /**
     * Specify the name of the desired firewallservice custom.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getCustom.
 */
export interface GetCustomResult {
    /**
     * Application category ID. The structure of `appCategory` block is documented below.
     */
    readonly appCategories: outputs.firewall.service.GetCustomAppCategory[];
    /**
     * Application service type.
     */
    readonly appServiceType: string;
    /**
     * Application ID. The structure of `application` block is documented below.
     */
    readonly applications: outputs.firewall.service.GetCustomApplication[];
    /**
     * Service category.
     */
    readonly category: string;
    /**
     * Configure the type of ICMP error message verification.
     */
    readonly checkResetRange: string;
    /**
     * Color of icon on the GUI.
     */
    readonly color: number;
    /**
     * Comment.
     */
    readonly comment: string;
    /**
     * Security Fabric global object setting.
     */
    readonly fabricObject: string;
    /**
     * Fully qualified domain name.
     */
    readonly fqdn: string;
    /**
     * Helper name.
     */
    readonly helper: string;
    /**
     * ICMP code.
     */
    readonly icmpcode: number;
    /**
     * ICMP type.
     */
    readonly icmptype: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Start and end of the IP range associated with service.
     */
    readonly iprange: string;
    /**
     * Custom service name.
     */
    readonly name: string;
    /**
     * Protocol type based on IANA numbers.
     */
    readonly protocol: string;
    /**
     * IP protocol number.
     */
    readonly protocolNumber: number;
    /**
     * Enable/disable web proxy service.
     */
    readonly proxy: string;
    /**
     * Multiple SCTP port ranges.
     */
    readonly sctpPortrange: string;
    /**
     * Session TTL (300 - 604800, 0 = default).
     */
    readonly sessionTtl: number;
    /**
     * Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
     */
    readonly tcpHalfcloseTimer: number;
    /**
     * Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
     */
    readonly tcpHalfopenTimer: number;
    /**
     * Multiple TCP port ranges.
     */
    readonly tcpPortrange: string;
    /**
     * Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
     */
    readonly tcpRstTimer: number;
    /**
     * Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
     */
    readonly tcpTimewaitTimer: number;
    /**
     * UDP half close timeout (0 - 86400 sec, 0 = default).
     */
    readonly udpIdleTimer: number;
    /**
     * Multiple UDP port ranges.
     */
    readonly udpPortrange: string;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    readonly uuid: string;
    readonly vdomparam?: string;
    /**
     * Enable/disable the visibility of the service on the GUI.
     */
    readonly visibility: string;
}
/**
 * Use this data source to get information on an fortios firewallservice custom
 */
export function getCustomOutput(args: GetCustomOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCustomResult> {
    return pulumi.output(args).apply((a: any) => getCustom(a, opts))
}

/**
 * A collection of arguments for invoking getCustom.
 */
export interface GetCustomOutputArgs {
    /**
     * Specify the name of the desired firewallservice custom.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
