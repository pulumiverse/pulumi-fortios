// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Hotspot20
{
    /// <summary>
    /// Configure WAN metrics.
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
    ///     var trname = new Fortios.Wirelesscontroller.Hotspot20.H2qpwanmetric("trname", new()
    ///     {
    ///         DownlinkLoad = 0,
    ///         DownlinkSpeed = 2400,
    ///         LinkAtCapacity = "disable",
    ///         LinkStatus = "up",
    ///         LoadMeasurementDuration = 0,
    ///         SymmetricWanLink = "symmetric",
    ///         UplinkLoad = 0,
    ///         UplinkSpeed = 2400,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// WirelessControllerHotspot20 H2QpWanMetric can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/hotspot20/h2qpwanmetric:H2qpwanmetric labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/hotspot20/h2qpwanmetric:H2qpwanmetric labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontroller/hotspot20/h2qpwanmetric:H2qpwanmetric")]
    public partial class H2qpwanmetric : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Downlink load.
        /// </summary>
        [Output("downlinkLoad")]
        public Output<int> DownlinkLoad { get; private set; } = null!;

        /// <summary>
        /// Downlink speed (in kilobits/s).
        /// </summary>
        [Output("downlinkSpeed")]
        public Output<int> DownlinkSpeed { get; private set; } = null!;

        /// <summary>
        /// Link at capacity. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("linkAtCapacity")]
        public Output<string> LinkAtCapacity { get; private set; } = null!;

        /// <summary>
        /// Link status. Valid values: `up`, `down`, `in-test`.
        /// </summary>
        [Output("linkStatus")]
        public Output<string> LinkStatus { get; private set; } = null!;

        /// <summary>
        /// Load measurement duration (in tenths of a second).
        /// </summary>
        [Output("loadMeasurementDuration")]
        public Output<int> LoadMeasurementDuration { get; private set; } = null!;

        /// <summary>
        /// WAN metric name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
        /// </summary>
        [Output("symmetricWanLink")]
        public Output<string> SymmetricWanLink { get; private set; } = null!;

        /// <summary>
        /// Uplink load.
        /// </summary>
        [Output("uplinkLoad")]
        public Output<int> UplinkLoad { get; private set; } = null!;

        /// <summary>
        /// Uplink speed (in kilobits/s).
        /// </summary>
        [Output("uplinkSpeed")]
        public Output<int> UplinkSpeed { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a H2qpwanmetric resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public H2qpwanmetric(string name, H2qpwanmetricArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/hotspot20/h2qpwanmetric:H2qpwanmetric", name, args ?? new H2qpwanmetricArgs(), MakeResourceOptions(options, ""))
        {
        }

        private H2qpwanmetric(string name, Input<string> id, H2qpwanmetricState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/hotspot20/h2qpwanmetric:H2qpwanmetric", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing H2qpwanmetric resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static H2qpwanmetric Get(string name, Input<string> id, H2qpwanmetricState? state = null, CustomResourceOptions? options = null)
        {
            return new H2qpwanmetric(name, id, state, options);
        }
    }

    public sealed class H2qpwanmetricArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Downlink load.
        /// </summary>
        [Input("downlinkLoad")]
        public Input<int>? DownlinkLoad { get; set; }

        /// <summary>
        /// Downlink speed (in kilobits/s).
        /// </summary>
        [Input("downlinkSpeed")]
        public Input<int>? DownlinkSpeed { get; set; }

        /// <summary>
        /// Link at capacity. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("linkAtCapacity")]
        public Input<string>? LinkAtCapacity { get; set; }

        /// <summary>
        /// Link status. Valid values: `up`, `down`, `in-test`.
        /// </summary>
        [Input("linkStatus")]
        public Input<string>? LinkStatus { get; set; }

        /// <summary>
        /// Load measurement duration (in tenths of a second).
        /// </summary>
        [Input("loadMeasurementDuration")]
        public Input<int>? LoadMeasurementDuration { get; set; }

        /// <summary>
        /// WAN metric name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
        /// </summary>
        [Input("symmetricWanLink")]
        public Input<string>? SymmetricWanLink { get; set; }

        /// <summary>
        /// Uplink load.
        /// </summary>
        [Input("uplinkLoad")]
        public Input<int>? UplinkLoad { get; set; }

        /// <summary>
        /// Uplink speed (in kilobits/s).
        /// </summary>
        [Input("uplinkSpeed")]
        public Input<int>? UplinkSpeed { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public H2qpwanmetricArgs()
        {
        }
        public static new H2qpwanmetricArgs Empty => new H2qpwanmetricArgs();
    }

    public sealed class H2qpwanmetricState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Downlink load.
        /// </summary>
        [Input("downlinkLoad")]
        public Input<int>? DownlinkLoad { get; set; }

        /// <summary>
        /// Downlink speed (in kilobits/s).
        /// </summary>
        [Input("downlinkSpeed")]
        public Input<int>? DownlinkSpeed { get; set; }

        /// <summary>
        /// Link at capacity. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("linkAtCapacity")]
        public Input<string>? LinkAtCapacity { get; set; }

        /// <summary>
        /// Link status. Valid values: `up`, `down`, `in-test`.
        /// </summary>
        [Input("linkStatus")]
        public Input<string>? LinkStatus { get; set; }

        /// <summary>
        /// Load measurement duration (in tenths of a second).
        /// </summary>
        [Input("loadMeasurementDuration")]
        public Input<int>? LoadMeasurementDuration { get; set; }

        /// <summary>
        /// WAN metric name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
        /// </summary>
        [Input("symmetricWanLink")]
        public Input<string>? SymmetricWanLink { get; set; }

        /// <summary>
        /// Uplink load.
        /// </summary>
        [Input("uplinkLoad")]
        public Input<int>? UplinkLoad { get; set; }

        /// <summary>
        /// Uplink speed (in kilobits/s).
        /// </summary>
        [Input("uplinkSpeed")]
        public Input<int>? UplinkSpeed { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public H2qpwanmetricState()
        {
        }
        public static new H2qpwanmetricState Empty => new H2qpwanmetricState();
    }
}
