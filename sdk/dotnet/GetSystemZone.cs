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
    public static class GetSystemZone
    {
        /// <summary>
        /// Use this data source to get information on an fortios system zone
        /// </summary>
        public static Task<GetSystemZoneResult> InvokeAsync(GetSystemZoneArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemZoneResult>("fortios:index/getSystemZone:getSystemZone", args ?? new GetSystemZoneArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system zone
        /// </summary>
        public static Output<GetSystemZoneResult> Invoke(GetSystemZoneInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemZoneResult>("fortios:index/getSystemZone:getSystemZone", args ?? new GetSystemZoneInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemZoneArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system zone.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemZoneArgs()
        {
        }
        public static new GetSystemZoneArgs Empty => new GetSystemZoneArgs();
    }

    public sealed class GetSystemZoneInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system zone.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemZoneInvokeArgs()
        {
        }
        public static new GetSystemZoneInvokeArgs Empty => new GetSystemZoneInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemZoneResult
    {
        /// <summary>
        /// Description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemZoneInterfaceResult> Interfaces;
        /// <summary>
        /// Allow or deny traffic routing between different interfaces in the same zone (default = deny).
        /// </summary>
        public readonly string Intrazone;
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemZoneTaggingResult> Taggings;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemZoneResult(
            string description,

            string id,

            ImmutableArray<Outputs.GetSystemZoneInterfaceResult> interfaces,

            string intrazone,

            string name,

            ImmutableArray<Outputs.GetSystemZoneTaggingResult> taggings,

            string? vdomparam)
        {
            Description = description;
            Id = id;
            Interfaces = interfaces;
            Intrazone = intrazone;
            Name = name;
            Taggings = taggings;
            Vdomparam = vdomparam;
        }
    }
}