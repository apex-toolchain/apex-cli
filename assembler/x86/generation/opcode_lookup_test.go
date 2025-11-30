package generation

import (
	"reflect"
	"testing"
)

func TestEncodeMovR64R64(t *testing.T) {
	// MOV r64, r64 : mov rax, rbx -> 48 8B C3
	out := EncodeOperation("MOV_r64_r64", []Operand{{Type: Reg, Name: "rax"}, {Type: Reg, Name: "rbx"}})
	expected := []byte{0x48, 0x8B, 0xC3}
	if !reflect.DeepEqual(out, expected) {
		t.Fatalf("unexpected bytes for MOV_r64_r64: got % X, want % X", out, expected)
	}
}

func TestEncodeMovR64Imm64(t *testing.T) {
	// MOV r64, imm64 : mov rax, 0x1122334455667788 -> 48 B8 8877665544332211
	imm := uint64(0x1122334455667788)
	out := EncodeOperation("MOV_r64_imm64", []Operand{{Type: Reg, Name: "rax"}, {Type: Imm, UImm64: imm}})
	expected := []byte{0x48, 0xB8, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11}
	if !reflect.DeepEqual(out, expected) {
		t.Fatalf("unexpected bytes for MOV_r64_imm64: got % X, want % X", out, expected)
	}
}

func TestEncodeMovRm64R64(t *testing.T) {
	// MOV r/m64, r64 : mov rax, rbx -> 48 89 D8
	out := EncodeOperation("MOV_rm64_r64", []Operand{{Type: Reg, Name: "rax"}, {Type: Reg, Name: "rbx"}})
	expected := []byte{0x48, 0x89, 0xD8}
	if !reflect.DeepEqual(out, expected) {
		t.Fatalf("unexpected bytes for MOV_rm64_r64: got % X, want % X", out, expected)
	}
}
