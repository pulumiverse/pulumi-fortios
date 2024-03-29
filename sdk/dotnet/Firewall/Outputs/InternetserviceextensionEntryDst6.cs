// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Outputs
{

    [OutputType]
    public sealed class InternetserviceextensionEntryDst6
    {
        /// <summary>
        /// Select the destination address6 or address group object from available options.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private InternetserviceextensionEntryDst6(string? name)
        {
            Name = name;
        }
    }
}
