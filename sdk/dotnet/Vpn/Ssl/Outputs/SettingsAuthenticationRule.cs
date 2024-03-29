// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ssl.Outputs
{

    [OutputType]
    public sealed class SettingsAuthenticationRule
    {
        /// <summary>
        /// SSL VPN authentication method restriction.
        /// </summary>
        public readonly string? Auth;
        /// <summary>
        /// SSL VPN cipher strength. Valid values: `any`, `high`, `medium`.
        /// </summary>
        public readonly string? Cipher;
        /// <summary>
        /// Enable/disable SSL VPN client certificate restrictive. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ClientCert;
        /// <summary>
        /// User groups. The structure of `groups` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.SettingsAuthenticationRuleGroup> Groups;
        /// <summary>
        /// ID (0 - 4294967295).
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// SSL VPN portal.
        /// </summary>
        public readonly string? Portal;
        /// <summary>
        /// SSL VPN realm.
        /// </summary>
        public readonly string? Realm;
        /// <summary>
        /// Enable/disable negated source IPv6 address match. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? SourceAddress6Negate;
        /// <summary>
        /// IPv6 source address of incoming traffic. The structure of `source_address6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.SettingsAuthenticationRuleSourceAddress6> SourceAddress6s;
        /// <summary>
        /// Enable/disable negated source address match. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? SourceAddressNegate;
        /// <summary>
        /// Source address of incoming traffic. The structure of `source_address` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.SettingsAuthenticationRuleSourceAddress> SourceAddresses;
        /// <summary>
        /// SSL VPN source interface of incoming traffic. The structure of `source_interface` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.SettingsAuthenticationRuleSourceInterface> SourceInterfaces;
        /// <summary>
        /// Name of user peer.
        /// </summary>
        public readonly string? UserPeer;
        /// <summary>
        /// User name. The structure of `users` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.SettingsAuthenticationRuleUser> Users;

        [OutputConstructor]
        private SettingsAuthenticationRule(
            string? auth,

            string? cipher,

            string? clientCert,

            ImmutableArray<Outputs.SettingsAuthenticationRuleGroup> groups,

            int? id,

            string? portal,

            string? realm,

            string? sourceAddress6Negate,

            ImmutableArray<Outputs.SettingsAuthenticationRuleSourceAddress6> sourceAddress6s,

            string? sourceAddressNegate,

            ImmutableArray<Outputs.SettingsAuthenticationRuleSourceAddress> sourceAddresses,

            ImmutableArray<Outputs.SettingsAuthenticationRuleSourceInterface> sourceInterfaces,

            string? userPeer,

            ImmutableArray<Outputs.SettingsAuthenticationRuleUser> users)
        {
            Auth = auth;
            Cipher = cipher;
            ClientCert = clientCert;
            Groups = groups;
            Id = id;
            Portal = portal;
            Realm = realm;
            SourceAddress6Negate = sourceAddress6Negate;
            SourceAddress6s = sourceAddress6s;
            SourceAddressNegate = sourceAddressNegate;
            SourceAddresses = sourceAddresses;
            SourceInterfaces = sourceInterfaces;
            UserPeer = userPeer;
            Users = users;
        }
    }
}
