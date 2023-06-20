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
    public static class GetRouterCommunitylist
    {
        /// <summary>
        /// Use this data source to get information on an fortios router communitylist
        /// </summary>
        public static Task<GetRouterCommunitylistResult> InvokeAsync(GetRouterCommunitylistArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouterCommunitylistResult>("fortios:index/getRouterCommunitylist:getRouterCommunitylist", args ?? new GetRouterCommunitylistArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios router communitylist
        /// </summary>
        public static Output<GetRouterCommunitylistResult> Invoke(GetRouterCommunitylistInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouterCommunitylistResult>("fortios:index/getRouterCommunitylist:getRouterCommunitylist", args ?? new GetRouterCommunitylistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouterCommunitylistArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router communitylist.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetRouterCommunitylistArgs()
        {
        }
        public static new GetRouterCommunitylistArgs Empty => new GetRouterCommunitylistArgs();
    }

    public sealed class GetRouterCommunitylistInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router communitylist.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetRouterCommunitylistInvokeArgs()
        {
        }
        public static new GetRouterCommunitylistInvokeArgs Empty => new GetRouterCommunitylistInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouterCommunitylistResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Community list name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Community list rule. The structure of `rule` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterCommunitylistRuleResult> Rules;
        /// <summary>
        /// Community list type (standard or expanded).
        /// </summary>
        public readonly string Type;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetRouterCommunitylistResult(
            string id,

            string name,

            ImmutableArray<Outputs.GetRouterCommunitylistRuleResult> rules,

            string type,

            string? vdomparam)
        {
            Id = id;
            Name = name;
            Rules = rules;
            Type = type;
            Vdomparam = vdomparam;
        }
    }
}