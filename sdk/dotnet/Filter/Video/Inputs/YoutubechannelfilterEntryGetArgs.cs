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

    public sealed class YoutubechannelfilterEntryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// YouTube channel filter action. Valid values: `allow`, `monitor`, `block`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Channel ID.
        /// </summary>
        [Input("channelId")]
        public Input<string>? ChannelId { get; set; }

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

        public YoutubechannelfilterEntryGetArgs()
        {
        }
        public static new YoutubechannelfilterEntryGetArgs Empty => new YoutubechannelfilterEntryGetArgs();
    }
}
