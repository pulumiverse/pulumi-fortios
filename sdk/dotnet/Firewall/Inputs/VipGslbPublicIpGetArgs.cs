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

    public sealed class VipGslbPublicIpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Index of this public IP setting.
        /// </summary>
        [Input("index")]
        public Input<int>? Index { get; set; }

        /// <summary>
        /// The publicly accessible IP address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public VipGslbPublicIpGetArgs()
        {
        }
        public static new VipGslbPublicIpGetArgs Empty => new VipGslbPublicIpGetArgs();
    }
}
