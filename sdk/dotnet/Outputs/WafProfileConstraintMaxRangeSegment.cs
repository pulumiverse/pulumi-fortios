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
    public sealed class WafProfileConstraintMaxRangeSegment
    {
        /// <summary>
        /// Action. Valid values: `allow`, `block`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Log;
        /// <summary>
        /// Maximum number of range segments in HTTP range line (0 to 2147483647).
        /// </summary>
        public readonly int? MaxRangeSegment;
        /// <summary>
        /// Severity. Valid values: `high`, `medium`, `low`.
        /// </summary>
        public readonly string? Severity;
        /// <summary>
        /// Enable/disable the constraint. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private WafProfileConstraintMaxRangeSegment(
            string? action,

            string? log,

            int? maxRangeSegment,

            string? severity,

            string? status)
        {
            Action = action;
            Log = log;
            MaxRangeSegment = maxRangeSegment;
            Severity = severity;
            Status = status;
        }
    }
}