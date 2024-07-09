// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure protocol options.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Profileprotocoloptions("trname", {
 *     dns: {
 *         ports: 53,
 *         status: "enable",
 *     },
 *     ftp: {
 *         comfortAmount: 1,
 *         comfortInterval: 10,
 *         inspectAll: "disable",
 *         options: "splice",
 *         oversizeLimit: 10,
 *         ports: 21,
 *         scanBzip2: "enable",
 *         status: "enable",
 *         uncompressedNestLimit: 12,
 *         uncompressedOversizeLimit: 10,
 *     },
 *     http: {
 *         blockPageStatusCode: 403,
 *         comfortAmount: 1,
 *         comfortInterval: 10,
 *         fortinetBar: "disable",
 *         fortinetBarPort: 8011,
 *         httpPolicy: "disable",
 *         inspectAll: "disable",
 *         oversizeLimit: 10,
 *         ports: 80,
 *         rangeBlock: "disable",
 *         retryCount: 0,
 *         scanBzip2: "enable",
 *         status: "enable",
 *         streamingContentBypass: "enable",
 *         stripXForwardedFor: "disable",
 *         switchingProtocols: "bypass",
 *         uncompressedNestLimit: 12,
 *         uncompressedOversizeLimit: 10,
 *     },
 *     imap: {
 *         inspectAll: "disable",
 *         options: "fragmail",
 *         oversizeLimit: 10,
 *         ports: 143,
 *         scanBzip2: "enable",
 *         status: "enable",
 *         uncompressedNestLimit: 12,
 *         uncompressedOversizeLimit: 10,
 *     },
 *     mailSignature: {
 *         status: "disable",
 *     },
 *     mapi: {
 *         options: "fragmail",
 *         oversizeLimit: 10,
 *         ports: 135,
 *         scanBzip2: "enable",
 *         status: "enable",
 *         uncompressedNestLimit: 12,
 *         uncompressedOversizeLimit: 10,
 *     },
 *     nntp: {
 *         inspectAll: "disable",
 *         options: "splice",
 *         oversizeLimit: 10,
 *         ports: 119,
 *         scanBzip2: "enable",
 *         status: "enable",
 *         uncompressedNestLimit: 12,
 *         uncompressedOversizeLimit: 10,
 *     },
 *     oversizeLog: "disable",
 *     pop3: {
 *         inspectAll: "disable",
 *         options: "fragmail",
 *         oversizeLimit: 10,
 *         ports: 110,
 *         scanBzip2: "enable",
 *         status: "enable",
 *         uncompressedNestLimit: 12,
 *         uncompressedOversizeLimit: 10,
 *     },
 *     rpcOverHttp: "disable",
 *     smtp: {
 *         inspectAll: "disable",
 *         options: "fragmail splice",
 *         oversizeLimit: 10,
 *         ports: 25,
 *         scanBzip2: "enable",
 *         serverBusy: "disable",
 *         status: "enable",
 *         uncompressedNestLimit: 12,
 *         uncompressedOversizeLimit: 10,
 *     },
 *     switchingProtocolsLog: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall ProfileProtocolOptions can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/profileprotocoloptions:Profileprotocoloptions labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/profileprotocoloptions:Profileprotocoloptions labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Profileprotocoloptions extends pulumi.CustomResource {
    /**
     * Get an existing Profileprotocoloptions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileprotocoloptionsState, opts?: pulumi.CustomResourceOptions): Profileprotocoloptions {
        return new Profileprotocoloptions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/profileprotocoloptions:Profileprotocoloptions';

    /**
     * Returns true if the given object is an instance of Profileprotocoloptions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Profileprotocoloptions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Profileprotocoloptions.__pulumiType;
    }

    /**
     * Configure CIFS protocol options. The structure of `cifs` block is documented below.
     */
    public readonly cifs!: pulumi.Output<outputs.firewall.ProfileprotocoloptionsCifs>;
    /**
     * Optional comments.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Configure DNS protocol options. The structure of `dns` block is documented below.
     */
    public readonly dns!: pulumi.Output<outputs.firewall.ProfileprotocoloptionsDns>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    public readonly featureSet!: pulumi.Output<string>;
    /**
     * Configure FTP protocol options. The structure of `ftp` block is documented below.
     */
    public readonly ftp!: pulumi.Output<outputs.firewall.ProfileprotocoloptionsFtp>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Configure HTTP protocol options. The structure of `http` block is documented below.
     */
    public readonly http!: pulumi.Output<outputs.firewall.ProfileprotocoloptionsHttp>;
    /**
     * Configure IMAP protocol options. The structure of `imap` block is documented below.
     */
    public readonly imap!: pulumi.Output<outputs.firewall.ProfileprotocoloptionsImap>;
    /**
     * Configure Mail signature. The structure of `mailSignature` block is documented below.
     */
    public readonly mailSignature!: pulumi.Output<outputs.firewall.ProfileprotocoloptionsMailSignature>;
    /**
     * Configure MAPI protocol options. The structure of `mapi` block is documented below.
     */
    public readonly mapi!: pulumi.Output<outputs.firewall.ProfileprotocoloptionsMapi>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configure NNTP protocol options. The structure of `nntp` block is documented below.
     */
    public readonly nntp!: pulumi.Output<outputs.firewall.ProfileprotocoloptionsNntp>;
    /**
     * Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
     */
    public readonly oversizeLog!: pulumi.Output<string>;
    /**
     * Configure POP3 protocol options. The structure of `pop3` block is documented below.
     */
    public readonly pop3!: pulumi.Output<outputs.firewall.ProfileprotocoloptionsPop3>;
    /**
     * Name of the replacement message group to be used
     */
    public readonly replacemsgGroup!: pulumi.Output<string>;
    /**
     * Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.
     */
    public readonly rpcOverHttp!: pulumi.Output<string>;
    /**
     * Configure SMTP protocol options. The structure of `smtp` block is documented below.
     */
    public readonly smtp!: pulumi.Output<outputs.firewall.ProfileprotocoloptionsSmtp>;
    /**
     * Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
     */
    public readonly ssh!: pulumi.Output<outputs.firewall.ProfileprotocoloptionsSsh>;
    /**
     * Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
     */
    public readonly switchingProtocolsLog!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Profileprotocoloptions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProfileprotocoloptionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileprotocoloptionsArgs | ProfileprotocoloptionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProfileprotocoloptionsState | undefined;
            resourceInputs["cifs"] = state ? state.cifs : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dns"] = state ? state.dns : undefined;
            resourceInputs["featureSet"] = state ? state.featureSet : undefined;
            resourceInputs["ftp"] = state ? state.ftp : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["http"] = state ? state.http : undefined;
            resourceInputs["imap"] = state ? state.imap : undefined;
            resourceInputs["mailSignature"] = state ? state.mailSignature : undefined;
            resourceInputs["mapi"] = state ? state.mapi : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nntp"] = state ? state.nntp : undefined;
            resourceInputs["oversizeLog"] = state ? state.oversizeLog : undefined;
            resourceInputs["pop3"] = state ? state.pop3 : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["rpcOverHttp"] = state ? state.rpcOverHttp : undefined;
            resourceInputs["smtp"] = state ? state.smtp : undefined;
            resourceInputs["ssh"] = state ? state.ssh : undefined;
            resourceInputs["switchingProtocolsLog"] = state ? state.switchingProtocolsLog : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ProfileprotocoloptionsArgs | undefined;
            resourceInputs["cifs"] = args ? args.cifs : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dns"] = args ? args.dns : undefined;
            resourceInputs["featureSet"] = args ? args.featureSet : undefined;
            resourceInputs["ftp"] = args ? args.ftp : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["http"] = args ? args.http : undefined;
            resourceInputs["imap"] = args ? args.imap : undefined;
            resourceInputs["mailSignature"] = args ? args.mailSignature : undefined;
            resourceInputs["mapi"] = args ? args.mapi : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nntp"] = args ? args.nntp : undefined;
            resourceInputs["oversizeLog"] = args ? args.oversizeLog : undefined;
            resourceInputs["pop3"] = args ? args.pop3 : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["rpcOverHttp"] = args ? args.rpcOverHttp : undefined;
            resourceInputs["smtp"] = args ? args.smtp : undefined;
            resourceInputs["ssh"] = args ? args.ssh : undefined;
            resourceInputs["switchingProtocolsLog"] = args ? args.switchingProtocolsLog : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Profileprotocoloptions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Profileprotocoloptions resources.
 */
export interface ProfileprotocoloptionsState {
    /**
     * Configure CIFS protocol options. The structure of `cifs` block is documented below.
     */
    cifs?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsCifs>;
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Configure DNS protocol options. The structure of `dns` block is documented below.
     */
    dns?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsDns>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    featureSet?: pulumi.Input<string>;
    /**
     * Configure FTP protocol options. The structure of `ftp` block is documented below.
     */
    ftp?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsFtp>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Configure HTTP protocol options. The structure of `http` block is documented below.
     */
    http?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsHttp>;
    /**
     * Configure IMAP protocol options. The structure of `imap` block is documented below.
     */
    imap?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsImap>;
    /**
     * Configure Mail signature. The structure of `mailSignature` block is documented below.
     */
    mailSignature?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsMailSignature>;
    /**
     * Configure MAPI protocol options. The structure of `mapi` block is documented below.
     */
    mapi?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsMapi>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Configure NNTP protocol options. The structure of `nntp` block is documented below.
     */
    nntp?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsNntp>;
    /**
     * Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
     */
    oversizeLog?: pulumi.Input<string>;
    /**
     * Configure POP3 protocol options. The structure of `pop3` block is documented below.
     */
    pop3?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsPop3>;
    /**
     * Name of the replacement message group to be used
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.
     */
    rpcOverHttp?: pulumi.Input<string>;
    /**
     * Configure SMTP protocol options. The structure of `smtp` block is documented below.
     */
    smtp?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsSmtp>;
    /**
     * Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
     */
    ssh?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsSsh>;
    /**
     * Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
     */
    switchingProtocolsLog?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Profileprotocoloptions resource.
 */
export interface ProfileprotocoloptionsArgs {
    /**
     * Configure CIFS protocol options. The structure of `cifs` block is documented below.
     */
    cifs?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsCifs>;
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Configure DNS protocol options. The structure of `dns` block is documented below.
     */
    dns?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsDns>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    featureSet?: pulumi.Input<string>;
    /**
     * Configure FTP protocol options. The structure of `ftp` block is documented below.
     */
    ftp?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsFtp>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Configure HTTP protocol options. The structure of `http` block is documented below.
     */
    http?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsHttp>;
    /**
     * Configure IMAP protocol options. The structure of `imap` block is documented below.
     */
    imap?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsImap>;
    /**
     * Configure Mail signature. The structure of `mailSignature` block is documented below.
     */
    mailSignature?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsMailSignature>;
    /**
     * Configure MAPI protocol options. The structure of `mapi` block is documented below.
     */
    mapi?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsMapi>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Configure NNTP protocol options. The structure of `nntp` block is documented below.
     */
    nntp?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsNntp>;
    /**
     * Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
     */
    oversizeLog?: pulumi.Input<string>;
    /**
     * Configure POP3 protocol options. The structure of `pop3` block is documented below.
     */
    pop3?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsPop3>;
    /**
     * Name of the replacement message group to be used
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.
     */
    rpcOverHttp?: pulumi.Input<string>;
    /**
     * Configure SMTP protocol options. The structure of `smtp` block is documented below.
     */
    smtp?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsSmtp>;
    /**
     * Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
     */
    ssh?: pulumi.Input<inputs.firewall.ProfileprotocoloptionsSsh>;
    /**
     * Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
     */
    switchingProtocolsLog?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
