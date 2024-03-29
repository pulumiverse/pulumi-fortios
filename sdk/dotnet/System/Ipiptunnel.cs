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
    /// Configure IP in IP Tunneling.
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
    ///     var trname = new Fortios.System.Ipiptunnel("trname", new()
    ///     {
    ///         Interface = "port3",
    ///         LocalGw = "1.1.1.1",
    ///         RemoteGw = "2.2.2.2",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// System IpipTunnel can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/ipiptunnel:Ipiptunnel labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/ipiptunnel:Ipiptunnel labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/ipiptunnel:Ipiptunnel")]
    public partial class Ipiptunnel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("autoAsicOffload")]
        public Output<string> AutoAsicOffload { get; private set; } = null!;

        /// <summary>
        /// Interface name that is associated with the incoming traffic from available options.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// IPv4 address for the local gateway.
        /// </summary>
        [Output("localGw")]
        public Output<string> LocalGw { get; private set; } = null!;

        /// <summary>
        /// IPIP Tunnel name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// IPv4 address for the remote gateway.
        /// </summary>
        [Output("remoteGw")]
        public Output<string> RemoteGw { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("useSdwan")]
        public Output<string> UseSdwan { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ipiptunnel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipiptunnel(string name, IpiptunnelArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/ipiptunnel:Ipiptunnel", name, args ?? new IpiptunnelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipiptunnel(string name, Input<string> id, IpiptunnelState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/ipiptunnel:Ipiptunnel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipiptunnel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipiptunnel Get(string name, Input<string> id, IpiptunnelState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipiptunnel(name, id, state, options);
        }
    }

    public sealed class IpiptunnelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoAsicOffload")]
        public Input<string>? AutoAsicOffload { get; set; }

        /// <summary>
        /// Interface name that is associated with the incoming traffic from available options.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// IPv4 address for the local gateway.
        /// </summary>
        [Input("localGw", required: true)]
        public Input<string> LocalGw { get; set; } = null!;

        /// <summary>
        /// IPIP Tunnel name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// IPv4 address for the remote gateway.
        /// </summary>
        [Input("remoteGw", required: true)]
        public Input<string> RemoteGw { get; set; } = null!;

        /// <summary>
        /// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("useSdwan")]
        public Input<string>? UseSdwan { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IpiptunnelArgs()
        {
        }
        public static new IpiptunnelArgs Empty => new IpiptunnelArgs();
    }

    public sealed class IpiptunnelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoAsicOffload")]
        public Input<string>? AutoAsicOffload { get; set; }

        /// <summary>
        /// Interface name that is associated with the incoming traffic from available options.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// IPv4 address for the local gateway.
        /// </summary>
        [Input("localGw")]
        public Input<string>? LocalGw { get; set; }

        /// <summary>
        /// IPIP Tunnel name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// IPv4 address for the remote gateway.
        /// </summary>
        [Input("remoteGw")]
        public Input<string>? RemoteGw { get; set; }

        /// <summary>
        /// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("useSdwan")]
        public Input<string>? UseSdwan { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IpiptunnelState()
        {
        }
        public static new IpiptunnelState Empty => new IpiptunnelState();
    }
}
