// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure VXLAN devices.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Vxlan("trname", {
 *     dstport: 4789,
 *     "interface": "port3",
 *     ipVersion: "ipv4-unicast",
 *     remoteIps: [{
 *         ip: "1.1.1.1",
 *     }],
 *     vni: 3,
 * });
 * ```
 *
 * ## Import
 *
 * System Vxlan can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/vxlan:Vxlan labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/vxlan:Vxlan labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Vxlan extends pulumi.CustomResource {
    /**
     * Get an existing Vxlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VxlanState, opts?: pulumi.CustomResourceOptions): Vxlan {
        return new Vxlan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/vxlan:Vxlan';

    /**
     * Returns true if the given object is an instance of Vxlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vxlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vxlan.__pulumiType;
    }

    /**
     * VXLAN destination port (1 - 65535, default = 4789).
     */
    public readonly dstport!: pulumi.Output<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * EVPN instance.
     */
    public readonly evpnId!: pulumi.Output<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Outgoing interface for VXLAN encapsulated traffic.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast. Valid values: `ipv4-unicast`, `ipv6-unicast`, `ipv4-multicast`, `ipv6-multicast`.
     */
    public readonly ipVersion!: pulumi.Output<string>;
    /**
     * Enable/disable VXLAN MAC learning from traffic. Valid values: `enable`, `disable`.
     */
    public readonly learnFromTraffic!: pulumi.Output<string>;
    /**
     * VXLAN multicast TTL (1-255, default = 0).
     */
    public readonly multicastTtl!: pulumi.Output<number>;
    /**
     * VXLAN device or interface name. Must be a unique interface name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp6` block is documented below.
     */
    public readonly remoteIp6s!: pulumi.Output<outputs.system.VxlanRemoteIp6[] | undefined>;
    /**
     * IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp` block is documented below.
     */
    public readonly remoteIps!: pulumi.Output<outputs.system.VxlanRemoteIp[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * VXLAN network ID.
     */
    public readonly vni!: pulumi.Output<number>;

    /**
     * Create a Vxlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VxlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VxlanArgs | VxlanState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VxlanState | undefined;
            resourceInputs["dstport"] = state ? state.dstport : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["evpnId"] = state ? state.evpnId : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["ipVersion"] = state ? state.ipVersion : undefined;
            resourceInputs["learnFromTraffic"] = state ? state.learnFromTraffic : undefined;
            resourceInputs["multicastTtl"] = state ? state.multicastTtl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["remoteIp6s"] = state ? state.remoteIp6s : undefined;
            resourceInputs["remoteIps"] = state ? state.remoteIps : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vni"] = state ? state.vni : undefined;
        } else {
            const args = argsOrState as VxlanArgs | undefined;
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            if ((!args || args.ipVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipVersion'");
            }
            if ((!args || args.vni === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vni'");
            }
            resourceInputs["dstport"] = args ? args.dstport : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["evpnId"] = args ? args.evpnId : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["ipVersion"] = args ? args.ipVersion : undefined;
            resourceInputs["learnFromTraffic"] = args ? args.learnFromTraffic : undefined;
            resourceInputs["multicastTtl"] = args ? args.multicastTtl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["remoteIp6s"] = args ? args.remoteIp6s : undefined;
            resourceInputs["remoteIps"] = args ? args.remoteIps : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vni"] = args ? args.vni : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vxlan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vxlan resources.
 */
export interface VxlanState {
    /**
     * VXLAN destination port (1 - 65535, default = 4789).
     */
    dstport?: pulumi.Input<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * EVPN instance.
     */
    evpnId?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Outgoing interface for VXLAN encapsulated traffic.
     */
    interface?: pulumi.Input<string>;
    /**
     * IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast. Valid values: `ipv4-unicast`, `ipv6-unicast`, `ipv4-multicast`, `ipv6-multicast`.
     */
    ipVersion?: pulumi.Input<string>;
    /**
     * Enable/disable VXLAN MAC learning from traffic. Valid values: `enable`, `disable`.
     */
    learnFromTraffic?: pulumi.Input<string>;
    /**
     * VXLAN multicast TTL (1-255, default = 0).
     */
    multicastTtl?: pulumi.Input<number>;
    /**
     * VXLAN device or interface name. Must be a unique interface name.
     */
    name?: pulumi.Input<string>;
    /**
     * IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp6` block is documented below.
     */
    remoteIp6s?: pulumi.Input<pulumi.Input<inputs.system.VxlanRemoteIp6>[]>;
    /**
     * IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp` block is documented below.
     */
    remoteIps?: pulumi.Input<pulumi.Input<inputs.system.VxlanRemoteIp>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * VXLAN network ID.
     */
    vni?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Vxlan resource.
 */
export interface VxlanArgs {
    /**
     * VXLAN destination port (1 - 65535, default = 4789).
     */
    dstport?: pulumi.Input<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * EVPN instance.
     */
    evpnId?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Outgoing interface for VXLAN encapsulated traffic.
     */
    interface: pulumi.Input<string>;
    /**
     * IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast. Valid values: `ipv4-unicast`, `ipv6-unicast`, `ipv4-multicast`, `ipv6-multicast`.
     */
    ipVersion: pulumi.Input<string>;
    /**
     * Enable/disable VXLAN MAC learning from traffic. Valid values: `enable`, `disable`.
     */
    learnFromTraffic?: pulumi.Input<string>;
    /**
     * VXLAN multicast TTL (1-255, default = 0).
     */
    multicastTtl?: pulumi.Input<number>;
    /**
     * VXLAN device or interface name. Must be a unique interface name.
     */
    name?: pulumi.Input<string>;
    /**
     * IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp6` block is documented below.
     */
    remoteIp6s?: pulumi.Input<pulumi.Input<inputs.system.VxlanRemoteIp6>[]>;
    /**
     * IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp` block is documented below.
     */
    remoteIps?: pulumi.Input<pulumi.Input<inputs.system.VxlanRemoteIp>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * VXLAN network ID.
     */
    vni: pulumi.Input<number>;
}
