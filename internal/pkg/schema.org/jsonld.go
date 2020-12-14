package schemaorg

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	ldjsonrxp *regexp.Regexp
)

func init() {
	ldjsonrxp = regexp.MustCompile(`type="application/ld\+json"[^>]*>(([^<]*)"@type":"Product"([^<]*))</`)
}

type Product struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       []string `json:"image"`
}

func FindProductCtx(ctx context.Context, url string) (*Product, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't get: %w", err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read body: %s", err)
	}

	found := ldjsonrxp.FindAllSubmatch(b, -1)

	var (
		ldstr []byte
	)

	for _, str := range found {
		if len(str) > 2 {
			ldstr = str[1]
		}
	}

	doc := &Product{}

	err = json.Unmarshal(ldstr, &doc)
	if err != nil {
		return nil, fmt.Errorf("can't umarshal: %s", err)
	}

	return doc, nil
}
