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

    public sealed class WtpprofileRadio4ChannelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Channel number.
        /// </summary>
        [Input("chan")]
        public Input<string>? Chan { get; set; }

        public WtpprofileRadio4ChannelArgs()
        {
        }
        public static new WtpprofileRadio4ChannelArgs Empty => new WtpprofileRadio4ChannelArgs();
    }
}
