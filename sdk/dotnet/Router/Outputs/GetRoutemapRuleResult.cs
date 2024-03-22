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
    public sealed class GetRoutemapRuleResult
    {
        /// <summary>
        /// Action.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Rule ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Match BGP AS path list.
        /// </summary>
        public readonly string MatchAsPath;
        /// <summary>
        /// Match BGP community list.
        /// </summary>
        public readonly string MatchCommunity;
        /// <summary>
        /// Enable/disable exact matching of communities.
        /// </summary>
        public readonly string MatchCommunityExact;
        /// <summary>
        /// Match BGP extended community list.
        /// </summary>
        public readonly string MatchExtcommunity;
        /// <summary>
        /// Enable/disable exact matching of extended communities.
        /// </summary>
        public readonly string MatchExtcommunityExact;
        /// <summary>
        /// BGP flag value to match (0 - 65535)
        /// </summary>
        public readonly int MatchFlags;
        /// <summary>
        /// Match interface configuration.
        /// </summary>
        public readonly string MatchInterface;
        /// <summary>
        /// Match IPv6 address permitted by access-list6 or prefix-list6.
        /// </summary>
        public readonly string MatchIp6Address;
        /// <summary>
        /// Match next hop IPv6 address passed by access-list6 or prefix-list6.
        /// </summary>
        public readonly string MatchIp6Nexthop;
        /// <summary>
        /// Match IP address permitted by access-list or prefix-list.
        /// </summary>
        public readonly string MatchIpAddress;
        /// <summary>
        /// Match next hop IP address passed by access-list or prefix-list.
        /// </summary>
        public readonly string MatchIpNexthop;
        /// <summary>
        /// Match metric for redistribute routes.
        /// </summary>
        public readonly int MatchMetric;
        /// <summary>
        /// Match BGP origin code.
        /// </summary>
        public readonly string MatchOrigin;
        /// <summary>
        /// Match route type.
        /// </summary>
        public readonly string MatchRouteType;
        /// <summary>
        /// Match tag.
        /// </summary>
        public readonly int MatchTag;
        /// <summary>
        /// Match VRF ID.
        /// </summary>
        public readonly int MatchVrf;
        /// <summary>
        /// BGP aggregator AS.
        /// </summary>
        public readonly int SetAggregatorAs;
        /// <summary>
        /// BGP aggregator IP.
        /// </summary>
        public readonly string SetAggregatorIp;
        /// <summary>
        /// Specify preferred action of set-aspath.
        /// </summary>
        public readonly string SetAspathAction;
        /// <summary>
        /// Prepend BGP AS path attribute. The structure of `set_aspath` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRoutemapRuleSetAspathResult> SetAspaths;
        /// <summary>
        /// Enable/disable BGP atomic aggregate attribute.
        /// </summary>
        public readonly string SetAtomicAggregate;
        /// <summary>
        /// BGP community attribute. The structure of `set_community` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRoutemapRuleSetCommunityResult> SetCommunities;
        /// <summary>
        /// Enable/disable adding set-community to existing community.
        /// </summary>
        public readonly string SetCommunityAdditive;
        /// <summary>
        /// Delete communities matching community list.
        /// </summary>
        public readonly string SetCommunityDelete;
        /// <summary>
        /// Maximum duration to suppress a route (1 - 255 min, 0 = unset).
        /// </summary>
        public readonly int SetDampeningMaxSuppress;
        /// <summary>
        /// Reachability half-life time for the penalty (1 - 45 min, 0 = unset).
        /// </summary>
        public readonly int SetDampeningReachabilityHalfLife;
        /// <summary>
        /// Value to start reusing a route (1 - 20000, 0 = unset).
        /// </summary>
        public readonly int SetDampeningReuse;
        /// <summary>
        /// Value to start suppressing a route (1 - 20000, 0 = unset).
        /// </summary>
        public readonly int SetDampeningSuppress;
        /// <summary>
        /// Unreachability Half-life time for the penalty (1 - 45 min, 0 = unset)
        /// </summary>
        public readonly int SetDampeningUnreachabilityHalfLife;
        /// <summary>
        /// Route Target extended community. The structure of `set_extcommunity_rt` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRoutemapRuleSetExtcommunityRtResult> SetExtcommunityRts;
        /// <summary>
        /// Site-of-Origin extended community. The structure of `set_extcommunity_soo` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRoutemapRuleSetExtcommunitySooResult> SetExtcommunitySoos;
        /// <summary>
        /// BGP flags value (0 - 65535)
        /// </summary>
        public readonly int SetFlags;
        /// <summary>
        /// IPv6 global address of next hop.
        /// </summary>
        public readonly string SetIp6Nexthop;
        /// <summary>
        /// IPv6 local address of next hop.
        /// </summary>
        public readonly string SetIp6NexthopLocal;
        /// <summary>
        /// IP address of next hop.
        /// </summary>
        public readonly string SetIpNexthop;
        /// <summary>
        /// IP address of preferred source.
        /// </summary>
        public readonly string SetIpPrefsrc;
        /// <summary>
        /// BGP local preference path attribute.
        /// </summary>
        public readonly int SetLocalPreference;
        /// <summary>
        /// Metric value.
        /// </summary>
        public readonly int SetMetric;
        /// <summary>
        /// Metric type.
        /// </summary>
        public readonly string SetMetricType;
        /// <summary>
        /// BGP origin code.
        /// </summary>
        public readonly string SetOrigin;
        /// <summary>
        /// BGP originator ID attribute.
        /// </summary>
        public readonly string SetOriginatorId;
        /// <summary>
        /// Priority for routing table.
        /// </summary>
        public readonly int SetPriority;
        /// <summary>
        /// Route tag for routing table.
        /// </summary>
        public readonly int SetRouteTag;
        /// <summary>
        /// Tag value.
        /// </summary>
        public readonly int SetTag;
        /// <summary>
        /// IP address of VPNv4 next-hop.
        /// </summary>
        public readonly string SetVpnv4Nexthop;
        /// <summary>
        /// IPv6 global address of VPNv6 next-hop.
        /// </summary>
        public readonly string SetVpnv6Nexthop;
        /// <summary>
        /// IPv6 link-local address of VPNv6 next-hop.
        /// </summary>
        public readonly string SetVpnv6NexthopLocal;
        /// <summary>
        /// BGP weight for routing table.
        /// </summary>
        public readonly int SetWeight;

        [OutputConstructor]
        private GetRoutemapRuleResult(
            string action,

            int id,

            string matchAsPath,

            string matchCommunity,

            string matchCommunityExact,

            string matchExtcommunity,

            string matchExtcommunityExact,

            int matchFlags,

            string matchInterface,

            string matchIp6Address,

            string matchIp6Nexthop,

            string matchIpAddress,

            string matchIpNexthop,

            int matchMetric,

            string matchOrigin,

            string matchRouteType,

            int matchTag,

            int matchVrf,

            int setAggregatorAs,

            string setAggregatorIp,

            string setAspathAction,

            ImmutableArray<Outputs.GetRoutemapRuleSetAspathResult> setAspaths,

            string setAtomicAggregate,

            ImmutableArray<Outputs.GetRoutemapRuleSetCommunityResult> setCommunities,

            string setCommunityAdditive,

            string setCommunityDelete,

            int setDampeningMaxSuppress,

            int setDampeningReachabilityHalfLife,

            int setDampeningReuse,

            int setDampeningSuppress,

            int setDampeningUnreachabilityHalfLife,

            ImmutableArray<Outputs.GetRoutemapRuleSetExtcommunityRtResult> setExtcommunityRts,

            ImmutableArray<Outputs.GetRoutemapRuleSetExtcommunitySooResult> setExtcommunitySoos,

            int setFlags,

            string setIp6Nexthop,

            string setIp6NexthopLocal,

            string setIpNexthop,

            string setIpPrefsrc,

            int setLocalPreference,

            int setMetric,

            string setMetricType,

            string setOrigin,

            string setOriginatorId,

            int setPriority,

            int setRouteTag,

            int setTag,

            string setVpnv4Nexthop,

            string setVpnv6Nexthop,

            string setVpnv6NexthopLocal,

            int setWeight)
        {
            Action = action;
            Id = id;
            MatchAsPath = matchAsPath;
            MatchCommunity = matchCommunity;
            MatchCommunityExact = matchCommunityExact;
            MatchExtcommunity = matchExtcommunity;
            MatchExtcommunityExact = matchExtcommunityExact;
            MatchFlags = matchFlags;
            MatchInterface = matchInterface;
            MatchIp6Address = matchIp6Address;
            MatchIp6Nexthop = matchIp6Nexthop;
            MatchIpAddress = matchIpAddress;
            MatchIpNexthop = matchIpNexthop;
            MatchMetric = matchMetric;
            MatchOrigin = matchOrigin;
            MatchRouteType = matchRouteType;
            MatchTag = matchTag;
            MatchVrf = matchVrf;
            SetAggregatorAs = setAggregatorAs;
            SetAggregatorIp = setAggregatorIp;
            SetAspathAction = setAspathAction;
            SetAspaths = setAspaths;
            SetAtomicAggregate = setAtomicAggregate;
            SetCommunities = setCommunities;
            SetCommunityAdditive = setCommunityAdditive;
            SetCommunityDelete = setCommunityDelete;
            SetDampeningMaxSuppress = setDampeningMaxSuppress;
            SetDampeningReachabilityHalfLife = setDampeningReachabilityHalfLife;
            SetDampeningReuse = setDampeningReuse;
            SetDampeningSuppress = setDampeningSuppress;
            SetDampeningUnreachabilityHalfLife = setDampeningUnreachabilityHalfLife;
            SetExtcommunityRts = setExtcommunityRts;
            SetExtcommunitySoos = setExtcommunitySoos;
            SetFlags = setFlags;
            SetIp6Nexthop = setIp6Nexthop;
            SetIp6NexthopLocal = setIp6NexthopLocal;
            SetIpNexthop = setIpNexthop;
            SetIpPrefsrc = setIpPrefsrc;
            SetLocalPreference = setLocalPreference;
            SetMetric = setMetric;
            SetMetricType = setMetricType;
            SetOrigin = setOrigin;
            SetOriginatorId = setOriginatorId;
            SetPriority = setPriority;
            SetRouteTag = setRouteTag;
            SetTag = setTag;
            SetVpnv4Nexthop = setVpnv4Nexthop;
            SetVpnv6Nexthop = setVpnv6Nexthop;
            SetVpnv6NexthopLocal = setVpnv6NexthopLocal;
            SetWeight = setWeight;
        }
    }
}
