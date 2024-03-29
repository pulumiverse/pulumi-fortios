// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller
{
    /// <summary>
    /// Configure WiFi quality of service (QoS) profiles.
    /// 
    /// ## Import
    /// 
    /// WirelessController QosProfile can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/qosprofile:Qosprofile labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/qosprofile:Qosprofile labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontroller/qosprofile:Qosprofile")]
    public partial class Qosprofile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable WMM bandwidth admission control. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("bandwidthAdmissionControl")]
        public Output<string> BandwidthAdmissionControl { get; private set; } = null!;

        /// <summary>
        /// Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
        /// </summary>
        [Output("bandwidthCapacity")]
        public Output<int> BandwidthCapacity { get; private set; } = null!;

        /// <summary>
        /// Enable/disable client rate burst. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("burst")]
        public Output<string> Burst { get; private set; } = null!;

        /// <summary>
        /// Enable/disable WMM call admission control. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("callAdmissionControl")]
        public Output<string> CallAdmissionControl { get; private set; } = null!;

        /// <summary>
        /// Maximum number of Voice over WLAN (VoWLAN) phones allowed (0 - 60, default = 10).
        /// </summary>
        [Output("callCapacity")]
        public Output<int> CallCapacity { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string> Comment { get; private set; } = null!;

        /// <summary>
        /// Maximum downlink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).
        /// </summary>
        [Output("downlink")]
        public Output<int> Downlink { get; private set; } = null!;

        /// <summary>
        /// Maximum downlink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).
        /// </summary>
        [Output("downlinkSta")]
        public Output<int> DownlinkSta { get; private set; } = null!;

        /// <summary>
        /// DSCP mapping for best effort access (default = 0 24). The structure of `dscp_wmm_be` block is documented below.
        /// </summary>
        [Output("dscpWmmBes")]
        public Output<ImmutableArray<Outputs.QosprofileDscpWmmBe>> DscpWmmBes { get; private set; } = null!;

        /// <summary>
        /// DSCP mapping for background access (default = 8 16). The structure of `dscp_wmm_bk` block is documented below.
        /// </summary>
        [Output("dscpWmmBks")]
        public Output<ImmutableArray<Outputs.QosprofileDscpWmmBk>> DscpWmmBks { get; private set; } = null!;

        /// <summary>
        /// Enable/disable Differentiated Services Code Point (DSCP) mapping. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dscpWmmMapping")]
        public Output<string> DscpWmmMapping { get; private set; } = null!;

        /// <summary>
        /// DSCP mapping for video access (default = 32 40). The structure of `dscp_wmm_vi` block is documented below.
        /// </summary>
        [Output("dscpWmmVis")]
        public Output<ImmutableArray<Outputs.QosprofileDscpWmmVi>> DscpWmmVis { get; private set; } = null!;

        /// <summary>
        /// DSCP mapping for voice access (default = 48 56). The structure of `dscp_wmm_vo` block is documented below.
        /// </summary>
        [Output("dscpWmmVos")]
        public Output<ImmutableArray<Outputs.QosprofileDscpWmmVo>> DscpWmmVos { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// WiFi QoS profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Maximum uplink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).
        /// </summary>
        [Output("uplink")]
        public Output<int> Uplink { get; private set; } = null!;

        /// <summary>
        /// Maximum uplink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).
        /// </summary>
        [Output("uplinkSta")]
        public Output<int> UplinkSta { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable WiFi multi-media (WMM) control. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("wmm")]
        public Output<string> Wmm { get; private set; } = null!;

        /// <summary>
        /// DSCP marking for best effort access (default = 0).
        /// </summary>
        [Output("wmmBeDscp")]
        public Output<int> WmmBeDscp { get; private set; } = null!;

        /// <summary>
        /// DSCP marking for background access (default = 8).
        /// </summary>
        [Output("wmmBkDscp")]
        public Output<int> WmmBkDscp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable WMM Differentiated Services Code Point (DSCP) marking. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("wmmDscpMarking")]
        public Output<string> WmmDscpMarking { get; private set; } = null!;

        /// <summary>
        /// Enable/disable WMM Unscheduled Automatic Power Save Delivery (U-APSD) power save mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("wmmUapsd")]
        public Output<string> WmmUapsd { get; private set; } = null!;

        /// <summary>
        /// DSCP marking for video access (default = 32).
        /// </summary>
        [Output("wmmViDscp")]
        public Output<int> WmmViDscp { get; private set; } = null!;

        /// <summary>
        /// DSCP marking for voice access (default = 48).
        /// </summary>
        [Output("wmmVoDscp")]
        public Output<int> WmmVoDscp { get; private set; } = null!;


        /// <summary>
        /// Create a Qosprofile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Qosprofile(string name, QosprofileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/qosprofile:Qosprofile", name, args ?? new QosprofileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Qosprofile(string name, Input<string> id, QosprofileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/qosprofile:Qosprofile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Qosprofile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Qosprofile Get(string name, Input<string> id, QosprofileState? state = null, CustomResourceOptions? options = null)
        {
            return new Qosprofile(name, id, state, options);
        }
    }

    public sealed class QosprofileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable WMM bandwidth admission control. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bandwidthAdmissionControl")]
        public Input<string>? BandwidthAdmissionControl { get; set; }

        /// <summary>
        /// Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
        /// </summary>
        [Input("bandwidthCapacity")]
        public Input<int>? BandwidthCapacity { get; set; }

        /// <summary>
        /// Enable/disable client rate burst. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("burst")]
        public Input<string>? Burst { get; set; }

        /// <summary>
        /// Enable/disable WMM call admission control. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("callAdmissionControl")]
        public Input<string>? CallAdmissionControl { get; set; }

        /// <summary>
        /// Maximum number of Voice over WLAN (VoWLAN) phones allowed (0 - 60, default = 10).
        /// </summary>
        [Input("callCapacity")]
        public Input<int>? CallCapacity { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Maximum downlink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).
        /// </summary>
        [Input("downlink")]
        public Input<int>? Downlink { get; set; }

        /// <summary>
        /// Maximum downlink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).
        /// </summary>
        [Input("downlinkSta")]
        public Input<int>? DownlinkSta { get; set; }

        [Input("dscpWmmBes")]
        private InputList<Inputs.QosprofileDscpWmmBeArgs>? _dscpWmmBes;

        /// <summary>
        /// DSCP mapping for best effort access (default = 0 24). The structure of `dscp_wmm_be` block is documented below.
        /// </summary>
        public InputList<Inputs.QosprofileDscpWmmBeArgs> DscpWmmBes
        {
            get => _dscpWmmBes ?? (_dscpWmmBes = new InputList<Inputs.QosprofileDscpWmmBeArgs>());
            set => _dscpWmmBes = value;
        }

        [Input("dscpWmmBks")]
        private InputList<Inputs.QosprofileDscpWmmBkArgs>? _dscpWmmBks;

        /// <summary>
        /// DSCP mapping for background access (default = 8 16). The structure of `dscp_wmm_bk` block is documented below.
        /// </summary>
        public InputList<Inputs.QosprofileDscpWmmBkArgs> DscpWmmBks
        {
            get => _dscpWmmBks ?? (_dscpWmmBks = new InputList<Inputs.QosprofileDscpWmmBkArgs>());
            set => _dscpWmmBks = value;
        }

        /// <summary>
        /// Enable/disable Differentiated Services Code Point (DSCP) mapping. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dscpWmmMapping")]
        public Input<string>? DscpWmmMapping { get; set; }

        [Input("dscpWmmVis")]
        private InputList<Inputs.QosprofileDscpWmmViArgs>? _dscpWmmVis;

        /// <summary>
        /// DSCP mapping for video access (default = 32 40). The structure of `dscp_wmm_vi` block is documented below.
        /// </summary>
        public InputList<Inputs.QosprofileDscpWmmViArgs> DscpWmmVis
        {
            get => _dscpWmmVis ?? (_dscpWmmVis = new InputList<Inputs.QosprofileDscpWmmViArgs>());
            set => _dscpWmmVis = value;
        }

        [Input("dscpWmmVos")]
        private InputList<Inputs.QosprofileDscpWmmVoArgs>? _dscpWmmVos;

        /// <summary>
        /// DSCP mapping for voice access (default = 48 56). The structure of `dscp_wmm_vo` block is documented below.
        /// </summary>
        public InputList<Inputs.QosprofileDscpWmmVoArgs> DscpWmmVos
        {
            get => _dscpWmmVos ?? (_dscpWmmVos = new InputList<Inputs.QosprofileDscpWmmVoArgs>());
            set => _dscpWmmVos = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// WiFi QoS profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Maximum uplink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).
        /// </summary>
        [Input("uplink")]
        public Input<int>? Uplink { get; set; }

        /// <summary>
        /// Maximum uplink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).
        /// </summary>
        [Input("uplinkSta")]
        public Input<int>? UplinkSta { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable WiFi multi-media (WMM) control. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wmm")]
        public Input<string>? Wmm { get; set; }

        /// <summary>
        /// DSCP marking for best effort access (default = 0).
        /// </summary>
        [Input("wmmBeDscp")]
        public Input<int>? WmmBeDscp { get; set; }

        /// <summary>
        /// DSCP marking for background access (default = 8).
        /// </summary>
        [Input("wmmBkDscp")]
        public Input<int>? WmmBkDscp { get; set; }

        /// <summary>
        /// Enable/disable WMM Differentiated Services Code Point (DSCP) marking. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wmmDscpMarking")]
        public Input<string>? WmmDscpMarking { get; set; }

        /// <summary>
        /// Enable/disable WMM Unscheduled Automatic Power Save Delivery (U-APSD) power save mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wmmUapsd")]
        public Input<string>? WmmUapsd { get; set; }

        /// <summary>
        /// DSCP marking for video access (default = 32).
        /// </summary>
        [Input("wmmViDscp")]
        public Input<int>? WmmViDscp { get; set; }

        /// <summary>
        /// DSCP marking for voice access (default = 48).
        /// </summary>
        [Input("wmmVoDscp")]
        public Input<int>? WmmVoDscp { get; set; }

        public QosprofileArgs()
        {
        }
        public static new QosprofileArgs Empty => new QosprofileArgs();
    }

    public sealed class QosprofileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable WMM bandwidth admission control. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bandwidthAdmissionControl")]
        public Input<string>? BandwidthAdmissionControl { get; set; }

        /// <summary>
        /// Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
        /// </summary>
        [Input("bandwidthCapacity")]
        public Input<int>? BandwidthCapacity { get; set; }

        /// <summary>
        /// Enable/disable client rate burst. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("burst")]
        public Input<string>? Burst { get; set; }

        /// <summary>
        /// Enable/disable WMM call admission control. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("callAdmissionControl")]
        public Input<string>? CallAdmissionControl { get; set; }

        /// <summary>
        /// Maximum number of Voice over WLAN (VoWLAN) phones allowed (0 - 60, default = 10).
        /// </summary>
        [Input("callCapacity")]
        public Input<int>? CallCapacity { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Maximum downlink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).
        /// </summary>
        [Input("downlink")]
        public Input<int>? Downlink { get; set; }

        /// <summary>
        /// Maximum downlink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).
        /// </summary>
        [Input("downlinkSta")]
        public Input<int>? DownlinkSta { get; set; }

        [Input("dscpWmmBes")]
        private InputList<Inputs.QosprofileDscpWmmBeGetArgs>? _dscpWmmBes;

        /// <summary>
        /// DSCP mapping for best effort access (default = 0 24). The structure of `dscp_wmm_be` block is documented below.
        /// </summary>
        public InputList<Inputs.QosprofileDscpWmmBeGetArgs> DscpWmmBes
        {
            get => _dscpWmmBes ?? (_dscpWmmBes = new InputList<Inputs.QosprofileDscpWmmBeGetArgs>());
            set => _dscpWmmBes = value;
        }

        [Input("dscpWmmBks")]
        private InputList<Inputs.QosprofileDscpWmmBkGetArgs>? _dscpWmmBks;

        /// <summary>
        /// DSCP mapping for background access (default = 8 16). The structure of `dscp_wmm_bk` block is documented below.
        /// </summary>
        public InputList<Inputs.QosprofileDscpWmmBkGetArgs> DscpWmmBks
        {
            get => _dscpWmmBks ?? (_dscpWmmBks = new InputList<Inputs.QosprofileDscpWmmBkGetArgs>());
            set => _dscpWmmBks = value;
        }

        /// <summary>
        /// Enable/disable Differentiated Services Code Point (DSCP) mapping. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dscpWmmMapping")]
        public Input<string>? DscpWmmMapping { get; set; }

        [Input("dscpWmmVis")]
        private InputList<Inputs.QosprofileDscpWmmViGetArgs>? _dscpWmmVis;

        /// <summary>
        /// DSCP mapping for video access (default = 32 40). The structure of `dscp_wmm_vi` block is documented below.
        /// </summary>
        public InputList<Inputs.QosprofileDscpWmmViGetArgs> DscpWmmVis
        {
            get => _dscpWmmVis ?? (_dscpWmmVis = new InputList<Inputs.QosprofileDscpWmmViGetArgs>());
            set => _dscpWmmVis = value;
        }

        [Input("dscpWmmVos")]
        private InputList<Inputs.QosprofileDscpWmmVoGetArgs>? _dscpWmmVos;

        /// <summary>
        /// DSCP mapping for voice access (default = 48 56). The structure of `dscp_wmm_vo` block is documented below.
        /// </summary>
        public InputList<Inputs.QosprofileDscpWmmVoGetArgs> DscpWmmVos
        {
            get => _dscpWmmVos ?? (_dscpWmmVos = new InputList<Inputs.QosprofileDscpWmmVoGetArgs>());
            set => _dscpWmmVos = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// WiFi QoS profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Maximum uplink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).
        /// </summary>
        [Input("uplink")]
        public Input<int>? Uplink { get; set; }

        /// <summary>
        /// Maximum uplink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).
        /// </summary>
        [Input("uplinkSta")]
        public Input<int>? UplinkSta { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable WiFi multi-media (WMM) control. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wmm")]
        public Input<string>? Wmm { get; set; }

        /// <summary>
        /// DSCP marking for best effort access (default = 0).
        /// </summary>
        [Input("wmmBeDscp")]
        public Input<int>? WmmBeDscp { get; set; }

        /// <summary>
        /// DSCP marking for background access (default = 8).
        /// </summary>
        [Input("wmmBkDscp")]
        public Input<int>? WmmBkDscp { get; set; }

        /// <summary>
        /// Enable/disable WMM Differentiated Services Code Point (DSCP) marking. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wmmDscpMarking")]
        public Input<string>? WmmDscpMarking { get; set; }

        /// <summary>
        /// Enable/disable WMM Unscheduled Automatic Power Save Delivery (U-APSD) power save mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wmmUapsd")]
        public Input<string>? WmmUapsd { get; set; }

        /// <summary>
        /// DSCP marking for video access (default = 32).
        /// </summary>
        [Input("wmmViDscp")]
        public Input<int>? WmmViDscp { get; set; }

        /// <summary>
        /// DSCP marking for voice access (default = 48).
        /// </summary>
        [Input("wmmVoDscp")]
        public Input<int>? WmmVoDscp { get; set; }

        public QosprofileState()
        {
        }
        public static new QosprofileState Empty => new QosprofileState();
    }
}
