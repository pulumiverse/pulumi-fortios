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
    public static class GetSystemAutomationdestinationlist
    {
        /// <summary>
        /// Provides a list of `fortios.SystemAutomationdestination`.
        /// </summary>
        public static Task<GetSystemAutomationdestinationlistResult> InvokeAsync(GetSystemAutomationdestinationlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemAutomationdestinationlistResult>("fortios:index/getSystemAutomationdestinationlist:getSystemAutomationdestinationlist", args ?? new GetSystemAutomationdestinationlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.SystemAutomationdestination`.
        /// </summary>
        public static Output<GetSystemAutomationdestinationlistResult> Invoke(GetSystemAutomationdestinationlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemAutomationdestinationlistResult>("fortios:index/getSystemAutomationdestinationlist:getSystemAutomationdestinationlist", args ?? new GetSystemAutomationdestinationlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemAutomationdestinationlistArgs : global::Pulumi.InvokeArgs
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

        public GetSystemAutomationdestinationlistArgs()
        {
        }
        public static new GetSystemAutomationdestinationlistArgs Empty => new GetSystemAutomationdestinationlistArgs();
    }

    public sealed class GetSystemAutomationdestinationlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSystemAutomationdestinationlistInvokeArgs()
        {
        }
        public static new GetSystemAutomationdestinationlistInvokeArgs Empty => new GetSystemAutomationdestinationlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemAutomationdestinationlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.SystemAutomationdestination`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemAutomationdestinationlistResult(
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
