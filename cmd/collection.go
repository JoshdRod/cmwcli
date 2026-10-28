/*
Copyright © 2026 Josh Rodriguez josh.rodriguez@arm.com
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/veraison/cmw"
	"os"
)

// collectionCmd represents the collection command
var collectionCmd = NewCollectionCmd()

/*
Create a new CMW from the evidence files input.
If only one evidence file input, then create a CMW with a record.
If multiple evidence files input, then create a CMW with a collection of records.
All records have:

	media type = application/eat+cbor
	value = hash of evidence file
	indicator = 4 (evidence)
*/
func NewCollectionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collection",
		Short: "Create a Conceptual Message Wrapper from (CBOR only) evidence file(s)",
		Long: `Create a Conceptual Message Wrapper from (CBOR only) evidence file(s)
	For example:

	cmwcli collection --data evidence.cbor --label "evidence" --cmwc_t "https://example.com" --output out.cbor`,
		Run: func(cmd *cobra.Command, args []string) {
			// Read flags
			data, err := cmd.Flags().GetStringSlice("data")
			if err != nil || data == nil || len(data) == 0 {
				fmt.Println("No data inputted")
				return
			}

			label, err := cmd.Flags().GetStringSlice("label")
			if err != nil {
				fmt.Println("Error parsing labels")
				fmt.Println(err)
				return
			}
			if len(data) > 1 && len(data) != len(label) {
				fmt.Println("Failed to create container - Non-equal amount of files and labels input.")
				fmt.Println("Given", len(data), "files, but", len(label), "labels.")
				fmt.Println("Each file must have its own label.")
				return
			}

			output, err := cmd.Flags().GetString("output")
			if err != nil || output == "" {
				fmt.Println("No output file specified")
				return
			}

			cmwc_t, err := cmd.Flags().GetString("cmwc_t")
			if err != nil {
				fmt.Println("Invalid __cmwc_t specified")
				fmt.Println(err)
				return
			}

			var cmwOutput cmw.CMW
			// Check how many files there are
			if len(data) > 1 {
				cmwOutputRef, err := cmw.NewCollection(cmwc_t)
				if err != nil {
					fmt.Println("Error creating collection")
					fmt.Println(err)
					return
				}
				cmwOutput = *cmwOutputRef

				for i, d := range data {
					file, err := os.ReadFile(d)
					if err != nil {
						fmt.Println("Error accessing cbor data in file", d)
						fmt.Println(err)
						return
					}

					newCMWItem, err := cmw.NewMonad("application/eat+cbor", file, 4)
					err = cmwOutput.AddCollectionItem(label[i], newCMWItem)
					if err != nil {
						fmt.Println("Error adding item from file", d, "to CMW collection")
						fmt.Println(err)
						return
					}
				}
			} else {
				file, err := os.ReadFile(data[0])
				if err != nil {
					fmt.Println("Error accessing cbor data in file", data[0])
					fmt.Println(err)
					return
				}
				cmwOutputRef, err := cmw.NewMonad("application/eat+cbor", file, 4)
				if err != nil {
					fmt.Println("Error creating CMW Record from", data[0])
					fmt.Println(err)
					return
				}
				cmwOutput = *cmwOutputRef
			}

			// Write to output file
			fmt.Println(cmwOutput)

			cmwBytes, err := cmwOutput.MarshalCBOR()
			if err != nil {
				fmt.Println("Error marshalling CMW collection to a CBOR file. Maybe check the input file is valid CBOR?")
				return
			}

			err = os.WriteFile(output, cmwBytes, 0666)
			if err != nil {
				fmt.Println("Error writing to file")
				return
			}
		},
	}
	cmd.Flags().StringSliceP("data", "d", nil, "Location of input data")
	cmd.Flags().StringSliceP("label", "l", nil, "Label for input data")
	cmd.Flags().StringP("output", "o", "", "Location of output file")
	cmd.Flags().StringP("cmwc_t", "c", "", "The __cmwc_t of the parent collection")
	return cmd
}

func init() {
	rootCmd.AddCommand(collectionCmd)
}
