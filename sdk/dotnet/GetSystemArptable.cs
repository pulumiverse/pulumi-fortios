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
    public static class GetSystemArptable
    {
        /// <summary>
        /// Use this data source to get information on an fortios system arptable
        /// </summary>
        public static Task<GetSystemArptableResult> InvokeAsync(GetSystemArptableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemArptableResult>("fortios:index/getSystemArptable:getSystemArptable", args ?? new GetSystemArptableArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system arptable
        /// </summary>
        public static Output<GetSystemArptableResult> Invoke(GetSystemArptableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemArptableResult>("fortios:index/getSystemArptable:getSystemArptable", args ?? new GetSystemArptableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemArptableArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the fosid of the desired system arptable.
        /// </summary>
        [Input("fosid", required: true)]
        public int Fosid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemArptableArgs()
        {
        }
        public static new GetSystemArptableArgs Empty => new GetSystemArptableArgs();
    }

    public sealed class GetSystemArptableInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the fosid of the desired system arptable.
        /// </summary>
        [Input("fosid", required: true)]
        public Input<int> Fosid { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemArptableInvokeArgs()
        {
        }
        public static new GetSystemArptableInvokeArgs Empty => new GetSystemArptableInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemArptableResult
    {
        /// <summary>
        /// Unique integer ID of the entry.
        /// </summary>
        public readonly int Fosid;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// IP address.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// MAC address.
        /// </summary>
        public readonly string Mac;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemArptableResult(
            int fosid,

            string id,

            string @interface,

            string ip,

            string mac,

            string? vdomparam)
        {
            Fosid = fosid;
            Id = id;
            Interface = @interface;
            Ip = ip;
            Mac = mac;
            Vdomparam = vdomparam;
        }
    }
}
