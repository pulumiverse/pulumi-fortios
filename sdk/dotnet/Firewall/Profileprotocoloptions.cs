// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    /// <summary>
    /// Configure protocol options.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.Firewall.Profileprotocoloptions("trname", new()
    ///     {
    ///         Dns = new Fortios.Firewall.Inputs.ProfileprotocoloptionsDnsArgs
    ///         {
    ///             Ports = 53,
    ///             Status = "enable",
    ///         },
    ///         Ftp = new Fortios.Firewall.Inputs.ProfileprotocoloptionsFtpArgs
    ///         {
    ///             ComfortAmount = 1,
    ///             ComfortInterval = 10,
    ///             InspectAll = "disable",
    ///             Options = "splice",
    ///             OversizeLimit = 10,
    ///             Ports = 21,
    ///             ScanBzip2 = "enable",
    ///             Status = "enable",
    ///             UncompressedNestLimit = 12,
    ///             UncompressedOversizeLimit = 10,
    ///         },
    ///         Http = new Fortios.Firewall.Inputs.ProfileprotocoloptionsHttpArgs
    ///         {
    ///             BlockPageStatusCode = 403,
    ///             ComfortAmount = 1,
    ///             ComfortInterval = 10,
    ///             FortinetBar = "disable",
    ///             FortinetBarPort = 8011,
    ///             HttpPolicy = "disable",
    ///             InspectAll = "disable",
    ///             OversizeLimit = 10,
    ///             Ports = 80,
    ///             RangeBlock = "disable",
    ///             RetryCount = 0,
    ///             ScanBzip2 = "enable",
    ///             Status = "enable",
    ///             StreamingContentBypass = "enable",
    ///             StripXForwardedFor = "disable",
    ///             SwitchingProtocols = "bypass",
    ///             UncompressedNestLimit = 12,
    ///             UncompressedOversizeLimit = 10,
    ///         },
    ///         Imap = new Fortios.Firewall.Inputs.ProfileprotocoloptionsImapArgs
    ///         {
    ///             InspectAll = "disable",
    ///             Options = "fragmail",
    ///             OversizeLimit = 10,
    ///             Ports = 143,
    ///             ScanBzip2 = "enable",
    ///             Status = "enable",
    ///             UncompressedNestLimit = 12,
    ///             UncompressedOversizeLimit = 10,
    ///         },
    ///         MailSignature = new Fortios.Firewall.Inputs.ProfileprotocoloptionsMailSignatureArgs
    ///         {
    ///             Status = "disable",
    ///         },
    ///         Mapi = new Fortios.Firewall.Inputs.ProfileprotocoloptionsMapiArgs
    ///         {
    ///             Options = "fragmail",
    ///             OversizeLimit = 10,
    ///             Ports = 135,
    ///             ScanBzip2 = "enable",
    ///             Status = "enable",
    ///             UncompressedNestLimit = 12,
    ///             UncompressedOversizeLimit = 10,
    ///         },
    ///         Nntp = new Fortios.Firewall.Inputs.ProfileprotocoloptionsNntpArgs
    ///         {
    ///             InspectAll = "disable",
    ///             Options = "splice",
    ///             OversizeLimit = 10,
    ///             Ports = 119,
    ///             ScanBzip2 = "enable",
    ///             Status = "enable",
    ///             UncompressedNestLimit = 12,
    ///             UncompressedOversizeLimit = 10,
    ///         },
    ///         OversizeLog = "disable",
    ///         Pop3 = new Fortios.Firewall.Inputs.ProfileprotocoloptionsPop3Args
    ///         {
    ///             InspectAll = "disable",
    ///             Options = "fragmail",
    ///             OversizeLimit = 10,
    ///             Ports = 110,
    ///             ScanBzip2 = "enable",
    ///             Status = "enable",
    ///             UncompressedNestLimit = 12,
    ///             UncompressedOversizeLimit = 10,
    ///         },
    ///         RpcOverHttp = "disable",
    ///         Smtp = new Fortios.Firewall.Inputs.ProfileprotocoloptionsSmtpArgs
    ///         {
    ///             InspectAll = "disable",
    ///             Options = "fragmail splice",
    ///             OversizeLimit = 10,
    ///             Ports = 25,
    ///             ScanBzip2 = "enable",
    ///             ServerBusy = "disable",
    ///             Status = "enable",
    ///             UncompressedNestLimit = 12,
    ///             UncompressedOversizeLimit = 10,
    ///         },
    ///         SwitchingProtocolsLog = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall ProfileProtocolOptions can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/profileprotocoloptions:Profileprotocoloptions labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/profileprotocoloptions:Profileprotocoloptions labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/profileprotocoloptions:Profileprotocoloptions")]
    public partial class Profileprotocoloptions : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configure CIFS protocol options. The structure of `cifs` block is documented below.
        /// </summary>
        [Output("cifs")]
        public Output<Outputs.ProfileprotocoloptionsCifs> Cifs { get; private set; } = null!;

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Configure DNS protocol options. The structure of `dns` block is documented below.
        /// </summary>
        [Output("dns")]
        public Output<Outputs.ProfileprotocoloptionsDns> Dns { get; private set; } = null!;

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Output("featureSet")]
        public Output<string> FeatureSet { get; private set; } = null!;

        /// <summary>
        /// Configure FTP protocol options. The structure of `ftp` block is documented below.
        /// </summary>
        [Output("ftp")]
        public Output<Outputs.ProfileprotocoloptionsFtp> Ftp { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Configure HTTP protocol options. The structure of `http` block is documented below.
        /// </summary>
        [Output("http")]
        public Output<Outputs.ProfileprotocoloptionsHttp> Http { get; private set; } = null!;

        /// <summary>
        /// Configure IMAP protocol options. The structure of `imap` block is documented below.
        /// </summary>
        [Output("imap")]
        public Output<Outputs.ProfileprotocoloptionsImap> Imap { get; private set; } = null!;

        /// <summary>
        /// Configure Mail signature. The structure of `mail_signature` block is documented below.
        /// </summary>
        [Output("mailSignature")]
        public Output<Outputs.ProfileprotocoloptionsMailSignature> MailSignature { get; private set; } = null!;

        /// <summary>
        /// Configure MAPI protocol options. The structure of `mapi` block is documented below.
        /// </summary>
        [Output("mapi")]
        public Output<Outputs.ProfileprotocoloptionsMapi> Mapi { get; private set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configure NNTP protocol options. The structure of `nntp` block is documented below.
        /// </summary>
        [Output("nntp")]
        public Output<Outputs.ProfileprotocoloptionsNntp> Nntp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("oversizeLog")]
        public Output<string> OversizeLog { get; private set; } = null!;

        /// <summary>
        /// Configure POP3 protocol options. The structure of `pop3` block is documented below.
        /// </summary>
        [Output("pop3")]
        public Output<Outputs.ProfileprotocoloptionsPop3> Pop3 { get; private set; } = null!;

        /// <summary>
        /// Name of the replacement message group to be used
        /// </summary>
        [Output("replacemsgGroup")]
        public Output<string> ReplacemsgGroup { get; private set; } = null!;

        /// <summary>
        /// Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("rpcOverHttp")]
        public Output<string> RpcOverHttp { get; private set; } = null!;

        /// <summary>
        /// Configure SMTP protocol options. The structure of `smtp` block is documented below.
        /// </summary>
        [Output("smtp")]
        public Output<Outputs.ProfileprotocoloptionsSmtp> Smtp { get; private set; } = null!;

        /// <summary>
        /// Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
        /// </summary>
        [Output("ssh")]
        public Output<Outputs.ProfileprotocoloptionsSsh> Ssh { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("switchingProtocolsLog")]
        public Output<string> SwitchingProtocolsLog { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Profileprotocoloptions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Profileprotocoloptions(string name, ProfileprotocoloptionsArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/profileprotocoloptions:Profileprotocoloptions", name, args ?? new ProfileprotocoloptionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Profileprotocoloptions(string name, Input<string> id, ProfileprotocoloptionsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/profileprotocoloptions:Profileprotocoloptions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Profileprotocoloptions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Profileprotocoloptions Get(string name, Input<string> id, ProfileprotocoloptionsState? state = null, CustomResourceOptions? options = null)
        {
            return new Profileprotocoloptions(name, id, state, options);
        }
    }

    public sealed class ProfileprotocoloptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configure CIFS protocol options. The structure of `cifs` block is documented below.
        /// </summary>
        [Input("cifs")]
        public Input<Inputs.ProfileprotocoloptionsCifsArgs>? Cifs { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Configure DNS protocol options. The structure of `dns` block is documented below.
        /// </summary>
        [Input("dns")]
        public Input<Inputs.ProfileprotocoloptionsDnsArgs>? Dns { get; set; }

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Input("featureSet")]
        public Input<string>? FeatureSet { get; set; }

        /// <summary>
        /// Configure FTP protocol options. The structure of `ftp` block is documented below.
        /// </summary>
        [Input("ftp")]
        public Input<Inputs.ProfileprotocoloptionsFtpArgs>? Ftp { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Configure HTTP protocol options. The structure of `http` block is documented below.
        /// </summary>
        [Input("http")]
        public Input<Inputs.ProfileprotocoloptionsHttpArgs>? Http { get; set; }

        /// <summary>
        /// Configure IMAP protocol options. The structure of `imap` block is documented below.
        /// </summary>
        [Input("imap")]
        public Input<Inputs.ProfileprotocoloptionsImapArgs>? Imap { get; set; }

        /// <summary>
        /// Configure Mail signature. The structure of `mail_signature` block is documented below.
        /// </summary>
        [Input("mailSignature")]
        public Input<Inputs.ProfileprotocoloptionsMailSignatureArgs>? MailSignature { get; set; }

        /// <summary>
        /// Configure MAPI protocol options. The structure of `mapi` block is documented below.
        /// </summary>
        [Input("mapi")]
        public Input<Inputs.ProfileprotocoloptionsMapiArgs>? Mapi { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configure NNTP protocol options. The structure of `nntp` block is documented below.
        /// </summary>
        [Input("nntp")]
        public Input<Inputs.ProfileprotocoloptionsNntpArgs>? Nntp { get; set; }

        /// <summary>
        /// Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("oversizeLog")]
        public Input<string>? OversizeLog { get; set; }

        /// <summary>
        /// Configure POP3 protocol options. The structure of `pop3` block is documented below.
        /// </summary>
        [Input("pop3")]
        public Input<Inputs.ProfileprotocoloptionsPop3Args>? Pop3 { get; set; }

        /// <summary>
        /// Name of the replacement message group to be used
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        /// <summary>
        /// Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("rpcOverHttp")]
        public Input<string>? RpcOverHttp { get; set; }

        /// <summary>
        /// Configure SMTP protocol options. The structure of `smtp` block is documented below.
        /// </summary>
        [Input("smtp")]
        public Input<Inputs.ProfileprotocoloptionsSmtpArgs>? Smtp { get; set; }

        /// <summary>
        /// Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.ProfileprotocoloptionsSshArgs>? Ssh { get; set; }

        /// <summary>
        /// Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("switchingProtocolsLog")]
        public Input<string>? SwitchingProtocolsLog { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ProfileprotocoloptionsArgs()
        {
        }
        public static new ProfileprotocoloptionsArgs Empty => new ProfileprotocoloptionsArgs();
    }

    public sealed class ProfileprotocoloptionsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configure CIFS protocol options. The structure of `cifs` block is documented below.
        /// </summary>
        [Input("cifs")]
        public Input<Inputs.ProfileprotocoloptionsCifsGetArgs>? Cifs { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Configure DNS protocol options. The structure of `dns` block is documented below.
        /// </summary>
        [Input("dns")]
        public Input<Inputs.ProfileprotocoloptionsDnsGetArgs>? Dns { get; set; }

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Input("featureSet")]
        public Input<string>? FeatureSet { get; set; }

        /// <summary>
        /// Configure FTP protocol options. The structure of `ftp` block is documented below.
        /// </summary>
        [Input("ftp")]
        public Input<Inputs.ProfileprotocoloptionsFtpGetArgs>? Ftp { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Configure HTTP protocol options. The structure of `http` block is documented below.
        /// </summary>
        [Input("http")]
        public Input<Inputs.ProfileprotocoloptionsHttpGetArgs>? Http { get; set; }

        /// <summary>
        /// Configure IMAP protocol options. The structure of `imap` block is documented below.
        /// </summary>
        [Input("imap")]
        public Input<Inputs.ProfileprotocoloptionsImapGetArgs>? Imap { get; set; }

        /// <summary>
        /// Configure Mail signature. The structure of `mail_signature` block is documented below.
        /// </summary>
        [Input("mailSignature")]
        public Input<Inputs.ProfileprotocoloptionsMailSignatureGetArgs>? MailSignature { get; set; }

        /// <summary>
        /// Configure MAPI protocol options. The structure of `mapi` block is documented below.
        /// </summary>
        [Input("mapi")]
        public Input<Inputs.ProfileprotocoloptionsMapiGetArgs>? Mapi { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configure NNTP protocol options. The structure of `nntp` block is documented below.
        /// </summary>
        [Input("nntp")]
        public Input<Inputs.ProfileprotocoloptionsNntpGetArgs>? Nntp { get; set; }

        /// <summary>
        /// Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("oversizeLog")]
        public Input<string>? OversizeLog { get; set; }

        /// <summary>
        /// Configure POP3 protocol options. The structure of `pop3` block is documented below.
        /// </summary>
        [Input("pop3")]
        public Input<Inputs.ProfileprotocoloptionsPop3GetArgs>? Pop3 { get; set; }

        /// <summary>
        /// Name of the replacement message group to be used
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        /// <summary>
        /// Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("rpcOverHttp")]
        public Input<string>? RpcOverHttp { get; set; }

        /// <summary>
        /// Configure SMTP protocol options. The structure of `smtp` block is documented below.
        /// </summary>
        [Input("smtp")]
        public Input<Inputs.ProfileprotocoloptionsSmtpGetArgs>? Smtp { get; set; }

        /// <summary>
        /// Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.ProfileprotocoloptionsSshGetArgs>? Ssh { get; set; }

        /// <summary>
        /// Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("switchingProtocolsLog")]
        public Input<string>? SwitchingProtocolsLog { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ProfileprotocoloptionsState()
        {
        }
        public static new ProfileprotocoloptionsState Empty => new ProfileprotocoloptionsState();
    }
}
