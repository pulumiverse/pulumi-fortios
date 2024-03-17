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
    /// Configure Fortinet Single Sign On (FSSO) server.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.System.Fssopolling("trname", new()
    ///     {
    ///         Authentication = "disable",
    ///         ListeningPort = 8000,
    ///         Status = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// System FssoPolling can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/fssopolling:Fssopolling labelname SystemFssoPolling
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/fssopolling:Fssopolling labelname SystemFssoPolling
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/fssopolling:Fssopolling")]
    public partial class Fssopolling : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Password to connect to FSSO Agent.
        /// </summary>
        [Output("authPassword")]
        public Output<string?> AuthPassword { get; private set; } = null!;

        /// <summary>
        /// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("authentication")]
        public Output<string> Authentication { get; private set; } = null!;

        /// <summary>
        /// Listening port to accept clients (1 - 65535).
        /// </summary>
        [Output("listeningPort")]
        public Output<int> ListeningPort { get; private set; } = null!;

        /// <summary>
        /// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Fssopolling resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fssopolling(string name, FssopollingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/fssopolling:Fssopolling", name, args ?? new FssopollingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fssopolling(string name, Input<string> id, FssopollingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/fssopolling:Fssopolling", name, state, MakeResourceOptions(options, id))
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
                    "authPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Fssopolling resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fssopolling Get(string name, Input<string> id, FssopollingState? state = null, CustomResourceOptions? options = null)
        {
            return new Fssopolling(name, id, state, options);
        }
    }

    public sealed class FssopollingArgs : global::Pulumi.ResourceArgs
    {
        [Input("authPassword")]
        private Input<string>? _authPassword;

        /// <summary>
        /// Password to connect to FSSO Agent.
        /// </summary>
        public Input<string>? AuthPassword
        {
            get => _authPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Listening port to accept clients (1 - 65535).
        /// </summary>
        [Input("listeningPort")]
        public Input<int>? ListeningPort { get; set; }

        /// <summary>
        /// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FssopollingArgs()
        {
        }
        public static new FssopollingArgs Empty => new FssopollingArgs();
    }

    public sealed class FssopollingState : global::Pulumi.ResourceArgs
    {
        [Input("authPassword")]
        private Input<string>? _authPassword;

        /// <summary>
        /// Password to connect to FSSO Agent.
        /// </summary>
        public Input<string>? AuthPassword
        {
            get => _authPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Listening port to accept clients (1 - 65535).
        /// </summary>
        [Input("listeningPort")]
        public Input<int>? ListeningPort { get; set; }

        /// <summary>
        /// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FssopollingState()
        {
        }
        public static new FssopollingState Empty => new FssopollingState();
    }
}
