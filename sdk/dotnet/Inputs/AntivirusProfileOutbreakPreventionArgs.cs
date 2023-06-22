// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class AntivirusProfileOutbreakPreventionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable external malware blocklist. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("externalBlocklist")]
        public Input<string>? ExternalBlocklist { get; set; }

        /// <summary>
        /// Enable/disable FortiGuard Virus outbreak prevention service. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("ftgdService")]
        public Input<string>? FtgdService { get; set; }

        public AntivirusProfileOutbreakPreventionArgs()
        {
        }
        public static new AntivirusProfileOutbreakPreventionArgs Empty => new AntivirusProfileOutbreakPreventionArgs();
    }
}
