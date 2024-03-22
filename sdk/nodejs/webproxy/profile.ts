// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure web proxy profiles.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.webproxy.Profile("trname", {
 *     headerClientIp: "pass",
 *     headerFrontEndHttps: "pass",
 *     headerViaRequest: "add",
 *     headerViaResponse: "pass",
 *     headerXAuthenticatedGroups: "pass",
 *     headerXAuthenticatedUser: "pass",
 *     headerXForwardedFor: "pass",
 *     logHeaderChange: "disable",
 *     stripEncoding: "disable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * WebProxy Profile can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:webproxy/profile:Profile labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:webproxy/profile:Profile labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Profile extends pulumi.CustomResource {
    /**
     * Get an existing Profile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileState, opts?: pulumi.CustomResourceOptions): Profile {
        return new Profile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:webproxy/profile:Profile';

    /**
     * Returns true if the given object is an instance of Profile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Profile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Profile.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    public readonly headerClientIp!: pulumi.Output<string>;
    /**
     * Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    public readonly headerFrontEndHttps!: pulumi.Output<string>;
    /**
     * Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    public readonly headerViaRequest!: pulumi.Output<string>;
    /**
     * Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    public readonly headerViaResponse!: pulumi.Output<string>;
    /**
     * Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    public readonly headerXAuthenticatedGroups!: pulumi.Output<string>;
    /**
     * Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    public readonly headerXAuthenticatedUser!: pulumi.Output<string>;
    /**
     * Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    public readonly headerXForwardedClientCert!: pulumi.Output<string>;
    /**
     * Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    public readonly headerXForwardedFor!: pulumi.Output<string>;
    /**
     * Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
     */
    public readonly headers!: pulumi.Output<outputs.webproxy.ProfileHeader[] | undefined>;
    /**
     * Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
     */
    public readonly logHeaderChange!: pulumi.Output<string>;
    /**
     * Profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
     */
    public readonly stripEncoding!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Profile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileArgs | ProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProfileState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["headerClientIp"] = state ? state.headerClientIp : undefined;
            resourceInputs["headerFrontEndHttps"] = state ? state.headerFrontEndHttps : undefined;
            resourceInputs["headerViaRequest"] = state ? state.headerViaRequest : undefined;
            resourceInputs["headerViaResponse"] = state ? state.headerViaResponse : undefined;
            resourceInputs["headerXAuthenticatedGroups"] = state ? state.headerXAuthenticatedGroups : undefined;
            resourceInputs["headerXAuthenticatedUser"] = state ? state.headerXAuthenticatedUser : undefined;
            resourceInputs["headerXForwardedClientCert"] = state ? state.headerXForwardedClientCert : undefined;
            resourceInputs["headerXForwardedFor"] = state ? state.headerXForwardedFor : undefined;
            resourceInputs["headers"] = state ? state.headers : undefined;
            resourceInputs["logHeaderChange"] = state ? state.logHeaderChange : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["stripEncoding"] = state ? state.stripEncoding : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ProfileArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["headerClientIp"] = args ? args.headerClientIp : undefined;
            resourceInputs["headerFrontEndHttps"] = args ? args.headerFrontEndHttps : undefined;
            resourceInputs["headerViaRequest"] = args ? args.headerViaRequest : undefined;
            resourceInputs["headerViaResponse"] = args ? args.headerViaResponse : undefined;
            resourceInputs["headerXAuthenticatedGroups"] = args ? args.headerXAuthenticatedGroups : undefined;
            resourceInputs["headerXAuthenticatedUser"] = args ? args.headerXAuthenticatedUser : undefined;
            resourceInputs["headerXForwardedClientCert"] = args ? args.headerXForwardedClientCert : undefined;
            resourceInputs["headerXForwardedFor"] = args ? args.headerXForwardedFor : undefined;
            resourceInputs["headers"] = args ? args.headers : undefined;
            resourceInputs["logHeaderChange"] = args ? args.logHeaderChange : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["stripEncoding"] = args ? args.stripEncoding : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Profile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Profile resources.
 */
export interface ProfileState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerClientIp?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerFrontEndHttps?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerViaRequest?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerViaResponse?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerXAuthenticatedGroups?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerXAuthenticatedUser?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerXForwardedClientCert?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerXForwardedFor?: pulumi.Input<string>;
    /**
     * Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
     */
    headers?: pulumi.Input<pulumi.Input<inputs.webproxy.ProfileHeader>[]>;
    /**
     * Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
     */
    logHeaderChange?: pulumi.Input<string>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
     */
    stripEncoding?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Profile resource.
 */
export interface ProfileArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerClientIp?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerFrontEndHttps?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerViaRequest?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerViaResponse?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerXAuthenticatedGroups?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerXAuthenticatedUser?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerXForwardedClientCert?: pulumi.Input<string>;
    /**
     * Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
     */
    headerXForwardedFor?: pulumi.Input<string>;
    /**
     * Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
     */
    headers?: pulumi.Input<pulumi.Input<inputs.webproxy.ProfileHeader>[]>;
    /**
     * Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
     */
    logHeaderChange?: pulumi.Input<string>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
     */
    stripEncoding?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
