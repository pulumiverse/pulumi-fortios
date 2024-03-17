// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class AccprofileLoggrpPermission
    {
        /// <summary>
        /// Log &amp; Report configuration. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        public readonly string? Config;
        /// <summary>
        /// Log &amp; Report Data Access. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        public readonly string? DataAccess;
        /// <summary>
        /// Log &amp; Report Report Access. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        public readonly string? ReportAccess;
        /// <summary>
        /// Log &amp; Report Threat Weight. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        public readonly string? ThreatWeight;

        [OutputConstructor]
        private AccprofileLoggrpPermission(
            string? config,

            string? dataAccess,

            string? reportAccess,

            string? threatWeight)
        {
            Config = config;
            DataAccess = dataAccess;
            ReportAccess = reportAccess;
            ThreatWeight = threatWeight;
        }
    }
}
