// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Networking
{
    /// <summary>
    /// Provides a resource to configure static route of FortiOS.
    /// 
    /// !&gt; **Warning:** The resource will be deprecated and replaced by new resource `fortios.router.Static`, we recommend that you use the new resource.
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
    ///     var subnet = new Fortios.Networking.RouteStatic("subnet", new()
    ///     {
    ///         Blackhole = "disable",
    ///         Comment = "Terraform test",
    ///         Device = "port2",
    ///         Distance = "22",
    ///         Dst = "110.2.2.122/32",
    ///         Gateway = "2.2.2.2",
    ///         Priority = "3",
    ///         Status = "enable",
    ///         Weight = "3",
    ///     });
    /// 
    ///     var internetService = new Fortios.Networking.RouteStatic("internetService", new()
    ///     {
    ///         Blackhole = "disable",
    ///         Comment = "Terraform Test",
    ///         Device = "port2",
    ///         Distance = "22",
    ///         Gateway = "2.2.2.2",
    ///         InternetService = 5242881,
    ///         Priority = "3",
    ///         Status = "enable",
    ///         Weight = "3",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:networking/routeStatic:RouteStatic")]
    public partial class RouteStatic : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable black hole.
        /// </summary>
        [Output("blackhole")]
        public Output<string> Blackhole { get; private set; } = null!;

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Gateway out interface or tunnel.
        /// </summary>
        [Output("device")]
        public Output<string> Device { get; private set; } = null!;

        /// <summary>
        /// Administrative distance.
        /// </summary>
        [Output("distance")]
        public Output<string> Distance { get; private set; } = null!;

        /// <summary>
        /// Destination IP and mask for this route.
        /// </summary>
        [Output("dst")]
        public Output<string> Dst { get; private set; } = null!;

        /// <summary>
        /// Gateway IP for this route.
        /// </summary>
        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// Application ID in the Internet service database.
        /// </summary>
        [Output("internetService")]
        public Output<int> InternetService { get; private set; } = null!;

        /// <summary>
        /// Administrative priority.
        /// </summary>
        [Output("priority")]
        public Output<string> Priority { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this static route. default is "enable".
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Administrative weight.
        /// </summary>
        [Output("weight")]
        public Output<string> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a RouteStatic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteStatic(string name, RouteStaticArgs args, CustomResourceOptions? options = null)
            : base("fortios:networking/routeStatic:RouteStatic", name, args ?? new RouteStaticArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteStatic(string name, Input<string> id, RouteStaticState? state = null, CustomResourceOptions? options = null)
            : base("fortios:networking/routeStatic:RouteStatic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouteStatic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteStatic Get(string name, Input<string> id, RouteStaticState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteStatic(name, id, state, options);
        }
    }

    public sealed class RouteStaticArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable black hole.
        /// </summary>
        [Input("blackhole")]
        public Input<string>? Blackhole { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Gateway out interface or tunnel.
        /// </summary>
        [Input("device", required: true)]
        public Input<string> Device { get; set; } = null!;

        /// <summary>
        /// Administrative distance.
        /// </summary>
        [Input("distance")]
        public Input<string>? Distance { get; set; }

        /// <summary>
        /// Destination IP and mask for this route.
        /// </summary>
        [Input("dst")]
        public Input<string>? Dst { get; set; }

        /// <summary>
        /// Gateway IP for this route.
        /// </summary>
        [Input("gateway", required: true)]
        public Input<string> Gateway { get; set; } = null!;

        /// <summary>
        /// Application ID in the Internet service database.
        /// </summary>
        [Input("internetService")]
        public Input<int>? InternetService { get; set; }

        /// <summary>
        /// Administrative priority.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Enable/disable this static route. default is "enable".
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Administrative weight.
        /// </summary>
        [Input("weight")]
        public Input<string>? Weight { get; set; }

        public RouteStaticArgs()
        {
        }
        public static new RouteStaticArgs Empty => new RouteStaticArgs();
    }

    public sealed class RouteStaticState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable black hole.
        /// </summary>
        [Input("blackhole")]
        public Input<string>? Blackhole { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Gateway out interface or tunnel.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Administrative distance.
        /// </summary>
        [Input("distance")]
        public Input<string>? Distance { get; set; }

        /// <summary>
        /// Destination IP and mask for this route.
        /// </summary>
        [Input("dst")]
        public Input<string>? Dst { get; set; }

        /// <summary>
        /// Gateway IP for this route.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Application ID in the Internet service database.
        /// </summary>
        [Input("internetService")]
        public Input<int>? InternetService { get; set; }

        /// <summary>
        /// Administrative priority.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Enable/disable this static route. default is "enable".
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Administrative weight.
        /// </summary>
        [Input("weight")]
        public Input<string>? Weight { get; set; }

        public RouteStaticState()
        {
        }
        public static new RouteStaticState Empty => new RouteStaticState();
    }
}
