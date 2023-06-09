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
    public static class GetFirewallDoSpolicy6list
    {
        /// <summary>
        /// Provides a list of `fortios_firewall_DoSpolicy6`.
        /// </summary>
        public static Task<GetFirewallDoSpolicy6listResult> InvokeAsync(GetFirewallDoSpolicy6listArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallDoSpolicy6listResult>("fortios:index/getFirewallDoSpolicy6list:getFirewallDoSpolicy6list", args ?? new GetFirewallDoSpolicy6listArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios_firewall_DoSpolicy6`.
        /// </summary>
        public static Output<GetFirewallDoSpolicy6listResult> Invoke(GetFirewallDoSpolicy6listInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallDoSpolicy6listResult>("fortios:index/getFirewallDoSpolicy6list:getFirewallDoSpolicy6list", args ?? new GetFirewallDoSpolicy6listInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallDoSpolicy6listArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter used to scope the list. See Filter results of datasource.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallDoSpolicy6listArgs()
        {
        }
        public static new GetFirewallDoSpolicy6listArgs Empty => new GetFirewallDoSpolicy6listArgs();
    }

    public sealed class GetFirewallDoSpolicy6listInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter used to scope the list. See Filter results of datasource.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallDoSpolicy6listInvokeArgs()
        {
        }
        public static new GetFirewallDoSpolicy6listInvokeArgs Empty => new GetFirewallDoSpolicy6listInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallDoSpolicy6listResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios_firewall_DoSpolicy6`.
        /// </summary>
        public readonly ImmutableArray<int> Policyidlists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallDoSpolicy6listResult(
            string? filter,

            string id,

            ImmutableArray<int> policyidlists,

            string? vdomparam)
        {
            Filter = filter;
            Id = id;
            Policyidlists = policyidlists;
            Vdomparam = vdomparam;
        }
    }
}
