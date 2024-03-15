// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Bgp
{
    /// <summary>
    /// BGP network table.
    /// 
    /// &gt; The provider supports the definition of Network in Router Bgp `fortios.router.Bgp`, and also allows the definition of separate Network resources `fortios.router/bgp.Network`, but do not use a `fortios.router.Bgp` with in-line Network in conjunction with any `fortios.router/bgp.Network` resources, otherwise conflicts and overwrite will occur.
    /// 
    /// ## Import
    /// 
    /// Routerbgp Network can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/bgp/network:Network labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/bgp/network:Network labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:router/bgp/network:Network")]
    public partial class Network : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("backdoor")]
        public Output<string> Backdoor { get; private set; } = null!;

        /// <summary>
        /// ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        /// </summary>
        [Output("networkImportCheck")]
        public Output<string> NetworkImportCheck { get; private set; } = null!;

        /// <summary>
        /// Network prefix.
        /// </summary>
        [Output("prefix")]
        public Output<string> Prefix { get; private set; } = null!;

        /// <summary>
        /// Route map to modify generated route.
        /// </summary>
        [Output("routeMap")]
        public Output<string> RouteMap { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Network resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Network(string name, NetworkArgs args, CustomResourceOptions? options = null)
            : base("fortios:router/bgp/network:Network", name, args ?? new NetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Network(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
            : base("fortios:router/bgp/network:Network", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Network resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Network Get(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new Network(name, id, state, options);
        }
    }

    public sealed class NetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("backdoor")]
        public Input<string>? Backdoor { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        /// </summary>
        [Input("networkImportCheck")]
        public Input<string>? NetworkImportCheck { get; set; }

        /// <summary>
        /// Network prefix.
        /// </summary>
        [Input("prefix", required: true)]
        public Input<string> Prefix { get; set; } = null!;

        /// <summary>
        /// Route map to modify generated route.
        /// </summary>
        [Input("routeMap")]
        public Input<string>? RouteMap { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public NetworkArgs()
        {
        }
        public static new NetworkArgs Empty => new NetworkArgs();
    }

    public sealed class NetworkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("backdoor")]
        public Input<string>? Backdoor { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        /// </summary>
        [Input("networkImportCheck")]
        public Input<string>? NetworkImportCheck { get; set; }

        /// <summary>
        /// Network prefix.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// Route map to modify generated route.
        /// </summary>
        [Input("routeMap")]
        public Input<string>? RouteMap { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public NetworkState()
        {
        }
        public static new NetworkState Empty => new NetworkState();
    }
}