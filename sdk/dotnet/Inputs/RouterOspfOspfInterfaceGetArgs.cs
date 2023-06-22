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

    public sealed class RouterOspfOspfInterfaceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication type.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        [Input("authenticationKey")]
        private Input<string>? _authenticationKey;

        /// <summary>
        /// Authentication key.
        /// </summary>
        public Input<string>? AuthenticationKey
        {
            get => _authenticationKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authenticationKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
        /// </summary>
        [Input("bfd")]
        public Input<string>? Bfd { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
        /// </summary>
        [Input("cost")]
        public Input<int>? Cost { get; set; }

        /// <summary>
        /// Enable/disable control of flooding out LSAs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("databaseFilterOut")]
        public Input<string>? DatabaseFilterOut { get; set; }

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
        /// Number of hello packets within dead interval.
        /// </summary>
        [Input("helloMultiplier")]
        public Input<int>? HelloMultiplier { get; set; }

        /// <summary>
        /// Configuration interface name.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// IP address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Message-digest key-chain name.
        /// </summary>
        [Input("keychain")]
        public Input<string>? Keychain { get; set; }

        [Input("md5Key")]
        private Input<string>? _md5Key;

        /// <summary>
        /// MD5 key.
        /// </summary>
        public Input<string>? Md5Key
        {
            get => _md5Key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _md5Key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Authentication MD5 key-chain name.
        /// </summary>
        [Input("md5Keychain")]
        public Input<string>? Md5Keychain { get; set; }

        [Input("md5Keys")]
        private InputList<Inputs.RouterOspfOspfInterfaceMd5KeyGetArgs>? _md5Keys;

        /// <summary>
        /// MD5 key. The structure of `md5_keys` block is documented below.
        /// 
        /// The `md5_keys` block supports:
        /// </summary>
        public InputList<Inputs.RouterOspfOspfInterfaceMd5KeyGetArgs> Md5Keys
        {
            get => _md5Keys ?? (_md5Keys = new InputList<Inputs.RouterOspfOspfInterfaceMd5KeyGetArgs>());
            set => _md5Keys = value;
        }

        /// <summary>
        /// MTU for database description packets.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// Enable/disable ignore MTU. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("mtuIgnore")]
        public Input<string>? MtuIgnore { get; set; }

        /// <summary>
        /// Interface entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network type. Valid values: `broadcast`, `non-broadcast`, `point-to-point`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// Prefix length.
        /// </summary>
        [Input("prefixLength")]
        public Input<int>? PrefixLength { get; set; }

        /// <summary>
        /// Priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Graceful restart neighbor resynchronization timeout.
        /// </summary>
        [Input("resyncTimeout")]
        public Input<int>? ResyncTimeout { get; set; }

        /// <summary>
        /// Retransmit interval.
        /// </summary>
        [Input("retransmitInterval")]
        public Input<int>? RetransmitInterval { get; set; }

        /// <summary>
        /// Enable/disable status. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Transmit delay.
        /// </summary>
        [Input("transmitDelay")]
        public Input<int>? TransmitDelay { get; set; }

        public RouterOspfOspfInterfaceGetArgs()
        {
        }
        public static new RouterOspfOspfInterfaceGetArgs Empty => new RouterOspfOspfInterfaceGetArgs();
    }
}
