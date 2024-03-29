// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Inputs
{

    public sealed class ManagedswitchDhcpSnoopingStaticClientGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client static IP address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Client MAC address.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// Client name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// VLAN name.
        /// </summary>
        [Input("vlan")]
        public Input<string>? Vlan { get; set; }

        public ManagedswitchDhcpSnoopingStaticClientGetArgs()
        {
        }
        public static new ManagedswitchDhcpSnoopingStaticClientGetArgs Empty => new ManagedswitchDhcpSnoopingStaticClientGetArgs();
    }
}
