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

	// EnterAdminArchivos is called when entering the adminArchivos production.
	EnterAdminArchivos(c *AdminArchivosContext)

	// EnterAdminUsuarios is called when entering the adminUsuarios production.
	EnterAdminUsuarios(c *AdminUsuariosContext)

	// EnterAdminCarpetas is called when entering the adminCarpetas production.
	EnterAdminCarpetas(c *AdminCarpetasContext)

	// EnterPmkfile is called when entering the pmkfile production.
	EnterPmkfile(c *PmkfileContext)

	// EnterParametrosmkfile is called when entering the parametrosmkfile production.
	EnterParametrosmkfile(c *ParametrosmkfileContext)

	// EnterPmkdisk is called when entering the pmkdisk production.
	EnterPmkdisk(c *PmkdiskContext)

	// EnterParametrosmkdisk is called when entering the parametrosmkdisk production.
	EnterParametrosmkdisk(c *ParametrosmkdiskContext)

	// EnterPrmdisk is called when entering the prmdisk production.
	EnterPrmdisk(c *PrmdiskContext)

	// EnterPmkgrp is called when entering the pmkgrp production.
	EnterPmkgrp(c *PmkgrpContext)

	// EnterPrmkgrp is called when entering the prmkgrp production.
	EnterPrmkgrp(c *PrmkgrpContext)

	// EnterPrmkusr is called when entering the prmkusr production.
	EnterPrmkusr(c *PrmkusrContext)

	// EnterPfdisk is called when entering the pfdisk production.
	EnterPfdisk(c *PfdiskContext)

	// EnterParametrosfdisk is called when entering the parametrosfdisk production.
	EnterParametrosfdisk(c *ParametrosfdiskContext)

	// EnterPmount is called when entering the pmount production.
	EnterPmount(c *PmountContext)

	// EnterParametrosmount is called when entering the parametrosmount production.
	EnterParametrosmount(c *ParametrosmountContext)

	// EnterPmkfs is called when entering the pmkfs production.
	EnterPmkfs(c *PmkfsContext)

	// EnterParametrosmkfs is called when entering the parametrosmkfs production.
	EnterParametrosmkfs(c *ParametrosmkfsContext)

	// EnterPlogin is called when entering the plogin production.
	EnterPlogin(c *PloginContext)

	// EnterParametroslogin is called when entering the parametroslogin production.
	EnterParametroslogin(c *ParametrosloginContext)

	// EnterPmkuser is called when entering the pmkuser production.
	EnterPmkuser(c *PmkuserContext)

	// EnterParametrosmkuser is called when entering the parametrosmkuser production.
	EnterParametrosmkuser(c *ParametrosmkuserContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitAdminDiscos is called when exiting the adminDiscos production.
	ExitAdminDiscos(c *AdminDiscosContext)

	// ExitAdminArchivos is called when exiting the adminArchivos production.
	ExitAdminArchivos(c *AdminArchivosContext)

	// ExitAdminUsuarios is called when exiting the adminUsuarios production.
	ExitAdminUsuarios(c *AdminUsuariosContext)

	// ExitAdminCarpetas is called when exiting the adminCarpetas production.
	ExitAdminCarpetas(c *AdminCarpetasContext)

	// ExitPmkfile is called when exiting the pmkfile production.
	ExitPmkfile(c *PmkfileContext)

	// ExitParametrosmkfile is called when exiting the parametrosmkfile production.
	ExitParametrosmkfile(c *ParametrosmkfileContext)

	// ExitPmkdisk is called when exiting the pmkdisk production.
	ExitPmkdisk(c *PmkdiskContext)

	// ExitParametrosmkdisk is called when exiting the parametrosmkdisk production.
	ExitParametrosmkdisk(c *ParametrosmkdiskContext)

	// ExitPrmdisk is called when exiting the prmdisk production.
	ExitPrmdisk(c *PrmdiskContext)

	// ExitPmkgrp is called when exiting the pmkgrp production.
	ExitPmkgrp(c *PmkgrpContext)

	// ExitPrmkgrp is called when exiting the prmkgrp production.
	ExitPrmkgrp(c *PrmkgrpContext)

	// ExitPrmkusr is called when exiting the prmkusr production.
	ExitPrmkusr(c *PrmkusrContext)

	// ExitPfdisk is called when exiting the pfdisk production.
	ExitPfdisk(c *PfdiskContext)

	// ExitParametrosfdisk is called when exiting the parametrosfdisk production.
	ExitParametrosfdisk(c *ParametrosfdiskContext)

	// ExitPmount is called when exiting the pmount production.
	ExitPmount(c *PmountContext)

	// ExitParametrosmount is called when exiting the parametrosmount production.
	ExitParametrosmount(c *ParametrosmountContext)

	// ExitPmkfs is called when exiting the pmkfs production.
	ExitPmkfs(c *PmkfsContext)

	// ExitParametrosmkfs is called when exiting the parametrosmkfs production.
	ExitParametrosmkfs(c *ParametrosmkfsContext)

	// ExitPlogin is called when exiting the plogin production.
	ExitPlogin(c *PloginContext)

	// ExitParametroslogin is called when exiting the parametroslogin production.
	ExitParametroslogin(c *ParametrosloginContext)

	// ExitPmkuser is called when exiting the pmkuser production.
	ExitPmkuser(c *PmkuserContext)

	// ExitParametrosmkuser is called when exiting the parametrosmkuser production.
	ExitParametrosmkuser(c *ParametrosmkuserContext)
}
