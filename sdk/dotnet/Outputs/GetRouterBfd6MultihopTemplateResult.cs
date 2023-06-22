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
    public sealed class GetRouterBfd6MultihopTemplateResult
    {
        /// <summary>
        /// Authentication mode.
        /// </summary>
        public readonly string AuthMode;
        /// <summary>
        /// BFD desired minimal transmit interval (milliseconds).
        /// </summary>
        public readonly int BfdDesiredMinTx;
        /// <summary>
        /// BFD detection multiplier.
        /// </summary>
        public readonly int BfdDetectMult;
        /// <summary>
        /// BFD required minimal receive interval (milliseconds).
        /// </summary>
        public readonly int BfdRequiredMinRx;
        /// <summary>
        /// Destination prefix.
        /// </summary>
        public readonly string Dst;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// MD5 key of key ID 1.
        /// </summary>
        public readonly string Md5Key;
        /// <summary>
        /// Source prefix.
        /// </summary>
        public readonly string Src;

        [OutputConstructor]
        private GetRouterBfd6MultihopTemplateResult(
            string authMode,

            int bfdDesiredMinTx,

            int bfdDetectMult,

            int bfdRequiredMinRx,

            string dst,

            int id,

            string md5Key,

            string src)
        {
            AuthMode = authMode;
            BfdDesiredMinTx = bfdDesiredMinTx;
            BfdDetectMult = bfdDetectMult;
            BfdRequiredMinRx = bfdRequiredMinRx;
            Dst = dst;
            Id = id;
            Md5Key = md5Key;
            Src = src;
        }
    }
}
