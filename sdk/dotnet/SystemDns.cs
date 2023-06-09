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
    /// Configure DNS.
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
    ///     var trname = new Fortios.SystemDns("trname", new()
    ///     {
    ///         CacheNotfoundResponses = "disable",
    ///         DnsCacheLimit = 5000,
    ///         DnsCacheTtl = 1800,
    ///         Ip6Primary = "::",
    ///         Ip6Secondary = "::",
    ///         Primary = "208.91.112.53",
    ///         Retry = 2,
    ///         Secondary = "208.91.112.51",
    ///         SourceIp = "0.0.0.0",
    ///         Timeout = 5,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System Dns can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemDns:SystemDns labelname SystemDns
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemDns:SystemDns labelname SystemDns
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemDns:SystemDns")]
    public partial class SystemDns : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Alternate primary DNS server. (This is not used as a failover DNS server.)
        /// </summary>
        [Output("altPrimary")]
        public Output<string> AltPrimary { get; private set; } = null!;

        /// <summary>
        /// Alternate secondary DNS server. (This is not used as a failover DNS server.)
        /// </summary>
        [Output("altSecondary")]
        public Output<string> AltSecondary { get; private set; } = null!;

        /// <summary>
        /// Enable/disable response from the DNS server when a record is not in cache. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("cacheNotfoundResponses")]
        public Output<string> CacheNotfoundResponses { get; private set; } = null!;

        /// <summary>
        /// Maximum number of records in the DNS cache.
        /// </summary>
        [Output("dnsCacheLimit")]
        public Output<int> DnsCacheLimit { get; private set; } = null!;

        /// <summary>
        /// Duration in seconds that the DNS cache retains information.
        /// </summary>
        [Output("dnsCacheTtl")]
        public Output<int> DnsCacheTtl { get; private set; } = null!;

        /// <summary>
        /// Enable/disable/enforce DNS over TLS. Valid values: `disable`, `enable`, `enforce`.
        /// </summary>
        [Output("dnsOverTls")]
        public Output<string> DnsOverTls { get; private set; } = null!;

        /// <summary>
        /// Search suffix list for hostname lookup. The structure of `domain` block is documented below.
        /// </summary>
        [Output("domains")]
        public Output<ImmutableArray<Outputs.SystemDnsDomain>> Domains { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// FQDN cache time to live in seconds (0 - 86400, default = 0).
        /// </summary>
        [Output("fqdnCacheTtl")]
        public Output<int> FqdnCacheTtl { get; private set; } = null!;

        /// <summary>
        /// FQDN cache minimum refresh time in seconds (10 - 3600, default = 60).
        /// </summary>
        [Output("fqdnMinRefresh")]
        public Output<int> FqdnMinRefresh { get; private set; } = null!;

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Output("interfaceSelectMethod")]
        public Output<string> InterfaceSelectMethod { get; private set; } = null!;

        /// <summary>
        /// Primary DNS server IPv6 address.
        /// </summary>
        [Output("ip6Primary")]
        public Output<string> Ip6Primary { get; private set; } = null!;

        /// <summary>
        /// Secondary DNS server IPv6 address.
        /// </summary>
        [Output("ip6Secondary")]
        public Output<string> Ip6Secondary { get; private set; } = null!;

        /// <summary>
        /// Local DNS log setting. Valid values: `disable`, `error`, `all`.
        /// </summary>
        [Output("log")]
        public Output<string> Log { get; private set; } = null!;

        /// <summary>
        /// Primary DNS server IP address.
        /// </summary>
        [Output("primary")]
        public Output<string> Primary { get; private set; } = null!;

        /// <summary>
        /// DNS protocols. Valid values: `cleartext`, `dot`, `doh`.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Number of times to retry (0 - 5).
        /// </summary>
        [Output("retry")]
        public Output<int> Retry { get; private set; } = null!;

        /// <summary>
        /// Secondary DNS server IP address.
        /// </summary>
        [Output("secondary")]
        public Output<string> Secondary { get; private set; } = null!;

        /// <summary>
        /// DNS server host name list. The structure of `server_hostname` block is documented below.
        /// </summary>
        [Output("serverHostnames")]
        public Output<ImmutableArray<Outputs.SystemDnsServerHostname>> ServerHostnames { get; private set; } = null!;

        /// <summary>
        /// Specify how configured servers are prioritized. Valid values: `least-rtt`, `failover`.
        /// </summary>
        [Output("serverSelectMethod")]
        public Output<string> ServerSelectMethod { get; private set; } = null!;

        /// <summary>
        /// IP address used by the DNS server as its source IP.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Name of local certificate for SSL connections.
        /// </summary>
        [Output("sslCertificate")]
        public Output<string> SslCertificate { get; private set; } = null!;

        /// <summary>
        /// DNS query timeout interval in seconds (1 - 10).
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemDns resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemDns(string name, SystemDnsArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/systemDns:SystemDns", name, args ?? new SystemDnsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemDns(string name, Input<string> id, SystemDnsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemDns:SystemDns", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemDns resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemDns Get(string name, Input<string> id, SystemDnsState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemDns(name, id, state, options);
        }
    }

    public sealed class SystemDnsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alternate primary DNS server. (This is not used as a failover DNS server.)
        /// </summary>
        [Input("altPrimary")]
        public Input<string>? AltPrimary { get; set; }

        /// <summary>
        /// Alternate secondary DNS server. (This is not used as a failover DNS server.)
        /// </summary>
        [Input("altSecondary")]
        public Input<string>? AltSecondary { get; set; }

        /// <summary>
        /// Enable/disable response from the DNS server when a record is not in cache. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("cacheNotfoundResponses")]
        public Input<string>? CacheNotfoundResponses { get; set; }

        /// <summary>
        /// Maximum number of records in the DNS cache.
        /// </summary>
        [Input("dnsCacheLimit")]
        public Input<int>? DnsCacheLimit { get; set; }

        /// <summary>
        /// Duration in seconds that the DNS cache retains information.
        /// </summary>
        [Input("dnsCacheTtl")]
        public Input<int>? DnsCacheTtl { get; set; }

        /// <summary>
        /// Enable/disable/enforce DNS over TLS. Valid values: `disable`, `enable`, `enforce`.
        /// </summary>
        [Input("dnsOverTls")]
        public Input<string>? DnsOverTls { get; set; }

        [Input("domains")]
        private InputList<Inputs.SystemDnsDomainArgs>? _domains;

        /// <summary>
        /// Search suffix list for hostname lookup. The structure of `domain` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemDnsDomainArgs> Domains
        {
            get => _domains ?? (_domains = new InputList<Inputs.SystemDnsDomainArgs>());
            set => _domains = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// FQDN cache time to live in seconds (0 - 86400, default = 0).
        /// </summary>
        [Input("fqdnCacheTtl")]
        public Input<int>? FqdnCacheTtl { get; set; }

        /// <summary>
        /// FQDN cache minimum refresh time in seconds (10 - 3600, default = 60).
        /// </summary>
        [Input("fqdnMinRefresh")]
        public Input<int>? FqdnMinRefresh { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Primary DNS server IPv6 address.
        /// </summary>
        [Input("ip6Primary")]
        public Input<string>? Ip6Primary { get; set; }

        /// <summary>
        /// Secondary DNS server IPv6 address.
        /// </summary>
        [Input("ip6Secondary")]
        public Input<string>? Ip6Secondary { get; set; }

        /// <summary>
        /// Local DNS log setting. Valid values: `disable`, `error`, `all`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Primary DNS server IP address.
        /// </summary>
        [Input("primary", required: true)]
        public Input<string> Primary { get; set; } = null!;

        /// <summary>
        /// DNS protocols. Valid values: `cleartext`, `dot`, `doh`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Number of times to retry (0 - 5).
        /// </summary>
        [Input("retry")]
        public Input<int>? Retry { get; set; }

        /// <summary>
        /// Secondary DNS server IP address.
        /// </summary>
        [Input("secondary")]
        public Input<string>? Secondary { get; set; }

        [Input("serverHostnames")]
        private InputList<Inputs.SystemDnsServerHostnameArgs>? _serverHostnames;

        /// <summary>
        /// DNS server host name list. The structure of `server_hostname` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemDnsServerHostnameArgs> ServerHostnames
        {
            get => _serverHostnames ?? (_serverHostnames = new InputList<Inputs.SystemDnsServerHostnameArgs>());
            set => _serverHostnames = value;
        }

        /// <summary>
        /// Specify how configured servers are prioritized. Valid values: `least-rtt`, `failover`.
        /// </summary>
        [Input("serverSelectMethod")]
        public Input<string>? ServerSelectMethod { get; set; }

        /// <summary>
        /// IP address used by the DNS server as its source IP.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Name of local certificate for SSL connections.
        /// </summary>
        [Input("sslCertificate")]
        public Input<string>? SslCertificate { get; set; }

        /// <summary>
        /// DNS query timeout interval in seconds (1 - 10).
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemDnsArgs()
        {
        }
        public static new SystemDnsArgs Empty => new SystemDnsArgs();
    }

    public sealed class SystemDnsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alternate primary DNS server. (This is not used as a failover DNS server.)
        /// </summary>
        [Input("altPrimary")]
        public Input<string>? AltPrimary { get; set; }

        /// <summary>
        /// Alternate secondary DNS server. (This is not used as a failover DNS server.)
        /// </summary>
        [Input("altSecondary")]
        public Input<string>? AltSecondary { get; set; }

        /// <summary>
        /// Enable/disable response from the DNS server when a record is not in cache. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("cacheNotfoundResponses")]
        public Input<string>? CacheNotfoundResponses { get; set; }

        /// <summary>
        /// Maximum number of records in the DNS cache.
        /// </summary>
        [Input("dnsCacheLimit")]
        public Input<int>? DnsCacheLimit { get; set; }

        /// <summary>
        /// Duration in seconds that the DNS cache retains information.
        /// </summary>
        [Input("dnsCacheTtl")]
        public Input<int>? DnsCacheTtl { get; set; }

        /// <summary>
        /// Enable/disable/enforce DNS over TLS. Valid values: `disable`, `enable`, `enforce`.
        /// </summary>
        [Input("dnsOverTls")]
        public Input<string>? DnsOverTls { get; set; }

        [Input("domains")]
        private InputList<Inputs.SystemDnsDomainGetArgs>? _domains;

        /// <summary>
        /// Search suffix list for hostname lookup. The structure of `domain` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemDnsDomainGetArgs> Domains
        {
            get => _domains ?? (_domains = new InputList<Inputs.SystemDnsDomainGetArgs>());
            set => _domains = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// FQDN cache time to live in seconds (0 - 86400, default = 0).
        /// </summary>
        [Input("fqdnCacheTtl")]
        public Input<int>? FqdnCacheTtl { get; set; }

        /// <summary>
        /// FQDN cache minimum refresh time in seconds (10 - 3600, default = 60).
        /// </summary>
        [Input("fqdnMinRefresh")]
        public Input<int>? FqdnMinRefresh { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Primary DNS server IPv6 address.
        /// </summary>
        [Input("ip6Primary")]
        public Input<string>? Ip6Primary { get; set; }

        /// <summary>
        /// Secondary DNS server IPv6 address.
        /// </summary>
        [Input("ip6Secondary")]
        public Input<string>? Ip6Secondary { get; set; }

        /// <summary>
        /// Local DNS log setting. Valid values: `disable`, `error`, `all`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Primary DNS server IP address.
        /// </summary>
        [Input("primary")]
        public Input<string>? Primary { get; set; }

        /// <summary>
        /// DNS protocols. Valid values: `cleartext`, `dot`, `doh`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Number of times to retry (0 - 5).
        /// </summary>
        [Input("retry")]
        public Input<int>? Retry { get; set; }

        /// <summary>
        /// Secondary DNS server IP address.
        /// </summary>
        [Input("secondary")]
        public Input<string>? Secondary { get; set; }

        [Input("serverHostnames")]
        private InputList<Inputs.SystemDnsServerHostnameGetArgs>? _serverHostnames;

        /// <summary>
        /// DNS server host name list. The structure of `server_hostname` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemDnsServerHostnameGetArgs> ServerHostnames
        {
            get => _serverHostnames ?? (_serverHostnames = new InputList<Inputs.SystemDnsServerHostnameGetArgs>());
            set => _serverHostnames = value;
        }

        /// <summary>
        /// Specify how configured servers are prioritized. Valid values: `least-rtt`, `failover`.
        /// </summary>
        [Input("serverSelectMethod")]
        public Input<string>? ServerSelectMethod { get; set; }

        /// <summary>
        /// IP address used by the DNS server as its source IP.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Name of local certificate for SSL connections.
        /// </summary>
        [Input("sslCertificate")]
        public Input<string>? SslCertificate { get; set; }

        /// <summary>
        /// DNS query timeout interval in seconds (1 - 10).
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemDnsState()
        {
        }
        public static new SystemDnsState Empty => new SystemDnsState();
    }
}
