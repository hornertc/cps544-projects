## Advanced UNIX

- data

---

## Isolation Mechanisms

----

## Namespaces

----

### Sub Users and Groups

- Often used in containers
- Sub-users in `/etc/subuid`
- Sub-groups in `/etc/subgid`

---

## Kernel Modules

- A way of extending the Linux kernel
  - New filesystem
  - Device driver
- Can be dynamically loaded at runtime
- `lsmod` lists currently enabled modules
- `modprobe` can be used to load a module
  - To persist across reboot add it to `/etc/modprobe.d/my-mod.conf`

---

## IP Tables

---

## `direnv`

- Safe way to have dynamic environment variables based on *working directory*
- Integrates with version control
- Allows custom scripts/hooks to modify the environment dynamically
- Create a `.envrc` file in your project:

```bash
export MY_VARIABLE=123
```

- Enable it with `direnv allow .`

---

## Static vs Dynamically Linked Libraries/Executables

- `nm`
- `ldd`
- `go tool nm`
