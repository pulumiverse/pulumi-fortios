// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Aspyrmedia.Fortios.Antivirus.Outputs
{

    [OutputType]
    public sealed class AntivirusProfileOutbreakPrevention
    {
        public readonly string? ExternalBlocklist;
        public readonly string? FtgdService;

        [OutputConstructor]
        private AntivirusProfileOutbreakPrevention(
            string? externalBlocklist,

            string? ftgdService)
        {
            ExternalBlocklist = externalBlocklist;
            FtgdService = ftgdService;
        }
    }
}
