// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure IPv4 IP pools.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Ippool("trname", {
 *     arpReply: "enable",
 *     blockSize: 128,
 *     endip: "1.0.0.20",
 *     numBlocksPerUser: 8,
 *     pbaTimeout: 30,
 *     permitAnyHost: "disable",
 *     sourceEndip: "0.0.0.0",
 *     sourceStartip: "0.0.0.0",
 *     startip: "1.0.0.0",
 *     type: "overload",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Firewall Ippool can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/ippool:Ippool labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/ippool:Ippool labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ippool extends pulumi.CustomResource {
    /**
     * Get an existing Ippool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IppoolState, opts?: pulumi.CustomResourceOptions): Ippool {
        return new Ippool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/ippool:Ippool';

    /**
     * Returns true if the given object is an instance of Ippool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ippool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ippool.__pulumiType;
    }

    /**
     * Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.
     */
    public readonly addNat64Route!: pulumi.Output<string>;
    /**
     * Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
     */
    public readonly arpIntf!: pulumi.Output<string>;
    /**
     * Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
     */
    public readonly arpReply!: pulumi.Output<string>;
    /**
     * Associated interface name.
     */
    public readonly associatedInterface!: pulumi.Output<string>;
    /**
     * Number of addresses in a block (64 - 4096, default = 128).
     */
    public readonly blockSize!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
     */
    public readonly endip!: pulumi.Output<string>;
    /**
     * Final port number (inclusive) in the range for the address pool (Default: 65533).
     */
    public readonly endport!: pulumi.Output<number>;
    /**
     * IP pool name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable NAT64. Valid values: `disable`, `enable`.
     */
    public readonly nat64!: pulumi.Output<string>;
    /**
     * Number of addresses blocks that can be used by a user (1 to 128, default = 8).
     */
    public readonly numBlocksPerUser!: pulumi.Output<number>;
    /**
     * Port block allocation timeout (seconds).
     */
    public readonly pbaTimeout!: pulumi.Output<number>;
    /**
     * Enable/disable full cone NAT. Valid values: `disable`, `enable`.
     */
    public readonly permitAnyHost!: pulumi.Output<string>;
    /**
     * Number of port for each user (32 - 60416, default = 0, which is auto).
     */
    public readonly portPerUser!: pulumi.Output<number>;
    /**
     * Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
     */
    public readonly sourceEndip!: pulumi.Output<string>;
    /**
     * First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
     */
    public readonly sourceStartip!: pulumi.Output<string>;
    /**
     * First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
     */
    public readonly startip!: pulumi.Output<string>;
    /**
     * First port number (inclusive) in the range for the address pool (Default: 5117).
     */
    public readonly startport!: pulumi.Output<number>;
    /**
     * Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.
     */
    public readonly subnetBroadcastInIppool!: pulumi.Output<string>;
    /**
     * IP pool type. On FortiOS versions 6.2.0-7.4.1: overload, one-to-one, fixed port range, or port block allocation. On FortiOS versions >= 7.4.2: overload, one-to-one, fixed-port-range, port-block-allocation, cgn-resource-allocation (hyperscale vdom only). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Ippool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IppoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IppoolArgs | IppoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IppoolState | undefined;
            resourceInputs["addNat64Route"] = state ? state.addNat64Route : undefined;
            resourceInputs["arpIntf"] = state ? state.arpIntf : undefined;
            resourceInputs["arpReply"] = state ? state.arpReply : undefined;
            resourceInputs["associatedInterface"] = state ? state.associatedInterface : undefined;
            resourceInputs["blockSize"] = state ? state.blockSize : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["endip"] = state ? state.endip : undefined;
            resourceInputs["endport"] = state ? state.endport : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nat64"] = state ? state.nat64 : undefined;
            resourceInputs["numBlocksPerUser"] = state ? state.numBlocksPerUser : undefined;
            resourceInputs["pbaTimeout"] = state ? state.pbaTimeout : undefined;
            resourceInputs["permitAnyHost"] = state ? state.permitAnyHost : undefined;
            resourceInputs["portPerUser"] = state ? state.portPerUser : undefined;
            resourceInputs["sourceEndip"] = state ? state.sourceEndip : undefined;
            resourceInputs["sourceStartip"] = state ? state.sourceStartip : undefined;
            resourceInputs["startip"] = state ? state.startip : undefined;
            resourceInputs["startport"] = state ? state.startport : undefined;
            resourceInputs["subnetBroadcastInIppool"] = state ? state.subnetBroadcastInIppool : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as IppoolArgs | undefined;
            if ((!args || args.endip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endip'");
            }
            if ((!args || args.startip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startip'");
            }
            resourceInputs["addNat64Route"] = args ? args.addNat64Route : undefined;
            resourceInputs["arpIntf"] = args ? args.arpIntf : undefined;
            resourceInputs["arpReply"] = args ? args.arpReply : undefined;
            resourceInputs["associatedInterface"] = args ? args.associatedInterface : undefined;
            resourceInputs["blockSize"] = args ? args.blockSize : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["endip"] = args ? args.endip : undefined;
            resourceInputs["endport"] = args ? args.endport : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nat64"] = args ? args.nat64 : undefined;
            resourceInputs["numBlocksPerUser"] = args ? args.numBlocksPerUser : undefined;
            resourceInputs["pbaTimeout"] = args ? args.pbaTimeout : undefined;
            resourceInputs["permitAnyHost"] = args ? args.permitAnyHost : undefined;
            resourceInputs["portPerUser"] = args ? args.portPerUser : undefined;
            resourceInputs["sourceEndip"] = args ? args.sourceEndip : undefined;
            resourceInputs["sourceStartip"] = args ? args.sourceStartip : undefined;
            resourceInputs["startip"] = args ? args.startip : undefined;
            resourceInputs["startport"] = args ? args.startport : undefined;
            resourceInputs["subnetBroadcastInIppool"] = args ? args.subnetBroadcastInIppool : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ippool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ippool resources.
 */
export interface IppoolState {
    /**
     * Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.
     */
    addNat64Route?: pulumi.Input<string>;
    /**
     * Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
     */
    arpIntf?: pulumi.Input<string>;
    /**
     * Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
     */
    arpReply?: pulumi.Input<string>;
    /**
     * Associated interface name.
     */
    associatedInterface?: pulumi.Input<string>;
    /**
     * Number of addresses in a block (64 - 4096, default = 128).
     */
    blockSize?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
     */
    endip?: pulumi.Input<string>;
    /**
     * Final port number (inclusive) in the range for the address pool (Default: 65533).
     */
    endport?: pulumi.Input<number>;
    /**
     * IP pool name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable NAT64. Valid values: `disable`, `enable`.
     */
    nat64?: pulumi.Input<string>;
    /**
     * Number of addresses blocks that can be used by a user (1 to 128, default = 8).
     */
    numBlocksPerUser?: pulumi.Input<number>;
    /**
     * Port block allocation timeout (seconds).
     */
    pbaTimeout?: pulumi.Input<number>;
    /**
     * Enable/disable full cone NAT. Valid values: `disable`, `enable`.
     */
    permitAnyHost?: pulumi.Input<string>;
    /**
     * Number of port for each user (32 - 60416, default = 0, which is auto).
     */
    portPerUser?: pulumi.Input<number>;
    /**
     * Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
     */
    sourceEndip?: pulumi.Input<string>;
    /**
     * First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
     */
    sourceStartip?: pulumi.Input<string>;
    /**
     * First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
     */
    startip?: pulumi.Input<string>;
    /**
     * First port number (inclusive) in the range for the address pool (Default: 5117).
     */
    startport?: pulumi.Input<number>;
    /**
     * Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.
     */
    subnetBroadcastInIppool?: pulumi.Input<string>;
    /**
     * IP pool type. On FortiOS versions 6.2.0-7.4.1: overload, one-to-one, fixed port range, or port block allocation. On FortiOS versions >= 7.4.2: overload, one-to-one, fixed-port-range, port-block-allocation, cgn-resource-allocation (hyperscale vdom only). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ippool resource.
 */
export interface IppoolArgs {
    /**
     * Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.
     */
    addNat64Route?: pulumi.Input<string>;
    /**
     * Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
     */
    arpIntf?: pulumi.Input<string>;
    /**
     * Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
     */
    arpReply?: pulumi.Input<string>;
    /**
     * Associated interface name.
     */
    associatedInterface?: pulumi.Input<string>;
    /**
     * Number of addresses in a block (64 - 4096, default = 128).
     */
    blockSize?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
     */
    endip: pulumi.Input<string>;
    /**
     * Final port number (inclusive) in the range for the address pool (Default: 65533).
     */
    endport?: pulumi.Input<number>;
    /**
     * IP pool name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable NAT64. Valid values: `disable`, `enable`.
     */
    nat64?: pulumi.Input<string>;
    /**
     * Number of addresses blocks that can be used by a user (1 to 128, default = 8).
     */
    numBlocksPerUser?: pulumi.Input<number>;
    /**
     * Port block allocation timeout (seconds).
     */
    pbaTimeout?: pulumi.Input<number>;
    /**
     * Enable/disable full cone NAT. Valid values: `disable`, `enable`.
     */
    permitAnyHost?: pulumi.Input<string>;
    /**
     * Number of port for each user (32 - 60416, default = 0, which is auto).
     */
    portPerUser?: pulumi.Input<number>;
    /**
     * Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
     */
    sourceEndip?: pulumi.Input<string>;
    /**
     * First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
     */
    sourceStartip?: pulumi.Input<string>;
    /**
     * First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
     */
    startip: pulumi.Input<string>;
    /**
     * First port number (inclusive) in the range for the address pool (Default: 5117).
     */
    startport?: pulumi.Input<number>;
    /**
     * Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.
     */
    subnetBroadcastInIppool?: pulumi.Input<string>;
    /**
     * IP pool type. On FortiOS versions 6.2.0-7.4.1: overload, one-to-one, fixed port range, or port block allocation. On FortiOS versions >= 7.4.2: overload, one-to-one, fixed-port-range, port-block-allocation, cgn-resource-allocation (hyperscale vdom only). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
