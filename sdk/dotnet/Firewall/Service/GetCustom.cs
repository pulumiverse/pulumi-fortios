// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Service
{
    public static class GetCustom
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewallservice custom
        /// </summary>
        public static Task<GetCustomResult> InvokeAsync(GetCustomArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCustomResult>("fortios:firewall/service/getCustom:getCustom", args ?? new GetCustomArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewallservice custom
        /// </summary>
        public static Output<GetCustomResult> Invoke(GetCustomInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomResult>("fortios:firewall/service/getCustom:getCustom", args ?? new GetCustomInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewallservice custom.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetCustomArgs()
        {
        }
        public static new GetCustomArgs Empty => new GetCustomArgs();
    }

    public sealed class GetCustomInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewallservice custom.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetCustomInvokeArgs()
        {
        }
        public static new GetCustomInvokeArgs Empty => new GetCustomInvokeArgs();
    }


    [OutputType]
    public sealed class GetCustomResult
    {
        /// <summary>
        /// Application category ID. The structure of `app_category` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCustomAppCategoryResult> AppCategories;
        /// <summary>
        /// Application service type.
        /// </summary>
        public readonly string AppServiceType;
        /// <summary>
        /// Application ID. The structure of `application` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCustomApplicationResult> Applications;
        /// <summary>
        /// Service category.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Configure the type of ICMP error message verification.
        /// </summary>
        public readonly string CheckResetRange;
        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        public readonly int Color;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// Security Fabric global object setting.
        /// </summary>
        public readonly string FabricObject;
        /// <summary>
        /// Fully qualified domain name.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// Helper name.
        /// </summary>
        public readonly string Helper;
        /// <summary>
        /// ICMP code.
        /// </summary>
        public readonly int Icmpcode;
        /// <summary>
        /// ICMP type.
        /// </summary>
        public readonly int Icmptype;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Start and end of the IP range associated with service.
        /// </summary>
        public readonly string Iprange;
        /// <summary>
        /// Custom service name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Protocol type based on IANA numbers.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// IP protocol number.
        /// </summary>
        public readonly int ProtocolNumber;
        /// <summary>
        /// Enable/disable web proxy service.
        /// </summary>
        public readonly string Proxy;
        /// <summary>
        /// Multiple SCTP port ranges.
        /// </summary>
        public readonly string SctpPortrange;
        /// <summary>
        /// Session TTL (300 - 604800, 0 = default).
        /// </summary>
        public readonly int SessionTtl;
        /// <summary>
        /// Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
        /// </summary>
        public readonly int TcpHalfcloseTimer;
        /// <summary>
        /// Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
        /// </summary>
        public readonly int TcpHalfopenTimer;
        /// <summary>
        /// Multiple TCP port ranges.
        /// </summary>
        public readonly string TcpPortrange;
        /// <summary>
        /// Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
        /// </summary>
        public readonly int TcpRstTimer;
        /// <summary>
        /// Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
        /// </summary>
        public readonly int TcpTimewaitTimer;
        /// <summary>
        /// UDP half close timeout (0 - 86400 sec, 0 = default).
        /// </summary>
        public readonly int UdpIdleTimer;
        /// <summary>
        /// Multiple UDP port ranges.
        /// </summary>
        public readonly string UdpPortrange;
        public readonly string? Vdomparam;
        /// <summary>
        /// Enable/disable the visibility of the service on the GUI.
        /// </summary>
        public readonly string Visibility;

        [OutputConstructor]
        private GetCustomResult(
            ImmutableArray<Outputs.GetCustomAppCategoryResult> appCategories,

            string appServiceType,

            ImmutableArray<Outputs.GetCustomApplicationResult> applications,

            string category,

            string checkResetRange,

            int color,

            string comment,

            string fabricObject,

            string fqdn,

            string helper,

            int icmpcode,

            int icmptype,

            string id,

            string iprange,

            string name,

            string protocol,

            int protocolNumber,

            string proxy,

            string sctpPortrange,

            int sessionTtl,

            int tcpHalfcloseTimer,

            int tcpHalfopenTimer,

            string tcpPortrange,

            int tcpRstTimer,

            int tcpTimewaitTimer,

            int udpIdleTimer,

            string udpPortrange,

            string? vdomparam,

            string visibility)
        {
            AppCategories = appCategories;
            AppServiceType = appServiceType;
            Applications = applications;
            Category = category;
            CheckResetRange = checkResetRange;
            Color = color;
            Comment = comment;
            FabricObject = fabricObject;
            Fqdn = fqdn;
            Helper = helper;
            Icmpcode = icmpcode;
            Icmptype = icmptype;
            Id = id;
            Iprange = iprange;
            Name = name;
            Protocol = protocol;
            ProtocolNumber = protocolNumber;
            Proxy = proxy;
            SctpPortrange = sctpPortrange;
            SessionTtl = sessionTtl;
            TcpHalfcloseTimer = tcpHalfcloseTimer;
            TcpHalfopenTimer = tcpHalfopenTimer;
            TcpPortrange = tcpPortrange;
            TcpRstTimer = tcpRstTimer;
            TcpTimewaitTimer = tcpTimewaitTimer;
            UdpIdleTimer = udpIdleTimer;
            UdpPortrange = udpPortrange;
            Vdomparam = vdomparam;
            Visibility = visibility;
        }
    }
}