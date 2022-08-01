package core

import (
	// "os"

	"git.bitcubix.io/go/validation"
)

func (c *Core) newValidator() *validation.Validator {

	validator := validation.NewValidator()

	// Todo: add return err to validation.NewValidator()
	// validator, err := validation.NewValidator()
	// if err != nil {
	// 	log.Error().Err(err).Msg("Setup validator error")
	// 	os.Exit(2)
	// }

	// Todo: Setup Validator --> check stuff in git.bitcubix.io/go/validation
	// Todo: use validator in example and user service

	c.Log.Info().Msg("Setup validator")
	c.Log.Warn().Msg("TODO: Setup validator")
	return validator
}
