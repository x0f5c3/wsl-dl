package internal

import "errors"

var (
	ErrInvalidDistro  = errors.New("invalid distro")
	ErrInvalidVersion = errors.New("invalid version")
	LinkMap           = map[string]string{
		"Ubuntu":                              UbuntuLink,
		"Ubuntu 22.04 LTS":                    Ubuntu2204lts,
		"Ubuntu 20.04":                        Ubuntu2004,
		"Ubuntu 20.04 ARM":                    Ubuntu2004Arm,
		"Ubuntu 18.04":                        Ubuntu1804link,
		"Ubuntu 18.04 ARM":                    Ubuntu1804Arm,
		"Ubuntu 16.04":                        Ubuntu1604link,
		"Debian GNU/Linux":                    DebianGnuLinux,
		"Kali Linux":                          KaliLinux,
		"SUSE Linux Enterprise Server 12":     SuseLinuxEnterpriseServer12,
		"SUSE Linux Enterprise Server 15 SP2": SuseLinuxEnterpriseServer15Sp2,
		"SUSE Linux Enterprise Server 15 SP3": SuseLinuxEnterpriseServer15Sp3,
		"openSUSE Tumbleweed":                 opensuseTumbleweed,
		"openSUSE Leap 15.3":                  opensuseLeap153,
		"openSUSE Leap 15.2":                  opensuseLeap152,
		"Oracle Linux 8.5":                    OracleLinux85,
		"Oracle Linux 7.9":                    OracleLinux79,
	}
	FilenameMap = map[string]string{
		"Ubuntu":                              "ubuntu.zip",
		"Ubuntu 22.04 LTS":                    "ubuntu-22.04.zip",
		"Ubuntu 20.04":                        "ubuntu-20.04.zip",
		"Ubuntu 20.04 ARM":                    "ubuntu-20.04-arm.zip",
		"Ubuntu 18.04":                        "ubuntu-18.04.zip",
		"Ubuntu 18.04 ARM":                    "ubuntu-18.04-arm.zip",
		"Ubuntu 16.04":                        "ubuntu-16.04.zip",
		"Debian GNU/Linux":                    "debian.zip",
		"Kali Linux":                          "kali.zip",
		"SUSE Linux Enterprise Server 12":     "suse-12.zip",
		"SUSE Linux Enterprise Server 15 SP2": "suse-15-sp2.zip",
		"SUSE Linux Enterprise Server 15 SP3": "suse-15-sp3.zip",
		"openSUSE Tumbleweed":                 "opensuse-tumbleweed.zip",
		"openSUSE Leap 15.3":                  "opensuse-leap-15.3.zip",
		"openSUSE Leap 15.2":                  "opensuse-leap-15.2.zip",
		"Oracle Linux 8.5":                    "oracle-8.5.zip",
		"Oracle Linux 7.9":                    "oracle-7.9.zip",
	}
	LinkKeys = func() []string {
		var keys []string
		for k := range LinkMap {
			keys = append(keys, k)
		}
		return keys
	}()
)

const (
	UbuntuLink                     = "https://aka.ms/wslubuntu"
	Ubuntu2204lts                  = "https://aka.ms/wslubuntu2204"
	Ubuntu2004                     = "https://aka.ms/wslubuntu2004"
	Ubuntu2004Arm                  = "https://aka.ms/wslubuntu2004arm"
	Ubuntu1804link                 = "https://aka.ms/wsl-ubuntu-1804"
	Ubuntu1804Arm                  = "https://aka.ms/wsl-ubuntu-1804-arm"
	Ubuntu1604link                 = "https://aka.ms/wsl-ubuntu-1604"
	DebianGnuLinux                 = "https://aka.ms/wsl-debian-gnulinux"
	KaliLinux                      = "https://aka.ms/wsl-kali-linux-new"
	SuseLinuxEnterpriseServer12    = "https://aka.ms/wsl-sles-12"
	SuseLinuxEnterpriseServer15Sp2 = "https://aka.ms/wsl-SUSELinuxEnterpriseServer15SP2"
	SuseLinuxEnterpriseServer15Sp3 = "https://aka.ms/wsl-SUSELinuxEnterpriseServer15SP3"
	opensuseTumbleweed             = "https://aka.ms/wsl-opensuse-tumbleweed"
	opensuseLeap153                = "https://aka.ms/wsl-opensuseleap15-3"
	opensuseLeap152                = "https://aka.ms/wsl-opensuseleap15-2"
	OracleLinux85                  = "https://aka.ms/wsl-oraclelinux-8-5"
	OracleLinux79                  = "https://aka.ms/wsl-oraclelinux-7-9"
)
