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
    public static class GetObjecttagginglist
    {
        /// <summary>
        /// Provides a list of `fortios.system.Objecttagging`.
        /// </summary>
        public static Task<GetObjecttagginglistResult> InvokeAsync(GetObjecttagginglistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetObjecttagginglistResult>("fortios:system/getObjecttagginglist:getObjecttagginglist", args ?? new GetObjecttagginglistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.system.Objecttagging`.
        /// </summary>
        public static Output<GetObjecttagginglistResult> Invoke(GetObjecttagginglistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetObjecttagginglistResult>("fortios:system/getObjecttagginglist:getObjecttagginglist", args ?? new GetObjecttagginglistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetObjecttagginglistArgs : global::Pulumi.InvokeArgs
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

        public GetObjecttagginglistArgs()
        {
        }
        public static new GetObjecttagginglistArgs Empty => new GetObjecttagginglistArgs();
    }

    public sealed class GetObjecttagginglistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetObjecttagginglistInvokeArgs()
        {
        }
        public static new GetObjecttagginglistInvokeArgs Empty => new GetObjecttagginglistInvokeArgs();
    }


    [OutputType]
    public sealed class GetObjecttagginglistResult
    {
        /// <summary>
        /// A list of the `fortios.system.Objecttagging`.
        /// </summary>
        public readonly ImmutableArray<string> Categorylists;
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetObjecttagginglistResult(
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
