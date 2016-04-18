package color_test

import (
	"fmt"
	"os"

	"github.com/nhooyr/color"
)

func Example_attributes() {
	// "panic:" with a red foreground then normal "rip".
	color.Fprintf(os.Stdout, "%h[fgRed]panic:%r %s\n", "rip")

	// "panic:" with a brightRed background then normal "rip".
	color.Hprintf("%h[bgBrightRed]panic:%r %s\n", "rip")

	// Bold "panic:" then normal "rip".
	color.Hprintf("%h[bold]panic:%r %s\n", "rip")

	// Underlined "panic:" with then normal "rip".
	color.Hprintf("%h[underline]panic:%r %s\n", "rip")

	// "panic:" using color 83 as the foreground then normal "rip".
	color.Hprintf("%h[fg83]panic:%r %s\n", "rip")

	// "panic:" using color 158 as the background then normal "rip".
	color.Hprintf("%h[bg158]panic:%r %s\n", "rip")
}

func Example_mixing() {
	// Bolded "panic:" with a green foreground then normal "rip".
	color.Hprintf("%h[fgGreen+bold]panic:%r %s\n", "rip")

	// Underlined "panic:" with a bright black background then normal "rip".
	color.Hprintf("%h[bg8+underline]panic:%r %s\n", "rip")
}

func Example_reset() {
	// "rip" will be printed with a blue foreground and bright black background
	// because we never reset the highlighting after "panic:". The blue foreground is
	// carried on from "panic:".
	color.Hprintf("%h[fgBlue+bgBlack]panic: %h[bg8]%s\n", "rip")

	// The attributes carry onto anything written to the terminal until reset.
	// This prints "rip" in the same attributes as above.
	fmt.Println("rip")

	// Resets the highlighting and then prints "hello" normally.
	color.Hprintf("%r%s", "hello")
}

func ExamplePrepare() {
	// Prepare processes the highlight verbs in the string only once,
	// letting you print it repeatedly with performance.
	f := color.Prepare("%h[fgRed+bold]panic:%r %s\n")

	// Each prints bolded "panic:" with a red foreground and some normal text after.
	color.Hprintf(f, "rip")
	color.Hprintf(f, "yippie")
	color.Hprintf(f, "dsda")
}

func ExamplePrinter() {
	// "hi" with red foreground.
	p := color.NewPrinter(os.Stderr, color.EnableColor)
	p.Hprintf("%h[fgRed]%s%r\n", "hi")

	// normal "hi", the highlight verbs are ignored.
	p = color.NewPrinter(os.Stderr, color.DisableColor)
	p.Hprintf("%h[fgRed]%s%r\n", "hi")

	// If os.Stderr is a terminal, this will print in color.
	// Otherwise it will be a normal "hi".
	p = color.NewPrinter(os.Stderr, color.PerformCheck)
	p.Hprintf("%h[fgRed]%s%r\n", "hi")
}

func ExampleLogger() {
	// "hi" with a red foreground.
	l := color.NewLogger(os.Stderr, "%h[bold]color:%r ", 0, color.EnableColor)
	l.Hprintf("%h[fgRed]%s%r", "hi")

	// normal "hi", the highlight verbs are ignored.
	l = color.NewLogger(os.Stderr, "%h[bold]color:%r ", 0, color.DisableColor)
	l.Hprintf("%h[fgRed]%s%r", "hi")

	// If os.Stderr is a terminal, this will print in color.
	// Otherwise it will be a normal "hi".
	l = color.NewLogger(os.Stderr, "%h[bold]color:%r ", 0, color.PerformCheck)
	l.Fatalf("%h[fgRed]%s%r", "hi")
}
