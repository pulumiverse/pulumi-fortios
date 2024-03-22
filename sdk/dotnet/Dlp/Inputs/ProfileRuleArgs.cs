// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Dlp.Inputs
{

    public sealed class ProfileRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take with content that this DLP profile matches. Valid values: `allow`, `log-only`, `block`, `quarantine-ip`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Enable/disable DLP archiving. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("archive")]
        public Input<string>? Archive { get; set; }

        /// <summary>
        /// Quarantine duration in days, hours, minutes (format = dddhhmm).
        /// </summary>
        [Input("expiry")]
        public Input<string>? Expiry { get; set; }

        /// <summary>
        /// Match files greater than or equal to this size (KB).
        /// </summary>
        [Input("fileSize")]
        public Input<int>? FileSize { get; set; }

        /// <summary>
        /// Select the number of a DLP file pattern table to match.
        /// </summary>
        [Input("fileType")]
        public Input<int>? FileType { get; set; }

        /// <summary>
        /// Select the type of content to match. Valid values: `sensor`, `mip`, `fingerprint`, `encrypted`, `none`.
        /// </summary>
        [Input("filterBy")]
        public Input<string>? FilterBy { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// MIP label dictionary.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Percentage of fingerprints in the fingerprint databases designated with the selected sensitivity to match.
        /// </summary>
        [Input("matchPercentage")]
        public Input<int>? MatchPercentage { get; set; }

        /// <summary>
        /// Filter name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Check messages or files over one or more of these protocols. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        /// </summary>
        [Input("proto")]
        public Input<string>? Proto { get; set; }

        [Input("sensitivities")]
        private InputList<Inputs.ProfileRuleSensitivityArgs>? _sensitivities;

        /// <summary>
        /// Select a DLP file pattern sensitivity to match. The structure of `sensitivity` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileRuleSensitivityArgs> Sensitivities
        {
            get => _sensitivities ?? (_sensitivities = new InputList<Inputs.ProfileRuleSensitivityArgs>());
            set => _sensitivities = value;
        }

        [Input("sensors")]
        private InputList<Inputs.ProfileRuleSensorArgs>? _sensors;

        /// <summary>
        /// Select DLP sensors. The structure of `sensor` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileRuleSensorArgs> Sensors
        {
            get => _sensors ?? (_sensors = new InputList<Inputs.ProfileRuleSensorArgs>());
            set => _sensors = value;
        }

        /// <summary>
        /// Select the severity or threat level that matches this filter. Valid values: `info`, `low`, `medium`, `high`, `critical`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Select whether to check the content of messages (an email message) or files (downloaded files or email attachments). Valid values: `file`, `message`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ProfileRuleArgs()
        {
        }
        public static new ProfileRuleArgs Empty => new ProfileRuleArgs();
    }
}
