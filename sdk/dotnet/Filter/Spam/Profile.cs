// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Spam
{
    /// <summary>
    /// Configure AntiSpam profiles. Applies to FortiOS Version `&lt;= 6.2.0`.
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
    ///     var trname = new Fortios.Filter.Spam.Profile("trname", new()
    ///     {
    ///         Comment = "terraform test",
    ///         External = "disable",
    ///         FlowBased = "disable",
    ///         Gmail = new Fortios.Filter.Spam.Inputs.ProfileGmailArgs
    ///         {
    ///             Log = "disable",
    ///         },
    ///         Imap = new Fortios.Filter.Spam.Inputs.ProfileImapArgs
    ///         {
    ///             Action = "tag",
    ///             Log = "disable",
    ///             TagMsg = "Spam",
    ///             TagType = "subject spaminfo",
    ///         },
    ///         Mapi = new Fortios.Filter.Spam.Inputs.ProfileMapiArgs
    ///         {
    ///             Action = "discard",
    ///             Log = "disable",
    ///         },
    ///         MsnHotmail = new Fortios.Filter.Spam.Inputs.ProfileMsnHotmailArgs
    ///         {
    ///             Log = "disable",
    ///         },
    ///         Pop3 = new Fortios.Filter.Spam.Inputs.ProfilePop3Args
    ///         {
    ///             Action = "tag",
    ///             Log = "disable",
    ///             TagMsg = "Spam",
    ///             TagType = "subject spaminfo",
    ///         },
    ///         Smtp = new Fortios.Filter.Spam.Inputs.ProfileSmtpArgs
    ///         {
    ///             Action = "discard",
    ///             Hdrip = "disable",
    ///             LocalOverride = "disable",
    ///             Log = "disable",
    ///             TagMsg = "Spam",
    ///             TagType = "subject spaminfo",
    ///         },
    ///         SpamBwlTable = 0,
    ///         SpamBwordTable = 0,
    ///         SpamBwordThreshold = 10,
    ///         SpamFiltering = "disable",
    ///         SpamIptrustTable = 0,
    ///         SpamLog = "enable",
    ///         SpamLogFortiguardResponse = "disable",
    ///         SpamMheaderTable = 0,
    ///         SpamRblTable = 0,
    ///         YahooMail = new Fortios.Filter.Spam.Inputs.ProfileYahooMailArgs
    ///         {
    ///             Log = "disable",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Spamfilter Profile can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:filter/spam/profile:Profile labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:filter/spam/profile:Profile labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:filter/spam/profile:Profile")]
    public partial class Profile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Enable/disable external Email inspection. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("external")]
        public Output<string> External { get; private set; } = null!;

        /// <summary>
        /// Enable/disable flow-based spam filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("flowBased")]
        public Output<string> FlowBased { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Gmail. The structure of `gmail` block is documented below.
        /// </summary>
        [Output("gmail")]
        public Output<Outputs.ProfileGmail> Gmail { get; private set; } = null!;

        /// <summary>
        /// IMAP. The structure of `imap` block is documented below.
        /// </summary>
        [Output("imap")]
        public Output<Outputs.ProfileImap> Imap { get; private set; } = null!;

        /// <summary>
        /// MAPI. The structure of `mapi` block is documented below.
        /// </summary>
        [Output("mapi")]
        public Output<Outputs.ProfileMapi> Mapi { get; private set; } = null!;

        /// <summary>
        /// MSN Hotmail. The structure of `msn_hotmail` block is documented below.
        /// </summary>
        [Output("msnHotmail")]
        public Output<Outputs.ProfileMsnHotmail> MsnHotmail { get; private set; } = null!;

        /// <summary>
        /// Profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Options. Valid values: `bannedword`, `spambwl`, `spamfsip`, `spamfssubmit`, `spamfschksum`, `spamfsurl`, `spamhelodns`, `spamraddrdns`, `spamrbl`, `spamhdrcheck`, `spamfsphish`.
        /// </summary>
        [Output("options")]
        public Output<string> Options { get; private set; } = null!;

        /// <summary>
        /// POP3. The structure of `pop3` block is documented below.
        /// </summary>
        [Output("pop3")]
        public Output<Outputs.ProfilePop3> Pop3 { get; private set; } = null!;

        /// <summary>
        /// Replacement message group.
        /// </summary>
        [Output("replacemsgGroup")]
        public Output<string> ReplacemsgGroup { get; private set; } = null!;

        /// <summary>
        /// SMTP. The structure of `smtp` block is documented below.
        /// </summary>
        [Output("smtp")]
        public Output<Outputs.ProfileSmtp> Smtp { get; private set; } = null!;

        /// <summary>
        /// Anti-spam black/white list table ID.
        /// </summary>
        [Output("spamBwlTable")]
        public Output<int> SpamBwlTable { get; private set; } = null!;

        /// <summary>
        /// Anti-spam banned word table ID.
        /// </summary>
        [Output("spamBwordTable")]
        public Output<int> SpamBwordTable { get; private set; } = null!;

        /// <summary>
        /// Spam banned word threshold.
        /// </summary>
        [Output("spamBwordThreshold")]
        public Output<int> SpamBwordThreshold { get; private set; } = null!;

        /// <summary>
        /// Enable/disable spam filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("spamFiltering")]
        public Output<string> SpamFiltering { get; private set; } = null!;

        /// <summary>
        /// Anti-spam IP trust table ID.
        /// </summary>
        [Output("spamIptrustTable")]
        public Output<int> SpamIptrustTable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("spamLog")]
        public Output<string> SpamLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("spamLogFortiguardResponse")]
        public Output<string> SpamLogFortiguardResponse { get; private set; } = null!;

        /// <summary>
        /// Anti-spam MIME header table ID.
        /// </summary>
        [Output("spamMheaderTable")]
        public Output<int> SpamMheaderTable { get; private set; } = null!;

        /// <summary>
        /// Anti-spam DNSBL table ID.
        /// </summary>
        [Output("spamRblTable")]
        public Output<int> SpamRblTable { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Yahoo! Mail. The structure of `yahoo_mail` block is documented below.
        /// </summary>
        [Output("yahooMail")]
        public Output<Outputs.ProfileYahooMail> YahooMail { get; private set; } = null!;


        /// <summary>
        /// Create a Profile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Profile(string name, ProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:filter/spam/profile:Profile", name, args ?? new ProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Profile(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:filter/spam/profile:Profile", name, state, MakeResourceOptions(options, id))
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
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Enable/disable external Email inspection. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("external")]
        public Input<string>? External { get; set; }

        /// <summary>
        /// Enable/disable flow-based spam filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("flowBased")]
        public Input<string>? FlowBased { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Gmail. The structure of `gmail` block is documented below.
        /// </summary>
        [Input("gmail")]
        public Input<Inputs.ProfileGmailArgs>? Gmail { get; set; }

        /// <summary>
        /// IMAP. The structure of `imap` block is documented below.
        /// </summary>
        [Input("imap")]
        public Input<Inputs.ProfileImapArgs>? Imap { get; set; }

        /// <summary>
        /// MAPI. The structure of `mapi` block is documented below.
        /// </summary>
        [Input("mapi")]
        public Input<Inputs.ProfileMapiArgs>? Mapi { get; set; }

        /// <summary>
        /// MSN Hotmail. The structure of `msn_hotmail` block is documented below.
        /// </summary>
        [Input("msnHotmail")]
        public Input<Inputs.ProfileMsnHotmailArgs>? MsnHotmail { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Options. Valid values: `bannedword`, `spambwl`, `spamfsip`, `spamfssubmit`, `spamfschksum`, `spamfsurl`, `spamhelodns`, `spamraddrdns`, `spamrbl`, `spamhdrcheck`, `spamfsphish`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// POP3. The structure of `pop3` block is documented below.
        /// </summary>
        [Input("pop3")]
        public Input<Inputs.ProfilePop3Args>? Pop3 { get; set; }

        /// <summary>
        /// Replacement message group.
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        /// <summary>
        /// SMTP. The structure of `smtp` block is documented below.
        /// </summary>
        [Input("smtp")]
        public Input<Inputs.ProfileSmtpArgs>? Smtp { get; set; }

        /// <summary>
        /// Anti-spam black/white list table ID.
        /// </summary>
        [Input("spamBwlTable")]
        public Input<int>? SpamBwlTable { get; set; }

        /// <summary>
        /// Anti-spam banned word table ID.
        /// </summary>
        [Input("spamBwordTable")]
        public Input<int>? SpamBwordTable { get; set; }

        /// <summary>
        /// Spam banned word threshold.
        /// </summary>
        [Input("spamBwordThreshold")]
        public Input<int>? SpamBwordThreshold { get; set; }

        /// <summary>
        /// Enable/disable spam filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("spamFiltering")]
        public Input<string>? SpamFiltering { get; set; }

        /// <summary>
        /// Anti-spam IP trust table ID.
        /// </summary>
        [Input("spamIptrustTable")]
        public Input<int>? SpamIptrustTable { get; set; }

        /// <summary>
        /// Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("spamLog")]
        public Input<string>? SpamLog { get; set; }

        /// <summary>
        /// Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("spamLogFortiguardResponse")]
        public Input<string>? SpamLogFortiguardResponse { get; set; }

        /// <summary>
        /// Anti-spam MIME header table ID.
        /// </summary>
        [Input("spamMheaderTable")]
        public Input<int>? SpamMheaderTable { get; set; }

        /// <summary>
        /// Anti-spam DNSBL table ID.
        /// </summary>
        [Input("spamRblTable")]
        public Input<int>? SpamRblTable { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Yahoo! Mail. The structure of `yahoo_mail` block is documented below.
        /// </summary>
        [Input("yahooMail")]
        public Input<Inputs.ProfileYahooMailArgs>? YahooMail { get; set; }

        public ProfileArgs()
        {
        }
        public static new ProfileArgs Empty => new ProfileArgs();
    }

    public sealed class ProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Enable/disable external Email inspection. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("external")]
        public Input<string>? External { get; set; }

        /// <summary>
        /// Enable/disable flow-based spam filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("flowBased")]
        public Input<string>? FlowBased { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Gmail. The structure of `gmail` block is documented below.
        /// </summary>
        [Input("gmail")]
        public Input<Inputs.ProfileGmailGetArgs>? Gmail { get; set; }

        /// <summary>
        /// IMAP. The structure of `imap` block is documented below.
        /// </summary>
        [Input("imap")]
        public Input<Inputs.ProfileImapGetArgs>? Imap { get; set; }

        /// <summary>
        /// MAPI. The structure of `mapi` block is documented below.
        /// </summary>
        [Input("mapi")]
        public Input<Inputs.ProfileMapiGetArgs>? Mapi { get; set; }

        /// <summary>
        /// MSN Hotmail. The structure of `msn_hotmail` block is documented below.
        /// </summary>
        [Input("msnHotmail")]
        public Input<Inputs.ProfileMsnHotmailGetArgs>? MsnHotmail { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Options. Valid values: `bannedword`, `spambwl`, `spamfsip`, `spamfssubmit`, `spamfschksum`, `spamfsurl`, `spamhelodns`, `spamraddrdns`, `spamrbl`, `spamhdrcheck`, `spamfsphish`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// POP3. The structure of `pop3` block is documented below.
        /// </summary>
        [Input("pop3")]
        public Input<Inputs.ProfilePop3GetArgs>? Pop3 { get; set; }

        /// <summary>
        /// Replacement message group.
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        /// <summary>
        /// SMTP. The structure of `smtp` block is documented below.
        /// </summary>
        [Input("smtp")]
        public Input<Inputs.ProfileSmtpGetArgs>? Smtp { get; set; }

        /// <summary>
        /// Anti-spam black/white list table ID.
        /// </summary>
        [Input("spamBwlTable")]
        public Input<int>? SpamBwlTable { get; set; }

        /// <summary>
        /// Anti-spam banned word table ID.
        /// </summary>
        [Input("spamBwordTable")]
        public Input<int>? SpamBwordTable { get; set; }

        /// <summary>
        /// Spam banned word threshold.
        /// </summary>
        [Input("spamBwordThreshold")]
        public Input<int>? SpamBwordThreshold { get; set; }

        /// <summary>
        /// Enable/disable spam filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("spamFiltering")]
        public Input<string>? SpamFiltering { get; set; }

        /// <summary>
        /// Anti-spam IP trust table ID.
        /// </summary>
        [Input("spamIptrustTable")]
        public Input<int>? SpamIptrustTable { get; set; }

        /// <summary>
        /// Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("spamLog")]
        public Input<string>? SpamLog { get; set; }

        /// <summary>
        /// Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("spamLogFortiguardResponse")]
        public Input<string>? SpamLogFortiguardResponse { get; set; }

        /// <summary>
        /// Anti-spam MIME header table ID.
        /// </summary>
        [Input("spamMheaderTable")]
        public Input<int>? SpamMheaderTable { get; set; }

        /// <summary>
        /// Anti-spam DNSBL table ID.
        /// </summary>
        [Input("spamRblTable")]
        public Input<int>? SpamRblTable { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Yahoo! Mail. The structure of `yahoo_mail` block is documented below.
        /// </summary>
        [Input("yahooMail")]
        public Input<Inputs.ProfileYahooMailGetArgs>? YahooMail { get; set; }

        public ProfileState()
        {
        }
        public static new ProfileState Empty => new ProfileState();
    }
}
