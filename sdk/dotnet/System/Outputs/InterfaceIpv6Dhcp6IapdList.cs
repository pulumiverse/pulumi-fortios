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
    public sealed class InterfaceIpv6Dhcp6IapdList
    {
        public readonly int? Iaid;
        public readonly string? PrefixHint;
        public readonly int? PrefixHintPlt;
        public readonly int? PrefixHintVlt;

        [OutputConstructor]
        private InterfaceIpv6Dhcp6IapdList(
            int? iaid,

            string? prefixHint,

            int? prefixHintPlt,

            int? prefixHintVlt)
        {
            Iaid = iaid;
            PrefixHint = prefixHint;
            PrefixHintPlt = prefixHintPlt;
            PrefixHintVlt = prefixHintVlt;
        }
    }
}
