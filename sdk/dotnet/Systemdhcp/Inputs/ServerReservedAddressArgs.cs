// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Systemdhcp.Inputs
{

    public sealed class ServerReservedAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Options for the DHCP server to configure the client with the reserved MAC address. Valid values: `assign`, `block`, `reserved`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Option 82 circuit-ID of the client that will get the reserved IP address.
        /// </summary>
        [Input("circuitId")]
        public Input<string>? CircuitId { get; set; }

        /// <summary>
        /// DHCP option type. Valid values: `hex`, `string`.
        /// </summary>
        [Input("circuitIdType")]
        public Input<string>? CircuitIdType { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// IP address to be reserved for the MAC address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// MAC address of the client that will get the reserved IP address.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// Option 82 remote-ID of the client that will get the reserved IP address.
        /// </summary>
        [Input("remoteId")]
        public Input<string>? RemoteId { get; set; }

        /// <summary>
        /// DHCP option type. Valid values: `hex`, `string`.
        /// </summary>
        [Input("remoteIdType")]
        public Input<string>? RemoteIdType { get; set; }

        /// <summary>
        /// DHCP reserved-address type. Valid values: `mac`, `option82`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ServerReservedAddressArgs()
        {
        }
        public static new ServerReservedAddressArgs Empty => new ServerReservedAddressArgs();
    }
}
