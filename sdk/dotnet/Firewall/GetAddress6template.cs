// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    public static class GetAddress6template
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall address6template
        /// </summary>
        public static Task<GetAddress6templateResult> InvokeAsync(GetAddress6templateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAddress6templateResult>("fortios:firewall/getAddress6template:getAddress6template", args ?? new GetAddress6templateArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall address6template
        /// </summary>
        public static Output<GetAddress6templateResult> Invoke(GetAddress6templateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAddress6templateResult>("fortios:firewall/getAddress6template:getAddress6template", args ?? new GetAddress6templateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAddress6templateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall address6template.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetAddress6templateArgs()
        {
        }
        public static new GetAddress6templateArgs Empty => new GetAddress6templateArgs();
    }

    public sealed class GetAddress6templateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall address6template.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetAddress6templateInvokeArgs()
        {
        }
        public static new GetAddress6templateInvokeArgs Empty => new GetAddress6templateInvokeArgs();
    }


    [OutputType]
    public sealed class GetAddress6templateResult
    {
        /// <summary>
        /// Security Fabric global object setting.
        /// </summary>
        public readonly string FabricObject;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IPv6 address prefix.
        /// </summary>
        public readonly string Ip6;
        /// <summary>
        /// Subnet segment value name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number of IPv6 subnet segments.
        /// </summary>
        public readonly int SubnetSegmentCount;
        /// <summary>
        /// IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAddress6templateSubnetSegmentResult> SubnetSegments;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetAddress6templateResult(
            string fabricObject,

            string id,

            string ip6,

            string name,

            int subnetSegmentCount,

            ImmutableArray<Outputs.GetAddress6templateSubnetSegmentResult> subnetSegments,

            string? vdomparam)
        {
            FabricObject = fabricObject;
            Id = id;
            Ip6 = ip6;
            Name = name;
            SubnetSegmentCount = subnetSegmentCount;
            SubnetSegments = subnetSegments;
            Vdomparam = vdomparam;
        }
    }
}
