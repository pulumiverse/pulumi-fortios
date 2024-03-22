// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router
{
    /// <summary>
    /// Configure IPv4 static routing tables.
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
    ///     var trname = new Fortios.Router.Static("trname", new()
    ///     {
    ///         Bfd = "disable",
    ///         Blackhole = "disable",
    ///         Device = "port4",
    ///         Distance = 10,
    ///         Dst = "1.0.0.0 255.240.0.0",
    ///         DynamicGateway = "disable",
    ///         Gateway = "0.0.0.0",
    ///         InternetService = 0,
    ///         LinkMonitorExempt = "disable",
    ///         Priority = 22,
    ///         SeqNum = 1,
    ///         Src = "0.0.0.0 0.0.0.0",
    ///         Status = "enable",
    ///         VirtualWanLink = "disable",
    ///         Vrf = 0,
    ///         Weight = 2,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Router Static can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/static:Static labelname {{seq_num}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/static:Static labelname {{seq_num}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:router/static:Static")]
    public partial class Static : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("bfd")]
        public Output<string> Bfd { get; private set; } = null!;

        /// <summary>
        /// Enable/disable black hole. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("blackhole")]
        public Output<string> Blackhole { get; private set; } = null!;

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Gateway out interface or tunnel.
        /// </summary>
        [Output("device")]
        public Output<string> Device { get; private set; } = null!;

        /// <summary>
        /// Administrative distance (1 - 255).
        /// </summary>
        [Output("distance")]
        public Output<int> Distance { get; private set; } = null!;

        /// <summary>
        /// Destination IP and mask for this route.
        /// </summary>
        [Output("dst")]
        public Output<string> Dst { get; private set; } = null!;

        /// <summary>
        /// Name of firewall address or address group.
        /// </summary>
        [Output("dstaddr")]
        public Output<string> Dstaddr { get; private set; } = null!;

        /// <summary>
        /// Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dynamicGateway")]
        public Output<string> DynamicGateway { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Gateway IP for this route.
        /// </summary>
        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Application ID in the Internet service database.
        /// </summary>
        [Output("internetService")]
        public Output<int> InternetService { get; private set; } = null!;

        /// <summary>
        /// Application name in the Internet service custom database.
        /// </summary>
        [Output("internetServiceCustom")]
        public Output<string> InternetServiceCustom { get; private set; } = null!;

        /// <summary>
        /// Enable/disable withdrawing this route when link monitor or health check is down. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("linkMonitorExempt")]
        public Output<string> LinkMonitorExempt { get; private set; } = null!;

        /// <summary>
        /// Preferred source IP for this route.
        /// </summary>
        [Output("preferredSource")]
        public Output<string> PreferredSource { get; private set; } = null!;

        /// <summary>
        /// Administrative priority. On FortiOS versions 6.2.0-6.4.1: 0 - 4294967295. On FortiOS versions 6.4.2-7.0.3: 0 - 65535. On FortiOS versions &gt;= 7.0.4: 1 - 65535.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Enable/disable egress through SD-WAN. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("sdwan")]
        public Output<string> Sdwan { get; private set; } = null!;

        /// <summary>
        /// Choose SD-WAN Zone. The structure of `sdwan_zone` block is documented below.
        /// </summary>
        [Output("sdwanZones")]
        public Output<ImmutableArray<Outputs.StaticSdwanZone>> SdwanZones { get; private set; } = null!;

        /// <summary>
        /// Sequence number.
        /// </summary>
        [Output("seqNum")]
        public Output<int> SeqNum { get; private set; } = null!;

        /// <summary>
        /// Source prefix for this route.
        /// </summary>
        [Output("src")]
        public Output<string> Src { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this static route. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Route tag.
        /// </summary>
        [Output("tag")]
        public Output<int> Tag { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("virtualWanLink")]
        public Output<string> VirtualWanLink { get; private set; } = null!;

        /// <summary>
        /// Virtual Routing Forwarding ID.
        /// </summary>
        [Output("vrf")]
        public Output<int> Vrf { get; private set; } = null!;

        /// <summary>
        /// Administrative weight (0 - 255).
        /// </summary>
        [Output("weight")]
        public Output<int> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a Static resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Static(string name, StaticArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:router/static:Static", name, args ?? new StaticArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Static(string name, Input<string> id, StaticState? state = null, CustomResourceOptions? options = null)
            : base("fortios:router/static:Static", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Static resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Static Get(string name, Input<string> id, StaticState? state = null, CustomResourceOptions? options = null)
        {
            return new Static(name, id, state, options);
        }
    }

    public sealed class StaticArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bfd")]
        public Input<string>? Bfd { get; set; }

        /// <summary>
        /// Enable/disable black hole. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("blackhole")]
        public Input<string>? Blackhole { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Gateway out interface or tunnel.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Administrative distance (1 - 255).
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// Destination IP and mask for this route.
        /// </summary>
        [Input("dst")]
        public Input<string>? Dst { get; set; }

        /// <summary>
        /// Name of firewall address or address group.
        /// </summary>
        [Input("dstaddr")]
        public Input<string>? Dstaddr { get; set; }

        /// <summary>
        /// Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dynamicGateway")]
        public Input<string>? DynamicGateway { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Gateway IP for this route.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Application ID in the Internet service database.
        /// </summary>
        [Input("internetService")]
        public Input<int>? InternetService { get; set; }

        /// <summary>
        /// Application name in the Internet service custom database.
        /// </summary>
        [Input("internetServiceCustom")]
        public Input<string>? InternetServiceCustom { get; set; }

        /// <summary>
        /// Enable/disable withdrawing this route when link monitor or health check is down. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("linkMonitorExempt")]
        public Input<string>? LinkMonitorExempt { get; set; }

        /// <summary>
        /// Preferred source IP for this route.
        /// </summary>
        [Input("preferredSource")]
        public Input<string>? PreferredSource { get; set; }

        /// <summary>
        /// Administrative priority. On FortiOS versions 6.2.0-6.4.1: 0 - 4294967295. On FortiOS versions 6.4.2-7.0.3: 0 - 65535. On FortiOS versions &gt;= 7.0.4: 1 - 65535.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Enable/disable egress through SD-WAN. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sdwan")]
        public Input<string>? Sdwan { get; set; }

        [Input("sdwanZones")]
        private InputList<Inputs.StaticSdwanZoneArgs>? _sdwanZones;

        /// <summary>
        /// Choose SD-WAN Zone. The structure of `sdwan_zone` block is documented below.
        /// </summary>
        public InputList<Inputs.StaticSdwanZoneArgs> SdwanZones
        {
            get => _sdwanZones ?? (_sdwanZones = new InputList<Inputs.StaticSdwanZoneArgs>());
            set => _sdwanZones = value;
        }

        /// <summary>
        /// Sequence number.
        /// </summary>
        [Input("seqNum")]
        public Input<int>? SeqNum { get; set; }

        /// <summary>
        /// Source prefix for this route.
        /// </summary>
        [Input("src")]
        public Input<string>? Src { get; set; }

        /// <summary>
        /// Enable/disable this static route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Route tag.
        /// </summary>
        [Input("tag")]
        public Input<int>? Tag { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("virtualWanLink")]
        public Input<string>? VirtualWanLink { get; set; }

        /// <summary>
        /// Virtual Routing Forwarding ID.
        /// </summary>
        [Input("vrf")]
        public Input<int>? Vrf { get; set; }

        /// <summary>
        /// Administrative weight (0 - 255).
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public StaticArgs()
        {
        }
        public static new StaticArgs Empty => new StaticArgs();
    }

    public sealed class StaticState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bfd")]
        public Input<string>? Bfd { get; set; }

        /// <summary>
        /// Enable/disable black hole. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("blackhole")]
        public Input<string>? Blackhole { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Gateway out interface or tunnel.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Administrative distance (1 - 255).
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// Destination IP and mask for this route.
        /// </summary>
        [Input("dst")]
        public Input<string>? Dst { get; set; }

        /// <summary>
        /// Name of firewall address or address group.
        /// </summary>
        [Input("dstaddr")]
        public Input<string>? Dstaddr { get; set; }

        /// <summary>
        /// Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dynamicGateway")]
        public Input<string>? DynamicGateway { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Gateway IP for this route.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Application ID in the Internet service database.
        /// </summary>
        [Input("internetService")]
        public Input<int>? InternetService { get; set; }

        /// <summary>
        /// Application name in the Internet service custom database.
        /// </summary>
        [Input("internetServiceCustom")]
        public Input<string>? InternetServiceCustom { get; set; }

        /// <summary>
        /// Enable/disable withdrawing this route when link monitor or health check is down. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("linkMonitorExempt")]
        public Input<string>? LinkMonitorExempt { get; set; }

        /// <summary>
        /// Preferred source IP for this route.
        /// </summary>
        [Input("preferredSource")]
        public Input<string>? PreferredSource { get; set; }

        /// <summary>
        /// Administrative priority. On FortiOS versions 6.2.0-6.4.1: 0 - 4294967295. On FortiOS versions 6.4.2-7.0.3: 0 - 65535. On FortiOS versions &gt;= 7.0.4: 1 - 65535.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Enable/disable egress through SD-WAN. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sdwan")]
        public Input<string>? Sdwan { get; set; }

        [Input("sdwanZones")]
        private InputList<Inputs.StaticSdwanZoneGetArgs>? _sdwanZones;

        /// <summary>
        /// Choose SD-WAN Zone. The structure of `sdwan_zone` block is documented below.
        /// </summary>
        public InputList<Inputs.StaticSdwanZoneGetArgs> SdwanZones
        {
            get => _sdwanZones ?? (_sdwanZones = new InputList<Inputs.StaticSdwanZoneGetArgs>());
            set => _sdwanZones = value;
        }

        /// <summary>
        /// Sequence number.
        /// </summary>
        [Input("seqNum")]
        public Input<int>? SeqNum { get; set; }

        /// <summary>
        /// Source prefix for this route.
        /// </summary>
        [Input("src")]
        public Input<string>? Src { get; set; }

        /// <summary>
        /// Enable/disable this static route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Route tag.
        /// </summary>
        [Input("tag")]
        public Input<int>? Tag { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("virtualWanLink")]
        public Input<string>? VirtualWanLink { get; set; }

        /// <summary>
        /// Virtual Routing Forwarding ID.
        /// </summary>
        [Input("vrf")]
        public Input<int>? Vrf { get; set; }

        /// <summary>
        /// Administrative weight (0 - 255).
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public StaticState()
        {
        }
        public static new StaticState Empty => new StaticState();
    }
}
