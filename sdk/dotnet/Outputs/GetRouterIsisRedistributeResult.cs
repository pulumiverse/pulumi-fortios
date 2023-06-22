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
    public sealed class GetRouterIsisRedistributeResult
    {
        /// <summary>
        /// Level.
        /// </summary>
        public readonly string Level;
        /// <summary>
        /// Metric.
        /// </summary>
        public readonly int Metric;
        /// <summary>
        /// Metric type.
        /// </summary>
        public readonly string MetricType;
        /// <summary>
        /// Protocol name.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// Route map name.
        /// </summary>
        public readonly string Routemap;
        /// <summary>
        /// Enable/disable redistribution.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetRouterIsisRedistributeResult(
            string level,

            int metric,

            string metricType,

            string protocol,

            string routemap,

            string status)
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
