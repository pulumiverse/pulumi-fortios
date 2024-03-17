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
    /// Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.
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
    ///     var trname = new Fortios.System.Mobiletunnel("trname", new()
    ///     {
    ///         HashAlgorithm = "hmac-md5",
    ///         HomeAddress = "0.0.0.0",
    ///         HomeAgent = "1.1.1.1",
    ///         Lifetime = 65535,
    ///         NMhaeKey = "'ENC M2wyM3DcnUhqgich7vsLk5oVuPAI9LTkcFNt0c3jI1ujC6w1XBot7gsRAf2S8X5dagfUnJGhZ5LrQxw21e4y8oXuCOLp8MmaRZbCkxYCAl1wm/wVY3aNzVk2+jE='",
    ///         NMhaeKeyType = "ascii",
    ///         NMhaeSpi = 256,
    ///         RegInterval = 5,
    ///         RegRetry = 3,
    ///         RenewInterval = 60,
    ///         RoamingInterface = "port3",
    ///         Status = "disable",
    ///         TunnelMode = "gre",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// System MobileTunnel can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/mobiletunnel:Mobiletunnel labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/mobiletunnel:Mobiletunnel labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/mobiletunnel:Mobiletunnel")]
    public partial class Mobiletunnel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Hash Algorithm (Keyed MD5). Valid values: `hmac-md5`.
        /// </summary>
        [Output("hashAlgorithm")]
        public Output<string> HashAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Home IP address (Format: xxx.xxx.xxx.xxx).
        /// </summary>
        [Output("homeAddress")]
        public Output<string> HomeAddress { get; private set; } = null!;

        /// <summary>
        /// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
        /// </summary>
        [Output("homeAgent")]
        public Output<string> HomeAgent { get; private set; } = null!;

        /// <summary>
        /// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
        /// </summary>
        [Output("lifetime")]
        public Output<int> Lifetime { get; private set; } = null!;

        /// <summary>
        /// NEMO authentication key.
        /// </summary>
        [Output("nMhaeKey")]
        public Output<string> NMhaeKey { get; private set; } = null!;

        /// <summary>
        /// NEMO authentication key type (ascii or base64). Valid values: `ascii`, `base64`.
        /// </summary>
        [Output("nMhaeKeyType")]
        public Output<string> NMhaeKeyType { get; private set; } = null!;

        /// <summary>
        /// NEMO authentication SPI (default: 256).
        /// </summary>
        [Output("nMhaeSpi")]
        public Output<int> NMhaeSpi { get; private set; } = null!;

        /// <summary>
        /// Tunnel name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// NEMO network configuration. The structure of `network` block is documented below.
        /// </summary>
        [Output("networks")]
        public Output<ImmutableArray<Outputs.MobiletunnelNetwork>> Networks { get; private set; } = null!;

        /// <summary>
        /// NMMO HA registration interval (5 - 300, default = 5).
        /// </summary>
        [Output("regInterval")]
        public Output<int> RegInterval { get; private set; } = null!;

        /// <summary>
        /// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
        /// </summary>
        [Output("regRetry")]
        public Output<int> RegRetry { get; private set; } = null!;

        /// <summary>
        /// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
        /// </summary>
        [Output("renewInterval")]
        public Output<int> RenewInterval { get; private set; } = null!;

        /// <summary>
        /// Select the associated interface name from available options.
        /// </summary>
        [Output("roamingInterface")]
        public Output<string> RoamingInterface { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this mobile tunnel. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// NEMO tunnnel mode (GRE tunnel). Valid values: `gre`.
        /// </summary>
        [Output("tunnelMode")]
        public Output<string> TunnelMode { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Mobiletunnel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Mobiletunnel(string name, MobiletunnelArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/mobiletunnel:Mobiletunnel", name, args ?? new MobiletunnelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Mobiletunnel(string name, Input<string> id, MobiletunnelState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/mobiletunnel:Mobiletunnel", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
                AdditionalSecretOutputs =
                {
                    "nMhaeKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Mobiletunnel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Mobiletunnel Get(string name, Input<string> id, MobiletunnelState? state = null, CustomResourceOptions? options = null)
        {
            return new Mobiletunnel(name, id, state, options);
        }
    }

    public sealed class MobiletunnelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Hash Algorithm (Keyed MD5). Valid values: `hmac-md5`.
        /// </summary>
        [Input("hashAlgorithm", required: true)]
        public Input<string> HashAlgorithm { get; set; } = null!;

        /// <summary>
        /// Home IP address (Format: xxx.xxx.xxx.xxx).
        /// </summary>
        [Input("homeAddress")]
        public Input<string>? HomeAddress { get; set; }

        /// <summary>
        /// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
        /// </summary>
        [Input("homeAgent", required: true)]
        public Input<string> HomeAgent { get; set; } = null!;

        /// <summary>
        /// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
        /// </summary>
        [Input("lifetime", required: true)]
        public Input<int> Lifetime { get; set; } = null!;

        [Input("nMhaeKey")]
        private Input<string>? _nMhaeKey;

        /// <summary>
        /// NEMO authentication key.
        /// </summary>
        public Input<string>? NMhaeKey
        {
            get => _nMhaeKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _nMhaeKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// NEMO authentication key type (ascii or base64). Valid values: `ascii`, `base64`.
        /// </summary>
        [Input("nMhaeKeyType", required: true)]
        public Input<string> NMhaeKeyType { get; set; } = null!;

        /// <summary>
        /// NEMO authentication SPI (default: 256).
        /// </summary>
        [Input("nMhaeSpi", required: true)]
        public Input<int> NMhaeSpi { get; set; } = null!;

        /// <summary>
        /// Tunnel name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networks")]
        private InputList<Inputs.MobiletunnelNetworkArgs>? _networks;

        /// <summary>
        /// NEMO network configuration. The structure of `network` block is documented below.
        /// </summary>
        public InputList<Inputs.MobiletunnelNetworkArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.MobiletunnelNetworkArgs>());
            set => _networks = value;
        }

        /// <summary>
        /// NMMO HA registration interval (5 - 300, default = 5).
        /// </summary>
        [Input("regInterval", required: true)]
        public Input<int> RegInterval { get; set; } = null!;

        /// <summary>
        /// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
        /// </summary>
        [Input("regRetry", required: true)]
        public Input<int> RegRetry { get; set; } = null!;

        /// <summary>
        /// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
        /// </summary>
        [Input("renewInterval", required: true)]
        public Input<int> RenewInterval { get; set; } = null!;

        /// <summary>
        /// Select the associated interface name from available options.
        /// </summary>
        [Input("roamingInterface", required: true)]
        public Input<string> RoamingInterface { get; set; } = null!;

        /// <summary>
        /// Enable/disable this mobile tunnel. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// NEMO tunnnel mode (GRE tunnel). Valid values: `gre`.
        /// </summary>
        [Input("tunnelMode", required: true)]
        public Input<string> TunnelMode { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public MobiletunnelArgs()
        {
        }
        public static new MobiletunnelArgs Empty => new MobiletunnelArgs();
    }

    public sealed class MobiletunnelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Hash Algorithm (Keyed MD5). Valid values: `hmac-md5`.
        /// </summary>
        [Input("hashAlgorithm")]
        public Input<string>? HashAlgorithm { get; set; }

        /// <summary>
        /// Home IP address (Format: xxx.xxx.xxx.xxx).
        /// </summary>
        [Input("homeAddress")]
        public Input<string>? HomeAddress { get; set; }

        /// <summary>
        /// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
        /// </summary>
        [Input("homeAgent")]
        public Input<string>? HomeAgent { get; set; }

        /// <summary>
        /// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
        /// </summary>
        [Input("lifetime")]
        public Input<int>? Lifetime { get; set; }

        [Input("nMhaeKey")]
        private Input<string>? _nMhaeKey;

        /// <summary>
        /// NEMO authentication key.
        /// </summary>
        public Input<string>? NMhaeKey
        {
            get => _nMhaeKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _nMhaeKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// NEMO authentication key type (ascii or base64). Valid values: `ascii`, `base64`.
        /// </summary>
        [Input("nMhaeKeyType")]
        public Input<string>? NMhaeKeyType { get; set; }

        /// <summary>
        /// NEMO authentication SPI (default: 256).
        /// </summary>
        [Input("nMhaeSpi")]
        public Input<int>? NMhaeSpi { get; set; }

        /// <summary>
        /// Tunnel name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networks")]
        private InputList<Inputs.MobiletunnelNetworkGetArgs>? _networks;

        /// <summary>
        /// NEMO network configuration. The structure of `network` block is documented below.
        /// </summary>
        public InputList<Inputs.MobiletunnelNetworkGetArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.MobiletunnelNetworkGetArgs>());
            set => _networks = value;
        }

        /// <summary>
        /// NMMO HA registration interval (5 - 300, default = 5).
        /// </summary>
        [Input("regInterval")]
        public Input<int>? RegInterval { get; set; }

        /// <summary>
        /// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
        /// </summary>
        [Input("regRetry")]
        public Input<int>? RegRetry { get; set; }

        /// <summary>
        /// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
        /// </summary>
        [Input("renewInterval")]
        public Input<int>? RenewInterval { get; set; }

        /// <summary>
        /// Select the associated interface name from available options.
        /// </summary>
        [Input("roamingInterface")]
        public Input<string>? RoamingInterface { get; set; }

        /// <summary>
        /// Enable/disable this mobile tunnel. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// NEMO tunnnel mode (GRE tunnel). Valid values: `gre`.
        /// </summary>
        [Input("tunnelMode")]
        public Input<string>? TunnelMode { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public MobiletunnelState()
        {
        }
        public static new MobiletunnelState Empty => new MobiletunnelState();
    }
}
