// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Outputs
{

    [OutputType]
    public sealed class GetAccprofileLoggrpPermissionResult
    {
        /// <summary>
        /// Log &amp; Report configuration.
        /// </summary>
        public readonly string Config;
        /// <summary>
        /// Log &amp; Report Data Access.
        /// </summary>
        public readonly string DataAccess;
        /// <summary>
        /// Log &amp; Report Report Access.
        /// </summary>
        public readonly string ReportAccess;
        /// <summary>
        /// Log &amp; Report Threat Weight.
        /// </summary>
        public readonly string ThreatWeight;

        [OutputConstructor]
        private GetAccprofileLoggrpPermissionResult(
            string config,

            string dataAccess,

            string reportAccess,

            string threatWeight)
        {
            Config = config;
            DataAccess = dataAccess;
            ReportAccess = reportAccess;
            ThreatWeight = threatWeight;
        }
    }
}
