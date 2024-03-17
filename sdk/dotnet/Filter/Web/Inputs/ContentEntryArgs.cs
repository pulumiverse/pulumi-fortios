// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Web.Inputs
{

    public sealed class ContentEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Block or exempt word when a match is found. Valid values: `block`, `exempt`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Language of banned word. Valid values: `western`, `simch`, `trach`, `japanese`, `korean`, `french`, `thai`, `spanish`, `cyrillic`.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// Banned word.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Banned word pattern type: wildcard pattern or Perl regular expression. Valid values: `wildcard`, `regexp`.
        /// </summary>
        [Input("patternType")]
        public Input<string>? PatternType { get; set; }

        /// <summary>
        /// Score, to be applied every time the word appears on a web page (0 - 4294967295, default = 10).
        /// </summary>
        [Input("score")]
        public Input<int>? Score { get; set; }

        /// <summary>
        /// Enable/disable banned word. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ContentEntryArgs()
        {
        }
        public static new ContentEntryArgs Empty => new ContentEntryArgs();
    }
}
