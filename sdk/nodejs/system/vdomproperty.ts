// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure VDOM property.
 *
 * ## Import
 *
 * System VdomProperty can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/vdomproperty:Vdomproperty labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/vdomproperty:Vdomproperty labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Vdomproperty extends pulumi.CustomResource {
    /**
     * Get an existing Vdomproperty resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VdompropertyState, opts?: pulumi.CustomResourceOptions): Vdomproperty {
        return new Vdomproperty(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/vdomproperty:Vdomproperty';

    /**
     * Returns true if the given object is an instance of Vdomproperty.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vdomproperty {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vdomproperty.__pulumiType;
    }

    /**
     * Maximum guaranteed number of firewall custom services.
     */
    public readonly customService!: pulumi.Output<string>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of dial-up tunnels.
     */
    public readonly dialupTunnel!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
     */
    public readonly firewallAddress!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of firewall address groups (IPv4, IPv6).
     */
    public readonly firewallAddrgrp!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
     */
    public readonly firewallPolicy!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of VPN IPsec phase 1 tunnels.
     */
    public readonly ipsecPhase1!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
     */
    public readonly ipsecPhase1Interface!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of VPN IPsec phase 2 tunnels.
     */
    public readonly ipsecPhase2!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
     */
    public readonly ipsecPhase2Interface!: pulumi.Output<string>;
    /**
     * Log disk quota in MB (range depends on how much disk space is available).
     */
    public readonly logDiskQuota!: pulumi.Output<string>;
    /**
     * VDOM name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of firewall one-time schedules.
     */
    public readonly onetimeSchedule!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of concurrent proxy users.
     */
    public readonly proxy!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of firewall recurring schedules.
     */
    public readonly recurringSchedule!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of firewall service groups.
     */
    public readonly serviceGroup!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of sessions.
     */
    public readonly session!: pulumi.Output<string>;
    /**
     * Permanent SNMP Index of the virtual domain (0 - 4294967295).
     */
    public readonly snmpIndex!: pulumi.Output<number>;
    /**
     * Maximum guaranteed number of SSL-VPNs.
     */
    public readonly sslvpn!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of local users.
     */
    public readonly user!: pulumi.Output<string>;
    /**
     * Maximum guaranteed number of user groups.
     */
    public readonly userGroup!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Vdomproperty resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VdompropertyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VdompropertyArgs | VdompropertyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VdompropertyState | undefined;
            resourceInputs["customService"] = state ? state.customService : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dialupTunnel"] = state ? state.dialupTunnel : undefined;
            resourceInputs["firewallAddress"] = state ? state.firewallAddress : undefined;
            resourceInputs["firewallAddrgrp"] = state ? state.firewallAddrgrp : undefined;
            resourceInputs["firewallPolicy"] = state ? state.firewallPolicy : undefined;
            resourceInputs["ipsecPhase1"] = state ? state.ipsecPhase1 : undefined;
            resourceInputs["ipsecPhase1Interface"] = state ? state.ipsecPhase1Interface : undefined;
            resourceInputs["ipsecPhase2"] = state ? state.ipsecPhase2 : undefined;
            resourceInputs["ipsecPhase2Interface"] = state ? state.ipsecPhase2Interface : undefined;
            resourceInputs["logDiskQuota"] = state ? state.logDiskQuota : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["onetimeSchedule"] = state ? state.onetimeSchedule : undefined;
            resourceInputs["proxy"] = state ? state.proxy : undefined;
            resourceInputs["recurringSchedule"] = state ? state.recurringSchedule : undefined;
            resourceInputs["serviceGroup"] = state ? state.serviceGroup : undefined;
            resourceInputs["session"] = state ? state.session : undefined;
            resourceInputs["snmpIndex"] = state ? state.snmpIndex : undefined;
            resourceInputs["sslvpn"] = state ? state.sslvpn : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
            resourceInputs["userGroup"] = state ? state.userGroup : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as VdompropertyArgs | undefined;
            resourceInputs["customService"] = args ? args.customService : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dialupTunnel"] = args ? args.dialupTunnel : undefined;
            resourceInputs["firewallAddress"] = args ? args.firewallAddress : undefined;
            resourceInputs["firewallAddrgrp"] = args ? args.firewallAddrgrp : undefined;
            resourceInputs["firewallPolicy"] = args ? args.firewallPolicy : undefined;
            resourceInputs["ipsecPhase1"] = args ? args.ipsecPhase1 : undefined;
            resourceInputs["ipsecPhase1Interface"] = args ? args.ipsecPhase1Interface : undefined;
            resourceInputs["ipsecPhase2"] = args ? args.ipsecPhase2 : undefined;
            resourceInputs["ipsecPhase2Interface"] = args ? args.ipsecPhase2Interface : undefined;
            resourceInputs["logDiskQuota"] = args ? args.logDiskQuota : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["onetimeSchedule"] = args ? args.onetimeSchedule : undefined;
            resourceInputs["proxy"] = args ? args.proxy : undefined;
            resourceInputs["recurringSchedule"] = args ? args.recurringSchedule : undefined;
            resourceInputs["serviceGroup"] = args ? args.serviceGroup : undefined;
            resourceInputs["session"] = args ? args.session : undefined;
            resourceInputs["snmpIndex"] = args ? args.snmpIndex : undefined;
            resourceInputs["sslvpn"] = args ? args.sslvpn : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["userGroup"] = args ? args.userGroup : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vdomproperty.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vdomproperty resources.
 */
export interface VdompropertyState {
    /**
     * Maximum guaranteed number of firewall custom services.
     */
    customService?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of dial-up tunnels.
     */
    dialupTunnel?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
     */
    firewallAddress?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of firewall address groups (IPv4, IPv6).
     */
    firewallAddrgrp?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
     */
    firewallPolicy?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of VPN IPsec phase 1 tunnels.
     */
    ipsecPhase1?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
     */
    ipsecPhase1Interface?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of VPN IPsec phase 2 tunnels.
     */
    ipsecPhase2?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
     */
    ipsecPhase2Interface?: pulumi.Input<string>;
    /**
     * Log disk quota in MB (range depends on how much disk space is available).
     */
    logDiskQuota?: pulumi.Input<string>;
    /**
     * VDOM name.
     */
    name?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of firewall one-time schedules.
     */
    onetimeSchedule?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of concurrent proxy users.
     */
    proxy?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of firewall recurring schedules.
     */
    recurringSchedule?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of firewall service groups.
     */
    serviceGroup?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of sessions.
     */
    session?: pulumi.Input<string>;
    /**
     * Permanent SNMP Index of the virtual domain (0 - 4294967295).
     */
    snmpIndex?: pulumi.Input<number>;
    /**
     * Maximum guaranteed number of SSL-VPNs.
     */
    sslvpn?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of local users.
     */
    user?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of user groups.
     */
    userGroup?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Vdomproperty resource.
 */
export interface VdompropertyArgs {
    /**
     * Maximum guaranteed number of firewall custom services.
     */
    customService?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of dial-up tunnels.
     */
    dialupTunnel?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
     */
    firewallAddress?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of firewall address groups (IPv4, IPv6).
     */
    firewallAddrgrp?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
     */
    firewallPolicy?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of VPN IPsec phase 1 tunnels.
     */
    ipsecPhase1?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
     */
    ipsecPhase1Interface?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of VPN IPsec phase 2 tunnels.
     */
    ipsecPhase2?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
     */
    ipsecPhase2Interface?: pulumi.Input<string>;
    /**
     * Log disk quota in MB (range depends on how much disk space is available).
     */
    logDiskQuota?: pulumi.Input<string>;
    /**
     * VDOM name.
     */
    name?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of firewall one-time schedules.
     */
    onetimeSchedule?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of concurrent proxy users.
     */
    proxy?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of firewall recurring schedules.
     */
    recurringSchedule?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of firewall service groups.
     */
    serviceGroup?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of sessions.
     */
    session?: pulumi.Input<string>;
    /**
     * Permanent SNMP Index of the virtual domain (0 - 4294967295).
     */
    snmpIndex?: pulumi.Input<number>;
    /**
     * Maximum guaranteed number of SSL-VPNs.
     */
    sslvpn?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of local users.
     */
    user?: pulumi.Input<string>;
    /**
     * Maximum guaranteed number of user groups.
     */
    userGroup?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
