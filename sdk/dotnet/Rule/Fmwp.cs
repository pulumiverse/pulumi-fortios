// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Rule
{
    /// <summary>
    /// Show FMWP signatures. Applies to FortiOS Version `&gt;= 7.4.2`.
    /// 
    /// ## Import
    /// 
    /// Rule Fmwp can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:rule/fmwp:Fmwp labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:rule/fmwp:Fmwp labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:rule/fmwp:Fmwp")]
    public partial class Fmwp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action. Valid values: `pass`, `block`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// Vulnerable applications.
        /// </summary>
        [Output("application")]
        public Output<string> Application { get; private set; } = null!;

        /// <summary>
        /// Date.
        /// </summary>
        [Output("date")]
        public Output<int> Date { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Group.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// Vulnerable location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("log")]
        public Output<string> Log { get; private set; } = null!;

        /// <summary>
        /// Enable/disable packet logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("logPacket")]
        public Output<string> LogPacket { get; private set; } = null!;

        /// <summary>
        /// Meta data. The structure of `metadata` block is documented below.
        /// </summary>
        [Output("metadatas")]
        public Output<ImmutableArray<Outputs.FmwpMetadata>> Metadatas { get; private set; } = null!;

        /// <summary>
        /// Rule name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Vulnerable operation systems.
        /// </summary>
        [Output("os")]
        public Output<string> Os { get; private set; } = null!;

        /// <summary>
        /// Revision.
        /// </summary>
        [Output("rev")]
        public Output<int> Rev { get; private set; } = null!;

        /// <summary>
        /// Rule ID.
        /// </summary>
        [Output("ruleId")]
        public Output<int> RuleId { get; private set; } = null!;

        /// <summary>
        /// Vulnerable service.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;

        /// <summary>
        /// Severity.
        /// </summary>
        [Output("severity")]
        public Output<string> Severity { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Fmwp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fmwp(string name, FmwpArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:rule/fmwp:Fmwp", name, args ?? new FmwpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fmwp(string name, Input<string> id, FmwpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:rule/fmwp:Fmwp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Fmwp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fmwp Get(string name, Input<string> id, FmwpState? state = null, CustomResourceOptions? options = null)
        {
            return new Fmwp(name, id, state, options);
        }
    }

    public sealed class FmwpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action. Valid values: `pass`, `block`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Vulnerable applications.
        /// </summary>
        [Input("application")]
        public Input<string>? Application { get; set; }

        /// <summary>
        /// Date.
        /// </summary>
        [Input("date")]
        public Input<int>? Date { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Group.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Vulnerable location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Enable/disable logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Enable/disable packet logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logPacket")]
        public Input<string>? LogPacket { get; set; }

        [Input("metadatas")]
        private InputList<Inputs.FmwpMetadataArgs>? _metadatas;

        /// <summary>
        /// Meta data. The structure of `metadata` block is documented below.
        /// </summary>
        public InputList<Inputs.FmwpMetadataArgs> Metadatas
        {
            get => _metadatas ?? (_metadatas = new InputList<Inputs.FmwpMetadataArgs>());
            set => _metadatas = value;
        }

        /// <summary>
        /// Rule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Vulnerable operation systems.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// Revision.
        /// </summary>
        [Input("rev")]
        public Input<int>? Rev { get; set; }

        /// <summary>
        /// Rule ID.
        /// </summary>
        [Input("ruleId")]
        public Input<int>? RuleId { get; set; }

        /// <summary>
        /// Vulnerable service.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// Severity.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FmwpArgs()
        {
        }
        public static new FmwpArgs Empty => new FmwpArgs();
    }

    public sealed class FmwpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action. Valid values: `pass`, `block`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Vulnerable applications.
        /// </summary>
        [Input("application")]
        public Input<string>? Application { get; set; }

        /// <summary>
        /// Date.
        /// </summary>
        [Input("date")]
        public Input<int>? Date { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Group.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Vulnerable location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Enable/disable logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Enable/disable packet logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logPacket")]
        public Input<string>? LogPacket { get; set; }

        [Input("metadatas")]
        private InputList<Inputs.FmwpMetadataGetArgs>? _metadatas;

        /// <summary>
        /// Meta data. The structure of `metadata` block is documented below.
        /// </summary>
        public InputList<Inputs.FmwpMetadataGetArgs> Metadatas
        {
            get => _metadatas ?? (_metadatas = new InputList<Inputs.FmwpMetadataGetArgs>());
            set => _metadatas = value;
        }

        /// <summary>
        /// Rule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Vulnerable operation systems.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// Revision.
        /// </summary>
        [Input("rev")]
        public Input<int>? Rev { get; set; }

        /// <summary>
        /// Rule ID.
        /// </summary>
        [Input("ruleId")]
        public Input<int>? RuleId { get; set; }

        /// <summary>
        /// Vulnerable service.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// Severity.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FmwpState()
        {
        }
        public static new FmwpState Empty => new FmwpState();
    }
}
