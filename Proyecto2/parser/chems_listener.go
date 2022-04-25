// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChemsListener is a complete listener for a parse tree produced by Chems.
type ChemsListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterAdminDiscos is called when entering the adminDiscos production.
	EnterAdminDiscos(c *AdminDiscosContext)

	// EnterMkdisk is called when entering the mkdisk production.
	EnterMkdisk(c *MkdiskContext)

	// EnterParametrosmkdisk is called when entering the parametrosmkdisk production.
	EnterParametrosmkdisk(c *ParametrosmkdiskContext)

	// EnterAjuste is called when entering the ajuste production.
	EnterAjuste(c *AjusteContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitAdminDiscos is called when exiting the adminDiscos production.
	ExitAdminDiscos(c *AdminDiscosContext)

	// ExitMkdisk is called when exiting the mkdisk production.
	ExitMkdisk(c *MkdiskContext)

	// ExitParametrosmkdisk is called when exiting the parametrosmkdisk production.
	ExitParametrosmkdisk(c *ParametrosmkdiskContext)

	// ExitAjuste is called when exiting the ajuste production.
	ExitAjuste(c *AjusteContext)
}
