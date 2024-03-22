// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    /// <summary>
    /// Configure replacement message groups.
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
    ///     var trname = new Fortios.System.Replacemsggroup("trname", new()
    ///     {
    ///         Comment = "sgh",
    ///         GroupType = "utm",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// System ReplacemsgGroup can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/replacemsggroup:Replacemsggroup labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/replacemsggroup:Replacemsggroup labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/replacemsggroup:Replacemsggroup")]
    public partial class Replacemsggroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Replacement message table entries. The structure of `admin` block is documented below.
        /// </summary>
        [Output("admins")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupAdmin>> Admins { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `alertmail` block is documented below.
        /// </summary>
        [Output("alertmails")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupAlertmail>> Alertmails { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `auth` block is documented below.
        /// </summary>
        [Output("auths")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupAuth>> Auths { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `automation` block is documented below.
        /// </summary>
        [Output("automations")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupAutomation>> Automations { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `custom_message` block is documented below.
        /// </summary>
        [Output("customMessages")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupCustomMessage>> CustomMessages { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `device_detection_portal` block is documented below.
        /// </summary>
        [Output("deviceDetectionPortals")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupDeviceDetectionPortal>> DeviceDetectionPortals { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `ec` block is documented below.
        /// </summary>
        [Output("ecs")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupEc>> Ecs { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `fortiguard_wf` block is documented below.
        /// </summary>
        [Output("fortiguardWfs")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupFortiguardWf>> FortiguardWfs { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `ftp` block is documented below.
        /// </summary>
        [Output("ftps")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupFtp>> Ftps { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Group type.
        /// </summary>
        [Output("groupType")]
        public Output<string> GroupType { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `http` block is documented below.
        /// </summary>
        [Output("https")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupHttp>> Https { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `icap` block is documented below.
        /// </summary>
        [Output("icaps")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupIcap>> Icaps { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `mail` block is documented below.
        /// </summary>
        [Output("mails")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupMail>> Mails { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `nac_quar` block is documented below.
        /// </summary>
        [Output("nacQuars")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupNacQuar>> NacQuars { get; private set; } = null!;

        /// <summary>
        /// Group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `nntp` block is documented below.
        /// </summary>
        [Output("nntps")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupNntp>> Nntps { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `spam` block is documented below.
        /// </summary>
        [Output("spams")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupSpam>> Spams { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `sslvpn` block is documented below.
        /// </summary>
        [Output("sslvpns")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupSslvpn>> Sslvpns { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `traffic_quota` block is documented below.
        /// </summary>
        [Output("trafficQuotas")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupTrafficQuota>> TrafficQuotas { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `utm` block is documented below.
        /// </summary>
        [Output("utms")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupUtm>> Utms { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Replacement message table entries. The structure of `webproxy` block is documented below.
        /// </summary>
        [Output("webproxies")]
        public Output<ImmutableArray<Outputs.ReplacemsggroupWebproxy>> Webproxies { get; private set; } = null!;


        /// <summary>
        /// Create a Replacemsggroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Replacemsggroup(string name, ReplacemsggroupArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/replacemsggroup:Replacemsggroup", name, args ?? new ReplacemsggroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Replacemsggroup(string name, Input<string> id, ReplacemsggroupState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/replacemsggroup:Replacemsggroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Replacemsggroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Replacemsggroup Get(string name, Input<string> id, ReplacemsggroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Replacemsggroup(name, id, state, options);
        }
    }

    public sealed class ReplacemsggroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("admins")]
        private InputList<Inputs.ReplacemsggroupAdminArgs>? _admins;

        /// <summary>
        /// Replacement message table entries. The structure of `admin` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupAdminArgs> Admins
        {
            get => _admins ?? (_admins = new InputList<Inputs.ReplacemsggroupAdminArgs>());
            set => _admins = value;
        }

        [Input("alertmails")]
        private InputList<Inputs.ReplacemsggroupAlertmailArgs>? _alertmails;

        /// <summary>
        /// Replacement message table entries. The structure of `alertmail` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupAlertmailArgs> Alertmails
        {
            get => _alertmails ?? (_alertmails = new InputList<Inputs.ReplacemsggroupAlertmailArgs>());
            set => _alertmails = value;
        }

        [Input("auths")]
        private InputList<Inputs.ReplacemsggroupAuthArgs>? _auths;

        /// <summary>
        /// Replacement message table entries. The structure of `auth` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupAuthArgs> Auths
        {
            get => _auths ?? (_auths = new InputList<Inputs.ReplacemsggroupAuthArgs>());
            set => _auths = value;
        }

        [Input("automations")]
        private InputList<Inputs.ReplacemsggroupAutomationArgs>? _automations;

        /// <summary>
        /// Replacement message table entries. The structure of `automation` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupAutomationArgs> Automations
        {
            get => _automations ?? (_automations = new InputList<Inputs.ReplacemsggroupAutomationArgs>());
            set => _automations = value;
        }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("customMessages")]
        private InputList<Inputs.ReplacemsggroupCustomMessageArgs>? _customMessages;

        /// <summary>
        /// Replacement message table entries. The structure of `custom_message` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupCustomMessageArgs> CustomMessages
        {
            get => _customMessages ?? (_customMessages = new InputList<Inputs.ReplacemsggroupCustomMessageArgs>());
            set => _customMessages = value;
        }

        [Input("deviceDetectionPortals")]
        private InputList<Inputs.ReplacemsggroupDeviceDetectionPortalArgs>? _deviceDetectionPortals;

        /// <summary>
        /// Replacement message table entries. The structure of `device_detection_portal` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupDeviceDetectionPortalArgs> DeviceDetectionPortals
        {
            get => _deviceDetectionPortals ?? (_deviceDetectionPortals = new InputList<Inputs.ReplacemsggroupDeviceDetectionPortalArgs>());
            set => _deviceDetectionPortals = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("ecs")]
        private InputList<Inputs.ReplacemsggroupEcArgs>? _ecs;

        /// <summary>
        /// Replacement message table entries. The structure of `ec` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupEcArgs> Ecs
        {
            get => _ecs ?? (_ecs = new InputList<Inputs.ReplacemsggroupEcArgs>());
            set => _ecs = value;
        }

        [Input("fortiguardWfs")]
        private InputList<Inputs.ReplacemsggroupFortiguardWfArgs>? _fortiguardWfs;

        /// <summary>
        /// Replacement message table entries. The structure of `fortiguard_wf` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupFortiguardWfArgs> FortiguardWfs
        {
            get => _fortiguardWfs ?? (_fortiguardWfs = new InputList<Inputs.ReplacemsggroupFortiguardWfArgs>());
            set => _fortiguardWfs = value;
        }

        [Input("ftps")]
        private InputList<Inputs.ReplacemsggroupFtpArgs>? _ftps;

        /// <summary>
        /// Replacement message table entries. The structure of `ftp` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupFtpArgs> Ftps
        {
            get => _ftps ?? (_ftps = new InputList<Inputs.ReplacemsggroupFtpArgs>());
            set => _ftps = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Group type.
        /// </summary>
        [Input("groupType", required: true)]
        public Input<string> GroupType { get; set; } = null!;

        [Input("https")]
        private InputList<Inputs.ReplacemsggroupHttpArgs>? _https;

        /// <summary>
        /// Replacement message table entries. The structure of `http` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupHttpArgs> Https
        {
            get => _https ?? (_https = new InputList<Inputs.ReplacemsggroupHttpArgs>());
            set => _https = value;
        }

        [Input("icaps")]
        private InputList<Inputs.ReplacemsggroupIcapArgs>? _icaps;

        /// <summary>
        /// Replacement message table entries. The structure of `icap` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupIcapArgs> Icaps
        {
            get => _icaps ?? (_icaps = new InputList<Inputs.ReplacemsggroupIcapArgs>());
            set => _icaps = value;
        }

        [Input("mails")]
        private InputList<Inputs.ReplacemsggroupMailArgs>? _mails;

        /// <summary>
        /// Replacement message table entries. The structure of `mail` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupMailArgs> Mails
        {
            get => _mails ?? (_mails = new InputList<Inputs.ReplacemsggroupMailArgs>());
            set => _mails = value;
        }

        [Input("nacQuars")]
        private InputList<Inputs.ReplacemsggroupNacQuarArgs>? _nacQuars;

        /// <summary>
        /// Replacement message table entries. The structure of `nac_quar` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupNacQuarArgs> NacQuars
        {
            get => _nacQuars ?? (_nacQuars = new InputList<Inputs.ReplacemsggroupNacQuarArgs>());
            set => _nacQuars = value;
        }

        /// <summary>
        /// Group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nntps")]
        private InputList<Inputs.ReplacemsggroupNntpArgs>? _nntps;

        /// <summary>
        /// Replacement message table entries. The structure of `nntp` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupNntpArgs> Nntps
        {
            get => _nntps ?? (_nntps = new InputList<Inputs.ReplacemsggroupNntpArgs>());
            set => _nntps = value;
        }

        [Input("spams")]
        private InputList<Inputs.ReplacemsggroupSpamArgs>? _spams;

        /// <summary>
        /// Replacement message table entries. The structure of `spam` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupSpamArgs> Spams
        {
            get => _spams ?? (_spams = new InputList<Inputs.ReplacemsggroupSpamArgs>());
            set => _spams = value;
        }

        [Input("sslvpns")]
        private InputList<Inputs.ReplacemsggroupSslvpnArgs>? _sslvpns;

        /// <summary>
        /// Replacement message table entries. The structure of `sslvpn` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupSslvpnArgs> Sslvpns
        {
            get => _sslvpns ?? (_sslvpns = new InputList<Inputs.ReplacemsggroupSslvpnArgs>());
            set => _sslvpns = value;
        }

        [Input("trafficQuotas")]
        private InputList<Inputs.ReplacemsggroupTrafficQuotaArgs>? _trafficQuotas;

        /// <summary>
        /// Replacement message table entries. The structure of `traffic_quota` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupTrafficQuotaArgs> TrafficQuotas
        {
            get => _trafficQuotas ?? (_trafficQuotas = new InputList<Inputs.ReplacemsggroupTrafficQuotaArgs>());
            set => _trafficQuotas = value;
        }

        [Input("utms")]
        private InputList<Inputs.ReplacemsggroupUtmArgs>? _utms;

        /// <summary>
        /// Replacement message table entries. The structure of `utm` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupUtmArgs> Utms
        {
            get => _utms ?? (_utms = new InputList<Inputs.ReplacemsggroupUtmArgs>());
            set => _utms = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("webproxies")]
        private InputList<Inputs.ReplacemsggroupWebproxyArgs>? _webproxies;

        /// <summary>
        /// Replacement message table entries. The structure of `webproxy` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupWebproxyArgs> Webproxies
        {
            get => _webproxies ?? (_webproxies = new InputList<Inputs.ReplacemsggroupWebproxyArgs>());
            set => _webproxies = value;
        }

        public ReplacemsggroupArgs()
        {
        }
        public static new ReplacemsggroupArgs Empty => new ReplacemsggroupArgs();
    }

    public sealed class ReplacemsggroupState : global::Pulumi.ResourceArgs
    {
        [Input("admins")]
        private InputList<Inputs.ReplacemsggroupAdminGetArgs>? _admins;

        /// <summary>
        /// Replacement message table entries. The structure of `admin` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupAdminGetArgs> Admins
        {
            get => _admins ?? (_admins = new InputList<Inputs.ReplacemsggroupAdminGetArgs>());
            set => _admins = value;
        }

        [Input("alertmails")]
        private InputList<Inputs.ReplacemsggroupAlertmailGetArgs>? _alertmails;

        /// <summary>
        /// Replacement message table entries. The structure of `alertmail` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupAlertmailGetArgs> Alertmails
        {
            get => _alertmails ?? (_alertmails = new InputList<Inputs.ReplacemsggroupAlertmailGetArgs>());
            set => _alertmails = value;
        }

        [Input("auths")]
        private InputList<Inputs.ReplacemsggroupAuthGetArgs>? _auths;

        /// <summary>
        /// Replacement message table entries. The structure of `auth` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupAuthGetArgs> Auths
        {
            get => _auths ?? (_auths = new InputList<Inputs.ReplacemsggroupAuthGetArgs>());
            set => _auths = value;
        }

        [Input("automations")]
        private InputList<Inputs.ReplacemsggroupAutomationGetArgs>? _automations;

        /// <summary>
        /// Replacement message table entries. The structure of `automation` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupAutomationGetArgs> Automations
        {
            get => _automations ?? (_automations = new InputList<Inputs.ReplacemsggroupAutomationGetArgs>());
            set => _automations = value;
        }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("customMessages")]
        private InputList<Inputs.ReplacemsggroupCustomMessageGetArgs>? _customMessages;

        /// <summary>
        /// Replacement message table entries. The structure of `custom_message` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupCustomMessageGetArgs> CustomMessages
        {
            get => _customMessages ?? (_customMessages = new InputList<Inputs.ReplacemsggroupCustomMessageGetArgs>());
            set => _customMessages = value;
        }

        [Input("deviceDetectionPortals")]
        private InputList<Inputs.ReplacemsggroupDeviceDetectionPortalGetArgs>? _deviceDetectionPortals;

        /// <summary>
        /// Replacement message table entries. The structure of `device_detection_portal` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupDeviceDetectionPortalGetArgs> DeviceDetectionPortals
        {
            get => _deviceDetectionPortals ?? (_deviceDetectionPortals = new InputList<Inputs.ReplacemsggroupDeviceDetectionPortalGetArgs>());
            set => _deviceDetectionPortals = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("ecs")]
        private InputList<Inputs.ReplacemsggroupEcGetArgs>? _ecs;

        /// <summary>
        /// Replacement message table entries. The structure of `ec` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupEcGetArgs> Ecs
        {
            get => _ecs ?? (_ecs = new InputList<Inputs.ReplacemsggroupEcGetArgs>());
            set => _ecs = value;
        }

        [Input("fortiguardWfs")]
        private InputList<Inputs.ReplacemsggroupFortiguardWfGetArgs>? _fortiguardWfs;

        /// <summary>
        /// Replacement message table entries. The structure of `fortiguard_wf` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupFortiguardWfGetArgs> FortiguardWfs
        {
            get => _fortiguardWfs ?? (_fortiguardWfs = new InputList<Inputs.ReplacemsggroupFortiguardWfGetArgs>());
            set => _fortiguardWfs = value;
        }

        [Input("ftps")]
        private InputList<Inputs.ReplacemsggroupFtpGetArgs>? _ftps;

        /// <summary>
        /// Replacement message table entries. The structure of `ftp` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupFtpGetArgs> Ftps
        {
            get => _ftps ?? (_ftps = new InputList<Inputs.ReplacemsggroupFtpGetArgs>());
            set => _ftps = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Group type.
        /// </summary>
        [Input("groupType")]
        public Input<string>? GroupType { get; set; }

        [Input("https")]
        private InputList<Inputs.ReplacemsggroupHttpGetArgs>? _https;

        /// <summary>
        /// Replacement message table entries. The structure of `http` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupHttpGetArgs> Https
        {
            get => _https ?? (_https = new InputList<Inputs.ReplacemsggroupHttpGetArgs>());
            set => _https = value;
        }

        [Input("icaps")]
        private InputList<Inputs.ReplacemsggroupIcapGetArgs>? _icaps;

        /// <summary>
        /// Replacement message table entries. The structure of `icap` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupIcapGetArgs> Icaps
        {
            get => _icaps ?? (_icaps = new InputList<Inputs.ReplacemsggroupIcapGetArgs>());
            set => _icaps = value;
        }

        [Input("mails")]
        private InputList<Inputs.ReplacemsggroupMailGetArgs>? _mails;

        /// <summary>
        /// Replacement message table entries. The structure of `mail` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupMailGetArgs> Mails
        {
            get => _mails ?? (_mails = new InputList<Inputs.ReplacemsggroupMailGetArgs>());
            set => _mails = value;
        }

        [Input("nacQuars")]
        private InputList<Inputs.ReplacemsggroupNacQuarGetArgs>? _nacQuars;

        /// <summary>
        /// Replacement message table entries. The structure of `nac_quar` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupNacQuarGetArgs> NacQuars
        {
            get => _nacQuars ?? (_nacQuars = new InputList<Inputs.ReplacemsggroupNacQuarGetArgs>());
            set => _nacQuars = value;
        }

        /// <summary>
        /// Group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nntps")]
        private InputList<Inputs.ReplacemsggroupNntpGetArgs>? _nntps;

        /// <summary>
        /// Replacement message table entries. The structure of `nntp` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupNntpGetArgs> Nntps
        {
            get => _nntps ?? (_nntps = new InputList<Inputs.ReplacemsggroupNntpGetArgs>());
            set => _nntps = value;
        }

        [Input("spams")]
        private InputList<Inputs.ReplacemsggroupSpamGetArgs>? _spams;

        /// <summary>
        /// Replacement message table entries. The structure of `spam` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupSpamGetArgs> Spams
        {
            get => _spams ?? (_spams = new InputList<Inputs.ReplacemsggroupSpamGetArgs>());
            set => _spams = value;
        }

        [Input("sslvpns")]
        private InputList<Inputs.ReplacemsggroupSslvpnGetArgs>? _sslvpns;

        /// <summary>
        /// Replacement message table entries. The structure of `sslvpn` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupSslvpnGetArgs> Sslvpns
        {
            get => _sslvpns ?? (_sslvpns = new InputList<Inputs.ReplacemsggroupSslvpnGetArgs>());
            set => _sslvpns = value;
        }

        [Input("trafficQuotas")]
        private InputList<Inputs.ReplacemsggroupTrafficQuotaGetArgs>? _trafficQuotas;

        /// <summary>
        /// Replacement message table entries. The structure of `traffic_quota` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupTrafficQuotaGetArgs> TrafficQuotas
        {
            get => _trafficQuotas ?? (_trafficQuotas = new InputList<Inputs.ReplacemsggroupTrafficQuotaGetArgs>());
            set => _trafficQuotas = value;
        }

        [Input("utms")]
        private InputList<Inputs.ReplacemsggroupUtmGetArgs>? _utms;

        /// <summary>
        /// Replacement message table entries. The structure of `utm` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupUtmGetArgs> Utms
        {
            get => _utms ?? (_utms = new InputList<Inputs.ReplacemsggroupUtmGetArgs>());
            set => _utms = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("webproxies")]
        private InputList<Inputs.ReplacemsggroupWebproxyGetArgs>? _webproxies;

        /// <summary>
        /// Replacement message table entries. The structure of `webproxy` block is documented below.
        /// </summary>
        public InputList<Inputs.ReplacemsggroupWebproxyGetArgs> Webproxies
        {
            get => _webproxies ?? (_webproxies = new InputList<Inputs.ReplacemsggroupWebproxyGetArgs>());
            set => _webproxies = value;
        }

        public ReplacemsggroupState()
        {
        }
        public static new ReplacemsggroupState Empty => new ReplacemsggroupState();
    }
}
