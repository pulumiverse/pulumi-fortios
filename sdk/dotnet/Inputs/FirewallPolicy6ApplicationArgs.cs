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

    public sealed class FirewallPolicy6ApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application IDs.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public FirewallPolicy6ApplicationArgs()
        {
        }
        public static new FirewallPolicy6ApplicationArgs Empty => new FirewallPolicy6ApplicationArgs();
    }
}
