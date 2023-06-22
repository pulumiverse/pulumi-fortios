// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class VideofilterYoutubechannelfilterEntry
    {
        /// <summary>
        /// YouTube channel filter action. Valid values: `allow`, `monitor`, `block`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Channel ID.
        /// </summary>
        public readonly string? ChannelId;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string? Comment;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;

        [OutputConstructor]
        private VideofilterYoutubechannelfilterEntry(
            string? action,

            string? channelId,

            string? comment,

            int? id)
        {
            Action = action;
            ChannelId = channelId;
            Comment = comment;
            Id = id;
        }
    }
}
