// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Hotspot20.Outputs
{

    [OutputType]
    public sealed class H2qpadviceofchargeAocListPlanInfo
    {
        /// <summary>
        /// Currency code.
        /// </summary>
        public readonly string? Currency;
        /// <summary>
        /// Info file.
        /// </summary>
        public readonly string? InfoFile;
        /// <summary>
        /// Languague code.
        /// </summary>
        public readonly string? Lang;
        /// <summary>
        /// Plan name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private H2qpadviceofchargeAocListPlanInfo(
            string? currency,

            string? infoFile,

            string? lang,

            string? name)
        {
            Currency = currency;
            InfoFile = infoFile;
            Lang = lang;
            Name = name;
        }
    }
}
