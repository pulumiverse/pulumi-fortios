// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Diameterfilter
{
    /// <summary>
    /// Configure Diameter filter profiles. Applies to FortiOS Version `&gt;= 7.4.2`.
    /// 
    /// ## Import
    /// 
    /// DiameterFilter Profile can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:diameterfilter/profile:Profile labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:diameterfilter/profile:Profile labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:diameterfilter/profile:Profile")]
    public partial class Profile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action to be taken for messages with cmd flag reserve bits set. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Output("cmdFlagsReserveSet")]
        public Output<string> CmdFlagsReserveSet { get; private set; } = null!;

        /// <summary>
        /// Action to be taken for messages with invalid command code. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Output("commandCodeInvalid")]
        public Output<string> CommandCodeInvalid { get; private set; } = null!;

        /// <summary>
        /// Valid range for command codes (0-16777215).
        /// </summary>
        [Output("commandCodeRange")]
        public Output<string> CommandCodeRange { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Enable/disable packet log for triggered diameter settings. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("logPacket")]
        public Output<string> LogPacket { get; private set; } = null!;

        /// <summary>
        /// Action to be taken for invalid message length. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Output("messageLengthInvalid")]
        public Output<string> MessageLengthInvalid { get; private set; } = null!;

        /// <summary>
        /// Action to be taken for answers without corresponding request. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Output("missingRequestAction")]
        public Output<string> MissingRequestAction { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging for all User Name and Result Code AVP messages. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("monitorAllMessages")]
        public Output<string> MonitorAllMessages { get; private set; } = null!;

        /// <summary>
        /// Profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Action to be taken for invalid protocol version. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Output("protocolVersionInvalid")]
        public Output<string> ProtocolVersionInvalid { get; private set; } = null!;

        /// <summary>
        /// Action to be taken for request messages with error flag set. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Output("requestErrorFlagSet")]
        public Output<string> RequestErrorFlagSet { get; private set; } = null!;

        /// <summary>
        /// Enable/disable validation that each answer has a corresponding request. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("trackRequestsAnswers")]
        public Output<string> TrackRequestsAnswers { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Profile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Profile(string name, ProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:diameterfilter/profile:Profile", name, args ?? new ProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Profile(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:diameterfilter/profile:Profile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Profile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Profile Get(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new Profile(name, id, state, options);
        }
    }

    public sealed class ProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to be taken for messages with cmd flag reserve bits set. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("cmdFlagsReserveSet")]
        public Input<string>? CmdFlagsReserveSet { get; set; }

        /// <summary>
        /// Action to be taken for messages with invalid command code. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("commandCodeInvalid")]
        public Input<string>? CommandCodeInvalid { get; set; }

        /// <summary>
        /// Valid range for command codes (0-16777215).
        /// </summary>
        [Input("commandCodeRange")]
        public Input<string>? CommandCodeRange { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Enable/disable packet log for triggered diameter settings. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logPacket")]
        public Input<string>? LogPacket { get; set; }

        /// <summary>
        /// Action to be taken for invalid message length. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("messageLengthInvalid")]
        public Input<string>? MessageLengthInvalid { get; set; }

        /// <summary>
        /// Action to be taken for answers without corresponding request. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("missingRequestAction")]
        public Input<string>? MissingRequestAction { get; set; }

        /// <summary>
        /// Enable/disable logging for all User Name and Result Code AVP messages. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("monitorAllMessages")]
        public Input<string>? MonitorAllMessages { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Action to be taken for invalid protocol version. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("protocolVersionInvalid")]
        public Input<string>? ProtocolVersionInvalid { get; set; }

        /// <summary>
        /// Action to be taken for request messages with error flag set. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("requestErrorFlagSet")]
        public Input<string>? RequestErrorFlagSet { get; set; }

        /// <summary>
        /// Enable/disable validation that each answer has a corresponding request. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("trackRequestsAnswers")]
        public Input<string>? TrackRequestsAnswers { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ProfileArgs()
        {
        }
        public static new ProfileArgs Empty => new ProfileArgs();
    }

    public sealed class ProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to be taken for messages with cmd flag reserve bits set. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("cmdFlagsReserveSet")]
        public Input<string>? CmdFlagsReserveSet { get; set; }

        /// <summary>
        /// Action to be taken for messages with invalid command code. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("commandCodeInvalid")]
        public Input<string>? CommandCodeInvalid { get; set; }

        /// <summary>
        /// Valid range for command codes (0-16777215).
        /// </summary>
        [Input("commandCodeRange")]
        public Input<string>? CommandCodeRange { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Enable/disable packet log for triggered diameter settings. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logPacket")]
        public Input<string>? LogPacket { get; set; }

        /// <summary>
        /// Action to be taken for invalid message length. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("messageLengthInvalid")]
        public Input<string>? MessageLengthInvalid { get; set; }

        /// <summary>
        /// Action to be taken for answers without corresponding request. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("missingRequestAction")]
        public Input<string>? MissingRequestAction { get; set; }

        /// <summary>
        /// Enable/disable logging for all User Name and Result Code AVP messages. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("monitorAllMessages")]
        public Input<string>? MonitorAllMessages { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Action to be taken for invalid protocol version. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("protocolVersionInvalid")]
        public Input<string>? ProtocolVersionInvalid { get; set; }

        /// <summary>
        /// Action to be taken for request messages with error flag set. Valid values: `allow`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("requestErrorFlagSet")]
        public Input<string>? RequestErrorFlagSet { get; set; }

        /// <summary>
        /// Enable/disable validation that each answer has a corresponding request. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("trackRequestsAnswers")]
        public Input<string>? TrackRequestsAnswers { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ProfileState()
        {
        }
        public static new ProfileState Empty => new ProfileState();
    }
}