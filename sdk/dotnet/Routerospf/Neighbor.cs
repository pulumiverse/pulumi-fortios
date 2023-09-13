// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Routerospf
{
    /// <summary>
    /// OSPF neighbor configuration are used when OSPF runs on non-broadcast media
    /// 
    /// &gt; The provider supports the definition of Neighbor in Router Ospf `fortios.router.Ospf`, and also allows the definition of separate Neighbor resources `fortios.routerospf.Neighbor`, but do not use a `fortios.router.Ospf` with in-line Neighbor in conjunction with any `fortios.routerospf.Neighbor` resources, otherwise conflicts and overwrite will occur.
    /// 
    /// ## Import
    /// 
    /// Routerospf Neighbor can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:routerospf/neighbor:Neighbor labelname {{fosid}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:routerospf/neighbor:Neighbor labelname {{fosid}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:routerospf/neighbor:Neighbor")]
    public partial class Neighbor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
        /// </summary>
        [Output("cost")]
        public Output<int> Cost { get; private set; } = null!;

        /// <summary>
        /// Neighbor entry ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Interface IP address of the neighbor.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// Poll interval time in seconds.
        /// </summary>
        [Output("pollInterval")]
        public Output<int> PollInterval { get; private set; } = null!;

        /// <summary>
        /// Priority.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Neighbor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Neighbor(string name, NeighborArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:routerospf/neighbor:Neighbor", name, args ?? new NeighborArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Neighbor(string name, Input<string> id, NeighborState? state = null, CustomResourceOptions? options = null)
            : base("fortios:routerospf/neighbor:Neighbor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Neighbor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Neighbor Get(string name, Input<string> id, NeighborState? state = null, CustomResourceOptions? options = null)
        {
            return new Neighbor(name, id, state, options);
        }
    }

    public sealed class NeighborArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
        /// </summary>
        [Input("cost")]
        public Input<int>? Cost { get; set; }

        /// <summary>
        /// Neighbor entry ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Interface IP address of the neighbor.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Poll interval time in seconds.
        /// </summary>
        [Input("pollInterval")]
        public Input<int>? PollInterval { get; set; }

        /// <summary>
        /// Priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public NeighborArgs()
        {
        }
        public static new NeighborArgs Empty => new NeighborArgs();
    }

    public sealed class NeighborState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
        /// </summary>
        [Input("cost")]
        public Input<int>? Cost { get; set; }

        /// <summary>
        /// Neighbor entry ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Interface IP address of the neighbor.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Poll interval time in seconds.
        /// </summary>
        [Input("pollInterval")]
        public Input<int>? PollInterval { get; set; }

        /// <summary>
        /// Priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public NeighborState()
        {
        }
        public static new NeighborState Empty => new NeighborState();
    }
}
