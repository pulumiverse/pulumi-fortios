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
    /// Configure community lists.
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
    ///     var trname = new Fortios.Router.Communitylist("trname", new()
    ///     {
    ///         Rules = new[]
    ///         {
    ///             new Fortios.Router.Inputs.CommunitylistRuleArgs
    ///             {
    ///                 Action = "deny",
    ///                 Match = "123:234 345:456",
    ///                 Regexp = "123:234 345:456",
    ///             },
    ///         },
    ///         Type = "standard",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Router CommunityList can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/communitylist:Communitylist labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/communitylist:Communitylist labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:router/communitylist:Communitylist")]
    public partial class Communitylist : global::Pulumi.CustomResource
    {
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
        /// Community list name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Community list rule. The structure of `rule` block is documented below.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.CommunitylistRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Communitylist resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Communitylist(string name, CommunitylistArgs args, CustomResourceOptions? options = null)
            : base("fortios:router/communitylist:Communitylist", name, args ?? new CommunitylistArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Communitylist(string name, Input<string> id, CommunitylistState? state = null, CustomResourceOptions? options = null)
            : base("fortios:router/communitylist:Communitylist", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Communitylist resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Communitylist Get(string name, Input<string> id, CommunitylistState? state = null, CustomResourceOptions? options = null)
        {
            return new Communitylist(name, id, state, options);
        }
    }

    public sealed class CommunitylistArgs : global::Pulumi.ResourceArgs
    {
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
        /// Community list name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rules")]
        private InputList<Inputs.CommunitylistRuleArgs>? _rules;

        /// <summary>
        /// Community list rule. The structure of `rule` block is documented below.
        /// </summary>
        public InputList<Inputs.CommunitylistRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.CommunitylistRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CommunitylistArgs()
        {
        }
        public static new CommunitylistArgs Empty => new CommunitylistArgs();
    }

    public sealed class CommunitylistState : global::Pulumi.ResourceArgs
    {
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
        /// Community list name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rules")]
        private InputList<Inputs.CommunitylistRuleGetArgs>? _rules;

        /// <summary>
        /// Community list rule. The structure of `rule` block is documented below.
        /// </summary>
        public InputList<Inputs.CommunitylistRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.CommunitylistRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CommunitylistState()
        {
        }
        public static new CommunitylistState Empty => new CommunitylistState();
    }
}
