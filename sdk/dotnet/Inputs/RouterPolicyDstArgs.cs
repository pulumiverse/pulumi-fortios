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

    public sealed class RouterPolicyDstArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP and mask.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        public RouterPolicyDstArgs()
        {
        }
        public static new RouterPolicyDstArgs Empty => new RouterPolicyDstArgs();
    }
}
