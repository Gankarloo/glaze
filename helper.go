package glaze

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Glaze struct{
	lipgloss.
}
// BorderWithTitle renders a string inside a box with a title embedded in the top border.
// position: 0.0 (left), 0.5 (center), 1.0 (right)
func BorderWithTitle(content string, title string, style lipgloss.Style, position lipgloss.Position) string {
	// 1. Render the internal content to find out how wide the box should be
	// if the style has a fixed width, use that; otherwise, calculate from content.
	renderedBox := style.Render(content)
	lines := strings.Split(renderedBox, "\n")
	boxWidth := lipgloss.Width(lines[0])

	// 2. Style the border and title
	borderStyle := lipgloss.NewStyle().
		Foreground(style.GetBorderTopForeground()).
		Render

	// We add a little padding so the border doesn't touch the text
	titleStyle := lipgloss.NewStyle().
		Foreground(style.GetBorderTopForeground()).
		Padding(0, 1).
		Render

	renderedTitle := titleStyle(title)
	titleWidth := lipgloss.Width(renderedTitle)

	// 3. Create the custom Top Border using JoinHorizontal
	// We subtract 2 from the boxWidth to account for the corners (TopLeft, TopRight)
	remainingSpace := boxWidth - 2 - titleWidth
	if remainingSpace < 0 {
		return renderedBox // Title too long, return default
	}

	// Calculate how many border characters go on the left vs right
	// leftCount := int(float64(remainingSpace) * position)
	leftCount := int(float64(remainingSpace) * float64(position))
	rightCount := remainingSpace - leftCount

	// Adjust title position for left and right alignment
	switch position {
	case lipgloss.Left:
		leftCount += 2
		rightCount -= 2
	case lipgloss.Right:
		leftCount -= 2
		rightCount += 2
	}

	leftBar := strings.Repeat(style.GetBorderStyle().Top, leftCount)
	rightBar := strings.Repeat(style.GetBorderStyle().Top, rightCount)

	// Build the top line: [Corner][LeftBar][Title][RightBar][Corner]
	topLine := lipgloss.JoinHorizontal(lipgloss.Bottom,
		borderStyle(style.GetBorderStyle().TopLeft),
		borderStyle(leftBar),
		renderedTitle,
		borderStyle(rightBar),
		borderStyle(style.GetBorderStyle().TopRight),
	)

	// 4. Swap the original top line with our new titled line
	lines[0] = topLine
	return strings.Join(lines, "\n")
}
