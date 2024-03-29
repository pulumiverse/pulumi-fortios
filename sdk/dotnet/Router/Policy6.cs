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
    /// Configure IPv6 routing policies.
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
    ///     var trname = new Fortios.Router.Policy6("trname", new()
    ///     {
    ///         Dst = "::/0",
    ///         EndPort = 65535,
    ///         Gateway = "::",
    ///         InputDevice = "port1",
    ///         OutputDevice = "port3",
    ///         Protocol = 33,
    ///         SeqNum = 1,
    ///         Src = "2001:db8:85a3::8a2e:370:7334/128",
    ///         StartPort = 1,
    ///         Status = "enable",
    ///         Tos = "0x00",
    ///         TosMask = "0x00",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Router Policy6 can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/policy6:Policy6 labelname {{seq_num}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/policy6:Policy6 labelname {{seq_num}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:router/policy6:Policy6")]
    public partial class Policy6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action of the policy route. Valid values: `deny`, `permit`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// Destination IPv6 prefix.
        /// </summary>
        [Output("dst")]
        public Output<string> Dst { get; private set; } = null!;

        /// <summary>
        /// Enable/disable negating destination address match. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dstNegate")]
        public Output<string> DstNegate { get; private set; } = null!;

        /// <summary>
        /// Destination address name. The structure of `dstaddr` block is documented below.
        /// </summary>
        [Output("dstaddrs")]
        public Output<ImmutableArray<Outputs.Policy6Dstaddr>> Dstaddrs { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// End destination port number (1 - 65535).
        /// </summary>
        [Output("endPort")]
        public Output<int> EndPort { get; private set; } = null!;

        /// <summary>
        /// End source port number (1 - 65535).
        /// </summary>
        [Output("endSourcePort")]
        public Output<int> EndSourcePort { get; private set; } = null!;

        /// <summary>
        /// IPv6 address of the gateway.
        /// </summary>
        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Incoming interface name.
        /// </summary>
        [Output("inputDevice")]
        public Output<string> InputDevice { get; private set; } = null!;

        /// <summary>
        /// Enable/disable negation of input device match. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("inputDeviceNegate")]
        public Output<string> InputDeviceNegate { get; private set; } = null!;

        /// <summary>
        /// Custom Destination Internet Service name. The structure of `internet_service_custom` block is documented below.
        /// </summary>
        [Output("internetServiceCustoms")]
        public Output<ImmutableArray<Outputs.Policy6InternetServiceCustom>> InternetServiceCustoms { get; private set; } = null!;

        /// <summary>
        /// Destination Internet Service ID. The structure of `internet_service_id` block is documented below.
        /// </summary>
        [Output("internetServiceIds")]
        public Output<ImmutableArray<Outputs.Policy6InternetServiceId>> InternetServiceIds { get; private set; } = null!;

        /// <summary>
        /// Outgoing interface name.
        /// </summary>
        [Output("outputDevice")]
        public Output<string> OutputDevice { get; private set; } = null!;

        /// <summary>
        /// Protocol number (0 - 255).
        /// </summary>
        [Output("protocol")]
        public Output<int> Protocol { get; private set; } = null!;

        /// <summary>
        /// Sequence number.
        /// </summary>
        [Output("seqNum")]
        public Output<int> SeqNum { get; private set; } = null!;

        /// <summary>
        /// Source IPv6 prefix.
        /// </summary>
        [Output("src")]
        public Output<string> Src { get; private set; } = null!;

        /// <summary>
        /// Enable/disable negating source address match. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("srcNegate")]
        public Output<string> SrcNegate { get; private set; } = null!;

        /// <summary>
        /// Source address name. The structure of `srcaddr` block is documented below.
        /// </summary>
        [Output("srcaddrs")]
        public Output<ImmutableArray<Outputs.Policy6Srcaddr>> Srcaddrs { get; private set; } = null!;

        /// <summary>
        /// Start destination port number (1 - 65535).
        /// </summary>
        [Output("startPort")]
        public Output<int> StartPort { get; private set; } = null!;

        /// <summary>
        /// Start source port number (1 - 65535).
        /// </summary>
        [Output("startSourcePort")]
        public Output<int> StartSourcePort { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this policy route. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Type of service bit pattern.
        /// </summary>
        [Output("tos")]
        public Output<string> Tos { get; private set; } = null!;

        /// <summary>
        /// Type of service evaluated bits.
        /// </summary>
        [Output("tosMask")]
        public Output<string> TosMask { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Policy6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy6(string name, Policy6Args args, CustomResourceOptions? options = null)
            : base("fortios:router/policy6:Policy6", name, args ?? new Policy6Args(), MakeResourceOptions(options, ""))
        {
        }

        private Policy6(string name, Input<string> id, Policy6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:router/policy6:Policy6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Policy6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy6 Get(string name, Input<string> id, Policy6State? state = null, CustomResourceOptions? options = null)
        {
            return new Policy6(name, id, state, options);
        }
    }

    public sealed class Policy6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action of the policy route. Valid values: `deny`, `permit`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Destination IPv6 prefix.
        /// </summary>
        [Input("dst")]
        public Input<string>? Dst { get; set; }

        /// <summary>
        /// Enable/disable negating destination address match. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dstNegate")]
        public Input<string>? DstNegate { get; set; }

        [Input("dstaddrs")]
        private InputList<Inputs.Policy6DstaddrArgs>? _dstaddrs;

        /// <summary>
        /// Destination address name. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.Policy6DstaddrArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.Policy6DstaddrArgs>());
            set => _dstaddrs = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// End destination port number (1 - 65535).
        /// </summary>
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        /// <summary>
        /// End source port number (1 - 65535).
        /// </summary>
        [Input("endSourcePort")]
        public Input<int>? EndSourcePort { get; set; }

        /// <summary>
        /// IPv6 address of the gateway.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Incoming interface name.
        /// </summary>
        [Input("inputDevice", required: true)]
        public Input<string> InputDevice { get; set; } = null!;

        /// <summary>
        /// Enable/disable negation of input device match. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("inputDeviceNegate")]
        public Input<string>? InputDeviceNegate { get; set; }

        [Input("internetServiceCustoms")]
        private InputList<Inputs.Policy6InternetServiceCustomArgs>? _internetServiceCustoms;

        /// <summary>
        /// Custom Destination Internet Service name. The structure of `internet_service_custom` block is documented below.
        /// </summary>
        public InputList<Inputs.Policy6InternetServiceCustomArgs> InternetServiceCustoms
        {
            get => _internetServiceCustoms ?? (_internetServiceCustoms = new InputList<Inputs.Policy6InternetServiceCustomArgs>());
            set => _internetServiceCustoms = value;
        }

        [Input("internetServiceIds")]
        private InputList<Inputs.Policy6InternetServiceIdArgs>? _internetServiceIds;

        /// <summary>
        /// Destination Internet Service ID. The structure of `internet_service_id` block is documented below.
        /// </summary>
        public InputList<Inputs.Policy6InternetServiceIdArgs> InternetServiceIds
        {
            get => _internetServiceIds ?? (_internetServiceIds = new InputList<Inputs.Policy6InternetServiceIdArgs>());
            set => _internetServiceIds = value;
        }

        /// <summary>
        /// Outgoing interface name.
        /// </summary>
        [Input("outputDevice")]
        public Input<string>? OutputDevice { get; set; }

        /// <summary>
        /// Protocol number (0 - 255).
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        /// <summary>
        /// Sequence number.
        /// </summary>
        [Input("seqNum")]
        public Input<int>? SeqNum { get; set; }

        /// <summary>
        /// Source IPv6 prefix.
        /// </summary>
        [Input("src")]
        public Input<string>? Src { get; set; }

        /// <summary>
        /// Enable/disable negating source address match. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("srcNegate")]
        public Input<string>? SrcNegate { get; set; }

        [Input("srcaddrs")]
        private InputList<Inputs.Policy6SrcaddrArgs>? _srcaddrs;

        /// <summary>
        /// Source address name. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.Policy6SrcaddrArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.Policy6SrcaddrArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// Start destination port number (1 - 65535).
        /// </summary>
        [Input("startPort")]
        public Input<int>? StartPort { get; set; }

        /// <summary>
        /// Start source port number (1 - 65535).
        /// </summary>
        [Input("startSourcePort")]
        public Input<int>? StartSourcePort { get; set; }

        /// <summary>
        /// Enable/disable this policy route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Type of service bit pattern.
        /// </summary>
        [Input("tos")]
        public Input<string>? Tos { get; set; }

        /// <summary>
        /// Type of service evaluated bits.
        /// </summary>
        [Input("tosMask")]
        public Input<string>? TosMask { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Policy6Args()
        {
        }
        public static new Policy6Args Empty => new Policy6Args();
    }

    public sealed class Policy6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action of the policy route. Valid values: `deny`, `permit`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Destination IPv6 prefix.
        /// </summary>
        [Input("dst")]
        public Input<string>? Dst { get; set; }

        /// <summary>
        /// Enable/disable negating destination address match. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dstNegate")]
        public Input<string>? DstNegate { get; set; }

        [Input("dstaddrs")]
        private InputList<Inputs.Policy6DstaddrGetArgs>? _dstaddrs;

        /// <summary>
        /// Destination address name. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.Policy6DstaddrGetArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.Policy6DstaddrGetArgs>());
            set => _dstaddrs = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// End destination port number (1 - 65535).
        /// </summary>
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        /// <summary>
        /// End source port number (1 - 65535).
        /// </summary>
        [Input("endSourcePort")]
        public Input<int>? EndSourcePort { get; set; }

        /// <summary>
        /// IPv6 address of the gateway.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Incoming interface name.
        /// </summary>
        [Input("inputDevice")]
        public Input<string>? InputDevice { get; set; }

        /// <summary>
        /// Enable/disable negation of input device match. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("inputDeviceNegate")]
        public Input<string>? InputDeviceNegate { get; set; }

        [Input("internetServiceCustoms")]
        private InputList<Inputs.Policy6InternetServiceCustomGetArgs>? _internetServiceCustoms;

        /// <summary>
        /// Custom Destination Internet Service name. The structure of `internet_service_custom` block is documented below.
        /// </summary>
        public InputList<Inputs.Policy6InternetServiceCustomGetArgs> InternetServiceCustoms
        {
            get => _internetServiceCustoms ?? (_internetServiceCustoms = new InputList<Inputs.Policy6InternetServiceCustomGetArgs>());
            set => _internetServiceCustoms = value;
        }

        [Input("internetServiceIds")]
        private InputList<Inputs.Policy6InternetServiceIdGetArgs>? _internetServiceIds;

        /// <summary>
        /// Destination Internet Service ID. The structure of `internet_service_id` block is documented below.
        /// </summary>
        public InputList<Inputs.Policy6InternetServiceIdGetArgs> InternetServiceIds
        {
            get => _internetServiceIds ?? (_internetServiceIds = new InputList<Inputs.Policy6InternetServiceIdGetArgs>());
            set => _internetServiceIds = value;
        }

        /// <summary>
        /// Outgoing interface name.
        /// </summary>
        [Input("outputDevice")]
        public Input<string>? OutputDevice { get; set; }

        /// <summary>
        /// Protocol number (0 - 255).
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        /// <summary>
        /// Sequence number.
        /// </summary>
        [Input("seqNum")]
        public Input<int>? SeqNum { get; set; }

        /// <summary>
        /// Source IPv6 prefix.
        /// </summary>
        [Input("src")]
        public Input<string>? Src { get; set; }

        /// <summary>
        /// Enable/disable negating source address match. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("srcNegate")]
        public Input<string>? SrcNegate { get; set; }

        [Input("srcaddrs")]
        private InputList<Inputs.Policy6SrcaddrGetArgs>? _srcaddrs;

        /// <summary>
        /// Source address name. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.Policy6SrcaddrGetArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.Policy6SrcaddrGetArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// Start destination port number (1 - 65535).
        /// </summary>
        [Input("startPort")]
        public Input<int>? StartPort { get; set; }

        /// <summary>
        /// Start source port number (1 - 65535).
        /// </summary>
        [Input("startSourcePort")]
        public Input<int>? StartSourcePort { get; set; }

        /// <summary>
        /// Enable/disable this policy route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Type of service bit pattern.
        /// </summary>
        [Input("tos")]
        public Input<string>? Tos { get; set; }

        /// <summary>
        /// Type of service evaluated bits.
        /// </summary>
        [Input("tosMask")]
        public Input<string>? TosMask { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Policy6State()
        {
        }
        public static new Policy6State Empty => new Policy6State();
    }
}
