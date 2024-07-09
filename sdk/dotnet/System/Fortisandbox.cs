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
    /// Configure FortiSandbox.
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
    ///     var trname = new Fortios.System.Fortisandbox("trname", new()
    ///     {
    ///         EncAlgorithm = "default",
    ///         SslMinProtoVersion = "default",
    ///         Status = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System Fortisandbox can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/fortisandbox:Fortisandbox labelname SystemFortisandbox
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/fortisandbox:Fortisandbox labelname SystemFortisandbox
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/fortisandbox:Fortisandbox")]
    public partial class Fortisandbox : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Notifier email address.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
        /// </summary>
        [Output("encAlgorithm")]
        public Output<string> EncAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("forticloud")]
        public Output<string> Forticloud { get; private set; } = null!;

        /// <summary>
        /// Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("inlineScan")]
        public Output<string> InlineScan { get; private set; } = null!;

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Output("interfaceSelectMethod")]
        public Output<string> InterfaceSelectMethod { get; private set; } = null!;

        /// <summary>
        /// IPv4 or IPv6 address of the remote FortiSandbox.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Source IP address for communications to FortiSandbox.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        [Output("sslMinProtoVersion")]
        public Output<string> SslMinProtoVersion { get; private set; } = null!;

        /// <summary>
        /// Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Fortisandbox resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fortisandbox(string name, FortisandboxArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/fortisandbox:Fortisandbox", name, args ?? new FortisandboxArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fortisandbox(string name, Input<string> id, FortisandboxState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/fortisandbox:Fortisandbox", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Fortisandbox resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fortisandbox Get(string name, Input<string> id, FortisandboxState? state = null, CustomResourceOptions? options = null)
        {
            return new Fortisandbox(name, id, state, options);
        }
    }

    public sealed class FortisandboxArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Notifier email address.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
        /// </summary>
        [Input("encAlgorithm")]
        public Input<string>? EncAlgorithm { get; set; }

        /// <summary>
        /// Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticloud")]
        public Input<string>? Forticloud { get; set; }

        /// <summary>
        /// Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("inlineScan")]
        public Input<string>? InlineScan { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// IPv4 or IPv6 address of the remote FortiSandbox.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Source IP address for communications to FortiSandbox.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        [Input("sslMinProtoVersion")]
        public Input<string>? SslMinProtoVersion { get; set; }

        /// <summary>
        /// Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FortisandboxArgs()
        {
        }
        public static new FortisandboxArgs Empty => new FortisandboxArgs();
    }

    public sealed class FortisandboxState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Notifier email address.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
        /// </summary>
        [Input("encAlgorithm")]
        public Input<string>? EncAlgorithm { get; set; }

        /// <summary>
        /// Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticloud")]
        public Input<string>? Forticloud { get; set; }

        /// <summary>
        /// Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("inlineScan")]
        public Input<string>? InlineScan { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// IPv4 or IPv6 address of the remote FortiSandbox.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Source IP address for communications to FortiSandbox.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        [Input("sslMinProtoVersion")]
        public Input<string>? SslMinProtoVersion { get; set; }

        /// <summary>
        /// Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FortisandboxState()
        {
        }
        public static new FortisandboxState Empty => new FortisandboxState();
    }
}
