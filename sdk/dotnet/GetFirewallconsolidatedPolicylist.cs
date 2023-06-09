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
    public static class GetFirewallconsolidatedPolicylist
    {
        /// <summary>
        /// Provides a list of `fortios.FirewallconsolidatedPolicy`.
        /// </summary>
        public static Task<GetFirewallconsolidatedPolicylistResult> InvokeAsync(GetFirewallconsolidatedPolicylistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallconsolidatedPolicylistResult>("fortios:index/getFirewallconsolidatedPolicylist:getFirewallconsolidatedPolicylist", args ?? new GetFirewallconsolidatedPolicylistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.FirewallconsolidatedPolicy`.
        /// </summary>
        public static Output<GetFirewallconsolidatedPolicylistResult> Invoke(GetFirewallconsolidatedPolicylistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallconsolidatedPolicylistResult>("fortios:index/getFirewallconsolidatedPolicylist:getFirewallconsolidatedPolicylist", args ?? new GetFirewallconsolidatedPolicylistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallconsolidatedPolicylistArgs : global::Pulumi.InvokeArgs
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

        public GetFirewallconsolidatedPolicylistArgs()
        {
        }
        public static new GetFirewallconsolidatedPolicylistArgs Empty => new GetFirewallconsolidatedPolicylistArgs();
    }

    public sealed class GetFirewallconsolidatedPolicylistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetFirewallconsolidatedPolicylistInvokeArgs()
        {
        }
        public static new GetFirewallconsolidatedPolicylistInvokeArgs Empty => new GetFirewallconsolidatedPolicylistInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallconsolidatedPolicylistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.FirewallconsolidatedPolicy`.
        /// </summary>
        public readonly ImmutableArray<int> Policyidlists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallconsolidatedPolicylistResult(
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
