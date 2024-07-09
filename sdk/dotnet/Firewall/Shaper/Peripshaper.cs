// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Shaper
{
    /// <summary>
    /// Configure per-IP traffic shaper.
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
    ///     var trname = new Fortios.Firewall.Shaper.Peripshaper("trname", new()
    ///     {
    ///         BandwidthUnit = "kbps",
    ///         DiffservForward = "disable",
    ///         DiffservReverse = "disable",
    ///         DiffservcodeForward = "000000",
    ///         DiffservcodeRev = "000000",
    ///         MaxBandwidth = 1024,
    ///         MaxConcurrentSession = 33,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// FirewallShaper PerIpShaper can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/shaper/peripshaper:Peripshaper labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/shaper/peripshaper:Peripshaper labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/shaper/peripshaper:Peripshaper")]
    public partial class Peripshaper : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        /// </summary>
        [Output("bandwidthUnit")]
        public Output<string> BandwidthUnit { get; private set; } = null!;

        /// <summary>
        /// Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("diffservForward")]
        public Output<string> DiffservForward { get; private set; } = null!;

        /// <summary>
        /// Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("diffservReverse")]
        public Output<string> DiffservReverse { get; private set; } = null!;

        /// <summary>
        /// Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
        /// </summary>
        [Output("diffservcodeForward")]
        public Output<string> DiffservcodeForward { get; private set; } = null!;

        /// <summary>
        /// Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
        /// </summary>
        [Output("diffservcodeRev")]
        public Output<string> DiffservcodeRev { get; private set; } = null!;

        /// <summary>
        /// Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.14, 7.0.6-7.0.13, &gt;= 7.2.1: 0 - 80000000.
        /// </summary>
        [Output("maxBandwidth")]
        public Output<int> MaxBandwidth { get; private set; } = null!;

        /// <summary>
        /// Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        /// </summary>
        [Output("maxConcurrentSession")]
        public Output<int> MaxConcurrentSession { get; private set; } = null!;

        /// <summary>
        /// Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        /// </summary>
        [Output("maxConcurrentTcpSession")]
        public Output<int> MaxConcurrentTcpSession { get; private set; } = null!;

        /// <summary>
        /// Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        /// </summary>
        [Output("maxConcurrentUdpSession")]
        public Output<int> MaxConcurrentUdpSession { get; private set; } = null!;

        /// <summary>
        /// Traffic shaper name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Peripshaper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Peripshaper(string name, PeripshaperArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/shaper/peripshaper:Peripshaper", name, args ?? new PeripshaperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Peripshaper(string name, Input<string> id, PeripshaperState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/shaper/peripshaper:Peripshaper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Peripshaper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Peripshaper Get(string name, Input<string> id, PeripshaperState? state = null, CustomResourceOptions? options = null)
        {
            return new Peripshaper(name, id, state, options);
        }
    }

    public sealed class PeripshaperArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        /// </summary>
        [Input("bandwidthUnit")]
        public Input<string>? BandwidthUnit { get; set; }

        /// <summary>
        /// Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("diffservForward")]
        public Input<string>? DiffservForward { get; set; }

        /// <summary>
        /// Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("diffservReverse")]
        public Input<string>? DiffservReverse { get; set; }

        /// <summary>
        /// Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
        /// </summary>
        [Input("diffservcodeForward")]
        public Input<string>? DiffservcodeForward { get; set; }

        /// <summary>
        /// Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
        /// </summary>
        [Input("diffservcodeRev")]
        public Input<string>? DiffservcodeRev { get; set; }

        /// <summary>
        /// Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.14, 7.0.6-7.0.13, &gt;= 7.2.1: 0 - 80000000.
        /// </summary>
        [Input("maxBandwidth")]
        public Input<int>? MaxBandwidth { get; set; }

        /// <summary>
        /// Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        /// </summary>
        [Input("maxConcurrentSession")]
        public Input<int>? MaxConcurrentSession { get; set; }

        /// <summary>
        /// Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        /// </summary>
        [Input("maxConcurrentTcpSession")]
        public Input<int>? MaxConcurrentTcpSession { get; set; }

        /// <summary>
        /// Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        /// </summary>
        [Input("maxConcurrentUdpSession")]
        public Input<int>? MaxConcurrentUdpSession { get; set; }

        /// <summary>
        /// Traffic shaper name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PeripshaperArgs()
        {
        }
        public static new PeripshaperArgs Empty => new PeripshaperArgs();
    }

    public sealed class PeripshaperState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        /// </summary>
        [Input("bandwidthUnit")]
        public Input<string>? BandwidthUnit { get; set; }

        /// <summary>
        /// Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("diffservForward")]
        public Input<string>? DiffservForward { get; set; }

        /// <summary>
        /// Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("diffservReverse")]
        public Input<string>? DiffservReverse { get; set; }

        /// <summary>
        /// Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
        /// </summary>
        [Input("diffservcodeForward")]
        public Input<string>? DiffservcodeForward { get; set; }

        /// <summary>
        /// Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
        /// </summary>
        [Input("diffservcodeRev")]
        public Input<string>? DiffservcodeRev { get; set; }

        /// <summary>
        /// Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.14, 7.0.6-7.0.13, &gt;= 7.2.1: 0 - 80000000.
        /// </summary>
        [Input("maxBandwidth")]
        public Input<int>? MaxBandwidth { get; set; }

        /// <summary>
        /// Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        /// </summary>
        [Input("maxConcurrentSession")]
        public Input<int>? MaxConcurrentSession { get; set; }

        /// <summary>
        /// Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        /// </summary>
        [Input("maxConcurrentTcpSession")]
        public Input<int>? MaxConcurrentTcpSession { get; set; }

        /// <summary>
        /// Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        /// </summary>
        [Input("maxConcurrentUdpSession")]
        public Input<int>? MaxConcurrentUdpSession { get; set; }

        /// <summary>
        /// Traffic shaper name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PeripshaperState()
        {
        }
        public static new PeripshaperState Empty => new PeripshaperState();
    }
}
