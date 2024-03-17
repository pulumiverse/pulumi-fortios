// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Dhcp.Outputs
{

    [OutputType]
    public sealed class GetServerOptionVciStringResult
    {
        /// <summary>
        /// VCI strings.
        /// </summary>
        public readonly string VciString;

        [OutputConstructor]
        private GetServerOptionVciStringResult(string vciString)
        {
            VciString = vciString;
        }
    }
}
