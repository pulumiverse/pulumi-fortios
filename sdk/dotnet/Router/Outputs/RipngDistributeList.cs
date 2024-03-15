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
    public sealed class RipngDistributeList
    {
        /// <summary>
        /// Distribute list direction. Valid values: `in`, `out`.
        /// </summary>
        public readonly string? Direction;
        /// <summary>
        /// Distribute list ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Distribute list interface name.
        /// </summary>
        public readonly string? Interface;
        /// <summary>
        /// Distribute access/prefix list name.
        /// </summary>
        public readonly string? Listname;
        /// <summary>
        /// status Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private RipngDistributeList(
            string? direction,

            int? id,

            string? @interface,

            string? listname,

            string? status)
        {
            Direction = direction;
            Id = id;
            Interface = @interface;
            Listname = listname;
            Status = status;
        }
    }
}