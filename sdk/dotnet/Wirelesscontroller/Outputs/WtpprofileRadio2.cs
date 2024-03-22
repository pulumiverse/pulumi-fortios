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
    public sealed class WtpprofileRadio2
    {
        /// <summary>
        /// Enable/disable airtime fairness (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? AirtimeFairness;
        /// <summary>
        /// Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Amsdu;
        /// <summary>
        /// Enable/disable AP handoff of clients to other APs (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ApHandoff;
        /// <summary>
        /// MAC address to monitor.
        /// </summary>
        public readonly string? ApSnifferAddr;
        /// <summary>
        /// Sniffer buffer size (1 - 32 MB, default = 16).
        /// </summary>
        public readonly int? ApSnifferBufsize;
        /// <summary>
        /// Channel on which to operate the sniffer (default = 6).
        /// </summary>
        public readonly int? ApSnifferChan;
        /// <summary>
        /// Enable/disable sniffer on WiFi control frame (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ApSnifferCtl;
        /// <summary>
        /// Enable/disable sniffer on WiFi data frame (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ApSnifferData;
        /// <summary>
        /// Enable/disable sniffer on WiFi management Beacon frames (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ApSnifferMgmtBeacon;
        /// <summary>
        /// Enable/disable sniffer on WiFi management other frames  (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ApSnifferMgmtOther;
        /// <summary>
        /// Enable/disable sniffer on WiFi management probe frames (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ApSnifferMgmtProbe;
        /// <summary>
        /// Distributed Automatic Radio Resource Provisioning (DARRP) profile name to assign to the radio.
        /// </summary>
        public readonly string? ArrpProfile;
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
        /// WiFi band that Radio 3 operates on.
        /// </summary>
        public readonly string? Band;
        /// <summary>
        /// WiFi 5G band type. Valid values: `5g-full`, `5g-high`, `5g-low`.
        /// </summary>
        public readonly string? Band5gType;
        /// <summary>
        /// Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? BandwidthAdmissionControl;
        /// <summary>
        /// Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
        /// </summary>
        public readonly int? BandwidthCapacity;
        /// <summary>
        /// Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).
        /// </summary>
        public readonly int? BeaconInterval;
        /// <summary>
        /// BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).
        /// </summary>
        public readonly int? BssColor;
        /// <summary>
        /// BSS color mode for this 11ax radio (default = auto). Valid values: `auto`, `static`.
        /// </summary>
        public readonly string? BssColorMode;
        /// <summary>
        /// Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? CallAdmissionControl;
        /// <summary>
        /// Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
        /// </summary>
        public readonly int? CallCapacity;
        /// <summary>
        /// Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence. Valid values: `160MHz`, `80MHz`, `40MHz`, `20MHz`.
        /// </summary>
        public readonly string? ChannelBonding;
        /// <summary>
        /// Enable/disable measuring channel utilization. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ChannelUtilization;
        /// <summary>
        /// Selected list of wireless radio channels. The structure of `channel` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.WtpprofileRadio2Channel> Channels;
        /// <summary>
        /// Enable/disable allowing both HT20 and HT40 on the same radio (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Coexistence;
        /// <summary>
        /// Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Darrp;
        /// <summary>
        /// Enable/disable dynamic radio mode assignment (DRMA) (default = disable). Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Drma;
        /// <summary>
        /// Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low). Valid values: `low`, `medium`, `high`.
        /// </summary>
        public readonly string? DrmaSensitivity;
        /// <summary>
        /// Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
        /// </summary>
        public readonly int? Dtim;
        /// <summary>
        /// Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
        /// </summary>
        public readonly int? FragThreshold;
        /// <summary>
        /// Enable/disable frequency handoff of clients to other channels (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? FrequencyHandoff;
        /// <summary>
        /// Iperf test protocol (default = "UDP"). Valid values: `udp`, `tcp`.
        /// </summary>
        public readonly string? IperfProtocol;
        /// <summary>
        /// Iperf service port number.
        /// </summary>
        public readonly int? IperfServerPort;
        /// <summary>
        /// Maximum number of stations (STAs) supported by the WTP (default = 0, meaning no client limitation).
        /// </summary>
        public readonly int? MaxClients;
        /// <summary>
        /// Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
        /// </summary>
        public readonly int? MaxDistance;
        /// <summary>
        /// Configure radio MIMO mode (default = default). Valid values: `default`, `1x1`, `2x2`, `3x3`, `4x4`, `8x8`.
        /// </summary>
        public readonly string? MimoMode;
        /// <summary>
        /// Mode of radio 3. Radio 3 can be disabled, configured as an access point, a rogue AP monitor, or a sniffer.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// Enable/disable 802.11d countryie(default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? N80211d;
        /// <summary>
        /// Optional antenna used on FAP (default = none).
        /// </summary>
        public readonly string? OptionalAntenna;
        /// <summary>
        /// Optional antenna gain in dBi (0 to 20, default = 0).
        /// </summary>
        public readonly string? OptionalAntennaGain;
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
        /// Enable client power-saving features such as TIM, AC VO, and OBSS etc. Valid values: `tim`, `ac-vo`, `no-obss-scan`, `no-11b-rate`, `client-rate-follow`.
        /// </summary>
        public readonly string? PowersaveOptimize;
        /// <summary>
        /// Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable). Valid values: `rtscts`, `ctsonly`, `disable`.
        /// </summary>
        public readonly string? ProtectionMode;
        /// <summary>
        /// radio-id
        /// </summary>
        public readonly int? RadioId;
        /// <summary>
        /// Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
        /// </summary>
        public readonly int? RtsThreshold;
        /// <summary>
        /// BSSID for WiFi network.
        /// </summary>
        public readonly string? SamBssid;
        /// <summary>
        /// CA certificate for WPA2/WPA3-ENTERPRISE.
        /// </summary>
        public readonly string? SamCaCertificate;
        /// <summary>
        /// Enable/disable Captive Portal Authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? SamCaptivePortal;
        /// <summary>
        /// Client certificate for WPA2/WPA3-ENTERPRISE.
        /// </summary>
        public readonly string? SamClientCertificate;
        /// <summary>
        /// Failure identification on the page after an incorrect login.
        /// </summary>
        public readonly string? SamCwpFailureString;
        /// <summary>
        /// Identification string from the captive portal login form.
        /// </summary>
        public readonly string? SamCwpMatchString;
        /// <summary>
        /// Password for captive portal authentication.
        /// </summary>
        public readonly string? SamCwpPassword;
        /// <summary>
        /// Success identification on the page after a successful login.
        /// </summary>
        public readonly string? SamCwpSuccessString;
        /// <summary>
        /// Website the client is trying to access.
        /// </summary>
        public readonly string? SamCwpTestUrl;
        /// <summary>
        /// Username for captive portal authentication.
        /// </summary>
        public readonly string? SamCwpUsername;
        /// <summary>
        /// Select WPA2/WPA3-ENTERPRISE EAP Method (default = PEAP). Valid values: `both`, `tls`, `peap`.
        /// </summary>
        public readonly string? SamEapMethod;
        /// <summary>
        /// Passphrase for WiFi network connection.
        /// </summary>
        public readonly string? SamPassword;
        /// <summary>
        /// Private key for WPA2/WPA3-ENTERPRISE.
        /// </summary>
        public readonly string? SamPrivateKey;
        /// <summary>
        /// Password for private key file for WPA2/WPA3-ENTERPRISE.
        /// </summary>
        public readonly string? SamPrivateKeyPassword;
        /// <summary>
        /// SAM report interval (sec), 0 for a one-time report.
        /// </summary>
        public readonly int? SamReportIntv;
        /// <summary>
        /// Select WiFi network security type (default = "wpa-personal").
        /// </summary>
        public readonly string? SamSecurityType;
        /// <summary>
        /// SAM test server domain name.
        /// </summary>
        public readonly string? SamServerFqdn;
        /// <summary>
        /// SAM test server IP address.
        /// </summary>
        public readonly string? SamServerIp;
        /// <summary>
        /// Select SAM server type (default = "IP"). Valid values: `ip`, `fqdn`.
        /// </summary>
        public readonly string? SamServerType;
        /// <summary>
        /// SSID for WiFi network.
        /// </summary>
        public readonly string? SamSsid;
        /// <summary>
        /// Select SAM test type (default = "PING"). Valid values: `ping`, `iperf`.
        /// </summary>
        public readonly string? SamTest;
        /// <summary>
        /// Username for WiFi network connection.
        /// </summary>
        public readonly string? SamUsername;
        /// <summary>
        /// Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ShortGuardInterval;
        /// <summary>
        /// Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
        /// </summary>
        public readonly string? SpectrumAnalysis;
        /// <summary>
        /// Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default. Valid values: `disable`, `power-save`, `aggr-limit`, `retry-limit`, `send-bar`.
        /// </summary>
        public readonly string? TransmitOptimize;
        /// <summary>
        /// Enable/disable the automatic inheritance of all Virtual Access Points (VAPs) (default = enable).
        /// </summary>
        public readonly string? VapAll;
        /// <summary>
        /// Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.WtpprofileRadio2Vap> Vaps;
        /// <summary>
        /// Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.
        /// </summary>
        public readonly string? WidsProfile;
        /// <summary>
        /// Enable/disable zero wait DFS on radio (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ZeroWaitDfs;

        [OutputConstructor]
        private WtpprofileRadio2(
            string? airtimeFairness,

            string? amsdu,

            string? apHandoff,

            string? apSnifferAddr,

            int? apSnifferBufsize,

            int? apSnifferChan,

            string? apSnifferCtl,

            string? apSnifferData,

            string? apSnifferMgmtBeacon,

            string? apSnifferMgmtOther,

            string? apSnifferMgmtProbe,

            string? arrpProfile,

            int? autoPowerHigh,

            string? autoPowerLevel,

            int? autoPowerLow,

            string? autoPowerTarget,

            string? band,

            string? band5gType,

            string? bandwidthAdmissionControl,

            int? bandwidthCapacity,

            int? beaconInterval,

            int? bssColor,

            string? bssColorMode,

            string? callAdmissionControl,

            int? callCapacity,

            string? channelBonding,

            string? channelUtilization,

            ImmutableArray<Outputs.WtpprofileRadio2Channel> channels,

            string? coexistence,

            string? darrp,

            string? drma,

            string? drmaSensitivity,

            int? dtim,

            int? fragThreshold,

            string? frequencyHandoff,

            string? iperfProtocol,

            int? iperfServerPort,

            int? maxClients,

            int? maxDistance,

            string? mimoMode,

            string? mode,

            string? n80211d,

            string? optionalAntenna,

            string? optionalAntennaGain,

            int? powerLevel,

            string? powerMode,

            int? powerValue,

            string? powersaveOptimize,

            string? protectionMode,

            int? radioId,

            int? rtsThreshold,

            string? samBssid,

            string? samCaCertificate,

            string? samCaptivePortal,

            string? samClientCertificate,

            string? samCwpFailureString,

            string? samCwpMatchString,

            string? samCwpPassword,

            string? samCwpSuccessString,

            string? samCwpTestUrl,

            string? samCwpUsername,

            string? samEapMethod,

            string? samPassword,

            string? samPrivateKey,

            string? samPrivateKeyPassword,

            int? samReportIntv,

            string? samSecurityType,

            string? samServerFqdn,

            string? samServerIp,

            string? samServerType,

            string? samSsid,

            string? samTest,

            string? samUsername,

            string? shortGuardInterval,

            string? spectrumAnalysis,

            string? transmitOptimize,

            string? vapAll,

            ImmutableArray<Outputs.WtpprofileRadio2Vap> vaps,

            string? widsProfile,

            string? zeroWaitDfs)
        {
            AirtimeFairness = airtimeFairness;
            Amsdu = amsdu;
            ApHandoff = apHandoff;
            ApSnifferAddr = apSnifferAddr;
            ApSnifferBufsize = apSnifferBufsize;
            ApSnifferChan = apSnifferChan;
            ApSnifferCtl = apSnifferCtl;
            ApSnifferData = apSnifferData;
            ApSnifferMgmtBeacon = apSnifferMgmtBeacon;
            ApSnifferMgmtOther = apSnifferMgmtOther;
            ApSnifferMgmtProbe = apSnifferMgmtProbe;
            ArrpProfile = arrpProfile;
            AutoPowerHigh = autoPowerHigh;
            AutoPowerLevel = autoPowerLevel;
            AutoPowerLow = autoPowerLow;
            AutoPowerTarget = autoPowerTarget;
            Band = band;
            Band5gType = band5gType;
            BandwidthAdmissionControl = bandwidthAdmissionControl;
            BandwidthCapacity = bandwidthCapacity;
            BeaconInterval = beaconInterval;
            BssColor = bssColor;
            BssColorMode = bssColorMode;
            CallAdmissionControl = callAdmissionControl;
            CallCapacity = callCapacity;
            ChannelBonding = channelBonding;
            ChannelUtilization = channelUtilization;
            Channels = channels;
            Coexistence = coexistence;
            Darrp = darrp;
            Drma = drma;
            DrmaSensitivity = drmaSensitivity;
            Dtim = dtim;
            FragThreshold = fragThreshold;
            FrequencyHandoff = frequencyHandoff;
            IperfProtocol = iperfProtocol;
            IperfServerPort = iperfServerPort;
            MaxClients = maxClients;
            MaxDistance = maxDistance;
            MimoMode = mimoMode;
            Mode = mode;
            N80211d = n80211d;
            OptionalAntenna = optionalAntenna;
            OptionalAntennaGain = optionalAntennaGain;
            PowerLevel = powerLevel;
            PowerMode = powerMode;
            PowerValue = powerValue;
            PowersaveOptimize = powersaveOptimize;
            ProtectionMode = protectionMode;
            RadioId = radioId;
            RtsThreshold = rtsThreshold;
            SamBssid = samBssid;
            SamCaCertificate = samCaCertificate;
            SamCaptivePortal = samCaptivePortal;
            SamClientCertificate = samClientCertificate;
            SamCwpFailureString = samCwpFailureString;
            SamCwpMatchString = samCwpMatchString;
            SamCwpPassword = samCwpPassword;
            SamCwpSuccessString = samCwpSuccessString;
            SamCwpTestUrl = samCwpTestUrl;
            SamCwpUsername = samCwpUsername;
            SamEapMethod = samEapMethod;
            SamPassword = samPassword;
            SamPrivateKey = samPrivateKey;
            SamPrivateKeyPassword = samPrivateKeyPassword;
            SamReportIntv = samReportIntv;
            SamSecurityType = samSecurityType;
            SamServerFqdn = samServerFqdn;
            SamServerIp = samServerIp;
            SamServerType = samServerType;
            SamSsid = samSsid;
            SamTest = samTest;
            SamUsername = samUsername;
            ShortGuardInterval = shortGuardInterval;
            SpectrumAnalysis = spectrumAnalysis;
            TransmitOptimize = transmitOptimize;
            VapAll = vapAll;
            Vaps = vaps;
            WidsProfile = widsProfile;
            ZeroWaitDfs = zeroWaitDfs;
        }
    }
}
