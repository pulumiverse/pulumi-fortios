// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure FortiSwitch SNMP v3 users globally. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * SwitchController SnmpUser can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/snmpuser:Snmpuser labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/snmpuser:Snmpuser labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Snmpuser extends pulumi.CustomResource {
    /**
     * Get an existing Snmpuser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnmpuserState, opts?: pulumi.CustomResourceOptions): Snmpuser {
        return new Snmpuser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/snmpuser:Snmpuser';

    /**
     * Returns true if the given object is an instance of Snmpuser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Snmpuser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Snmpuser.__pulumiType;
    }

    /**
     * Authentication protocol.
     */
    public readonly authProto!: pulumi.Output<string>;
    /**
     * Password for authentication protocol.
     */
    public readonly authPwd!: pulumi.Output<string | undefined>;
    /**
     * SNMP user name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Privacy (encryption) protocol.
     */
    public readonly privProto!: pulumi.Output<string>;
    /**
     * Password for privacy (encryption) protocol.
     */
    public readonly privPwd!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable SNMP queries for this user. Valid values: `disable`, `enable`.
     */
    public readonly queries!: pulumi.Output<string>;
    /**
     * SNMPv3 query port (default = 161).
     */
    public readonly queryPort!: pulumi.Output<number>;
    /**
     * Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
     */
    public readonly securityLevel!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Snmpuser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SnmpuserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnmpuserArgs | SnmpuserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnmpuserState | undefined;
            resourceInputs["authProto"] = state ? state.authProto : undefined;
            resourceInputs["authPwd"] = state ? state.authPwd : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["privProto"] = state ? state.privProto : undefined;
            resourceInputs["privPwd"] = state ? state.privPwd : undefined;
            resourceInputs["queries"] = state ? state.queries : undefined;
            resourceInputs["queryPort"] = state ? state.queryPort : undefined;
            resourceInputs["securityLevel"] = state ? state.securityLevel : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SnmpuserArgs | undefined;
            resourceInputs["authProto"] = args ? args.authProto : undefined;
            resourceInputs["authPwd"] = args?.authPwd ? pulumi.secret(args.authPwd) : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["privProto"] = args ? args.privProto : undefined;
            resourceInputs["privPwd"] = args?.privPwd ? pulumi.secret(args.privPwd) : undefined;
            resourceInputs["queries"] = args ? args.queries : undefined;
            resourceInputs["queryPort"] = args ? args.queryPort : undefined;
            resourceInputs["securityLevel"] = args ? args.securityLevel : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["authPwd", "privPwd"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Snmpuser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Snmpuser resources.
 */
export interface SnmpuserState {
    /**
     * Authentication protocol.
     */
    authProto?: pulumi.Input<string>;
    /**
     * Password for authentication protocol.
     */
    authPwd?: pulumi.Input<string>;
    /**
     * SNMP user name.
     */
    name?: pulumi.Input<string>;
    /**
     * Privacy (encryption) protocol.
     */
    privProto?: pulumi.Input<string>;
    /**
     * Password for privacy (encryption) protocol.
     */
    privPwd?: pulumi.Input<string>;
    /**
     * Enable/disable SNMP queries for this user. Valid values: `disable`, `enable`.
     */
    queries?: pulumi.Input<string>;
    /**
     * SNMPv3 query port (default = 161).
     */
    queryPort?: pulumi.Input<number>;
    /**
     * Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
     */
    securityLevel?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Snmpuser resource.
 */
export interface SnmpuserArgs {
    /**
     * Authentication protocol.
     */
    authProto?: pulumi.Input<string>;
    /**
     * Password for authentication protocol.
     */
    authPwd?: pulumi.Input<string>;
    /**
     * SNMP user name.
     */
    name?: pulumi.Input<string>;
    /**
     * Privacy (encryption) protocol.
     */
    privProto?: pulumi.Input<string>;
    /**
     * Password for privacy (encryption) protocol.
     */
    privPwd?: pulumi.Input<string>;
    /**
     * Enable/disable SNMP queries for this user. Valid values: `disable`, `enable`.
     */
    queries?: pulumi.Input<string>;
    /**
     * SNMPv3 query port (default = 161).
     */
    queryPort?: pulumi.Input<number>;
    /**
     * Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
     */
    securityLevel?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
