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
    /// Configure sFlow.
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
    ///     var trname = new Fortios.System.Sflow("trname", new()
    ///     {
    ///         CollectorIp = "0.0.0.0",
    ///         CollectorPort = 6343,
    ///         SourceIp = "0.0.0.0",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System Sflow can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/sflow:Sflow labelname SystemSflow
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/sflow:Sflow labelname SystemSflow
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/sflow:Sflow")]
    public partial class Sflow : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
        /// </summary>
        [Output("collectorIp")]
        public Output<string> CollectorIp { get; private set; } = null!;

        /// <summary>
        /// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
        /// </summary>
        [Output("collectorPort")]
        public Output<int> CollectorPort { get; private set; } = null!;

        /// <summary>
        /// sFlow collectors. The structure of `collectors` block is documented below.
        /// </summary>
        [Output("collectors")]
        public Output<ImmutableArray<Outputs.SflowCollector>> Collectors { get; private set; } = null!;

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
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Output("interfaceSelectMethod")]
        public Output<string> InterfaceSelectMethod { get; private set; } = null!;

        /// <summary>
        /// Source IP address for sFlow agent.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Sflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Sflow(string name, SflowArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/sflow:Sflow", name, args ?? new SflowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Sflow(string name, Input<string> id, SflowState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/sflow:Sflow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Sflow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Sflow Get(string name, Input<string> id, SflowState? state = null, CustomResourceOptions? options = null)
        {
            return new Sflow(name, id, state, options);
        }
    }

    public sealed class SflowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
        /// </summary>
        [Input("collectorIp", required: true)]
        public Input<string> CollectorIp { get; set; } = null!;

        /// <summary>
        /// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
        /// </summary>
        [Input("collectorPort")]
        public Input<int>? CollectorPort { get; set; }

        [Input("collectors")]
        private InputList<Inputs.SflowCollectorArgs>? _collectors;

        /// <summary>
        /// sFlow collectors. The structure of `collectors` block is documented below.
        /// </summary>
        public InputList<Inputs.SflowCollectorArgs> Collectors
        {
            get => _collectors ?? (_collectors = new InputList<Inputs.SflowCollectorArgs>());
            set => _collectors = value;
        }

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
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Source IP address for sFlow agent.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SflowArgs()
        {
        }
        public static new SflowArgs Empty => new SflowArgs();
    }

    public sealed class SflowState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
        /// </summary>
        [Input("collectorIp")]
        public Input<string>? CollectorIp { get; set; }

        /// <summary>
        /// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
        /// </summary>
        [Input("collectorPort")]
        public Input<int>? CollectorPort { get; set; }

        [Input("collectors")]
        private InputList<Inputs.SflowCollectorGetArgs>? _collectors;

        /// <summary>
        /// sFlow collectors. The structure of `collectors` block is documented below.
        /// </summary>
        public InputList<Inputs.SflowCollectorGetArgs> Collectors
        {
            get => _collectors ?? (_collectors = new InputList<Inputs.SflowCollectorGetArgs>());
            set => _collectors = value;
        }

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
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Source IP address for sFlow agent.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SflowState()
        {
        }
        public static new SflowState Empty => new SflowState();
    }
}
