// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios firewall proxyaddress
 */
export function getProxyaddress(args: GetProxyaddressArgs, opts?: pulumi.InvokeOptions): Promise<GetProxyaddressResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:firewall/getProxyaddress:getProxyaddress", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getProxyaddress.
 */
export interface GetProxyaddressArgs {
    /**
     * Specify the name of the desired firewall proxyaddress.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getProxyaddress.
 */
export interface GetProxyaddressResult {
    /**
     * SaaS application. The structure of `application` block is documented below.
     */
    readonly applications: outputs.firewall.GetProxyaddressApplication[];
    /**
     * Case sensitivity in pattern.
     */
    readonly caseSensitivity: string;
    /**
     * Tag category.
     */
    readonly categories: outputs.firewall.GetProxyaddressCategory[];
    /**
     * Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
     */
    readonly color: number;
    /**
     * Optional comments.
     */
    readonly comment: string;
    /**
     * HTTP header regular expression.
     */
    readonly header: string;
    /**
     * HTTP header group. The structure of `headerGroup` block is documented below.
     */
    readonly headerGroups: outputs.firewall.GetProxyaddressHeaderGroup[];
    /**
     * HTTP header.
     */
    readonly headerName: string;
    /**
     * Address object for the host.
     */
    readonly host: string;
    /**
     * Host name as a regular expression.
     */
    readonly hostRegex: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * HTTP request methods to be used.
     */
    readonly method: string;
    /**
     * SaaS applicaton name.
     */
    readonly name: string;
    /**
     * URL path as a regular expression.
     */
    readonly path: string;
    /**
     * Match the query part of the URL as a regular expression.
     */
    readonly query: string;
    /**
     * Enable/disable use of referrer field in the HTTP header to match the address.
     */
    readonly referrer: string;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    readonly taggings: outputs.firewall.GetProxyaddressTagging[];
    /**
     * Proxy address type.
     */
    readonly type: string;
    /**
     * Names of browsers to be used as user agent.
     */
    readonly ua: string;
    /**
     * Maximum version of the user agent specified in dotted notation. For example, use 120 with the ua field set to "chrome" to require Google Chrome's maximum version must be 120.
     */
    readonly uaMaxVer: string;
    /**
     * Minimum version of the user agent specified in dotted notation. For example, use 90.0.1 with the ua field set to "chrome" to require Google Chrome's minimum version must be 90.0.1.
     */
    readonly uaMinVer: string;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    readonly uuid: string;
    readonly vdomparam?: string;
    /**
     * Enable/disable visibility of the object in the GUI.
     */
    readonly visibility: string;
}
/**
 * Use this data source to get information on an fortios firewall proxyaddress
 */
export function getProxyaddressOutput(args: GetProxyaddressOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProxyaddressResult> {
    return pulumi.output(args).apply((a: any) => getProxyaddress(a, opts))
}

/**
 * A collection of arguments for invoking getProxyaddress.
 */
export interface GetProxyaddressOutputArgs {
    /**
     * Specify the name of the desired firewall proxyaddress.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
