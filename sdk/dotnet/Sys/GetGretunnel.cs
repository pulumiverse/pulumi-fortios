// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys
{
    public static class GetGretunnel
    {
        /// <summary>
        /// Use this data source to get information on an fortios system gretunnel
        /// </summary>
        public static Task<GetGretunnelResult> InvokeAsync(GetGretunnelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGretunnelResult>("fortios:sys/getGretunnel:getGretunnel", args ?? new GetGretunnelArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system gretunnel
        /// </summary>
        public static Output<GetGretunnelResult> Invoke(GetGretunnelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGretunnelResult>("fortios:sys/getGretunnel:getGretunnel", args ?? new GetGretunnelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGretunnelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system gretunnel.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetGretunnelArgs()
        {
        }
        public static new GetGretunnelArgs Empty => new GetGretunnelArgs();
    }

    public sealed class GetGretunnelInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system gretunnel.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetGretunnelInvokeArgs()
        {
        }
        public static new GetGretunnelInvokeArgs Empty => new GetGretunnelInvokeArgs();
    }


    [OutputType]
    public sealed class GetGretunnelResult
    {
        /// <summary>
        /// Enable/disable validating checksums in received GRE packets.
        /// </summary>
        public readonly string ChecksumReception;
        /// <summary>
        /// Enable/disable including checksums in transmitted GRE packets.
        /// </summary>
        public readonly string ChecksumTransmission;
        /// <summary>
        /// DiffServ setting to be applied to GRE tunnel outer IP header.
        /// </summary>
        public readonly string Diffservcode;
        /// <summary>
        /// Enable/disable DSCP copying.
        /// </summary>
        public readonly string DscpCopying;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// IP version to use for VPN interface.
        /// </summary>
        public readonly string IpVersion;
        /// <summary>
        /// Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
        /// </summary>
        public readonly int KeepaliveFailtimes;
        /// <summary>
        /// Keepalive message interval (0 - 32767, 0 = disabled).
        /// </summary>
        public readonly int KeepaliveInterval;
        /// <summary>
        /// Require received GRE packets contain this key (0 - 4294967295).
        /// </summary>
        public readonly int KeyInbound;
        /// <summary>
        /// Include this key in transmitted GRE packets (0 - 4294967295).
        /// </summary>
        public readonly int KeyOutbound;
        /// <summary>
        /// IP address of the local gateway.
        /// </summary>
        public readonly string LocalGw;
        /// <summary>
        /// IPv6 address of the local gateway.
        /// </summary>
        public readonly string LocalGw6;
        /// <summary>
        /// Tunnel name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// IP address of the remote gateway.
        /// </summary>
        public readonly string RemoteGw;
        /// <summary>
        /// IPv6 address of the remote gateway.
        /// </summary>
        public readonly string RemoteGw6;
        /// <summary>
        /// Enable/disable validating sequence numbers in received GRE packets.
        /// </summary>
        public readonly string SequenceNumberReception;
        /// <summary>
        /// Enable/disable including of sequence numbers in transmitted GRE packets.
        /// </summary>
        public readonly string SequenceNumberTransmission;
        /// <summary>
        /// Enable/disable use of SD-WAN to reach remote gateway.
        /// </summary>
        public readonly string UseSdwan;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetGretunnelResult(
            string checksumReception,

            string checksumTransmission,

            string diffservcode,

            string dscpCopying,

            string id,

            string @interface,

            string ipVersion,

            int keepaliveFailtimes,

            int keepaliveInterval,

            int keyInbound,

            int keyOutbound,

            string localGw,

            string localGw6,

            string name,

            string remoteGw,

            string remoteGw6,

            string sequenceNumberReception,

            string sequenceNumberTransmission,

            string useSdwan,

            string? vdomparam)
        {
            ChecksumReception = checksumReception;
            ChecksumTransmission = checksumTransmission;
            Diffservcode = diffservcode;
            DscpCopying = dscpCopying;
            Id = id;
            Interface = @interface;
            IpVersion = ipVersion;
            KeepaliveFailtimes = keepaliveFailtimes;
            KeepaliveInterval = keepaliveInterval;
            KeyInbound = keyInbound;
            KeyOutbound = keyOutbound;
            LocalGw = localGw;
            LocalGw6 = localGw6;
            Name = name;
            RemoteGw = remoteGw;
            RemoteGw6 = remoteGw6;
            SequenceNumberReception = sequenceNumberReception;
            SequenceNumberTransmission = sequenceNumberTransmission;
            UseSdwan = useSdwan;
            Vdomparam = vdomparam;
        }
    }
}
