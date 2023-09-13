// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios system automationaction
 */
export function getAutomationaction(args: GetAutomationactionArgs, opts?: pulumi.InvokeOptions): Promise<GetAutomationactionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:sys/getAutomationaction:getAutomationaction", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getAutomationaction.
 */
export interface GetAutomationactionArgs {
    /**
     * Specify the name of the desired system automationaction.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getAutomationaction.
 */
export interface GetAutomationactionResult {
    /**
     * Access profile for CLI script action to access FortiGate features.
     */
    readonly accprofile: string;
    /**
     * Action type.
     */
    readonly actionType: string;
    /**
     * AliCloud AccessKey ID.
     */
    readonly alicloudAccessKeyId: string;
    /**
     * AliCloud AccessKey secret.
     */
    readonly alicloudAccessKeySecret: string;
    /**
     * AliCloud account ID.
     */
    readonly alicloudAccountId: string;
    /**
     * AliCloud function name.
     */
    readonly alicloudFunction: string;
    /**
     * AliCloud function authorization type.
     */
    readonly alicloudFunctionAuthorization: string;
    /**
     * AliCloud function domain.
     */
    readonly alicloudFunctionDomain: string;
    /**
     * AliCloud region.
     */
    readonly alicloudRegion: string;
    /**
     * AliCloud service name.
     */
    readonly alicloudService: string;
    /**
     * AliCloud version.
     */
    readonly alicloudVersion: string;
    /**
     * AWS API Gateway ID.
     */
    readonly awsApiId: string;
    /**
     * AWS API Gateway API key.
     */
    readonly awsApiKey: string;
    /**
     * AWS API Gateway path.
     */
    readonly awsApiPath: string;
    /**
     * AWS API Gateway deployment stage name.
     */
    readonly awsApiStage: string;
    /**
     * AWS domain.
     */
    readonly awsDomain: string;
    /**
     * AWS region.
     */
    readonly awsRegion: string;
    /**
     * Azure function API key.
     */
    readonly azureApiKey: string;
    /**
     * Azure function application name.
     */
    readonly azureApp: string;
    /**
     * Azure function domain.
     */
    readonly azureDomain: string;
    /**
     * Azure function name.
     */
    readonly azureFunction: string;
    /**
     * Azure function authorization level.
     */
    readonly azureFunctionAuthorization: string;
    /**
     * Delay before execution (in seconds).
     */
    readonly delay: number;
    /**
     * Description.
     */
    readonly description: string;
    /**
     * Email body.
     */
    readonly emailBody: string;
    /**
     * Email sender name.
     */
    readonly emailFrom: string;
    /**
     * Email subject.
     */
    readonly emailSubject: string;
    /**
     * Email addresses. The structure of `emailTo` block is documented below.
     */
    readonly emailTos: outputs.sys.GetAutomationactionEmailTo[];
    /**
     * Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric.
     */
    readonly executeSecurityFabric: string;
    /**
     * Google Cloud function name.
     */
    readonly gcpFunction: string;
    /**
     * Google Cloud function domain.
     */
    readonly gcpFunctionDomain: string;
    /**
     * Google Cloud function region.
     */
    readonly gcpFunctionRegion: string;
    /**
     * Google Cloud Platform project name.
     */
    readonly gcpProject: string;
    /**
     * Request headers. The structure of `headers` block is documented below.
     */
    readonly headers: outputs.sys.GetAutomationactionHeader[];
    /**
     * Request body (if necessary). Should be serialized json string.
     */
    readonly httpBody: string;
    /**
     * Request headers. The structure of `httpHeaders` block is documented below.
     */
    readonly httpHeaders: outputs.sys.GetAutomationactionHttpHeader[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Message content.
     */
    readonly message: string;
    /**
     * Message type.
     */
    readonly messageType: string;
    /**
     * Request method (POST, PUT, GET, PATCH or DELETE).
     */
    readonly method: string;
    /**
     * Limit execution to no more than once in this interval (in seconds).
     */
    readonly minimumInterval: number;
    /**
     * SDN connector name.
     */
    readonly name: string;
    /**
     * Number of megabytes to limit script output to (1 - 1024, default = 10).
     */
    readonly outputSize: number;
    /**
     * Protocol port.
     */
    readonly port: number;
    /**
     * Request protocol.
     */
    readonly protocol: string;
    /**
     * Enable/disable replacement message.
     */
    readonly replacementMessage: string;
    /**
     * Replacement message group.
     */
    readonly replacemsgGroup: string;
    /**
     * Required in action chain.
     */
    readonly required: string;
    /**
     * CLI script.
     */
    readonly script: string;
    /**
     * NSX SDN connector names. The structure of `sdnConnector` block is documented below.
     */
    readonly sdnConnectors: outputs.sys.GetAutomationactionSdnConnector[];
    /**
     * NSX security tag.
     */
    readonly securityTag: string;
    /**
     * System action type.
     */
    readonly systemAction: string;
    /**
     * Maximum running time for this script in seconds (0 = no timeout).
     */
    readonly timeout: number;
    /**
     * Custom TLS certificate for API request.
     */
    readonly tlsCertificate: string;
    /**
     * Request API URI.
     */
    readonly uri: string;
    readonly vdomparam?: string;
    /**
     * Enable/disable verification of the remote host certificate.
     */
    readonly verifyHostCert: string;
}
/**
 * Use this data source to get information on an fortios system automationaction
 */
export function getAutomationactionOutput(args: GetAutomationactionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAutomationactionResult> {
    return pulumi.output(args).apply((a: any) => getAutomationaction(a, opts))
}

/**
 * A collection of arguments for invoking getAutomationaction.
 */
export interface GetAutomationactionOutputArgs {
    /**
     * Specify the name of the desired system automationaction.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
