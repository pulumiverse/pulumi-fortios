// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class SystemVirtualwanlinkServiceInternetServiceGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internet Service group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SystemVirtualwanlinkServiceInternetServiceGroupArgs()
        {
        }
        public static new SystemVirtualwanlinkServiceInternetServiceGroupArgs Empty => new SystemVirtualwanlinkServiceInternetServiceGroupArgs();
    }
}
