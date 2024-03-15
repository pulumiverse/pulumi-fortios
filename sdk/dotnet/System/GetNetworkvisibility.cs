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
    public static class GetNetworkvisibility
    {
        /// <summary>
        /// Use this data source to get information on fortios system networkvisibility
        /// </summary>
        public static Task<GetNetworkvisibilityResult> InvokeAsync(GetNetworkvisibilityArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkvisibilityResult>("fortios:system/getNetworkvisibility:getNetworkvisibility", args ?? new GetNetworkvisibilityArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system networkvisibility
        /// </summary>
        public static Output<GetNetworkvisibilityResult> Invoke(GetNetworkvisibilityInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkvisibilityResult>("fortios:system/getNetworkvisibility:getNetworkvisibility", args ?? new GetNetworkvisibilityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkvisibilityArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetNetworkvisibilityArgs()
        {
        }
        public static new GetNetworkvisibilityArgs Empty => new GetNetworkvisibilityArgs();
    }

    public sealed class GetNetworkvisibilityInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetNetworkvisibilityInvokeArgs()
        {
        }
        public static new GetNetworkvisibilityInvokeArgs Empty => new GetNetworkvisibilityInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkvisibilityResult
    {
        /// <summary>
        /// Enable/disable logging of destination hostname visibility.
        /// </summary>
        public readonly string DestinationHostnameVisibility;
        /// <summary>
        /// Enable/disable logging of destination geographical location visibility.
        /// </summary>
        public readonly string DestinationLocation;
        /// <summary>
        /// Enable/disable logging of destination visibility.
        /// </summary>
        public readonly string DestinationVisibility;
        /// <summary>
        /// Limit of the number of hostname table entries (0 - 50000).
        /// </summary>
        public readonly int HostnameLimit;
        /// <summary>
        /// TTL of hostname table entries (60 - 86400).
        /// </summary>
        public readonly int HostnameTtl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable logging of source geographical location visibility.
        /// </summary>
        public readonly string SourceLocation;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetNetworkvisibilityResult(
            string destinationHostnameVisibility,

            string destinationLocation,

            string destinationVisibility,

            int hostnameLimit,

            int hostnameTtl,

            string id,

            string sourceLocation,

            string? vdomparam)
        {
            DestinationHostnameVisibility = destinationHostnameVisibility;
            DestinationLocation = destinationLocation;
            DestinationVisibility = destinationVisibility;
            HostnameLimit = hostnameLimit;
            HostnameTtl = hostnameTtl;
            Id = id;
            SourceLocation = sourceLocation;
            Vdomparam = vdomparam;
        }
    }
}