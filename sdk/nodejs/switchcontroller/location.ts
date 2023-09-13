// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure FortiSwitch location services. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * SwitchController Location can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:switchcontroller/location:Location labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:switchcontroller/location:Location labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Location extends pulumi.CustomResource {
    /**
     * Get an existing Location resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocationState, opts?: pulumi.CustomResourceOptions): Location {
        return new Location(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/location:Location';

    /**
     * Returns true if the given object is an instance of Location.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Location {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Location.__pulumiType;
    }

    /**
     * Configure location civic address. The structure of `addressCivic` block is documented below.
     */
    public readonly addressCivic!: pulumi.Output<outputs.switchcontroller.LocationAddressCivic>;
    /**
     * Configure location GPS coordinates. The structure of `coordinates` block is documented below.
     */
    public readonly coordinates!: pulumi.Output<outputs.switchcontroller.LocationCoordinates>;
    /**
     * Configure location ELIN number. The structure of `elinNumber` block is documented below.
     */
    public readonly elinNumber!: pulumi.Output<outputs.switchcontroller.LocationElinNumber>;
    /**
     * Unique location item name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Location resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LocationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocationArgs | LocationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LocationState | undefined;
            resourceInputs["addressCivic"] = state ? state.addressCivic : undefined;
            resourceInputs["coordinates"] = state ? state.coordinates : undefined;
            resourceInputs["elinNumber"] = state ? state.elinNumber : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LocationArgs | undefined;
            resourceInputs["addressCivic"] = args ? args.addressCivic : undefined;
            resourceInputs["coordinates"] = args ? args.coordinates : undefined;
            resourceInputs["elinNumber"] = args ? args.elinNumber : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Location.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Location resources.
 */
export interface LocationState {
    /**
     * Configure location civic address. The structure of `addressCivic` block is documented below.
     */
    addressCivic?: pulumi.Input<inputs.switchcontroller.LocationAddressCivic>;
    /**
     * Configure location GPS coordinates. The structure of `coordinates` block is documented below.
     */
    coordinates?: pulumi.Input<inputs.switchcontroller.LocationCoordinates>;
    /**
     * Configure location ELIN number. The structure of `elinNumber` block is documented below.
     */
    elinNumber?: pulumi.Input<inputs.switchcontroller.LocationElinNumber>;
    /**
     * Unique location item name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Location resource.
 */
export interface LocationArgs {
    /**
     * Configure location civic address. The structure of `addressCivic` block is documented below.
     */
    addressCivic?: pulumi.Input<inputs.switchcontroller.LocationAddressCivic>;
    /**
     * Configure location GPS coordinates. The structure of `coordinates` block is documented below.
     */
    coordinates?: pulumi.Input<inputs.switchcontroller.LocationCoordinates>;
    /**
     * Configure location ELIN number. The structure of `elinNumber` block is documented below.
     */
    elinNumber?: pulumi.Input<inputs.switchcontroller.LocationElinNumber>;
    /**
     * Unique location item name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
