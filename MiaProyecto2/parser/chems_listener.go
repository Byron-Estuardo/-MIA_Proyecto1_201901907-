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

	// EnterPmkdisk is called when entering the pmkdisk production.
	EnterPmkdisk(c *PmkdiskContext)

	// EnterParametrosmkdisk is called when entering the parametrosmkdisk production.
	EnterParametrosmkdisk(c *ParametrosmkdiskContext)

	// EnterPrmdisk is called when entering the prmdisk production.
	EnterPrmdisk(c *PrmdiskContext)

	// EnterPfdisk is called when entering the pfdisk production.
	EnterPfdisk(c *PfdiskContext)

	// EnterParametrosfdisk is called when entering the parametrosfdisk production.
	EnterParametrosfdisk(c *ParametrosfdiskContext)

	// EnterPmount is called when entering the pmount production.
	EnterPmount(c *PmountContext)

	// EnterParametrosmount is called when entering the parametrosmount production.
	EnterParametrosmount(c *ParametrosmountContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitAdminDiscos is called when exiting the adminDiscos production.
	ExitAdminDiscos(c *AdminDiscosContext)

	// ExitPmkdisk is called when exiting the pmkdisk production.
	ExitPmkdisk(c *PmkdiskContext)

	// ExitParametrosmkdisk is called when exiting the parametrosmkdisk production.
	ExitParametrosmkdisk(c *ParametrosmkdiskContext)

	// ExitPrmdisk is called when exiting the prmdisk production.
	ExitPrmdisk(c *PrmdiskContext)

	// ExitPfdisk is called when exiting the pfdisk production.
	ExitPfdisk(c *PfdiskContext)

	// ExitParametrosfdisk is called when exiting the parametrosfdisk production.
	ExitParametrosfdisk(c *ParametrosfdiskContext)

	// ExitPmount is called when exiting the pmount production.
	ExitPmount(c *PmountContext)

	// ExitParametrosmount is called when exiting the parametrosmount production.
	ExitParametrosmount(c *ParametrosmountContext)
}
