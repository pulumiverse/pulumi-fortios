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

    public sealed class FirewallAddressFssoGroupGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FSSO group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallAddressFssoGroupGetArgs()
        {
        }
        public static new FirewallAddressFssoGroupGetArgs Empty => new FirewallAddressFssoGroupGetArgs();
    }
}
