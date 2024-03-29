// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Log.Webtrends.Inputs
{

    public sealed class FilterFreeStyleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Log category.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Free style filter string.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        /// </summary>
        [Input("filterType")]
        public Input<string>? FilterType { get; set; }

        /// <summary>
        /// Entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public FilterFreeStyleArgs()
        {
        }
        public static new FilterFreeStyleArgs Empty => new FilterFreeStyleArgs();
    }
}
