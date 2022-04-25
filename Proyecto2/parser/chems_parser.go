// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Proyecto2/nodos"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 58, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 2, 5, 2, 17, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 24, 10, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 34, 10, 4, 12, 4,
	14, 4, 37, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	5, 5, 48, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 56, 10, 6, 3,
	6, 2, 3, 6, 7, 2, 4, 6, 8, 10, 2, 2, 2, 58, 2, 16, 3, 2, 2, 2, 4, 23, 3,
	2, 2, 2, 6, 25, 3, 2, 2, 2, 8, 47, 3, 2, 2, 2, 10, 55, 3, 2, 2, 2, 12,
	13, 5, 4, 3, 2, 13, 14, 8, 2, 1, 2, 14, 17, 3, 2, 2, 2, 15, 17, 7, 34,
	2, 2, 16, 12, 3, 2, 2, 2, 16, 15, 3, 2, 2, 2, 17, 3, 3, 2, 2, 2, 18, 19,
	7, 35, 2, 2, 19, 20, 5, 6, 4, 2, 20, 21, 8, 3, 1, 2, 21, 24, 3, 2, 2, 2,
	22, 24, 7, 36, 2, 2, 23, 18, 3, 2, 2, 2, 23, 22, 3, 2, 2, 2, 24, 5, 3,
	2, 2, 2, 25, 26, 8, 4, 1, 2, 26, 27, 5, 8, 5, 2, 27, 28, 8, 4, 1, 2, 28,
	35, 3, 2, 2, 2, 29, 30, 12, 4, 2, 2, 30, 31, 5, 8, 5, 2, 31, 32, 8, 4,
	1, 2, 32, 34, 3, 2, 2, 2, 33, 29, 3, 2, 2, 2, 34, 37, 3, 2, 2, 2, 35, 33,
	3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 7, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2,
	38, 39, 7, 37, 2, 2, 39, 40, 7, 18, 2, 2, 40, 41, 7, 38, 2, 2, 41, 48,
	8, 5, 1, 2, 42, 43, 7, 39, 2, 2, 43, 44, 7, 18, 2, 2, 44, 45, 5, 10, 6,
	2, 45, 46, 8, 5, 1, 2, 46, 48, 3, 2, 2, 2, 47, 38, 3, 2, 2, 2, 47, 42,
	3, 2, 2, 2, 48, 9, 3, 2, 2, 2, 49, 50, 7, 40, 2, 2, 50, 56, 8, 6, 1, 2,
	51, 52, 7, 41, 2, 2, 52, 56, 8, 6, 1, 2, 53, 54, 7, 42, 2, 2, 54, 56, 8,
	6, 1, 2, 55, 49, 3, 2, 2, 2, 55, 51, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 56,
	11, 3, 2, 2, 2, 7, 16, 23, 35, 47, 55,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'console'", "'log'", "'number'", "'string'", "'if'", "'while'", "'struct'",
	"", "", "", "'.'", "';'", "','", "':'", "'!'", "'='", "'>='", "'<='", "'>'",
	"'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "CONSOLE", "LOG", "P_NUMBER", "P_STRING", "P_IF", "P_WHILE", "P_STRUCT",
	"NUMBER", "STRING", "ID", "PUNTO", "PTCOMA", "COMA", "DOSPT", "DIFERENTE",
	"IGUAL", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", "DIV", "ADD",
	"SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORIZQ", "CORDER",
	"WHITESPACE", "PAUSA", "MKDISK", "RMDISK", "SIZE", "ENTERO", "FIT", "UNIT",
	"CARACTER", "PATH", "RUTA", "BF", "FF", "WF",
}

