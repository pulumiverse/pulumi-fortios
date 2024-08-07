// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure resource limits.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Resourcelimits("trname", {
 *     customService: 0,
 *     dialupTunnel: 0,
 *     firewallAddress: 41024,
 *     firewallAddrgrp: 10692,
 *     firewallPolicy: 41024,
 *     ipsecPhase1: 2000,
 *     ipsecPhase1Interface: 0,
 *     ipsecPhase2: 2000,
 *     ipsecPhase2Interface: 0,
 *     logDiskQuota: 30235,
 *     onetimeSchedule: 0,
 *     proxy: 64000,
 *     recurringSchedule: 0,
 *     serviceGroup: 0,
 *     session: 0,
 *     sslvpn: 0,
 *     user: 0,
 *     userGroup: 0,
 * });
 * ```
 *
 * ## Import
 *
 * System ResourceLimits can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/resourcelimits:Resourcelimits labelname SystemResourceLimits
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/resourcelimits:Resourcelimits labelname SystemResourceLimits
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Resourcelimits extends pulumi.CustomResource {
    /**
     * Get an existing Resourcelimits resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourcelimitsState, opts?: pulumi.CustomResourceOptions): Resourcelimits {
        return new Resourcelimits(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/resourcelimits:Resourcelimits';

    /**
     * Returns true if the given object is an instance of Resourcelimits.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Resourcelimits {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Resourcelimits.__pulumiType;
    }

    /**
     * Maximum number of firewall custom services.
     */
    public readonly customService!: pulumi.Output<number>;
    /**
     * Maximum number of dial-up tunnels.
     */
    public readonly dialupTunnel!: pulumi.Output<number>;
    /**
     * Maximum number of firewall addresses (IPv4, IPv6, multicast).
     */
    public readonly firewallAddress!: pulumi.Output<number>;
    /**
     * Maximum number of firewall address groups (IPv4, IPv6).
     */
    public readonly firewallAddrgrp!: pulumi.Output<number>;
    /**
     * Maximum number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).
     */
    public readonly firewallPolicy!: pulumi.Output<number>;
    /**
     * Maximum number of VPN IPsec phase1 tunnels.
     */
    public readonly ipsecPhase1!: pulumi.Output<number>;
    /**
     * Maximum number of VPN IPsec phase1 interface tunnels.
     */
    public readonly ipsecPhase1Interface!: pulumi.Output<number>;
    /**
     * Maximum number of VPN IPsec phase2 tunnels.
     */
    public readonly ipsecPhase2!: pulumi.Output<number>;
    /**
     * Maximum number of VPN IPsec phase2 interface tunnels.
     */
    public readonly ipsecPhase2Interface!: pulumi.Output<number>;
    /**
     * Log disk quota in megabytes (MB).
     */
    public readonly logDiskQuota!: pulumi.Output<number>;
    /**
     * Maximum number of firewall one-time schedules.
     */
    public readonly onetimeSchedule!: pulumi.Output<number>;
    /**
     * Maximum number of concurrent proxy users.
     */
    public readonly proxy!: pulumi.Output<number>;
    /**
     * Maximum number of firewall recurring schedules.
     */
    public readonly recurringSchedule!: pulumi.Output<number>;
    /**
     * Maximum number of firewall service groups.
     */
    public readonly serviceGroup!: pulumi.Output<number>;
    /**
     * Maximum number of sessions.
     */
    public readonly session!: pulumi.Output<number>;
    /**
     * Maximum number of SSL-VPN.
     */
    public readonly sslvpn!: pulumi.Output<number>;
    /**
     * Maximum number of local users.
     */
    public readonly user!: pulumi.Output<number>;
    /**
     * Maximum number of user groups.
     */
    public readonly userGroup!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Resourcelimits resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResourcelimitsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourcelimitsArgs | ResourcelimitsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourcelimitsState | undefined;
            resourceInputs["customService"] = state ? state.customService : undefined;
            resourceInputs["dialupTunnel"] = state ? state.dialupTunnel : undefined;
            resourceInputs["firewallAddress"] = state ? state.firewallAddress : undefined;
            resourceInputs["firewallAddrgrp"] = state ? state.firewallAddrgrp : undefined;
            resourceInputs["firewallPolicy"] = state ? state.firewallPolicy : undefined;
            resourceInputs["ipsecPhase1"] = state ? state.ipsecPhase1 : undefined;
            resourceInputs["ipsecPhase1Interface"] = state ? state.ipsecPhase1Interface : undefined;
            resourceInputs["ipsecPhase2"] = state ? state.ipsecPhase2 : undefined;
            resourceInputs["ipsecPhase2Interface"] = state ? state.ipsecPhase2Interface : undefined;
            resourceInputs["logDiskQuota"] = state ? state.logDiskQuota : undefined;
            resourceInputs["onetimeSchedule"] = state ? state.onetimeSchedule : undefined;
            resourceInputs["proxy"] = state ? state.proxy : undefined;
            resourceInputs["recurringSchedule"] = state ? state.recurringSchedule : undefined;
            resourceInputs["serviceGroup"] = state ? state.serviceGroup : undefined;
            resourceInputs["session"] = state ? state.session : undefined;
            resourceInputs["sslvpn"] = state ? state.sslvpn : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
            resourceInputs["userGroup"] = state ? state.userGroup : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ResourcelimitsArgs | undefined;
            resourceInputs["customService"] = args ? args.customService : undefined;
            resourceInputs["dialupTunnel"] = args ? args.dialupTunnel : undefined;
            resourceInputs["firewallAddress"] = args ? args.firewallAddress : undefined;
            resourceInputs["firewallAddrgrp"] = args ? args.firewallAddrgrp : undefined;
            resourceInputs["firewallPolicy"] = args ? args.firewallPolicy : undefined;
            resourceInputs["ipsecPhase1"] = args ? args.ipsecPhase1 : undefined;
            resourceInputs["ipsecPhase1Interface"] = args ? args.ipsecPhase1Interface : undefined;
            resourceInputs["ipsecPhase2"] = args ? args.ipsecPhase2 : undefined;
            resourceInputs["ipsecPhase2Interface"] = args ? args.ipsecPhase2Interface : undefined;
            resourceInputs["logDiskQuota"] = args ? args.logDiskQuota : undefined;
            resourceInputs["onetimeSchedule"] = args ? args.onetimeSchedule : undefined;
            resourceInputs["proxy"] = args ? args.proxy : undefined;
            resourceInputs["recurringSchedule"] = args ? args.recurringSchedule : undefined;
            resourceInputs["serviceGroup"] = args ? args.serviceGroup : undefined;
            resourceInputs["session"] = args ? args.session : undefined;
            resourceInputs["sslvpn"] = args ? args.sslvpn : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["userGroup"] = args ? args.userGroup : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Resourcelimits.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Resourcelimits resources.
 */
export interface ResourcelimitsState {
    /**
     * Maximum number of firewall custom services.
     */
    customService?: pulumi.Input<number>;
    /**
     * Maximum number of dial-up tunnels.
     */
    dialupTunnel?: pulumi.Input<number>;
    /**
     * Maximum number of firewall addresses (IPv4, IPv6, multicast).
     */
    firewallAddress?: pulumi.Input<number>;
    /**
     * Maximum number of firewall address groups (IPv4, IPv6).
     */
    firewallAddrgrp?: pulumi.Input<number>;
    /**
     * Maximum number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).
     */
    firewallPolicy?: pulumi.Input<number>;
    /**
     * Maximum number of VPN IPsec phase1 tunnels.
     */
    ipsecPhase1?: pulumi.Input<number>;
    /**
     * Maximum number of VPN IPsec phase1 interface tunnels.
     */
    ipsecPhase1Interface?: pulumi.Input<number>;
    /**
     * Maximum number of VPN IPsec phase2 tunnels.
     */
    ipsecPhase2?: pulumi.Input<number>;
    /**
     * Maximum number of VPN IPsec phase2 interface tunnels.
     */
    ipsecPhase2Interface?: pulumi.Input<number>;
    /**
     * Log disk quota in megabytes (MB).
     */
    logDiskQuota?: pulumi.Input<number>;
    /**
     * Maximum number of firewall one-time schedules.
     */
    onetimeSchedule?: pulumi.Input<number>;
    /**
     * Maximum number of concurrent proxy users.
     */
    proxy?: pulumi.Input<number>;
    /**
     * Maximum number of firewall recurring schedules.
     */
    recurringSchedule?: pulumi.Input<number>;
    /**
     * Maximum number of firewall service groups.
     */
    serviceGroup?: pulumi.Input<number>;
    /**
     * Maximum number of sessions.
     */
    session?: pulumi.Input<number>;
    /**
     * Maximum number of SSL-VPN.
     */
    sslvpn?: pulumi.Input<number>;
    /**
     * Maximum number of local users.
     */
    user?: pulumi.Input<number>;
    /**
     * Maximum number of user groups.
     */
    userGroup?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Resourcelimits resource.
 */
export interface ResourcelimitsArgs {
    /**
     * Maximum number of firewall custom services.
     */
    customService?: pulumi.Input<number>;
    /**
     * Maximum number of dial-up tunnels.
     */
    dialupTunnel?: pulumi.Input<number>;
    /**
     * Maximum number of firewall addresses (IPv4, IPv6, multicast).
     */
    firewallAddress?: pulumi.Input<number>;
    /**
     * Maximum number of firewall address groups (IPv4, IPv6).
     */
    firewallAddrgrp?: pulumi.Input<number>;
    /**
     * Maximum number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).
     */
    firewallPolicy?: pulumi.Input<number>;
    /**
     * Maximum number of VPN IPsec phase1 tunnels.
     */
    ipsecPhase1?: pulumi.Input<number>;
    /**
     * Maximum number of VPN IPsec phase1 interface tunnels.
     */
    ipsecPhase1Interface?: pulumi.Input<number>;
    /**
     * Maximum number of VPN IPsec phase2 tunnels.
     */
    ipsecPhase2?: pulumi.Input<number>;
    /**
     * Maximum number of VPN IPsec phase2 interface tunnels.
     */
    ipsecPhase2Interface?: pulumi.Input<number>;
    /**
     * Log disk quota in megabytes (MB).
     */
    logDiskQuota?: pulumi.Input<number>;
    /**
     * Maximum number of firewall one-time schedules.
     */
    onetimeSchedule?: pulumi.Input<number>;
    /**
     * Maximum number of concurrent proxy users.
     */
    proxy?: pulumi.Input<number>;
    /**
     * Maximum number of firewall recurring schedules.
     */
    recurringSchedule?: pulumi.Input<number>;
    /**
     * Maximum number of firewall service groups.
     */
    serviceGroup?: pulumi.Input<number>;
    /**
     * Maximum number of sessions.
     */
    session?: pulumi.Input<number>;
    /**
     * Maximum number of SSL-VPN.
     */
    sslvpn?: pulumi.Input<number>;
    /**
     * Maximum number of local users.
     */
    user?: pulumi.Input<number>;
    /**
     * Maximum number of user groups.
     */
    userGroup?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
