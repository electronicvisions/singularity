// Copyright (c) 2017-2019, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package cli

import (
	"github.com/spf13/cobra"
	"github.com/sylabs/singularity/docs"
	"github.com/sylabs/singularity/internal/pkg/sylog"
	"github.com/sylabs/singularity/pkg/sypgp"
)

func init() {
	KeyNewPairCmd.Flags().SetInterspersed(false)
}

// KeyNewPairCmd is `singularity key newpair' and generate a new OpenPGP key pair
var KeyNewPairCmd = &cobra.Command{
	Args:                  cobra.ExactArgs(0),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := sypgp.GenKeyPair(); err != nil {
			sylog.Fatalf("creating newpair failed: %v", err)
		}
	},

	Use:     docs.KeyNewPairUse,
	Short:   docs.KeyNewPairShort,
	Long:    docs.KeyNewPairLong,
	Example: docs.KeyNewPairExample,
}
