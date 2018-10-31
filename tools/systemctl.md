### 使用 systemctl 管理系统服务

##### 启动和关闭服务

```bash
systemctl start application.service / application
systemctl stop application.service / application
```

##### 重启和重新加载服务

```bash
systemctl restart application.service / application
systemctl reload application.service /application

# 如果 reload 可用则 reload 应用，否则重启应用
systemctl reload-or-restart application.service
```

##### Enable and disable service

```bash
# enable 后服务开机自动启
systemctl enable application.service
systemctl disable application.service
```

##### 检查服务状态

```bash
systemctl status application.service
```

##### 检查服务是否处于运行状态

```bash
systemctl is-active application.service
```

##### 检查服务是否 enabled

```bash
systemctl is-enabled application.service
```

##### 检查服务是否处于失败状态

```bash
systemctl is-failed application.service
```

##### 列出目前所有服务单元

```bash
systemctl list-units
```

```bash
UNIT                                      LOAD   ACTIVE SUB     DESCRIPTION
atd.service                               loaded active running ATD daemon
avahi-daemon.service                      loaded active running Avahi mDNS/DNS-SD Stack
dbus.service                              loaded active running D-Bus System Message Bus
dcron.service                             loaded active running Periodic Command Scheduler
dkms.service                              loaded active exited  Dynamic Kernel Modules System
getty@tty1.service                        loaded active running Getty on tty1
. . .
UNIT: The systemd unit name
LOAD: Whether the unit's configuration has been parsed by systemd. The configuration of loaded units is kept in memory.
ACTIVE: A summary state about whether the unit is active. This is usually a fairly basic way to tell if the unit has started successfully or not.
SUB: This is a lower-level state that indicates more detailed information about the unit. This often varies by unit type, state, and the actual method in which the unit runs.
DESCRIPTION: A short textual description of what the unit is/does.

# list-units 支持多个参数，比如
systemctl list-units --all
systemctl list-units --all --state=inactive
systemctl list-units --type=service

# 
systemctl list-unit-files
```

##### 单元管理

```bash
# 显示一个单元文件
systemctl cat atd.service
```

```bash
[Unit]
Description=ATD daemon
[Service]
Type=forking
ExecStart=/usr/bin/atd
[Install]
WantedBy=multi-user.target
```

##### 显示依赖

```bash
systemctl list-dependencies sshd.service
```

```bash
sshd.service
├─system.slice
└─basic.target
  ├─microcode.service
  ├─rhel-autorelabel-mark.service
  ├─rhel-autorelabel.service
  ├─rhel-configure.service
  ├─rhel-dmesg.service
  ├─rhel-loadmodules.service
  ├─paths.target
  ├─slices.target
. . .
```

##### 检查单元属性

```bash
systemctl show sshd.service
```

```bash
Id=sshd.service
Names=sshd.service
Requires=basic.target
Wants=system.slice
WantedBy=multi-user.target
Conflicts=shutdown.target
Before=shutdown.target multi-user.target
After=syslog.target network.target auditd.service systemd-journald.socket basic.target system.slice
Description=OpenSSH server daemon
. . .
```

```bash
# 检查单元单一属性
systemctl show sshd.service -p Conflicts

Conflicts=shutdown.target
```

[How To Use Systemctl to Manage Systemd Services and Units]: https://www.digitalocean.com/community/tutorials/how-to-use-systemctl-to-manage-systemd-services-and-units

