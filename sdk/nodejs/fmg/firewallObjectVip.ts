// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource supports Create/Read/Update/Delete firewall object virtual ip for FortiManager.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.fmg.FirewallObjectVip("test1", {
 *     arpReply: "enable",
 *     comment: "test obj vip",
 *     configDefault: "enable",
 *     extIntf: "any",
 *     extIp: "2.2.2.2",
 *     mappedIp: "1.1.1.1",
 *     type: "static-nat",
 * });
 * const test2 = new fortios.fmg.FirewallObjectVip("test2", {
 *     arpReply: "disable",
 *     comment: "test obj vip",
 *     configDefault: "enable",
 *     extIp: "2.2.2.2-2.2.2.100",
 *     mappedAddr: "update.microsoft.com",
 *     type: "fqdn",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class FirewallObjectVip extends pulumi.CustomResource {
    /**
     * Get an existing FirewallObjectVip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallObjectVipState, opts?: pulumi.CustomResourceOptions): FirewallObjectVip {
        return new FirewallObjectVip(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:fmg/firewallObjectVip:FirewallObjectVip';

    /**
     * Returns true if the given object is an instance of FirewallObjectVip.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallObjectVip {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallObjectVip.__pulumiType;
    }

    /**
     * ADOM name. default is 'root'.
     */
    public readonly adom!: pulumi.Output<string | undefined>;
    /**
     * Enable to respond to ARP requests for this virtual IP address. Enabled by default.
     */
    public readonly arpReply!: pulumi.Output<string | undefined>;
    /**
     * Comments.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Enable/Disable config default value. enabled by default.
     */
    public readonly configDefault!: pulumi.Output<string | undefined>;
    /**
     * Interface connected to the source network that receives the packets that will be forwarded to the destination network.
     */
    public readonly extIntf!: pulumi.Output<string | undefined>;
    /**
     * IP address or address range on the external interface that you want to map to an address or address range on the destination network.
     */
    public readonly extIp!: pulumi.Output<string | undefined>;
    /**
     * Mapped FQDN address name.
     */
    public readonly mappedAddr!: pulumi.Output<string | undefined>;
    /**
     * Mapped Ip address.
     */
    public readonly mappedIp!: pulumi.Output<string | undefined>;
    /**
     * Virtual IP name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallObjectVip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallObjectVipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallObjectVipArgs | FirewallObjectVipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallObjectVipState | undefined;
            resourceInputs["adom"] = state ? state.adom : undefined;
            resourceInputs["arpReply"] = state ? state.arpReply : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["configDefault"] = state ? state.configDefault : undefined;
            resourceInputs["extIntf"] = state ? state.extIntf : undefined;
            resourceInputs["extIp"] = state ? state.extIp : undefined;
            resourceInputs["mappedAddr"] = state ? state.mappedAddr : undefined;
            resourceInputs["mappedIp"] = state ? state.mappedIp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as FirewallObjectVipArgs | undefined;
            resourceInputs["adom"] = args ? args.adom : undefined;
            resourceInputs["arpReply"] = args ? args.arpReply : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["configDefault"] = args ? args.configDefault : undefined;
            resourceInputs["extIntf"] = args ? args.extIntf : undefined;
            resourceInputs["extIp"] = args ? args.extIp : undefined;
            resourceInputs["mappedAddr"] = args ? args.mappedAddr : undefined;
            resourceInputs["mappedIp"] = args ? args.mappedIp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallObjectVip.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallObjectVip resources.
 */
export interface FirewallObjectVipState {
    /**
     * ADOM name. default is 'root'.
     */
    adom?: pulumi.Input<string>;
    /**
     * Enable to respond to ARP requests for this virtual IP address. Enabled by default.
     */
    arpReply?: pulumi.Input<string>;
    /**
     * Comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable/Disable config default value. enabled by default.
     */
    configDefault?: pulumi.Input<string>;
    /**
     * Interface connected to the source network that receives the packets that will be forwarded to the destination network.
     */
    extIntf?: pulumi.Input<string>;
    /**
     * IP address or address range on the external interface that you want to map to an address or address range on the destination network.
     */
    extIp?: pulumi.Input<string>;
    /**
     * Mapped FQDN address name.
     */
    mappedAddr?: pulumi.Input<string>;
    /**
     * Mapped Ip address.
     */
    mappedIp?: pulumi.Input<string>;
    /**
     * Virtual IP name.
     */
    name?: pulumi.Input<string>;
    /**
     * Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallObjectVip resource.
 */
export interface FirewallObjectVipArgs {
    /**
     * ADOM name. default is 'root'.
     */
    adom?: pulumi.Input<string>;
    /**
     * Enable to respond to ARP requests for this virtual IP address. Enabled by default.
     */
    arpReply?: pulumi.Input<string>;
    /**
     * Comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable/Disable config default value. enabled by default.
     */
    configDefault?: pulumi.Input<string>;
    /**
     * Interface connected to the source network that receives the packets that will be forwarded to the destination network.
     */
    extIntf?: pulumi.Input<string>;
    /**
     * IP address or address range on the external interface that you want to map to an address or address range on the destination network.
     */
    extIp?: pulumi.Input<string>;
    /**
     * Mapped FQDN address name.
     */
    mappedAddr?: pulumi.Input<string>;
    /**
     * Mapped Ip address.
     */
    mappedIp?: pulumi.Input<string>;
    /**
     * Virtual IP name.
     */
    name?: pulumi.Input<string>;
    /**
     * Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
     */
    type?: pulumi.Input<string>;
}
