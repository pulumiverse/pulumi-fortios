# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ForticlientemsArgs', 'Forticlientems']

@pulumi.input_type
class ForticlientemsArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 admin_username: pulumi.Input[str],
                 serial_number: pulumi.Input[str],
                 admin_password: Optional[pulumi.Input[str]] = None,
                 admin_type: Optional[pulumi.Input[str]] = None,
                 https_port: Optional[pulumi.Input[int]] = None,
                 listen_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rest_api_auth: Optional[pulumi.Input[str]] = None,
                 upload_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Forticlientems resource.
        :param pulumi.Input[str] address: Firewall address name.
        :param pulumi.Input[str] admin_username: FortiClient EMS admin username.
        :param pulumi.Input[str] serial_number: FortiClient EMS Serial Number.
        :param pulumi.Input[str] admin_password: FortiClient EMS admin password.
        :param pulumi.Input[str] admin_type: FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
        :param pulumi.Input[int] https_port: FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        :param pulumi.Input[int] listen_port: FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
        :param pulumi.Input[str] name: FortiClient Enterprise Management Server (EMS) name.
        :param pulumi.Input[str] rest_api_auth: FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
        :param pulumi.Input[int] upload_port: FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "admin_username", admin_username)
        pulumi.set(__self__, "serial_number", serial_number)
        if admin_password is not None:
            pulumi.set(__self__, "admin_password", admin_password)
        if admin_type is not None:
            pulumi.set(__self__, "admin_type", admin_type)
        if https_port is not None:
            pulumi.set(__self__, "https_port", https_port)
        if listen_port is not None:
            pulumi.set(__self__, "listen_port", listen_port)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rest_api_auth is not None:
            pulumi.set(__self__, "rest_api_auth", rest_api_auth)
        if upload_port is not None:
            pulumi.set(__self__, "upload_port", upload_port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        Firewall address name.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> pulumi.Input[str]:
        """
        FortiClient EMS admin username.
        """
        return pulumi.get(self, "admin_username")

    @admin_username.setter
    def admin_username(self, value: pulumi.Input[str]):
        pulumi.set(self, "admin_username", value)

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> pulumi.Input[str]:
        """
        FortiClient EMS Serial Number.
        """
        return pulumi.get(self, "serial_number")

    @serial_number.setter
    def serial_number(self, value: pulumi.Input[str]):
        pulumi.set(self, "serial_number", value)

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS admin password.
        """
        return pulumi.get(self, "admin_password")

    @admin_password.setter
    def admin_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_password", value)

    @property
    @pulumi.getter(name="adminType")
    def admin_type(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
        """
        return pulumi.get(self, "admin_type")

    @admin_type.setter
    def admin_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_type", value)

    @property
    @pulumi.getter(name="httpsPort")
    def https_port(self) -> Optional[pulumi.Input[int]]:
        """
        FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        """
        return pulumi.get(self, "https_port")

    @https_port.setter
    def https_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "https_port", value)

    @property
    @pulumi.getter(name="listenPort")
    def listen_port(self) -> Optional[pulumi.Input[int]]:
        """
        FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
        """
        return pulumi.get(self, "listen_port")

    @listen_port.setter
    def listen_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "listen_port", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient Enterprise Management Server (EMS) name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="restApiAuth")
    def rest_api_auth(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
        """
        return pulumi.get(self, "rest_api_auth")

    @rest_api_auth.setter
    def rest_api_auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rest_api_auth", value)

    @property
    @pulumi.getter(name="uploadPort")
    def upload_port(self) -> Optional[pulumi.Input[int]]:
        """
        FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
        """
        return pulumi.get(self, "upload_port")

    @upload_port.setter
    def upload_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "upload_port", value)

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
class _ForticlientemsState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 admin_password: Optional[pulumi.Input[str]] = None,
                 admin_type: Optional[pulumi.Input[str]] = None,
                 admin_username: Optional[pulumi.Input[str]] = None,
                 https_port: Optional[pulumi.Input[int]] = None,
                 listen_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rest_api_auth: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[str]] = None,
                 upload_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Forticlientems resources.
        :param pulumi.Input[str] address: Firewall address name.
        :param pulumi.Input[str] admin_password: FortiClient EMS admin password.
        :param pulumi.Input[str] admin_type: FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
        :param pulumi.Input[str] admin_username: FortiClient EMS admin username.
        :param pulumi.Input[int] https_port: FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        :param pulumi.Input[int] listen_port: FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
        :param pulumi.Input[str] name: FortiClient Enterprise Management Server (EMS) name.
        :param pulumi.Input[str] rest_api_auth: FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
        :param pulumi.Input[str] serial_number: FortiClient EMS Serial Number.
        :param pulumi.Input[int] upload_port: FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if admin_password is not None:
            pulumi.set(__self__, "admin_password", admin_password)
        if admin_type is not None:
            pulumi.set(__self__, "admin_type", admin_type)
        if admin_username is not None:
            pulumi.set(__self__, "admin_username", admin_username)
        if https_port is not None:
            pulumi.set(__self__, "https_port", https_port)
        if listen_port is not None:
            pulumi.set(__self__, "listen_port", listen_port)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rest_api_auth is not None:
            pulumi.set(__self__, "rest_api_auth", rest_api_auth)
        if serial_number is not None:
            pulumi.set(__self__, "serial_number", serial_number)
        if upload_port is not None:
            pulumi.set(__self__, "upload_port", upload_port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Firewall address name.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS admin password.
        """
        return pulumi.get(self, "admin_password")

    @admin_password.setter
    def admin_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_password", value)

    @property
    @pulumi.getter(name="adminType")
    def admin_type(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
        """
        return pulumi.get(self, "admin_type")

    @admin_type.setter
    def admin_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_type", value)

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS admin username.
        """
        return pulumi.get(self, "admin_username")

    @admin_username.setter
    def admin_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_username", value)

    @property
    @pulumi.getter(name="httpsPort")
    def https_port(self) -> Optional[pulumi.Input[int]]:
        """
        FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        """
        return pulumi.get(self, "https_port")

    @https_port.setter
    def https_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "https_port", value)

    @property
    @pulumi.getter(name="listenPort")
    def listen_port(self) -> Optional[pulumi.Input[int]]:
        """
        FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
        """
        return pulumi.get(self, "listen_port")

    @listen_port.setter
    def listen_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "listen_port", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient Enterprise Management Server (EMS) name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="restApiAuth")
    def rest_api_auth(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
        """
        return pulumi.get(self, "rest_api_auth")

    @rest_api_auth.setter
    def rest_api_auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rest_api_auth", value)

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS Serial Number.
        """
        return pulumi.get(self, "serial_number")

    @serial_number.setter
    def serial_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "serial_number", value)

    @property
    @pulumi.getter(name="uploadPort")
    def upload_port(self) -> Optional[pulumi.Input[int]]:
        """
        FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
        """
        return pulumi.get(self, "upload_port")

    @upload_port.setter
    def upload_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "upload_port", value)

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


