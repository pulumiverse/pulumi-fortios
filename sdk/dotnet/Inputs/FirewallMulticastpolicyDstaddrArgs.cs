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

    public sealed class FirewallMulticastpolicyDstaddrArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Destination address objects.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallMulticastpolicyDstaddrArgs()
        {
        }
        public static new FirewallMulticastpolicyDstaddrArgs Empty => new FirewallMulticastpolicyDstaddrArgs();
    }
}
