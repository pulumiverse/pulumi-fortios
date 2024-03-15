// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    /// <summary>
    /// Internet Service definition.
    /// 
    /// ## Import
    /// 
    /// Firewall InternetServiceDefinition can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/internetservicedefinition:Internetservicedefinition labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/internetservicedefinition:Internetservicedefinition labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/internetservicedefinition:Internetservicedefinition")]
    public partial class Internetservicedefinition : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Protocol and port information in an Internet Service entry. The structure of `entry` block is documented below.
        /// </summary>
        [Output("entries")]
        public Output<ImmutableArray<Outputs.InternetservicedefinitionEntry>> Entries { get; private set; } = null!;

        /// <summary>
        /// Internet Service application list ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Internetservicedefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Internetservicedefinition(string name, InternetservicedefinitionArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/internetservicedefinition:Internetservicedefinition", name, args ?? new InternetservicedefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Internetservicedefinition(string name, Input<string> id, InternetservicedefinitionState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/internetservicedefinition:Internetservicedefinition", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Internetservicedefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Internetservicedefinition Get(string name, Input<string> id, InternetservicedefinitionState? state = null, CustomResourceOptions? options = null)
        {
            return new Internetservicedefinition(name, id, state, options);
        }
    }

    public sealed class InternetservicedefinitionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("entries")]
        private InputList<Inputs.InternetservicedefinitionEntryArgs>? _entries;

        /// <summary>
        /// Protocol and port information in an Internet Service entry. The structure of `entry` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetservicedefinitionEntryArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.InternetservicedefinitionEntryArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// Internet Service application list ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public InternetservicedefinitionArgs()
        {
        }
        public static new InternetservicedefinitionArgs Empty => new InternetservicedefinitionArgs();
    }

    public sealed class InternetservicedefinitionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("entries")]
        private InputList<Inputs.InternetservicedefinitionEntryGetArgs>? _entries;

        /// <summary>
        /// Protocol and port information in an Internet Service entry. The structure of `entry` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetservicedefinitionEntryGetArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.InternetservicedefinitionEntryGetArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// Internet Service application list ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public InternetservicedefinitionState()
        {
        }
        public static new InternetservicedefinitionState Empty => new InternetservicedefinitionState();
    }
}
