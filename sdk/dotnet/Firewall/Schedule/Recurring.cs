// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Schedule
{
    /// <summary>
    /// Recurring schedule configuration.
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
    ///     var trname = new Fortios.Firewall.Schedule.Recurring("trname", new()
    ///     {
    ///         Color = 0,
    ///         Day = "sunday",
    ///         End = "00:00",
    ///         Start = "00:00",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// FirewallSchedule Recurring can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/schedule/recurring:Recurring labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/schedule/recurring:Recurring labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/schedule/recurring:Recurring")]
    public partial class Recurring : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
        /// </summary>
        [Output("day")]
        public Output<string> Day { get; private set; } = null!;

        /// <summary>
        /// Time of day to end the schedule, format hh:mm.
        /// </summary>
        [Output("end")]
        public Output<string> End { get; private set; } = null!;

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fabricObject")]
        public Output<string> FabricObject { get; private set; } = null!;

        /// <summary>
        /// Recurring schedule name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Time of day to start the schedule, format hh:mm.
        /// </summary>
        [Output("start")]
        public Output<string> Start { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Recurring resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Recurring(string name, RecurringArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/schedule/recurring:Recurring", name, args ?? new RecurringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Recurring(string name, Input<string> id, RecurringState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/schedule/recurring:Recurring", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Recurring resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Recurring Get(string name, Input<string> id, RecurringState? state = null, CustomResourceOptions? options = null)
        {
            return new Recurring(name, id, state, options);
        }
    }

    public sealed class RecurringArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
        /// </summary>
        [Input("day")]
        public Input<string>? Day { get; set; }

        /// <summary>
        /// Time of day to end the schedule, format hh:mm.
        /// </summary>
        [Input("end", required: true)]
        public Input<string> End { get; set; } = null!;

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Recurring schedule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Time of day to start the schedule, format hh:mm.
        /// </summary>
        [Input("start", required: true)]
        public Input<string> Start { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public RecurringArgs()
        {
        }
        public static new RecurringArgs Empty => new RecurringArgs();
    }

    public sealed class RecurringState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
        /// </summary>
        [Input("day")]
        public Input<string>? Day { get; set; }

        /// <summary>
        /// Time of day to end the schedule, format hh:mm.
        /// </summary>
        [Input("end")]
        public Input<string>? End { get; set; }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Recurring schedule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Time of day to start the schedule, format hh:mm.
        /// </summary>
        [Input("start")]
        public Input<string>? Start { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public RecurringState()
        {
        }
        public static new RecurringState Empty => new RecurringState();
    }
}
