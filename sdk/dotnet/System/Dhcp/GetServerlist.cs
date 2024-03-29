// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Dhcp
{
    public static class GetServerlist
    {
        /// <summary>
        /// Provides a list of `fortios.system/dhcp.Server`.
        /// </summary>
        public static Task<GetServerlistResult> InvokeAsync(GetServerlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerlistResult>("fortios:system/dhcp/getServerlist:getServerlist", args ?? new GetServerlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.system/dhcp.Server`.
        /// </summary>
        public static Output<GetServerlistResult> Invoke(GetServerlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerlistResult>("fortios:system/dhcp/getServerlist:getServerlist", args ?? new GetServerlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerlistArgs : global::Pulumi.InvokeArgs
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

        public GetServerlistArgs()
        {
        }
        public static new GetServerlistArgs Empty => new GetServerlistArgs();
    }

    public sealed class GetServerlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetServerlistInvokeArgs()
        {
        }
        public static new GetServerlistInvokeArgs Empty => new GetServerlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// A list of the `fortios.system/dhcp.Server`.
        /// </summary>
        public readonly ImmutableArray<int> Fosidlists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetServerlistResult(
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
