// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Outputs
{

    [OutputType]
    public sealed class GeoipoverrideIp6Range
    {
        /// <summary>
        /// Final IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
        /// 
        /// The `ip6_range` block supports:
        /// </summary>
        public readonly string? EndIp;
        /// <summary>
        /// ID number for individual entry in the IP-Range table.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Starting IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
        /// </summary>
        public readonly string? StartIp;

        [OutputConstructor]
        private GeoipoverrideIp6Range(
            string? endIp,

            int? id,

            string? startIp)
        {
            EndIp = endIp;
            Id = id;
            StartIp = startIp;
        }
    }
}
