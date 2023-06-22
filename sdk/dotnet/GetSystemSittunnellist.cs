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
    public static class GetSystemSittunnellist
    {
        /// <summary>
        /// Provides a list of `fortios.SystemSittunnel`.
        /// </summary>
        public static Task<GetSystemSittunnellistResult> InvokeAsync(GetSystemSittunnellistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemSittunnellistResult>("fortios:index/getSystemSittunnellist:getSystemSittunnellist", args ?? new GetSystemSittunnellistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.SystemSittunnel`.
        /// </summary>
        public static Output<GetSystemSittunnellistResult> Invoke(GetSystemSittunnellistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemSittunnellistResult>("fortios:index/getSystemSittunnellist:getSystemSittunnellist", args ?? new GetSystemSittunnellistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemSittunnellistArgs : global::Pulumi.InvokeArgs
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

        public GetSystemSittunnellistArgs()
        {
        }
        public static new GetSystemSittunnellistArgs Empty => new GetSystemSittunnellistArgs();
    }

    public sealed class GetSystemSittunnellistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSystemSittunnellistInvokeArgs()
        {
        }
        public static new GetSystemSittunnellistInvokeArgs Empty => new GetSystemSittunnellistInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemSittunnellistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.SystemSittunnel`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemSittunnellistResult(
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
