// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Configure web filter search engines.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.filter.web.Searchengine("trname", {
 *     charset: "utf-8",
 *     hostname: "sg.eiwuc.com",
 *     query: "sc=",
 *     safesearch: "disable",
 *     url: "^\\/f",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Webfilter SearchEngine can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:filter/web/searchengine:Searchengine labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:filter/web/searchengine:Searchengine labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Searchengine extends pulumi.CustomResource {
    /**
     * Get an existing Searchengine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SearchengineState, opts?: pulumi.CustomResourceOptions): Searchengine {
        return new Searchengine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:filter/web/searchengine:Searchengine';

    /**
     * Returns true if the given object is an instance of Searchengine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Searchengine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Searchengine.__pulumiType;
    }

    /**
     * Search engine charset. Valid values: `utf-8`, `gb2312`.
     */
    public readonly charset!: pulumi.Output<string>;
    /**
     * Hostname (regular expression).
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * Search engine name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Code used to prefix a query (must end with an equals character).
     */
    public readonly query!: pulumi.Output<string>;
    /**
     * Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
     */
    public readonly safesearch!: pulumi.Output<string>;
    /**
     * Safe search parameter used in the URL.
     */
    public readonly safesearchStr!: pulumi.Output<string>;
    /**
     * URL (regular expression).
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Searchengine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SearchengineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SearchengineArgs | SearchengineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SearchengineState | undefined;
            resourceInputs["charset"] = state ? state.charset : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["query"] = state ? state.query : undefined;
            resourceInputs["safesearch"] = state ? state.safesearch : undefined;
            resourceInputs["safesearchStr"] = state ? state.safesearchStr : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SearchengineArgs | undefined;
            resourceInputs["charset"] = args ? args.charset : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["query"] = args ? args.query : undefined;
            resourceInputs["safesearch"] = args ? args.safesearch : undefined;
            resourceInputs["safesearchStr"] = args ? args.safesearchStr : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Searchengine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Searchengine resources.
 */
export interface SearchengineState {
    /**
     * Search engine charset. Valid values: `utf-8`, `gb2312`.
     */
    charset?: pulumi.Input<string>;
    /**
     * Hostname (regular expression).
     */
    hostname?: pulumi.Input<string>;
    /**
     * Search engine name.
     */
    name?: pulumi.Input<string>;
    /**
     * Code used to prefix a query (must end with an equals character).
     */
    query?: pulumi.Input<string>;
    /**
     * Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
     */
    safesearch?: pulumi.Input<string>;
    /**
     * Safe search parameter used in the URL.
     */
    safesearchStr?: pulumi.Input<string>;
    /**
     * URL (regular expression).
     */
    url?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Searchengine resource.
 */
export interface SearchengineArgs {
    /**
     * Search engine charset. Valid values: `utf-8`, `gb2312`.
     */
    charset?: pulumi.Input<string>;
    /**
     * Hostname (regular expression).
     */
    hostname?: pulumi.Input<string>;
    /**
     * Search engine name.
     */
    name?: pulumi.Input<string>;
    /**
     * Code used to prefix a query (must end with an equals character).
     */
    query?: pulumi.Input<string>;
    /**
     * Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
     */
    safesearch?: pulumi.Input<string>;
    /**
     * Safe search parameter used in the URL.
     */
    safesearchStr?: pulumi.Input<string>;
    /**
     * URL (regular expression).
     */
    url?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
