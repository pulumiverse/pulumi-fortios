// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Inputs
{

    public sealed class WtpRadio2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
        /// </summary>
        [Input("autoPowerHigh")]
        public Input<int>? AutoPowerHigh { get; set; }

        /// <summary>
        /// Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoPowerLevel")]
        public Input<string>? AutoPowerLevel { get; set; }

        /// <summary>
        /// The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
        /// </summary>
        [Input("autoPowerLow")]
        public Input<int>? AutoPowerLow { get; set; }

        /// <summary>
        /// The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
        /// </summary>
        [Input("autoPowerTarget")]
        public Input<string>? AutoPowerTarget { get; set; }

        /// <summary>
        /// WiFi band that Radio 4 operates on.
        /// </summary>
        [Input("band")]
        public Input<string>? Band { get; set; }

        [Input("channels")]
        private InputList<Inputs.WtpRadio2ChannelArgs>? _channels;

        /// <summary>
        /// Selected list of wireless radio channels. The structure of `channel` block is documented below.
        /// </summary>
        public InputList<Inputs.WtpRadio2ChannelArgs> Channels
        {
            get => _channels ?? (_channels = new InputList<Inputs.WtpRadio2ChannelArgs>());
            set => _channels = value;
        }

        /// <summary>
        /// Radio mode to be used for DRMA manual mode (default = ncf). Valid values: `ap`, `monitor`, `ncf`, `ncf-peek`.
        /// </summary>
        [Input("drmaManualMode")]
        public Input<string>? DrmaManualMode { get; set; }

        /// <summary>
        /// Enable to override the WTP profile spectrum analysis configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("overrideAnalysis")]
        public Input<string>? OverrideAnalysis { get; set; }

        /// <summary>
        /// Enable to override the WTP profile band setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("overrideBand")]
        public Input<string>? OverrideBand { get; set; }

        /// <summary>
        /// Enable to override WTP profile channel settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("overrideChannel")]
        public Input<string>? OverrideChannel { get; set; }

        /// <summary>
        /// Enable to override the WTP profile power level configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("overrideTxpower")]
        public Input<string>? OverrideTxpower { get; set; }

        /// <summary>
        /// Enable to override WTP profile Virtual Access Point (VAP) settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("overrideVaps")]
        public Input<string>? OverrideVaps { get; set; }

        /// <summary>
        /// Radio power level as a percentage of the maximum transmit power (0 - 100, default = 100).
        /// </summary>
        [Input("powerLevel")]
        public Input<int>? PowerLevel { get; set; }

        /// <summary>
        /// Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm`, `percentage`.
        /// </summary>
        [Input("powerMode")]
        public Input<string>? PowerMode { get; set; }

        /// <summary>
        /// Radio EIRP power in dBm (1 - 33, default = 27).
        /// </summary>
        [Input("powerValue")]
        public Input<int>? PowerValue { get; set; }

        /// <summary>
        /// radio-id
        /// </summary>
        [Input("radioId")]
        public Input<int>? RadioId { get; set; }

        /// <summary>
        /// Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
        /// </summary>
        [Input("spectrumAnalysis")]
        public Input<string>? SpectrumAnalysis { get; set; }

        /// <summary>
        /// Enable/disable the automatic inheritance of all Virtual Access Points (VAPs) (default = enable).
        /// </summary>
        [Input("vapAll")]
        public Input<string>? VapAll { get; set; }

        [Input("vaps")]
        private InputList<Inputs.WtpRadio2VapArgs>? _vaps;

        /// <summary>
        /// Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.
        /// </summary>
        public InputList<Inputs.WtpRadio2VapArgs> Vaps
        {
            get => _vaps ?? (_vaps = new InputList<Inputs.WtpRadio2VapArgs>());
            set => _vaps = value;
        }

        public WtpRadio2Args()
        {
        }
        public static new WtpRadio2Args Empty => new WtpRadio2Args();
    }
}
