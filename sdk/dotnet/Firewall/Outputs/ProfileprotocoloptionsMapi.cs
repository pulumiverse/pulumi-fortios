// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Outputs
{

    [OutputType]
    public sealed class ProfileprotocoloptionsMapi
    {
        /// <summary>
        /// One or more options that can be applied to the session. Valid values: `fragmail`, `oversize`.
        /// </summary>
        public readonly string? Options;
        /// <summary>
        /// Maximum in-memory file size that can be scanned (MB).
        /// </summary>
        public readonly int? OversizeLimit;
        /// <summary>
        /// Ports to scan for content (1 - 65535, default = 110).
        /// </summary>
        public readonly int? Ports;
        /// <summary>
        /// Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ScanBzip2;
        /// <summary>
        /// Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
        /// </summary>
        public readonly int? UncompressedNestLimit;
        /// <summary>
        /// Maximum in-memory uncompressed file size that can be scanned (MB).
        /// </summary>
        public readonly int? UncompressedOversizeLimit;

        [OutputConstructor]
        private ProfileprotocoloptionsMapi(
            string? options,

            int? oversizeLimit,

            int? ports,

            string? scanBzip2,

            string? status,

            int? uncompressedNestLimit,

            int? uncompressedOversizeLimit)
        {
            Options = options;
            OversizeLimit = oversizeLimit;
            Ports = ports;
            ScanBzip2 = scanBzip2;
            Status = status;
            UncompressedNestLimit = uncompressedNestLimit;
            UncompressedOversizeLimit = uncompressedOversizeLimit;
        }
    }
}
