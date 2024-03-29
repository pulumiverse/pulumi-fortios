// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Web.Inputs
{

    public sealed class ProfileAntiphishArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication methods. Valid values: `domain-controller`, `ldap`.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Enable/disable checking of HTTP Basic Auth field for known credentials. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("checkBasicAuth")]
        public Input<string>? CheckBasicAuth { get; set; }

        /// <summary>
        /// Enable/disable checking of GET URI parameters for known credentials. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("checkUri")]
        public Input<string>? CheckUri { get; set; }

        /// <summary>
        /// Enable/disable acting only on valid username credentials. Action will be taken for valid usernames regardless of password validity. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("checkUsernameOnly")]
        public Input<string>? CheckUsernameOnly { get; set; }

        [Input("customPatterns")]
        private InputList<Inputs.ProfileAntiphishCustomPatternArgs>? _customPatterns;

        /// <summary>
        /// Custom username and password regex patterns. The structure of `custom_patterns` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileAntiphishCustomPatternArgs> CustomPatterns
        {
            get => _customPatterns ?? (_customPatterns = new InputList<Inputs.ProfileAntiphishCustomPatternArgs>());
            set => _customPatterns = value;
        }

        /// <summary>
        /// Action to be taken when there is no matching rule. Valid values: `exempt`, `log`, `block`.
        /// </summary>
        [Input("defaultAction")]
        public Input<string>? DefaultAction { get; set; }

        /// <summary>
        /// Domain for which to verify received credentials against.
        /// </summary>
        [Input("domainController")]
        public Input<string>? DomainController { get; set; }

        [Input("inspectionEntries")]
        private InputList<Inputs.ProfileAntiphishInspectionEntryArgs>? _inspectionEntries;

        /// <summary>
        /// AntiPhishing entries. The structure of `inspection_entries` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileAntiphishInspectionEntryArgs> InspectionEntries
        {
            get => _inspectionEntries ?? (_inspectionEntries = new InputList<Inputs.ProfileAntiphishInspectionEntryArgs>());
            set => _inspectionEntries = value;
        }

        /// <summary>
        /// LDAP server for which to verify received credentials against.
        /// </summary>
        [Input("ldap")]
        public Input<string>? Ldap { get; set; }

        /// <summary>
        /// Maximum size of a POST body to check for credentials.
        /// </summary>
        [Input("maxBodyLen")]
        public Input<int>? MaxBodyLen { get; set; }

        /// <summary>
        /// Toggle AntiPhishing functionality. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ProfileAntiphishArgs()
        {
        }
        public static new ProfileAntiphishArgs Empty => new ProfileAntiphishArgs();
    }
}
