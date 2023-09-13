// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure wireless access gateway (WAG) profiles used for tunnels on AP. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * WirelessController WagProfile can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:wirelesscontroller/wagprofile:Wagprofile labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:wirelesscontroller/wagprofile:Wagprofile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Wagprofile extends pulumi.CustomResource {
    /**
     * Get an existing Wagprofile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WagprofileState, opts?: pulumi.CustomResourceOptions): Wagprofile {
        return new Wagprofile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:wirelesscontroller/wagprofile:Wagprofile';

    /**
     * Returns true if the given object is an instance of Wagprofile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Wagprofile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Wagprofile.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * IP address of the monitoring DHCP request packet sent through the tunnel.
     */
    public readonly dhcpIpAddr!: pulumi.Output<string>;
    /**
     * Tunnel profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
     */
    public readonly pingInterval!: pulumi.Output<number>;
    /**
     * Number of the tunnel monitoring echo packets (1 - 65535, default = 5).
     */
    public readonly pingNumber!: pulumi.Output<number>;
    /**
     * Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
     */
    public readonly returnPacketTimeout!: pulumi.Output<number>;
    /**
     * Tunnel type. Valid values: `l2tpv3`, `gre`.
     */
    public readonly tunnelType!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * IP Address of the wireless access gateway.
     */
    public readonly wagIp!: pulumi.Output<string>;
    /**
     * UDP port of the wireless access gateway.
     */
    public readonly wagPort!: pulumi.Output<number>;

    /**
     * Create a Wagprofile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WagprofileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WagprofileArgs | WagprofileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WagprofileState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dhcpIpAddr"] = state ? state.dhcpIpAddr : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pingInterval"] = state ? state.pingInterval : undefined;
            resourceInputs["pingNumber"] = state ? state.pingNumber : undefined;
            resourceInputs["returnPacketTimeout"] = state ? state.returnPacketTimeout : undefined;
            resourceInputs["tunnelType"] = state ? state.tunnelType : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["wagIp"] = state ? state.wagIp : undefined;
            resourceInputs["wagPort"] = state ? state.wagPort : undefined;
        } else {
            const args = argsOrState as WagprofileArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dhcpIpAddr"] = args ? args.dhcpIpAddr : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pingInterval"] = args ? args.pingInterval : undefined;
            resourceInputs["pingNumber"] = args ? args.pingNumber : undefined;
            resourceInputs["returnPacketTimeout"] = args ? args.returnPacketTimeout : undefined;
            resourceInputs["tunnelType"] = args ? args.tunnelType : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["wagIp"] = args ? args.wagIp : undefined;
            resourceInputs["wagPort"] = args ? args.wagPort : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Wagprofile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Wagprofile resources.
 */
export interface WagprofileState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * IP address of the monitoring DHCP request packet sent through the tunnel.
     */
    dhcpIpAddr?: pulumi.Input<string>;
    /**
     * Tunnel profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
     */
    pingInterval?: pulumi.Input<number>;
    /**
     * Number of the tunnel monitoring echo packets (1 - 65535, default = 5).
     */
    pingNumber?: pulumi.Input<number>;
    /**
     * Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
     */
    returnPacketTimeout?: pulumi.Input<number>;
    /**
     * Tunnel type. Valid values: `l2tpv3`, `gre`.
     */
    tunnelType?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * IP Address of the wireless access gateway.
     */
    wagIp?: pulumi.Input<string>;
    /**
     * UDP port of the wireless access gateway.
     */
    wagPort?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Wagprofile resource.
 */
export interface WagprofileArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * IP address of the monitoring DHCP request packet sent through the tunnel.
     */
    dhcpIpAddr?: pulumi.Input<string>;
    /**
     * Tunnel profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
     */
    pingInterval?: pulumi.Input<number>;
    /**
     * Number of the tunnel monitoring echo packets (1 - 65535, default = 5).
     */
    pingNumber?: pulumi.Input<number>;
    /**
     * Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
     */
    returnPacketTimeout?: pulumi.Input<number>;
    /**
     * Tunnel type. Valid values: `l2tpv3`, `gre`.
     */
    tunnelType?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * IP Address of the wireless access gateway.
     */
    wagIp?: pulumi.Input<string>;
    /**
     * UDP port of the wireless access gateway.
     */
    wagPort?: pulumi.Input<number>;
}
