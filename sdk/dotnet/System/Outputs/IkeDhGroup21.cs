// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class IkeDhGroup21
    {
        /// <summary>
        /// Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
        /// </summary>
        public readonly string? KeypairCache;
        /// <summary>
        /// Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
        /// </summary>
        public readonly int? KeypairCount;
        /// <summary>
        /// Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
        /// </summary>
        public readonly string? Mode;

        [OutputConstructor]
        private IkeDhGroup21(
            string? keypairCache,

            int? keypairCount,

            string? mode)
        {
            KeypairCache = keypairCache;
            KeypairCount = keypairCount;
            Mode = mode;
        }
    }
}
