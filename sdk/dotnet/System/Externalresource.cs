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
    /// Configure external resource.
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
    ///     var trname = new Fortios.System.Externalresource("trname", new()
    ///     {
    ///         Category = 199,
    ///         RefreshRate = 5,
    ///         Resource = "https://tmpxxxxx.com",
    ///         Status = "enable",
    ///         Type = "category",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// System ExternalResource can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/externalresource:Externalresource labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/externalresource:Externalresource labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/externalresource:Externalresource")]
    public partial class Externalresource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// User resource category.
        /// </summary>
        [Output("category")]
        public Output<int> Category { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

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
        /// External resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// HTTP basic authentication password.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Time interval to refresh external resource (1 - 43200 min, default = 5 min).
        /// </summary>
        [Output("refreshRate")]
        public Output<int> RefreshRate { get; private set; } = null!;

        /// <summary>
        /// URI of external resource.
        /// </summary>
        [Output("resource")]
        public Output<string> Resource { get; private set; } = null!;

        /// <summary>
        /// Certificate verification option. Valid values: `none`, `basic`, `full`.
        /// </summary>
        [Output("serverIdentityCheck")]
        public Output<string> ServerIdentityCheck { get; private set; } = null!;

        /// <summary>
        /// Source IPv4 address used to communicate with server.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable user resource. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// User resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// External resource update method. Valid values: `feed`, `push`.
        /// </summary>
        [Output("updateMethod")]
        public Output<string> UpdateMethod { get; private set; } = null!;

        /// <summary>
        /// Override HTTP User-Agent header used when retrieving this external resource.
        /// </summary>
        [Output("userAgent")]
        public Output<string> UserAgent { get; private set; } = null!;

        /// <summary>
        /// HTTP basic authentication user name.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Externalresource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Externalresource(string name, ExternalresourceArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/externalresource:Externalresource", name, args ?? new ExternalresourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Externalresource(string name, Input<string> id, ExternalresourceState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/externalresource:Externalresource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Externalresource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Externalresource Get(string name, Input<string> id, ExternalresourceState? state = null, CustomResourceOptions? options = null)
        {
            return new Externalresource(name, id, state, options);
        }
    }

    public sealed class ExternalresourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User resource category.
        /// </summary>
        [Input("category")]
        public Input<int>? Category { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

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
        /// External resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// HTTP basic authentication password.
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
        /// Time interval to refresh external resource (1 - 43200 min, default = 5 min).
        /// </summary>
        [Input("refreshRate", required: true)]
        public Input<int> RefreshRate { get; set; } = null!;

        /// <summary>
        /// URI of external resource.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// Certificate verification option. Valid values: `none`, `basic`, `full`.
        /// </summary>
        [Input("serverIdentityCheck")]
        public Input<string>? ServerIdentityCheck { get; set; }

        /// <summary>
        /// Source IPv4 address used to communicate with server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Enable/disable user resource. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// User resource type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// External resource update method. Valid values: `feed`, `push`.
        /// </summary>
        [Input("updateMethod")]
        public Input<string>? UpdateMethod { get; set; }

        /// <summary>
        /// Override HTTP User-Agent header used when retrieving this external resource.
        /// </summary>
        [Input("userAgent")]
        public Input<string>? UserAgent { get; set; }

        /// <summary>
        /// HTTP basic authentication user name.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ExternalresourceArgs()
        {
        }
        public static new ExternalresourceArgs Empty => new ExternalresourceArgs();
    }

    public sealed class ExternalresourceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User resource category.
        /// </summary>
        [Input("category")]
        public Input<int>? Category { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

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
        /// External resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// HTTP basic authentication password.
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
        /// Time interval to refresh external resource (1 - 43200 min, default = 5 min).
        /// </summary>
        [Input("refreshRate")]
        public Input<int>? RefreshRate { get; set; }

        /// <summary>
        /// URI of external resource.
        /// </summary>
        [Input("resource")]
        public Input<string>? Resource { get; set; }

        /// <summary>
        /// Certificate verification option. Valid values: `none`, `basic`, `full`.
        /// </summary>
        [Input("serverIdentityCheck")]
        public Input<string>? ServerIdentityCheck { get; set; }

        /// <summary>
        /// Source IPv4 address used to communicate with server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Enable/disable user resource. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// User resource type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// External resource update method. Valid values: `feed`, `push`.
        /// </summary>
        [Input("updateMethod")]
        public Input<string>? UpdateMethod { get; set; }

        /// <summary>
        /// Override HTTP User-Agent header used when retrieving this external resource.
        /// </summary>
        [Input("userAgent")]
        public Input<string>? UserAgent { get; set; }

        /// <summary>
        /// HTTP basic authentication user name.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ExternalresourceState()
        {
        }
        public static new ExternalresourceState Empty => new ExternalresourceState();
    }
}
