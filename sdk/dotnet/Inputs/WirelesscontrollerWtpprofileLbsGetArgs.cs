// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class WirelesscontrollerWtpprofileLbsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable AeroScout Real Time Location Service (RTLS) support. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("aeroscout")]
        public Input<string>? Aeroscout { get; set; }

        /// <summary>
        /// Use BSSID or board MAC address as AP MAC address in the Aeroscout AP message. Valid values: `bssid`, `board-mac`.
        /// </summary>
        [Input("aeroscoutApMac")]
        public Input<string>? AeroscoutApMac { get; set; }

        /// <summary>
        /// Enable/disable MU compounded report. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("aeroscoutMmuReport")]
        public Input<string>? AeroscoutMmuReport { get; set; }

        /// <summary>
        /// Enable/disable AeroScout support. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("aeroscoutMu")]
        public Input<string>? AeroscoutMu { get; set; }

        /// <summary>
        /// AeroScout Mobile Unit (MU) mode dilution factor (default = 20).
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
        /// Enable/disable Ekahua blink mode (also called AiRISTA Flow Blink Mode) to find the location of devices connected to a wireless LAN (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ekahauBlinkMode")]
        public Input<string>? EkahauBlinkMode { get; set; }

        /// <summary>
        /// WiFi frame MAC address or WiFi Tag.
        /// </summary>
        [Input("ekahauTag")]
        public Input<string>? EkahauTag { get; set; }

        /// <summary>
        /// IP address of Ekahua RTLS Controller (ERC).
        /// </summary>
        [Input("ercServerIp")]
        public Input<string>? ErcServerIp { get; set; }

        /// <summary>
        /// Ekahua RTLS Controller (ERC) UDP listening port.
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
        /// FortiPresence server UDP listening port (default = 3000).
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
        /// Enable/disable client station locating services for all clients, whether associated or not (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("stationLocate")]
        public Input<string>? StationLocate { get; set; }

        public WirelesscontrollerWtpprofileLbsGetArgs()
        {
        }
        public static new WirelesscontrollerWtpprofileLbsGetArgs Empty => new WirelesscontrollerWtpprofileLbsGetArgs();
    }
}
