// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseChemsListener is a complete listener for a parse tree produced by Chems.
type BaseChemsListener struct{}

var _ ChemsListener = &BaseChemsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChemsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChemsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChemsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChemsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseChemsListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseChemsListener) ExitStart(ctx *StartContext) {}

// EnterAdminDiscos is called when production adminDiscos is entered.
func (s *BaseChemsListener) EnterAdminDiscos(ctx *AdminDiscosContext) {}

// ExitAdminDiscos is called when production adminDiscos is exited.
func (s *BaseChemsListener) ExitAdminDiscos(ctx *AdminDiscosContext) {}

// EnterAdminArchivos is called when production adminArchivos is entered.
func (s *BaseChemsListener) EnterAdminArchivos(ctx *AdminArchivosContext) {}

// ExitAdminArchivos is called when production adminArchivos is exited.
func (s *BaseChemsListener) ExitAdminArchivos(ctx *AdminArchivosContext) {}

// EnterAdminUsuarios is called when production adminUsuarios is entered.
func (s *BaseChemsListener) EnterAdminUsuarios(ctx *AdminUsuariosContext) {}

// ExitAdminUsuarios is called when production adminUsuarios is exited.
func (s *BaseChemsListener) ExitAdminUsuarios(ctx *AdminUsuariosContext) {}

// EnterAdminCarpetas is called when production adminCarpetas is entered.
func (s *BaseChemsListener) EnterAdminCarpetas(ctx *AdminCarpetasContext) {}

// ExitAdminCarpetas is called when production adminCarpetas is exited.
func (s *BaseChemsListener) ExitAdminCarpetas(ctx *AdminCarpetasContext) {}

// EnterPmkfile is called when production pmkfile is entered.
func (s *BaseChemsListener) EnterPmkfile(ctx *PmkfileContext) {}

// ExitPmkfile is called when production pmkfile is exited.
func (s *BaseChemsListener) ExitPmkfile(ctx *PmkfileContext) {}

// EnterParametrosmkfile is called when production parametrosmkfile is entered.
func (s *BaseChemsListener) EnterParametrosmkfile(ctx *ParametrosmkfileContext) {}

// ExitParametrosmkfile is called when production parametrosmkfile is exited.
func (s *BaseChemsListener) ExitParametrosmkfile(ctx *ParametrosmkfileContext) {}

// EnterPmkdisk is called when production pmkdisk is entered.
func (s *BaseChemsListener) EnterPmkdisk(ctx *PmkdiskContext) {}

// ExitPmkdisk is called when production pmkdisk is exited.
func (s *BaseChemsListener) ExitPmkdisk(ctx *PmkdiskContext) {}

// EnterParametrosmkdisk is called when production parametrosmkdisk is entered.
func (s *BaseChemsListener) EnterParametrosmkdisk(ctx *ParametrosmkdiskContext) {}

// ExitParametrosmkdisk is called when production parametrosmkdisk is exited.
func (s *BaseChemsListener) ExitParametrosmkdisk(ctx *ParametrosmkdiskContext) {}

// EnterPrmdisk is called when production prmdisk is entered.
func (s *BaseChemsListener) EnterPrmdisk(ctx *PrmdiskContext) {}

// ExitPrmdisk is called when production prmdisk is exited.
func (s *BaseChemsListener) ExitPrmdisk(ctx *PrmdiskContext) {}

// EnterPmkgrp is called when production pmkgrp is entered.
func (s *BaseChemsListener) EnterPmkgrp(ctx *PmkgrpContext) {}

// ExitPmkgrp is called when production pmkgrp is exited.
func (s *BaseChemsListener) ExitPmkgrp(ctx *PmkgrpContext) {}

// EnterPrmkgrp is called when production prmkgrp is entered.
func (s *BaseChemsListener) EnterPrmkgrp(ctx *PrmkgrpContext) {}

// ExitPrmkgrp is called when production prmkgrp is exited.
func (s *BaseChemsListener) ExitPrmkgrp(ctx *PrmkgrpContext) {}

// EnterPrmkusr is called when production prmkusr is entered.
func (s *BaseChemsListener) EnterPrmkusr(ctx *PrmkusrContext) {}

// ExitPrmkusr is called when production prmkusr is exited.
func (s *BaseChemsListener) ExitPrmkusr(ctx *PrmkusrContext) {}

// EnterPfdisk is called when production pfdisk is entered.
func (s *BaseChemsListener) EnterPfdisk(ctx *PfdiskContext) {}

// ExitPfdisk is called when production pfdisk is exited.
func (s *BaseChemsListener) ExitPfdisk(ctx *PfdiskContext) {}

// EnterParametrosfdisk is called when production parametrosfdisk is entered.
func (s *BaseChemsListener) EnterParametrosfdisk(ctx *ParametrosfdiskContext) {}

// ExitParametrosfdisk is called when production parametrosfdisk is exited.
func (s *BaseChemsListener) ExitParametrosfdisk(ctx *ParametrosfdiskContext) {}

// EnterPmount is called when production pmount is entered.
func (s *BaseChemsListener) EnterPmount(ctx *PmountContext) {}

// ExitPmount is called when production pmount is exited.
func (s *BaseChemsListener) ExitPmount(ctx *PmountContext) {}

// EnterParametrosmount is called when production parametrosmount is entered.
func (s *BaseChemsListener) EnterParametrosmount(ctx *ParametrosmountContext) {}

// ExitParametrosmount is called when production parametrosmount is exited.
func (s *BaseChemsListener) ExitParametrosmount(ctx *ParametrosmountContext) {}

// EnterPmkfs is called when production pmkfs is entered.
func (s *BaseChemsListener) EnterPmkfs(ctx *PmkfsContext) {}

// ExitPmkfs is called when production pmkfs is exited.
func (s *BaseChemsListener) ExitPmkfs(ctx *PmkfsContext) {}

// EnterParametrosmkfs is called when production parametrosmkfs is entered.
func (s *BaseChemsListener) EnterParametrosmkfs(ctx *ParametrosmkfsContext) {}

// ExitParametrosmkfs is called when production parametrosmkfs is exited.
func (s *BaseChemsListener) ExitParametrosmkfs(ctx *ParametrosmkfsContext) {}

// EnterPlogin is called when production plogin is entered.
func (s *BaseChemsListener) EnterPlogin(ctx *PloginContext) {}

// ExitPlogin is called when production plogin is exited.
func (s *BaseChemsListener) ExitPlogin(ctx *PloginContext) {}

// EnterParametroslogin is called when production parametroslogin is entered.
func (s *BaseChemsListener) EnterParametroslogin(ctx *ParametrosloginContext) {}

// ExitParametroslogin is called when production parametroslogin is exited.
func (s *BaseChemsListener) ExitParametroslogin(ctx *ParametrosloginContext) {}

// EnterPmkuser is called when production pmkuser is entered.
func (s *BaseChemsListener) EnterPmkuser(ctx *PmkuserContext) {}

// ExitPmkuser is called when production pmkuser is exited.
func (s *BaseChemsListener) ExitPmkuser(ctx *PmkuserContext) {}

// EnterParametrosmkuser is called when production parametrosmkuser is entered.
func (s *BaseChemsListener) EnterParametrosmkuser(ctx *ParametrosmkuserContext) {}

// ExitParametrosmkuser is called when production parametrosmkuser is exited.
func (s *BaseChemsListener) ExitParametrosmkuser(ctx *ParametrosmkuserContext) {}
