// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class InterfaceSecurityGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Names of user groups that can authenticate with the captive portal.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public InterfaceSecurityGroupArgs()
        {
        }
        public static new InterfaceSecurityGroupArgs Empty => new InterfaceSecurityGroupArgs();
    }
}
