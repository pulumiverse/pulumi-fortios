// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Dlp
{
    /// <summary>
    /// Configure DLP profiles. Applies to FortiOS Version `&gt;= 7.2.0`.
    /// 
    /// ## Import
    /// 
    /// Dlp Profile can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:dlp/profile:Profile labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:dlp/profile:Profile labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:dlp/profile:Profile")]
    public partial class Profile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DLP logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dlpLog")]
        public Output<string> DlpLog { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("extendedLog")]
        public Output<string> ExtendedLog { get; private set; } = null!;

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Output("featureSet")]
        public Output<string> FeatureSet { get; private set; } = null!;

        /// <summary>
        /// Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        /// </summary>
        [Output("fullArchiveProto")]
        public Output<string> FullArchiveProto { get; private set; } = null!;

        /// <summary>
        /// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("nacQuarLog")]
        public Output<string> NacQuarLog { get; private set; } = null!;

        /// <summary>
        /// Name of the DLP profile.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Replacement message group used by this DLP profile.
        /// </summary>
        [Output("replacemsgGroup")]
        public Output<string> ReplacemsgGroup { get; private set; } = null!;

        /// <summary>
        /// Set up DLP rules for this profile. The structure of `rule` block is documented below.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.ProfileRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        /// </summary>
        [Output("summaryProto")]
        public Output<string> SummaryProto { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Profile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Profile(string name, ProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:dlp/profile:Profile", name, args ?? new ProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Profile(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:dlp/profile:Profile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Profile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Profile Get(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new Profile(name, id, state, options);
        }
    }

    public sealed class ProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Enable/disable DLP logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpLog")]
        public Input<string>? DlpLog { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Input("featureSet")]
        public Input<string>? FeatureSet { get; set; }

        /// <summary>
        /// Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        /// </summary>
        [Input("fullArchiveProto")]
        public Input<string>? FullArchiveProto { get; set; }

        /// <summary>
        /// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nacQuarLog")]
        public Input<string>? NacQuarLog { get; set; }

        /// <summary>
        /// Name of the DLP profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Replacement message group used by this DLP profile.
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        [Input("rules")]
        private InputList<Inputs.ProfileRuleArgs>? _rules;

        /// <summary>
        /// Set up DLP rules for this profile. The structure of `rule` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.ProfileRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        /// </summary>
        [Input("summaryProto")]
        public Input<string>? SummaryProto { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ProfileArgs()
        {
        }
        public static new ProfileArgs Empty => new ProfileArgs();
    }

    public sealed class ProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Enable/disable DLP logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpLog")]
        public Input<string>? DlpLog { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Input("featureSet")]
        public Input<string>? FeatureSet { get; set; }

        /// <summary>
        /// Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        /// </summary>
        [Input("fullArchiveProto")]
        public Input<string>? FullArchiveProto { get; set; }

        /// <summary>
        /// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nacQuarLog")]
        public Input<string>? NacQuarLog { get; set; }

        /// <summary>
        /// Name of the DLP profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Replacement message group used by this DLP profile.
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        [Input("rules")]
        private InputList<Inputs.ProfileRuleGetArgs>? _rules;

        /// <summary>
        /// Set up DLP rules for this profile. The structure of `rule` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.ProfileRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        /// </summary>
        [Input("summaryProto")]
        public Input<string>? SummaryProto { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ProfileState()
        {
        }
        public static new ProfileState Empty => new ProfileState();
    }
}
