// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Inputs
{

    public sealed class WtpprofilePlatformArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable use of one radio for dedicated dual-band scanning to detect RF characterization and wireless threat management. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ddscan")]
        public Input<string>? Ddscan { get; set; }

        /// <summary>
        /// Configure operation mode of 5G radios (default = single-5G). Valid values: `single-5G`, `dual-5G`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// WTP, FortiAP or AP platform type. There are built-in WTP profiles for all supported FortiAP models. You can select a built-in profile and customize it or create a new profile.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public WtpprofilePlatformArgs()
        {
        }
        public static new WtpprofilePlatformArgs Empty => new WtpprofilePlatformArgs();
    }
}
