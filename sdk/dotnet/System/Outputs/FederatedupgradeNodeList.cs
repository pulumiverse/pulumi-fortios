// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class FederatedupgradeNodeList
    {
        /// <summary>
        /// The serial of the FortiGate that controls this device
        /// </summary>
        public readonly string? CoordinatingFortigate;
        /// <summary>
        /// What type of device this node represents.
        /// </summary>
        public readonly string? DeviceType;
        /// <summary>
        /// Serial number of the node to include.
        /// </summary>
        public readonly string? Serial;
        /// <summary>
        /// When the upgrade was configured. Format hh:mm yyyy/mm/dd UTC.
        /// </summary>
        public readonly string? SetupTime;
        /// <summary>
        /// Scheduled time for the upgrade. Format hh:mm yyyy/mm/dd UTC.
        /// </summary>
        public readonly string? Time;
        /// <summary>
        /// Whether the upgrade should be run immediately, or at a scheduled time. Valid values: `immediate`, `scheduled`.
        /// </summary>
        public readonly string? Timing;
        /// <summary>
        /// Image IDs to upgrade through.
        /// </summary>
        public readonly string? UpgradePath;

        [OutputConstructor]
        private FederatedupgradeNodeList(
            string? coordinatingFortigate,

            string? deviceType,

            string? serial,

            string? setupTime,

            string? time,

            string? timing,

            string? upgradePath)
        {
            CoordinatingFortigate = coordinatingFortigate;
            DeviceType = deviceType;
            Serial = serial;
            SetupTime = setupTime;
            Time = time;
            Timing = timing;
            UpgradePath = upgradePath;
        }
    }
}
