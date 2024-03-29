// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios system dnsdatabase
 */
export function getDnsdatabase(args: GetDnsdatabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetDnsdatabaseResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:system/getDnsdatabase:getDnsdatabase", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getDnsdatabase.
 */
export interface GetDnsdatabaseArgs {
    /**
     * Specify the name of the desired system dnsdatabase.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getDnsdatabase.
 */
export interface GetDnsdatabaseResult {
    /**
     * DNS zone transfer IP address list.
     */
    readonly allowTransfer: string;
    /**
     * Enable/disable authoritative zone.
     */
    readonly authoritative: string;
    /**
     * Email address of the administrator for this zone.
     * You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
     * When using a simple username, the domain of the email will be this zone.
     */
    readonly contact: string;
    /**
     * DNS entry. The structure of `dnsEntry` block is documented below.
     */
    readonly dnsEntries: outputs.system.GetDnsdatabaseDnsEntry[];
    /**
     * Domain name.
     */
    readonly domain: string;
    /**
     * DNS zone forwarder IP address list.
     */
    readonly forwarder: string;
    /**
     * Forwarder IPv6 address.
     */
    readonly forwarder6: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
     */
    readonly ipMaster: string;
    /**
     * IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
     */
    readonly ipPrimary: string;
    /**
     * Zone name.
     */
    readonly name: string;
    /**
     * Domain name of the default DNS server for this zone.
     */
    readonly primaryName: string;
    /**
     * Maximum number of resource records (10 - 65536, 0 means infinite).
     */
    readonly rrMax: number;
    /**
     * Source IP for forwarding to DNS server.
     */
    readonly sourceIp: string;
    /**
     * IPv6 source IP address for forwarding to DNS server.
     */
    readonly sourceIp6: string;
    /**
     * Enable/disable resource record status.
     */
    readonly status: string;
    /**
     * Time-to-live for this entry (0 to 2147483647 sec, default = 0).
     */
    readonly ttl: number;
    /**
     * Resource record type.
     */
    readonly type: string;
    readonly vdomparam?: string;
    /**
     * Zone view (public to serve public clients, shadow to serve internal clients).
     */
    readonly view: string;
}
/**
 * Use this data source to get information on an fortios system dnsdatabase
 */
export function getDnsdatabaseOutput(args: GetDnsdatabaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDnsdatabaseResult> {
    return pulumi.output(args).apply((a: any) => getDnsdatabase(a, opts))
}

/**
 * A collection of arguments for invoking getDnsdatabase.
 */
export interface GetDnsdatabaseOutputArgs {
    /**
     * Specify the name of the desired system dnsdatabase.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
