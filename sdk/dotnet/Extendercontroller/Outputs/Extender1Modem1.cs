// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Extendercontroller.Outputs
{

    [OutputType]
    public sealed class Extender1Modem1
    {
        public readonly Outputs.Extender1Modem1AutoSwitch? AutoSwitch;
        public readonly int? ConnStatus;
        public readonly string? DefaultSim;
        public readonly string? Gps;
        public readonly string? Ifname;
        public readonly string? PreferredCarrier;
        public readonly string? RedundantIntf;
        public readonly string? RedundantMode;
        public readonly string? Sim1Pin;
        public readonly string? Sim1PinCode;
        public readonly string? Sim2Pin;
        public readonly string? Sim2PinCode;

        [OutputConstructor]
        private Extender1Modem1(
            Outputs.Extender1Modem1AutoSwitch? autoSwitch,

            int? connStatus,

            string? defaultSim,

            string? gps,

            string? ifname,

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
            Ifname = ifname;
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
