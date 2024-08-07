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
    public sealed class GetRipDistanceResult
    {
        /// <summary>
        /// Access list name.
        /// </summary>
        public readonly string AccessList;
        /// <summary>
        /// Distance (1 - 255).
        /// </summary>
        public readonly int Distance;
        /// <summary>
        /// Offset-list ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Network prefix.
        /// </summary>
        public readonly string Prefix;

        [OutputConstructor]
        private GetRipDistanceResult(
            string accessList,

            int distance,

            int id,

            string prefix)
        {
            AccessList = accessList;
            Distance = distance;
            Id = id;
            Prefix = prefix;
        }
    }
}
