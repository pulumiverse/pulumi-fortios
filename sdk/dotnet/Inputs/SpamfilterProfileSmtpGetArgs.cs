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

    public sealed class SpamfilterProfileSmtpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action for spam email. Valid values: `pass`, `tag`, `discard`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Enable/disable SMTP email header IP checks for spamfsip, spamrbl and spambwl filters. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("hdrip")]
        public Input<string>? Hdrip { get; set; }

        /// <summary>
        /// Enable/disable local filter to override SMTP remote check result. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("localOverride")]
        public Input<string>? LocalOverride { get; set; }

        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Subject text or header added to spam email.
        /// </summary>
        [Input("tagMsg")]
        public Input<string>? TagMsg { get; set; }

        /// <summary>
        /// Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
        /// </summary>
        [Input("tagType")]
        public Input<string>? TagType { get; set; }

        public SpamfilterProfileSmtpGetArgs()
        {
        }
        public static new SpamfilterProfileSmtpGetArgs Empty => new SpamfilterProfileSmtpGetArgs();
    }
}
