// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Dns.Inputs
{

    public sealed class ProfileExternalIpBlocklistArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// External domain block list name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ProfileExternalIpBlocklistArgs()
        {
        }
        public static new ProfileExternalIpBlocklistArgs Empty => new ProfileExternalIpBlocklistArgs();
    }
}
