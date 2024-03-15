// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Application
{
    /// <summary>
    /// Configure application signatures.
    /// 
    /// ## Import
    /// 
    /// Application Name can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:application/name:Name labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:application/name:Name labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:application/name:Name")]
    public partial class ApplicationName : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Application behavior.
        /// </summary>
        [Output("behavior")]
        public Output<string> Behavior { get; private set; } = null!;

        /// <summary>
        /// Application category ID.
        /// </summary>
        [Output("category")]
        public Output<int> Category { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Application ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Meta data. The structure of `metadata` block is documented below.
        /// </summary>
        [Output("metadatas")]
        public Output<ImmutableArray<Outputs.NameMetadata>> Metadatas { get; private set; } = null!;

        /// <summary>
        /// Application name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Application parameter name.
        /// </summary>
        [Output("parameter")]
        public Output<string> Parameter { get; private set; } = null!;

        /// <summary>
        /// Application parameters. The structure of `parameters` block is documented below.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.NameParameter>> Parameters { get; private set; } = null!;

        /// <summary>
        /// Application popularity.
        /// </summary>
        [Output("popularity")]
        public Output<int> Popularity { get; private set; } = null!;

        /// <summary>
        /// Application protocol.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Application risk.
        /// </summary>
        [Output("risk")]
        public Output<int> Risk { get; private set; } = null!;

        /// <summary>
        /// Application sub-category ID.
        /// </summary>
        [Output("subCategory")]
        public Output<int> SubCategory { get; private set; } = null!;

        /// <summary>
        /// Application technology.
        /// </summary>
        [Output("technology")]
        public Output<string> Technology { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Application vendor.
        /// </summary>
        [Output("vendor")]
        public Output<string> Vendor { get; private set; } = null!;

        /// <summary>
        /// Application weight.
        /// </summary>
        [Output("weight")]
        public Output<int> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationName resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationName(string name, ApplicationNameArgs args, CustomResourceOptions? options = null)
            : base("fortios:application/name:Name", name, args ?? new ApplicationNameArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationName(string name, Input<string> id, ApplicationNameState? state = null, CustomResourceOptions? options = null)
            : base("fortios:application/name:Name", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationName resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationName Get(string name, Input<string> id, ApplicationNameState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationName(name, id, state, options);
        }
    }

    public sealed class ApplicationNameArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application behavior.
        /// </summary>
        [Input("behavior")]
        public Input<string>? Behavior { get; set; }

        /// <summary>
        /// Application category ID.
        /// </summary>
        [Input("category", required: true)]
        public Input<int> Category { get; set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Application ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        [Input("metadatas")]
        private InputList<Inputs.NameMetadataArgs>? _metadatas;

        /// <summary>
        /// Meta data. The structure of `metadata` block is documented below.
        /// </summary>
        public InputList<Inputs.NameMetadataArgs> Metadatas
        {
            get => _metadatas ?? (_metadatas = new InputList<Inputs.NameMetadataArgs>());
            set => _metadatas = value;
        }

        /// <summary>
        /// Application name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Application parameter name.
        /// </summary>
        [Input("parameter")]
        public Input<string>? Parameter { get; set; }

        [Input("parameters")]
        private InputList<Inputs.NameParameterArgs>? _parameters;

        /// <summary>
        /// Application parameters. The structure of `parameters` block is documented below.
        /// </summary>
        public InputList<Inputs.NameParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.NameParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Application popularity.
        /// </summary>
        [Input("popularity")]
        public Input<int>? Popularity { get; set; }

        /// <summary>
        /// Application protocol.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Application risk.
        /// </summary>
        [Input("risk")]
        public Input<int>? Risk { get; set; }

        /// <summary>
        /// Application sub-category ID.
        /// </summary>
        [Input("subCategory")]
        public Input<int>? SubCategory { get; set; }

        /// <summary>
        /// Application technology.
        /// </summary>
        [Input("technology")]
        public Input<string>? Technology { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Application vendor.
        /// </summary>
        [Input("vendor")]
        public Input<string>? Vendor { get; set; }

        /// <summary>
        /// Application weight.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ApplicationNameArgs()
        {
        }
        public static new ApplicationNameArgs Empty => new ApplicationNameArgs();
    }

    public sealed class ApplicationNameState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application behavior.
        /// </summary>
        [Input("behavior")]
        public Input<string>? Behavior { get; set; }

        /// <summary>
        /// Application category ID.
        /// </summary>
        [Input("category")]
        public Input<int>? Category { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Application ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        [Input("metadatas")]
        private InputList<Inputs.NameMetadataGetArgs>? _metadatas;

        /// <summary>
        /// Meta data. The structure of `metadata` block is documented below.
        /// </summary>
        public InputList<Inputs.NameMetadataGetArgs> Metadatas
        {
            get => _metadatas ?? (_metadatas = new InputList<Inputs.NameMetadataGetArgs>());
            set => _metadatas = value;
        }

        /// <summary>
        /// Application name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Application parameter name.
        /// </summary>
        [Input("parameter")]
        public Input<string>? Parameter { get; set; }

        [Input("parameters")]
        private InputList<Inputs.NameParameterGetArgs>? _parameters;

        /// <summary>
        /// Application parameters. The structure of `parameters` block is documented below.
        /// </summary>
        public InputList<Inputs.NameParameterGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.NameParameterGetArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Application popularity.
        /// </summary>
        [Input("popularity")]
        public Input<int>? Popularity { get; set; }

        /// <summary>
        /// Application protocol.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Application risk.
        /// </summary>
        [Input("risk")]
        public Input<int>? Risk { get; set; }

        /// <summary>
        /// Application sub-category ID.
        /// </summary>
        [Input("subCategory")]
        public Input<int>? SubCategory { get; set; }

        /// <summary>
        /// Application technology.
        /// </summary>
        [Input("technology")]
        public Input<string>? Technology { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Application vendor.
        /// </summary>
        [Input("vendor")]
        public Input<string>? Vendor { get; set; }

        /// <summary>
        /// Application weight.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ApplicationNameState()
        {
        }
        public static new ApplicationNameState Empty => new ApplicationNameState();
    }
}
