# Virtualization on Fedora

```sh
# List related packages in virtualization package group
$ dnf groupinfo virtualization
```

```sh
# Install packages
$ sudo dnf install @virtualization
```


```sh
# Start libvirtd
$ sudo systemctl start libvirtd
```

```sh
# Then you can use packages
$ virsh list
```

```sh
# Install libraries before we can actually compile mediumkube
$ sudo dnf install libvirt-devel-6.6.0-5.fc33.x86_64
```