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

    public sealed class ProfileYoutubeChannelFilterGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// YouTube channel ID to be filtered.
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

        public ProfileYoutubeChannelFilterGetArgs()
        {
        }
        public static new ProfileYoutubeChannelFilterGetArgs Empty => new ProfileYoutubeChannelFilterGetArgs();
    }
}
