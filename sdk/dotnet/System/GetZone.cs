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
    public static class GetZone
    {
        /// <summary>
        /// Use this data source to get information on an fortios system zone
        /// </summary>
        public static Task<GetZoneResult> InvokeAsync(GetZoneArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetZoneResult>("fortios:system/getZone:getZone", args ?? new GetZoneArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system zone
        /// </summary>
        public static Output<GetZoneResult> Invoke(GetZoneInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetZoneResult>("fortios:system/getZone:getZone", args ?? new GetZoneInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZoneArgs : global::Pulumi.InvokeArgs
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

        public GetZoneArgs()
        {
        }
        public static new GetZoneArgs Empty => new GetZoneArgs();
    }

    public sealed class GetZoneInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetZoneInvokeArgs()
        {
        }
        public static new GetZoneInvokeArgs Empty => new GetZoneInvokeArgs();
    }


    [OutputType]
    public sealed class GetZoneResult
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
        public readonly ImmutableArray<Outputs.GetZoneInterfaceResult> Interfaces;
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
        public readonly ImmutableArray<Outputs.GetZoneTaggingResult> Taggings;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetZoneResult(
            string description,

            string id,

            ImmutableArray<Outputs.GetZoneInterfaceResult> interfaces,

            string intrazone,

            string name,

            ImmutableArray<Outputs.GetZoneTaggingResult> taggings,

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
