package calendar

type Event struct {
	Title string
	Date  // Embed a Date using an anonymous field
}
