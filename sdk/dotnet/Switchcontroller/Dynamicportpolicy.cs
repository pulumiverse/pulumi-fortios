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
    /// Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device. Applies to FortiOS Version `&gt;= 7.0.1`.
    /// 
    /// ## Import
    /// 
    /// SwitchController DynamicPortPolicy can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:switchcontroller/dynamicportpolicy:Dynamicportpolicy labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:switchcontroller/dynamicportpolicy:Dynamicportpolicy labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/dynamicportpolicy:Dynamicportpolicy")]
    public partial class Dynamicportpolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description for the Dynamic port policy.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// FortiLink interface for which this Dynamic port policy belongs to.
        /// </summary>
        [Output("fortilink")]
        public Output<string> Fortilink { get; private set; } = null!;

        /// <summary>
        /// Dynamic port policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Port policies with matching criteria and actions. The structure of `policy` block is documented below.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<Outputs.DynamicportpolicyPolicy>> Policies { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Dynamicportpolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dynamicportpolicy(string name, DynamicportpolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/dynamicportpolicy:Dynamicportpolicy", name, args ?? new DynamicportpolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dynamicportpolicy(string name, Input<string> id, DynamicportpolicyState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/dynamicportpolicy:Dynamicportpolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Dynamicportpolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dynamicportpolicy Get(string name, Input<string> id, DynamicportpolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Dynamicportpolicy(name, id, state, options);
        }
    }

    public sealed class DynamicportpolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the Dynamic port policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// FortiLink interface for which this Dynamic port policy belongs to.
        /// </summary>
        [Input("fortilink")]
        public Input<string>? Fortilink { get; set; }

        /// <summary>
        /// Dynamic port policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policies")]
        private InputList<Inputs.DynamicportpolicyPolicyArgs>? _policies;

        /// <summary>
        /// Port policies with matching criteria and actions. The structure of `policy` block is documented below.
        /// </summary>
        public InputList<Inputs.DynamicportpolicyPolicyArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.DynamicportpolicyPolicyArgs>());
            set => _policies = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DynamicportpolicyArgs()
        {
        }
        public static new DynamicportpolicyArgs Empty => new DynamicportpolicyArgs();
    }

    public sealed class DynamicportpolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the Dynamic port policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// FortiLink interface for which this Dynamic port policy belongs to.
        /// </summary>
        [Input("fortilink")]
        public Input<string>? Fortilink { get; set; }

        /// <summary>
        /// Dynamic port policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policies")]
        private InputList<Inputs.DynamicportpolicyPolicyGetArgs>? _policies;

        /// <summary>
        /// Port policies with matching criteria and actions. The structure of `policy` block is documented below.
        /// </summary>
        public InputList<Inputs.DynamicportpolicyPolicyGetArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.DynamicportpolicyPolicyGetArgs>());
            set => _policies = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DynamicportpolicyState()
        {
        }
        public static new DynamicportpolicyState Empty => new DynamicportpolicyState();
    }
}
