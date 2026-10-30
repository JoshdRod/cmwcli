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

func NewCollectionCmd() *cobra.Command {
	cmd := &cobra.Command {
		Use:   "collection",
		Short: "Create a Conceptual Message Wrapper for a (CBOR only) collection",
		Long: `Create a Conceptual Message Wrapper for a (CBOR only) collection
	For example:

	cmwcli collection --data collection.cbor`,
		Run: func(cmd *cobra.Command, args []string) {
			// Read flags
			data, err := cmd.Flags().GetStringSlice("data")
			if err != nil || data == nil {
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
				// If only one file, create new CMW
				// Create a record on that CMW
				// Add media type = application/eat+cbor
				// Add value = hash of file
				// Add indicator = 4 (evidence)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// collectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// collectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//collectionCmd.Flags().StringSliceP("data", "d", nil, "Location of input data")
	//collectionCmd.Flags().StringSliceP("label", "l", nil, "Label for input data")
	//collectionCmd.Flags().StringP("output", "o", "", "Location of output file")
	//collectionCmd.Flags().StringP("cmwc_t", "c", "", "The __cmwc_t of the parent collection")
}
