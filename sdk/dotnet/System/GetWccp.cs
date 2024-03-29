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
    public static class GetWccp
    {
        /// <summary>
        /// Use this data source to get information on an fortios system wccp
        /// </summary>
        public static Task<GetWccpResult> InvokeAsync(GetWccpArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWccpResult>("fortios:system/getWccp:getWccp", args ?? new GetWccpArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system wccp
        /// </summary>
        public static Output<GetWccpResult> Invoke(GetWccpInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWccpResult>("fortios:system/getWccp:getWccp", args ?? new GetWccpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWccpArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the service_id of the desired system wccp.
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetWccpArgs()
        {
        }
        public static new GetWccpArgs Empty => new GetWccpArgs();
    }

    public sealed class GetWccpInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the service_id of the desired system wccp.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetWccpInvokeArgs()
        {
        }
        public static new GetWccpInvokeArgs Empty => new GetWccpInvokeArgs();
    }


    [OutputType]
    public sealed class GetWccpResult
    {
        /// <summary>
        /// Assignment bucket format for the WCCP cache engine.
        /// </summary>
        public readonly string AssignmentBucketFormat;
        /// <summary>
        /// Assignment destination address mask.
        /// </summary>
        public readonly string AssignmentDstaddrMask;
        /// <summary>
        /// Hash key assignment preference.
        /// </summary>
        public readonly string AssignmentMethod;
        /// <summary>
        /// Assignment source address mask.
        /// </summary>
        public readonly string AssignmentSrcaddrMask;
        /// <summary>
        /// Assignment of hash weight/ratio for the WCCP cache engine.
        /// </summary>
        public readonly int AssignmentWeight;
        /// <summary>
        /// Enable/disable MD5 authentication.
        /// </summary>
        public readonly string Authentication;
        /// <summary>
        /// Method used to forward traffic to the routers or to return to the cache engine.
        /// </summary>
        public readonly string CacheEngineMethod;
        /// <summary>
        /// IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
        /// </summary>
        public readonly string CacheId;
        /// <summary>
        /// Method used to forward traffic to the cache servers.
        /// </summary>
        public readonly string ForwardMethod;
        /// <summary>
        /// IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
        /// </summary>
        public readonly string GroupAddress;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Password for MD5 authentication.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Service ports.
        /// </summary>
        public readonly string Ports;
        /// <summary>
        /// Match method.
        /// </summary>
        public readonly string PortsDefined;
        /// <summary>
        /// Hash method.
        /// </summary>
        public readonly string PrimaryHash;
        /// <summary>
        /// Service priority.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// Service protocol.
        /// </summary>
        public readonly int Protocol;
        /// <summary>
        /// Method used to decline a redirected packet and return it to the FortiGate.
        /// </summary>
        public readonly string ReturnMethod;
        /// <summary>
        /// IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
        /// </summary>
        public readonly string RouterId;
        /// <summary>
        /// IP addresses of one or more WCCP routers.
        /// </summary>
        public readonly string RouterList;
        /// <summary>
        /// IP addresses and netmasks for up to four cache servers.
        /// </summary>
        public readonly string ServerList;
        /// <summary>
        /// Cache server type.
        /// </summary>
        public readonly string ServerType;
        /// <summary>
        /// Service ID.
        /// </summary>
        public readonly string ServiceId;
        /// <summary>
        /// WCCP service type used by the cache server for logical interception and redirection of traffic.
        /// </summary>
        public readonly string ServiceType;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetWccpResult(
            string assignmentBucketFormat,

            string assignmentDstaddrMask,

            string assignmentMethod,

            string assignmentSrcaddrMask,

            int assignmentWeight,

            string authentication,

            string cacheEngineMethod,

            string cacheId,

            string forwardMethod,

            string groupAddress,

            string id,

            string password,

            string ports,

            string portsDefined,

            string primaryHash,

            int priority,

            int protocol,

            string returnMethod,

            string routerId,

            string routerList,

            string serverList,

            string serverType,

            string serviceId,

            string serviceType,

            string? vdomparam)
        {
            AssignmentBucketFormat = assignmentBucketFormat;
            AssignmentDstaddrMask = assignmentDstaddrMask;
            AssignmentMethod = assignmentMethod;
            AssignmentSrcaddrMask = assignmentSrcaddrMask;
            AssignmentWeight = assignmentWeight;
            Authentication = authentication;
            CacheEngineMethod = cacheEngineMethod;
            CacheId = cacheId;
            ForwardMethod = forwardMethod;
            GroupAddress = groupAddress;
            Id = id;
            Password = password;
            Ports = ports;
            PortsDefined = portsDefined;
            PrimaryHash = primaryHash;
            Priority = priority;
            Protocol = protocol;
            ReturnMethod = returnMethod;
            RouterId = routerId;
            RouterList = routerList;
            ServerList = serverList;
            ServerType = serverType;
            ServiceId = serviceId;
            ServiceType = serviceType;
            Vdomparam = vdomparam;
        }
    }
}
