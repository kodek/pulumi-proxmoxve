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

__all__ = [
    'ContainerClone',
    'ContainerConsole',
    'ContainerCpu',
    'ContainerDisk',
    'ContainerFeatures',
    'ContainerInitialization',
    'ContainerInitializationDns',
    'ContainerInitializationIpConfig',
    'ContainerInitializationIpConfigIpv4',
    'ContainerInitializationIpConfigIpv6',
    'ContainerInitializationUserAccount',
    'ContainerMemory',
    'ContainerMountPoint',
    'ContainerNetworkInterface',
    'ContainerOperatingSystem',
    'ContainerStartup',
]

@pulumi.output_type
class ContainerClone(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "vmId":
            suggest = "vm_id"
        elif key == "datastoreId":
            suggest = "datastore_id"
        elif key == "nodeName":
            suggest = "node_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerClone. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerClone.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerClone.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 vm_id: int,
                 datastore_id: Optional[str] = None,
                 node_name: Optional[str] = None):
        """
        :param int vm_id: The container identifier
        :param str datastore_id: The identifier for the datastore to create the
               disk in (defaults to `local`).
        :param str node_name: The name of the node to assign the container to.
        """
        pulumi.set(__self__, "vm_id", vm_id)
        if datastore_id is not None:
            pulumi.set(__self__, "datastore_id", datastore_id)
        if node_name is not None:
            pulumi.set(__self__, "node_name", node_name)

    @property
    @pulumi.getter(name="vmId")
    def vm_id(self) -> int:
        """
        The container identifier
        """
        return pulumi.get(self, "vm_id")

    @property
    @pulumi.getter(name="datastoreId")
    def datastore_id(self) -> Optional[str]:
        """
        The identifier for the datastore to create the
        disk in (defaults to `local`).
        """
        return pulumi.get(self, "datastore_id")

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> Optional[str]:
        """
        The name of the node to assign the container to.
        """
        return pulumi.get(self, "node_name")


@pulumi.output_type
class ContainerConsole(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ttyCount":
            suggest = "tty_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerConsole. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerConsole.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerConsole.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: Optional[bool] = None,
                 tty_count: Optional[int] = None,
                 type: Optional[str] = None):
        """
        :param bool enabled: Whether to enable the network device (defaults
               to `true`).
        :param int tty_count: The number of available TTY (defaults to `2`).
        :param str type: The type (defaults to `unmanaged`).
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if tty_count is not None:
            pulumi.set(__self__, "tty_count", tty_count)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Whether to enable the network device (defaults
        to `true`).
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="ttyCount")
    def tty_count(self) -> Optional[int]:
        """
        The number of available TTY (defaults to `2`).
        """
        return pulumi.get(self, "tty_count")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type (defaults to `unmanaged`).
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class ContainerCpu(dict):
    def __init__(__self__, *,
                 architecture: Optional[str] = None,
                 cores: Optional[int] = None,
                 units: Optional[int] = None):
        """
        :param str architecture: The CPU architecture (defaults to `amd64`).
        :param int cores: The number of CPU cores (defaults to `1`).
        :param int units: The CPU units (defaults to `1024`).
        """
        if architecture is not None:
            pulumi.set(__self__, "architecture", architecture)
        if cores is not None:
            pulumi.set(__self__, "cores", cores)
        if units is not None:
            pulumi.set(__self__, "units", units)

    @property
    @pulumi.getter
    def architecture(self) -> Optional[str]:
        """
        The CPU architecture (defaults to `amd64`).
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter
    def cores(self) -> Optional[int]:
        """
        The number of CPU cores (defaults to `1`).
        """
        return pulumi.get(self, "cores")

    @property
    @pulumi.getter
    def units(self) -> Optional[int]:
        """
        The CPU units (defaults to `1024`).
        """
        return pulumi.get(self, "units")


@pulumi.output_type
class ContainerDisk(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "datastoreId":
            suggest = "datastore_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerDisk. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerDisk.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerDisk.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 datastore_id: Optional[str] = None,
                 size: Optional[int] = None):
        """
        :param str datastore_id: The identifier for the datastore to create the
               disk in (defaults to `local`).
        :param int size: Volume size (only for volume mount points).
               Can be specified with a unit suffix (e.g. `10G`).
        """
        if datastore_id is not None:
            pulumi.set(__self__, "datastore_id", datastore_id)
        if size is not None:
            pulumi.set(__self__, "size", size)

    @property
    @pulumi.getter(name="datastoreId")
    def datastore_id(self) -> Optional[str]:
        """
        The identifier for the datastore to create the
        disk in (defaults to `local`).
        """
        return pulumi.get(self, "datastore_id")

    @property
    @pulumi.getter
    def size(self) -> Optional[int]:
        """
        Volume size (only for volume mount points).
        Can be specified with a unit suffix (e.g. `10G`).
        """
        return pulumi.get(self, "size")


