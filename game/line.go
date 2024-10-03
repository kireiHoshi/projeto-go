package game

/**
Line of the Board, has * places
 */
type Line []string

/**
Creates a new Line with an argument size to be added to the Board
 */
func (l Line) GetNewLine(axisY int) Line{

	for x:=0; x<axisY; x++{
		l = append(l, " - ")
	}

	return l
}
