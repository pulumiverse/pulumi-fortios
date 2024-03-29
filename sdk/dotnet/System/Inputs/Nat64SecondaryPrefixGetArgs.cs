// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class Nat64SecondaryPrefixGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// NAT64 prefix name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// NAT64 prefix.
        /// </summary>
        [Input("nat64Prefix")]
        public Input<string>? Nat64Prefix { get; set; }

        public Nat64SecondaryPrefixGetArgs()
        {
        }
        public static new Nat64SecondaryPrefixGetArgs Empty => new Nat64SecondaryPrefixGetArgs();
    }
}
