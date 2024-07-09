// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure Fortinet Single Sign On (FSSO) agents.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.user.Fsso("trname", {
 *     port: 32381,
 *     port2: 8000,
 *     port3: 8000,
 *     port4: 8000,
 *     port5: 8000,
 *     server: "1.1.1.1",
 *     sourceIp: "0.0.0.0",
 *     sourceIp6: "::",
 * });
 * ```
 *
 * ## Import
 *
 * User Fsso can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:user/fsso:Fsso labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:user/fsso:Fsso labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Fsso extends pulumi.CustomResource {
    /**
     * Get an existing Fsso resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FssoState, opts?: pulumi.CustomResourceOptions): Fsso {
        return new Fsso(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:user/fsso:Fsso';

    /**
     * Returns true if the given object is an instance of Fsso.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Fsso {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Fsso.__pulumiType;
    }

    /**
     * Interval in minutes within to fetch groups from FSSO server, or unset to disable.
     */
    public readonly groupPollInterval!: pulumi.Output<number>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    /**
     * Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable`, `disable`.
     */
    public readonly ldapPoll!: pulumi.Output<string>;
    /**
     * Filter used to fetch groups.
     */
    public readonly ldapPollFilter!: pulumi.Output<string>;
    /**
     * Interval in minutes within to fetch groups from LDAP server.
     */
    public readonly ldapPollInterval!: pulumi.Output<number>;
    /**
     * LDAP server to get group information.
     */
    public readonly ldapServer!: pulumi.Output<string>;
    /**
     * Interval in minutes to keep logons after FSSO server down.
     */
    public readonly logonTimeout!: pulumi.Output<number>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password of the first FSSO collector agent.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Password of the second FSSO collector agent.
     */
    public readonly password2!: pulumi.Output<string | undefined>;
    /**
     * Password of the third FSSO collector agent.
     */
    public readonly password3!: pulumi.Output<string | undefined>;
    /**
     * Password of the fourth FSSO collector agent.
     */
    public readonly password4!: pulumi.Output<string | undefined>;
    /**
     * Password of the fifth FSSO collector agent.
     */
    public readonly password5!: pulumi.Output<string | undefined>;
    /**
     * Port of the first FSSO collector agent.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Port of the second FSSO collector agent.
     */
    public readonly port2!: pulumi.Output<number>;
    /**
     * Port of the third FSSO collector agent.
     */
    public readonly port3!: pulumi.Output<number>;
    /**
     * Port of the fourth FSSO collector agent.
     */
    public readonly port4!: pulumi.Output<number>;
    /**
     * Port of the fifth FSSO collector agent.
     */
    public readonly port5!: pulumi.Output<number>;
    /**
     * Domain name or IP address of the first FSSO collector agent.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * Domain name or IP address of the second FSSO collector agent.
     */
    public readonly server2!: pulumi.Output<string>;
    /**
     * Domain name or IP address of the third FSSO collector agent.
     */
    public readonly server3!: pulumi.Output<string>;
    /**
     * Domain name or IP address of the fourth FSSO collector agent.
     */
    public readonly server4!: pulumi.Output<string>;
    /**
     * Domain name or IP address of the fifth FSSO collector agent.
     */
    public readonly server5!: pulumi.Output<string>;
    /**
     * Server Name Indication.
     */
    public readonly sni!: pulumi.Output<string>;
    /**
     * Source IP for communications to FSSO agent.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * IPv6 source for communications to FSSO agent.
     */
    public readonly sourceIp6!: pulumi.Output<string>;
    /**
     * Enable/disable use of SSL. Valid values: `enable`, `disable`.
     */
    public readonly ssl!: pulumi.Output<string>;
    /**
     * Enable/disable server host/IP verification. Valid values: `enable`, `disable`.
     */
    public readonly sslServerHostIpCheck!: pulumi.Output<string>;
    /**
     * Trusted server certificate or CA certificate.
     */
    public readonly sslTrustedCert!: pulumi.Output<string>;
    /**
     * Server type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * LDAP server to get user information.
     */
    public readonly userInfoServer!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Fsso resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FssoArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FssoArgs | FssoState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FssoState | undefined;
            resourceInputs["groupPollInterval"] = state ? state.groupPollInterval : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["ldapPoll"] = state ? state.ldapPoll : undefined;
            resourceInputs["ldapPollFilter"] = state ? state.ldapPollFilter : undefined;
            resourceInputs["ldapPollInterval"] = state ? state.ldapPollInterval : undefined;
            resourceInputs["ldapServer"] = state ? state.ldapServer : undefined;
            resourceInputs["logonTimeout"] = state ? state.logonTimeout : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["password2"] = state ? state.password2 : undefined;
            resourceInputs["password3"] = state ? state.password3 : undefined;
            resourceInputs["password4"] = state ? state.password4 : undefined;
            resourceInputs["password5"] = state ? state.password5 : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["port2"] = state ? state.port2 : undefined;
            resourceInputs["port3"] = state ? state.port3 : undefined;
            resourceInputs["port4"] = state ? state.port4 : undefined;
            resourceInputs["port5"] = state ? state.port5 : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["server2"] = state ? state.server2 : undefined;
            resourceInputs["server3"] = state ? state.server3 : undefined;
            resourceInputs["server4"] = state ? state.server4 : undefined;
            resourceInputs["server5"] = state ? state.server5 : undefined;
            resourceInputs["sni"] = state ? state.sni : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sourceIp6"] = state ? state.sourceIp6 : undefined;
            resourceInputs["ssl"] = state ? state.ssl : undefined;
            resourceInputs["sslServerHostIpCheck"] = state ? state.sslServerHostIpCheck : undefined;
            resourceInputs["sslTrustedCert"] = state ? state.sslTrustedCert : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["userInfoServer"] = state ? state.userInfoServer : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FssoArgs | undefined;
            if ((!args || args.server === undefined) && !opts.urn) {
                throw new Error("Missing required property 'server'");
            }
            resourceInputs["groupPollInterval"] = args ? args.groupPollInterval : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["ldapPoll"] = args ? args.ldapPoll : undefined;
            resourceInputs["ldapPollFilter"] = args ? args.ldapPollFilter : undefined;
            resourceInputs["ldapPollInterval"] = args ? args.ldapPollInterval : undefined;
            resourceInputs["ldapServer"] = args ? args.ldapServer : undefined;
            resourceInputs["logonTimeout"] = args ? args.logonTimeout : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["password2"] = args?.password2 ? pulumi.secret(args.password2) : undefined;
            resourceInputs["password3"] = args?.password3 ? pulumi.secret(args.password3) : undefined;
            resourceInputs["password4"] = args?.password4 ? pulumi.secret(args.password4) : undefined;
            resourceInputs["password5"] = args?.password5 ? pulumi.secret(args.password5) : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["port2"] = args ? args.port2 : undefined;
            resourceInputs["port3"] = args ? args.port3 : undefined;
            resourceInputs["port4"] = args ? args.port4 : undefined;
            resourceInputs["port5"] = args ? args.port5 : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["server2"] = args ? args.server2 : undefined;
            resourceInputs["server3"] = args ? args.server3 : undefined;
            resourceInputs["server4"] = args ? args.server4 : undefined;
            resourceInputs["server5"] = args ? args.server5 : undefined;
            resourceInputs["sni"] = args ? args.sni : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sourceIp6"] = args ? args.sourceIp6 : undefined;
            resourceInputs["ssl"] = args ? args.ssl : undefined;
            resourceInputs["sslServerHostIpCheck"] = args ? args.sslServerHostIpCheck : undefined;
            resourceInputs["sslTrustedCert"] = args ? args.sslTrustedCert : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["userInfoServer"] = args ? args.userInfoServer : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password", "password2", "password3", "password4", "password5"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Fsso.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Fsso resources.
 */
