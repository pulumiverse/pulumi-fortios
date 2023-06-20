// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch global settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.SwitchcontrollerGlobal("trname", {
 *     allowMultipleInterfaces: "disable",
 *     httpsImagePush: "disable",
 *     logMacLimitViolations: "disable",
 *     macAgingInterval: 332,
 *     macRetentionPeriod: 24,
 *     macViolationTimer: 0,
 * });
 * ```
 *
 * ## Import
 *
 * SwitchController Global can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/switchcontrollerGlobal:SwitchcontrollerGlobal labelname SwitchControllerGlobal
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchcontrollerGlobal:SwitchcontrollerGlobal labelname SwitchControllerGlobal
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchcontrollerGlobal extends pulumi.CustomResource {
    /**
     * Get an existing SwitchcontrollerGlobal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchcontrollerGlobalState, opts?: pulumi.CustomResourceOptions): SwitchcontrollerGlobal {
        return new SwitchcontrollerGlobal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchcontrollerGlobal:SwitchcontrollerGlobal';

    /**
     * Returns true if the given object is an instance of SwitchcontrollerGlobal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchcontrollerGlobal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchcontrollerGlobal.__pulumiType;
    }

    /**
     * Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
     */
    public readonly allowMultipleInterfaces!: pulumi.Output<string>;
    /**
     * Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
     */
    public readonly bounceQuarantinedLink!: pulumi.Output<string>;
    /**
     * List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
     */
    public readonly customCommands!: pulumi.Output<outputs.SwitchcontrollerGlobalCustomCommand[] | undefined>;
    /**
     * Default VLAN for ports when added to the virtual-switch.
     */
    public readonly defaultVirtualSwitchVlan!: pulumi.Output<string>;
    /**
     * Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
     */
    public readonly dhcpServerAccessList!: pulumi.Output<string>;
    /**
     * Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
     */
    public readonly disableDiscoveries!: pulumi.Output<outputs.SwitchcontrollerGlobalDisableDiscovery[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
     */
    public readonly fipsEnforce!: pulumi.Output<string>;
    /**
     * Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
     */
    public readonly firmwareProvisionOnAuthorization!: pulumi.Output<string>;
    /**
     * Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
     */
    public readonly httpsImagePush!: pulumi.Output<string>;
    /**
     * Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
     */
    public readonly logMacLimitViolations!: pulumi.Output<string>;
    /**
     * Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
     */
    public readonly macAgingInterval!: pulumi.Output<number>;
    /**
     * Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
     */
    public readonly macEventLogging!: pulumi.Output<string>;
    /**
     * Time in hours after which an inactive MAC is removed from client DB.
     */
    public readonly macRetentionPeriod!: pulumi.Output<number>;
    /**
     * Set timeout for Learning Limit Violations (0 = disabled).
     */
    public readonly macViolationTimer!: pulumi.Output<number>;
    /**
     * Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
     */
    public readonly quarantineMode!: pulumi.Output<string>;
    /**
     * Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
     */
    public readonly snDnsResolution!: pulumi.Output<string>;
    /**
     * Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
     */
    public readonly updateUserDevice!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
     */
    public readonly vlanAllMode!: pulumi.Output<string>;
    /**
     * FortiLink VLAN optimization. Valid values: `enable`, `disable`.
     */
    public readonly vlanOptimization!: pulumi.Output<string>;

    /**
     * Create a SwitchcontrollerGlobal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchcontrollerGlobalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchcontrollerGlobalArgs | SwitchcontrollerGlobalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchcontrollerGlobalState | undefined;
            resourceInputs["allowMultipleInterfaces"] = state ? state.allowMultipleInterfaces : undefined;
            resourceInputs["bounceQuarantinedLink"] = state ? state.bounceQuarantinedLink : undefined;
            resourceInputs["customCommands"] = state ? state.customCommands : undefined;
            resourceInputs["defaultVirtualSwitchVlan"] = state ? state.defaultVirtualSwitchVlan : undefined;
            resourceInputs["dhcpServerAccessList"] = state ? state.dhcpServerAccessList : undefined;
            resourceInputs["disableDiscoveries"] = state ? state.disableDiscoveries : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fipsEnforce"] = state ? state.fipsEnforce : undefined;
            resourceInputs["firmwareProvisionOnAuthorization"] = state ? state.firmwareProvisionOnAuthorization : undefined;
            resourceInputs["httpsImagePush"] = state ? state.httpsImagePush : undefined;
            resourceInputs["logMacLimitViolations"] = state ? state.logMacLimitViolations : undefined;
            resourceInputs["macAgingInterval"] = state ? state.macAgingInterval : undefined;
            resourceInputs["macEventLogging"] = state ? state.macEventLogging : undefined;
            resourceInputs["macRetentionPeriod"] = state ? state.macRetentionPeriod : undefined;
            resourceInputs["macViolationTimer"] = state ? state.macViolationTimer : undefined;
            resourceInputs["quarantineMode"] = state ? state.quarantineMode : undefined;
            resourceInputs["snDnsResolution"] = state ? state.snDnsResolution : undefined;
            resourceInputs["updateUserDevice"] = state ? state.updateUserDevice : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vlanAllMode"] = state ? state.vlanAllMode : undefined;
            resourceInputs["vlanOptimization"] = state ? state.vlanOptimization : undefined;
        } else {
            const args = argsOrState as SwitchcontrollerGlobalArgs | undefined;
            resourceInputs["allowMultipleInterfaces"] = args ? args.allowMultipleInterfaces : undefined;
            resourceInputs["bounceQuarantinedLink"] = args ? args.bounceQuarantinedLink : undefined;
            resourceInputs["customCommands"] = args ? args.customCommands : undefined;
            resourceInputs["defaultVirtualSwitchVlan"] = args ? args.defaultVirtualSwitchVlan : undefined;
            resourceInputs["dhcpServerAccessList"] = args ? args.dhcpServerAccessList : undefined;
            resourceInputs["disableDiscoveries"] = args ? args.disableDiscoveries : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fipsEnforce"] = args ? args.fipsEnforce : undefined;
            resourceInputs["firmwareProvisionOnAuthorization"] = args ? args.firmwareProvisionOnAuthorization : undefined;
            resourceInputs["httpsImagePush"] = args ? args.httpsImagePush : undefined;
            resourceInputs["logMacLimitViolations"] = args ? args.logMacLimitViolations : undefined;
            resourceInputs["macAgingInterval"] = args ? args.macAgingInterval : undefined;
            resourceInputs["macEventLogging"] = args ? args.macEventLogging : undefined;
            resourceInputs["macRetentionPeriod"] = args ? args.macRetentionPeriod : undefined;
            resourceInputs["macViolationTimer"] = args ? args.macViolationTimer : undefined;
            resourceInputs["quarantineMode"] = args ? args.quarantineMode : undefined;
            resourceInputs["snDnsResolution"] = args ? args.snDnsResolution : undefined;
            resourceInputs["updateUserDevice"] = args ? args.updateUserDevice : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vlanAllMode"] = args ? args.vlanAllMode : undefined;
            resourceInputs["vlanOptimization"] = args ? args.vlanOptimization : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchcontrollerGlobal.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchcontrollerGlobal resources.
 */
export interface SwitchcontrollerGlobalState {
    /**
     * Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
     */
    allowMultipleInterfaces?: pulumi.Input<string>;
    /**
     * Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
     */
    bounceQuarantinedLink?: pulumi.Input<string>;
    /**
     * List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
     */
    customCommands?: pulumi.Input<pulumi.Input<inputs.SwitchcontrollerGlobalCustomCommand>[]>;
    /**
     * Default VLAN for ports when added to the virtual-switch.
     */
    defaultVirtualSwitchVlan?: pulumi.Input<string>;
    /**
     * Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
     */
    dhcpServerAccessList?: pulumi.Input<string>;
    /**
     * Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
     */
    disableDiscoveries?: pulumi.Input<pulumi.Input<inputs.SwitchcontrollerGlobalDisableDiscovery>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
     */
    fipsEnforce?: pulumi.Input<string>;
    /**
     * Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
     */
    firmwareProvisionOnAuthorization?: pulumi.Input<string>;
    /**
     * Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
     */
    httpsImagePush?: pulumi.Input<string>;
    /**
     * Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
     */
    logMacLimitViolations?: pulumi.Input<string>;
    /**
     * Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
     */
    macAgingInterval?: pulumi.Input<number>;
    /**
     * Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
     */
    macEventLogging?: pulumi.Input<string>;
    /**
     * Time in hours after which an inactive MAC is removed from client DB.
     */
    macRetentionPeriod?: pulumi.Input<number>;
    /**
     * Set timeout for Learning Limit Violations (0 = disabled).
     */
    macViolationTimer?: pulumi.Input<number>;
    /**
     * Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
     */
    quarantineMode?: pulumi.Input<string>;
    /**
     * Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
     */
    snDnsResolution?: pulumi.Input<string>;
    /**
     * Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
     */
    updateUserDevice?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
     */
    vlanAllMode?: pulumi.Input<string>;
    /**
     * FortiLink VLAN optimization. Valid values: `enable`, `disable`.
     */
    vlanOptimization?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchcontrollerGlobal resource.
 */
export interface SwitchcontrollerGlobalArgs {
    /**
     * Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
     */
    allowMultipleInterfaces?: pulumi.Input<string>;
    /**
     * Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
     */
    bounceQuarantinedLink?: pulumi.Input<string>;
    /**
     * List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
     */
    customCommands?: pulumi.Input<pulumi.Input<inputs.SwitchcontrollerGlobalCustomCommand>[]>;
    /**
     * Default VLAN for ports when added to the virtual-switch.
     */
    defaultVirtualSwitchVlan?: pulumi.Input<string>;
    /**
     * Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
     */
    dhcpServerAccessList?: pulumi.Input<string>;
    /**
     * Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
     */
    disableDiscoveries?: pulumi.Input<pulumi.Input<inputs.SwitchcontrollerGlobalDisableDiscovery>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
     */
    fipsEnforce?: pulumi.Input<string>;
    /**
     * Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
     */
    firmwareProvisionOnAuthorization?: pulumi.Input<string>;
    /**
     * Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
     */
    httpsImagePush?: pulumi.Input<string>;
    /**
     * Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
     */
    logMacLimitViolations?: pulumi.Input<string>;
    /**
     * Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
     */
    macAgingInterval?: pulumi.Input<number>;
    /**
     * Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
     */
    macEventLogging?: pulumi.Input<string>;
    /**
     * Time in hours after which an inactive MAC is removed from client DB.
     */
    macRetentionPeriod?: pulumi.Input<number>;
    /**
     * Set timeout for Learning Limit Violations (0 = disabled).
     */
    macViolationTimer?: pulumi.Input<number>;
    /**
     * Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
     */
    quarantineMode?: pulumi.Input<string>;
    /**
     * Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
     */
    snDnsResolution?: pulumi.Input<string>;
    /**
     * Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
     */
    updateUserDevice?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
     */
    vlanAllMode?: pulumi.Input<string>;
    /**
     * FortiLink VLAN optimization. Valid values: `enable`, `disable`.
     */
    vlanOptimization?: pulumi.Input<string>;
}