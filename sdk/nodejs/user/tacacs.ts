// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure TACACS+ server entries.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.user.Tacacs("trname", {
 *     authenType: "auto",
 *     authorization: "disable",
 *     port: 2342,
 *     server: "1.1.1.1",
 * });
 * ```
 *
 * ## Import
 *
 * User Tacacs can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:user/tacacs:Tacacs labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:user/tacacs:Tacacs labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Tacacs extends pulumi.CustomResource {
    /**
     * Get an existing Tacacs resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TacacsState, opts?: pulumi.CustomResourceOptions): Tacacs {
        return new Tacacs(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:user/tacacs:Tacacs';

    /**
     * Returns true if the given object is an instance of Tacacs.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Tacacs {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Tacacs.__pulumiType;
    }

    /**
     * Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
     */
    public readonly authenType!: pulumi.Output<string>;
    /**
     * Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
     */
    public readonly authorization!: pulumi.Output<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    /**
     * Key to access the primary server.
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * TACACS+ server entry name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Port number of the TACACS+ server.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Key to access the secondary server.
     */
    public readonly secondaryKey!: pulumi.Output<string | undefined>;
    /**
     * Secondary TACACS+ server CN domain name or IP address.
     */
    public readonly secondaryServer!: pulumi.Output<string>;
    /**
     * Primary TACACS+ server CN domain name or IP address.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * source IP for communications to TACACS+ server.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Key to access the tertiary server.
     */
    public readonly tertiaryKey!: pulumi.Output<string | undefined>;
    /**
     * Tertiary TACACS+ server CN domain name or IP address.
     */
    public readonly tertiaryServer!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Tacacs resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TacacsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TacacsArgs | TacacsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TacacsState | undefined;
            resourceInputs["authenType"] = state ? state.authenType : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["secondaryKey"] = state ? state.secondaryKey : undefined;
            resourceInputs["secondaryServer"] = state ? state.secondaryServer : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["tertiaryKey"] = state ? state.tertiaryKey : undefined;
            resourceInputs["tertiaryServer"] = state ? state.tertiaryServer : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as TacacsArgs | undefined;
            resourceInputs["authenType"] = args ? args.authenType : undefined;
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["key"] = args?.key ? pulumi.secret(args.key) : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["secondaryKey"] = args?.secondaryKey ? pulumi.secret(args.secondaryKey) : undefined;
            resourceInputs["secondaryServer"] = args ? args.secondaryServer : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["tertiaryKey"] = args?.tertiaryKey ? pulumi.secret(args.tertiaryKey) : undefined;
            resourceInputs["tertiaryServer"] = args ? args.tertiaryServer : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["key", "secondaryKey", "tertiaryKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Tacacs.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Tacacs resources.
 */
export interface TacacsState {
    /**
     * Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
     */
    authenType?: pulumi.Input<string>;
    /**
     * Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
     */
    authorization?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Key to access the primary server.
     */
    key?: pulumi.Input<string>;
    /**
     * TACACS+ server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Port number of the TACACS+ server.
     */
    port?: pulumi.Input<number>;
    /**
     * Key to access the secondary server.
     */
    secondaryKey?: pulumi.Input<string>;
    /**
     * Secondary TACACS+ server CN domain name or IP address.
     */
    secondaryServer?: pulumi.Input<string>;
    /**
     * Primary TACACS+ server CN domain name or IP address.
     */
    server?: pulumi.Input<string>;
    /**
     * source IP for communications to TACACS+ server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Key to access the tertiary server.
     */
    tertiaryKey?: pulumi.Input<string>;
    /**
     * Tertiary TACACS+ server CN domain name or IP address.
     */
    tertiaryServer?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Tacacs resource.
 */
export interface TacacsArgs {
    /**
     * Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
     */
    authenType?: pulumi.Input<string>;
    /**
     * Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
     */
    authorization?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Key to access the primary server.
     */
    key?: pulumi.Input<string>;
    /**
     * TACACS+ server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Port number of the TACACS+ server.
     */
    port?: pulumi.Input<number>;
    /**
     * Key to access the secondary server.
     */
    secondaryKey?: pulumi.Input<string>;
    /**
     * Secondary TACACS+ server CN domain name or IP address.
     */
    secondaryServer?: pulumi.Input<string>;
    /**
     * Primary TACACS+ server CN domain name or IP address.
     */
    server?: pulumi.Input<string>;
    /**
     * source IP for communications to TACACS+ server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Key to access the tertiary server.
     */
    tertiaryKey?: pulumi.Input<string>;
    /**
     * Tertiary TACACS+ server CN domain name or IP address.
     */
    tertiaryServer?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
