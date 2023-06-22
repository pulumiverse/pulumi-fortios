// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class SystemlldpNetworkpolicySoftphone
    {
        /// <summary>
        /// Differentiated Services Code Point (DSCP) value to advertise.
        /// </summary>
        public readonly int? Dscp;
        /// <summary>
        /// 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// Enable/disable advertising this policy. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Advertise tagged or untagged traffic. Valid values: `none`, `dot1q`, `dot1p`.
        /// </summary>
        public readonly string? Tag;
        /// <summary>
        /// 802.1Q VLAN ID to advertise (1 - 4094).
        /// </summary>
        public readonly int? Vlan;

        [OutputConstructor]
        private SystemlldpNetworkpolicySoftphone(
            int? dscp,

            int? priority,

            string? status,

            string? tag,

            int? vlan)
        {
            Dscp = dscp;
            Priority = priority;
            Status = status;
            Tag = tag;
            Vlan = vlan;
        }
    }
}
