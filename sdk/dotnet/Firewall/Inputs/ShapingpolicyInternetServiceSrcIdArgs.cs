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

    public sealed class ShapingpolicyInternetServiceSrcIdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internet Service ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public ShapingpolicyInternetServiceSrcIdArgs()
        {
        }
        public static new ShapingpolicyInternetServiceSrcIdArgs Empty => new ShapingpolicyInternetServiceSrcIdArgs();
    }
}
