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
    public sealed class Nat64SecondaryPrefix
    {
        /// <summary>
        /// NAT64 prefix name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// NAT64 prefix.
        /// </summary>
        public readonly string? Nat64Prefix;

        [OutputConstructor]
        private Nat64SecondaryPrefix(
            string? name,

            string? nat64Prefix)
        {
            Name = name;
            Nat64Prefix = nat64Prefix;
        }
    }
}
