// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Web.Outputs
{

    [OutputType]
    public sealed class ProfileFtgdWfQuota
    {
        /// <summary>
        /// FortiGuard categories to apply quota to (category action must be set to monitor).
        /// </summary>
        public readonly string? Category;
        /// <summary>
        /// Duration of quota.
        /// </summary>
        public readonly string? Duration;
        /// <summary>
        /// ID number.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Override replacement message.
        /// </summary>
        public readonly string? OverrideReplacemsg;
        /// <summary>
        /// Quota type. Valid values: `time`, `traffic`.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Traffic quota unit of measurement. Valid values: `B`, `KB`, `MB`, `GB`.
        /// </summary>
        public readonly string? Unit;
        /// <summary>
        /// Traffic quota value.
        /// </summary>
        public readonly int? Value;

        [OutputConstructor]
        private ProfileFtgdWfQuota(
            string? category,

            string? duration,

            int? id,

            string? overrideReplacemsg,

            string? type,

            string? unit,

            int? value)
        {
            Category = category;
            Duration = duration;
            Id = id;
            OverrideReplacemsg = overrideReplacemsg;
            Type = type;
            Unit = unit;
            Value = value;
        }
    }
}
