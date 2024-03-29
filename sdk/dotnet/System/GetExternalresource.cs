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
    public static class GetExternalresource
    {
        /// <summary>
        /// Use this data source to get information on an fortios system externalresource
        /// </summary>
        public static Task<GetExternalresourceResult> InvokeAsync(GetExternalresourceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExternalresourceResult>("fortios:system/getExternalresource:getExternalresource", args ?? new GetExternalresourceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system externalresource
        /// </summary>
        public static Output<GetExternalresourceResult> Invoke(GetExternalresourceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExternalresourceResult>("fortios:system/getExternalresource:getExternalresource", args ?? new GetExternalresourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExternalresourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system externalresource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetExternalresourceArgs()
        {
        }
        public static new GetExternalresourceArgs Empty => new GetExternalresourceArgs();
    }

    public sealed class GetExternalresourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system externalresource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetExternalresourceInvokeArgs()
        {
        }
        public static new GetExternalresourceInvokeArgs Empty => new GetExternalresourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetExternalresourceResult
    {
        /// <summary>
        /// User resource category.
        /// </summary>
        public readonly int Category;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comments;
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
        /// External resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// HTTP basic authentication password.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Time interval to refresh external resource (1 - 43200 min, default = 5 min).
        /// </summary>
        public readonly int RefreshRate;
        /// <summary>
        /// URI of external resource.
        /// </summary>
        public readonly string Resource;
        /// <summary>
        /// Certificate verification option.
        /// </summary>
        public readonly string ServerIdentityCheck;
        /// <summary>
        /// Source IPv4 address used to communicate with server.
        /// </summary>
        public readonly string SourceIp;
        /// <summary>
        /// Enable/disable user resource.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// User resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// External resource update method.
        /// </summary>
        public readonly string UpdateMethod;
        /// <summary>
        /// Override HTTP User-Agent header used when retrieving this external resource.
        /// </summary>
        public readonly string UserAgent;
        /// <summary>
        /// HTTP basic authentication user name.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        public readonly string Uuid;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetExternalresourceResult(
            int category,

            string comments,

            string id,

            string @interface,

            string interfaceSelectMethod,

            string name,

            string password,

            int refreshRate,

            string resource,

            string serverIdentityCheck,

            string sourceIp,

            string status,

            string type,

            string updateMethod,

            string userAgent,

            string username,

            string uuid,

            string? vdomparam)
        {
            Category = category;
            Comments = comments;
            Id = id;
            Interface = @interface;
            InterfaceSelectMethod = interfaceSelectMethod;
            Name = name;
            Password = password;
            RefreshRate = refreshRate;
            Resource = resource;
            ServerIdentityCheck = serverIdentityCheck;
            SourceIp = sourceIp;
            Status = status;
            Type = type;
            UpdateMethod = updateMethod;
            UserAgent = userAgent;
            Username = username;
            Uuid = uuid;
            Vdomparam = vdomparam;
        }
    }
}
