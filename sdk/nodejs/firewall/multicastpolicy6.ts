// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure IPv6 multicast NAT policies.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Multicastpolicy6("trname", {
 *     action: "accept",
 *     dstaddrs: [{
 *         name: "all",
 *     }],
 *     dstintf: "port4",
 *     endPort: 65535,
 *     fosid: 1,
 *     logtraffic: "disable",
 *     protocol: 0,
 *     srcaddrs: [{
 *         name: "all",
 *     }],
 *     srcintf: "port3",
 *     startPort: 1,
 *     status: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall MulticastPolicy6 can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/multicastpolicy6:Multicastpolicy6 labelname {{fosid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/multicastpolicy6:Multicastpolicy6 labelname {{fosid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Multicastpolicy6 extends pulumi.CustomResource {
    /**
     * Get an existing Multicastpolicy6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Multicastpolicy6State, opts?: pulumi.CustomResourceOptions): Multicastpolicy6 {
        return new Multicastpolicy6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/multicastpolicy6:Multicastpolicy6';

    /**
     * Returns true if the given object is an instance of Multicastpolicy6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Multicastpolicy6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Multicastpolicy6.__pulumiType;
    }

    /**
     * Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
     */
    public readonly autoAsicOffload!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * IPv6 destination address name. The structure of `dstaddr` block is documented below.
     */
    public readonly dstaddrs!: pulumi.Output<outputs.firewall.Multicastpolicy6Dstaddr[]>;
    /**
     * IPv6 destination interface name.
     */
    public readonly dstintf!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
     */
    public readonly endPort!: pulumi.Output<number>;
    /**
     * Policy ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Name of an existing IPS sensor.
     */
    public readonly ipsSensor!: pulumi.Output<string>;
    /**
     * Enable/disable logging traffic accepted by this policy.
     */
    public readonly logtraffic!: pulumi.Output<string>;
    /**
     * Policy name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
     */
    public readonly protocol!: pulumi.Output<number>;
    /**
     * IPv6 source address name. The structure of `srcaddr` block is documented below.
     */
    public readonly srcaddrs!: pulumi.Output<outputs.firewall.Multicastpolicy6Srcaddr[]>;
    /**
     * IPv6 source interface name.
     */
    public readonly srcintf!: pulumi.Output<string>;
    /**
     * Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
     */
    public readonly startPort!: pulumi.Output<number>;
    /**
     * Enable/disable this policy. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
     */
    public readonly utmStatus!: pulumi.Output<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    public readonly uuid!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Multicastpolicy6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Multicastpolicy6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Multicastpolicy6Args | Multicastpolicy6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Multicastpolicy6State | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["autoAsicOffload"] = state ? state.autoAsicOffload : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dstaddrs"] = state ? state.dstaddrs : undefined;
            resourceInputs["dstintf"] = state ? state.dstintf : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["endPort"] = state ? state.endPort : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["ipsSensor"] = state ? state.ipsSensor : undefined;
            resourceInputs["logtraffic"] = state ? state.logtraffic : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["srcaddrs"] = state ? state.srcaddrs : undefined;
            resourceInputs["srcintf"] = state ? state.srcintf : undefined;
            resourceInputs["startPort"] = state ? state.startPort : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["utmStatus"] = state ? state.utmStatus : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Multicastpolicy6Args | undefined;
            if ((!args || args.dstaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstaddrs'");
            }
            if ((!args || args.dstintf === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstintf'");
            }
            if ((!args || args.srcaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcaddrs'");
            }
            if ((!args || args.srcintf === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcintf'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["autoAsicOffload"] = args ? args.autoAsicOffload : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dstaddrs"] = args ? args.dstaddrs : undefined;
            resourceInputs["dstintf"] = args ? args.dstintf : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["endPort"] = args ? args.endPort : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["ipsSensor"] = args ? args.ipsSensor : undefined;
            resourceInputs["logtraffic"] = args ? args.logtraffic : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["srcaddrs"] = args ? args.srcaddrs : undefined;
            resourceInputs["srcintf"] = args ? args.srcintf : undefined;
            resourceInputs["startPort"] = args ? args.startPort : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["utmStatus"] = args ? args.utmStatus : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Multicastpolicy6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Multicastpolicy6 resources.
 */
export interface Multicastpolicy6State {
    /**
     * Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
     */
    action?: pulumi.Input<string>;
    /**
     * Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
     */
    autoAsicOffload?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * IPv6 destination address name. The structure of `dstaddr` block is documented below.
     */
    dstaddrs?: pulumi.Input<pulumi.Input<inputs.firewall.Multicastpolicy6Dstaddr>[]>;
    /**
     * IPv6 destination interface name.
     */
    dstintf?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
     */
    endPort?: pulumi.Input<number>;
    /**
     * Policy ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Name of an existing IPS sensor.
     */
    ipsSensor?: pulumi.Input<string>;
    /**
     * Enable/disable logging traffic accepted by this policy.
     */
    logtraffic?: pulumi.Input<string>;
    /**
     * Policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
     */
    protocol?: pulumi.Input<number>;
    /**
     * IPv6 source address name. The structure of `srcaddr` block is documented below.
     */
    srcaddrs?: pulumi.Input<pulumi.Input<inputs.firewall.Multicastpolicy6Srcaddr>[]>;
    /**
     * IPv6 source interface name.
     */
    srcintf?: pulumi.Input<string>;
    /**
     * Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
     */
    startPort?: pulumi.Input<number>;
    /**
     * Enable/disable this policy. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
     */
    utmStatus?: pulumi.Input<string>;
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
 * The set of arguments for constructing a Multicastpolicy6 resource.
 */
export interface Multicastpolicy6Args {
    /**
     * Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
     */
    action?: pulumi.Input<string>;
    /**
     * Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
     */
    autoAsicOffload?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * IPv6 destination address name. The structure of `dstaddr` block is documented below.
     */
    dstaddrs: pulumi.Input<pulumi.Input<inputs.firewall.Multicastpolicy6Dstaddr>[]>;
    /**
     * IPv6 destination interface name.
     */
    dstintf: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
     */
    endPort?: pulumi.Input<number>;
    /**
     * Policy ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Name of an existing IPS sensor.
     */
    ipsSensor?: pulumi.Input<string>;
    /**
     * Enable/disable logging traffic accepted by this policy.
     */
    logtraffic?: pulumi.Input<string>;
    /**
     * Policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
     */
    protocol?: pulumi.Input<number>;
    /**
     * IPv6 source address name. The structure of `srcaddr` block is documented below.
     */
    srcaddrs: pulumi.Input<pulumi.Input<inputs.firewall.Multicastpolicy6Srcaddr>[]>;
    /**
     * IPv6 source interface name.
     */
    srcintf: pulumi.Input<string>;
    /**
     * Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
     */
    startPort?: pulumi.Input<number>;
    /**
     * Enable/disable this policy. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
     */
    utmStatus?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
