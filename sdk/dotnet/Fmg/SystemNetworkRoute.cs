// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Fmg
{
    /// <summary>
    /// This resource supports updating system network route for FortiManager.
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
    ///     var test1 = new Fortios.Fmg.SystemNetworkRoute("test1", new()
    ///     {
    ///         Destination = "192.168.2.0 255.255.255.0",
    ///         Device = "port4",
    ///         Gateway = "192.168.2.1",
    ///         RouteId = 5,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [FortiosResourceType("fortios:fmg/systemNetworkRoute:SystemNetworkRoute")]
    public partial class SystemNetworkRoute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Destination Ip/Mask.
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// Gateway out interface.
        /// </summary>
        [Output("device")]
        public Output<string> Device { get; private set; } = null!;

        /// <summary>
        /// Gateway Ip.
        /// </summary>
        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// Route id.
        /// </summary>
        [Output("routeId")]
        public Output<int> RouteId { get; private set; } = null!;


        /// <summary>
        /// Create a SystemNetworkRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemNetworkRoute(string name, SystemNetworkRouteArgs args, CustomResourceOptions? options = null)
            : base("fortios:fmg/systemNetworkRoute:SystemNetworkRoute", name, args ?? new SystemNetworkRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemNetworkRoute(string name, Input<string> id, SystemNetworkRouteState? state = null, CustomResourceOptions? options = null)
            : base("fortios:fmg/systemNetworkRoute:SystemNetworkRoute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemNetworkRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemNetworkRoute Get(string name, Input<string> id, SystemNetworkRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemNetworkRoute(name, id, state, options);
        }
    }

    public sealed class SystemNetworkRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Destination Ip/Mask.
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        /// <summary>
        /// Gateway out interface.
        /// </summary>
        [Input("device", required: true)]
        public Input<string> Device { get; set; } = null!;

        /// <summary>
        /// Gateway Ip.
        /// </summary>
        [Input("gateway", required: true)]
        public Input<string> Gateway { get; set; } = null!;

        /// <summary>
        /// Route id.
        /// </summary>
        [Input("routeId", required: true)]
        public Input<int> RouteId { get; set; } = null!;

        public SystemNetworkRouteArgs()
        {
        }
        public static new SystemNetworkRouteArgs Empty => new SystemNetworkRouteArgs();
    }

    public sealed class SystemNetworkRouteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Destination Ip/Mask.
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// Gateway out interface.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Gateway Ip.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Route id.
        /// </summary>
        [Input("routeId")]
        public Input<int>? RouteId { get; set; }

        public SystemNetworkRouteState()
        {
        }
        public static new SystemNetworkRouteState Empty => new SystemNetworkRouteState();
    }
}
