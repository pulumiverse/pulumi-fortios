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
    public sealed class WtpRadio1
    {
        /// <summary>
        /// The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
        /// </summary>
        public readonly int? AutoPowerHigh;
        /// <summary>
        /// Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? AutoPowerLevel;
        /// <summary>
        /// The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
        /// </summary>
        public readonly int? AutoPowerLow;
        /// <summary>
        /// The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
        /// </summary>
        public readonly string? AutoPowerTarget;
        /// <summary>
        /// WiFi band that Radio 4 operates on.
        /// </summary>
        public readonly string? Band;
        /// <summary>
        /// Selected list of wireless radio channels. The structure of `channel` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.WtpRadio1Channel> Channels;
        /// <summary>
        /// Radio mode to be used for DRMA manual mode (default = ncf). Valid values: `ap`, `monitor`, `ncf`, `ncf-peek`.
        /// </summary>
        public readonly string? DrmaManualMode;
        /// <summary>
        /// Enable to override the WTP profile spectrum analysis configuration. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? OverrideAnalysis;
        /// <summary>
        /// Enable to override the WTP profile band setting. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? OverrideBand;
        /// <summary>
        /// Enable to override WTP profile channel settings. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? OverrideChannel;
        /// <summary>
        /// Enable to override the WTP profile power level configuration. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? OverrideTxpower;
        /// <summary>
        /// Enable to override WTP profile Virtual Access Point (VAP) settings. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? OverrideVaps;
        /// <summary>
        /// Radio power level as a percentage of the maximum transmit power (0 - 100, default = 100).
        /// </summary>
        public readonly int? PowerLevel;
        /// <summary>
        /// Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm`, `percentage`.
        /// </summary>
        public readonly string? PowerMode;
        /// <summary>
        /// Radio EIRP power in dBm (1 - 33, default = 27).
        /// </summary>
        public readonly int? PowerValue;
        /// <summary>
        /// radio-id
        /// </summary>
        public readonly int? RadioId;
        /// <summary>
        /// Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
        /// </summary>
        public readonly string? SpectrumAnalysis;
        /// <summary>
        /// Enable/disable the automatic inheritance of all Virtual Access Points (VAPs) (default = enable).
        /// </summary>
        public readonly string? VapAll;
        /// <summary>
        /// Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.WtpRadio1Vap> Vaps;

        [OutputConstructor]
        private WtpRadio1(
            int? autoPowerHigh,

            string? autoPowerLevel,

            int? autoPowerLow,

            string? autoPowerTarget,

            string? band,

            ImmutableArray<Outputs.WtpRadio1Channel> channels,

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

            ImmutableArray<Outputs.WtpRadio1Vap> vaps)
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
