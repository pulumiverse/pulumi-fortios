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
    /// Configure file-filter profiles. Applies to FortiOS Version `&gt;= 6.4.1`.
    /// 
    /// ## Import
    /// 
    /// FileFilter Profile can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/filefilterProfile:FilefilterProfile labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/filefilterProfile:FilefilterProfile labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/filefilterProfile:FilefilterProfile")]
    public partial class FilefilterProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable file-filter extended logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("extendedLog")]
        public Output<string> ExtendedLog { get; private set; } = null!;

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Output("featureSet")]
        public Output<string> FeatureSet { get; private set; } = null!;

        /// <summary>
        /// Enable/disable file-filter logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("log")]
        public Output<string> Log { get; private set; } = null!;

        /// <summary>
        /// Profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Replacement message group
        /// </summary>
        [Output("replacemsgGroup")]
        public Output<string> ReplacemsgGroup { get; private set; } = null!;

        /// <summary>
        /// File filter rules. The structure of `rules` block is documented below.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.FilefilterProfileRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// Enable/disable archive contents scan. (Not for CIFS) Valid values: `disable`, `enable`.
        /// </summary>
        [Output("scanArchiveContents")]
        public Output<string> ScanArchiveContents { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a FilefilterProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FilefilterProfile(string name, FilefilterProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/filefilterProfile:FilefilterProfile", name, args ?? new FilefilterProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FilefilterProfile(string name, Input<string> id, FilefilterProfileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/filefilterProfile:FilefilterProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FilefilterProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FilefilterProfile Get(string name, Input<string> id, FilefilterProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new FilefilterProfile(name, id, state, options);
        }
    }

    public sealed class FilefilterProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable file-filter extended logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Input("featureSet")]
        public Input<string>? FeatureSet { get; set; }

        /// <summary>
        /// Enable/disable file-filter logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Replacement message group
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        [Input("rules")]
        private InputList<Inputs.FilefilterProfileRuleArgs>? _rules;

        /// <summary>
        /// File filter rules. The structure of `rules` block is documented below.
        /// </summary>
        public InputList<Inputs.FilefilterProfileRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FilefilterProfileRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Enable/disable archive contents scan. (Not for CIFS) Valid values: `disable`, `enable`.
        /// </summary>
        [Input("scanArchiveContents")]
        public Input<string>? ScanArchiveContents { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FilefilterProfileArgs()
        {
        }
        public static new FilefilterProfileArgs Empty => new FilefilterProfileArgs();
    }

    public sealed class FilefilterProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable file-filter extended logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Input("featureSet")]
        public Input<string>? FeatureSet { get; set; }

        /// <summary>
        /// Enable/disable file-filter logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Replacement message group
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        [Input("rules")]
        private InputList<Inputs.FilefilterProfileRuleGetArgs>? _rules;

        /// <summary>
        /// File filter rules. The structure of `rules` block is documented below.
        /// </summary>
        public InputList<Inputs.FilefilterProfileRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FilefilterProfileRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Enable/disable archive contents scan. (Not for CIFS) Valid values: `disable`, `enable`.
        /// </summary>
        [Input("scanArchiveContents")]
        public Input<string>? ScanArchiveContents { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FilefilterProfileState()
        {
        }
        public static new FilefilterProfileState Empty => new FilefilterProfileState();
    }
}
