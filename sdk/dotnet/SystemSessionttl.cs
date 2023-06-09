// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Configure global session TTL timers for this FortiGate.
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
    ///     var trname = new Fortios.SystemSessionttl("trname", new()
    ///     {
    ///         Default = "3600",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System SessionTtl can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemSessionttl:SystemSessionttl labelname SystemSessionTtl
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemSessionttl:SystemSessionttl labelname SystemSessionTtl
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemSessionttl:SystemSessionttl")]
    public partial class SystemSessionttl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Default timeout.
        /// </summary>
        [Output("default")]
        public Output<string> Default { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Session TTL port. The structure of `port` block is documented below.
        /// </summary>
        [Output("ports")]
        public Output<ImmutableArray<Outputs.SystemSessionttlPort>> Ports { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemSessionttl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemSessionttl(string name, SystemSessionttlArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemSessionttl:SystemSessionttl", name, args ?? new SystemSessionttlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemSessionttl(string name, Input<string> id, SystemSessionttlState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemSessionttl:SystemSessionttl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemSessionttl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemSessionttl Get(string name, Input<string> id, SystemSessionttlState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemSessionttl(name, id, state, options);
        }
    }

    public sealed class SystemSessionttlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default timeout.
        /// </summary>
        [Input("default")]
        public Input<string>? Default { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("ports")]
        private InputList<Inputs.SystemSessionttlPortArgs>? _ports;

        /// <summary>
        /// Session TTL port. The structure of `port` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSessionttlPortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.SystemSessionttlPortArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemSessionttlArgs()
        {
        }
        public static new SystemSessionttlArgs Empty => new SystemSessionttlArgs();
    }

    public sealed class SystemSessionttlState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default timeout.
        /// </summary>
        [Input("default")]
        public Input<string>? Default { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("ports")]
        private InputList<Inputs.SystemSessionttlPortGetArgs>? _ports;

        /// <summary>
        /// Session TTL port. The structure of `port` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSessionttlPortGetArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.SystemSessionttlPortGetArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemSessionttlState()
        {
        }
        public static new SystemSessionttlState Empty => new SystemSessionttlState();
    }
}
