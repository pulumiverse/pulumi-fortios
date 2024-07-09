// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure quarantine options.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.antivirus.Quarantine("trname", {
 *     agelimit: 0,
 *     destination: "disk",
 *     lowspace: "ovrw-old",
 *     maxfilesize: 0,
 *     quarantineQuota: 0,
 *     storeBlocked: "imap smtp pop3 http ftp nntp imaps smtps pop3s ftps mapi cifs",
 *     storeHeuristic: "imap smtp pop3 http ftp nntp imaps smtps pop3s https ftps mapi cifs",
 *     storeInfected: "imap smtp pop3 http ftp nntp imaps smtps pop3s https ftps mapi cifs",
 * });
 * ```
 *
 * ## Import
 *
 * Antivirus Quarantine can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:antivirus/quarantine:Quarantine labelname AntivirusQuarantine
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:antivirus/quarantine:Quarantine labelname AntivirusQuarantine
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Quarantine extends pulumi.CustomResource {
    /**
     * Get an existing Quarantine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QuarantineState, opts?: pulumi.CustomResourceOptions): Quarantine {
        return new Quarantine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:antivirus/quarantine:Quarantine';

    /**
     * Returns true if the given object is an instance of Quarantine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Quarantine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Quarantine.__pulumiType;
    }

    /**
     * Age limit for quarantined files (0 - 479 hours, 0 means forever).
     */
    public readonly agelimit!: pulumi.Output<number>;
    /**
     * Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
     */
    public readonly dropBlocked!: pulumi.Output<string>;
    /**
     * Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
     */
    public readonly dropHeuristic!: pulumi.Output<string>;
    /**
     * Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
     */
    public readonly dropInfected!: pulumi.Output<string>;
    /**
     * Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
     */
    public readonly dropMachineLearning!: pulumi.Output<string>;
    /**
     * Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
     */
    public readonly lowspace!: pulumi.Output<string>;
    /**
     * Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
     */
    public readonly maxfilesize!: pulumi.Output<number>;
    /**
     * The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
     */
    public readonly quarantineQuota!: pulumi.Output<number>;
    /**
     * Quarantine blocked files found in sessions using the selected protocols.
     */
    public readonly storeBlocked!: pulumi.Output<string>;
    /**
     * Quarantine files detected by heuristics found in sessions using the selected protocols.
     */
    public readonly storeHeuristic!: pulumi.Output<string>;
    /**
     * Quarantine infected files found in sessions using the selected protocols.
     */
    public readonly storeInfected!: pulumi.Output<string>;
    /**
     * Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
     */
    public readonly storeMachineLearning!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Quarantine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: QuarantineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QuarantineArgs | QuarantineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QuarantineState | undefined;
            resourceInputs["agelimit"] = state ? state.agelimit : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["dropBlocked"] = state ? state.dropBlocked : undefined;
            resourceInputs["dropHeuristic"] = state ? state.dropHeuristic : undefined;
            resourceInputs["dropInfected"] = state ? state.dropInfected : undefined;
            resourceInputs["dropMachineLearning"] = state ? state.dropMachineLearning : undefined;
            resourceInputs["lowspace"] = state ? state.lowspace : undefined;
            resourceInputs["maxfilesize"] = state ? state.maxfilesize : undefined;
            resourceInputs["quarantineQuota"] = state ? state.quarantineQuota : undefined;
            resourceInputs["storeBlocked"] = state ? state.storeBlocked : undefined;
            resourceInputs["storeHeuristic"] = state ? state.storeHeuristic : undefined;
            resourceInputs["storeInfected"] = state ? state.storeInfected : undefined;
            resourceInputs["storeMachineLearning"] = state ? state.storeMachineLearning : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as QuarantineArgs | undefined;
            resourceInputs["agelimit"] = args ? args.agelimit : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["dropBlocked"] = args ? args.dropBlocked : undefined;
            resourceInputs["dropHeuristic"] = args ? args.dropHeuristic : undefined;
            resourceInputs["dropInfected"] = args ? args.dropInfected : undefined;
            resourceInputs["dropMachineLearning"] = args ? args.dropMachineLearning : undefined;
            resourceInputs["lowspace"] = args ? args.lowspace : undefined;
            resourceInputs["maxfilesize"] = args ? args.maxfilesize : undefined;
            resourceInputs["quarantineQuota"] = args ? args.quarantineQuota : undefined;
            resourceInputs["storeBlocked"] = args ? args.storeBlocked : undefined;
            resourceInputs["storeHeuristic"] = args ? args.storeHeuristic : undefined;
            resourceInputs["storeInfected"] = args ? args.storeInfected : undefined;
            resourceInputs["storeMachineLearning"] = args ? args.storeMachineLearning : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Quarantine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Quarantine resources.
 */
export interface QuarantineState {
    /**
     * Age limit for quarantined files (0 - 479 hours, 0 means forever).
     */
    agelimit?: pulumi.Input<number>;
    /**
     * Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
     */
    destination?: pulumi.Input<string>;
    /**
     * Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
     */
    dropBlocked?: pulumi.Input<string>;
    /**
     * Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
     */
    dropHeuristic?: pulumi.Input<string>;
    /**
     * Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
     */
    dropInfected?: pulumi.Input<string>;
    /**
     * Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
     */
    dropMachineLearning?: pulumi.Input<string>;
    /**
     * Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
     */
    lowspace?: pulumi.Input<string>;
    /**
     * Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
     */
    maxfilesize?: pulumi.Input<number>;
    /**
     * The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
     */
    quarantineQuota?: pulumi.Input<number>;
    /**
     * Quarantine blocked files found in sessions using the selected protocols.
     */
    storeBlocked?: pulumi.Input<string>;
    /**
     * Quarantine files detected by heuristics found in sessions using the selected protocols.
     */
    storeHeuristic?: pulumi.Input<string>;
    /**
     * Quarantine infected files found in sessions using the selected protocols.
     */
    storeInfected?: pulumi.Input<string>;
    /**
     * Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
     */
    storeMachineLearning?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Quarantine resource.
 */
export interface QuarantineArgs {
    /**
     * Age limit for quarantined files (0 - 479 hours, 0 means forever).
     */
    agelimit?: pulumi.Input<number>;
    /**
     * Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
     */
    destination?: pulumi.Input<string>;
    /**
     * Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
     */
    dropBlocked?: pulumi.Input<string>;
    /**
     * Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
     */
    dropHeuristic?: pulumi.Input<string>;
    /**
     * Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
     */
    dropInfected?: pulumi.Input<string>;
    /**
     * Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
     */
    dropMachineLearning?: pulumi.Input<string>;
    /**
     * Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
     */
    lowspace?: pulumi.Input<string>;
    /**
     * Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
     */
    maxfilesize?: pulumi.Input<number>;
    /**
     * The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
     */
    quarantineQuota?: pulumi.Input<number>;
    /**
     * Quarantine blocked files found in sessions using the selected protocols.
     */
    storeBlocked?: pulumi.Input<string>;
    /**
     * Quarantine files detected by heuristics found in sessions using the selected protocols.
     */
    storeHeuristic?: pulumi.Input<string>;
    /**
     * Quarantine infected files found in sessions using the selected protocols.
     */
    storeInfected?: pulumi.Input<string>;
    /**
     * Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
     */
    storeMachineLearning?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
