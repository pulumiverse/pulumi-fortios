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
    public static class GetIpv6tunnel
    {
        /// <summary>
        /// Use this data source to get information on an fortios system ipv6tunnel
        /// </summary>
        public static Task<GetIpv6tunnelResult> InvokeAsync(GetIpv6tunnelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpv6tunnelResult>("fortios:system/getIpv6tunnel:getIpv6tunnel", args ?? new GetIpv6tunnelArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system ipv6tunnel
        /// </summary>
        public static Output<GetIpv6tunnelResult> Invoke(GetIpv6tunnelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpv6tunnelResult>("fortios:system/getIpv6tunnel:getIpv6tunnel", args ?? new GetIpv6tunnelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpv6tunnelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system ipv6tunnel.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetIpv6tunnelArgs()
        {
        }
        public static new GetIpv6tunnelArgs Empty => new GetIpv6tunnelArgs();
    }

    public sealed class GetIpv6tunnelInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system ipv6tunnel.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetIpv6tunnelInvokeArgs()
        {
        }
        public static new GetIpv6tunnelInvokeArgs Empty => new GetIpv6tunnelInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpv6tunnelResult
    {
        /// <summary>
        /// Enable/disable tunnel ASIC offloading.
        /// </summary>
        public readonly string AutoAsicOffload;
        /// <summary>
        /// Remote IPv6 address of the tunnel.
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
        /// IPv6 tunnel name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Local IPv6 address of the tunnel.
        /// </summary>
        public readonly string Source;
        /// <summary>
        /// Enable/disable use of SD-WAN to reach remote gateway.
        /// </summary>
        public readonly string UseSdwan;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetIpv6tunnelResult(
            string autoAsicOffload,

            string destination,

            string id,

            string @interface,

            string name,

            string source,

            string useSdwan,

            string? vdomparam)
        {
            AutoAsicOffload = autoAsicOffload;
            Destination = destination;
            Id = id;
            Interface = @interface;
            Name = name;
            Source = source;
            UseSdwan = useSdwan;
            Vdomparam = vdomparam;
        }
    }
}
