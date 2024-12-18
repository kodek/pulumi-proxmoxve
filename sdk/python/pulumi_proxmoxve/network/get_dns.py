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
    'GetDNSResult',
    'AwaitableGetDNSResult',
    'get_dns',
    'get_dns_output',
]

@pulumi.output_type
class GetDNSResult:
    """
    A collection of values returned by getDNS.
    """
    def __init__(__self__, domain=None, id=None, node_name=None, servers=None):
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if node_name and not isinstance(node_name, str):
            raise TypeError("Expected argument 'node_name' to be a str")
        pulumi.set(__self__, "node_name", node_name)
        if servers and not isinstance(servers, list):
            raise TypeError("Expected argument 'servers' to be a list")
        pulumi.set(__self__, "servers", servers)

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        The DNS search domain.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> str:
        return pulumi.get(self, "node_name")

    @property
    @pulumi.getter
    def servers(self) -> Sequence[str]:
        """
        The DNS servers.
        """
        return pulumi.get(self, "servers")


class AwaitableGetDNSResult(GetDNSResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDNSResult(
            domain=self.domain,
            id=self.id,
            node_name=self.node_name,
            servers=self.servers)


def get_dns(node_name: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDNSResult:
    """
    Retrieves the DNS configuration for a specific node.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_proxmoxve as proxmoxve

    first_node = proxmoxve.Network.get_dns(node_name="first-node")
    ```


    :param str node_name: A node name.
    """
    __args__ = dict()
    __args__['nodeName'] = node_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('proxmoxve:Network/getDNS:getDNS', __args__, opts=opts, typ=GetDNSResult).value

    return AwaitableGetDNSResult(
        domain=pulumi.get(__ret__, 'domain'),
        id=pulumi.get(__ret__, 'id'),
        node_name=pulumi.get(__ret__, 'node_name'),
        servers=pulumi.get(__ret__, 'servers'))
def get_dns_output(node_name: Optional[pulumi.Input[str]] = None,
                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDNSResult]:
    """
    Retrieves the DNS configuration for a specific node.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_proxmoxve as proxmoxve

    first_node = proxmoxve.Network.get_dns(node_name="first-node")
    ```


    :param str node_name: A node name.
    """
    __args__ = dict()
    __args__['nodeName'] = node_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('proxmoxve:Network/getDNS:getDNS', __args__, opts=opts, typ=GetDNSResult)
    return __ret__.apply(lambda __response__: GetDNSResult(
        domain=pulumi.get(__response__, 'domain'),
        id=pulumi.get(__response__, 'id'),
        node_name=pulumi.get(__response__, 'node_name'),
        servers=pulumi.get(__response__, 'servers')))
