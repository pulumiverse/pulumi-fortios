// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Report.Outputs
{

    [OutputType]
    public sealed class ChartCategorySeries
    {
        /// <summary>
        /// Category series value expression.
        /// </summary>
        public readonly string? Databind;
        /// <summary>
        /// Font size of category-series title.
        /// </summary>
        public readonly int? FontSize;

        [OutputConstructor]
        private ChartCategorySeries(
            string? databind,

            int? fontSize)
        {
            Databind = databind;
            FontSize = fontSize;
        }
    }
}
