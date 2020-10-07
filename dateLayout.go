package iotmakerdbmongodbutilschema

var dateLayout = "Jan 2, 2006 at 3:04pm (MST)"

// DefineNewDateLayout (English): Define a new date layout for parse. Default layout is
// long form, "Jan 2, 2006 at 3:04pm (MST)".
//
// Documentation: https://golang.org/pkg/time/#example_Parse
//
//   Golang default layout are:
//   time.ANSIC       = "Mon Jan _2 15:04:05 2006"
//   time.UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
//   time.RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
//   time.RFC822      = "02 Jan 06 15:04 MST"
//   time.RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
//   time.RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
//   time.RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
//   time.RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
//   time.RFC3339     = "2006-01-02T15:04:05Z07:00"
//   time.RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
//   time.Kitchen     = "3:04PM"
//   // Handy time stamps.
//   time.Stamp      = "Jan _2 15:04:05"
//   time.StampMilli = "Jan _2 15:04:05.000"
//   time.StampMicro = "Jan _2 15:04:05.000000"
//   time.StampNano  = "Jan _2 15:04:05.000000000"
//
// DefineNewDateLayout (Português): Define um nodo layout para o parser de data. O layout
// padrão é a forma long, "Jan 2, 2006 at 3:04pm (MST)".
//
// documentação: https://golang.org/pkg/time/#example_Parse
//
//   Os layouts padrão do Golang são:
//   time.ANSIC       = "Mon Jan _2 15:04:05 2006"
//   time.UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
//   time.RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
//   time.RFC822      = "02 Jan 06 15:04 MST"
//   time.RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
//   time.RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
//   time.RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
//   time.RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
//   time.RFC3339     = "2006-01-02T15:04:05Z07:00"
//   time.RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
//   time.Kitchen     = "3:04PM"
//   // Handy time stamps.
//   time.Stamp      = "Jan _2 15:04:05"
//   time.StampMilli = "Jan _2 15:04:05.000"
//   time.StampMicro = "Jan _2 15:04:05.000000"
//   time.StampNano  = "Jan _2 15:04:05.000000000"
func DefineNewDateLayout(layout string) {
	dateLayout = layout
}
