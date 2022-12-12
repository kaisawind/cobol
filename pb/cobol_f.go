package pb

func (x *CompilationGroup) UnmarshalF(f string) (err error) {
	for {
		unit := &CompilationUnit{}
		err = unit.UnmarshalF(f)
		if err != nil {
			break
		}
		x.Units = append(x.Units, unit)
	}
	return
}

func (x *CompilationUnit) UnmarshalF(f string) (err error) {
	return
}

func (x *SourceUnit) UnmarshalF(f string) (err error) {
	const NEW_LINE = '\n'
	contents := []*Content{}
	isComment := false

	content := ""
	// 合并注释 和 多行代码
	for _, v := range f {
		if v == NEW_LINE {
			content += string(v)
			if len(content) > 7 {
				indicator := content[6:7]
				content = content[7:]
				switch indicator {
				case COMMENT_LINE, COMMENT_LINE_NEW:
					if isComment {
						contents[len(contents)-1].Content += content
					} else {
						contents = append(contents, &Content{
							Type:    Content_COMMENT,
							Content: content,
						})
					}
					isComment = true
				case CONTINUATION_LINE:
					contents[len(contents)-1].Content += content
					isComment = false
				case DEBUGGING_MODE:
					contents = append(contents, &Content{
						Type:    Content_DEBUG,
						Content: content,
					})
					isComment = false
				default:
					contents = append(contents, &Content{
						Type:    Content_CODE,
						Content: content,
					})
					isComment = false
				}
			}
			content = ""
		} else {
			content += string(v)
		}
	}

	x.Contents = contents
	return
}
