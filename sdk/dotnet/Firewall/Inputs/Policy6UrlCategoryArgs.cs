// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class Policy6UrlCategoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URL category ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public Policy6UrlCategoryArgs()
        {
        }
        public static new Policy6UrlCategoryArgs Empty => new Policy6UrlCategoryArgs();
    }
}
