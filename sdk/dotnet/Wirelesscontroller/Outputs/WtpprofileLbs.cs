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
    public sealed class WtpprofileLbs
    {
        /// <summary>
        /// Enable/disable AeroScout Real Time Location Service (RTLS) support. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Aeroscout;
        /// <summary>
        /// Use BSSID or board MAC address as AP MAC address in the Aeroscout AP message. Valid values: `bssid`, `board-mac`.
        /// </summary>
        public readonly string? AeroscoutApMac;
        /// <summary>
        /// Enable/disable MU compounded report. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? AeroscoutMmuReport;
        /// <summary>
        /// Enable/disable AeroScout support. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? AeroscoutMu;
        /// <summary>
        /// AeroScout Mobile Unit (MU) mode dilution factor (default = 20).
        /// </summary>
        public readonly int? AeroscoutMuFactor;
        /// <summary>
        /// AeroScout MU mode timeout (0 - 65535 sec, default = 5).
        /// </summary>
        public readonly int? AeroscoutMuTimeout;
        /// <summary>
        /// IP address of AeroScout server.
        /// </summary>
        public readonly string? AeroscoutServerIp;
        /// <summary>
        /// AeroScout server UDP listening port.
        /// </summary>
        public readonly int? AeroscoutServerPort;
        /// <summary>
        /// Enable/disable Ekahua blink mode (also called AiRISTA Flow Blink Mode) to find the location of devices connected to a wireless LAN (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? EkahauBlinkMode;
        /// <summary>
        /// WiFi frame MAC address or WiFi Tag.
        /// </summary>
        public readonly string? EkahauTag;
        /// <summary>
        /// IP address of Ekahua RTLS Controller (ERC).
        /// </summary>
        public readonly string? ErcServerIp;
        /// <summary>
        /// Ekahua RTLS Controller (ERC) UDP listening port.
        /// </summary>
        public readonly int? ErcServerPort;
        /// <summary>
        /// Enable/disable FortiPresence to monitor the location and activity of WiFi clients even if they don't connect to this WiFi network (default = disable). Valid values: `foreign`, `both`, `disable`.
        /// </summary>
        public readonly string? Fortipresence;
        /// <summary>
        /// Enable/disable FortiPresence finding and reporting BLE devices. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? FortipresenceBle;
        /// <summary>
        /// FortiPresence report transmit frequency (5 - 65535 sec, default = 30).
        /// </summary>
        public readonly int? FortipresenceFrequency;
        /// <summary>
        /// FortiPresence server UDP listening port (default = 3000).
        /// </summary>
        public readonly int? FortipresencePort;
        /// <summary>
        /// FortiPresence project name (max. 16 characters, default = fortipresence).
        /// </summary>
        public readonly string? FortipresenceProject;
        /// <summary>
        /// Enable/disable FortiPresence finding and reporting rogue APs. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? FortipresenceRogue;
        /// <summary>
        /// FortiPresence secret password (max. 16 characters).
        /// </summary>
        public readonly string? FortipresenceSecret;
        /// <summary>
        /// FortiPresence server IP address.
        /// </summary>
        public readonly string? FortipresenceServer;
        /// <summary>
        /// FortiPresence server address type (default = ipv4). Valid values: `ipv4`, `fqdn`.
        /// </summary>
        public readonly string? FortipresenceServerAddrType;
        /// <summary>
        /// FQDN of FortiPresence server.
        /// </summary>
        public readonly string? FortipresenceServerFqdn;
        /// <summary>
        /// Enable/disable FortiPresence finding and reporting unassociated stations. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? FortipresenceUnassoc;
        /// <summary>
        /// Enable/disable PoleStar BLE NAO Track Real Time Location Service (RTLS) support (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Polestar;
        /// <summary>
        /// Time that measurements should be accumulated in seconds (default = 2).
        /// </summary>
        public readonly int? PolestarAccumulationInterval;
        /// <summary>
        /// Tags and asset addrgrp list to be reported.
        /// </summary>
        public readonly string? PolestarAssetAddrgrpList;
        /// <summary>
        /// Tags and asset UUID list 1 to be reported (string in the format of 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX').
        /// </summary>
        public readonly string? PolestarAssetUuidList1;
        /// <summary>
        /// Tags and asset UUID list 2 to be reported (string in the format of 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX').
        /// </summary>
        public readonly string? PolestarAssetUuidList2;
        /// <summary>
        /// Tags and asset UUID list 3 to be reported (string in the format of 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX').
        /// </summary>
        public readonly string? PolestarAssetUuidList3;
        /// <summary>
        /// Tags and asset UUID list 4 to be reported (string in the format of 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX').
        /// </summary>
        public readonly string? PolestarAssetUuidList4;
        /// <summary>
        /// Select the protocol to report Measurements, Advertising Data, or Location Data to NAO Cloud. (default = WSS). Valid values: `WSS`.
        /// </summary>
        public readonly string? PolestarProtocol;
        /// <summary>
        /// Time between reporting accumulated measurements in seconds (default = 2).
        /// </summary>
        public readonly int? PolestarReportingInterval;
        /// <summary>
        /// FQDN of PoleStar Nao Track Server (default = ws.nao-cloud.com).
        /// </summary>
        public readonly string? PolestarServerFqdn;
        /// <summary>
        /// Path of PoleStar Nao Track Server (default = /v1/token/&lt;access_token&gt;/pst-v2).
        /// </summary>
        public readonly string? PolestarServerPath;
        /// <summary>
        /// Port of PoleStar Nao Track Server (default = 443).
        /// </summary>
        public readonly int? PolestarServerPort;
        /// <summary>
        /// Access Token of PoleStar Nao Track Server.
        /// </summary>
        public readonly string? PolestarServerToken;
        /// <summary>
        /// Enable/disable client station locating services for all clients, whether associated or not (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? StationLocate;

        [OutputConstructor]
        private WtpprofileLbs(
            string? aeroscout,

            string? aeroscoutApMac,

            string? aeroscoutMmuReport,

            string? aeroscoutMu,

            int? aeroscoutMuFactor,

            int? aeroscoutMuTimeout,

            string? aeroscoutServerIp,

            int? aeroscoutServerPort,

            string? ekahauBlinkMode,

            string? ekahauTag,

            string? ercServerIp,

            int? ercServerPort,

            string? fortipresence,

            string? fortipresenceBle,

            int? fortipresenceFrequency,

            int? fortipresencePort,

            string? fortipresenceProject,

            string? fortipresenceRogue,

            string? fortipresenceSecret,

            string? fortipresenceServer,

            string? fortipresenceServerAddrType,

            string? fortipresenceServerFqdn,

            string? fortipresenceUnassoc,

            string? polestar,

            int? polestarAccumulationInterval,

            string? polestarAssetAddrgrpList,

            string? polestarAssetUuidList1,

            string? polestarAssetUuidList2,

            string? polestarAssetUuidList3,

            string? polestarAssetUuidList4,

            string? polestarProtocol,

            int? polestarReportingInterval,

            string? polestarServerFqdn,

            string? polestarServerPath,

            int? polestarServerPort,

            string? polestarServerToken,

            string? stationLocate)
        {
            Aeroscout = aeroscout;
            AeroscoutApMac = aeroscoutApMac;
            AeroscoutMmuReport = aeroscoutMmuReport;
            AeroscoutMu = aeroscoutMu;
            AeroscoutMuFactor = aeroscoutMuFactor;
            AeroscoutMuTimeout = aeroscoutMuTimeout;
            AeroscoutServerIp = aeroscoutServerIp;
            AeroscoutServerPort = aeroscoutServerPort;
            EkahauBlinkMode = ekahauBlinkMode;
            EkahauTag = ekahauTag;
            ErcServerIp = ercServerIp;
            ErcServerPort = ercServerPort;
            Fortipresence = fortipresence;
            FortipresenceBle = fortipresenceBle;
            FortipresenceFrequency = fortipresenceFrequency;
            FortipresencePort = fortipresencePort;
            FortipresenceProject = fortipresenceProject;
            FortipresenceRogue = fortipresenceRogue;
            FortipresenceSecret = fortipresenceSecret;
            FortipresenceServer = fortipresenceServer;
            FortipresenceServerAddrType = fortipresenceServerAddrType;
            FortipresenceServerFqdn = fortipresenceServerFqdn;
            FortipresenceUnassoc = fortipresenceUnassoc;
            Polestar = polestar;
            PolestarAccumulationInterval = polestarAccumulationInterval;
            PolestarAssetAddrgrpList = polestarAssetAddrgrpList;
            PolestarAssetUuidList1 = polestarAssetUuidList1;
            PolestarAssetUuidList2 = polestarAssetUuidList2;
            PolestarAssetUuidList3 = polestarAssetUuidList3;
            PolestarAssetUuidList4 = polestarAssetUuidList4;
            PolestarProtocol = polestarProtocol;
            PolestarReportingInterval = polestarReportingInterval;
            PolestarServerFqdn = polestarServerFqdn;
            PolestarServerPath = polestarServerPath;
            PolestarServerPort = polestarServerPort;
            PolestarServerToken = polestarServerToken;
            StationLocate = stationLocate;
        }
    }
}
