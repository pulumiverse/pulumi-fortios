# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ObjectAdomRevisionArgs', 'ObjectAdomRevision']

@pulumi.input_type
class ObjectAdomRevisionArgs:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 created_by: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ObjectAdomRevision resource.
        :param pulumi.Input[str] adom: ADOM name. default is 'root'.
        :param pulumi.Input[str] created_by: Who created this adom revision.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[int] locked: lock. 0 means unlock and 1 means locked.
        :param pulumi.Input[str] name: Adom revision name.
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if created_by is not None:
            pulumi.set(__self__, "created_by", created_by)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if locked is not None:
            pulumi.set(__self__, "locked", locked)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        """
        ADOM name. default is 'root'.
        """
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[pulumi.Input[str]]:
        """
        Who created this adom revision.
        """
        return pulumi.get(self, "created_by")

    @created_by.setter
    def created_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_by", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def locked(self) -> Optional[pulumi.Input[int]]:
        """
        lock. 0 means unlock and 1 means locked.
        """
        return pulumi.get(self, "locked")

    @locked.setter
    def locked(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "locked", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Adom revision name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ObjectAdomRevisionState:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 created_by: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ObjectAdomRevision resources.
        :param pulumi.Input[str] adom: ADOM name. default is 'root'.
        :param pulumi.Input[str] created_by: Who created this adom revision.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[int] locked: lock. 0 means unlock and 1 means locked.
        :param pulumi.Input[str] name: Adom revision name.
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if created_by is not None:
            pulumi.set(__self__, "created_by", created_by)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if locked is not None:
            pulumi.set(__self__, "locked", locked)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        """
        ADOM name. default is 'root'.
        """
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[pulumi.Input[str]]:
        """
        Who created this adom revision.
        """
        return pulumi.get(self, "created_by")

    @created_by.setter
    def created_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_by", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def locked(self) -> Optional[pulumi.Input[int]]:
        """
        lock. 0 means unlock and 1 means locked.
        """
        return pulumi.get(self, "locked")

    @locked.setter
    def locked(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "locked", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Adom revision name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class ObjectAdomRevision(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 created_by: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource supports Create/Read/Update/Delete object adom revision for FortiManager.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        test1 = fortios.fmg.ObjectAdomRevision("test1",
            created_by="fortinet",
            description="adom revision",
            locked=0)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adom: ADOM name. default is 'root'.
        :param pulumi.Input[str] created_by: Who created this adom revision.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[int] locked: lock. 0 means unlock and 1 means locked.
        :param pulumi.Input[str] name: Adom revision name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ObjectAdomRevisionArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource supports Create/Read/Update/Delete object adom revision for FortiManager.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        test1 = fortios.fmg.ObjectAdomRevision("test1",
            created_by="fortinet",
            description="adom revision",
            locked=0)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param ObjectAdomRevisionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ObjectAdomRevisionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 created_by: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ObjectAdomRevisionArgs.__new__(ObjectAdomRevisionArgs)

            __props__.__dict__["adom"] = adom
            __props__.__dict__["created_by"] = created_by
            __props__.__dict__["description"] = description
            __props__.__dict__["locked"] = locked
            __props__.__dict__["name"] = name
        super(ObjectAdomRevision, __self__).__init__(
            'fortios:fmg/objectAdomRevision:ObjectAdomRevision',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            adom: Optional[pulumi.Input[str]] = None,
            created_by: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            locked: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'ObjectAdomRevision':
        """
        Get an existing ObjectAdomRevision resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adom: ADOM name. default is 'root'.
        :param pulumi.Input[str] created_by: Who created this adom revision.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[int] locked: lock. 0 means unlock and 1 means locked.
        :param pulumi.Input[str] name: Adom revision name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ObjectAdomRevisionState.__new__(_ObjectAdomRevisionState)

        __props__.__dict__["adom"] = adom
        __props__.__dict__["created_by"] = created_by
        __props__.__dict__["description"] = description
        __props__.__dict__["locked"] = locked
        __props__.__dict__["name"] = name
        return ObjectAdomRevision(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def adom(self) -> pulumi.Output[Optional[str]]:
        """
        ADOM name. default is 'root'.
        """
        return pulumi.get(self, "adom")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> pulumi.Output[Optional[str]]:
        """
        Who created this adom revision.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def locked(self) -> pulumi.Output[Optional[int]]:
        """
        lock. 0 means unlock and 1 means locked.
        """
        return pulumi.get(self, "locked")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Adom revision name.
        """
        return pulumi.get(self, "name")

