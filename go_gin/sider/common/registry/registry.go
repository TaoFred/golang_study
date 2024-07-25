package registry

import "golang.org/x/sys/windows/registry"

const (
	REGISTRYPATH    = "SOFTWARE\\SUPCON\\Integrity"
	DATAPATH        = "DataPath"
	INSTALLDIR      = "InstallDir"
	VERSION         = "Version"
	DATABASEVERSION = "DatabaseVersion"
)
const (
	DEFAULT_INSTALLDIR      = "C:\\Integrity"
	DEFAULT_DATAPATH        = "D:\\IntegrityData"
	DEFAULT_VERSION         = "V2.00.00.00-C"
	DEFAULT_DATABASEVERSION = ""
)

type RegistryInfo struct {
	DataPath        string
	InstallDir      string
	Version         string
	DatabaseVersion string
}

func ReadRegistry() (RegistryInfo, error) {
	registryInfo := RegistryInfo{
		DataPath:        DEFAULT_DATAPATH,
		InstallDir:      DEFAULT_INSTALLDIR,
		Version:         DEFAULT_VERSION,
		DatabaseVersion: DEFAULT_DATABASEVERSION,
	}

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, REGISTRYPATH, registry.READ|registry.WOW64_32KEY)
	if err != nil {
		return registryInfo, err
	}
	defer k.Close()

	value, _, err := k.GetStringValue(DATAPATH)
	if err == nil {
		registryInfo.DataPath = value
	}

	value, _, err = k.GetStringValue(INSTALLDIR)
	if err == nil {
		registryInfo.InstallDir = value
	}

	value, _, err = k.GetStringValue(VERSION)
	if err == nil {
		registryInfo.Version = value
	}

	value, _, err = k.GetStringValue(DATABASEVERSION)
	if err == nil {
		registryInfo.DatabaseVersion = value
	}
	return registryInfo, err
}
