// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    /// <summary>
    /// Configure user defined IPv4 local-in policies.
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
    ///     var trname = new Fortios.Firewall.Localinpolicy("trname", new()
    ///     {
    ///         Action = "accept",
    ///         Dstaddrs = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.LocalinpolicyDstaddrArgs
    ///             {
    ///                 Name = "all",
    ///             },
    ///         },
    ///         HaMgmtIntfOnly = "disable",
    ///         Intf = "port4",
    ///         Policyid = 1,
    ///         Schedule = "always",
    ///         Services = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.LocalinpolicyServiceArgs
    ///             {
    ///                 Name = "ALL",
    ///             },
    ///         },
    ///         Srcaddrs = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.LocalinpolicySrcaddrArgs
    ///             {
    ///                 Name = "all",
    ///             },
    ///         },
    ///         Status = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall LocalInPolicy can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/localinpolicy:Localinpolicy labelname {{policyid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/localinpolicy:Localinpolicy labelname {{policyid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/localinpolicy:Localinpolicy")]
    public partial class Localinpolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action performed on traffic matching the policy (default = deny). Valid values: `accept`, `deny`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dstaddrNegate")]
        public Output<string> DstaddrNegate { get; private set; } = null!;

        /// <summary>
        /// Destination address object from available options. The structure of `dstaddr` block is documented below.
        /// </summary>
        [Output("dstaddrs")]
        public Output<ImmutableArray<Outputs.LocalinpolicyDstaddr>> Dstaddrs { get; private set; } = null!;

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
        /// Enable/disable dedicating the HA management interface only for local-in policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("haMgmtIntfOnly")]
        public Output<string> HaMgmtIntfOnly { get; private set; } = null!;

        /// <summary>
        /// Incoming interface name from available options.
        /// </summary>
        [Output("intf")]
        public Output<string> Intf { get; private set; } = null!;

        /// <summary>
        /// User defined local in policy ID.
        /// </summary>
        [Output("policyid")]
        public Output<int> Policyid { get; private set; } = null!;

        /// <summary>
        /// Schedule object from available options.
        /// </summary>
        [Output("schedule")]
        public Output<string> Schedule { get; private set; } = null!;

        /// <summary>
        /// When enabled service specifies what the service must NOT be. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("serviceNegate")]
        public Output<string> ServiceNegate { get; private set; } = null!;

        /// <summary>
        /// Service object from available options. The structure of `service` block is documented below.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<Outputs.LocalinpolicyService>> Services { get; private set; } = null!;

        /// <summary>
        /// When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("srcaddrNegate")]
        public Output<string> SrcaddrNegate { get; private set; } = null!;

        /// <summary>
        /// Source address object from available options. The structure of `srcaddr` block is documented below.
        /// </summary>
        [Output("srcaddrs")]
        public Output<ImmutableArray<Outputs.LocalinpolicySrcaddr>> Srcaddrs { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this local-in policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable virtual patching. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("virtualPatch")]
        public Output<string> VirtualPatch { get; private set; } = null!;


        /// <summary>
        /// Create a Localinpolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Localinpolicy(string name, LocalinpolicyArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/localinpolicy:Localinpolicy", name, args ?? new LocalinpolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Localinpolicy(string name, Input<string> id, LocalinpolicyState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/localinpolicy:Localinpolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Localinpolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Localinpolicy Get(string name, Input<string> id, LocalinpolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Localinpolicy(name, id, state, options);
        }
    }

    public sealed class LocalinpolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action performed on traffic matching the policy (default = deny). Valid values: `accept`, `deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dstaddrNegate")]
        public Input<string>? DstaddrNegate { get; set; }

        [Input("dstaddrs", required: true)]
        private InputList<Inputs.LocalinpolicyDstaddrArgs>? _dstaddrs;

        /// <summary>
        /// Destination address object from available options. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.LocalinpolicyDstaddrArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.LocalinpolicyDstaddrArgs>());
            set => _dstaddrs = value;
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
        /// Enable/disable dedicating the HA management interface only for local-in policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("haMgmtIntfOnly")]
        public Input<string>? HaMgmtIntfOnly { get; set; }

        /// <summary>
        /// Incoming interface name from available options.
        /// </summary>
        [Input("intf")]
        public Input<string>? Intf { get; set; }

        /// <summary>
        /// User defined local in policy ID.
        /// </summary>
        [Input("policyid")]
        public Input<int>? Policyid { get; set; }

        /// <summary>
        /// Schedule object from available options.
        /// </summary>
        [Input("schedule", required: true)]
        public Input<string> Schedule { get; set; } = null!;

        /// <summary>
        /// When enabled service specifies what the service must NOT be. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("serviceNegate")]
        public Input<string>? ServiceNegate { get; set; }

        [Input("services")]
        private InputList<Inputs.LocalinpolicyServiceArgs>? _services;

        /// <summary>
        /// Service object from available options. The structure of `service` block is documented below.
        /// </summary>
        public InputList<Inputs.LocalinpolicyServiceArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.LocalinpolicyServiceArgs>());
            set => _services = value;
        }

        /// <summary>
        /// When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("srcaddrNegate")]
        public Input<string>? SrcaddrNegate { get; set; }

        [Input("srcaddrs", required: true)]
        private InputList<Inputs.LocalinpolicySrcaddrArgs>? _srcaddrs;

        /// <summary>
        /// Source address object from available options. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.LocalinpolicySrcaddrArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.LocalinpolicySrcaddrArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// Enable/disable this local-in policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable virtual patching. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("virtualPatch")]
        public Input<string>? VirtualPatch { get; set; }

        public LocalinpolicyArgs()
        {
        }
        public static new LocalinpolicyArgs Empty => new LocalinpolicyArgs();
    }

    public sealed class LocalinpolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action performed on traffic matching the policy (default = deny). Valid values: `accept`, `deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dstaddrNegate")]
        public Input<string>? DstaddrNegate { get; set; }

        [Input("dstaddrs")]
        private InputList<Inputs.LocalinpolicyDstaddrGetArgs>? _dstaddrs;

        /// <summary>
        /// Destination address object from available options. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.LocalinpolicyDstaddrGetArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.LocalinpolicyDstaddrGetArgs>());
            set => _dstaddrs = value;
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
        /// Enable/disable dedicating the HA management interface only for local-in policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("haMgmtIntfOnly")]
        public Input<string>? HaMgmtIntfOnly { get; set; }

        /// <summary>
        /// Incoming interface name from available options.
        /// </summary>
        [Input("intf")]
        public Input<string>? Intf { get; set; }

        /// <summary>
        /// User defined local in policy ID.
        /// </summary>
        [Input("policyid")]
        public Input<int>? Policyid { get; set; }

        /// <summary>
        /// Schedule object from available options.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// When enabled service specifies what the service must NOT be. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("serviceNegate")]
        public Input<string>? ServiceNegate { get; set; }

        [Input("services")]
        private InputList<Inputs.LocalinpolicyServiceGetArgs>? _services;

        /// <summary>
        /// Service object from available options. The structure of `service` block is documented below.
        /// </summary>
        public InputList<Inputs.LocalinpolicyServiceGetArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.LocalinpolicyServiceGetArgs>());
            set => _services = value;
        }

        /// <summary>
        /// When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("srcaddrNegate")]
        public Input<string>? SrcaddrNegate { get; set; }

        [Input("srcaddrs")]
        private InputList<Inputs.LocalinpolicySrcaddrGetArgs>? _srcaddrs;

        /// <summary>
        /// Source address object from available options. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.LocalinpolicySrcaddrGetArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.LocalinpolicySrcaddrGetArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// Enable/disable this local-in policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable virtual patching. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("virtualPatch")]
        public Input<string>? VirtualPatch { get; set; }

        public LocalinpolicyState()
        {
        }
        public static new LocalinpolicyState Empty => new LocalinpolicyState();
    }
}
