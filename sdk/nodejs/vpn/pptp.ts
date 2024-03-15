// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure PPTP.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.vpn.Pptp("trname", {
 *     eip: "1.1.1.22",
 *     ipMode: "range",
 *     localIp: "0.0.0.0",
 *     sip: "1.1.1.1",
 *     status: "enable",
 *     usrgrp: "Guest-group",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Vpn Pptp can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:vpn/pptp:Pptp labelname VpnPptp
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:vpn/pptp:Pptp labelname VpnPptp
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Pptp extends pulumi.CustomResource {
    /**
     * Get an existing Pptp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PptpState, opts?: pulumi.CustomResourceOptions): Pptp {
        return new Pptp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:vpn/pptp:Pptp';

    /**
     * Returns true if the given object is an instance of Pptp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pptp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pptp.__pulumiType;
    }

    /**
     * End IP.
     */
    public readonly eip!: pulumi.Output<string>;
    /**
     * IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
     */
    public readonly ipMode!: pulumi.Output<string>;
    /**
     * Local IP to be used for peer's remote IP.
     */
    public readonly localIp!: pulumi.Output<string>;
    /**
     * Start IP.
     */
    public readonly sip!: pulumi.Output<string>;
    /**
     * Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * User group.
     */
    public readonly usrgrp!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Pptp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PptpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PptpArgs | PptpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PptpState | undefined;
            resourceInputs["eip"] = state ? state.eip : undefined;
            resourceInputs["ipMode"] = state ? state.ipMode : undefined;
            resourceInputs["localIp"] = state ? state.localIp : undefined;
            resourceInputs["sip"] = state ? state.sip : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["usrgrp"] = state ? state.usrgrp : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as PptpArgs | undefined;
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["eip"] = args ? args.eip : undefined;
            resourceInputs["ipMode"] = args ? args.ipMode : undefined;
            resourceInputs["localIp"] = args ? args.localIp : undefined;
            resourceInputs["sip"] = args ? args.sip : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["usrgrp"] = args ? args.usrgrp : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pptp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pptp resources.
 */
export interface PptpState {
    /**
     * End IP.
     */
    eip?: pulumi.Input<string>;
    /**
     * IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
     */
    ipMode?: pulumi.Input<string>;
    /**
     * Local IP to be used for peer's remote IP.
     */
    localIp?: pulumi.Input<string>;
    /**
     * Start IP.
     */
    sip?: pulumi.Input<string>;
    /**
     * Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * User group.
     */
    usrgrp?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pptp resource.
 */
export interface PptpArgs {
    /**
     * End IP.
     */
    eip?: pulumi.Input<string>;
    /**
     * IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
     */
    ipMode?: pulumi.Input<string>;
    /**
     * Local IP to be used for peer's remote IP.
     */
    localIp?: pulumi.Input<string>;
    /**
     * Start IP.
     */
    sip?: pulumi.Input<string>;
    /**
     * Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
     */
    status: pulumi.Input<string>;
    /**
     * User group.
     */
    usrgrp?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
