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

    public sealed class ProfileFortiguardCategoryArgs : global::Pulumi.ResourceArgs
    {
        [Input("filters")]
        private InputList<Inputs.ProfileFortiguardCategoryFilterArgs>? _filters;

        /// <summary>
        /// Configure VideoFilter FortiGuard category. The structure of `filters` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileFortiguardCategoryFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.ProfileFortiguardCategoryFilterArgs>());
            set => _filters = value;
        }

        public ProfileFortiguardCategoryArgs()
        {
        }
        public static new ProfileFortiguardCategoryArgs Empty => new ProfileFortiguardCategoryArgs();
    }
}
