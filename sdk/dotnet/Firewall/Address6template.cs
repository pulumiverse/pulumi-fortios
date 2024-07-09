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
    /// Configure IPv6 address templates.
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
    ///     var trname = new Fortios.Firewall.Address6template("trname", new()
    ///     {
    ///         Ip6 = "2001:db8:0:b::/64",
    ///         SubnetSegments = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.Address6templateSubnetSegmentArgs
    ///             {
    ///                 Bits = 4,
    ///                 Exclusive = "disable",
    ///                 Id = 1,
    ///                 Name = "country",
    ///             },
    ///             new Fortios.Firewall.Inputs.Address6templateSubnetSegmentArgs
    ///             {
    ///                 Bits = 4,
    ///                 Exclusive = "disable",
    ///                 Id = 2,
    ///                 Name = "state",
    ///             },
    ///         },
    ///         SubnetSegmentCount = 2,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Address6Template can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/address6template:Address6template labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/address6template:Address6template labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/address6template:Address6template")]
    public partial class Address6template : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fabricObject")]
        public Output<string> FabricObject { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// IPv6 address prefix.
        /// </summary>
        [Output("ip6")]
        public Output<string> Ip6 { get; private set; } = null!;

        /// <summary>
        /// IPv6 address template name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of IPv6 subnet segments.
        /// </summary>
        [Output("subnetSegmentCount")]
        public Output<int> SubnetSegmentCount { get; private set; } = null!;

        /// <summary>
        /// IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        /// </summary>
        [Output("subnetSegments")]
        public Output<ImmutableArray<Outputs.Address6templateSubnetSegment>> SubnetSegments { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Address6template resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Address6template(string name, Address6templateArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/address6template:Address6template", name, args ?? new Address6templateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Address6template(string name, Input<string> id, Address6templateState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/address6template:Address6template", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Address6template resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Address6template Get(string name, Input<string> id, Address6templateState? state = null, CustomResourceOptions? options = null)
        {
            return new Address6template(name, id, state, options);
        }
    }

    public sealed class Address6templateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// IPv6 address prefix.
        /// </summary>
        [Input("ip6", required: true)]
        public Input<string> Ip6 { get; set; } = null!;

        /// <summary>
        /// IPv6 address template name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of IPv6 subnet segments.
        /// </summary>
        [Input("subnetSegmentCount", required: true)]
        public Input<int> SubnetSegmentCount { get; set; } = null!;

        [Input("subnetSegments")]
        private InputList<Inputs.Address6templateSubnetSegmentArgs>? _subnetSegments;

        /// <summary>
        /// IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        /// </summary>
        public InputList<Inputs.Address6templateSubnetSegmentArgs> SubnetSegments
        {
            get => _subnetSegments ?? (_subnetSegments = new InputList<Inputs.Address6templateSubnetSegmentArgs>());
            set => _subnetSegments = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Address6templateArgs()
        {
        }
        public static new Address6templateArgs Empty => new Address6templateArgs();
    }

    public sealed class Address6templateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// IPv6 address prefix.
        /// </summary>
        [Input("ip6")]
        public Input<string>? Ip6 { get; set; }

        /// <summary>
        /// IPv6 address template name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of IPv6 subnet segments.
        /// </summary>
        [Input("subnetSegmentCount")]
        public Input<int>? SubnetSegmentCount { get; set; }

        [Input("subnetSegments")]
        private InputList<Inputs.Address6templateSubnetSegmentGetArgs>? _subnetSegments;

        /// <summary>
        /// IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        /// </summary>
        public InputList<Inputs.Address6templateSubnetSegmentGetArgs> SubnetSegments
        {
            get => _subnetSegments ?? (_subnetSegments = new InputList<Inputs.Address6templateSubnetSegmentGetArgs>());
            set => _subnetSegments = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Address6templateState()
        {
        }
        public static new Address6templateState Empty => new Address6templateState();
    }
}
