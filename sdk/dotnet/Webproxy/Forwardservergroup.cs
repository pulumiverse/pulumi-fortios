// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Webproxy
{
    /// <summary>
    /// Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.
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
    ///     var trname1Forwardserver = new Fortios.Webproxy.Forwardserver("trname1Forwardserver", new()
    ///     {
    ///         AddrType = "fqdn",
    ///         Healthcheck = "disable",
    ///         Ip = "0.0.0.0",
    ///         Monitor = "http://www.google.com",
    ///         Port = 1128,
    ///         ServerDownOption = "block",
    ///     });
    /// 
    ///     var trname1Forwardservergroup = new Fortios.Webproxy.Forwardservergroup("trname1Forwardservergroup", new()
    ///     {
    ///         Affinity = "disable",
    ///         GroupDownOption = "block",
    ///         LdbMethod = "weighted",
    ///         ServerLists = new[]
    ///         {
    ///             new Fortios.Webproxy.Inputs.ForwardservergroupServerListArgs
    ///             {
    ///                 Name = trname1Forwardserver.Name,
    ///                 Weight = 12,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// WebProxy ForwardServerGroup can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:webproxy/forwardservergroup:Forwardservergroup labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:webproxy/forwardservergroup:Forwardservergroup labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:webproxy/forwardservergroup:Forwardservergroup")]
    public partial class Forwardservergroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("affinity")]
        public Output<string> Affinity { get; private set; } = null!;

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
        /// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
        /// </summary>
        [Output("groupDownOption")]
        public Output<string> GroupDownOption { get; private set; } = null!;

        /// <summary>
        /// Load balance method: weighted or least-session.
        /// </summary>
        [Output("ldbMethod")]
        public Output<string> LdbMethod { get; private set; } = null!;

        /// <summary>
        /// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        /// </summary>
        [Output("serverLists")]
        public Output<ImmutableArray<Outputs.ForwardservergroupServerList>> ServerLists { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Forwardservergroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Forwardservergroup(string name, ForwardservergroupArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:webproxy/forwardservergroup:Forwardservergroup", name, args ?? new ForwardservergroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Forwardservergroup(string name, Input<string> id, ForwardservergroupState? state = null, CustomResourceOptions? options = null)
            : base("fortios:webproxy/forwardservergroup:Forwardservergroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Forwardservergroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Forwardservergroup Get(string name, Input<string> id, ForwardservergroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Forwardservergroup(name, id, state, options);
        }
    }

    public sealed class ForwardservergroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("affinity")]
        public Input<string>? Affinity { get; set; }

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
        /// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
        /// </summary>
        [Input("groupDownOption")]
        public Input<string>? GroupDownOption { get; set; }

        /// <summary>
        /// Load balance method: weighted or least-session.
        /// </summary>
        [Input("ldbMethod")]
        public Input<string>? LdbMethod { get; set; }

        /// <summary>
        /// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("serverLists")]
        private InputList<Inputs.ForwardservergroupServerListArgs>? _serverLists;

        /// <summary>
        /// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        /// </summary>
        public InputList<Inputs.ForwardservergroupServerListArgs> ServerLists
        {
            get => _serverLists ?? (_serverLists = new InputList<Inputs.ForwardservergroupServerListArgs>());
            set => _serverLists = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ForwardservergroupArgs()
        {
        }
        public static new ForwardservergroupArgs Empty => new ForwardservergroupArgs();
    }

    public sealed class ForwardservergroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("affinity")]
        public Input<string>? Affinity { get; set; }

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
        /// Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
        /// </summary>
        [Input("groupDownOption")]
        public Input<string>? GroupDownOption { get; set; }

        /// <summary>
        /// Load balance method: weighted or least-session.
        /// </summary>
        [Input("ldbMethod")]
        public Input<string>? LdbMethod { get; set; }

        /// <summary>
        /// Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("serverLists")]
        private InputList<Inputs.ForwardservergroupServerListGetArgs>? _serverLists;

        /// <summary>
        /// Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        /// </summary>
        public InputList<Inputs.ForwardservergroupServerListGetArgs> ServerLists
        {
            get => _serverLists ?? (_serverLists = new InputList<Inputs.ForwardservergroupServerListGetArgs>());
            set => _serverLists = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ForwardservergroupState()
        {
        }
        public static new ForwardservergroupState Empty => new ForwardservergroupState();
    }
}
