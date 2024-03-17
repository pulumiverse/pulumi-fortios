// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User.Outputs
{

    [OutputType]
    public sealed class GroupMatch
    {
        /// <summary>
        /// Name of matching group on remote auththentication server.
        /// </summary>
        public readonly string? GroupName;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Name of remote auth server.
        /// </summary>
        public readonly string? ServerName;

        [OutputConstructor]
        private GroupMatch(
            string? groupName,

            int? id,

            string? serverName)
        {
            GroupName = groupName;
            Id = id;
            ServerName = serverName;
        }
    }
}
