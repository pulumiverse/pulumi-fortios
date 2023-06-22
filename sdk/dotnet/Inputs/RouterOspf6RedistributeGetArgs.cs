// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class RouterOspf6RedistributeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Redistribute metric setting.
        /// </summary>
        [Input("metric")]
        public Input<int>? Metric { get; set; }

        /// <summary>
        /// Metric type. Valid values: `1`, `2`.
        /// </summary>
        [Input("metricType")]
        public Input<string>? MetricType { get; set; }

        /// <summary>
        /// Redistribute name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Route map name.
        /// </summary>
        [Input("routemap")]
        public Input<string>? Routemap { get; set; }

        /// <summary>
        /// status Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public RouterOspf6RedistributeGetArgs()
        {
        }
        public static new RouterOspf6RedistributeGetArgs Empty => new RouterOspf6RedistributeGetArgs();
    }
}
