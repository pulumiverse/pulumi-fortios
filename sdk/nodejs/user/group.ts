// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure user groups.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.user.Group("trname", {
 *     company: "optional",
 *     email: "enable",
 *     expire: 14400,
 *     expireType: "immediately",
 *     groupType: "firewall",
 *     maxAccounts: 0,
 *     members: [{
 *         name: "guest",
 *     }],
 *     mobilePhone: "disable",
 *     multipleGuestAdd: "disable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * User Group can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:user/group:Group labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:user/group:Group labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:user/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * Enable/disable overriding the global number of concurrent authentication sessions for this user group. Valid values: `enable`, `disable`.
     */
    public readonly authConcurrentOverride!: pulumi.Output<string>;
    /**
     * Maximum number of concurrent authenticated connections per user (0 - 100).
     */
    public readonly authConcurrentValue!: pulumi.Output<number>;
    /**
     * Authentication timeout in minutes for this user group. 0 to use the global user setting auth-timeout.
     */
    public readonly authtimeout!: pulumi.Output<number>;
    /**
     * Set the action for the company guest user field. Valid values: `optional`, `mandatory`, `disabled`.
     */
    public readonly company!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable the guest user email address field. Valid values: `disable`, `enable`.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Time in seconds before guest user accounts expire. (1 - 31536000 sec)
     */
    public readonly expire!: pulumi.Output<number>;
    /**
     * Determine when the expiration countdown begins. Valid values: `immediately`, `first-successful-login`.
     */
    public readonly expireType!: pulumi.Output<string>;
    /**
     * Group ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Set the group to be for firewall authentication, FSSO, RSSO, or guest users. Valid values: `firewall`, `fsso-service`, `rsso`, `guest`.
     */
    public readonly groupType!: pulumi.Output<string>;
    /**
     * Guest User. The structure of `guest` block is documented below.
     */
    public readonly guests!: pulumi.Output<outputs.user.GroupGuest[] | undefined>;
    /**
     * Realm attribute for MD5-digest authentication.
     */
    public readonly httpDigestRealm!: pulumi.Output<string>;
    /**
     * Group matches. The structure of `match` block is documented below.
     */
    public readonly matches!: pulumi.Output<outputs.user.GroupMatch[] | undefined>;
    /**
     * Maximum number of guest accounts that can be created for this group (0 means unlimited).
     */
    public readonly maxAccounts!: pulumi.Output<number>;
    /**
     * Names of users, peers, LDAP severs, or RADIUS servers to add to the user group. The structure of `member` block is documented below.
     */
    public readonly members!: pulumi.Output<outputs.user.GroupMember[] | undefined>;
    /**
     * Enable/disable the guest user mobile phone number field. Valid values: `disable`, `enable`.
     */
    public readonly mobilePhone!: pulumi.Output<string>;
    /**
     * Enable/disable addition of multiple guests. Valid values: `disable`, `enable`.
     */
    public readonly multipleGuestAdd!: pulumi.Output<string>;
    /**
     * Group name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Guest user password type. Valid values: `auto-generate`, `specify`, `disable`.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * SMS server.
     */
    public readonly smsCustomServer!: pulumi.Output<string>;
    /**
     * Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
     */
    public readonly smsServer!: pulumi.Output<string>;
    /**
     * Set the action for the sponsor guest user field. Valid values: `optional`, `mandatory`, `disabled`.
     */
    public readonly sponsor!: pulumi.Output<string>;
    /**
     * Name of the RADIUS user group that this local user group represents.
     */
    public readonly ssoAttributeValue!: pulumi.Output<string>;
    /**
     * Guest user ID type. Valid values: `email`, `auto-generate`, `specify`.
     */
    public readonly userId!: pulumi.Output<string>;
    /**
     * Enable/disable the guest user name entry. Valid values: `disable`, `enable`.
     */
    public readonly userName!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupState | undefined;
            resourceInputs["authConcurrentOverride"] = state ? state.authConcurrentOverride : undefined;
            resourceInputs["authConcurrentValue"] = state ? state.authConcurrentValue : undefined;
            resourceInputs["authtimeout"] = state ? state.authtimeout : undefined;
            resourceInputs["company"] = state ? state.company : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["expire"] = state ? state.expire : undefined;
            resourceInputs["expireType"] = state ? state.expireType : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["groupType"] = state ? state.groupType : undefined;
            resourceInputs["guests"] = state ? state.guests : undefined;
            resourceInputs["httpDigestRealm"] = state ? state.httpDigestRealm : undefined;
            resourceInputs["matches"] = state ? state.matches : undefined;
            resourceInputs["maxAccounts"] = state ? state.maxAccounts : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["mobilePhone"] = state ? state.mobilePhone : undefined;
            resourceInputs["multipleGuestAdd"] = state ? state.multipleGuestAdd : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["smsCustomServer"] = state ? state.smsCustomServer : undefined;
            resourceInputs["smsServer"] = state ? state.smsServer : undefined;
            resourceInputs["sponsor"] = state ? state.sponsor : undefined;
            resourceInputs["ssoAttributeValue"] = state ? state.ssoAttributeValue : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            resourceInputs["authConcurrentOverride"] = args ? args.authConcurrentOverride : undefined;
            resourceInputs["authConcurrentValue"] = args ? args.authConcurrentValue : undefined;
            resourceInputs["authtimeout"] = args ? args.authtimeout : undefined;
            resourceInputs["company"] = args ? args.company : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["expire"] = args ? args.expire : undefined;
            resourceInputs["expireType"] = args ? args.expireType : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["groupType"] = args ? args.groupType : undefined;
            resourceInputs["guests"] = args ? args.guests : undefined;
            resourceInputs["httpDigestRealm"] = args ? args.httpDigestRealm : undefined;
            resourceInputs["matches"] = args ? args.matches : undefined;
            resourceInputs["maxAccounts"] = args ? args.maxAccounts : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["mobilePhone"] = args ? args.mobilePhone : undefined;
            resourceInputs["multipleGuestAdd"] = args ? args.multipleGuestAdd : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["smsCustomServer"] = args ? args.smsCustomServer : undefined;
            resourceInputs["smsServer"] = args ? args.smsServer : undefined;
            resourceInputs["sponsor"] = args ? args.sponsor : undefined;
            resourceInputs["ssoAttributeValue"] = args ? args.ssoAttributeValue : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * Enable/disable overriding the global number of concurrent authentication sessions for this user group. Valid values: `enable`, `disable`.
     */
    authConcurrentOverride?: pulumi.Input<string>;
    /**
     * Maximum number of concurrent authenticated connections per user (0 - 100).
     */
    authConcurrentValue?: pulumi.Input<number>;
    /**
     * Authentication timeout in minutes for this user group. 0 to use the global user setting auth-timeout.
     */
    authtimeout?: pulumi.Input<number>;
    /**
     * Set the action for the company guest user field. Valid values: `optional`, `mandatory`, `disabled`.
     */
    company?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable the guest user email address field. Valid values: `disable`, `enable`.
     */
    email?: pulumi.Input<string>;
    /**
     * Time in seconds before guest user accounts expire. (1 - 31536000 sec)
     */
    expire?: pulumi.Input<number>;
    /**
     * Determine when the expiration countdown begins. Valid values: `immediately`, `first-successful-login`.
     */
    expireType?: pulumi.Input<string>;
    /**
     * Group ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Set the group to be for firewall authentication, FSSO, RSSO, or guest users. Valid values: `firewall`, `fsso-service`, `rsso`, `guest`.
     */
    groupType?: pulumi.Input<string>;
    /**
     * Guest User. The structure of `guest` block is documented below.
     */
    guests?: pulumi.Input<pulumi.Input<inputs.user.GroupGuest>[]>;
    /**
     * Realm attribute for MD5-digest authentication.
     */
    httpDigestRealm?: pulumi.Input<string>;
    /**
     * Group matches. The structure of `match` block is documented below.
     */
    matches?: pulumi.Input<pulumi.Input<inputs.user.GroupMatch>[]>;
    /**
     * Maximum number of guest accounts that can be created for this group (0 means unlimited).
     */
    maxAccounts?: pulumi.Input<number>;
    /**
     * Names of users, peers, LDAP severs, or RADIUS servers to add to the user group. The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.user.GroupMember>[]>;
    /**
     * Enable/disable the guest user mobile phone number field. Valid values: `disable`, `enable`.
     */
    mobilePhone?: pulumi.Input<string>;
    /**
     * Enable/disable addition of multiple guests. Valid values: `disable`, `enable`.
     */
    multipleGuestAdd?: pulumi.Input<string>;
    /**
     * Group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Guest user password type. Valid values: `auto-generate`, `specify`, `disable`.
     */
    password?: pulumi.Input<string>;
    /**
     * SMS server.
     */
    smsCustomServer?: pulumi.Input<string>;
    /**
     * Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
     */
    smsServer?: pulumi.Input<string>;
    /**
     * Set the action for the sponsor guest user field. Valid values: `optional`, `mandatory`, `disabled`.
     */
    sponsor?: pulumi.Input<string>;
    /**
     * Name of the RADIUS user group that this local user group represents.
     */
    ssoAttributeValue?: pulumi.Input<string>;
    /**
     * Guest user ID type. Valid values: `email`, `auto-generate`, `specify`.
     */
    userId?: pulumi.Input<string>;
    /**
     * Enable/disable the guest user name entry. Valid values: `disable`, `enable`.
     */
    userName?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * Enable/disable overriding the global number of concurrent authentication sessions for this user group. Valid values: `enable`, `disable`.
     */
    authConcurrentOverride?: pulumi.Input<string>;
    /**
     * Maximum number of concurrent authenticated connections per user (0 - 100).
     */
    authConcurrentValue?: pulumi.Input<number>;
    /**
     * Authentication timeout in minutes for this user group. 0 to use the global user setting auth-timeout.
     */
    authtimeout?: pulumi.Input<number>;
    /**
     * Set the action for the company guest user field. Valid values: `optional`, `mandatory`, `disabled`.
     */
    company?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable the guest user email address field. Valid values: `disable`, `enable`.
     */
    email?: pulumi.Input<string>;
    /**
     * Time in seconds before guest user accounts expire. (1 - 31536000 sec)
     */
    expire?: pulumi.Input<number>;
    /**
     * Determine when the expiration countdown begins. Valid values: `immediately`, `first-successful-login`.
     */
    expireType?: pulumi.Input<string>;
    /**
     * Group ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Set the group to be for firewall authentication, FSSO, RSSO, or guest users. Valid values: `firewall`, `fsso-service`, `rsso`, `guest`.
     */
    groupType?: pulumi.Input<string>;
    /**
     * Guest User. The structure of `guest` block is documented below.
     */
    guests?: pulumi.Input<pulumi.Input<inputs.user.GroupGuest>[]>;
    /**
     * Realm attribute for MD5-digest authentication.
     */
    httpDigestRealm?: pulumi.Input<string>;
    /**
     * Group matches. The structure of `match` block is documented below.
     */
    matches?: pulumi.Input<pulumi.Input<inputs.user.GroupMatch>[]>;
    /**
     * Maximum number of guest accounts that can be created for this group (0 means unlimited).
     */
    maxAccounts?: pulumi.Input<number>;
    /**
     * Names of users, peers, LDAP severs, or RADIUS servers to add to the user group. The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.user.GroupMember>[]>;
    /**
     * Enable/disable the guest user mobile phone number field. Valid values: `disable`, `enable`.
     */
    mobilePhone?: pulumi.Input<string>;
    /**
     * Enable/disable addition of multiple guests. Valid values: `disable`, `enable`.
     */
    multipleGuestAdd?: pulumi.Input<string>;
    /**
     * Group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Guest user password type. Valid values: `auto-generate`, `specify`, `disable`.
     */
    password?: pulumi.Input<string>;
    /**
     * SMS server.
     */
    smsCustomServer?: pulumi.Input<string>;
    /**
     * Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
     */
    smsServer?: pulumi.Input<string>;
    /**
     * Set the action for the sponsor guest user field. Valid values: `optional`, `mandatory`, `disabled`.
     */
    sponsor?: pulumi.Input<string>;
    /**
     * Name of the RADIUS user group that this local user group represents.
     */
    ssoAttributeValue?: pulumi.Input<string>;
    /**
     * Guest user ID type. Valid values: `email`, `auto-generate`, `specify`.
     */
    userId?: pulumi.Input<string>;
    /**
     * Enable/disable the guest user name entry. Valid values: `disable`, `enable`.
     */
    userName?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
