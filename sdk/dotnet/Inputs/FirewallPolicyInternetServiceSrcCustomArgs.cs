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

    public sealed class FirewallPolicyInternetServiceSrcCustomArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom Internet Service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallPolicyInternetServiceSrcCustomArgs()
        {
        }
        public static new FirewallPolicyInternetServiceSrcCustomArgs Empty => new FirewallPolicyInternetServiceSrcCustomArgs();
    }
}
