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
    public static class GetDnsserver
    {
        /// <summary>
        /// Use this data source to get information on an fortios system dnsserver
        /// </summary>
        public static Task<GetDnsserverResult> InvokeAsync(GetDnsserverArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDnsserverResult>("fortios:system/getDnsserver:getDnsserver", args ?? new GetDnsserverArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system dnsserver
        /// </summary>
        public static Output<GetDnsserverResult> Invoke(GetDnsserverInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDnsserverResult>("fortios:system/getDnsserver:getDnsserver", args ?? new GetDnsserverInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDnsserverArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system dnsserver.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetDnsserverArgs()
        {
        }
        public static new GetDnsserverArgs Empty => new GetDnsserverArgs();
    }

    public sealed class GetDnsserverInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system dnsserver.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetDnsserverInvokeArgs()
        {
        }
        public static new GetDnsserverInvokeArgs Empty => new GetDnsserverInvokeArgs();
    }


    [OutputType]
    public sealed class GetDnsserverResult
    {
        /// <summary>
        /// DNS filter profile.
        /// </summary>
        public readonly string DnsfilterProfile;
        /// <summary>
        /// DNS over HTTPS.
        /// </summary>
        public readonly string Doh;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// DNS server mode.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// DNS server name.
        /// </summary>
        public readonly string Name;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetDnsserverResult(
            string dnsfilterProfile,

            string doh,

            string id,

            string mode,

            string name,

            string? vdomparam)
        {
            DnsfilterProfile = dnsfilterProfile;
            Doh = doh;
            Id = id;
            Mode = mode;
            Name = name;
            Vdomparam = vdomparam;
        }
    }
}