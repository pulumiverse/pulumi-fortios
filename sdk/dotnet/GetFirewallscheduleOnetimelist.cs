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
    public static class GetFirewallscheduleOnetimelist
    {
        /// <summary>
        /// Provides a list of `fortios.FirewallscheduleOnetime`.
        /// </summary>
        public static Task<GetFirewallscheduleOnetimelistResult> InvokeAsync(GetFirewallscheduleOnetimelistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallscheduleOnetimelistResult>("fortios:index/getFirewallscheduleOnetimelist:getFirewallscheduleOnetimelist", args ?? new GetFirewallscheduleOnetimelistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.FirewallscheduleOnetime`.
        /// </summary>
        public static Output<GetFirewallscheduleOnetimelistResult> Invoke(GetFirewallscheduleOnetimelistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallscheduleOnetimelistResult>("fortios:index/getFirewallscheduleOnetimelist:getFirewallscheduleOnetimelist", args ?? new GetFirewallscheduleOnetimelistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallscheduleOnetimelistArgs : global::Pulumi.InvokeArgs
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

        public GetFirewallscheduleOnetimelistArgs()
        {
        }
        public static new GetFirewallscheduleOnetimelistArgs Empty => new GetFirewallscheduleOnetimelistArgs();
    }

    public sealed class GetFirewallscheduleOnetimelistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetFirewallscheduleOnetimelistInvokeArgs()
        {
        }
        public static new GetFirewallscheduleOnetimelistInvokeArgs Empty => new GetFirewallscheduleOnetimelistInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallscheduleOnetimelistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.FirewallscheduleOnetime`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallscheduleOnetimelistResult(
            string? filter,

            string id,

            ImmutableArray<string> namelists,

            string? vdomparam)
        {
            Filter = filter;
            Id = id;
            Namelists = namelists;
            Vdomparam = vdomparam;
        }
    }
}
