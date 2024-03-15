// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class GetRipngDistanceResult
    {
        /// <summary>
        /// IPv6 access list name.
        /// </summary>
        public readonly string AccessList6;
        /// <summary>
        /// Distance (1 - 255).
        /// </summary>
        public readonly int Distance;
        /// <summary>
        /// Offset-list ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Aggregate address prefix.
        /// </summary>
        public readonly string Prefix6;

        [OutputConstructor]
        private GetRipngDistanceResult(
            string accessList6,

            int distance,

            int id,

            string prefix6)
        {
            AccessList6 = accessList6;
            Distance = distance;
            Id = id;
            Prefix6 = prefix6;
        }
    }
}