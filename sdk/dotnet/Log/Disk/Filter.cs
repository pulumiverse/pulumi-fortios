// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Log.Disk
{
    /// <summary>
    /// Configure filters for local disk logging. Use these filters to determine the log messages to record according to severity and type.
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
    ///     var trname = new Fortios.Log.Disk.Filter("trname", new()
    ///     {
    ///         Anomaly = "enable",
    ///         DlpArchive = "enable",
    ///         Dns = "enable",
    ///         FilterType = "include",
    ///         ForwardTraffic = "enable",
    ///         Gtp = "enable",
    ///         LocalTraffic = "enable",
    ///         MulticastTraffic = "enable",
    ///         Severity = "information",
    ///         SnifferTraffic = "enable",
    ///         Ssh = "enable",
    ///         Voip = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// LogDisk Filter can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/disk/filter:Filter labelname LogDiskFilter
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/disk/filter:Filter labelname LogDiskFilter
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:log/disk/filter:Filter")]
    public partial class Filter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable admin login/logout logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("admin")]
        public Output<string> Admin { get; private set; } = null!;

        /// <summary>
        /// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("anomaly")]
        public Output<string> Anomaly { get; private set; } = null!;

        /// <summary>
        /// Enable/disable firewall authentication logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("auth")]
        public Output<string> Auth { get; private set; } = null!;

        /// <summary>
        /// Enable/disable CPU &amp; memory usage logging every 5 minutes. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("cpuMemoryUsage")]
        public Output<string> CpuMemoryUsage { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DHCP service messages logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dhcp")]
        public Output<string> Dhcp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dlpArchive")]
        public Output<string> DlpArchive { get; private set; } = null!;

        /// <summary>
        /// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dns")]
        public Output<string> Dns { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("event")]
        public Output<string> Event { get; private set; } = null!;

        /// <summary>
        /// Disk log filter.
        /// </summary>
        [Output("filter")]
        public Output<string> Definition { get; private set; } = null!;

        /// <summary>
        /// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        /// </summary>
        [Output("filterType")]
        public Output<string> FilterType { get; private set; } = null!;

        /// <summary>
        /// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fortiSwitch")]
        public Output<string> FortiSwitch { get; private set; } = null!;

        /// <summary>
        /// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("forwardTraffic")]
        public Output<string> ForwardTraffic { get; private set; } = null!;

        /// <summary>
        /// Free Style Filters The structure of `free_style` block is documented below.
        /// </summary>
        [Output("freeStyles")]
        public Output<ImmutableArray<Outputs.FilterFreeStyle>> FreeStyles { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("gtp")]
        public Output<string> Gtp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable HA logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ha")]
        public Output<string> Ha { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IPsec negotiation messages logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipsec")]
        public Output<string> Ipsec { get; private set; } = null!;

        /// <summary>
        /// Enable/disable VIP real server health monitoring logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ldbMonitor")]
        public Output<string> LdbMonitor { get; private set; } = null!;

        /// <summary>
        /// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("localTraffic")]
        public Output<string> LocalTraffic { get; private set; } = null!;

        /// <summary>
        /// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("multicastTraffic")]
        public Output<string> MulticastTraffic { get; private set; } = null!;

        /// <summary>
        /// Enable/disable netscan discovery event logging.
        /// </summary>
        [Output("netscanDiscovery")]
        public Output<string> NetscanDiscovery { get; private set; } = null!;

        /// <summary>
        /// Enable/disable netscan vulnerability event logging.
        /// </summary>
        [Output("netscanVulnerability")]
        public Output<string> NetscanVulnerability { get; private set; } = null!;

        /// <summary>
        /// Enable/disable pattern update logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("pattern")]
        public Output<string> Pattern { get; private set; } = null!;

        /// <summary>
        /// Enable/disable L2TP/PPTP/PPPoE logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ppp")]
        public Output<string> Ppp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable RADIUS messages logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("radius")]
        public Output<string> Radius { get; private set; } = null!;

        /// <summary>
        /// Log to disk every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        /// </summary>
        [Output("severity")]
        public Output<string> Severity { get; private set; } = null!;

        /// <summary>
        /// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("snifferTraffic")]
        public Output<string> SnifferTraffic { get; private set; } = null!;

        /// <summary>
        /// Enable/disable SSH logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ssh")]
        public Output<string> Ssh { get; private set; } = null!;

        /// <summary>
        /// Enable/disable SSL administrator login logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("sslvpnLogAdm")]
        public Output<string> SslvpnLogAdm { get; private set; } = null!;

        /// <summary>
        /// Enable/disable SSL user authentication logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("sslvpnLogAuth")]
        public Output<string> SslvpnLogAuth { get; private set; } = null!;

        /// <summary>
        /// Enable/disable SSL session logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("sslvpnLogSession")]
        public Output<string> SslvpnLogSession { get; private set; } = null!;

        /// <summary>
        /// Enable/disable system activity logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("system")]
        public Output<string> System { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable VIP SSL logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("vipSsl")]
        public Output<string> VipSsl { get; private set; } = null!;

        /// <summary>
        /// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("voip")]
        public Output<string> Voip { get; private set; } = null!;

        /// <summary>
        /// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("wanOpt")]
        public Output<string> WanOpt { get; private set; } = null!;

        /// <summary>
        /// Enable/disable wireless activity event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("wirelessActivity")]
        public Output<string> WirelessActivity { get; private set; } = null!;

        /// <summary>
        /// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ztnaTraffic")]
        public Output<string> ZtnaTraffic { get; private set; } = null!;


        /// <summary>
        /// Create a Filter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Filter(string name, FilterArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:log/disk/filter:Filter", name, args ?? new FilterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Filter(string name, Input<string> id, FilterState? state = null, CustomResourceOptions? options = null)
            : base("fortios:log/disk/filter:Filter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Filter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Filter Get(string name, Input<string> id, FilterState? state = null, CustomResourceOptions? options = null)
        {
            return new Filter(name, id, state, options);
        }
    }

    public sealed class FilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable admin login/logout logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("admin")]
        public Input<string>? Admin { get; set; }

        /// <summary>
        /// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("anomaly")]
        public Input<string>? Anomaly { get; set; }

        /// <summary>
        /// Enable/disable firewall authentication logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("auth")]
        public Input<string>? Auth { get; set; }

        /// <summary>
        /// Enable/disable CPU &amp; memory usage logging every 5 minutes. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cpuMemoryUsage")]
        public Input<string>? CpuMemoryUsage { get; set; }

        /// <summary>
        /// Enable/disable DHCP service messages logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dhcp")]
        public Input<string>? Dhcp { get; set; }

        /// <summary>
        /// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpArchive")]
        public Input<string>? DlpArchive { get; set; }

        /// <summary>
        /// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dns")]
        public Input<string>? Dns { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("event")]
        public Input<string>? Event { get; set; }

        /// <summary>
        /// Disk log filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Definition { get; set; }

        /// <summary>
        /// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        /// </summary>
        [Input("filterType")]
        public Input<string>? FilterType { get; set; }

        /// <summary>
        /// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fortiSwitch")]
        public Input<string>? FortiSwitch { get; set; }

        /// <summary>
        /// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forwardTraffic")]
        public Input<string>? ForwardTraffic { get; set; }

        [Input("freeStyles")]
        private InputList<Inputs.FilterFreeStyleArgs>? _freeStyles;

        /// <summary>
        /// Free Style Filters The structure of `free_style` block is documented below.
        /// </summary>
        public InputList<Inputs.FilterFreeStyleArgs> FreeStyles
        {
            get => _freeStyles ?? (_freeStyles = new InputList<Inputs.FilterFreeStyleArgs>());
            set => _freeStyles = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("gtp")]
        public Input<string>? Gtp { get; set; }

        /// <summary>
        /// Enable/disable HA logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ha")]
        public Input<string>? Ha { get; set; }

        /// <summary>
        /// Enable/disable IPsec negotiation messages logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsec")]
        public Input<string>? Ipsec { get; set; }

        /// <summary>
        /// Enable/disable VIP real server health monitoring logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ldbMonitor")]
        public Input<string>? LdbMonitor { get; set; }

        /// <summary>
        /// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localTraffic")]
        public Input<string>? LocalTraffic { get; set; }

        /// <summary>
        /// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("multicastTraffic")]
        public Input<string>? MulticastTraffic { get; set; }

        /// <summary>
        /// Enable/disable netscan discovery event logging.
        /// </summary>
        [Input("netscanDiscovery")]
        public Input<string>? NetscanDiscovery { get; set; }

        /// <summary>
        /// Enable/disable netscan vulnerability event logging.
        /// </summary>
        [Input("netscanVulnerability")]
        public Input<string>? NetscanVulnerability { get; set; }

        /// <summary>
        /// Enable/disable pattern update logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Enable/disable L2TP/PPTP/PPPoE logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ppp")]
        public Input<string>? Ppp { get; set; }

        /// <summary>
        /// Enable/disable RADIUS messages logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("radius")]
        public Input<string>? Radius { get; set; }

        /// <summary>
        /// Log to disk every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("snifferTraffic")]
        public Input<string>? SnifferTraffic { get; set; }

        /// <summary>
        /// Enable/disable SSH logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ssh")]
        public Input<string>? Ssh { get; set; }

        /// <summary>
        /// Enable/disable SSL administrator login logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslvpnLogAdm")]
        public Input<string>? SslvpnLogAdm { get; set; }

        /// <summary>
        /// Enable/disable SSL user authentication logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslvpnLogAuth")]
        public Input<string>? SslvpnLogAuth { get; set; }

        /// <summary>
        /// Enable/disable SSL session logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslvpnLogSession")]
        public Input<string>? SslvpnLogSession { get; set; }

        /// <summary>
        /// Enable/disable system activity logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("system")]
        public Input<string>? System { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable VIP SSL logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("vipSsl")]
        public Input<string>? VipSsl { get; set; }

        /// <summary>
        /// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("voip")]
        public Input<string>? Voip { get; set; }

        /// <summary>
        /// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wanOpt")]
        public Input<string>? WanOpt { get; set; }

        /// <summary>
        /// Enable/disable wireless activity event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wirelessActivity")]
        public Input<string>? WirelessActivity { get; set; }

        /// <summary>
        /// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ztnaTraffic")]
        public Input<string>? ZtnaTraffic { get; set; }

        public FilterArgs()
        {
        }
        public static new FilterArgs Empty => new FilterArgs();
    }

    public sealed class FilterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable admin login/logout logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("admin")]
        public Input<string>? Admin { get; set; }

        /// <summary>
        /// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("anomaly")]
        public Input<string>? Anomaly { get; set; }

        /// <summary>
        /// Enable/disable firewall authentication logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("auth")]
        public Input<string>? Auth { get; set; }

        /// <summary>
        /// Enable/disable CPU &amp; memory usage logging every 5 minutes. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cpuMemoryUsage")]
        public Input<string>? CpuMemoryUsage { get; set; }

        /// <summary>
        /// Enable/disable DHCP service messages logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dhcp")]
        public Input<string>? Dhcp { get; set; }

        /// <summary>
        /// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpArchive")]
        public Input<string>? DlpArchive { get; set; }

        /// <summary>
        /// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dns")]
        public Input<string>? Dns { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("event")]
        public Input<string>? Event { get; set; }

        /// <summary>
        /// Disk log filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Definition { get; set; }

        /// <summary>
        /// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        /// </summary>
        [Input("filterType")]
        public Input<string>? FilterType { get; set; }

        /// <summary>
        /// Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fortiSwitch")]
        public Input<string>? FortiSwitch { get; set; }

        /// <summary>
        /// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forwardTraffic")]
        public Input<string>? ForwardTraffic { get; set; }

        [Input("freeStyles")]
        private InputList<Inputs.FilterFreeStyleGetArgs>? _freeStyles;

        /// <summary>
        /// Free Style Filters The structure of `free_style` block is documented below.
        /// </summary>
        public InputList<Inputs.FilterFreeStyleGetArgs> FreeStyles
        {
            get => _freeStyles ?? (_freeStyles = new InputList<Inputs.FilterFreeStyleGetArgs>());
            set => _freeStyles = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("gtp")]
        public Input<string>? Gtp { get; set; }

        /// <summary>
        /// Enable/disable HA logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ha")]
        public Input<string>? Ha { get; set; }

        /// <summary>
        /// Enable/disable IPsec negotiation messages logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsec")]
        public Input<string>? Ipsec { get; set; }

        /// <summary>
        /// Enable/disable VIP real server health monitoring logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ldbMonitor")]
        public Input<string>? LdbMonitor { get; set; }

        /// <summary>
        /// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localTraffic")]
        public Input<string>? LocalTraffic { get; set; }

        /// <summary>
        /// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("multicastTraffic")]
        public Input<string>? MulticastTraffic { get; set; }

        /// <summary>
        /// Enable/disable netscan discovery event logging.
        /// </summary>
        [Input("netscanDiscovery")]
        public Input<string>? NetscanDiscovery { get; set; }

        /// <summary>
        /// Enable/disable netscan vulnerability event logging.
        /// </summary>
        [Input("netscanVulnerability")]
        public Input<string>? NetscanVulnerability { get; set; }

        /// <summary>
        /// Enable/disable pattern update logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Enable/disable L2TP/PPTP/PPPoE logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ppp")]
        public Input<string>? Ppp { get; set; }

        /// <summary>
        /// Enable/disable RADIUS messages logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("radius")]
        public Input<string>? Radius { get; set; }

        /// <summary>
        /// Log to disk every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("snifferTraffic")]
        public Input<string>? SnifferTraffic { get; set; }

        /// <summary>
        /// Enable/disable SSH logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ssh")]
        public Input<string>? Ssh { get; set; }

        /// <summary>
        /// Enable/disable SSL administrator login logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslvpnLogAdm")]
        public Input<string>? SslvpnLogAdm { get; set; }

        /// <summary>
        /// Enable/disable SSL user authentication logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslvpnLogAuth")]
        public Input<string>? SslvpnLogAuth { get; set; }

        /// <summary>
        /// Enable/disable SSL session logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslvpnLogSession")]
        public Input<string>? SslvpnLogSession { get; set; }

        /// <summary>
        /// Enable/disable system activity logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("system")]
        public Input<string>? System { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable VIP SSL logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("vipSsl")]
        public Input<string>? VipSsl { get; set; }

        /// <summary>
        /// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("voip")]
        public Input<string>? Voip { get; set; }

        /// <summary>
        /// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wanOpt")]
        public Input<string>? WanOpt { get; set; }

        /// <summary>
        /// Enable/disable wireless activity event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wirelessActivity")]
        public Input<string>? WirelessActivity { get; set; }

        /// <summary>
        /// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ztnaTraffic")]
        public Input<string>? ZtnaTraffic { get; set; }

        public FilterState()
        {
        }
        public static new FilterState Empty => new FilterState();
    }
}
