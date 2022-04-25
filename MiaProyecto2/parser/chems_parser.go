// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import AD "MiaProyecto2/AdminDisks"

var ValSize string = ""
var ValFit string = ""
var ValUnit string = ""
var ValPath string = ""
var ValType string = ""
var ValName string = ""

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 48, 129,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 25,
	10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 35, 10, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 45, 10, 4, 12, 4,
	14, 4, 48, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 74, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 90, 10, 7, 12, 7,
	14, 7, 93, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	5, 8, 104, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9,
	114, 10, 9, 12, 9, 14, 9, 117, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 5, 10, 127, 10, 10, 3, 10, 2, 5, 6, 12, 16, 11, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 2, 2, 2, 135, 2, 24, 3, 2, 2, 2, 4, 34, 3,
	2, 2, 2, 6, 36, 3, 2, 2, 2, 8, 73, 3, 2, 2, 2, 10, 75, 3, 2, 2, 2, 12,
	81, 3, 2, 2, 2, 14, 103, 3, 2, 2, 2, 16, 105, 3, 2, 2, 2, 18, 126, 3, 2,
	2, 2, 20, 25, 5, 4, 3, 2, 21, 22, 7, 37, 2, 2, 22, 25, 8, 2, 1, 2, 23,
	25, 7, 2, 2, 3, 24, 20, 3, 2, 2, 2, 24, 21, 3, 2, 2, 2, 24, 23, 3, 2, 2,
	2, 25, 3, 3, 2, 2, 2, 26, 27, 7, 3, 2, 2, 27, 35, 5, 6, 4, 2, 28, 29, 7,
	4, 2, 2, 29, 35, 5, 10, 6, 2, 30, 31, 7, 5, 2, 2, 31, 35, 5, 12, 7, 2,
	32, 33, 7, 6, 2, 2, 33, 35, 5, 16, 9, 2, 34, 26, 3, 2, 2, 2, 34, 28, 3,
	2, 2, 2, 34, 30, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 35, 5, 3, 2, 2, 2, 36,
	37, 8, 4, 1, 2, 37, 38, 5, 8, 5, 2, 38, 39, 8, 4, 1, 2, 39, 46, 3, 2, 2,
	2, 40, 41, 12, 4, 2, 2, 41, 42, 5, 8, 5, 2, 42, 43, 8, 4, 1, 2, 43, 45,
	3, 2, 2, 2, 44, 40, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2,
	46, 47, 3, 2, 2, 2, 47, 7, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 7, 18,
	2, 2, 50, 51, 7, 46, 2, 2, 51, 52, 7, 42, 2, 2, 52, 74, 8, 5, 1, 2, 53,
	54, 7, 19, 2, 2, 54, 55, 7, 46, 2, 2, 55, 56, 7, 25, 2, 2, 56, 74, 8, 5,
	1, 2, 57, 58, 7, 19, 2, 2, 58, 59, 7, 46, 2, 2, 59, 60, 7, 26, 2, 2, 60,
	74, 8, 5, 1, 2, 61, 62, 7, 19, 2, 2, 62, 63, 7, 46, 2, 2, 63, 64, 7, 27,
	2, 2, 64, 74, 8, 5, 1, 2, 65, 66, 7, 20, 2, 2, 66, 67, 7, 46, 2, 2, 67,
	68, 7, 43, 2, 2, 68, 74, 8, 5, 1, 2, 69, 70, 7, 21, 2, 2, 70, 71, 7, 46,
	2, 2, 71, 72, 7, 45, 2, 2, 72, 74, 8, 5, 1, 2, 73, 49, 3, 2, 2, 2, 73,
	53, 3, 2, 2, 2, 73, 57, 3, 2, 2, 2, 73, 61, 3, 2, 2, 2, 73, 65, 3, 2, 2,
	2, 73, 69, 3, 2, 2, 2, 74, 9, 3, 2, 2, 2, 75, 76, 7, 4, 2, 2, 76, 77, 7,
	21, 2, 2, 77, 78, 7, 46, 2, 2, 78, 79, 7, 45, 2, 2, 79, 80, 8, 6, 1, 2,
	80, 11, 3, 2, 2, 2, 81, 82, 8, 7, 1, 2, 82, 83, 5, 14, 8, 2, 83, 84, 8,
	7, 1, 2, 84, 91, 3, 2, 2, 2, 85, 86, 12, 4, 2, 2, 86, 87, 5, 14, 8, 2,
	87, 88, 8, 7, 1, 2, 88, 90, 3, 2, 2, 2, 89, 85, 3, 2, 2, 2, 90, 93, 3,
	2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 13, 3, 2, 2, 2, 93,
	91, 3, 2, 2, 2, 94, 104, 5, 8, 5, 2, 95, 96, 7, 23, 2, 2, 96, 97, 7, 46,
	2, 2, 97, 98, 7, 43, 2, 2, 98, 104, 8, 8, 1, 2, 99, 100, 7, 22, 2, 2, 100,
	101, 7, 46, 2, 2, 101, 102, 7, 43, 2, 2, 102, 104, 8, 8, 1, 2, 103, 94,
	3, 2, 2, 2, 103, 95, 3, 2, 2, 2, 103, 99, 3, 2, 2, 2, 104, 15, 3, 2, 2,
	2, 105, 106, 8, 9, 1, 2, 106, 107, 5, 18, 10, 2, 107, 108, 8, 9, 1, 2,
	108, 115, 3, 2, 2, 2, 109, 110, 12, 4, 2, 2, 110, 111, 5, 18, 10, 2, 111,
	112, 8, 9, 1, 2, 112, 114, 3, 2, 2, 2, 113, 109, 3, 2, 2, 2, 114, 117,
	3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 17, 3, 2,
	2, 2, 117, 115, 3, 2, 2, 2, 118, 119, 7, 21, 2, 2, 119, 120, 7, 46, 2,
	2, 120, 121, 7, 45, 2, 2, 121, 127, 8, 10, 1, 2, 122, 123, 7, 22, 2, 2,
	123, 124, 7, 46, 2, 2, 124, 125, 7, 43, 2, 2, 125, 127, 8, 10, 1, 2, 126,
	118, 3, 2, 2, 2, 126, 122, 3, 2, 2, 2, 127, 19, 3, 2, 2, 2, 10, 24, 34,
	46, 73, 91, 103, 115, 126,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "'='",
}
var symbolicNames = []string{
	"", "MKDISK", "RMDISK", "FDISK", "MOUNT", "MKFS", "LOGIN", "LOGOUT", "MKGRP",
	"RMGRP", "MKUSER", "RMUSR", "MKFILE", "MKDIR", "EXEC", "REP", "SIZE", "FIT",
	"UNIT", "PATH", "NAME", "TYPE", "ID", "BF", "FF", "WF", "FAST", "FULL",
	"USR", "PWD", "PWD1", "GRP", "GR", "CONT", "GP", "PAUSA", "DISK", "TREE",
	"FILE", "LETRA", "ENTERO", "CADENA", "IDENTIFICADOR", "RUTA", "IGUAL",
	"WHITESPACE", "COMENTARIOS",
}

