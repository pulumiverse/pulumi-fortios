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
    public static class GetSystemSessionhelperlist
    {
        /// <summary>
        /// Provides a list of `fortios.SystemSessionhelper`.
        /// </summary>
        public static Task<GetSystemSessionhelperlistResult> InvokeAsync(GetSystemSessionhelperlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemSessionhelperlistResult>("fortios:index/getSystemSessionhelperlist:getSystemSessionhelperlist", args ?? new GetSystemSessionhelperlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.SystemSessionhelper`.
        /// </summary>
        public static Output<GetSystemSessionhelperlistResult> Invoke(GetSystemSessionhelperlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemSessionhelperlistResult>("fortios:index/getSystemSessionhelperlist:getSystemSessionhelperlist", args ?? new GetSystemSessionhelperlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemSessionhelperlistArgs : global::Pulumi.InvokeArgs
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

        public GetSystemSessionhelperlistArgs()
        {
        }
        public static new GetSystemSessionhelperlistArgs Empty => new GetSystemSessionhelperlistArgs();
    }

    public sealed class GetSystemSessionhelperlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSystemSessionhelperlistInvokeArgs()
        {
        }
        public static new GetSystemSessionhelperlistInvokeArgs Empty => new GetSystemSessionhelperlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemSessionhelperlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// A list of the `fortios.SystemSessionhelper`.
        /// </summary>
        public readonly ImmutableArray<int> Fosidlists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemSessionhelperlistResult(
            string? filter,

            ImmutableArray<int> fosidlists,

            string id,

            string? vdomparam)
        {
            Filter = filter;
            Fosidlists = fosidlists;
            Id = id;
            Vdomparam = vdomparam;
        }
    }
}
