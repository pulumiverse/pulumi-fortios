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
    public static class GetRoutemaplist
    {
        /// <summary>
        /// Provides a list of `fortios.router.Routemap`.
        /// </summary>
        public static Task<GetRoutemaplistResult> InvokeAsync(GetRoutemaplistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRoutemaplistResult>("fortios:router/getRoutemaplist:getRoutemaplist", args ?? new GetRoutemaplistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.router.Routemap`.
        /// </summary>
        public static Output<GetRoutemaplistResult> Invoke(GetRoutemaplistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRoutemaplistResult>("fortios:router/getRoutemaplist:getRoutemaplist", args ?? new GetRoutemaplistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRoutemaplistArgs : global::Pulumi.InvokeArgs
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

        public GetRoutemaplistArgs()
        {
        }
        public static new GetRoutemaplistArgs Empty => new GetRoutemaplistArgs();
    }

    public sealed class GetRoutemaplistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetRoutemaplistInvokeArgs()
        {
        }
        public static new GetRoutemaplistInvokeArgs Empty => new GetRoutemaplistInvokeArgs();
    }


    [OutputType]
    public sealed class GetRoutemaplistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.router.Routemap`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetRoutemaplistResult(
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