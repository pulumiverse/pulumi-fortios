// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class WanoptContentdeliverynetworkruleRule
    {
        /// <summary>
        /// Content ID settings. The structure of `content_id` block is documented below.
        /// </summary>
        public readonly Outputs.WanoptContentdeliverynetworkruleRuleContentId? ContentId;
        /// <summary>
        /// List of entries to match. The structure of `match_entries` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.WanoptContentdeliverynetworkruleRuleMatchEntry> MatchEntries;
        /// <summary>
        /// Match criteria for collecting content ID. Valid values: `all`, `any`.
        /// </summary>
        public readonly string? MatchMode;
        /// <summary>
        /// WAN optimization content delivery network rule name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// List of entries to skip. The structure of `skip_entries` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.WanoptContentdeliverynetworkruleRuleSkipEntry> SkipEntries;
        /// <summary>
        /// Skip mode when evaluating skip-rules. Valid values: `all`, `any`.
        /// </summary>
        public readonly string? SkipRuleMode;

        [OutputConstructor]
        private WanoptContentdeliverynetworkruleRule(
            Outputs.WanoptContentdeliverynetworkruleRuleContentId? contentId,

            ImmutableArray<Outputs.WanoptContentdeliverynetworkruleRuleMatchEntry> matchEntries,

            string? matchMode,

            string? name,

            ImmutableArray<Outputs.WanoptContentdeliverynetworkruleRuleSkipEntry> skipEntries,

            string? skipRuleMode)
        {
            ContentId = contentId;
            MatchEntries = matchEntries;
            MatchMode = matchMode;
            Name = name;
            SkipEntries = skipEntries;
            SkipRuleMode = skipRuleMode;
        }
    }
}
