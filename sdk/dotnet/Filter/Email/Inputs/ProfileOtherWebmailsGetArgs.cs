// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Email.Inputs
{

    public sealed class ProfileOtherWebmailsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logAll")]
        public Input<string>? LogAll { get; set; }

        public ProfileOtherWebmailsGetArgs()
        {
        }
        public static new ProfileOtherWebmailsGetArgs Empty => new ProfileOtherWebmailsGetArgs();
    }
}
