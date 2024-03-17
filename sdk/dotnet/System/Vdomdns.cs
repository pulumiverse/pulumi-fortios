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
    /// Configure DNS servers for a non-management VDOM.
    /// 
    /// ## Import
    /// 
    /// System VdomDns can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/vdomdns:Vdomdns labelname SystemVdomDns
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/vdomdns:Vdomdns labelname SystemVdomDns
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/vdomdns:Vdomdns")]
    public partial class Vdomdns : global::Pulumi.CustomResource
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
        /// Enable/disable/enforce DNS over TLS. Valid values: `disable`, `enable`, `enforce`.
        /// </summary>
        [Output("dnsOverTls")]
        public Output<string> DnsOverTls { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

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
        /// Primary IPv6 DNS server IP address for the VDOM.
        /// </summary>
        [Output("ip6Primary")]
        public Output<string> Ip6Primary { get; private set; } = null!;

        /// <summary>
        /// Secondary IPv6 DNS server IP address for the VDOM.
        /// </summary>
        [Output("ip6Secondary")]
        public Output<string> Ip6Secondary { get; private set; } = null!;

        /// <summary>
        /// Primary DNS server IP address for the VDOM.
        /// </summary>
        [Output("primary")]
        public Output<string> Primary { get; private set; } = null!;

        /// <summary>
        /// DNS protocols. Valid values: `cleartext`, `dot`, `doh`.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Secondary DNS server IP address for the VDOM.
        /// </summary>
        [Output("secondary")]
        public Output<string> Secondary { get; private set; } = null!;

        /// <summary>
        /// DNS server host name list. The structure of `server_hostname` block is documented below.
        /// </summary>
        [Output("serverHostnames")]
        public Output<ImmutableArray<Outputs.VdomdnsServerHostname>> ServerHostnames { get; private set; } = null!;

        /// <summary>
        /// Specify how configured servers are prioritized. Valid values: `least-rtt`, `failover`.
        /// </summary>
        [Output("serverSelectMethod")]
        public Output<string> ServerSelectMethod { get; private set; } = null!;

        /// <summary>
        /// Source IP for communications with the DNS server.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Name of local certificate for SSL connections.
        /// </summary>
        [Output("sslCertificate")]
        public Output<string> SslCertificate { get; private set; } = null!;

        /// <summary>
        /// Enable/disable configuring DNS servers for the current VDOM. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("vdomDns")]
        public Output<string> VdomDns { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Vdomdns resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vdomdns(string name, VdomdnsArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/vdomdns:Vdomdns", name, args ?? new VdomdnsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vdomdns(string name, Input<string> id, VdomdnsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/vdomdns:Vdomdns", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vdomdns resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vdomdns Get(string name, Input<string> id, VdomdnsState? state = null, CustomResourceOptions? options = null)
        {
            return new Vdomdns(name, id, state, options);
        }
    }

    public sealed class VdomdnsArgs : global::Pulumi.ResourceArgs
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
        /// Enable/disable/enforce DNS over TLS. Valid values: `disable`, `enable`, `enforce`.
        /// </summary>
        [Input("dnsOverTls")]
        public Input<string>? DnsOverTls { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

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
        /// Primary IPv6 DNS server IP address for the VDOM.
        /// </summary>
        [Input("ip6Primary")]
        public Input<string>? Ip6Primary { get; set; }

        /// <summary>
        /// Secondary IPv6 DNS server IP address for the VDOM.
        /// </summary>
        [Input("ip6Secondary")]
        public Input<string>? Ip6Secondary { get; set; }

        /// <summary>
        /// Primary DNS server IP address for the VDOM.
        /// </summary>
        [Input("primary")]
        public Input<string>? Primary { get; set; }

        /// <summary>
        /// DNS protocols. Valid values: `cleartext`, `dot`, `doh`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Secondary DNS server IP address for the VDOM.
        /// </summary>
        [Input("secondary")]
        public Input<string>? Secondary { get; set; }

        [Input("serverHostnames")]
        private InputList<Inputs.VdomdnsServerHostnameArgs>? _serverHostnames;

        /// <summary>
        /// DNS server host name list. The structure of `server_hostname` block is documented below.
        /// </summary>
        public InputList<Inputs.VdomdnsServerHostnameArgs> ServerHostnames
        {
            get => _serverHostnames ?? (_serverHostnames = new InputList<Inputs.VdomdnsServerHostnameArgs>());
            set => _serverHostnames = value;
        }

        /// <summary>
        /// Specify how configured servers are prioritized. Valid values: `least-rtt`, `failover`.
        /// </summary>
        [Input("serverSelectMethod")]
        public Input<string>? ServerSelectMethod { get; set; }

        /// <summary>
        /// Source IP for communications with the DNS server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Name of local certificate for SSL connections.
        /// </summary>
        [Input("sslCertificate")]
        public Input<string>? SslCertificate { get; set; }

        /// <summary>
        /// Enable/disable configuring DNS servers for the current VDOM. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("vdomDns")]
        public Input<string>? VdomDns { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public VdomdnsArgs()
        {
        }
        public static new VdomdnsArgs Empty => new VdomdnsArgs();
    }

    public sealed class VdomdnsState : global::Pulumi.ResourceArgs
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
        /// Enable/disable/enforce DNS over TLS. Valid values: `disable`, `enable`, `enforce`.
        /// </summary>
        [Input("dnsOverTls")]
        public Input<string>? DnsOverTls { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

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
        /// Primary IPv6 DNS server IP address for the VDOM.
        /// </summary>
        [Input("ip6Primary")]
        public Input<string>? Ip6Primary { get; set; }

        /// <summary>
        /// Secondary IPv6 DNS server IP address for the VDOM.
        /// </summary>
        [Input("ip6Secondary")]
        public Input<string>? Ip6Secondary { get; set; }

        /// <summary>
        /// Primary DNS server IP address for the VDOM.
        /// </summary>
        [Input("primary")]
        public Input<string>? Primary { get; set; }

        /// <summary>
        /// DNS protocols. Valid values: `cleartext`, `dot`, `doh`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Secondary DNS server IP address for the VDOM.
        /// </summary>
        [Input("secondary")]
        public Input<string>? Secondary { get; set; }

        [Input("serverHostnames")]
        private InputList<Inputs.VdomdnsServerHostnameGetArgs>? _serverHostnames;

        /// <summary>
        /// DNS server host name list. The structure of `server_hostname` block is documented below.
        /// </summary>
        public InputList<Inputs.VdomdnsServerHostnameGetArgs> ServerHostnames
        {
            get => _serverHostnames ?? (_serverHostnames = new InputList<Inputs.VdomdnsServerHostnameGetArgs>());
            set => _serverHostnames = value;
        }

        /// <summary>
        /// Specify how configured servers are prioritized. Valid values: `least-rtt`, `failover`.
        /// </summary>
        [Input("serverSelectMethod")]
        public Input<string>? ServerSelectMethod { get; set; }

        /// <summary>
        /// Source IP for communications with the DNS server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Name of local certificate for SSL connections.
        /// </summary>
        [Input("sslCertificate")]
        public Input<string>? SslCertificate { get; set; }

        /// <summary>
        /// Enable/disable configuring DNS servers for the current VDOM. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("vdomDns")]
        public Input<string>? VdomDns { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public VdomdnsState()
        {
        }
        public static new VdomdnsState Empty => new VdomdnsState();
    }
}
