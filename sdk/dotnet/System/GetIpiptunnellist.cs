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
    public static class GetIpiptunnellist
    {
        /// <summary>
        /// Provides a list of `fortios.system.Ipiptunnel`.
        /// </summary>
        public static Task<GetIpiptunnellistResult> InvokeAsync(GetIpiptunnellistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpiptunnellistResult>("fortios:system/getIpiptunnellist:getIpiptunnellist", args ?? new GetIpiptunnellistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.system.Ipiptunnel`.
        /// </summary>
        public static Output<GetIpiptunnellistResult> Invoke(GetIpiptunnellistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpiptunnellistResult>("fortios:system/getIpiptunnellist:getIpiptunnellist", args ?? new GetIpiptunnellistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpiptunnellistArgs : global::Pulumi.InvokeArgs
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

        public GetIpiptunnellistArgs()
        {
        }
        public static new GetIpiptunnellistArgs Empty => new GetIpiptunnellistArgs();
    }

    public sealed class GetIpiptunnellistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetIpiptunnellistInvokeArgs()
        {
        }
        public static new GetIpiptunnellistInvokeArgs Empty => new GetIpiptunnellistInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpiptunnellistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.system.Ipiptunnel`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetIpiptunnellistResult(
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
