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
    /// Configure shaping profiles.
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
    ///     var trname = new Fortios.FirewallShapingprofile("trname", new()
    ///     {
    ///         DefaultClassId = 2,
    ///         ProfileName = "shapingprofiles1",
    ///         ShapingEntries = new[]
    ///         {
    ///             new Fortios.Inputs.FirewallShapingprofileShapingEntryArgs
    ///             {
    ///                 ClassId = 2,
    ///                 GuaranteedBandwidthPercentage = 33,
    ///                 Id = 1,
    ///                 MaximumBandwidthPercentage = 88,
    ///                 Priority = "high",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall ShapingProfile can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallShapingprofile:FirewallShapingprofile labelname {{profile_name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallShapingprofile:FirewallShapingprofile labelname {{profile_name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/firewallShapingprofile:FirewallShapingprofile")]
    public partial class FirewallShapingprofile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Default class ID to handle unclassified packets (including all local traffic).
        /// </summary>
        [Output("defaultClassId")]
        public Output<int> DefaultClassId { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Shaping profile name.
        /// </summary>
        [Output("profileName")]
        public Output<string> ProfileName { get; private set; } = null!;

        /// <summary>
        /// Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.
        /// </summary>
        [Output("shapingEntries")]
        public Output<ImmutableArray<Outputs.FirewallShapingprofileShapingEntry>> ShapingEntries { get; private set; } = null!;

        /// <summary>
        /// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallShapingprofile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallShapingprofile(string name, FirewallShapingprofileArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/firewallShapingprofile:FirewallShapingprofile", name, args ?? new FirewallShapingprofileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallShapingprofile(string name, Input<string> id, FirewallShapingprofileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallShapingprofile:FirewallShapingprofile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallShapingprofile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallShapingprofile Get(string name, Input<string> id, FirewallShapingprofileState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallShapingprofile(name, id, state, options);
        }
    }

    public sealed class FirewallShapingprofileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Default class ID to handle unclassified packets (including all local traffic).
        /// </summary>
        [Input("defaultClassId", required: true)]
        public Input<int> DefaultClassId { get; set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Shaping profile name.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        [Input("shapingEntries")]
        private InputList<Inputs.FirewallShapingprofileShapingEntryArgs>? _shapingEntries;

        /// <summary>
        /// Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallShapingprofileShapingEntryArgs> ShapingEntries
        {
            get => _shapingEntries ?? (_shapingEntries = new InputList<Inputs.FirewallShapingprofileShapingEntryArgs>());
            set => _shapingEntries = value;
        }

        /// <summary>
        /// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallShapingprofileArgs()
        {
        }
        public static new FirewallShapingprofileArgs Empty => new FirewallShapingprofileArgs();
    }

    public sealed class FirewallShapingprofileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Default class ID to handle unclassified packets (including all local traffic).
        /// </summary>
        [Input("defaultClassId")]
        public Input<int>? DefaultClassId { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Shaping profile name.
        /// </summary>
        [Input("profileName")]
        public Input<string>? ProfileName { get; set; }

        [Input("shapingEntries")]
        private InputList<Inputs.FirewallShapingprofileShapingEntryGetArgs>? _shapingEntries;

        /// <summary>
        /// Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallShapingprofileShapingEntryGetArgs> ShapingEntries
        {
            get => _shapingEntries ?? (_shapingEntries = new InputList<Inputs.FirewallShapingprofileShapingEntryGetArgs>());
            set => _shapingEntries = value;
        }

        /// <summary>
        /// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallShapingprofileState()
        {
        }
        public static new FirewallShapingprofileState Empty => new FirewallShapingprofileState();
    }
}
