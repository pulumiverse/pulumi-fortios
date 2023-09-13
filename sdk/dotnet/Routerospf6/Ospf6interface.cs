// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Routerospf6
{
    /// <summary>
    /// OSPF6 interface configuration.
    /// 
    /// &gt; The provider supports the definition of Ospf6-Interface in Router Ospf6 `fortios.router.Ospf6`, and also allows the definition of separate Ospf6-Interface resources `fortios.routerospf6.Ospf6interface`, but do not use a `fortios.router.Ospf6` with in-line Ospf6-Interface in conjunction with any `fortios.routerospf6.Ospf6interface` resources, otherwise conflicts and overwrite will occur.
    /// 
    /// ## Import
    /// 
    /// Routerospf6 Ospf6Interface can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:routerospf6/ospf6interface:Ospf6interface labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:routerospf6/ospf6interface:Ospf6interface labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:routerospf6/ospf6interface:Ospf6interface")]
    public partial class Ospf6interface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A.B.C.D, in IPv4 address format.
        /// </summary>
        [Output("areaId")]
        public Output<string> AreaId { get; private set; } = null!;

        /// <summary>
        /// Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
        /// </summary>
        [Output("authentication")]
        public Output<string> Authentication { get; private set; } = null!;

        /// <summary>
        /// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
        /// </summary>
        [Output("bfd")]
        public Output<string> Bfd { get; private set; } = null!;

        /// <summary>
        /// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
        /// </summary>
        [Output("cost")]
        public Output<int> Cost { get; private set; } = null!;

        /// <summary>
        /// Dead interval.
        /// </summary>
        [Output("deadInterval")]
        public Output<int> DeadInterval { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Hello interval.
        /// </summary>
        [Output("helloInterval")]
        public Output<int> HelloInterval { get; private set; } = null!;

        /// <summary>
        /// Configuration interface name.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
        /// </summary>
        [Output("ipsecAuthAlg")]
        public Output<string> IpsecAuthAlg { get; private set; } = null!;

        /// <summary>
        /// Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
        /// </summary>
        [Output("ipsecEncAlg")]
        public Output<string> IpsecEncAlg { get; private set; } = null!;

        /// <summary>
        /// IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.
        /// </summary>
        [Output("ipsecKeys")]
        public Output<ImmutableArray<Outputs.Ospf6interfaceIpsecKey>> IpsecKeys { get; private set; } = null!;

        /// <summary>
        /// Key roll-over interval.
        /// </summary>
        [Output("keyRolloverInterval")]
        public Output<int> KeyRolloverInterval { get; private set; } = null!;

        /// <summary>
        /// MTU for OSPFv3 packets.
        /// </summary>
        [Output("mtu")]
        public Output<int> Mtu { get; private set; } = null!;

        /// <summary>
        /// Enable/disable ignoring MTU field in DBD packets. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("mtuIgnore")]
        public Output<string> MtuIgnore { get; private set; } = null!;

        /// <summary>
        /// Interface entry name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
        /// </summary>
        [Output("neighbors")]
        public Output<ImmutableArray<Outputs.Ospf6interfaceNeighbor>> Neighbors { get; private set; } = null!;

        /// <summary>
        /// Network type. Valid values: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
        /// </summary>
        [Output("networkType")]
        public Output<string> NetworkType { get; private set; } = null!;

        /// <summary>
        /// priority
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Retransmit interval.
        /// </summary>
        [Output("retransmitInterval")]
        public Output<int> RetransmitInterval { get; private set; } = null!;

        /// <summary>
        /// Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Transmit delay.
        /// </summary>
        [Output("transmitDelay")]
        public Output<int> TransmitDelay { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ospf6interface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ospf6interface(string name, Ospf6interfaceArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:routerospf6/ospf6interface:Ospf6interface", name, args ?? new Ospf6interfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ospf6interface(string name, Input<string> id, Ospf6interfaceState? state = null, CustomResourceOptions? options = null)
            : base("fortios:routerospf6/ospf6interface:Ospf6interface", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Ospf6interface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ospf6interface Get(string name, Input<string> id, Ospf6interfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Ospf6interface(name, id, state, options);
        }
    }

    public sealed class Ospf6interfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A.B.C.D, in IPv4 address format.
        /// </summary>
        [Input("areaId")]
        public Input<string>? AreaId { get; set; }

        /// <summary>
        /// Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
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
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

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
        private InputList<Inputs.Ospf6interfaceIpsecKeyArgs>? _ipsecKeys;

        /// <summary>
        /// IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6interfaceIpsecKeyArgs> IpsecKeys
        {
            get => _ipsecKeys ?? (_ipsecKeys = new InputList<Inputs.Ospf6interfaceIpsecKeyArgs>());
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
        /// Interface entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("neighbors")]
        private InputList<Inputs.Ospf6interfaceNeighborArgs>? _neighbors;

        /// <summary>
        /// OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6interfaceNeighborArgs> Neighbors
        {
            get => _neighbors ?? (_neighbors = new InputList<Inputs.Ospf6interfaceNeighborArgs>());
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

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Ospf6interfaceArgs()
        {
        }
        public static new Ospf6interfaceArgs Empty => new Ospf6interfaceArgs();
    }

    public sealed class Ospf6interfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A.B.C.D, in IPv4 address format.
        /// </summary>
        [Input("areaId")]
        public Input<string>? AreaId { get; set; }

        /// <summary>
        /// Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
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
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

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
        private InputList<Inputs.Ospf6interfaceIpsecKeyGetArgs>? _ipsecKeys;

        /// <summary>
        /// IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6interfaceIpsecKeyGetArgs> IpsecKeys
        {
            get => _ipsecKeys ?? (_ipsecKeys = new InputList<Inputs.Ospf6interfaceIpsecKeyGetArgs>());
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
        /// Interface entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("neighbors")]
        private InputList<Inputs.Ospf6interfaceNeighborGetArgs>? _neighbors;

        /// <summary>
        /// OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6interfaceNeighborGetArgs> Neighbors
        {
            get => _neighbors ?? (_neighbors = new InputList<Inputs.Ospf6interfaceNeighborGetArgs>());
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

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Ospf6interfaceState()
        {
        }
        public static new Ospf6interfaceState Empty => new Ospf6interfaceState();
    }
}
