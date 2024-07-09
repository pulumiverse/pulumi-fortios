// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on fortios system csf
 */
export function getCsf(args?: GetCsfArgs, opts?: pulumi.InvokeOptions): Promise<GetCsfResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:system/getCsf:getCsf", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getCsf.
 */
export interface GetCsfArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getCsf.
 */
export interface GetCsfResult {
    /**
     * Accept connections with unknown certificates and ask admin for approval.
     */
    readonly acceptAuthByCert: string;
    /**
     * Authorization request type.
     */
    readonly authorizationRequestType: string;
    /**
     * Certificate.
     */
    readonly certificate: string;
    /**
     * Configuration sync mode.
     */
    readonly configurationSync: string;
    /**
     * Enable/disable downstream device access to this device's configuration and data.
     */
    readonly downstreamAccess: string;
    /**
     * Default access profile for requests from downstream devices.
     */
    readonly downstreamAccprofile: string;
    /**
     * Fabric connector configuration. The structure of `fabricConnector` block is documented below.
     */
    readonly fabricConnectors: outputs.system.GetCsfFabricConnector[];
    /**
     * Fabric device configuration. The structure of `fabricDevice` block is documented below.
     */
    readonly fabricDevices: outputs.system.GetCsfFabricDevice[];
    /**
     * Fabric CMDB Object Unification
     */
    readonly fabricObjectUnification: string;
    /**
     * Number of worker processes for Security Fabric daemon.
     */
    readonly fabricWorkers: number;
    /**
     * Enable/disable Security Fabric daemon file management.
     */
    readonly fileMgmt: string;
    /**
     * Maximum amount of memory that can be used by the daemon files (in bytes).
     */
    readonly fileQuota: number;
    /**
     * Warn when the set percentage of quota has been used.
     */
    readonly fileQuotaWarning: number;
    /**
     * Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
     */
    readonly fixedKey: string;
    /**
     * Fabric FortiCloud account unification.
     */
    readonly forticloudAccountEnforcement: string;
    /**
     * Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
     */
    readonly groupName: string;
    /**
     * Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
     */
    readonly groupPassword: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Enable/disable broadcast of discovery messages for log unification.
     */
    readonly logUnification: string;
    /**
     * Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
     */
    readonly managementIp: string;
    /**
     * Overriding port for management connection (Overrides admin port).
     */
    readonly managementPort: number;
    /**
     * SAML setting configuration synchronization.
     */
    readonly samlConfigurationSync: string;
    /**
     * Source IP address for communication with the upstream FortiGate.
     */
    readonly sourceIp: string;
    /**
     * Enable/disable Security Fabric.
     */
    readonly status: string;
    /**
     * Pre-authorized and blocked security fabric nodes. The structure of `trustedList` block is documented below.
     */
    readonly trustedLists: outputs.system.GetCsfTrustedList[];
    /**
     * Unique ID of the current CSF node
     */
    readonly uid: string;
    /**
     * IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
     */
    readonly upstream: string;
    /**
     * Specify outgoing interface to reach server.
     */
    readonly upstreamInterface: string;
    /**
     * Specify how to select outgoing interface to reach server.
     */
    readonly upstreamInterfaceSelectMethod: string;
    /**
     * IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
     */
    readonly upstreamIp: string;
    /**
     * The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
     */
    readonly upstreamPort: number;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on fortios system csf
 */
export function getCsfOutput(args?: GetCsfOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCsfResult> {
    return pulumi.output(args).apply((a: any) => getCsf(a, opts))
}

/**
 * A collection of arguments for invoking getCsf.
 */
export interface GetCsfOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
