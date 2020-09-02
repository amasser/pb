package pb

import (
	"fmt"
	"strings"
	"sync"

	"github.com/verabull/pb/internal/app"
)

// LsdBar - Just Progress Bar
type LsdBar struct {
	length       int
	title        string
	prefix       rune
	suffix       rune
	MaxIteration int
	lock         sync.Mutex

	progress *progress
}

type progress struct {
	maxFilled int
	filled    int
	unfilled  int
	padding   int

	currentIteration int
	currentPercent   float64

	unfilledSymbol rune
	filledSymbol   rune

	finish bool
}

// DefaultLsdBar - Init DefaultLsdBar without customization
func DefaultLsdBar(MaxIteration int, options ...Option) (*LsdBar, error) {
	bar := &LsdBar{
		length:       app.GetWidth(),
		title:        "Test Lsd Bar",
		prefix:       '|',
		suffix:       '|',
		MaxIteration: MaxIteration,

		progress: &progress{
			maxFilled: 0,
			filled:    0,
			unfilled:  0,
			padding:   9,

			currentIteration: 0,
			currentPercent:   0,

			unfilledSymbol: ' ',
			filledSymbol:   '█',

			finish: false,
		},
	}

	for _, o := range options {
		o(bar)
	}

	bar.progress.unfilled = bar.length - len(bar.title) - bar.progress.padding
	bar.progress.maxFilled = bar.progress.unfilled

	err := bar.render()
	if err != nil {
		return nil, err
	}

	return bar, nil
}

// Add - Add new iteration and count percent
func (b *LsdBar) Add() error {

	if !b.progress.finish {
		b.lock.Lock()
		defer b.lock.Unlock()

		b.progress.currentIteration++
		b.progress.currentPercent = float64(b.progress.currentIteration) / float64(b.MaxIteration) * 100

		b.progress.filled = int(float64(b.progress.currentIteration) * float64(b.progress.maxFilled) / float64(b.MaxIteration))
		b.progress.unfilled = b.progress.maxFilled - b.progress.filled

		err := b.render()
		if err != nil {
			return err
		}

		if b.progress.currentPercent >= 100 {
			_, err := fmt.Print("\n")
			if err != nil {
				return err
			}

			b.progress.finish = true

		}
	}

	return nil
}

func (b *LsdBar) render() error {

	err := b.clear()
	if err != nil {
		return err
	}

	filled := app.TextRainbow(strings.Repeat(string(b.progress.filledSymbol), b.progress.filled))

	_, err = fmt.Printf("\x1b[1m%v\x1b[1m %c%v%v%c \x1b[1m%d%%\x1b[1m",
		b.title,
		b.prefix,
		filled,
		strings.Repeat(string(b.progress.unfilledSymbol), b.progress.unfilled),
		b.suffix,
		int(b.progress.currentPercent),
	)

	if err != nil {
		return err
	}

	return nil
}

func (b *LsdBar) clear() error {

	_, err := fmt.Printf("\r%s\r", strings.Repeat(" ", b.length))
	if err != nil {
		return err
	}

	return nil
}

// Option - Options struct
type Option func(b *LsdBar)

// SetWidth - Set Bar Width
func SetWidth(w uint) Option {
	return func(b *LsdBar) {
		if b.length > int(w) {
			b.length = int(w)
		}
	}
}

// SetTitle - Set Bar Title
func SetTitle(t string) Option {
	return func(b *LsdBar) {
		b.title = t
	}
}

// SetPrefix - Set Bar Prefix
func SetPrefix(p rune) Option {
	return func(b *LsdBar) {
		b.prefix = p
	}
}

// SetSuffix - Set Bar Suffix
func SetSuffix(s rune) Option {
	return func(b *LsdBar) {
		b.suffix = s
	}
}

// SetFilledSymbol - Set Filled Symbol
func SetFilledSymbol(f rune) Option {
	return func(b *LsdBar) {
		b.progress.filledSymbol = f
	}
}

// SetUnfilledSymbol - Set Unfilled Symbol
func SetUnfilledSymbol(uf rune) Option {
	return func(b *LsdBar) {
		b.progress.unfilledSymbol = uf
	}
}
