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

    public sealed class SystemGeoipoverrideIp6RangeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Final IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
        /// 
        /// The `ip6_range` block supports:
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// ID number for individual entry in the IP-Range table.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Starting IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        public SystemGeoipoverrideIp6RangeArgs()
        {
        }
        public static new SystemGeoipoverrideIp6RangeArgs Empty => new SystemGeoipoverrideIp6RangeArgs();
    }
}