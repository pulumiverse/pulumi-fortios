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
    public static class GetAlias
    {
        /// <summary>
        /// Use this data source to get information on an fortios system alias
        /// </summary>
        public static Task<GetAliasResult> InvokeAsync(GetAliasArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAliasResult>("fortios:system/getAlias:getAlias", args ?? new GetAliasArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system alias
        /// </summary>
        public static Output<GetAliasResult> Invoke(GetAliasInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAliasResult>("fortios:system/getAlias:getAlias", args ?? new GetAliasInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAliasArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system alias.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetAliasArgs()
        {
        }
        public static new GetAliasArgs Empty => new GetAliasArgs();
    }

    public sealed class GetAliasInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system alias.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetAliasInvokeArgs()
        {
        }
        public static new GetAliasInvokeArgs Empty => new GetAliasInvokeArgs();
    }


    [OutputType]
    public sealed class GetAliasResult
    {
        /// <summary>
        /// Command list to execute.
        /// </summary>
        public readonly string Command;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Alias command name.
        /// </summary>
        public readonly string Name;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetAliasResult(
            string command,

            string id,

            string name,

            string? vdomparam)
        {
            Command = command;
            Id = id;
            Name = name;
            Vdomparam = vdomparam;
        }
    }
}
