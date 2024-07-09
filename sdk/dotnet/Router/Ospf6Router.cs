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
    /// Configure IPv6 OSPF.
    /// 
    /// &gt; The provider supports the definition of Ospf6-Interface in Router Ospf6 `fortios.router.Ospf6`, and also allows the definition of separate Ospf6-Interface resources `fortios.router/ospf6.Ospf6interface`, but do not use a `fortios.router.Ospf6` with in-line Ospf6-Interface in conjunction with any `fortios.router/ospf6.Ospf6interface` resources, otherwise conflicts and overwrite will occur.
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
    ///     var trname = new Fortios.Router.Ospf6Router("trname", new()
    ///     {
    ///         AbrType = "standard",
    ///         AutoCostRefBandwidth = 1000,
    ///         Bfd = "disable",
    ///         DefaultInformationMetric = 10,
    ///         DefaultInformationMetricType = "2",
    ///         DefaultInformationOriginate = "disable",
    ///         DefaultMetric = 10,
    ///         LogNeighbourChanges = "enable",
    ///         Redistributes = new[]
    ///         {
    ///             new Fortios.Router.Inputs.Ospf6RedistributeArgs
    ///             {
    ///                 Metric = 0,
    ///                 MetricType = "2",
    ///                 Name = "connected",
    ///                 Status = "disable",
    ///             },
    ///             new Fortios.Router.Inputs.Ospf6RedistributeArgs
    ///             {
    ///                 Metric = 0,
    ///                 MetricType = "2",
    ///                 Name = "static",
    ///                 Status = "disable",
    ///             },
    ///             new Fortios.Router.Inputs.Ospf6RedistributeArgs
    ///             {
    ///                 Metric = 0,
    ///                 MetricType = "2",
    ///                 Name = "rip",
    ///                 Status = "disable",
    ///             },
    ///             new Fortios.Router.Inputs.Ospf6RedistributeArgs
    ///             {
    ///                 Metric = 0,
    ///                 MetricType = "2",
    ///                 Name = "bgp",
    ///                 Status = "disable",
    ///             },
    ///             new Fortios.Router.Inputs.Ospf6RedistributeArgs
    ///             {
    ///                 Metric = 0,
    ///                 MetricType = "2",
    ///                 Name = "isis",
    ///                 Status = "disable",
    ///             },
    ///         },
    ///         RouterId = "0.0.0.0",
    ///         SpfTimers = "5 10",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Router Ospf6 can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/ospf6:Ospf6 labelname RouterOspf6
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/ospf6:Ospf6 labelname RouterOspf6
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:router/ospf6:Ospf6")]
    public partial class Ospf6Router : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Area border router type. Valid values: `cisco`, `ibm`, `standard`.
        /// </summary>
        [Output("abrType")]
        public Output<string> AbrType { get; private set; } = null!;

        /// <summary>
        /// OSPF6 area configuration. The structure of `area` block is documented below.
        /// </summary>
        [Output("areas")]
        public Output<ImmutableArray<Outputs.Ospf6Area>> Areas { get; private set; } = null!;

        /// <summary>
        /// Reference bandwidth in terms of megabits per second.
        /// </summary>
        [Output("autoCostRefBandwidth")]
        public Output<int> AutoCostRefBandwidth { get; private set; } = null!;

        /// <summary>
        /// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("bfd")]
        public Output<string> Bfd { get; private set; } = null!;

        /// <summary>
        /// Default information metric.
        /// </summary>
        [Output("defaultInformationMetric")]
        public Output<int> DefaultInformationMetric { get; private set; } = null!;

        /// <summary>
        /// Default information metric type. Valid values: `1`, `2`.
        /// </summary>
        [Output("defaultInformationMetricType")]
        public Output<string> DefaultInformationMetricType { get; private set; } = null!;

        /// <summary>
        /// Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
        /// </summary>
        [Output("defaultInformationOriginate")]
        public Output<string> DefaultInformationOriginate { get; private set; } = null!;

        /// <summary>
        /// Default information route map.
        /// </summary>
        [Output("defaultInformationRouteMap")]
        public Output<string> DefaultInformationRouteMap { get; private set; } = null!;

        /// <summary>
        /// Default metric of redistribute routes.
        /// </summary>
        [Output("defaultMetric")]
        public Output<int> DefaultMetric { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Enable logging of OSPFv3 neighbour's changes Valid values: `enable`, `disable`.
        /// </summary>
        [Output("logNeighbourChanges")]
        public Output<string> LogNeighbourChanges { get; private set; } = null!;

        /// <summary>
        /// OSPF6 interface configuration. The structure of `ospf6_interface` block is documented below.
        /// </summary>
        [Output("ospf6Interfaces")]
        public Output<ImmutableArray<Outputs.Ospf6Ospf6Interface>> Ospf6Interfaces { get; private set; } = null!;

        /// <summary>
        /// Passive interface configuration. The structure of `passive_interface` block is documented below.
        /// </summary>
        [Output("passiveInterfaces")]
        public Output<ImmutableArray<Outputs.Ospf6PassiveInterface>> PassiveInterfaces { get; private set; } = null!;

        /// <summary>
        /// Redistribute configuration. The structure of `redistribute` block is documented below.
        /// </summary>
        [Output("redistributes")]
        public Output<ImmutableArray<Outputs.Ospf6Redistribute>> Redistributes { get; private set; } = null!;

        /// <summary>
        /// OSPFv3 restart mode (graceful or none). Valid values: `none`, `graceful-restart`.
        /// </summary>
        [Output("restartMode")]
        public Output<string> RestartMode { get; private set; } = null!;

        /// <summary>
        /// Enable/disable continuing graceful restart upon topology change. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("restartOnTopologyChange")]
        public Output<string> RestartOnTopologyChange { get; private set; } = null!;

        /// <summary>
        /// Graceful restart period in seconds.
        /// </summary>
        [Output("restartPeriod")]
        public Output<int> RestartPeriod { get; private set; } = null!;

        /// <summary>
        /// A.B.C.D, in IPv4 address format.
        /// </summary>
        [Output("routerId")]
        public Output<string> RouterId { get; private set; } = null!;

        /// <summary>
        /// SPF calculation frequency.
        /// </summary>
        [Output("spfTimers")]
        public Output<string> SpfTimers { get; private set; } = null!;

        /// <summary>
        /// IPv6 address summary configuration. The structure of `summary_address` block is documented below.
        /// </summary>
        [Output("summaryAddresses")]
        public Output<ImmutableArray<Outputs.Ospf6SummaryAddress>> SummaryAddresses { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ospf6Router resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ospf6Router(string name, Ospf6RouterArgs args, CustomResourceOptions? options = null)
            : base("fortios:router/ospf6:Ospf6", name, args ?? new Ospf6RouterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ospf6Router(string name, Input<string> id, Ospf6RouterState? state = null, CustomResourceOptions? options = null)
            : base("fortios:router/ospf6:Ospf6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ospf6Router resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ospf6Router Get(string name, Input<string> id, Ospf6RouterState? state = null, CustomResourceOptions? options = null)
        {
            return new Ospf6Router(name, id, state, options);
        }
    }

    public sealed class Ospf6RouterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Area border router type. Valid values: `cisco`, `ibm`, `standard`.
        /// </summary>
        [Input("abrType")]
        public Input<string>? AbrType { get; set; }

        [Input("areas")]
        private InputList<Inputs.Ospf6AreaArgs>? _areas;

        /// <summary>
        /// OSPF6 area configuration. The structure of `area` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6AreaArgs> Areas
        {
            get => _areas ?? (_areas = new InputList<Inputs.Ospf6AreaArgs>());
            set => _areas = value;
        }

        /// <summary>
        /// Reference bandwidth in terms of megabits per second.
        /// </summary>
        [Input("autoCostRefBandwidth")]
        public Input<int>? AutoCostRefBandwidth { get; set; }

        /// <summary>
        /// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bfd")]
        public Input<string>? Bfd { get; set; }

        /// <summary>
        /// Default information metric.
        /// </summary>
        [Input("defaultInformationMetric")]
        public Input<int>? DefaultInformationMetric { get; set; }

        /// <summary>
        /// Default information metric type. Valid values: `1`, `2`.
        /// </summary>
        [Input("defaultInformationMetricType")]
        public Input<string>? DefaultInformationMetricType { get; set; }

        /// <summary>
        /// Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
        /// </summary>
        [Input("defaultInformationOriginate")]
        public Input<string>? DefaultInformationOriginate { get; set; }

        /// <summary>
        /// Default information route map.
        /// </summary>
        [Input("defaultInformationRouteMap")]
        public Input<string>? DefaultInformationRouteMap { get; set; }

        /// <summary>
        /// Default metric of redistribute routes.
        /// </summary>
        [Input("defaultMetric")]
        public Input<int>? DefaultMetric { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Enable logging of OSPFv3 neighbour's changes Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logNeighbourChanges")]
        public Input<string>? LogNeighbourChanges { get; set; }

        [Input("ospf6Interfaces")]
        private InputList<Inputs.Ospf6Ospf6InterfaceArgs>? _ospf6Interfaces;

        /// <summary>
        /// OSPF6 interface configuration. The structure of `ospf6_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6Ospf6InterfaceArgs> Ospf6Interfaces
        {
            get => _ospf6Interfaces ?? (_ospf6Interfaces = new InputList<Inputs.Ospf6Ospf6InterfaceArgs>());
            set => _ospf6Interfaces = value;
        }

        [Input("passiveInterfaces")]
        private InputList<Inputs.Ospf6PassiveInterfaceArgs>? _passiveInterfaces;

        /// <summary>
        /// Passive interface configuration. The structure of `passive_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6PassiveInterfaceArgs> PassiveInterfaces
        {
            get => _passiveInterfaces ?? (_passiveInterfaces = new InputList<Inputs.Ospf6PassiveInterfaceArgs>());
            set => _passiveInterfaces = value;
        }

        [Input("redistributes")]
        private InputList<Inputs.Ospf6RedistributeArgs>? _redistributes;

        /// <summary>
        /// Redistribute configuration. The structure of `redistribute` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6RedistributeArgs> Redistributes
        {
            get => _redistributes ?? (_redistributes = new InputList<Inputs.Ospf6RedistributeArgs>());
            set => _redistributes = value;
        }

        /// <summary>
        /// OSPFv3 restart mode (graceful or none). Valid values: `none`, `graceful-restart`.
        /// </summary>
        [Input("restartMode")]
        public Input<string>? RestartMode { get; set; }

        /// <summary>
        /// Enable/disable continuing graceful restart upon topology change. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("restartOnTopologyChange")]
        public Input<string>? RestartOnTopologyChange { get; set; }

        /// <summary>
        /// Graceful restart period in seconds.
        /// </summary>
        [Input("restartPeriod")]
        public Input<int>? RestartPeriod { get; set; }

        /// <summary>
        /// A.B.C.D, in IPv4 address format.
        /// </summary>
        [Input("routerId", required: true)]
        public Input<string> RouterId { get; set; } = null!;

        /// <summary>
        /// SPF calculation frequency.
        /// </summary>
        [Input("spfTimers")]
        public Input<string>? SpfTimers { get; set; }

        [Input("summaryAddresses")]
        private InputList<Inputs.Ospf6SummaryAddressArgs>? _summaryAddresses;

        /// <summary>
        /// IPv6 address summary configuration. The structure of `summary_address` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6SummaryAddressArgs> SummaryAddresses
        {
            get => _summaryAddresses ?? (_summaryAddresses = new InputList<Inputs.Ospf6SummaryAddressArgs>());
            set => _summaryAddresses = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Ospf6RouterArgs()
        {
        }
        public static new Ospf6RouterArgs Empty => new Ospf6RouterArgs();
    }

    public sealed class Ospf6RouterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Area border router type. Valid values: `cisco`, `ibm`, `standard`.
        /// </summary>
        [Input("abrType")]
        public Input<string>? AbrType { get; set; }

        [Input("areas")]
        private InputList<Inputs.Ospf6AreaGetArgs>? _areas;

        /// <summary>
        /// OSPF6 area configuration. The structure of `area` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6AreaGetArgs> Areas
        {
            get => _areas ?? (_areas = new InputList<Inputs.Ospf6AreaGetArgs>());
            set => _areas = value;
        }

        /// <summary>
        /// Reference bandwidth in terms of megabits per second.
        /// </summary>
        [Input("autoCostRefBandwidth")]
        public Input<int>? AutoCostRefBandwidth { get; set; }

        /// <summary>
        /// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bfd")]
        public Input<string>? Bfd { get; set; }

        /// <summary>
        /// Default information metric.
        /// </summary>
        [Input("defaultInformationMetric")]
        public Input<int>? DefaultInformationMetric { get; set; }

        /// <summary>
        /// Default information metric type. Valid values: `1`, `2`.
        /// </summary>
        [Input("defaultInformationMetricType")]
        public Input<string>? DefaultInformationMetricType { get; set; }

        /// <summary>
        /// Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
        /// </summary>
        [Input("defaultInformationOriginate")]
        public Input<string>? DefaultInformationOriginate { get; set; }

        /// <summary>
        /// Default information route map.
        /// </summary>
        [Input("defaultInformationRouteMap")]
        public Input<string>? DefaultInformationRouteMap { get; set; }

        /// <summary>
        /// Default metric of redistribute routes.
        /// </summary>
        [Input("defaultMetric")]
        public Input<int>? DefaultMetric { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Enable logging of OSPFv3 neighbour's changes Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logNeighbourChanges")]
        public Input<string>? LogNeighbourChanges { get; set; }

        [Input("ospf6Interfaces")]
        private InputList<Inputs.Ospf6Ospf6InterfaceGetArgs>? _ospf6Interfaces;

        /// <summary>
        /// OSPF6 interface configuration. The structure of `ospf6_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6Ospf6InterfaceGetArgs> Ospf6Interfaces
        {
            get => _ospf6Interfaces ?? (_ospf6Interfaces = new InputList<Inputs.Ospf6Ospf6InterfaceGetArgs>());
            set => _ospf6Interfaces = value;
        }

        [Input("passiveInterfaces")]
        private InputList<Inputs.Ospf6PassiveInterfaceGetArgs>? _passiveInterfaces;

        /// <summary>
        /// Passive interface configuration. The structure of `passive_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6PassiveInterfaceGetArgs> PassiveInterfaces
        {
            get => _passiveInterfaces ?? (_passiveInterfaces = new InputList<Inputs.Ospf6PassiveInterfaceGetArgs>());
            set => _passiveInterfaces = value;
        }

        [Input("redistributes")]
        private InputList<Inputs.Ospf6RedistributeGetArgs>? _redistributes;

        /// <summary>
        /// Redistribute configuration. The structure of `redistribute` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6RedistributeGetArgs> Redistributes
        {
            get => _redistributes ?? (_redistributes = new InputList<Inputs.Ospf6RedistributeGetArgs>());
            set => _redistributes = value;
        }

        /// <summary>
        /// OSPFv3 restart mode (graceful or none). Valid values: `none`, `graceful-restart`.
        /// </summary>
        [Input("restartMode")]
        public Input<string>? RestartMode { get; set; }

        /// <summary>
        /// Enable/disable continuing graceful restart upon topology change. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("restartOnTopologyChange")]
        public Input<string>? RestartOnTopologyChange { get; set; }

        /// <summary>
        /// Graceful restart period in seconds.
        /// </summary>
        [Input("restartPeriod")]
        public Input<int>? RestartPeriod { get; set; }

        /// <summary>
        /// A.B.C.D, in IPv4 address format.
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        /// <summary>
        /// SPF calculation frequency.
        /// </summary>
        [Input("spfTimers")]
        public Input<string>? SpfTimers { get; set; }

        [Input("summaryAddresses")]
        private InputList<Inputs.Ospf6SummaryAddressGetArgs>? _summaryAddresses;

        /// <summary>
        /// IPv6 address summary configuration. The structure of `summary_address` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6SummaryAddressGetArgs> SummaryAddresses
        {
            get => _summaryAddresses ?? (_summaryAddresses = new InputList<Inputs.Ospf6SummaryAddressGetArgs>());
            set => _summaryAddresses = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Ospf6RouterState()
        {
        }
        public static new Ospf6RouterState Empty => new Ospf6RouterState();
    }
}
