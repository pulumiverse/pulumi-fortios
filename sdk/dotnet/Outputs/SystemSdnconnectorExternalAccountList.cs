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
    public sealed class SystemSdnconnectorExternalAccountList
    {
        /// <summary>
        /// AWS external ID.
        /// </summary>
        public readonly string? ExternalId;
        /// <summary>
        /// AWS region name list. The structure of `region_list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.SystemSdnconnectorExternalAccountListRegionList> RegionLists;
        /// <summary>
        /// AWS role ARN to assume.
        /// </summary>
        public readonly string? RoleArn;

        [OutputConstructor]
        private SystemSdnconnectorExternalAccountList(
            string? externalId,

            ImmutableArray<Outputs.SystemSdnconnectorExternalAccountListRegionList> regionLists,

            string? roleArn)
        {
            ExternalId = externalId;
            RegionLists = regionLists;
            RoleArn = roleArn;
        }
    }
}
