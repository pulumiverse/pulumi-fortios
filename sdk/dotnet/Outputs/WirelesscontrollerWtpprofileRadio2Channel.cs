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
    public sealed class WirelesscontrollerWtpprofileRadio2Channel
    {
        /// <summary>
        /// Channel number.
        /// </summary>
        public readonly string? Chan;

        [OutputConstructor]
        private WirelesscontrollerWtpprofileRadio2Channel(string? chan)
        {
            Chan = chan;
        }
    }
}
