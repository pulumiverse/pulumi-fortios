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
    public sealed class GetProxypolicyInternetServiceIdResult
    {
        /// <summary>
        /// Internet Service ID.
        /// </summary>
        public readonly int Id;

        [OutputConstructor]
        private GetProxypolicyInternetServiceIdResult(int id)
        {
            Id = id;
        }
    }
}
