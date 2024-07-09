// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure system NTP information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Ntp("trname", {
 *     ntpsync: "enable",
 *     serverMode: "disable",
 *     sourceIp: "0.0.0.0",
 *     sourceIp6: "::",
 *     syncinterval: 1,
 *     type: "fortiguard",
 * });
 * ```
 *
 * ## Import
 *
 * System Ntp can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/ntp:Ntp labelname SystemNtp
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/ntp:Ntp labelname SystemNtp
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ntp extends pulumi.CustomResource {
    /**
     * Get an existing Ntp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NtpState, opts?: pulumi.CustomResourceOptions): Ntp {
        return new Ntp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/ntp:Ntp';

    /**
     * Returns true if the given object is an instance of Ntp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ntp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ntp.__pulumiType;
    }

    /**
     * Enable/disable authentication. Valid values: `enable`, `disable`.
     */
    public readonly authentication!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * FortiGate interface(s) with NTP server mode enabled. Devices on your network can contact these interfaces for NTP services. The structure of `interface` block is documented below.
     */
    public readonly interfaces!: pulumi.Output<outputs.system.NtpInterface[] | undefined>;
    /**
     * Key for authentication.
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * Key ID for authentication.
     */
    public readonly keyId!: pulumi.Output<number>;
    /**
     * Key type for authentication. On FortiOS versions 6.2.4-7.4.3: MD5, SHA1. On FortiOS versions >= 7.4.4: MD5, SHA1, SHA256.
     */
    public readonly keyType!: pulumi.Output<string>;
    /**
     * Configure the FortiGate to connect to any available third-party NTP server. The structure of `ntpserver` block is documented below.
     */
    public readonly ntpservers!: pulumi.Output<outputs.system.NtpNtpserver[] | undefined>;
    /**
     * Enable/disable setting the FortiGate system time by synchronizing with an NTP Server. Valid values: `enable`, `disable`.
     */
    public readonly ntpsync!: pulumi.Output<string>;
    /**
     * Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server. Valid values: `enable`, `disable`.
     */
    public readonly serverMode!: pulumi.Output<string>;
    /**
     * Source IP address for communication to the NTP server.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Source IPv6 address for communication to the NTP server.
     */
    public readonly sourceIp6!: pulumi.Output<string>;
    /**
     * NTP synchronization interval (1 - 1440 min).
     */
    public readonly syncinterval!: pulumi.Output<number>;
    /**
     * Use the FortiGuard NTP server or any other available NTP Server. Valid values: `fortiguard`, `custom`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Ntp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NtpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NtpArgs | NtpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NtpState | undefined;
            resourceInputs["authentication"] = state ? state.authentication : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["keyType"] = state ? state.keyType : undefined;
            resourceInputs["ntpservers"] = state ? state.ntpservers : undefined;
            resourceInputs["ntpsync"] = state ? state.ntpsync : undefined;
            resourceInputs["serverMode"] = state ? state.serverMode : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sourceIp6"] = state ? state.sourceIp6 : undefined;
            resourceInputs["syncinterval"] = state ? state.syncinterval : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as NtpArgs | undefined;
            resourceInputs["authentication"] = args ? args.authentication : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["key"] = args?.key ? pulumi.secret(args.key) : undefined;
            resourceInputs["keyId"] = args ? args.keyId : undefined;
            resourceInputs["keyType"] = args ? args.keyType : undefined;
            resourceInputs["ntpservers"] = args ? args.ntpservers : undefined;
            resourceInputs["ntpsync"] = args ? args.ntpsync : undefined;
            resourceInputs["serverMode"] = args ? args.serverMode : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sourceIp6"] = args ? args.sourceIp6 : undefined;
            resourceInputs["syncinterval"] = args ? args.syncinterval : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["key"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Ntp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ntp resources.
 */
export interface NtpState {
    /**
     * Enable/disable authentication. Valid values: `enable`, `disable`.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * FortiGate interface(s) with NTP server mode enabled. Devices on your network can contact these interfaces for NTP services. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.system.NtpInterface>[]>;
    /**
     * Key for authentication.
     */
    key?: pulumi.Input<string>;
    /**
     * Key ID for authentication.
     */
    keyId?: pulumi.Input<number>;
    /**
     * Key type for authentication. On FortiOS versions 6.2.4-7.4.3: MD5, SHA1. On FortiOS versions >= 7.4.4: MD5, SHA1, SHA256.
     */
    keyType?: pulumi.Input<string>;
    /**
     * Configure the FortiGate to connect to any available third-party NTP server. The structure of `ntpserver` block is documented below.
     */
    ntpservers?: pulumi.Input<pulumi.Input<inputs.system.NtpNtpserver>[]>;
    /**
     * Enable/disable setting the FortiGate system time by synchronizing with an NTP Server. Valid values: `enable`, `disable`.
     */
    ntpsync?: pulumi.Input<string>;
    /**
     * Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server. Valid values: `enable`, `disable`.
     */
    serverMode?: pulumi.Input<string>;
    /**
     * Source IP address for communication to the NTP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Source IPv6 address for communication to the NTP server.
     */
    sourceIp6?: pulumi.Input<string>;
    /**
     * NTP synchronization interval (1 - 1440 min).
     */
    syncinterval?: pulumi.Input<number>;
    /**
     * Use the FortiGuard NTP server or any other available NTP Server. Valid values: `fortiguard`, `custom`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ntp resource.
 */
export interface NtpArgs {
    /**
     * Enable/disable authentication. Valid values: `enable`, `disable`.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * FortiGate interface(s) with NTP server mode enabled. Devices on your network can contact these interfaces for NTP services. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.system.NtpInterface>[]>;
    /**
     * Key for authentication.
     */
    key?: pulumi.Input<string>;
    /**
     * Key ID for authentication.
     */
    keyId?: pulumi.Input<number>;
    /**
     * Key type for authentication. On FortiOS versions 6.2.4-7.4.3: MD5, SHA1. On FortiOS versions >= 7.4.4: MD5, SHA1, SHA256.
     */
    keyType?: pulumi.Input<string>;
    /**
     * Configure the FortiGate to connect to any available third-party NTP server. The structure of `ntpserver` block is documented below.
     */
    ntpservers?: pulumi.Input<pulumi.Input<inputs.system.NtpNtpserver>[]>;
    /**
     * Enable/disable setting the FortiGate system time by synchronizing with an NTP Server. Valid values: `enable`, `disable`.
     */
    ntpsync?: pulumi.Input<string>;
    /**
     * Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server. Valid values: `enable`, `disable`.
     */
    serverMode?: pulumi.Input<string>;
    /**
     * Source IP address for communication to the NTP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Source IPv6 address for communication to the NTP server.
     */
    sourceIp6?: pulumi.Input<string>;
    /**
     * NTP synchronization interval (1 - 1440 min).
     */
    syncinterval?: pulumi.Input<number>;
    /**
     * Use the FortiGuard NTP server or any other available NTP Server. Valid values: `fortiguard`, `custom`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
