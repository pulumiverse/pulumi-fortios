// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Web.Outputs
{

    [OutputType]
    public sealed class ProfileYoutubeChannelFilter
    {
        /// <summary>
        /// YouTube channel ID to be filtered.
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
        private ProfileYoutubeChannelFilter(
            string? channelId,

            string? comment,

            int? id)
        {
            ChannelId = channelId;
            Comment = comment;
            Id = id;
        }
    }
}
