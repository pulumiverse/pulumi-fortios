// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Outputs
{

    [OutputType]
    public sealed class WtpRadio2
    {
        public readonly int? AutoPowerHigh;
        public readonly string? AutoPowerLevel;
        public readonly int? AutoPowerLow;
        public readonly string? AutoPowerTarget;
        public readonly string? Band;
        public readonly ImmutableArray<Outputs.WtpRadio2Channel> Channels;
        public readonly string? DrmaManualMode;
        public readonly string? OverrideAnalysis;
        public readonly string? OverrideBand;
        public readonly string? OverrideChannel;
        public readonly string? OverrideTxpower;
        public readonly string? OverrideVaps;
        public readonly int? PowerLevel;
        public readonly string? PowerMode;
        public readonly int? PowerValue;
        public readonly int? RadioId;
        public readonly string? SpectrumAnalysis;
        public readonly string? VapAll;
        public readonly ImmutableArray<Outputs.WtpRadio2Vap> Vaps;

        [OutputConstructor]
        private WtpRadio2(
            int? autoPowerHigh,

            string? autoPowerLevel,

            int? autoPowerLow,

            string? autoPowerTarget,

            string? band,

            ImmutableArray<Outputs.WtpRadio2Channel> channels,

            string? drmaManualMode,

            string? overrideAnalysis,

            string? overrideBand,

            string? overrideChannel,

            string? overrideTxpower,

            string? overrideVaps,

            int? powerLevel,

            string? powerMode,

            int? powerValue,

            int? radioId,

            string? spectrumAnalysis,

            string? vapAll,

            ImmutableArray<Outputs.WtpRadio2Vap> vaps)
        {
            AutoPowerHigh = autoPowerHigh;
            AutoPowerLevel = autoPowerLevel;
            AutoPowerLow = autoPowerLow;
            AutoPowerTarget = autoPowerTarget;
            Band = band;
            Channels = channels;
            DrmaManualMode = drmaManualMode;
            OverrideAnalysis = overrideAnalysis;
            OverrideBand = overrideBand;
            OverrideChannel = overrideChannel;
            OverrideTxpower = overrideTxpower;
            OverrideVaps = overrideVaps;
            PowerLevel = powerLevel;
            PowerMode = powerMode;
            PowerValue = powerValue;
            RadioId = radioId;
            SpectrumAnalysis = spectrumAnalysis;
            VapAll = vapAll;
            Vaps = vaps;
        }
    }
}
