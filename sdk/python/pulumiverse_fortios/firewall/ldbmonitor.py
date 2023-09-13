# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LdbmonitorArgs', 'Ldbmonitor']

@pulumi.input_type
class LdbmonitorArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 dns_match_ip: Optional[pulumi.Input[str]] = None,
                 dns_protocol: Optional[pulumi.Input[str]] = None,
                 dns_request_domain: Optional[pulumi.Input[str]] = None,
                 http_get: Optional[pulumi.Input[str]] = None,
                 http_match: Optional[pulumi.Input[str]] = None,
                 http_max_redirects: Optional[pulumi.Input[int]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 retry: Optional[pulumi.Input[int]] = None,
                 src_ip: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Ldbmonitor resource.
        :param pulumi.Input[str] type: Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP).
        :param pulumi.Input[str] dns_match_ip: Response IP expected from DNS server.
        :param pulumi.Input[str] dns_protocol: Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
        :param pulumi.Input[str] dns_request_domain: Fully qualified domain name to resolve for the DNS probe.
        :param pulumi.Input[str] http_get: URL used to send a GET request to check the health of an HTTP server.
        :param pulumi.Input[str] http_match: String to match the value expected in response to an HTTP-GET request.
        :param pulumi.Input[int] http_max_redirects: The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
        :param pulumi.Input[int] interval: Time between health checks (5 - 65635 sec, default = 10).
        :param pulumi.Input[str] name: Monitor name.
        :param pulumi.Input[int] port: Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65635, default = 0).
        :param pulumi.Input[int] retry: Number health check attempts before the server is considered down (1 - 255, default = 3).
        :param pulumi.Input[str] src_ip: Source IP for ldb-monitor.
        :param pulumi.Input[int] timeout: Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "type", type)
        if dns_match_ip is not None:
            pulumi.set(__self__, "dns_match_ip", dns_match_ip)
        if dns_protocol is not None:
            pulumi.set(__self__, "dns_protocol", dns_protocol)
        if dns_request_domain is not None:
            pulumi.set(__self__, "dns_request_domain", dns_request_domain)
        if http_get is not None:
            pulumi.set(__self__, "http_get", http_get)
        if http_match is not None:
            pulumi.set(__self__, "http_match", http_match)
        if http_max_redirects is not None:
            pulumi.set(__self__, "http_max_redirects", http_max_redirects)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if retry is not None:
            pulumi.set(__self__, "retry", retry)
        if src_ip is not None:
            pulumi.set(__self__, "src_ip", src_ip)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="dnsMatchIp")
    def dns_match_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Response IP expected from DNS server.
        """
        return pulumi.get(self, "dns_match_ip")

    @dns_match_ip.setter
    def dns_match_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_match_ip", value)

    @property
    @pulumi.getter(name="dnsProtocol")
    def dns_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
        """
        return pulumi.get(self, "dns_protocol")

    @dns_protocol.setter
    def dns_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_protocol", value)

    @property
    @pulumi.getter(name="dnsRequestDomain")
    def dns_request_domain(self) -> Optional[pulumi.Input[str]]:
        """
        Fully qualified domain name to resolve for the DNS probe.
        """
        return pulumi.get(self, "dns_request_domain")

    @dns_request_domain.setter
    def dns_request_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_request_domain", value)

    @property
    @pulumi.getter(name="httpGet")
    def http_get(self) -> Optional[pulumi.Input[str]]:
        """
        URL used to send a GET request to check the health of an HTTP server.
        """
        return pulumi.get(self, "http_get")

    @http_get.setter
    def http_get(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_get", value)

    @property
    @pulumi.getter(name="httpMatch")
    def http_match(self) -> Optional[pulumi.Input[str]]:
        """
        String to match the value expected in response to an HTTP-GET request.
        """
        return pulumi.get(self, "http_match")

    @http_match.setter
    def http_match(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_match", value)

    @property
    @pulumi.getter(name="httpMaxRedirects")
    def http_max_redirects(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
        """
        return pulumi.get(self, "http_max_redirects")

    @http_max_redirects.setter
    def http_max_redirects(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "http_max_redirects", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        Time between health checks (5 - 65635 sec, default = 10).
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Monitor name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65635, default = 0).
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def retry(self) -> Optional[pulumi.Input[int]]:
        """
        Number health check attempts before the server is considered down (1 - 255, default = 3).
        """
        return pulumi.get(self, "retry")

    @retry.setter
    def retry(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retry", value)

    @property
    @pulumi.getter(name="srcIp")
    def src_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IP for ldb-monitor.
        """
        return pulumi.get(self, "src_ip")

    @src_ip.setter
    def src_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src_ip", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _LdbmonitorState:
    def __init__(__self__, *,
                 dns_match_ip: Optional[pulumi.Input[str]] = None,
                 dns_protocol: Optional[pulumi.Input[str]] = None,
                 dns_request_domain: Optional[pulumi.Input[str]] = None,
                 http_get: Optional[pulumi.Input[str]] = None,
                 http_match: Optional[pulumi.Input[str]] = None,
                 http_max_redirects: Optional[pulumi.Input[int]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 retry: Optional[pulumi.Input[int]] = None,
                 src_ip: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ldbmonitor resources.
        :param pulumi.Input[str] dns_match_ip: Response IP expected from DNS server.
        :param pulumi.Input[str] dns_protocol: Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
        :param pulumi.Input[str] dns_request_domain: Fully qualified domain name to resolve for the DNS probe.
        :param pulumi.Input[str] http_get: URL used to send a GET request to check the health of an HTTP server.
        :param pulumi.Input[str] http_match: String to match the value expected in response to an HTTP-GET request.
        :param pulumi.Input[int] http_max_redirects: The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
        :param pulumi.Input[int] interval: Time between health checks (5 - 65635 sec, default = 10).
        :param pulumi.Input[str] name: Monitor name.
        :param pulumi.Input[int] port: Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65635, default = 0).
        :param pulumi.Input[int] retry: Number health check attempts before the server is considered down (1 - 255, default = 3).
        :param pulumi.Input[str] src_ip: Source IP for ldb-monitor.
        :param pulumi.Input[int] timeout: Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
        :param pulumi.Input[str] type: Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dns_match_ip is not None:
            pulumi.set(__self__, "dns_match_ip", dns_match_ip)
        if dns_protocol is not None:
            pulumi.set(__self__, "dns_protocol", dns_protocol)
        if dns_request_domain is not None:
            pulumi.set(__self__, "dns_request_domain", dns_request_domain)
        if http_get is not None:
            pulumi.set(__self__, "http_get", http_get)
        if http_match is not None:
            pulumi.set(__self__, "http_match", http_match)
        if http_max_redirects is not None:
            pulumi.set(__self__, "http_max_redirects", http_max_redirects)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if retry is not None:
            pulumi.set(__self__, "retry", retry)
        if src_ip is not None:
            pulumi.set(__self__, "src_ip", src_ip)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dnsMatchIp")
    def dns_match_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Response IP expected from DNS server.
        """
        return pulumi.get(self, "dns_match_ip")

    @dns_match_ip.setter
    def dns_match_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_match_ip", value)

    @property
    @pulumi.getter(name="dnsProtocol")
    def dns_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
        """
        return pulumi.get(self, "dns_protocol")

    @dns_protocol.setter
    def dns_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_protocol", value)

    @property
    @pulumi.getter(name="dnsRequestDomain")
    def dns_request_domain(self) -> Optional[pulumi.Input[str]]:
        """
        Fully qualified domain name to resolve for the DNS probe.
        """
        return pulumi.get(self, "dns_request_domain")

    @dns_request_domain.setter
    def dns_request_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_request_domain", value)

    @property
    @pulumi.getter(name="httpGet")
    def http_get(self) -> Optional[pulumi.Input[str]]:
        """
        URL used to send a GET request to check the health of an HTTP server.
        """
        return pulumi.get(self, "http_get")

    @http_get.setter
    def http_get(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_get", value)

    @property
    @pulumi.getter(name="httpMatch")
    def http_match(self) -> Optional[pulumi.Input[str]]:
        """
        String to match the value expected in response to an HTTP-GET request.
        """
        return pulumi.get(self, "http_match")

    @http_match.setter
    def http_match(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_match", value)

    @property
    @pulumi.getter(name="httpMaxRedirects")
    def http_max_redirects(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
        """
        return pulumi.get(self, "http_max_redirects")

    @http_max_redirects.setter
    def http_max_redirects(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "http_max_redirects", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        Time between health checks (5 - 65635 sec, default = 10).
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Monitor name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65635, default = 0).
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def retry(self) -> Optional[pulumi.Input[int]]:
        """
        Number health check attempts before the server is considered down (1 - 255, default = 3).
        """
        return pulumi.get(self, "retry")

    @retry.setter
    def retry(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retry", value)

    @property
    @pulumi.getter(name="srcIp")
    def src_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IP for ldb-monitor.
        """
        return pulumi.get(self, "src_ip")

    @src_ip.setter
    def src_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src_ip", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class Ldbmonitor(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_match_ip: Optional[pulumi.Input[str]] = None,
                 dns_protocol: Optional[pulumi.Input[str]] = None,
                 dns_request_domain: Optional[pulumi.Input[str]] = None,
                 http_get: Optional[pulumi.Input[str]] = None,
                 http_match: Optional[pulumi.Input[str]] = None,
                 http_max_redirects: Optional[pulumi.Input[int]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 retry: Optional[pulumi.Input[int]] = None,
                 src_ip: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure server load balancing health monitors.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.firewall.Ldbmonitor("trname",
            http_max_redirects=0,
            interval=10,
            port=0,
            retry=3,
            timeout=2,
            type="ping")
        ```

        ## Import

        Firewall LdbMonitor can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:firewall/ldbmonitor:Ldbmonitor labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:firewall/ldbmonitor:Ldbmonitor labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dns_match_ip: Response IP expected from DNS server.
        :param pulumi.Input[str] dns_protocol: Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
        :param pulumi.Input[str] dns_request_domain: Fully qualified domain name to resolve for the DNS probe.
        :param pulumi.Input[str] http_get: URL used to send a GET request to check the health of an HTTP server.
        :param pulumi.Input[str] http_match: String to match the value expected in response to an HTTP-GET request.
        :param pulumi.Input[int] http_max_redirects: The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
        :param pulumi.Input[int] interval: Time between health checks (5 - 65635 sec, default = 10).
        :param pulumi.Input[str] name: Monitor name.
        :param pulumi.Input[int] port: Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65635, default = 0).
        :param pulumi.Input[int] retry: Number health check attempts before the server is considered down (1 - 255, default = 3).
        :param pulumi.Input[str] src_ip: Source IP for ldb-monitor.
        :param pulumi.Input[int] timeout: Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
        :param pulumi.Input[str] type: Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LdbmonitorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure server load balancing health monitors.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.firewall.Ldbmonitor("trname",
            http_max_redirects=0,
            interval=10,
            port=0,
            retry=3,
            timeout=2,
            type="ping")
        ```

        ## Import

        Firewall LdbMonitor can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:firewall/ldbmonitor:Ldbmonitor labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:firewall/ldbmonitor:Ldbmonitor labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param LdbmonitorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LdbmonitorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_match_ip: Optional[pulumi.Input[str]] = None,
                 dns_protocol: Optional[pulumi.Input[str]] = None,
                 dns_request_domain: Optional[pulumi.Input[str]] = None,
                 http_get: Optional[pulumi.Input[str]] = None,
                 http_match: Optional[pulumi.Input[str]] = None,
                 http_max_redirects: Optional[pulumi.Input[int]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 retry: Optional[pulumi.Input[int]] = None,
                 src_ip: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LdbmonitorArgs.__new__(LdbmonitorArgs)

            __props__.__dict__["dns_match_ip"] = dns_match_ip
            __props__.__dict__["dns_protocol"] = dns_protocol
            __props__.__dict__["dns_request_domain"] = dns_request_domain
            __props__.__dict__["http_get"] = http_get
            __props__.__dict__["http_match"] = http_match
            __props__.__dict__["http_max_redirects"] = http_max_redirects
            __props__.__dict__["interval"] = interval
            __props__.__dict__["name"] = name
            __props__.__dict__["port"] = port
            __props__.__dict__["retry"] = retry
            __props__.__dict__["src_ip"] = src_ip
            __props__.__dict__["timeout"] = timeout
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["vdomparam"] = vdomparam
        super(Ldbmonitor, __self__).__init__(
            'fortios:firewall/ldbmonitor:Ldbmonitor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dns_match_ip: Optional[pulumi.Input[str]] = None,
            dns_protocol: Optional[pulumi.Input[str]] = None,
            dns_request_domain: Optional[pulumi.Input[str]] = None,
            http_get: Optional[pulumi.Input[str]] = None,
            http_match: Optional[pulumi.Input[str]] = None,
            http_max_redirects: Optional[pulumi.Input[int]] = None,
            interval: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            retry: Optional[pulumi.Input[int]] = None,
            src_ip: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Ldbmonitor':
        """
        Get an existing Ldbmonitor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dns_match_ip: Response IP expected from DNS server.
        :param pulumi.Input[str] dns_protocol: Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
        :param pulumi.Input[str] dns_request_domain: Fully qualified domain name to resolve for the DNS probe.
        :param pulumi.Input[str] http_get: URL used to send a GET request to check the health of an HTTP server.
        :param pulumi.Input[str] http_match: String to match the value expected in response to an HTTP-GET request.
        :param pulumi.Input[int] http_max_redirects: The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
        :param pulumi.Input[int] interval: Time between health checks (5 - 65635 sec, default = 10).
        :param pulumi.Input[str] name: Monitor name.
        :param pulumi.Input[int] port: Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65635, default = 0).
        :param pulumi.Input[int] retry: Number health check attempts before the server is considered down (1 - 255, default = 3).
        :param pulumi.Input[str] src_ip: Source IP for ldb-monitor.
        :param pulumi.Input[int] timeout: Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
        :param pulumi.Input[str] type: Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LdbmonitorState.__new__(_LdbmonitorState)

        __props__.__dict__["dns_match_ip"] = dns_match_ip
        __props__.__dict__["dns_protocol"] = dns_protocol
        __props__.__dict__["dns_request_domain"] = dns_request_domain
        __props__.__dict__["http_get"] = http_get
        __props__.__dict__["http_match"] = http_match
        __props__.__dict__["http_max_redirects"] = http_max_redirects
        __props__.__dict__["interval"] = interval
        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["retry"] = retry
        __props__.__dict__["src_ip"] = src_ip
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["type"] = type
        __props__.__dict__["vdomparam"] = vdomparam
        return Ldbmonitor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dnsMatchIp")
    def dns_match_ip(self) -> pulumi.Output[str]:
        """
        Response IP expected from DNS server.
        """
        return pulumi.get(self, "dns_match_ip")

    @property
    @pulumi.getter(name="dnsProtocol")
    def dns_protocol(self) -> pulumi.Output[str]:
        """
        Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
        """
        return pulumi.get(self, "dns_protocol")

    @property
    @pulumi.getter(name="dnsRequestDomain")
    def dns_request_domain(self) -> pulumi.Output[str]:
        """
        Fully qualified domain name to resolve for the DNS probe.
        """
        return pulumi.get(self, "dns_request_domain")

    @property
    @pulumi.getter(name="httpGet")
    def http_get(self) -> pulumi.Output[str]:
        """
        URL used to send a GET request to check the health of an HTTP server.
        """
        return pulumi.get(self, "http_get")

    @property
    @pulumi.getter(name="httpMatch")
    def http_match(self) -> pulumi.Output[str]:
        """
        String to match the value expected in response to an HTTP-GET request.
        """
        return pulumi.get(self, "http_match")

    @property
    @pulumi.getter(name="httpMaxRedirects")
    def http_max_redirects(self) -> pulumi.Output[int]:
        """
        The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
        """
        return pulumi.get(self, "http_max_redirects")

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Output[int]:
        """
        Time between health checks (5 - 65635 sec, default = 10).
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Monitor name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65635, default = 0).
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def retry(self) -> pulumi.Output[int]:
        """
        Number health check attempts before the server is considered down (1 - 255, default = 3).
        """
        return pulumi.get(self, "retry")

    @property
    @pulumi.getter(name="srcIp")
    def src_ip(self) -> pulumi.Output[str]:
        """
        Source IP for ldb-monitor.
        """
        return pulumi.get(self, "src_ip")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[int]:
        """
        Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

