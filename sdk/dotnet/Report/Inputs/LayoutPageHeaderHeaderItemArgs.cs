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

    public sealed class LayoutPageHeaderHeaderItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Report item text content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Report item ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Report item image file name.
        /// </summary>
        [Input("imgSrc")]
        public Input<string>? ImgSrc { get; set; }

        /// <summary>
        /// Report item style.
        /// </summary>
        [Input("style")]
        public Input<string>? Style { get; set; }

        /// <summary>
        /// Report item type. Valid values: `text`, `image`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public LayoutPageHeaderHeaderItemArgs()
        {
        }
        public static new LayoutPageHeaderHeaderItemArgs Empty => new LayoutPageHeaderHeaderItemArgs();
    }
}
