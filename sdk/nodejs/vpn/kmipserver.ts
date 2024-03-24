// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * KMIP server entry configuration. Applies to FortiOS Version `>= 7.4.0`.
 *
 * ## Import
 *
 * Vpn KmipServer can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:vpn/kmipserver:Kmipserver labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:vpn/kmipserver:Kmipserver labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Kmipserver extends pulumi.CustomResource {
    /**
     * Get an existing Kmipserver resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KmipserverState, opts?: pulumi.CustomResourceOptions): Kmipserver {
        return new Kmipserver(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:vpn/kmipserver:Kmipserver';

    /**
     * Returns true if the given object is an instance of Kmipserver.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Kmipserver {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Kmipserver.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    /**
     * KMIP server entry name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password to use for connectivity to the KMIP server.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable KMIP server identity check (verify server FQDN/IP address against the server certificate). Valid values: `enable`, `disable`.
     */
    public readonly serverIdentityCheck!: pulumi.Output<string>;
    /**
     * KMIP server list. The structure of `serverList` block is documented below.
     */
    public readonly serverLists!: pulumi.Output<outputs.vpn.KmipserverServerList[] | undefined>;
    /**
     * FortiGate IP address to be used for communication with the KMIP server.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
     */
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    /**
     * User name to use for connectivity to the KMIP server.
     */
    public readonly username!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Kmipserver resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: KmipserverArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KmipserverArgs | KmipserverState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KmipserverState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["serverIdentityCheck"] = state ? state.serverIdentityCheck : undefined;
            resourceInputs["serverLists"] = state ? state.serverLists : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as KmipserverArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["serverIdentityCheck"] = args ? args.serverIdentityCheck : undefined;
            resourceInputs["serverLists"] = args ? args.serverLists : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Kmipserver.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Kmipserver resources.
 */
export interface KmipserverState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * KMIP server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Password to use for connectivity to the KMIP server.
     */
    password?: pulumi.Input<string>;
    /**
     * Enable/disable KMIP server identity check (verify server FQDN/IP address against the server certificate). Valid values: `enable`, `disable`.
     */
    serverIdentityCheck?: pulumi.Input<string>;
    /**
     * KMIP server list. The structure of `serverList` block is documented below.
     */
    serverLists?: pulumi.Input<pulumi.Input<inputs.vpn.KmipserverServerList>[]>;
    /**
     * FortiGate IP address to be used for communication with the KMIP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * User name to use for connectivity to the KMIP server.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Kmipserver resource.
 */
export interface KmipserverArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * KMIP server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Password to use for connectivity to the KMIP server.
     */
    password?: pulumi.Input<string>;
    /**
     * Enable/disable KMIP server identity check (verify server FQDN/IP address against the server certificate). Valid values: `enable`, `disable`.
     */
    serverIdentityCheck?: pulumi.Input<string>;
    /**
     * KMIP server list. The structure of `serverList` block is documented below.
     */
    serverLists?: pulumi.Input<pulumi.Input<inputs.vpn.KmipserverServerList>[]>;
    /**
     * FortiGate IP address to be used for communication with the KMIP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * User name to use for connectivity to the KMIP server.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}