// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class SystemIpamRuleDeviceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fortigate serial number or wildcard.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SystemIpamRuleDeviceGetArgs()
        {
        }
        public static new SystemIpamRuleDeviceGetArgs Empty => new SystemIpamRuleDeviceGetArgs();
    }
}