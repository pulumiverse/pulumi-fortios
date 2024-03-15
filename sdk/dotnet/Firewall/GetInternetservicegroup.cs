// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    public static class GetInternetservicegroup
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall internetservicegroup
        /// </summary>
        public static Task<GetInternetservicegroupResult> InvokeAsync(GetInternetservicegroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInternetservicegroupResult>("fortios:firewall/getInternetservicegroup:getInternetservicegroup", args ?? new GetInternetservicegroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall internetservicegroup
        /// </summary>
        public static Output<GetInternetservicegroupResult> Invoke(GetInternetservicegroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInternetservicegroupResult>("fortios:firewall/getInternetservicegroup:getInternetservicegroup", args ?? new GetInternetservicegroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInternetservicegroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall internetservicegroup.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetInternetservicegroupArgs()
        {
        }
        public static new GetInternetservicegroupArgs Empty => new GetInternetservicegroupArgs();
    }

    public sealed class GetInternetservicegroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall internetservicegroup.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetInternetservicegroupInvokeArgs()
        {
        }
        public static new GetInternetservicegroupInvokeArgs Empty => new GetInternetservicegroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetInternetservicegroupResult
    {
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// How this service may be used (source, destination or both).
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Internet Service group member. The structure of `member` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInternetservicegroupMemberResult> Members;
        /// <summary>
        /// Internet Service name.
        /// </summary>
        public readonly string Name;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetInternetservicegroupResult(
            string comment,

            string direction,

            string id,

            ImmutableArray<Outputs.GetInternetservicegroupMemberResult> members,

            string name,

            string? vdomparam)
        {
            Comment = comment;
            Direction = direction;
            Id = id;
            Members = members;
            Name = name;
            Vdomparam = vdomparam;
        }
    }
}