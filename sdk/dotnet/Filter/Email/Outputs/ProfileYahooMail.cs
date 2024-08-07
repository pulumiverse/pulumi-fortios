// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Email.Outputs
{

    [OutputType]
    public sealed class ProfileYahooMail
    {
        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Log;
        /// <summary>
        /// Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? LogAll;

        [OutputConstructor]
        private ProfileYahooMail(
            string? log,

            string? logAll)
        {
            Log = log;
            LogAll = logAll;
        }
    }
}
