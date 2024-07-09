// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure SSH config. Applies to FortiOS Version `>= 7.4.4`.
 *
 * ## Import
 *
 * System SshConfig can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/sshconfig:Sshconfig labelname SystemSshConfig
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/sshconfig:Sshconfig labelname SystemSshConfig
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Sshconfig extends pulumi.CustomResource {
    /**
     * Get an existing Sshconfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SshconfigState, opts?: pulumi.CustomResourceOptions): Sshconfig {
        return new Sshconfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/sshconfig:Sshconfig';

    /**
     * Returns true if the given object is an instance of Sshconfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Sshconfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Sshconfig.__pulumiType;
    }

    /**
     * Select one or more SSH ciphers. Valid values: `chacha20-poly1305@openssh.com`, `aes128-ctr`, `aes192-ctr`, `aes256-ctr`, `arcfour256`, `arcfour128`, `aes128-cbc`, `3des-cbc`, `blowfish-cbc`, `cast128-cbc`, `aes192-cbc`, `aes256-cbc`, `arcfour`, `rijndael-cbc@lysator.liu.se`, `aes128-gcm@openssh.com`, `aes256-gcm@openssh.com`.
     */
    public readonly sshEncAlgo!: pulumi.Output<string>;
    /**
     * Config SSH host key.
     */
    public readonly sshHsk!: pulumi.Output<string>;
    /**
     * Select one or more SSH hostkey algorithms. Valid values: `ssh-rsa`, `ecdsa-sha2-nistp521`, `ecdsa-sha2-nistp384`, `ecdsa-sha2-nistp256`, `rsa-sha2-256`, `rsa-sha2-512`, `ssh-ed25519`.
     */
    public readonly sshHskAlgo!: pulumi.Output<string>;
    /**
     * Enable/disable SSH host key override in SSH daemon. Valid values: `disable`, `enable`.
     */
    public readonly sshHskOverride!: pulumi.Output<string>;
    /**
     * Password for ssh-hostkey.
     */
    public readonly sshHskPassword!: pulumi.Output<string | undefined>;
    /**
     * Select one or more SSH kex algorithms. Valid values: `diffie-hellman-group1-sha1`, `diffie-hellman-group14-sha1`, `diffie-hellman-group14-sha256`, `diffie-hellman-group16-sha512`, `diffie-hellman-group18-sha512`, `diffie-hellman-group-exchange-sha1`, `diffie-hellman-group-exchange-sha256`, `curve25519-sha256@libssh.org`, `ecdh-sha2-nistp256`, `ecdh-sha2-nistp384`, `ecdh-sha2-nistp521`.
     */
    public readonly sshKexAlgo!: pulumi.Output<string>;
    /**
     * Select one or more SSH MAC algorithms. Valid values: `hmac-md5`, `hmac-md5-etm@openssh.com`, `hmac-md5-96`, `hmac-md5-96-etm@openssh.com`, `hmac-sha1`, `hmac-sha1-etm@openssh.com`, `hmac-sha2-256`, `hmac-sha2-256-etm@openssh.com`, `hmac-sha2-512`, `hmac-sha2-512-etm@openssh.com`, `hmac-ripemd160`, `hmac-ripemd160@openssh.com`, `hmac-ripemd160-etm@openssh.com`, `umac-64@openssh.com`, `umac-128@openssh.com`, `umac-64-etm@openssh.com`, `umac-128-etm@openssh.com`.
     */
    public readonly sshMacAlgo!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Sshconfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SshconfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SshconfigArgs | SshconfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SshconfigState | undefined;
            resourceInputs["sshEncAlgo"] = state ? state.sshEncAlgo : undefined;
            resourceInputs["sshHsk"] = state ? state.sshHsk : undefined;
            resourceInputs["sshHskAlgo"] = state ? state.sshHskAlgo : undefined;
            resourceInputs["sshHskOverride"] = state ? state.sshHskOverride : undefined;
            resourceInputs["sshHskPassword"] = state ? state.sshHskPassword : undefined;
            resourceInputs["sshKexAlgo"] = state ? state.sshKexAlgo : undefined;
            resourceInputs["sshMacAlgo"] = state ? state.sshMacAlgo : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SshconfigArgs | undefined;
            resourceInputs["sshEncAlgo"] = args ? args.sshEncAlgo : undefined;
            resourceInputs["sshHsk"] = args ? args.sshHsk : undefined;
            resourceInputs["sshHskAlgo"] = args ? args.sshHskAlgo : undefined;
            resourceInputs["sshHskOverride"] = args ? args.sshHskOverride : undefined;
            resourceInputs["sshHskPassword"] = args ? args.sshHskPassword : undefined;
            resourceInputs["sshKexAlgo"] = args ? args.sshKexAlgo : undefined;
            resourceInputs["sshMacAlgo"] = args ? args.sshMacAlgo : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Sshconfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Sshconfig resources.
 */
export interface SshconfigState {
    /**
     * Select one or more SSH ciphers. Valid values: `chacha20-poly1305@openssh.com`, `aes128-ctr`, `aes192-ctr`, `aes256-ctr`, `arcfour256`, `arcfour128`, `aes128-cbc`, `3des-cbc`, `blowfish-cbc`, `cast128-cbc`, `aes192-cbc`, `aes256-cbc`, `arcfour`, `rijndael-cbc@lysator.liu.se`, `aes128-gcm@openssh.com`, `aes256-gcm@openssh.com`.
     */
    sshEncAlgo?: pulumi.Input<string>;
    /**
     * Config SSH host key.
     */
    sshHsk?: pulumi.Input<string>;
    /**
     * Select one or more SSH hostkey algorithms. Valid values: `ssh-rsa`, `ecdsa-sha2-nistp521`, `ecdsa-sha2-nistp384`, `ecdsa-sha2-nistp256`, `rsa-sha2-256`, `rsa-sha2-512`, `ssh-ed25519`.
     */
    sshHskAlgo?: pulumi.Input<string>;
    /**
     * Enable/disable SSH host key override in SSH daemon. Valid values: `disable`, `enable`.
     */
    sshHskOverride?: pulumi.Input<string>;
    /**
     * Password for ssh-hostkey.
     */
    sshHskPassword?: pulumi.Input<string>;
    /**
     * Select one or more SSH kex algorithms. Valid values: `diffie-hellman-group1-sha1`, `diffie-hellman-group14-sha1`, `diffie-hellman-group14-sha256`, `diffie-hellman-group16-sha512`, `diffie-hellman-group18-sha512`, `diffie-hellman-group-exchange-sha1`, `diffie-hellman-group-exchange-sha256`, `curve25519-sha256@libssh.org`, `ecdh-sha2-nistp256`, `ecdh-sha2-nistp384`, `ecdh-sha2-nistp521`.
     */
    sshKexAlgo?: pulumi.Input<string>;
    /**
     * Select one or more SSH MAC algorithms. Valid values: `hmac-md5`, `hmac-md5-etm@openssh.com`, `hmac-md5-96`, `hmac-md5-96-etm@openssh.com`, `hmac-sha1`, `hmac-sha1-etm@openssh.com`, `hmac-sha2-256`, `hmac-sha2-256-etm@openssh.com`, `hmac-sha2-512`, `hmac-sha2-512-etm@openssh.com`, `hmac-ripemd160`, `hmac-ripemd160@openssh.com`, `hmac-ripemd160-etm@openssh.com`, `umac-64@openssh.com`, `umac-128@openssh.com`, `umac-64-etm@openssh.com`, `umac-128-etm@openssh.com`.
     */
    sshMacAlgo?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Sshconfig resource.
 */
export interface SshconfigArgs {
    /**
     * Select one or more SSH ciphers. Valid values: `chacha20-poly1305@openssh.com`, `aes128-ctr`, `aes192-ctr`, `aes256-ctr`, `arcfour256`, `arcfour128`, `aes128-cbc`, `3des-cbc`, `blowfish-cbc`, `cast128-cbc`, `aes192-cbc`, `aes256-cbc`, `arcfour`, `rijndael-cbc@lysator.liu.se`, `aes128-gcm@openssh.com`, `aes256-gcm@openssh.com`.
     */
    sshEncAlgo?: pulumi.Input<string>;
    /**
     * Config SSH host key.
     */
    sshHsk?: pulumi.Input<string>;
    /**
     * Select one or more SSH hostkey algorithms. Valid values: `ssh-rsa`, `ecdsa-sha2-nistp521`, `ecdsa-sha2-nistp384`, `ecdsa-sha2-nistp256`, `rsa-sha2-256`, `rsa-sha2-512`, `ssh-ed25519`.
     */
    sshHskAlgo?: pulumi.Input<string>;
    /**
     * Enable/disable SSH host key override in SSH daemon. Valid values: `disable`, `enable`.
     */
    sshHskOverride?: pulumi.Input<string>;
    /**
     * Password for ssh-hostkey.
     */
    sshHskPassword?: pulumi.Input<string>;
    /**
     * Select one or more SSH kex algorithms. Valid values: `diffie-hellman-group1-sha1`, `diffie-hellman-group14-sha1`, `diffie-hellman-group14-sha256`, `diffie-hellman-group16-sha512`, `diffie-hellman-group18-sha512`, `diffie-hellman-group-exchange-sha1`, `diffie-hellman-group-exchange-sha256`, `curve25519-sha256@libssh.org`, `ecdh-sha2-nistp256`, `ecdh-sha2-nistp384`, `ecdh-sha2-nistp521`.
     */
    sshKexAlgo?: pulumi.Input<string>;
    /**
     * Select one or more SSH MAC algorithms. Valid values: `hmac-md5`, `hmac-md5-etm@openssh.com`, `hmac-md5-96`, `hmac-md5-96-etm@openssh.com`, `hmac-sha1`, `hmac-sha1-etm@openssh.com`, `hmac-sha2-256`, `hmac-sha2-256-etm@openssh.com`, `hmac-sha2-512`, `hmac-sha2-512-etm@openssh.com`, `hmac-ripemd160`, `hmac-ripemd160@openssh.com`, `hmac-ripemd160-etm@openssh.com`, `umac-64@openssh.com`, `umac-128@openssh.com`, `umac-64-etm@openssh.com`, `umac-128-etm@openssh.com`.
     */
    sshMacAlgo?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
