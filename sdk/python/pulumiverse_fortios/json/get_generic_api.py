# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetGenericApiResult',
    'AwaitableGetGenericApiResult',
    'get_generic_api',
    'get_generic_api_output',
]

@pulumi.output_type
class GetGenericApiResult:
    """
    A collection of values returned by getGenericApi.
    """
    def __init__(__self__, id=None, path=None, response=None, specialparams=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if response and not isinstance(response, str):
            raise TypeError("Expected argument 'response' to be a str")
        pulumi.set(__self__, "response", response)
        if specialparams and not isinstance(specialparams, str):
            raise TypeError("Expected argument 'specialparams' to be a str")
        pulumi.set(__self__, "specialparams", specialparams)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        FortiAPI URL path.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def response(self) -> str:
        """
        FortiAPI returns results.
        """
        return pulumi.get(self, "response")

    @property
    @pulumi.getter
    def specialparams(self) -> Optional[str]:
        """
        URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
        """
        return pulumi.get(self, "specialparams")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetGenericApiResult(GetGenericApiResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGenericApiResult(
            id=self.id,
            path=self.path,
            response=self.response,
            specialparams=self.specialparams,
            vdomparam=self.vdomparam)


def get_generic_api(path: Optional[str] = None,
                    specialparams: Optional[str] = None,
                    vdomparam: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGenericApiResult:
    """
    Provides a FortiAPI Generic Interface data source.


    :param str path: FortiAPI URL path.
    :param str specialparams: URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
    """
    __args__ = dict()
    __args__['path'] = path
    __args__['specialparams'] = specialparams
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:json/getGenericApi:getGenericApi', __args__, opts=opts, typ=GetGenericApiResult).value

    return AwaitableGetGenericApiResult(
        id=pulumi.get(__ret__, 'id'),
        path=pulumi.get(__ret__, 'path'),
        response=pulumi.get(__ret__, 'response'),
        specialparams=pulumi.get(__ret__, 'specialparams'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_generic_api)
def get_generic_api_output(path: Optional[pulumi.Input[str]] = None,
                           specialparams: Optional[pulumi.Input[Optional[str]]] = None,
                           vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGenericApiResult]:
    """
    Provides a FortiAPI Generic Interface data source.


    :param str path: FortiAPI URL path.
    :param str specialparams: URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
    """
    ...
