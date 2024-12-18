# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'GetHAResourcesResult',
    'AwaitableGetHAResourcesResult',
    'get_ha_resources',
    'get_ha_resources_output',
]

@pulumi.output_type
class GetHAResourcesResult:
    """
    A collection of values returned by getHAResources.
    """
    def __init__(__self__, id=None, resource_ids=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if resource_ids and not isinstance(resource_ids, list):
            raise TypeError("Expected argument 'resource_ids' to be a list")
        pulumi.set(__self__, "resource_ids", resource_ids)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The unique identifier of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resourceIds")
    def resource_ids(self) -> Sequence[str]:
        """
        The identifiers of the High Availability resources.
        """
        return pulumi.get(self, "resource_ids")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of High Availability resources to fetch (`vm` or `ct`). All resources will be fetched if this option is unset.
        """
        return pulumi.get(self, "type")


class AwaitableGetHAResourcesResult(GetHAResourcesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHAResourcesResult(
            id=self.id,
            resource_ids=self.resource_ids,
            type=self.type)


def get_ha_resources(type: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHAResourcesResult:
    """
    Retrieves the list of High Availability resources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_proxmoxve as proxmoxve

    example_all = proxmoxve.HA.get_ha_resources()
    example_vm = proxmoxve.HA.get_ha_resources(type="vm")
    pulumi.export("dataProxmoxVirtualEnvironmentHaresources", {
        "all": example_all.resource_ids,
        "vms": example_vm.resource_ids,
    })
    ```


    :param str type: The type of High Availability resources to fetch (`vm` or `ct`). All resources will be fetched if this option is unset.
    """
    __args__ = dict()
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('proxmoxve:HA/getHAResources:getHAResources', __args__, opts=opts, typ=GetHAResourcesResult).value

    return AwaitableGetHAResourcesResult(
        id=pulumi.get(__ret__, 'id'),
        resource_ids=pulumi.get(__ret__, 'resource_ids'),
        type=pulumi.get(__ret__, 'type'))
def get_ha_resources_output(type: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetHAResourcesResult]:
    """
    Retrieves the list of High Availability resources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_proxmoxve as proxmoxve

    example_all = proxmoxve.HA.get_ha_resources()
    example_vm = proxmoxve.HA.get_ha_resources(type="vm")
    pulumi.export("dataProxmoxVirtualEnvironmentHaresources", {
        "all": example_all.resource_ids,
        "vms": example_vm.resource_ids,
    })
    ```


    :param str type: The type of High Availability resources to fetch (`vm` or `ct`). All resources will be fetched if this option is unset.
    """
    __args__ = dict()
    __args__['type'] = type
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('proxmoxve:HA/getHAResources:getHAResources', __args__, opts=opts, typ=GetHAResourcesResult)
    return __ret__.apply(lambda __response__: GetHAResourcesResult(
        id=pulumi.get(__response__, 'id'),
        resource_ids=pulumi.get(__response__, 'resource_ids'),
        type=pulumi.get(__response__, 'type')))
