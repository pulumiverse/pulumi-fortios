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

    public sealed class SystemNtpInterfaceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interfaceName")]
        public Input<string>? InterfaceName { get; set; }

        public SystemNtpInterfaceGetArgs()
        {
        }
        public static new SystemNtpInterfaceGetArgs Empty => new SystemNtpInterfaceGetArgs();
    }
}
