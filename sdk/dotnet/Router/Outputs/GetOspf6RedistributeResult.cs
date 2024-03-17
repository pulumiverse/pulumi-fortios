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
    public sealed class GetOspf6RedistributeResult
    {
        /// <summary>
        /// Redistribute metric setting.
        /// </summary>
        public readonly int Metric;
        /// <summary>
        /// Metric type.
        /// </summary>
        public readonly string MetricType;
        /// <summary>
        /// Passive interface name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Route map name.
        /// </summary>
        public readonly string Routemap;
        /// <summary>
        /// status
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetOspf6RedistributeResult(
            int metric,

            string metricType,

            string name,

            string routemap,

            string status)
        {
            Metric = metric;
            MetricType = metricType;
            Name = name;
            Routemap = routemap;
            Status = status;
        }
    }
}
