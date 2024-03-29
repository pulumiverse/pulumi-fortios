// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Outputs
{

    [OutputType]
    public sealed class ManagedswitchVlan
    {
        /// <summary>
        /// 802.1x Radius (Tunnel-Private-Group-Id) VLANID assign-by-name priority. A smaller value has a higher priority.
        /// </summary>
        public readonly int? AssignmentPriority;
        /// <summary>
        /// VLAN name.
        /// </summary>
        public readonly string? VlanName;

        [OutputConstructor]
        private ManagedswitchVlan(
            int? assignmentPriority,

            string? vlanName)
        {
            AssignmentPriority = assignmentPriority;
            VlanName = vlanName;
        }
    }
}
