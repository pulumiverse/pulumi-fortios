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
    public sealed class CentralmanagementServerList
    {
        /// <summary>
        /// Indicate whether the FortiGate communicates with the override server using an IPv4 address, an IPv6 address or a FQDN. Valid values: `ipv4`, `ipv6`, `fqdn`.
        /// </summary>
        public readonly string? AddrType;
        /// <summary>
        /// FQDN address of override server.
        /// </summary>
        public readonly string? Fqdn;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// IPv4 address of override server.
        /// </summary>
        public readonly string? ServerAddress;
        /// <summary>
        /// IPv6 address of override server.
        /// </summary>
        public readonly string? ServerAddress6;
        /// <summary>
        /// FortiGuard service type.
        /// </summary>
        public readonly string? ServerType;

        [OutputConstructor]
        private CentralmanagementServerList(
            string? addrType,

            string? fqdn,

            int? id,

            string? serverAddress,

            string? serverAddress6,

            string? serverType)
        {
            AddrType = addrType;
            Fqdn = fqdn;
            Id = id;
            ServerAddress = serverAddress;
            ServerAddress6 = serverAddress6;
            ServerType = serverType;
        }
    }
}
