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
    /// Configure IPv6 multicast NAT policies.
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
    ///     var trname = new Fortios.Firewall.Multicastpolicy6("trname", new()
    ///     {
    ///         Action = "accept",
    ///         Dstaddrs = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.Multicastpolicy6DstaddrArgs
    ///             {
    ///                 Name = "all",
    ///             },
    ///         },
    ///         Dstintf = "port4",
    ///         EndPort = 65535,
    ///         Fosid = 1,
    ///         Logtraffic = "disable",
    ///         Protocol = 0,
    ///         Srcaddrs = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.Multicastpolicy6SrcaddrArgs
    ///             {
    ///                 Name = "all",
    ///             },
    ///         },
    ///         Srcintf = "port3",
    ///         StartPort = 1,
    ///         Status = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Firewall MulticastPolicy6 can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/multicastpolicy6:Multicastpolicy6 labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/multicastpolicy6:Multicastpolicy6 labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/multicastpolicy6:Multicastpolicy6")]
    public partial class Multicastpolicy6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("autoAsicOffload")]
        public Output<string> AutoAsicOffload { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// IPv6 destination address name. The structure of `dstaddr` block is documented below.
        /// </summary>
        [Output("dstaddrs")]
        public Output<ImmutableArray<Outputs.Multicastpolicy6Dstaddr>> Dstaddrs { get; private set; } = null!;

        /// <summary>
        /// IPv6 destination interface name.
        /// </summary>
        [Output("dstintf")]
        public Output<string> Dstintf { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
        /// </summary>
        [Output("endPort")]
        public Output<int> EndPort { get; private set; } = null!;

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Name of an existing IPS sensor.
        /// </summary>
        [Output("ipsSensor")]
        public Output<string> IpsSensor { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging traffic accepted by this policy.
        /// </summary>
        [Output("logtraffic")]
        public Output<string> Logtraffic { get; private set; } = null!;

        /// <summary>
        /// Policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
        /// </summary>
        [Output("protocol")]
        public Output<int> Protocol { get; private set; } = null!;

        /// <summary>
        /// IPv6 source address name. The structure of `srcaddr` block is documented below.
        /// </summary>
        [Output("srcaddrs")]
        public Output<ImmutableArray<Outputs.Multicastpolicy6Srcaddr>> Srcaddrs { get; private set; } = null!;

        /// <summary>
        /// IPv6 source interface name.
        /// </summary>
        [Output("srcintf")]
        public Output<string> Srcintf { get; private set; } = null!;

        /// <summary>
        /// Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
        /// </summary>
        [Output("startPort")]
        public Output<int> StartPort { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("utmStatus")]
        public Output<string> UtmStatus { get; private set; } = null!;

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
        /// Create a Multicastpolicy6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Multicastpolicy6(string name, Multicastpolicy6Args args, CustomResourceOptions? options = null)
            : base("fortios:firewall/multicastpolicy6:Multicastpolicy6", name, args ?? new Multicastpolicy6Args(), MakeResourceOptions(options, ""))
        {
        }

        private Multicastpolicy6(string name, Input<string> id, Multicastpolicy6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/multicastpolicy6:Multicastpolicy6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Multicastpolicy6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Multicastpolicy6 Get(string name, Input<string> id, Multicastpolicy6State? state = null, CustomResourceOptions? options = null)
        {
            return new Multicastpolicy6(name, id, state, options);
        }
    }

    public sealed class Multicastpolicy6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoAsicOffload")]
        public Input<string>? AutoAsicOffload { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("dstaddrs", required: true)]
        private InputList<Inputs.Multicastpolicy6DstaddrArgs>? _dstaddrs;

        /// <summary>
        /// IPv6 destination address name. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.Multicastpolicy6DstaddrArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.Multicastpolicy6DstaddrArgs>());
            set => _dstaddrs = value;
        }

        /// <summary>
        /// IPv6 destination interface name.
        /// </summary>
        [Input("dstintf", required: true)]
        public Input<string> Dstintf { get; set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
        /// </summary>
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Name of an existing IPS sensor.
        /// </summary>
        [Input("ipsSensor")]
        public Input<string>? IpsSensor { get; set; }

        /// <summary>
        /// Enable/disable logging traffic accepted by this policy.
        /// </summary>
        [Input("logtraffic")]
        public Input<string>? Logtraffic { get; set; }

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        [Input("srcaddrs", required: true)]
        private InputList<Inputs.Multicastpolicy6SrcaddrArgs>? _srcaddrs;

        /// <summary>
        /// IPv6 source address name. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.Multicastpolicy6SrcaddrArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.Multicastpolicy6SrcaddrArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// IPv6 source interface name.
        /// </summary>
        [Input("srcintf", required: true)]
        public Input<string> Srcintf { get; set; } = null!;

        /// <summary>
        /// Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
        /// </summary>
        [Input("startPort")]
        public Input<int>? StartPort { get; set; }

        /// <summary>
        /// Enable/disable this policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("utmStatus")]
        public Input<string>? UtmStatus { get; set; }

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

        public Multicastpolicy6Args()
        {
        }
        public static new Multicastpolicy6Args Empty => new Multicastpolicy6Args();
    }

    public sealed class Multicastpolicy6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoAsicOffload")]
        public Input<string>? AutoAsicOffload { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("dstaddrs")]
        private InputList<Inputs.Multicastpolicy6DstaddrGetArgs>? _dstaddrs;

        /// <summary>
        /// IPv6 destination address name. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.Multicastpolicy6DstaddrGetArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.Multicastpolicy6DstaddrGetArgs>());
            set => _dstaddrs = value;
        }

        /// <summary>
        /// IPv6 destination interface name.
        /// </summary>
        [Input("dstintf")]
        public Input<string>? Dstintf { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
        /// </summary>
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Name of an existing IPS sensor.
        /// </summary>
        [Input("ipsSensor")]
        public Input<string>? IpsSensor { get; set; }

        /// <summary>
        /// Enable/disable logging traffic accepted by this policy.
        /// </summary>
        [Input("logtraffic")]
        public Input<string>? Logtraffic { get; set; }

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        [Input("srcaddrs")]
        private InputList<Inputs.Multicastpolicy6SrcaddrGetArgs>? _srcaddrs;

        /// <summary>
        /// IPv6 source address name. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.Multicastpolicy6SrcaddrGetArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.Multicastpolicy6SrcaddrGetArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// IPv6 source interface name.
        /// </summary>
        [Input("srcintf")]
        public Input<string>? Srcintf { get; set; }

        /// <summary>
        /// Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
        /// </summary>
        [Input("startPort")]
        public Input<int>? StartPort { get; set; }

        /// <summary>
        /// Enable/disable this policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("utmStatus")]
        public Input<string>? UtmStatus { get; set; }

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

        public Multicastpolicy6State()
        {
        }
        public static new Multicastpolicy6State Empty => new Multicastpolicy6State();
    }
}
