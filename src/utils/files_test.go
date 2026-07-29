package files

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestTruncateCommentBodyWithinLimit(t *testing.T) {
	body := strings.Repeat("a", MaxCommentBodyLength)

	got := TruncateCommentBody(body)
	if got != body {
		t.Fatalf("expected unchanged body of length %d, got %d", len(body), len(got))
	}
}

func TestTruncateCommentBodyOverLimit(t *testing.T) {
	body := strings.Repeat("b", MaxCommentBodyLength+1000)

	got := TruncateCommentBody(body)
	if utf8.RuneCountInString(got) != MaxCommentBodyLength {
		t.Fatalf("expected truncated body length %d runes, got %d", MaxCommentBodyLength, utf8.RuneCountInString(got))
	}
	if !strings.HasSuffix(got, truncationNotice) {
		t.Fatalf("expected truncation notice suffix, got %q", got[len(got)-len(truncationNotice):])
	}
}

func TestGetCommentBodyTruncatesContent(t *testing.T) {
	content := strings.Repeat("c", MaxCommentBodyLength+500)

	got := GetCommentBody(content, "")
	if utf8.RuneCountInString(got) != MaxCommentBodyLength {
		t.Fatalf("expected truncated body length %d runes, got %d", MaxCommentBodyLength, utf8.RuneCountInString(got))
	}
}

func TestTruncateCommentBodyUTF8Safe(t *testing.T) {
	body := strings.Repeat("中", MaxCommentBodyLength+100)

	got := TruncateCommentBody(body)
	if !utf8.ValidString(got) {
		t.Fatal("truncated body contains invalid UTF-8")
	}
	if utf8.RuneCountInString(got) != MaxCommentBodyLength {
		t.Fatalf("expected truncated body length %d runes, got %d", MaxCommentBodyLength, utf8.RuneCountInString(got))
	}
	if !strings.HasSuffix(got, truncationNotice) {
		t.Fatalf("expected truncation notice suffix")
	}
}
