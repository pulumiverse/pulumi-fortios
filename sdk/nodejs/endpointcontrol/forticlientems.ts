// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure FortiClient Enterprise Management Server (EMS) entries. Applies to FortiOS Version `<= 6.2.0`.
 *
 * ## Import
 *
 * EndpointControl ForticlientEms can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:endpointcontrol/forticlientems:Forticlientems labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:endpointcontrol/forticlientems:Forticlientems labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Forticlientems extends pulumi.CustomResource {
    /**
     * Get an existing Forticlientems resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ForticlientemsState, opts?: pulumi.CustomResourceOptions): Forticlientems {
        return new Forticlientems(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:endpointcontrol/forticlientems:Forticlientems';

    /**
     * Returns true if the given object is an instance of Forticlientems.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Forticlientems {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Forticlientems.__pulumiType;
    }

    /**
     * Firewall address name.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * FortiClient EMS admin password.
     */
    public readonly adminPassword!: pulumi.Output<string | undefined>;
    /**
     * FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
     */
    public readonly adminType!: pulumi.Output<string>;
    /**
     * FortiClient EMS admin username.
     */
    public readonly adminUsername!: pulumi.Output<string>;
    /**
     * FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
     */
    public readonly httpsPort!: pulumi.Output<number>;
    /**
     * FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
     */
    public readonly listenPort!: pulumi.Output<number>;
    /**
     * FortiClient Enterprise Management Server (EMS) name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
     */
    public readonly restApiAuth!: pulumi.Output<string>;
    /**
     * FortiClient EMS Serial Number.
     */
    public readonly serialNumber!: pulumi.Output<string>;
    /**
     * FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
     */
    public readonly uploadPort!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Forticlientems resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ForticlientemsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ForticlientemsArgs | ForticlientemsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ForticlientemsState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["adminPassword"] = state ? state.adminPassword : undefined;
            resourceInputs["adminType"] = state ? state.adminType : undefined;
            resourceInputs["adminUsername"] = state ? state.adminUsername : undefined;
            resourceInputs["httpsPort"] = state ? state.httpsPort : undefined;
            resourceInputs["listenPort"] = state ? state.listenPort : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["restApiAuth"] = state ? state.restApiAuth : undefined;
            resourceInputs["serialNumber"] = state ? state.serialNumber : undefined;
            resourceInputs["uploadPort"] = state ? state.uploadPort : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ForticlientemsArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.adminUsername === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminUsername'");
            }
            if ((!args || args.serialNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serialNumber'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["adminPassword"] = args?.adminPassword ? pulumi.secret(args.adminPassword) : undefined;
            resourceInputs["adminType"] = args ? args.adminType : undefined;
            resourceInputs["adminUsername"] = args ? args.adminUsername : undefined;
            resourceInputs["httpsPort"] = args ? args.httpsPort : undefined;
            resourceInputs["listenPort"] = args ? args.listenPort : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["restApiAuth"] = args ? args.restApiAuth : undefined;
            resourceInputs["serialNumber"] = args ? args.serialNumber : undefined;
            resourceInputs["uploadPort"] = args ? args.uploadPort : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["adminPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Forticlientems.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Forticlientems resources.
 */
export interface ForticlientemsState {
    /**
     * Firewall address name.
     */
    address?: pulumi.Input<string>;
    /**
     * FortiClient EMS admin password.
     */
    adminPassword?: pulumi.Input<string>;
    /**
     * FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
     */
    adminType?: pulumi.Input<string>;
    /**
     * FortiClient EMS admin username.
     */
    adminUsername?: pulumi.Input<string>;
    /**
     * FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
     */
    httpsPort?: pulumi.Input<number>;
    /**
     * FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
     */
    listenPort?: pulumi.Input<number>;
    /**
     * FortiClient Enterprise Management Server (EMS) name.
     */
    name?: pulumi.Input<string>;
    /**
     * FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
     */
    restApiAuth?: pulumi.Input<string>;
    /**
     * FortiClient EMS Serial Number.
     */
    serialNumber?: pulumi.Input<string>;
    /**
     * FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
     */
    uploadPort?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Forticlientems resource.
 */
export interface ForticlientemsArgs {
    /**
     * Firewall address name.
     */
    address: pulumi.Input<string>;
    /**
     * FortiClient EMS admin password.
     */
    adminPassword?: pulumi.Input<string>;
    /**
     * FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
     */
    adminType?: pulumi.Input<string>;
    /**
     * FortiClient EMS admin username.
     */
    adminUsername: pulumi.Input<string>;
    /**
     * FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
     */
    httpsPort?: pulumi.Input<number>;
    /**
     * FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
     */
    listenPort?: pulumi.Input<number>;
    /**
     * FortiClient Enterprise Management Server (EMS) name.
     */
    name?: pulumi.Input<string>;
    /**
     * FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
     */
    restApiAuth?: pulumi.Input<string>;
    /**
     * FortiClient EMS Serial Number.
     */
    serialNumber: pulumi.Input<string>;
    /**
     * FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
     */
    uploadPort?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
