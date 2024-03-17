// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Endpointcontrol.Inputs
{

    public sealed class ProfileForticlientWinmacSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable FortiClient AntiVirus real-time protection. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("avRealtimeProtection")]
        public Input<string>? AvRealtimeProtection { get; set; }

        /// <summary>
        /// Enable/disable FortiClient AV signature updates. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("avSignatureUpToDate")]
        public Input<string>? AvSignatureUpToDate { get; set; }

        /// <summary>
        /// Enable/disable the FortiClient application firewall. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientApplicationFirewall")]
        public Input<string>? ForticlientApplicationFirewall { get; set; }

        /// <summary>
        /// FortiClient application firewall rule list.
        /// </summary>
        [Input("forticlientApplicationFirewallList")]
        public Input<string>? ForticlientApplicationFirewallList { get; set; }

        /// <summary>
        /// Enable/disable FortiClient AntiVirus scanning. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientAv")]
        public Input<string>? ForticlientAv { get; set; }

        /// <summary>
        /// Enable/disable FortiClient Enterprise Management Server (EMS) compliance. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientEmsCompliance")]
        public Input<string>? ForticlientEmsCompliance { get; set; }

        /// <summary>
        /// FortiClient EMS compliance action. Valid values: `block`, `warning`.
        /// </summary>
        [Input("forticlientEmsComplianceAction")]
        public Input<string>? ForticlientEmsComplianceAction { get; set; }

        [Input("forticlientEmsEntries")]
        private InputList<Inputs.ProfileForticlientWinmacSettingsForticlientEmsEntryGetArgs>? _forticlientEmsEntries;

        /// <summary>
        /// FortiClient EMS entries. The structure of `forticlient_ems_entries` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileForticlientWinmacSettingsForticlientEmsEntryGetArgs> ForticlientEmsEntries
        {
            get => _forticlientEmsEntries ?? (_forticlientEmsEntries = new InputList<Inputs.ProfileForticlientWinmacSettingsForticlientEmsEntryGetArgs>());
            set => _forticlientEmsEntries = value;
        }

        /// <summary>
        /// Minimum FortiClient Linux version.
        /// </summary>
        [Input("forticlientLinuxVer")]
        public Input<string>? ForticlientLinuxVer { get; set; }

        /// <summary>
        /// Enable/disable uploading FortiClient logs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientLogUpload")]
        public Input<string>? ForticlientLogUpload { get; set; }

        /// <summary>
        /// Select the FortiClient logs to upload. Valid values: `traffic`, `vulnerability`, `event`.
        /// </summary>
        [Input("forticlientLogUploadLevel")]
        public Input<string>? ForticlientLogUploadLevel { get; set; }

        /// <summary>
        /// IP address or FQDN of the server to which to upload FortiClient logs.
        /// </summary>
        [Input("forticlientLogUploadServer")]
        public Input<string>? ForticlientLogUploadServer { get; set; }

        /// <summary>
        /// Minimum FortiClient Mac OS version.
        /// </summary>
        [Input("forticlientMacVer")]
        public Input<string>? ForticlientMacVer { get; set; }

        /// <summary>
        /// Enable/disable requiring clients to run FortiClient with a minimum software version number. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientMinimumSoftwareVersion")]
        public Input<string>? ForticlientMinimumSoftwareVersion { get; set; }

        [Input("forticlientOperatingSystems")]
        private InputList<Inputs.ProfileForticlientWinmacSettingsForticlientOperatingSystemGetArgs>? _forticlientOperatingSystems;

        /// <summary>
        /// FortiClient operating system. The structure of `forticlient_operating_system` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileForticlientWinmacSettingsForticlientOperatingSystemGetArgs> ForticlientOperatingSystems
        {
            get => _forticlientOperatingSystems ?? (_forticlientOperatingSystems = new InputList<Inputs.ProfileForticlientWinmacSettingsForticlientOperatingSystemGetArgs>());
            set => _forticlientOperatingSystems = value;
        }

        [Input("forticlientOwnFiles")]
        private InputList<Inputs.ProfileForticlientWinmacSettingsForticlientOwnFileGetArgs>? _forticlientOwnFiles;

        /// <summary>
        /// Checking the path and filename of the FortiClient application. The structure of `forticlient_own_file` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileForticlientWinmacSettingsForticlientOwnFileGetArgs> ForticlientOwnFiles
        {
            get => _forticlientOwnFiles ?? (_forticlientOwnFiles = new InputList<Inputs.ProfileForticlientWinmacSettingsForticlientOwnFileGetArgs>());
            set => _forticlientOwnFiles = value;
        }

        /// <summary>
        /// FortiClient registration compliance action. Valid values: `block`, `warning`.
        /// </summary>
        [Input("forticlientRegistrationComplianceAction")]
        public Input<string>? ForticlientRegistrationComplianceAction { get; set; }

        [Input("forticlientRegistryEntries")]
        private InputList<Inputs.ProfileForticlientWinmacSettingsForticlientRegistryEntryGetArgs>? _forticlientRegistryEntries;

        /// <summary>
        /// FortiClient registry entry. The structure of `forticlient_registry_entry` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileForticlientWinmacSettingsForticlientRegistryEntryGetArgs> ForticlientRegistryEntries
        {
            get => _forticlientRegistryEntries ?? (_forticlientRegistryEntries = new InputList<Inputs.ProfileForticlientWinmacSettingsForticlientRegistryEntryGetArgs>());
            set => _forticlientRegistryEntries = value;
        }

        [Input("forticlientRunningApps")]
        private InputList<Inputs.ProfileForticlientWinmacSettingsForticlientRunningAppGetArgs>? _forticlientRunningApps;

        /// <summary>
        /// Use FortiClient to verify if the listed applications are running on the client. The structure of `forticlient_running_app` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileForticlientWinmacSettingsForticlientRunningAppGetArgs> ForticlientRunningApps
        {
            get => _forticlientRunningApps ?? (_forticlientRunningApps = new InputList<Inputs.ProfileForticlientWinmacSettingsForticlientRunningAppGetArgs>());
            set => _forticlientRunningApps = value;
        }

        /// <summary>
        /// Enable/disable FortiClient security posture check options. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientSecurityPosture")]
        public Input<string>? ForticlientSecurityPosture { get; set; }

        /// <summary>
        /// FortiClient security posture compliance action. Valid values: `block`, `warning`.
        /// </summary>
        [Input("forticlientSecurityPostureComplianceAction")]
        public Input<string>? ForticlientSecurityPostureComplianceAction { get; set; }

        /// <summary>
        /// Enable/disable enforcement of FortiClient system compliance. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientSystemCompliance")]
        public Input<string>? ForticlientSystemCompliance { get; set; }

        /// <summary>
        /// Block or warn clients not compliant with FortiClient requirements. Valid values: `block`, `warning`.
        /// </summary>
        [Input("forticlientSystemComplianceAction")]
        public Input<string>? ForticlientSystemComplianceAction { get; set; }

        /// <summary>
        /// Enable/disable FortiClient vulnerability scanning. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientVulnScan")]
        public Input<string>? ForticlientVulnScan { get; set; }

        /// <summary>
        /// FortiClient vulnerability compliance action. Valid values: `block`, `warning`.
        /// </summary>
        [Input("forticlientVulnScanComplianceAction")]
        public Input<string>? ForticlientVulnScanComplianceAction { get; set; }

        /// <summary>
        /// Configure the level of the vulnerability found that causes a FortiClient vulnerability compliance action. Valid values: `critical`, `high`, `medium`, `low`, `info`.
        /// </summary>
        [Input("forticlientVulnScanEnforce")]
        public Input<string>? ForticlientVulnScanEnforce { get; set; }

        /// <summary>
        /// FortiClient vulnerability scan enforcement grace period (0 - 30 days, default = 1).
        /// </summary>
        [Input("forticlientVulnScanEnforceGrace")]
        public Input<int>? ForticlientVulnScanEnforceGrace { get; set; }

        /// <summary>
        /// Enable/disable compliance exemption for vulnerabilities that cannot be patched automatically. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientVulnScanExempt")]
        public Input<string>? ForticlientVulnScanExempt { get; set; }

        /// <summary>
        /// Enable/disable FortiClient web filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientWf")]
        public Input<string>? ForticlientWf { get; set; }

        /// <summary>
        /// The FortiClient web filter profile to apply.
        /// </summary>
        [Input("forticlientWfProfile")]
        public Input<string>? ForticlientWfProfile { get; set; }

        /// <summary>
        /// Minimum FortiClient Windows version.
        /// </summary>
        [Input("forticlientWinVer")]
        public Input<string>? ForticlientWinVer { get; set; }

        /// <summary>
        /// Enable/disable checking for OS recognized AntiVirus software. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("osAvSoftwareInstalled")]
        public Input<string>? OsAvSoftwareInstalled { get; set; }

        /// <summary>
        /// FortiSandbox address.
        /// </summary>
        [Input("sandboxAddress")]
        public Input<string>? SandboxAddress { get; set; }

        /// <summary>
        /// Enable/disable sending files to FortiSandbox for analysis. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sandboxAnalysis")]
        public Input<string>? SandboxAnalysis { get; set; }

        public ProfileForticlientWinmacSettingsGetArgs()
        {
        }
        public static new ProfileForticlientWinmacSettingsGetArgs Empty => new ProfileForticlientWinmacSettingsGetArgs();
    }
}
