// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Outputs
{

    [OutputType]
    public sealed class ManagedswitchDhcpSnoopingStaticClient
    {
        /// <summary>
        /// Client static IP address.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Client MAC address.
        /// </summary>
        public readonly string? Mac;
        /// <summary>
        /// Client name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// VLAN name.
        /// </summary>
        public readonly string? Vlan;

        [OutputConstructor]
        private ManagedswitchDhcpSnoopingStaticClient(
            string? ip,

            string? mac,

            string? name,

            string? port,

            string? vlan)
        {
            Ip = ip;
            Mac = mac;
            Name = name;
            Port = port;
            Vlan = vlan;
        }
    }
}
