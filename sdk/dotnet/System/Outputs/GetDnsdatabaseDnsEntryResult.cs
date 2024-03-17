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
    public sealed class GetDnsdatabaseDnsEntryResult
    {
        /// <summary>
        /// Canonical name of the host.
        /// </summary>
        public readonly string CanonicalName;
        /// <summary>
        /// Name of the host.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// DNS entry ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// IPv4 address of the host.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// IPv6 address of the host.
        /// </summary>
        public readonly string Ipv6;
        /// <summary>
        /// DNS entry preference, 0 is the highest preference (0 - 65535, default = 10)
        /// </summary>
        public readonly int Preference;
        /// <summary>
        /// Enable/disable resource record status.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Time-to-live for this entry (0 to 2147483647 sec, default = 0).
        /// </summary>
        public readonly int Ttl;
        /// <summary>
        /// Resource record type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDnsdatabaseDnsEntryResult(
            string canonicalName,

            string hostname,

            int id,

            string ip,

            string ipv6,

            int preference,

            string status,

            int ttl,

            string type)
        {
            CanonicalName = canonicalName;
            Hostname = hostname;
            Id = id;
            Ip = ip;
            Ipv6 = ipv6;
            Preference = preference;
            Status = status;
            Ttl = ttl;
            Type = type;
        }
    }
}
