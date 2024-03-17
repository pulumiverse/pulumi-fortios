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
    /// Configure GRE tunnel.
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
    ///     var trname = new Fortios.System.Gretunnel("trname", new()
    ///     {
    ///         ChecksumReception = "disable",
    ///         ChecksumTransmission = "disable",
    ///         DscpCopying = "disable",
    ///         Interface = "port3",
    ///         IpVersion = "4",
    ///         KeepaliveFailtimes = 10,
    ///         KeepaliveInterval = 0,
    ///         KeyInbound = 0,
    ///         KeyOutbound = 0,
    ///         LocalGw = "3.3.3.3",
    ///         LocalGw6 = "::",
    ///         RemoteGw = "1.1.1.1",
    ///         RemoteGw6 = "::",
    ///         SequenceNumberReception = "disable",
    ///         SequenceNumberTransmission = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// System GreTunnel can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/gretunnel:Gretunnel labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/gretunnel:Gretunnel labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/gretunnel:Gretunnel")]
    public partial class Gretunnel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable validating checksums in received GRE packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("checksumReception")]
        public Output<string> ChecksumReception { get; private set; } = null!;

        /// <summary>
        /// Enable/disable including checksums in transmitted GRE packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("checksumTransmission")]
        public Output<string> ChecksumTransmission { get; private set; } = null!;

        /// <summary>
        /// DiffServ setting to be applied to GRE tunnel outer IP header.
        /// </summary>
        [Output("diffservcode")]
        public Output<string> Diffservcode { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DSCP copying. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("dscpCopying")]
        public Output<string> DscpCopying { get; private set; } = null!;

        /// <summary>
        /// Interface name.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// IP version to use for VPN interface. Valid values: `4`, `6`.
        /// </summary>
        [Output("ipVersion")]
        public Output<string> IpVersion { get; private set; } = null!;

        /// <summary>
        /// Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
        /// </summary>
        [Output("keepaliveFailtimes")]
        public Output<int> KeepaliveFailtimes { get; private set; } = null!;

        /// <summary>
        /// Keepalive message interval (0 - 32767, 0 = disabled).
        /// </summary>
        [Output("keepaliveInterval")]
        public Output<int> KeepaliveInterval { get; private set; } = null!;

        /// <summary>
        /// Require received GRE packets contain this key (0 - 4294967295).
        /// </summary>
        [Output("keyInbound")]
        public Output<int> KeyInbound { get; private set; } = null!;

        /// <summary>
        /// Include this key in transmitted GRE packets (0 - 4294967295).
        /// </summary>
        [Output("keyOutbound")]
        public Output<int> KeyOutbound { get; private set; } = null!;

        /// <summary>
        /// IP address of the local gateway.
        /// </summary>
        [Output("localGw")]
        public Output<string> LocalGw { get; private set; } = null!;

        /// <summary>
        /// IPv6 address of the local gateway.
        /// </summary>
        [Output("localGw6")]
        public Output<string> LocalGw6 { get; private set; } = null!;

        /// <summary>
        /// Tunnel name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// IP address of the remote gateway.
        /// </summary>
        [Output("remoteGw")]
        public Output<string> RemoteGw { get; private set; } = null!;

        /// <summary>
        /// IPv6 address of the remote gateway.
        /// </summary>
        [Output("remoteGw6")]
        public Output<string> RemoteGw6 { get; private set; } = null!;

        /// <summary>
        /// Enable/disable validating sequence numbers in received GRE packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("sequenceNumberReception")]
        public Output<string> SequenceNumberReception { get; private set; } = null!;

        /// <summary>
        /// Enable/disable including of sequence numbers in transmitted GRE packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("sequenceNumberTransmission")]
        public Output<string> SequenceNumberTransmission { get; private set; } = null!;

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
        /// Create a Gretunnel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Gretunnel(string name, GretunnelArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/gretunnel:Gretunnel", name, args ?? new GretunnelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Gretunnel(string name, Input<string> id, GretunnelState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/gretunnel:Gretunnel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Gretunnel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Gretunnel Get(string name, Input<string> id, GretunnelState? state = null, CustomResourceOptions? options = null)
        {
            return new Gretunnel(name, id, state, options);
        }
    }

    public sealed class GretunnelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable validating checksums in received GRE packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("checksumReception")]
        public Input<string>? ChecksumReception { get; set; }

        /// <summary>
        /// Enable/disable including checksums in transmitted GRE packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("checksumTransmission")]
        public Input<string>? ChecksumTransmission { get; set; }

        /// <summary>
        /// DiffServ setting to be applied to GRE tunnel outer IP header.
        /// </summary>
        [Input("diffservcode")]
        public Input<string>? Diffservcode { get; set; }

        /// <summary>
        /// Enable/disable DSCP copying. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("dscpCopying")]
        public Input<string>? DscpCopying { get; set; }

        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// IP version to use for VPN interface. Valid values: `4`, `6`.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        /// <summary>
        /// Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
        /// </summary>
        [Input("keepaliveFailtimes")]
        public Input<int>? KeepaliveFailtimes { get; set; }

        /// <summary>
        /// Keepalive message interval (0 - 32767, 0 = disabled).
        /// </summary>
        [Input("keepaliveInterval")]
        public Input<int>? KeepaliveInterval { get; set; }

        /// <summary>
        /// Require received GRE packets contain this key (0 - 4294967295).
        /// </summary>
        [Input("keyInbound")]
        public Input<int>? KeyInbound { get; set; }

        /// <summary>
        /// Include this key in transmitted GRE packets (0 - 4294967295).
        /// </summary>
        [Input("keyOutbound")]
        public Input<int>? KeyOutbound { get; set; }

        /// <summary>
        /// IP address of the local gateway.
        /// </summary>
        [Input("localGw", required: true)]
        public Input<string> LocalGw { get; set; } = null!;

        /// <summary>
        /// IPv6 address of the local gateway.
        /// </summary>
        [Input("localGw6")]
        public Input<string>? LocalGw6 { get; set; }

        /// <summary>
        /// Tunnel name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// IP address of the remote gateway.
        /// </summary>
        [Input("remoteGw", required: true)]
        public Input<string> RemoteGw { get; set; } = null!;

        /// <summary>
        /// IPv6 address of the remote gateway.
        /// </summary>
        [Input("remoteGw6")]
        public Input<string>? RemoteGw6 { get; set; }

        /// <summary>
        /// Enable/disable validating sequence numbers in received GRE packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sequenceNumberReception")]
        public Input<string>? SequenceNumberReception { get; set; }

        /// <summary>
        /// Enable/disable including of sequence numbers in transmitted GRE packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sequenceNumberTransmission")]
        public Input<string>? SequenceNumberTransmission { get; set; }

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

        public GretunnelArgs()
        {
        }
        public static new GretunnelArgs Empty => new GretunnelArgs();
    }

    public sealed class GretunnelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable validating checksums in received GRE packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("checksumReception")]
        public Input<string>? ChecksumReception { get; set; }

        /// <summary>
        /// Enable/disable including checksums in transmitted GRE packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("checksumTransmission")]
        public Input<string>? ChecksumTransmission { get; set; }

        /// <summary>
        /// DiffServ setting to be applied to GRE tunnel outer IP header.
        /// </summary>
        [Input("diffservcode")]
        public Input<string>? Diffservcode { get; set; }

        /// <summary>
        /// Enable/disable DSCP copying. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("dscpCopying")]
        public Input<string>? DscpCopying { get; set; }

        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// IP version to use for VPN interface. Valid values: `4`, `6`.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        /// <summary>
        /// Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
        /// </summary>
        [Input("keepaliveFailtimes")]
        public Input<int>? KeepaliveFailtimes { get; set; }

        /// <summary>
        /// Keepalive message interval (0 - 32767, 0 = disabled).
        /// </summary>
        [Input("keepaliveInterval")]
        public Input<int>? KeepaliveInterval { get; set; }

        /// <summary>
        /// Require received GRE packets contain this key (0 - 4294967295).
        /// </summary>
        [Input("keyInbound")]
        public Input<int>? KeyInbound { get; set; }

        /// <summary>
        /// Include this key in transmitted GRE packets (0 - 4294967295).
        /// </summary>
        [Input("keyOutbound")]
        public Input<int>? KeyOutbound { get; set; }

        /// <summary>
        /// IP address of the local gateway.
        /// </summary>
        [Input("localGw")]
        public Input<string>? LocalGw { get; set; }

        /// <summary>
        /// IPv6 address of the local gateway.
        /// </summary>
        [Input("localGw6")]
        public Input<string>? LocalGw6 { get; set; }

        /// <summary>
        /// Tunnel name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// IP address of the remote gateway.
        /// </summary>
        [Input("remoteGw")]
        public Input<string>? RemoteGw { get; set; }

        /// <summary>
        /// IPv6 address of the remote gateway.
        /// </summary>
        [Input("remoteGw6")]
        public Input<string>? RemoteGw6 { get; set; }

        /// <summary>
        /// Enable/disable validating sequence numbers in received GRE packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sequenceNumberReception")]
        public Input<string>? SequenceNumberReception { get; set; }

        /// <summary>
        /// Enable/disable including of sequence numbers in transmitted GRE packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sequenceNumberTransmission")]
        public Input<string>? SequenceNumberTransmission { get; set; }

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

        public GretunnelState()
        {
        }
        public static new GretunnelState Empty => new GretunnelState();
    }
}
