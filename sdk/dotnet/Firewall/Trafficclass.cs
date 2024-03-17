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
    /// Configure names for shaping classes. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// Firewall TrafficClass can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/trafficclass:Trafficclass labelname {{class_id}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/trafficclass:Trafficclass labelname {{class_id}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/trafficclass:Trafficclass")]
    public partial class Trafficclass : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Class ID to be named.
        /// </summary>
        [Output("classId")]
        public Output<int> ClassId { get; private set; } = null!;

        /// <summary>
        /// Define the name for this class-id.
        /// </summary>
        [Output("className")]
        public Output<string> ClassName { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Trafficclass resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trafficclass(string name, TrafficclassArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/trafficclass:Trafficclass", name, args ?? new TrafficclassArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Trafficclass(string name, Input<string> id, TrafficclassState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/trafficclass:Trafficclass", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Trafficclass resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trafficclass Get(string name, Input<string> id, TrafficclassState? state = null, CustomResourceOptions? options = null)
        {
            return new Trafficclass(name, id, state, options);
        }
    }

    public sealed class TrafficclassArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Class ID to be named.
        /// </summary>
        [Input("classId")]
        public Input<int>? ClassId { get; set; }

        /// <summary>
        /// Define the name for this class-id.
        /// </summary>
        [Input("className")]
        public Input<string>? ClassName { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TrafficclassArgs()
        {
        }
        public static new TrafficclassArgs Empty => new TrafficclassArgs();
    }

    public sealed class TrafficclassState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Class ID to be named.
        /// </summary>
        [Input("classId")]
        public Input<int>? ClassId { get; set; }

        /// <summary>
        /// Define the name for this class-id.
        /// </summary>
        [Input("className")]
        public Input<string>? ClassName { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TrafficclassState()
        {
        }
        public static new TrafficclassState Empty => new TrafficclassState();
    }
}