var ruleNames = []string{
	"start", "adminDiscos", "pmkdisk", "parametrosmkdisk", "prmdisk", "pfdisk",
	"parametrosfdisk", "pmount", "parametrosmount",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Chems struct {
	*antlr.BaseParser
}

func NewChems(input antlr.TokenStream) *Chems {
	this := new(Chems)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Chems.g4"

	return this
}

// Chems tokens.
const (
	ChemsEOF           = antlr.TokenEOF
	ChemsMKDISK        = 1
	ChemsRMDISK        = 2
	ChemsFDISK         = 3
	ChemsMOUNT         = 4
	ChemsMKFS          = 5
	ChemsLOGIN         = 6
	ChemsLOGOUT        = 7
	ChemsMKGRP         = 8
	ChemsRMGRP         = 9
	ChemsMKUSER        = 10
	ChemsRMUSR         = 11
	ChemsMKFILE        = 12
	ChemsMKDIR         = 13
	ChemsEXEC          = 14
	ChemsREP           = 15
	ChemsSIZE          = 16
	ChemsFIT           = 17
	ChemsUNIT          = 18
	ChemsPATH          = 19
	ChemsNAME          = 20
	ChemsTYPE          = 21
	ChemsID            = 22
	ChemsBF            = 23
	ChemsFF            = 24
	ChemsWF            = 25
	ChemsFAST          = 26
	ChemsFULL          = 27
	ChemsUSR           = 28
	ChemsPWD           = 29
	ChemsPWD1          = 30
	ChemsGRP           = 31
	ChemsGR            = 32
	ChemsCONT          = 33
	ChemsGP            = 34
	ChemsPAUSA         = 35
	ChemsDISK          = 36
	ChemsTREE          = 37
	ChemsFILE          = 38
	ChemsLETRA         = 39
	ChemsENTERO        = 40
	ChemsCADENA        = 41
	ChemsIDENTIFICADOR = 42
	ChemsRUTA          = 43
	ChemsIGUAL         = 44
	ChemsWHITESPACE    = 45
	ChemsCOMENTARIOS   = 46
)

// Chems rules.
const (
	ChemsRULE_start            = 0
	ChemsRULE_adminDiscos      = 1
	ChemsRULE_pmkdisk          = 2
	ChemsRULE_parametrosmkdisk = 3
	ChemsRULE_prmdisk          = 4
	ChemsRULE_pfdisk           = 5
	ChemsRULE_parametrosfdisk  = 6
	ChemsRULE_pmount           = 7
	ChemsRULE_parametrosmount  = 8
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) AdminDiscos() IAdminDiscosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdminDiscosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdminDiscosContext)
}

