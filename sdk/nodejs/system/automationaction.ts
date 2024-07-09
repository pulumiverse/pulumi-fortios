// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Action for automation stitches.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Automationaction("trname", {
 *     actionType: "email",
 *     awsDomain: "amazonaws.com",
 *     delay: 0,
 *     emailSubject: "SUBJECT1",
 *     method: "post",
 *     minimumInterval: 1,
 *     protocol: "http",
 *     required: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * System AutomationAction can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/automationaction:Automationaction labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/automationaction:Automationaction labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Automationaction extends pulumi.CustomResource {
    /**
     * Get an existing Automationaction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutomationactionState, opts?: pulumi.CustomResourceOptions): Automationaction {
        return new Automationaction(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/automationaction:Automationaction';

    /**
     * Returns true if the given object is an instance of Automationaction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Automationaction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Automationaction.__pulumiType;
    }

    /**
     * Access profile for CLI script action to access FortiGate features.
     */
    public readonly accprofile!: pulumi.Output<string>;
    /**
     * Action type.
     */
    public readonly actionType!: pulumi.Output<string>;
    /**
     * AliCloud AccessKey ID.
     */
    public readonly alicloudAccessKeyId!: pulumi.Output<string>;
    /**
     * AliCloud AccessKey secret.
     */
    public readonly alicloudAccessKeySecret!: pulumi.Output<string | undefined>;
    /**
     * AliCloud account ID.
     */
    public readonly alicloudAccountId!: pulumi.Output<string>;
    /**
     * AliCloud function name.
     */
    public readonly alicloudFunction!: pulumi.Output<string>;
    /**
     * AliCloud function authorization type. Valid values: `anonymous`, `function`.
     */
    public readonly alicloudFunctionAuthorization!: pulumi.Output<string>;
    /**
     * AliCloud function domain.
     */
    public readonly alicloudFunctionDomain!: pulumi.Output<string>;
    /**
     * AliCloud region.
     */
    public readonly alicloudRegion!: pulumi.Output<string>;
    /**
     * AliCloud service name.
     */
    public readonly alicloudService!: pulumi.Output<string>;
    /**
     * AliCloud version.
     */
    public readonly alicloudVersion!: pulumi.Output<string>;
    /**
     * AWS API Gateway ID.
     */
    public readonly awsApiId!: pulumi.Output<string>;
    /**
     * AWS API Gateway API key.
     */
    public readonly awsApiKey!: pulumi.Output<string | undefined>;
    /**
     * AWS API Gateway path.
     */
    public readonly awsApiPath!: pulumi.Output<string>;
    /**
     * AWS API Gateway deployment stage name.
     */
    public readonly awsApiStage!: pulumi.Output<string>;
    /**
     * AWS domain.
     */
    public readonly awsDomain!: pulumi.Output<string>;
    /**
     * AWS region.
     */
    public readonly awsRegion!: pulumi.Output<string>;
    /**
     * Azure function API key.
     */
    public readonly azureApiKey!: pulumi.Output<string | undefined>;
    /**
     * Azure function application name.
     */
    public readonly azureApp!: pulumi.Output<string>;
    /**
     * Azure function domain.
     */
    public readonly azureDomain!: pulumi.Output<string>;
    /**
     * Azure function name.
     */
    public readonly azureFunction!: pulumi.Output<string>;
    /**
     * Azure function authorization level. Valid values: `anonymous`, `function`, `admin`.
     */
    public readonly azureFunctionAuthorization!: pulumi.Output<string>;
    /**
     * Delay before execution (in seconds).
     */
    public readonly delay!: pulumi.Output<number>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Email body.
     */
    public readonly emailBody!: pulumi.Output<string>;
    /**
     * Email sender name.
     */
    public readonly emailFrom!: pulumi.Output<string | undefined>;
    /**
     * Email subject.
     */
    public readonly emailSubject!: pulumi.Output<string | undefined>;
    /**
     * Email addresses. The structure of `emailTo` block is documented below.
     */
    public readonly emailTos!: pulumi.Output<outputs.system.AutomationactionEmailTo[] | undefined>;
    /**
     * Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric. Valid values: `enable`, `disable`.
     */
    public readonly executeSecurityFabric!: pulumi.Output<string>;
    /**
     * Enable/disable use of your FortiCare email address as the email-to address. Valid values: `enable`, `disable`.
     */
    public readonly forticareEmail!: pulumi.Output<string>;
    /**
     * Google Cloud function name.
     */
    public readonly gcpFunction!: pulumi.Output<string>;
    /**
     * Google Cloud function domain.
     */
    public readonly gcpFunctionDomain!: pulumi.Output<string>;
    /**
     * Google Cloud function region.
     */
    public readonly gcpFunctionRegion!: pulumi.Output<string>;
    /**
     * Google Cloud Platform project name.
     */
    public readonly gcpProject!: pulumi.Output<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Request headers. The structure of `headers` block is documented below.
     */
    public readonly headers!: pulumi.Output<outputs.system.AutomationactionHeader[] | undefined>;
    /**
     * Request body (if necessary). Should be serialized json string.
     */
    public readonly httpBody!: pulumi.Output<string | undefined>;
    /**
     * Request headers. The structure of `httpHeaders` block is documented below.
     */
    public readonly httpHeaders!: pulumi.Output<outputs.system.AutomationactionHttpHeader[] | undefined>;
    /**
     * Message content.
     */
    public readonly message!: pulumi.Output<string>;
    /**
     * Message type. Valid values: `text`, `json`.
     */
    public readonly messageType!: pulumi.Output<string>;
    /**
     * Request method (POST, PUT, GET, PATCH or DELETE). Valid values: `post`, `put`, `get`, `patch`, `delete`.
     */
    public readonly method!: pulumi.Output<string>;
    /**
     * Limit execution to no more than once in this interval (in seconds).
     */
    public readonly minimumInterval!: pulumi.Output<number>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of megabytes to limit script output to (1 - 1024, default = 10).
     */
    public readonly outputSize!: pulumi.Output<number>;
    /**
     * Protocol port.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Request protocol. Valid values: `http`, `https`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Enable/disable replacement message. Valid values: `enable`, `disable`.
     */
    public readonly replacementMessage!: pulumi.Output<string>;
    /**
     * Replacement message group.
     */
    public readonly replacemsgGroup!: pulumi.Output<string>;
    /**
     * Required in action chain. Valid values: `enable`, `disable`.
     */
    public readonly required!: pulumi.Output<string>;
    /**
     * CLI script.
     */
    public readonly script!: pulumi.Output<string | undefined>;
    /**
     * NSX SDN connector names. The structure of `sdnConnector` block is documented below.
     */
    public readonly sdnConnectors!: pulumi.Output<outputs.system.AutomationactionSdnConnector[] | undefined>;
    /**
     * NSX security tag.
     */
    public readonly securityTag!: pulumi.Output<string>;
    /**
     * System action type. Valid values: `reboot`, `shutdown`, `backup-config`.
     */
    public readonly systemAction!: pulumi.Output<string>;
    /**
     * Maximum running time for this script in seconds (0 = no timeout).
     */
    public readonly timeout!: pulumi.Output<number>;
    /**
     * Custom TLS certificate for API request.
     */
    public readonly tlsCertificate!: pulumi.Output<string>;
    /**
     * Request API URI.
     */
    public readonly uri!: pulumi.Output<string | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;
    /**
     * Enable/disable verification of the remote host certificate. Valid values: `enable`, `disable`.
     */
    public readonly verifyHostCert!: pulumi.Output<string>;

    /**
     * Create a Automationaction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AutomationactionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutomationactionArgs | AutomationactionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutomationactionState | undefined;
            resourceInputs["accprofile"] = state ? state.accprofile : undefined;
            resourceInputs["actionType"] = state ? state.actionType : undefined;
            resourceInputs["alicloudAccessKeyId"] = state ? state.alicloudAccessKeyId : undefined;
            resourceInputs["alicloudAccessKeySecret"] = state ? state.alicloudAccessKeySecret : undefined;
            resourceInputs["alicloudAccountId"] = state ? state.alicloudAccountId : undefined;
            resourceInputs["alicloudFunction"] = state ? state.alicloudFunction : undefined;
            resourceInputs["alicloudFunctionAuthorization"] = state ? state.alicloudFunctionAuthorization : undefined;
            resourceInputs["alicloudFunctionDomain"] = state ? state.alicloudFunctionDomain : undefined;
            resourceInputs["alicloudRegion"] = state ? state.alicloudRegion : undefined;
            resourceInputs["alicloudService"] = state ? state.alicloudService : undefined;
            resourceInputs["alicloudVersion"] = state ? state.alicloudVersion : undefined;
            resourceInputs["awsApiId"] = state ? state.awsApiId : undefined;
            resourceInputs["awsApiKey"] = state ? state.awsApiKey : undefined;
            resourceInputs["awsApiPath"] = state ? state.awsApiPath : undefined;
            resourceInputs["awsApiStage"] = state ? state.awsApiStage : undefined;
            resourceInputs["awsDomain"] = state ? state.awsDomain : undefined;
            resourceInputs["awsRegion"] = state ? state.awsRegion : undefined;
            resourceInputs["azureApiKey"] = state ? state.azureApiKey : undefined;
            resourceInputs["azureApp"] = state ? state.azureApp : undefined;
            resourceInputs["azureDomain"] = state ? state.azureDomain : undefined;
            resourceInputs["azureFunction"] = state ? state.azureFunction : undefined;
            resourceInputs["azureFunctionAuthorization"] = state ? state.azureFunctionAuthorization : undefined;
            resourceInputs["delay"] = state ? state.delay : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["emailBody"] = state ? state.emailBody : undefined;
            resourceInputs["emailFrom"] = state ? state.emailFrom : undefined;
            resourceInputs["emailSubject"] = state ? state.emailSubject : undefined;
            resourceInputs["emailTos"] = state ? state.emailTos : undefined;
            resourceInputs["executeSecurityFabric"] = state ? state.executeSecurityFabric : undefined;
            resourceInputs["forticareEmail"] = state ? state.forticareEmail : undefined;
            resourceInputs["gcpFunction"] = state ? state.gcpFunction : undefined;
            resourceInputs["gcpFunctionDomain"] = state ? state.gcpFunctionDomain : undefined;
            resourceInputs["gcpFunctionRegion"] = state ? state.gcpFunctionRegion : undefined;
            resourceInputs["gcpProject"] = state ? state.gcpProject : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["headers"] = state ? state.headers : undefined;
            resourceInputs["httpBody"] = state ? state.httpBody : undefined;
            resourceInputs["httpHeaders"] = state ? state.httpHeaders : undefined;
            resourceInputs["message"] = state ? state.message : undefined;
            resourceInputs["messageType"] = state ? state.messageType : undefined;
            resourceInputs["method"] = state ? state.method : undefined;
            resourceInputs["minimumInterval"] = state ? state.minimumInterval : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outputSize"] = state ? state.outputSize : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["replacementMessage"] = state ? state.replacementMessage : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["required"] = state ? state.required : undefined;
            resourceInputs["script"] = state ? state.script : undefined;
            resourceInputs["sdnConnectors"] = state ? state.sdnConnectors : undefined;
            resourceInputs["securityTag"] = state ? state.securityTag : undefined;
            resourceInputs["systemAction"] = state ? state.systemAction : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["tlsCertificate"] = state ? state.tlsCertificate : undefined;
            resourceInputs["uri"] = state ? state.uri : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["verifyHostCert"] = state ? state.verifyHostCert : undefined;
        } else {
            const args = argsOrState as AutomationactionArgs | undefined;
            resourceInputs["accprofile"] = args ? args.accprofile : undefined;
            resourceInputs["actionType"] = args ? args.actionType : undefined;
            resourceInputs["alicloudAccessKeyId"] = args ? args.alicloudAccessKeyId : undefined;
            resourceInputs["alicloudAccessKeySecret"] = args?.alicloudAccessKeySecret ? pulumi.secret(args.alicloudAccessKeySecret) : undefined;
            resourceInputs["alicloudAccountId"] = args ? args.alicloudAccountId : undefined;
            resourceInputs["alicloudFunction"] = args ? args.alicloudFunction : undefined;
            resourceInputs["alicloudFunctionAuthorization"] = args ? args.alicloudFunctionAuthorization : undefined;
            resourceInputs["alicloudFunctionDomain"] = args ? args.alicloudFunctionDomain : undefined;
            resourceInputs["alicloudRegion"] = args ? args.alicloudRegion : undefined;
            resourceInputs["alicloudService"] = args ? args.alicloudService : undefined;
            resourceInputs["alicloudVersion"] = args ? args.alicloudVersion : undefined;
            resourceInputs["awsApiId"] = args ? args.awsApiId : undefined;
            resourceInputs["awsApiKey"] = args?.awsApiKey ? pulumi.secret(args.awsApiKey) : undefined;
            resourceInputs["awsApiPath"] = args ? args.awsApiPath : undefined;
            resourceInputs["awsApiStage"] = args ? args.awsApiStage : undefined;
            resourceInputs["awsDomain"] = args ? args.awsDomain : undefined;
            resourceInputs["awsRegion"] = args ? args.awsRegion : undefined;
            resourceInputs["azureApiKey"] = args?.azureApiKey ? pulumi.secret(args.azureApiKey) : undefined;
            resourceInputs["azureApp"] = args ? args.azureApp : undefined;
            resourceInputs["azureDomain"] = args ? args.azureDomain : undefined;
            resourceInputs["azureFunction"] = args ? args.azureFunction : undefined;
            resourceInputs["azureFunctionAuthorization"] = args ? args.azureFunctionAuthorization : undefined;
            resourceInputs["delay"] = args ? args.delay : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["emailBody"] = args ? args.emailBody : undefined;
            resourceInputs["emailFrom"] = args ? args.emailFrom : undefined;
            resourceInputs["emailSubject"] = args ? args.emailSubject : undefined;
            resourceInputs["emailTos"] = args ? args.emailTos : undefined;
            resourceInputs["executeSecurityFabric"] = args ? args.executeSecurityFabric : undefined;
            resourceInputs["forticareEmail"] = args ? args.forticareEmail : undefined;
            resourceInputs["gcpFunction"] = args ? args.gcpFunction : undefined;
            resourceInputs["gcpFunctionDomain"] = args ? args.gcpFunctionDomain : undefined;
            resourceInputs["gcpFunctionRegion"] = args ? args.gcpFunctionRegion : undefined;
            resourceInputs["gcpProject"] = args ? args.gcpProject : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["headers"] = args ? args.headers : undefined;
            resourceInputs["httpBody"] = args ? args.httpBody : undefined;
            resourceInputs["httpHeaders"] = args ? args.httpHeaders : undefined;
            resourceInputs["message"] = args ? args.message : undefined;
            resourceInputs["messageType"] = args ? args.messageType : undefined;
            resourceInputs["method"] = args ? args.method : undefined;
            resourceInputs["minimumInterval"] = args ? args.minimumInterval : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outputSize"] = args ? args.outputSize : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["replacementMessage"] = args ? args.replacementMessage : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["required"] = args ? args.required : undefined;
            resourceInputs["script"] = args ? args.script : undefined;
            resourceInputs["sdnConnectors"] = args ? args.sdnConnectors : undefined;
            resourceInputs["securityTag"] = args ? args.securityTag : undefined;
            resourceInputs["systemAction"] = args ? args.systemAction : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["tlsCertificate"] = args ? args.tlsCertificate : undefined;
            resourceInputs["uri"] = args ? args.uri : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["verifyHostCert"] = args ? args.verifyHostCert : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["alicloudAccessKeySecret", "awsApiKey", "azureApiKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Automationaction.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Automationaction resources.
 */
export interface AutomationactionState {
    /**
     * Access profile for CLI script action to access FortiGate features.
     */
    accprofile?: pulumi.Input<string>;
    /**
     * Action type.
     */
    actionType?: pulumi.Input<string>;
    /**
     * AliCloud AccessKey ID.
     */
    alicloudAccessKeyId?: pulumi.Input<string>;
    /**
     * AliCloud AccessKey secret.
     */
    alicloudAccessKeySecret?: pulumi.Input<string>;
    /**
     * AliCloud account ID.
     */
    alicloudAccountId?: pulumi.Input<string>;
    /**
     * AliCloud function name.
     */
    alicloudFunction?: pulumi.Input<string>;
    /**
     * AliCloud function authorization type. Valid values: `anonymous`, `function`.
     */
    alicloudFunctionAuthorization?: pulumi.Input<string>;
    /**
     * AliCloud function domain.
     */
    alicloudFunctionDomain?: pulumi.Input<string>;
    /**
     * AliCloud region.
     */
    alicloudRegion?: pulumi.Input<string>;
    /**
     * AliCloud service name.
     */
    alicloudService?: pulumi.Input<string>;
    /**
     * AliCloud version.
     */
    alicloudVersion?: pulumi.Input<string>;
    /**
     * AWS API Gateway ID.
     */
    awsApiId?: pulumi.Input<string>;
    /**
     * AWS API Gateway API key.
     */
    awsApiKey?: pulumi.Input<string>;
    /**
     * AWS API Gateway path.
     */
    awsApiPath?: pulumi.Input<string>;
    /**
     * AWS API Gateway deployment stage name.
     */
    awsApiStage?: pulumi.Input<string>;
    /**
     * AWS domain.
     */
    awsDomain?: pulumi.Input<string>;
    /**
     * AWS region.
     */
    awsRegion?: pulumi.Input<string>;
    /**
     * Azure function API key.
     */
    azureApiKey?: pulumi.Input<string>;
    /**
     * Azure function application name.
     */
    azureApp?: pulumi.Input<string>;
    /**
     * Azure function domain.
     */
    azureDomain?: pulumi.Input<string>;
    /**
     * Azure function name.
     */
    azureFunction?: pulumi.Input<string>;
    /**
     * Azure function authorization level. Valid values: `anonymous`, `function`, `admin`.
     */
    azureFunctionAuthorization?: pulumi.Input<string>;
    /**
     * Delay before execution (in seconds).
     */
    delay?: pulumi.Input<number>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Email body.
     */
    emailBody?: pulumi.Input<string>;
    /**
     * Email sender name.
     */
    emailFrom?: pulumi.Input<string>;
    /**
     * Email subject.
     */
    emailSubject?: pulumi.Input<string>;
    /**
     * Email addresses. The structure of `emailTo` block is documented below.
     */
    emailTos?: pulumi.Input<pulumi.Input<inputs.system.AutomationactionEmailTo>[]>;
    /**
     * Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric. Valid values: `enable`, `disable`.
     */
    executeSecurityFabric?: pulumi.Input<string>;
    /**
     * Enable/disable use of your FortiCare email address as the email-to address. Valid values: `enable`, `disable`.
     */
    forticareEmail?: pulumi.Input<string>;
    /**
     * Google Cloud function name.
     */
    gcpFunction?: pulumi.Input<string>;
    /**
     * Google Cloud function domain.
     */
    gcpFunctionDomain?: pulumi.Input<string>;
    /**
     * Google Cloud function region.
     */
    gcpFunctionRegion?: pulumi.Input<string>;
    /**
     * Google Cloud Platform project name.
     */
    gcpProject?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Request headers. The structure of `headers` block is documented below.
     */
    headers?: pulumi.Input<pulumi.Input<inputs.system.AutomationactionHeader>[]>;
    /**
     * Request body (if necessary). Should be serialized json string.
     */
    httpBody?: pulumi.Input<string>;
    /**
     * Request headers. The structure of `httpHeaders` block is documented below.
     */
    httpHeaders?: pulumi.Input<pulumi.Input<inputs.system.AutomationactionHttpHeader>[]>;
    /**
     * Message content.
     */
    message?: pulumi.Input<string>;
    /**
     * Message type. Valid values: `text`, `json`.
     */
    messageType?: pulumi.Input<string>;
    /**
     * Request method (POST, PUT, GET, PATCH or DELETE). Valid values: `post`, `put`, `get`, `patch`, `delete`.
     */
    method?: pulumi.Input<string>;
    /**
     * Limit execution to no more than once in this interval (in seconds).
     */
    minimumInterval?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of megabytes to limit script output to (1 - 1024, default = 10).
     */
    outputSize?: pulumi.Input<number>;
    /**
     * Protocol port.
     */
    port?: pulumi.Input<number>;
    /**
     * Request protocol. Valid values: `http`, `https`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Enable/disable replacement message. Valid values: `enable`, `disable`.
     */
    replacementMessage?: pulumi.Input<string>;
    /**
     * Replacement message group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Required in action chain. Valid values: `enable`, `disable`.
     */
    required?: pulumi.Input<string>;
    /**
     * CLI script.
     */
    script?: pulumi.Input<string>;
    /**
     * NSX SDN connector names. The structure of `sdnConnector` block is documented below.
     */
    sdnConnectors?: pulumi.Input<pulumi.Input<inputs.system.AutomationactionSdnConnector>[]>;
    /**
     * NSX security tag.
     */
    securityTag?: pulumi.Input<string>;
    /**
     * System action type. Valid values: `reboot`, `shutdown`, `backup-config`.
     */
    systemAction?: pulumi.Input<string>;
    /**
     * Maximum running time for this script in seconds (0 = no timeout).
     */
    timeout?: pulumi.Input<number>;
    /**
     * Custom TLS certificate for API request.
     */
    tlsCertificate?: pulumi.Input<string>;
    /**
     * Request API URI.
     */
    uri?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable verification of the remote host certificate. Valid values: `enable`, `disable`.
     */
    verifyHostCert?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Automationaction resource.
 */
export interface AutomationactionArgs {
    /**
     * Access profile for CLI script action to access FortiGate features.
     */
    accprofile?: pulumi.Input<string>;
    /**
     * Action type.
     */
    actionType?: pulumi.Input<string>;
    /**
     * AliCloud AccessKey ID.
     */
    alicloudAccessKeyId?: pulumi.Input<string>;
    /**
     * AliCloud AccessKey secret.
     */
    alicloudAccessKeySecret?: pulumi.Input<string>;
    /**
     * AliCloud account ID.
     */
    alicloudAccountId?: pulumi.Input<string>;
    /**
     * AliCloud function name.
     */
    alicloudFunction?: pulumi.Input<string>;
    /**
     * AliCloud function authorization type. Valid values: `anonymous`, `function`.
     */
    alicloudFunctionAuthorization?: pulumi.Input<string>;
    /**
     * AliCloud function domain.
     */
    alicloudFunctionDomain?: pulumi.Input<string>;
    /**
     * AliCloud region.
     */
    alicloudRegion?: pulumi.Input<string>;
    /**
     * AliCloud service name.
     */
    alicloudService?: pulumi.Input<string>;
    /**
     * AliCloud version.
     */
    alicloudVersion?: pulumi.Input<string>;
    /**
     * AWS API Gateway ID.
     */
    awsApiId?: pulumi.Input<string>;
    /**
     * AWS API Gateway API key.
     */
    awsApiKey?: pulumi.Input<string>;
    /**
     * AWS API Gateway path.
     */
    awsApiPath?: pulumi.Input<string>;
    /**
     * AWS API Gateway deployment stage name.
     */
    awsApiStage?: pulumi.Input<string>;
    /**
     * AWS domain.
     */
    awsDomain?: pulumi.Input<string>;
    /**
     * AWS region.
     */
    awsRegion?: pulumi.Input<string>;
    /**
     * Azure function API key.
     */
    azureApiKey?: pulumi.Input<string>;
    /**
     * Azure function application name.
     */
    azureApp?: pulumi.Input<string>;
    /**
     * Azure function domain.
     */
    azureDomain?: pulumi.Input<string>;
    /**
     * Azure function name.
     */
    azureFunction?: pulumi.Input<string>;
    /**
     * Azure function authorization level. Valid values: `anonymous`, `function`, `admin`.
     */
    azureFunctionAuthorization?: pulumi.Input<string>;
    /**
     * Delay before execution (in seconds).
     */
    delay?: pulumi.Input<number>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Email body.
     */
    emailBody?: pulumi.Input<string>;
    /**
     * Email sender name.
     */
    emailFrom?: pulumi.Input<string>;
    /**
     * Email subject.
     */
    emailSubject?: pulumi.Input<string>;
    /**
     * Email addresses. The structure of `emailTo` block is documented below.
     */
    emailTos?: pulumi.Input<pulumi.Input<inputs.system.AutomationactionEmailTo>[]>;
    /**
     * Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric. Valid values: `enable`, `disable`.
     */
    executeSecurityFabric?: pulumi.Input<string>;
    /**
     * Enable/disable use of your FortiCare email address as the email-to address. Valid values: `enable`, `disable`.
     */
    forticareEmail?: pulumi.Input<string>;
    /**
     * Google Cloud function name.
     */
    gcpFunction?: pulumi.Input<string>;
    /**
     * Google Cloud function domain.
     */
    gcpFunctionDomain?: pulumi.Input<string>;
    /**
     * Google Cloud function region.
     */
    gcpFunctionRegion?: pulumi.Input<string>;
    /**
     * Google Cloud Platform project name.
     */
    gcpProject?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Request headers. The structure of `headers` block is documented below.
     */
    headers?: pulumi.Input<pulumi.Input<inputs.system.AutomationactionHeader>[]>;
    /**
     * Request body (if necessary). Should be serialized json string.
     */
    httpBody?: pulumi.Input<string>;
    /**
     * Request headers. The structure of `httpHeaders` block is documented below.
     */
    httpHeaders?: pulumi.Input<pulumi.Input<inputs.system.AutomationactionHttpHeader>[]>;
    /**
     * Message content.
     */
    message?: pulumi.Input<string>;
    /**
     * Message type. Valid values: `text`, `json`.
     */
    messageType?: pulumi.Input<string>;
    /**
     * Request method (POST, PUT, GET, PATCH or DELETE). Valid values: `post`, `put`, `get`, `patch`, `delete`.
     */
    method?: pulumi.Input<string>;
    /**
     * Limit execution to no more than once in this interval (in seconds).
     */
    minimumInterval?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of megabytes to limit script output to (1 - 1024, default = 10).
     */
    outputSize?: pulumi.Input<number>;
    /**
     * Protocol port.
     */
    port?: pulumi.Input<number>;
    /**
     * Request protocol. Valid values: `http`, `https`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Enable/disable replacement message. Valid values: `enable`, `disable`.
     */
    replacementMessage?: pulumi.Input<string>;
    /**
     * Replacement message group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Required in action chain. Valid values: `enable`, `disable`.
     */
    required?: pulumi.Input<string>;
    /**
     * CLI script.
     */
    script?: pulumi.Input<string>;
    /**
     * NSX SDN connector names. The structure of `sdnConnector` block is documented below.
     */
    sdnConnectors?: pulumi.Input<pulumi.Input<inputs.system.AutomationactionSdnConnector>[]>;
    /**
     * NSX security tag.
     */
    securityTag?: pulumi.Input<string>;
    /**
     * System action type. Valid values: `reboot`, `shutdown`, `backup-config`.
     */
    systemAction?: pulumi.Input<string>;
    /**
     * Maximum running time for this script in seconds (0 = no timeout).
     */
    timeout?: pulumi.Input<number>;
    /**
     * Custom TLS certificate for API request.
     */
    tlsCertificate?: pulumi.Input<string>;
    /**
     * Request API URI.
     */
    uri?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable verification of the remote host certificate. Valid values: `enable`, `disable`.
     */
    verifyHostCert?: pulumi.Input<string>;
}