class Forticlientems(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 admin_password: Optional[pulumi.Input[str]] = None,
                 admin_type: Optional[pulumi.Input[str]] = None,
                 admin_username: Optional[pulumi.Input[str]] = None,
                 https_port: Optional[pulumi.Input[int]] = None,
                 listen_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rest_api_auth: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[str]] = None,
                 upload_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiClient Enterprise Management Server (EMS) entries. Applies to FortiOS Version `<= 6.2.0`.

        ## Import

        EndpointControl ForticlientEms can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:endpointcontrol/forticlientems:Forticlientems labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:endpointcontrol/forticlientems:Forticlientems labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Firewall address name.
        :param pulumi.Input[str] admin_password: FortiClient EMS admin password.
        :param pulumi.Input[str] admin_type: FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
        :param pulumi.Input[str] admin_username: FortiClient EMS admin username.
        :param pulumi.Input[int] https_port: FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        :param pulumi.Input[int] listen_port: FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
        :param pulumi.Input[str] name: FortiClient Enterprise Management Server (EMS) name.
        :param pulumi.Input[str] rest_api_auth: FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
        :param pulumi.Input[str] serial_number: FortiClient EMS Serial Number.
        :param pulumi.Input[int] upload_port: FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ForticlientemsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiClient Enterprise Management Server (EMS) entries. Applies to FortiOS Version `<= 6.2.0`.

        ## Import

        EndpointControl ForticlientEms can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:endpointcontrol/forticlientems:Forticlientems labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:endpointcontrol/forticlientems:Forticlientems labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ForticlientemsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ForticlientemsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 admin_password: Optional[pulumi.Input[str]] = None,
                 admin_type: Optional[pulumi.Input[str]] = None,
                 admin_username: Optional[pulumi.Input[str]] = None,
                 https_port: Optional[pulumi.Input[int]] = None,
                 listen_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rest_api_auth: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[str]] = None,
                 upload_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ForticlientemsArgs.__new__(ForticlientemsArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            __props__.__dict__["admin_password"] = None if admin_password is None else pulumi.Output.secret(admin_password)
            __props__.__dict__["admin_type"] = admin_type
            if admin_username is None and not opts.urn:
                raise TypeError("Missing required property 'admin_username'")
            __props__.__dict__["admin_username"] = admin_username
            __props__.__dict__["https_port"] = https_port
            __props__.__dict__["listen_port"] = listen_port
            __props__.__dict__["name"] = name
            __props__.__dict__["rest_api_auth"] = rest_api_auth
            if serial_number is None and not opts.urn:
                raise TypeError("Missing required property 'serial_number'")
            __props__.__dict__["serial_number"] = serial_number
            __props__.__dict__["upload_port"] = upload_port
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["adminPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Forticlientems, __self__).__init__(
            'fortios:endpointcontrol/forticlientems:Forticlientems',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            admin_password: Optional[pulumi.Input[str]] = None,
            admin_type: Optional[pulumi.Input[str]] = None,
            admin_username: Optional[pulumi.Input[str]] = None,
            https_port: Optional[pulumi.Input[int]] = None,
            listen_port: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rest_api_auth: Optional[pulumi.Input[str]] = None,
            serial_number: Optional[pulumi.Input[str]] = None,
            upload_port: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Forticlientems':
        """
        Get an existing Forticlientems resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Firewall address name.
        :param pulumi.Input[str] admin_password: FortiClient EMS admin password.
        :param pulumi.Input[str] admin_type: FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
        :param pulumi.Input[str] admin_username: FortiClient EMS admin username.
        :param pulumi.Input[int] https_port: FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        :param pulumi.Input[int] listen_port: FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
        :param pulumi.Input[str] name: FortiClient Enterprise Management Server (EMS) name.
        :param pulumi.Input[str] rest_api_auth: FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
        :param pulumi.Input[str] serial_number: FortiClient EMS Serial Number.
        :param pulumi.Input[int] upload_port: FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ForticlientemsState.__new__(_ForticlientemsState)

        __props__.__dict__["address"] = address
        __props__.__dict__["admin_password"] = admin_password
        __props__.__dict__["admin_type"] = admin_type
        __props__.__dict__["admin_username"] = admin_username
        __props__.__dict__["https_port"] = https_port
        __props__.__dict__["listen_port"] = listen_port
        __props__.__dict__["name"] = name
        __props__.__dict__["rest_api_auth"] = rest_api_auth
        __props__.__dict__["serial_number"] = serial_number
        __props__.__dict__["upload_port"] = upload_port
        __props__.__dict__["vdomparam"] = vdomparam
        return Forticlientems(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        Firewall address name.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> pulumi.Output[Optional[str]]:
        """
        FortiClient EMS admin password.
        """
        return pulumi.get(self, "admin_password")

    @property
    @pulumi.getter(name="adminType")
    def admin_type(self) -> pulumi.Output[str]:
        """
        FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
        """
        return pulumi.get(self, "admin_type")

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> pulumi.Output[str]:
        """
        FortiClient EMS admin username.
        """
        return pulumi.get(self, "admin_username")

    @property
    @pulumi.getter(name="httpsPort")
    def https_port(self) -> pulumi.Output[int]:
        """
        FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        """
        return pulumi.get(self, "https_port")

    @property
    @pulumi.getter(name="listenPort")
    def listen_port(self) -> pulumi.Output[int]:
        """
        FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
        """
        return pulumi.get(self, "listen_port")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        FortiClient Enterprise Management Server (EMS) name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="restApiAuth")
    def rest_api_auth(self) -> pulumi.Output[str]:
        """
        FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
        """
        return pulumi.get(self, "rest_api_auth")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> pulumi.Output[str]:
        """
        FortiClient EMS Serial Number.
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="uploadPort")
    def upload_port(self) -> pulumi.Output[int]:
        """
        FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
        """
        return pulumi.get(self, "upload_port")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

