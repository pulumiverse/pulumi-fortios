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
    /// Configure SNMP. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// WirelessController Snmp can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/snmp:Snmp labelname WirelessControllerSnmp
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/snmp:Snmp labelname WirelessControllerSnmp
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontroller/snmp:Snmp")]
    public partial class Snmp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// SNMP Community Configuration. The structure of `community` block is documented below.
        /// </summary>
        [Output("communities")]
        public Output<ImmutableArray<Outputs.SnmpCommunity>> Communities { get; private set; } = null!;

        /// <summary>
        /// Contact Information.
        /// </summary>
        [Output("contactInfo")]
        public Output<string> ContactInfo { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// AC SNMP engineId string (maximum 24 characters).
        /// </summary>
        [Output("engineId")]
        public Output<string> EngineId { get; private set; } = null!;

        /// <summary>
        /// CPU usage when trap is sent.
        /// </summary>
        [Output("trapHighCpuThreshold")]
        public Output<int> TrapHighCpuThreshold { get; private set; } = null!;

        /// <summary>
        /// Memory usage when trap is sent.
        /// </summary>
        [Output("trapHighMemThreshold")]
        public Output<int> TrapHighMemThreshold { get; private set; } = null!;

        /// <summary>
        /// SNMP User Configuration. The structure of `user` block is documented below.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.SnmpUser>> Users { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Snmp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snmp(string name, SnmpArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/snmp:Snmp", name, args ?? new SnmpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Snmp(string name, Input<string> id, SnmpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/snmp:Snmp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Snmp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snmp Get(string name, Input<string> id, SnmpState? state = null, CustomResourceOptions? options = null)
        {
            return new Snmp(name, id, state, options);
        }
    }

    public sealed class SnmpArgs : global::Pulumi.ResourceArgs
    {
        [Input("communities")]
        private InputList<Inputs.SnmpCommunityArgs>? _communities;

        /// <summary>
        /// SNMP Community Configuration. The structure of `community` block is documented below.
        /// </summary>
        public InputList<Inputs.SnmpCommunityArgs> Communities
        {
            get => _communities ?? (_communities = new InputList<Inputs.SnmpCommunityArgs>());
            set => _communities = value;
        }

        /// <summary>
        /// Contact Information.
        /// </summary>
        [Input("contactInfo")]
        public Input<string>? ContactInfo { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// AC SNMP engineId string (maximum 24 characters).
        /// </summary>
        [Input("engineId")]
        public Input<string>? EngineId { get; set; }

        /// <summary>
        /// CPU usage when trap is sent.
        /// </summary>
        [Input("trapHighCpuThreshold")]
        public Input<int>? TrapHighCpuThreshold { get; set; }

        /// <summary>
        /// Memory usage when trap is sent.
        /// </summary>
        [Input("trapHighMemThreshold")]
        public Input<int>? TrapHighMemThreshold { get; set; }

        [Input("users")]
        private InputList<Inputs.SnmpUserArgs>? _users;

        /// <summary>
        /// SNMP User Configuration. The structure of `user` block is documented below.
        /// </summary>
        public InputList<Inputs.SnmpUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.SnmpUserArgs>());
            set => _users = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SnmpArgs()
        {
        }
        public static new SnmpArgs Empty => new SnmpArgs();
    }

    public sealed class SnmpState : global::Pulumi.ResourceArgs
    {
        [Input("communities")]
        private InputList<Inputs.SnmpCommunityGetArgs>? _communities;

        /// <summary>
        /// SNMP Community Configuration. The structure of `community` block is documented below.
        /// </summary>
        public InputList<Inputs.SnmpCommunityGetArgs> Communities
        {
            get => _communities ?? (_communities = new InputList<Inputs.SnmpCommunityGetArgs>());
            set => _communities = value;
        }

        /// <summary>
        /// Contact Information.
        /// </summary>
        [Input("contactInfo")]
        public Input<string>? ContactInfo { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// AC SNMP engineId string (maximum 24 characters).
        /// </summary>
        [Input("engineId")]
        public Input<string>? EngineId { get; set; }

        /// <summary>
        /// CPU usage when trap is sent.
        /// </summary>
        [Input("trapHighCpuThreshold")]
        public Input<int>? TrapHighCpuThreshold { get; set; }

        /// <summary>
        /// Memory usage when trap is sent.
        /// </summary>
        [Input("trapHighMemThreshold")]
        public Input<int>? TrapHighMemThreshold { get; set; }

        [Input("users")]
        private InputList<Inputs.SnmpUserGetArgs>? _users;

        /// <summary>
        /// SNMP User Configuration. The structure of `user` block is documented below.
        /// </summary>
        public InputList<Inputs.SnmpUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.SnmpUserGetArgs>());
            set => _users = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SnmpState()
        {
        }
        public static new SnmpState Empty => new SnmpState();
    }
}
