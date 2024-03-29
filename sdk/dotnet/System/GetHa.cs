// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    public static class GetHa
    {
        /// <summary>
        /// Use this data source to get information on fortios system ha
        /// </summary>
        public static Task<GetHaResult> InvokeAsync(GetHaArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHaResult>("fortios:system/getHa:getHa", args ?? new GetHaArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system ha
        /// </summary>
        public static Output<GetHaResult> Invoke(GetHaInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHaResult>("fortios:system/getHa:getHa", args ?? new GetHaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHaArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetHaArgs()
        {
        }
        public static new GetHaArgs Empty => new GetHaArgs();
    }

    public sealed class GetHaInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetHaInvokeArgs()
        {
        }
        public static new GetHaInvokeArgs Empty => new GetHaInvokeArgs();
    }


    [OutputType]
    public sealed class GetHaResult
    {
        /// <summary>
        /// Number of gratuitous ARPs (1 - 60). Lower to reduce traffic. Higher to reduce failover time.
        /// </summary>
        public readonly int Arps;
        /// <summary>
        /// Time between gratuitous ARPs  (1 - 20 sec). Lower to reduce failover time. Higher to reduce traffic.
        /// </summary>
        public readonly int ArpsInterval;
        /// <summary>
        /// Enable/disable heartbeat message authentication.
        /// </summary>
        public readonly string Authentication;
        /// <summary>
        /// Dynamic weighted load balancing CPU usage weight and high and low thresholds.
        /// </summary>
        public readonly string CpuThreshold;
        /// <summary>
        /// Enable/disable heartbeat message encryption.
        /// </summary>
        public readonly string Encryption;
        /// <summary>
        /// HA EVPN FDB TTL on primary box (5 - 3600 sec).
        /// </summary>
        public readonly int EvpnTtl;
        /// <summary>
        /// Time to wait before failover (0 - 300 sec, default = 0), to avoid flip.
        /// </summary>
        public readonly int FailoverHoldTime;
        /// <summary>
        /// Dynamic weighted load balancing weight and high and low number of FTP proxy sessions.
        /// </summary>
        public readonly string FtpProxyThreshold;
        /// <summary>
        /// Enable/disable gratuitous ARPs. Disable if link-failed-signal enabled.
        /// </summary>
        public readonly string GratuitousArps;
        /// <summary>
        /// Cluster group ID  (0 - 255). Must be the same for all members.
        /// </summary>
        public readonly int GroupId;
        /// <summary>
        /// Cluster group name. Must be the same for all members.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// Enable/disable using ha-mgmt interface for syslog, SNMP, remote authentication (RADIUS), FortiAnalyzer, and FortiSandbox.
        /// </summary>
        public readonly string HaDirect;
        /// <summary>
        /// HA heartbeat packet Ethertype (4-digit hex).
        /// </summary>
        public readonly string HaEthType;
        /// <summary>
        /// Reserve interfaces to manage individual cluster units. The structure of `ha_mgmt_interfaces` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHaHaMgmtInterfaceResult> HaMgmtInterfaces;
        /// <summary>
        /// Enable to reserve interfaces to manage individual cluster units.
        /// </summary>
        public readonly string HaMgmtStatus;
        /// <summary>
        /// Normally you would only reduce this value for failover testing.
        /// </summary>
        public readonly int HaUptimeDiffMargin;
        /// <summary>
        /// Time between sending heartbeat packets (1 - 20 (100*ms)). Increase to reduce false positives.
        /// </summary>
        public readonly int HbInterval;
        /// <summary>
        /// Number of milliseconds for each heartbeat interval: 100ms or 10ms.
        /// </summary>
        public readonly string HbIntervalInMilliseconds;
        /// <summary>
        /// Number of lost heartbeats to signal a failure (1 - 60). Increase to reduce false positives.
        /// </summary>
        public readonly int HbLostThreshold;
        /// <summary>
        /// Heartbeat interfaces. Must be the same for all members.
        /// </summary>
        public readonly string Hbdev;
        /// <summary>
        /// Transparent mode HA heartbeat packet Ethertype (4-digit hex).
        /// </summary>
        public readonly string HcEthType;
        /// <summary>
        /// Time to wait before changing from hello to work state (5 - 300 sec).
        /// </summary>
        public readonly int HelloHolddown;
        /// <summary>
        /// Dynamic weighted load balancing weight and high and low number of HTTP proxy sessions.
        /// </summary>
        public readonly string HttpProxyThreshold;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Dynamic weighted load balancing weight and high and low number of IMAP proxy sessions.
        /// </summary>
        public readonly string ImapProxyThreshold;
        /// <summary>
        /// Enable/disable synchronization of sessions among HA clusters.
        /// </summary>
        public readonly string InterClusterSessionSync;
        /// <summary>
        /// IPsec phase2 proposal.
        /// </summary>
        public readonly string IpsecPhase2Proposal;
        /// <summary>
        /// key
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Telnet session HA heartbeat packet Ethertype (4-digit hex).
        /// </summary>
        public readonly string L2epEthType;
        /// <summary>
        /// Enable to shut down all interfaces for 1 sec after a failover. Use if gratuitous ARPs do not update network.
        /// </summary>
        public readonly string LinkFailedSignal;
        /// <summary>
        /// Enable to load balance TCP sessions. Disable to load balance proxy sessions only.
        /// </summary>
        public readonly string LoadBalanceAll;
        /// <summary>
        /// Enable/disable usage of the logical serial number.
        /// </summary>
        public readonly string LogicalSn;
        /// <summary>
        /// Enable/disable memory based failover.
        /// </summary>
        public readonly string MemoryBasedFailover;
        /// <summary>
        /// Enable/disable memory compatible mode.
        /// </summary>
        public readonly string MemoryCompatibleMode;
        /// <summary>
        /// Time to wait between subsequent memory based failovers in minutes (6 - 2147483647, default = 6).
        /// </summary>
        public readonly int MemoryFailoverFlipTimeout;
        /// <summary>
        /// Duration of high memory usage before memory based failover is triggered in seconds (1 - 300, default = 60).
        /// </summary>
        public readonly int MemoryFailoverMonitorPeriod;
        /// <summary>
        /// Rate at which memory usage is sampled in order to measure memory usage in seconds (1 - 60, default = 1).
        /// </summary>
        public readonly int MemoryFailoverSampleRate;
        /// <summary>
        /// Memory usage threshold to trigger memory based failover (0 means using conserve mode threshold in system.global).
        /// </summary>
        public readonly int MemoryFailoverThreshold;
        /// <summary>
        /// Dynamic weighted load balancing memory usage weight and high and low thresholds.
        /// </summary>
        public readonly string MemoryThreshold;
        /// <summary>
        /// HA mode. Must be the same for all members. FGSP requires standalone.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// Interfaces to check for port monitoring (or link failure).
        /// </summary>
        public readonly string Monitor;
        /// <summary>
        /// HA multicast TTL on master (5 - 3600 sec).
        /// </summary>
        public readonly int MulticastTtl;
        /// <summary>
        /// Dynamic weighted load balancing weight and high and low number of NNTP proxy sessions.
        /// </summary>
        public readonly string NntpProxyThreshold;
        /// <summary>
        /// Enable and increase the priority of the unit that should always be primary (master).
        /// </summary>
        public readonly string Override;
        /// <summary>
        /// Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
        /// </summary>
        public readonly int OverrideWaitTime;
        /// <summary>
        /// Cluster password. Must be the same for all members.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Remote IP monitoring failover threshold (0 - 50).
        /// </summary>
        public readonly int PingserverFailoverThreshold;
        /// <summary>
        /// Time to wait in minutes before renegotiating after a remote IP monitoring failover.
        /// </summary>
        public readonly int PingserverFlipTimeout;
        /// <summary>
        /// Interfaces to check for remote IP monitoring.
        /// </summary>
        public readonly string PingserverMonitorInterface;
        /// <summary>
        /// Enable to force the cluster to negotiate after a remote IP monitoring failover.
        /// </summary>
        public readonly string PingserverSecondaryForceReset;
        /// <summary>
        /// Enable to force the cluster to negotiate after a remote IP monitoring failover.
        /// </summary>
        public readonly string PingserverSlaveForceReset;
        /// <summary>
        /// Dynamic weighted load balancing weight and high and low number of POP3 proxy sessions.
        /// </summary>
        public readonly string Pop3ProxyThreshold;
        /// <summary>
        /// Increase the priority to select the primary unit (0 - 255).
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// Time to wait between routing table updates to the cluster (0 - 3600 sec).
        /// </summary>
        public readonly int RouteHold;
        /// <summary>
        /// TTL for primary unit routes (5 - 3600 sec). Increase to maintain active routes during failover.
        /// </summary>
        public readonly int RouteTtl;
        /// <summary>
        /// Time to wait before sending new routes to the cluster (0 - 3600 sec).
        /// </summary>
        public readonly int RouteWait;
        /// <summary>
        /// Type of A-A load balancing. Use none if you have external load balancers.
        /// </summary>
        public readonly string Schedule;
        /// <summary>
        /// Configure virtual cluster 2. The structure of `secondary_vcluster` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHaSecondaryVclusterResult> SecondaryVclusters;
        /// <summary>
        /// Enable/disable session pickup. Enabling it can reduce session down time when fail over happens.
        /// </summary>
        public readonly string SessionPickup;
        /// <summary>
        /// Enable/disable UDP and ICMP session sync.
        /// </summary>
        public readonly string SessionPickupConnectionless;
        /// <summary>
        /// Enable to sync sessions longer than 30 sec. Only longer lived sessions need to be synced.
        /// </summary>
        public readonly string SessionPickupDelay;
        /// <summary>
        /// Enable/disable session helper expectation session sync for FGSP.
        /// </summary>
        public readonly string SessionPickupExpectation;
        /// <summary>
        /// Enable/disable NAT session sync for FGSP.
        /// </summary>
        public readonly string SessionPickupNat;
        /// <summary>
        /// Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        /// </summary>
        public readonly string SessionSyncDev;
        /// <summary>
        /// Dynamic weighted load balancing weight and high and low number of SMTP proxy sessions.
        /// </summary>
        public readonly string SmtpProxyThreshold;
        /// <summary>
        /// Enable/disable automatic HA failover on SSD disk failure.
        /// </summary>
        public readonly string SsdFailover;
        /// <summary>
        /// Enable/disable FGSP configuration synchronization.
        /// </summary>
        public readonly string StandaloneConfigSync;
        /// <summary>
        /// Enable/disable standalone management VDOM.
        /// </summary>
        public readonly string StandaloneMgmtVdom;
        /// <summary>
        /// Enable/disable configuration synchronization.
        /// </summary>
        public readonly string SyncConfig;
        /// <summary>
        /// Enable/disable HA packet distribution to multiple CPUs.
        /// </summary>
        public readonly string SyncPacketBalance;
        /// <summary>
        /// Default route gateway for unicast interface.
        /// </summary>
        public readonly string UnicastGateway;
        /// <summary>
        /// Enable/disable unicast heartbeat.
        /// </summary>
        public readonly string UnicastHb;
        /// <summary>
        /// Unicast heartbeat netmask.
        /// </summary>
        public readonly string UnicastHbNetmask;
        /// <summary>
        /// Unicast heartbeat peer IP.
        /// </summary>
        public readonly string UnicastHbPeerip;
        /// <summary>
        /// Number of unicast peers. The structure of `unicast_peers` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHaUnicastPeerResult> UnicastPeers;
        /// <summary>
        /// Enable/disable unicast connection.
        /// </summary>
        public readonly string UnicastStatus;
        /// <summary>
        /// Number of minutes the primary HA unit waits before the secondary HA unit is considered upgraded and the system is started before starting its own upgrade (1 - 300, default = 30).
        /// </summary>
        public readonly int UninterruptiblePrimaryWait;
        /// <summary>
        /// Enable to upgrade a cluster without blocking network traffic.
        /// </summary>
        public readonly string UninterruptibleUpgrade;
        /// <summary>
        /// The mode to upgrade a cluster.
        /// </summary>
        public readonly string UpgradeMode;
        /// <summary>
        /// Enable/disable virtual cluster 2 for virtual clustering.
        /// </summary>
        public readonly string Vcluster2;
        /// <summary>
        /// Cluster ID.
        /// </summary>
        public readonly int VclusterId;
        /// <summary>
        /// Enable/disable virtual cluster for virtual clustering.
        /// </summary>
        public readonly string VclusterStatus;
        /// <summary>
        /// Virtual cluster table. The structure of `vcluster` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHaVclusterResult> Vclusters;
        /// <summary>
        /// VDOMs in virtual cluster 2.
        /// </summary>
        public readonly string Vdom;
        public readonly string? Vdomparam;
        /// <summary>
        /// Weight-round-robin weight for each cluster unit. Syntax &lt;priority&gt; &lt;weight&gt;.
        /// </summary>
        public readonly string Weight;

        [OutputConstructor]
        private GetHaResult(
            int arps,

            int arpsInterval,

            string authentication,

            string cpuThreshold,

            string encryption,

            int evpnTtl,

            int failoverHoldTime,

            string ftpProxyThreshold,

            string gratuitousArps,

            int groupId,

            string groupName,

            string haDirect,

            string haEthType,

            ImmutableArray<Outputs.GetHaHaMgmtInterfaceResult> haMgmtInterfaces,

            string haMgmtStatus,

            int haUptimeDiffMargin,

            int hbInterval,

            string hbIntervalInMilliseconds,

            int hbLostThreshold,

            string hbdev,

            string hcEthType,

            int helloHolddown,

            string httpProxyThreshold,

            string id,

            string imapProxyThreshold,

            string interClusterSessionSync,

            string ipsecPhase2Proposal,

            string key,

            string l2epEthType,

            string linkFailedSignal,

            string loadBalanceAll,

            string logicalSn,

            string memoryBasedFailover,

            string memoryCompatibleMode,

            int memoryFailoverFlipTimeout,

            int memoryFailoverMonitorPeriod,

            int memoryFailoverSampleRate,

            int memoryFailoverThreshold,

            string memoryThreshold,

            string mode,

            string monitor,

            int multicastTtl,

            string nntpProxyThreshold,

            string @override,

            int overrideWaitTime,

            string password,

            int pingserverFailoverThreshold,

            int pingserverFlipTimeout,

            string pingserverMonitorInterface,

            string pingserverSecondaryForceReset,

            string pingserverSlaveForceReset,

            string pop3ProxyThreshold,

            int priority,

            int routeHold,

            int routeTtl,

            int routeWait,

            string schedule,

            ImmutableArray<Outputs.GetHaSecondaryVclusterResult> secondaryVclusters,

            string sessionPickup,

            string sessionPickupConnectionless,

            string sessionPickupDelay,

            string sessionPickupExpectation,

            string sessionPickupNat,

            string sessionSyncDev,

            string smtpProxyThreshold,

            string ssdFailover,

            string standaloneConfigSync,

            string standaloneMgmtVdom,

            string syncConfig,

            string syncPacketBalance,

            string unicastGateway,

            string unicastHb,

            string unicastHbNetmask,

            string unicastHbPeerip,

            ImmutableArray<Outputs.GetHaUnicastPeerResult> unicastPeers,

            string unicastStatus,

            int uninterruptiblePrimaryWait,

            string uninterruptibleUpgrade,

            string upgradeMode,

            string vcluster2,

            int vclusterId,

            string vclusterStatus,

            ImmutableArray<Outputs.GetHaVclusterResult> vclusters,

            string vdom,

            string? vdomparam,

            string weight)
        {
            Arps = arps;
            ArpsInterval = arpsInterval;
            Authentication = authentication;
            CpuThreshold = cpuThreshold;
            Encryption = encryption;
            EvpnTtl = evpnTtl;
            FailoverHoldTime = failoverHoldTime;
            FtpProxyThreshold = ftpProxyThreshold;
            GratuitousArps = gratuitousArps;
            GroupId = groupId;
            GroupName = groupName;
            HaDirect = haDirect;
            HaEthType = haEthType;
            HaMgmtInterfaces = haMgmtInterfaces;
            HaMgmtStatus = haMgmtStatus;
            HaUptimeDiffMargin = haUptimeDiffMargin;
            HbInterval = hbInterval;
            HbIntervalInMilliseconds = hbIntervalInMilliseconds;
            HbLostThreshold = hbLostThreshold;
            Hbdev = hbdev;
            HcEthType = hcEthType;
            HelloHolddown = helloHolddown;
            HttpProxyThreshold = httpProxyThreshold;
            Id = id;
            ImapProxyThreshold = imapProxyThreshold;
            InterClusterSessionSync = interClusterSessionSync;
            IpsecPhase2Proposal = ipsecPhase2Proposal;
            Key = key;
            L2epEthType = l2epEthType;
            LinkFailedSignal = linkFailedSignal;
            LoadBalanceAll = loadBalanceAll;
            LogicalSn = logicalSn;
            MemoryBasedFailover = memoryBasedFailover;
            MemoryCompatibleMode = memoryCompatibleMode;
            MemoryFailoverFlipTimeout = memoryFailoverFlipTimeout;
            MemoryFailoverMonitorPeriod = memoryFailoverMonitorPeriod;
            MemoryFailoverSampleRate = memoryFailoverSampleRate;
            MemoryFailoverThreshold = memoryFailoverThreshold;
            MemoryThreshold = memoryThreshold;
            Mode = mode;
            Monitor = monitor;
            MulticastTtl = multicastTtl;
            NntpProxyThreshold = nntpProxyThreshold;
            Override = @override;
            OverrideWaitTime = overrideWaitTime;
            Password = password;
            PingserverFailoverThreshold = pingserverFailoverThreshold;
            PingserverFlipTimeout = pingserverFlipTimeout;
            PingserverMonitorInterface = pingserverMonitorInterface;
            PingserverSecondaryForceReset = pingserverSecondaryForceReset;
            PingserverSlaveForceReset = pingserverSlaveForceReset;
            Pop3ProxyThreshold = pop3ProxyThreshold;
            Priority = priority;
            RouteHold = routeHold;
            RouteTtl = routeTtl;
            RouteWait = routeWait;
            Schedule = schedule;
            SecondaryVclusters = secondaryVclusters;
            SessionPickup = sessionPickup;
            SessionPickupConnectionless = sessionPickupConnectionless;
            SessionPickupDelay = sessionPickupDelay;
            SessionPickupExpectation = sessionPickupExpectation;
            SessionPickupNat = sessionPickupNat;
            SessionSyncDev = sessionSyncDev;
            SmtpProxyThreshold = smtpProxyThreshold;
            SsdFailover = ssdFailover;
            StandaloneConfigSync = standaloneConfigSync;
            StandaloneMgmtVdom = standaloneMgmtVdom;
            SyncConfig = syncConfig;
            SyncPacketBalance = syncPacketBalance;
            UnicastGateway = unicastGateway;
            UnicastHb = unicastHb;
            UnicastHbNetmask = unicastHbNetmask;
            UnicastHbPeerip = unicastHbPeerip;
            UnicastPeers = unicastPeers;
            UnicastStatus = unicastStatus;
            UninterruptiblePrimaryWait = uninterruptiblePrimaryWait;
            UninterruptibleUpgrade = uninterruptibleUpgrade;
            UpgradeMode = upgradeMode;
            Vcluster2 = vcluster2;
            VclusterId = vclusterId;
            VclusterStatus = vclusterStatus;
            Vclusters = vclusters;
            Vdom = vdom;
            Vdomparam = vdomparam;
            Weight = weight;
        }
    }
}
