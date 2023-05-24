// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("fortios");

/**
 * CA Bundle file content
 */
export declare const cabundlecontent: string | undefined;
Object.defineProperty(exports, "cabundlecontent", {
    get() {
        return __config.get("cabundlecontent");
    },
    enumerable: true,
});

/**
 * CA Bundle file
 */
export declare const cabundlefile: string | undefined;
Object.defineProperty(exports, "cabundlefile", {
    get() {
        return __config.get("cabundlefile");
    },
    enumerable: true,
});

/**
 * CA certtificate(Optional)
 */
export declare const cacert: string | undefined;
Object.defineProperty(exports, "cacert", {
    get() {
        return __config.get("cacert");
    },
    enumerable: true,
});

/**
 * User certificate
 */
export declare const clientcert: string | undefined;
Object.defineProperty(exports, "clientcert", {
    get() {
        return __config.get("clientcert");
    },
    enumerable: true,
});

/**
 * User private key
 */
export declare const clientkey: string | undefined;
Object.defineProperty(exports, "clientkey", {
    get() {
        return __config.get("clientkey");
    },
    enumerable: true,
});

/**
 * CA Bundle file
 */
export declare const fmgCabundlefile: string | undefined;
Object.defineProperty(exports, "fmgCabundlefile", {
    get() {
        return __config.get("fmgCabundlefile");
    },
    enumerable: true,
});

/**
 * Hostname/IP address of the FortiManager to connect to
 */
export declare const fmgHostname: string | undefined;
Object.defineProperty(exports, "fmgHostname", {
    get() {
        return __config.get("fmgHostname");
    },
    enumerable: true,
});

export declare const fmgInsecure: boolean | undefined;
Object.defineProperty(exports, "fmgInsecure", {
    get() {
        return __config.getObject<boolean>("fmgInsecure");
    },
    enumerable: true,
});

export declare const fmgPasswd: string | undefined;
Object.defineProperty(exports, "fmgPasswd", {
    get() {
        return __config.get("fmgPasswd");
    },
    enumerable: true,
});

export declare const fmgUsername: string | undefined;
Object.defineProperty(exports, "fmgUsername", {
    get() {
        return __config.get("fmgUsername");
    },
    enumerable: true,
});

/**
 * The hostname/IP address of the FortiOS to be connected
 */
export declare const hostname: string | undefined;
Object.defineProperty(exports, "hostname", {
    get() {
        return __config.get("hostname");
    },
    enumerable: true,
});

/**
 * HTTP proxy address
 */
export declare const httpProxy: string | undefined;
Object.defineProperty(exports, "httpProxy", {
    get() {
        return __config.get("httpProxy");
    },
    enumerable: true,
});

export declare const insecure: boolean | undefined;
Object.defineProperty(exports, "insecure", {
    get() {
        return __config.getObject<boolean>("insecure");
    },
    enumerable: true,
});

/**
 * Enable/disable peer authentication, can be 'enable' or 'disable'
 */
export declare const peerauth: string | undefined;
Object.defineProperty(exports, "peerauth", {
    get() {
        return __config.get("peerauth");
    },
    enumerable: true,
});

export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token");
    },
    enumerable: true,
});

export declare const vdom: string | undefined;
Object.defineProperty(exports, "vdom", {
    get() {
        return __config.get("vdom");
    },
    enumerable: true,
});
