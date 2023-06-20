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
    /// Configure Spanning Tree Protocol (STP). Applies to FortiOS Version `7.0.4`.
    /// 
    /// ## Import
    /// 
    /// System Stp can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemStp:SystemStp labelname SystemStp
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemStp:SystemStp labelname SystemStp
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemStp:SystemStp")]
    public partial class SystemStp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Forward delay (4 - 30 sec, default = 15).
        /// </summary>
        [Output("forwardDelay")]
        public Output<int> ForwardDelay { get; private set; } = null!;

        /// <summary>
        /// Hello time (1 - 10 sec, default = 2).
        /// </summary>
        [Output("helloTime")]
        public Output<int> HelloTime { get; private set; } = null!;

        /// <summary>
        /// Maximum packet age (6 - 40 sec, default = 20).
        /// </summary>
        [Output("maxAge")]
        public Output<int> MaxAge { get; private set; } = null!;

        /// <summary>
        /// Maximum number of hops (1 - 40, default = 20).
        /// </summary>
        [Output("maxHops")]
        public Output<int> MaxHops { get; private set; } = null!;

        /// <summary>
        /// STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
        /// </summary>
        [Output("switchPriority")]
        public Output<string> SwitchPriority { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemStp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemStp(string name, SystemStpArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemStp:SystemStp", name, args ?? new SystemStpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemStp(string name, Input<string> id, SystemStpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemStp:SystemStp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemStp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemStp Get(string name, Input<string> id, SystemStpState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemStp(name, id, state, options);
        }
    }

    public sealed class SystemStpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Forward delay (4 - 30 sec, default = 15).
        /// </summary>
        [Input("forwardDelay")]
        public Input<int>? ForwardDelay { get; set; }

        /// <summary>
        /// Hello time (1 - 10 sec, default = 2).
        /// </summary>
        [Input("helloTime")]
        public Input<int>? HelloTime { get; set; }

        /// <summary>
        /// Maximum packet age (6 - 40 sec, default = 20).
        /// </summary>
        [Input("maxAge")]
        public Input<int>? MaxAge { get; set; }

        /// <summary>
        /// Maximum number of hops (1 - 40, default = 20).
        /// </summary>
        [Input("maxHops")]
        public Input<int>? MaxHops { get; set; }

        /// <summary>
        /// STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
        /// </summary>
        [Input("switchPriority")]
        public Input<string>? SwitchPriority { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemStpArgs()
        {
        }
        public static new SystemStpArgs Empty => new SystemStpArgs();
    }

    public sealed class SystemStpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Forward delay (4 - 30 sec, default = 15).
        /// </summary>
        [Input("forwardDelay")]
        public Input<int>? ForwardDelay { get; set; }

        /// <summary>
        /// Hello time (1 - 10 sec, default = 2).
        /// </summary>
        [Input("helloTime")]
        public Input<int>? HelloTime { get; set; }

        /// <summary>
        /// Maximum packet age (6 - 40 sec, default = 20).
        /// </summary>
        [Input("maxAge")]
        public Input<int>? MaxAge { get; set; }

        /// <summary>
        /// Maximum number of hops (1 - 40, default = 20).
        /// </summary>
        [Input("maxHops")]
        public Input<int>? MaxHops { get; set; }

        /// <summary>
        /// STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
        /// </summary>
        [Input("switchPriority")]
        public Input<string>? SwitchPriority { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemStpState()
        {
        }
        public static new SystemStpState Empty => new SystemStpState();
    }
}