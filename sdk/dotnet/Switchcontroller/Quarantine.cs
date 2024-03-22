// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller
{
    /// <summary>
    /// Configure FortiSwitch quarantine support.
    /// 
    /// ## Import
    /// 
    /// SwitchController Quarantine can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/quarantine:Quarantine labelname SwitchControllerQuarantine
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/quarantine:Quarantine labelname SwitchControllerQuarantine
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/quarantine:Quarantine")]
    public partial class Quarantine : global::Pulumi.CustomResource
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
        /// Enable/disable quarantine. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("quarantine")]
        public Output<string> Data { get; private set; } = null!;

        /// <summary>
        /// Quarantine MACs. The structure of `targets` block is documented below.
        /// </summary>
        [Output("targets")]
        public Output<ImmutableArray<Outputs.QuarantineTarget>> Targets { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Quarantine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Quarantine(string name, QuarantineArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/quarantine:Quarantine", name, args ?? new QuarantineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Quarantine(string name, Input<string> id, QuarantineState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/quarantine:Quarantine", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Quarantine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Quarantine Get(string name, Input<string> id, QuarantineState? state = null, CustomResourceOptions? options = null)
        {
            return new Quarantine(name, id, state, options);
        }
    }

    public sealed class QuarantineArgs : global::Pulumi.ResourceArgs
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
        /// Enable/disable quarantine. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("quarantine")]
        public Input<string>? Data { get; set; }

        [Input("targets")]
        private InputList<Inputs.QuarantineTargetArgs>? _targets;

        /// <summary>
        /// Quarantine MACs. The structure of `targets` block is documented below.
        /// </summary>
        public InputList<Inputs.QuarantineTargetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.QuarantineTargetArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public QuarantineArgs()
        {
        }
        public static new QuarantineArgs Empty => new QuarantineArgs();
    }

    public sealed class QuarantineState : global::Pulumi.ResourceArgs
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
        /// Enable/disable quarantine. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("quarantine")]
        public Input<string>? Data { get; set; }

        [Input("targets")]
        private InputList<Inputs.QuarantineTargetGetArgs>? _targets;

        /// <summary>
        /// Quarantine MACs. The structure of `targets` block is documented below.
        /// </summary>
        public InputList<Inputs.QuarantineTargetGetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.QuarantineTargetGetArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public QuarantineState()
        {
        }
        public static new QuarantineState Empty => new QuarantineState();
    }
}
