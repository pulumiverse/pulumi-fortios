// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    /// <summary>
    /// Configure sniffer.
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
    ///     var trname = new Fortios.Firewall.Sniffer("trname", new()
    ///     {
    ///         ApplicationListStatus = "disable",
    ///         AvProfileStatus = "disable",
    ///         DlpSensorStatus = "disable",
    ///         Dsri = "disable",
    ///         Fosid = 1,
    ///         Interface = "port4",
    ///         IpsDosStatus = "disable",
    ///         IpsSensorStatus = "disable",
    ///         Ipv6 = "disable",
    ///         Logtraffic = "utm",
    ///         MaxPacketCount = 4000,
    ///         NonIp = "enable",
    ///         ScanBotnetConnections = "disable",
    ///         SpamfilterProfileStatus = "disable",
    ///         Status = "enable",
    ///         WebfilterProfileStatus = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Sniffer can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/sniffer:Sniffer labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/sniffer:Sniffer labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/sniffer:Sniffer")]
    public partial class Sniffer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.
        /// </summary>
        [Output("anomalies")]
        public Output<ImmutableArray<Outputs.SnifferAnomaly>> Anomalies { get; private set; } = null!;

        /// <summary>
        /// Name of an existing application list.
        /// </summary>
        [Output("applicationList")]
        public Output<string> ApplicationList { get; private set; } = null!;

        /// <summary>
        /// Enable/disable application control profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("applicationListStatus")]
        public Output<string> ApplicationListStatus { get; private set; } = null!;

        /// <summary>
        /// Name of an existing antivirus profile.
        /// </summary>
        [Output("avProfile")]
        public Output<string> AvProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable antivirus profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("avProfileStatus")]
        public Output<string> AvProfileStatus { get; private set; } = null!;

        /// <summary>
        /// Name of an existing CASB profile.
        /// </summary>
        [Output("casbProfile")]
        public Output<string> CasbProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable CASB profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("casbProfileStatus")]
        public Output<string> CasbProfileStatus { get; private set; } = null!;

        /// <summary>
        /// Name of an existing DLP profile.
        /// </summary>
        [Output("dlpProfile")]
        public Output<string> DlpProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DLP profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dlpProfileStatus")]
        public Output<string> DlpProfileStatus { get; private set; } = null!;

        /// <summary>
        /// Name of an existing DLP sensor.
        /// </summary>
        [Output("dlpSensor")]
        public Output<string> DlpSensor { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DLP sensor. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dlpSensorStatus")]
        public Output<string> DlpSensorStatus { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DSRI. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dsri")]
        public Output<string> Dsri { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Name of an existing email filter profile.
        /// </summary>
        [Output("emailfilterProfile")]
        public Output<string> EmailfilterProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable emailfilter. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("emailfilterProfileStatus")]
        public Output<string> EmailfilterProfileStatus { get; private set; } = null!;

        /// <summary>
        /// Name of an existing file-filter profile.
        /// </summary>
        [Output("fileFilterProfile")]
        public Output<string> FileFilterProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable file filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fileFilterProfileStatus")]
        public Output<string> FileFilterProfileStatus { get; private set; } = null!;

        /// <summary>
        /// Sniffer ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// Interface name that traffic sniffing will take place on.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IP threat feed. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipThreatfeedStatus")]
        public Output<string> IpThreatfeedStatus { get; private set; } = null!;

        /// <summary>
        /// Name of an existing IP threat feed. The structure of `ip_threatfeed` block is documented below.
        /// </summary>
        [Output("ipThreatfeeds")]
        public Output<ImmutableArray<Outputs.SnifferIpThreatfeed>> IpThreatfeeds { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IPS DoS anomaly detection. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipsDosStatus")]
        public Output<string> IpsDosStatus { get; private set; } = null!;

        /// <summary>
        /// Name of an existing IPS sensor.
        /// </summary>
        [Output("ipsSensor")]
        public Output<string> IpsSensor { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IPS sensor. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipsSensorStatus")]
        public Output<string> IpsSensorStatus { get; private set; } = null!;

        /// <summary>
        /// Enable/disable sniffing IPv6 packets. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipv6")]
        public Output<string> Ipv6 { get; private set; } = null!;

        /// <summary>
        /// Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all`, `utm`, `disable`.
        /// </summary>
        [Output("logtraffic")]
        public Output<string> Logtraffic { get; private set; } = null!;

        /// <summary>
        /// Maximum packet count. On FortiOS versions 6.2.0: 1 - 1000000, default = 10000. On FortiOS versions 6.2.4-6.4.2, 7.0.0: 1 - 10000, default = 4000. On FortiOS versions 6.4.10-6.4.15, 7.0.1-7.0.15: 1 - 1000000, default = 4000.
        /// </summary>
        [Output("maxPacketCount")]
        public Output<int> MaxPacketCount { get; private set; } = null!;

        /// <summary>
        /// Enable/disable sniffing non-IP packets. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("nonIp")]
        public Output<string> NonIp { get; private set; } = null!;

        /// <summary>
        /// Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255).
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        [Output("scanBotnetConnections")]
        public Output<string> ScanBotnetConnections { get; private set; } = null!;

        /// <summary>
        /// Name of an existing spam filter profile.
        /// </summary>
        [Output("spamfilterProfile")]
        public Output<string> SpamfilterProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable spam filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("spamfilterProfileStatus")]
        public Output<string> SpamfilterProfileStatus { get; private set; } = null!;

        /// <summary>
        /// Enable/disable the active status of the sniffer. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// List of VLANs to sniff.
        /// </summary>
        [Output("vlan")]
        public Output<string> Vlan { get; private set; } = null!;

        /// <summary>
        /// Name of an existing web filter profile.
        /// </summary>
        [Output("webfilterProfile")]
        public Output<string> WebfilterProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable web filter profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webfilterProfileStatus")]
        public Output<string> WebfilterProfileStatus { get; private set; } = null!;


        /// <summary>
        /// Create a Sniffer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Sniffer(string name, SnifferArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/sniffer:Sniffer", name, args ?? new SnifferArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Sniffer(string name, Input<string> id, SnifferState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/sniffer:Sniffer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Sniffer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Sniffer Get(string name, Input<string> id, SnifferState? state = null, CustomResourceOptions? options = null)
        {
            return new Sniffer(name, id, state, options);
        }
    }

    public sealed class SnifferArgs : global::Pulumi.ResourceArgs
    {
        [Input("anomalies")]
        private InputList<Inputs.SnifferAnomalyArgs>? _anomalies;

        /// <summary>
        /// Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.
        /// </summary>
        public InputList<Inputs.SnifferAnomalyArgs> Anomalies
        {
            get => _anomalies ?? (_anomalies = new InputList<Inputs.SnifferAnomalyArgs>());
            set => _anomalies = value;
        }

        /// <summary>
        /// Name of an existing application list.
        /// </summary>
        [Input("applicationList")]
        public Input<string>? ApplicationList { get; set; }

        /// <summary>
        /// Enable/disable application control profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("applicationListStatus")]
        public Input<string>? ApplicationListStatus { get; set; }

        /// <summary>
        /// Name of an existing antivirus profile.
        /// </summary>
        [Input("avProfile")]
        public Input<string>? AvProfile { get; set; }

        /// <summary>
        /// Enable/disable antivirus profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("avProfileStatus")]
        public Input<string>? AvProfileStatus { get; set; }

        /// <summary>
        /// Name of an existing CASB profile.
        /// </summary>
        [Input("casbProfile")]
        public Input<string>? CasbProfile { get; set; }

        /// <summary>
        /// Enable/disable CASB profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("casbProfileStatus")]
        public Input<string>? CasbProfileStatus { get; set; }

        /// <summary>
        /// Name of an existing DLP profile.
        /// </summary>
        [Input("dlpProfile")]
        public Input<string>? DlpProfile { get; set; }

        /// <summary>
        /// Enable/disable DLP profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpProfileStatus")]
        public Input<string>? DlpProfileStatus { get; set; }

        /// <summary>
        /// Name of an existing DLP sensor.
        /// </summary>
        [Input("dlpSensor")]
        public Input<string>? DlpSensor { get; set; }

        /// <summary>
        /// Enable/disable DLP sensor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpSensorStatus")]
        public Input<string>? DlpSensorStatus { get; set; }

        /// <summary>
        /// Enable/disable DSRI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dsri")]
        public Input<string>? Dsri { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Name of an existing email filter profile.
        /// </summary>
        [Input("emailfilterProfile")]
        public Input<string>? EmailfilterProfile { get; set; }

        /// <summary>
        /// Enable/disable emailfilter. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("emailfilterProfileStatus")]
        public Input<string>? EmailfilterProfileStatus { get; set; }

        /// <summary>
        /// Name of an existing file-filter profile.
        /// </summary>
        [Input("fileFilterProfile")]
        public Input<string>? FileFilterProfile { get; set; }

        /// <summary>
        /// Enable/disable file filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fileFilterProfileStatus")]
        public Input<string>? FileFilterProfileStatus { get; set; }

        /// <summary>
        /// Sniffer ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Interface name that traffic sniffing will take place on.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// Enable/disable IP threat feed. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipThreatfeedStatus")]
        public Input<string>? IpThreatfeedStatus { get; set; }

        [Input("ipThreatfeeds")]
        private InputList<Inputs.SnifferIpThreatfeedArgs>? _ipThreatfeeds;

        /// <summary>
        /// Name of an existing IP threat feed. The structure of `ip_threatfeed` block is documented below.
        /// </summary>
        public InputList<Inputs.SnifferIpThreatfeedArgs> IpThreatfeeds
        {
            get => _ipThreatfeeds ?? (_ipThreatfeeds = new InputList<Inputs.SnifferIpThreatfeedArgs>());
            set => _ipThreatfeeds = value;
        }

        /// <summary>
        /// Enable/disable IPS DoS anomaly detection. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsDosStatus")]
        public Input<string>? IpsDosStatus { get; set; }

        /// <summary>
        /// Name of an existing IPS sensor.
        /// </summary>
        [Input("ipsSensor")]
        public Input<string>? IpsSensor { get; set; }

        /// <summary>
        /// Enable/disable IPS sensor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsSensorStatus")]
        public Input<string>? IpsSensorStatus { get; set; }

        /// <summary>
        /// Enable/disable sniffing IPv6 packets. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipv6")]
        public Input<string>? Ipv6 { get; set; }

        /// <summary>
        /// Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all`, `utm`, `disable`.
        /// </summary>
        [Input("logtraffic")]
        public Input<string>? Logtraffic { get; set; }

        /// <summary>
        /// Maximum packet count. On FortiOS versions 6.2.0: 1 - 1000000, default = 10000. On FortiOS versions 6.2.4-6.4.2, 7.0.0: 1 - 10000, default = 4000. On FortiOS versions 6.4.10-6.4.15, 7.0.1-7.0.15: 1 - 1000000, default = 4000.
        /// </summary>
        [Input("maxPacketCount")]
        public Input<int>? MaxPacketCount { get; set; }

        /// <summary>
        /// Enable/disable sniffing non-IP packets. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nonIp")]
        public Input<string>? NonIp { get; set; }

        /// <summary>
        /// Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255).
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        [Input("scanBotnetConnections")]
        public Input<string>? ScanBotnetConnections { get; set; }

        /// <summary>
        /// Name of an existing spam filter profile.
        /// </summary>
        [Input("spamfilterProfile")]
        public Input<string>? SpamfilterProfile { get; set; }

        /// <summary>
        /// Enable/disable spam filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("spamfilterProfileStatus")]
        public Input<string>? SpamfilterProfileStatus { get; set; }

        /// <summary>
        /// Enable/disable the active status of the sniffer. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

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

        /// <summary>
        /// List of VLANs to sniff.
        /// </summary>
        [Input("vlan")]
        public Input<string>? Vlan { get; set; }

        /// <summary>
        /// Name of an existing web filter profile.
        /// </summary>
        [Input("webfilterProfile")]
        public Input<string>? WebfilterProfile { get; set; }

        /// <summary>
        /// Enable/disable web filter profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webfilterProfileStatus")]
        public Input<string>? WebfilterProfileStatus { get; set; }

        public SnifferArgs()
        {
        }
        public static new SnifferArgs Empty => new SnifferArgs();
    }

    public sealed class SnifferState : global::Pulumi.ResourceArgs
    {
        [Input("anomalies")]
        private InputList<Inputs.SnifferAnomalyGetArgs>? _anomalies;

        /// <summary>
        /// Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.
        /// </summary>
        public InputList<Inputs.SnifferAnomalyGetArgs> Anomalies
        {
            get => _anomalies ?? (_anomalies = new InputList<Inputs.SnifferAnomalyGetArgs>());
            set => _anomalies = value;
        }

        /// <summary>
        /// Name of an existing application list.
        /// </summary>
        [Input("applicationList")]
        public Input<string>? ApplicationList { get; set; }

        /// <summary>
        /// Enable/disable application control profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("applicationListStatus")]
        public Input<string>? ApplicationListStatus { get; set; }

        /// <summary>
        /// Name of an existing antivirus profile.
        /// </summary>
        [Input("avProfile")]
        public Input<string>? AvProfile { get; set; }

        /// <summary>
        /// Enable/disable antivirus profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("avProfileStatus")]
        public Input<string>? AvProfileStatus { get; set; }

        /// <summary>
        /// Name of an existing CASB profile.
        /// </summary>
        [Input("casbProfile")]
        public Input<string>? CasbProfile { get; set; }

        /// <summary>
        /// Enable/disable CASB profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("casbProfileStatus")]
        public Input<string>? CasbProfileStatus { get; set; }

        /// <summary>
        /// Name of an existing DLP profile.
        /// </summary>
        [Input("dlpProfile")]
        public Input<string>? DlpProfile { get; set; }

        /// <summary>
        /// Enable/disable DLP profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpProfileStatus")]
        public Input<string>? DlpProfileStatus { get; set; }

        /// <summary>
        /// Name of an existing DLP sensor.
        /// </summary>
        [Input("dlpSensor")]
        public Input<string>? DlpSensor { get; set; }

        /// <summary>
        /// Enable/disable DLP sensor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpSensorStatus")]
        public Input<string>? DlpSensorStatus { get; set; }

        /// <summary>
        /// Enable/disable DSRI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dsri")]
        public Input<string>? Dsri { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Name of an existing email filter profile.
        /// </summary>
        [Input("emailfilterProfile")]
        public Input<string>? EmailfilterProfile { get; set; }

        /// <summary>
        /// Enable/disable emailfilter. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("emailfilterProfileStatus")]
        public Input<string>? EmailfilterProfileStatus { get; set; }

        /// <summary>
        /// Name of an existing file-filter profile.
        /// </summary>
        [Input("fileFilterProfile")]
        public Input<string>? FileFilterProfile { get; set; }

        /// <summary>
        /// Enable/disable file filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fileFilterProfileStatus")]
        public Input<string>? FileFilterProfileStatus { get; set; }

        /// <summary>
        /// Sniffer ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Interface name that traffic sniffing will take place on.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Enable/disable IP threat feed. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipThreatfeedStatus")]
        public Input<string>? IpThreatfeedStatus { get; set; }

        [Input("ipThreatfeeds")]
        private InputList<Inputs.SnifferIpThreatfeedGetArgs>? _ipThreatfeeds;

        /// <summary>
        /// Name of an existing IP threat feed. The structure of `ip_threatfeed` block is documented below.
        /// </summary>
        public InputList<Inputs.SnifferIpThreatfeedGetArgs> IpThreatfeeds
        {
            get => _ipThreatfeeds ?? (_ipThreatfeeds = new InputList<Inputs.SnifferIpThreatfeedGetArgs>());
            set => _ipThreatfeeds = value;
        }

        /// <summary>
        /// Enable/disable IPS DoS anomaly detection. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsDosStatus")]
        public Input<string>? IpsDosStatus { get; set; }

        /// <summary>
        /// Name of an existing IPS sensor.
        /// </summary>
        [Input("ipsSensor")]
        public Input<string>? IpsSensor { get; set; }

        /// <summary>
        /// Enable/disable IPS sensor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsSensorStatus")]
        public Input<string>? IpsSensorStatus { get; set; }

        /// <summary>
        /// Enable/disable sniffing IPv6 packets. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipv6")]
        public Input<string>? Ipv6 { get; set; }

        /// <summary>
        /// Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all`, `utm`, `disable`.
        /// </summary>
        [Input("logtraffic")]
        public Input<string>? Logtraffic { get; set; }

        /// <summary>
        /// Maximum packet count. On FortiOS versions 6.2.0: 1 - 1000000, default = 10000. On FortiOS versions 6.2.4-6.4.2, 7.0.0: 1 - 10000, default = 4000. On FortiOS versions 6.4.10-6.4.15, 7.0.1-7.0.15: 1 - 1000000, default = 4000.
        /// </summary>
        [Input("maxPacketCount")]
        public Input<int>? MaxPacketCount { get; set; }

        /// <summary>
        /// Enable/disable sniffing non-IP packets. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nonIp")]
        public Input<string>? NonIp { get; set; }

        /// <summary>
        /// Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255).
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        [Input("scanBotnetConnections")]
        public Input<string>? ScanBotnetConnections { get; set; }

        /// <summary>
        /// Name of an existing spam filter profile.
        /// </summary>
        [Input("spamfilterProfile")]
        public Input<string>? SpamfilterProfile { get; set; }

        /// <summary>
        /// Enable/disable spam filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("spamfilterProfileStatus")]
        public Input<string>? SpamfilterProfileStatus { get; set; }

        /// <summary>
        /// Enable/disable the active status of the sniffer. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

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

        /// <summary>
        /// List of VLANs to sniff.
        /// </summary>
        [Input("vlan")]
        public Input<string>? Vlan { get; set; }

        /// <summary>
        /// Name of an existing web filter profile.
        /// </summary>
        [Input("webfilterProfile")]
        public Input<string>? WebfilterProfile { get; set; }

        /// <summary>
        /// Enable/disable web filter profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webfilterProfileStatus")]
        public Input<string>? WebfilterProfileStatus { get; set; }

        public SnifferState()
        {
        }
        public static new SnifferState Empty => new SnifferState();
    }
}