var ruleNames = []string{
	"start", "adminDiscos", "mkdisk", "parametrosmkdisk", "ajuste",
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
	ChemsEOF        = antlr.TokenEOF
	ChemsCONSOLE    = 1
	ChemsLOG        = 2
	ChemsP_NUMBER   = 3
	ChemsP_STRING   = 4
	ChemsP_IF       = 5
	ChemsP_WHILE    = 6
	ChemsP_STRUCT   = 7
	ChemsNUMBER     = 8
	ChemsSTRING     = 9
	ChemsID         = 10
	ChemsPUNTO      = 11
	ChemsPTCOMA     = 12
	ChemsCOMA       = 13
	ChemsDOSPT      = 14
	ChemsDIFERENTE  = 15
	ChemsIGUAL      = 16
	ChemsMAYORIGUAL = 17
	ChemsMENORIGUAL = 18
	ChemsMAYOR      = 19
	ChemsMENOR      = 20
	ChemsMUL        = 21
	ChemsDIV        = 22
	ChemsADD        = 23
	ChemsSUB        = 24
	ChemsPARIZQ     = 25
	ChemsPARDER     = 26
	ChemsLLAVEIZQ   = 27
	ChemsLLAVEDER   = 28
	ChemsCORIZQ     = 29
	ChemsCORDER     = 30
	ChemsWHITESPACE = 31
	ChemsPAUSA      = 32
	ChemsMKDISK     = 33
	ChemsRMDISK     = 34
	ChemsSIZE       = 35
	ChemsENTERO     = 36
	ChemsFIT        = 37
	ChemsUNIT       = 38
	ChemsCARACTER   = 39
	ChemsPATH       = 40
	ChemsRUTA       = 41
	ChemsBF         = 42
	ChemsFF         = 43
	ChemsWF         = 44
)

// Chems rules.
const (
	ChemsRULE_start            = 0
	ChemsRULE_adminDiscos      = 1
	ChemsRULE_mkdisk           = 2
	ChemsRULE_parametrosmkdisk = 3
	ChemsRULE_ajuste           = 4
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_adminDiscos returns the _adminDiscos rule contexts.
	Get_adminDiscos() IAdminDiscosContext

	// Set_adminDiscos sets the _adminDiscos rule contexts.
	Set_adminDiscos(IAdminDiscosContext)

	// GetEnvio returns the Envio attribute.
	GetEnvio() *nodos.Nodo

	// SetEnvio sets the Envio attribute.
	SetEnvio(*nodos.Nodo)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	Envio        *nodos.Nodo
	_adminDiscos IAdminDiscosContext
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

func (s *StartContext) Get_adminDiscos() IAdminDiscosContext { return s._adminDiscos }

func (s *StartContext) Set_adminDiscos(v IAdminDiscosContext) { s._adminDiscos = v }

func (s *StartContext) GetEnvio() *nodos.Nodo { return s.Envio }

func (s *StartContext) SetEnvio(v *nodos.Nodo) { s.Envio = v }

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

	p.SetState(14)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsMKDISK, ChemsRMDISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)

			var _x = p.AdminDiscos()

			localctx.(*StartContext)._adminDiscos = _x
		}
		localctx.(*StartContext).Envio = localctx.(*StartContext).Get_adminDiscos().GetRaiz()

	case ChemsPAUSA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)
			p.Match(ChemsPAUSA)
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

	// GetRaiz returns the Raiz attribute.
	GetRaiz() *nodos.Nodo

	// SetRaiz sets the Raiz attribute.
	SetRaiz(*nodos.Nodo)

	// IsAdminDiscosContext differentiates from other interfaces.
	IsAdminDiscosContext()
}

type AdminDiscosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Raiz   *nodos.Nodo
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

func (s *AdminDiscosContext) GetRaiz() *nodos.Nodo { return s.Raiz }

func (s *AdminDiscosContext) SetRaiz(v *nodos.Nodo) { s.Raiz = v }

func (s *AdminDiscosContext) MKDISK() antlr.TerminalNode {
	return s.GetToken(ChemsMKDISK, 0)
}

func (s *AdminDiscosContext) Mkdisk() IMkdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMkdiskContext)
}

func (s *AdminDiscosContext) RMDISK() antlr.TerminalNode {
	return s.GetToken(ChemsRMDISK, 0)
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

	p.SetState(21)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsMKDISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(16)
			p.Match(ChemsMKDISK)
		}
		{
			p.SetState(17)
			p.mkdisk(0)
		}
		fmt.Println("hOLA")
		nodos.NewNodo("MKDISK", "")
		localctx.(*AdminDiscosContext).Raiz = nodos.Raiz

	case ChemsRMDISK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(20)
			p.Match(ChemsRMDISK)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMkdiskContext is an interface to support dynamic dispatch.
type IMkdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_parametrosmkdisk returns the _parametrosmkdisk rule contexts.
	Get_parametrosmkdisk() IParametrosmkdiskContext

	// Set_parametrosmkdisk sets the _parametrosmkdisk rule contexts.
	Set_parametrosmkdisk(IParametrosmkdiskContext)

	// IsMkdiskContext differentiates from other interfaces.
	IsMkdiskContext()
}

type MkdiskContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	_parametrosmkdisk IParametrosmkdiskContext
}

func NewEmptyMkdiskContext() *MkdiskContext {
	var p = new(MkdiskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_mkdisk
	return p
}

func (*MkdiskContext) IsMkdiskContext() {}

func NewMkdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdiskContext {
	var p = new(MkdiskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_mkdisk

	return p
}

func (s *MkdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdiskContext) Get_parametrosmkdisk() IParametrosmkdiskContext { return s._parametrosmkdisk }

func (s *MkdiskContext) Set_parametrosmkdisk(v IParametrosmkdiskContext) { s._parametrosmkdisk = v }

func (s *MkdiskContext) Parametrosmkdisk() IParametrosmkdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosmkdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosmkdiskContext)
}

func (s *MkdiskContext) Mkdisk() IMkdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMkdiskContext)
}

func (s *MkdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterMkdisk(s)
	}
}

func (s *MkdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitMkdisk(s)
	}
}

func (p *Chems) Mkdisk() (localctx IMkdiskContext) {
	return p.mkdisk(0)
}

func (p *Chems) mkdisk(_p int) (localctx IMkdiskContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMkdiskContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMkdiskContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, ChemsRULE_mkdisk, _p)

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
		p.SetState(24)

		var _x = p.Parametrosmkdisk()

		localctx.(*MkdiskContext)._parametrosmkdisk = _x
	}
	fmt.Println("Parametros")
	nodos.Add(nodos.Nodoo("PARAMETRO", ""))
	nodos.AddParameters(localctx.(*MkdiskContext).Get_parametrosmkdisk().GetNODS())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMkdiskContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_mkdisk)
			p.SetState(27)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(28)

				var _x = p.Parametrosmkdisk()

				localctx.(*MkdiskContext)._parametrosmkdisk = _x
			}
			fmt.Println("hOLA 2")
			nodos.AddParameters(localctx.(*MkdiskContext).Get_parametrosmkdisk().GetNODS())

		}
		p.SetState(35)
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

	// Get_CARACTER returns the _CARACTER token.
	Get_CARACTER() antlr.Token

	// Get_RUTA returns the _RUTA token.
	Get_RUTA() antlr.Token

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_CARACTER sets the _CARACTER token.
	Set_CARACTER(antlr.Token)

	// Set_RUTA sets the _RUTA token.
	Set_RUTA(antlr.Token)

	// Get_ajuste returns the _ajuste rule contexts.
	Get_ajuste() IAjusteContext

	// Set_ajuste sets the _ajuste rule contexts.
	Set_ajuste(IAjusteContext)

	// GetNODS returns the NODS attribute.
	GetNODS() nodos.Nodo

	// SetNODS sets the NODS attribute.
	SetNODS(nodos.Nodo)

	// IsParametrosmkdiskContext differentiates from other interfaces.
	IsParametrosmkdiskContext()
}

type ParametrosmkdiskContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	NODS      nodos.Nodo
	_ENTERO   antlr.Token
	_ajuste   IAjusteContext
	_CARACTER antlr.Token
	_RUTA     antlr.Token
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

func (s *ParametrosmkdiskContext) Get_CARACTER() antlr.Token { return s._CARACTER }

func (s *ParametrosmkdiskContext) Get_RUTA() antlr.Token { return s._RUTA }

func (s *ParametrosmkdiskContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *ParametrosmkdiskContext) Set_CARACTER(v antlr.Token) { s._CARACTER = v }

func (s *ParametrosmkdiskContext) Set_RUTA(v antlr.Token) { s._RUTA = v }

func (s *ParametrosmkdiskContext) Get_ajuste() IAjusteContext { return s._ajuste }

func (s *ParametrosmkdiskContext) Set_ajuste(v IAjusteContext) { s._ajuste = v }

func (s *ParametrosmkdiskContext) GetNODS() nodos.Nodo { return s.NODS }

func (s *ParametrosmkdiskContext) SetNODS(v nodos.Nodo) { s.NODS = v }

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

func (s *ParametrosmkdiskContext) Ajuste() IAjusteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAjusteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAjusteContext)
}

