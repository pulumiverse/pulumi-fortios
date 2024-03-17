// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Autoupdate
{
    /// <summary>
    /// Configure web proxy tunnelling for the FDN.
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
    ///     var trname = new Fortios.System.Autoupdate.Tunneling("trname", new()
    ///     {
    ///         Port = 0,
    ///         Status = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// SystemAutoupdate Tunneling can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/autoupdate/tunneling:Tunneling labelname SystemAutoupdateTunneling
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/autoupdate/tunneling:Tunneling labelname SystemAutoupdateTunneling
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/autoupdate/tunneling:Tunneling")]
    public partial class Tunneling : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Web proxy IP address or FQDN.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Web proxy password.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Web proxy port.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Web proxy username.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Tunneling resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Tunneling(string name, TunnelingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/autoupdate/tunneling:Tunneling", name, args ?? new TunnelingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Tunneling(string name, Input<string> id, TunnelingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/autoupdate/tunneling:Tunneling", name, state, MakeResourceOptions(options, id))
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
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Tunneling resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Tunneling Get(string name, Input<string> id, TunnelingState? state = null, CustomResourceOptions? options = null)
        {
            return new Tunneling(name, id, state, options);
        }
    }

    public sealed class TunnelingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Web proxy IP address or FQDN.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Web proxy password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Web proxy port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Web proxy username.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TunnelingArgs()
        {
        }
        public static new TunnelingArgs Empty => new TunnelingArgs();
    }

    public sealed class TunnelingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Web proxy IP address or FQDN.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Web proxy password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Web proxy port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Web proxy username.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TunnelingState()
        {
        }
        public static new TunnelingState Empty => new TunnelingState();
    }
}
