// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Outputs
{

    [OutputType]
    public sealed class GetVirtualwanlinkHealthCheckMemberResult
    {
        /// <summary>
        /// Member sequence number.
        /// </summary>
        public readonly int SeqNum;

        [OutputConstructor]
        private GetVirtualwanlinkHealthCheckMemberResult(int seqNum)
        {
            SeqNum = seqNum;
        }
    }
}
