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
	"encoding/json"
	entity "es-cli/odfe-cli/entity/ad"
	"es-cli/odfe-cli/handler/ad"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

const (
	getDetectorsCommandName = "get"
	getDetectorIDFlagName   = "id"
)

//getDetectorsCmd prints detectors configuration based on id, name or name regex pattern.
//default input is name pattern, one can change this format to be id by passing --id flag
var getDetectorsCmd = &cobra.Command{
	Use:   getDetectorsCommandName + " detector_name ..." + " [flags] ",
	Short: "Get detectors based on list of ids, names or name regex patterns ",
	Long: fmt.Sprintf("Description:\n  " +
		`Get detectors based on list of user input. Use "" to make sure the name does not match with pwd lists'.
  The default input is detector name, use --id flag if input is detector id instead of name`),
	Run: func(cmd *cobra.Command, args []string) {
		//If no args, display usage
		if len(args) < 1 {
			fmt.Println(cmd.Usage())
			return
		}
		err := printDetectors(Println, cmd, args)
		if err != nil {
			DisplayError(err, getDetectorsCommandName)
		}
	},
}

type Display func(*cobra.Command, *entity.DetectorOutput) error

//printDetectors print detectors
func printDetectors(display Display, cmd *cobra.Command, detectors []string) error {
	idStatus, _ := cmd.Flags().GetBool(getDetectorIDFlagName)
	commandHandler, err := GetADHandler()
	if err != nil {
		return err
	}
	// default is name
	action := ad.GetAnomalyDetectorsByNamePattern
	if idStatus {
		action = getDetectorsByID
	}
	results, err := getDetectors(commandHandler, detectors, action)
	if err != nil {
		return err
	}
	return fprint(cmd, display, results)
}

//getDetectors fetch detector from controller
func getDetectors(
	commandHandler *ad.Handler, args []string, get func(*ad.Handler, string) (
		[]*entity.DetectorOutput, error)) ([]*entity.DetectorOutput, error) {
	var results []*entity.DetectorOutput
	for _, detector := range args {
		output, err := get(commandHandler, detector)
		if err != nil {
			return nil, err
		}
		results = append(results, output...)
	}
	return results, nil
}

//getDetectorsByID gets detector output based on ID as argument
func getDetectorsByID(commandHandler *ad.Handler, ID string) ([]*entity.DetectorOutput, error) {

	output, err := ad.GetAnomalyDetectorByID(commandHandler, ID)
	if err != nil {
		return nil, err
	}
	return []*entity.DetectorOutput{output}, nil
}

//fprint displays the list of detectors.
func fprint(cmd *cobra.Command, display Display, results []*entity.DetectorOutput) error {
	if results == nil {
		return nil
	}
	for _, d := range results {
		if err := display(cmd, d); err != nil {
			return err
		}
	}
	return nil
}

//FPrint prints detector configuration on writer
//Since this is json format, use indent function to pretty print before printing on writer
func FPrint(writer io.Writer, d *entity.DetectorOutput) error {
	formattedOutput, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(writer, string(formattedOutput))
	return err
}

//Println prints detector configuration on stdout
func Println(cmd *cobra.Command, d *entity.DetectorOutput) error {
	return FPrint(os.Stdout, d)
}

func init() {
	GetADCommand().AddCommand(getDetectorsCmd)
	getDetectorsCmd.Flags().BoolP(getDetectorIDFlagName, "", false, "input is detector's id")
}
