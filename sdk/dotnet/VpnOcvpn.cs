// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Configure Overlay Controller VPN settings. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// Vpn Ocvpn can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/vpnOcvpn:VpnOcvpn labelname VpnOcvpn
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/vpnOcvpn:VpnOcvpn labelname VpnOcvpn
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/vpnOcvpn:VpnOcvpn")]
    public partial class VpnOcvpn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable auto-discovery shortcuts. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("autoDiscovery")]
        public Output<string> AutoDiscovery { get; private set; } = null!;

        /// <summary>
        /// Control deletion of child short-cut tunnels when the parent tunnel goes down. Valid values: `independent`, `dependent`.
        /// </summary>
        [Output("autoDiscoveryShortcutMode")]
        public Output<string> AutoDiscoveryShortcutMode { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable EAP client authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("eap")]
        public Output<string> Eap { get; private set; } = null!;

        /// <summary>
        /// EAP authentication user group.
        /// </summary>
        [Output("eapUsers")]
        public Output<string> EapUsers { get; private set; } = null!;

        /// <summary>
        /// Configure FortiClient settings. The structure of `forticlient_access` block is documented below.
        /// </summary>
        [Output("forticlientAccess")]
        public Output<Outputs.VpnOcvpnForticlientAccess> ForticlientAccess { get; private set; } = null!;

        /// <summary>
        /// Class B subnet reserved for private IP address assignment.
        /// </summary>
        [Output("ipAllocationBlock")]
        public Output<string> IpAllocationBlock { get; private set; } = null!;

        /// <summary>
        /// Enable/disable multipath redundancy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("multipath")]
        public Output<string> Multipath { get; private set; } = null!;

        /// <summary>
        /// Enable/disable inter-overlay source NAT. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("nat")]
        public Output<string> Nat { get; private set; } = null!;

        /// <summary>
        /// Network overlays to register with Overlay Controller VPN service. The structure of `overlays` block is documented below.
        /// </summary>
        [Output("overlays")]
        public Output<ImmutableArray<Outputs.VpnOcvpnOverlay>> Overlays { get; private set; } = null!;

        /// <summary>
        /// Overlay Controller VPN polling interval.
        /// </summary>
        [Output("pollInterval")]
        public Output<int> PollInterval { get; private set; } = null!;

        /// <summary>
        /// Set device role. Valid values: `spoke`, `primary-hub`, `secondary-hub`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// Enable/disable adding OCVPN tunnels to SDWAN. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("sdwan")]
        public Output<string> Sdwan { get; private set; } = null!;

        /// <summary>
        /// Set SD-WAN zone.
        /// </summary>
        [Output("sdwanZone")]
        public Output<string> SdwanZone { get; private set; } = null!;

        /// <summary>
        /// Enable/disable Overlay Controller cloud assisted VPN. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// FortiGate WAN interfaces to use with OCVPN. The structure of `wan_interface` block is documented below.
        /// </summary>
        [Output("wanInterfaces")]
        public Output<ImmutableArray<Outputs.VpnOcvpnWanInterface>> WanInterfaces { get; private set; } = null!;


        /// <summary>
        /// Create a VpnOcvpn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnOcvpn(string name, VpnOcvpnArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/vpnOcvpn:VpnOcvpn", name, args ?? new VpnOcvpnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpnOcvpn(string name, Input<string> id, VpnOcvpnState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/vpnOcvpn:VpnOcvpn", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpnOcvpn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpnOcvpn Get(string name, Input<string> id, VpnOcvpnState? state = null, CustomResourceOptions? options = null)
        {
            return new VpnOcvpn(name, id, state, options);
        }
    }

    public sealed class VpnOcvpnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable auto-discovery shortcuts. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoDiscovery")]
        public Input<string>? AutoDiscovery { get; set; }

        /// <summary>
        /// Control deletion of child short-cut tunnels when the parent tunnel goes down. Valid values: `independent`, `dependent`.
        /// </summary>
        [Input("autoDiscoveryShortcutMode")]
        public Input<string>? AutoDiscoveryShortcutMode { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable EAP client authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("eap")]
        public Input<string>? Eap { get; set; }

        /// <summary>
        /// EAP authentication user group.
        /// </summary>
        [Input("eapUsers")]
        public Input<string>? EapUsers { get; set; }

        /// <summary>
        /// Configure FortiClient settings. The structure of `forticlient_access` block is documented below.
        /// </summary>
        [Input("forticlientAccess")]
        public Input<Inputs.VpnOcvpnForticlientAccessArgs>? ForticlientAccess { get; set; }

        /// <summary>
        /// Class B subnet reserved for private IP address assignment.
        /// </summary>
        [Input("ipAllocationBlock")]
        public Input<string>? IpAllocationBlock { get; set; }

        /// <summary>
        /// Enable/disable multipath redundancy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("multipath")]
        public Input<string>? Multipath { get; set; }

        /// <summary>
        /// Enable/disable inter-overlay source NAT. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nat")]
        public Input<string>? Nat { get; set; }

        [Input("overlays")]
        private InputList<Inputs.VpnOcvpnOverlayArgs>? _overlays;

        /// <summary>
        /// Network overlays to register with Overlay Controller VPN service. The structure of `overlays` block is documented below.
        /// </summary>
        public InputList<Inputs.VpnOcvpnOverlayArgs> Overlays
        {
            get => _overlays ?? (_overlays = new InputList<Inputs.VpnOcvpnOverlayArgs>());
            set => _overlays = value;
        }

        /// <summary>
        /// Overlay Controller VPN polling interval.
        /// </summary>
        [Input("pollInterval")]
        public Input<int>? PollInterval { get; set; }

        /// <summary>
        /// Set device role. Valid values: `spoke`, `primary-hub`, `secondary-hub`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Enable/disable adding OCVPN tunnels to SDWAN. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sdwan")]
        public Input<string>? Sdwan { get; set; }

        /// <summary>
        /// Set SD-WAN zone.
        /// </summary>
        [Input("sdwanZone")]
        public Input<string>? SdwanZone { get; set; }

        /// <summary>
        /// Enable/disable Overlay Controller cloud assisted VPN. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("wanInterfaces")]
        private InputList<Inputs.VpnOcvpnWanInterfaceArgs>? _wanInterfaces;

        /// <summary>
        /// FortiGate WAN interfaces to use with OCVPN. The structure of `wan_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.VpnOcvpnWanInterfaceArgs> WanInterfaces
        {
            get => _wanInterfaces ?? (_wanInterfaces = new InputList<Inputs.VpnOcvpnWanInterfaceArgs>());
            set => _wanInterfaces = value;
        }

        public VpnOcvpnArgs()
        {
        }
        public static new VpnOcvpnArgs Empty => new VpnOcvpnArgs();
    }

    public sealed class VpnOcvpnState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable auto-discovery shortcuts. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoDiscovery")]
        public Input<string>? AutoDiscovery { get; set; }

        /// <summary>
        /// Control deletion of child short-cut tunnels when the parent tunnel goes down. Valid values: `independent`, `dependent`.
        /// </summary>
        [Input("autoDiscoveryShortcutMode")]
        public Input<string>? AutoDiscoveryShortcutMode { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable EAP client authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("eap")]
        public Input<string>? Eap { get; set; }

        /// <summary>
        /// EAP authentication user group.
        /// </summary>
        [Input("eapUsers")]
        public Input<string>? EapUsers { get; set; }

        /// <summary>
        /// Configure FortiClient settings. The structure of `forticlient_access` block is documented below.
        /// </summary>
        [Input("forticlientAccess")]
        public Input<Inputs.VpnOcvpnForticlientAccessGetArgs>? ForticlientAccess { get; set; }

        /// <summary>
        /// Class B subnet reserved for private IP address assignment.
        /// </summary>
        [Input("ipAllocationBlock")]
        public Input<string>? IpAllocationBlock { get; set; }

        /// <summary>
        /// Enable/disable multipath redundancy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("multipath")]
        public Input<string>? Multipath { get; set; }

        /// <summary>
        /// Enable/disable inter-overlay source NAT. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nat")]
        public Input<string>? Nat { get; set; }

        [Input("overlays")]
        private InputList<Inputs.VpnOcvpnOverlayGetArgs>? _overlays;

        /// <summary>
        /// Network overlays to register with Overlay Controller VPN service. The structure of `overlays` block is documented below.
        /// </summary>
        public InputList<Inputs.VpnOcvpnOverlayGetArgs> Overlays
        {
            get => _overlays ?? (_overlays = new InputList<Inputs.VpnOcvpnOverlayGetArgs>());
            set => _overlays = value;
        }

        /// <summary>
        /// Overlay Controller VPN polling interval.
        /// </summary>
        [Input("pollInterval")]
        public Input<int>? PollInterval { get; set; }

        /// <summary>
        /// Set device role. Valid values: `spoke`, `primary-hub`, `secondary-hub`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Enable/disable adding OCVPN tunnels to SDWAN. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sdwan")]
        public Input<string>? Sdwan { get; set; }

        /// <summary>
        /// Set SD-WAN zone.
        /// </summary>
        [Input("sdwanZone")]
        public Input<string>? SdwanZone { get; set; }

        /// <summary>
        /// Enable/disable Overlay Controller cloud assisted VPN. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("wanInterfaces")]
        private InputList<Inputs.VpnOcvpnWanInterfaceGetArgs>? _wanInterfaces;

        /// <summary>
        /// FortiGate WAN interfaces to use with OCVPN. The structure of `wan_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.VpnOcvpnWanInterfaceGetArgs> WanInterfaces
        {
            get => _wanInterfaces ?? (_wanInterfaces = new InputList<Inputs.VpnOcvpnWanInterfaceGetArgs>());
            set => _wanInterfaces = value;
        }

        public VpnOcvpnState()
        {
        }
        public static new VpnOcvpnState Empty => new VpnOcvpnState();
    }
}
