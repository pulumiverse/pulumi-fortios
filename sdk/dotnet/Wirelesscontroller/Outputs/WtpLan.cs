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
    public sealed class WtpLan
    {
        /// <summary>
        /// LAN port 1 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        public readonly string? Port1Mode;
        /// <summary>
        /// Bridge LAN port 1 to SSID.
        /// </summary>
        public readonly string? Port1Ssid;
        /// <summary>
        /// LAN port 2 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        public readonly string? Port2Mode;
        /// <summary>
        /// Bridge LAN port 2 to SSID.
        /// </summary>
        public readonly string? Port2Ssid;
        /// <summary>
        /// LAN port 3 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        public readonly string? Port3Mode;
        /// <summary>
        /// Bridge LAN port 3 to SSID.
        /// </summary>
        public readonly string? Port3Ssid;
        /// <summary>
        /// LAN port 4 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        public readonly string? Port4Mode;
        /// <summary>
        /// Bridge LAN port 4 to SSID.
        /// </summary>
        public readonly string? Port4Ssid;
        /// <summary>
        /// LAN port 5 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        public readonly string? Port5Mode;
        /// <summary>
        /// Bridge LAN port 5 to SSID.
        /// </summary>
        public readonly string? Port5Ssid;
        /// <summary>
        /// LAN port 6 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        public readonly string? Port6Mode;
        /// <summary>
        /// Bridge LAN port 6 to SSID.
        /// </summary>
        public readonly string? Port6Ssid;
        /// <summary>
        /// LAN port 7 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        public readonly string? Port7Mode;
        /// <summary>
        /// Bridge LAN port 7 to SSID.
        /// </summary>
        public readonly string? Port7Ssid;
        /// <summary>
        /// LAN port 8 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        public readonly string? Port8Mode;
        /// <summary>
        /// Bridge LAN port 8 to SSID.
        /// </summary>
        public readonly string? Port8Ssid;
        /// <summary>
        /// ESL port mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        public readonly string? PortEslMode;
        /// <summary>
        /// Bridge ESL port to SSID.
        /// 
        /// The `radio_1` block supports:
        /// </summary>
        public readonly string? PortEslSsid;
        /// <summary>
        /// LAN port mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        public readonly string? PortMode;
        /// <summary>
        /// Bridge LAN port to SSID.
        /// </summary>
        public readonly string? PortSsid;

        [OutputConstructor]
        private WtpLan(
            string? port1Mode,

            string? port1Ssid,

            string? port2Mode,

            string? port2Ssid,

            string? port3Mode,

            string? port3Ssid,

            string? port4Mode,

            string? port4Ssid,

            string? port5Mode,

            string? port5Ssid,

            string? port6Mode,

            string? port6Ssid,

            string? port7Mode,

            string? port7Ssid,

            string? port8Mode,

            string? port8Ssid,

            string? portEslMode,

            string? portEslSsid,

            string? portMode,

            string? portSsid)
        {
            Port1Mode = port1Mode;
            Port1Ssid = port1Ssid;
            Port2Mode = port2Mode;
            Port2Ssid = port2Ssid;
            Port3Mode = port3Mode;
            Port3Ssid = port3Ssid;
            Port4Mode = port4Mode;
            Port4Ssid = port4Ssid;
            Port5Mode = port5Mode;
            Port5Ssid = port5Ssid;
            Port6Mode = port6Mode;
            Port6Ssid = port6Ssid;
            Port7Mode = port7Mode;
            Port7Ssid = port7Ssid;
            Port8Mode = port8Mode;
            Port8Ssid = port8Ssid;
            PortEslMode = portEslMode;
            PortEslSsid = portEslSsid;
            PortMode = portMode;
            PortSsid = portSsid;
        }
    }
}