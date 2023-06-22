// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class GetRouterBfd6NeighborResult
    {
        /// <summary>
        /// Interface to the BFD neighbor.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// IPv6 address of the BFD neighbor.
        /// </summary>
        public readonly string Ip6Address;

        [OutputConstructor]
        private GetRouterBfd6NeighborResult(
            string @interface,

            string ip6Address)
        {
            Interface = @interface;
            Ip6Address = ip6Address;
        }
    }
}
