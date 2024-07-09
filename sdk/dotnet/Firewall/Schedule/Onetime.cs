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
    /// Onetime schedule configuration.
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
    ///     var trname = new Fortios.Firewall.Schedule.Onetime("trname", new()
    ///     {
    ///         Color = 0,
    ///         End = "00:00 2020/12/12",
    ///         ExpirationDays = 2,
    ///         Start = "00:00 2010/12/12",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// FirewallSchedule Onetime can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/schedule/onetime:Onetime labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/schedule/onetime:Onetime labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/schedule/onetime:Onetime")]
    public partial class Onetime : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// Schedule end date and time, format hh:mm yyyy/mm/dd.
        /// </summary>
        [Output("end")]
        public Output<string> End { get; private set; } = null!;

        /// <summary>
        /// Schedule end date and time, in epoch format.
        /// </summary>
        [Output("endUtc")]
        public Output<string> EndUtc { get; private set; } = null!;

        /// <summary>
        /// Write an event log message this many days before the schedule expires.
        /// </summary>
        [Output("expirationDays")]
        public Output<int> ExpirationDays { get; private set; } = null!;

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fabricObject")]
        public Output<string> FabricObject { get; private set; } = null!;

        /// <summary>
        /// Onetime schedule name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Schedule start date and time, format hh:mm yyyy/mm/dd.
        /// </summary>
        [Output("start")]
        public Output<string> Start { get; private set; } = null!;

        /// <summary>
        /// Schedule start date and time, in epoch format.
        /// </summary>
        [Output("startUtc")]
        public Output<string> StartUtc { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Onetime resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Onetime(string name, OnetimeArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/schedule/onetime:Onetime", name, args ?? new OnetimeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Onetime(string name, Input<string> id, OnetimeState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/schedule/onetime:Onetime", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Onetime resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Onetime Get(string name, Input<string> id, OnetimeState? state = null, CustomResourceOptions? options = null)
        {
            return new Onetime(name, id, state, options);
        }
    }

    public sealed class OnetimeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Schedule end date and time, format hh:mm yyyy/mm/dd.
        /// </summary>
        [Input("end", required: true)]
        public Input<string> End { get; set; } = null!;

        /// <summary>
        /// Schedule end date and time, in epoch format.
        /// </summary>
        [Input("endUtc")]
        public Input<string>? EndUtc { get; set; }

        /// <summary>
        /// Write an event log message this many days before the schedule expires.
        /// </summary>
        [Input("expirationDays")]
        public Input<int>? ExpirationDays { get; set; }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Onetime schedule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Schedule start date and time, format hh:mm yyyy/mm/dd.
        /// </summary>
        [Input("start", required: true)]
        public Input<string> Start { get; set; } = null!;

        /// <summary>
        /// Schedule start date and time, in epoch format.
        /// </summary>
        [Input("startUtc")]
        public Input<string>? StartUtc { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public OnetimeArgs()
        {
        }
        public static new OnetimeArgs Empty => new OnetimeArgs();
    }

    public sealed class OnetimeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Schedule end date and time, format hh:mm yyyy/mm/dd.
        /// </summary>
        [Input("end")]
        public Input<string>? End { get; set; }

        /// <summary>
        /// Schedule end date and time, in epoch format.
        /// </summary>
        [Input("endUtc")]
        public Input<string>? EndUtc { get; set; }

        /// <summary>
        /// Write an event log message this many days before the schedule expires.
        /// </summary>
        [Input("expirationDays")]
        public Input<int>? ExpirationDays { get; set; }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Onetime schedule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Schedule start date and time, format hh:mm yyyy/mm/dd.
        /// </summary>
        [Input("start")]
        public Input<string>? Start { get; set; }

        /// <summary>
        /// Schedule start date and time, in epoch format.
        /// </summary>
        [Input("startUtc")]
        public Input<string>? StartUtc { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public OnetimeState()
        {
        }
        public static new OnetimeState Empty => new OnetimeState();
    }
}
