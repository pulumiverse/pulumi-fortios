// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewallconsolidated.Inputs
{

    public sealed class PolicyInternetServiceCustomGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom Internet Service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PolicyInternetServiceCustomGetArgs()
        {
        }
        public static new PolicyInternetServiceCustomGetArgs Empty => new PolicyInternetServiceCustomGetArgs();
    }
}
