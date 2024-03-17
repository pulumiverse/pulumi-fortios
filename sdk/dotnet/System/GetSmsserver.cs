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
    public static class GetSmsserver
    {
        /// <summary>
        /// Use this data source to get information on an fortios system smsserver
        /// </summary>
        public static Task<GetSmsserverResult> InvokeAsync(GetSmsserverArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSmsserverResult>("fortios:system/getSmsserver:getSmsserver", args ?? new GetSmsserverArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system smsserver
        /// </summary>
        public static Output<GetSmsserverResult> Invoke(GetSmsserverInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSmsserverResult>("fortios:system/getSmsserver:getSmsserver", args ?? new GetSmsserverInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSmsserverArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system smsserver.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSmsserverArgs()
        {
        }
        public static new GetSmsserverArgs Empty => new GetSmsserverArgs();
    }

    public sealed class GetSmsserverInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system smsserver.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSmsserverInvokeArgs()
        {
        }
        public static new GetSmsserverInvokeArgs Empty => new GetSmsserverInvokeArgs();
    }


    [OutputType]
    public sealed class GetSmsserverResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Email-to-SMS server domain name.
        /// </summary>
        public readonly string MailServer;
        /// <summary>
        /// Name of SMS server.
        /// </summary>
        public readonly string Name;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSmsserverResult(
            string id,

            string mailServer,

            string name,

            string? vdomparam)
        {
            Id = id;
            MailServer = mailServer;
            Name = name;
            Vdomparam = vdomparam;
        }
    }
}
