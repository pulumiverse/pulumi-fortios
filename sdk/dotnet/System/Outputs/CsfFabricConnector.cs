// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class CsfFabricConnector
    {
        /// <summary>
        /// Override access profile.
        /// </summary>
        public readonly string? Accprofile;
        /// <summary>
        /// Enable/disable downstream device write access to configuration. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ConfigurationWriteAccess;
        /// <summary>
        /// Serial.
        /// </summary>
        public readonly string? Serial;
        /// <summary>
        /// Virtual domains that the connector has access to. If none are set, the connector will only have access to the VDOM that it joins the Security Fabric through. The structure of `vdom` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.CsfFabricConnectorVdom> Vdoms;

        [OutputConstructor]
        private CsfFabricConnector(
            string? accprofile,

            string? configurationWriteAccess,

            string? serial,

            ImmutableArray<Outputs.CsfFabricConnectorVdom> vdoms)
        {
            Accprofile = accprofile;
            ConfigurationWriteAccess = configurationWriteAccess;
            Serial = serial;
            Vdoms = vdoms;
        }
    }
}
