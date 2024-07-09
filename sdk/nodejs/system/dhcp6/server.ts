// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Configure DHCPv6 servers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.dhcp6.Server("trname", {
 *     fosid: 1,
 *     "interface": "port3",
 *     leaseTime: 604800,
 *     rapidCommit: "disable",
 *     status: "enable",
 *     subnet: "2001:db8:1234:113::/64",
 * });
 * ```
 *
 * ## Import
 *
 * SystemDhcp6 Server can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/dhcp6/server:Server labelname {{fosid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/dhcp6/server:Server labelname {{fosid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Server extends pulumi.CustomResource {
    /**
     * Get an existing Server resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerState, opts?: pulumi.CustomResourceOptions): Server {
        return new Server(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/dhcp6/server:Server';

    /**
     * Returns true if the given object is an instance of Server.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Server {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Server.__pulumiType;
    }

    /**
     * IAID of obtained delegated-prefix from the upstream interface.
     */
    public readonly delegatedPrefixIaid!: pulumi.Output<number>;
    /**
     * DNS search list options. Valid values: `delegated`, `specify`.
     */
    public readonly dnsSearchList!: pulumi.Output<string>;
    /**
     * DNS server 1.
     */
    public readonly dnsServer1!: pulumi.Output<string>;
    /**
     * DNS server 2.
     */
    public readonly dnsServer2!: pulumi.Output<string>;
    /**
     * DNS server 3.
     */
    public readonly dnsServer3!: pulumi.Output<string>;
    /**
     * DNS server 4.
     */
    public readonly dnsServer4!: pulumi.Output<string>;
    /**
     * Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
     */
    public readonly dnsService!: pulumi.Output<string>;
    /**
     * Domain name suffix for the IP addresses that the DHCP server assigns to clients.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * DHCP server can assign IP configurations to clients connected to this interface.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Method used to assign client IP. Valid values: `range`, `delegated`.
     */
    public readonly ipMode!: pulumi.Output<string>;
    /**
     * DHCP IP range configuration. The structure of `ipRange` block is documented below.
     */
    public readonly ipRanges!: pulumi.Output<outputs.system.dhcp6.ServerIpRange[] | undefined>;
    /**
     * Lease time in seconds, 0 means unlimited.
     */
    public readonly leaseTime!: pulumi.Output<number>;
    /**
     * Option 1.
     */
    public readonly option1!: pulumi.Output<string>;
    /**
     * Option 2.
     */
    public readonly option2!: pulumi.Output<string>;
    /**
     * Option 3.
     */
    public readonly option3!: pulumi.Output<string>;
    /**
     * Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
     */
    public readonly prefixMode!: pulumi.Output<string>;
    /**
     * DHCP prefix configuration. The structure of `prefixRange` block is documented below.
     */
    public readonly prefixRanges!: pulumi.Output<outputs.system.dhcp6.ServerPrefixRange[] | undefined>;
    /**
     * Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
     */
    public readonly rapidCommit!: pulumi.Output<string>;
    /**
     * Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Subnet or subnet-id if the IP mode is delegated.
     */
    public readonly subnet!: pulumi.Output<string>;
    /**
     * Interface name from where delegated information is provided.
     */
    public readonly upstreamInterface!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Server resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerArgs | ServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerState | undefined;
            resourceInputs["delegatedPrefixIaid"] = state ? state.delegatedPrefixIaid : undefined;
            resourceInputs["dnsSearchList"] = state ? state.dnsSearchList : undefined;
            resourceInputs["dnsServer1"] = state ? state.dnsServer1 : undefined;
            resourceInputs["dnsServer2"] = state ? state.dnsServer2 : undefined;
            resourceInputs["dnsServer3"] = state ? state.dnsServer3 : undefined;
            resourceInputs["dnsServer4"] = state ? state.dnsServer4 : undefined;
            resourceInputs["dnsService"] = state ? state.dnsService : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["ipMode"] = state ? state.ipMode : undefined;
            resourceInputs["ipRanges"] = state ? state.ipRanges : undefined;
            resourceInputs["leaseTime"] = state ? state.leaseTime : undefined;
            resourceInputs["option1"] = state ? state.option1 : undefined;
            resourceInputs["option2"] = state ? state.option2 : undefined;
            resourceInputs["option3"] = state ? state.option3 : undefined;
            resourceInputs["prefixMode"] = state ? state.prefixMode : undefined;
            resourceInputs["prefixRanges"] = state ? state.prefixRanges : undefined;
            resourceInputs["rapidCommit"] = state ? state.rapidCommit : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnet"] = state ? state.subnet : undefined;
            resourceInputs["upstreamInterface"] = state ? state.upstreamInterface : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ServerArgs | undefined;
            if ((!args || args.fosid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fosid'");
            }
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            if ((!args || args.subnet === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnet'");
            }
            resourceInputs["delegatedPrefixIaid"] = args ? args.delegatedPrefixIaid : undefined;
            resourceInputs["dnsSearchList"] = args ? args.dnsSearchList : undefined;
            resourceInputs["dnsServer1"] = args ? args.dnsServer1 : undefined;
            resourceInputs["dnsServer2"] = args ? args.dnsServer2 : undefined;
            resourceInputs["dnsServer3"] = args ? args.dnsServer3 : undefined;
            resourceInputs["dnsServer4"] = args ? args.dnsServer4 : undefined;
            resourceInputs["dnsService"] = args ? args.dnsService : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["ipMode"] = args ? args.ipMode : undefined;
            resourceInputs["ipRanges"] = args ? args.ipRanges : undefined;
            resourceInputs["leaseTime"] = args ? args.leaseTime : undefined;
            resourceInputs["option1"] = args ? args.option1 : undefined;
            resourceInputs["option2"] = args ? args.option2 : undefined;
            resourceInputs["option3"] = args ? args.option3 : undefined;
            resourceInputs["prefixMode"] = args ? args.prefixMode : undefined;
            resourceInputs["prefixRanges"] = args ? args.prefixRanges : undefined;
            resourceInputs["rapidCommit"] = args ? args.rapidCommit : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["subnet"] = args ? args.subnet : undefined;
            resourceInputs["upstreamInterface"] = args ? args.upstreamInterface : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Server.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Server resources.
 */
export interface ServerState {
    /**
     * IAID of obtained delegated-prefix from the upstream interface.
     */
    delegatedPrefixIaid?: pulumi.Input<number>;
    /**
     * DNS search list options. Valid values: `delegated`, `specify`.
     */
    dnsSearchList?: pulumi.Input<string>;
    /**
     * DNS server 1.
     */
    dnsServer1?: pulumi.Input<string>;
    /**
     * DNS server 2.
     */
    dnsServer2?: pulumi.Input<string>;
    /**
     * DNS server 3.
     */
    dnsServer3?: pulumi.Input<string>;
    /**
     * DNS server 4.
     */
    dnsServer4?: pulumi.Input<string>;
    /**
     * Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
     */
    dnsService?: pulumi.Input<string>;
    /**
     * Domain name suffix for the IP addresses that the DHCP server assigns to clients.
     */
    domain?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * DHCP server can assign IP configurations to clients connected to this interface.
     */
    interface?: pulumi.Input<string>;
    /**
     * Method used to assign client IP. Valid values: `range`, `delegated`.
     */
    ipMode?: pulumi.Input<string>;
    /**
     * DHCP IP range configuration. The structure of `ipRange` block is documented below.
     */
    ipRanges?: pulumi.Input<pulumi.Input<inputs.system.dhcp6.ServerIpRange>[]>;
    /**
     * Lease time in seconds, 0 means unlimited.
     */
    leaseTime?: pulumi.Input<number>;
    /**
     * Option 1.
     */
    option1?: pulumi.Input<string>;
    /**
     * Option 2.
     */
    option2?: pulumi.Input<string>;
    /**
     * Option 3.
     */
    option3?: pulumi.Input<string>;
    /**
     * Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
     */
    prefixMode?: pulumi.Input<string>;
    /**
     * DHCP prefix configuration. The structure of `prefixRange` block is documented below.
     */
    prefixRanges?: pulumi.Input<pulumi.Input<inputs.system.dhcp6.ServerPrefixRange>[]>;
    /**
     * Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
     */
    rapidCommit?: pulumi.Input<string>;
    /**
     * Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Subnet or subnet-id if the IP mode is delegated.
     */
    subnet?: pulumi.Input<string>;
    /**
     * Interface name from where delegated information is provided.
     */
    upstreamInterface?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Server resource.
 */
export interface ServerArgs {
    /**
     * IAID of obtained delegated-prefix from the upstream interface.
     */
    delegatedPrefixIaid?: pulumi.Input<number>;
    /**
     * DNS search list options. Valid values: `delegated`, `specify`.
     */
    dnsSearchList?: pulumi.Input<string>;
    /**
     * DNS server 1.
     */
    dnsServer1?: pulumi.Input<string>;
    /**
     * DNS server 2.
     */
    dnsServer2?: pulumi.Input<string>;
    /**
     * DNS server 3.
     */
    dnsServer3?: pulumi.Input<string>;
    /**
     * DNS server 4.
     */
    dnsServer4?: pulumi.Input<string>;
    /**
     * Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
     */
    dnsService?: pulumi.Input<string>;
    /**
     * Domain name suffix for the IP addresses that the DHCP server assigns to clients.
     */
    domain?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * DHCP server can assign IP configurations to clients connected to this interface.
     */
    interface: pulumi.Input<string>;
    /**
     * Method used to assign client IP. Valid values: `range`, `delegated`.
     */
    ipMode?: pulumi.Input<string>;
    /**
     * DHCP IP range configuration. The structure of `ipRange` block is documented below.
     */
    ipRanges?: pulumi.Input<pulumi.Input<inputs.system.dhcp6.ServerIpRange>[]>;
    /**
     * Lease time in seconds, 0 means unlimited.
     */
    leaseTime?: pulumi.Input<number>;
    /**
     * Option 1.
     */
    option1?: pulumi.Input<string>;
    /**
     * Option 2.
     */
    option2?: pulumi.Input<string>;
    /**
     * Option 3.
     */
    option3?: pulumi.Input<string>;
    /**
     * Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
     */
    prefixMode?: pulumi.Input<string>;
    /**
     * DHCP prefix configuration. The structure of `prefixRange` block is documented below.
     */
    prefixRanges?: pulumi.Input<pulumi.Input<inputs.system.dhcp6.ServerPrefixRange>[]>;
    /**
     * Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
     */
    rapidCommit?: pulumi.Input<string>;
    /**
     * Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Subnet or subnet-id if the IP mode is delegated.
     */
    subnet: pulumi.Input<string>;
    /**
     * Interface name from where delegated information is provided.
     */
    upstreamInterface?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
