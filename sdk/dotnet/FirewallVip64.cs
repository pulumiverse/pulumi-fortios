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
    /// Configure IPv6 to IPv4 virtual IPs. Applies to FortiOS Version `&lt;= 7.0.0`.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.FirewallVip64("trname", new()
    ///     {
    ///         ArpReply = "enable",
    ///         Color = 0,
    ///         Extip = "2001:db8:99:203::22",
    ///         Extport = "0-65535",
    ///         Fosid = 0,
    ///         LdbMethod = "static",
    ///         Mappedip = "1.1.1.1",
    ///         Mappedport = "0-65535",
    ///         Portforward = "disable",
    ///         Protocol = "tcp",
    ///         Type = "static-nat",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Vip64 can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallVip64:FirewallVip64 labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallVip64:FirewallVip64 labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/firewallVip64:FirewallVip64")]
    public partial class FirewallVip64 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable ARP reply. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("arpReply")]
        public Output<string> ArpReply { get; private set; } = null!;

        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Start-external-IP [-end-external-IP].
        /// </summary>
        [Output("extip")]
        public Output<string> Extip { get; private set; } = null!;

        /// <summary>
        /// External service port.
        /// </summary>
        [Output("extport")]
        public Output<string> Extport { get; private set; } = null!;

        /// <summary>
        /// Custom defined id.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
        /// </summary>
        [Output("ldbMethod")]
        public Output<string> LdbMethod { get; private set; } = null!;

        /// <summary>
        /// Start-mapped-IP [-end-mapped-IP].
        /// </summary>
        [Output("mappedip")]
        public Output<string> Mappedip { get; private set; } = null!;

        /// <summary>
        /// Mapped service port.
        /// </summary>
        [Output("mappedport")]
        public Output<string> Mappedport { get; private set; } = null!;

        /// <summary>
        /// Health monitors. The structure of `monitor` block is documented below.
        /// </summary>
        [Output("monitors")]
        public Output<ImmutableArray<Outputs.FirewallVip64Monitor>> Monitors { get; private set; } = null!;

        /// <summary>
        /// VIP64 name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable port forwarding. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("portforward")]
        public Output<string> Portforward { get; private set; } = null!;

        /// <summary>
        /// Mapped port protocol. Valid values: `tcp`, `udp`.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Real servers. The structure of `realservers` block is documented below.
        /// </summary>
        [Output("realservers")]
        public Output<ImmutableArray<Outputs.FirewallVip64Realserver>> Realservers { get; private set; } = null!;

        /// <summary>
        /// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
        /// </summary>
        [Output("serverType")]
        public Output<string> ServerType { get; private set; } = null!;

        /// <summary>
        /// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `src_filter` block is documented below.
        /// </summary>
        [Output("srcFilters")]
        public Output<ImmutableArray<Outputs.FirewallVip64SrcFilter>> SrcFilters { get; private set; } = null!;

        /// <summary>
        /// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallVip64 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallVip64(string name, FirewallVip64Args args, CustomResourceOptions? options = null)
            : base("fortios:index/firewallVip64:FirewallVip64", name, args ?? new FirewallVip64Args(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallVip64(string name, Input<string> id, FirewallVip64State? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallVip64:FirewallVip64", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallVip64 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallVip64 Get(string name, Input<string> id, FirewallVip64State? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallVip64(name, id, state, options);
        }
    }

    public sealed class FirewallVip64Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable ARP reply. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("arpReply")]
        public Input<string>? ArpReply { get; set; }

        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Start-external-IP [-end-external-IP].
        /// </summary>
        [Input("extip", required: true)]
        public Input<string> Extip { get; set; } = null!;

        /// <summary>
        /// External service port.
        /// </summary>
        [Input("extport")]
        public Input<string>? Extport { get; set; }

        /// <summary>
        /// Custom defined id.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
        /// </summary>
        [Input("ldbMethod")]
        public Input<string>? LdbMethod { get; set; }

        /// <summary>
        /// Start-mapped-IP [-end-mapped-IP].
        /// </summary>
        [Input("mappedip", required: true)]
        public Input<string> Mappedip { get; set; } = null!;

        /// <summary>
        /// Mapped service port.
        /// </summary>
        [Input("mappedport")]
        public Input<string>? Mappedport { get; set; }

        [Input("monitors")]
        private InputList<Inputs.FirewallVip64MonitorArgs>? _monitors;

        /// <summary>
        /// Health monitors. The structure of `monitor` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallVip64MonitorArgs> Monitors
        {
            get => _monitors ?? (_monitors = new InputList<Inputs.FirewallVip64MonitorArgs>());
            set => _monitors = value;
        }

        /// <summary>
        /// VIP64 name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable port forwarding. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("portforward")]
        public Input<string>? Portforward { get; set; }

        /// <summary>
        /// Mapped port protocol. Valid values: `tcp`, `udp`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("realservers")]
        private InputList<Inputs.FirewallVip64RealserverArgs>? _realservers;

        /// <summary>
        /// Real servers. The structure of `realservers` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallVip64RealserverArgs> Realservers
        {
            get => _realservers ?? (_realservers = new InputList<Inputs.FirewallVip64RealserverArgs>());
            set => _realservers = value;
        }

        /// <summary>
        /// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
        /// </summary>
        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        [Input("srcFilters")]
        private InputList<Inputs.FirewallVip64SrcFilterArgs>? _srcFilters;

        /// <summary>
        /// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `src_filter` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallVip64SrcFilterArgs> SrcFilters
        {
            get => _srcFilters ?? (_srcFilters = new InputList<Inputs.FirewallVip64SrcFilterArgs>());
            set => _srcFilters = value;
        }

        /// <summary>
        /// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallVip64Args()
        {
        }
        public static new FirewallVip64Args Empty => new FirewallVip64Args();
    }

    public sealed class FirewallVip64State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable ARP reply. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("arpReply")]
        public Input<string>? ArpReply { get; set; }

        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Start-external-IP [-end-external-IP].
        /// </summary>
        [Input("extip")]
        public Input<string>? Extip { get; set; }

        /// <summary>
        /// External service port.
        /// </summary>
        [Input("extport")]
        public Input<string>? Extport { get; set; }

        /// <summary>
        /// Custom defined id.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
        /// </summary>
        [Input("ldbMethod")]
        public Input<string>? LdbMethod { get; set; }

        /// <summary>
        /// Start-mapped-IP [-end-mapped-IP].
        /// </summary>
        [Input("mappedip")]
        public Input<string>? Mappedip { get; set; }

        /// <summary>
        /// Mapped service port.
        /// </summary>
        [Input("mappedport")]
        public Input<string>? Mappedport { get; set; }

        [Input("monitors")]
        private InputList<Inputs.FirewallVip64MonitorGetArgs>? _monitors;

        /// <summary>
        /// Health monitors. The structure of `monitor` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallVip64MonitorGetArgs> Monitors
        {
            get => _monitors ?? (_monitors = new InputList<Inputs.FirewallVip64MonitorGetArgs>());
            set => _monitors = value;
        }

        /// <summary>
        /// VIP64 name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable port forwarding. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("portforward")]
        public Input<string>? Portforward { get; set; }

        /// <summary>
        /// Mapped port protocol. Valid values: `tcp`, `udp`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("realservers")]
        private InputList<Inputs.FirewallVip64RealserverGetArgs>? _realservers;

        /// <summary>
        /// Real servers. The structure of `realservers` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallVip64RealserverGetArgs> Realservers
        {
            get => _realservers ?? (_realservers = new InputList<Inputs.FirewallVip64RealserverGetArgs>());
            set => _realservers = value;
        }

        /// <summary>
        /// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
        /// </summary>
        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        [Input("srcFilters")]
        private InputList<Inputs.FirewallVip64SrcFilterGetArgs>? _srcFilters;

        /// <summary>
        /// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `src_filter` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallVip64SrcFilterGetArgs> SrcFilters
        {
            get => _srcFilters ?? (_srcFilters = new InputList<Inputs.FirewallVip64SrcFilterGetArgs>());
            set => _srcFilters = value;
        }

        /// <summary>
        /// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallVip64State()
        {
        }
        public static new FirewallVip64State Empty => new FirewallVip64State();
    }
}