// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router
{
    /// <summary>
    /// Configure multicast-flow.
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
    ///     var trname = new Fortios.Router.Multicastflow("trname", new()
    ///     {
    ///         Flows = new[]
    ///         {
    ///             new Fortios.Router.Inputs.MulticastflowFlowArgs
    ///             {
    ///                 GroupAddr = "224.252.0.0",
    ///                 SourceAddr = "224.112.0.0",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Router MulticastFlow can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/multicastflow:Multicastflow labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/multicastflow:Multicastflow labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:router/multicastflow:Multicastflow")]
    public partial class Multicastflow : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string> Comments { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Multicast-flow entries. The structure of `flows` block is documented below.
        /// </summary>
        [Output("flows")]
        public Output<ImmutableArray<Outputs.MulticastflowFlow>> Flows { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Multicastflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Multicastflow(string name, MulticastflowArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:router/multicastflow:Multicastflow", name, args ?? new MulticastflowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Multicastflow(string name, Input<string> id, MulticastflowState? state = null, CustomResourceOptions? options = null)
            : base("fortios:router/multicastflow:Multicastflow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Multicastflow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Multicastflow Get(string name, Input<string> id, MulticastflowState? state = null, CustomResourceOptions? options = null)
        {
            return new Multicastflow(name, id, state, options);
        }
    }

    public sealed class MulticastflowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("flows")]
        private InputList<Inputs.MulticastflowFlowArgs>? _flows;

        /// <summary>
        /// Multicast-flow entries. The structure of `flows` block is documented below.
        /// </summary>
        public InputList<Inputs.MulticastflowFlowArgs> Flows
        {
            get => _flows ?? (_flows = new InputList<Inputs.MulticastflowFlowArgs>());
            set => _flows = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public MulticastflowArgs()
        {
        }
        public static new MulticastflowArgs Empty => new MulticastflowArgs();
    }

    public sealed class MulticastflowState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("flows")]
        private InputList<Inputs.MulticastflowFlowGetArgs>? _flows;

        /// <summary>
        /// Multicast-flow entries. The structure of `flows` block is documented below.
        /// </summary>
        public InputList<Inputs.MulticastflowFlowGetArgs> Flows
        {
            get => _flows ?? (_flows = new InputList<Inputs.MulticastflowFlowGetArgs>());
            set => _flows = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public MulticastflowState()
        {
        }
        public static new MulticastflowState Empty => new MulticastflowState();
    }
}
