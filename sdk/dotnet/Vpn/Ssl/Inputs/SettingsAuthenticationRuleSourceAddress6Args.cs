// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ssl.Inputs
{

    public sealed class SettingsAuthenticationRuleSourceAddress6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SettingsAuthenticationRuleSourceAddress6Args()
        {
        }
        public static new SettingsAuthenticationRuleSourceAddress6Args Empty => new SettingsAuthenticationRuleSourceAddress6Args();
    }
}