func (s *ParametrosmkdiskContext) UNIT() antlr.TerminalNode {
	return s.GetToken(ChemsUNIT, 0)
}

func (s *ParametrosmkdiskContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(ChemsCARACTER, 0)
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

	p.SetState(53)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsSIZE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Match(ChemsSIZE)
		}
		{
			p.SetState(37)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(38)

			var _m = p.Match(ChemsENTERO)

			localctx.(*ParametrosmkdiskContext)._ENTERO = _m
		}
		localctx.(*ParametrosmkdiskContext).NODS = nodos.Nodoo("size", (func() string {
			if localctx.(*ParametrosmkdiskContext).Get_ENTERO() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkdiskContext).Get_ENTERO().GetText()
			}
		}()))

	case ChemsFIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Match(ChemsFIT)
		}
		{
			p.SetState(41)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(42)

			var _x = p.Ajuste()

			localctx.(*ParametrosmkdiskContext)._ajuste = _x
		}
		localctx.(*ParametrosmkdiskContext).NODS = nodos.Nodoo("fit", "")
		nodos.AddParameters(localctx.(*ParametrosmkdiskContext).Get_ajuste().GetAJ())

	case ChemsUNIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(45)
			p.Match(ChemsUNIT)
		}
		{
			p.SetState(46)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(47)

			var _m = p.Match(ChemsCARACTER)

			localctx.(*ParametrosmkdiskContext)._CARACTER = _m
		}
		localctx.(*ParametrosmkdiskContext).NODS = nodos.Nodoo("unit", (func() string {
			if localctx.(*ParametrosmkdiskContext).Get_CARACTER() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkdiskContext).Get_CARACTER().GetText()
			}
		}()))

	case ChemsPATH:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.Match(ChemsPATH)
		}
		{
			p.SetState(50)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(51)

			var _m = p.Match(ChemsRUTA)

			localctx.(*ParametrosmkdiskContext)._RUTA = _m
		}
		localctx.(*ParametrosmkdiskContext).NODS = nodos.Nodoo("path", (func() string {
			if localctx.(*ParametrosmkdiskContext).Get_RUTA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkdiskContext).Get_RUTA().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAjusteContext is an interface to support dynamic dispatch.
type IAjusteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAJ returns the AJ attribute.
	GetAJ() nodos.Nodo

	// SetAJ sets the AJ attribute.
	SetAJ(nodos.Nodo)

	// IsAjusteContext differentiates from other interfaces.
	IsAjusteContext()
}

type AjusteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	AJ     nodos.Nodo
}

func NewEmptyAjusteContext() *AjusteContext {
	var p = new(AjusteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_ajuste
	return p
}

func (*AjusteContext) IsAjusteContext() {}

func NewAjusteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AjusteContext {
	var p = new(AjusteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_ajuste

	return p
}

func (s *AjusteContext) GetParser() antlr.Parser { return s.parser }

func (s *AjusteContext) GetAJ() nodos.Nodo { return s.AJ }

func (s *AjusteContext) SetAJ(v nodos.Nodo) { s.AJ = v }

func (s *AjusteContext) BF() antlr.TerminalNode {
	return s.GetToken(ChemsBF, 0)
}

func (s *AjusteContext) FF() antlr.TerminalNode {
	return s.GetToken(ChemsFF, 0)
}

func (s *AjusteContext) WF() antlr.TerminalNode {
	return s.GetToken(ChemsWF, 0)
}

func (s *AjusteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AjusteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AjusteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterAjuste(s)
	}
}

func (s *AjusteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitAjuste(s)
	}
}

func (p *Chems) Ajuste() (localctx IAjusteContext) {
	localctx = NewAjusteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ChemsRULE_ajuste)

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

	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsBF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Match(ChemsBF)
		}
		localctx.(*AjusteContext).AJ = nodos.Nodoo("ajuste", "bf")

	case ChemsFF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(ChemsFF)
		}
		localctx.(*AjusteContext).AJ = nodos.Nodoo("ajuste", "ff")

	case ChemsWF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(59)
			p.Match(ChemsWF)
		}
		localctx.(*AjusteContext).AJ = nodos.Nodoo("ajuste", "wf")

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *MkdiskContext = nil
		if localctx != nil {
			t = localctx.(*MkdiskContext)
		}
		return p.Mkdisk_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Chems) Mkdisk_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
