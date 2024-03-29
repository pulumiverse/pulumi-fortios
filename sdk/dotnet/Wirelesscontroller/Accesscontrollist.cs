// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller
{
    /// <summary>
    /// Configure WiFi bridge access control list. Applies to FortiOS Version `&gt;= 6.4.0`.
    /// 
    /// ## Import
    /// 
    /// WirelessController AccessControlList can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/accesscontrollist:Accesscontrollist labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/accesscontrollist:Accesscontrollist labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontroller/accesscontrollist:Accesscontrollist")]
    public partial class Accesscontrollist : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description.
        /// </summary>
        [Output("comment")]
        public Output<string> Comment { get; private set; } = null!;

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
        /// AP ACL layer3 ipv4 rule list. The structure of `layer3_ipv4_rules` block is documented below.
        /// </summary>
        [Output("layer3Ipv4Rules")]
        public Output<ImmutableArray<Outputs.AccesscontrollistLayer3Ipv4Rule>> Layer3Ipv4Rules { get; private set; } = null!;

        /// <summary>
        /// AP ACL layer3 ipv6 rule list. The structure of `layer3_ipv6_rules` block is documented below.
        /// </summary>
        [Output("layer3Ipv6Rules")]
        public Output<ImmutableArray<Outputs.AccesscontrollistLayer3Ipv6Rule>> Layer3Ipv6Rules { get; private set; } = null!;

        /// <summary>
        /// AP access control list name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// 
        /// The `layer3_ipv4_rules` block supports:
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Accesscontrollist resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Accesscontrollist(string name, AccesscontrollistArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/accesscontrollist:Accesscontrollist", name, args ?? new AccesscontrollistArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Accesscontrollist(string name, Input<string> id, AccesscontrollistState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/accesscontrollist:Accesscontrollist", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Accesscontrollist resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Accesscontrollist Get(string name, Input<string> id, AccesscontrollistState? state = null, CustomResourceOptions? options = null)
        {
            return new Accesscontrollist(name, id, state, options);
        }
    }

    public sealed class AccesscontrollistArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

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

        [Input("layer3Ipv4Rules")]
        private InputList<Inputs.AccesscontrollistLayer3Ipv4RuleArgs>? _layer3Ipv4Rules;

        /// <summary>
        /// AP ACL layer3 ipv4 rule list. The structure of `layer3_ipv4_rules` block is documented below.
        /// </summary>
        public InputList<Inputs.AccesscontrollistLayer3Ipv4RuleArgs> Layer3Ipv4Rules
        {
            get => _layer3Ipv4Rules ?? (_layer3Ipv4Rules = new InputList<Inputs.AccesscontrollistLayer3Ipv4RuleArgs>());
            set => _layer3Ipv4Rules = value;
        }

        [Input("layer3Ipv6Rules")]
        private InputList<Inputs.AccesscontrollistLayer3Ipv6RuleArgs>? _layer3Ipv6Rules;

        /// <summary>
        /// AP ACL layer3 ipv6 rule list. The structure of `layer3_ipv6_rules` block is documented below.
        /// </summary>
        public InputList<Inputs.AccesscontrollistLayer3Ipv6RuleArgs> Layer3Ipv6Rules
        {
            get => _layer3Ipv6Rules ?? (_layer3Ipv6Rules = new InputList<Inputs.AccesscontrollistLayer3Ipv6RuleArgs>());
            set => _layer3Ipv6Rules = value;
        }

        /// <summary>
        /// AP access control list name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// 
        /// The `layer3_ipv4_rules` block supports:
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AccesscontrollistArgs()
        {
        }
        public static new AccesscontrollistArgs Empty => new AccesscontrollistArgs();
    }

    public sealed class AccesscontrollistState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

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

        [Input("layer3Ipv4Rules")]
        private InputList<Inputs.AccesscontrollistLayer3Ipv4RuleGetArgs>? _layer3Ipv4Rules;

        /// <summary>
        /// AP ACL layer3 ipv4 rule list. The structure of `layer3_ipv4_rules` block is documented below.
        /// </summary>
        public InputList<Inputs.AccesscontrollistLayer3Ipv4RuleGetArgs> Layer3Ipv4Rules
        {
            get => _layer3Ipv4Rules ?? (_layer3Ipv4Rules = new InputList<Inputs.AccesscontrollistLayer3Ipv4RuleGetArgs>());
            set => _layer3Ipv4Rules = value;
        }

        [Input("layer3Ipv6Rules")]
        private InputList<Inputs.AccesscontrollistLayer3Ipv6RuleGetArgs>? _layer3Ipv6Rules;

        /// <summary>
        /// AP ACL layer3 ipv6 rule list. The structure of `layer3_ipv6_rules` block is documented below.
        /// </summary>
        public InputList<Inputs.AccesscontrollistLayer3Ipv6RuleGetArgs> Layer3Ipv6Rules
        {
            get => _layer3Ipv6Rules ?? (_layer3Ipv6Rules = new InputList<Inputs.AccesscontrollistLayer3Ipv6RuleGetArgs>());
            set => _layer3Ipv6Rules = value;
        }

        /// <summary>
        /// AP access control list name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// 
        /// The `layer3_ipv4_rules` block supports:
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AccesscontrollistState()
        {
        }
        public static new AccesscontrollistState Empty => new AccesscontrollistState();
    }
}
