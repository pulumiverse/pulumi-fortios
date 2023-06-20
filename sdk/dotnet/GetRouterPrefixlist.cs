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
    public static class GetRouterPrefixlist
    {
        /// <summary>
        /// Use this data source to get information on an fortios router prefixlist
        /// </summary>
        public static Task<GetRouterPrefixlistResult> InvokeAsync(GetRouterPrefixlistArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouterPrefixlistResult>("fortios:index/getRouterPrefixlist:getRouterPrefixlist", args ?? new GetRouterPrefixlistArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios router prefixlist
        /// </summary>
        public static Output<GetRouterPrefixlistResult> Invoke(GetRouterPrefixlistInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouterPrefixlistResult>("fortios:index/getRouterPrefixlist:getRouterPrefixlist", args ?? new GetRouterPrefixlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouterPrefixlistArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router prefixlist.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetRouterPrefixlistArgs()
        {
        }
        public static new GetRouterPrefixlistArgs Empty => new GetRouterPrefixlistArgs();
    }

    public sealed class GetRouterPrefixlistInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router prefixlist.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetRouterPrefixlistInvokeArgs()
        {
        }
        public static new GetRouterPrefixlistInvokeArgs Empty => new GetRouterPrefixlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouterPrefixlistResult
    {
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// IPv4 prefix list rule. The structure of `rule` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterPrefixlistRuleResult> Rules;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetRouterPrefixlistResult(
            string comments,

            string id,

            string name,

            ImmutableArray<Outputs.GetRouterPrefixlistRuleResult> rules,

            string? vdomparam)
        {
            Comments = comments;
            Id = id;
            Name = name;
            Rules = rules;
            Vdomparam = vdomparam;
        }
    }
}