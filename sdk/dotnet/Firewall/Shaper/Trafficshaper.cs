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
    /// Configure shared traffic shaper.
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
    ///     var trname = new Fortios.Firewall.Shaper.Trafficshaper("trname", new()
    ///     {
    ///         BandwidthUnit = "kbps",
    ///         Diffserv = "disable",
    ///         Diffservcode = "000000",
    ///         GuaranteedBandwidth = 0,
    ///         MaximumBandwidth = 1024,
    ///         PerPolicy = "disable",
    ///         Priority = "low",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// FirewallShaper TrafficShaper can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/shaper/trafficshaper:Trafficshaper labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/shaper/trafficshaper:Trafficshaper labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/shaper/trafficshaper:Trafficshaper")]
    public partial class Trafficshaper : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        /// </summary>
        [Output("bandwidthUnit")]
        public Output<string> BandwidthUnit { get; private set; } = null!;

        /// <summary>
        /// Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("diffserv")]
        public Output<string> Diffserv { get; private set; } = null!;

        /// <summary>
        /// DiffServ setting to be applied to traffic accepted by this shaper.
        /// </summary>
        [Output("diffservcode")]
        public Output<string> Diffservcode { get; private set; } = null!;

        /// <summary>
        /// Select DSCP marking method. Valid values: `multi-stage`, `static`.
        /// </summary>
        [Output("dscpMarkingMethod")]
        public Output<string> DscpMarkingMethod { get; private set; } = null!;

        /// <summary>
        /// Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
        /// </summary>
        [Output("exceedBandwidth")]
        public Output<int> ExceedBandwidth { get; private set; } = null!;

        /// <summary>
        /// Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].
        /// </summary>
        [Output("exceedClassId")]
        public Output<int> ExceedClassId { get; private set; } = null!;

        /// <summary>
        /// DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
        /// </summary>
        [Output("exceedDscp")]
        public Output<string> ExceedDscp { get; private set; } = null!;

        /// <summary>
        /// Amount of bandwidth guaranteed for this shaper (0 - 16776000). Units depend on the bandwidth-unit setting.
        /// </summary>
        [Output("guaranteedBandwidth")]
        public Output<int> GuaranteedBandwidth { get; private set; } = null!;

        /// <summary>
        /// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
        /// </summary>
        [Output("maximumBandwidth")]
        public Output<int> MaximumBandwidth { get; private set; } = null!;

        /// <summary>
        /// DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
        /// </summary>
        [Output("maximumDscp")]
        public Output<string> MaximumDscp { get; private set; } = null!;

        /// <summary>
        /// Traffic shaper name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Per-packet size overhead used in rate computations.
        /// </summary>
        [Output("overhead")]
        public Output<int> Overhead { get; private set; } = null!;

        /// <summary>
        /// Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("perPolicy")]
        public Output<string> PerPolicy { get; private set; } = null!;

        /// <summary>
        /// Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth. Valid values: `low`, `medium`, `high`.
        /// </summary>
        [Output("priority")]
        public Output<string> Priority { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Trafficshaper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trafficshaper(string name, TrafficshaperArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/shaper/trafficshaper:Trafficshaper", name, args ?? new TrafficshaperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Trafficshaper(string name, Input<string> id, TrafficshaperState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/shaper/trafficshaper:Trafficshaper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Trafficshaper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trafficshaper Get(string name, Input<string> id, TrafficshaperState? state = null, CustomResourceOptions? options = null)
        {
            return new Trafficshaper(name, id, state, options);
        }
    }

    public sealed class TrafficshaperArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        /// </summary>
        [Input("bandwidthUnit")]
        public Input<string>? BandwidthUnit { get; set; }

        /// <summary>
        /// Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("diffserv")]
        public Input<string>? Diffserv { get; set; }

        /// <summary>
        /// DiffServ setting to be applied to traffic accepted by this shaper.
        /// </summary>
        [Input("diffservcode")]
        public Input<string>? Diffservcode { get; set; }

        /// <summary>
        /// Select DSCP marking method. Valid values: `multi-stage`, `static`.
        /// </summary>
        [Input("dscpMarkingMethod")]
        public Input<string>? DscpMarkingMethod { get; set; }

        /// <summary>
        /// Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
        /// </summary>
        [Input("exceedBandwidth")]
        public Input<int>? ExceedBandwidth { get; set; }

        /// <summary>
        /// Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].
        /// </summary>
        [Input("exceedClassId")]
        public Input<int>? ExceedClassId { get; set; }

        /// <summary>
        /// DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
        /// </summary>
        [Input("exceedDscp")]
        public Input<string>? ExceedDscp { get; set; }

        /// <summary>
        /// Amount of bandwidth guaranteed for this shaper (0 - 16776000). Units depend on the bandwidth-unit setting.
        /// </summary>
        [Input("guaranteedBandwidth")]
        public Input<int>? GuaranteedBandwidth { get; set; }

        /// <summary>
        /// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
        /// </summary>
        [Input("maximumBandwidth")]
        public Input<int>? MaximumBandwidth { get; set; }

        /// <summary>
        /// DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
        /// </summary>
        [Input("maximumDscp")]
        public Input<string>? MaximumDscp { get; set; }

        /// <summary>
        /// Traffic shaper name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Per-packet size overhead used in rate computations.
        /// </summary>
        [Input("overhead")]
        public Input<int>? Overhead { get; set; }

        /// <summary>
        /// Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("perPolicy")]
        public Input<string>? PerPolicy { get; set; }

        /// <summary>
        /// Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth. Valid values: `low`, `medium`, `high`.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TrafficshaperArgs()
        {
        }
        public static new TrafficshaperArgs Empty => new TrafficshaperArgs();
    }

    public sealed class TrafficshaperState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        /// </summary>
        [Input("bandwidthUnit")]
        public Input<string>? BandwidthUnit { get; set; }

        /// <summary>
        /// Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("diffserv")]
        public Input<string>? Diffserv { get; set; }

        /// <summary>
        /// DiffServ setting to be applied to traffic accepted by this shaper.
        /// </summary>
        [Input("diffservcode")]
        public Input<string>? Diffservcode { get; set; }

        /// <summary>
        /// Select DSCP marking method. Valid values: `multi-stage`, `static`.
        /// </summary>
        [Input("dscpMarkingMethod")]
        public Input<string>? DscpMarkingMethod { get; set; }

        /// <summary>
        /// Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
        /// </summary>
        [Input("exceedBandwidth")]
        public Input<int>? ExceedBandwidth { get; set; }

        /// <summary>
        /// Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].
        /// </summary>
        [Input("exceedClassId")]
        public Input<int>? ExceedClassId { get; set; }

        /// <summary>
        /// DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
        /// </summary>
        [Input("exceedDscp")]
        public Input<string>? ExceedDscp { get; set; }

        /// <summary>
        /// Amount of bandwidth guaranteed for this shaper (0 - 16776000). Units depend on the bandwidth-unit setting.
        /// </summary>
        [Input("guaranteedBandwidth")]
        public Input<int>? GuaranteedBandwidth { get; set; }

        /// <summary>
        /// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
        /// </summary>
        [Input("maximumBandwidth")]
        public Input<int>? MaximumBandwidth { get; set; }

        /// <summary>
        /// DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
        /// </summary>
        [Input("maximumDscp")]
        public Input<string>? MaximumDscp { get; set; }

        /// <summary>
        /// Traffic shaper name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Per-packet size overhead used in rate computations.
        /// </summary>
        [Input("overhead")]
        public Input<int>? Overhead { get; set; }

        /// <summary>
        /// Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("perPolicy")]
        public Input<string>? PerPolicy { get; set; }

        /// <summary>
        /// Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth. Valid values: `low`, `medium`, `high`.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TrafficshaperState()
        {
        }
        public static new TrafficshaperState Empty => new TrafficshaperState();
    }
}
