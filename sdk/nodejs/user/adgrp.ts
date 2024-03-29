// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure FSSO groups.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname1 = new fortios.user.Fsso("trname1", {
 *     port: 32381,
 *     port2: 8000,
 *     port3: 8000,
 *     port4: 8000,
 *     port5: 8000,
 *     server: "1.1.1.1",
 *     sourceIp: "0.0.0.0",
 *     sourceIp6: "::",
 * });
 * const trname = new fortios.user.Adgrp("trname", {serverName: trname1.name});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * User Adgrp can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:user/adgrp:Adgrp labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:user/adgrp:Adgrp labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Adgrp extends pulumi.CustomResource {
    /**
     * Get an existing Adgrp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdgrpState, opts?: pulumi.CustomResourceOptions): Adgrp {
        return new Adgrp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:user/adgrp:Adgrp';

    /**
     * Returns true if the given object is an instance of Adgrp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Adgrp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Adgrp.__pulumiType;
    }

    /**
     * FSSO connector source.
     */
    public readonly connectorSource!: pulumi.Output<string>;
    /**
     * Group ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * FSSO agent name.
     */
    public readonly serverName!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Adgrp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AdgrpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdgrpArgs | AdgrpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AdgrpState | undefined;
            resourceInputs["connectorSource"] = state ? state.connectorSource : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["serverName"] = state ? state.serverName : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as AdgrpArgs | undefined;
            resourceInputs["connectorSource"] = args ? args.connectorSource : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serverName"] = args ? args.serverName : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Adgrp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Adgrp resources.
 */
export interface AdgrpState {
    /**
     * FSSO connector source.
     */
    connectorSource?: pulumi.Input<string>;
    /**
     * Group ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * FSSO agent name.
     */
    serverName?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Adgrp resource.
 */
export interface AdgrpArgs {
    /**
     * FSSO connector source.
     */
    connectorSource?: pulumi.Input<string>;
    /**
     * Group ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * FSSO agent name.
     */
    serverName?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
