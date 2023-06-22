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
    public static class GetSystemPppoeinterfacelist
    {
        /// <summary>
        /// Provides a list of `fortios.SystemPppoeinterface`.
        /// </summary>
        public static Task<GetSystemPppoeinterfacelistResult> InvokeAsync(GetSystemPppoeinterfacelistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemPppoeinterfacelistResult>("fortios:index/getSystemPppoeinterfacelist:getSystemPppoeinterfacelist", args ?? new GetSystemPppoeinterfacelistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.SystemPppoeinterface`.
        /// </summary>
        public static Output<GetSystemPppoeinterfacelistResult> Invoke(GetSystemPppoeinterfacelistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemPppoeinterfacelistResult>("fortios:index/getSystemPppoeinterfacelist:getSystemPppoeinterfacelist", args ?? new GetSystemPppoeinterfacelistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemPppoeinterfacelistArgs : global::Pulumi.InvokeArgs
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

        public GetSystemPppoeinterfacelistArgs()
        {
        }
        public static new GetSystemPppoeinterfacelistArgs Empty => new GetSystemPppoeinterfacelistArgs();
    }

    public sealed class GetSystemPppoeinterfacelistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSystemPppoeinterfacelistInvokeArgs()
        {
        }
        public static new GetSystemPppoeinterfacelistInvokeArgs Empty => new GetSystemPppoeinterfacelistInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemPppoeinterfacelistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.SystemPppoeinterface`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemPppoeinterfacelistResult(
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
