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

    public sealed class AcmeInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interfaceName")]
        public Input<string>? InterfaceName { get; set; }

        public AcmeInterfaceArgs()
        {
        }
        public static new AcmeInterfaceArgs Empty => new AcmeInterfaceArgs();
    }
}
