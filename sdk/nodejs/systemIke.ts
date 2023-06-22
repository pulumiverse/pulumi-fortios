// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure IKE global attributes. Applies to FortiOS Version `>= 7.0.1`.
 *
 * ## Import
 *
 * System Ike can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemIke:SystemIke labelname SystemIke
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemIke:SystemIke labelname SystemIke
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemIke extends pulumi.CustomResource {
    /**
     * Get an existing SystemIke resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemIkeState, opts?: pulumi.CustomResourceOptions): SystemIke {
        return new SystemIke(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemIke:SystemIke';

    /**
     * Returns true if the given object is an instance of SystemIke.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemIke {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemIke.__pulumiType;
    }

    /**
     * Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
     */
    public readonly dhGroup1!: pulumi.Output<outputs.SystemIkeDhGroup1>;
    /**
     * Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
     */
    public readonly dhGroup14!: pulumi.Output<outputs.SystemIkeDhGroup14>;
    /**
     * Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
     */
    public readonly dhGroup15!: pulumi.Output<outputs.SystemIkeDhGroup15>;
    /**
     * Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
     */
    public readonly dhGroup16!: pulumi.Output<outputs.SystemIkeDhGroup16>;
    /**
     * Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
     */
    public readonly dhGroup17!: pulumi.Output<outputs.SystemIkeDhGroup17>;
    /**
     * Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
     */
    public readonly dhGroup18!: pulumi.Output<outputs.SystemIkeDhGroup18>;
    /**
     * Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
     */
    public readonly dhGroup19!: pulumi.Output<outputs.SystemIkeDhGroup19>;
    /**
     * Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
     */
    public readonly dhGroup2!: pulumi.Output<outputs.SystemIkeDhGroup2>;
    /**
     * Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
     */
    public readonly dhGroup20!: pulumi.Output<outputs.SystemIkeDhGroup20>;
    /**
     * Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
     */
    public readonly dhGroup21!: pulumi.Output<outputs.SystemIkeDhGroup21>;
    /**
     * Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
     */
    public readonly dhGroup27!: pulumi.Output<outputs.SystemIkeDhGroup27>;
    /**
     * Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
     */
    public readonly dhGroup28!: pulumi.Output<outputs.SystemIkeDhGroup28>;
    /**
     * Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
     */
    public readonly dhGroup29!: pulumi.Output<outputs.SystemIkeDhGroup29>;
    /**
     * Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
     */
    public readonly dhGroup30!: pulumi.Output<outputs.SystemIkeDhGroup30>;
    /**
     * Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
     */
    public readonly dhGroup31!: pulumi.Output<outputs.SystemIkeDhGroup31>;
    /**
     * Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
     */
    public readonly dhGroup32!: pulumi.Output<outputs.SystemIkeDhGroup32>;
    /**
     * Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
     */
    public readonly dhGroup5!: pulumi.Output<outputs.SystemIkeDhGroup5>;
    /**
     * Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
     */
    public readonly dhKeypairCache!: pulumi.Output<string>;
    /**
     * Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
     */
    public readonly dhKeypairCount!: pulumi.Output<number>;
    /**
     * Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
     */
    public readonly dhKeypairThrottle!: pulumi.Output<string>;
    /**
     * Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
     */
    public readonly dhMode!: pulumi.Output<string>;
    /**
     * Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
     */
    public readonly dhMultiprocess!: pulumi.Output<string>;
    /**
     * Number of Diffie-Hellman workers to start.
     */
    public readonly dhWorkerCount!: pulumi.Output<number>;
    /**
     * Maximum number of IPsec tunnels to negotiate simultaneously.
     */
    public readonly embryonicLimit!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     *
     * The `dhGroup1` block supports:
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemIke resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemIkeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemIkeArgs | SystemIkeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemIkeState | undefined;
            resourceInputs["dhGroup1"] = state ? state.dhGroup1 : undefined;
            resourceInputs["dhGroup14"] = state ? state.dhGroup14 : undefined;
            resourceInputs["dhGroup15"] = state ? state.dhGroup15 : undefined;
            resourceInputs["dhGroup16"] = state ? state.dhGroup16 : undefined;
            resourceInputs["dhGroup17"] = state ? state.dhGroup17 : undefined;
            resourceInputs["dhGroup18"] = state ? state.dhGroup18 : undefined;
            resourceInputs["dhGroup19"] = state ? state.dhGroup19 : undefined;
            resourceInputs["dhGroup2"] = state ? state.dhGroup2 : undefined;
            resourceInputs["dhGroup20"] = state ? state.dhGroup20 : undefined;
            resourceInputs["dhGroup21"] = state ? state.dhGroup21 : undefined;
            resourceInputs["dhGroup27"] = state ? state.dhGroup27 : undefined;
            resourceInputs["dhGroup28"] = state ? state.dhGroup28 : undefined;
            resourceInputs["dhGroup29"] = state ? state.dhGroup29 : undefined;
            resourceInputs["dhGroup30"] = state ? state.dhGroup30 : undefined;
            resourceInputs["dhGroup31"] = state ? state.dhGroup31 : undefined;
            resourceInputs["dhGroup32"] = state ? state.dhGroup32 : undefined;
            resourceInputs["dhGroup5"] = state ? state.dhGroup5 : undefined;
            resourceInputs["dhKeypairCache"] = state ? state.dhKeypairCache : undefined;
            resourceInputs["dhKeypairCount"] = state ? state.dhKeypairCount : undefined;
            resourceInputs["dhKeypairThrottle"] = state ? state.dhKeypairThrottle : undefined;
            resourceInputs["dhMode"] = state ? state.dhMode : undefined;
            resourceInputs["dhMultiprocess"] = state ? state.dhMultiprocess : undefined;
            resourceInputs["dhWorkerCount"] = state ? state.dhWorkerCount : undefined;
            resourceInputs["embryonicLimit"] = state ? state.embryonicLimit : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemIkeArgs | undefined;
            resourceInputs["dhGroup1"] = args ? args.dhGroup1 : undefined;
            resourceInputs["dhGroup14"] = args ? args.dhGroup14 : undefined;
            resourceInputs["dhGroup15"] = args ? args.dhGroup15 : undefined;
            resourceInputs["dhGroup16"] = args ? args.dhGroup16 : undefined;
            resourceInputs["dhGroup17"] = args ? args.dhGroup17 : undefined;
            resourceInputs["dhGroup18"] = args ? args.dhGroup18 : undefined;
            resourceInputs["dhGroup19"] = args ? args.dhGroup19 : undefined;
            resourceInputs["dhGroup2"] = args ? args.dhGroup2 : undefined;
            resourceInputs["dhGroup20"] = args ? args.dhGroup20 : undefined;
            resourceInputs["dhGroup21"] = args ? args.dhGroup21 : undefined;
            resourceInputs["dhGroup27"] = args ? args.dhGroup27 : undefined;
            resourceInputs["dhGroup28"] = args ? args.dhGroup28 : undefined;
            resourceInputs["dhGroup29"] = args ? args.dhGroup29 : undefined;
            resourceInputs["dhGroup30"] = args ? args.dhGroup30 : undefined;
            resourceInputs["dhGroup31"] = args ? args.dhGroup31 : undefined;
            resourceInputs["dhGroup32"] = args ? args.dhGroup32 : undefined;
            resourceInputs["dhGroup5"] = args ? args.dhGroup5 : undefined;
            resourceInputs["dhKeypairCache"] = args ? args.dhKeypairCache : undefined;
            resourceInputs["dhKeypairCount"] = args ? args.dhKeypairCount : undefined;
            resourceInputs["dhKeypairThrottle"] = args ? args.dhKeypairThrottle : undefined;
            resourceInputs["dhMode"] = args ? args.dhMode : undefined;
            resourceInputs["dhMultiprocess"] = args ? args.dhMultiprocess : undefined;
            resourceInputs["dhWorkerCount"] = args ? args.dhWorkerCount : undefined;
            resourceInputs["embryonicLimit"] = args ? args.embryonicLimit : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemIke.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemIke resources.
 */
export interface SystemIkeState {
    /**
     * Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
     */
    dhGroup1?: pulumi.Input<inputs.SystemIkeDhGroup1>;
    /**
     * Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
     */
    dhGroup14?: pulumi.Input<inputs.SystemIkeDhGroup14>;
    /**
     * Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
     */
    dhGroup15?: pulumi.Input<inputs.SystemIkeDhGroup15>;
    /**
     * Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
     */
    dhGroup16?: pulumi.Input<inputs.SystemIkeDhGroup16>;
    /**
     * Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
     */
    dhGroup17?: pulumi.Input<inputs.SystemIkeDhGroup17>;
    /**
     * Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
     */
    dhGroup18?: pulumi.Input<inputs.SystemIkeDhGroup18>;
    /**
     * Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
     */
    dhGroup19?: pulumi.Input<inputs.SystemIkeDhGroup19>;
    /**
     * Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
     */
    dhGroup2?: pulumi.Input<inputs.SystemIkeDhGroup2>;
    /**
     * Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
     */
    dhGroup20?: pulumi.Input<inputs.SystemIkeDhGroup20>;
    /**
     * Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
     */
    dhGroup21?: pulumi.Input<inputs.SystemIkeDhGroup21>;
    /**
     * Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
     */
    dhGroup27?: pulumi.Input<inputs.SystemIkeDhGroup27>;
    /**
     * Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
     */
    dhGroup28?: pulumi.Input<inputs.SystemIkeDhGroup28>;
    /**
     * Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
     */
    dhGroup29?: pulumi.Input<inputs.SystemIkeDhGroup29>;
    /**
     * Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
     */
    dhGroup30?: pulumi.Input<inputs.SystemIkeDhGroup30>;
    /**
     * Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
     */
    dhGroup31?: pulumi.Input<inputs.SystemIkeDhGroup31>;
    /**
     * Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
     */
    dhGroup32?: pulumi.Input<inputs.SystemIkeDhGroup32>;
    /**
     * Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
     */
    dhGroup5?: pulumi.Input<inputs.SystemIkeDhGroup5>;
    /**
     * Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
     */
    dhKeypairCache?: pulumi.Input<string>;
    /**
     * Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
     */
    dhKeypairCount?: pulumi.Input<number>;
    /**
     * Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
     */
    dhKeypairThrottle?: pulumi.Input<string>;
    /**
     * Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
     */
    dhMode?: pulumi.Input<string>;
    /**
     * Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
     */
    dhMultiprocess?: pulumi.Input<string>;
    /**
     * Number of Diffie-Hellman workers to start.
     */
    dhWorkerCount?: pulumi.Input<number>;
    /**
     * Maximum number of IPsec tunnels to negotiate simultaneously.
     */
    embryonicLimit?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     *
     * The `dhGroup1` block supports:
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemIke resource.
 */
export interface SystemIkeArgs {
    /**
     * Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
     */
    dhGroup1?: pulumi.Input<inputs.SystemIkeDhGroup1>;
    /**
     * Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
     */
    dhGroup14?: pulumi.Input<inputs.SystemIkeDhGroup14>;
    /**
     * Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
     */
    dhGroup15?: pulumi.Input<inputs.SystemIkeDhGroup15>;
    /**
     * Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
     */
    dhGroup16?: pulumi.Input<inputs.SystemIkeDhGroup16>;
    /**
     * Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
     */
    dhGroup17?: pulumi.Input<inputs.SystemIkeDhGroup17>;
    /**
     * Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
     */
    dhGroup18?: pulumi.Input<inputs.SystemIkeDhGroup18>;
    /**
     * Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
     */
    dhGroup19?: pulumi.Input<inputs.SystemIkeDhGroup19>;
    /**
     * Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
     */
    dhGroup2?: pulumi.Input<inputs.SystemIkeDhGroup2>;
    /**
     * Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
     */
    dhGroup20?: pulumi.Input<inputs.SystemIkeDhGroup20>;
    /**
     * Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
     */
    dhGroup21?: pulumi.Input<inputs.SystemIkeDhGroup21>;
    /**
     * Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
     */
    dhGroup27?: pulumi.Input<inputs.SystemIkeDhGroup27>;
    /**
     * Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
     */
    dhGroup28?: pulumi.Input<inputs.SystemIkeDhGroup28>;
    /**
     * Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
     */
    dhGroup29?: pulumi.Input<inputs.SystemIkeDhGroup29>;
    /**
     * Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
     */
    dhGroup30?: pulumi.Input<inputs.SystemIkeDhGroup30>;
    /**
     * Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
     */
    dhGroup31?: pulumi.Input<inputs.SystemIkeDhGroup31>;
    /**
     * Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
     */
    dhGroup32?: pulumi.Input<inputs.SystemIkeDhGroup32>;
    /**
     * Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
     */
    dhGroup5?: pulumi.Input<inputs.SystemIkeDhGroup5>;
    /**
     * Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
     */
    dhKeypairCache?: pulumi.Input<string>;
    /**
     * Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
     */
    dhKeypairCount?: pulumi.Input<number>;
    /**
     * Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
     */
    dhKeypairThrottle?: pulumi.Input<string>;
    /**
     * Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
     */
    dhMode?: pulumi.Input<string>;
    /**
     * Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
     */
    dhMultiprocess?: pulumi.Input<string>;
    /**
     * Number of Diffie-Hellman workers to start.
     */
    dhWorkerCount?: pulumi.Input<number>;
    /**
     * Maximum number of IPsec tunnels to negotiate simultaneously.
     */
    embryonicLimit?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     *
     * The `dhGroup1` block supports:
     */
    vdomparam?: pulumi.Input<string>;
}
