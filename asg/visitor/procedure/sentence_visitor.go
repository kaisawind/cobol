package procedure

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/conv"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type SentenceVisitor struct {
	cobol85.BaseCobol85Visitor
	sentence *pb.Sentence
}

func NewSentenceVisitor(sentence *pb.Sentence) *SentenceVisitor {
	return &SentenceVisitor{
		sentence: sentence,
	}
}

func (v *SentenceVisitor) VisitMoveStatement(ctx *cobol85.MoveStatementContext) any {
	statement := &pb.MoveStatement{}
	if ictx := ctx.MoveToStatement(); ictx != nil {
		cctx := ictx.(*cobol85.MoveToStatementContext)
		mts := &pb.MoveToStatement{}
		if imts := cctx.MoveToSendingArea(); imts != nil {
			cmts := imts.(*cobol85.MoveToSendingAreaContext)
			if cmts.Identifier() != nil {
				mts.SendingArea = &pb.MoveToStatement_Identifier{
					Identifier: conv.Identifier(cmts.Identifier()),
				}
			} else if cmts.Literal() != nil {
				mts.SendingArea = &pb.MoveToStatement_Literal{
					Literal: conv.Literal(cmts.Literal()),
				}
			}
		}
		for _, i := range cctx.AllIdentifier() {
			mts.To = append(mts.To, conv.Identifier(i))
		}
		statement.OneOf = &pb.MoveStatement_MoveTo{
			MoveTo: mts,
		}
	} else if ictx := ctx.MoveCorrespondingToStatement(); ictx != nil {
		cctx := ictx.(*cobol85.MoveCorrespondingToStatementContext)
		mts := &pb.MoveCorrespondingToStatement{}
		if imts := cctx.MoveCorrespondingToSendingArea(); imts != nil {
			cmts := imts.(*cobol85.MoveCorrespondingToSendingAreaContext)
			if cmts.Identifier() != nil {
				mts.SendingArea = conv.Identifier(cmts.Identifier())
			}
		}
		for _, i := range cctx.AllIdentifier() {
			mts.To = append(mts.To, conv.Identifier(i))
		}
		statement.OneOf = &pb.MoveStatement_MoveCorrespondingTo{
			MoveCorrespondingTo: mts,
		}
	}
	v.sentence.Statements = append(v.sentence.Statements, &pb.Statement{
		OneOf: &pb.Statement_MoveStatement{
			MoveStatement: statement,
		},
	})
	return v.VisitChildren(ctx)
}

func (v *SentenceVisitor) VisitStatement(ctx *cobol85.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SentenceVisitor) VisitSentence(ctx *cobol85.SentenceContext) any {
	return v.VisitChildren(ctx)
}

func (v *SentenceVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *SentenceVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
