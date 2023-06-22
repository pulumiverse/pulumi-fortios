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
    public sealed class SystemAcmeAccount
    {
        /// <summary>
        /// Account ca_url.
        /// </summary>
        public readonly string? CaUrl;
        /// <summary>
        /// Account email.
        /// </summary>
        public readonly string? Email;
        /// <summary>
        /// Account id.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Account Private Key.
        /// </summary>
        public readonly string? Privatekey;
        /// <summary>
        /// Account status.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Account url.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private SystemAcmeAccount(
            string? caUrl,

            string? email,

            string? id,

            string? privatekey,

            string? status,

            string? url)
        {
            CaUrl = caUrl;
            Email = email;
            Id = id;
            Privatekey = privatekey;
            Status = status;
            Url = url;
        }
    }
}
