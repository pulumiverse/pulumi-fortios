// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontrollerhotspot20
{
    /// <summary>
    /// Configure venue name duple.
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
    ///     var trname = new Fortios.Wirelesscontrollerhotspot20.Anqpvenuename("trname", new()
    ///     {
    ///         ValueLists = new[]
    ///         {
    ///             new Fortios.Wirelesscontrollerhotspot20.Inputs.AnqpvenuenameValueListArgs
    ///             {
    ///                 Index = 1,
    ///                 Lang = "CN",
    ///                 Value = "3",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// WirelessControllerHotspot20 AnqpVenueName can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:wirelesscontrollerhotspot20/anqpvenuename:Anqpvenuename labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:wirelesscontrollerhotspot20/anqpvenuename:Anqpvenuename labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontrollerhotspot20/anqpvenuename:Anqpvenuename")]
    public partial class Anqpvenuename : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Name of venue name duple.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name list. The structure of `value_list` block is documented below.
        /// </summary>
        [Output("valueLists")]
        public Output<ImmutableArray<Outputs.AnqpvenuenameValueList>> ValueLists { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Anqpvenuename resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Anqpvenuename(string name, AnqpvenuenameArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontrollerhotspot20/anqpvenuename:Anqpvenuename", name, args ?? new AnqpvenuenameArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Anqpvenuename(string name, Input<string> id, AnqpvenuenameState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontrollerhotspot20/anqpvenuename:Anqpvenuename", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Anqpvenuename resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Anqpvenuename Get(string name, Input<string> id, AnqpvenuenameState? state = null, CustomResourceOptions? options = null)
        {
            return new Anqpvenuename(name, id, state, options);
        }
    }

    public sealed class AnqpvenuenameArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Name of venue name duple.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("valueLists")]
        private InputList<Inputs.AnqpvenuenameValueListArgs>? _valueLists;

        /// <summary>
        /// Name list. The structure of `value_list` block is documented below.
        /// </summary>
        public InputList<Inputs.AnqpvenuenameValueListArgs> ValueLists
        {
            get => _valueLists ?? (_valueLists = new InputList<Inputs.AnqpvenuenameValueListArgs>());
            set => _valueLists = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AnqpvenuenameArgs()
        {
        }
        public static new AnqpvenuenameArgs Empty => new AnqpvenuenameArgs();
    }

    public sealed class AnqpvenuenameState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Name of venue name duple.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("valueLists")]
        private InputList<Inputs.AnqpvenuenameValueListGetArgs>? _valueLists;

        /// <summary>
        /// Name list. The structure of `value_list` block is documented below.
        /// </summary>
        public InputList<Inputs.AnqpvenuenameValueListGetArgs> ValueLists
        {
            get => _valueLists ?? (_valueLists = new InputList<Inputs.AnqpvenuenameValueListGetArgs>());
            set => _valueLists = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AnqpvenuenameState()
        {
        }
        public static new AnqpvenuenameState Empty => new AnqpvenuenameState();
    }
}
