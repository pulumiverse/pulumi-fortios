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

    public sealed class WtpprofileRadio3Args : global::Pulumi.ResourceArgs
    {
        [Input("airtimeFairness")]
        public Input<string>? AirtimeFairness { get; set; }

        [Input("amsdu")]
        public Input<string>? Amsdu { get; set; }

        /// <summary>
        /// Enable/disable AP handoff of clients to other APs (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("apHandoff")]
        public Input<string>? ApHandoff { get; set; }

        [Input("apSnifferAddr")]
        public Input<string>? ApSnifferAddr { get; set; }

        [Input("apSnifferBufsize")]
        public Input<int>? ApSnifferBufsize { get; set; }

        [Input("apSnifferChan")]
        public Input<int>? ApSnifferChan { get; set; }

        [Input("apSnifferCtl")]
        public Input<string>? ApSnifferCtl { get; set; }

        [Input("apSnifferData")]
        public Input<string>? ApSnifferData { get; set; }

        [Input("apSnifferMgmtBeacon")]
        public Input<string>? ApSnifferMgmtBeacon { get; set; }

        [Input("apSnifferMgmtOther")]
        public Input<string>? ApSnifferMgmtOther { get; set; }

        [Input("apSnifferMgmtProbe")]
        public Input<string>? ApSnifferMgmtProbe { get; set; }

        [Input("arrpProfile")]
        public Input<string>? ArrpProfile { get; set; }

        [Input("autoPowerHigh")]
        public Input<int>? AutoPowerHigh { get; set; }

        [Input("autoPowerLevel")]
        public Input<string>? AutoPowerLevel { get; set; }

        [Input("autoPowerLow")]
        public Input<int>? AutoPowerLow { get; set; }

        [Input("autoPowerTarget")]
        public Input<string>? AutoPowerTarget { get; set; }

        [Input("band")]
        public Input<string>? Band { get; set; }

        [Input("band5gType")]
        public Input<string>? Band5gType { get; set; }

        [Input("bandwidthAdmissionControl")]
        public Input<string>? BandwidthAdmissionControl { get; set; }

        [Input("bandwidthCapacity")]
        public Input<int>? BandwidthCapacity { get; set; }

        [Input("beaconInterval")]
        public Input<int>? BeaconInterval { get; set; }

        [Input("bssColor")]
        public Input<int>? BssColor { get; set; }

        [Input("bssColorMode")]
        public Input<string>? BssColorMode { get; set; }

        [Input("callAdmissionControl")]
        public Input<string>? CallAdmissionControl { get; set; }

        [Input("callCapacity")]
        public Input<int>? CallCapacity { get; set; }

        [Input("channelBonding")]
        public Input<string>? ChannelBonding { get; set; }

        [Input("channelBondingExt")]
        public Input<string>? ChannelBondingExt { get; set; }

        [Input("channelUtilization")]
        public Input<string>? ChannelUtilization { get; set; }

        [Input("channels")]
        private InputList<Inputs.WtpprofileRadio3ChannelArgs>? _channels;
        public InputList<Inputs.WtpprofileRadio3ChannelArgs> Channels
        {
            get => _channels ?? (_channels = new InputList<Inputs.WtpprofileRadio3ChannelArgs>());
            set => _channels = value;
        }

        [Input("coexistence")]
        public Input<string>? Coexistence { get; set; }

        [Input("darrp")]
        public Input<string>? Darrp { get; set; }

        [Input("drma")]
        public Input<string>? Drma { get; set; }

        [Input("drmaSensitivity")]
        public Input<string>? DrmaSensitivity { get; set; }

        [Input("dtim")]
        public Input<int>? Dtim { get; set; }

        [Input("fragThreshold")]
        public Input<int>? FragThreshold { get; set; }

        /// <summary>
        /// Enable/disable frequency handoff of clients to other channels (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("frequencyHandoff")]
        public Input<string>? FrequencyHandoff { get; set; }

        [Input("iperfProtocol")]
        public Input<string>? IperfProtocol { get; set; }

        [Input("iperfServerPort")]
        public Input<int>? IperfServerPort { get; set; }

        /// <summary>
        /// Maximum number of stations (STAs) supported by the WTP (default = 0, meaning no client limitation).
        /// </summary>
        [Input("maxClients")]
        public Input<int>? MaxClients { get; set; }

        [Input("maxDistance")]
        public Input<int>? MaxDistance { get; set; }

        [Input("mimoMode")]
        public Input<string>? MimoMode { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("n80211d")]
        public Input<string>? N80211d { get; set; }

        [Input("optionalAntenna")]
        public Input<string>? OptionalAntenna { get; set; }

        [Input("optionalAntennaGain")]
        public Input<string>? OptionalAntennaGain { get; set; }

        [Input("powerLevel")]
        public Input<int>? PowerLevel { get; set; }

        [Input("powerMode")]
        public Input<string>? PowerMode { get; set; }

        [Input("powerValue")]
        public Input<int>? PowerValue { get; set; }

        [Input("powersaveOptimize")]
        public Input<string>? PowersaveOptimize { get; set; }

        [Input("protectionMode")]
        public Input<string>? ProtectionMode { get; set; }

        [Input("rtsThreshold")]
        public Input<int>? RtsThreshold { get; set; }

        [Input("samBssid")]
        public Input<string>? SamBssid { get; set; }

        [Input("samCaCertificate")]
        public Input<string>? SamCaCertificate { get; set; }

        [Input("samCaptivePortal")]
        public Input<string>? SamCaptivePortal { get; set; }

        [Input("samClientCertificate")]
        public Input<string>? SamClientCertificate { get; set; }

        [Input("samCwpFailureString")]
        public Input<string>? SamCwpFailureString { get; set; }

        [Input("samCwpMatchString")]
        public Input<string>? SamCwpMatchString { get; set; }

        [Input("samCwpPassword")]
        public Input<string>? SamCwpPassword { get; set; }

        [Input("samCwpSuccessString")]
        public Input<string>? SamCwpSuccessString { get; set; }

        [Input("samCwpTestUrl")]
        public Input<string>? SamCwpTestUrl { get; set; }

        [Input("samCwpUsername")]
        public Input<string>? SamCwpUsername { get; set; }

        [Input("samEapMethod")]
        public Input<string>? SamEapMethod { get; set; }

        [Input("samPassword")]
        public Input<string>? SamPassword { get; set; }

        [Input("samPrivateKey")]
        public Input<string>? SamPrivateKey { get; set; }

        [Input("samPrivateKeyPassword")]
        public Input<string>? SamPrivateKeyPassword { get; set; }

        [Input("samReportIntv")]
        public Input<int>? SamReportIntv { get; set; }

        [Input("samSecurityType")]
        public Input<string>? SamSecurityType { get; set; }

        [Input("samServerFqdn")]
        public Input<string>? SamServerFqdn { get; set; }

        [Input("samServerIp")]
        public Input<string>? SamServerIp { get; set; }

        [Input("samServerType")]
        public Input<string>? SamServerType { get; set; }

        [Input("samSsid")]
        public Input<string>? SamSsid { get; set; }

        [Input("samTest")]
        public Input<string>? SamTest { get; set; }

        [Input("samUsername")]
        public Input<string>? SamUsername { get; set; }

        [Input("shortGuardInterval")]
        public Input<string>? ShortGuardInterval { get; set; }

        [Input("spectrumAnalysis")]
        public Input<string>? SpectrumAnalysis { get; set; }

        [Input("transmitOptimize")]
        public Input<string>? TransmitOptimize { get; set; }

        [Input("vapAll")]
        public Input<string>? VapAll { get; set; }

        [Input("vaps")]
        private InputList<Inputs.WtpprofileRadio3VapArgs>? _vaps;
        public InputList<Inputs.WtpprofileRadio3VapArgs> Vaps
        {
            get => _vaps ?? (_vaps = new InputList<Inputs.WtpprofileRadio3VapArgs>());
            set => _vaps = value;
        }

        [Input("widsProfile")]
        public Input<string>? WidsProfile { get; set; }

        [Input("zeroWaitDfs")]
        public Input<string>? ZeroWaitDfs { get; set; }

        public WtpprofileRadio3Args()
        {
        }
        public static new WtpprofileRadio3Args Empty => new WtpprofileRadio3Args();
    }
}
