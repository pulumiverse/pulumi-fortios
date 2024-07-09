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
    /// Configure Type of Service (ToS) based priority table to set network traffic priorities.
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
    ///     var trname = new Fortios.System.Tosbasedpriority("trname", new()
    ///     {
    ///         Fosid = 1,
    ///         Priority = "low",
    ///         Tos = 11,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System TosBasedPriority can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/tosbasedpriority:Tosbasedpriority labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/tosbasedpriority:Tosbasedpriority labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/tosbasedpriority:Tosbasedpriority")]
    public partial class Tosbasedpriority : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Item ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// ToS based priority level to low, medium or high (these priorities match firewall traffic shaping priorities) (default = medium). Valid values: `low`, `medium`, `high`.
        /// </summary>
        [Output("priority")]
        public Output<string> Priority { get; private set; } = null!;

        /// <summary>
        /// Value of the ToS byte in the IP datagram header (0-15, 8: minimize delay, 4: maximize throughput, 2: maximize reliability, 1: minimize monetary cost, and 0: default service).
        /// </summary>
        [Output("tos")]
        public Output<int> Tos { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Tosbasedpriority resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Tosbasedpriority(string name, TosbasedpriorityArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/tosbasedpriority:Tosbasedpriority", name, args ?? new TosbasedpriorityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Tosbasedpriority(string name, Input<string> id, TosbasedpriorityState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/tosbasedpriority:Tosbasedpriority", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Tosbasedpriority resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Tosbasedpriority Get(string name, Input<string> id, TosbasedpriorityState? state = null, CustomResourceOptions? options = null)
        {
            return new Tosbasedpriority(name, id, state, options);
        }
    }

    public sealed class TosbasedpriorityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Item ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// ToS based priority level to low, medium or high (these priorities match firewall traffic shaping priorities) (default = medium). Valid values: `low`, `medium`, `high`.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Value of the ToS byte in the IP datagram header (0-15, 8: minimize delay, 4: maximize throughput, 2: maximize reliability, 1: minimize monetary cost, and 0: default service).
        /// </summary>
        [Input("tos")]
        public Input<int>? Tos { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TosbasedpriorityArgs()
        {
        }
        public static new TosbasedpriorityArgs Empty => new TosbasedpriorityArgs();
    }

    public sealed class TosbasedpriorityState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Item ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// ToS based priority level to low, medium or high (these priorities match firewall traffic shaping priorities) (default = medium). Valid values: `low`, `medium`, `high`.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Value of the ToS byte in the IP datagram header (0-15, 8: minimize delay, 4: maximize throughput, 2: maximize reliability, 1: minimize monetary cost, and 0: default service).
        /// </summary>
        [Input("tos")]
        public Input<int>? Tos { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TosbasedpriorityState()
        {
        }
        public static new TosbasedpriorityState Empty => new TosbasedpriorityState();
    }
}
