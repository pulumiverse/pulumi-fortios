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
    /// Configure virtual hardware switch interfaces. Applies to FortiOS Version `7.0.4`.
    /// 
    /// ## Import
    /// 
    /// System VirtualSwitch can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/virtualswitch:Virtualswitch labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/virtualswitch:Virtualswitch labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/virtualswitch:Virtualswitch")]
    public partial class Virtualswitch : global::Pulumi.CustomResource
    {
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
        /// Name of the virtual switch.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Physical switch parent.
        /// </summary>
        [Output("physicalSwitch")]
        public Output<string> PhysicalSwitch { get; private set; } = null!;

        /// <summary>
        /// Configure member ports. The structure of `port` block is documented below.
        /// </summary>
        [Output("ports")]
        public Output<ImmutableArray<Outputs.VirtualswitchPort>> Ports { get; private set; } = null!;

        /// <summary>
        /// Enable/disable SPAN. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("span")]
        public Output<string> Span { get; private set; } = null!;

        /// <summary>
        /// SPAN destination port.
        /// </summary>
        [Output("spanDestPort")]
        public Output<string> SpanDestPort { get; private set; } = null!;

        /// <summary>
        /// SPAN direction. Valid values: `rx`, `tx`, `both`.
        /// </summary>
        [Output("spanDirection")]
        public Output<string> SpanDirection { get; private set; } = null!;

        /// <summary>
        /// SPAN source port.
        /// </summary>
        [Output("spanSourcePort")]
        public Output<string> SpanSourcePort { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// VLAN.
        /// </summary>
        [Output("vlan")]
        public Output<int> Vlan { get; private set; } = null!;


        /// <summary>
        /// Create a Virtualswitch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Virtualswitch(string name, VirtualswitchArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/virtualswitch:Virtualswitch", name, args ?? new VirtualswitchArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Virtualswitch(string name, Input<string> id, VirtualswitchState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/virtualswitch:Virtualswitch", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Virtualswitch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Virtualswitch Get(string name, Input<string> id, VirtualswitchState? state = null, CustomResourceOptions? options = null)
        {
            return new Virtualswitch(name, id, state, options);
        }
    }

    public sealed class VirtualswitchArgs : global::Pulumi.ResourceArgs
    {
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
        /// Name of the virtual switch.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Physical switch parent.
        /// </summary>
        [Input("physicalSwitch")]
        public Input<string>? PhysicalSwitch { get; set; }

        [Input("ports")]
        private InputList<Inputs.VirtualswitchPortArgs>? _ports;

        /// <summary>
        /// Configure member ports. The structure of `port` block is documented below.
        /// </summary>
        public InputList<Inputs.VirtualswitchPortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.VirtualswitchPortArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// Enable/disable SPAN. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("span")]
        public Input<string>? Span { get; set; }

        /// <summary>
        /// SPAN destination port.
        /// </summary>
        [Input("spanDestPort")]
        public Input<string>? SpanDestPort { get; set; }

        /// <summary>
        /// SPAN direction. Valid values: `rx`, `tx`, `both`.
        /// </summary>
        [Input("spanDirection")]
        public Input<string>? SpanDirection { get; set; }

        /// <summary>
        /// SPAN source port.
        /// </summary>
        [Input("spanSourcePort")]
        public Input<string>? SpanSourcePort { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// VLAN.
        /// </summary>
        [Input("vlan")]
        public Input<int>? Vlan { get; set; }

        public VirtualswitchArgs()
        {
        }
        public static new VirtualswitchArgs Empty => new VirtualswitchArgs();
    }

    public sealed class VirtualswitchState : global::Pulumi.ResourceArgs
    {
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
        /// Name of the virtual switch.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Physical switch parent.
        /// </summary>
        [Input("physicalSwitch")]
        public Input<string>? PhysicalSwitch { get; set; }

        [Input("ports")]
        private InputList<Inputs.VirtualswitchPortGetArgs>? _ports;

        /// <summary>
        /// Configure member ports. The structure of `port` block is documented below.
        /// </summary>
        public InputList<Inputs.VirtualswitchPortGetArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.VirtualswitchPortGetArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// Enable/disable SPAN. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("span")]
        public Input<string>? Span { get; set; }

        /// <summary>
        /// SPAN destination port.
        /// </summary>
        [Input("spanDestPort")]
        public Input<string>? SpanDestPort { get; set; }

        /// <summary>
        /// SPAN direction. Valid values: `rx`, `tx`, `both`.
        /// </summary>
        [Input("spanDirection")]
        public Input<string>? SpanDirection { get; set; }

        /// <summary>
        /// SPAN source port.
        /// </summary>
        [Input("spanSourcePort")]
        public Input<string>? SpanSourcePort { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// VLAN.
        /// </summary>
        [Input("vlan")]
        public Input<int>? Vlan { get; set; }

        public VirtualswitchState()
        {
        }
        public static new VirtualswitchState Empty => new VirtualswitchState();
    }
}
