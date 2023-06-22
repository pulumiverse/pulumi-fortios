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
    public static class GetSystemEmailserver
    {
        /// <summary>
        /// Use this data source to get information on fortios system emailserver
        /// </summary>
        public static Task<GetSystemEmailserverResult> InvokeAsync(GetSystemEmailserverArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemEmailserverResult>("fortios:index/getSystemEmailserver:getSystemEmailserver", args ?? new GetSystemEmailserverArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system emailserver
        /// </summary>
        public static Output<GetSystemEmailserverResult> Invoke(GetSystemEmailserverInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemEmailserverResult>("fortios:index/getSystemEmailserver:getSystemEmailserver", args ?? new GetSystemEmailserverInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemEmailserverArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemEmailserverArgs()
        {
        }
        public static new GetSystemEmailserverArgs Empty => new GetSystemEmailserverArgs();
    }

    public sealed class GetSystemEmailserverInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemEmailserverInvokeArgs()
        {
        }
        public static new GetSystemEmailserverInvokeArgs Empty => new GetSystemEmailserverInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemEmailserverResult
    {
        /// <summary>
        /// Enable/disable authentication.
        /// </summary>
        public readonly string Authenticate;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// Specify how to select outgoing interface to reach server.
        /// </summary>
        public readonly string InterfaceSelectMethod;
        /// <summary>
        /// SMTP server user password for authentication.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// SMTP server port.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Reply-To email address.
        /// </summary>
        public readonly string ReplyTo;
        /// <summary>
        /// Connection security used by the email server.
        /// </summary>
        public readonly string Security;
        /// <summary>
        /// SMTP server IP address or hostname.
        /// </summary>
        public readonly string Server;
        /// <summary>
        /// SMTP server IPv4 source IP.
        /// </summary>
        public readonly string SourceIp;
        /// <summary>
        /// SMTP server IPv6 source IP.
        /// </summary>
        public readonly string SourceIp6;
        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        public readonly string SslMinProtoVersion;
        /// <summary>
        /// Use FortiGuard Message service or custom email server.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// SMTP server user name for authentication.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// Enable/disable validation of server certificate.
        /// </summary>
        public readonly string ValidateServer;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemEmailserverResult(
            string authenticate,

            string id,

            string @interface,

            string interfaceSelectMethod,

            string password,

            int port,

            string replyTo,

            string security,

            string server,

            string sourceIp,

            string sourceIp6,

            string sslMinProtoVersion,

            string type,

            string username,

            string validateServer,

            string? vdomparam)
        {
            Authenticate = authenticate;
            Id = id;
            Interface = @interface;
            InterfaceSelectMethod = interfaceSelectMethod;
            Password = password;
            Port = port;
            ReplyTo = replyTo;
            Security = security;
            Server = server;
            SourceIp = sourceIp;
            SourceIp6 = sourceIp6;
            SslMinProtoVersion = sslMinProtoVersion;
            Type = type;
            Username = username;
            ValidateServer = validateServer;
            Vdomparam = vdomparam;
        }
    }
}
