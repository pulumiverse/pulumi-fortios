// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure IPv6 routing policies.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.router.Policy6("trname", {
 *     dst: "::/0",
 *     endPort: 65535,
 *     gateway: "::",
 *     inputDevice: "port1",
 *     outputDevice: "port3",
 *     protocol: 33,
 *     seqNum: 1,
 *     src: "2001:db8:85a3::8a2e:370:7334/128",
 *     startPort: 1,
 *     status: "enable",
 *     tos: "0x00",
 *     tosMask: "0x00",
 * });
 * ```
 *
 * ## Import
 *
 * Router Policy6 can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:router/policy6:Policy6 labelname {{seq_num}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:router/policy6:Policy6 labelname {{seq_num}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Policy6 extends pulumi.CustomResource {
    /**
     * Get an existing Policy6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Policy6State, opts?: pulumi.CustomResourceOptions): Policy6 {
        return new Policy6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:router/policy6:Policy6';

    /**
     * Returns true if the given object is an instance of Policy6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy6.__pulumiType;
    }

    /**
     * Action of the policy route. Valid values: `deny`, `permit`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Optional comments.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Destination IPv6 prefix.
     */
    public readonly dst!: pulumi.Output<string>;
    /**
     * Enable/disable negating destination address match. Valid values: `enable`, `disable`.
     */
    public readonly dstNegate!: pulumi.Output<string>;
    /**
     * Destination address name. The structure of `dstaddr` block is documented below.
     */
    public readonly dstaddrs!: pulumi.Output<outputs.router.Policy6Dstaddr[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * End destination port number (1 - 65535).
     */
    public readonly endPort!: pulumi.Output<number>;
    /**
     * End source port number (1 - 65535).
     */
    public readonly endSourcePort!: pulumi.Output<number>;
    /**
     * IPv6 address of the gateway.
     */
    public readonly gateway!: pulumi.Output<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Incoming interface name.
     */
    public readonly inputDevice!: pulumi.Output<string>;
    /**
     * Enable/disable negation of input device match. Valid values: `enable`, `disable`.
     */
    public readonly inputDeviceNegate!: pulumi.Output<string>;
    /**
     * Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
     */
    public readonly internetServiceCustoms!: pulumi.Output<outputs.router.Policy6InternetServiceCustom[] | undefined>;
    /**
     * Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
     */
    public readonly internetServiceIds!: pulumi.Output<outputs.router.Policy6InternetServiceId[] | undefined>;
    /**
     * Outgoing interface name.
     */
    public readonly outputDevice!: pulumi.Output<string>;
    /**
     * Protocol number (0 - 255).
     */
    public readonly protocol!: pulumi.Output<number>;
    /**
     * Sequence number.
     */
    public readonly seqNum!: pulumi.Output<number>;
    /**
     * Source IPv6 prefix.
     */
    public readonly src!: pulumi.Output<string>;
    /**
     * Enable/disable negating source address match. Valid values: `enable`, `disable`.
     */
    public readonly srcNegate!: pulumi.Output<string>;
    /**
     * Source address name. The structure of `srcaddr` block is documented below.
     */
    public readonly srcaddrs!: pulumi.Output<outputs.router.Policy6Srcaddr[] | undefined>;
    /**
     * Start destination port number (1 - 65535).
     */
    public readonly startPort!: pulumi.Output<number>;
    /**
     * Start source port number (1 - 65535).
     */
    public readonly startSourcePort!: pulumi.Output<number>;
    /**
     * Enable/disable this policy route. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Type of service bit pattern.
     */
    public readonly tos!: pulumi.Output<string>;
    /**
     * Type of service evaluated bits.
     */
    public readonly tosMask!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Policy6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Policy6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Policy6Args | Policy6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Policy6State | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dst"] = state ? state.dst : undefined;
            resourceInputs["dstNegate"] = state ? state.dstNegate : undefined;
            resourceInputs["dstaddrs"] = state ? state.dstaddrs : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["endPort"] = state ? state.endPort : undefined;
            resourceInputs["endSourcePort"] = state ? state.endSourcePort : undefined;
            resourceInputs["gateway"] = state ? state.gateway : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["inputDevice"] = state ? state.inputDevice : undefined;
            resourceInputs["inputDeviceNegate"] = state ? state.inputDeviceNegate : undefined;
            resourceInputs["internetServiceCustoms"] = state ? state.internetServiceCustoms : undefined;
            resourceInputs["internetServiceIds"] = state ? state.internetServiceIds : undefined;
            resourceInputs["outputDevice"] = state ? state.outputDevice : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["seqNum"] = state ? state.seqNum : undefined;
            resourceInputs["src"] = state ? state.src : undefined;
            resourceInputs["srcNegate"] = state ? state.srcNegate : undefined;
            resourceInputs["srcaddrs"] = state ? state.srcaddrs : undefined;
            resourceInputs["startPort"] = state ? state.startPort : undefined;
            resourceInputs["startSourcePort"] = state ? state.startSourcePort : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tos"] = state ? state.tos : undefined;
            resourceInputs["tosMask"] = state ? state.tosMask : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Policy6Args | undefined;
            if ((!args || args.inputDevice === undefined) && !opts.urn) {
                throw new Error("Missing required property 'inputDevice'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dst"] = args ? args.dst : undefined;
            resourceInputs["dstNegate"] = args ? args.dstNegate : undefined;
            resourceInputs["dstaddrs"] = args ? args.dstaddrs : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["endPort"] = args ? args.endPort : undefined;
            resourceInputs["endSourcePort"] = args ? args.endSourcePort : undefined;
            resourceInputs["gateway"] = args ? args.gateway : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["inputDevice"] = args ? args.inputDevice : undefined;
            resourceInputs["inputDeviceNegate"] = args ? args.inputDeviceNegate : undefined;
            resourceInputs["internetServiceCustoms"] = args ? args.internetServiceCustoms : undefined;
            resourceInputs["internetServiceIds"] = args ? args.internetServiceIds : undefined;
            resourceInputs["outputDevice"] = args ? args.outputDevice : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["seqNum"] = args ? args.seqNum : undefined;
            resourceInputs["src"] = args ? args.src : undefined;
            resourceInputs["srcNegate"] = args ? args.srcNegate : undefined;
            resourceInputs["srcaddrs"] = args ? args.srcaddrs : undefined;
            resourceInputs["startPort"] = args ? args.startPort : undefined;
            resourceInputs["startSourcePort"] = args ? args.startSourcePort : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tos"] = args ? args.tos : undefined;
            resourceInputs["tosMask"] = args ? args.tosMask : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Policy6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy6 resources.
 */
export interface Policy6State {
    /**
     * Action of the policy route. Valid values: `deny`, `permit`.
     */
    action?: pulumi.Input<string>;
    /**
     * Optional comments.
     */
    comments?: pulumi.Input<string>;
    /**
     * Destination IPv6 prefix.
     */
    dst?: pulumi.Input<string>;
    /**
     * Enable/disable negating destination address match. Valid values: `enable`, `disable`.
     */
    dstNegate?: pulumi.Input<string>;
    /**
     * Destination address name. The structure of `dstaddr` block is documented below.
     */
    dstaddrs?: pulumi.Input<pulumi.Input<inputs.router.Policy6Dstaddr>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * End destination port number (1 - 65535).
     */
    endPort?: pulumi.Input<number>;
    /**
     * End source port number (1 - 65535).
     */
    endSourcePort?: pulumi.Input<number>;
    /**
     * IPv6 address of the gateway.
     */
    gateway?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Incoming interface name.
     */
    inputDevice?: pulumi.Input<string>;
    /**
     * Enable/disable negation of input device match. Valid values: `enable`, `disable`.
     */
    inputDeviceNegate?: pulumi.Input<string>;
    /**
     * Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
     */
    internetServiceCustoms?: pulumi.Input<pulumi.Input<inputs.router.Policy6InternetServiceCustom>[]>;
    /**
     * Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
     */
    internetServiceIds?: pulumi.Input<pulumi.Input<inputs.router.Policy6InternetServiceId>[]>;
    /**
     * Outgoing interface name.
     */
    outputDevice?: pulumi.Input<string>;
    /**
     * Protocol number (0 - 255).
     */
    protocol?: pulumi.Input<number>;
    /**
     * Sequence number.
     */
    seqNum?: pulumi.Input<number>;
    /**
     * Source IPv6 prefix.
     */
    src?: pulumi.Input<string>;
    /**
     * Enable/disable negating source address match. Valid values: `enable`, `disable`.
     */
    srcNegate?: pulumi.Input<string>;
    /**
     * Source address name. The structure of `srcaddr` block is documented below.
     */
    srcaddrs?: pulumi.Input<pulumi.Input<inputs.router.Policy6Srcaddr>[]>;
    /**
     * Start destination port number (1 - 65535).
     */
    startPort?: pulumi.Input<number>;
    /**
     * Start source port number (1 - 65535).
     */
    startSourcePort?: pulumi.Input<number>;
    /**
     * Enable/disable this policy route. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Type of service bit pattern.
     */
    tos?: pulumi.Input<string>;
    /**
     * Type of service evaluated bits.
     */
    tosMask?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy6 resource.
 */
export interface Policy6Args {
    /**
     * Action of the policy route. Valid values: `deny`, `permit`.
     */
    action?: pulumi.Input<string>;
    /**
     * Optional comments.
     */
    comments?: pulumi.Input<string>;
    /**
     * Destination IPv6 prefix.
     */
    dst?: pulumi.Input<string>;
    /**
     * Enable/disable negating destination address match. Valid values: `enable`, `disable`.
     */
    dstNegate?: pulumi.Input<string>;
    /**
     * Destination address name. The structure of `dstaddr` block is documented below.
     */
    dstaddrs?: pulumi.Input<pulumi.Input<inputs.router.Policy6Dstaddr>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * End destination port number (1 - 65535).
     */
    endPort?: pulumi.Input<number>;
    /**
     * End source port number (1 - 65535).
     */
    endSourcePort?: pulumi.Input<number>;
    /**
     * IPv6 address of the gateway.
     */
    gateway?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Incoming interface name.
     */
    inputDevice: pulumi.Input<string>;
    /**
     * Enable/disable negation of input device match. Valid values: `enable`, `disable`.
     */
    inputDeviceNegate?: pulumi.Input<string>;
    /**
     * Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
     */
    internetServiceCustoms?: pulumi.Input<pulumi.Input<inputs.router.Policy6InternetServiceCustom>[]>;
    /**
     * Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
     */
    internetServiceIds?: pulumi.Input<pulumi.Input<inputs.router.Policy6InternetServiceId>[]>;
    /**
     * Outgoing interface name.
     */
    outputDevice?: pulumi.Input<string>;
    /**
     * Protocol number (0 - 255).
     */
    protocol?: pulumi.Input<number>;
    /**
     * Sequence number.
     */
    seqNum?: pulumi.Input<number>;
    /**
     * Source IPv6 prefix.
     */
    src?: pulumi.Input<string>;
    /**
     * Enable/disable negating source address match. Valid values: `enable`, `disable`.
     */
    srcNegate?: pulumi.Input<string>;
    /**
     * Source address name. The structure of `srcaddr` block is documented below.
     */
    srcaddrs?: pulumi.Input<pulumi.Input<inputs.router.Policy6Srcaddr>[]>;
    /**
     * Start destination port number (1 - 65535).
     */
    startPort?: pulumi.Input<number>;
    /**
     * Start source port number (1 - 65535).
     */
    startSourcePort?: pulumi.Input<number>;
    /**
     * Enable/disable this policy route. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Type of service bit pattern.
     */
    tos?: pulumi.Input<string>;
    /**
     * Type of service evaluated bits.
     */
    tosMask?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
