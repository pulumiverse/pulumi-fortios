// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Webproxy
{
    /// <summary>
    /// Configure explicit Web proxy settings.
    /// 
    /// ## Import
    /// 
    /// WebProxy Explicit can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:webproxy/explicit:Explicit labelname WebProxyExplicit
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:webproxy/explicit:Explicit labelname WebProxyExplicit
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:webproxy/explicit:Explicit")]
    public partial class Explicit : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
        /// </summary>
        [Output("ftpIncomingPort")]
        public Output<string> FtpIncomingPort { get; private set; } = null!;

        /// <summary>
        /// Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ftpOverHttp")]
        public Output<string> FtpOverHttp { get; private set; } = null!;

        /// <summary>
        /// Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
        /// </summary>
        [Output("httpIncomingPort")]
        public Output<string> HttpIncomingPort { get; private set; } = null!;

        /// <summary>
        /// Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
        /// </summary>
        [Output("httpsIncomingPort")]
        public Output<string> HttpsIncomingPort { get; private set; } = null!;

        /// <summary>
        /// Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("httpsReplacementMessage")]
        public Output<string> HttpsReplacementMessage { get; private set; } = null!;

        /// <summary>
        /// Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
        /// </summary>
        [Output("incomingIp")]
        public Output<string> IncomingIp { get; private set; } = null!;

        /// <summary>
        /// Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
        /// </summary>
        [Output("incomingIp6")]
        public Output<string> IncomingIp6 { get; private set; } = null!;

        /// <summary>
        /// Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipv6Status")]
        public Output<string> Ipv6Status { get; private set; } = null!;

        /// <summary>
        /// Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("messageUponServerError")]
        public Output<string> MessageUponServerError { get; private set; } = null!;

        /// <summary>
        /// Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
        /// </summary>
        [Output("outgoingIp")]
        public Output<string> OutgoingIp { get; private set; } = null!;

        /// <summary>
        /// Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
        /// </summary>
        [Output("outgoingIp6")]
        public Output<string> OutgoingIp6 { get; private set; } = null!;

        /// <summary>
        /// PAC file contents enclosed in quotes (maximum of 256K bytes).
        /// </summary>
        [Output("pacFileData")]
        public Output<string> PacFileData { get; private set; } = null!;

        /// <summary>
        /// Pac file name.
        /// </summary>
        [Output("pacFileName")]
        public Output<string> PacFileName { get; private set; } = null!;

        /// <summary>
        /// Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
        /// </summary>
        [Output("pacFileServerPort")]
        public Output<string> PacFileServerPort { get; private set; } = null!;

        /// <summary>
        /// Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("pacFileServerStatus")]
        public Output<string> PacFileServerStatus { get; private set; } = null!;

        /// <summary>
        /// Enable/disable to get Proxy Auto-Configuration (PAC) through HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("pacFileThroughHttps")]
        public Output<string> PacFileThroughHttps { get; private set; } = null!;

        /// <summary>
        /// PAC file access URL.
        /// </summary>
        [Output("pacFileUrl")]
        public Output<string> PacFileUrl { get; private set; } = null!;

        /// <summary>
        /// PAC policies. The structure of `pac_policy` block is documented below.
        /// </summary>
        [Output("pacPolicies")]
        public Output<ImmutableArray<Outputs.ExplicitPacPolicy>> PacPolicies { get; private set; } = null!;

        /// <summary>
        /// Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Output("prefDnsResult")]
        public Output<string> PrefDnsResult { get; private set; } = null!;

        /// <summary>
        /// Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
        /// </summary>
        [Output("realm")]
        public Output<string> Realm { get; private set; } = null!;

        /// <summary>
        /// Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept`, `deny`.
        /// </summary>
        [Output("secDefaultAction")]
        public Output<string> SecDefaultAction { get; private set; } = null!;

        /// <summary>
        /// Enable/disable the SOCKS proxy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("socks")]
        public Output<string> Socks { get; private set; } = null!;

        /// <summary>
        /// Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
        /// </summary>
        [Output("socksIncomingPort")]
        public Output<string> SocksIncomingPort { get; private set; } = null!;

        /// <summary>
        /// Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.
        /// </summary>
        [Output("sslAlgorithm")]
        public Output<string> SslAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Enable/disable the explicit Web proxy for HTTP and HTTPS session. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("strictGuest")]
        public Output<string> StrictGuest { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging timed-out authentication requests. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("traceAuthNoRsp")]
        public Output<string> TraceAuthNoRsp { get; private set; } = null!;

        /// <summary>
        /// Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
        /// </summary>
        [Output("unknownHttpVersion")]
        public Output<string> UnknownHttpVersion { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Explicit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Explicit(string name, ExplicitArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:webproxy/explicit:Explicit", name, args ?? new ExplicitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Explicit(string name, Input<string> id, ExplicitState? state = null, CustomResourceOptions? options = null)
            : base("fortios:webproxy/explicit:Explicit", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Explicit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Explicit Get(string name, Input<string> id, ExplicitState? state = null, CustomResourceOptions? options = null)
        {
            return new Explicit(name, id, state, options);
        }
    }

    public sealed class ExplicitArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
        /// </summary>
        [Input("ftpIncomingPort")]
        public Input<string>? FtpIncomingPort { get; set; }

        /// <summary>
        /// Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ftpOverHttp")]
        public Input<string>? FtpOverHttp { get; set; }

        /// <summary>
        /// Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
        /// </summary>
        [Input("httpIncomingPort")]
        public Input<string>? HttpIncomingPort { get; set; }

        /// <summary>
        /// Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
        /// </summary>
        [Input("httpsIncomingPort")]
        public Input<string>? HttpsIncomingPort { get; set; }

        /// <summary>
        /// Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("httpsReplacementMessage")]
        public Input<string>? HttpsReplacementMessage { get; set; }

        /// <summary>
        /// Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
        /// </summary>
        [Input("incomingIp")]
        public Input<string>? IncomingIp { get; set; }

        /// <summary>
        /// Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
        /// </summary>
        [Input("incomingIp6")]
        public Input<string>? IncomingIp6 { get; set; }

        /// <summary>
        /// Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipv6Status")]
        public Input<string>? Ipv6Status { get; set; }

        /// <summary>
        /// Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("messageUponServerError")]
        public Input<string>? MessageUponServerError { get; set; }

        /// <summary>
        /// Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
        /// </summary>
        [Input("outgoingIp")]
        public Input<string>? OutgoingIp { get; set; }

        /// <summary>
        /// Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
        /// </summary>
        [Input("outgoingIp6")]
        public Input<string>? OutgoingIp6 { get; set; }

        /// <summary>
        /// PAC file contents enclosed in quotes (maximum of 256K bytes).
        /// </summary>
        [Input("pacFileData")]
        public Input<string>? PacFileData { get; set; }

        /// <summary>
        /// Pac file name.
        /// </summary>
        [Input("pacFileName")]
        public Input<string>? PacFileName { get; set; }

        /// <summary>
        /// Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
        /// </summary>
        [Input("pacFileServerPort")]
        public Input<string>? PacFileServerPort { get; set; }

        /// <summary>
        /// Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pacFileServerStatus")]
        public Input<string>? PacFileServerStatus { get; set; }

        /// <summary>
        /// Enable/disable to get Proxy Auto-Configuration (PAC) through HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pacFileThroughHttps")]
        public Input<string>? PacFileThroughHttps { get; set; }

        /// <summary>
        /// PAC file access URL.
        /// </summary>
        [Input("pacFileUrl")]
        public Input<string>? PacFileUrl { get; set; }

        [Input("pacPolicies")]
        private InputList<Inputs.ExplicitPacPolicyArgs>? _pacPolicies;

        /// <summary>
        /// PAC policies. The structure of `pac_policy` block is documented below.
        /// </summary>
        public InputList<Inputs.ExplicitPacPolicyArgs> PacPolicies
        {
            get => _pacPolicies ?? (_pacPolicies = new InputList<Inputs.ExplicitPacPolicyArgs>());
            set => _pacPolicies = value;
        }

        /// <summary>
        /// Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("prefDnsResult")]
        public Input<string>? PrefDnsResult { get; set; }

        /// <summary>
        /// Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
        /// </summary>
        [Input("realm")]
        public Input<string>? Realm { get; set; }

        /// <summary>
        /// Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept`, `deny`.
        /// </summary>
        [Input("secDefaultAction")]
        public Input<string>? SecDefaultAction { get; set; }

        /// <summary>
        /// Enable/disable the SOCKS proxy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("socks")]
        public Input<string>? Socks { get; set; }

        /// <summary>
        /// Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
        /// </summary>
        [Input("socksIncomingPort")]
        public Input<string>? SocksIncomingPort { get; set; }

        /// <summary>
        /// Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.
        /// </summary>
        [Input("sslAlgorithm")]
        public Input<string>? SslAlgorithm { get; set; }

        /// <summary>
        /// Enable/disable the explicit Web proxy for HTTP and HTTPS session. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("strictGuest")]
        public Input<string>? StrictGuest { get; set; }

        /// <summary>
        /// Enable/disable logging timed-out authentication requests. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("traceAuthNoRsp")]
        public Input<string>? TraceAuthNoRsp { get; set; }

        /// <summary>
        /// Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
        /// </summary>
        [Input("unknownHttpVersion")]
        public Input<string>? UnknownHttpVersion { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ExplicitArgs()
        {
        }
        public static new ExplicitArgs Empty => new ExplicitArgs();
    }

    public sealed class ExplicitState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
        /// </summary>
        [Input("ftpIncomingPort")]
        public Input<string>? FtpIncomingPort { get; set; }

        /// <summary>
        /// Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ftpOverHttp")]
        public Input<string>? FtpOverHttp { get; set; }

        /// <summary>
        /// Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
        /// </summary>
        [Input("httpIncomingPort")]
        public Input<string>? HttpIncomingPort { get; set; }

        /// <summary>
        /// Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
        /// </summary>
        [Input("httpsIncomingPort")]
        public Input<string>? HttpsIncomingPort { get; set; }

        /// <summary>
        /// Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("httpsReplacementMessage")]
        public Input<string>? HttpsReplacementMessage { get; set; }

        /// <summary>
        /// Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
        /// </summary>
        [Input("incomingIp")]
        public Input<string>? IncomingIp { get; set; }

        /// <summary>
        /// Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
        /// </summary>
        [Input("incomingIp6")]
        public Input<string>? IncomingIp6 { get; set; }

        /// <summary>
        /// Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipv6Status")]
        public Input<string>? Ipv6Status { get; set; }

        /// <summary>
        /// Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("messageUponServerError")]
        public Input<string>? MessageUponServerError { get; set; }

        /// <summary>
        /// Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
        /// </summary>
        [Input("outgoingIp")]
        public Input<string>? OutgoingIp { get; set; }

        /// <summary>
        /// Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
        /// </summary>
        [Input("outgoingIp6")]
        public Input<string>? OutgoingIp6 { get; set; }

        /// <summary>
        /// PAC file contents enclosed in quotes (maximum of 256K bytes).
        /// </summary>
        [Input("pacFileData")]
        public Input<string>? PacFileData { get; set; }

        /// <summary>
        /// Pac file name.
        /// </summary>
        [Input("pacFileName")]
        public Input<string>? PacFileName { get; set; }

        /// <summary>
        /// Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
        /// </summary>
        [Input("pacFileServerPort")]
        public Input<string>? PacFileServerPort { get; set; }

        /// <summary>
        /// Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pacFileServerStatus")]
        public Input<string>? PacFileServerStatus { get; set; }

        /// <summary>
        /// Enable/disable to get Proxy Auto-Configuration (PAC) through HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pacFileThroughHttps")]
        public Input<string>? PacFileThroughHttps { get; set; }

        /// <summary>
        /// PAC file access URL.
        /// </summary>
        [Input("pacFileUrl")]
        public Input<string>? PacFileUrl { get; set; }

        [Input("pacPolicies")]
        private InputList<Inputs.ExplicitPacPolicyGetArgs>? _pacPolicies;

        /// <summary>
        /// PAC policies. The structure of `pac_policy` block is documented below.
        /// </summary>
        public InputList<Inputs.ExplicitPacPolicyGetArgs> PacPolicies
        {
            get => _pacPolicies ?? (_pacPolicies = new InputList<Inputs.ExplicitPacPolicyGetArgs>());
            set => _pacPolicies = value;
        }

        /// <summary>
        /// Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("prefDnsResult")]
        public Input<string>? PrefDnsResult { get; set; }

        /// <summary>
        /// Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
        /// </summary>
        [Input("realm")]
        public Input<string>? Realm { get; set; }

        /// <summary>
        /// Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept`, `deny`.
        /// </summary>
        [Input("secDefaultAction")]
        public Input<string>? SecDefaultAction { get; set; }

        /// <summary>
        /// Enable/disable the SOCKS proxy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("socks")]
        public Input<string>? Socks { get; set; }

        /// <summary>
        /// Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
        /// </summary>
        [Input("socksIncomingPort")]
        public Input<string>? SocksIncomingPort { get; set; }

        /// <summary>
        /// Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.
        /// </summary>
        [Input("sslAlgorithm")]
        public Input<string>? SslAlgorithm { get; set; }

        /// <summary>
        /// Enable/disable the explicit Web proxy for HTTP and HTTPS session. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("strictGuest")]
        public Input<string>? StrictGuest { get; set; }

        /// <summary>
        /// Enable/disable logging timed-out authentication requests. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("traceAuthNoRsp")]
        public Input<string>? TraceAuthNoRsp { get; set; }

        /// <summary>
        /// Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
        /// </summary>
        [Input("unknownHttpVersion")]
        public Input<string>? UnknownHttpVersion { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ExplicitState()
        {
        }
        public static new ExplicitState Empty => new ExplicitState();
    }
}
