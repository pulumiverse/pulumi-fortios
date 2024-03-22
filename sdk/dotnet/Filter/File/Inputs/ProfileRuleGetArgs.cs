// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.File.Inputs
{

    public sealed class ProfileRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action taken for matched file. Valid values: `log-only`, `block`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Traffic direction. On FortiOS versions 6.4.1-7.4.1: HTTP, FTP, SSH, CIFS only. On FortiOS versions &gt;= 7.4.2: HTTP, FTP, SSH, CIFS, and MAPI only. Valid values: `incoming`, `outgoing`, `any`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        [Input("fileTypes")]
        private InputList<Inputs.ProfileRuleFileTypeGetArgs>? _fileTypes;

        /// <summary>
        /// Select file type. The structure of `file_type` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileRuleFileTypeGetArgs> FileTypes
        {
            get => _fileTypes ?? (_fileTypes = new InputList<Inputs.ProfileRuleFileTypeGetArgs>());
            set => _fileTypes = value;
        }

        /// <summary>
        /// File-filter rule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Match password-protected files. Valid values: `yes`, `any`.
        /// </summary>
        [Input("passwordProtected")]
        public Input<string>? PasswordProtected { get; set; }

        /// <summary>
        /// Protocols to apply rule to. Valid values: `http`, `ftp`, `smtp`, `imap`, `pop3`, `mapi`, `cifs`, `ssh`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public ProfileRuleGetArgs()
        {
        }
        public static new ProfileRuleGetArgs Empty => new ProfileRuleGetArgs();
    }
}
