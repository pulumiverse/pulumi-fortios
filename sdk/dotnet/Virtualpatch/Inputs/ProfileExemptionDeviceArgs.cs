// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Virtualpatch.Inputs
{

    public sealed class ProfileExemptionDeviceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Device MAC address.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        public ProfileExemptionDeviceArgs()
        {
        }
        public static new ProfileExemptionDeviceArgs Empty => new ProfileExemptionDeviceArgs();
    }
}