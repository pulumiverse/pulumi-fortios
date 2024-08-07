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
    public static class GetDdnslist
    {
        /// <summary>
        /// Provides a list of `fortios.system.Ddns`.
        /// </summary>
        public static Task<GetDdnslistResult> InvokeAsync(GetDdnslistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDdnslistResult>("fortios:system/getDdnslist:getDdnslist", args ?? new GetDdnslistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.system.Ddns`.
        /// </summary>
        public static Output<GetDdnslistResult> Invoke(GetDdnslistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDdnslistResult>("fortios:system/getDdnslist:getDdnslist", args ?? new GetDdnslistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDdnslistArgs : global::Pulumi.InvokeArgs
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

        public GetDdnslistArgs()
        {
        }
        public static new GetDdnslistArgs Empty => new GetDdnslistArgs();
    }

    public sealed class GetDdnslistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetDdnslistInvokeArgs()
        {
        }
        public static new GetDdnslistInvokeArgs Empty => new GetDdnslistInvokeArgs();
    }


    [OutputType]
    public sealed class GetDdnslistResult
    {
        /// <summary>
        /// A list of the `fortios.system.Ddns`.
        /// </summary>
        public readonly ImmutableArray<int> Ddnsidlists;
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetDdnslistResult(
            ImmutableArray<int> ddnsidlists,

            string? filter,

            string id,

            string? vdomparam)
        {
            Ddnsidlists = ddnsidlists;
            Filter = filter;
            Id = id;
            Vdomparam = vdomparam;
        }
    }
}
