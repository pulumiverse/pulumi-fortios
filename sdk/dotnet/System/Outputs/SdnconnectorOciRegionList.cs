// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class SdnconnectorOciRegionList
    {
        /// <summary>
        /// OCI region.
        /// </summary>
        public readonly string? Region;

        [OutputConstructor]
        private SdnconnectorOciRegionList(string? region)
        {
            Region = region;
        }
    }
}
