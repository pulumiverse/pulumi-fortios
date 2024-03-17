// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * OSPF interface configuration.
 *
 * > The provider supports the definition of Ospf-Interface in Router Ospf `fortios.router.Ospf`, and also allows the definition of separate Ospf-Interface resources `fortios.router/ospf.Ospfinterface`, but do not use a `fortios.router.Ospf` with in-line Ospf-Interface in conjunction with any `fortios.router/ospf.Ospfinterface` resources, otherwise conflicts and overwrite will occur.
 *
 * ## Import
 *
 * Routerospf OspfInterface can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:router/ospf/ospfinterface:Ospfinterface labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:router/ospf/ospfinterface:Ospfinterface labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ospfinterface extends pulumi.CustomResource {
    /**
     * Get an existing Ospfinterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OspfinterfaceState, opts?: pulumi.CustomResourceOptions): Ospfinterface {
        return new Ospfinterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:router/ospf/ospfinterface:Ospfinterface';

    /**
     * Returns true if the given object is an instance of Ospfinterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ospfinterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ospfinterface.__pulumiType;
    }

    /**
     * Authentication type.
     */
    public readonly authentication!: pulumi.Output<string>;
    /**
     * Authentication key.
     */
    public readonly authenticationKey!: pulumi.Output<string | undefined>;
    /**
     * Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
     */
    public readonly bfd!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
     */
    public readonly cost!: pulumi.Output<number>;
    /**
     * Enable/disable control of flooding out LSAs. Valid values: `enable`, `disable`.
     */
    public readonly databaseFilterOut!: pulumi.Output<string>;
    /**
     * Dead interval.
     */
    public readonly deadInterval!: pulumi.Output<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Hello interval.
     */
    public readonly helloInterval!: pulumi.Output<number>;
    /**
     * Number of hello packets within dead interval.
     */
    public readonly helloMultiplier!: pulumi.Output<number>;
    /**
     * Configuration interface name.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * IP address.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Message-digest key-chain name.
     */
    public readonly keychain!: pulumi.Output<string>;
    /**
     * MD5 key.
     */
    public readonly md5Key!: pulumi.Output<string>;
    /**
     * Authentication MD5 key-chain name.
     */
    public readonly md5Keychain!: pulumi.Output<string>;
    /**
     * MD5 key. The structure of `md5Keys` block is documented below.
     */
    public readonly md5Keys!: pulumi.Output<outputs.router.ospf.OspfinterfaceMd5Key[] | undefined>;
    /**
     * MTU for database description packets.
     */
    public readonly mtu!: pulumi.Output<number>;
    /**
     * Enable/disable ignore MTU. Valid values: `enable`, `disable`.
     */
    public readonly mtuIgnore!: pulumi.Output<string>;
    /**
     * Interface entry name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Network type. Valid values: `broadcast`, `non-broadcast`, `point-to-point`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
     */
    public readonly networkType!: pulumi.Output<string>;
    /**
     * Prefix length.
     */
    public readonly prefixLength!: pulumi.Output<number>;
    /**
     * Priority.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Graceful restart neighbor resynchronization timeout.
     */
    public readonly resyncTimeout!: pulumi.Output<number>;
    /**
     * Retransmit interval.
     */
    public readonly retransmitInterval!: pulumi.Output<number>;
    /**
     * Enable/disable status. Valid values: `disable`, `enable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Transmit delay.
     */
    public readonly transmitDelay!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     *
     * The `md5Keys` block supports:
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Ospfinterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OspfinterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OspfinterfaceArgs | OspfinterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OspfinterfaceState | undefined;
            resourceInputs["authentication"] = state ? state.authentication : undefined;
            resourceInputs["authenticationKey"] = state ? state.authenticationKey : undefined;
            resourceInputs["bfd"] = state ? state.bfd : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["cost"] = state ? state.cost : undefined;
            resourceInputs["databaseFilterOut"] = state ? state.databaseFilterOut : undefined;
            resourceInputs["deadInterval"] = state ? state.deadInterval : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["helloInterval"] = state ? state.helloInterval : undefined;
            resourceInputs["helloMultiplier"] = state ? state.helloMultiplier : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["keychain"] = state ? state.keychain : undefined;
            resourceInputs["md5Key"] = state ? state.md5Key : undefined;
            resourceInputs["md5Keychain"] = state ? state.md5Keychain : undefined;
            resourceInputs["md5Keys"] = state ? state.md5Keys : undefined;
            resourceInputs["mtu"] = state ? state.mtu : undefined;
            resourceInputs["mtuIgnore"] = state ? state.mtuIgnore : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["prefixLength"] = state ? state.prefixLength : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["resyncTimeout"] = state ? state.resyncTimeout : undefined;
            resourceInputs["retransmitInterval"] = state ? state.retransmitInterval : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["transmitDelay"] = state ? state.transmitDelay : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as OspfinterfaceArgs | undefined;
            resourceInputs["authentication"] = args ? args.authentication : undefined;
            resourceInputs["authenticationKey"] = args ? args.authenticationKey : undefined;
            resourceInputs["bfd"] = args ? args.bfd : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["cost"] = args ? args.cost : undefined;
            resourceInputs["databaseFilterOut"] = args ? args.databaseFilterOut : undefined;
            resourceInputs["deadInterval"] = args ? args.deadInterval : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["helloInterval"] = args ? args.helloInterval : undefined;
            resourceInputs["helloMultiplier"] = args ? args.helloMultiplier : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["keychain"] = args ? args.keychain : undefined;
            resourceInputs["md5Key"] = args ? args.md5Key : undefined;
            resourceInputs["md5Keychain"] = args ? args.md5Keychain : undefined;
            resourceInputs["md5Keys"] = args ? args.md5Keys : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["mtuIgnore"] = args ? args.mtuIgnore : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["prefixLength"] = args ? args.prefixLength : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["resyncTimeout"] = args ? args.resyncTimeout : undefined;
            resourceInputs["retransmitInterval"] = args ? args.retransmitInterval : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["transmitDelay"] = args ? args.transmitDelay : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ospfinterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ospfinterface resources.
 */
export interface OspfinterfaceState {
    /**
     * Authentication type.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Authentication key.
     */
    authenticationKey?: pulumi.Input<string>;
    /**
     * Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
     */
    bfd?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
     */
    cost?: pulumi.Input<number>;
    /**
     * Enable/disable control of flooding out LSAs. Valid values: `enable`, `disable`.
     */
    databaseFilterOut?: pulumi.Input<string>;
    /**
     * Dead interval.
     */
    deadInterval?: pulumi.Input<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Hello interval.
     */
    helloInterval?: pulumi.Input<number>;
    /**
     * Number of hello packets within dead interval.
     */
    helloMultiplier?: pulumi.Input<number>;
    /**
     * Configuration interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * IP address.
     */
    ip?: pulumi.Input<string>;
    /**
     * Message-digest key-chain name.
     */
    keychain?: pulumi.Input<string>;
    /**
     * MD5 key.
     */
    md5Key?: pulumi.Input<string>;
    /**
     * Authentication MD5 key-chain name.
     */
    md5Keychain?: pulumi.Input<string>;
    /**
     * MD5 key. The structure of `md5Keys` block is documented below.
     */
    md5Keys?: pulumi.Input<pulumi.Input<inputs.router.ospf.OspfinterfaceMd5Key>[]>;
    /**
     * MTU for database description packets.
     */
    mtu?: pulumi.Input<number>;
    /**
     * Enable/disable ignore MTU. Valid values: `enable`, `disable`.
     */
    mtuIgnore?: pulumi.Input<string>;
    /**
     * Interface entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Network type. Valid values: `broadcast`, `non-broadcast`, `point-to-point`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * Prefix length.
     */
    prefixLength?: pulumi.Input<number>;
    /**
     * Priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * Graceful restart neighbor resynchronization timeout.
     */
    resyncTimeout?: pulumi.Input<number>;
    /**
     * Retransmit interval.
     */
    retransmitInterval?: pulumi.Input<number>;
    /**
     * Enable/disable status. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Transmit delay.
     */
    transmitDelay?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     *
     * The `md5Keys` block supports:
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ospfinterface resource.
 */
export interface OspfinterfaceArgs {
    /**
     * Authentication type.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Authentication key.
     */
    authenticationKey?: pulumi.Input<string>;
    /**
     * Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
     */
    bfd?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
     */
    cost?: pulumi.Input<number>;
    /**
     * Enable/disable control of flooding out LSAs. Valid values: `enable`, `disable`.
     */
    databaseFilterOut?: pulumi.Input<string>;
    /**
     * Dead interval.
     */
    deadInterval?: pulumi.Input<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Hello interval.
     */
    helloInterval?: pulumi.Input<number>;
    /**
     * Number of hello packets within dead interval.
     */
    helloMultiplier?: pulumi.Input<number>;
    /**
     * Configuration interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * IP address.
     */
    ip?: pulumi.Input<string>;
    /**
     * Message-digest key-chain name.
     */
    keychain?: pulumi.Input<string>;
    /**
     * MD5 key.
     */
    md5Key?: pulumi.Input<string>;
    /**
     * Authentication MD5 key-chain name.
     */
    md5Keychain?: pulumi.Input<string>;
    /**
     * MD5 key. The structure of `md5Keys` block is documented below.
     */
    md5Keys?: pulumi.Input<pulumi.Input<inputs.router.ospf.OspfinterfaceMd5Key>[]>;
    /**
     * MTU for database description packets.
     */
    mtu?: pulumi.Input<number>;
    /**
     * Enable/disable ignore MTU. Valid values: `enable`, `disable`.
     */
    mtuIgnore?: pulumi.Input<string>;
    /**
     * Interface entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Network type. Valid values: `broadcast`, `non-broadcast`, `point-to-point`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * Prefix length.
     */
    prefixLength?: pulumi.Input<number>;
    /**
     * Priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * Graceful restart neighbor resynchronization timeout.
     */
    resyncTimeout?: pulumi.Input<number>;
    /**
     * Retransmit interval.
     */
    retransmitInterval?: pulumi.Input<number>;
    /**
     * Enable/disable status. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Transmit delay.
     */
    transmitDelay?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     *
     * The `md5Keys` block supports:
     */
    vdomparam?: pulumi.Input<string>;
}
