// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class MobiletunnelNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Network entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Select the associated interface name from available options.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Class IP and Netmask with correction (Format:xxx.xxx.xxx.xxx xxx.xxx.xxx.xxx or xxx.xxx.xxx.xxx/x).
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public MobiletunnelNetworkGetArgs()
        {
        }
        public static new MobiletunnelNetworkGetArgs Empty => new MobiletunnelNetworkGetArgs();
    }
}
