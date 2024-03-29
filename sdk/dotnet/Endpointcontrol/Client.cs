// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Endpointcontrol
{
    /// <summary>
    /// Configure endpoint control client lists. Applies to FortiOS Version `&lt;= 6.2.0`.
    /// 
    /// ## Import
    /// 
    /// EndpointControl Client can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:endpointcontrol/client:Client labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:endpointcontrol/client:Client labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:endpointcontrol/client:Client")]
    public partial class Client : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Endpoint client AD logon groups.
        /// </summary>
        [Output("adGroups")]
        public Output<string?> AdGroups { get; private set; } = null!;

        /// <summary>
        /// Endpoint client ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Endpoint FortiClient UID.
        /// </summary>
        [Output("ftclUid")]
        public Output<string> FtclUid { get; private set; } = null!;

        /// <summary>
        /// Endpoint client information.
        /// </summary>
        [Output("info")]
        public Output<string> Info { get; private set; } = null!;

        /// <summary>
        /// Endpoint client IP address.
        /// </summary>
        [Output("srcIp")]
        public Output<string> SrcIp { get; private set; } = null!;

        /// <summary>
        /// Endpoint client MAC address.
        /// </summary>
        [Output("srcMac")]
        public Output<string> SrcMac { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Client resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Client(string name, ClientArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:endpointcontrol/client:Client", name, args ?? new ClientArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Client(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
            : base("fortios:endpointcontrol/client:Client", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Client resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Client Get(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
        {
            return new Client(name, id, state, options);
        }
    }

    public sealed class ClientArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Endpoint client AD logon groups.
        /// </summary>
        [Input("adGroups")]
        public Input<string>? AdGroups { get; set; }

        /// <summary>
        /// Endpoint client ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Endpoint FortiClient UID.
        /// </summary>
        [Input("ftclUid")]
        public Input<string>? FtclUid { get; set; }

        /// <summary>
        /// Endpoint client information.
        /// </summary>
        [Input("info")]
        public Input<string>? Info { get; set; }

        /// <summary>
        /// Endpoint client IP address.
        /// </summary>
        [Input("srcIp")]
        public Input<string>? SrcIp { get; set; }

        /// <summary>
        /// Endpoint client MAC address.
        /// </summary>
        [Input("srcMac")]
        public Input<string>? SrcMac { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ClientArgs()
        {
        }
        public static new ClientArgs Empty => new ClientArgs();
    }

    public sealed class ClientState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Endpoint client AD logon groups.
        /// </summary>
        [Input("adGroups")]
        public Input<string>? AdGroups { get; set; }

        /// <summary>
        /// Endpoint client ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Endpoint FortiClient UID.
        /// </summary>
        [Input("ftclUid")]
        public Input<string>? FtclUid { get; set; }

        /// <summary>
        /// Endpoint client information.
        /// </summary>
        [Input("info")]
        public Input<string>? Info { get; set; }

        /// <summary>
        /// Endpoint client IP address.
        /// </summary>
        [Input("srcIp")]
        public Input<string>? SrcIp { get; set; }

        /// <summary>
        /// Endpoint client MAC address.
        /// </summary>
        [Input("srcMac")]
        public Input<string>? SrcMac { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ClientState()
        {
        }
        public static new ClientState Empty => new ClientState();
    }
}
