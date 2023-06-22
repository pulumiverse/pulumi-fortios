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

    public sealed class Systemdhcp6ServerIpRangeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// End of IP range.
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Start of IP range.
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        public Systemdhcp6ServerIpRangeArgs()
        {
        }
        public static new Systemdhcp6ServerIpRangeArgs Empty => new Systemdhcp6ServerIpRangeArgs();
    }
}
