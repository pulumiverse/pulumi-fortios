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
    public static class GetFirewallProfileprotocoloptionslist
    {
        /// <summary>
        /// Provides a list of `fortios.FirewallProfileprotocoloptions`.
        /// </summary>
        public static Task<GetFirewallProfileprotocoloptionslistResult> InvokeAsync(GetFirewallProfileprotocoloptionslistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallProfileprotocoloptionslistResult>("fortios:index/getFirewallProfileprotocoloptionslist:getFirewallProfileprotocoloptionslist", args ?? new GetFirewallProfileprotocoloptionslistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.FirewallProfileprotocoloptions`.
        /// </summary>
        public static Output<GetFirewallProfileprotocoloptionslistResult> Invoke(GetFirewallProfileprotocoloptionslistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallProfileprotocoloptionslistResult>("fortios:index/getFirewallProfileprotocoloptionslist:getFirewallProfileprotocoloptionslist", args ?? new GetFirewallProfileprotocoloptionslistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallProfileprotocoloptionslistArgs : global::Pulumi.InvokeArgs
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

        public GetFirewallProfileprotocoloptionslistArgs()
        {
        }
        public static new GetFirewallProfileprotocoloptionslistArgs Empty => new GetFirewallProfileprotocoloptionslistArgs();
    }

    public sealed class GetFirewallProfileprotocoloptionslistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetFirewallProfileprotocoloptionslistInvokeArgs()
        {
        }
        public static new GetFirewallProfileprotocoloptionslistInvokeArgs Empty => new GetFirewallProfileprotocoloptionslistInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallProfileprotocoloptionslistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.FirewallProfileprotocoloptions`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallProfileprotocoloptionslistResult(
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
