// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ssl.Web.Outputs
{

    [OutputType]
    public sealed class PortalOsCheckList
    {
        /// <summary>
        /// OS check options. Valid values: `deny`, `allow`, `check-up-to-date`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Latest OS patch level.
        /// </summary>
        public readonly string? LatestPatchLevel;
        /// <summary>
        /// Name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// OS patch level tolerance.
        /// </summary>
        public readonly int? Tolerance;

        [OutputConstructor]
        private PortalOsCheckList(
            string? action,

            string? latestPatchLevel,

            string? name,

            int? tolerance)
        {
            Action = action;
            LatestPatchLevel = latestPatchLevel;
            Name = name;
            Tolerance = tolerance;
        }
    }
}
