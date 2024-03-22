// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Lldp
{
    /// <summary>
    /// Configure LLDP network policy.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.System.Lldp.Networkpolicy("trname", new()
    ///     {
    ///         Comment = "test",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// SystemLldp NetworkPolicy can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/lldp/networkpolicy:Networkpolicy labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/lldp/networkpolicy:Networkpolicy labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/lldp/networkpolicy:Networkpolicy")]
    public partial class Networkpolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Guest. The structure of `guest` block is documented below.
        /// </summary>
        [Output("guest")]
        public Output<Outputs.NetworkpolicyGuest> Guest { get; private set; } = null!;

        /// <summary>
        /// Guest Voice Signaling. The structure of `guest_voice_signaling` block is documented below.
        /// </summary>
        [Output("guestVoiceSignaling")]
        public Output<Outputs.NetworkpolicyGuestVoiceSignaling> GuestVoiceSignaling { get; private set; } = null!;

        /// <summary>
        /// LLDP network policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Softphone. The structure of `softphone` block is documented below.
        /// </summary>
        [Output("softphone")]
        public Output<Outputs.NetworkpolicySoftphone> Softphone { get; private set; } = null!;

        /// <summary>
        /// Streaming Video. The structure of `streaming_video` block is documented below.
        /// </summary>
        [Output("streamingVideo")]
        public Output<Outputs.NetworkpolicyStreamingVideo> StreamingVideo { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Video Conferencing. The structure of `video_conferencing` block is documented below.
        /// </summary>
        [Output("videoConferencing")]
        public Output<Outputs.NetworkpolicyVideoConferencing> VideoConferencing { get; private set; } = null!;

        /// <summary>
        /// Video Signaling. The structure of `video_signaling` block is documented below.
        /// </summary>
        [Output("videoSignaling")]
        public Output<Outputs.NetworkpolicyVideoSignaling> VideoSignaling { get; private set; } = null!;

        /// <summary>
        /// Voice. The structure of `voice` block is documented below.
        /// </summary>
        [Output("voice")]
        public Output<Outputs.NetworkpolicyVoice> Voice { get; private set; } = null!;

        /// <summary>
        /// Voice signaling. The structure of `voice_signaling` block is documented below.
        /// </summary>
        [Output("voiceSignaling")]
        public Output<Outputs.NetworkpolicyVoiceSignaling> VoiceSignaling { get; private set; } = null!;


        /// <summary>
        /// Create a Networkpolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Networkpolicy(string name, NetworkpolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/lldp/networkpolicy:Networkpolicy", name, args ?? new NetworkpolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Networkpolicy(string name, Input<string> id, NetworkpolicyState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/lldp/networkpolicy:Networkpolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Networkpolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Networkpolicy Get(string name, Input<string> id, NetworkpolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Networkpolicy(name, id, state, options);
        }
    }

    public sealed class NetworkpolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Guest. The structure of `guest` block is documented below.
        /// </summary>
        [Input("guest")]
        public Input<Inputs.NetworkpolicyGuestArgs>? Guest { get; set; }

        /// <summary>
        /// Guest Voice Signaling. The structure of `guest_voice_signaling` block is documented below.
        /// </summary>
        [Input("guestVoiceSignaling")]
        public Input<Inputs.NetworkpolicyGuestVoiceSignalingArgs>? GuestVoiceSignaling { get; set; }

        /// <summary>
        /// LLDP network policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Softphone. The structure of `softphone` block is documented below.
        /// </summary>
        [Input("softphone")]
        public Input<Inputs.NetworkpolicySoftphoneArgs>? Softphone { get; set; }

        /// <summary>
        /// Streaming Video. The structure of `streaming_video` block is documented below.
        /// </summary>
        [Input("streamingVideo")]
        public Input<Inputs.NetworkpolicyStreamingVideoArgs>? StreamingVideo { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Video Conferencing. The structure of `video_conferencing` block is documented below.
        /// </summary>
        [Input("videoConferencing")]
        public Input<Inputs.NetworkpolicyVideoConferencingArgs>? VideoConferencing { get; set; }

        /// <summary>
        /// Video Signaling. The structure of `video_signaling` block is documented below.
        /// </summary>
        [Input("videoSignaling")]
        public Input<Inputs.NetworkpolicyVideoSignalingArgs>? VideoSignaling { get; set; }

        /// <summary>
        /// Voice. The structure of `voice` block is documented below.
        /// </summary>
        [Input("voice")]
        public Input<Inputs.NetworkpolicyVoiceArgs>? Voice { get; set; }

        /// <summary>
        /// Voice signaling. The structure of `voice_signaling` block is documented below.
        /// </summary>
        [Input("voiceSignaling")]
        public Input<Inputs.NetworkpolicyVoiceSignalingArgs>? VoiceSignaling { get; set; }

        public NetworkpolicyArgs()
        {
        }
        public static new NetworkpolicyArgs Empty => new NetworkpolicyArgs();
    }

    public sealed class NetworkpolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Guest. The structure of `guest` block is documented below.
        /// </summary>
        [Input("guest")]
        public Input<Inputs.NetworkpolicyGuestGetArgs>? Guest { get; set; }

        /// <summary>
        /// Guest Voice Signaling. The structure of `guest_voice_signaling` block is documented below.
        /// </summary>
        [Input("guestVoiceSignaling")]
        public Input<Inputs.NetworkpolicyGuestVoiceSignalingGetArgs>? GuestVoiceSignaling { get; set; }

        /// <summary>
        /// LLDP network policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Softphone. The structure of `softphone` block is documented below.
        /// </summary>
        [Input("softphone")]
        public Input<Inputs.NetworkpolicySoftphoneGetArgs>? Softphone { get; set; }

        /// <summary>
        /// Streaming Video. The structure of `streaming_video` block is documented below.
        /// </summary>
        [Input("streamingVideo")]
        public Input<Inputs.NetworkpolicyStreamingVideoGetArgs>? StreamingVideo { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Video Conferencing. The structure of `video_conferencing` block is documented below.
        /// </summary>
        [Input("videoConferencing")]
        public Input<Inputs.NetworkpolicyVideoConferencingGetArgs>? VideoConferencing { get; set; }

        /// <summary>
        /// Video Signaling. The structure of `video_signaling` block is documented below.
        /// </summary>
        [Input("videoSignaling")]
        public Input<Inputs.NetworkpolicyVideoSignalingGetArgs>? VideoSignaling { get; set; }

        /// <summary>
        /// Voice. The structure of `voice` block is documented below.
        /// </summary>
        [Input("voice")]
        public Input<Inputs.NetworkpolicyVoiceGetArgs>? Voice { get; set; }

        /// <summary>
        /// Voice signaling. The structure of `voice_signaling` block is documented below.
        /// </summary>
        [Input("voiceSignaling")]
        public Input<Inputs.NetworkpolicyVoiceSignalingGetArgs>? VoiceSignaling { get; set; }

        public NetworkpolicyState()
        {
        }
        public static new NetworkpolicyState Empty => new NetworkpolicyState();
    }
}
