// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Extensioncontroller.Outputs
{

    [OutputType]
    public sealed class ExtenderprofileWifiRadio2LocalVap
    {
        /// <summary>
        /// Wi-Fi local VAP name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private ExtenderprofileWifiRadio2LocalVap(string? name)
        {
            Name = name;
        }
    }
}
