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
    public static class GetFirewallAddrgrplist
    {
        /// <summary>
        /// Provides a list of `fortios.FirewallAddrgrp`.
        /// </summary>
        public static Task<GetFirewallAddrgrplistResult> InvokeAsync(GetFirewallAddrgrplistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallAddrgrplistResult>("fortios:index/getFirewallAddrgrplist:getFirewallAddrgrplist", args ?? new GetFirewallAddrgrplistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.FirewallAddrgrp`.
        /// </summary>
        public static Output<GetFirewallAddrgrplistResult> Invoke(GetFirewallAddrgrplistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallAddrgrplistResult>("fortios:index/getFirewallAddrgrplist:getFirewallAddrgrplist", args ?? new GetFirewallAddrgrplistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallAddrgrplistArgs : global::Pulumi.InvokeArgs
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

        public GetFirewallAddrgrplistArgs()
        {
        }
        public static new GetFirewallAddrgrplistArgs Empty => new GetFirewallAddrgrplistArgs();
    }

    public sealed class GetFirewallAddrgrplistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetFirewallAddrgrplistInvokeArgs()
        {
        }
        public static new GetFirewallAddrgrplistInvokeArgs Empty => new GetFirewallAddrgrplistInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallAddrgrplistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.FirewallAddrgrp`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallAddrgrplistResult(
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