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
    public sealed class StandaloneclusterClusterPeerSyncvd
    {
        /// <summary>
        /// VDOM name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private StandaloneclusterClusterPeerSyncvd(string? name)
        {
            Name = name;
        }
    }
}
