// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Report.Inputs
{

    public sealed class ChartValueSeriesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Value series value expression.
        /// </summary>
        [Input("databind")]
        public Input<string>? Databind { get; set; }

        public ChartValueSeriesArgs()
        {
        }
        public static new ChartValueSeriesArgs Empty => new ChartValueSeriesArgs();
    }
}