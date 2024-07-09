// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Report
{
    /// <summary>
    /// Report dataset configuration. Applies to FortiOS Version `&lt;= 7.0.0`.
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
    ///     var trname = new Fortios.Report.Dataset("trname", new()
    ///     {
    ///         Policy = 0,
    ///         Query = "select * from testdb",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Report Dataset can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:report/dataset:Dataset labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:report/dataset:Dataset labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:report/dataset:Dataset")]
    public partial class Dataset : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Fields. The structure of `field` block is documented below.
        /// </summary>
        [Output("fields")]
        public Output<ImmutableArray<Outputs.DatasetField>> Fields { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Parameters. The structure of `parameters` block is documented below.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.DatasetParameter>> Parameters { get; private set; } = null!;

        /// <summary>
        /// Used by monitor policy.
        /// </summary>
        [Output("policy")]
        public Output<int> Policy { get; private set; } = null!;

        /// <summary>
        /// SQL query statement.
        /// </summary>
        [Output("query")]
        public Output<string> Query { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Dataset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dataset(string name, DatasetArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:report/dataset:Dataset", name, args ?? new DatasetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dataset(string name, Input<string> id, DatasetState? state = null, CustomResourceOptions? options = null)
            : base("fortios:report/dataset:Dataset", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Dataset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dataset Get(string name, Input<string> id, DatasetState? state = null, CustomResourceOptions? options = null)
        {
            return new Dataset(name, id, state, options);
        }
    }

    public sealed class DatasetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("fields")]
        private InputList<Inputs.DatasetFieldArgs>? _fields;

        /// <summary>
        /// Fields. The structure of `field` block is documented below.
        /// </summary>
        public InputList<Inputs.DatasetFieldArgs> Fields
        {
            get => _fields ?? (_fields = new InputList<Inputs.DatasetFieldArgs>());
            set => _fields = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputList<Inputs.DatasetParameterArgs>? _parameters;

        /// <summary>
        /// Parameters. The structure of `parameters` block is documented below.
        /// </summary>
        public InputList<Inputs.DatasetParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.DatasetParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Used by monitor policy.
        /// </summary>
        [Input("policy")]
        public Input<int>? Policy { get; set; }

        /// <summary>
        /// SQL query statement.
        /// </summary>
        [Input("query")]
        public Input<string>? Query { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DatasetArgs()
        {
        }
        public static new DatasetArgs Empty => new DatasetArgs();
    }

    public sealed class DatasetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("fields")]
        private InputList<Inputs.DatasetFieldGetArgs>? _fields;

        /// <summary>
        /// Fields. The structure of `field` block is documented below.
        /// </summary>
        public InputList<Inputs.DatasetFieldGetArgs> Fields
        {
            get => _fields ?? (_fields = new InputList<Inputs.DatasetFieldGetArgs>());
            set => _fields = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputList<Inputs.DatasetParameterGetArgs>? _parameters;

        /// <summary>
        /// Parameters. The structure of `parameters` block is documented below.
        /// </summary>
        public InputList<Inputs.DatasetParameterGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.DatasetParameterGetArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Used by monitor policy.
        /// </summary>
        [Input("policy")]
        public Input<int>? Policy { get; set; }

        /// <summary>
        /// SQL query statement.
        /// </summary>
        [Input("query")]
        public Input<string>? Query { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DatasetState()
        {
        }
        public static new DatasetState Empty => new DatasetState();
    }
}
