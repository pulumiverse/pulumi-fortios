// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    /// <summary>
    /// Configure object tagging.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.System.Objecttagging("trname", new()
    ///     {
    ///         Address = "disable",
    ///         Category = "s1",
    ///         Color = 0,
    ///         Device = "mandatory",
    ///         Interface = "disable",
    ///         Multiple = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// System ObjectTagging can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/objecttagging:Objecttagging labelname {{category}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/objecttagging:Objecttagging labelname {{category}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/objecttagging:Objecttagging")]
    public partial class Objecttagging : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Address. Valid values: `disable`, `mandatory`, `optional`.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Tag Category.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// Device. Valid values: `disable`, `mandatory`, `optional`.
        /// </summary>
        [Output("device")]
        public Output<string> Device { get; private set; } = null!;

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
        /// Interface. Valid values: `disable`, `mandatory`, `optional`.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Allow multiple tag selection. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("multiple")]
        public Output<string> Multiple { get; private set; } = null!;

        /// <summary>
        /// Tags. The structure of `tags` block is documented below.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ObjecttaggingTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Objecttagging resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Objecttagging(string name, ObjecttaggingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/objecttagging:Objecttagging", name, args ?? new ObjecttaggingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Objecttagging(string name, Input<string> id, ObjecttaggingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/objecttagging:Objecttagging", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Objecttagging resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Objecttagging Get(string name, Input<string> id, ObjecttaggingState? state = null, CustomResourceOptions? options = null)
        {
            return new Objecttagging(name, id, state, options);
        }
    }

    public sealed class ObjecttaggingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address. Valid values: `disable`, `mandatory`, `optional`.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Tag Category.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Device. Valid values: `disable`, `mandatory`, `optional`.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

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
        /// Interface. Valid values: `disable`, `mandatory`, `optional`.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Allow multiple tag selection. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("multiple")]
        public Input<string>? Multiple { get; set; }

        [Input("tags")]
        private InputList<Inputs.ObjecttaggingTagArgs>? _tags;

        /// <summary>
        /// Tags. The structure of `tags` block is documented below.
        /// </summary>
        public InputList<Inputs.ObjecttaggingTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ObjecttaggingTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ObjecttaggingArgs()
        {
        }
        public static new ObjecttaggingArgs Empty => new ObjecttaggingArgs();
    }

    public sealed class ObjecttaggingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address. Valid values: `disable`, `mandatory`, `optional`.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Tag Category.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Device. Valid values: `disable`, `mandatory`, `optional`.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

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
        /// Interface. Valid values: `disable`, `mandatory`, `optional`.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Allow multiple tag selection. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("multiple")]
        public Input<string>? Multiple { get; set; }

        [Input("tags")]
        private InputList<Inputs.ObjecttaggingTagGetArgs>? _tags;

        /// <summary>
        /// Tags. The structure of `tags` block is documented below.
        /// </summary>
        public InputList<Inputs.ObjecttaggingTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ObjecttaggingTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ObjecttaggingState()
        {
        }
        public static new ObjecttaggingState Empty => new ObjecttaggingState();
    }
}
