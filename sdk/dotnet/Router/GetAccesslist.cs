// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router
{
    public static class GetAccesslist
    {
        /// <summary>
        /// Use this data source to get information on an fortios router accesslist
        /// </summary>
        public static Task<GetAccesslistResult> InvokeAsync(GetAccesslistArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccesslistResult>("fortios:router/getAccesslist:getAccesslist", args ?? new GetAccesslistArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios router accesslist
        /// </summary>
        public static Output<GetAccesslistResult> Invoke(GetAccesslistInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccesslistResult>("fortios:router/getAccesslist:getAccesslist", args ?? new GetAccesslistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccesslistArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router accesslist.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetAccesslistArgs()
        {
        }
        public static new GetAccesslistArgs Empty => new GetAccesslistArgs();
    }

    public sealed class GetAccesslistInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router accesslist.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetAccesslistInvokeArgs()
        {
        }
        public static new GetAccesslistInvokeArgs Empty => new GetAccesslistInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccesslistResult
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
        /// Rule. The structure of `rule` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccesslistRuleResult> Rules;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetAccesslistResult(
            string comments,

            string id,

            string name,

            ImmutableArray<Outputs.GetAccesslistRuleResult> rules,

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