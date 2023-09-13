// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class BgpNeighbor
    {
        /// <summary>
        /// Enable/disable address family IPv4 for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Activate;
        /// <summary>
        /// Enable/disable address family IPv6 for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Activate6;
        /// <summary>
        /// Enable/disable address family VPNv4 for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ActivateVpnv4;
        /// <summary>
        /// Enable/disable IPv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
        /// </summary>
        public readonly string? AdditionalPath;
        /// <summary>
        /// Enable/disable IPv6 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
        /// </summary>
        public readonly string? AdditionalPath6;
        /// <summary>
        /// Enable/disable VPNv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
        /// </summary>
        public readonly string? AdditionalPathVpnv4;
        /// <summary>
        /// Number of IPv4 additional paths that can be advertised to this neighbor.
        /// </summary>
        public readonly int? AdvAdditionalPath;
        /// <summary>
        /// Number of IPv6 additional paths that can be advertised to this neighbor.
        /// </summary>
        public readonly int? AdvAdditionalPath6;
        /// <summary>
        /// Number of VPNv4 additional paths that can be advertised to this neighbor.
        /// </summary>
        public readonly int? AdvAdditionalPathVpnv4;
        /// <summary>
        /// Minimum interval (sec) between sending updates.
        /// </summary>
        public readonly int? AdvertisementInterval;
        /// <summary>
        /// IPv4 The maximum number of occurrence of my AS number allowed.
        /// </summary>
        public readonly int? AllowasIn;
        /// <summary>
        /// IPv6 The maximum number of occurrence of my AS number allowed.
        /// </summary>
        public readonly int? AllowasIn6;
        /// <summary>
        /// Enable/disable IPv4 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? AllowasInEnable;
        /// <summary>
        /// Enable/disable IPv6 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? AllowasInEnable6;
        /// <summary>
        /// The maximum number of occurrence of my AS number allowed for VPNv4 route.
        /// </summary>
        public readonly int? AllowasInVpnv4;
        /// <summary>
        /// Enable/disable replace peer AS with own AS for IPv4. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? AsOverride;
        /// <summary>
        /// Enable/disable replace peer AS with own AS for IPv6. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? AsOverride6;
        /// <summary>
        /// IPv4 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
        /// </summary>
        public readonly string? AttributeUnchanged;
        /// <summary>
        /// IPv6 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
        /// </summary>
        public readonly string? AttributeUnchanged6;
        /// <summary>
        /// List of attributes that should be unchanged for VPNv4 route. Valid values: `as-path`, `med`, `next-hop`.
        /// </summary>
        public readonly string? AttributeUnchangedVpnv4;
        /// <summary>
        /// Enable/disable BFD for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Bfd;
        /// <summary>
        /// Enable/disable advertise default IPv4 route to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? CapabilityDefaultOriginate;
        /// <summary>
        /// Enable/disable advertise default IPv6 route to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? CapabilityDefaultOriginate6;
        /// <summary>
        /// Enable/disable advertise dynamic capability to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? CapabilityDynamic;
        /// <summary>
        /// Enable/disable advertise IPv4 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? CapabilityGracefulRestart;
        /// <summary>
        /// Enable/disable advertise IPv6 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? CapabilityGracefulRestart6;
        /// <summary>
        /// Enable/disable advertise VPNv4 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? CapabilityGracefulRestartVpnv4;
        /// <summary>
        /// Accept/Send IPv4 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
        /// </summary>
        public readonly string? CapabilityOrf;
        /// <summary>
        /// Accept/Send IPv6 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
        /// </summary>
        public readonly string? CapabilityOrf6;
        /// <summary>
        /// Enable/disable advertise route refresh capability to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? CapabilityRouteRefresh;
        /// <summary>
        /// IPv6 conditional advertisement. The structure of `conditional_advertise6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.BgpNeighborConditionalAdvertise6> ConditionalAdvertise6s;
        /// <summary>
        /// Conditional advertisement. The structure of `conditional_advertise` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.BgpNeighborConditionalAdvertise> ConditionalAdvertises;
        /// <summary>
        /// Interval (sec) for connect timer.
        /// </summary>
        public readonly int? ConnectTimer;
        /// <summary>
        /// Route map to specify criteria to originate IPv4 default.
        /// </summary>
        public readonly string? DefaultOriginateRoutemap;
        /// <summary>
        /// Route map to specify criteria to originate IPv6 default.
        /// </summary>
        public readonly string? DefaultOriginateRoutemap6;
        /// <summary>
        /// Description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Filter for IPv4 updates from this neighbor.
        /// </summary>
        public readonly string? DistributeListIn;
        /// <summary>
        /// Filter for IPv6 updates from this neighbor.
        /// </summary>
        public readonly string? DistributeListIn6;
        /// <summary>
        /// Filter for VPNv4 updates from this neighbor.
        /// </summary>
        public readonly string? DistributeListInVpnv4;
        /// <summary>
        /// Filter for IPv4 updates to this neighbor.
        /// </summary>
        public readonly string? DistributeListOut;
        /// <summary>
        /// Filter for IPv6 updates to this neighbor.
        /// </summary>
        public readonly string? DistributeListOut6;
        /// <summary>
        /// Filter for VPNv4 updates to this neighbor.
        /// </summary>
        public readonly string? DistributeListOutVpnv4;
        /// <summary>
        /// Don't negotiate capabilities with this neighbor Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? DontCapabilityNegotiate;
        /// <summary>
        /// Enable/disable allow multi-hop EBGP neighbors. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? EbgpEnforceMultihop;
        /// <summary>
        /// EBGP multihop TTL for this peer.
        /// </summary>
        public readonly int? EbgpMultihopTtl;
        /// <summary>
        /// BGP filter for IPv4 inbound routes.
        /// </summary>
        public readonly string? FilterListIn;
        /// <summary>
        /// BGP filter for IPv6 inbound routes.
        /// </summary>
        public readonly string? FilterListIn6;
        /// <summary>
        /// BGP filter for IPv4 outbound routes.
        /// </summary>
        public readonly string? FilterListOut;
        /// <summary>
        /// BGP filter for IPv6 outbound routes.
        /// </summary>
        public readonly string? FilterListOut6;
        /// <summary>
        /// Interval (sec) before peer considered dead.
        /// </summary>
        public readonly int? HoldtimeTimer;
        /// <summary>
        /// Interface
        /// </summary>
        public readonly string? Interface;
        /// <summary>
        /// IP/IPv6 address of neighbor.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Keep alive timer interval (sec).
        /// </summary>
        public readonly int? KeepAliveTimer;
        /// <summary>
        /// Enable/disable failover upon link down. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? LinkDownFailover;
        /// <summary>
        /// Local AS number of neighbor.
        /// </summary>
        public readonly int? LocalAs;
        /// <summary>
        /// Do not prepend local-as to incoming updates. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? LocalAsNoPrepend;
        /// <summary>
        /// Replace real AS with local-as in outgoing updates. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? LocalAsReplaceAs;
        /// <summary>
        /// Maximum number of IPv4 prefixes to accept from this peer.
        /// </summary>
        public readonly int? MaximumPrefix;
        /// <summary>
        /// Maximum number of IPv6 prefixes to accept from this peer.
        /// </summary>
        public readonly int? MaximumPrefix6;
        /// <summary>
        /// Maximum IPv4 prefix threshold value (1 - 100 percent).
        /// </summary>
        public readonly int? MaximumPrefixThreshold;
        /// <summary>
        /// Maximum IPv6 prefix threshold value (1 - 100 percent).
        /// </summary>
        public readonly int? MaximumPrefixThreshold6;
        /// <summary>
        /// Maximum VPNv4 prefix threshold value (1 - 100 percent).
        /// </summary>
        public readonly int? MaximumPrefixThresholdVpnv4;
        /// <summary>
        /// Maximum number of VPNv4 prefixes to accept from this peer.
        /// </summary>
        public readonly int? MaximumPrefixVpnv4;
        /// <summary>
        /// Enable/disable IPv4 Only give warning message when limit is exceeded. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? MaximumPrefixWarningOnly;
        /// <summary>
        /// Enable/disable IPv6 Only give warning message when limit is exceeded. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? MaximumPrefixWarningOnly6;
        /// <summary>
        /// Enable/disable only giving warning message when limit is exceeded for VPNv4 routes. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? MaximumPrefixWarningOnlyVpnv4;
        /// <summary>
        /// Enable/disable IPv4 next-hop calculation for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? NextHopSelf;
        /// <summary>
        /// Enable/disable IPv6 next-hop calculation for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? NextHopSelf6;
        /// <summary>
        /// Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? NextHopSelfRr;
        /// <summary>
        /// Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? NextHopSelfRr6;
        /// <summary>
        /// Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? NextHopSelfVpnv4;
        /// <summary>
        /// Enable/disable override result of capability negotiation. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? OverrideCapability;
        /// <summary>
        /// Enable/disable sending of open messages to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Passive;
        /// <summary>
        /// Password used in MD5 authentication.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// IPv4 Inbound filter for updates from this neighbor.
        /// </summary>
        public readonly string? PrefixListIn;
        /// <summary>
        /// IPv6 Inbound filter for updates from this neighbor.
        /// </summary>
        public readonly string? PrefixListIn6;
        /// <summary>
        /// Inbound filter for VPNv4 updates from this neighbor.
        /// </summary>
        public readonly string? PrefixListInVpnv4;
        /// <summary>
        /// IPv4 Outbound filter for updates to this neighbor.
        /// </summary>
        public readonly string? PrefixListOut;
        /// <summary>
        /// IPv6 Outbound filter for updates to this neighbor.
        /// </summary>
        public readonly string? PrefixListOut6;
        /// <summary>
        /// Outbound filter for VPNv4 updates to this neighbor.
        /// </summary>
        public readonly string? PrefixListOutVpnv4;
        /// <summary>
        /// AS number of neighbor.
        /// </summary>
        public readonly int? RemoteAs;
        /// <summary>
        /// Enable/disable remove private AS number from IPv4 outbound updates. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? RemovePrivateAs;
        /// <summary>
        /// Enable/disable remove private AS number from IPv6 outbound updates. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? RemovePrivateAs6;
        /// <summary>
        /// Enable/disable remove private AS number from VPNv4 outbound updates. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? RemovePrivateAsVpnv4;
        /// <summary>
        /// Graceful restart delay time (sec, 0 = global default).
        /// </summary>
        public readonly int? RestartTime;
        /// <summary>
        /// Time to retain stale routes.
        /// </summary>
        public readonly int? RetainStaleTime;
        /// <summary>
        /// IPv4 Inbound route map filter.
        /// </summary>
        public readonly string? RouteMapIn;
        /// <summary>
        /// IPv6 Inbound route map filter.
        /// </summary>
        public readonly string? RouteMapIn6;
        /// <summary>
        /// VPNv4 inbound route map filter.
        /// </summary>
        public readonly string? RouteMapInVpnv4;
        /// <summary>
        /// IPv4 Outbound route map filter.
        /// </summary>
        public readonly string? RouteMapOut;
        /// <summary>
        /// IPv6 Outbound route map filter.
        /// </summary>
        public readonly string? RouteMapOut6;
        /// <summary>
        /// IPv6 outbound route map filter if the peer is preferred.
        /// </summary>
        public readonly string? RouteMapOut6Preferable;
        /// <summary>
        /// IPv4 outbound route map filter if the peer is preferred.
        /// </summary>
        public readonly string? RouteMapOutPreferable;
        /// <summary>
        /// VPNv4 outbound route map filter.
        /// </summary>
        public readonly string? RouteMapOutVpnv4;
        /// <summary>
        /// VPNv4 outbound route map filter if the peer is preferred.
        /// </summary>
        public readonly string? RouteMapOutVpnv4Preferable;
        /// <summary>
        /// Enable/disable IPv4 AS route reflector client. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? RouteReflectorClient;
        /// <summary>
        /// Enable/disable IPv6 AS route reflector client. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? RouteReflectorClient6;
        /// <summary>
        /// Enable/disable VPNv4 AS route reflector client for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? RouteReflectorClientVpnv4;
        /// <summary>
        /// Enable/disable IPv4 AS route server client. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? RouteServerClient;
        /// <summary>
        /// Enable/disable IPv6 AS route server client. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? RouteServerClient6;
        /// <summary>
        /// Enable/disable VPNv4 AS route server client for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? RouteServerClientVpnv4;
        /// <summary>
        /// IPv4 Send community attribute to neighbor. Valid values: `standard`, `extended`, `both`, `disable`.
        /// </summary>
        public readonly string? SendCommunity;
        /// <summary>
        /// IPv6 Send community attribute to neighbor. Valid values: `standard`, `extended`, `both`, `disable`.
        /// </summary>
        public readonly string? SendCommunity6;
        /// <summary>
        /// Send community attribute to neighbor for VPNv4 address family. Valid values: `standard`, `extended`, `both`, `disable`.
        /// </summary>
        public readonly string? SendCommunityVpnv4;
        /// <summary>
        /// Enable/disable shutdown this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Shutdown;
        /// <summary>
        /// Enable/disable allow IPv4 inbound soft reconfiguration. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? SoftReconfiguration;
        /// <summary>
        /// Enable/disable allow IPv6 inbound soft reconfiguration. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? SoftReconfiguration6;
        /// <summary>
        /// Enable/disable allow VPNv4 inbound soft reconfiguration. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? SoftReconfigurationVpnv4;
        /// <summary>
        /// Enable/disable stale route after neighbor down. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? StaleRoute;
        /// <summary>
        /// Enable/disable strict capability matching. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? StrictCapabilityMatch;
        /// <summary>
        /// IPv4 Route map to selectively unsuppress suppressed routes.
        /// </summary>
        public readonly string? UnsuppressMap;
        /// <summary>
        /// IPv6 Route map to selectively unsuppress suppressed routes.
        /// </summary>
        public readonly string? UnsuppressMap6;
        /// <summary>
        /// Interface to use as source IP/IPv6 address of TCP connections.
        /// </summary>
        public readonly string? UpdateSource;
        /// <summary>
        /// Neighbor weight.
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private BgpNeighbor(
            string? activate,

            string? activate6,

            string? activateVpnv4,

            string? additionalPath,

            string? additionalPath6,

            string? additionalPathVpnv4,

            int? advAdditionalPath,

            int? advAdditionalPath6,

            int? advAdditionalPathVpnv4,

            int? advertisementInterval,

            int? allowasIn,

            int? allowasIn6,

            string? allowasInEnable,

            string? allowasInEnable6,

            int? allowasInVpnv4,

            string? asOverride,

            string? asOverride6,

            string? attributeUnchanged,

            string? attributeUnchanged6,

            string? attributeUnchangedVpnv4,

            string? bfd,

            string? capabilityDefaultOriginate,

            string? capabilityDefaultOriginate6,

            string? capabilityDynamic,

            string? capabilityGracefulRestart,

            string? capabilityGracefulRestart6,

            string? capabilityGracefulRestartVpnv4,

            string? capabilityOrf,

            string? capabilityOrf6,

            string? capabilityRouteRefresh,

            ImmutableArray<Outputs.BgpNeighborConditionalAdvertise6> conditionalAdvertise6s,

            ImmutableArray<Outputs.BgpNeighborConditionalAdvertise> conditionalAdvertises,

            int? connectTimer,

            string? defaultOriginateRoutemap,

            string? defaultOriginateRoutemap6,

            string? description,

            string? distributeListIn,

            string? distributeListIn6,

            string? distributeListInVpnv4,

            string? distributeListOut,

            string? distributeListOut6,

            string? distributeListOutVpnv4,

            string? dontCapabilityNegotiate,

            string? ebgpEnforceMultihop,

            int? ebgpMultihopTtl,

            string? filterListIn,

            string? filterListIn6,

            string? filterListOut,

            string? filterListOut6,

            int? holdtimeTimer,

            string? @interface,

            string? ip,

            int? keepAliveTimer,

            string? linkDownFailover,

            int? localAs,

            string? localAsNoPrepend,

            string? localAsReplaceAs,

            int? maximumPrefix,

            int? maximumPrefix6,

            int? maximumPrefixThreshold,

            int? maximumPrefixThreshold6,

            int? maximumPrefixThresholdVpnv4,

            int? maximumPrefixVpnv4,

            string? maximumPrefixWarningOnly,

            string? maximumPrefixWarningOnly6,

            string? maximumPrefixWarningOnlyVpnv4,

            string? nextHopSelf,

            string? nextHopSelf6,

            string? nextHopSelfRr,

            string? nextHopSelfRr6,

            string? nextHopSelfVpnv4,

            string? overrideCapability,

            string? passive,

            string? password,

            string? prefixListIn,

            string? prefixListIn6,

            string? prefixListInVpnv4,

            string? prefixListOut,

            string? prefixListOut6,

            string? prefixListOutVpnv4,

            int? remoteAs,

            string? removePrivateAs,

            string? removePrivateAs6,

            string? removePrivateAsVpnv4,

            int? restartTime,

            int? retainStaleTime,

            string? routeMapIn,

            string? routeMapIn6,

            string? routeMapInVpnv4,

            string? routeMapOut,

            string? routeMapOut6,

            string? routeMapOut6Preferable,

            string? routeMapOutPreferable,

            string? routeMapOutVpnv4,

            string? routeMapOutVpnv4Preferable,

            string? routeReflectorClient,

            string? routeReflectorClient6,

            string? routeReflectorClientVpnv4,

            string? routeServerClient,

            string? routeServerClient6,

            string? routeServerClientVpnv4,

            string? sendCommunity,

            string? sendCommunity6,

            string? sendCommunityVpnv4,

            string? shutdown,

            string? softReconfiguration,

            string? softReconfiguration6,

            string? softReconfigurationVpnv4,

            string? staleRoute,

            string? strictCapabilityMatch,

            string? unsuppressMap,

            string? unsuppressMap6,

            string? updateSource,

            int? weight)
        {
            Activate = activate;
            Activate6 = activate6;
            ActivateVpnv4 = activateVpnv4;
            AdditionalPath = additionalPath;
            AdditionalPath6 = additionalPath6;
            AdditionalPathVpnv4 = additionalPathVpnv4;
            AdvAdditionalPath = advAdditionalPath;
            AdvAdditionalPath6 = advAdditionalPath6;
            AdvAdditionalPathVpnv4 = advAdditionalPathVpnv4;
            AdvertisementInterval = advertisementInterval;
            AllowasIn = allowasIn;
            AllowasIn6 = allowasIn6;
            AllowasInEnable = allowasInEnable;
            AllowasInEnable6 = allowasInEnable6;
            AllowasInVpnv4 = allowasInVpnv4;
            AsOverride = asOverride;
            AsOverride6 = asOverride6;
            AttributeUnchanged = attributeUnchanged;
            AttributeUnchanged6 = attributeUnchanged6;
            AttributeUnchangedVpnv4 = attributeUnchangedVpnv4;
            Bfd = bfd;
            CapabilityDefaultOriginate = capabilityDefaultOriginate;
            CapabilityDefaultOriginate6 = capabilityDefaultOriginate6;
            CapabilityDynamic = capabilityDynamic;
            CapabilityGracefulRestart = capabilityGracefulRestart;
            CapabilityGracefulRestart6 = capabilityGracefulRestart6;
            CapabilityGracefulRestartVpnv4 = capabilityGracefulRestartVpnv4;
            CapabilityOrf = capabilityOrf;
            CapabilityOrf6 = capabilityOrf6;
            CapabilityRouteRefresh = capabilityRouteRefresh;
            ConditionalAdvertise6s = conditionalAdvertise6s;
            ConditionalAdvertises = conditionalAdvertises;
            ConnectTimer = connectTimer;
            DefaultOriginateRoutemap = defaultOriginateRoutemap;
            DefaultOriginateRoutemap6 = defaultOriginateRoutemap6;
            Description = description;
            DistributeListIn = distributeListIn;
            DistributeListIn6 = distributeListIn6;
            DistributeListInVpnv4 = distributeListInVpnv4;
            DistributeListOut = distributeListOut;
            DistributeListOut6 = distributeListOut6;
            DistributeListOutVpnv4 = distributeListOutVpnv4;
            DontCapabilityNegotiate = dontCapabilityNegotiate;
            EbgpEnforceMultihop = ebgpEnforceMultihop;
            EbgpMultihopTtl = ebgpMultihopTtl;
            FilterListIn = filterListIn;
            FilterListIn6 = filterListIn6;
            FilterListOut = filterListOut;
            FilterListOut6 = filterListOut6;
            HoldtimeTimer = holdtimeTimer;
            Interface = @interface;
            Ip = ip;
            KeepAliveTimer = keepAliveTimer;
            LinkDownFailover = linkDownFailover;
            LocalAs = localAs;
            LocalAsNoPrepend = localAsNoPrepend;
            LocalAsReplaceAs = localAsReplaceAs;
            MaximumPrefix = maximumPrefix;
            MaximumPrefix6 = maximumPrefix6;
            MaximumPrefixThreshold = maximumPrefixThreshold;
            MaximumPrefixThreshold6 = maximumPrefixThreshold6;
            MaximumPrefixThresholdVpnv4 = maximumPrefixThresholdVpnv4;
            MaximumPrefixVpnv4 = maximumPrefixVpnv4;
            MaximumPrefixWarningOnly = maximumPrefixWarningOnly;
            MaximumPrefixWarningOnly6 = maximumPrefixWarningOnly6;
            MaximumPrefixWarningOnlyVpnv4 = maximumPrefixWarningOnlyVpnv4;
            NextHopSelf = nextHopSelf;
            NextHopSelf6 = nextHopSelf6;
            NextHopSelfRr = nextHopSelfRr;
            NextHopSelfRr6 = nextHopSelfRr6;
            NextHopSelfVpnv4 = nextHopSelfVpnv4;
            OverrideCapability = overrideCapability;
            Passive = passive;
            Password = password;
            PrefixListIn = prefixListIn;
            PrefixListIn6 = prefixListIn6;
            PrefixListInVpnv4 = prefixListInVpnv4;
            PrefixListOut = prefixListOut;
            PrefixListOut6 = prefixListOut6;
            PrefixListOutVpnv4 = prefixListOutVpnv4;
            RemoteAs = remoteAs;
            RemovePrivateAs = removePrivateAs;
            RemovePrivateAs6 = removePrivateAs6;
            RemovePrivateAsVpnv4 = removePrivateAsVpnv4;
            RestartTime = restartTime;
            RetainStaleTime = retainStaleTime;
            RouteMapIn = routeMapIn;
            RouteMapIn6 = routeMapIn6;
            RouteMapInVpnv4 = routeMapInVpnv4;
            RouteMapOut = routeMapOut;
            RouteMapOut6 = routeMapOut6;
            RouteMapOut6Preferable = routeMapOut6Preferable;
            RouteMapOutPreferable = routeMapOutPreferable;
            RouteMapOutVpnv4 = routeMapOutVpnv4;
            RouteMapOutVpnv4Preferable = routeMapOutVpnv4Preferable;
            RouteReflectorClient = routeReflectorClient;
            RouteReflectorClient6 = routeReflectorClient6;
            RouteReflectorClientVpnv4 = routeReflectorClientVpnv4;
            RouteServerClient = routeServerClient;
            RouteServerClient6 = routeServerClient6;
            RouteServerClientVpnv4 = routeServerClientVpnv4;
            SendCommunity = sendCommunity;
            SendCommunity6 = sendCommunity6;
            SendCommunityVpnv4 = sendCommunityVpnv4;
            Shutdown = shutdown;
            SoftReconfiguration = softReconfiguration;
            SoftReconfiguration6 = softReconfiguration6;
            SoftReconfigurationVpnv4 = softReconfigurationVpnv4;
            StaleRoute = staleRoute;
            StrictCapabilityMatch = strictCapabilityMatch;
            UnsuppressMap = unsuppressMap;
            UnsuppressMap6 = unsuppressMap6;
            UpdateSource = updateSource;
            Weight = weight;
        }
    }
}
