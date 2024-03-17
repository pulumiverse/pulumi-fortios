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
    public sealed class ProfilePop3
    {
        /// <summary>
        /// Action for spam email. Valid values: `pass`, `tag`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Log;
        /// <summary>
        /// Subject text or header added to spam email.
        /// </summary>
        public readonly string? TagMsg;
        /// <summary>
        /// Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
        /// </summary>
        public readonly string? TagType;

        [OutputConstructor]
        private ProfilePop3(
            string? action,

            string? log,

            string? tagMsg,

            string? tagType)
        {
            Action = action;
            Log = log;
            TagMsg = tagMsg;
            TagType = tagType;
        }
    }
}
