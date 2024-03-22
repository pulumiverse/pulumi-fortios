// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Antivirus.Inputs
{

    public sealed class ProfilePop3GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Select the archive types to block.
        /// </summary>
        [Input("archiveBlock")]
        public Input<string>? ArchiveBlock { get; set; }

        /// <summary>
        /// Select the archive types to log.
        /// </summary>
        [Input("archiveLog")]
        public Input<string>? ArchiveLog { get; set; }

        /// <summary>
        /// Enable AntiVirus scan service. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        [Input("avScan")]
        public Input<string>? AvScan { get; set; }

        /// <summary>
        /// AV Content Disarm and Reconstruction settings. The structure of `content_disarm` block is documented below.
        /// </summary>
        [Input("contentDisarm")]
        public Input<string>? ContentDisarm { get; set; }

        /// <summary>
        /// Enable/disable the virus emulator. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("emulator")]
        public Input<string>? Emulator { get; set; }

        /// <summary>
        /// Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.
        /// </summary>
        [Input("executables")]
        public Input<string>? Executables { get; set; }

        /// <summary>
        /// One or more external malware block lists. The structure of `external_blocklist` block is documented below.
        /// </summary>
        [Input("externalBlocklist")]
        public Input<string>? ExternalBlocklist { get; set; }

        /// <summary>
        /// Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        [Input("fortiai")]
        public Input<string>? Fortiai { get; set; }

        /// <summary>
        /// Enable/disable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        [Input("fortindr")]
        public Input<string>? Fortindr { get; set; }

        /// <summary>
        /// Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        [Input("fortisandbox")]
        public Input<string>? Fortisandbox { get; set; }

        /// <summary>
        /// Enable/disable CIFS AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Configure Virus Outbreak Prevention settings. The structure of `outbreak_prevention` block is documented below.
        /// </summary>
        [Input("outbreakPrevention")]
        public Input<string>? OutbreakPrevention { get; set; }

        /// <summary>
        /// Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("quarantine")]
        public Input<string>? Quarantine { get; set; }

        public ProfilePop3GetArgs()
        {
        }
        public static new ProfilePop3GetArgs Empty => new ProfilePop3GetArgs();
    }
}
