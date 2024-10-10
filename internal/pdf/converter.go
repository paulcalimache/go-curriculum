package pdf

import (
	"bytes"
	"context"
	"log"
	"os"
	"sync"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func ConvertHtmlToPdf(file bytes.Buffer) error {
	ctx, cancelCtx := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancelCtx()

	var wg sync.WaitGroup
	wg.Add(1)

	if err := chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			lctx, cancel := context.WithCancel(ctx)
			chromedp.ListenTarget(lctx, func(ev interface{}) {
				if _, ok := ev.(*page.EventLoadEventFired); ok {
					wg.Done()
					// remove event listener
					cancel()
				}
			})
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, file.String()).Do(ctx)
		}),
		// wait for the page.EventLoadEventFired
		chromedp.ActionFunc(func(ctx context.Context) error {
			wg.Wait()
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.
				PrintToPDF().
				WithMarginBottom(0).
				WithMarginTop(0).
				WithMarginRight(0).
				WithMarginLeft(0).
				WithPaperHeight(11.69).
				WithPaperWidth(8.27).
				WithPrintBackground(true).
				Do(ctx)
			if err != nil {
				return err
			}
			return os.WriteFile("curriculum.pdf", buf, 0644)
		}),
	); err != nil {
		return err
	}
	return nil
}
