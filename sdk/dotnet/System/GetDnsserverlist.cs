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
    public static class GetDnsserverlist
    {
        /// <summary>
        /// Provides a list of `fortios.system.Dnsserver`.
        /// </summary>
        public static Task<GetDnsserverlistResult> InvokeAsync(GetDnsserverlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDnsserverlistResult>("fortios:system/getDnsserverlist:getDnsserverlist", args ?? new GetDnsserverlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.system.Dnsserver`.
        /// </summary>
        public static Output<GetDnsserverlistResult> Invoke(GetDnsserverlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDnsserverlistResult>("fortios:system/getDnsserverlist:getDnsserverlist", args ?? new GetDnsserverlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDnsserverlistArgs : global::Pulumi.InvokeArgs
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

        public GetDnsserverlistArgs()
        {
        }
        public static new GetDnsserverlistArgs Empty => new GetDnsserverlistArgs();
    }

    public sealed class GetDnsserverlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetDnsserverlistInvokeArgs()
        {
        }
        public static new GetDnsserverlistInvokeArgs Empty => new GetDnsserverlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetDnsserverlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.system.Dnsserver`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetDnsserverlistResult(
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
