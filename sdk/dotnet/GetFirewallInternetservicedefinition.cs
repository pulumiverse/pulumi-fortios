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
    public static class GetFirewallInternetservicedefinition
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall internetservicedefinition
        /// </summary>
        public static Task<GetFirewallInternetservicedefinitionResult> InvokeAsync(GetFirewallInternetservicedefinitionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallInternetservicedefinitionResult>("fortios:index/getFirewallInternetservicedefinition:getFirewallInternetservicedefinition", args ?? new GetFirewallInternetservicedefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall internetservicedefinition
        /// </summary>
        public static Output<GetFirewallInternetservicedefinitionResult> Invoke(GetFirewallInternetservicedefinitionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallInternetservicedefinitionResult>("fortios:index/getFirewallInternetservicedefinition:getFirewallInternetservicedefinition", args ?? new GetFirewallInternetservicedefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallInternetservicedefinitionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the fosid of the desired firewall internetservicedefinition.
        /// </summary>
        [Input("fosid", required: true)]
        public int Fosid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallInternetservicedefinitionArgs()
        {
        }
        public static new GetFirewallInternetservicedefinitionArgs Empty => new GetFirewallInternetservicedefinitionArgs();
    }

    public sealed class GetFirewallInternetservicedefinitionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the fosid of the desired firewall internetservicedefinition.
        /// </summary>
        [Input("fosid", required: true)]
        public Input<int> Fosid { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallInternetservicedefinitionInvokeArgs()
        {
        }
        public static new GetFirewallInternetservicedefinitionInvokeArgs Empty => new GetFirewallInternetservicedefinitionInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallInternetservicedefinitionResult
    {
        /// <summary>
        /// Protocol and port information in an Internet Service entry. The structure of `entry` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallInternetservicedefinitionEntryResult> Entries;
        /// <summary>
        /// Internet Service application list ID.
        /// </summary>
        public readonly int Fosid;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallInternetservicedefinitionResult(
            ImmutableArray<Outputs.GetFirewallInternetservicedefinitionEntryResult> entries,

            int fosid,

            string id,

            string? vdomparam)
        {
            Entries = entries;
            Fosid = fosid;
            Id = id;
            Vdomparam = vdomparam;
        }
    }
}
