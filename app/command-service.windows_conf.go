// +build windows

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type setting struct {
	path         string
	defaultValue interface{}
}

var settingInputSdCardTemplatePath = setting{
	"input.sdcard_template.path",
	"./data/sdcard-template",
}

var settingInputBaseConfigurationPath = setting{
	"input.base_configuration.path",
	"./data/configs/default.atv",
}

var settingInputMergeConfigurationPath = setting{
	"input.merge_configuration.path",
	"",
}

var settingInputHotfolderPath = setting{
	"input.hotfolder.path",
	"./data/input",
}

var settingOutputMergedConfigurationsPath = setting{
	"output.merged_configurations.path",
	"./data/output-merged-configs",
}

var settingOutputMergedConfigurationsWriteAtv = setting{
	"output.merged_configurations.write_atv",
	true,
}

var settingOutputMergedConfigurationsWriteEcs = setting{
	"output.merged_configurations.write_ecs",
	false,
}

var settingOutputUpdatePackagesPath = setting{
	"output.update_packages.path",
	"./data/output-update-packages",
}

var allSettings = []setting{
	settingInputSdCardTemplatePath,
	settingInputBaseConfigurationPath,
	settingInputMergeConfigurationPath,
	settingInputHotfolderPath,
	settingOutputMergedConfigurationsPath,
	settingOutputMergedConfigurationsWriteAtv,
	settingOutputMergedConfigurationsWriteEcs,
	settingOutputUpdatePackagesPath,
}

