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
    public static class GetWccplist
    {
        /// <summary>
        /// Provides a list of `fortios.system.Wccp`.
        /// </summary>
        public static Task<GetWccplistResult> InvokeAsync(GetWccplistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWccplistResult>("fortios:system/getWccplist:getWccplist", args ?? new GetWccplistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.system.Wccp`.
        /// </summary>
        public static Output<GetWccplistResult> Invoke(GetWccplistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWccplistResult>("fortios:system/getWccplist:getWccplist", args ?? new GetWccplistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWccplistArgs : global::Pulumi.InvokeArgs
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

        public GetWccplistArgs()
        {
        }
        public static new GetWccplistArgs Empty => new GetWccplistArgs();
    }

    public sealed class GetWccplistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetWccplistInvokeArgs()
        {
        }
        public static new GetWccplistInvokeArgs Empty => new GetWccplistInvokeArgs();
    }


    [OutputType]
    public sealed class GetWccplistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.system.Wccp`.
        /// </summary>
        public readonly ImmutableArray<string> ServiceIdlists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetWccplistResult(
            string? filter,

            string id,

            ImmutableArray<string> serviceIdlists,

            string? vdomparam)
        {
            Filter = filter;
            Id = id;
            ServiceIdlists = serviceIdlists;
            Vdomparam = vdomparam;
        }
    }
}
