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

    public sealed class WtpprofileLbsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable AeroScout Real Time Location Service (RTLS) support. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("aeroscout")]
        public Input<string>? Aeroscout { get; set; }

        /// <summary>
        /// Use BSSID or board MAC address as AP MAC address in AeroScout AP messages (default = bssid). Valid values: `bssid`, `board-mac`.
        /// </summary>
        [Input("aeroscoutApMac")]
        public Input<string>? AeroscoutApMac { get; set; }

        /// <summary>
        /// Enable/disable compounded AeroScout tag and MU report (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("aeroscoutMmuReport")]
        public Input<string>? AeroscoutMmuReport { get; set; }

        /// <summary>
        /// Enable/disable AeroScout Mobile Unit (MU) support (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("aeroscoutMu")]
        public Input<string>? AeroscoutMu { get; set; }

        /// <summary>
        /// eroScout MU mode dilution factor (default = 20).
        /// </summary>
        [Input("aeroscoutMuFactor")]
        public Input<int>? AeroscoutMuFactor { get; set; }

        /// <summary>
        /// AeroScout MU mode timeout (0 - 65535 sec, default = 5).
        /// </summary>
        [Input("aeroscoutMuTimeout")]
        public Input<int>? AeroscoutMuTimeout { get; set; }

        /// <summary>
        /// IP address of AeroScout server.
        /// </summary>
        [Input("aeroscoutServerIp")]
        public Input<string>? AeroscoutServerIp { get; set; }

        /// <summary>
        /// AeroScout server UDP listening port.
        /// </summary>
        [Input("aeroscoutServerPort")]
        public Input<int>? AeroscoutServerPort { get; set; }

        /// <summary>
        /// Enable/disable Ekahau blink mode (now known as AiRISTA Flow) to track and locate WiFi tags (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ekahauBlinkMode")]
        public Input<string>? EkahauBlinkMode { get; set; }

        /// <summary>
        /// WiFi frame MAC address or WiFi Tag.
        /// </summary>
        [Input("ekahauTag")]
        public Input<string>? EkahauTag { get; set; }

        /// <summary>
        /// IP address of Ekahau RTLS Controller (ERC).
        /// </summary>
        [Input("ercServerIp")]
        public Input<string>? ErcServerIp { get; set; }

        /// <summary>
        /// Ekahau RTLS Controller (ERC) UDP listening port.
        /// </summary>
        [Input("ercServerPort")]
        public Input<int>? ErcServerPort { get; set; }

        /// <summary>
        /// Enable/disable FortiPresence to monitor the location and activity of WiFi clients even if they don't connect to this WiFi network (default = disable). Valid values: `foreign`, `both`, `disable`.
        /// </summary>
        [Input("fortipresence")]
        public Input<string>? Fortipresence { get; set; }

        /// <summary>
        /// Enable/disable FortiPresence finding and reporting BLE devices. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fortipresenceBle")]
        public Input<string>? FortipresenceBle { get; set; }

        /// <summary>
        /// FortiPresence report transmit frequency (5 - 65535 sec, default = 30).
        /// </summary>
        [Input("fortipresenceFrequency")]
        public Input<int>? FortipresenceFrequency { get; set; }

        /// <summary>
        /// UDP listening port of FortiPresence server (default = 3000).
        /// </summary>
        [Input("fortipresencePort")]
        public Input<int>? FortipresencePort { get; set; }

        /// <summary>
        /// FortiPresence project name (max. 16 characters, default = fortipresence).
        /// </summary>
        [Input("fortipresenceProject")]
        public Input<string>? FortipresenceProject { get; set; }

        /// <summary>
        /// Enable/disable FortiPresence finding and reporting rogue APs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fortipresenceRogue")]
        public Input<string>? FortipresenceRogue { get; set; }

        [Input("fortipresenceSecret")]
        private Input<string>? _fortipresenceSecret;

        /// <summary>
        /// FortiPresence secret password (max. 16 characters).
        /// </summary>
        public Input<string>? FortipresenceSecret
        {
            get => _fortipresenceSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _fortipresenceSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// FortiPresence server IP address.
        /// </summary>
        [Input("fortipresenceServer")]
        public Input<string>? FortipresenceServer { get; set; }

        /// <summary>
        /// FortiPresence server address type (default = ipv4). Valid values: `ipv4`, `fqdn`.
        /// </summary>
        [Input("fortipresenceServerAddrType")]
        public Input<string>? FortipresenceServerAddrType { get; set; }

        /// <summary>
        /// FQDN of FortiPresence server.
        /// </summary>
        [Input("fortipresenceServerFqdn")]
        public Input<string>? FortipresenceServerFqdn { get; set; }

        /// <summary>
        /// Enable/disable FortiPresence finding and reporting unassociated stations. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fortipresenceUnassoc")]
        public Input<string>? FortipresenceUnassoc { get; set; }

        /// <summary>
        /// Enable/disable PoleStar BLE NAO Track Real Time Location Service (RTLS) support (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("polestar")]
        public Input<string>? Polestar { get; set; }

        /// <summary>
        /// Time that measurements should be accumulated in seconds (default = 2).
        /// </summary>
        [Input("polestarAccumulationInterval")]
        public Input<int>? PolestarAccumulationInterval { get; set; }

        /// <summary>
        /// Tags and asset addrgrp list to be reported.
        /// </summary>
        [Input("polestarAssetAddrgrpList")]
        public Input<string>? PolestarAssetAddrgrpList { get; set; }

        /// <summary>
        /// Tags and asset UUID list 1 to be reported (string in the format of 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX').
        /// </summary>
        [Input("polestarAssetUuidList1")]
        public Input<string>? PolestarAssetUuidList1 { get; set; }

        /// <summary>
        /// Tags and asset UUID list 2 to be reported (string in the format of 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX').
        /// </summary>
        [Input("polestarAssetUuidList2")]
        public Input<string>? PolestarAssetUuidList2 { get; set; }

        /// <summary>
        /// Tags and asset UUID list 3 to be reported (string in the format of 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX').
        /// </summary>
        [Input("polestarAssetUuidList3")]
        public Input<string>? PolestarAssetUuidList3 { get; set; }

        /// <summary>
        /// Tags and asset UUID list 4 to be reported (string in the format of 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX').
        /// </summary>
        [Input("polestarAssetUuidList4")]
        public Input<string>? PolestarAssetUuidList4 { get; set; }

        /// <summary>
        /// Select the protocol to report Measurements, Advertising Data, or Location Data to NAO Cloud. (default = WSS). Valid values: `WSS`.
        /// </summary>
        [Input("polestarProtocol")]
        public Input<string>? PolestarProtocol { get; set; }

        /// <summary>
        /// Time between reporting accumulated measurements in seconds (default = 2).
        /// </summary>
        [Input("polestarReportingInterval")]
        public Input<int>? PolestarReportingInterval { get; set; }

        /// <summary>
        /// FQDN of PoleStar Nao Track Server (default = ws.nao-cloud.com).
        /// </summary>
        [Input("polestarServerFqdn")]
        public Input<string>? PolestarServerFqdn { get; set; }

        /// <summary>
        /// Path of PoleStar Nao Track Server (default = /v1/token/&lt;access_token&gt;/pst-v2).
        /// </summary>
        [Input("polestarServerPath")]
        public Input<string>? PolestarServerPath { get; set; }

        /// <summary>
        /// Port of PoleStar Nao Track Server (default = 443).
        /// </summary>
        [Input("polestarServerPort")]
        public Input<int>? PolestarServerPort { get; set; }

        /// <summary>
        /// Access Token of PoleStar Nao Track Server.
        /// </summary>
        [Input("polestarServerToken")]
        public Input<string>? PolestarServerToken { get; set; }

        /// <summary>
        /// Enable/disable client station locating services for all clients, whether associated or not (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("stationLocate")]
        public Input<string>? StationLocate { get; set; }

        public WtpprofileLbsGetArgs()
        {
        }
        public static new WtpprofileLbsGetArgs Empty => new WtpprofileLbsGetArgs();
    }
}
