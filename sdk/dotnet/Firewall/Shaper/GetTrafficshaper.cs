// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Shaper
{
    public static class GetTrafficshaper
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewallshaper trafficshaper
        /// </summary>
        public static Task<GetTrafficshaperResult> InvokeAsync(GetTrafficshaperArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTrafficshaperResult>("fortios:firewall/shaper/getTrafficshaper:getTrafficshaper", args ?? new GetTrafficshaperArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewallshaper trafficshaper
        /// </summary>
        public static Output<GetTrafficshaperResult> Invoke(GetTrafficshaperInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrafficshaperResult>("fortios:firewall/shaper/getTrafficshaper:getTrafficshaper", args ?? new GetTrafficshaperInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTrafficshaperArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewallshaper trafficshaper.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetTrafficshaperArgs()
        {
        }
        public static new GetTrafficshaperArgs Empty => new GetTrafficshaperArgs();
    }

    public sealed class GetTrafficshaperInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewallshaper trafficshaper.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetTrafficshaperInvokeArgs()
        {
        }
        public static new GetTrafficshaperInvokeArgs Empty => new GetTrafficshaperInvokeArgs();
    }


    [OutputType]
    public sealed class GetTrafficshaperResult
    {
        /// <summary>
        /// Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps).
        /// </summary>
        public readonly string BandwidthUnit;
        /// <summary>
        /// VLAN CoS mark.
        /// </summary>
        public readonly string Cos;
        /// <summary>
        /// Enable/disable VLAN CoS marking.
        /// </summary>
        public readonly string CosMarking;
        /// <summary>
        /// Select VLAN CoS marking method.
        /// </summary>
        public readonly string CosMarkingMethod;
        /// <summary>
        /// Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper.
        /// </summary>
        public readonly string Diffserv;
        /// <summary>
        /// DiffServ setting to be applied to traffic accepted by this shaper.
        /// </summary>
        public readonly string Diffservcode;
        /// <summary>
        /// Select DSCP marking method.
        /// </summary>
        public readonly string DscpMarkingMethod;
        /// <summary>
        /// Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
        /// </summary>
        public readonly int ExceedBandwidth;
        /// <summary>
        /// Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].
        /// </summary>
        public readonly int ExceedClassId;
        /// <summary>
        /// VLAN CoS mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
        /// </summary>
        public readonly string ExceedCos;
        /// <summary>
        /// DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
        /// </summary>
        public readonly string ExceedDscp;
        /// <summary>
        /// Amount of bandwidth guaranteed for this shaper (0 - 16776000). Units depend on the bandwidth-unit setting.
        /// </summary>
        public readonly int GuaranteedBandwidth;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
        /// </summary>
        public readonly int MaximumBandwidth;
        /// <summary>
        /// VLAN CoS mark for traffic in [exceed-bandwidth, maximum-bandwidth].
        /// </summary>
        public readonly string MaximumCos;
        /// <summary>
        /// DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
        /// </summary>
        public readonly string MaximumDscp;
        /// <summary>
        /// Traffic shaper name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Per-packet size overhead used in rate computations.
        /// </summary>
        public readonly int Overhead;
        /// <summary>
        /// Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy.
        /// </summary>
        public readonly string PerPolicy;
        /// <summary>
        /// Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth.
        /// </summary>
        public readonly string Priority;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetTrafficshaperResult(
            string bandwidthUnit,

            string cos,

            string cosMarking,

            string cosMarkingMethod,

            string diffserv,

            string diffservcode,

            string dscpMarkingMethod,

            int exceedBandwidth,

            int exceedClassId,

            string exceedCos,

            string exceedDscp,

            int guaranteedBandwidth,

            string id,

            int maximumBandwidth,

            string maximumCos,

            string maximumDscp,

            string name,

            int overhead,

            string perPolicy,

            string priority,

            string? vdomparam)
        {
            BandwidthUnit = bandwidthUnit;
            Cos = cos;
            CosMarking = cosMarking;
            CosMarkingMethod = cosMarkingMethod;
            Diffserv = diffserv;
            Diffservcode = diffservcode;
            DscpMarkingMethod = dscpMarkingMethod;
            ExceedBandwidth = exceedBandwidth;
            ExceedClassId = exceedClassId;
            ExceedCos = exceedCos;
            ExceedDscp = exceedDscp;
            GuaranteedBandwidth = guaranteedBandwidth;
            Id = id;
            MaximumBandwidth = maximumBandwidth;
            MaximumCos = maximumCos;
            MaximumDscp = maximumDscp;
            Name = name;
            Overhead = overhead;
            PerPolicy = perPolicy;
            Priority = priority;
            Vdomparam = vdomparam;
        }
    }
}
