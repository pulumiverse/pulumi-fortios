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
    public sealed class GetSystemAccprofileUtmgrpPermissionResult
    {
        /// <summary>
        /// Antivirus profiles and settings.
        /// </summary>
        public readonly string Antivirus;
        /// <summary>
        /// Application Control profiles and settings.
        /// </summary>
        public readonly string ApplicationControl;
        /// <summary>
        /// DLP profiles and settings.
        /// </summary>
        public readonly string DataLossPrevention;
        /// <summary>
        /// DNS Filter profiles and settings.
        /// </summary>
        public readonly string Dnsfilter;
        /// <summary>
        /// AntiSpam filter and settings.
        /// </summary>
        public readonly string Emailfilter;
        /// <summary>
        /// FortiClient Profiles.
        /// </summary>
        public readonly string EndpointControl;
        /// <summary>
        /// File-filter profiles and settings.
        /// </summary>
        public readonly string FileFilter;
        /// <summary>
        /// ICAP profiles and settings.
        /// </summary>
        public readonly string Icap;
        /// <summary>
        /// IPS profiles and settings.
        /// </summary>
        public readonly string Ips;
        /// <summary>
        /// AntiSpam filter and settings.
        /// </summary>
        public readonly string Spamfilter;
        /// <summary>
        /// Video filter profiles and settings.
        /// </summary>
        public readonly string Videofilter;
        /// <summary>
        /// VoIP profiles and settings.
        /// </summary>
        public readonly string Voip;
        /// <summary>
        /// Web Application Firewall profiles and settings.
        /// </summary>
        public readonly string Waf;
        /// <summary>
        /// Web Filter profiles and settings.
        /// </summary>
        public readonly string Webfilter;

        [OutputConstructor]
        private GetSystemAccprofileUtmgrpPermissionResult(
            string antivirus,

            string applicationControl,

            string dataLossPrevention,

            string dnsfilter,

            string emailfilter,

            string endpointControl,

            string fileFilter,

            string icap,

            string ips,

            string spamfilter,

            string videofilter,

            string voip,

            string waf,

            string webfilter)
        {
            Antivirus = antivirus;
            ApplicationControl = applicationControl;
            DataLossPrevention = dataLossPrevention;
            Dnsfilter = dnsfilter;
            Emailfilter = emailfilter;
            EndpointControl = endpointControl;
            FileFilter = fileFilter;
            Icap = icap;
            Ips = ips;
            Spamfilter = spamfilter;
            Videofilter = videofilter;
            Voip = voip;
            Waf = waf;
            Webfilter = webfilter;
        }
    }
}
