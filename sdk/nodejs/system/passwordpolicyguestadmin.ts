// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure the password policy for guest administrators.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Passwordpolicyguestadmin("trname", {
 *     applyTo: "guest-admin-password",
 *     change4Characters: "disable",
 *     expireDay: 90,
 *     expireStatus: "disable",
 *     minLowerCaseLetter: 0,
 *     minNonAlphanumeric: 0,
 *     minNumber: 0,
 *     minUpperCaseLetter: 0,
 *     minimumLength: 8,
 *     reusePassword: "enable",
 *     status: "disable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * System PasswordPolicyGuestAdmin can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/passwordpolicyguestadmin:Passwordpolicyguestadmin labelname SystemPasswordPolicyGuestAdmin
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/passwordpolicyguestadmin:Passwordpolicyguestadmin labelname SystemPasswordPolicyGuestAdmin
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Passwordpolicyguestadmin extends pulumi.CustomResource {
    /**
     * Get an existing Passwordpolicyguestadmin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PasswordpolicyguestadminState, opts?: pulumi.CustomResourceOptions): Passwordpolicyguestadmin {
        return new Passwordpolicyguestadmin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/passwordpolicyguestadmin:Passwordpolicyguestadmin';

    /**
     * Returns true if the given object is an instance of Passwordpolicyguestadmin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Passwordpolicyguestadmin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Passwordpolicyguestadmin.__pulumiType;
    }

    /**
     * Guest administrator to which this password policy applies. Valid values: `guest-admin-password`.
     */
    public readonly applyTo!: pulumi.Output<string>;
    /**
     * Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
     */
    public readonly change4Characters!: pulumi.Output<string>;
    /**
     * Number of days after which passwords expire (1 - 999 days, default = 90).
     */
    public readonly expireDay!: pulumi.Output<number>;
    /**
     * Enable/disable password expiration. Valid values: `enable`, `disable`.
     */
    public readonly expireStatus!: pulumi.Output<string>;
    /**
     * Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
     */
    public readonly minChangeCharacters!: pulumi.Output<number>;
    /**
     * Minimum number of lowercase characters in password (0 - 128, default = 0).
     */
    public readonly minLowerCaseLetter!: pulumi.Output<number>;
    /**
     * Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
     */
    public readonly minNonAlphanumeric!: pulumi.Output<number>;
    /**
     * Minimum number of numeric characters in password (0 - 128, default = 0).
     */
    public readonly minNumber!: pulumi.Output<number>;
    /**
     * Minimum number of uppercase characters in password (0 - 128, default = 0).
     */
    public readonly minUpperCaseLetter!: pulumi.Output<number>;
    /**
     * Minimum password length (8 - 128, default = 8).
     */
    public readonly minimumLength!: pulumi.Output<number>;
    /**
     * Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
     */
    public readonly reusePassword!: pulumi.Output<string>;
    /**
     * Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Passwordpolicyguestadmin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PasswordpolicyguestadminArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PasswordpolicyguestadminArgs | PasswordpolicyguestadminState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PasswordpolicyguestadminState | undefined;
            resourceInputs["applyTo"] = state ? state.applyTo : undefined;
            resourceInputs["change4Characters"] = state ? state.change4Characters : undefined;
            resourceInputs["expireDay"] = state ? state.expireDay : undefined;
            resourceInputs["expireStatus"] = state ? state.expireStatus : undefined;
            resourceInputs["minChangeCharacters"] = state ? state.minChangeCharacters : undefined;
            resourceInputs["minLowerCaseLetter"] = state ? state.minLowerCaseLetter : undefined;
            resourceInputs["minNonAlphanumeric"] = state ? state.minNonAlphanumeric : undefined;
            resourceInputs["minNumber"] = state ? state.minNumber : undefined;
            resourceInputs["minUpperCaseLetter"] = state ? state.minUpperCaseLetter : undefined;
            resourceInputs["minimumLength"] = state ? state.minimumLength : undefined;
            resourceInputs["reusePassword"] = state ? state.reusePassword : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as PasswordpolicyguestadminArgs | undefined;
            resourceInputs["applyTo"] = args ? args.applyTo : undefined;
            resourceInputs["change4Characters"] = args ? args.change4Characters : undefined;
            resourceInputs["expireDay"] = args ? args.expireDay : undefined;
            resourceInputs["expireStatus"] = args ? args.expireStatus : undefined;
            resourceInputs["minChangeCharacters"] = args ? args.minChangeCharacters : undefined;
            resourceInputs["minLowerCaseLetter"] = args ? args.minLowerCaseLetter : undefined;
            resourceInputs["minNonAlphanumeric"] = args ? args.minNonAlphanumeric : undefined;
            resourceInputs["minNumber"] = args ? args.minNumber : undefined;
            resourceInputs["minUpperCaseLetter"] = args ? args.minUpperCaseLetter : undefined;
            resourceInputs["minimumLength"] = args ? args.minimumLength : undefined;
            resourceInputs["reusePassword"] = args ? args.reusePassword : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Passwordpolicyguestadmin.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Passwordpolicyguestadmin resources.
 */
export interface PasswordpolicyguestadminState {
    /**
     * Guest administrator to which this password policy applies. Valid values: `guest-admin-password`.
     */
    applyTo?: pulumi.Input<string>;
    /**
     * Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
     */
    change4Characters?: pulumi.Input<string>;
    /**
     * Number of days after which passwords expire (1 - 999 days, default = 90).
     */
    expireDay?: pulumi.Input<number>;
    /**
     * Enable/disable password expiration. Valid values: `enable`, `disable`.
     */
    expireStatus?: pulumi.Input<string>;
    /**
     * Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
     */
    minChangeCharacters?: pulumi.Input<number>;
    /**
     * Minimum number of lowercase characters in password (0 - 128, default = 0).
     */
    minLowerCaseLetter?: pulumi.Input<number>;
    /**
     * Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
     */
    minNonAlphanumeric?: pulumi.Input<number>;
    /**
     * Minimum number of numeric characters in password (0 - 128, default = 0).
     */
    minNumber?: pulumi.Input<number>;
    /**
     * Minimum number of uppercase characters in password (0 - 128, default = 0).
     */
    minUpperCaseLetter?: pulumi.Input<number>;
    /**
     * Minimum password length (8 - 128, default = 8).
     */
    minimumLength?: pulumi.Input<number>;
    /**
     * Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
     */
    reusePassword?: pulumi.Input<string>;
    /**
     * Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Passwordpolicyguestadmin resource.
 */
export interface PasswordpolicyguestadminArgs {
    /**
     * Guest administrator to which this password policy applies. Valid values: `guest-admin-password`.
     */
    applyTo?: pulumi.Input<string>;
    /**
     * Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
     */
    change4Characters?: pulumi.Input<string>;
    /**
     * Number of days after which passwords expire (1 - 999 days, default = 90).
     */
    expireDay?: pulumi.Input<number>;
    /**
     * Enable/disable password expiration. Valid values: `enable`, `disable`.
     */
    expireStatus?: pulumi.Input<string>;
    /**
     * Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
     */
    minChangeCharacters?: pulumi.Input<number>;
    /**
     * Minimum number of lowercase characters in password (0 - 128, default = 0).
     */
    minLowerCaseLetter?: pulumi.Input<number>;
    /**
     * Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
     */
    minNonAlphanumeric?: pulumi.Input<number>;
    /**
     * Minimum number of numeric characters in password (0 - 128, default = 0).
     */
    minNumber?: pulumi.Input<number>;
    /**
     * Minimum number of uppercase characters in password (0 - 128, default = 0).
     */
    minUpperCaseLetter?: pulumi.Input<number>;
    /**
     * Minimum password length (8 - 128, default = 8).
     */
    minimumLength?: pulumi.Input<number>;
    /**
     * Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
     */
    reusePassword?: pulumi.Input<string>;
    /**
     * Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}