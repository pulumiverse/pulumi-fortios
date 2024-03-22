// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Casb.Inputs
{

    public sealed class ProfileSaasApplicationAccessRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CASB access rule action. Valid values: `bypass`, `block`, `monitor`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// CASB bypass options. Valid values: `av`, `dlp`, `web-filter`, `file-filter`, `video-filter`.
        /// </summary>
        [Input("bypass")]
        public Input<string>? Bypass { get; set; }

        /// <summary>
        /// CASB access rule activity name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ProfileSaasApplicationAccessRuleGetArgs()
        {
        }
        public static new ProfileSaasApplicationAccessRuleGetArgs Empty => new ProfileSaasApplicationAccessRuleGetArgs();
    }
}
