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
    public sealed class RouterRipngInterface
    {
        /// <summary>
        /// Flags.
        /// </summary>
        public readonly int? Flags;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Enable/disable split horizon. Valid values: `poisoned`, `regular`.
        /// </summary>
        public readonly string? SplitHorizon;
        /// <summary>
        /// Enable/disable split horizon. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? SplitHorizonStatus;

        [OutputConstructor]
        private RouterRipngInterface(
            int? flags,

            string? name,

            string? splitHorizon,

            string? splitHorizonStatus)
        {
            Flags = flags;
            Name = name;
            SplitHorizon = splitHorizon;
            SplitHorizonStatus = splitHorizonStatus;
        }
    }
}
