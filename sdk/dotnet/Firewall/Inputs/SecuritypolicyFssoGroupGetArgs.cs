// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class SecuritypolicyFssoGroupGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Names of FSSO groups.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SecuritypolicyFssoGroupGetArgs()
        {
        }
        public static new SecuritypolicyFssoGroupGetArgs Empty => new SecuritypolicyFssoGroupGetArgs();
    }
}
