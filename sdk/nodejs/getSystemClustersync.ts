// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios system clustersync
 */
export function getSystemClustersync(args: GetSystemClustersyncArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemClustersyncResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemClustersync:getSystemClustersync", {
        "syncId": args.syncId,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemClustersync.
 */
export interface GetSystemClustersyncArgs {
    /**
     * Specify the syncId of the desired system clustersync.
     */
    syncId: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getSystemClustersync.
 */
export interface GetSystemClustersyncResult {
    /**
     * List of interfaces to be turned down before session synchronization is complete. The structure of `downIntfsBeforeSessSync` block is documented below.
     */
    readonly downIntfsBeforeSessSyncs: outputs.GetSystemClustersyncDownIntfsBeforeSessSync[];
    /**
     * Heartbeat interval (1 - 10 sec).
     */
    readonly hbInterval: number;
    /**
     * Lost heartbeat threshold (1 - 10).
     */
    readonly hbLostThreshold: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * IKE heartbeat interval (1 - 60 secs).
     */
    readonly ikeHeartbeatInterval: number;
    /**
     * Enable/disable IKE HA monitor.
     */
    readonly ikeMonitor: string;
    /**
     * IKE HA monitor interval (10 - 300 secs).
     */
    readonly ikeMonitorInterval: number;
    /**
     * Enable/disable IPsec tunnel synchronization.
     */
    readonly ipsecTunnelSync: string;
    /**
     * IP address of the interface on the peer unit that is used for the session synchronization link.
     */
    readonly peerip: string;
    /**
     * VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
     */
    readonly peervd: string;
    /**
     * Enable/disable IKE route announcement on the backup unit.
     */
    readonly secondaryAddIpsecRoutes: string;
    /**
     * Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `sessionSyncFilter` block is documented below.
     */
    readonly sessionSyncFilters: outputs.GetSystemClustersyncSessionSyncFilter[];
    /**
     * Enable/disable IKE route announcement on the backup unit.
     */
    readonly slaveAddIkeRoutes: string;
    /**
     * Sync ID.
     */
    readonly syncId: number;
    /**
     * Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
     */
    readonly syncvds: outputs.GetSystemClustersyncSyncvd[];
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios system clustersync
 */
export function getSystemClustersyncOutput(args: GetSystemClustersyncOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemClustersyncResult> {
    return pulumi.output(args).apply((a: any) => getSystemClustersync(a, opts))
}

/**
 * A collection of arguments for invoking getSystemClustersync.
 */
export interface GetSystemClustersyncOutputArgs {
    /**
     * Specify the syncId of the desired system clustersync.
     */
    syncId: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
