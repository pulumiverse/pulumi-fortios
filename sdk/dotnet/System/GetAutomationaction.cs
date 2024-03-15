// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    public static class GetAutomationaction
    {
        /// <summary>
        /// Use this data source to get information on an fortios system automationaction
        /// </summary>
        public static Task<GetAutomationactionResult> InvokeAsync(GetAutomationactionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAutomationactionResult>("fortios:system/getAutomationaction:getAutomationaction", args ?? new GetAutomationactionArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system automationaction
        /// </summary>
        public static Output<GetAutomationactionResult> Invoke(GetAutomationactionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutomationactionResult>("fortios:system/getAutomationaction:getAutomationaction", args ?? new GetAutomationactionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAutomationactionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system automationaction.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetAutomationactionArgs()
        {
        }
        public static new GetAutomationactionArgs Empty => new GetAutomationactionArgs();
    }

    public sealed class GetAutomationactionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system automationaction.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetAutomationactionInvokeArgs()
        {
        }
        public static new GetAutomationactionInvokeArgs Empty => new GetAutomationactionInvokeArgs();
    }


    [OutputType]
    public sealed class GetAutomationactionResult
    {
        /// <summary>
        /// Access profile for CLI script action to access FortiGate features.
        /// </summary>
        public readonly string Accprofile;
        /// <summary>
        /// Action type.
        /// </summary>
        public readonly string ActionType;
        /// <summary>
        /// AliCloud AccessKey ID.
        /// </summary>
        public readonly string AlicloudAccessKeyId;
        /// <summary>
        /// AliCloud AccessKey secret.
        /// </summary>
        public readonly string AlicloudAccessKeySecret;
        /// <summary>
        /// AliCloud account ID.
        /// </summary>
        public readonly string AlicloudAccountId;
        /// <summary>
        /// AliCloud function name.
        /// </summary>
        public readonly string AlicloudFunction;
        /// <summary>
        /// AliCloud function authorization type.
        /// </summary>
        public readonly string AlicloudFunctionAuthorization;
        /// <summary>
        /// AliCloud function domain.
        /// </summary>
        public readonly string AlicloudFunctionDomain;
        /// <summary>
        /// AliCloud region.
        /// </summary>
        public readonly string AlicloudRegion;
        /// <summary>
        /// AliCloud service name.
        /// </summary>
        public readonly string AlicloudService;
        /// <summary>
        /// AliCloud version.
        /// </summary>
        public readonly string AlicloudVersion;
        /// <summary>
        /// AWS API Gateway ID.
        /// </summary>
        public readonly string AwsApiId;
        /// <summary>
        /// AWS API Gateway API key.
        /// </summary>
        public readonly string AwsApiKey;
        /// <summary>
        /// AWS API Gateway path.
        /// </summary>
        public readonly string AwsApiPath;
        /// <summary>
        /// AWS API Gateway deployment stage name.
        /// </summary>
        public readonly string AwsApiStage;
        /// <summary>
        /// AWS domain.
        /// </summary>
        public readonly string AwsDomain;
        /// <summary>
        /// AWS region.
        /// </summary>
        public readonly string AwsRegion;
        /// <summary>
        /// Azure function API key.
        /// </summary>
        public readonly string AzureApiKey;
        /// <summary>
        /// Azure function application name.
        /// </summary>
        public readonly string AzureApp;
        /// <summary>
        /// Azure function domain.
        /// </summary>
        public readonly string AzureDomain;
        /// <summary>
        /// Azure function name.
        /// </summary>
        public readonly string AzureFunction;
        /// <summary>
        /// Azure function authorization level.
        /// </summary>
        public readonly string AzureFunctionAuthorization;
        /// <summary>
        /// Delay before execution (in seconds).
        /// </summary>
        public readonly int Delay;
        /// <summary>
        /// Description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Email body.
        /// </summary>
        public readonly string EmailBody;
        /// <summary>
        /// Email sender name.
        /// </summary>
        public readonly string EmailFrom;
        /// <summary>
        /// Email subject.
        /// </summary>
        public readonly string EmailSubject;
        /// <summary>
        /// Email addresses. The structure of `email_to` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAutomationactionEmailToResult> EmailTos;
        /// <summary>
        /// Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric.
        /// </summary>
        public readonly string ExecuteSecurityFabric;
        /// <summary>
        /// Google Cloud function name.
        /// </summary>
        public readonly string GcpFunction;
        /// <summary>
        /// Google Cloud function domain.
        /// </summary>
        public readonly string GcpFunctionDomain;
        /// <summary>
        /// Google Cloud function region.
        /// </summary>
        public readonly string GcpFunctionRegion;
        /// <summary>
        /// Google Cloud Platform project name.
        /// </summary>
        public readonly string GcpProject;
        /// <summary>
        /// Request headers. The structure of `headers` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAutomationactionHeaderResult> Headers;
        /// <summary>
        /// Request body (if necessary). Should be serialized json string.
        /// </summary>
        public readonly string HttpBody;
        /// <summary>
        /// Request headers. The structure of `http_headers` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAutomationactionHttpHeaderResult> HttpHeaders;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Message content.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// Message type.
        /// </summary>
        public readonly string MessageType;
        /// <summary>
        /// Request method (POST, PUT, GET, PATCH or DELETE).
        /// </summary>
        public readonly string Method;
        /// <summary>
        /// Limit execution to no more than once in this interval (in seconds).
        /// </summary>
        public readonly int MinimumInterval;
        /// <summary>
        /// SDN connector name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number of megabytes to limit script output to (1 - 1024, default = 10).
        /// </summary>
        public readonly int OutputSize;
        /// <summary>
        /// Protocol port.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Request protocol.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// Enable/disable replacement message.
        /// </summary>
        public readonly string ReplacementMessage;
        /// <summary>
        /// Replacement message group.
        /// </summary>
        public readonly string ReplacemsgGroup;
        /// <summary>
        /// Required in action chain.
        /// </summary>
        public readonly string Required;
        /// <summary>
        /// CLI script.
        /// </summary>
        public readonly string Script;
        /// <summary>
        /// NSX SDN connector names. The structure of `sdn_connector` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAutomationactionSdnConnectorResult> SdnConnectors;
        /// <summary>
        /// NSX security tag.
        /// </summary>
        public readonly string SecurityTag;
        /// <summary>
        /// System action type.
        /// </summary>
        public readonly string SystemAction;
        /// <summary>
        /// Maximum running time for this script in seconds (0 = no timeout).
        /// </summary>
        public readonly int Timeout;
        /// <summary>
        /// Custom TLS certificate for API request.
        /// </summary>
        public readonly string TlsCertificate;
        /// <summary>
        /// Request API URI.
        /// </summary>
        public readonly string Uri;
        public readonly string? Vdomparam;
        /// <summary>
        /// Enable/disable verification of the remote host certificate.
        /// </summary>
        public readonly string VerifyHostCert;

        [OutputConstructor]
        private GetAutomationactionResult(
            string accprofile,

            string actionType,

            string alicloudAccessKeyId,

            string alicloudAccessKeySecret,

            string alicloudAccountId,

            string alicloudFunction,

            string alicloudFunctionAuthorization,

            string alicloudFunctionDomain,

            string alicloudRegion,

            string alicloudService,

            string alicloudVersion,

            string awsApiId,

            string awsApiKey,

            string awsApiPath,

            string awsApiStage,

            string awsDomain,

            string awsRegion,

            string azureApiKey,

            string azureApp,

            string azureDomain,

            string azureFunction,

            string azureFunctionAuthorization,

            int delay,

            string description,

            string emailBody,

            string emailFrom,

            string emailSubject,

            ImmutableArray<Outputs.GetAutomationactionEmailToResult> emailTos,

            string executeSecurityFabric,

            string gcpFunction,

            string gcpFunctionDomain,

            string gcpFunctionRegion,

            string gcpProject,

            ImmutableArray<Outputs.GetAutomationactionHeaderResult> headers,

            string httpBody,

            ImmutableArray<Outputs.GetAutomationactionHttpHeaderResult> httpHeaders,

            string id,

            string message,

            string messageType,

            string method,

            int minimumInterval,

            string name,

            int outputSize,

            int port,

            string protocol,

            string replacementMessage,

            string replacemsgGroup,

            string required,

            string script,

            ImmutableArray<Outputs.GetAutomationactionSdnConnectorResult> sdnConnectors,

            string securityTag,

            string systemAction,

            int timeout,

            string tlsCertificate,

            string uri,

            string? vdomparam,

            string verifyHostCert)
        {
            Accprofile = accprofile;
            ActionType = actionType;
            AlicloudAccessKeyId = alicloudAccessKeyId;
            AlicloudAccessKeySecret = alicloudAccessKeySecret;
            AlicloudAccountId = alicloudAccountId;
            AlicloudFunction = alicloudFunction;
            AlicloudFunctionAuthorization = alicloudFunctionAuthorization;
            AlicloudFunctionDomain = alicloudFunctionDomain;
            AlicloudRegion = alicloudRegion;
            AlicloudService = alicloudService;
            AlicloudVersion = alicloudVersion;
            AwsApiId = awsApiId;
            AwsApiKey = awsApiKey;
            AwsApiPath = awsApiPath;
            AwsApiStage = awsApiStage;
            AwsDomain = awsDomain;
            AwsRegion = awsRegion;
            AzureApiKey = azureApiKey;
            AzureApp = azureApp;
            AzureDomain = azureDomain;
            AzureFunction = azureFunction;
            AzureFunctionAuthorization = azureFunctionAuthorization;
            Delay = delay;
            Description = description;
            EmailBody = emailBody;
            EmailFrom = emailFrom;
            EmailSubject = emailSubject;
            EmailTos = emailTos;
            ExecuteSecurityFabric = executeSecurityFabric;
            GcpFunction = gcpFunction;
            GcpFunctionDomain = gcpFunctionDomain;
            GcpFunctionRegion = gcpFunctionRegion;
            GcpProject = gcpProject;
            Headers = headers;
            HttpBody = httpBody;
            HttpHeaders = httpHeaders;
            Id = id;
            Message = message;
            MessageType = messageType;
            Method = method;
            MinimumInterval = minimumInterval;
            Name = name;
            OutputSize = outputSize;
            Port = port;
            Protocol = protocol;
            ReplacementMessage = replacementMessage;
            ReplacemsgGroup = replacemsgGroup;
            Required = required;
            Script = script;
            SdnConnectors = sdnConnectors;
            SecurityTag = securityTag;
            SystemAction = systemAction;
            Timeout = timeout;
            TlsCertificate = tlsCertificate;
            Uri = uri;
            Vdomparam = vdomparam;
            VerifyHostCert = verifyHostCert;
        }
    }
}