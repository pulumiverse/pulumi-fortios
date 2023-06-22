// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Configure DSCP based priority table.
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
    ///     var trname = new Fortios.SystemDscpbasedpriority("trname", new()
    ///     {
    ///         Ds = 1,
    ///         Fosid = 1,
    ///         Priority = "low",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System DscpBasedPriority can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemDscpbasedpriority:SystemDscpbasedpriority labelname {{fosid}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemDscpbasedpriority:SystemDscpbasedpriority labelname {{fosid}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemDscpbasedpriority:SystemDscpbasedpriority")]
    public partial class SystemDscpbasedpriority : global::Pulumi.CustomResource
    {
        /// <summary>
        /// DSCP(DiffServ) DS value (0 - 63).
        /// </summary>
        [Output("ds")]
        public Output<int> Ds { get; private set; } = null!;

        /// <summary>
        /// Item ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// DSCP based priority level. Valid values: `low`, `medium`, `high`.
        /// </summary>
        [Output("priority")]
        public Output<string> Priority { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemDscpbasedpriority resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemDscpbasedpriority(string name, SystemDscpbasedpriorityArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemDscpbasedpriority:SystemDscpbasedpriority", name, args ?? new SystemDscpbasedpriorityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemDscpbasedpriority(string name, Input<string> id, SystemDscpbasedpriorityState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemDscpbasedpriority:SystemDscpbasedpriority", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemDscpbasedpriority resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemDscpbasedpriority Get(string name, Input<string> id, SystemDscpbasedpriorityState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemDscpbasedpriority(name, id, state, options);
        }
    }

    public sealed class SystemDscpbasedpriorityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DSCP(DiffServ) DS value (0 - 63).
        /// </summary>
        [Input("ds")]
        public Input<int>? Ds { get; set; }

        /// <summary>
        /// Item ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// DSCP based priority level. Valid values: `low`, `medium`, `high`.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemDscpbasedpriorityArgs()
        {
        }
        public static new SystemDscpbasedpriorityArgs Empty => new SystemDscpbasedpriorityArgs();
    }

    public sealed class SystemDscpbasedpriorityState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DSCP(DiffServ) DS value (0 - 63).
        /// </summary>
        [Input("ds")]
        public Input<int>? Ds { get; set; }

        /// <summary>
        /// Item ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// DSCP based priority level. Valid values: `low`, `medium`, `high`.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemDscpbasedpriorityState()
        {
        }
        public static new SystemDscpbasedpriorityState Empty => new SystemDscpbasedpriorityState();
    }
}
