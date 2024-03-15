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
    public sealed class IsisRedistribute
    {
        /// <summary>
        /// Level. Valid values: `level-1-2`, `level-1`, `level-2`.
        /// </summary>
        public readonly string? Level;
        /// <summary>
        /// Metric.
        /// </summary>
        public readonly int? Metric;
        /// <summary>
        /// Metric type. Valid values: `external`, `internal`.
        /// </summary>
        public readonly string? MetricType;
        /// <summary>
        /// Protocol name.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// Route map name.
        /// </summary>
        public readonly string? Routemap;
        /// <summary>
        /// Enable/disable redistribution. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private IsisRedistribute(
            string? level,

            int? metric,

            string? metricType,

            string? protocol,

            string? routemap,

            string? status)
        {
            Level = level;
            Metric = metric;
            MetricType = metricType;
            Protocol = protocol;
            Routemap = routemap;
            Status = status;
        }
    }
}