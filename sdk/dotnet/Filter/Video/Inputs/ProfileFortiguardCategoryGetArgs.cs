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

    public sealed class ProfileFortiguardCategoryGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("filters")]
        private InputList<Inputs.ProfileFortiguardCategoryFilterGetArgs>? _filters;

        /// <summary>
        /// Configure VideoFilter FortiGuard category. The structure of `filters` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileFortiguardCategoryFilterGetArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.ProfileFortiguardCategoryFilterGetArgs>());
            set => _filters = value;
        }

        public ProfileFortiguardCategoryGetArgs()
        {
        }
        public static new ProfileFortiguardCategoryGetArgs Empty => new ProfileFortiguardCategoryGetArgs();
    }
}
