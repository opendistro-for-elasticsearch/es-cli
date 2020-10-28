/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package commands

import (
	handler "es-cli/odfe-cli/handler/ad"
	"fmt"

	"github.com/spf13/cobra"
)

const (
	updateDetectorsCommandName = "update"
	forceFlagName              = "force"
	startFlagName              = "start"
)

//updateDetectorsCmd updates detectors with configuration from input file
var updateDetectorsCmd = &cobra.Command{
	Use:   updateDetectorsCommandName + " json-file-path ... [flags]",
	Short: "Update detectors based on JSON files",
	Long: fmt.Sprintf("Description:\n  " +
		"Update detectors based on JSON files. To begin, use `odfe-cli ad get detector-name > detector_to_be_updated.json` to download detector and update it for your use case." +
		"Then use `odfe-cli ad update file-path` to update detector"),
	Run: func(cmd *cobra.Command, args []string) {
		//If no args, display usage
		if len(args) < 1 {
			if err := cmd.Usage(); err != nil {
				fmt.Println(err)
			}
			return
		}
		force, _ := cmd.Flags().GetBool(forceFlagName)
		start, _ := cmd.Flags().GetBool(startFlagName)
		err := updateDetectors(args, force, start)
		if err != nil {
			DisplayError(err, updateDetectorsCommandName)
		}
	},
}

func init() {
	GetADCommand().AddCommand(updateDetectorsCmd)
	updateDetectorsCmd.Flags().BoolP(forceFlagName, "f", false, "Stop detector and update forcefully")
	updateDetectorsCmd.Flags().BoolP(startFlagName, "s", false, "Start detector if update is successful")
	updateDetectorsCmd.Flags().StringP(flagProfileName, "p", "", "Use a specific profile from your configuration file.")
	updateDetectorsCmd.Flags().BoolP("help", "h", false, "Help for "+updateDetectorsCommandName)
}

func updateDetectors(fileNames []string, force bool, start bool) error {
	commandHandler, err := GetADHandler()
	if err != nil {
		return err
	}
	for _, name := range fileNames {
		err = handler.UpdateAnomalyDetector(commandHandler, name, force, start)
		if err != nil {
			return err
		}
	}
	return nil
}
