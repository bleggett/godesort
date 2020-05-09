package saturn

type SaturnImage struct {
	Title string
	DiscNumber int
	DiscCount int
	Region string
	Version string
	Date string
	Order string //zero padded subdir image is in, 01, 02, 03, etc
}
