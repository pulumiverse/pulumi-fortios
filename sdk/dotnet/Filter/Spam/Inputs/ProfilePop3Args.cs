// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Spam.Inputs
{

    public sealed class ProfilePop3Args : global::Pulumi.ResourceArgs
    {
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("log")]
        public Input<string>? Log { get; set; }

        [Input("tagMsg")]
        public Input<string>? TagMsg { get; set; }

        [Input("tagType")]
        public Input<string>? TagType { get; set; }

        public ProfilePop3Args()
        {
        }
        public static new ProfilePop3Args Empty => new ProfilePop3Args();
    }
}
