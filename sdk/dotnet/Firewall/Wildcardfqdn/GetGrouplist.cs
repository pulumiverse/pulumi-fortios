// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Wildcardfqdn
{
    public static class GetGrouplist
    {
        /// <summary>
        /// Provides a list of `fortios.firewall/wildcardfqdn.Group`.
        /// </summary>
        public static Task<GetGrouplistResult> InvokeAsync(GetGrouplistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGrouplistResult>("fortios:firewall/wildcardfqdn/getGrouplist:getGrouplist", args ?? new GetGrouplistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.firewall/wildcardfqdn.Group`.
        /// </summary>
        public static Output<GetGrouplistResult> Invoke(GetGrouplistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGrouplistResult>("fortios:firewall/wildcardfqdn/getGrouplist:getGrouplist", args ?? new GetGrouplistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGrouplistArgs : global::Pulumi.InvokeArgs
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

        public GetGrouplistArgs()
        {
        }
        public static new GetGrouplistArgs Empty => new GetGrouplistArgs();
    }

    public sealed class GetGrouplistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetGrouplistInvokeArgs()
        {
        }
        public static new GetGrouplistInvokeArgs Empty => new GetGrouplistInvokeArgs();
    }


    [OutputType]
    public sealed class GetGrouplistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.firewall/wildcardfqdn.Group`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetGrouplistResult(
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
