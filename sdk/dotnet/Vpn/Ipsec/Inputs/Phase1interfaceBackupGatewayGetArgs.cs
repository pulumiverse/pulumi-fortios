// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ipsec.Inputs
{

    public sealed class Phase1interfaceBackupGatewayGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address of backup gateway.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        public Phase1interfaceBackupGatewayGetArgs()
        {
        }
        public static new Phase1interfaceBackupGatewayGetArgs Empty => new Phase1interfaceBackupGatewayGetArgs();
    }
}