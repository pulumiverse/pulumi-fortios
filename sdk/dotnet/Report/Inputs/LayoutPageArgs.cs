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

    public sealed class LayoutPageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Report page auto column break before heading. Valid values: `heading1`, `heading2`, `heading3`.
        /// </summary>
        [Input("columnBreakBefore")]
        public Input<string>? ColumnBreakBefore { get; set; }

        /// <summary>
        /// Configure report page footer. The structure of `footer` block is documented below.
        /// </summary>
        [Input("footer")]
        public Input<Inputs.LayoutPageFooterArgs>? Footer { get; set; }

        /// <summary>
        /// Configure report page header. The structure of `header` block is documented below.
        /// </summary>
        [Input("header")]
        public Input<Inputs.LayoutPageHeaderArgs>? Header { get; set; }

        /// <summary>
        /// Report page options. Valid values: `header-on-first-page`, `footer-on-first-page`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Report page auto page break before heading. Valid values: `heading1`, `heading2`, `heading3`.
        /// </summary>
        [Input("pageBreakBefore")]
        public Input<string>? PageBreakBefore { get; set; }

        /// <summary>
        /// Report page paper. Valid values: `a4`, `letter`.
        /// </summary>
        [Input("paper")]
        public Input<string>? Paper { get; set; }

        public LayoutPageArgs()
        {
        }
        public static new LayoutPageArgs Empty => new LayoutPageArgs();
    }
}