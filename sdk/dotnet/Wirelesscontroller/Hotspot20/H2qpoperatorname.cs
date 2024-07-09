// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Hotspot20
{
    /// <summary>
    /// Configure operator friendly name.
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
    ///     var trname = new Fortios.Wirelesscontroller.Hotspot20.H2qpoperatorname("trname");
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// WirelessControllerHotspot20 H2QpOperatorName can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/hotspot20/h2qpoperatorname:H2qpoperatorname labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/hotspot20/h2qpoperatorname:H2qpoperatorname labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontroller/hotspot20/h2qpoperatorname:H2qpoperatorname")]
    public partial class H2qpoperatorname : global::Pulumi.CustomResource
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
        /// Friendly name ID.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name list. The structure of `value_list` block is documented below.
        /// </summary>
        [Output("valueLists")]
        public Output<ImmutableArray<Outputs.H2qpoperatornameValueList>> ValueLists { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a H2qpoperatorname resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public H2qpoperatorname(string name, H2qpoperatornameArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/hotspot20/h2qpoperatorname:H2qpoperatorname", name, args ?? new H2qpoperatornameArgs(), MakeResourceOptions(options, ""))
        {
        }

        private H2qpoperatorname(string name, Input<string> id, H2qpoperatornameState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/hotspot20/h2qpoperatorname:H2qpoperatorname", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing H2qpoperatorname resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static H2qpoperatorname Get(string name, Input<string> id, H2qpoperatornameState? state = null, CustomResourceOptions? options = null)
        {
            return new H2qpoperatorname(name, id, state, options);
        }
    }

    public sealed class H2qpoperatornameArgs : global::Pulumi.ResourceArgs
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
        /// Friendly name ID.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("valueLists")]
        private InputList<Inputs.H2qpoperatornameValueListArgs>? _valueLists;

        /// <summary>
        /// Name list. The structure of `value_list` block is documented below.
        /// </summary>
        public InputList<Inputs.H2qpoperatornameValueListArgs> ValueLists
        {
            get => _valueLists ?? (_valueLists = new InputList<Inputs.H2qpoperatornameValueListArgs>());
            set => _valueLists = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public H2qpoperatornameArgs()
        {
        }
        public static new H2qpoperatornameArgs Empty => new H2qpoperatornameArgs();
    }

    public sealed class H2qpoperatornameState : global::Pulumi.ResourceArgs
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
        /// Friendly name ID.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("valueLists")]
        private InputList<Inputs.H2qpoperatornameValueListGetArgs>? _valueLists;

        /// <summary>
        /// Name list. The structure of `value_list` block is documented below.
        /// </summary>
        public InputList<Inputs.H2qpoperatornameValueListGetArgs> ValueLists
        {
            get => _valueLists ?? (_valueLists = new InputList<Inputs.H2qpoperatornameValueListGetArgs>());
            set => _valueLists = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public H2qpoperatornameState()
        {
        }
        public static new H2qpoperatornameState Empty => new H2qpoperatornameState();
    }
}
