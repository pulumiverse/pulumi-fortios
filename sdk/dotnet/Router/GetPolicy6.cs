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
    public static class GetPolicy6
    {
        /// <summary>
        /// Use this data source to get information on an fortios router policy6
        /// </summary>
        public static Task<GetPolicy6Result> InvokeAsync(GetPolicy6Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPolicy6Result>("fortios:router/getPolicy6:getPolicy6", args ?? new GetPolicy6Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios router policy6
        /// </summary>
        public static Output<GetPolicy6Result> Invoke(GetPolicy6InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPolicy6Result>("fortios:router/getPolicy6:getPolicy6", args ?? new GetPolicy6InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPolicy6Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the seq_num of the desired router policy6.
        /// </summary>
        [Input("seqNum", required: true)]
        public int SeqNum { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetPolicy6Args()
        {
        }
        public static new GetPolicy6Args Empty => new GetPolicy6Args();
    }

    public sealed class GetPolicy6InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the seq_num of the desired router policy6.
        /// </summary>
        [Input("seqNum", required: true)]
        public Input<int> SeqNum { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetPolicy6InvokeArgs()
        {
        }
        public static new GetPolicy6InvokeArgs Empty => new GetPolicy6InvokeArgs();
    }


    [OutputType]
    public sealed class GetPolicy6Result
    {
        /// <summary>
        /// Action of the policy route.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Optional comments.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// Destination IPv6 prefix.
        /// </summary>
        public readonly string Dst;
        /// <summary>
        /// Enable/disable negating destination address match.
        /// </summary>
        public readonly string DstNegate;
        /// <summary>
        /// Destination address name. The structure of `dstaddr` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPolicy6DstaddrResult> Dstaddrs;
        /// <summary>
        /// End destination port number (1 - 65535).
        /// </summary>
        public readonly int EndPort;
        /// <summary>
        /// End source port number (1 - 65535).
        /// </summary>
        public readonly int EndSourcePort;
        /// <summary>
        /// IPv6 address of the gateway.
        /// </summary>
        public readonly string Gateway;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Incoming interface name.
        /// </summary>
        public readonly string InputDevice;
        /// <summary>
        /// Enable/disable negation of input device match.
        /// </summary>
        public readonly string InputDeviceNegate;
        /// <summary>
        /// Custom Destination Internet Service name. The structure of `internet_service_custom` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPolicy6InternetServiceCustomResult> InternetServiceCustoms;
        /// <summary>
        /// Destination Internet Service ID. The structure of `internet_service_id` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPolicy6InternetServiceIdResult> InternetServiceIds;
        /// <summary>
        /// Outgoing interface name.
        /// </summary>
        public readonly string OutputDevice;
        /// <summary>
        /// Protocol number (0 - 255).
        /// </summary>
        public readonly int Protocol;
        /// <summary>
        /// Sequence number.
        /// </summary>
        public readonly int SeqNum;
        /// <summary>
        /// Source IPv6 prefix.
        /// </summary>
        public readonly string Src;
        /// <summary>
        /// Enable/disable negating source address match.
        /// </summary>
        public readonly string SrcNegate;
        /// <summary>
        /// Source address name. The structure of `srcaddr` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPolicy6SrcaddrResult> Srcaddrs;
        /// <summary>
        /// Start destination port number (1 - 65535).
        /// </summary>
        public readonly int StartPort;
        /// <summary>
        /// Start source port number (1 - 65535).
        /// </summary>
        public readonly int StartSourcePort;
        /// <summary>
        /// Enable/disable this policy route.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Type of service bit pattern.
        /// </summary>
        public readonly string Tos;
        /// <summary>
        /// Type of service evaluated bits.
        /// </summary>
        public readonly string TosMask;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetPolicy6Result(
            string action,

            string comments,

            string dst,

            string dstNegate,

            ImmutableArray<Outputs.GetPolicy6DstaddrResult> dstaddrs,

            int endPort,

            int endSourcePort,

            string gateway,

            string id,

            string inputDevice,

            string inputDeviceNegate,

            ImmutableArray<Outputs.GetPolicy6InternetServiceCustomResult> internetServiceCustoms,

            ImmutableArray<Outputs.GetPolicy6InternetServiceIdResult> internetServiceIds,

            string outputDevice,

            int protocol,

            int seqNum,

            string src,

            string srcNegate,

            ImmutableArray<Outputs.GetPolicy6SrcaddrResult> srcaddrs,

            int startPort,

            int startSourcePort,

            string status,

            string tos,

            string tosMask,

            string? vdomparam)
        {
            Action = action;
            Comments = comments;
            Dst = dst;
            DstNegate = dstNegate;
            Dstaddrs = dstaddrs;
            EndPort = endPort;
            EndSourcePort = endSourcePort;
            Gateway = gateway;
            Id = id;
            InputDevice = inputDevice;
            InputDeviceNegate = inputDeviceNegate;
            InternetServiceCustoms = internetServiceCustoms;
            InternetServiceIds = internetServiceIds;
            OutputDevice = outputDevice;
            Protocol = protocol;
            SeqNum = seqNum;
            Src = src;
            SrcNegate = srcNegate;
            Srcaddrs = srcaddrs;
            StartPort = startPort;
            StartSourcePort = startSourcePort;
            Status = status;
            Tos = tos;
            TosMask = tosMask;
            Vdomparam = vdomparam;
        }
    }
}
