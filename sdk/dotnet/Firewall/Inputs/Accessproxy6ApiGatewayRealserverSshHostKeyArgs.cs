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

    public sealed class Accessproxy6ApiGatewayRealserverSshHostKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Server host key name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Accessproxy6ApiGatewayRealserverSshHostKeyArgs()
        {
        }
        public static new Accessproxy6ApiGatewayRealserverSshHostKeyArgs Empty => new Accessproxy6ApiGatewayRealserverSshHostKeyArgs();
    }
}
