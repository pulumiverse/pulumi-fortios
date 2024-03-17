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
    public sealed class GetMulticastInterfaceResult
    {
        /// <summary>
        /// Enable/disable Protocol Independent Multicast (PIM) Bidirectional Forwarding Detection (BFD).
        /// </summary>
        public readonly string Bfd;
        /// <summary>
        /// Exclude GenID from hello packets (compatibility with old Cisco IOS).
        /// </summary>
        public readonly string CiscoExcludeGenid;
        /// <summary>
        /// DR election priority.
        /// </summary>
        public readonly int DrPriority;
        /// <summary>
        /// Time before old neighbor information expires (0 - 65535 sec, default = 105).
        /// </summary>
        public readonly int HelloHoldtime;
        /// <summary>
        /// Interval between sending PIM hello messages (0 - 65535 sec, default = 30).
        /// </summary>
        public readonly int HelloInterval;
        /// <summary>
        /// IGMP configuration options. The structure of `igmp` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMulticastInterfaceIgmpResult> Igmps;
        /// <summary>
        /// Join multicast groups. The structure of `join_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMulticastInterfaceJoinGroupResult> JoinGroups;
        /// <summary>
        /// Acceptable source for multicast group.
        /// </summary>
        public readonly string MulticastFlow;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Routers acknowledged as neighbor routers.
        /// </summary>
        public readonly string NeighbourFilter;
        /// <summary>
        /// Enable/disable listening to IGMP but not participating in PIM.
        /// </summary>
        public readonly string Passive;
        /// <summary>
        /// PIM operation mode.
        /// </summary>
        public readonly string PimMode;
        /// <summary>
        /// Delay flooding packets on this interface (100 - 5000 msec, default = 500).
        /// </summary>
        public readonly int PropagationDelay;
        /// <summary>
        /// Enable/disable compete to become RP in elections.
        /// </summary>
        public readonly string RpCandidate;
        /// <summary>
        /// Multicast groups managed by this RP.
        /// </summary>
        public readonly string RpCandidateGroup;
        /// <summary>
        /// RP candidate advertisement interval (1 - 16383 sec, default = 60).
        /// </summary>
        public readonly int RpCandidateInterval;
        /// <summary>
        /// Router's priority as RP.
        /// </summary>
        public readonly int RpCandidatePriority;
        /// <summary>
        /// Enable/disable fail back for RPF neighbor query.
        /// </summary>
        public readonly string RpfNbrFailBack;
        /// <summary>
        /// Filter for fail back RPF neighbors.
        /// </summary>
        public readonly string RpfNbrFailBackFilter;
        /// <summary>
        /// Interval between sending state-refresh packets (1 - 100 sec, default = 60).
        /// </summary>
        public readonly int StateRefreshInterval;
        /// <summary>
        /// Statically set multicast groups to forward out.
        /// </summary>
        public readonly string StaticGroup;
        /// <summary>
        /// Minimum TTL of multicast packets that will be forwarded (applied only to new multicast routes) (1 - 255, default = 1).
        /// </summary>
        public readonly int TtlThreshold;

        [OutputConstructor]
        private GetMulticastInterfaceResult(
            string bfd,

            string ciscoExcludeGenid,

            int drPriority,

            int helloHoldtime,

            int helloInterval,

            ImmutableArray<Outputs.GetMulticastInterfaceIgmpResult> igmps,

            ImmutableArray<Outputs.GetMulticastInterfaceJoinGroupResult> joinGroups,

            string multicastFlow,

            string name,

            string neighbourFilter,

            string passive,

            string pimMode,

            int propagationDelay,

            string rpCandidate,

            string rpCandidateGroup,

            int rpCandidateInterval,

            int rpCandidatePriority,

            string rpfNbrFailBack,

            string rpfNbrFailBackFilter,

            int stateRefreshInterval,

            string staticGroup,

            int ttlThreshold)
        {
            Bfd = bfd;
            CiscoExcludeGenid = ciscoExcludeGenid;
            DrPriority = drPriority;
            HelloHoldtime = helloHoldtime;
            HelloInterval = helloInterval;
            Igmps = igmps;
            JoinGroups = joinGroups;
            MulticastFlow = multicastFlow;
            Name = name;
            NeighbourFilter = neighbourFilter;
            Passive = passive;
            PimMode = pimMode;
            PropagationDelay = propagationDelay;
            RpCandidate = rpCandidate;
            RpCandidateGroup = rpCandidateGroup;
            RpCandidateInterval = rpCandidateInterval;
            RpCandidatePriority = rpCandidatePriority;
            RpfNbrFailBack = rpfNbrFailBack;
            RpfNbrFailBackFilter = rpfNbrFailBackFilter;
            StateRefreshInterval = stateRefreshInterval;
            StaticGroup = staticGroup;
            TtlThreshold = ttlThreshold;
        }
    }
}
