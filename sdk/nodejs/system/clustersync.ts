// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. Applies to FortiOS Version `<= 7.2.0`.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Clustersync("trname", {
 *     hbInterval: 3,
 *     hbLostThreshold: 3,
 *     peerip: "1.1.1.1",
 *     peervd: "root",
 *     slaveAddIkeRoutes: "enable",
 *     syncId: 1,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * System ClusterSync can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/clustersync:Clustersync labelname {{sync_id}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/clustersync:Clustersync labelname {{sync_id}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Clustersync extends pulumi.CustomResource {
    /**
     * Get an existing Clustersync resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClustersyncState, opts?: pulumi.CustomResourceOptions): Clustersync {
        return new Clustersync(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/clustersync:Clustersync';

    /**
     * Returns true if the given object is an instance of Clustersync.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Clustersync {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Clustersync.__pulumiType;
    }

    /**
     * List of interfaces to be turned down before session synchronization is complete. The structure of `downIntfsBeforeSessSync` block is documented below.
     */
    public readonly downIntfsBeforeSessSyncs!: pulumi.Output<outputs.system.ClustersyncDownIntfsBeforeSessSync[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Heartbeat interval (1 - 10 sec).
     */
    public readonly hbInterval!: pulumi.Output<number>;
    /**
     * Lost heartbeat threshold (1 - 10).
     */
    public readonly hbLostThreshold!: pulumi.Output<number>;
    /**
     * IKE heartbeat interval (1 - 60 secs).
     */
    public readonly ikeHeartbeatInterval!: pulumi.Output<number>;
    /**
     * Enable/disable IKE HA monitor. Valid values: `enable`, `disable`.
     */
    public readonly ikeMonitor!: pulumi.Output<string>;
    /**
     * IKE HA monitor interval (10 - 300 secs).
     */
    public readonly ikeMonitorInterval!: pulumi.Output<number>;
    /**
     * Enable/disable IPsec tunnel synchronization. Valid values: `enable`, `disable`.
     */
    public readonly ipsecTunnelSync!: pulumi.Output<string>;
    /**
     * IP address of the interface on the peer unit that is used for the session synchronization link.
     */
    public readonly peerip!: pulumi.Output<string>;
    /**
     * VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
     */
    public readonly peervd!: pulumi.Output<string>;
    /**
     * Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
     */
    public readonly secondaryAddIpsecRoutes!: pulumi.Output<string>;
    /**
     * Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `sessionSyncFilter` block is documented below.
     */
    public readonly sessionSyncFilter!: pulumi.Output<outputs.system.ClustersyncSessionSyncFilter>;
    /**
     * Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
     */
    public readonly slaveAddIkeRoutes!: pulumi.Output<string>;
    /**
     * Sync ID.
     */
    public readonly syncId!: pulumi.Output<number>;
    /**
     * Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
     */
    public readonly syncvds!: pulumi.Output<outputs.system.ClustersyncSyncvd[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Clustersync resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClustersyncArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClustersyncArgs | ClustersyncState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClustersyncState | undefined;
            resourceInputs["downIntfsBeforeSessSyncs"] = state ? state.downIntfsBeforeSessSyncs : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["hbInterval"] = state ? state.hbInterval : undefined;
            resourceInputs["hbLostThreshold"] = state ? state.hbLostThreshold : undefined;
            resourceInputs["ikeHeartbeatInterval"] = state ? state.ikeHeartbeatInterval : undefined;
            resourceInputs["ikeMonitor"] = state ? state.ikeMonitor : undefined;
            resourceInputs["ikeMonitorInterval"] = state ? state.ikeMonitorInterval : undefined;
            resourceInputs["ipsecTunnelSync"] = state ? state.ipsecTunnelSync : undefined;
            resourceInputs["peerip"] = state ? state.peerip : undefined;
            resourceInputs["peervd"] = state ? state.peervd : undefined;
            resourceInputs["secondaryAddIpsecRoutes"] = state ? state.secondaryAddIpsecRoutes : undefined;
            resourceInputs["sessionSyncFilter"] = state ? state.sessionSyncFilter : undefined;
            resourceInputs["slaveAddIkeRoutes"] = state ? state.slaveAddIkeRoutes : undefined;
            resourceInputs["syncId"] = state ? state.syncId : undefined;
            resourceInputs["syncvds"] = state ? state.syncvds : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ClustersyncArgs | undefined;
            resourceInputs["downIntfsBeforeSessSyncs"] = args ? args.downIntfsBeforeSessSyncs : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["hbInterval"] = args ? args.hbInterval : undefined;
            resourceInputs["hbLostThreshold"] = args ? args.hbLostThreshold : undefined;
            resourceInputs["ikeHeartbeatInterval"] = args ? args.ikeHeartbeatInterval : undefined;
            resourceInputs["ikeMonitor"] = args ? args.ikeMonitor : undefined;
            resourceInputs["ikeMonitorInterval"] = args ? args.ikeMonitorInterval : undefined;
            resourceInputs["ipsecTunnelSync"] = args ? args.ipsecTunnelSync : undefined;
            resourceInputs["peerip"] = args ? args.peerip : undefined;
            resourceInputs["peervd"] = args ? args.peervd : undefined;
            resourceInputs["secondaryAddIpsecRoutes"] = args ? args.secondaryAddIpsecRoutes : undefined;
            resourceInputs["sessionSyncFilter"] = args ? args.sessionSyncFilter : undefined;
            resourceInputs["slaveAddIkeRoutes"] = args ? args.slaveAddIkeRoutes : undefined;
            resourceInputs["syncId"] = args ? args.syncId : undefined;
            resourceInputs["syncvds"] = args ? args.syncvds : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Clustersync.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Clustersync resources.
 */
export interface ClustersyncState {
    /**
     * List of interfaces to be turned down before session synchronization is complete. The structure of `downIntfsBeforeSessSync` block is documented below.
     */
    downIntfsBeforeSessSyncs?: pulumi.Input<pulumi.Input<inputs.system.ClustersyncDownIntfsBeforeSessSync>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Heartbeat interval (1 - 10 sec).
     */
    hbInterval?: pulumi.Input<number>;
    /**
     * Lost heartbeat threshold (1 - 10).
     */
    hbLostThreshold?: pulumi.Input<number>;
    /**
     * IKE heartbeat interval (1 - 60 secs).
     */
    ikeHeartbeatInterval?: pulumi.Input<number>;
    /**
     * Enable/disable IKE HA monitor. Valid values: `enable`, `disable`.
     */
    ikeMonitor?: pulumi.Input<string>;
    /**
     * IKE HA monitor interval (10 - 300 secs).
     */
    ikeMonitorInterval?: pulumi.Input<number>;
    /**
     * Enable/disable IPsec tunnel synchronization. Valid values: `enable`, `disable`.
     */
    ipsecTunnelSync?: pulumi.Input<string>;
    /**
     * IP address of the interface on the peer unit that is used for the session synchronization link.
     */
    peerip?: pulumi.Input<string>;
    /**
     * VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
     */
    peervd?: pulumi.Input<string>;
    /**
     * Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
     */
    secondaryAddIpsecRoutes?: pulumi.Input<string>;
    /**
     * Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `sessionSyncFilter` block is documented below.
     */
    sessionSyncFilter?: pulumi.Input<inputs.system.ClustersyncSessionSyncFilter>;
    /**
     * Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
     */
    slaveAddIkeRoutes?: pulumi.Input<string>;
    /**
     * Sync ID.
     */
    syncId?: pulumi.Input<number>;
    /**
     * Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
     */
    syncvds?: pulumi.Input<pulumi.Input<inputs.system.ClustersyncSyncvd>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Clustersync resource.
 */
export interface ClustersyncArgs {
    /**
     * List of interfaces to be turned down before session synchronization is complete. The structure of `downIntfsBeforeSessSync` block is documented below.
     */
    downIntfsBeforeSessSyncs?: pulumi.Input<pulumi.Input<inputs.system.ClustersyncDownIntfsBeforeSessSync>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Heartbeat interval (1 - 10 sec).
     */
    hbInterval?: pulumi.Input<number>;
    /**
     * Lost heartbeat threshold (1 - 10).
     */
    hbLostThreshold?: pulumi.Input<number>;
    /**
     * IKE heartbeat interval (1 - 60 secs).
     */
    ikeHeartbeatInterval?: pulumi.Input<number>;
    /**
     * Enable/disable IKE HA monitor. Valid values: `enable`, `disable`.
     */
    ikeMonitor?: pulumi.Input<string>;
    /**
     * IKE HA monitor interval (10 - 300 secs).
     */
    ikeMonitorInterval?: pulumi.Input<number>;
    /**
     * Enable/disable IPsec tunnel synchronization. Valid values: `enable`, `disable`.
     */
    ipsecTunnelSync?: pulumi.Input<string>;
    /**
     * IP address of the interface on the peer unit that is used for the session synchronization link.
     */
    peerip?: pulumi.Input<string>;
    /**
     * VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
     */
    peervd?: pulumi.Input<string>;
    /**
     * Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
     */
    secondaryAddIpsecRoutes?: pulumi.Input<string>;
    /**
     * Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `sessionSyncFilter` block is documented below.
     */
    sessionSyncFilter?: pulumi.Input<inputs.system.ClustersyncSessionSyncFilter>;
    /**
     * Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
     */
    slaveAddIkeRoutes?: pulumi.Input<string>;
    /**
     * Sync ID.
     */
    syncId?: pulumi.Input<number>;
    /**
     * Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
     */
    syncvds?: pulumi.Input<pulumi.Input<inputs.system.ClustersyncSyncvd>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
