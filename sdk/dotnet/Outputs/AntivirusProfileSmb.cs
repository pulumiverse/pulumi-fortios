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
    public sealed class AntivirusProfileSmb
    {
        /// <summary>
        /// Select the archive types to block. Valid values: `encrypted`, `corrupted`, `partiallycorrupted`, `multipart`, `nested`, `mailbomb`, `fileslimit`, `timeout`, `unhandled`.
        /// </summary>
        public readonly string? ArchiveBlock;
        /// <summary>
        /// Select the archive types to log. Valid values: `encrypted`, `corrupted`, `partiallycorrupted`, `multipart`, `nested`, `mailbomb`, `fileslimit`, `timeout`, `unhandled`.
        /// </summary>
        public readonly string? ArchiveLog;
        /// <summary>
        /// Enable/disable the virus emulator. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Emulator;
        /// <summary>
        /// Enable/disable SMB AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
        /// </summary>
        public readonly string? Options;
        /// <summary>
        /// Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`.
        /// </summary>
        public readonly string? OutbreakPrevention;

        [OutputConstructor]
        private AntivirusProfileSmb(
            string? archiveBlock,

            string? archiveLog,

            string? emulator,

            string? options,

            string? outbreakPrevention)
        {
            ArchiveBlock = archiveBlock;
            ArchiveLog = archiveLog;
            Emulator = emulator;
            Options = options;
            OutbreakPrevention = outbreakPrevention;
        }
    }
}