func (s *StartContext) PAUSA() antlr.TerminalNode {
	return s.GetToken(ChemsPAUSA, 0)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ChemsEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Chems) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChemsRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(22)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsMKDISK, ChemsRMDISK, ChemsFDISK, ChemsMOUNT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.AdminDiscos()
		}

	case ChemsPAUSA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(19)
			p.Match(ChemsPAUSA)
		}
		AD.Pausa()

	case ChemsEOF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(21)
			p.Match(ChemsEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAdminDiscosContext is an interface to support dynamic dispatch.
type IAdminDiscosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdminDiscosContext differentiates from other interfaces.
	IsAdminDiscosContext()
}

type AdminDiscosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdminDiscosContext() *AdminDiscosContext {
	var p = new(AdminDiscosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_adminDiscos
	return p
}

func (*AdminDiscosContext) IsAdminDiscosContext() {}

func NewAdminDiscosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdminDiscosContext {
	var p = new(AdminDiscosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_adminDiscos

	return p
}

func (s *AdminDiscosContext) GetParser() antlr.Parser { return s.parser }

func (s *AdminDiscosContext) MKDISK() antlr.TerminalNode {
	return s.GetToken(ChemsMKDISK, 0)
}

func (s *AdminDiscosContext) Pmkdisk() IPmkdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmkdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmkdiskContext)
}

func (s *AdminDiscosContext) RMDISK() antlr.TerminalNode {
	return s.GetToken(ChemsRMDISK, 0)
}

func (s *AdminDiscosContext) Prmdisk() IPrmdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrmdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrmdiskContext)
}

func (s *AdminDiscosContext) FDISK() antlr.TerminalNode {
	return s.GetToken(ChemsFDISK, 0)
}

func (s *AdminDiscosContext) Pfdisk() IPfdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPfdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPfdiskContext)
}

func (s *AdminDiscosContext) MOUNT() antlr.TerminalNode {
	return s.GetToken(ChemsMOUNT, 0)
}

func (s *AdminDiscosContext) Pmount() IPmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmountContext)
}

func (s *AdminDiscosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdminDiscosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdminDiscosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterAdminDiscos(s)
	}
}

