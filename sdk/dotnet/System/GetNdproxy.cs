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
    public static class GetNdproxy
    {
        /// <summary>
        /// Use this data source to get information on fortios system ndproxy
        /// </summary>
        public static Task<GetNdproxyResult> InvokeAsync(GetNdproxyArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNdproxyResult>("fortios:system/getNdproxy:getNdproxy", args ?? new GetNdproxyArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system ndproxy
        /// </summary>
        public static Output<GetNdproxyResult> Invoke(GetNdproxyInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNdproxyResult>("fortios:system/getNdproxy:getNdproxy", args ?? new GetNdproxyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNdproxyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetNdproxyArgs()
        {
        }
        public static new GetNdproxyArgs Empty => new GetNdproxyArgs();
    }

    public sealed class GetNdproxyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetNdproxyInvokeArgs()
        {
        }
        public static new GetNdproxyInvokeArgs Empty => new GetNdproxyInvokeArgs();
    }


    [OutputType]
    public sealed class GetNdproxyResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Interfaces using the neighbor discovery proxy. The structure of `member` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNdproxyMemberResult> Members;
        /// <summary>
        /// Enable/disable neighbor discovery proxy.
        /// </summary>
        public readonly string Status;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetNdproxyResult(
            string id,

            ImmutableArray<Outputs.GetNdproxyMemberResult> members,

            string status,

            string? vdomparam)
        {
            Id = id;
            Members = members;
            Status = status;
            Vdomparam = vdomparam;
        }
    }
}
