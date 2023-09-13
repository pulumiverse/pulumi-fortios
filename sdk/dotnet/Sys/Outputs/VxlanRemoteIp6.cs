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
    public sealed class VxlanRemoteIp6
    {
        /// <summary>
        /// IPv6 address.
        /// </summary>
        public readonly string? Ip6;

        [OutputConstructor]
        private VxlanRemoteIp6(string? ip6)
        {
            Ip6 = ip6;
        }
    }
}
