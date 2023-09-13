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
    public static class GetCommunitylistlist
    {
        /// <summary>
        /// Provides a list of `fortios.router.Communitylist`.
        /// </summary>
        public static Task<GetCommunitylistlistResult> InvokeAsync(GetCommunitylistlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCommunitylistlistResult>("fortios:router/getCommunitylistlist:getCommunitylistlist", args ?? new GetCommunitylistlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.router.Communitylist`.
        /// </summary>
        public static Output<GetCommunitylistlistResult> Invoke(GetCommunitylistlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCommunitylistlistResult>("fortios:router/getCommunitylistlist:getCommunitylistlist", args ?? new GetCommunitylistlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCommunitylistlistArgs : global::Pulumi.InvokeArgs
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

        public GetCommunitylistlistArgs()
        {
        }
        public static new GetCommunitylistlistArgs Empty => new GetCommunitylistlistArgs();
    }

    public sealed class GetCommunitylistlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetCommunitylistlistInvokeArgs()
        {
        }
        public static new GetCommunitylistlistInvokeArgs Empty => new GetCommunitylistlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetCommunitylistlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.router.Communitylist`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetCommunitylistlistResult(
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
