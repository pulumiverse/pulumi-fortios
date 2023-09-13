// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller
{
    /// <summary>
    /// Configure FortiSwitch traffic policy.
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
    ///     var trname = new Fortios.Switchcontroller.Trafficpolicy("trname", new()
    ///     {
    ///         GuaranteedBandwidth = 0,
    ///         GuaranteedBurst = 0,
    ///         MaximumBurst = 0,
    ///         PolicerStatus = "enable",
    ///         Type = "ingress",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SwitchController TrafficPolicy can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:switchcontroller/trafficpolicy:Trafficpolicy labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:switchcontroller/trafficpolicy:Trafficpolicy labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/trafficpolicy:Trafficpolicy")]
    public partial class Trafficpolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// COS queue(0 - 7), or unset to disable.
        /// </summary>
        [Output("cos")]
        public Output<int> Cos { get; private set; } = null!;

        /// <summary>
        /// COS queue(0 - 7), or unset to disable.
        /// </summary>
        [Output("cosQueue")]
        public Output<int> CosQueue { get; private set; } = null!;

        /// <summary>
        /// Description of the traffic policy.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// FSW Policer id
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Guaranteed bandwidth in kbps (max value = 524287000).
        /// </summary>
        [Output("guaranteedBandwidth")]
        public Output<int> GuaranteedBandwidth { get; private set; } = null!;

        /// <summary>
        /// Guaranteed burst size in bytes (max value = 4294967295).
        /// </summary>
        [Output("guaranteedBurst")]
        public Output<int> GuaranteedBurst { get; private set; } = null!;

        /// <summary>
        /// Maximum burst size in bytes (max value = 4294967295).
        /// </summary>
        [Output("maximumBurst")]
        public Output<int> MaximumBurst { get; private set; } = null!;

        /// <summary>
        /// Traffic policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable policer config on the traffic policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("policerStatus")]
        public Output<string> PolicerStatus { get; private set; } = null!;

        /// <summary>
        /// Configure type of policy(ingress/egress). Valid values: `ingress`, `egress`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Trafficpolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trafficpolicy(string name, TrafficpolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/trafficpolicy:Trafficpolicy", name, args ?? new TrafficpolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Trafficpolicy(string name, Input<string> id, TrafficpolicyState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/trafficpolicy:Trafficpolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Trafficpolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trafficpolicy Get(string name, Input<string> id, TrafficpolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Trafficpolicy(name, id, state, options);
        }
    }

    public sealed class TrafficpolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// COS queue(0 - 7), or unset to disable.
        /// </summary>
        [Input("cos")]
        public Input<int>? Cos { get; set; }

        /// <summary>
        /// COS queue(0 - 7), or unset to disable.
        /// </summary>
        [Input("cosQueue")]
        public Input<int>? CosQueue { get; set; }

        /// <summary>
        /// Description of the traffic policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// FSW Policer id
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Guaranteed bandwidth in kbps (max value = 524287000).
        /// </summary>
        [Input("guaranteedBandwidth")]
        public Input<int>? GuaranteedBandwidth { get; set; }

        /// <summary>
        /// Guaranteed burst size in bytes (max value = 4294967295).
        /// </summary>
        [Input("guaranteedBurst")]
        public Input<int>? GuaranteedBurst { get; set; }

        /// <summary>
        /// Maximum burst size in bytes (max value = 4294967295).
        /// </summary>
        [Input("maximumBurst")]
        public Input<int>? MaximumBurst { get; set; }

        /// <summary>
        /// Traffic policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable policer config on the traffic policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("policerStatus")]
        public Input<string>? PolicerStatus { get; set; }

        /// <summary>
        /// Configure type of policy(ingress/egress). Valid values: `ingress`, `egress`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TrafficpolicyArgs()
        {
        }
        public static new TrafficpolicyArgs Empty => new TrafficpolicyArgs();
    }

    public sealed class TrafficpolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// COS queue(0 - 7), or unset to disable.
        /// </summary>
        [Input("cos")]
        public Input<int>? Cos { get; set; }

        /// <summary>
        /// COS queue(0 - 7), or unset to disable.
        /// </summary>
        [Input("cosQueue")]
        public Input<int>? CosQueue { get; set; }

        /// <summary>
        /// Description of the traffic policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// FSW Policer id
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Guaranteed bandwidth in kbps (max value = 524287000).
        /// </summary>
        [Input("guaranteedBandwidth")]
        public Input<int>? GuaranteedBandwidth { get; set; }

        /// <summary>
        /// Guaranteed burst size in bytes (max value = 4294967295).
        /// </summary>
        [Input("guaranteedBurst")]
        public Input<int>? GuaranteedBurst { get; set; }

        /// <summary>
        /// Maximum burst size in bytes (max value = 4294967295).
        /// </summary>
        [Input("maximumBurst")]
        public Input<int>? MaximumBurst { get; set; }

        /// <summary>
        /// Traffic policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable policer config on the traffic policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("policerStatus")]
        public Input<string>? PolicerStatus { get; set; }

        /// <summary>
        /// Configure type of policy(ingress/egress). Valid values: `ingress`, `egress`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TrafficpolicyState()
        {
        }
        public static new TrafficpolicyState Empty => new TrafficpolicyState();
    }
}
