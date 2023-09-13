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
    /// Configure FortiSwitch switch profile.
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
    ///     var trname = new Fortios.Switchcontroller.Switchprofile("trname", new()
    ///     {
    ///         LoginPasswdOverride = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SwitchController SwitchProfile can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:switchcontroller/switchprofile:Switchprofile labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:switchcontroller/switchprofile:Switchprofile labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/switchprofile:Switchprofile")]
    public partial class Switchprofile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable FortiSwitch serial console. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("login")]
        public Output<string> Login { get; private set; } = null!;

        /// <summary>
        /// Login password of managed FortiSwitch.
        /// </summary>
        [Output("loginPasswd")]
        public Output<string?> LoginPasswd { get; private set; } = null!;

        /// <summary>
        /// Enable/disable overriding the admin administrator password for a managed FortiSwitch with the FortiGate admin administrator account password. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("loginPasswdOverride")]
        public Output<string> LoginPasswdOverride { get; private set; } = null!;

        /// <summary>
        /// FortiSwitch Profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable automatic revision backup upon logout from FortiSwitch. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("revisionBackupOnLogout")]
        public Output<string> RevisionBackupOnLogout { get; private set; } = null!;

        /// <summary>
        /// Enable/disable automatic revision backup upon FortiSwitch image upgrade. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("revisionBackupOnUpgrade")]
        public Output<string> RevisionBackupOnUpgrade { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Switchprofile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Switchprofile(string name, SwitchprofileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/switchprofile:Switchprofile", name, args ?? new SwitchprofileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Switchprofile(string name, Input<string> id, SwitchprofileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/switchprofile:Switchprofile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
                AdditionalSecretOutputs =
                {
                    "loginPasswd",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Switchprofile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Switchprofile Get(string name, Input<string> id, SwitchprofileState? state = null, CustomResourceOptions? options = null)
        {
            return new Switchprofile(name, id, state, options);
        }
    }

    public sealed class SwitchprofileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable FortiSwitch serial console. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("login")]
        public Input<string>? Login { get; set; }

        [Input("loginPasswd")]
        private Input<string>? _loginPasswd;

        /// <summary>
        /// Login password of managed FortiSwitch.
        /// </summary>
        public Input<string>? LoginPasswd
        {
            get => _loginPasswd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _loginPasswd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable overriding the admin administrator password for a managed FortiSwitch with the FortiGate admin administrator account password. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("loginPasswdOverride")]
        public Input<string>? LoginPasswdOverride { get; set; }

        /// <summary>
        /// FortiSwitch Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable automatic revision backup upon logout from FortiSwitch. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("revisionBackupOnLogout")]
        public Input<string>? RevisionBackupOnLogout { get; set; }

        /// <summary>
        /// Enable/disable automatic revision backup upon FortiSwitch image upgrade. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("revisionBackupOnUpgrade")]
        public Input<string>? RevisionBackupOnUpgrade { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchprofileArgs()
        {
        }
        public static new SwitchprofileArgs Empty => new SwitchprofileArgs();
    }

    public sealed class SwitchprofileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable FortiSwitch serial console. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("login")]
        public Input<string>? Login { get; set; }

        [Input("loginPasswd")]
        private Input<string>? _loginPasswd;

        /// <summary>
        /// Login password of managed FortiSwitch.
        /// </summary>
        public Input<string>? LoginPasswd
        {
            get => _loginPasswd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _loginPasswd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable overriding the admin administrator password for a managed FortiSwitch with the FortiGate admin administrator account password. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("loginPasswdOverride")]
        public Input<string>? LoginPasswdOverride { get; set; }

        /// <summary>
        /// FortiSwitch Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable automatic revision backup upon logout from FortiSwitch. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("revisionBackupOnLogout")]
        public Input<string>? RevisionBackupOnLogout { get; set; }

        /// <summary>
        /// Enable/disable automatic revision backup upon FortiSwitch image upgrade. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("revisionBackupOnUpgrade")]
        public Input<string>? RevisionBackupOnUpgrade { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchprofileState()
        {
        }
        public static new SwitchprofileState Empty => new SwitchprofileState();
    }
}
