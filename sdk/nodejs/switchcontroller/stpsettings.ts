// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure FortiSwitch spanning tree protocol (STP).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.switchcontroller.Stpsettings("trname", {
 *     forwardTime: 15,
 *     helloTime: 2,
 *     maxAge: 20,
 *     maxHops: 20,
 *     pendingTimer: 4,
 *     revision: 0,
 *     status: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * SwitchController StpSettings can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/stpsettings:Stpsettings labelname SwitchControllerStpSettings
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/stpsettings:Stpsettings labelname SwitchControllerStpSettings
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Stpsettings extends pulumi.CustomResource {
    /**
     * Get an existing Stpsettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StpsettingsState, opts?: pulumi.CustomResourceOptions): Stpsettings {
        return new Stpsettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/stpsettings:Stpsettings';

    /**
     * Returns true if the given object is an instance of Stpsettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stpsettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stpsettings.__pulumiType;
    }

    /**
     * Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
     */
    public readonly forwardTime!: pulumi.Output<number>;
    /**
     * Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
     */
    public readonly helloTime!: pulumi.Output<number>;
    /**
     * Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
     */
    public readonly maxAge!: pulumi.Output<number>;
    /**
     * Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
     */
    public readonly maxHops!: pulumi.Output<number>;
    /**
     * Name of global STP settings configuration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Pending time (1 - 15 sec, default = 4).
     */
    public readonly pendingTimer!: pulumi.Output<number>;
    /**
     * STP revision number (0 - 65535).
     */
    public readonly revision!: pulumi.Output<number>;
    /**
     * Enable/disable STP. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Stpsettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StpsettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StpsettingsArgs | StpsettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StpsettingsState | undefined;
            resourceInputs["forwardTime"] = state ? state.forwardTime : undefined;
            resourceInputs["helloTime"] = state ? state.helloTime : undefined;
            resourceInputs["maxAge"] = state ? state.maxAge : undefined;
            resourceInputs["maxHops"] = state ? state.maxHops : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pendingTimer"] = state ? state.pendingTimer : undefined;
            resourceInputs["revision"] = state ? state.revision : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as StpsettingsArgs | undefined;
            resourceInputs["forwardTime"] = args ? args.forwardTime : undefined;
            resourceInputs["helloTime"] = args ? args.helloTime : undefined;
            resourceInputs["maxAge"] = args ? args.maxAge : undefined;
            resourceInputs["maxHops"] = args ? args.maxHops : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pendingTimer"] = args ? args.pendingTimer : undefined;
            resourceInputs["revision"] = args ? args.revision : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Stpsettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Stpsettings resources.
 */
export interface StpsettingsState {
    /**
     * Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
     */
    forwardTime?: pulumi.Input<number>;
    /**
     * Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
     */
    helloTime?: pulumi.Input<number>;
    /**
     * Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
     */
    maxAge?: pulumi.Input<number>;
    /**
     * Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
     */
    maxHops?: pulumi.Input<number>;
    /**
     * Name of global STP settings configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * Pending time (1 - 15 sec, default = 4).
     */
    pendingTimer?: pulumi.Input<number>;
    /**
     * STP revision number (0 - 65535).
     */
    revision?: pulumi.Input<number>;
    /**
     * Enable/disable STP. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Stpsettings resource.
 */
export interface StpsettingsArgs {
    /**
     * Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
     */
    forwardTime?: pulumi.Input<number>;
    /**
     * Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
     */
    helloTime?: pulumi.Input<number>;
    /**
     * Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
     */
    maxAge?: pulumi.Input<number>;
    /**
     * Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
     */
    maxHops?: pulumi.Input<number>;
    /**
     * Name of global STP settings configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * Pending time (1 - 15 sec, default = 4).
     */
    pendingTimer?: pulumi.Input<number>;
    /**
     * STP revision number (0 - 65535).
     */
    revision?: pulumi.Input<number>;
    /**
     * Enable/disable STP. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
