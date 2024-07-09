// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure multicast addresses.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Multicastaddress("trname", {
 *     color: 0,
 *     endIp: "224.0.0.22",
 *     startIp: "224.0.0.11",
 *     subnet: "224.0.0.11 224.0.0.22",
 *     type: "multicastrange",
 *     visibility: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall MulticastAddress can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/multicastaddress:Multicastaddress labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/multicastaddress:Multicastaddress labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Multicastaddress extends pulumi.CustomResource {
    /**
     * Get an existing Multicastaddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MulticastaddressState, opts?: pulumi.CustomResourceOptions): Multicastaddress {
        return new Multicastaddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/multicastaddress:Multicastaddress';

    /**
     * Returns true if the given object is an instance of Multicastaddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Multicastaddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Multicastaddress.__pulumiType;
    }

    /**
     * Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
     */
    public readonly associatedInterface!: pulumi.Output<string>;
    /**
     * Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
     */
    public readonly color!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Final IPv4 address (inclusive) in the range for the address.
     */
    public readonly endIp!: pulumi.Output<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Multicast address name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * First IPv4 address (inclusive) in the range for the address.
     */
    public readonly startIp!: pulumi.Output<string>;
    /**
     * Broadcast address and subnet.
     */
    public readonly subnet!: pulumi.Output<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    public readonly taggings!: pulumi.Output<outputs.firewall.MulticastaddressTagging[] | undefined>;
    /**
     * Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address. Valid values: `multicastrange`, `broadcastmask`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable visibility of the multicast address on the GUI. Valid values: `enable`, `disable`.
     */
    public readonly visibility!: pulumi.Output<string>;

    /**
     * Create a Multicastaddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MulticastaddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MulticastaddressArgs | MulticastaddressState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MulticastaddressState | undefined;
            resourceInputs["associatedInterface"] = state ? state.associatedInterface : undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["endIp"] = state ? state.endIp : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["startIp"] = state ? state.startIp : undefined;
            resourceInputs["subnet"] = state ? state.subnet : undefined;
            resourceInputs["taggings"] = state ? state.taggings : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as MulticastaddressArgs | undefined;
            resourceInputs["associatedInterface"] = args ? args.associatedInterface : undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["endIp"] = args ? args.endIp : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["startIp"] = args ? args.startIp : undefined;
            resourceInputs["subnet"] = args ? args.subnet : undefined;
            resourceInputs["taggings"] = args ? args.taggings : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Multicastaddress.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Multicastaddress resources.
 */
export interface MulticastaddressState {
    /**
     * Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
     */
    associatedInterface?: pulumi.Input<string>;
    /**
     * Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Final IPv4 address (inclusive) in the range for the address.
     */
    endIp?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Multicast address name.
     */
    name?: pulumi.Input<string>;
    /**
     * First IPv4 address (inclusive) in the range for the address.
     */
    startIp?: pulumi.Input<string>;
    /**
     * Broadcast address and subnet.
     */
    subnet?: pulumi.Input<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.firewall.MulticastaddressTagging>[]>;
    /**
     * Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address. Valid values: `multicastrange`, `broadcastmask`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable visibility of the multicast address on the GUI. Valid values: `enable`, `disable`.
     */
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Multicastaddress resource.
 */
export interface MulticastaddressArgs {
    /**
     * Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
     */
    associatedInterface?: pulumi.Input<string>;
    /**
     * Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Final IPv4 address (inclusive) in the range for the address.
     */
    endIp?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Multicast address name.
     */
    name?: pulumi.Input<string>;
    /**
     * First IPv4 address (inclusive) in the range for the address.
     */
    startIp?: pulumi.Input<string>;
    /**
     * Broadcast address and subnet.
     */
    subnet?: pulumi.Input<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.firewall.MulticastaddressTagging>[]>;
    /**
     * Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address. Valid values: `multicastrange`, `broadcastmask`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable visibility of the multicast address on the GUI. Valid values: `enable`, `disable`.
     */
    visibility?: pulumi.Input<string>;
}
