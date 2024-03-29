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
    public static class GetFssopolling
    {
        /// <summary>
        /// Use this data source to get information on fortios system fssopolling
        /// </summary>
        public static Task<GetFssopollingResult> InvokeAsync(GetFssopollingArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFssopollingResult>("fortios:system/getFssopolling:getFssopolling", args ?? new GetFssopollingArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system fssopolling
        /// </summary>
        public static Output<GetFssopollingResult> Invoke(GetFssopollingInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFssopollingResult>("fortios:system/getFssopolling:getFssopolling", args ?? new GetFssopollingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFssopollingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFssopollingArgs()
        {
        }
        public static new GetFssopollingArgs Empty => new GetFssopollingArgs();
    }

    public sealed class GetFssopollingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFssopollingInvokeArgs()
        {
        }
        public static new GetFssopollingInvokeArgs Empty => new GetFssopollingInvokeArgs();
    }


    [OutputType]
    public sealed class GetFssopollingResult
    {
        /// <summary>
        /// Password to connect to FSSO Agent.
        /// </summary>
        public readonly string AuthPassword;
        /// <summary>
        /// Enable/disable FSSO Agent Authentication.
        /// </summary>
        public readonly string Authentication;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Listening port to accept clients (1 - 65535).
        /// </summary>
        public readonly int ListeningPort;
        /// <summary>
        /// Enable/disable FSSO Polling Mode.
        /// </summary>
        public readonly string Status;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFssopollingResult(
            string authPassword,

            string authentication,

            string id,

            int listeningPort,

            string status,

            string? vdomparam)
        {
            AuthPassword = authPassword;
            Authentication = authentication;
            Id = id;
            ListeningPort = listeningPort;
            Status = status;
            Vdomparam = vdomparam;
        }
    }
}
