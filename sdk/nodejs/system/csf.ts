// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Csf("trname", {
 *     configurationSync: "default",
 *     groupPassword: "tmp",
 *     managementIp: "0.0.0.0",
 *     managementPort: 33,
 *     status: "disable",
 *     upstreamIp: "0.0.0.0",
 *     upstreamPort: 8013,
 * });
 * ```
 *
 * ## Import
 *
 * System Csf can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/csf:Csf labelname SystemCsf
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/csf:Csf labelname SystemCsf
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Csf extends pulumi.CustomResource {
    /**
     * Get an existing Csf resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CsfState, opts?: pulumi.CustomResourceOptions): Csf {
        return new Csf(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/csf:Csf';

    /**
     * Returns true if the given object is an instance of Csf.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Csf {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Csf.__pulumiType;
    }

    /**
     * Accept connections with unknown certificates and ask admin for approval. Valid values: `disable`, `enable`.
     */
    public readonly acceptAuthByCert!: pulumi.Output<string>;
    /**
     * Authorization request type. Valid values: `serial`, `certificate`.
     */
    public readonly authorizationRequestType!: pulumi.Output<string>;
    /**
     * Certificate.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * Configuration sync mode. Valid values: `default`, `local`.
     */
    public readonly configurationSync!: pulumi.Output<string>;
    /**
     * Enable/disable downstream device access to this device's configuration and data. Valid values: `enable`, `disable`.
     */
    public readonly downstreamAccess!: pulumi.Output<string>;
    /**
     * Default access profile for requests from downstream devices.
     */
    public readonly downstreamAccprofile!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Fabric connector configuration. The structure of `fabricConnector` block is documented below.
     */
    public readonly fabricConnectors!: pulumi.Output<outputs.system.CsfFabricConnector[] | undefined>;
    /**
     * Fabric device configuration. The structure of `fabricDevice` block is documented below.
     */
    public readonly fabricDevices!: pulumi.Output<outputs.system.CsfFabricDevice[] | undefined>;
    /**
     * Fabric CMDB Object Unification Valid values: `default`, `local`.
     */
    public readonly fabricObjectUnification!: pulumi.Output<string>;
    /**
     * Number of worker processes for Security Fabric daemon.
     */
    public readonly fabricWorkers!: pulumi.Output<number>;
    /**
     * Enable/disable Security Fabric daemon file management. Valid values: `enable`, `disable`.
     */
    public readonly fileMgmt!: pulumi.Output<string>;
    /**
     * Maximum amount of memory that can be used by the daemon files (in bytes).
     */
    public readonly fileQuota!: pulumi.Output<number>;
    /**
     * Warn when the set percentage of quota has been used.
     */
    public readonly fileQuotaWarning!: pulumi.Output<number>;
    /**
     * Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
     */
    public readonly fixedKey!: pulumi.Output<string | undefined>;
    /**
     * Fabric FortiCloud account unification. Valid values: `enable`, `disable`.
     */
    public readonly forticloudAccountEnforcement!: pulumi.Output<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
     */
    public readonly groupPassword!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable broadcast of discovery messages for log unification. Valid values: `disable`, `enable`.
     */
    public readonly logUnification!: pulumi.Output<string>;
    /**
     * Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
     */
    public readonly managementIp!: pulumi.Output<string>;
    /**
     * Overriding port for management connection (Overrides admin port).
     */
    public readonly managementPort!: pulumi.Output<number>;
    /**
     * SAML setting configuration synchronization. Valid values: `default`, `local`.
     */
    public readonly samlConfigurationSync!: pulumi.Output<string>;
    /**
     * Source IP address for communication with the upstream FortiGate.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Enable/disable Security Fabric. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Pre-authorized and blocked security fabric nodes. The structure of `trustedList` block is documented below.
     */
    public readonly trustedLists!: pulumi.Output<outputs.system.CsfTrustedList[] | undefined>;
    /**
     * Unique ID of the current CSF node
     */
    public readonly uid!: pulumi.Output<string>;
    /**
     * IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
     */
    public readonly upstream!: pulumi.Output<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly upstreamInterface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly upstreamInterfaceSelectMethod!: pulumi.Output<string>;
    /**
     * IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
     */
    public readonly upstreamIp!: pulumi.Output<string>;
    /**
     * The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
     */
    public readonly upstreamPort!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Csf resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CsfArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CsfArgs | CsfState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CsfState | undefined;
            resourceInputs["acceptAuthByCert"] = state ? state.acceptAuthByCert : undefined;
            resourceInputs["authorizationRequestType"] = state ? state.authorizationRequestType : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["configurationSync"] = state ? state.configurationSync : undefined;
            resourceInputs["downstreamAccess"] = state ? state.downstreamAccess : undefined;
            resourceInputs["downstreamAccprofile"] = state ? state.downstreamAccprofile : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fabricConnectors"] = state ? state.fabricConnectors : undefined;
            resourceInputs["fabricDevices"] = state ? state.fabricDevices : undefined;
            resourceInputs["fabricObjectUnification"] = state ? state.fabricObjectUnification : undefined;
            resourceInputs["fabricWorkers"] = state ? state.fabricWorkers : undefined;
            resourceInputs["fileMgmt"] = state ? state.fileMgmt : undefined;
            resourceInputs["fileQuota"] = state ? state.fileQuota : undefined;
            resourceInputs["fileQuotaWarning"] = state ? state.fileQuotaWarning : undefined;
            resourceInputs["fixedKey"] = state ? state.fixedKey : undefined;
            resourceInputs["forticloudAccountEnforcement"] = state ? state.forticloudAccountEnforcement : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["groupPassword"] = state ? state.groupPassword : undefined;
            resourceInputs["logUnification"] = state ? state.logUnification : undefined;
            resourceInputs["managementIp"] = state ? state.managementIp : undefined;
            resourceInputs["managementPort"] = state ? state.managementPort : undefined;
            resourceInputs["samlConfigurationSync"] = state ? state.samlConfigurationSync : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["trustedLists"] = state ? state.trustedLists : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["upstream"] = state ? state.upstream : undefined;
            resourceInputs["upstreamInterface"] = state ? state.upstreamInterface : undefined;
            resourceInputs["upstreamInterfaceSelectMethod"] = state ? state.upstreamInterfaceSelectMethod : undefined;
            resourceInputs["upstreamIp"] = state ? state.upstreamIp : undefined;
            resourceInputs["upstreamPort"] = state ? state.upstreamPort : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as CsfArgs | undefined;
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["acceptAuthByCert"] = args ? args.acceptAuthByCert : undefined;
            resourceInputs["authorizationRequestType"] = args ? args.authorizationRequestType : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["configurationSync"] = args ? args.configurationSync : undefined;
            resourceInputs["downstreamAccess"] = args ? args.downstreamAccess : undefined;
            resourceInputs["downstreamAccprofile"] = args ? args.downstreamAccprofile : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fabricConnectors"] = args ? args.fabricConnectors : undefined;
            resourceInputs["fabricDevices"] = args ? args.fabricDevices : undefined;
            resourceInputs["fabricObjectUnification"] = args ? args.fabricObjectUnification : undefined;
            resourceInputs["fabricWorkers"] = args ? args.fabricWorkers : undefined;
            resourceInputs["fileMgmt"] = args ? args.fileMgmt : undefined;
            resourceInputs["fileQuota"] = args ? args.fileQuota : undefined;
            resourceInputs["fileQuotaWarning"] = args ? args.fileQuotaWarning : undefined;
            resourceInputs["fixedKey"] = args?.fixedKey ? pulumi.secret(args.fixedKey) : undefined;
            resourceInputs["forticloudAccountEnforcement"] = args ? args.forticloudAccountEnforcement : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["groupPassword"] = args?.groupPassword ? pulumi.secret(args.groupPassword) : undefined;
            resourceInputs["logUnification"] = args ? args.logUnification : undefined;
            resourceInputs["managementIp"] = args ? args.managementIp : undefined;
            resourceInputs["managementPort"] = args ? args.managementPort : undefined;
            resourceInputs["samlConfigurationSync"] = args ? args.samlConfigurationSync : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["trustedLists"] = args ? args.trustedLists : undefined;
            resourceInputs["uid"] = args ? args.uid : undefined;
            resourceInputs["upstream"] = args ? args.upstream : undefined;
            resourceInputs["upstreamInterface"] = args ? args.upstreamInterface : undefined;
            resourceInputs["upstreamInterfaceSelectMethod"] = args ? args.upstreamInterfaceSelectMethod : undefined;
            resourceInputs["upstreamIp"] = args ? args.upstreamIp : undefined;
            resourceInputs["upstreamPort"] = args ? args.upstreamPort : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["fixedKey", "groupPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Csf.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Csf resources.
 */
export interface CsfState {
    /**
     * Accept connections with unknown certificates and ask admin for approval. Valid values: `disable`, `enable`.
     */
    acceptAuthByCert?: pulumi.Input<string>;
    /**
     * Authorization request type. Valid values: `serial`, `certificate`.
     */
    authorizationRequestType?: pulumi.Input<string>;
    /**
     * Certificate.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Configuration sync mode. Valid values: `default`, `local`.
     */
    configurationSync?: pulumi.Input<string>;
    /**
     * Enable/disable downstream device access to this device's configuration and data. Valid values: `enable`, `disable`.
     */
    downstreamAccess?: pulumi.Input<string>;
    /**
     * Default access profile for requests from downstream devices.
     */
    downstreamAccprofile?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Fabric connector configuration. The structure of `fabricConnector` block is documented below.
     */
    fabricConnectors?: pulumi.Input<pulumi.Input<inputs.system.CsfFabricConnector>[]>;
    /**
     * Fabric device configuration. The structure of `fabricDevice` block is documented below.
     */
    fabricDevices?: pulumi.Input<pulumi.Input<inputs.system.CsfFabricDevice>[]>;
    /**
     * Fabric CMDB Object Unification Valid values: `default`, `local`.
     */
    fabricObjectUnification?: pulumi.Input<string>;
    /**
     * Number of worker processes for Security Fabric daemon.
     */
    fabricWorkers?: pulumi.Input<number>;
    /**
     * Enable/disable Security Fabric daemon file management. Valid values: `enable`, `disable`.
     */
    fileMgmt?: pulumi.Input<string>;
    /**
     * Maximum amount of memory that can be used by the daemon files (in bytes).
     */
    fileQuota?: pulumi.Input<number>;
    /**
     * Warn when the set percentage of quota has been used.
     */
    fileQuotaWarning?: pulumi.Input<number>;
    /**
     * Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
     */
    fixedKey?: pulumi.Input<string>;
    /**
     * Fabric FortiCloud account unification. Valid values: `enable`, `disable`.
     */
    forticloudAccountEnforcement?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
     */
    groupName?: pulumi.Input<string>;
    /**
     * Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
     */
    groupPassword?: pulumi.Input<string>;
    /**
     * Enable/disable broadcast of discovery messages for log unification. Valid values: `disable`, `enable`.
     */
    logUnification?: pulumi.Input<string>;
    /**
     * Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
     */
    managementIp?: pulumi.Input<string>;
    /**
     * Overriding port for management connection (Overrides admin port).
     */
    managementPort?: pulumi.Input<number>;
    /**
     * SAML setting configuration synchronization. Valid values: `default`, `local`.
     */
    samlConfigurationSync?: pulumi.Input<string>;
    /**
     * Source IP address for communication with the upstream FortiGate.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable Security Fabric. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Pre-authorized and blocked security fabric nodes. The structure of `trustedList` block is documented below.
     */
    trustedLists?: pulumi.Input<pulumi.Input<inputs.system.CsfTrustedList>[]>;
    /**
     * Unique ID of the current CSF node
     */
    uid?: pulumi.Input<string>;
    /**
     * IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
     */
    upstream?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    upstreamInterface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    upstreamInterfaceSelectMethod?: pulumi.Input<string>;
    /**
     * IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
     */
    upstreamIp?: pulumi.Input<string>;
    /**
     * The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
     */
    upstreamPort?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Csf resource.
 */
export interface CsfArgs {
    /**
     * Accept connections with unknown certificates and ask admin for approval. Valid values: `disable`, `enable`.
     */
    acceptAuthByCert?: pulumi.Input<string>;
    /**
     * Authorization request type. Valid values: `serial`, `certificate`.
     */
    authorizationRequestType?: pulumi.Input<string>;
    /**
     * Certificate.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Configuration sync mode. Valid values: `default`, `local`.
     */
    configurationSync?: pulumi.Input<string>;
    /**
     * Enable/disable downstream device access to this device's configuration and data. Valid values: `enable`, `disable`.
     */
    downstreamAccess?: pulumi.Input<string>;
    /**
     * Default access profile for requests from downstream devices.
     */
    downstreamAccprofile?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Fabric connector configuration. The structure of `fabricConnector` block is documented below.
     */
    fabricConnectors?: pulumi.Input<pulumi.Input<inputs.system.CsfFabricConnector>[]>;
    /**
     * Fabric device configuration. The structure of `fabricDevice` block is documented below.
     */
    fabricDevices?: pulumi.Input<pulumi.Input<inputs.system.CsfFabricDevice>[]>;
    /**
     * Fabric CMDB Object Unification Valid values: `default`, `local`.
     */
    fabricObjectUnification?: pulumi.Input<string>;
    /**
     * Number of worker processes for Security Fabric daemon.
     */
    fabricWorkers?: pulumi.Input<number>;
    /**
     * Enable/disable Security Fabric daemon file management. Valid values: `enable`, `disable`.
     */
    fileMgmt?: pulumi.Input<string>;
    /**
     * Maximum amount of memory that can be used by the daemon files (in bytes).
     */
    fileQuota?: pulumi.Input<number>;
    /**
     * Warn when the set percentage of quota has been used.
     */
    fileQuotaWarning?: pulumi.Input<number>;
    /**
     * Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
     */
    fixedKey?: pulumi.Input<string>;
    /**
     * Fabric FortiCloud account unification. Valid values: `enable`, `disable`.
     */
    forticloudAccountEnforcement?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
     */
    groupName?: pulumi.Input<string>;
    /**
     * Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
     */
    groupPassword?: pulumi.Input<string>;
    /**
     * Enable/disable broadcast of discovery messages for log unification. Valid values: `disable`, `enable`.
     */
    logUnification?: pulumi.Input<string>;
    /**
     * Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
     */
    managementIp?: pulumi.Input<string>;
    /**
     * Overriding port for management connection (Overrides admin port).
     */
    managementPort?: pulumi.Input<number>;
    /**
     * SAML setting configuration synchronization. Valid values: `default`, `local`.
     */
    samlConfigurationSync?: pulumi.Input<string>;
    /**
     * Source IP address for communication with the upstream FortiGate.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable Security Fabric. Valid values: `enable`, `disable`.
     */
    status: pulumi.Input<string>;
    /**
     * Pre-authorized and blocked security fabric nodes. The structure of `trustedList` block is documented below.
     */
    trustedLists?: pulumi.Input<pulumi.Input<inputs.system.CsfTrustedList>[]>;
    /**
     * Unique ID of the current CSF node
     */
    uid?: pulumi.Input<string>;
    /**
     * IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
     */
    upstream?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    upstreamInterface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    upstreamInterfaceSelectMethod?: pulumi.Input<string>;
    /**
     * IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
     */
    upstreamIp?: pulumi.Input<string>;
    /**
     * The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
     */
    upstreamPort?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
