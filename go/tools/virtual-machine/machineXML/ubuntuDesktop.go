package machineXML

const XmlUbuntuDesktop = `
<domain type="kvm">
    <name>%s</name>
    <memory unit="MiB">%d</memory>
    <vcpu placement="static">%d</vcpu>
    <os firmware="efi">
        <type arch="x86_64" machine="pc-q35-6.2">hvm</type>
        <smbios mode='sysinfo'/>
        <boot dev="hd"/>
    </os>
    <sysinfo type='smbios'>
      <bios>
        <entry name='vendor'>LENOVO</entry>
      </bios>
      <system>
        <entry name='manufacturer'>LENOVO</entry>
        <entry name='product'>Virt-Manager</entry>
        <entry name='version'>0.9.4</entry>
      </system>
      <baseBoard>
        <entry name='manufacturer'>LENOVO</entry>
        <entry name='product'>20BE0061MC</entry>
        <entry name='version'>0B98401 Pro</entry>
        <entry name='serial'>W1KS427111E</entry>
      </baseBoard>
      <chassis>
        <entry name='manufacturer'>LENOVO</entry>
        <entry name='product'>Virt-Manager</entry>
        <entry name='version'>2.12</entry>
        <entry name='serial'>65X0XF2</entry>
        <entry name='asset'>40000101</entry>
        <entry name='sku'>Type3Sku1</entry>
      </chassis>
      <oemStrings>
        <entry>myappname:virtualization</entry>
        <entry>code:https://github.com/sam-caldwell/monorepo</entry>
      </oemStrings>
    </sysinfo>
    <features>
        <acpi/>
        <apic/>
        <vmport state="off"/>
    </features>
    <cpu mode="host-passthrough" check="none" migratable="on"/>
    <clock offset="utc">
        <timer name="rtc" tickpolicy="catchup"/>
        <timer name="pit" tickpolicy="delay"/>
        <timer name="hpet" present="no"/>
    </clock>
    <on_poweroff>destroy</on_poweroff>
    <on_reboot>restart</on_reboot>
    <on_crash>destroy</on_crash>
    <pm>
        <suspend-to-mem enabled="no"/>
        <suspend-to-disk enabled="no"/>
    </pm>
    <devices>
        <emulator>/usr/bin/qemu-system-x86_64</emulator>
        <disk type="file" device="disk">
            <driver name="qemu" type="qcow2" discard="unmap" />
            <source file="%s" />
            <target dev="vda" bus="virtio" />
            <address type="pci" domain="0x0000" bus="0x04" slot="0x00" function="0x0" />
        </disk>
        <disk type="file" device="cdrom">
            <driver name="qemu" type="raw" />
            <source file='%s' />
            <target dev="sda" bus="sata" />
            <readonly />
            <address type="drive" controller="0" bus="0" target="0" unit="0" />
        </disk>
        <filesystem type='mount'>
          <source dir='%s'/>
          <target dir='/home/git'/>
        </filesystem>
        <controller type="usb" index="0" model="qemu-xhci" ports="15">
            <address type="pci" domain="0x0000" bus="0x02" slot="0x00" function="0x0" />
        </controller>
        <controller type="pci" index="0" model="pcie-root" />
        <controller type="pci" index="1" model="pcie-root-port">
            <model name="pcie-root-port"/>
            <target chassis="1" port="0x10"/>
            <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x0" multifunction="on" />
        </controller>
        <controller type="pci" index="2" model="pcie-root-port">
            <model name="pcie-root-port"/>
            <target chassis="2" port="0x11"/>
            <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x1" />
        </controller>
        <controller type="pci" index="3" model="pcie-root-port">
            <model name="pcie-root-port"/>
            <target chassis="3" port="0x12"/>
            <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x2" />
        </controller>
        <controller type="pci" index="4" model="pcie-root-port">
            <model name="pcie-root-port"/>
            <target chassis="4" port="0x13"/>
            <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x3" />
        </controller>
        <controller type="pci" index="5" model="pcie-root-port">
            <model name="pcie-root-port"/>
            <target chassis="5" port="0x14"/>
            <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x4" />
        </controller>
        <controller type="pci" index="6" model="pcie-root-port">
            <model name="pcie-root-port"/>
            <target chassis="6" port="0x15"/>
            <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x5" />
        </controller>
        <controller type="pci" index="7" model="pcie-root-port">
            <model name="pcie-root-port"/>
            <target chassis="7" port="0x16"/>
            <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x6" />
        </controller>
        <controller type="pci" index="8" model="pcie-root-port">
            <model name="pcie-root-port"/>
            <target chassis="8" port="0x17"/>
            <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x7" />
        </controller>
        <controller type="pci" index="9" model="pcie-root-port">
            <model name="pcie-root-port"/>
            <target chassis="9" port="0x18"/>
            <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x0" multifunction="on" />
        </controller>
        <controller type="pci" index="10" model="pcie-root-port">
            <model name="pcie-root-port"/>
            <target chassis="10" port="0x19"/>
            <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x1" />
        </controller>
        <controller type="pci" index="11" model="pcie-root-port">
            <model name="pcie-root-port"/>
            <target chassis="11" port="0x1a"/>
            <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x2" />
        </controller>
        <controller type="pci" index="12" model="pcie-root-port">
            <model name="pcie-root-port" />
            <target chassis="12" port="0x1b" />
            <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x3" />
        </controller>
        <controller type="pci" index="13" model="pcie-root-port">
            <model name="pcie-root-port" />
            <target chassis="13" port="0x1c" />
            <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x4" />
        </controller>
        <controller type="pci" index="14" model="pcie-root-port">
            <model name="pcie-root-port" />
            <target chassis="14" port="0x1d" />
            <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x5" />
        </controller>
        <controller type="sata" index="0">
            <address type="pci" domain="0x0000" bus="0x00" slot="0x1f" function="0x2" />
        </controller>
        <controller type="virtio-serial" index="0">
            <address type="pci" domain="0x0000" bus="0x03" slot="0x00" function="0x0" />
        </controller>
        <interface type="network">
            <mac address="52:54:00:91:8a:4d" />
            <source network="default" />
            <model type="virtio" />
            <link state="down" />
            <address type="pci" domain="0x0000" bus="0x01" slot="0x00" function="0x0" />
        </interface>
        <serial type="pty">
            <target type="isa-serial" port="0">
                <model name="isa-serial" />
            </target>
        </serial>
        <console type="pty">
            <target type="serial" port="0" />
        </console>
        <channel type="unix">
            <target type="virtio" name="org.qemu.guest_agent.0" />
            <address type="virtio-serial" controller="0" bus="0" port="1" />
        </channel>
        <channel type="spicevmc">
            <target type="virtio" name="com.redhat.spice.0" />
            <address type="virtio-serial" controller="0" bus="0" port="2" />
        </channel>
        <input type="tablet" bus="usb">
            <address type="usb" bus="0" port="1" />
        </input>
        <input type="mouse" bus="ps2" />
        <input type="keyboard" bus="ps2" />
            <graphics type="spice">
                <listen type="none" />
                <image compression="off" />
            </graphics>
            <sound model="ich9">
                <address type="pci" domain="0x0000" bus="0x00" slot="0x1b" function="0x0" />
            </sound>
            <audio id="1" type="spice"/>
            <video>
                <model type="qxl" ram="65536" vram="65536" vgamem="16384" heads="1" primary="yes"/>
                <address type="pci" domain="0x0000" bus="0x00" slot="0x01" function="0x0"/>
            </video>
            <redirdev bus="usb" type="spicevmc">
                <address type="usb" bus="0" port="2"/>
            </redirdev>
            <redirdev bus="usb" type="spicevmc">
                <address type="usb" bus="0" port="3"/>
            </redirdev>
            <memballoon model="virtio">
                <address type="pci" domain="0x0000" bus="0x05" slot="0x00" function="0x0"/>
            </memballoon>
            <rng model="virtio">
                <backend model="random">/dev/urandom</backend>
                <address type="pci" domain="0x0000" bus="0x06" slot="0x00" function="0x0"/>
            </rng>
    </devices>
</domain>`