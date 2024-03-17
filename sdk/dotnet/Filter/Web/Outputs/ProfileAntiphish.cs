// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Web.Outputs
{

    [OutputType]
    public sealed class ProfileAntiphish
    {
        /// <summary>
        /// Authentication methods. Valid values: `domain-controller`, `ldap`.
        /// </summary>
        public readonly string? Authentication;
        /// <summary>
        /// Enable/disable checking of HTTP Basic Auth field for known credentials. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? CheckBasicAuth;
        /// <summary>
        /// Enable/disable checking of GET URI parameters for known credentials. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? CheckUri;
        /// <summary>
        /// Enable/disable acting only on valid username credentials. Action will be taken for valid usernames regardless of password validity. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? CheckUsernameOnly;
        /// <summary>
        /// Custom username and password regex patterns. The structure of `custom_patterns` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProfileAntiphishCustomPattern> CustomPatterns;
        /// <summary>
        /// Action to be taken when there is no matching rule. Valid values: `exempt`, `log`, `block`.
        /// </summary>
        public readonly string? DefaultAction;
        /// <summary>
        /// Domain for which to verify received credentials against.
        /// </summary>
        public readonly string? DomainController;
        /// <summary>
        /// AntiPhishing entries. The structure of `inspection_entries` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProfileAntiphishInspectionEntry> InspectionEntries;
        /// <summary>
        /// LDAP server for which to verify received credentials against.
        /// </summary>
        public readonly string? Ldap;
        /// <summary>
        /// Maximum size of a POST body to check for credentials.
        /// </summary>
        public readonly int? MaxBodyLen;
        /// <summary>
        /// Toggle AntiPhishing functionality. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private ProfileAntiphish(
            string? authentication,

            string? checkBasicAuth,

            string? checkUri,

            string? checkUsernameOnly,

            ImmutableArray<Outputs.ProfileAntiphishCustomPattern> customPatterns,

            string? defaultAction,

            string? domainController,

            ImmutableArray<Outputs.ProfileAntiphishInspectionEntry> inspectionEntries,

            string? ldap,

            int? maxBodyLen,

            string? status)
        {
            Authentication = authentication;
            CheckBasicAuth = checkBasicAuth;
            CheckUri = checkUri;
            CheckUsernameOnly = checkUsernameOnly;
            CustomPatterns = customPatterns;
            DefaultAction = defaultAction;
            DomainController = domainController;
            InspectionEntries = inspectionEntries;
            Ldap = ldap;
            MaxBodyLen = maxBodyLen;
            Status = status;
        }
    }
}
