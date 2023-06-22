// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class AntivirusProfileImap
    {
        /// <summary>
        /// Select the archive types to block.
        /// </summary>
        public readonly string? ArchiveBlock;
        /// <summary>
        /// Select the archive types to log.
        /// </summary>
        public readonly string? ArchiveLog;
        /// <summary>
        /// Enable AntiVirus scan service. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        public readonly string? AvScan;
        /// <summary>
        /// Enable Content Disarm and Reconstruction for this protocol. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? ContentDisarm;
        /// <summary>
        /// Enable/disable the virus emulator. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Emulator;
        /// <summary>
        /// Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.
        /// </summary>
        public readonly string? Executables;
        /// <summary>
        /// Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        public readonly string? ExternalBlocklist;
        /// <summary>
        /// Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        public readonly string? Fortiai;
        /// <summary>
        /// Enable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        public readonly string? Fortindr;
        /// <summary>
        /// Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        public readonly string? Fortisandbox;
        /// <summary>
        /// Enable/disable POP3 AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
        /// </summary>
        public readonly string? Options;
        /// <summary>
        /// Enable Virus Outbreak Prevention service.
        /// </summary>
        public readonly string? OutbreakPrevention;
        /// <summary>
        /// Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Quarantine;

        [OutputConstructor]
        private AntivirusProfileImap(
            string? archiveBlock,

            string? archiveLog,

            string? avScan,

            string? contentDisarm,

            string? emulator,

            string? executables,

            string? externalBlocklist,

            string? fortiai,

            string? fortindr,

            string? fortisandbox,

            string? options,

            string? outbreakPrevention,

            string? quarantine)
        {
            ArchiveBlock = archiveBlock;
            ArchiveLog = archiveLog;
            AvScan = avScan;
            ContentDisarm = contentDisarm;
            Emulator = emulator;
            Executables = executables;
            ExternalBlocklist = externalBlocklist;
            Fortiai = fortiai;
            Fortindr = fortindr;
            Fortisandbox = fortisandbox;
            Options = options;
            OutbreakPrevention = outbreakPrevention;
            Quarantine = quarantine;
        }
    }
}
