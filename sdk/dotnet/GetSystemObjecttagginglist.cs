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
    public static class GetSystemObjecttagginglist
    {
        /// <summary>
        /// Provides a list of `fortios.SystemObjecttagging`.
        /// </summary>
        public static Task<GetSystemObjecttagginglistResult> InvokeAsync(GetSystemObjecttagginglistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemObjecttagginglistResult>("fortios:index/getSystemObjecttagginglist:getSystemObjecttagginglist", args ?? new GetSystemObjecttagginglistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.SystemObjecttagging`.
        /// </summary>
        public static Output<GetSystemObjecttagginglistResult> Invoke(GetSystemObjecttagginglistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemObjecttagginglistResult>("fortios:index/getSystemObjecttagginglist:getSystemObjecttagginglist", args ?? new GetSystemObjecttagginglistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemObjecttagginglistArgs : global::Pulumi.InvokeArgs
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

        public GetSystemObjecttagginglistArgs()
        {
        }
        public static new GetSystemObjecttagginglistArgs Empty => new GetSystemObjecttagginglistArgs();
    }

    public sealed class GetSystemObjecttagginglistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSystemObjecttagginglistInvokeArgs()
        {
        }
        public static new GetSystemObjecttagginglistInvokeArgs Empty => new GetSystemObjecttagginglistInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemObjecttagginglistResult
    {
        /// <summary>
        /// A list of the `fortios.SystemObjecttagging`.
        /// </summary>
        public readonly ImmutableArray<string> Categorylists;
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemObjecttagginglistResult(
            ImmutableArray<string> categorylists,

            string? filter,

            string id,

            string? vdomparam)
        {
            Categorylists = categorylists;
            Filter = filter;
            Id = id;
            Vdomparam = vdomparam;
        }
    }
}
