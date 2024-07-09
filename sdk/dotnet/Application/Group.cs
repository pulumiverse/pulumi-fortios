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
    /// Configure firewall application groups.
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
    ///     var trname = new Fortios.Application.Group("trname", new()
    ///     {
    ///         Categories = new[]
    ///         {
    ///             new Fortios.Application.Inputs.GroupCategoryArgs
    ///             {
    ///                 Id = 2,
    ///             },
    ///         },
    ///         Comment = "group1 test",
    ///         Type = "category",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Application Group can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:application/group:Group labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:application/group:Group labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:application/group:Group")]
    public partial class Group : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Application ID list. The structure of `application` block is documented below.
        /// </summary>
        [Output("applications")]
        public Output<ImmutableArray<Outputs.GroupApplication>> Applications { get; private set; } = null!;

        /// <summary>
        /// Application behavior filter.
        /// </summary>
        [Output("behavior")]
        public Output<string> Behavior { get; private set; } = null!;

        /// <summary>
        /// Application category ID list. The structure of `category` block is documented below.
        /// </summary>
        [Output("categories")]
        public Output<ImmutableArray<Outputs.GroupCategory>> Categories { get; private set; } = null!;

        /// <summary>
        /// Comment
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

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
        /// Application group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
        /// </summary>
        [Output("popularity")]
        public Output<string> Popularity { get; private set; } = null!;

        /// <summary>
        /// Application protocol filter.
        /// </summary>
        [Output("protocols")]
        public Output<string> Protocols { get; private set; } = null!;

        /// <summary>
        /// Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
        /// </summary>
        [Output("risks")]
        public Output<ImmutableArray<Outputs.GroupRisk>> Risks { get; private set; } = null!;

        /// <summary>
        /// Application technology filter.
        /// </summary>
        [Output("technology")]
        public Output<string> Technology { get; private set; } = null!;

        /// <summary>
        /// Application group type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Application vendor filter.
        /// </summary>
        [Output("vendor")]
        public Output<string> Vendor { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:application/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("fortios:application/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("applications")]
        private InputList<Inputs.GroupApplicationArgs>? _applications;

        /// <summary>
        /// Application ID list. The structure of `application` block is documented below.
        /// </summary>
        public InputList<Inputs.GroupApplicationArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.GroupApplicationArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// Application behavior filter.
        /// </summary>
        [Input("behavior")]
        public Input<string>? Behavior { get; set; }

        [Input("categories")]
        private InputList<Inputs.GroupCategoryArgs>? _categories;

        /// <summary>
        /// Application category ID list. The structure of `category` block is documented below.
        /// </summary>
        public InputList<Inputs.GroupCategoryArgs> Categories
        {
            get => _categories ?? (_categories = new InputList<Inputs.GroupCategoryArgs>());
            set => _categories = value;
        }

        /// <summary>
        /// Comment
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

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
        /// Application group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
        /// </summary>
        [Input("popularity")]
        public Input<string>? Popularity { get; set; }

        /// <summary>
        /// Application protocol filter.
        /// </summary>
        [Input("protocols")]
        public Input<string>? Protocols { get; set; }

        [Input("risks")]
        private InputList<Inputs.GroupRiskArgs>? _risks;

        /// <summary>
        /// Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
        /// </summary>
        public InputList<Inputs.GroupRiskArgs> Risks
        {
            get => _risks ?? (_risks = new InputList<Inputs.GroupRiskArgs>());
            set => _risks = value;
        }

        /// <summary>
        /// Application technology filter.
        /// </summary>
        [Input("technology")]
        public Input<string>? Technology { get; set; }

        /// <summary>
        /// Application group type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Application vendor filter.
        /// </summary>
        [Input("vendor")]
        public Input<string>? Vendor { get; set; }

        public GroupArgs()
        {
        }
        public static new GroupArgs Empty => new GroupArgs();
    }

    public sealed class GroupState : global::Pulumi.ResourceArgs
    {
        [Input("applications")]
        private InputList<Inputs.GroupApplicationGetArgs>? _applications;

        /// <summary>
        /// Application ID list. The structure of `application` block is documented below.
        /// </summary>
        public InputList<Inputs.GroupApplicationGetArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.GroupApplicationGetArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// Application behavior filter.
        /// </summary>
        [Input("behavior")]
        public Input<string>? Behavior { get; set; }

        [Input("categories")]
        private InputList<Inputs.GroupCategoryGetArgs>? _categories;

        /// <summary>
        /// Application category ID list. The structure of `category` block is documented below.
        /// </summary>
        public InputList<Inputs.GroupCategoryGetArgs> Categories
        {
            get => _categories ?? (_categories = new InputList<Inputs.GroupCategoryGetArgs>());
            set => _categories = value;
        }

        /// <summary>
        /// Comment
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

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
        /// Application group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
        /// </summary>
        [Input("popularity")]
        public Input<string>? Popularity { get; set; }

        /// <summary>
        /// Application protocol filter.
        /// </summary>
        [Input("protocols")]
        public Input<string>? Protocols { get; set; }

        [Input("risks")]
        private InputList<Inputs.GroupRiskGetArgs>? _risks;

        /// <summary>
        /// Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
        /// </summary>
        public InputList<Inputs.GroupRiskGetArgs> Risks
        {
            get => _risks ?? (_risks = new InputList<Inputs.GroupRiskGetArgs>());
            set => _risks = value;
        }

        /// <summary>
        /// Application technology filter.
        /// </summary>
        [Input("technology")]
        public Input<string>? Technology { get; set; }

        /// <summary>
        /// Application group type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Application vendor filter.
        /// </summary>
        [Input("vendor")]
        public Input<string>? Vendor { get; set; }

        public GroupState()
        {
        }
        public static new GroupState Empty => new GroupState();
    }
}
