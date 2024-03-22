// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    /// <summary>
    /// Configure an aggregate of IPsec tunnels.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname1Phase1interface = new Fortios.Vpn.Ipsec.Phase1interface("trname1Phase1interface", new()
    ///     {
    ///         AcctVerify = "disable",
    ///         AddGwRoute = "disable",
    ///         AddRoute = "enable",
    ///         AssignIp = "enable",
    ///         AssignIpFrom = "range",
    ///         Authmethod = "psk",
    ///         AutoDiscoveryForwarder = "disable",
    ///         AutoDiscoveryPsk = "disable",
    ///         AutoDiscoveryReceiver = "disable",
    ///         AutoDiscoverySender = "disable",
    ///         AutoNegotiate = "enable",
    ///         CertIdValidation = "enable",
    ///         ChildlessIke = "disable",
    ///         ClientAutoNegotiate = "disable",
    ///         ClientKeepAlive = "disable",
    ///         DefaultGw = "0.0.0.0",
    ///         DefaultGwPriority = 0,
    ///         Dhgrp = "14 5",
    ///         DigitalSignatureAuth = "disable",
    ///         Distance = 15,
    ///         DnsMode = "manual",
    ///         Dpd = "on-demand",
    ///         DpdRetrycount = 3,
    ///         DpdRetryinterval = "20",
    ///         Eap = "disable",
    ///         EapIdentity = "use-id-payload",
    ///         EncapLocalGw4 = "0.0.0.0",
    ///         EncapLocalGw6 = "::",
    ///         EncapRemoteGw4 = "0.0.0.0",
    ///         EncapRemoteGw6 = "::",
    ///         Encapsulation = "none",
    ///         EncapsulationAddress = "ike",
    ///         EnforceUniqueId = "disable",
    ///         ExchangeInterfaceIp = "disable",
    ///         ExchangeIpAddr4 = "0.0.0.0",
    ///         ExchangeIpAddr6 = "::",
    ///         ForticlientEnforcement = "disable",
    ///         Fragmentation = "enable",
    ///         FragmentationMtu = 1200,
    ///         GroupAuthentication = "disable",
    ///         HaSyncEspSeqno = "enable",
    ///         IdleTimeout = "disable",
    ///         IdleTimeoutinterval = 15,
    ///         IkeVersion = "1",
    ///         IncludeLocalLan = "disable",
    ///         Interface = "port3",
    ///         IpVersion = "4",
    ///         Ipv4DnsServer1 = "0.0.0.0",
    ///         Ipv4DnsServer2 = "0.0.0.0",
    ///         Ipv4DnsServer3 = "0.0.0.0",
    ///         Ipv4EndIp = "0.0.0.0",
    ///         Ipv4Netmask = "255.255.255.255",
    ///         Ipv4StartIp = "0.0.0.0",
    ///         Ipv4WinsServer1 = "0.0.0.0",
    ///         Ipv4WinsServer2 = "0.0.0.0",
    ///         Ipv6DnsServer1 = "::",
    ///         Ipv6DnsServer2 = "::",
    ///         Ipv6DnsServer3 = "::",
    ///         Ipv6EndIp = "::",
    ///         Ipv6Prefix = 128,
    ///         Ipv6StartIp = "::",
    ///         Keepalive = 10,
    ///         Keylife = 86400,
    ///         LocalGw = "0.0.0.0",
    ///         LocalGw6 = "::",
    ///         LocalidType = "auto",
    ///         MeshSelectorType = "disable",
    ///         Mode = "main",
    ///         ModeCfg = "disable",
    ///         MonitorHoldDownDelay = 0,
    ///         MonitorHoldDownTime = "00:00",
    ///         MonitorHoldDownType = "immediate",
    ///         MonitorHoldDownWeekday = "sunday",
    ///         Nattraversal = "enable",
    ///         NegotiateTimeout = 30,
    ///         NetDevice = "disable",
    ///         PassiveMode = "disable",
    ///         Peertype = "any",
    ///         Ppk = "disable",
    ///         Priority = 0,
    ///         Proposal = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1",
    ///         Psksecret = "eweeeeeeeecee",
    ///         Reauth = "disable",
    ///         Rekey = "enable",
    ///         RemoteGw = "2.2.2.2",
    ///         RemoteGw6 = "::",
    ///         RsaSignatureFormat = "pkcs1",
    ///         SavePassword = "disable",
    ///         SendCertChain = "enable",
    ///         SignatureHashAlg = "sha2-512 sha2-384 sha2-256 sha1",
    ///         SuiteB = "disable",
    ///         TunnelSearch = "selectors",
    ///         Type = "static",
    ///         UnitySupport = "enable",
    ///         WizardType = "custom",
    ///         Xauthtype = "disable",
    ///     });
    /// 
    ///     var trname1Phase2interface = new Fortios.Vpn.Ipsec.Phase2interface("trname1Phase2interface", new()
    ///     {
    ///         AddRoute = "phase1",
    ///         AutoDiscoveryForwarder = "phase1",
    ///         AutoDiscoverySender = "phase1",
    ///         AutoNegotiate = "disable",
    ///         DhcpIpsec = "disable",
    ///         Dhgrp = "14 5",
    ///         DstAddrType = "subnet",
    ///         DstEndIp6 = "::",
    ///         DstPort = 0,
    ///         DstSubnet = "0.0.0.0 0.0.0.0",
    ///         Encapsulation = "tunnel-mode",
    ///         Keepalive = "disable",
    ///         KeylifeType = "seconds",
    ///         Keylifekbs = 5120,
    ///         Keylifeseconds = 43200,
    ///         L2tp = "disable",
    ///         Pfs = "enable",
    ///         Phase1name = trname1Phase1interface.Name,
    ///         Proposal = "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305",
    ///         Protocol = 0,
    ///         Replay = "enable",
    ///         RouteOverlap = "use-new",
    ///         SingleSource = "disable",
    ///         SrcAddrType = "subnet",
    ///         SrcEndIp6 = "::",
    ///         SrcPort = 0,
    ///         SrcSubnet = "0.0.0.0 0.0.0.0",
    ///     });
    /// 
    ///     var trname = new Fortios.System.Ipsecaggregate("trname", new()
    ///     {
    ///         Algorithm = "round-robin",
    ///         Members = new[]
    ///         {
    ///             new Fortios.System.Inputs.IpsecaggregateMemberArgs
    ///             {
    ///                 TunnelName = trname1Phase1interface.Name,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// System IpsecAggregate can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/ipsecaggregate:Ipsecaggregate labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/ipsecaggregate:Ipsecaggregate labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/ipsecaggregate:Ipsecaggregate")]
    public partial class Ipsecaggregate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Frame distribution algorithm.
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Member tunnels of the aggregate. The structure of `member` block is documented below.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.IpsecaggregateMember>> Members { get; private set; } = null!;

        /// <summary>
        /// IPsec aggregate name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ipsecaggregate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipsecaggregate(string name, IpsecaggregateArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/ipsecaggregate:Ipsecaggregate", name, args ?? new IpsecaggregateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipsecaggregate(string name, Input<string> id, IpsecaggregateState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/ipsecaggregate:Ipsecaggregate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipsecaggregate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipsecaggregate Get(string name, Input<string> id, IpsecaggregateState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipsecaggregate(name, id, state, options);
        }
    }

    public sealed class IpsecaggregateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Frame distribution algorithm.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        [Input("members", required: true)]
        private InputList<Inputs.IpsecaggregateMemberArgs>? _members;

        /// <summary>
        /// Member tunnels of the aggregate. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.IpsecaggregateMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.IpsecaggregateMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// IPsec aggregate name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IpsecaggregateArgs()
        {
        }
        public static new IpsecaggregateArgs Empty => new IpsecaggregateArgs();
    }

    public sealed class IpsecaggregateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Frame distribution algorithm.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        [Input("members")]
        private InputList<Inputs.IpsecaggregateMemberGetArgs>? _members;

        /// <summary>
        /// Member tunnels of the aggregate. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.IpsecaggregateMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.IpsecaggregateMemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// IPsec aggregate name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IpsecaggregateState()
        {
        }
        public static new IpsecaggregateState Empty => new IpsecaggregateState();
    }
}
