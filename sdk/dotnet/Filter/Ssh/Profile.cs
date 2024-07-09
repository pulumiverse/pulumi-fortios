// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Ssh
{
    /// <summary>
    /// SSH filter profile.
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
    ///     var trname = new Fortios.Filter.Ssh.Profile("trname", new()
    ///     {
    ///         Block = "x11",
    ///         DefaultCommandLog = "enable",
    ///         Log = "x11",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SshFilter Profile can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:filter/ssh/profile:Profile labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:filter/ssh/profile:Profile labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:filter/ssh/profile:Profile")]
    public partial class Profile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// SSH blocking options.
        /// </summary>
        [Output("block")]
        public Output<string> Block { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("defaultCommandLog")]
        public Output<string> DefaultCommandLog { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// File filter. The structure of `file_filter` block is documented below.
        /// </summary>
        [Output("fileFilter")]
        public Output<Outputs.ProfileFileFilter> FileFilter { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// SSH logging options.
        /// </summary>
        [Output("log")]
        public Output<string> Log { get; private set; } = null!;

        /// <summary>
        /// SSH filter profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// SSH command filter. The structure of `shell_commands` block is documented below.
        /// </summary>
        [Output("shellCommands")]
        public Output<ImmutableArray<Outputs.ProfileShellCommand>> ShellCommands { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Profile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Profile(string name, ProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:filter/ssh/profile:Profile", name, args ?? new ProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Profile(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:filter/ssh/profile:Profile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Profile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Profile Get(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new Profile(name, id, state, options);
        }
    }

    public sealed class ProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SSH blocking options.
        /// </summary>
        [Input("block")]
        public Input<string>? Block { get; set; }

        /// <summary>
        /// Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("defaultCommandLog")]
        public Input<string>? DefaultCommandLog { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// File filter. The structure of `file_filter` block is documented below.
        /// </summary>
        [Input("fileFilter")]
        public Input<Inputs.ProfileFileFilterArgs>? FileFilter { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// SSH logging options.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// SSH filter profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("shellCommands")]
        private InputList<Inputs.ProfileShellCommandArgs>? _shellCommands;

        /// <summary>
        /// SSH command filter. The structure of `shell_commands` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileShellCommandArgs> ShellCommands
        {
            get => _shellCommands ?? (_shellCommands = new InputList<Inputs.ProfileShellCommandArgs>());
            set => _shellCommands = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ProfileArgs()
        {
        }
        public static new ProfileArgs Empty => new ProfileArgs();
    }

    public sealed class ProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SSH blocking options.
        /// </summary>
        [Input("block")]
        public Input<string>? Block { get; set; }

        /// <summary>
        /// Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("defaultCommandLog")]
        public Input<string>? DefaultCommandLog { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// File filter. The structure of `file_filter` block is documented below.
        /// </summary>
        [Input("fileFilter")]
        public Input<Inputs.ProfileFileFilterGetArgs>? FileFilter { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// SSH logging options.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// SSH filter profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("shellCommands")]
        private InputList<Inputs.ProfileShellCommandGetArgs>? _shellCommands;

        /// <summary>
        /// SSH command filter. The structure of `shell_commands` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileShellCommandGetArgs> ShellCommands
        {
            get => _shellCommands ?? (_shellCommands = new InputList<Inputs.ProfileShellCommandGetArgs>());
            set => _shellCommands = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ProfileState()
        {
        }
        public static new ProfileState Empty => new ProfileState();
    }
}
