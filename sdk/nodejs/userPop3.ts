// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * POP3 server entry configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.UserPop3("trname", {
 *     port: 0,
 *     secure: "pop3s",
 *     server: "1.1.1.1",
 *     sslMinProtoVersion: "default",
 * });
 * ```
 *
 * ## Import
 *
 * User Pop3 can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/userPop3:UserPop3 labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/userPop3:UserPop3 labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class UserPop3 extends pulumi.CustomResource {
    /**
     * Get an existing UserPop3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserPop3State, opts?: pulumi.CustomResourceOptions): UserPop3 {
        return new UserPop3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/userPop3:UserPop3';

    /**
     * Returns true if the given object is an instance of UserPop3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPop3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPop3.__pulumiType;
    }

    /**
     * POP3 server entry name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * POP3 service port number.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * SSL connection. Valid values: `none`, `starttls`, `pop3s`.
     */
    public readonly secure!: pulumi.Output<string>;
    /**
     * {<name_str|ip_str>} server domain name or IP.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a UserPop3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserPop3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserPop3Args | UserPop3State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserPop3State | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["secure"] = state ? state.secure : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as UserPop3Args | undefined;
            if ((!args || args.server === undefined) && !opts.urn) {
                throw new Error("Missing required property 'server'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["secure"] = args ? args.secure : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserPop3.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserPop3 resources.
 */
export interface UserPop3State {
    /**
     * POP3 server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * POP3 service port number.
     */
    port?: pulumi.Input<number>;
    /**
     * SSL connection. Valid values: `none`, `starttls`, `pop3s`.
     */
    secure?: pulumi.Input<string>;
    /**
     * {<name_str|ip_str>} server domain name or IP.
     */
    server?: pulumi.Input<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserPop3 resource.
 */
export interface UserPop3Args {
    /**
     * POP3 server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * POP3 service port number.
     */
    port?: pulumi.Input<number>;
    /**
     * SSL connection. Valid values: `none`, `starttls`, `pop3s`.
     */
    secure?: pulumi.Input<string>;
    /**
     * {<name_str|ip_str>} server domain name or IP.
     */
    server: pulumi.Input<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
