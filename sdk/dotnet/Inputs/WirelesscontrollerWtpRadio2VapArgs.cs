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

    public sealed class WirelesscontrollerWtpRadio2VapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Virtual Access Point (VAP) name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public WirelesscontrollerWtpRadio2VapArgs()
        {
        }
        public static new WirelesscontrollerWtpRadio2VapArgs Empty => new WirelesscontrollerWtpRadio2VapArgs();
    }
}
