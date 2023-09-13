// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Ips.Inputs
{

    public sealed class SensorEntryExemptIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Destination IP address and netmask.
        /// </summary>
        [Input("dstIp")]
        public Input<string>? DstIp { get; set; }

        /// <summary>
        /// Exempt IP ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Source IP address and netmask.
        /// </summary>
        [Input("srcIp")]
        public Input<string>? SrcIp { get; set; }

        public SensorEntryExemptIpArgs()
        {
        }
        public static new SensorEntryExemptIpArgs Empty => new SensorEntryExemptIpArgs();
    }
}
