// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Video.Inputs
{

    public sealed class ProfileFilterGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VideoFilter action. Valid values: `allow`, `monitor`, `block`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// FortiGuard category ID.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Channel ID.
        /// </summary>
        [Input("channel")]
        public Input<string>? Channel { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Video filter keyword ID.
        /// </summary>
        [Input("keyword")]
        public Input<int>? Keyword { get; set; }

        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Filter type. Valid values: `category`, `channel`, `title`, `description`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ProfileFilterGetArgs()
        {
        }
        public static new ProfileFilterGetArgs Empty => new ProfileFilterGetArgs();
    }
}
