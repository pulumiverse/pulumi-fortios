// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the fortios package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'fortios';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * CA Bundle file content
     */
    public readonly cabundlecontent!: pulumi.Output<string | undefined>;
    /**
     * CA Bundle file
     */
    public readonly cabundlefile!: pulumi.Output<string | undefined>;
    /**
     * CA certtificate(Optional)
     */
    public readonly cacert!: pulumi.Output<string | undefined>;
    /**
     * User certificate
     */
    public readonly clientcert!: pulumi.Output<string | undefined>;
    /**
     * User private key
     */
    public readonly clientkey!: pulumi.Output<string | undefined>;
    /**
     * CA Bundle file
     */
    public readonly fmgCabundlefile!: pulumi.Output<string | undefined>;
    /**
     * Hostname/IP address of the FortiManager to connect to
     */
    public readonly fmgHostname!: pulumi.Output<string | undefined>;
    public readonly fmgPasswd!: pulumi.Output<string | undefined>;
    public readonly fmgUsername!: pulumi.Output<string | undefined>;
    /**
     * The hostname/IP address of the FortiOS to be connected
     */
    public readonly hostname!: pulumi.Output<string | undefined>;
    /**
     * HTTP proxy address
     */
    public readonly httpProxy!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable peer authentication, can be 'enable' or 'disable'
     */
    public readonly peerauth!: pulumi.Output<string | undefined>;
    public readonly token!: pulumi.Output<string | undefined>;
    public readonly vdom!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["cabundlecontent"] = (args ? args.cabundlecontent : undefined) ?? utilities.getEnv("FORTIOS_CA_CABUNDLECONTENT");
            resourceInputs["cabundlefile"] = (args ? args.cabundlefile : undefined) ?? utilities.getEnv("FORTIOS_CA_CABUNDLE");
            resourceInputs["cacert"] = (args ? args.cacert : undefined) ?? utilities.getEnv("FORTIOS_CA_CACERT");
            resourceInputs["clientcert"] = (args ? args.clientcert : undefined) ?? utilities.getEnv("FORTIOS_CA_CLIENTCERT");
            resourceInputs["clientkey"] = (args?.clientkey ? pulumi.secret(args.clientkey) : undefined) ?? utilities.getEnv("FORTIOS_CA_CLIENTKEY");
            resourceInputs["fmgCabundlefile"] = (args ? args.fmgCabundlefile : undefined) ?? utilities.getEnv("FORTIOS_FMG_CABUNDLE");
            resourceInputs["fmgHostname"] = (args ? args.fmgHostname : undefined) ?? utilities.getEnv("FORTIOS_FMG_HOSTNAME");
            resourceInputs["fmgInsecure"] = pulumi.output((args ? args.fmgInsecure : undefined) ?? utilities.getEnvBoolean("FORTIOS_FMG_INSECURE")).apply(JSON.stringify);
            resourceInputs["fmgPasswd"] = (args?.fmgPasswd ? pulumi.secret(args.fmgPasswd) : undefined) ?? utilities.getEnv("FORTIOS_FMG_PASSWORD");
            resourceInputs["fmgUsername"] = (args ? args.fmgUsername : undefined) ?? utilities.getEnv("FORTIOS_FMG_USERNAME");
            resourceInputs["hostname"] = (args ? args.hostname : undefined) ?? utilities.getEnv("FORTIOS_ACCESS_HOSTNAME");
            resourceInputs["httpProxy"] = (args ? args.httpProxy : undefined) ?? utilities.getEnv("FORTIOS_HTTP_PROXY");
            resourceInputs["insecure"] = pulumi.output((args ? args.insecure : undefined) ?? utilities.getEnvBoolean("FORTIOS_INSECURE")).apply(JSON.stringify);
            resourceInputs["peerauth"] = (args ? args.peerauth : undefined) ?? utilities.getEnv("FORTIOS_CA_PEERAUTH");
            resourceInputs["token"] = (args?.token ? pulumi.secret(args.token) : undefined) ?? utilities.getEnv("FORTIOS_ACCESS_TOKEN");
            resourceInputs["vdom"] = (args ? args.vdom : undefined) ?? utilities.getEnv("FORTIOS_VDOM");
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientkey", "fmgPasswd", "token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * CA Bundle file content
     */
    cabundlecontent?: pulumi.Input<string>;
    /**
     * CA Bundle file
     */
    cabundlefile?: pulumi.Input<string>;
    /**
     * CA certtificate(Optional)
     */
    cacert?: pulumi.Input<string>;
    /**
     * User certificate
     */
    clientcert?: pulumi.Input<string>;
    /**
     * User private key
     */
    clientkey?: pulumi.Input<string>;
    /**
     * CA Bundle file
     */
    fmgCabundlefile?: pulumi.Input<string>;
    /**
     * Hostname/IP address of the FortiManager to connect to
     */
    fmgHostname?: pulumi.Input<string>;
    fmgInsecure?: pulumi.Input<boolean>;
    fmgPasswd?: pulumi.Input<string>;
    fmgUsername?: pulumi.Input<string>;
    /**
     * The hostname/IP address of the FortiOS to be connected
     */
    hostname?: pulumi.Input<string>;
    /**
     * HTTP proxy address
     */
    httpProxy?: pulumi.Input<string>;
    insecure?: pulumi.Input<boolean>;
    /**
     * Enable/disable peer authentication, can be 'enable' or 'disable'
     */
    peerauth?: pulumi.Input<string>;
    token?: pulumi.Input<string>;
    vdom?: pulumi.Input<string>;
}
