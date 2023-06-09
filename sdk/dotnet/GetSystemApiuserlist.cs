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
    public static class GetSystemApiuserlist
    {
        /// <summary>
        /// Provides a list of `fortios.SystemApiuser`.
        /// </summary>
        public static Task<GetSystemApiuserlistResult> InvokeAsync(GetSystemApiuserlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemApiuserlistResult>("fortios:index/getSystemApiuserlist:getSystemApiuserlist", args ?? new GetSystemApiuserlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.SystemApiuser`.
        /// </summary>
        public static Output<GetSystemApiuserlistResult> Invoke(GetSystemApiuserlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemApiuserlistResult>("fortios:index/getSystemApiuserlist:getSystemApiuserlist", args ?? new GetSystemApiuserlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemApiuserlistArgs : global::Pulumi.InvokeArgs
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

        public GetSystemApiuserlistArgs()
        {
        }
        public static new GetSystemApiuserlistArgs Empty => new GetSystemApiuserlistArgs();
    }

    public sealed class GetSystemApiuserlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSystemApiuserlistInvokeArgs()
        {
        }
        public static new GetSystemApiuserlistInvokeArgs Empty => new GetSystemApiuserlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemApiuserlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.SystemApiuser`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemApiuserlistResult(
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
