# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ProfileArgs', 'Profile']

@pulumi.input_type
class ProfileArgs:
    def __init__(__self__, *,
                 auth_group: Optional[pulumi.Input[str]] = None,
                 cifs: Optional[pulumi.Input['ProfileCifsArgs']] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 ftp: Optional[pulumi.Input['ProfileFtpArgs']] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 http: Optional[pulumi.Input['ProfileHttpArgs']] = None,
                 mapi: Optional[pulumi.Input['ProfileMapiArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tcp: Optional[pulumi.Input['ProfileTcpArgs']] = None,
                 transparent: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Profile resource.
        :param pulumi.Input[str] auth_group: Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        :param pulumi.Input['ProfileCifsArgs'] cifs: Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input['ProfileFtpArgs'] ftp: Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input['ProfileHttpArgs'] http: Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        :param pulumi.Input['ProfileMapiArgs'] mapi: Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input['ProfileTcpArgs'] tcp: Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        :param pulumi.Input[str] transparent: Enable/disable transparent mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auth_group is not None:
            pulumi.set(__self__, "auth_group", auth_group)
        if cifs is not None:
            pulumi.set(__self__, "cifs", cifs)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if ftp is not None:
            pulumi.set(__self__, "ftp", ftp)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if http is not None:
            pulumi.set(__self__, "http", http)
        if mapi is not None:
            pulumi.set(__self__, "mapi", mapi)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tcp is not None:
            pulumi.set(__self__, "tcp", tcp)
        if transparent is not None:
            pulumi.set(__self__, "transparent", transparent)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authGroup")
    def auth_group(self) -> Optional[pulumi.Input[str]]:
        """
        Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        """
        return pulumi.get(self, "auth_group")

    @auth_group.setter
    def auth_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_group", value)

    @property
    @pulumi.getter
    def cifs(self) -> Optional[pulumi.Input['ProfileCifsArgs']]:
        """
        Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        """
        return pulumi.get(self, "cifs")

    @cifs.setter
    def cifs(self, value: Optional[pulumi.Input['ProfileCifsArgs']]):
        pulumi.set(self, "cifs", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def ftp(self) -> Optional[pulumi.Input['ProfileFtpArgs']]:
        """
        Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        """
        return pulumi.get(self, "ftp")

    @ftp.setter
    def ftp(self, value: Optional[pulumi.Input['ProfileFtpArgs']]):
        pulumi.set(self, "ftp", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def http(self) -> Optional[pulumi.Input['ProfileHttpArgs']]:
        """
        Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        """
        return pulumi.get(self, "http")

    @http.setter
    def http(self, value: Optional[pulumi.Input['ProfileHttpArgs']]):
        pulumi.set(self, "http", value)

    @property
    @pulumi.getter
    def mapi(self) -> Optional[pulumi.Input['ProfileMapiArgs']]:
        """
        Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        """
        return pulumi.get(self, "mapi")

    @mapi.setter
    def mapi(self, value: Optional[pulumi.Input['ProfileMapiArgs']]):
        pulumi.set(self, "mapi", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tcp(self) -> Optional[pulumi.Input['ProfileTcpArgs']]:
        """
        Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        """
        return pulumi.get(self, "tcp")

    @tcp.setter
    def tcp(self, value: Optional[pulumi.Input['ProfileTcpArgs']]):
        pulumi.set(self, "tcp", value)

    @property
    @pulumi.getter
    def transparent(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable transparent mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "transparent")

    @transparent.setter
    def transparent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transparent", value)

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
class _ProfileState:
    def __init__(__self__, *,
                 auth_group: Optional[pulumi.Input[str]] = None,
                 cifs: Optional[pulumi.Input['ProfileCifsArgs']] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 ftp: Optional[pulumi.Input['ProfileFtpArgs']] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 http: Optional[pulumi.Input['ProfileHttpArgs']] = None,
                 mapi: Optional[pulumi.Input['ProfileMapiArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tcp: Optional[pulumi.Input['ProfileTcpArgs']] = None,
                 transparent: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Profile resources.
        :param pulumi.Input[str] auth_group: Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        :param pulumi.Input['ProfileCifsArgs'] cifs: Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input['ProfileFtpArgs'] ftp: Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input['ProfileHttpArgs'] http: Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        :param pulumi.Input['ProfileMapiArgs'] mapi: Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input['ProfileTcpArgs'] tcp: Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        :param pulumi.Input[str] transparent: Enable/disable transparent mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auth_group is not None:
            pulumi.set(__self__, "auth_group", auth_group)
        if cifs is not None:
            pulumi.set(__self__, "cifs", cifs)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if ftp is not None:
            pulumi.set(__self__, "ftp", ftp)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if http is not None:
            pulumi.set(__self__, "http", http)
        if mapi is not None:
            pulumi.set(__self__, "mapi", mapi)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tcp is not None:
            pulumi.set(__self__, "tcp", tcp)
        if transparent is not None:
            pulumi.set(__self__, "transparent", transparent)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authGroup")
    def auth_group(self) -> Optional[pulumi.Input[str]]:
        """
        Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        """
        return pulumi.get(self, "auth_group")

    @auth_group.setter
    def auth_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_group", value)

    @property
    @pulumi.getter
    def cifs(self) -> Optional[pulumi.Input['ProfileCifsArgs']]:
        """
        Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        """
        return pulumi.get(self, "cifs")

    @cifs.setter
    def cifs(self, value: Optional[pulumi.Input['ProfileCifsArgs']]):
        pulumi.set(self, "cifs", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def ftp(self) -> Optional[pulumi.Input['ProfileFtpArgs']]:
        """
        Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        """
        return pulumi.get(self, "ftp")

    @ftp.setter
    def ftp(self, value: Optional[pulumi.Input['ProfileFtpArgs']]):
        pulumi.set(self, "ftp", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def http(self) -> Optional[pulumi.Input['ProfileHttpArgs']]:
        """
        Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        """
        return pulumi.get(self, "http")

    @http.setter
    def http(self, value: Optional[pulumi.Input['ProfileHttpArgs']]):
        pulumi.set(self, "http", value)

    @property
    @pulumi.getter
    def mapi(self) -> Optional[pulumi.Input['ProfileMapiArgs']]:
        """
        Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        """
        return pulumi.get(self, "mapi")

    @mapi.setter
    def mapi(self, value: Optional[pulumi.Input['ProfileMapiArgs']]):
        pulumi.set(self, "mapi", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tcp(self) -> Optional[pulumi.Input['ProfileTcpArgs']]:
        """
        Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        """
        return pulumi.get(self, "tcp")

    @tcp.setter
    def tcp(self, value: Optional[pulumi.Input['ProfileTcpArgs']]):
        pulumi.set(self, "tcp", value)

    @property
    @pulumi.getter
    def transparent(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable transparent mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "transparent")

    @transparent.setter
    def transparent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transparent", value)

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


class Profile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_group: Optional[pulumi.Input[str]] = None,
                 cifs: Optional[pulumi.Input[pulumi.InputType['ProfileCifsArgs']]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 ftp: Optional[pulumi.Input[pulumi.InputType['ProfileFtpArgs']]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 http: Optional[pulumi.Input[pulumi.InputType['ProfileHttpArgs']]] = None,
                 mapi: Optional[pulumi.Input[pulumi.InputType['ProfileMapiArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tcp: Optional[pulumi.Input[pulumi.InputType['ProfileTcpArgs']]] = None,
                 transparent: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure WAN optimization profiles.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.wanopt.Profile("trname",
            cifs=fortios.wanopt.ProfileCifsArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=445,
                prefer_chunking="fix",
                secure_tunnel="disable",
                status="disable",
                tunnel_sharing="private",
            ),
            comments="test",
            ftp=fortios.wanopt.ProfileFtpArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=21,
                prefer_chunking="fix",
                secure_tunnel="disable",
                status="disable",
                tunnel_sharing="private",
            ),
            http=fortios.wanopt.ProfileHttpArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=80,
                prefer_chunking="fix",
                secure_tunnel="disable",
                ssl="disable",
                ssl_port=443,
                status="disable",
                tunnel_non_http="disable",
                tunnel_sharing="private",
                unknown_http_version="tunnel",
            ),
            mapi=fortios.wanopt.ProfileMapiArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=135,
                secure_tunnel="disable",
                status="disable",
                tunnel_sharing="private",
            ),
            tcp=fortios.wanopt.ProfileTcpArgs(
                byte_caching="disable",
                byte_caching_opt="mem-only",
                log_traffic="enable",
                port="1-65535",
                secure_tunnel="disable",
                ssl="disable",
                ssl_port=443,
                status="disable",
                tunnel_sharing="private",
            ),
            transparent="enable")
        ```

        ## Import

        Wanopt Profile can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wanopt/profile:Profile labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wanopt/profile:Profile labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_group: Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        :param pulumi.Input[pulumi.InputType['ProfileCifsArgs']] cifs: Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[pulumi.InputType['ProfileFtpArgs']] ftp: Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[pulumi.InputType['ProfileHttpArgs']] http: Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        :param pulumi.Input[pulumi.InputType['ProfileMapiArgs']] mapi: Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input[pulumi.InputType['ProfileTcpArgs']] tcp: Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        :param pulumi.Input[str] transparent: Enable/disable transparent mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure WAN optimization profiles.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.wanopt.Profile("trname",
            cifs=fortios.wanopt.ProfileCifsArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=445,
                prefer_chunking="fix",
                secure_tunnel="disable",
                status="disable",
                tunnel_sharing="private",
            ),
            comments="test",
            ftp=fortios.wanopt.ProfileFtpArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=21,
                prefer_chunking="fix",
                secure_tunnel="disable",
                status="disable",
                tunnel_sharing="private",
            ),
            http=fortios.wanopt.ProfileHttpArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=80,
                prefer_chunking="fix",
                secure_tunnel="disable",
                ssl="disable",
                ssl_port=443,
                status="disable",
                tunnel_non_http="disable",
                tunnel_sharing="private",
                unknown_http_version="tunnel",
            ),
            mapi=fortios.wanopt.ProfileMapiArgs(
                byte_caching="enable",
                log_traffic="enable",
                port=135,
                secure_tunnel="disable",
                status="disable",
                tunnel_sharing="private",
            ),
            tcp=fortios.wanopt.ProfileTcpArgs(
                byte_caching="disable",
                byte_caching_opt="mem-only",
                log_traffic="enable",
                port="1-65535",
                secure_tunnel="disable",
                ssl="disable",
                ssl_port=443,
                status="disable",
                tunnel_sharing="private",
            ),
            transparent="enable")
        ```

        ## Import

        Wanopt Profile can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wanopt/profile:Profile labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wanopt/profile:Profile labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_group: Optional[pulumi.Input[str]] = None,
                 cifs: Optional[pulumi.Input[pulumi.InputType['ProfileCifsArgs']]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 ftp: Optional[pulumi.Input[pulumi.InputType['ProfileFtpArgs']]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 http: Optional[pulumi.Input[pulumi.InputType['ProfileHttpArgs']]] = None,
                 mapi: Optional[pulumi.Input[pulumi.InputType['ProfileMapiArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tcp: Optional[pulumi.Input[pulumi.InputType['ProfileTcpArgs']]] = None,
                 transparent: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProfileArgs.__new__(ProfileArgs)

            __props__.__dict__["auth_group"] = auth_group
            __props__.__dict__["cifs"] = cifs
            __props__.__dict__["comments"] = comments
            __props__.__dict__["ftp"] = ftp
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["http"] = http
            __props__.__dict__["mapi"] = mapi
            __props__.__dict__["name"] = name
            __props__.__dict__["tcp"] = tcp
            __props__.__dict__["transparent"] = transparent
            __props__.__dict__["vdomparam"] = vdomparam
        super(Profile, __self__).__init__(
            'fortios:wanopt/profile:Profile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_group: Optional[pulumi.Input[str]] = None,
            cifs: Optional[pulumi.Input[pulumi.InputType['ProfileCifsArgs']]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            ftp: Optional[pulumi.Input[pulumi.InputType['ProfileFtpArgs']]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            http: Optional[pulumi.Input[pulumi.InputType['ProfileHttpArgs']]] = None,
            mapi: Optional[pulumi.Input[pulumi.InputType['ProfileMapiArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tcp: Optional[pulumi.Input[pulumi.InputType['ProfileTcpArgs']]] = None,
            transparent: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Profile':
        """
        Get an existing Profile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_group: Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        :param pulumi.Input[pulumi.InputType['ProfileCifsArgs']] cifs: Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[pulumi.InputType['ProfileFtpArgs']] ftp: Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[pulumi.InputType['ProfileHttpArgs']] http: Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        :param pulumi.Input[pulumi.InputType['ProfileMapiArgs']] mapi: Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input[pulumi.InputType['ProfileTcpArgs']] tcp: Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        :param pulumi.Input[str] transparent: Enable/disable transparent mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProfileState.__new__(_ProfileState)

        __props__.__dict__["auth_group"] = auth_group
        __props__.__dict__["cifs"] = cifs
        __props__.__dict__["comments"] = comments
        __props__.__dict__["ftp"] = ftp
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["http"] = http
        __props__.__dict__["mapi"] = mapi
        __props__.__dict__["name"] = name
        __props__.__dict__["tcp"] = tcp
        __props__.__dict__["transparent"] = transparent
        __props__.__dict__["vdomparam"] = vdomparam
        return Profile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authGroup")
    def auth_group(self) -> pulumi.Output[str]:
        """
        Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        """
        return pulumi.get(self, "auth_group")

    @property
    @pulumi.getter
    def cifs(self) -> pulumi.Output['outputs.ProfileCifs']:
        """
        Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        """
        return pulumi.get(self, "cifs")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def ftp(self) -> pulumi.Output['outputs.ProfileFtp']:
        """
        Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        """
        return pulumi.get(self, "ftp")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter
    def http(self) -> pulumi.Output['outputs.ProfileHttp']:
        """
        Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        """
        return pulumi.get(self, "http")

    @property
    @pulumi.getter
    def mapi(self) -> pulumi.Output['outputs.ProfileMapi']:
        """
        Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        """
        return pulumi.get(self, "mapi")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tcp(self) -> pulumi.Output['outputs.ProfileTcp']:
        """
        Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        """
        return pulumi.get(self, "tcp")

    @property
    @pulumi.getter
    def transparent(self) -> pulumi.Output[str]:
        """
        Enable/disable transparent mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "transparent")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