func (s *AdminDiscosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitAdminDiscos(s)
	}
}

func (p *Chems) AdminDiscos() (localctx IAdminDiscosContext) {
	localctx = NewAdminDiscosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChemsRULE_adminDiscos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsMKDISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Match(ChemsMKDISK)
		}
		{
			p.SetState(25)
			p.pmkdisk(0)
		}

	case ChemsRMDISK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)
			p.Match(ChemsRMDISK)
		}
		{
			p.SetState(27)
			p.Prmdisk()
		}

	case ChemsFDISK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(28)
			p.Match(ChemsFDISK)
		}
		{
			p.SetState(29)
			p.pfdisk(0)
		}

	case ChemsMOUNT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(30)
			p.Match(ChemsMOUNT)
		}
		{
			p.SetState(31)
			p.pmount(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPmkdiskContext is an interface to support dynamic dispatch.
type IPmkdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPmkdiskContext differentiates from other interfaces.
	IsPmkdiskContext()
}

type PmkdiskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPmkdiskContext() *PmkdiskContext {
	var p = new(PmkdiskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_pmkdisk
	return p
}

func (*PmkdiskContext) IsPmkdiskContext() {}

func NewPmkdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PmkdiskContext {
	var p = new(PmkdiskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_pmkdisk

	return p
}

func (s *PmkdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *PmkdiskContext) Parametrosmkdisk() IParametrosmkdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosmkdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosmkdiskContext)
}

func (s *PmkdiskContext) Pmkdisk() IPmkdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmkdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmkdiskContext)
}

func (s *PmkdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PmkdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PmkdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPmkdisk(s)
	}
}

func (s *PmkdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPmkdisk(s)
	}
}

func (p *Chems) Pmkdisk() (localctx IPmkdiskContext) {
	return p.pmkdisk(0)
}

func (p *Chems) pmkdisk(_p int) (localctx IPmkdiskContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPmkdiskContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPmkdiskContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, ChemsRULE_pmkdisk, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Parametrosmkdisk()
	}
	AD.Mkdisk(ValSize, ValFit, ValUnit, ValPath)
	ValSize = ""

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPmkdiskContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_pmkdisk)
			p.SetState(38)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(39)
				p.Parametrosmkdisk()
			}

		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IParametrosmkdiskContext is an interface to support dynamic dispatch.
type IParametrosmkdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_RUTA returns the _RUTA token.
	Get_RUTA() antlr.Token

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_RUTA sets the _RUTA token.
	Set_RUTA(antlr.Token)

	// IsParametrosmkdiskContext differentiates from other interfaces.
	IsParametrosmkdiskContext()
}

type ParametrosmkdiskContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_ENTERO antlr.Token
	_CADENA antlr.Token
	_RUTA   antlr.Token
}

func NewEmptyParametrosmkdiskContext() *ParametrosmkdiskContext {
	var p = new(ParametrosmkdiskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_parametrosmkdisk
	return p
}

func (*ParametrosmkdiskContext) IsParametrosmkdiskContext() {}

func NewParametrosmkdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosmkdiskContext {
	var p = new(ParametrosmkdiskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_parametrosmkdisk

	return p
}

func (s *ParametrosmkdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosmkdiskContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *ParametrosmkdiskContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ParametrosmkdiskContext) Get_RUTA() antlr.Token { return s._RUTA }

func (s *ParametrosmkdiskContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *ParametrosmkdiskContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ParametrosmkdiskContext) Set_RUTA(v antlr.Token) { s._RUTA = v }

func (s *ParametrosmkdiskContext) SIZE() antlr.TerminalNode {
	return s.GetToken(ChemsSIZE, 0)
}

func (s *ParametrosmkdiskContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *ParametrosmkdiskContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(ChemsENTERO, 0)
}

func (s *ParametrosmkdiskContext) FIT() antlr.TerminalNode {
	return s.GetToken(ChemsFIT, 0)
}

func (s *ParametrosmkdiskContext) BF() antlr.TerminalNode {
	return s.GetToken(ChemsBF, 0)
}

func (s *ParametrosmkdiskContext) FF() antlr.TerminalNode {
	return s.GetToken(ChemsFF, 0)
}

func (s *ParametrosmkdiskContext) WF() antlr.TerminalNode {
	return s.GetToken(ChemsWF, 0)
}

func (s *ParametrosmkdiskContext) UNIT() antlr.TerminalNode {
	return s.GetToken(ChemsUNIT, 0)
}

func (s *ParametrosmkdiskContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *ParametrosmkdiskContext) PATH() antlr.TerminalNode {
	return s.GetToken(ChemsPATH, 0)
}

func (s *ParametrosmkdiskContext) RUTA() antlr.TerminalNode {
	return s.GetToken(ChemsRUTA, 0)
}

func (s *ParametrosmkdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosmkdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosmkdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterParametrosmkdisk(s)
	}
}

func (s *ParametrosmkdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitParametrosmkdisk(s)
	}
}

