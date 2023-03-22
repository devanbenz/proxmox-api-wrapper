package pkg

import (
	"github.com/gorilla/schema"
	"net/url"
	"strings"
)

type VirtualMachineInput struct {
	Node              string `schema:"node,required"`
	Vmid              int    `schema:"vmid,required"`
	Acpi              bool   `schema:"acpi"`
	Affinity          string `schema:"affinity"`
	Agent             string `schema:"agent"`
	Arch              string `schema:"arch"`
	Archive           string `schema:"archive"`
	Args              string `schema:"args"`
	Audio0            string `schema:"audio0"`
	Autostart         bool   `schema:"autostart"`
	Balloon           int    `schema:"balloon"`
	Bios              string `schema:"bios"`
	Boot              string `schema:"boot"`
	Bootdisk          string `schema:"bootdisk"`
	Bwlimit           int    `schema:"bwlimit"`
	Cdrom             string `schema:"cdrom"`
	Cicustom          string `schema:"cicustom"`
	Cipassword        string `schema:"cipassword"`
	Citype            string `schema:"citype"`
	Ciuser            string `schema:"ciuser"`
	Cores             int    `schema:"cores"`
	Cpu               string `schema:"cpu"`
	Cpulimit          string `schema:"cpulimit"`
	Cpuunits          int    `schema:"cpuunits"`
	Description       string `schema:"description"`
	Efidisk0          string `schema:"efidisk0"`
	Force             bool   `schema:"force"`
	Freeze            bool   `schema:"freeze"`
	Hookscript        string `schema:"hookscript"`
	HostpciN          string `schema:"hostpci[n]"`
	Hotplug           string `schema:"hotplug"`
	Hugepages         string `schema:"hugepages"`
	IdeN              string `schema:"ide[n]"`
	IpconfigN         string `schema:"ipconfig[n]"`
	Ivshmem           string `schema:"ivshmem"`
	Keephugepages     bool   `schema:"keephugepages"`
	Keyboard          string `schema:"keyboard"`
	Kvm               bool   `schema:"kvm"`
	LiveRestore       bool   `schema:"live-restore"`
	Localtime         bool   `schema:"localtime"`
	Lock              string `schema:"lock"`
	Machine           string `schema:"machine"`
	Memory            int    `schema:"memory"`
	MigrateDowntime   string `schema:"migrate_downtime"`
	MigrateSpeed      int    `schema:"migrate_speed"`
	Name              string `schema:"name"`
	Nameserver        string `schema:"nameserver"`
	NetN              string `schema:"net[n]"`
	Numa              bool   `schema:"numa"`
	NumaN             string `schema:"numa[n]"`
	Onboot            bool   `schema:"onboot"`
	Ostype            string `schema:"ostype"`
	ParallelN         string `schema:"parallel[n]"`
	Pool              string `schema:"pool"`
	Protection        bool   `schema:"protection"`
	Reboot            bool   `schema:"reboot"`
	Rng0              string `schema:"rng0"`
	Sata              string `schema:"sata[0]"`
	Scsi              string `schema:"scsi[0]"`
	Scsihw            string `schema:"scsihw"`
	Searchdomain      string `schema:"searchdomain"`
	SerialN           string `schema:"serial[n]"`
	Shares            int    `schema:"shares"`
	Smbios1           string `schema:"smbios1"`
	Smp               int    `schema:"smp"`
	Sockets           int    `schema:"sockets"`
	SpiceEnhancements string `schema:"spice_enhancements"`
	Sshkeys           string `schema:"sshkeys"`
	Start             bool   `schema:"start"`
	Startdate         string `schema:"startdate"`
	Startup           string `schema:"startup"`
	Storage           string `schema:"storage"`
	Tablet            bool   `schema:"tablet"`
	Tags              string `schema:"tags"`
	Tdf               bool   `schema:"tdf"`
	Template          bool   `schema:"template"`
	Tpmstate0         string `schema:"tpmstate0"`
	Unique            bool   `schema:"unique"`
	UnusedN           string `schema:"unused[n]"`
	UsbN              string `schema:"usb[n]"`
	Vcpus             int    `schema:"vcpus"`
	Vga               string `schema:"vga"`
	VirtioN           string `schema:"virtio[n]"`
	Vmgenid           string `schema:"vmgenid"`
	Vmstatestorage    string `schema:"vmstatestorage"`
	Watchdog          string `schema:"watchdog"`
}

func (v *VirtualMachineInput) CreateVirtualMachine() *strings.Reader {
	encoder := schema.NewEncoder()
	form := url.Values{}

	err := encoder.Encode(v, form)
	CheckErr(err)
	return strings.NewReader(form.Encode())
}
