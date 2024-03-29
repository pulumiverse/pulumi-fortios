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
    /// Configure TTL policies.
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
    ///     var trname = new Fortios.Firewall.Ttlpolicy("trname", new()
    ///     {
    ///         Action = "accept",
    ///         Fosid = 1,
    ///         Schedule = "always",
    ///         Services = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.TtlpolicyServiceArgs
    ///             {
    ///                 Name = "ALL",
    ///             },
    ///         },
    ///         Srcaddrs = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.TtlpolicySrcaddrArgs
    ///             {
    ///                 Name = "all",
    ///             },
    ///         },
    ///         Srcintf = "port3",
    ///         Status = "enable",
    ///         Ttl = "23",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Firewall TtlPolicy can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/ttlpolicy:Ttlpolicy labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/ttlpolicy:Ttlpolicy labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/ttlpolicy:Ttlpolicy")]
    public partial class Ttlpolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action to be performed on traffic matching this policy (default = deny). Valid values: `accept`, `deny`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Schedule object from available options.
        /// </summary>
        [Output("schedule")]
        public Output<string> Schedule { get; private set; } = null!;

        /// <summary>
        /// Service object(s) from available options. Separate multiple names with a space. The structure of `service` block is documented below.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<Outputs.TtlpolicyService>> Services { get; private set; } = null!;

        /// <summary>
        /// Source address object(s) from available options. Separate multiple names with a space. The structure of `srcaddr` block is documented below.
        /// </summary>
        [Output("srcaddrs")]
        public Output<ImmutableArray<Outputs.TtlpolicySrcaddr>> Srcaddrs { get; private set; } = null!;

        /// <summary>
        /// Source interface name from available interfaces.
        /// </summary>
        [Output("srcintf")]
        public Output<string> Srcintf { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this TTL policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
        /// </summary>
        [Output("ttl")]
        public Output<string> Ttl { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ttlpolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ttlpolicy(string name, TtlpolicyArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/ttlpolicy:Ttlpolicy", name, args ?? new TtlpolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ttlpolicy(string name, Input<string> id, TtlpolicyState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/ttlpolicy:Ttlpolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ttlpolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ttlpolicy Get(string name, Input<string> id, TtlpolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Ttlpolicy(name, id, state, options);
        }
    }

    public sealed class TtlpolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to be performed on traffic matching this policy (default = deny). Valid values: `accept`, `deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid", required: true)]
        public Input<int> Fosid { get; set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Schedule object from available options.
        /// </summary>
        [Input("schedule", required: true)]
        public Input<string> Schedule { get; set; } = null!;

        [Input("services", required: true)]
        private InputList<Inputs.TtlpolicyServiceArgs>? _services;

        /// <summary>
        /// Service object(s) from available options. Separate multiple names with a space. The structure of `service` block is documented below.
        /// </summary>
        public InputList<Inputs.TtlpolicyServiceArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.TtlpolicyServiceArgs>());
            set => _services = value;
        }

        [Input("srcaddrs", required: true)]
        private InputList<Inputs.TtlpolicySrcaddrArgs>? _srcaddrs;

        /// <summary>
        /// Source address object(s) from available options. Separate multiple names with a space. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.TtlpolicySrcaddrArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.TtlpolicySrcaddrArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// Source interface name from available interfaces.
        /// </summary>
        [Input("srcintf", required: true)]
        public Input<string> Srcintf { get; set; } = null!;

        /// <summary>
        /// Enable/disable this TTL policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
        /// </summary>
        [Input("ttl", required: true)]
        public Input<string> Ttl { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TtlpolicyArgs()
        {
        }
        public static new TtlpolicyArgs Empty => new TtlpolicyArgs();
    }

    public sealed class TtlpolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to be performed on traffic matching this policy (default = deny). Valid values: `accept`, `deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Schedule object from available options.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        [Input("services")]
        private InputList<Inputs.TtlpolicyServiceGetArgs>? _services;

        /// <summary>
        /// Service object(s) from available options. Separate multiple names with a space. The structure of `service` block is documented below.
        /// </summary>
        public InputList<Inputs.TtlpolicyServiceGetArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.TtlpolicyServiceGetArgs>());
            set => _services = value;
        }

        [Input("srcaddrs")]
        private InputList<Inputs.TtlpolicySrcaddrGetArgs>? _srcaddrs;

        /// <summary>
        /// Source address object(s) from available options. Separate multiple names with a space. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.TtlpolicySrcaddrGetArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.TtlpolicySrcaddrGetArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// Source interface name from available interfaces.
        /// </summary>
        [Input("srcintf")]
        public Input<string>? Srcintf { get; set; }

        /// <summary>
        /// Enable/disable this TTL policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TtlpolicyState()
        {
        }
        public static new TtlpolicyState Empty => new TtlpolicyState();
    }
}
