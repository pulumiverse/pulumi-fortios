// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller
{
    /// <summary>
    /// Configure inter wireless controller operation.
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
    ///     var trname = new Fortios.Wirelesscontroller.Intercontroller("trname", new()
    ///     {
    ///         FastFailoverMax = 10,
    ///         FastFailoverWait = 10,
    ///         InterControllerKey = "ENC XXXX",
    ///         InterControllerMode = "disable",
    ///         InterControllerPri = "primary",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// WirelessController InterController can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/intercontroller:Intercontroller labelname WirelessControllerInterController
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/intercontroller:Intercontroller labelname WirelessControllerInterController
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontroller/intercontroller:Intercontroller")]
    public partial class Intercontroller : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
        /// </summary>
        [Output("fastFailoverMax")]
        public Output<int> FastFailoverMax { get; private set; } = null!;

        /// <summary>
        /// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
        /// </summary>
        [Output("fastFailoverWait")]
        public Output<int> FastFailoverWait { get; private set; } = null!;

        /// <summary>
        /// Secret key for inter-controller communications.
        /// </summary>
        [Output("interControllerKey")]
        public Output<string?> InterControllerKey { get; private set; } = null!;

        /// <summary>
        /// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
        /// </summary>
        [Output("interControllerMode")]
        public Output<string> InterControllerMode { get; private set; } = null!;

        /// <summary>
        /// Fast failover peer wireless controller list. The structure of `inter_controller_peer` block is documented below.
        /// </summary>
        [Output("interControllerPeers")]
        public Output<ImmutableArray<Outputs.IntercontrollerInterControllerPeer>> InterControllerPeers { get; private set; } = null!;

        /// <summary>
        /// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
        /// </summary>
        [Output("interControllerPri")]
        public Output<string> InterControllerPri { get; private set; } = null!;

        /// <summary>
        /// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("l3Roaming")]
        public Output<string> L3Roaming { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Intercontroller resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Intercontroller(string name, IntercontrollerArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/intercontroller:Intercontroller", name, args ?? new IntercontrollerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Intercontroller(string name, Input<string> id, IntercontrollerState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/intercontroller:Intercontroller", name, state, MakeResourceOptions(options, id))
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
                    "interControllerKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Intercontroller resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Intercontroller Get(string name, Input<string> id, IntercontrollerState? state = null, CustomResourceOptions? options = null)
        {
            return new Intercontroller(name, id, state, options);
        }
    }

    public sealed class IntercontrollerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
        /// </summary>
        [Input("fastFailoverMax")]
        public Input<int>? FastFailoverMax { get; set; }

        /// <summary>
        /// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
        /// </summary>
        [Input("fastFailoverWait")]
        public Input<int>? FastFailoverWait { get; set; }

        [Input("interControllerKey")]
        private Input<string>? _interControllerKey;

        /// <summary>
        /// Secret key for inter-controller communications.
        /// </summary>
        public Input<string>? InterControllerKey
        {
            get => _interControllerKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _interControllerKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
        /// </summary>
        [Input("interControllerMode")]
        public Input<string>? InterControllerMode { get; set; }

        [Input("interControllerPeers")]
        private InputList<Inputs.IntercontrollerInterControllerPeerArgs>? _interControllerPeers;

        /// <summary>
        /// Fast failover peer wireless controller list. The structure of `inter_controller_peer` block is documented below.
        /// </summary>
        public InputList<Inputs.IntercontrollerInterControllerPeerArgs> InterControllerPeers
        {
            get => _interControllerPeers ?? (_interControllerPeers = new InputList<Inputs.IntercontrollerInterControllerPeerArgs>());
            set => _interControllerPeers = value;
        }

        /// <summary>
        /// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
        /// </summary>
        [Input("interControllerPri")]
        public Input<string>? InterControllerPri { get; set; }

        /// <summary>
        /// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("l3Roaming")]
        public Input<string>? L3Roaming { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IntercontrollerArgs()
        {
        }
        public static new IntercontrollerArgs Empty => new IntercontrollerArgs();
    }

    public sealed class IntercontrollerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
        /// </summary>
        [Input("fastFailoverMax")]
        public Input<int>? FastFailoverMax { get; set; }

        /// <summary>
        /// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
        /// </summary>
        [Input("fastFailoverWait")]
        public Input<int>? FastFailoverWait { get; set; }

        [Input("interControllerKey")]
        private Input<string>? _interControllerKey;

        /// <summary>
        /// Secret key for inter-controller communications.
        /// </summary>
        public Input<string>? InterControllerKey
        {
            get => _interControllerKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _interControllerKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
        /// </summary>
        [Input("interControllerMode")]
        public Input<string>? InterControllerMode { get; set; }

        [Input("interControllerPeers")]
        private InputList<Inputs.IntercontrollerInterControllerPeerGetArgs>? _interControllerPeers;

        /// <summary>
        /// Fast failover peer wireless controller list. The structure of `inter_controller_peer` block is documented below.
        /// </summary>
        public InputList<Inputs.IntercontrollerInterControllerPeerGetArgs> InterControllerPeers
        {
            get => _interControllerPeers ?? (_interControllerPeers = new InputList<Inputs.IntercontrollerInterControllerPeerGetArgs>());
            set => _interControllerPeers = value;
        }

        /// <summary>
        /// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
        /// </summary>
        [Input("interControllerPri")]
        public Input<string>? InterControllerPri { get; set; }

        /// <summary>
        /// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("l3Roaming")]
        public Input<string>? L3Roaming { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IntercontrollerState()
        {
        }
        public static new IntercontrollerState Empty => new IntercontrollerState();
    }
}
