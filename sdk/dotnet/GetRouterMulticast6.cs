// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    public static class GetRouterMulticast6
    {
        /// <summary>
        /// Use this data source to get information on fortios router multicast6
        /// </summary>
        public static Task<GetRouterMulticast6Result> InvokeAsync(GetRouterMulticast6Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouterMulticast6Result>("fortios:index/getRouterMulticast6:getRouterMulticast6", args ?? new GetRouterMulticast6Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios router multicast6
        /// </summary>
        public static Output<GetRouterMulticast6Result> Invoke(GetRouterMulticast6InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouterMulticast6Result>("fortios:index/getRouterMulticast6:getRouterMulticast6", args ?? new GetRouterMulticast6InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouterMulticast6Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetRouterMulticast6Args()
        {
        }
        public static new GetRouterMulticast6Args Empty => new GetRouterMulticast6Args();
    }

    public sealed class GetRouterMulticast6InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetRouterMulticast6InvokeArgs()
        {
        }
        public static new GetRouterMulticast6InvokeArgs Empty => new GetRouterMulticast6InvokeArgs();
    }


    [OutputType]
    public sealed class GetRouterMulticast6Result
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterMulticast6InterfaceResult> Interfaces;
        /// <summary>
        /// Enable/disable PMTU for IPv6 multicast.
        /// </summary>
        public readonly string MulticastPmtu;
        /// <summary>
        /// Enable/disable IPv6 multicast routing.
        /// </summary>
        public readonly string MulticastRouting;
        /// <summary>
        /// PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterMulticast6PimSmGlobalResult> PimSmGlobals;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetRouterMulticast6Result(
            string id,

            ImmutableArray<Outputs.GetRouterMulticast6InterfaceResult> interfaces,

            string multicastPmtu,

            string multicastRouting,

            ImmutableArray<Outputs.GetRouterMulticast6PimSmGlobalResult> pimSmGlobals,

            string? vdomparam)
        {
            Id = id;
            Interfaces = interfaces;
            MulticastPmtu = multicastPmtu;
            MulticastRouting = multicastRouting;
            PimSmGlobals = pimSmGlobals;
            Vdomparam = vdomparam;
        }
    }
}
