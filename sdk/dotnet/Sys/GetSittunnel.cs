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
    public static class GetSittunnel
    {
        /// <summary>
        /// Use this data source to get information on an fortios system sittunnel
        /// </summary>
        public static Task<GetSittunnelResult> InvokeAsync(GetSittunnelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSittunnelResult>("fortios:sys/getSittunnel:getSittunnel", args ?? new GetSittunnelArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system sittunnel
        /// </summary>
        public static Output<GetSittunnelResult> Invoke(GetSittunnelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSittunnelResult>("fortios:sys/getSittunnel:getSittunnel", args ?? new GetSittunnelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSittunnelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system sittunnel.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSittunnelArgs()
        {
        }
        public static new GetSittunnelArgs Empty => new GetSittunnelArgs();
    }

    public sealed class GetSittunnelInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system sittunnel.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSittunnelInvokeArgs()
        {
        }
        public static new GetSittunnelInvokeArgs Empty => new GetSittunnelInvokeArgs();
    }


    [OutputType]
    public sealed class GetSittunnelResult
    {
        /// <summary>
        /// Enable/disable tunnel ASIC offloading.
        /// </summary>
        public readonly string AutoAsicOffload;
        /// <summary>
        /// Destination IP address of the tunnel.
        /// </summary>
        public readonly string Destination;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// IPv6 address of the tunnel.
        /// </summary>
        public readonly string Ip6;
        /// <summary>
        /// Tunnel name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Source IP address of the tunnel.
        /// </summary>
        public readonly string Source;
        /// <summary>
        /// Enable/disable use of SD-WAN to reach remote gateway.
        /// </summary>
        public readonly string UseSdwan;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSittunnelResult(
            string autoAsicOffload,

            string destination,

            string id,

            string @interface,

            string ip6,

            string name,

            string source,

            string useSdwan,

            string? vdomparam)
        {
            AutoAsicOffload = autoAsicOffload;
            Destination = destination;
            Id = id;
            Interface = @interface;
            Ip6 = ip6;
            Name = name;
            Source = source;
            UseSdwan = useSdwan;
            Vdomparam = vdomparam;
        }
    }
}
