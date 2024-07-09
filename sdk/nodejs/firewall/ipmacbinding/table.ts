// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Configure IP to MAC address pairs in the IP/MAC binding table.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.ipmacbinding.Table("trname", {
 *     ip: "1.1.1.1",
 *     mac: "00:01:6c:06:a6:29",
 *     seqNum: 1,
 *     status: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * FirewallIpmacbinding Table can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/ipmacbinding/table:Table labelname {{seq_num}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/ipmacbinding/table:Table labelname {{seq_num}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableState, opts?: pulumi.CustomResourceOptions): Table {
        return new Table(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/ipmacbinding/table:Table';

    /**
     * Returns true if the given object is an instance of Table.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Table {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Table.__pulumiType;
    }

    /**
     * IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * MAC address portion of the pair (format = xx:xx:xx:xx:xx:xx in hexadecimal).
     */
    public readonly mac!: pulumi.Output<string>;
    /**
     * Name of the pair (optional, default = no name).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Entry number.
     */
    public readonly seqNum!: pulumi.Output<number>;
    /**
     * Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableArgs | TableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TableState | undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["mac"] = state ? state.mac : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["seqNum"] = state ? state.seqNum : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as TableArgs | undefined;
            if ((!args || args.ip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip'");
            }
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["mac"] = args ? args.mac : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["seqNum"] = args ? args.seqNum : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Table.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Table resources.
 */
export interface TableState {
    /**
     * IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
     */
    ip?: pulumi.Input<string>;
    /**
     * MAC address portion of the pair (format = xx:xx:xx:xx:xx:xx in hexadecimal).
     */
    mac?: pulumi.Input<string>;
    /**
     * Name of the pair (optional, default = no name).
     */
    name?: pulumi.Input<string>;
    /**
     * Entry number.
     */
    seqNum?: pulumi.Input<number>;
    /**
     * Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    /**
     * IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
     */
    ip: pulumi.Input<string>;
    /**
     * MAC address portion of the pair (format = xx:xx:xx:xx:xx:xx in hexadecimal).
     */
    mac?: pulumi.Input<string>;
    /**
     * Name of the pair (optional, default = no name).
     */
    name?: pulumi.Input<string>;
    /**
     * Entry number.
     */
    seqNum?: pulumi.Input<number>;
    /**
     * Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
