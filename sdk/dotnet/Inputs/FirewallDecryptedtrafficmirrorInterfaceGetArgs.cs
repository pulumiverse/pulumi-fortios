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

    public sealed class FirewallDecryptedtrafficmirrorInterfaceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Decrypted traffic mirror interface.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallDecryptedtrafficmirrorInterfaceGetArgs()
        {
        }
        public static new FirewallDecryptedtrafficmirrorInterfaceGetArgs Empty => new FirewallDecryptedtrafficmirrorInterfaceGetArgs();
    }
}