export interface FssoState {
    /**
     * Interval in minutes within to fetch groups from FSSO server, or unset to disable.
     */
    groupPollInterval?: pulumi.Input<number>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable`, `disable`.
     */
    ldapPoll?: pulumi.Input<string>;
    /**
     * Filter used to fetch groups.
     */
    ldapPollFilter?: pulumi.Input<string>;
    /**
     * Interval in minutes within to fetch groups from LDAP server.
     */
    ldapPollInterval?: pulumi.Input<number>;
    /**
     * LDAP server to get group information.
     */
    ldapServer?: pulumi.Input<string>;
    /**
     * Interval in minutes to keep logons after FSSO server down.
     */
    logonTimeout?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Password of the first FSSO collector agent.
     */
    password?: pulumi.Input<string>;
    /**
     * Password of the second FSSO collector agent.
     */
    password2?: pulumi.Input<string>;
    /**
     * Password of the third FSSO collector agent.
     */
    password3?: pulumi.Input<string>;
    /**
     * Password of the fourth FSSO collector agent.
     */
    password4?: pulumi.Input<string>;
    /**
     * Password of the fifth FSSO collector agent.
     */
    password5?: pulumi.Input<string>;
    /**
     * Port of the first FSSO collector agent.
     */
    port?: pulumi.Input<number>;
    /**
     * Port of the second FSSO collector agent.
     */
    port2?: pulumi.Input<number>;
    /**
     * Port of the third FSSO collector agent.
     */
    port3?: pulumi.Input<number>;
    /**
     * Port of the fourth FSSO collector agent.
     */
    port4?: pulumi.Input<number>;
    /**
     * Port of the fifth FSSO collector agent.
     */
    port5?: pulumi.Input<number>;
    /**
     * Domain name or IP address of the first FSSO collector agent.
     */
    server?: pulumi.Input<string>;
    /**
     * Domain name or IP address of the second FSSO collector agent.
     */
    server2?: pulumi.Input<string>;
    /**
     * Domain name or IP address of the third FSSO collector agent.
     */
    server3?: pulumi.Input<string>;
    /**
     * Domain name or IP address of the fourth FSSO collector agent.
     */
    server4?: pulumi.Input<string>;
    /**
     * Domain name or IP address of the fifth FSSO collector agent.
     */
    server5?: pulumi.Input<string>;
    /**
     * Server Name Indication.
     */
    sni?: pulumi.Input<string>;
    /**
     * Source IP for communications to FSSO agent.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * IPv6 source for communications to FSSO agent.
     */
    sourceIp6?: pulumi.Input<string>;
    /**
     * Enable/disable use of SSL. Valid values: `enable`, `disable`.
     */
    ssl?: pulumi.Input<string>;
    /**
     * Enable/disable server host/IP verification. Valid values: `enable`, `disable`.
     */
    sslServerHostIpCheck?: pulumi.Input<string>;
    /**
     * Trusted server certificate or CA certificate.
     */
    sslTrustedCert?: pulumi.Input<string>;
    /**
     * Server type.
     */
    type?: pulumi.Input<string>;
    /**
     * LDAP server to get user information.
     */
    userInfoServer?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Fsso resource.
 */
export interface FssoArgs {
    /**
     * Interval in minutes within to fetch groups from FSSO server, or unset to disable.
     */
    groupPollInterval?: pulumi.Input<number>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable`, `disable`.
     */
    ldapPoll?: pulumi.Input<string>;
    /**
     * Filter used to fetch groups.
     */
    ldapPollFilter?: pulumi.Input<string>;
    /**
     * Interval in minutes within to fetch groups from LDAP server.
     */
    ldapPollInterval?: pulumi.Input<number>;
    /**
     * LDAP server to get group information.
     */
    ldapServer?: pulumi.Input<string>;
    /**
     * Interval in minutes to keep logons after FSSO server down.
     */
    logonTimeout?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Password of the first FSSO collector agent.
     */
    password?: pulumi.Input<string>;
    /**
     * Password of the second FSSO collector agent.
     */
    password2?: pulumi.Input<string>;
    /**
     * Password of the third FSSO collector agent.
     */
    password3?: pulumi.Input<string>;
    /**
     * Password of the fourth FSSO collector agent.
     */
    password4?: pulumi.Input<string>;
    /**
     * Password of the fifth FSSO collector agent.
     */
    password5?: pulumi.Input<string>;
    /**
     * Port of the first FSSO collector agent.
     */
    port?: pulumi.Input<number>;
    /**
     * Port of the second FSSO collector agent.
     */
    port2?: pulumi.Input<number>;
    /**
     * Port of the third FSSO collector agent.
     */
    port3?: pulumi.Input<number>;
    /**
     * Port of the fourth FSSO collector agent.
     */
    port4?: pulumi.Input<number>;
    /**
     * Port of the fifth FSSO collector agent.
     */
    port5?: pulumi.Input<number>;
    /**
     * Domain name or IP address of the first FSSO collector agent.
     */
    server: pulumi.Input<string>;
    /**
     * Domain name or IP address of the second FSSO collector agent.
     */
    server2?: pulumi.Input<string>;
    /**
     * Domain name or IP address of the third FSSO collector agent.
     */
    server3?: pulumi.Input<string>;
    /**
     * Domain name or IP address of the fourth FSSO collector agent.
     */
    server4?: pulumi.Input<string>;
    /**
     * Domain name or IP address of the fifth FSSO collector agent.
     */
    server5?: pulumi.Input<string>;
    /**
     * Server Name Indication.
     */
    sni?: pulumi.Input<string>;
    /**
     * Source IP for communications to FSSO agent.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * IPv6 source for communications to FSSO agent.
     */
    sourceIp6?: pulumi.Input<string>;
    /**
     * Enable/disable use of SSL. Valid values: `enable`, `disable`.
     */
    ssl?: pulumi.Input<string>;
    /**
     * Enable/disable server host/IP verification. Valid values: `enable`, `disable`.
     */
    sslServerHostIpCheck?: pulumi.Input<string>;
    /**
     * Trusted server certificate or CA certificate.
     */
    sslTrustedCert?: pulumi.Input<string>;
    /**
     * Server type.
     */
    type?: pulumi.Input<string>;
    /**
     * LDAP server to get user information.
     */
    userInfoServer?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
