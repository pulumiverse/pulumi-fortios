// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.
 *
 * ## Import
 *
 * WirelessController Wtp can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:wirelesscontroller/wtp:Wtp labelname {{wtp_id}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:wirelesscontroller/wtp:Wtp labelname {{wtp_id}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Wtp extends pulumi.CustomResource {
    /**
     * Get an existing Wtp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WtpState, opts?: pulumi.CustomResourceOptions): Wtp {
        return new Wtp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:wirelesscontroller/wtp:Wtp';

    /**
     * Returns true if the given object is an instance of Wtp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Wtp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Wtp.__pulumiType;
    }

    /**
     * Configure how the FortiGate operating as a wireless controller discovers and manages this WTP, AP or FortiAP. Valid values: `discovered`, `disable`, `enable`.
     */
    public readonly admin!: pulumi.Output<string>;
    /**
     * Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.
     */
    public readonly allowaccess!: pulumi.Output<string>;
    /**
     * AP local configuration profile name.
     */
    public readonly apcfgProfile!: pulumi.Output<string>;
    /**
     * Bonjour profile name.
     */
    public readonly bonjourProfile!: pulumi.Output<string>;
    /**
     * WTP latitude coordinate.
     */
    public readonly coordinateLatitude!: pulumi.Output<string>;
    /**
     * WTP longitude coordinate.
     */
    public readonly coordinateLongitude!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Firmware version to provision to this FortiAP on bootup (major.minor.build, i.e. 6.2.1234).
     */
    public readonly firmwareProvision!: pulumi.Output<string>;
    /**
     * Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
     */
    public readonly firmwareProvisionLatest!: pulumi.Output<string>;
    /**
     * Enable/disable WTP image download. Valid values: `enable`, `disable`.
     */
    public readonly imageDownload!: pulumi.Output<string>;
    /**
     * Index (0 - 4294967295).
     */
    public readonly index!: pulumi.Output<number>;
    /**
     * Method by which IP fragmentation is prevented for CAPWAP tunneled control and data packets (default = tcp-mss-adjust). Valid values: `tcp-mss-adjust`, `icmp-unreachable`.
     */
    public readonly ipFragmentPreventing!: pulumi.Output<string>;
    /**
     * WTP LAN port mapping. The structure of `lan` block is documented below.
     */
    public readonly lan!: pulumi.Output<outputs.wirelesscontroller.WtpLan>;
    /**
     * Enable to allow the FortiAPs LEDs to light. Disable to keep the LEDs off. You may want to keep the LEDs off so they are not distracting in low light areas etc. Valid values: `enable`, `disable`.
     */
    public readonly ledState!: pulumi.Output<string>;
    /**
     * Field for describing the physical location of the WTP, AP or FortiAP.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Set the managed WTP, FortiAP, or AP's administrator password.
     */
    public readonly loginPasswd!: pulumi.Output<string | undefined>;
    /**
     * Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
     */
    public readonly loginPasswdChange!: pulumi.Output<string>;
    /**
     * Enable/disable mesh Ethernet bridge when WTP is configured as a mesh branch/leaf AP. Valid values: `default`, `enable`, `disable`.
     */
    public readonly meshBridgeEnable!: pulumi.Output<string>;
    /**
     * WTP, AP or FortiAP configuration name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable to override the WTP profile management access configuration. Valid values: `enable`, `disable`.
     */
    public readonly overrideAllowaccess!: pulumi.Output<string>;
    /**
     * Enable/disable overriding the WTP profile IP fragment prevention setting. Valid values: `enable`, `disable`.
     */
    public readonly overrideIpFragment!: pulumi.Output<string>;
    /**
     * Enable to override the WTP profile LAN port setting. Valid values: `enable`, `disable`.
     */
    public readonly overrideLan!: pulumi.Output<string>;
    /**
     * Enable to override the profile LED state setting for this FortiAP. You must enable this option to use the led-state command to turn off the FortiAP's LEDs. Valid values: `enable`, `disable`.
     */
    public readonly overrideLedState!: pulumi.Output<string>;
    /**
     * Enable to override the WTP profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
     */
    public readonly overrideLoginPasswdChange!: pulumi.Output<string>;
    /**
     * Enable/disable overriding the WTP profile split tunneling setting. Valid values: `enable`, `disable`.
     */
    public readonly overrideSplitTunnel!: pulumi.Output<string>;
    /**
     * Enable/disable overriding the wan-port-mode in the WTP profile. Valid values: `enable`, `disable`.
     */
    public readonly overrideWanPortMode!: pulumi.Output<string>;
    /**
     * Configuration options for radio 1. The structure of `radio1` block is documented below.
     */
    public readonly radio1!: pulumi.Output<outputs.wirelesscontroller.WtpRadio1>;
    /**
     * Configuration options for radio 2. The structure of `radio2` block is documented below.
     */
    public readonly radio2!: pulumi.Output<outputs.wirelesscontroller.WtpRadio2>;
    /**
     * Configuration options for radio 3. The structure of `radio3` block is documented below.
     */
    public readonly radio3!: pulumi.Output<outputs.wirelesscontroller.WtpRadio3>;
    /**
     * Configuration options for radio 4. The structure of `radio4` block is documented below.
     */
    public readonly radio4!: pulumi.Output<outputs.wirelesscontroller.WtpRadio4>;
    /**
     * Region name WTP is associated with.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Relative horizontal region coordinate (between 0 and 1).
     */
    public readonly regionX!: pulumi.Output<string>;
    /**
     * Relative vertical region coordinate (between 0 and 1).
     */
    public readonly regionY!: pulumi.Output<string>;
    /**
     * Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly splitTunnelingAclLocalApSubnet!: pulumi.Output<string>;
    /**
     * Split tunneling ACL path is local/tunnel. Valid values: `tunnel`, `local`.
     */
    public readonly splitTunnelingAclPath!: pulumi.Output<string>;
    /**
     * Split tunneling ACL filter list. The structure of `splitTunnelingAcl` block is documented below.
     */
    public readonly splitTunnelingAcls!: pulumi.Output<outputs.wirelesscontroller.WtpSplitTunnelingAcl[] | undefined>;
    /**
     * Downlink tunnel MTU in octets. Set the value to either 0 (by default), 576, or 1500.
     */
    public readonly tunMtuDownlink!: pulumi.Output<number>;
    /**
     * Uplink tunnel maximum transmission unit (MTU) in octets (eight-bit bytes). Set the value to either 0 (by default), 576, or 1500.
     */
    public readonly tunMtuUplink!: pulumi.Output<number>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    public readonly uuid!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable using the FortiAP WAN port as a LAN port. Valid values: `wan-lan`, `wan-only`.
     */
    public readonly wanPortMode!: pulumi.Output<string>;
    /**
     * WTP ID.
     */
    public readonly wtpId!: pulumi.Output<string>;
    /**
     * WTP, AP, or FortiAP operating mode; normal (by default) or remote. A tunnel mode SSID can be assigned to an AP in normal mode but not remote mode, while a local-bridge mode SSID can be assigned to an AP in either normal mode or remote mode. Valid values: `normal`, `remote`.
     */
    public readonly wtpMode!: pulumi.Output<string>;
    /**
     * WTP profile name to apply to this WTP, AP or FortiAP.
     */
    public readonly wtpProfile!: pulumi.Output<string>;

    /**
     * Create a Wtp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WtpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WtpArgs | WtpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WtpState | undefined;
            resourceInputs["admin"] = state ? state.admin : undefined;
            resourceInputs["allowaccess"] = state ? state.allowaccess : undefined;
            resourceInputs["apcfgProfile"] = state ? state.apcfgProfile : undefined;
            resourceInputs["bonjourProfile"] = state ? state.bonjourProfile : undefined;
            resourceInputs["coordinateLatitude"] = state ? state.coordinateLatitude : undefined;
            resourceInputs["coordinateLongitude"] = state ? state.coordinateLongitude : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["firmwareProvision"] = state ? state.firmwareProvision : undefined;
            resourceInputs["firmwareProvisionLatest"] = state ? state.firmwareProvisionLatest : undefined;
            resourceInputs["imageDownload"] = state ? state.imageDownload : undefined;
            resourceInputs["index"] = state ? state.index : undefined;
            resourceInputs["ipFragmentPreventing"] = state ? state.ipFragmentPreventing : undefined;
            resourceInputs["lan"] = state ? state.lan : undefined;
            resourceInputs["ledState"] = state ? state.ledState : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["loginPasswd"] = state ? state.loginPasswd : undefined;
            resourceInputs["loginPasswdChange"] = state ? state.loginPasswdChange : undefined;
            resourceInputs["meshBridgeEnable"] = state ? state.meshBridgeEnable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["overrideAllowaccess"] = state ? state.overrideAllowaccess : undefined;
            resourceInputs["overrideIpFragment"] = state ? state.overrideIpFragment : undefined;
            resourceInputs["overrideLan"] = state ? state.overrideLan : undefined;
            resourceInputs["overrideLedState"] = state ? state.overrideLedState : undefined;
            resourceInputs["overrideLoginPasswdChange"] = state ? state.overrideLoginPasswdChange : undefined;
            resourceInputs["overrideSplitTunnel"] = state ? state.overrideSplitTunnel : undefined;
            resourceInputs["overrideWanPortMode"] = state ? state.overrideWanPortMode : undefined;
            resourceInputs["radio1"] = state ? state.radio1 : undefined;
            resourceInputs["radio2"] = state ? state.radio2 : undefined;
            resourceInputs["radio3"] = state ? state.radio3 : undefined;
            resourceInputs["radio4"] = state ? state.radio4 : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["regionX"] = state ? state.regionX : undefined;
            resourceInputs["regionY"] = state ? state.regionY : undefined;
            resourceInputs["splitTunnelingAclLocalApSubnet"] = state ? state.splitTunnelingAclLocalApSubnet : undefined;
            resourceInputs["splitTunnelingAclPath"] = state ? state.splitTunnelingAclPath : undefined;
            resourceInputs["splitTunnelingAcls"] = state ? state.splitTunnelingAcls : undefined;
            resourceInputs["tunMtuDownlink"] = state ? state.tunMtuDownlink : undefined;
            resourceInputs["tunMtuUplink"] = state ? state.tunMtuUplink : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["wanPortMode"] = state ? state.wanPortMode : undefined;
            resourceInputs["wtpId"] = state ? state.wtpId : undefined;
            resourceInputs["wtpMode"] = state ? state.wtpMode : undefined;
            resourceInputs["wtpProfile"] = state ? state.wtpProfile : undefined;
        } else {
            const args = argsOrState as WtpArgs | undefined;
            if ((!args || args.wtpProfile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'wtpProfile'");
            }
            resourceInputs["admin"] = args ? args.admin : undefined;
            resourceInputs["allowaccess"] = args ? args.allowaccess : undefined;
            resourceInputs["apcfgProfile"] = args ? args.apcfgProfile : undefined;
            resourceInputs["bonjourProfile"] = args ? args.bonjourProfile : undefined;
            resourceInputs["coordinateLatitude"] = args ? args.coordinateLatitude : undefined;
            resourceInputs["coordinateLongitude"] = args ? args.coordinateLongitude : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["firmwareProvision"] = args ? args.firmwareProvision : undefined;
            resourceInputs["firmwareProvisionLatest"] = args ? args.firmwareProvisionLatest : undefined;
            resourceInputs["imageDownload"] = args ? args.imageDownload : undefined;
            resourceInputs["index"] = args ? args.index : undefined;
            resourceInputs["ipFragmentPreventing"] = args ? args.ipFragmentPreventing : undefined;
            resourceInputs["lan"] = args ? args.lan : undefined;
            resourceInputs["ledState"] = args ? args.ledState : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["loginPasswd"] = args?.loginPasswd ? pulumi.secret(args.loginPasswd) : undefined;
            resourceInputs["loginPasswdChange"] = args ? args.loginPasswdChange : undefined;
            resourceInputs["meshBridgeEnable"] = args ? args.meshBridgeEnable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overrideAllowaccess"] = args ? args.overrideAllowaccess : undefined;
            resourceInputs["overrideIpFragment"] = args ? args.overrideIpFragment : undefined;
            resourceInputs["overrideLan"] = args ? args.overrideLan : undefined;
            resourceInputs["overrideLedState"] = args ? args.overrideLedState : undefined;
            resourceInputs["overrideLoginPasswdChange"] = args ? args.overrideLoginPasswdChange : undefined;
            resourceInputs["overrideSplitTunnel"] = args ? args.overrideSplitTunnel : undefined;
            resourceInputs["overrideWanPortMode"] = args ? args.overrideWanPortMode : undefined;
            resourceInputs["radio1"] = args ? args.radio1 : undefined;
            resourceInputs["radio2"] = args ? args.radio2 : undefined;
            resourceInputs["radio3"] = args ? args.radio3 : undefined;
            resourceInputs["radio4"] = args ? args.radio4 : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["regionX"] = args ? args.regionX : undefined;
            resourceInputs["regionY"] = args ? args.regionY : undefined;
            resourceInputs["splitTunnelingAclLocalApSubnet"] = args ? args.splitTunnelingAclLocalApSubnet : undefined;
            resourceInputs["splitTunnelingAclPath"] = args ? args.splitTunnelingAclPath : undefined;
            resourceInputs["splitTunnelingAcls"] = args ? args.splitTunnelingAcls : undefined;
            resourceInputs["tunMtuDownlink"] = args ? args.tunMtuDownlink : undefined;
            resourceInputs["tunMtuUplink"] = args ? args.tunMtuUplink : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["wanPortMode"] = args ? args.wanPortMode : undefined;
            resourceInputs["wtpId"] = args ? args.wtpId : undefined;
            resourceInputs["wtpMode"] = args ? args.wtpMode : undefined;
            resourceInputs["wtpProfile"] = args ? args.wtpProfile : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["loginPasswd"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Wtp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Wtp resources.
 */
export interface WtpState {
    /**
     * Configure how the FortiGate operating as a wireless controller discovers and manages this WTP, AP or FortiAP. Valid values: `discovered`, `disable`, `enable`.
     */
    admin?: pulumi.Input<string>;
    /**
     * Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.
     */
    allowaccess?: pulumi.Input<string>;
    /**
     * AP local configuration profile name.
     */
    apcfgProfile?: pulumi.Input<string>;
    /**
     * Bonjour profile name.
     */
    bonjourProfile?: pulumi.Input<string>;
    /**
     * WTP latitude coordinate.
     */
    coordinateLatitude?: pulumi.Input<string>;
    /**
     * WTP longitude coordinate.
     */
    coordinateLongitude?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Firmware version to provision to this FortiAP on bootup (major.minor.build, i.e. 6.2.1234).
     */
    firmwareProvision?: pulumi.Input<string>;
    /**
     * Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
     */
    firmwareProvisionLatest?: pulumi.Input<string>;
    /**
     * Enable/disable WTP image download. Valid values: `enable`, `disable`.
     */
    imageDownload?: pulumi.Input<string>;
    /**
     * Index (0 - 4294967295).
     */
    index?: pulumi.Input<number>;
    /**
     * Method by which IP fragmentation is prevented for CAPWAP tunneled control and data packets (default = tcp-mss-adjust). Valid values: `tcp-mss-adjust`, `icmp-unreachable`.
     */
    ipFragmentPreventing?: pulumi.Input<string>;
    /**
     * WTP LAN port mapping. The structure of `lan` block is documented below.
     */
    lan?: pulumi.Input<inputs.wirelesscontroller.WtpLan>;
    /**
     * Enable to allow the FortiAPs LEDs to light. Disable to keep the LEDs off. You may want to keep the LEDs off so they are not distracting in low light areas etc. Valid values: `enable`, `disable`.
     */
    ledState?: pulumi.Input<string>;
    /**
     * Field for describing the physical location of the WTP, AP or FortiAP.
     */
    location?: pulumi.Input<string>;
    /**
     * Set the managed WTP, FortiAP, or AP's administrator password.
     */
    loginPasswd?: pulumi.Input<string>;
    /**
     * Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
     */
    loginPasswdChange?: pulumi.Input<string>;
    /**
     * Enable/disable mesh Ethernet bridge when WTP is configured as a mesh branch/leaf AP. Valid values: `default`, `enable`, `disable`.
     */
    meshBridgeEnable?: pulumi.Input<string>;
    /**
     * WTP, AP or FortiAP configuration name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable to override the WTP profile management access configuration. Valid values: `enable`, `disable`.
     */
    overrideAllowaccess?: pulumi.Input<string>;
    /**
     * Enable/disable overriding the WTP profile IP fragment prevention setting. Valid values: `enable`, `disable`.
     */
    overrideIpFragment?: pulumi.Input<string>;
    /**
     * Enable to override the WTP profile LAN port setting. Valid values: `enable`, `disable`.
     */
    overrideLan?: pulumi.Input<string>;
    /**
     * Enable to override the profile LED state setting for this FortiAP. You must enable this option to use the led-state command to turn off the FortiAP's LEDs. Valid values: `enable`, `disable`.
     */
    overrideLedState?: pulumi.Input<string>;
    /**
     * Enable to override the WTP profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
     */
    overrideLoginPasswdChange?: pulumi.Input<string>;
    /**
     * Enable/disable overriding the WTP profile split tunneling setting. Valid values: `enable`, `disable`.
     */
    overrideSplitTunnel?: pulumi.Input<string>;
    /**
     * Enable/disable overriding the wan-port-mode in the WTP profile. Valid values: `enable`, `disable`.
     */
    overrideWanPortMode?: pulumi.Input<string>;
    /**
     * Configuration options for radio 1. The structure of `radio1` block is documented below.
     */
    radio1?: pulumi.Input<inputs.wirelesscontroller.WtpRadio1>;
    /**
     * Configuration options for radio 2. The structure of `radio2` block is documented below.
     */
    radio2?: pulumi.Input<inputs.wirelesscontroller.WtpRadio2>;
    /**
     * Configuration options for radio 3. The structure of `radio3` block is documented below.
     */
    radio3?: pulumi.Input<inputs.wirelesscontroller.WtpRadio3>;
    /**
     * Configuration options for radio 4. The structure of `radio4` block is documented below.
     */
    radio4?: pulumi.Input<inputs.wirelesscontroller.WtpRadio4>;
    /**
     * Region name WTP is associated with.
     */
    region?: pulumi.Input<string>;
    /**
     * Relative horizontal region coordinate (between 0 and 1).
     */
    regionX?: pulumi.Input<string>;
    /**
     * Relative vertical region coordinate (between 0 and 1).
     */
    regionY?: pulumi.Input<string>;
    /**
     * Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable). Valid values: `enable`, `disable`.
     */
    splitTunnelingAclLocalApSubnet?: pulumi.Input<string>;
    /**
     * Split tunneling ACL path is local/tunnel. Valid values: `tunnel`, `local`.
     */
    splitTunnelingAclPath?: pulumi.Input<string>;
    /**
     * Split tunneling ACL filter list. The structure of `splitTunnelingAcl` block is documented below.
     */
    splitTunnelingAcls?: pulumi.Input<pulumi.Input<inputs.wirelesscontroller.WtpSplitTunnelingAcl>[]>;
    /**
     * Downlink tunnel MTU in octets. Set the value to either 0 (by default), 576, or 1500.
     */
    tunMtuDownlink?: pulumi.Input<number>;
    /**
     * Uplink tunnel maximum transmission unit (MTU) in octets (eight-bit bytes). Set the value to either 0 (by default), 576, or 1500.
     */
    tunMtuUplink?: pulumi.Input<number>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable using the FortiAP WAN port as a LAN port. Valid values: `wan-lan`, `wan-only`.
     */
    wanPortMode?: pulumi.Input<string>;
    /**
     * WTP ID.
     */
    wtpId?: pulumi.Input<string>;
    /**
     * WTP, AP, or FortiAP operating mode; normal (by default) or remote. A tunnel mode SSID can be assigned to an AP in normal mode but not remote mode, while a local-bridge mode SSID can be assigned to an AP in either normal mode or remote mode. Valid values: `normal`, `remote`.
     */
    wtpMode?: pulumi.Input<string>;
    /**
     * WTP profile name to apply to this WTP, AP or FortiAP.
     */
    wtpProfile?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Wtp resource.
 */
export interface WtpArgs {
    /**
     * Configure how the FortiGate operating as a wireless controller discovers and manages this WTP, AP or FortiAP. Valid values: `discovered`, `disable`, `enable`.
     */
    admin?: pulumi.Input<string>;
    /**
     * Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.
     */
    allowaccess?: pulumi.Input<string>;
    /**
     * AP local configuration profile name.
     */
    apcfgProfile?: pulumi.Input<string>;
    /**
     * Bonjour profile name.
     */
    bonjourProfile?: pulumi.Input<string>;
    /**
     * WTP latitude coordinate.
     */
    coordinateLatitude?: pulumi.Input<string>;
    /**
     * WTP longitude coordinate.
     */
    coordinateLongitude?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Firmware version to provision to this FortiAP on bootup (major.minor.build, i.e. 6.2.1234).
     */
    firmwareProvision?: pulumi.Input<string>;
    /**
     * Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
     */
    firmwareProvisionLatest?: pulumi.Input<string>;
    /**
     * Enable/disable WTP image download. Valid values: `enable`, `disable`.
     */
    imageDownload?: pulumi.Input<string>;
    /**
     * Index (0 - 4294967295).
     */
    index?: pulumi.Input<number>;
    /**
     * Method by which IP fragmentation is prevented for CAPWAP tunneled control and data packets (default = tcp-mss-adjust). Valid values: `tcp-mss-adjust`, `icmp-unreachable`.
     */
    ipFragmentPreventing?: pulumi.Input<string>;
    /**
     * WTP LAN port mapping. The structure of `lan` block is documented below.
     */
    lan?: pulumi.Input<inputs.wirelesscontroller.WtpLan>;
    /**
     * Enable to allow the FortiAPs LEDs to light. Disable to keep the LEDs off. You may want to keep the LEDs off so they are not distracting in low light areas etc. Valid values: `enable`, `disable`.
     */
    ledState?: pulumi.Input<string>;
    /**
     * Field for describing the physical location of the WTP, AP or FortiAP.
     */
    location?: pulumi.Input<string>;
    /**
     * Set the managed WTP, FortiAP, or AP's administrator password.
     */
    loginPasswd?: pulumi.Input<string>;
    /**
     * Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
     */
    loginPasswdChange?: pulumi.Input<string>;
    /**
     * Enable/disable mesh Ethernet bridge when WTP is configured as a mesh branch/leaf AP. Valid values: `default`, `enable`, `disable`.
     */
    meshBridgeEnable?: pulumi.Input<string>;
    /**
     * WTP, AP or FortiAP configuration name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable to override the WTP profile management access configuration. Valid values: `enable`, `disable`.
     */
    overrideAllowaccess?: pulumi.Input<string>;
    /**
     * Enable/disable overriding the WTP profile IP fragment prevention setting. Valid values: `enable`, `disable`.
     */
    overrideIpFragment?: pulumi.Input<string>;
    /**
     * Enable to override the WTP profile LAN port setting. Valid values: `enable`, `disable`.
     */
    overrideLan?: pulumi.Input<string>;
    /**
     * Enable to override the profile LED state setting for this FortiAP. You must enable this option to use the led-state command to turn off the FortiAP's LEDs. Valid values: `enable`, `disable`.
     */
    overrideLedState?: pulumi.Input<string>;
    /**
     * Enable to override the WTP profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
     */
    overrideLoginPasswdChange?: pulumi.Input<string>;
    /**
     * Enable/disable overriding the WTP profile split tunneling setting. Valid values: `enable`, `disable`.
     */
    overrideSplitTunnel?: pulumi.Input<string>;
    /**
     * Enable/disable overriding the wan-port-mode in the WTP profile. Valid values: `enable`, `disable`.
     */
    overrideWanPortMode?: pulumi.Input<string>;
    /**
     * Configuration options for radio 1. The structure of `radio1` block is documented below.
     */
    radio1?: pulumi.Input<inputs.wirelesscontroller.WtpRadio1>;
    /**
     * Configuration options for radio 2. The structure of `radio2` block is documented below.
     */
    radio2?: pulumi.Input<inputs.wirelesscontroller.WtpRadio2>;
    /**
     * Configuration options for radio 3. The structure of `radio3` block is documented below.
     */
    radio3?: pulumi.Input<inputs.wirelesscontroller.WtpRadio3>;
    /**
     * Configuration options for radio 4. The structure of `radio4` block is documented below.
     */
    radio4?: pulumi.Input<inputs.wirelesscontroller.WtpRadio4>;
    /**
     * Region name WTP is associated with.
     */
    region?: pulumi.Input<string>;
    /**
     * Relative horizontal region coordinate (between 0 and 1).
     */
    regionX?: pulumi.Input<string>;
    /**
     * Relative vertical region coordinate (between 0 and 1).
     */
    regionY?: pulumi.Input<string>;
    /**
     * Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable). Valid values: `enable`, `disable`.
     */
    splitTunnelingAclLocalApSubnet?: pulumi.Input<string>;
    /**
     * Split tunneling ACL path is local/tunnel. Valid values: `tunnel`, `local`.
     */
    splitTunnelingAclPath?: pulumi.Input<string>;
    /**
     * Split tunneling ACL filter list. The structure of `splitTunnelingAcl` block is documented below.
     */
    splitTunnelingAcls?: pulumi.Input<pulumi.Input<inputs.wirelesscontroller.WtpSplitTunnelingAcl>[]>;
    /**
     * Downlink tunnel MTU in octets. Set the value to either 0 (by default), 576, or 1500.
     */
    tunMtuDownlink?: pulumi.Input<number>;
    /**
     * Uplink tunnel maximum transmission unit (MTU) in octets (eight-bit bytes). Set the value to either 0 (by default), 576, or 1500.
     */
    tunMtuUplink?: pulumi.Input<number>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable using the FortiAP WAN port as a LAN port. Valid values: `wan-lan`, `wan-only`.
     */
    wanPortMode?: pulumi.Input<string>;
    /**
     * WTP ID.
     */
    wtpId?: pulumi.Input<string>;
    /**
     * WTP, AP, or FortiAP operating mode; normal (by default) or remote. A tunnel mode SSID can be assigned to an AP in normal mode but not remote mode, while a local-bridge mode SSID can be assigned to an AP in either normal mode or remote mode. Valid values: `normal`, `remote`.
     */
    wtpMode?: pulumi.Input<string>;
    /**
     * WTP profile name to apply to this WTP, AP or FortiAP.
     */
    wtpProfile: pulumi.Input<string>;
}
