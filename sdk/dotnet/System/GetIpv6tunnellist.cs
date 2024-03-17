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
    public static class GetIpv6tunnellist
    {
        /// <summary>
        /// Provides a list of `fortios.system.Ipv6tunnel`.
        /// </summary>
        public static Task<GetIpv6tunnellistResult> InvokeAsync(GetIpv6tunnellistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpv6tunnellistResult>("fortios:system/getIpv6tunnellist:getIpv6tunnellist", args ?? new GetIpv6tunnellistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.system.Ipv6tunnel`.
        /// </summary>
        public static Output<GetIpv6tunnellistResult> Invoke(GetIpv6tunnellistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpv6tunnellistResult>("fortios:system/getIpv6tunnellist:getIpv6tunnellist", args ?? new GetIpv6tunnellistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpv6tunnellistArgs : global::Pulumi.InvokeArgs
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

        public GetIpv6tunnellistArgs()
        {
        }
        public static new GetIpv6tunnellistArgs Empty => new GetIpv6tunnellistArgs();
    }

    public sealed class GetIpv6tunnellistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetIpv6tunnellistInvokeArgs()
        {
        }
        public static new GetIpv6tunnellistInvokeArgs Empty => new GetIpv6tunnellistInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpv6tunnellistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.system.Ipv6tunnel`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetIpv6tunnellistResult(
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
