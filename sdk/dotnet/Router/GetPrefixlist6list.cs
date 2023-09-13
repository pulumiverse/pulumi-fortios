// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router
{
    public static class GetPrefixlist6list
    {
        /// <summary>
        /// Provides a list of `fortios.router.Prefixlist6`.
        /// </summary>
        public static Task<GetPrefixlist6listResult> InvokeAsync(GetPrefixlist6listArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrefixlist6listResult>("fortios:router/getPrefixlist6list:getPrefixlist6list", args ?? new GetPrefixlist6listArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.router.Prefixlist6`.
        /// </summary>
        public static Output<GetPrefixlist6listResult> Invoke(GetPrefixlist6listInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrefixlist6listResult>("fortios:router/getPrefixlist6list:getPrefixlist6list", args ?? new GetPrefixlist6listInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrefixlist6listArgs : global::Pulumi.InvokeArgs
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

        public GetPrefixlist6listArgs()
        {
        }
        public static new GetPrefixlist6listArgs Empty => new GetPrefixlist6listArgs();
    }

    public sealed class GetPrefixlist6listInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetPrefixlist6listInvokeArgs()
        {
        }
        public static new GetPrefixlist6listInvokeArgs Empty => new GetPrefixlist6listInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrefixlist6listResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.router.Prefixlist6`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetPrefixlist6listResult(
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
