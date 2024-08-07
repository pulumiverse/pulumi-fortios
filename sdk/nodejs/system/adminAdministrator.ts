// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to configure administrator accounts of FortiOS.
 *
 * !> **Warning:** The resource will be deprecated and replaced by new resource `fortios.system.Admin`, we recommend that you use the new resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const admintest = new fortios.system.AdminAdministrator("admintest", {
 *     accprofile: "3d3",
 *     comments: "comments",
 *     password: "cc37331AC1",
 *     trusthost1: "1.1.1.0 255.255.255.0",
 *     trusthost2: "2.2.2.0 255.255.255.0",
 *     vdoms: ["root"],
 * });
 * ```
 */
export class AdminAdministrator extends pulumi.CustomResource {
    /**
     * Get an existing AdminAdministrator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdminAdministratorState, opts?: pulumi.CustomResourceOptions): AdminAdministrator {
        return new AdminAdministrator(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/adminAdministrator:AdminAdministrator';

    /**
     * Returns true if the given object is an instance of AdminAdministrator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AdminAdministrator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AdminAdministrator.__pulumiType;
    }

    /**
     * Access profile for this administrator. Access profiles control administrator access to FortiGate features.
     */
    public readonly accprofile!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * User name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Admin user password.
     * * `trusthostN` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.
     */
    public readonly password!: pulumi.Output<string>;
    public readonly trusthost1!: pulumi.Output<string>;
    public readonly trusthost10!: pulumi.Output<string>;
    public readonly trusthost2!: pulumi.Output<string>;
    public readonly trusthost3!: pulumi.Output<string>;
    public readonly trusthost4!: pulumi.Output<string>;
    public readonly trusthost5!: pulumi.Output<string>;
    public readonly trusthost6!: pulumi.Output<string>;
    public readonly trusthost7!: pulumi.Output<string>;
    public readonly trusthost8!: pulumi.Output<string>;
    public readonly trusthost9!: pulumi.Output<string>;
    /**
     * Virtual domain(s) that the administrator can access.
     */
    public readonly vdoms!: pulumi.Output<string[]>;

    /**
     * Create a AdminAdministrator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AdminAdministratorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdminAdministratorArgs | AdminAdministratorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AdminAdministratorState | undefined;
            resourceInputs["accprofile"] = state ? state.accprofile : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["trusthost1"] = state ? state.trusthost1 : undefined;
            resourceInputs["trusthost10"] = state ? state.trusthost10 : undefined;
            resourceInputs["trusthost2"] = state ? state.trusthost2 : undefined;
            resourceInputs["trusthost3"] = state ? state.trusthost3 : undefined;
            resourceInputs["trusthost4"] = state ? state.trusthost4 : undefined;
            resourceInputs["trusthost5"] = state ? state.trusthost5 : undefined;
            resourceInputs["trusthost6"] = state ? state.trusthost6 : undefined;
            resourceInputs["trusthost7"] = state ? state.trusthost7 : undefined;
            resourceInputs["trusthost8"] = state ? state.trusthost8 : undefined;
            resourceInputs["trusthost9"] = state ? state.trusthost9 : undefined;
            resourceInputs["vdoms"] = state ? state.vdoms : undefined;
        } else {
            const args = argsOrState as AdminAdministratorArgs | undefined;
            if ((!args || args.accprofile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accprofile'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            resourceInputs["accprofile"] = args ? args.accprofile : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["trusthost1"] = args ? args.trusthost1 : undefined;
            resourceInputs["trusthost10"] = args ? args.trusthost10 : undefined;
            resourceInputs["trusthost2"] = args ? args.trusthost2 : undefined;
            resourceInputs["trusthost3"] = args ? args.trusthost3 : undefined;
            resourceInputs["trusthost4"] = args ? args.trusthost4 : undefined;
            resourceInputs["trusthost5"] = args ? args.trusthost5 : undefined;
            resourceInputs["trusthost6"] = args ? args.trusthost6 : undefined;
            resourceInputs["trusthost7"] = args ? args.trusthost7 : undefined;
            resourceInputs["trusthost8"] = args ? args.trusthost8 : undefined;
            resourceInputs["trusthost9"] = args ? args.trusthost9 : undefined;
            resourceInputs["vdoms"] = args ? args.vdoms : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AdminAdministrator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AdminAdministrator resources.
 */
export interface AdminAdministratorState {
    /**
     * Access profile for this administrator. Access profiles control administrator access to FortiGate features.
     */
    accprofile?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * User name.
     */
    name?: pulumi.Input<string>;
    /**
     * Admin user password.
     * * `trusthostN` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.
     */
    password?: pulumi.Input<string>;
    trusthost1?: pulumi.Input<string>;
    trusthost10?: pulumi.Input<string>;
    trusthost2?: pulumi.Input<string>;
    trusthost3?: pulumi.Input<string>;
    trusthost4?: pulumi.Input<string>;
    trusthost5?: pulumi.Input<string>;
    trusthost6?: pulumi.Input<string>;
    trusthost7?: pulumi.Input<string>;
    trusthost8?: pulumi.Input<string>;
    trusthost9?: pulumi.Input<string>;
    /**
     * Virtual domain(s) that the administrator can access.
     */
    vdoms?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AdminAdministrator resource.
 */
export interface AdminAdministratorArgs {
    /**
     * Access profile for this administrator. Access profiles control administrator access to FortiGate features.
     */
    accprofile: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * User name.
     */
    name?: pulumi.Input<string>;
    /**
     * Admin user password.
     * * `trusthostN` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.
     */
    password: pulumi.Input<string>;
    trusthost1?: pulumi.Input<string>;
    trusthost10?: pulumi.Input<string>;
    trusthost2?: pulumi.Input<string>;
    trusthost3?: pulumi.Input<string>;
    trusthost4?: pulumi.Input<string>;
    trusthost5?: pulumi.Input<string>;
    trusthost6?: pulumi.Input<string>;
    trusthost7?: pulumi.Input<string>;
    trusthost8?: pulumi.Input<string>;
    trusthost9?: pulumi.Input<string>;
    /**
     * Virtual domain(s) that the administrator can access.
     */
    vdoms?: pulumi.Input<pulumi.Input<string>[]>;
}