func (p *Chems) Parametrosmkdisk() (localctx IParametrosmkdiskContext) {
	localctx = NewParametrosmkdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChemsRULE_parametrosmkdisk)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.Match(ChemsSIZE)
		}
		{
			p.SetState(48)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(49)

			var _m = p.Match(ChemsENTERO)

			localctx.(*ParametrosmkdiskContext)._ENTERO = _m
		}
		ValSize = localctx.(*ParametrosmkdiskContext).Get_ENTERO()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Match(ChemsFIT)
		}
		{
			p.SetState(52)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(53)
			p.Match(ChemsBF)
		}
		ValFit = "b"

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Match(ChemsFIT)
		}
		{
			p.SetState(56)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(57)
			p.Match(ChemsFF)
		}
		ValFit = "f"

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(59)
			p.Match(ChemsFIT)
		}
		{
			p.SetState(60)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(61)
			p.Match(ChemsWF)
		}
		ValFit = "w"

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.Match(ChemsUNIT)
		}
		{
			p.SetState(64)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(65)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosmkdiskContext)._CADENA = _m
		}
		ValUnit = localctx.(*ParametrosmkdiskContext).Get_CADENA()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(67)
			p.Match(ChemsPATH)
		}
		{
			p.SetState(68)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(69)

			var _m = p.Match(ChemsRUTA)

			localctx.(*ParametrosmkdiskContext)._RUTA = _m
		}
		ValPath = localctx.(*ParametrosmkdiskContext).Get_RUTA()

	}

	return localctx
}

// IPrmdiskContext is an interface to support dynamic dispatch.
type IPrmdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RUTA returns the _RUTA token.
	Get_RUTA() antlr.Token

	// Set_RUTA sets the _RUTA token.
	Set_RUTA(antlr.Token)

	// IsPrmdiskContext differentiates from other interfaces.
	IsPrmdiskContext()
}

type PrmdiskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_RUTA  antlr.Token
}

func NewEmptyPrmdiskContext() *PrmdiskContext {
	var p = new(PrmdiskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_prmdisk
	return p
}

func (*PrmdiskContext) IsPrmdiskContext() {}

func NewPrmdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrmdiskContext {
	var p = new(PrmdiskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_prmdisk

	return p
}

func (s *PrmdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *PrmdiskContext) Get_RUTA() antlr.Token { return s._RUTA }

func (s *PrmdiskContext) Set_RUTA(v antlr.Token) { s._RUTA = v }

func (s *PrmdiskContext) RMDISK() antlr.TerminalNode {
	return s.GetToken(ChemsRMDISK, 0)
}

func (s *PrmdiskContext) PATH() antlr.TerminalNode {
	return s.GetToken(ChemsPATH, 0)
}

func (s *PrmdiskContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *PrmdiskContext) RUTA() antlr.TerminalNode {
	return s.GetToken(ChemsRUTA, 0)
}

func (s *PrmdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrmdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrmdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPrmdisk(s)
	}
}

func (s *PrmdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPrmdisk(s)
	}
}

