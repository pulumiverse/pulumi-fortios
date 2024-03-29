// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class GetInterfaceIpv6Ip6DelegatedPrefixListResult
    {
        /// <summary>
        /// Enable/disable the autonomous flag.
        /// </summary>
        public readonly string AutonomousFlag;
        /// <summary>
        /// IAID of obtained delegated-prefix from the upstream interface.
        /// </summary>
        public readonly int DelegatedPrefixIaid;
        /// <summary>
        /// Enable/disable the onlink flag.
        /// </summary>
        public readonly string OnlinkFlag;
        /// <summary>
        /// Prefix ID.
        /// </summary>
        public readonly int PrefixId;
        /// <summary>
        /// Recursive DNS server option.
        /// </summary>
        public readonly string Rdnss;
        /// <summary>
        /// Recursive DNS service option.
        /// </summary>
        public readonly string RdnssService;
        /// <summary>
        /// Add subnet ID to routing prefix.
        /// </summary>
        public readonly string Subnet;
        /// <summary>
        /// Name of the interface that provides delegated information.
        /// </summary>
        public readonly string UpstreamInterface;

        [OutputConstructor]
        private GetInterfaceIpv6Ip6DelegatedPrefixListResult(
            string autonomousFlag,

            int delegatedPrefixIaid,

            string onlinkFlag,

            int prefixId,

            string rdnss,

            string rdnssService,

            string subnet,

            string upstreamInterface)
        {
            AutonomousFlag = autonomousFlag;
            DelegatedPrefixIaid = delegatedPrefixIaid;
            OnlinkFlag = onlinkFlag;
            PrefixId = prefixId;
            Rdnss = rdnss;
            RdnssService = rdnssService;
            Subnet = subnet;
            UpstreamInterface = upstreamInterface;
        }
    }
}
