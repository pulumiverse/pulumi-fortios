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
    /// Configure server load balancing health monitors.
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
    ///     var trname = new Fortios.Firewall.Ldbmonitor("trname", new()
    ///     {
    ///         HttpMaxRedirects = 0,
    ///         Interval = 10,
    ///         Port = 0,
    ///         Retry = 3,
    ///         Timeout = 2,
    ///         Type = "ping",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Firewall LdbMonitor can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/ldbmonitor:Ldbmonitor labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/ldbmonitor:Ldbmonitor labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/ldbmonitor:Ldbmonitor")]
    public partial class Ldbmonitor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Response IP expected from DNS server.
        /// </summary>
        [Output("dnsMatchIp")]
        public Output<string> DnsMatchIp { get; private set; } = null!;

        /// <summary>
        /// Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
        /// </summary>
        [Output("dnsProtocol")]
        public Output<string> DnsProtocol { get; private set; } = null!;

        /// <summary>
        /// Fully qualified domain name to resolve for the DNS probe.
        /// </summary>
        [Output("dnsRequestDomain")]
        public Output<string> DnsRequestDomain { get; private set; } = null!;

        /// <summary>
        /// URL used to send a GET request to check the health of an HTTP server.
        /// </summary>
        [Output("httpGet")]
        public Output<string> HttpGet { get; private set; } = null!;

        /// <summary>
        /// String to match the value expected in response to an HTTP-GET request.
        /// </summary>
        [Output("httpMatch")]
        public Output<string> HttpMatch { get; private set; } = null!;

        /// <summary>
        /// The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
        /// </summary>
        [Output("httpMaxRedirects")]
        public Output<int> HttpMaxRedirects { get; private set; } = null!;

        /// <summary>
        /// Time between health checks (default = 10). On FortiOS versions 6.2.0-7.0.13: 5 - 65635 sec. On FortiOS versions &gt;= 7.2.0: 5 - 65535 sec.
        /// </summary>
        [Output("interval")]
        public Output<int> Interval { get; private set; } = null!;

        /// <summary>
        /// Monitor name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (default = 0). On FortiOS versions 6.2.0-7.0.13: 0 - 65635. On FortiOS versions &gt;= 7.2.0: 0 - 65535.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Number health check attempts before the server is considered down (1 - 255, default = 3).
        /// </summary>
        [Output("retry")]
        public Output<int> Retry { get; private set; } = null!;

        /// <summary>
        /// Source IP for ldb-monitor.
        /// </summary>
        [Output("srcIp")]
        public Output<string> SrcIp { get; private set; } = null!;

        /// <summary>
        /// Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// Select the Monitor type used by the health check monitor to check the health of the server. On FortiOS versions 6.2.0: PING | TCP | HTTP. On FortiOS versions 6.2.4-7.0.0: PING | TCP | HTTP | HTTPS. On FortiOS versions &gt;= 7.0.1: PING | TCP | HTTP | HTTPS | DNS.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ldbmonitor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ldbmonitor(string name, LdbmonitorArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/ldbmonitor:Ldbmonitor", name, args ?? new LdbmonitorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ldbmonitor(string name, Input<string> id, LdbmonitorState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/ldbmonitor:Ldbmonitor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ldbmonitor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ldbmonitor Get(string name, Input<string> id, LdbmonitorState? state = null, CustomResourceOptions? options = null)
        {
            return new Ldbmonitor(name, id, state, options);
        }
    }

    public sealed class LdbmonitorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Response IP expected from DNS server.
        /// </summary>
        [Input("dnsMatchIp")]
        public Input<string>? DnsMatchIp { get; set; }

        /// <summary>
        /// Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
        /// </summary>
        [Input("dnsProtocol")]
        public Input<string>? DnsProtocol { get; set; }

        /// <summary>
        /// Fully qualified domain name to resolve for the DNS probe.
        /// </summary>
        [Input("dnsRequestDomain")]
        public Input<string>? DnsRequestDomain { get; set; }

        /// <summary>
        /// URL used to send a GET request to check the health of an HTTP server.
        /// </summary>
        [Input("httpGet")]
        public Input<string>? HttpGet { get; set; }

        /// <summary>
        /// String to match the value expected in response to an HTTP-GET request.
        /// </summary>
        [Input("httpMatch")]
        public Input<string>? HttpMatch { get; set; }

        /// <summary>
        /// The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
        /// </summary>
        [Input("httpMaxRedirects")]
        public Input<int>? HttpMaxRedirects { get; set; }

        /// <summary>
        /// Time between health checks (default = 10). On FortiOS versions 6.2.0-7.0.13: 5 - 65635 sec. On FortiOS versions &gt;= 7.2.0: 5 - 65535 sec.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// Monitor name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (default = 0). On FortiOS versions 6.2.0-7.0.13: 0 - 65635. On FortiOS versions &gt;= 7.2.0: 0 - 65535.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Number health check attempts before the server is considered down (1 - 255, default = 3).
        /// </summary>
        [Input("retry")]
        public Input<int>? Retry { get; set; }

        /// <summary>
        /// Source IP for ldb-monitor.
        /// </summary>
        [Input("srcIp")]
        public Input<string>? SrcIp { get; set; }

        /// <summary>
        /// Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Select the Monitor type used by the health check monitor to check the health of the server. On FortiOS versions 6.2.0: PING | TCP | HTTP. On FortiOS versions 6.2.4-7.0.0: PING | TCP | HTTP | HTTPS. On FortiOS versions &gt;= 7.0.1: PING | TCP | HTTP | HTTPS | DNS.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public LdbmonitorArgs()
        {
        }
        public static new LdbmonitorArgs Empty => new LdbmonitorArgs();
    }

    public sealed class LdbmonitorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Response IP expected from DNS server.
        /// </summary>
        [Input("dnsMatchIp")]
        public Input<string>? DnsMatchIp { get; set; }

        /// <summary>
        /// Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
        /// </summary>
        [Input("dnsProtocol")]
        public Input<string>? DnsProtocol { get; set; }

        /// <summary>
        /// Fully qualified domain name to resolve for the DNS probe.
        /// </summary>
        [Input("dnsRequestDomain")]
        public Input<string>? DnsRequestDomain { get; set; }

        /// <summary>
        /// URL used to send a GET request to check the health of an HTTP server.
        /// </summary>
        [Input("httpGet")]
        public Input<string>? HttpGet { get; set; }

        /// <summary>
        /// String to match the value expected in response to an HTTP-GET request.
        /// </summary>
        [Input("httpMatch")]
        public Input<string>? HttpMatch { get; set; }

        /// <summary>
        /// The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
        /// </summary>
        [Input("httpMaxRedirects")]
        public Input<int>? HttpMaxRedirects { get; set; }

        /// <summary>
        /// Time between health checks (default = 10). On FortiOS versions 6.2.0-7.0.13: 5 - 65635 sec. On FortiOS versions &gt;= 7.2.0: 5 - 65535 sec.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// Monitor name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (default = 0). On FortiOS versions 6.2.0-7.0.13: 0 - 65635. On FortiOS versions &gt;= 7.2.0: 0 - 65535.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Number health check attempts before the server is considered down (1 - 255, default = 3).
        /// </summary>
        [Input("retry")]
        public Input<int>? Retry { get; set; }

        /// <summary>
        /// Source IP for ldb-monitor.
        /// </summary>
        [Input("srcIp")]
        public Input<string>? SrcIp { get; set; }

        /// <summary>
        /// Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Select the Monitor type used by the health check monitor to check the health of the server. On FortiOS versions 6.2.0: PING | TCP | HTTP. On FortiOS versions 6.2.4-7.0.0: PING | TCP | HTTP | HTTPS. On FortiOS versions &gt;= 7.0.1: PING | TCP | HTTP | HTTPS | DNS.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public LdbmonitorState()
        {
        }
        public static new LdbmonitorState Empty => new LdbmonitorState();
    }
}