@pulumi.output_type
class ContainerFeatures(dict):
    def __init__(__self__, *,
                 fuse: Optional[bool] = None,
                 keyctl: Optional[bool] = None,
                 mounts: Optional[Sequence[str]] = None,
                 nesting: Optional[bool] = None):
        """
        :param bool fuse: Whether the container supports FUSE mounts (defaults
               to `false`)
        :param bool keyctl: Whether the container supports `keyctl()` system
               call (defaults to `false`)
        :param Sequence[str] mounts: List of allowed mount types (`cifs` or `nfs`)
        :param bool nesting: Whether the container is nested (defaults
               to `false`)
        """
        if fuse is not None:
            pulumi.set(__self__, "fuse", fuse)
        if keyctl is not None:
            pulumi.set(__self__, "keyctl", keyctl)
        if mounts is not None:
            pulumi.set(__self__, "mounts", mounts)
        if nesting is not None:
            pulumi.set(__self__, "nesting", nesting)

    @property
    @pulumi.getter
    def fuse(self) -> Optional[bool]:
        """
        Whether the container supports FUSE mounts (defaults
        to `false`)
        """
        return pulumi.get(self, "fuse")

    @property
    @pulumi.getter
    def keyctl(self) -> Optional[bool]:
        """
        Whether the container supports `keyctl()` system
        call (defaults to `false`)
        """
        return pulumi.get(self, "keyctl")

    @property
    @pulumi.getter
    def mounts(self) -> Optional[Sequence[str]]:
        """
        List of allowed mount types (`cifs` or `nfs`)
        """
        return pulumi.get(self, "mounts")

    @property
    @pulumi.getter
    def nesting(self) -> Optional[bool]:
        """
        Whether the container is nested (defaults
        to `false`)
        """
        return pulumi.get(self, "nesting")


