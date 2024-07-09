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
    /// Configure NAT64. Applies to FortiOS Version `&lt;= 7.0.0`.
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
    ///     var trname = new Fortios.System.Nat64("trname", new()
    ///     {
    ///         AlwaysSynthesizeAaaaRecord = "enable",
    ///         GenerateIpv6FragmentHeader = "disable",
    ///         Nat46ForceIpv4PacketForwarding = "disable",
    ///         Nat64Prefix = "2001:1:2:3::/96",
    ///         SecondaryPrefixStatus = "disable",
    ///         Status = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System Nat64 can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/nat64:Nat64 labelname SystemNat64
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/nat64:Nat64 labelname SystemNat64
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/nat64:Nat64")]
    public partial class Nat64 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("alwaysSynthesizeAaaaRecord")]
        public Output<string> AlwaysSynthesizeAaaaRecord { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("generateIpv6FragmentHeader")]
        public Output<string> GenerateIpv6FragmentHeader { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("nat46ForceIpv4PacketForwarding")]
        public Output<string> Nat46ForceIpv4PacketForwarding { get; private set; } = null!;

        /// <summary>
        /// NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
        /// </summary>
        [Output("nat64Prefix")]
        public Output<string> Nat64Prefix { get; private set; } = null!;

        /// <summary>
        /// Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("secondaryPrefixStatus")]
        public Output<string> SecondaryPrefixStatus { get; private set; } = null!;

        /// <summary>
        /// Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
        /// </summary>
        [Output("secondaryPrefixes")]
        public Output<ImmutableArray<Outputs.Nat64SecondaryPrefix>> SecondaryPrefixes { get; private set; } = null!;

        /// <summary>
        /// Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Nat64 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Nat64(string name, Nat64Args args, CustomResourceOptions? options = null)
            : base("fortios:system/nat64:Nat64", name, args ?? new Nat64Args(), MakeResourceOptions(options, ""))
        {
        }

        private Nat64(string name, Input<string> id, Nat64State? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/nat64:Nat64", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Nat64 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Nat64 Get(string name, Input<string> id, Nat64State? state = null, CustomResourceOptions? options = null)
        {
            return new Nat64(name, id, state, options);
        }
    }

    public sealed class Nat64Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("alwaysSynthesizeAaaaRecord")]
        public Input<string>? AlwaysSynthesizeAaaaRecord { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("generateIpv6FragmentHeader")]
        public Input<string>? GenerateIpv6FragmentHeader { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nat46ForceIpv4PacketForwarding")]
        public Input<string>? Nat46ForceIpv4PacketForwarding { get; set; }

        /// <summary>
        /// NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
        /// </summary>
        [Input("nat64Prefix", required: true)]
        public Input<string> Nat64Prefix { get; set; } = null!;

        /// <summary>
        /// Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("secondaryPrefixStatus")]
        public Input<string>? SecondaryPrefixStatus { get; set; }

        [Input("secondaryPrefixes")]
        private InputList<Inputs.Nat64SecondaryPrefixArgs>? _secondaryPrefixes;

        /// <summary>
        /// Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
        /// </summary>
        public InputList<Inputs.Nat64SecondaryPrefixArgs> SecondaryPrefixes
        {
            get => _secondaryPrefixes ?? (_secondaryPrefixes = new InputList<Inputs.Nat64SecondaryPrefixArgs>());
            set => _secondaryPrefixes = value;
        }

        /// <summary>
        /// Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Nat64Args()
        {
        }
        public static new Nat64Args Empty => new Nat64Args();
    }

    public sealed class Nat64State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("alwaysSynthesizeAaaaRecord")]
        public Input<string>? AlwaysSynthesizeAaaaRecord { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("generateIpv6FragmentHeader")]
        public Input<string>? GenerateIpv6FragmentHeader { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nat46ForceIpv4PacketForwarding")]
        public Input<string>? Nat46ForceIpv4PacketForwarding { get; set; }

        /// <summary>
        /// NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
        /// </summary>
        [Input("nat64Prefix")]
        public Input<string>? Nat64Prefix { get; set; }

        /// <summary>
        /// Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("secondaryPrefixStatus")]
        public Input<string>? SecondaryPrefixStatus { get; set; }

        [Input("secondaryPrefixes")]
        private InputList<Inputs.Nat64SecondaryPrefixGetArgs>? _secondaryPrefixes;

        /// <summary>
        /// Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
        /// </summary>
        public InputList<Inputs.Nat64SecondaryPrefixGetArgs> SecondaryPrefixes
        {
            get => _secondaryPrefixes ?? (_secondaryPrefixes = new InputList<Inputs.Nat64SecondaryPrefixGetArgs>());
            set => _secondaryPrefixes = value;
        }

        /// <summary>
        /// Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Nat64State()
        {
        }
        public static new Nat64State Empty => new Nat64State();
    }
}
