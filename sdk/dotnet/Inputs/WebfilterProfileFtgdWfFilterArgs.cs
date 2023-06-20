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

    public sealed class WebfilterProfileFtgdWfFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take for matches. Valid values: `block`, `authenticate`, `monitor`, `warning`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("authUsrGrps")]
        private InputList<Inputs.WebfilterProfileFtgdWfFilterAuthUsrGrpArgs>? _authUsrGrps;

        /// <summary>
        /// Groups with permission to authenticate. The structure of `auth_usr_grp` block is documented below.
        /// </summary>
        public InputList<Inputs.WebfilterProfileFtgdWfFilterAuthUsrGrpArgs> AuthUsrGrps
        {
            get => _authUsrGrps ?? (_authUsrGrps = new InputList<Inputs.WebfilterProfileFtgdWfFilterAuthUsrGrpArgs>());
            set => _authUsrGrps = value;
        }

        /// <summary>
        /// Categories and groups the filter examines.
        /// </summary>
        [Input("category")]
        public Input<int>? Category { get; set; }

        /// <summary>
        /// ID number.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Override replacement message.
        /// </summary>
        [Input("overrideReplacemsg")]
        public Input<string>? OverrideReplacemsg { get; set; }

        /// <summary>
        /// Duration of warnings.
        /// </summary>
        [Input("warnDuration")]
        public Input<string>? WarnDuration { get; set; }

        /// <summary>
        /// Re-display warning after closing browser or after a timeout. Valid values: `session`, `timeout`.
        /// </summary>
        [Input("warningDurationType")]
        public Input<string>? WarningDurationType { get; set; }

        /// <summary>
        /// Warning prompts in each category or each domain. Valid values: `per-domain`, `per-category`.
        /// </summary>
        [Input("warningPrompt")]
        public Input<string>? WarningPrompt { get; set; }

        public WebfilterProfileFtgdWfFilterArgs()
        {
        }
        public static new WebfilterProfileFtgdWfFilterArgs Empty => new WebfilterProfileFtgdWfFilterArgs();
    }
}