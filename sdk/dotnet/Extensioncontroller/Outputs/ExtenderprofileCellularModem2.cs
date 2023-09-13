// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Extensioncontroller.Outputs
{

    [OutputType]
    public sealed class ExtenderprofileCellularModem2
    {
        /// <summary>
        /// FortiExtender auto switch configuration. The structure of `auto_switch` block is documented below.
        /// </summary>
        public readonly Outputs.ExtenderprofileCellularModem2AutoSwitch? AutoSwitch;
        /// <summary>
        /// Connection status.
        /// </summary>
        public readonly int? ConnStatus;
        /// <summary>
        /// Default SIM selection. Valid values: `sim1`, `sim2`, `carrier`, `cost`.
        /// </summary>
        public readonly string? DefaultSim;
        /// <summary>
        /// FortiExtender GPS enable/disable. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Gps;
        /// <summary>
        /// Preferred carrier.
        /// </summary>
        public readonly string? PreferredCarrier;
        /// <summary>
        /// Redundant interface.
        /// </summary>
        public readonly string? RedundantIntf;
        /// <summary>
        /// FortiExtender mode. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? RedundantMode;
        /// <summary>
        /// SIM #1 PIN status. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Sim1Pin;
        /// <summary>
        /// SIM #1 PIN password.
        /// </summary>
        public readonly string? Sim1PinCode;
        /// <summary>
        /// SIM #2 PIN status. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Sim2Pin;
        /// <summary>
        /// SIM #2 PIN password.
        /// </summary>
        public readonly string? Sim2PinCode;

        [OutputConstructor]
        private ExtenderprofileCellularModem2(
            Outputs.ExtenderprofileCellularModem2AutoSwitch? autoSwitch,

            int? connStatus,

            string? defaultSim,

            string? gps,

            string? preferredCarrier,

            string? redundantIntf,

            string? redundantMode,

            string? sim1Pin,

            string? sim1PinCode,

            string? sim2Pin,

            string? sim2PinCode)
        {
            AutoSwitch = autoSwitch;
            ConnStatus = connStatus;
            DefaultSim = defaultSim;
            Gps = gps;
            PreferredCarrier = preferredCarrier;
            RedundantIntf = redundantIntf;
            RedundantMode = redundantMode;
            Sim1Pin = sim1Pin;
            Sim1PinCode = sim1PinCode;
            Sim2Pin = sim2Pin;
            Sim2PinCode = sim2PinCode;
        }
    }
}
