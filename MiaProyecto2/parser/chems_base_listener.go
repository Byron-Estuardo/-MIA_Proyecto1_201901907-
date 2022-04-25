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
