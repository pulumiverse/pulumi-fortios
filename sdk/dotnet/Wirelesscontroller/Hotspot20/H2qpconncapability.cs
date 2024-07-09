// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Hotspot20
{
    /// <summary>
    /// Configure connection capability.
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
    ///     var trname = new Fortios.Wirelesscontroller.Hotspot20.H2qpconncapability("trname", new()
    ///     {
    ///         EspPort = "unknown",
    ///         FtpPort = "unknown",
    ///         HttpPort = "unknown",
    ///         IcmpPort = "closed",
    ///         Ikev2Port = "unknown",
    ///         Ikev2XxPort = "unknown",
    ///         PptpVpnPort = "unknown",
    ///         SshPort = "unknown",
    ///         TlsPort = "unknown",
    ///         VoipTcpPort = "unknown",
    ///         VoipUdpPort = "unknown",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// WirelessControllerHotspot20 H2QpConnCapability can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/hotspot20/h2qpconncapability:H2qpconncapability labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/hotspot20/h2qpconncapability:H2qpconncapability labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontroller/hotspot20/h2qpconncapability:H2qpconncapability")]
    public partial class H2qpconncapability : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Output("espPort")]
        public Output<string> EspPort { get; private set; } = null!;

        /// <summary>
        /// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Output("ftpPort")]
        public Output<string> FtpPort { get; private set; } = null!;

        /// <summary>
        /// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Output("httpPort")]
        public Output<string> HttpPort { get; private set; } = null!;

        /// <summary>
        /// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Output("icmpPort")]
        public Output<string> IcmpPort { get; private set; } = null!;

        /// <summary>
        /// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Output("ikev2Port")]
        public Output<string> Ikev2Port { get; private set; } = null!;

        /// <summary>
        /// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Output("ikev2XxPort")]
        public Output<string> Ikev2XxPort { get; private set; } = null!;

        /// <summary>
        /// Connection capability name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Output("pptpVpnPort")]
        public Output<string> PptpVpnPort { get; private set; } = null!;

        /// <summary>
        /// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Output("sshPort")]
        public Output<string> SshPort { get; private set; } = null!;

        /// <summary>
        /// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Output("tlsPort")]
        public Output<string> TlsPort { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Output("voipTcpPort")]
        public Output<string> VoipTcpPort { get; private set; } = null!;

        /// <summary>
        /// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Output("voipUdpPort")]
        public Output<string> VoipUdpPort { get; private set; } = null!;


        /// <summary>
        /// Create a H2qpconncapability resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public H2qpconncapability(string name, H2qpconncapabilityArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/hotspot20/h2qpconncapability:H2qpconncapability", name, args ?? new H2qpconncapabilityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private H2qpconncapability(string name, Input<string> id, H2qpconncapabilityState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/hotspot20/h2qpconncapability:H2qpconncapability", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing H2qpconncapability resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static H2qpconncapability Get(string name, Input<string> id, H2qpconncapabilityState? state = null, CustomResourceOptions? options = null)
        {
            return new H2qpconncapability(name, id, state, options);
        }
    }

    public sealed class H2qpconncapabilityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("espPort")]
        public Input<string>? EspPort { get; set; }

        /// <summary>
        /// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("ftpPort")]
        public Input<string>? FtpPort { get; set; }

        /// <summary>
        /// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("httpPort")]
        public Input<string>? HttpPort { get; set; }

        /// <summary>
        /// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("icmpPort")]
        public Input<string>? IcmpPort { get; set; }

        /// <summary>
        /// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("ikev2Port")]
        public Input<string>? Ikev2Port { get; set; }

        /// <summary>
        /// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("ikev2XxPort")]
        public Input<string>? Ikev2XxPort { get; set; }

        /// <summary>
        /// Connection capability name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("pptpVpnPort")]
        public Input<string>? PptpVpnPort { get; set; }

        /// <summary>
        /// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("sshPort")]
        public Input<string>? SshPort { get; set; }

        /// <summary>
        /// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("tlsPort")]
        public Input<string>? TlsPort { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("voipTcpPort")]
        public Input<string>? VoipTcpPort { get; set; }

        /// <summary>
        /// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("voipUdpPort")]
        public Input<string>? VoipUdpPort { get; set; }

        public H2qpconncapabilityArgs()
        {
        }
        public static new H2qpconncapabilityArgs Empty => new H2qpconncapabilityArgs();
    }

    public sealed class H2qpconncapabilityState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("espPort")]
        public Input<string>? EspPort { get; set; }

        /// <summary>
        /// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("ftpPort")]
        public Input<string>? FtpPort { get; set; }

        /// <summary>
        /// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("httpPort")]
        public Input<string>? HttpPort { get; set; }

        /// <summary>
        /// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("icmpPort")]
        public Input<string>? IcmpPort { get; set; }

        /// <summary>
        /// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("ikev2Port")]
        public Input<string>? Ikev2Port { get; set; }

        /// <summary>
        /// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("ikev2XxPort")]
        public Input<string>? Ikev2XxPort { get; set; }

        /// <summary>
        /// Connection capability name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("pptpVpnPort")]
        public Input<string>? PptpVpnPort { get; set; }

        /// <summary>
        /// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("sshPort")]
        public Input<string>? SshPort { get; set; }

        /// <summary>
        /// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("tlsPort")]
        public Input<string>? TlsPort { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("voipTcpPort")]
        public Input<string>? VoipTcpPort { get; set; }

        /// <summary>
        /// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
        /// </summary>
        [Input("voipUdpPort")]
        public Input<string>? VoipUdpPort { get; set; }

        public H2qpconncapabilityState()
        {
        }
        public static new H2qpconncapabilityState Empty => new H2qpconncapabilityState();
    }
}
