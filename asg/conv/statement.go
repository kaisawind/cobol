package conv

import (
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

func Statement(in cobol85.IStatementContext) (out *pb.Statement) {
	ctx := in.(*cobol85.StatementContext)
	out = &pb.Statement{}
	if t := ctx.MoveStatement(); t != nil {
		out.OneOf = &pb.Statement_MoveStatement{
			MoveStatement: MoveToStatement(t),
		}
	} else if t := ctx.AcceptStatement(); t != nil {
		out.OneOf = &pb.Statement_AcceptStatement{
			AcceptStatement: AcceptStatement(t),
		}
	} else if t := ctx.AddStatement(); t != nil {
		out.OneOf = &pb.Statement_AddStatement{
			AddStatement: AddStatement(t),
		}
	} else if t := ctx.DisplayStatement(); t != nil {
		out.OneOf = &pb.Statement_DisplayStatement{
			DisplayStatement: DisplayStatement(t),
		}
	}
	return
}

func DisplayStatement(in cobol85.IDisplayStatementContext) (out *pb.DisplayStatement) {
	ctx := in.(*cobol85.DisplayStatementContext)
	out = &pb.DisplayStatement{}
	if ctx.OnExceptionClause() != nil {
		out.OnExceptionClause = OnExceptionClause(ctx.OnExceptionClause())
	}
	if ctx.NotOnExceptionClause() != nil {
		out.NotOnExceptionClause = NotOnExceptionClause(ctx.NotOnExceptionClause())
	}
	for _, v := range ctx.AllDisplayOperand() {
		iv := v.(*cobol85.DisplayOperandContext)
		operand := &pb.DisplayOperand{}
		if iv.Identifier() != nil {
			operand.OneOf = &pb.DisplayOperand_Identifier{
				Identifier: Identifier(iv.Identifier()),
			}
		} else if iv.Literal() != nil {
			operand.OneOf = &pb.DisplayOperand_Literal{
				Literal: Literal(iv.Literal()),
			}
		}
		out.Operands = append(out.Operands, operand)
	}

	if ictx := ctx.DisplayAt(); ictx != nil {
		cctx := ictx.(*cobol85.DisplayAtContext)
		at := &pb.DisplayAt{}
		if cctx.Identifier() != nil {
			at.OneOf = &pb.DisplayAt_Identifier{
				Identifier: Identifier(cctx.Identifier()),
			}
		} else if cctx.Literal() != nil {
			at.OneOf = &pb.DisplayAt_Literal{
				Literal: Literal(cctx.Literal()),
			}
		}
		out.At = at
	}

	if ictx := ctx.DisplayUpon(); ictx != nil {
		cctx := ictx.(*cobol85.DisplayUponContext)
		upon := &pb.DisplayUpon{}
		if cctx.MnemonicName() != nil {
			upon.OneOf = &pb.DisplayUpon_MnemonicName{
				MnemonicName: MnemonicName(cctx.MnemonicName()),
			}
		} else if cctx.EnvironmentName() != nil {
			upon.OneOf = &pb.DisplayUpon_EnvironmentName{
				EnvironmentName: EnvironmentName(cctx.EnvironmentName()),
			}
		}
		out.Upon = upon
	}

	if ictx := ctx.DisplayWith(); ictx != nil {
		out.With = &pb.DisplayWith{}
	}
	return
}

func AddStatement(in cobol85.IAddStatementContext) (out *pb.AddStatement) {
	ctx := in.(*cobol85.AddStatementContext)
	out = &pb.AddStatement{}
	if ctx.OnSizeErrorPhrase() != nil {
		out.OnSizeErrorPhrase = OnSizeErrorPhrase(ctx.OnSizeErrorPhrase())
	}
	if ctx.NotOnSizeErrorPhrase() != nil {
		out.NotOnSizeErrorPhrase = NotOnSizeErrorPhrase(ctx.NotOnSizeErrorPhrase())
	}

	if ictx := ctx.AddToStatement(); ictx != nil {
		cctx := ictx.(*cobol85.AddToStatementContext)
		addTo := &pb.AddStatement_To{}
		for _, v := range cctx.AllAddFrom() {
			from := &pb.AddStatement_AddFrom{}
			iv := v.(*cobol85.AddFromContext)
			if iv.Identifier() != nil {
				from.OneOf = &pb.AddStatement_AddFrom_Identifier{
					Identifier: Identifier(iv.Identifier()),
				}
			} else if iv.Literal() != nil {
				from.OneOf = &pb.AddStatement_AddFrom_Literal{
					Literal: Literal(iv.Literal()),
				}
			}
			addTo.Froms = append(addTo.Froms, from)
		}
		for _, v := range cctx.AllAddTo() {
			iv := v.(*cobol85.AddToContext)
			if iv.Identifier() != nil {
				addTo.Tos = append(addTo.Tos, Identifier(iv.Identifier()))
			}
		}
		out.OneOf = &pb.AddStatement_To_{
			To: addTo,
		}
	} else if ictx := ctx.AddToGivingStatement(); ictx != nil {
		cctx := ictx.(*cobol85.AddToGivingStatementContext)
		addTo := &pb.AddStatement_ToGiving{}
		for _, v := range cctx.AllAddFrom() {
			iv := v.(*cobol85.AddFromContext)
			from := &pb.AddStatement_AddFrom{}
			if iv.Identifier() != nil {
				from.OneOf = &pb.AddStatement_AddFrom_Identifier{
					Identifier: Identifier(iv.Identifier()),
				}
			} else if iv.Literal() != nil {
				from.OneOf = &pb.AddStatement_AddFrom_Literal{
					Literal: Literal(iv.Literal()),
				}
			}
			addTo.Froms = append(addTo.Froms, from)
		}
		for _, v := range cctx.AllAddToGiving() {
			iv := v.(*cobol85.AddToGivingContext)
			to := &pb.AddStatement_AddToGiving{}
			if iv.Identifier() != nil {
				to.OneOf = &pb.AddStatement_AddToGiving_Identifier{
					Identifier: Identifier(iv.Identifier()),
				}
			} else if iv.Literal() != nil {
				to.OneOf = &pb.AddStatement_AddToGiving_Literal{
					Literal: Literal(iv.Literal()),
				}
			}
			addTo.Tos = append(addTo.Tos, to)
		}

		for _, v := range cctx.AllAddGiving() {
			iv := v.(*cobol85.AddGivingContext)
			if iv.Identifier() != nil {
				addTo.Givings = append(addTo.Givings, Identifier(iv.Identifier()))
			}
		}
		out.OneOf = &pb.AddStatement_ToGiving_{
			ToGiving: addTo,
		}
	} else if ictx := ctx.AddCorrespondingStatement(); ictx != nil {
		cctx := ictx.(*cobol85.AddCorrespondingStatementContext)
		addTo := &pb.AddStatement_Corresponding{}
		if cctx.Identifier() != nil {
			addTo.Corresponding = Identifier(cctx.Identifier())
		}
		if v := cctx.AddTo(); v != nil {
			iv := v.(*cobol85.AddToContext)
			addTo.To = Identifier(iv.Identifier())
		}
		out.OneOf = &pb.AddStatement_Corresponding_{
			Corresponding: addTo,
		}
	}
	return
}

func AcceptStatement(in cobol85.IAcceptStatementContext) (out *pb.AcceptStatement) {
	ctx := in.(*cobol85.AcceptStatementContext)
	out = &pb.AcceptStatement{
		Identifier: Identifier(ctx.Identifier()),
	}
	if ctx.OnExceptionClause() != nil {
		out.OnExceptionClause = OnExceptionClause(ctx.OnExceptionClause())
	}
	if ctx.NotOnExceptionClause() != nil {
		out.NotOnExceptionClause = NotOnExceptionClause(ctx.NotOnExceptionClause())
	}

	if ictx := ctx.AcceptFromDateStatement(); ictx != nil {
		cctx := ictx.(*cobol85.AcceptFromDateStatementContext)

		typ := pb.AcceptStatement_FromDate_DATE
		switch {
		case cctx.DATE() != nil:
			typ = pb.AcceptStatement_FromDate_DATE
		case cctx.DAY() != nil:
			typ = pb.AcceptStatement_FromDate_DAY
		case cctx.DAY_OF_WEEK() != nil:
			typ = pb.AcceptStatement_FromDate_DAY_OF_WEEK
		case cctx.TIME() != nil:
			typ = pb.AcceptStatement_FromDate_TIME
		case cctx.TIMER() != nil:
			typ = pb.AcceptStatement_FromDate_TIMER
		case cctx.TODAYS_DATE() != nil:
			typ = pb.AcceptStatement_FromDate_TODAYS_DATE
		case cctx.TODAYS_NAME() != nil:
			typ = pb.AcceptStatement_FromDate_TODAYS_NAME
		case cctx.YEAR() != nil:
			typ = pb.AcceptStatement_FromDate_YEAR
		case cctx.YYYYMMDD() != nil:
			typ = pb.AcceptStatement_FromDate_YYYYMMDD
		case cctx.YYYYDDD() != nil:
			typ = pb.AcceptStatement_FromDate_YYYYDDD
		}

		out.OneOf = &pb.AcceptStatement_FromDate_{
			FromDate: &pb.AcceptStatement_FromDate{
				Type: typ,
			},
		}
	} else if ictx := ctx.AcceptFromEscapeKeyStatement(); ictx != nil {
		out.OneOf = &pb.AcceptStatement_FromEscapeKey_{
			FromEscapeKey: &pb.AcceptStatement_FromEscapeKey{},
		}
	} else if ictx := ctx.AcceptFromMnemonicStatement(); ictx != nil {
		cctx := ictx.(*cobol85.AcceptFromMnemonicStatementContext)
		out.OneOf = &pb.AcceptStatement_FromMnemonic_{
			FromMnemonic: &pb.AcceptStatement_FromMnemonic{
				MnemonicName: MnemonicName(cctx.MnemonicName()),
			},
		}
	} else if ictx := ctx.AcceptMessageCountStatement(); ictx != nil {
		out.OneOf = &pb.AcceptStatement_MessageCount_{
			MessageCount: &pb.AcceptStatement_MessageCount{},
		}
	}
	return
}

func MoveToStatement(in cobol85.IMoveStatementContext) (out *pb.MoveStatement) {
	ctx := in.(*cobol85.MoveStatementContext)
	out = &pb.MoveStatement{}
	if ictx := ctx.MoveToStatement(); ictx != nil {
		cctx := ictx.(*cobol85.MoveToStatementContext)
		mts := &pb.MoveToStatement{}
		if imts := cctx.MoveToSendingArea(); imts != nil {
			cmts := imts.(*cobol85.MoveToSendingAreaContext)
			if cmts.Identifier() != nil {
				mts.SendingArea = &pb.MoveToStatement_Identifier{
					Identifier: Identifier(cmts.Identifier()),
				}
			} else if cmts.Literal() != nil {
				mts.SendingArea = &pb.MoveToStatement_Literal{
					Literal: Literal(cmts.Literal()),
				}
			}
		}
		for _, i := range cctx.AllIdentifier() {
			mts.To = append(mts.To, Identifier(i))
		}
		out.OneOf = &pb.MoveStatement_MoveTo{
			MoveTo: mts,
		}
	} else if ictx := ctx.MoveCorrespondingToStatement(); ictx != nil {
		cctx := ictx.(*cobol85.MoveCorrespondingToStatementContext)
		mts := &pb.MoveCorrespondingToStatement{}
		if imts := cctx.MoveCorrespondingToSendingArea(); imts != nil {
			cmts := imts.(*cobol85.MoveCorrespondingToSendingAreaContext)
			if cmts.Identifier() != nil {
				mts.SendingArea = Identifier(cmts.Identifier())
			}
		}
		for _, i := range cctx.AllIdentifier() {
			mts.To = append(mts.To, Identifier(i))
		}
		out.OneOf = &pb.MoveStatement_MoveCorrespondingTo{
			MoveCorrespondingTo: mts,
		}
	}
	return
}

func OnExceptionClause(in cobol85.IOnExceptionClauseContext) (out *pb.OnExceptionClause) {
	ctx := in.(*cobol85.OnExceptionClauseContext)
	out = &pb.OnExceptionClause{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func NotOnExceptionClause(in cobol85.INotOnExceptionClauseContext) (out *pb.NotOnExceptionClause) {
	ctx := in.(*cobol85.NotOnExceptionClauseContext)
	out = &pb.NotOnExceptionClause{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func OnSizeErrorPhrase(in cobol85.IOnSizeErrorPhraseContext) (out *pb.OnSizeErrorPhrase) {
	ctx := in.(*cobol85.OnSizeErrorPhraseContext)
	out = &pb.OnSizeErrorPhrase{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func NotOnSizeErrorPhrase(in cobol85.INotOnSizeErrorPhraseContext) (out *pb.NotOnSizeErrorPhrase) {
	ctx := in.(*cobol85.NotOnSizeErrorPhraseContext)
	out = &pb.NotOnSizeErrorPhrase{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}
