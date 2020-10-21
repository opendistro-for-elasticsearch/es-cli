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
	Use:   updateDetectorsCommandName + " [detector-configuration-file-path ...] [flags]",
	Short: "Update detectors with configuration from input file",
	Long: fmt.Sprintf("Description:\n  " +
		`Update detectors with configuration from input file.
  Use 'odfe-cli ad download detector-name' to download configurations. Make necessary changes to file and update detector using 'odfe-cli ad update file-path'`),
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
	updateDetectorsCmd.Flags().BoolP(forceFlagName, "f", false, "stop detector and update forcefully")
	updateDetectorsCmd.Flags().BoolP(startFlagName, "s", false, "start detector if update is successful")
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
