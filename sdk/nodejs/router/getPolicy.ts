// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios router policy
 */
export function getPolicy(args: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:router/getPolicy:getPolicy", {
        "seqNum": args.seqNum,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyArgs {
    /**
     * Specify the seqNum of the desired router policy.
     */
    seqNum: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getPolicy.
 */
export interface GetPolicyResult {
    /**
     * Action of the policy route.
     */
    readonly action: string;
    /**
     * Optional comments.
     */
    readonly comments: string;
    /**
     * Enable/disable negating destination address match.
     */
    readonly dstNegate: string;
    /**
     * Destination address name. The structure of `dstaddr` block is documented below.
     */
    readonly dstaddrs: outputs.router.GetPolicyDstaddr[];
    /**
     * Destination IP and mask (x.x.x.x/x). The structure of `dst` block is documented below.
     */
    readonly dsts: outputs.router.GetPolicyDst[];
    /**
     * End destination port number (0 - 65535).
     */
    readonly endPort: number;
    /**
     * End source port number (0 - 65535).
     */
    readonly endSourcePort: number;
    /**
     * IP address of the gateway.
     */
    readonly gateway: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Enable/disable negation of input device match.
     */
    readonly inputDeviceNegate: string;
    /**
     * Incoming interface name. The structure of `inputDevice` block is documented below.
     */
    readonly inputDevices: outputs.router.GetPolicyInputDevice[];
    /**
     * Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
     */
    readonly internetServiceCustoms: outputs.router.GetPolicyInternetServiceCustom[];
    /**
     * Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
     */
    readonly internetServiceIds: outputs.router.GetPolicyInternetServiceId[];
    /**
     * Outgoing interface name.
     */
    readonly outputDevice: string;
    /**
     * Protocol number (0 - 255).
     */
    readonly protocol: number;
    /**
     * Sequence number.
     */
    readonly seqNum: number;
    /**
     * Enable/disable negating source address match.
     */
    readonly srcNegate: string;
    /**
     * Source address name. The structure of `srcaddr` block is documented below.
     */
    readonly srcaddrs: outputs.router.GetPolicySrcaddr[];
    /**
     * Source IP and mask (x.x.x.x/x). The structure of `src` block is documented below.
     */
    readonly srcs: outputs.router.GetPolicySrc[];
    /**
     * Start destination port number (0 - 65535).
     */
    readonly startPort: number;
    /**
     * Start source port number (0 - 65535).
     */
    readonly startSourcePort: number;
    /**
     * Enable/disable this policy route.
     */
    readonly status: string;
    /**
     * Type of service bit pattern.
     */
    readonly tos: string;
    /**
     * Type of service evaluated bits.
     */
    readonly tosMask: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios router policy
 */
export function getPolicyOutput(args: GetPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicyResult> {
    return pulumi.output(args).apply((a: any) => getPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyOutputArgs {
    /**
     * Specify the seqNum of the desired router policy.
     */
    seqNum: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
