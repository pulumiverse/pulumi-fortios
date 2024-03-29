// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Video.Outputs
{

    [OutputType]
    public sealed class ProfileFortiguardCategory
    {
        /// <summary>
        /// Configure VideoFilter FortiGuard category. The structure of `filters` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProfileFortiguardCategoryFilter> Filters;

        [OutputConstructor]
        private ProfileFortiguardCategory(ImmutableArray<Outputs.ProfileFortiguardCategoryFilter> filters)
        {
            Filters = filters;
        }
    }
}
