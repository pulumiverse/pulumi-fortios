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
    /// Configure dictionaries used by DLP blocking. Applies to FortiOS Version `&gt;= 7.2.0`.
    /// 
    /// ## Import
    /// 
    /// Dlp Dictionary can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/dlpDictionary:DlpDictionary labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/dlpDictionary:DlpDictionary labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/dlpDictionary:DlpDictionary")]
    public partial class DlpDictionary : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional comments.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// DLP dictionary entries. The structure of `entries` block is documented below.
        /// </summary>
        [Output("entries")]
        public Output<ImmutableArray<Outputs.DlpDictionaryEntry>> Entries { get; private set; } = null!;

        /// <summary>
        /// Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`.
        /// </summary>
        [Output("matchType")]
        public Output<string> MatchType { get; private set; } = null!;

        /// <summary>
        /// Name of table containing the dictionary.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a DlpDictionary resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DlpDictionary(string name, DlpDictionaryArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/dlpDictionary:DlpDictionary", name, args ?? new DlpDictionaryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DlpDictionary(string name, Input<string> id, DlpDictionaryState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/dlpDictionary:DlpDictionary", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DlpDictionary resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DlpDictionary Get(string name, Input<string> id, DlpDictionaryState? state = null, CustomResourceOptions? options = null)
        {
            return new DlpDictionary(name, id, state, options);
        }
    }

    public sealed class DlpDictionaryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("entries")]
        private InputList<Inputs.DlpDictionaryEntryArgs>? _entries;

        /// <summary>
        /// DLP dictionary entries. The structure of `entries` block is documented below.
        /// </summary>
        public InputList<Inputs.DlpDictionaryEntryArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.DlpDictionaryEntryArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`.
        /// </summary>
        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        /// <summary>
        /// Name of table containing the dictionary.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DlpDictionaryArgs()
        {
        }
        public static new DlpDictionaryArgs Empty => new DlpDictionaryArgs();
    }

    public sealed class DlpDictionaryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("entries")]
        private InputList<Inputs.DlpDictionaryEntryGetArgs>? _entries;

        /// <summary>
        /// DLP dictionary entries. The structure of `entries` block is documented below.
        /// </summary>
        public InputList<Inputs.DlpDictionaryEntryGetArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.DlpDictionaryEntryGetArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`.
        /// </summary>
        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        /// <summary>
        /// Name of table containing the dictionary.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DlpDictionaryState()
        {
        }
        public static new DlpDictionaryState Empty => new DlpDictionaryState();
    }
}
