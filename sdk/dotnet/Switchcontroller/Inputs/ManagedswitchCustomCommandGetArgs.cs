// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Inputs
{

    public sealed class ManagedswitchCustomCommandGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// List of FortiSwitch commands.
        /// </summary>
        [Input("commandEntry")]
        public Input<string>? CommandEntry { get; set; }

        /// <summary>
        /// Names of commands to be pushed to this FortiSwitch device, as configured under config switch-controller custom-command.
        /// </summary>
        [Input("commandName")]
        public Input<string>? CommandName { get; set; }

        public ManagedswitchCustomCommandGetArgs()
        {
        }
        public static new ManagedswitchCustomCommandGetArgs Empty => new ManagedswitchCustomCommandGetArgs();
    }
}