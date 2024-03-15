// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Spam.Outputs
{

    [OutputType]
    public sealed class BwordEntry
    {
        /// <summary>
        /// Mark spam or good. Valid values: `spam`, `clear`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Banned word entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Language for the banned word. Valid values: `western`, `simch`, `trach`, `japanese`, `korean`, `french`, `thai`, `spanish`.
        /// </summary>
        public readonly string? Language;
        /// <summary>
        /// Pattern for the banned word.
        /// </summary>
        public readonly string? Pattern;
        /// <summary>
        /// Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.
        /// </summary>
        public readonly string? PatternType;
        /// <summary>
        /// Score value.
        /// </summary>
        public readonly int? Score;
        /// <summary>
        /// Enable/disable status. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Component of the email to be scanned. Valid values: `subject`, `body`, `all`.
        /// </summary>
        public readonly string? Where;

        [OutputConstructor]
        private BwordEntry(
            string? action,

            int? id,

            string? language,

            string? pattern,

            string? patternType,

            int? score,

            string? status,

            string? where)
        {
            Action = action;
            Id = id;
            Language = language;
            Pattern = pattern;
            PatternType = patternType;
            Score = score;
            Status = status;
            Where = where;
        }
    }
}
