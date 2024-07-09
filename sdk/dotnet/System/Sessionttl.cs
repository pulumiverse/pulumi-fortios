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
    ///     var trname = new Fortios.System.Sessionttl("trname", new()
    ///     {
    ///         Default = "3600",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System SessionTtl can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/sessionttl:Sessionttl labelname SystemSessionTtl
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/sessionttl:Sessionttl labelname SystemSessionTtl
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/sessionttl:Sessionttl")]
    public partial class Sessionttl : global::Pulumi.CustomResource
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
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Session TTL port. The structure of `port` block is documented below.
        /// </summary>
        [Output("ports")]
        public Output<ImmutableArray<Outputs.SessionttlPort>> Ports { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Sessionttl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Sessionttl(string name, SessionttlArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/sessionttl:Sessionttl", name, args ?? new SessionttlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Sessionttl(string name, Input<string> id, SessionttlState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/sessionttl:Sessionttl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Sessionttl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Sessionttl Get(string name, Input<string> id, SessionttlState? state = null, CustomResourceOptions? options = null)
        {
            return new Sessionttl(name, id, state, options);
        }
    }

    public sealed class SessionttlArgs : global::Pulumi.ResourceArgs
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

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        [Input("ports")]
        private InputList<Inputs.SessionttlPortArgs>? _ports;

        /// <summary>
        /// Session TTL port. The structure of `port` block is documented below.
        /// </summary>
        public InputList<Inputs.SessionttlPortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.SessionttlPortArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SessionttlArgs()
        {
        }
        public static new SessionttlArgs Empty => new SessionttlArgs();
    }

    public sealed class SessionttlState : global::Pulumi.ResourceArgs
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

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        [Input("ports")]
        private InputList<Inputs.SessionttlPortGetArgs>? _ports;

        /// <summary>
        /// Session TTL port. The structure of `port` block is documented below.
        /// </summary>
        public InputList<Inputs.SessionttlPortGetArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.SessionttlPortGetArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SessionttlState()
        {
        }
        public static new SessionttlState Empty => new SessionttlState();
    }
}
