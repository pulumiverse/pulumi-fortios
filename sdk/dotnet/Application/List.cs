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
    /// Configure application control lists.
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
    ///     var trname = new Fortios.Application.List("trname", new()
    ///     {
    ///         AppReplacemsg = "enable",
    ///         DeepAppInspection = "enable",
    ///         EnforceDefaultAppPort = "disable",
    ///         ExtendedLog = "disable",
    ///         Options = "allow-dns",
    ///         OtherApplicationAction = "pass",
    ///         OtherApplicationLog = "disable",
    ///         UnknownApplicationAction = "pass",
    ///         UnknownApplicationLog = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Application List can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:application/list:List labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:application/list:List labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:application/list:List")]
    public partial class List : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("appReplacemsg")]
        public Output<string> AppReplacemsg { get; private set; } = null!;

        /// <summary>
        /// comments
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("controlDefaultNetworkServices")]
        public Output<string> ControlDefaultNetworkServices { get; private set; } = null!;

        /// <summary>
        /// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("deepAppInspection")]
        public Output<string> DeepAppInspection { get; private set; } = null!;

        /// <summary>
        /// Default network service entries. The structure of `default_network_services` block is documented below.
        /// </summary>
        [Output("defaultNetworkServices")]
        public Output<ImmutableArray<Outputs.ListDefaultNetworkService>> DefaultNetworkServices { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("enforceDefaultAppPort")]
        public Output<string> EnforceDefaultAppPort { get; private set; } = null!;

        /// <summary>
        /// Application list entries. The structure of `entries` block is documented below.
        /// </summary>
        [Output("entries")]
        public Output<ImmutableArray<Outputs.ListEntry>> Entries { get; private set; } = null!;

        /// <summary>
        /// Enable/disable extended logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("extendedLog")]
        public Output<string> ExtendedLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("forceInclusionSslDiSigs")]
        public Output<string> ForceInclusionSslDiSigs { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// List name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Basic application protocol signatures allowed by default.
        /// </summary>
        [Output("options")]
        public Output<string> Options { get; private set; } = null!;

        /// <summary>
        /// Action for other applications. Valid values: `pass`, `block`.
        /// </summary>
        [Output("otherApplicationAction")]
        public Output<string> OtherApplicationAction { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("otherApplicationLog")]
        public Output<string> OtherApplicationLog { get; private set; } = null!;

        /// <summary>
        /// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
        /// </summary>
        [Output("p2pBlackList")]
        public Output<string> P2pBlackList { get; private set; } = null!;

        /// <summary>
        /// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
        /// </summary>
        [Output("p2pBlockList")]
        public Output<string> P2pBlockList { get; private set; } = null!;

        /// <summary>
        /// Replacement message group.
        /// </summary>
        [Output("replacemsgGroup")]
        public Output<string> ReplacemsgGroup { get; private set; } = null!;

        /// <summary>
        /// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
        /// </summary>
        [Output("unknownApplicationAction")]
        public Output<string> UnknownApplicationAction { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("unknownApplicationLog")]
        public Output<string> UnknownApplicationLog { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a List resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public List(string name, ListArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:application/list:List", name, args ?? new ListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private List(string name, Input<string> id, ListState? state = null, CustomResourceOptions? options = null)
            : base("fortios:application/list:List", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing List resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static List Get(string name, Input<string> id, ListState? state = null, CustomResourceOptions? options = null)
        {
            return new List(name, id, state, options);
        }
    }

    public sealed class ListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("appReplacemsg")]
        public Input<string>? AppReplacemsg { get; set; }

        /// <summary>
        /// comments
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("controlDefaultNetworkServices")]
        public Input<string>? ControlDefaultNetworkServices { get; set; }

        /// <summary>
        /// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("deepAppInspection")]
        public Input<string>? DeepAppInspection { get; set; }

        [Input("defaultNetworkServices")]
        private InputList<Inputs.ListDefaultNetworkServiceArgs>? _defaultNetworkServices;

        /// <summary>
        /// Default network service entries. The structure of `default_network_services` block is documented below.
        /// </summary>
        public InputList<Inputs.ListDefaultNetworkServiceArgs> DefaultNetworkServices
        {
            get => _defaultNetworkServices ?? (_defaultNetworkServices = new InputList<Inputs.ListDefaultNetworkServiceArgs>());
            set => _defaultNetworkServices = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("enforceDefaultAppPort")]
        public Input<string>? EnforceDefaultAppPort { get; set; }

        [Input("entries")]
        private InputList<Inputs.ListEntryArgs>? _entries;

        /// <summary>
        /// Application list entries. The structure of `entries` block is documented below.
        /// </summary>
        public InputList<Inputs.ListEntryArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.ListEntryArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// Enable/disable extended logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("forceInclusionSslDiSigs")]
        public Input<string>? ForceInclusionSslDiSigs { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// List name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Basic application protocol signatures allowed by default.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Action for other applications. Valid values: `pass`, `block`.
        /// </summary>
        [Input("otherApplicationAction")]
        public Input<string>? OtherApplicationAction { get; set; }

        /// <summary>
        /// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("otherApplicationLog")]
        public Input<string>? OtherApplicationLog { get; set; }

        /// <summary>
        /// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
        /// </summary>
        [Input("p2pBlackList")]
        public Input<string>? P2pBlackList { get; set; }

        /// <summary>
        /// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
        /// </summary>
        [Input("p2pBlockList")]
        public Input<string>? P2pBlockList { get; set; }

        /// <summary>
        /// Replacement message group.
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        /// <summary>
        /// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
        /// </summary>
        [Input("unknownApplicationAction")]
        public Input<string>? UnknownApplicationAction { get; set; }

        /// <summary>
        /// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("unknownApplicationLog")]
        public Input<string>? UnknownApplicationLog { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ListArgs()
        {
        }
        public static new ListArgs Empty => new ListArgs();
    }

    public sealed class ListState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("appReplacemsg")]
        public Input<string>? AppReplacemsg { get; set; }

        /// <summary>
        /// comments
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("controlDefaultNetworkServices")]
        public Input<string>? ControlDefaultNetworkServices { get; set; }

        /// <summary>
        /// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("deepAppInspection")]
        public Input<string>? DeepAppInspection { get; set; }

        [Input("defaultNetworkServices")]
        private InputList<Inputs.ListDefaultNetworkServiceGetArgs>? _defaultNetworkServices;

        /// <summary>
        /// Default network service entries. The structure of `default_network_services` block is documented below.
        /// </summary>
        public InputList<Inputs.ListDefaultNetworkServiceGetArgs> DefaultNetworkServices
        {
            get => _defaultNetworkServices ?? (_defaultNetworkServices = new InputList<Inputs.ListDefaultNetworkServiceGetArgs>());
            set => _defaultNetworkServices = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("enforceDefaultAppPort")]
        public Input<string>? EnforceDefaultAppPort { get; set; }

        [Input("entries")]
        private InputList<Inputs.ListEntryGetArgs>? _entries;

        /// <summary>
        /// Application list entries. The structure of `entries` block is documented below.
        /// </summary>
        public InputList<Inputs.ListEntryGetArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.ListEntryGetArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// Enable/disable extended logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("forceInclusionSslDiSigs")]
        public Input<string>? ForceInclusionSslDiSigs { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// List name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Basic application protocol signatures allowed by default.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Action for other applications. Valid values: `pass`, `block`.
        /// </summary>
        [Input("otherApplicationAction")]
        public Input<string>? OtherApplicationAction { get; set; }

        /// <summary>
        /// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("otherApplicationLog")]
        public Input<string>? OtherApplicationLog { get; set; }

        /// <summary>
        /// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
        /// </summary>
        [Input("p2pBlackList")]
        public Input<string>? P2pBlackList { get; set; }

        /// <summary>
        /// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
        /// </summary>
        [Input("p2pBlockList")]
        public Input<string>? P2pBlockList { get; set; }

        /// <summary>
        /// Replacement message group.
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        /// <summary>
        /// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
        /// </summary>
        [Input("unknownApplicationAction")]
        public Input<string>? UnknownApplicationAction { get; set; }

        /// <summary>
        /// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("unknownApplicationLog")]
        public Input<string>? UnknownApplicationLog { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ListState()
        {
        }
        public static new ListState Empty => new ListState();
    }
}
