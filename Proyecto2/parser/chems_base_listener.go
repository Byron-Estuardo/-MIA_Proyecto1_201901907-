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

// EnterMkdisk is called when production mkdisk is entered.
func (s *BaseChemsListener) EnterMkdisk(ctx *MkdiskContext) {}

// ExitMkdisk is called when production mkdisk is exited.
func (s *BaseChemsListener) ExitMkdisk(ctx *MkdiskContext) {}

// EnterParametrosmkdisk is called when production parametrosmkdisk is entered.
func (s *BaseChemsListener) EnterParametrosmkdisk(ctx *ParametrosmkdiskContext) {}

// ExitParametrosmkdisk is called when production parametrosmkdisk is exited.
func (s *BaseChemsListener) ExitParametrosmkdisk(ctx *ParametrosmkdiskContext) {}

// EnterAjuste is called when production ajuste is entered.
func (s *BaseChemsListener) EnterAjuste(ctx *AjusteContext) {}

// ExitAjuste is called when production ajuste is exited.
func (s *BaseChemsListener) ExitAjuste(ctx *AjusteContext) {}
