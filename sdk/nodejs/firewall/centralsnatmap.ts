// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure central SNAT policies.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Centralsnatmap("trname", {
 *     dstAddrs: [{
 *         name: "all",
 *     }],
 *     dstintfs: [{
 *         name: "port3",
 *     }],
 *     nat: "enable",
 *     natPort: "0",
 *     origAddrs: [{
 *         name: "all",
 *     }],
 *     origPort: "0",
 *     policyid: 1,
 *     protocol: 33,
 *     srcintfs: [{
 *         name: "port1",
 *     }],
 *     status: "enable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Firewall CentralSnatMap can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/centralsnatmap:Centralsnatmap labelname {{policyid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/centralsnatmap:Centralsnatmap labelname {{policyid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Centralsnatmap extends pulumi.CustomResource {
    /**
     * Get an existing Centralsnatmap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CentralsnatmapState, opts?: pulumi.CustomResourceOptions): Centralsnatmap {
        return new Centralsnatmap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/centralsnatmap:Centralsnatmap';

    /**
     * Returns true if the given object is an instance of Centralsnatmap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Centralsnatmap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Centralsnatmap.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * IPv6 Destination address. The structure of `dstAddr6` block is documented below.
     */
    public readonly dstAddr6s!: pulumi.Output<outputs.firewall.CentralsnatmapDstAddr6[] | undefined>;
    /**
     * Destination address name from available addresses. The structure of `dstAddr` block is documented below.
     */
    public readonly dstAddrs!: pulumi.Output<outputs.firewall.CentralsnatmapDstAddr[]>;
    /**
     * Destination port or port range (1 to 65535, 0 means any port).
     */
    public readonly dstPort!: pulumi.Output<string>;
    /**
     * Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
     */
    public readonly dstintfs!: pulumi.Output<outputs.firewall.CentralsnatmapDstintf[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable source NAT. Valid values: `disable`, `enable`.
     */
    public readonly nat!: pulumi.Output<string>;
    /**
     * Enable/disable NAT46. Valid values: `enable`, `disable`.
     */
    public readonly nat46!: pulumi.Output<string>;
    /**
     * Enable/disable NAT64. Valid values: `enable`, `disable`.
     */
    public readonly nat64!: pulumi.Output<string>;
    /**
     * IPv6 pools to be used for source NAT. The structure of `natIppool6` block is documented below.
     */
    public readonly natIppool6s!: pulumi.Output<outputs.firewall.CentralsnatmapNatIppool6[] | undefined>;
    /**
     * Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `natIppool` block is documented below.
     */
    public readonly natIppools!: pulumi.Output<outputs.firewall.CentralsnatmapNatIppool[] | undefined>;
    /**
     * Translated port or port range (0 to 65535, 0 means any port).
     */
    public readonly natPort!: pulumi.Output<string>;
    /**
     * IPv6 Original address. The structure of `origAddr6` block is documented below.
     */
    public readonly origAddr6s!: pulumi.Output<outputs.firewall.CentralsnatmapOrigAddr6[] | undefined>;
    /**
     * Original address. The structure of `origAddr` block is documented below.
     */
    public readonly origAddrs!: pulumi.Output<outputs.firewall.CentralsnatmapOrigAddr[]>;
    /**
     * Original TCP port (1 to 65535, 0 means any port).
     */
    public readonly origPort!: pulumi.Output<string>;
    /**
     * Policy ID.
     */
    public readonly policyid!: pulumi.Output<number>;
    /**
     * Integer value for the protocol type (0 - 255).
     */
    public readonly protocol!: pulumi.Output<number>;
    /**
     * Source interface name from available interfaces. The structure of `srcintf` block is documented below.
     */
    public readonly srcintfs!: pulumi.Output<outputs.firewall.CentralsnatmapSrcintf[]>;
    /**
     * Enable/disable the active status of this policy. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * IPv4/IPv6 source NAT. Valid values: `ipv4`, `ipv6`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    public readonly uuid!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Centralsnatmap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CentralsnatmapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CentralsnatmapArgs | CentralsnatmapState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CentralsnatmapState | undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dstAddr6s"] = state ? state.dstAddr6s : undefined;
            resourceInputs["dstAddrs"] = state ? state.dstAddrs : undefined;
            resourceInputs["dstPort"] = state ? state.dstPort : undefined;
            resourceInputs["dstintfs"] = state ? state.dstintfs : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["nat"] = state ? state.nat : undefined;
            resourceInputs["nat46"] = state ? state.nat46 : undefined;
            resourceInputs["nat64"] = state ? state.nat64 : undefined;
            resourceInputs["natIppool6s"] = state ? state.natIppool6s : undefined;
            resourceInputs["natIppools"] = state ? state.natIppools : undefined;
            resourceInputs["natPort"] = state ? state.natPort : undefined;
            resourceInputs["origAddr6s"] = state ? state.origAddr6s : undefined;
            resourceInputs["origAddrs"] = state ? state.origAddrs : undefined;
            resourceInputs["origPort"] = state ? state.origPort : undefined;
            resourceInputs["policyid"] = state ? state.policyid : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["srcintfs"] = state ? state.srcintfs : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as CentralsnatmapArgs | undefined;
            if ((!args || args.dstAddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstAddrs'");
            }
            if ((!args || args.dstintfs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstintfs'");
            }
            if ((!args || args.nat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nat'");
            }
            if ((!args || args.origAddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'origAddrs'");
            }
            if ((!args || args.origPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'origPort'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.srcintfs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcintfs'");
            }
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dstAddr6s"] = args ? args.dstAddr6s : undefined;
            resourceInputs["dstAddrs"] = args ? args.dstAddrs : undefined;
            resourceInputs["dstPort"] = args ? args.dstPort : undefined;
            resourceInputs["dstintfs"] = args ? args.dstintfs : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["nat"] = args ? args.nat : undefined;
            resourceInputs["nat46"] = args ? args.nat46 : undefined;
            resourceInputs["nat64"] = args ? args.nat64 : undefined;
            resourceInputs["natIppool6s"] = args ? args.natIppool6s : undefined;
            resourceInputs["natIppools"] = args ? args.natIppools : undefined;
            resourceInputs["natPort"] = args ? args.natPort : undefined;
            resourceInputs["origAddr6s"] = args ? args.origAddr6s : undefined;
            resourceInputs["origAddrs"] = args ? args.origAddrs : undefined;
            resourceInputs["origPort"] = args ? args.origPort : undefined;
            resourceInputs["policyid"] = args ? args.policyid : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["srcintfs"] = args ? args.srcintfs : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Centralsnatmap.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Centralsnatmap resources.
 */
export interface CentralsnatmapState {
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * IPv6 Destination address. The structure of `dstAddr6` block is documented below.
     */
    dstAddr6s?: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapDstAddr6>[]>;
    /**
     * Destination address name from available addresses. The structure of `dstAddr` block is documented below.
     */
    dstAddrs?: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapDstAddr>[]>;
    /**
     * Destination port or port range (1 to 65535, 0 means any port).
     */
    dstPort?: pulumi.Input<string>;
    /**
     * Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
     */
    dstintfs?: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapDstintf>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Enable/disable source NAT. Valid values: `disable`, `enable`.
     */
    nat?: pulumi.Input<string>;
    /**
     * Enable/disable NAT46. Valid values: `enable`, `disable`.
     */
    nat46?: pulumi.Input<string>;
    /**
     * Enable/disable NAT64. Valid values: `enable`, `disable`.
     */
    nat64?: pulumi.Input<string>;
    /**
     * IPv6 pools to be used for source NAT. The structure of `natIppool6` block is documented below.
     */
    natIppool6s?: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapNatIppool6>[]>;
    /**
     * Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `natIppool` block is documented below.
     */
    natIppools?: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapNatIppool>[]>;
    /**
     * Translated port or port range (0 to 65535, 0 means any port).
     */
    natPort?: pulumi.Input<string>;
    /**
     * IPv6 Original address. The structure of `origAddr6` block is documented below.
     */
    origAddr6s?: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapOrigAddr6>[]>;
    /**
     * Original address. The structure of `origAddr` block is documented below.
     */
    origAddrs?: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapOrigAddr>[]>;
    /**
     * Original TCP port (1 to 65535, 0 means any port).
     */
    origPort?: pulumi.Input<string>;
    /**
     * Policy ID.
     */
    policyid?: pulumi.Input<number>;
    /**
     * Integer value for the protocol type (0 - 255).
     */
    protocol?: pulumi.Input<number>;
    /**
     * Source interface name from available interfaces. The structure of `srcintf` block is documented below.
     */
    srcintfs?: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapSrcintf>[]>;
    /**
     * Enable/disable the active status of this policy. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * IPv4/IPv6 source NAT. Valid values: `ipv4`, `ipv6`.
     */
    type?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Centralsnatmap resource.
 */
export interface CentralsnatmapArgs {
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * IPv6 Destination address. The structure of `dstAddr6` block is documented below.
     */
    dstAddr6s?: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapDstAddr6>[]>;
    /**
     * Destination address name from available addresses. The structure of `dstAddr` block is documented below.
     */
    dstAddrs: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapDstAddr>[]>;
    /**
     * Destination port or port range (1 to 65535, 0 means any port).
     */
    dstPort?: pulumi.Input<string>;
    /**
     * Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
     */
    dstintfs: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapDstintf>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Enable/disable source NAT. Valid values: `disable`, `enable`.
     */
    nat: pulumi.Input<string>;
    /**
     * Enable/disable NAT46. Valid values: `enable`, `disable`.
     */
    nat46?: pulumi.Input<string>;
    /**
     * Enable/disable NAT64. Valid values: `enable`, `disable`.
     */
    nat64?: pulumi.Input<string>;
    /**
     * IPv6 pools to be used for source NAT. The structure of `natIppool6` block is documented below.
     */
    natIppool6s?: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapNatIppool6>[]>;
    /**
     * Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `natIppool` block is documented below.
     */
    natIppools?: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapNatIppool>[]>;
    /**
     * Translated port or port range (0 to 65535, 0 means any port).
     */
    natPort?: pulumi.Input<string>;
    /**
     * IPv6 Original address. The structure of `origAddr6` block is documented below.
     */
    origAddr6s?: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapOrigAddr6>[]>;
    /**
     * Original address. The structure of `origAddr` block is documented below.
     */
    origAddrs: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapOrigAddr>[]>;
    /**
     * Original TCP port (1 to 65535, 0 means any port).
     */
    origPort: pulumi.Input<string>;
    /**
     * Policy ID.
     */
    policyid?: pulumi.Input<number>;
    /**
     * Integer value for the protocol type (0 - 255).
     */
    protocol: pulumi.Input<number>;
    /**
     * Source interface name from available interfaces. The structure of `srcintf` block is documented below.
     */
    srcintfs: pulumi.Input<pulumi.Input<inputs.firewall.CentralsnatmapSrcintf>[]>;
    /**
     * Enable/disable the active status of this policy. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * IPv4/IPv6 source NAT. Valid values: `ipv4`, `ipv6`.
     */
    type?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
