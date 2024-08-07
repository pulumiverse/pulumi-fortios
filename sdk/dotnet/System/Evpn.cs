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
    /// Configure EVPN instance. Applies to FortiOS Version `&gt;= 7.4.0`.
    /// 
    /// ## Import
    /// 
    /// System Evpn can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/evpn:Evpn labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/evpn:Evpn labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/evpn:Evpn")]
    public partial class Evpn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable ARP suppression. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("arpSuppression")]
        public Output<string> ArpSuppression { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// List of export route targets. The structure of `export_rt` block is documented below.
        /// </summary>
        [Output("exportRts")]
        public Output<ImmutableArray<Outputs.EvpnExportRt>> ExportRts { get; private set; } = null!;

        /// <summary>
        /// ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// List of import route targets. The structure of `import_rt` block is documented below.
        /// </summary>
        [Output("importRts")]
        public Output<ImmutableArray<Outputs.EvpnImportRt>> ImportRts { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IP address local learning. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipLocalLearning")]
        public Output<string> IpLocalLearning { get; private set; } = null!;

        /// <summary>
        /// Route Distinguisher: AA|AA:NN|A.B.C.D:NN.
        /// </summary>
        [Output("rd")]
        public Output<string> Rd { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Evpn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Evpn(string name, EvpnArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/evpn:Evpn", name, args ?? new EvpnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Evpn(string name, Input<string> id, EvpnState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/evpn:Evpn", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Evpn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Evpn Get(string name, Input<string> id, EvpnState? state = null, CustomResourceOptions? options = null)
        {
            return new Evpn(name, id, state, options);
        }
    }

    public sealed class EvpnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable ARP suppression. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("arpSuppression")]
        public Input<string>? ArpSuppression { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("exportRts")]
        private InputList<Inputs.EvpnExportRtArgs>? _exportRts;

        /// <summary>
        /// List of export route targets. The structure of `export_rt` block is documented below.
        /// </summary>
        public InputList<Inputs.EvpnExportRtArgs> ExportRts
        {
            get => _exportRts ?? (_exportRts = new InputList<Inputs.EvpnExportRtArgs>());
            set => _exportRts = value;
        }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        [Input("importRts")]
        private InputList<Inputs.EvpnImportRtArgs>? _importRts;

        /// <summary>
        /// List of import route targets. The structure of `import_rt` block is documented below.
        /// </summary>
        public InputList<Inputs.EvpnImportRtArgs> ImportRts
        {
            get => _importRts ?? (_importRts = new InputList<Inputs.EvpnImportRtArgs>());
            set => _importRts = value;
        }

        /// <summary>
        /// Enable/disable IP address local learning. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipLocalLearning")]
        public Input<string>? IpLocalLearning { get; set; }

        /// <summary>
        /// Route Distinguisher: AA|AA:NN|A.B.C.D:NN.
        /// </summary>
        [Input("rd")]
        public Input<string>? Rd { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public EvpnArgs()
        {
        }
        public static new EvpnArgs Empty => new EvpnArgs();
    }

    public sealed class EvpnState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable ARP suppression. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("arpSuppression")]
        public Input<string>? ArpSuppression { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("exportRts")]
        private InputList<Inputs.EvpnExportRtGetArgs>? _exportRts;

        /// <summary>
        /// List of export route targets. The structure of `export_rt` block is documented below.
        /// </summary>
        public InputList<Inputs.EvpnExportRtGetArgs> ExportRts
        {
            get => _exportRts ?? (_exportRts = new InputList<Inputs.EvpnExportRtGetArgs>());
            set => _exportRts = value;
        }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        [Input("importRts")]
        private InputList<Inputs.EvpnImportRtGetArgs>? _importRts;

        /// <summary>
        /// List of import route targets. The structure of `import_rt` block is documented below.
        /// </summary>
        public InputList<Inputs.EvpnImportRtGetArgs> ImportRts
        {
            get => _importRts ?? (_importRts = new InputList<Inputs.EvpnImportRtGetArgs>());
            set => _importRts = value;
        }

        /// <summary>
        /// Enable/disable IP address local learning. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipLocalLearning")]
        public Input<string>? IpLocalLearning { get; set; }

        /// <summary>
        /// Route Distinguisher: AA|AA:NN|A.B.C.D:NN.
        /// </summary>
        [Input("rd")]
        public Input<string>? Rd { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public EvpnState()
        {
        }
        public static new EvpnState Empty => new EvpnState();
    }
}
