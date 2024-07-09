// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class Ospf6Ospf6InterfaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("areaId")]
        public Input<string>? AreaId { get; set; }

        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bfd")]
        public Input<string>? Bfd { get; set; }

        [Input("cost")]
        public Input<int>? Cost { get; set; }

        [Input("deadInterval")]
        public Input<int>? DeadInterval { get; set; }

        [Input("helloInterval")]
        public Input<int>? HelloInterval { get; set; }

        [Input("interface")]
        public Input<string>? Interface { get; set; }

        [Input("ipsecAuthAlg")]
        public Input<string>? IpsecAuthAlg { get; set; }

        [Input("ipsecEncAlg")]
        public Input<string>? IpsecEncAlg { get; set; }

        [Input("ipsecKeys")]
        private InputList<Inputs.Ospf6Ospf6InterfaceIpsecKeyArgs>? _ipsecKeys;
        public InputList<Inputs.Ospf6Ospf6InterfaceIpsecKeyArgs> IpsecKeys
        {
            get => _ipsecKeys ?? (_ipsecKeys = new InputList<Inputs.Ospf6Ospf6InterfaceIpsecKeyArgs>());
            set => _ipsecKeys = value;
        }

        [Input("keyRolloverInterval")]
        public Input<int>? KeyRolloverInterval { get; set; }

        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        [Input("mtuIgnore")]
        public Input<string>? MtuIgnore { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("neighbors")]
        private InputList<Inputs.Ospf6Ospf6InterfaceNeighborArgs>? _neighbors;
        public InputList<Inputs.Ospf6Ospf6InterfaceNeighborArgs> Neighbors
        {
            get => _neighbors ?? (_neighbors = new InputList<Inputs.Ospf6Ospf6InterfaceNeighborArgs>());
            set => _neighbors = value;
        }

        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("retransmitInterval")]
        public Input<int>? RetransmitInterval { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("transmitDelay")]
        public Input<int>? TransmitDelay { get; set; }

        public Ospf6Ospf6InterfaceArgs()
        {
        }
        public static new Ospf6Ospf6InterfaceArgs Empty => new Ospf6Ospf6InterfaceArgs();
    }
}
