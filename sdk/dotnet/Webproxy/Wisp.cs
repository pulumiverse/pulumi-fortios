// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Webproxy
{
    /// <summary>
    /// Configure Wireless Internet service provider (WISP) servers.
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
    ///     var trname = new Fortios.Webproxy.Wisp("trname", new()
    ///     {
    ///         MaxConnections = 64,
    ///         OutgoingIp = "0.0.0.0",
    ///         ServerIp = "1.1.1.1",
    ///         ServerPort = 15868,
    ///         Timeout = 5,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// WebProxy Wisp can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:webproxy/wisp:Wisp labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:webproxy/wisp:Wisp labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:webproxy/wisp:Wisp")]
    public partial class Wisp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Maximum number of web proxy WISP connections (4 - 4096, default = 64).
        /// </summary>
        [Output("maxConnections")]
        public Output<int> MaxConnections { get; private set; } = null!;

        /// <summary>
        /// Server name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// WISP outgoing IP address.
        /// </summary>
        [Output("outgoingIp")]
        public Output<string> OutgoingIp { get; private set; } = null!;

        /// <summary>
        /// WISP server IP address.
        /// </summary>
        [Output("serverIp")]
        public Output<string> ServerIp { get; private set; } = null!;

        /// <summary>
        /// WISP server port (1 - 65535, default = 15868).
        /// </summary>
        [Output("serverPort")]
        public Output<int> ServerPort { get; private set; } = null!;

        /// <summary>
        /// Period of time before WISP requests time out (1 - 15 sec, default = 5).
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Wisp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Wisp(string name, WispArgs args, CustomResourceOptions? options = null)
            : base("fortios:webproxy/wisp:Wisp", name, args ?? new WispArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Wisp(string name, Input<string> id, WispState? state = null, CustomResourceOptions? options = null)
            : base("fortios:webproxy/wisp:Wisp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Wisp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Wisp Get(string name, Input<string> id, WispState? state = null, CustomResourceOptions? options = null)
        {
            return new Wisp(name, id, state, options);
        }
    }

    public sealed class WispArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Maximum number of web proxy WISP connections (4 - 4096, default = 64).
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// Server name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// WISP outgoing IP address.
        /// </summary>
        [Input("outgoingIp")]
        public Input<string>? OutgoingIp { get; set; }

        /// <summary>
        /// WISP server IP address.
        /// </summary>
        [Input("serverIp", required: true)]
        public Input<string> ServerIp { get; set; } = null!;

        /// <summary>
        /// WISP server port (1 - 65535, default = 15868).
        /// </summary>
        [Input("serverPort", required: true)]
        public Input<int> ServerPort { get; set; } = null!;

        /// <summary>
        /// Period of time before WISP requests time out (1 - 15 sec, default = 5).
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WispArgs()
        {
        }
        public static new WispArgs Empty => new WispArgs();
    }

    public sealed class WispState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Maximum number of web proxy WISP connections (4 - 4096, default = 64).
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// Server name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// WISP outgoing IP address.
        /// </summary>
        [Input("outgoingIp")]
        public Input<string>? OutgoingIp { get; set; }

        /// <summary>
        /// WISP server IP address.
        /// </summary>
        [Input("serverIp")]
        public Input<string>? ServerIp { get; set; }

        /// <summary>
        /// WISP server port (1 - 65535, default = 15868).
        /// </summary>
        [Input("serverPort")]
        public Input<int>? ServerPort { get; set; }

        /// <summary>
        /// Period of time before WISP requests time out (1 - 15 sec, default = 5).
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WispState()
        {
        }
        public static new WispState Empty => new WispState();
    }
}
