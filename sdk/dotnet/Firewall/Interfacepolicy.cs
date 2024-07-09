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
    /// Configure IPv4 interface policies.
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
    ///     var trname = new Fortios.Firewall.Interfacepolicy("trname", new()
    ///     {
    ///         AddressType = "ipv4",
    ///         ApplicationListStatus = "disable",
    ///         AvProfileStatus = "disable",
    ///         DlpSensorStatus = "disable",
    ///         Dsri = "disable",
    ///         Dstaddrs = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.InterfacepolicyDstaddrArgs
    ///             {
    ///                 Name = "all",
    ///             },
    ///         },
    ///         Interface = "port4",
    ///         IpsSensorStatus = "disable",
    ///         Logtraffic = "all",
    ///         Policyid = 1,
    ///         ScanBotnetConnections = "block",
    ///         Services = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.InterfacepolicyServiceArgs
    ///             {
    ///                 Name = "ALL",
    ///             },
    ///         },
    ///         SpamfilterProfileStatus = "disable",
    ///         Srcaddrs = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.InterfacepolicySrcaddrArgs
    ///             {
    ///                 Name = "all",
    ///             },
    ///         },
    ///         Status = "enable",
    ///         WebfilterProfileStatus = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall InterfacePolicy can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/interfacepolicy:Interfacepolicy labelname {{policyid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/interfacepolicy:Interfacepolicy labelname {{policyid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/interfacepolicy:Interfacepolicy")]
    public partial class Interfacepolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Policy address type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Output("addressType")]
        public Output<string> AddressType { get; private set; } = null!;

        /// <summary>
        /// Application list name.
        /// </summary>
        [Output("applicationList")]
        public Output<string> ApplicationList { get; private set; } = null!;

        /// <summary>
        /// Enable/disable application control. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("applicationListStatus")]
        public Output<string> ApplicationListStatus { get; private set; } = null!;

        /// <summary>
        /// Antivirus profile.
        /// </summary>
        [Output("avProfile")]
        public Output<string> AvProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable antivirus. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("avProfileStatus")]
        public Output<string> AvProfileStatus { get; private set; } = null!;

        /// <summary>
        /// CASB profile.
        /// </summary>
        [Output("casbProfile")]
        public Output<string> CasbProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable CASB. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("casbProfileStatus")]
        public Output<string> CasbProfileStatus { get; private set; } = null!;

        /// <summary>
        /// Comments.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// DLP profile name.
        /// </summary>
        [Output("dlpProfile")]
        public Output<string> DlpProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DLP. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dlpProfileStatus")]
        public Output<string> DlpProfileStatus { get; private set; } = null!;

        /// <summary>
        /// DLP sensor name.
        /// </summary>
        [Output("dlpSensor")]
        public Output<string> DlpSensor { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DLP. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dlpSensorStatus")]
        public Output<string> DlpSensorStatus { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DSRI. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dsri")]
        public Output<string> Dsri { get; private set; } = null!;

        /// <summary>
        /// Address object to limit traffic monitoring to network traffic sent to the specified address or range. The structure of `dstaddr` block is documented below.
        /// </summary>
        [Output("dstaddrs")]
        public Output<ImmutableArray<Outputs.InterfacepolicyDstaddr>> Dstaddrs { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Email filter profile.
        /// </summary>
        [Output("emailfilterProfile")]
        public Output<string> EmailfilterProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable email filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("emailfilterProfileStatus")]
        public Output<string> EmailfilterProfileStatus { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Monitored interface name from available interfaces.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// IPS sensor name.
        /// </summary>
        [Output("ipsSensor")]
        public Output<string> IpsSensor { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipsSensorStatus")]
        public Output<string> IpsSensorStatus { get; private set; } = null!;

        /// <summary>
        /// Label.
        /// </summary>
        [Output("label")]
        public Output<string> Label { get; private set; } = null!;

        /// <summary>
        /// Logging type to be used in this policy (Options: all | utm | disable, Default: utm). Valid values: `all`, `utm`, `disable`.
        /// </summary>
        [Output("logtraffic")]
        public Output<string> Logtraffic { get; private set; } = null!;

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Output("policyid")]
        public Output<int> Policyid { get; private set; } = null!;

        /// <summary>
        /// Enable/disable scanning for connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        [Output("scanBotnetConnections")]
        public Output<string> ScanBotnetConnections { get; private set; } = null!;

        /// <summary>
        /// Service object from available options. The structure of `service` block is documented below.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<Outputs.InterfacepolicyService>> Services { get; private set; } = null!;

        /// <summary>
        /// Antispam profile.
        /// </summary>
        [Output("spamfilterProfile")]
        public Output<string> SpamfilterProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable antispam. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("spamfilterProfileStatus")]
        public Output<string> SpamfilterProfileStatus { get; private set; } = null!;

        /// <summary>
        /// Address object to limit traffic monitoring to network traffic sent from the specified address or range. The structure of `srcaddr` block is documented below.
        /// </summary>
        [Output("srcaddrs")]
        public Output<ImmutableArray<Outputs.InterfacepolicySrcaddr>> Srcaddrs { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this policy. Valid values: `enable`, `disable`.
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
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Web filter profile.
        /// </summary>
        [Output("webfilterProfile")]
        public Output<string> WebfilterProfile { get; private set; } = null!;

        /// <summary>
        /// Enable/disable web filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webfilterProfileStatus")]
        public Output<string> WebfilterProfileStatus { get; private set; } = null!;


        /// <summary>
        /// Create a Interfacepolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Interfacepolicy(string name, InterfacepolicyArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/interfacepolicy:Interfacepolicy", name, args ?? new InterfacepolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Interfacepolicy(string name, Input<string> id, InterfacepolicyState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/interfacepolicy:Interfacepolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Interfacepolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Interfacepolicy Get(string name, Input<string> id, InterfacepolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Interfacepolicy(name, id, state, options);
        }
    }

    public sealed class InterfacepolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy address type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("addressType")]
        public Input<string>? AddressType { get; set; }

        /// <summary>
        /// Application list name.
        /// </summary>
        [Input("applicationList")]
        public Input<string>? ApplicationList { get; set; }

        /// <summary>
        /// Enable/disable application control. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("applicationListStatus")]
        public Input<string>? ApplicationListStatus { get; set; }

        /// <summary>
        /// Antivirus profile.
        /// </summary>
        [Input("avProfile")]
        public Input<string>? AvProfile { get; set; }

        /// <summary>
        /// Enable/disable antivirus. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("avProfileStatus")]
        public Input<string>? AvProfileStatus { get; set; }

        /// <summary>
        /// CASB profile.
        /// </summary>
        [Input("casbProfile")]
        public Input<string>? CasbProfile { get; set; }

        /// <summary>
        /// Enable/disable CASB. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("casbProfileStatus")]
        public Input<string>? CasbProfileStatus { get; set; }

        /// <summary>
        /// Comments.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// DLP profile name.
        /// </summary>
        [Input("dlpProfile")]
        public Input<string>? DlpProfile { get; set; }

        /// <summary>
        /// Enable/disable DLP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpProfileStatus")]
        public Input<string>? DlpProfileStatus { get; set; }

        /// <summary>
        /// DLP sensor name.
        /// </summary>
        [Input("dlpSensor")]
        public Input<string>? DlpSensor { get; set; }

        /// <summary>
        /// Enable/disable DLP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpSensorStatus")]
        public Input<string>? DlpSensorStatus { get; set; }

        /// <summary>
        /// Enable/disable DSRI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dsri")]
        public Input<string>? Dsri { get; set; }

        [Input("dstaddrs", required: true)]
        private InputList<Inputs.InterfacepolicyDstaddrArgs>? _dstaddrs;

        /// <summary>
        /// Address object to limit traffic monitoring to network traffic sent to the specified address or range. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.InterfacepolicyDstaddrArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.InterfacepolicyDstaddrArgs>());
            set => _dstaddrs = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Email filter profile.
        /// </summary>
        [Input("emailfilterProfile")]
        public Input<string>? EmailfilterProfile { get; set; }

        /// <summary>
        /// Enable/disable email filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("emailfilterProfileStatus")]
        public Input<string>? EmailfilterProfileStatus { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Monitored interface name from available interfaces.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// IPS sensor name.
        /// </summary>
        [Input("ipsSensor")]
        public Input<string>? IpsSensor { get; set; }

        /// <summary>
        /// Enable/disable IPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsSensorStatus")]
        public Input<string>? IpsSensorStatus { get; set; }

        /// <summary>
        /// Label.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Logging type to be used in this policy (Options: all | utm | disable, Default: utm). Valid values: `all`, `utm`, `disable`.
        /// </summary>
        [Input("logtraffic")]
        public Input<string>? Logtraffic { get; set; }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("policyid")]
        public Input<int>? Policyid { get; set; }

        /// <summary>
        /// Enable/disable scanning for connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        [Input("scanBotnetConnections")]
        public Input<string>? ScanBotnetConnections { get; set; }

        [Input("services", required: true)]
        private InputList<Inputs.InterfacepolicyServiceArgs>? _services;

        /// <summary>
        /// Service object from available options. The structure of `service` block is documented below.
        /// </summary>
        public InputList<Inputs.InterfacepolicyServiceArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.InterfacepolicyServiceArgs>());
            set => _services = value;
        }

        /// <summary>
        /// Antispam profile.
        /// </summary>
        [Input("spamfilterProfile")]
        public Input<string>? SpamfilterProfile { get; set; }

        /// <summary>
        /// Enable/disable antispam. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("spamfilterProfileStatus")]
        public Input<string>? SpamfilterProfileStatus { get; set; }

        [Input("srcaddrs", required: true)]
        private InputList<Inputs.InterfacepolicySrcaddrArgs>? _srcaddrs;

        /// <summary>
        /// Address object to limit traffic monitoring to network traffic sent from the specified address or range. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.InterfacepolicySrcaddrArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.InterfacepolicySrcaddrArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// Enable/disable this policy. Valid values: `enable`, `disable`.
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
        /// Web filter profile.
        /// </summary>
        [Input("webfilterProfile")]
        public Input<string>? WebfilterProfile { get; set; }

        /// <summary>
        /// Enable/disable web filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webfilterProfileStatus")]
        public Input<string>? WebfilterProfileStatus { get; set; }

        public InterfacepolicyArgs()
        {
        }
        public static new InterfacepolicyArgs Empty => new InterfacepolicyArgs();
    }

    public sealed class InterfacepolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy address type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("addressType")]
        public Input<string>? AddressType { get; set; }

        /// <summary>
        /// Application list name.
        /// </summary>
        [Input("applicationList")]
        public Input<string>? ApplicationList { get; set; }

        /// <summary>
        /// Enable/disable application control. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("applicationListStatus")]
        public Input<string>? ApplicationListStatus { get; set; }

        /// <summary>
        /// Antivirus profile.
        /// </summary>
        [Input("avProfile")]
        public Input<string>? AvProfile { get; set; }

        /// <summary>
        /// Enable/disable antivirus. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("avProfileStatus")]
        public Input<string>? AvProfileStatus { get; set; }

        /// <summary>
        /// CASB profile.
        /// </summary>
        [Input("casbProfile")]
        public Input<string>? CasbProfile { get; set; }

        /// <summary>
        /// Enable/disable CASB. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("casbProfileStatus")]
        public Input<string>? CasbProfileStatus { get; set; }

        /// <summary>
        /// Comments.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// DLP profile name.
        /// </summary>
        [Input("dlpProfile")]
        public Input<string>? DlpProfile { get; set; }

        /// <summary>
        /// Enable/disable DLP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpProfileStatus")]
        public Input<string>? DlpProfileStatus { get; set; }

        /// <summary>
        /// DLP sensor name.
        /// </summary>
        [Input("dlpSensor")]
        public Input<string>? DlpSensor { get; set; }

        /// <summary>
        /// Enable/disable DLP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpSensorStatus")]
        public Input<string>? DlpSensorStatus { get; set; }

        /// <summary>
        /// Enable/disable DSRI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dsri")]
        public Input<string>? Dsri { get; set; }

        [Input("dstaddrs")]
        private InputList<Inputs.InterfacepolicyDstaddrGetArgs>? _dstaddrs;

        /// <summary>
        /// Address object to limit traffic monitoring to network traffic sent to the specified address or range. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.InterfacepolicyDstaddrGetArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.InterfacepolicyDstaddrGetArgs>());
            set => _dstaddrs = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Email filter profile.
        /// </summary>
        [Input("emailfilterProfile")]
        public Input<string>? EmailfilterProfile { get; set; }

        /// <summary>
        /// Enable/disable email filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("emailfilterProfileStatus")]
        public Input<string>? EmailfilterProfileStatus { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Monitored interface name from available interfaces.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// IPS sensor name.
        /// </summary>
        [Input("ipsSensor")]
        public Input<string>? IpsSensor { get; set; }

        /// <summary>
        /// Enable/disable IPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsSensorStatus")]
        public Input<string>? IpsSensorStatus { get; set; }

        /// <summary>
        /// Label.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Logging type to be used in this policy (Options: all | utm | disable, Default: utm). Valid values: `all`, `utm`, `disable`.
        /// </summary>
        [Input("logtraffic")]
        public Input<string>? Logtraffic { get; set; }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("policyid")]
        public Input<int>? Policyid { get; set; }

        /// <summary>
        /// Enable/disable scanning for connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        [Input("scanBotnetConnections")]
        public Input<string>? ScanBotnetConnections { get; set; }

        [Input("services")]
        private InputList<Inputs.InterfacepolicyServiceGetArgs>? _services;

        /// <summary>
        /// Service object from available options. The structure of `service` block is documented below.
        /// </summary>
        public InputList<Inputs.InterfacepolicyServiceGetArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.InterfacepolicyServiceGetArgs>());
            set => _services = value;
        }

        /// <summary>
        /// Antispam profile.
        /// </summary>
        [Input("spamfilterProfile")]
        public Input<string>? SpamfilterProfile { get; set; }

        /// <summary>
        /// Enable/disable antispam. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("spamfilterProfileStatus")]
        public Input<string>? SpamfilterProfileStatus { get; set; }

        [Input("srcaddrs")]
        private InputList<Inputs.InterfacepolicySrcaddrGetArgs>? _srcaddrs;

        /// <summary>
        /// Address object to limit traffic monitoring to network traffic sent from the specified address or range. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.InterfacepolicySrcaddrGetArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.InterfacepolicySrcaddrGetArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// Enable/disable this policy. Valid values: `enable`, `disable`.
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
        /// Web filter profile.
        /// </summary>
        [Input("webfilterProfile")]
        public Input<string>? WebfilterProfile { get; set; }

        /// <summary>
        /// Enable/disable web filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webfilterProfileStatus")]
        public Input<string>? WebfilterProfileStatus { get; set; }

        public InterfacepolicyState()
        {
        }
        public static new InterfacepolicyState Empty => new InterfacepolicyState();
    }
}