func (p *Chems) Prmdisk() (localctx IPrmdiskContext) {
	localctx = NewPrmdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ChemsRULE_prmdisk)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(ChemsRMDISK)
	}
	{
		p.SetState(74)
		p.Match(ChemsPATH)
	}
	{
		p.SetState(75)
		p.Match(ChemsIGUAL)
	}
	{
		p.SetState(76)

		var _m = p.Match(ChemsRUTA)

		localctx.(*PrmdiskContext)._RUTA = _m
	}
	ValPath = localctx.(*PrmdiskContext).Get_RUTA()

	return localctx
}

// IPfdiskContext is an interface to support dynamic dispatch.
type IPfdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPfdiskContext differentiates from other interfaces.
	IsPfdiskContext()
}

type PfdiskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPfdiskContext() *PfdiskContext {
	var p = new(PfdiskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_pfdisk
	return p
}

func (*PfdiskContext) IsPfdiskContext() {}

func NewPfdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PfdiskContext {
	var p = new(PfdiskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_pfdisk

	return p
}

func (s *PfdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *PfdiskContext) Parametrosfdisk() IParametrosfdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosfdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosfdiskContext)
}

func (s *PfdiskContext) Pfdisk() IPfdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPfdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPfdiskContext)
}

func (s *PfdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PfdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PfdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPfdisk(s)
	}
}

func (s *PfdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPfdisk(s)
	}
}

func (p *Chems) Pfdisk() (localctx IPfdiskContext) {
	return p.pfdisk(0)
}

func (p *Chems) pfdisk(_p int) (localctx IPfdiskContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPfdiskContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPfdiskContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ChemsRULE_pfdisk, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Parametrosfdisk()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPfdiskContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_pfdisk)
			p.SetState(83)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(84)
				p.Parametrosfdisk()
			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IParametrosfdiskContext is an interface to support dynamic dispatch.
type IParametrosfdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// IsParametrosfdiskContext differentiates from other interfaces.
	IsParametrosfdiskContext()
}

type ParametrosfdiskContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_CADENA antlr.Token
}

func NewEmptyParametrosfdiskContext() *ParametrosfdiskContext {
	var p = new(ParametrosfdiskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_parametrosfdisk
	return p
}

func (*ParametrosfdiskContext) IsParametrosfdiskContext() {}

func NewParametrosfdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosfdiskContext {
	var p = new(ParametrosfdiskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_parametrosfdisk

	return p
}

func (s *ParametrosfdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosfdiskContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ParametrosfdiskContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ParametrosfdiskContext) Parametrosmkdisk() IParametrosmkdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosmkdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosmkdiskContext)
}

func (s *ParametrosfdiskContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ChemsTYPE, 0)
}

func (s *ParametrosfdiskContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *ParametrosfdiskContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *ParametrosfdiskContext) NAME() antlr.TerminalNode {
	return s.GetToken(ChemsNAME, 0)
}

func (s *ParametrosfdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosfdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosfdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterParametrosfdisk(s)
	}
}

func (s *ParametrosfdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitParametrosfdisk(s)
	}
}

func (p *Chems) Parametrosfdisk() (localctx IParametrosfdiskContext) {
	localctx = NewParametrosfdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ChemsRULE_parametrosfdisk)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsSIZE, ChemsFIT, ChemsUNIT, ChemsPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Parametrosmkdisk()
		}

	case ChemsTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(ChemsTYPE)
		}
		{
			p.SetState(94)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(95)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosfdiskContext)._CADENA = _m
		}
		ValType = localctx.(*ParametrosfdiskContext).Get_CADENA()

	case ChemsNAME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(97)
			p.Match(ChemsNAME)
		}
		{
			p.SetState(98)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(99)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosfdiskContext)._CADENA = _m
		}
		ValName = localctx.(*ParametrosfdiskContext).Get_CADENA()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPmountContext is an interface to support dynamic dispatch.
type IPmountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPmountContext differentiates from other interfaces.
	IsPmountContext()
}

type PmountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPmountContext() *PmountContext {
	var p = new(PmountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_pmount
	return p
}

func (*PmountContext) IsPmountContext() {}

func NewPmountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PmountContext {
	var p = new(PmountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_pmount

	return p
}

func (s *PmountContext) GetParser() antlr.Parser { return s.parser }

func (s *PmountContext) Parametrosmount() IParametrosmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosmountContext)
}

func (s *PmountContext) Pmount() IPmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmountContext)
}

func (s *PmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PmountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PmountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPmount(s)
	}
}

func (s *PmountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPmount(s)
	}
}

func (p *Chems) Pmount() (localctx IPmountContext) {
	return p.pmount(0)
}

func (p *Chems) pmount(_p int) (localctx IPmountContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPmountContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPmountContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, ChemsRULE_pmount, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Parametrosmount()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPmountContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_pmount)
			p.SetState(107)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(108)
				p.Parametrosmount()
			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IParametrosmountContext is an interface to support dynamic dispatch.
type IParametrosmountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RUTA returns the _RUTA token.
	Get_RUTA() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Set_RUTA sets the _RUTA token.
	Set_RUTA(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// IsParametrosmountContext differentiates from other interfaces.
	IsParametrosmountContext()
}

type ParametrosmountContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_RUTA   antlr.Token
	_CADENA antlr.Token
}

func NewEmptyParametrosmountContext() *ParametrosmountContext {
	var p = new(ParametrosmountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_parametrosmount
	return p
}

func (*ParametrosmountContext) IsParametrosmountContext() {}

func NewParametrosmountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosmountContext {
	var p = new(ParametrosmountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_parametrosmount

	return p
}

func (s *ParametrosmountContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosmountContext) Get_RUTA() antlr.Token { return s._RUTA }

func (s *ParametrosmountContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ParametrosmountContext) Set_RUTA(v antlr.Token) { s._RUTA = v }

func (s *ParametrosmountContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ParametrosmountContext) PATH() antlr.TerminalNode {
	return s.GetToken(ChemsPATH, 0)
}

func (s *ParametrosmountContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *ParametrosmountContext) RUTA() antlr.TerminalNode {
	return s.GetToken(ChemsRUTA, 0)
}

func (s *ParametrosmountContext) NAME() antlr.TerminalNode {
	return s.GetToken(ChemsNAME, 0)
}

func (s *ParametrosmountContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *ParametrosmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosmountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosmountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterParametrosmount(s)
	}
}

func (s *ParametrosmountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitParametrosmount(s)
	}
}

func (p *Chems) Parametrosmount() (localctx IParametrosmountContext) {
	localctx = NewParametrosmountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ChemsRULE_parametrosmount)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(ChemsPATH)
		}
		{
			p.SetState(117)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(118)

			var _m = p.Match(ChemsRUTA)

			localctx.(*ParametrosmountContext)._RUTA = _m
		}
		ValPath = localctx.(*ParametrosmountContext).Get_RUTA()

	case ChemsNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Match(ChemsNAME)
		}
		{
			p.SetState(121)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(122)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosmountContext)._CADENA = _m
		}
		ValName = localctx.(*ParametrosmountContext).Get_CADENA()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *PmkdiskContext = nil
		if localctx != nil {
			t = localctx.(*PmkdiskContext)
		}
		return p.Pmkdisk_Sempred(t, predIndex)

	case 5:
		var t *PfdiskContext = nil
		if localctx != nil {
			t = localctx.(*PfdiskContext)
		}
		return p.Pfdisk_Sempred(t, predIndex)

	case 7:
		var t *PmountContext = nil
		if localctx != nil {
			t = localctx.(*PmountContext)
		}
		return p.Pmount_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Chems) Pmkdisk_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Pfdisk_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Pmount_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
