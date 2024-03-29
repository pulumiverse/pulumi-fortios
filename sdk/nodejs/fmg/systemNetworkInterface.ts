// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource supports updating system network interface for FortiManager.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.fmg.SystemNetworkInterface("test1", {
 *     allowAccesses: [
 *         "ping",
 *         "ssh",
 *         "https",
 *         "http",
 *     ],
 *     description: "",
 *     ip: "1.1.1.3 255.255.255.0",
 *     serviceAccesses: [
 *         "webfilter",
 *         "fgtupdates",
 *     ],
 *     status: "up",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class SystemNetworkInterface extends pulumi.CustomResource {
    /**
     * Get an existing SystemNetworkInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemNetworkInterfaceState, opts?: pulumi.CustomResourceOptions): SystemNetworkInterface {
        return new SystemNetworkInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:fmg/systemNetworkInterface:SystemNetworkInterface';

    /**
     * Returns true if the given object is an instance of SystemNetworkInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemNetworkInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemNetworkInterface.__pulumiType;
    }

    /**
     * Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
     */
    public readonly allowAccesses!: pulumi.Output<string[] | undefined>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Interface Ipaddress.
     */
    public readonly ip!: pulumi.Output<string | undefined>;
    /**
     * Interface port.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly serviceAccesses!: pulumi.Output<string[] | undefined>;
    /**
     * Interface status, Enum: ["down", "up"]
     */
    public readonly status!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemNetworkInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemNetworkInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemNetworkInterfaceArgs | SystemNetworkInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemNetworkInterfaceState | undefined;
            resourceInputs["allowAccesses"] = state ? state.allowAccesses : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["serviceAccesses"] = state ? state.serviceAccesses : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as SystemNetworkInterfaceArgs | undefined;
            resourceInputs["allowAccesses"] = args ? args.allowAccesses : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serviceAccesses"] = args ? args.serviceAccesses : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemNetworkInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemNetworkInterface resources.
 */
export interface SystemNetworkInterfaceState {
    /**
     * Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
     */
    allowAccesses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Interface Ipaddress.
     */
    ip?: pulumi.Input<string>;
    /**
     * Interface port.
     */
    name?: pulumi.Input<string>;
    serviceAccesses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Interface status, Enum: ["down", "up"]
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemNetworkInterface resource.
 */
export interface SystemNetworkInterfaceArgs {
    /**
     * Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
     */
    allowAccesses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Interface Ipaddress.
     */
    ip?: pulumi.Input<string>;
    /**
     * Interface port.
     */
    name?: pulumi.Input<string>;
    serviceAccesses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Interface status, Enum: ["down", "up"]
     */
    status?: pulumi.Input<string>;
}
