// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure IPv6 access proxy. Applies to FortiOS Version `>= 7.0.1`.
 *
 * ## Import
 *
 * Firewall AccessProxy6 can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallAccessproxy6:FirewallAccessproxy6 labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallAccessproxy6:FirewallAccessproxy6 labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallAccessproxy6 extends pulumi.CustomResource {
    /**
     * Get an existing FirewallAccessproxy6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallAccessproxy6State, opts?: pulumi.CustomResourceOptions): FirewallAccessproxy6 {
        return new FirewallAccessproxy6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallAccessproxy6:FirewallAccessproxy6';

    /**
     * Returns true if the given object is an instance of FirewallAccessproxy6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallAccessproxy6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallAccessproxy6.__pulumiType;
    }

    /**
     * Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
     */
    public readonly addVhostDomainToDnsdb!: pulumi.Output<string>;
    /**
     * Set IPv6 API Gateway. The structure of `apiGateway6` block is documented below.
     */
    public readonly apiGateway6s!: pulumi.Output<outputs.FirewallAccessproxy6ApiGateway6[] | undefined>;
    /**
     * Set IPv4 API Gateway. The structure of `apiGateway` block is documented below.
     */
    public readonly apiGateways!: pulumi.Output<outputs.FirewallAccessproxy6ApiGateway[] | undefined>;
    /**
     * Enable/disable authentication portal. Valid values: `disable`, `enable`.
     */
    public readonly authPortal!: pulumi.Output<string>;
    /**
     * Virtual host for authentication portal.
     */
    public readonly authVirtualHost!: pulumi.Output<string>;
    /**
     * Enable/disable to request client certificate. Valid values: `disable`, `enable`.
     */
    public readonly clientCert!: pulumi.Output<string>;
    /**
     * Decrypted traffic mirror.
     */
    public readonly decryptedTrafficMirror!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Action of an empty client certificate.
     */
    public readonly emptyCertAction!: pulumi.Output<string>;
    /**
     * Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
     */
    public readonly logBlockedTraffic!: pulumi.Output<string>;
    /**
     * Access Proxy name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
     */
    public readonly userAgentDetect!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Virtual IP name.
     */
    public readonly vip!: pulumi.Output<string>;

    /**
     * Create a FirewallAccessproxy6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallAccessproxy6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallAccessproxy6Args | FirewallAccessproxy6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallAccessproxy6State | undefined;
            resourceInputs["addVhostDomainToDnsdb"] = state ? state.addVhostDomainToDnsdb : undefined;
            resourceInputs["apiGateway6s"] = state ? state.apiGateway6s : undefined;
            resourceInputs["apiGateways"] = state ? state.apiGateways : undefined;
            resourceInputs["authPortal"] = state ? state.authPortal : undefined;
            resourceInputs["authVirtualHost"] = state ? state.authVirtualHost : undefined;
            resourceInputs["clientCert"] = state ? state.clientCert : undefined;
            resourceInputs["decryptedTrafficMirror"] = state ? state.decryptedTrafficMirror : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["emptyCertAction"] = state ? state.emptyCertAction : undefined;
            resourceInputs["logBlockedTraffic"] = state ? state.logBlockedTraffic : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["userAgentDetect"] = state ? state.userAgentDetect : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vip"] = state ? state.vip : undefined;
        } else {
            const args = argsOrState as FirewallAccessproxy6Args | undefined;
            resourceInputs["addVhostDomainToDnsdb"] = args ? args.addVhostDomainToDnsdb : undefined;
            resourceInputs["apiGateway6s"] = args ? args.apiGateway6s : undefined;
            resourceInputs["apiGateways"] = args ? args.apiGateways : undefined;
            resourceInputs["authPortal"] = args ? args.authPortal : undefined;
            resourceInputs["authVirtualHost"] = args ? args.authVirtualHost : undefined;
            resourceInputs["clientCert"] = args ? args.clientCert : undefined;
            resourceInputs["decryptedTrafficMirror"] = args ? args.decryptedTrafficMirror : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["emptyCertAction"] = args ? args.emptyCertAction : undefined;
            resourceInputs["logBlockedTraffic"] = args ? args.logBlockedTraffic : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["userAgentDetect"] = args ? args.userAgentDetect : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vip"] = args ? args.vip : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallAccessproxy6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallAccessproxy6 resources.
 */
export interface FirewallAccessproxy6State {
    /**
     * Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
     */
    addVhostDomainToDnsdb?: pulumi.Input<string>;
    /**
     * Set IPv6 API Gateway. The structure of `apiGateway6` block is documented below.
     */
    apiGateway6s?: pulumi.Input<pulumi.Input<inputs.FirewallAccessproxy6ApiGateway6>[]>;
    /**
     * Set IPv4 API Gateway. The structure of `apiGateway` block is documented below.
     */
    apiGateways?: pulumi.Input<pulumi.Input<inputs.FirewallAccessproxy6ApiGateway>[]>;
    /**
     * Enable/disable authentication portal. Valid values: `disable`, `enable`.
     */
    authPortal?: pulumi.Input<string>;
    /**
     * Virtual host for authentication portal.
     */
    authVirtualHost?: pulumi.Input<string>;
    /**
     * Enable/disable to request client certificate. Valid values: `disable`, `enable`.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * Decrypted traffic mirror.
     */
    decryptedTrafficMirror?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Action of an empty client certificate.
     */
    emptyCertAction?: pulumi.Input<string>;
    /**
     * Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
     */
    logBlockedTraffic?: pulumi.Input<string>;
    /**
     * Access Proxy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
     */
    userAgentDetect?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Virtual IP name.
     */
    vip?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallAccessproxy6 resource.
 */
export interface FirewallAccessproxy6Args {
    /**
     * Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
     */
    addVhostDomainToDnsdb?: pulumi.Input<string>;
    /**
     * Set IPv6 API Gateway. The structure of `apiGateway6` block is documented below.
     */
    apiGateway6s?: pulumi.Input<pulumi.Input<inputs.FirewallAccessproxy6ApiGateway6>[]>;
    /**
     * Set IPv4 API Gateway. The structure of `apiGateway` block is documented below.
     */
    apiGateways?: pulumi.Input<pulumi.Input<inputs.FirewallAccessproxy6ApiGateway>[]>;
    /**
     * Enable/disable authentication portal. Valid values: `disable`, `enable`.
     */
    authPortal?: pulumi.Input<string>;
    /**
     * Virtual host for authentication portal.
     */
    authVirtualHost?: pulumi.Input<string>;
    /**
     * Enable/disable to request client certificate. Valid values: `disable`, `enable`.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * Decrypted traffic mirror.
     */
    decryptedTrafficMirror?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Action of an empty client certificate.
     */
    emptyCertAction?: pulumi.Input<string>;
    /**
     * Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
     */
    logBlockedTraffic?: pulumi.Input<string>;
    /**
     * Access Proxy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
     */
    userAgentDetect?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Virtual IP name.
     */
    vip?: pulumi.Input<string>;
}
