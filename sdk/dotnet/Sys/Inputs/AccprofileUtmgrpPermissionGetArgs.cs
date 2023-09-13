// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Inputs
{

    public sealed class AccprofileUtmgrpPermissionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Antivirus profiles and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("antivirus")]
        public Input<string>? Antivirus { get; set; }

        /// <summary>
        /// Application Control profiles and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("applicationControl")]
        public Input<string>? ApplicationControl { get; set; }

        /// <summary>
        /// DLP profiles and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("dataLossPrevention")]
        public Input<string>? DataLossPrevention { get; set; }

        /// <summary>
        /// DNS Filter profiles and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("dnsfilter")]
        public Input<string>? Dnsfilter { get; set; }

        /// <summary>
        /// AntiSpam filter and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("emailfilter")]
        public Input<string>? Emailfilter { get; set; }

        /// <summary>
        /// FortiClient Profiles. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("endpointControl")]
        public Input<string>? EndpointControl { get; set; }

        /// <summary>
        /// File-filter profiles and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("fileFilter")]
        public Input<string>? FileFilter { get; set; }

        /// <summary>
        /// ICAP profiles and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("icap")]
        public Input<string>? Icap { get; set; }

        /// <summary>
        /// IPS profiles and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("ips")]
        public Input<string>? Ips { get; set; }

        /// <summary>
        /// AntiSpam filter and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("spamfilter")]
        public Input<string>? Spamfilter { get; set; }

        /// <summary>
        /// Video filter profiles and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("videofilter")]
        public Input<string>? Videofilter { get; set; }

        /// <summary>
        /// VoIP profiles and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("voip")]
        public Input<string>? Voip { get; set; }

        /// <summary>
        /// Web Application Firewall profiles and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("waf")]
        public Input<string>? Waf { get; set; }

        /// <summary>
        /// Web Filter profiles and settings. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("webfilter")]
        public Input<string>? Webfilter { get; set; }

        public AccprofileUtmgrpPermissionGetArgs()
        {
        }
        public static new AccprofileUtmgrpPermissionGetArgs Empty => new AccprofileUtmgrpPermissionGetArgs();
    }
}
