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
    public sealed class Logsyslogd4FilterFreeStyle
    {
        /// <summary>
        /// Log category.
        /// </summary>
        public readonly string? Category;
        /// <summary>
        /// Free style filter string.
        /// </summary>
        public readonly string? Filter;
        /// <summary>
        /// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        /// </summary>
        public readonly string? FilterType;
        /// <summary>
        /// Entry ID.
        /// </summary>
        public readonly int? Id;

        [OutputConstructor]
        private Logsyslogd4FilterFreeStyle(
            string? category,

            string? filter,

            string? filterType,

            int? id)
        {
            Category = category;
            Filter = filter;
            FilterType = filterType;
            Id = id;
        }
    }
}
