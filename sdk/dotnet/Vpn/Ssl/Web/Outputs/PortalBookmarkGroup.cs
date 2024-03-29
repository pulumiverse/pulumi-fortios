// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ssl.Web.Outputs
{

    [OutputType]
    public sealed class PortalBookmarkGroup
    {
        /// <summary>
        /// Bookmark table. The structure of `bookmarks` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.PortalBookmarkGroupBookmark> Bookmarks;
        /// <summary>
        /// Bookmark group name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private PortalBookmarkGroup(
            ImmutableArray<Outputs.PortalBookmarkGroupBookmark> bookmarks,

            string? name)
        {
            Bookmarks = bookmarks;
            Name = name;
        }
    }
}
