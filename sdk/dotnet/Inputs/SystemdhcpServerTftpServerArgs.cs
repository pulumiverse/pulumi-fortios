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

    public sealed class SystemdhcpServerTftpServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// TFTP server.
        /// </summary>
        [Input("tftpServer")]
        public Input<string>? TftpServer { get; set; }

        public SystemdhcpServerTftpServerArgs()
        {
        }
        public static new SystemdhcpServerTftpServerArgs Empty => new SystemdhcpServerTftpServerArgs();
    }
}
