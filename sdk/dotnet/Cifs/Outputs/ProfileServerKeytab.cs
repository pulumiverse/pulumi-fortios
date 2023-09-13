// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Cifs.Outputs
{

    [OutputType]
    public sealed class ProfileServerKeytab
    {
        /// <summary>
        /// Base64 encoded keytab file containing credential of the server.
        /// </summary>
        public readonly string? Keytab;
        /// <summary>
        /// Service principal.  For example, "host/cifsserver.example.com@example.com".
        /// </summary>
        public readonly string? Principal;

        [OutputConstructor]
        private ProfileServerKeytab(
            string? keytab,

            string? principal)
        {
            Keytab = keytab;
            Principal = principal;
        }
    }
}
