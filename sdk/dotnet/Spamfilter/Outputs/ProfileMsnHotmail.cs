// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Spamfilter.Outputs
{

    [OutputType]
    public sealed class ProfileMsnHotmail
    {
        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Log;

        [OutputConstructor]
        private ProfileMsnHotmail(string? log)
        {
            Log = log;
        }
    }
}
