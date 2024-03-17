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
    public sealed class GetReplacemsggroupEcResult
    {
        /// <summary>
        /// Message string.
        /// </summary>
        public readonly string Buffer;
        /// <summary>
        /// Format flag.
        /// </summary>
        public readonly string Format;
        /// <summary>
        /// Header flag.
        /// </summary>
        public readonly string Header;
        /// <summary>
        /// Message type.
        /// </summary>
        public readonly string MsgType;

        [OutputConstructor]
        private GetReplacemsggroupEcResult(
            string buffer,

            string format,

            string header,

            string msgType)
        {
            Buffer = buffer;
            Format = format;
            Header = header;
            MsgType = msgType;
        }
    }
}
