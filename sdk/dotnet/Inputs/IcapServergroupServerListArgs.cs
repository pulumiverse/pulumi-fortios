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

    public sealed class IcapServergroupServerListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ICAP server name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optionally assign a weight of the ICAP server for weighted load balancing (1 - 100, default = 10)
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public IcapServergroupServerListArgs()
        {
        }
        public static new IcapServergroupServerListArgs Empty => new IcapServergroupServerListArgs();
    }
}