// loadServiceConfiguration loads the service configuration from the specified file.
func (cmd *ServiceCommand) loadServiceConfiguration(path string, createIfNotExist bool) error {

	// set up new viper configuration with default settings
	conf := viper.New()
	for _, setting := range allSettings {
		conf.SetDefault(setting.path, setting.defaultValue)
	}

	// read configuration from file
	log.Debugf("Loading configuration file '%s'...", path)
	basename := filepath.Base(path)
	configName := strings.TrimSuffix(basename, filepath.Ext(basename))
	configDir := filepath.Dir(path) + string(os.PathSeparator)
	conf.SetConfigName(configName)
	conf.SetConfigType("yaml")
	conf.AddConfigPath(configDir)
	err := conf.ReadInConfig()
	if err != nil {
		switch err := err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Debugf("Configuration file '%s' does not exist.", path)
			if createIfNotExist {
				log.Debugf("Creating new configuration file '%s'...", path)
				err := conf.WriteConfigAs(path)
				if err != nil {
					log.Errorf("Saving configuration file '%s' failed: %v", path, err)
					// proceed with default settings...
				}
			}
		default:
			// some severe error, cannot proceed...
			return err
		}
	}

	// configuration is available now
	// => validate settings

	// input: sdcard template path (must be a directory)
	log.Debugf("Setting '%s': '%s'", settingInputSdCardTemplatePath.path, conf.GetString(settingInputSdCardTemplatePath.path))
	cmd.sdcardTemplateDirectory = conf.GetString(settingInputSdCardTemplatePath.path)
	if len(cmd.sdcardTemplateDirectory) > 0 { // setting is optional
		if filepath.IsAbs(cmd.sdcardTemplateDirectory) {
			cmd.sdcardTemplateDirectory = filepath.Clean(cmd.sdcardTemplateDirectory)
		} else {
			path, err := filepath.Abs(filepath.Join(configDir, cmd.sdcardTemplateDirectory))
			if err != nil {
				return err
			}
			cmd.sdcardTemplateDirectory = path
		}
	} else {
		return fmt.Errorf("setting '%s' is not set.", settingInputSdCardTemplatePath.path)
	}

	// input: base configuration file
	log.Debugf("Setting '%s': '%s'", settingInputBaseConfigurationPath.path, conf.GetString(settingInputBaseConfigurationPath.path))
	cmd.baseConfigurationPath = conf.GetString(settingInputBaseConfigurationPath.path)
	if len(cmd.baseConfigurationPath) > 0 {
		if filepath.IsAbs(cmd.baseConfigurationPath) {
			cmd.baseConfigurationPath = filepath.Clean(cmd.baseConfigurationPath)
		} else {
			path, err := filepath.Abs(filepath.Join(configDir, cmd.baseConfigurationPath))
			if err != nil {
				return err
			}
			cmd.baseConfigurationPath = path
		}
	} else {
		return fmt.Errorf("setting '%s' is not set.", settingInputBaseConfigurationPath.path)
	}

	// input: merge configuration file
	log.Debugf("Setting '%s': '%s'", settingInputMergeConfigurationPath.path, conf.GetString(settingInputMergeConfigurationPath.path))
	cmd.mergeConfigurationPath = conf.GetString(settingInputMergeConfigurationPath.path)
	if len(cmd.mergeConfigurationPath) > 0 { // setting is optional
		if filepath.IsAbs(cmd.mergeConfigurationPath) {
			cmd.mergeConfigurationPath = filepath.Clean(cmd.mergeConfigurationPath)
		} else {
			path, err := filepath.Abs(filepath.Join(configDir, cmd.mergeConfigurationPath))
			if err != nil {
				return err
			}
			cmd.mergeConfigurationPath = path
		}
	}

	// input: hot folder path
	log.Debugf("Setting '%s': '%s'", settingInputHotfolderPath.path, conf.GetString(settingInputHotfolderPath.path))
	cmd.hotFolderPath = conf.GetString(settingInputHotfolderPath.path)
	if len(cmd.hotFolderPath) > 0 {
		if filepath.IsAbs(cmd.hotFolderPath) {
			cmd.hotFolderPath = filepath.Clean(cmd.hotFolderPath)
		} else {
			path, err := filepath.Abs(filepath.Join(configDir, cmd.hotFolderPath))
			if err != nil {
				return err
			}
			cmd.hotFolderPath = path
		}
	} else {
		return fmt.Errorf("setting '%s' is not set.", settingInputHotfolderPath.path)
	}

	// output: merged configuration directory
	log.Debugf("Setting '%s': '%s'", settingOutputMergedConfigurationsPath.path, conf.GetString(settingOutputMergedConfigurationsPath.path))
	cmd.mergedConfigurationDirectory = conf.GetString(settingOutputMergedConfigurationsPath.path)
	if len(cmd.mergedConfigurationDirectory) > 0 { // setting is optional
		if filepath.IsAbs(cmd.mergedConfigurationDirectory) {
			cmd.mergedConfigurationDirectory = filepath.Clean(cmd.mergedConfigurationDirectory)
		} else {
			path, err := filepath.Abs(filepath.Join(configDir, cmd.mergedConfigurationDirectory))
			if err != nil {
				return err
			}
			cmd.mergedConfigurationDirectory = path
		}
	}

	// output: merged configuration directory - write atv
	log.Debugf("Setting '%s': '%s'", settingOutputMergedConfigurationsWriteAtv.path, conf.GetString(settingOutputMergedConfigurationsWriteAtv.path))
	cmd.mergedConfigurationsWriteAtv = conf.GetBool(settingOutputMergedConfigurationsWriteAtv.path)

	// output: merged configuration directory - write ecs
	log.Debugf("Setting '%s': '%s'", settingOutputMergedConfigurationsWriteEcs.path, conf.GetString(settingOutputMergedConfigurationsWriteEcs.path))
	cmd.mergedConfigurationsWriteEcs = conf.GetBool(settingOutputMergedConfigurationsWriteEcs.path)

	// output: update package directory
	log.Debugf("Setting '%s': '%s'", settingOutputUpdatePackagesPath.path, conf.GetString(settingOutputUpdatePackagesPath.path))
	cmd.updatePackageDirectory = conf.GetString(settingOutputUpdatePackagesPath.path)
	if len(cmd.updatePackageDirectory) > 0 { // setting is optional
		if filepath.IsAbs(cmd.updatePackageDirectory) {
			cmd.updatePackageDirectory = filepath.Clean(cmd.updatePackageDirectory)
		} else {
			path, err := filepath.Abs(filepath.Join(configDir, cmd.updatePackageDirectory))
			if err != nil {
				return err
			}
			cmd.updatePackageDirectory = path
		}
	}

	// log configuration
	logtext := strings.Builder{}
	logtext.WriteString(fmt.Sprintf("--- Configuration ---\n"))
	logtext.WriteString(fmt.Sprintf("SD Card Template Directory:       %s\n", cmd.sdcardTemplateDirectory))
	logtext.WriteString(fmt.Sprintf("Base Configuration File:          %s\n", cmd.baseConfigurationPath))
	logtext.WriteString(fmt.Sprintf("Merge Configuration File:         %s\n", cmd.mergeConfigurationPath))
	logtext.WriteString(fmt.Sprintf("Hot folder:                       %s\n", cmd.hotFolderPath))
	logtext.WriteString(fmt.Sprintf("Merged Configuration Directory:   %s\n", cmd.mergedConfigurationDirectory))
	logtext.WriteString(fmt.Sprintf("  - Write ATV:                    %v\n", cmd.mergedConfigurationsWriteAtv))
	logtext.WriteString(fmt.Sprintf("  - Write ECS:                    %v\n", cmd.mergedConfigurationsWriteEcs))
	logtext.WriteString(fmt.Sprintf("Update Package Directory:         %s\n", cmd.updatePackageDirectory))
	logtext.WriteString(fmt.Sprintf("--- Configuration End ---"))
	log.Info(logtext.String())

	return nil
}
