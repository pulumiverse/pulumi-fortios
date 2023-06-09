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
    public static class GetSystemFssopolling
    {
        /// <summary>
        /// Use this data source to get information on fortios system fssopolling
        /// </summary>
        public static Task<GetSystemFssopollingResult> InvokeAsync(GetSystemFssopollingArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemFssopollingResult>("fortios:index/getSystemFssopolling:getSystemFssopolling", args ?? new GetSystemFssopollingArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system fssopolling
        /// </summary>
        public static Output<GetSystemFssopollingResult> Invoke(GetSystemFssopollingInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemFssopollingResult>("fortios:index/getSystemFssopolling:getSystemFssopolling", args ?? new GetSystemFssopollingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemFssopollingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemFssopollingArgs()
        {
        }
        public static new GetSystemFssopollingArgs Empty => new GetSystemFssopollingArgs();
    }

    public sealed class GetSystemFssopollingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemFssopollingInvokeArgs()
        {
        }
        public static new GetSystemFssopollingInvokeArgs Empty => new GetSystemFssopollingInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemFssopollingResult
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
        private GetSystemFssopollingResult(
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