@pulumi.output_type
class ContainerInitialization(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ipConfigs":
            suggest = "ip_configs"
        elif key == "userAccount":
            suggest = "user_account"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerInitialization. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerInitialization.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerInitialization.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 dns: Optional['outputs.ContainerInitializationDns'] = None,
                 hostname: Optional[str] = None,
                 ip_configs: Optional[Sequence['outputs.ContainerInitializationIpConfig']] = None,
                 user_account: Optional['outputs.ContainerInitializationUserAccount'] = None):
        """
        :param 'ContainerInitializationDnsArgs' dns: The DNS configuration.
        :param str hostname: The hostname.
        :param Sequence['ContainerInitializationIpConfigArgs'] ip_configs: The IP configuration (one block per network
               device).
        :param 'ContainerInitializationUserAccountArgs' user_account: The user account configuration.
        """
        if dns is not None:
            pulumi.set(__self__, "dns", dns)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if ip_configs is not None:
            pulumi.set(__self__, "ip_configs", ip_configs)
        if user_account is not None:
            pulumi.set(__self__, "user_account", user_account)

    @property
    @pulumi.getter
    def dns(self) -> Optional['outputs.ContainerInitializationDns']:
        """
        The DNS configuration.
        """
        return pulumi.get(self, "dns")

    @property
    @pulumi.getter
    def hostname(self) -> Optional[str]:
        """
        The hostname.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="ipConfigs")
    def ip_configs(self) -> Optional[Sequence['outputs.ContainerInitializationIpConfig']]:
        """
        The IP configuration (one block per network
        device).
        """
        return pulumi.get(self, "ip_configs")

    @property
    @pulumi.getter(name="userAccount")
    def user_account(self) -> Optional['outputs.ContainerInitializationUserAccount']:
        """
        The user account configuration.
        """
        return pulumi.get(self, "user_account")


@pulumi.output_type
class ContainerInitializationDns(dict):
    def __init__(__self__, *,
                 domain: Optional[str] = None,
                 server: Optional[str] = None,
                 servers: Optional[Sequence[str]] = None):
        """
        :param str domain: The DNS search domain.
        :param str server: The DNS server. The `server` attribute is
               deprecated and will be removed in a future release. Please use
               the `servers` attribute instead.
        :param Sequence[str] servers: The list of DNS servers.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if servers is not None:
            pulumi.set(__self__, "servers", servers)

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        """
        The DNS search domain.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def server(self) -> Optional[str]:
        """
        The DNS server. The `server` attribute is
        deprecated and will be removed in a future release. Please use
        the `servers` attribute instead.
        """
        warnings.warn("""The `server` attribute is deprecated and will be removed in a future release. Please use the `servers` attribute instead.""", DeprecationWarning)
        pulumi.log.warn("""server is deprecated: The `server` attribute is deprecated and will be removed in a future release. Please use the `servers` attribute instead.""")

        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def servers(self) -> Optional[Sequence[str]]:
        """
        The list of DNS servers.
        """
        return pulumi.get(self, "servers")


@pulumi.output_type
class ContainerInitializationIpConfig(dict):
    def __init__(__self__, *,
                 ipv4: Optional['outputs.ContainerInitializationIpConfigIpv4'] = None,
                 ipv6: Optional['outputs.ContainerInitializationIpConfigIpv6'] = None):
        """
        :param 'ContainerInitializationIpConfigIpv4Args' ipv4: The IPv4 configuration.
        :param 'ContainerInitializationIpConfigIpv6Args' ipv6: The IPv4 configuration.
        """
        if ipv4 is not None:
            pulumi.set(__self__, "ipv4", ipv4)
        if ipv6 is not None:
            pulumi.set(__self__, "ipv6", ipv6)

    @property
    @pulumi.getter
    def ipv4(self) -> Optional['outputs.ContainerInitializationIpConfigIpv4']:
        """
        The IPv4 configuration.
        """
        return pulumi.get(self, "ipv4")

    @property
    @pulumi.getter
    def ipv6(self) -> Optional['outputs.ContainerInitializationIpConfigIpv6']:
        """
        The IPv4 configuration.
        """
        return pulumi.get(self, "ipv6")


@pulumi.output_type
class ContainerInitializationIpConfigIpv4(dict):
    def __init__(__self__, *,
                 address: Optional[str] = None,
                 gateway: Optional[str] = None):
        """
        :param str address: The IPv6 address (use `dhcp` for
               autodiscovery).
        :param str gateway: The IPv6 gateway (must be omitted
               when `dhcp` is used as the address).
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        """
        The IPv6 address (use `dhcp` for
        autodiscovery).
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def gateway(self) -> Optional[str]:
        """
        The IPv6 gateway (must be omitted
        when `dhcp` is used as the address).
        """
        return pulumi.get(self, "gateway")


@pulumi.output_type
class ContainerInitializationIpConfigIpv6(dict):
    def __init__(__self__, *,
                 address: Optional[str] = None,
                 gateway: Optional[str] = None):
        """
        :param str address: The IPv6 address (use `dhcp` for
               autodiscovery).
        :param str gateway: The IPv6 gateway (must be omitted
               when `dhcp` is used as the address).
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        """
        The IPv6 address (use `dhcp` for
        autodiscovery).
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def gateway(self) -> Optional[str]:
        """
        The IPv6 gateway (must be omitted
        when `dhcp` is used as the address).
        """
        return pulumi.get(self, "gateway")


@pulumi.output_type
class ContainerInitializationUserAccount(dict):
    def __init__(__self__, *,
                 keys: Optional[Sequence[str]] = None,
                 password: Optional[str] = None):
        """
        :param Sequence[str] keys: The SSH keys for the root account.
        :param str password: The password for the root account.
        """
        if keys is not None:
            pulumi.set(__self__, "keys", keys)
        if password is not None:
            pulumi.set(__self__, "password", password)

    @property
    @pulumi.getter
    def keys(self) -> Optional[Sequence[str]]:
        """
        The SSH keys for the root account.
        """
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        """
        The password for the root account.
        """
        return pulumi.get(self, "password")


@pulumi.output_type
class ContainerMemory(dict):
    def __init__(__self__, *,
                 dedicated: Optional[int] = None,
                 swap: Optional[int] = None):
        """
        :param int dedicated: The dedicated memory in megabytes (defaults
               to `512`).
        :param int swap: The swap size in megabytes (defaults to `0`).
        """
        if dedicated is not None:
            pulumi.set(__self__, "dedicated", dedicated)
        if swap is not None:
            pulumi.set(__self__, "swap", swap)

    @property
    @pulumi.getter
    def dedicated(self) -> Optional[int]:
        """
        The dedicated memory in megabytes (defaults
        to `512`).
        """
        return pulumi.get(self, "dedicated")

    @property
    @pulumi.getter
    def swap(self) -> Optional[int]:
        """
        The swap size in megabytes (defaults to `0`).
        """
        return pulumi.get(self, "swap")


@pulumi.output_type
class ContainerMountPoint(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "mountOptions":
            suggest = "mount_options"
        elif key == "readOnly":
            suggest = "read_only"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerMountPoint. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerMountPoint.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerMountPoint.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 path: str,
                 volume: str,
                 acl: Optional[bool] = None,
                 backup: Optional[bool] = None,
                 mount_options: Optional[Sequence[str]] = None,
                 quota: Optional[bool] = None,
                 read_only: Optional[bool] = None,
                 replicate: Optional[bool] = None,
                 shared: Optional[bool] = None,
                 size: Optional[str] = None):
        """
        :param str path: Path to the mount point as seen from inside the
               container.
        :param str volume: Volume, device or directory to mount into the
               container.
        :param bool acl: Explicitly enable or disable ACL support.
        :param bool backup: Whether to include the mount point in backups (only
               used for volume mount points).
        :param Sequence[str] mount_options: List of extra mount options.
        :param bool quota: Enable user quotas inside the container (not supported
               with ZFS subvolumes).
        :param bool read_only: Read-only mount point.
        :param bool replicate: Will include this volume to a storage replica job.
        :param bool shared: Mark this non-volume mount point as available on all
               nodes.
        :param str size: Volume size (only for volume mount points).
               Can be specified with a unit suffix (e.g. `10G`).
        """
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "volume", volume)
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if backup is not None:
            pulumi.set(__self__, "backup", backup)
        if mount_options is not None:
            pulumi.set(__self__, "mount_options", mount_options)
        if quota is not None:
            pulumi.set(__self__, "quota", quota)
        if read_only is not None:
            pulumi.set(__self__, "read_only", read_only)
        if replicate is not None:
            pulumi.set(__self__, "replicate", replicate)
        if shared is not None:
            pulumi.set(__self__, "shared", shared)
        if size is not None:
            pulumi.set(__self__, "size", size)

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        Path to the mount point as seen from inside the
        container.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def volume(self) -> str:
        """
        Volume, device or directory to mount into the
        container.
        """
        return pulumi.get(self, "volume")

    @property
    @pulumi.getter
    def acl(self) -> Optional[bool]:
        """
        Explicitly enable or disable ACL support.
        """
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def backup(self) -> Optional[bool]:
        """
        Whether to include the mount point in backups (only
        used for volume mount points).
        """
        return pulumi.get(self, "backup")

    @property
    @pulumi.getter(name="mountOptions")
    def mount_options(self) -> Optional[Sequence[str]]:
        """
        List of extra mount options.
        """
        return pulumi.get(self, "mount_options")

    @property
    @pulumi.getter
    def quota(self) -> Optional[bool]:
        """
        Enable user quotas inside the container (not supported
        with ZFS subvolumes).
        """
        return pulumi.get(self, "quota")

    @property
    @pulumi.getter(name="readOnly")
    def read_only(self) -> Optional[bool]:
        """
        Read-only mount point.
        """
        return pulumi.get(self, "read_only")

    @property
    @pulumi.getter
    def replicate(self) -> Optional[bool]:
        """
        Will include this volume to a storage replica job.
        """
        return pulumi.get(self, "replicate")

    @property
    @pulumi.getter
    def shared(self) -> Optional[bool]:
        """
        Mark this non-volume mount point as available on all
        nodes.
        """
        return pulumi.get(self, "shared")

    @property
    @pulumi.getter
    def size(self) -> Optional[str]:
        """
        Volume size (only for volume mount points).
        Can be specified with a unit suffix (e.g. `10G`).
        """
        return pulumi.get(self, "size")


@pulumi.output_type
class ContainerNetworkInterface(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "macAddress":
            suggest = "mac_address"
        elif key == "rateLimit":
            suggest = "rate_limit"
        elif key == "vlanId":
            suggest = "vlan_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerNetworkInterface. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerNetworkInterface.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerNetworkInterface.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 bridge: Optional[str] = None,
                 enabled: Optional[bool] = None,
                 firewall: Optional[bool] = None,
                 mac_address: Optional[str] = None,
                 mtu: Optional[int] = None,
                 rate_limit: Optional[float] = None,
                 vlan_id: Optional[int] = None):
        """
        :param str name: The network interface name.
        :param str bridge: The name of the network bridge (defaults
               to `vmbr0`).
        :param bool enabled: Whether to enable the network device (defaults
               to `true`).
        :param bool firewall: Whether this interface's firewall rules should be
               used (defaults to `false`).
        :param str mac_address: The MAC address.
        :param int mtu: Maximum transfer unit of the interface. Cannot be
               larger than the bridge's MTU.
        :param float rate_limit: The rate limit in megabytes per second.
        :param int vlan_id: The VLAN identifier.
        """
        pulumi.set(__self__, "name", name)
        if bridge is not None:
            pulumi.set(__self__, "bridge", bridge)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if firewall is not None:
            pulumi.set(__self__, "firewall", firewall)
        if mac_address is not None:
            pulumi.set(__self__, "mac_address", mac_address)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if rate_limit is not None:
            pulumi.set(__self__, "rate_limit", rate_limit)
        if vlan_id is not None:
            pulumi.set(__self__, "vlan_id", vlan_id)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The network interface name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def bridge(self) -> Optional[str]:
        """
        The name of the network bridge (defaults
        to `vmbr0`).
        """
        return pulumi.get(self, "bridge")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Whether to enable the network device (defaults
        to `true`).
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def firewall(self) -> Optional[bool]:
        """
        Whether this interface's firewall rules should be
        used (defaults to `false`).
        """
        return pulumi.get(self, "firewall")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[str]:
        """
        The MAC address.
        """
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter
    def mtu(self) -> Optional[int]:
        """
        Maximum transfer unit of the interface. Cannot be
        larger than the bridge's MTU.
        """
        return pulumi.get(self, "mtu")

    @property
    @pulumi.getter(name="rateLimit")
    def rate_limit(self) -> Optional[float]:
        """
        The rate limit in megabytes per second.
        """
        return pulumi.get(self, "rate_limit")

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> Optional[int]:
        """
        The VLAN identifier.
        """
        return pulumi.get(self, "vlan_id")


@pulumi.output_type
class ContainerOperatingSystem(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "templateFileId":
            suggest = "template_file_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerOperatingSystem. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerOperatingSystem.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerOperatingSystem.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 template_file_id: str,
                 type: Optional[str] = None):
        """
        :param str template_file_id: The identifier for an OS template file.
        :param str type: The type (defaults to `unmanaged`).
        """
        pulumi.set(__self__, "template_file_id", template_file_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="templateFileId")
    def template_file_id(self) -> str:
        """
        The identifier for an OS template file.
        """
        return pulumi.get(self, "template_file_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type (defaults to `unmanaged`).
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class ContainerStartup(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "downDelay":
            suggest = "down_delay"
        elif key == "upDelay":
            suggest = "up_delay"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerStartup. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerStartup.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerStartup.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 down_delay: Optional[int] = None,
                 order: Optional[int] = None,
                 up_delay: Optional[int] = None):
        """
        :param int down_delay: A non-negative number defining the delay in seconds before the next container is shut down
        :param int order: A non-negative number defining the general startup
               order.
        :param int up_delay: A non-negative number defining the delay in seconds before the next container is started
        """
        if down_delay is not None:
            pulumi.set(__self__, "down_delay", down_delay)
        if order is not None:
            pulumi.set(__self__, "order", order)
        if up_delay is not None:
            pulumi.set(__self__, "up_delay", up_delay)

    @property
    @pulumi.getter(name="downDelay")
    def down_delay(self) -> Optional[int]:
        """
        A non-negative number defining the delay in seconds before the next container is shut down
        """
        return pulumi.get(self, "down_delay")

    @property
    @pulumi.getter
    def order(self) -> Optional[int]:
        """
        A non-negative number defining the general startup
        order.
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="upDelay")
    def up_delay(self) -> Optional[int]:
        """
        A non-negative number defining the delay in seconds before the next container is started
        """
        return pulumi.get(self, "up_delay")


