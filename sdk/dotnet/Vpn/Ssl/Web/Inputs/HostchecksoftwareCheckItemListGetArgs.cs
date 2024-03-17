// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ssl.Web.Inputs
{

    public sealed class HostchecksoftwareCheckItemListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action. Valid values: `require`, `deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Hex string of MD5 checksum.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("md5s")]
        private InputList<Inputs.HostchecksoftwareCheckItemListMd5GetArgs>? _md5s;

        /// <summary>
        /// MD5 checksum. The structure of `md5s` block is documented below.
        /// 
        /// The `md5s` block supports:
        /// </summary>
        public InputList<Inputs.HostchecksoftwareCheckItemListMd5GetArgs> Md5s
        {
            get => _md5s ?? (_md5s = new InputList<Inputs.HostchecksoftwareCheckItemListMd5GetArgs>());
            set => _md5s = value;
        }

        /// <summary>
        /// Target.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// Type. Valid values: `file`, `registry`, `process`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public HostchecksoftwareCheckItemListGetArgs()
        {
        }
        public static new HostchecksoftwareCheckItemListGetArgs Empty => new HostchecksoftwareCheckItemListGetArgs();
    }
}
