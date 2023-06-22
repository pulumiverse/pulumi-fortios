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
    public static class GetRouterAuthpath
    {
        /// <summary>
        /// Use this data source to get information on an fortios router authpath
        /// </summary>
        public static Task<GetRouterAuthpathResult> InvokeAsync(GetRouterAuthpathArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouterAuthpathResult>("fortios:index/getRouterAuthpath:getRouterAuthpath", args ?? new GetRouterAuthpathArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios router authpath
        /// </summary>
        public static Output<GetRouterAuthpathResult> Invoke(GetRouterAuthpathInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouterAuthpathResult>("fortios:index/getRouterAuthpath:getRouterAuthpath", args ?? new GetRouterAuthpathInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouterAuthpathArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router authpath.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetRouterAuthpathArgs()
        {
        }
        public static new GetRouterAuthpathArgs Empty => new GetRouterAuthpathArgs();
    }

    public sealed class GetRouterAuthpathInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router authpath.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetRouterAuthpathInvokeArgs()
        {
        }
        public static new GetRouterAuthpathInvokeArgs Empty => new GetRouterAuthpathInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouterAuthpathResult
    {
        /// <summary>
        /// Outgoing interface.
        /// </summary>
        public readonly string Device;
        /// <summary>
        /// Gateway IP address.
        /// </summary>
        public readonly string Gateway;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the entry.
        /// </summary>
        public readonly string Name;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetRouterAuthpathResult(
            string device,

            string gateway,

            string id,

            string name,

            string? vdomparam)
        {
            Device = device;
            Gateway = gateway;
            Id = id;
            Name = name;
            Vdomparam = vdomparam;
        }
    }
}
