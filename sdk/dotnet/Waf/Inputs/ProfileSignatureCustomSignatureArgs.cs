// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Waf.Inputs
{

    public sealed class ProfileSignatureCustomSignatureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action. Valid values: `allow`, `block`, `erase`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Case sensitivity in pattern. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("caseSensitivity")]
        public Input<string>? CaseSensitivity { get; set; }

        /// <summary>
        /// Traffic direction. Valid values: `request`, `response`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Signature name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Match pattern.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Severity. Valid values: `high`, `medium`, `low`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Status. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Match HTTP target. Valid values: `arg`, `arg-name`, `req-body`, `req-cookie`, `req-cookie-name`, `req-filename`, `req-header`, `req-header-name`, `req-raw-uri`, `req-uri`, `resp-body`, `resp-hdr`, `resp-status`.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public ProfileSignatureCustomSignatureArgs()
        {
        }
        public static new ProfileSignatureCustomSignatureArgs Empty => new ProfileSignatureCustomSignatureArgs();
    }
}
