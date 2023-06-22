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
    public static class GetRouterMulticastflow
    {
        /// <summary>
        /// Use this data source to get information on an fortios router multicastflow
        /// </summary>
        public static Task<GetRouterMulticastflowResult> InvokeAsync(GetRouterMulticastflowArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouterMulticastflowResult>("fortios:index/getRouterMulticastflow:getRouterMulticastflow", args ?? new GetRouterMulticastflowArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios router multicastflow
        /// </summary>
        public static Output<GetRouterMulticastflowResult> Invoke(GetRouterMulticastflowInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouterMulticastflowResult>("fortios:index/getRouterMulticastflow:getRouterMulticastflow", args ?? new GetRouterMulticastflowInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouterMulticastflowArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router multicastflow.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetRouterMulticastflowArgs()
        {
        }
        public static new GetRouterMulticastflowArgs Empty => new GetRouterMulticastflowArgs();
    }

    public sealed class GetRouterMulticastflowInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router multicastflow.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetRouterMulticastflowInvokeArgs()
        {
        }
        public static new GetRouterMulticastflowInvokeArgs Empty => new GetRouterMulticastflowInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouterMulticastflowResult
    {
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// Multicast-flow entries. The structure of `flows` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterMulticastflowFlowResult> Flows;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name.
        /// </summary>
        public readonly string Name;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetRouterMulticastflowResult(
            string comments,

            ImmutableArray<Outputs.GetRouterMulticastflowFlowResult> flows,

            string id,

            string name,

            string? vdomparam)
        {
            Comments = comments;
            Flows = flows;
            Id = id;
            Name = name;
            Vdomparam = vdomparam;
        }
    }
}
