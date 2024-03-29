// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a list of `fortios.firewall.Internetserviceextension`.
 */
export function getInternetserviceextensionlist(args?: GetInternetserviceextensionlistArgs, opts?: pulumi.InvokeOptions): Promise<GetInternetserviceextensionlistResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:firewall/getInternetserviceextensionlist:getInternetserviceextensionlist", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getInternetserviceextensionlist.
 */
export interface GetInternetserviceextensionlistArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getInternetserviceextensionlist.
 */
export interface GetInternetserviceextensionlistResult {
    readonly filter?: string;
    /**
     * A list of the `fortios.firewall.Internetserviceextension`.
     */
    readonly fosidlists: number[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly vdomparam?: string;
}
/**
 * Provides a list of `fortios.firewall.Internetserviceextension`.
 */
export function getInternetserviceextensionlistOutput(args?: GetInternetserviceextensionlistOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInternetserviceextensionlistResult> {
    return pulumi.output(args).apply((a: any) => getInternetserviceextensionlist(a, opts))
}

/**
 * A collection of arguments for invoking getInternetserviceextensionlist.
 */
export interface GetInternetserviceextensionlistOutputArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
