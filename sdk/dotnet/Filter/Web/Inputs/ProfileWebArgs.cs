// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Web.Inputs
{

    public sealed class ProfileWebArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FortiGuard allowlist settings. Valid values: `exempt-av`, `exempt-webcontent`, `exempt-activex-java-cookie`, `exempt-dlp`, `exempt-rangeblock`, `extended-log-others`.
        /// </summary>
        [Input("allowlist")]
        public Input<string>? Allowlist { get; set; }

        /// <summary>
        /// Enable/disable automatic addition of URLs detected by FortiSandbox to blacklist. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("blacklist")]
        public Input<string>? Blacklist { get; set; }

        /// <summary>
        /// Enable/disable automatic addition of URLs detected by FortiSandbox to blocklist. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("blocklist")]
        public Input<string>? Blocklist { get; set; }

        /// <summary>
        /// Banned word table ID.
        /// </summary>
        [Input("bwordTable")]
        public Input<int>? BwordTable { get; set; }

        /// <summary>
        /// Banned word score threshold.
        /// </summary>
        [Input("bwordThreshold")]
        public Input<int>? BwordThreshold { get; set; }

        /// <summary>
        /// Content header list.
        /// </summary>
        [Input("contentHeaderList")]
        public Input<int>? ContentHeaderList { get; set; }

        [Input("keywordMatches")]
        private InputList<Inputs.ProfileWebKeywordMatchArgs>? _keywordMatches;

        /// <summary>
        /// Search keywords to log when match is found. The structure of `keyword_match` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileWebKeywordMatchArgs> KeywordMatches
        {
            get => _keywordMatches ?? (_keywordMatches = new InputList<Inputs.ProfileWebKeywordMatchArgs>());
            set => _keywordMatches = value;
        }

        /// <summary>
        /// Enable/disable logging all search phrases. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logSearch")]
        public Input<string>? LogSearch { get; set; }

        /// <summary>
        /// Safe search type. Valid values: `url`, `header`.
        /// </summary>
        [Input("safeSearch")]
        public Input<string>? SafeSearch { get; set; }

        /// <summary>
        /// URL filter table ID.
        /// </summary>
        [Input("urlfilterTable")]
        public Input<int>? UrlfilterTable { get; set; }

        /// <summary>
        /// Set Vimeo-restrict ("7" = don't show mature content, "134" = don't show unrated and mature content). A value of cookie "content_rating".
        /// </summary>
        [Input("vimeoRestrict")]
        public Input<string>? VimeoRestrict { get; set; }

        /// <summary>
        /// FortiGuard whitelist settings. Valid values: `exempt-av`, `exempt-webcontent`, `exempt-activex-java-cookie`, `exempt-dlp`, `exempt-rangeblock`, `extended-log-others`.
        /// </summary>
        [Input("whitelist")]
        public Input<string>? Whitelist { get; set; }

        /// <summary>
        /// YouTube EDU filter level. Valid values: `none`, `strict`, `moderate`.
        /// </summary>
        [Input("youtubeRestrict")]
        public Input<string>? YoutubeRestrict { get; set; }

        public ProfileWebArgs()
        {
        }
        public static new ProfileWebArgs Empty => new ProfileWebArgs();
    }
}
