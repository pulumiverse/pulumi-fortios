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
        /// <summary>
        /// A.B.C.D, in IPv4 address format.
        /// </summary>
        [Input("areaId")]
        public Input<string>? AreaId { get; set; }

        /// <summary>
        /// Authentication mode. Valid values: `none`, `ah`, `esp`.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bfd")]
        public Input<string>? Bfd { get; set; }

        /// <summary>
        /// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
        /// </summary>
        [Input("cost")]
        public Input<int>? Cost { get; set; }

        /// <summary>
        /// Dead interval.
        /// </summary>
        [Input("deadInterval")]
        public Input<int>? DeadInterval { get; set; }

        /// <summary>
        /// Hello interval.
        /// </summary>
        [Input("helloInterval")]
        public Input<int>? HelloInterval { get; set; }

        /// <summary>
        /// Configuration interface name.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
        /// </summary>
        [Input("ipsecAuthAlg")]
        public Input<string>? IpsecAuthAlg { get; set; }

        /// <summary>
        /// Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
        /// </summary>
        [Input("ipsecEncAlg")]
        public Input<string>? IpsecEncAlg { get; set; }

        [Input("ipsecKeys")]
        private InputList<Inputs.Ospf6Ospf6InterfaceIpsecKeyArgs>? _ipsecKeys;

        /// <summary>
        /// IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6Ospf6InterfaceIpsecKeyArgs> IpsecKeys
        {
            get => _ipsecKeys ?? (_ipsecKeys = new InputList<Inputs.Ospf6Ospf6InterfaceIpsecKeyArgs>());
            set => _ipsecKeys = value;
        }

        /// <summary>
        /// Key roll-over interval.
        /// </summary>
        [Input("keyRolloverInterval")]
        public Input<int>? KeyRolloverInterval { get; set; }

        /// <summary>
        /// MTU for OSPFv3 packets.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// Enable/disable ignoring MTU field in DBD packets. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("mtuIgnore")]
        public Input<string>? MtuIgnore { get; set; }

        /// <summary>
        /// Virtual link entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("neighbors")]
        private InputList<Inputs.Ospf6Ospf6InterfaceNeighborArgs>? _neighbors;

        /// <summary>
        /// OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6Ospf6InterfaceNeighborArgs> Neighbors
        {
            get => _neighbors ?? (_neighbors = new InputList<Inputs.Ospf6Ospf6InterfaceNeighborArgs>());
            set => _neighbors = value;
        }

        /// <summary>
        /// Network type. Valid values: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// priority
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Retransmit interval.
        /// </summary>
        [Input("retransmitInterval")]
        public Input<int>? RetransmitInterval { get; set; }

        /// <summary>
        /// Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Transmit delay.
        /// </summary>
        [Input("transmitDelay")]
        public Input<int>? TransmitDelay { get; set; }

        public Ospf6Ospf6InterfaceArgs()
        {
        }
        public static new Ospf6Ospf6InterfaceArgs Empty => new Ospf6Ospf6InterfaceArgs();
    }
}
