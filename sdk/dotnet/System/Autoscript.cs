// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    /// <summary>
    /// Configure auto script.
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
    ///     var auto2 = new Fortios.System.Autoscript("auto2", new()
    ///     {
    ///         Interval = 1,
    ///         OutputSize = 10,
    ///         Repeat = 1,
    ///         Script = @"config firewall address
    ///     edit ""111""
    ///         set color 3
    ///         set subnet 1.1.1.1 255.255.255.255
    ///     next
    /// end
    /// 
    /// ",
    ///         Start = "auto",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System AutoScript can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/autoscript:Autoscript labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/autoscript:Autoscript labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/autoscript:Autoscript")]
    public partial class Autoscript : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Repeat interval in seconds.
        /// </summary>
        [Output("interval")]
        public Output<int> Interval { get; private set; } = null!;

        /// <summary>
        /// Auto script name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of megabytes to limit script output to (10 - 1024, default = 10).
        /// </summary>
        [Output("outputSize")]
        public Output<int> OutputSize { get; private set; } = null!;

        /// <summary>
        /// Number of times to repeat this script (0 = infinite).
        /// </summary>
        [Output("repeat")]
        public Output<int> Repeat { get; private set; } = null!;

        /// <summary>
        /// List of FortiOS CLI commands to repeat.
        /// </summary>
        [Output("script")]
        public Output<string?> Script { get; private set; } = null!;

        /// <summary>
        /// Script starting mode. Valid values: `manual`, `auto`.
        /// </summary>
        [Output("start")]
        public Output<string> Start { get; private set; } = null!;

        /// <summary>
        /// Maximum running time for this script in seconds (0 = no timeout).
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Autoscript resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Autoscript(string name, AutoscriptArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/autoscript:Autoscript", name, args ?? new AutoscriptArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Autoscript(string name, Input<string> id, AutoscriptState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/autoscript:Autoscript", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Autoscript resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Autoscript Get(string name, Input<string> id, AutoscriptState? state = null, CustomResourceOptions? options = null)
        {
            return new Autoscript(name, id, state, options);
        }
    }

    public sealed class AutoscriptArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Repeat interval in seconds.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// Auto script name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of megabytes to limit script output to (10 - 1024, default = 10).
        /// </summary>
        [Input("outputSize")]
        public Input<int>? OutputSize { get; set; }

        /// <summary>
        /// Number of times to repeat this script (0 = infinite).
        /// </summary>
        [Input("repeat")]
        public Input<int>? Repeat { get; set; }

        /// <summary>
        /// List of FortiOS CLI commands to repeat.
        /// </summary>
        [Input("script")]
        public Input<string>? Script { get; set; }

        /// <summary>
        /// Script starting mode. Valid values: `manual`, `auto`.
        /// </summary>
        [Input("start")]
        public Input<string>? Start { get; set; }

        /// <summary>
        /// Maximum running time for this script in seconds (0 = no timeout).
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AutoscriptArgs()
        {
        }
        public static new AutoscriptArgs Empty => new AutoscriptArgs();
    }

    public sealed class AutoscriptState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Repeat interval in seconds.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// Auto script name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of megabytes to limit script output to (10 - 1024, default = 10).
        /// </summary>
        [Input("outputSize")]
        public Input<int>? OutputSize { get; set; }

        /// <summary>
        /// Number of times to repeat this script (0 = infinite).
        /// </summary>
        [Input("repeat")]
        public Input<int>? Repeat { get; set; }

        /// <summary>
        /// List of FortiOS CLI commands to repeat.
        /// </summary>
        [Input("script")]
        public Input<string>? Script { get; set; }

        /// <summary>
        /// Script starting mode. Valid values: `manual`, `auto`.
        /// </summary>
        [Input("start")]
        public Input<string>? Start { get; set; }

        /// <summary>
        /// Maximum running time for this script in seconds (0 = no timeout).
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AutoscriptState()
        {
        }
        public static new AutoscriptState Empty => new AutoscriptState();
    }
}
