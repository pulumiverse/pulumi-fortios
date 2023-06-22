// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class EndpointcontrolProfileForticlientWinmacSettingsForticlientRunningApp
    {
        /// <summary>
        /// Application name.
        /// </summary>
        public readonly string? AppName;
        /// <summary>
        /// App's SHA256 signature.
        /// </summary>
        public readonly string? AppSha256Signature;
        /// <summary>
        /// App's SHA256 Signature.
        /// </summary>
        public readonly string? AppSha256Signature2;
        /// <summary>
        /// App's SHA256 Signature.
        /// </summary>
        public readonly string? AppSha256Signature3;
        /// <summary>
        /// App's SHA256 Signature.
        /// </summary>
        public readonly string? AppSha256Signature4;
        /// <summary>
        /// Application check rule. Valid values: `present`, `absent`.
        /// </summary>
        public readonly string? ApplicationCheckRule;
        /// <summary>
        /// Application ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Process name.
        /// </summary>
        public readonly string? ProcessName;
        /// <summary>
        /// Process name.
        /// </summary>
        public readonly string? ProcessName2;
        /// <summary>
        /// Process name.
        /// </summary>
        public readonly string? ProcessName3;
        /// <summary>
        /// Process name.
        /// </summary>
        public readonly string? ProcessName4;

        [OutputConstructor]
        private EndpointcontrolProfileForticlientWinmacSettingsForticlientRunningApp(
            string? appName,

            string? appSha256Signature,

            string? appSha256Signature2,

            string? appSha256Signature3,

            string? appSha256Signature4,

            string? applicationCheckRule,

            int? id,

            string? processName,

            string? processName2,

            string? processName3,

            string? processName4)
        {
            AppName = appName;
            AppSha256Signature = appSha256Signature;
            AppSha256Signature2 = appSha256Signature2;
            AppSha256Signature3 = appSha256Signature3;
            AppSha256Signature4 = appSha256Signature4;
            ApplicationCheckRule = applicationCheckRule;
            Id = id;
            ProcessName = processName;
            ProcessName2 = processName2;
            ProcessName3 = processName3;
            ProcessName4 = processName4;
        }
    }
}
