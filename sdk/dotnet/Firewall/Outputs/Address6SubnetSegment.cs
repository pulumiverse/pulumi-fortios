// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Outputs
{

    [OutputType]
    public sealed class Address6SubnetSegment
    {
        /// <summary>
        /// Name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Subnet segment type. Valid values: `any`, `specific`.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Subnet segment value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private Address6SubnetSegment(
            string? name,

            string? type,

            string? value)
        {
            Name = name;
            Type = type;
            Value = value;
        }
    }
}
