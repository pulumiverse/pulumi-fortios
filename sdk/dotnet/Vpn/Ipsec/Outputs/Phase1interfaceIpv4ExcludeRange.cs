// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ipsec.Outputs
{

    [OutputType]
    public sealed class Phase1interfaceIpv4ExcludeRange
    {
        public readonly string? EndIp;
        /// <summary>
        /// an identifier for the resource with format {{name}}.
        /// </summary>
        public readonly int? Id;
        public readonly string? StartIp;

        [OutputConstructor]
        private Phase1interfaceIpv4ExcludeRange(
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
