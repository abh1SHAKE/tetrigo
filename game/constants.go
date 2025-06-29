package game

const (
	BlockSize = 30
	GridRows = 20
	GridColumns = 10

	GridHeight = GridRows * BlockSize
	GridWidth = GridColumns * BlockSize

	UIPanelWidth = 400
	TopPadding, LeftPadding, BottomPadding = 100, 60, 60

	ScreenHeight = GridHeight + TopPadding + BottomPadding
	ScreenWidth = GridWidth + UIPanelWidth
)