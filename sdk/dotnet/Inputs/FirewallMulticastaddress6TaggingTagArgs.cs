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

    public sealed class FirewallMulticastaddress6TaggingTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tag name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallMulticastaddress6TaggingTagArgs()
        {
        }
        public static new FirewallMulticastaddress6TaggingTagArgs Empty => new FirewallMulticastaddress6TaggingTagArgs();
    }
}