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
    /// Configure Link Health Monitor.
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
    ///     var trname = new Fortios.System.Linkmonitor("trname", new()
    ///     {
    ///         AddrMode = "ipv4",
    ///         Failtime = 5,
    ///         GatewayIp = "2.2.2.2",
    ///         GatewayIp6 = "::",
    ///         HaPriority = 1,
    ///         HttpAgent = "Chrome/ Safari/",
    ///         HttpGet = "/",
    ///         Interval = 1,
    ///         PacketSize = 64,
    ///         Port = 80,
    ///         Protocol = "ping",
    ///         Recoverytime = 5,
    ///         SecurityMode = "none",
    ///         Servers = new[]
    ///         {
    ///             new Fortios.System.Inputs.LinkmonitorServerArgs
    ///             {
    ///                 Address = "3.3.3.3",
    ///             },
    ///         },
    ///         SourceIp = "0.0.0.0",
    ///         SourceIp6 = "::",
    ///         Srcintf = "port4",
    ///         Status = "enable",
    ///         UpdateCascadeInterface = "enable",
    ///         UpdateStaticRoute = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System LinkMonitor can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/linkmonitor:Linkmonitor labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/linkmonitor:Linkmonitor labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/linkmonitor:Linkmonitor")]
    public partial class Linkmonitor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Output("addrMode")]
        public Output<string> AddrMode { get; private set; } = null!;

        /// <summary>
        /// Traffic class ID.
        /// </summary>
        [Output("classId")]
        public Output<int> ClassId { get; private set; } = null!;

        /// <summary>
        /// Differentiated services code point (DSCP) in the IP header of the probe packet.
        /// </summary>
        [Output("diffservcode")]
        public Output<string> Diffservcode { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Threshold weight to trigger link failure alert.
        /// </summary>
        [Output("failWeight")]
        public Output<int> FailWeight { get; private set; } = null!;

        /// <summary>
        /// Number of retry attempts before the server is considered down (default = 5). On FortiOS versions 6.2.0-7.0.5: 1 - 10. On FortiOS versions &gt;= 7.0.6: 1 - 3600.
        /// </summary>
        [Output("failtime")]
        public Output<int> Failtime { get; private set; } = null!;

        /// <summary>
        /// Gateway IP address used to probe the server.
        /// </summary>
        [Output("gatewayIp")]
        public Output<string> GatewayIp { get; private set; } = null!;

        /// <summary>
        /// Gateway IPv6 address used to probe the server.
        /// </summary>
        [Output("gatewayIp6")]
        public Output<string> GatewayIp6 { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// HA election priority (1 - 50).
        /// </summary>
        [Output("haPriority")]
        public Output<int> HaPriority { get; private set; } = null!;

        /// <summary>
        /// String in the http-agent field in the HTTP header.
        /// </summary>
        [Output("httpAgent")]
        public Output<string> HttpAgent { get; private set; } = null!;

        /// <summary>
        /// If you are monitoring an HTML server you can send an HTTP-GET request with a custom string. Use this option to define the string.
        /// </summary>
        [Output("httpGet")]
        public Output<string> HttpGet { get; private set; } = null!;

        /// <summary>
        /// String that you expect to see in the HTTP-GET requests of the traffic to be monitored.
        /// </summary>
        [Output("httpMatch")]
        public Output<string> HttpMatch { get; private set; } = null!;

        /// <summary>
        /// Detection interval. On FortiOS versions 6.2.0: 1 - 3600 sec, default = 5. On FortiOS versions 6.2.4-7.0.10, 7.2.0-7.2.4: 500 - 3600 * 1000 msec, default = 500. On FortiOS versions 7.0.11-7.0.15, &gt;= 7.2.6: 20 - 3600 * 1000 msec, default = 500.
        /// </summary>
        [Output("interval")]
        public Output<int> Interval { get; private set; } = null!;

        /// <summary>
        /// Link monitor name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Packet size of a TWAMP test session.
        /// </summary>
        [Output("packetSize")]
        public Output<int> PacketSize { get; private set; } = null!;

        /// <summary>
        /// Twamp controller password in authentication mode
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Port number of the traffic to be used to monitor the server.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
        /// </summary>
        [Output("probeCount")]
        public Output<int> ProbeCount { get; private set; } = null!;

        /// <summary>
        /// Time to wait before a probe packet is considered lost (default = 500). On FortiOS versions 6.2.4-7.0.10, 7.2.0-7.2.4: 500 - 5000 msec. On FortiOS versions 7.0.11-7.0.15, &gt;= 7.2.6: 20 - 5000 msec.
        /// </summary>
        [Output("probeTimeout")]
        public Output<int> ProbeTimeout { get; private set; } = null!;

        /// <summary>
        /// Protocols used to monitor the server.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Number of successful responses received before server is considered recovered (default = 5). On FortiOS versions 6.2.0-7.0.5: 1 - 10. On FortiOS versions &gt;= 7.0.6: 1 - 3600.
        /// </summary>
        [Output("recoverytime")]
        public Output<int> Recoverytime { get; private set; } = null!;

        /// <summary>
        /// Subnet to monitor. The structure of `route` block is documented below.
        /// </summary>
        [Output("routes")]
        public Output<ImmutableArray<Outputs.LinkmonitorRoute>> Routes { get; private set; } = null!;

        /// <summary>
        /// Twamp controller security mode. Valid values: `none`, `authentication`.
        /// </summary>
        [Output("securityMode")]
        public Output<string> SecurityMode { get; private set; } = null!;

        /// <summary>
        /// Mode of server configuration. Valid values: `default`, `individual`.
        /// </summary>
        [Output("serverConfig")]
        public Output<string> ServerConfig { get; private set; } = null!;

        /// <summary>
        /// Servers for link-monitor to monitor. The structure of `server_list` block is documented below.
        /// </summary>
        [Output("serverLists")]
        public Output<ImmutableArray<Outputs.LinkmonitorServerList>> ServerLists { get; private set; } = null!;

        /// <summary>
        /// Server type (static or dynamic). Valid values: `static`, `dynamic`.
        /// </summary>
        [Output("serverType")]
        public Output<string> ServerType { get; private set; } = null!;

        /// <summary>
        /// IP address of the server(s) to be monitored. The structure of `server` block is documented below.
        /// </summary>
        [Output("servers")]
        public Output<ImmutableArray<Outputs.LinkmonitorServer>> Servers { get; private set; } = null!;

        /// <summary>
        /// Only use monitor to read quality values. If enabled, static routes and cascade interfaces will not be updated. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("serviceDetection")]
        public Output<string> ServiceDetection { get; private set; } = null!;

        /// <summary>
        /// Source IP address used in packet to the server.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Source IPv6 address used in packet to the server.
        /// </summary>
        [Output("sourceIp6")]
        public Output<string> SourceIp6 { get; private set; } = null!;

        /// <summary>
        /// Interface that receives the traffic to be monitored.
        /// </summary>
        [Output("srcintf")]
        public Output<string> Srcintf { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this link monitor. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Enable/disable update cascade interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("updateCascadeInterface")]
        public Output<string> UpdateCascadeInterface { get; private set; } = null!;

        /// <summary>
        /// Enable/disable updating the policy route. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("updatePolicyRoute")]
        public Output<string> UpdatePolicyRoute { get; private set; } = null!;

        /// <summary>
        /// Enable/disable updating the static route. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("updateStaticRoute")]
        public Output<string> UpdateStaticRoute { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Linkmonitor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Linkmonitor(string name, LinkmonitorArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/linkmonitor:Linkmonitor", name, args ?? new LinkmonitorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Linkmonitor(string name, Input<string> id, LinkmonitorState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/linkmonitor:Linkmonitor", name, state, MakeResourceOptions(options, id))
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
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Linkmonitor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Linkmonitor Get(string name, Input<string> id, LinkmonitorState? state = null, CustomResourceOptions? options = null)
        {
            return new Linkmonitor(name, id, state, options);
        }
    }

    public sealed class LinkmonitorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("addrMode")]
        public Input<string>? AddrMode { get; set; }

        /// <summary>
        /// Traffic class ID.
        /// </summary>
        [Input("classId")]
        public Input<int>? ClassId { get; set; }

        /// <summary>
        /// Differentiated services code point (DSCP) in the IP header of the probe packet.
        /// </summary>
        [Input("diffservcode")]
        public Input<string>? Diffservcode { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Threshold weight to trigger link failure alert.
        /// </summary>
        [Input("failWeight")]
        public Input<int>? FailWeight { get; set; }

        /// <summary>
        /// Number of retry attempts before the server is considered down (default = 5). On FortiOS versions 6.2.0-7.0.5: 1 - 10. On FortiOS versions &gt;= 7.0.6: 1 - 3600.
        /// </summary>
        [Input("failtime")]
        public Input<int>? Failtime { get; set; }

        /// <summary>
        /// Gateway IP address used to probe the server.
        /// </summary>
        [Input("gatewayIp")]
        public Input<string>? GatewayIp { get; set; }

        /// <summary>
        /// Gateway IPv6 address used to probe the server.
        /// </summary>
        [Input("gatewayIp6")]
        public Input<string>? GatewayIp6 { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// HA election priority (1 - 50).
        /// </summary>
        [Input("haPriority")]
        public Input<int>? HaPriority { get; set; }

        /// <summary>
        /// String in the http-agent field in the HTTP header.
        /// </summary>
        [Input("httpAgent")]
        public Input<string>? HttpAgent { get; set; }

        /// <summary>
        /// If you are monitoring an HTML server you can send an HTTP-GET request with a custom string. Use this option to define the string.
        /// </summary>
        [Input("httpGet")]
        public Input<string>? HttpGet { get; set; }

        /// <summary>
        /// String that you expect to see in the HTTP-GET requests of the traffic to be monitored.
        /// </summary>
        [Input("httpMatch")]
        public Input<string>? HttpMatch { get; set; }

        /// <summary>
        /// Detection interval. On FortiOS versions 6.2.0: 1 - 3600 sec, default = 5. On FortiOS versions 6.2.4-7.0.10, 7.2.0-7.2.4: 500 - 3600 * 1000 msec, default = 500. On FortiOS versions 7.0.11-7.0.15, &gt;= 7.2.6: 20 - 3600 * 1000 msec, default = 500.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// Link monitor name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Packet size of a TWAMP test session.
        /// </summary>
        [Input("packetSize")]
        public Input<int>? PacketSize { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Twamp controller password in authentication mode
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port number of the traffic to be used to monitor the server.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
        /// </summary>
        [Input("probeCount")]
        public Input<int>? ProbeCount { get; set; }

        /// <summary>
        /// Time to wait before a probe packet is considered lost (default = 500). On FortiOS versions 6.2.4-7.0.10, 7.2.0-7.2.4: 500 - 5000 msec. On FortiOS versions 7.0.11-7.0.15, &gt;= 7.2.6: 20 - 5000 msec.
        /// </summary>
        [Input("probeTimeout")]
        public Input<int>? ProbeTimeout { get; set; }

        /// <summary>
        /// Protocols used to monitor the server.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Number of successful responses received before server is considered recovered (default = 5). On FortiOS versions 6.2.0-7.0.5: 1 - 10. On FortiOS versions &gt;= 7.0.6: 1 - 3600.
        /// </summary>
        [Input("recoverytime")]
        public Input<int>? Recoverytime { get; set; }

        [Input("routes")]
        private InputList<Inputs.LinkmonitorRouteArgs>? _routes;

        /// <summary>
        /// Subnet to monitor. The structure of `route` block is documented below.
        /// </summary>
        public InputList<Inputs.LinkmonitorRouteArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.LinkmonitorRouteArgs>());
            set => _routes = value;
        }

        /// <summary>
        /// Twamp controller security mode. Valid values: `none`, `authentication`.
        /// </summary>
        [Input("securityMode")]
        public Input<string>? SecurityMode { get; set; }

        /// <summary>
        /// Mode of server configuration. Valid values: `default`, `individual`.
        /// </summary>
        [Input("serverConfig")]
        public Input<string>? ServerConfig { get; set; }

        [Input("serverLists")]
        private InputList<Inputs.LinkmonitorServerListArgs>? _serverLists;

        /// <summary>
        /// Servers for link-monitor to monitor. The structure of `server_list` block is documented below.
        /// </summary>
        public InputList<Inputs.LinkmonitorServerListArgs> ServerLists
        {
            get => _serverLists ?? (_serverLists = new InputList<Inputs.LinkmonitorServerListArgs>());
            set => _serverLists = value;
        }

        /// <summary>
        /// Server type (static or dynamic). Valid values: `static`, `dynamic`.
        /// </summary>
        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        [Input("servers", required: true)]
        private InputList<Inputs.LinkmonitorServerArgs>? _servers;

        /// <summary>
        /// IP address of the server(s) to be monitored. The structure of `server` block is documented below.
        /// </summary>
        public InputList<Inputs.LinkmonitorServerArgs> Servers
        {
            get => _servers ?? (_servers = new InputList<Inputs.LinkmonitorServerArgs>());
            set => _servers = value;
        }

        /// <summary>
        /// Only use monitor to read quality values. If enabled, static routes and cascade interfaces will not be updated. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("serviceDetection")]
        public Input<string>? ServiceDetection { get; set; }

        /// <summary>
        /// Source IP address used in packet to the server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Source IPv6 address used in packet to the server.
        /// </summary>
        [Input("sourceIp6")]
        public Input<string>? SourceIp6 { get; set; }

        /// <summary>
        /// Interface that receives the traffic to be monitored.
        /// </summary>
        [Input("srcintf")]
        public Input<string>? Srcintf { get; set; }

        /// <summary>
        /// Enable/disable this link monitor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Enable/disable update cascade interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("updateCascadeInterface")]
        public Input<string>? UpdateCascadeInterface { get; set; }

        /// <summary>
        /// Enable/disable updating the policy route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("updatePolicyRoute")]
        public Input<string>? UpdatePolicyRoute { get; set; }

        /// <summary>
        /// Enable/disable updating the static route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("updateStaticRoute")]
        public Input<string>? UpdateStaticRoute { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public LinkmonitorArgs()
        {
        }
        public static new LinkmonitorArgs Empty => new LinkmonitorArgs();
    }

    public sealed class LinkmonitorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("addrMode")]
        public Input<string>? AddrMode { get; set; }

        /// <summary>
        /// Traffic class ID.
        /// </summary>
        [Input("classId")]
        public Input<int>? ClassId { get; set; }

        /// <summary>
        /// Differentiated services code point (DSCP) in the IP header of the probe packet.
        /// </summary>
        [Input("diffservcode")]
        public Input<string>? Diffservcode { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Threshold weight to trigger link failure alert.
        /// </summary>
        [Input("failWeight")]
        public Input<int>? FailWeight { get; set; }

        /// <summary>
        /// Number of retry attempts before the server is considered down (default = 5). On FortiOS versions 6.2.0-7.0.5: 1 - 10. On FortiOS versions &gt;= 7.0.6: 1 - 3600.
        /// </summary>
        [Input("failtime")]
        public Input<int>? Failtime { get; set; }

        /// <summary>
        /// Gateway IP address used to probe the server.
        /// </summary>
        [Input("gatewayIp")]
        public Input<string>? GatewayIp { get; set; }

        /// <summary>
        /// Gateway IPv6 address used to probe the server.
        /// </summary>
        [Input("gatewayIp6")]
        public Input<string>? GatewayIp6 { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// HA election priority (1 - 50).
        /// </summary>
        [Input("haPriority")]
        public Input<int>? HaPriority { get; set; }

        /// <summary>
        /// String in the http-agent field in the HTTP header.
        /// </summary>
        [Input("httpAgent")]
        public Input<string>? HttpAgent { get; set; }

        /// <summary>
        /// If you are monitoring an HTML server you can send an HTTP-GET request with a custom string. Use this option to define the string.
        /// </summary>
        [Input("httpGet")]
        public Input<string>? HttpGet { get; set; }

        /// <summary>
        /// String that you expect to see in the HTTP-GET requests of the traffic to be monitored.
        /// </summary>
        [Input("httpMatch")]
        public Input<string>? HttpMatch { get; set; }

        /// <summary>
        /// Detection interval. On FortiOS versions 6.2.0: 1 - 3600 sec, default = 5. On FortiOS versions 6.2.4-7.0.10, 7.2.0-7.2.4: 500 - 3600 * 1000 msec, default = 500. On FortiOS versions 7.0.11-7.0.15, &gt;= 7.2.6: 20 - 3600 * 1000 msec, default = 500.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// Link monitor name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Packet size of a TWAMP test session.
        /// </summary>
        [Input("packetSize")]
        public Input<int>? PacketSize { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Twamp controller password in authentication mode
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port number of the traffic to be used to monitor the server.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
        /// </summary>
        [Input("probeCount")]
        public Input<int>? ProbeCount { get; set; }

        /// <summary>
        /// Time to wait before a probe packet is considered lost (default = 500). On FortiOS versions 6.2.4-7.0.10, 7.2.0-7.2.4: 500 - 5000 msec. On FortiOS versions 7.0.11-7.0.15, &gt;= 7.2.6: 20 - 5000 msec.
        /// </summary>
        [Input("probeTimeout")]
        public Input<int>? ProbeTimeout { get; set; }

        /// <summary>
        /// Protocols used to monitor the server.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Number of successful responses received before server is considered recovered (default = 5). On FortiOS versions 6.2.0-7.0.5: 1 - 10. On FortiOS versions &gt;= 7.0.6: 1 - 3600.
        /// </summary>
        [Input("recoverytime")]
        public Input<int>? Recoverytime { get; set; }

        [Input("routes")]
        private InputList<Inputs.LinkmonitorRouteGetArgs>? _routes;

        /// <summary>
        /// Subnet to monitor. The structure of `route` block is documented below.
        /// </summary>
        public InputList<Inputs.LinkmonitorRouteGetArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.LinkmonitorRouteGetArgs>());
            set => _routes = value;
        }

        /// <summary>
        /// Twamp controller security mode. Valid values: `none`, `authentication`.
        /// </summary>
        [Input("securityMode")]
        public Input<string>? SecurityMode { get; set; }

        /// <summary>
        /// Mode of server configuration. Valid values: `default`, `individual`.
        /// </summary>
        [Input("serverConfig")]
        public Input<string>? ServerConfig { get; set; }

        [Input("serverLists")]
        private InputList<Inputs.LinkmonitorServerListGetArgs>? _serverLists;

        /// <summary>
        /// Servers for link-monitor to monitor. The structure of `server_list` block is documented below.
        /// </summary>
        public InputList<Inputs.LinkmonitorServerListGetArgs> ServerLists
        {
            get => _serverLists ?? (_serverLists = new InputList<Inputs.LinkmonitorServerListGetArgs>());
            set => _serverLists = value;
        }

        /// <summary>
        /// Server type (static or dynamic). Valid values: `static`, `dynamic`.
        /// </summary>
        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        [Input("servers")]
        private InputList<Inputs.LinkmonitorServerGetArgs>? _servers;

        /// <summary>
        /// IP address of the server(s) to be monitored. The structure of `server` block is documented below.
        /// </summary>
        public InputList<Inputs.LinkmonitorServerGetArgs> Servers
        {
            get => _servers ?? (_servers = new InputList<Inputs.LinkmonitorServerGetArgs>());
            set => _servers = value;
        }

        /// <summary>
        /// Only use monitor to read quality values. If enabled, static routes and cascade interfaces will not be updated. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("serviceDetection")]
        public Input<string>? ServiceDetection { get; set; }

        /// <summary>
        /// Source IP address used in packet to the server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Source IPv6 address used in packet to the server.
        /// </summary>
        [Input("sourceIp6")]
        public Input<string>? SourceIp6 { get; set; }

        /// <summary>
        /// Interface that receives the traffic to be monitored.
        /// </summary>
        [Input("srcintf")]
        public Input<string>? Srcintf { get; set; }

        /// <summary>
        /// Enable/disable this link monitor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Enable/disable update cascade interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("updateCascadeInterface")]
        public Input<string>? UpdateCascadeInterface { get; set; }

        /// <summary>
        /// Enable/disable updating the policy route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("updatePolicyRoute")]
        public Input<string>? UpdatePolicyRoute { get; set; }

        /// <summary>
        /// Enable/disable updating the static route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("updateStaticRoute")]
        public Input<string>? UpdateStaticRoute { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public LinkmonitorState()
        {
        }
        public static new LinkmonitorState Empty => new LinkmonitorState();
    }
}
