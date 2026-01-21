package core

import (
	"testing"
)

func TestStringToJsonArrayEmpty(t *testing.T) {
	result := StringToJsonArray("")
	if len(result) != 0 {
		t.Errorf("Expected empty array, got %v", result)
	}
}

func TestStringToJsonArraySingleValue(t *testing.T) {
	result := StringToJsonArray("iqn.1993-08.org.debian:01:abc123")
	if len(result) != 1 {
		t.Errorf("Expected 1 element, got %d", len(result))
	}
	if result[0] != "iqn.1993-08.org.debian:01:abc123" {
		t.Errorf("Expected iqn value, got %v", result[0])
	}
}

func TestStringToJsonArrayCommaSeparated(t *testing.T) {
	result := StringToJsonArray("iqn1,iqn2,iqn3")
	if len(result) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(result))
	}
	if result[0] != "iqn1" || result[1] != "iqn2" || result[2] != "iqn3" {
		t.Errorf("Expected [iqn1,iqn2,iqn3], got %v", result)
	}
}

func TestStringToJsonArrayCommaSeparatedWithSpaces(t *testing.T) {
	result := StringToJsonArray("iqn1, iqn2 , iqn3")
	if len(result) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(result))
	}
	if result[0] != "iqn1" || result[1] != "iqn2" || result[2] != "iqn3" {
		t.Errorf("Expected trimmed values, got %v", result)
	}
}

func TestStringToJsonArrayJsonArray(t *testing.T) {
	result := StringToJsonArray(`["iqn1","iqn2"]`)
	if len(result) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(result))
	}
	if result[0] != "iqn1" || result[1] != "iqn2" {
		t.Errorf("Expected [iqn1,iqn2], got %v", result)
	}
}

func TestStringToJsonArrayJsonArraySingle(t *testing.T) {
	result := StringToJsonArray(`["single"]`)
	if len(result) != 1 {
		t.Errorf("Expected 1 element, got %d", len(result))
	}
	if result[0] != "single" {
		t.Errorf("Expected 'single', got %v", result[0])
	}
}

func TestStringToJsonArrayInvalidJson(t *testing.T) {
	// Invalid JSON should fall back to comma-split
	result := StringToJsonArray(`[invalid`)
	if len(result) != 1 {
		t.Errorf("Expected 1 element (fallback), got %d", len(result))
	}
	if result[0] != "[invalid" {
		t.Errorf("Expected '[invalid', got %v", result[0])
	}
}

func TestStringToJsonArrayEmptyElements(t *testing.T) {
	// Empty elements should be skipped
	result := StringToJsonArray("iqn1,,iqn2,")
	if len(result) != 2 {
		t.Errorf("Expected 2 elements (empty skipped), got %d: %v", len(result), result)
	}
}

func TestStringToJsonArrayALL(t *testing.T) {
	result := StringToJsonArray("ALL")
	if len(result) != 1 {
		t.Errorf("Expected 1 element, got %d", len(result))
	}
	if result[0] != "ALL" {
		t.Errorf("Expected 'ALL', got %v", result[0])
	}
}
