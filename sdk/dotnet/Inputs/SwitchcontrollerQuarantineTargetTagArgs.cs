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

    public sealed class SwitchcontrollerQuarantineTargetTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tag string(eg. string1 string2 string3).
        /// </summary>
        [Input("tags")]
        public Input<string>? Tags { get; set; }

        public SwitchcontrollerQuarantineTargetTagArgs()
        {
        }
        public static new SwitchcontrollerQuarantineTargetTagArgs Empty => new SwitchcontrollerQuarantineTargetTagArgs();
    }
}
