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
    public sealed class MobiletunnelNetwork
    {
        /// <summary>
        /// Network entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Select the associated interface name from available options.
        /// </summary>
        public readonly string? Interface;
        /// <summary>
        /// Class IP and Netmask with correction (Format:xxx.xxx.xxx.xxx xxx.xxx.xxx.xxx or xxx.xxx.xxx.xxx/x).
        /// </summary>
        public readonly string? Prefix;

        [OutputConstructor]
        private MobiletunnelNetwork(
            int? id,

            string? @interface,

            string? prefix)
        {
            Id = id;
            Interface = @interface;
            Prefix = prefix;
        }
    }
}
