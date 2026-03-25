package expr

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestRenderExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		expr ast.Expr
		want string
	}{
		// Field
		{
			name: "Field returns selector expression",
			expr: Field("t", "Name"),
			want: "t.Name",
		},

		// Ident
		{
			name: "Ident returns identifier",
			expr: Ident("x"),
			want: "x",
		},

		// IntLit
		{
			name: "IntLit returns integer literal",
			expr: IntLit("42"),
			want: "42",
		},
		{
			name: "IntLit with zero",
			expr: IntLit("0"),
			want: "0",
		},
		{
			name: "IntLit with negative",
			expr: IntLit("-1"),
			want: "-1",
		},

		// StringLit
		{
			name: "StringLit returns quoted string",
			expr: StringLit("hello"),
			want: `"hello"`,
		},
		{
			name: "StringLit with empty string",
			expr: StringLit(""),
			want: `""`,
		},
		{
			name: "StringLit escapes special characters",
			expr: StringLit(`say "hi"`),
			want: `"say \"hi\""`,
		},

		// Nil
		{
			name: "Nil returns nil identifier",
			expr: Nil(),
			want: "nil",
		},

		// Not
		{
			name: "Not negates expression",
			expr: Not(Ident("x")),
			want: "!x",
		},
		{
			name: "Not with field access",
			expr: Not(Field("t", "Valid")),
			want: "!t.Valid",
		},

		// Paren
		{
			name: "Paren wraps expression in parentheses",
			expr: Paren(Ident("x")),
			want: "(x)",
		},
		{
			name: "Paren with binary expression",
			expr: Paren(And(Ident("a"), Ident("b"))),
			want: "(a && b)",
		},

		// Binary
		{
			name: "Binary with addition",
			expr: Binary(Ident("a"), token.ADD, Ident("b")),
			want: "a + b",
		},

		// GT
		{
			name: "GT with identifiers",
			expr: GT(Ident("a"), Ident("b")),
			want: "a > b",
		},
		{
			name: "GT with field and int literal",
			expr: GT(Field("t", "Age"), IntLit("18")),
			want: "t.Age > 18",
		},

		// GTE
		{
			name: "GTE with identifiers",
			expr: GTE(Ident("a"), Ident("b")),
			want: "a >= b",
		},
		{
			name: "GTE with field and int literal",
			expr: GTE(Field("t", "Score"), IntLit("0")),
			want: "t.Score >= 0",
		},

		// LT
		{
			name: "LT with identifiers",
			expr: LT(Ident("a"), Ident("b")),
			want: "a < b",
		},
		{
			name: "LT with field and int literal",
			expr: LT(Field("t", "Count"), IntLit("100")),
			want: "t.Count < 100",
		},

		// LTE
		{
			name: "LTE with identifiers",
			expr: LTE(Ident("a"), Ident("b")),
			want: "a <= b",
		},
		{
			name: "LTE with field and int literal",
			expr: LTE(Field("t", "Len"), IntLit("255")),
			want: "t.Len <= 255",
		},

		// Eq
		{
			name: "Eq with identifiers",
			expr: Eq(Ident("a"), Ident("b")),
			want: "a == b",
		},
		{
			name: "Eq with field and nil",
			expr: Eq(Field("t", "Ptr"), Nil()),
			want: "t.Ptr == nil",
		},
		{
			name: "Eq with field and string literal",
			expr: Eq(Field("t", "Status"), StringLit("active")),
			want: `t.Status == "active"`,
		},

		// NEq
		{
			name: "NEq with identifiers",
			expr: NEq(Ident("a"), Ident("b")),
			want: "a != b",
		},
		{
			name: "NEq with field and nil",
			expr: NEq(Field("t", "Err"), Nil()),
			want: "t.Err != nil",
		},

		// And
		{
			name: "And with identifiers",
			expr: And(Ident("a"), Ident("b")),
			want: "a && b",
		},
		{
			name: "And with comparisons",
			expr: And(GT(Ident("x"), IntLit("0")), LT(Ident("x"), IntLit("100"))),
			want: "x > 0 && x < 100",
		},

		// Or
		{
			name: "Or with identifiers",
			expr: Or(Ident("a"), Ident("b")),
			want: "a || b",
		},
		{
			name: "Or with comparisons",
			expr: Or(Eq(Ident("s"), StringLit("yes")), Eq(Ident("s"), StringLit("no"))),
			want: `s == "yes" || s == "no"`,
		},

		// Call without package
		{
			name: "Call without package prefix",
			expr: Call("", "len", Field("t", "Items")),
			want: "len(t.Items)",
		},
		{
			name: "Call without package and no args",
			expr: Call("", "foo"),
			want: "foo()",
		},
		{
			name: "Call without package and multiple args",
			expr: Call("", "append", Ident("s"), Ident("v")),
			want: "append(s, v)",
		},

		// Call with package
		{
			name: "Call with package prefix",
			expr: Call("utf8", "RuneCountInString", Field("t", "Name")),
			want: "utf8.RuneCountInString(t.Name)",
		},
		{
			name: "Call with package and multiple args",
			expr: Call("fmt", "Sprintf", StringLit("%d"), Ident("n")),
			want: `fmt.Sprintf("%d", n)`,
		},
		{
			name: "Call with net package",
			expr: Call("net", "ParseIP", Field("t", "IP")),
			want: "net.ParseIP(t.IP)",
		},

		// MethodCall
		{
			name: "MethodCall without args",
			expr: MethodCall(Ident("ip"), "To4"),
			want: "ip.To4()",
		},
		{
			name: "MethodCall with args",
			expr: MethodCall(Ident("s"), "Contains", StringLit("@")),
			want: `s.Contains("@")`,
		},
		{
			name: "MethodCall on field",
			expr: MethodCall(Field("t", "Name"), "String"),
			want: "t.Name.String()",
		},

		// Composed expressions
		{
			name: "Not with function call",
			expr: Not(Call("", "isValid", Field("t", "Email"))),
			want: "!isValid(t.Email)",
		},
		{
			name: "Eq with method call result and nil",
			expr: Eq(MethodCall(Ident("ip"), "To4"), Nil()),
			want: "ip.To4() == nil",
		},
		{
			name: "Nested binary expressions",
			expr: And(GT(Field("t", "Age"), IntLit("0")), LTE(Field("t", "Age"), IntLit("150"))),
			want: "t.Age > 0 && t.Age <= 150",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Render(tt.expr)
			if got != tt.want {
				t.Errorf("Render() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestAndAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		exprs []ast.Expr
		want  string
		isNil bool
	}{
		{
			name:  "zero expressions returns nil",
			exprs: nil,
			isNil: true,
		},
		{
			name:  "single expression returns it unchanged",
			exprs: []ast.Expr{Ident("a")},
			want:  "a",
		},
		{
			name:  "two expressions joined with AND",
			exprs: []ast.Expr{Ident("a"), Ident("b")},
			want:  "a && b",
		},
		{
			name:  "three expressions left-associative",
			exprs: []ast.Expr{Ident("a"), Ident("b"), Ident("c")},
			want:  "a && b && c",
		},
		{
			name: "four expressions with comparisons",
			exprs: []ast.Expr{
				GT(Ident("a"), IntLit("0")),
				LT(Ident("a"), IntLit("100")),
				NEq(Ident("b"), Nil()),
				Eq(Ident("c"), StringLit("ok")),
			},
			want: `a > 0 && a < 100 && b != nil && c == "ok"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := AndAll(tt.exprs...)
			if tt.isNil {
				if got != nil {
					t.Errorf("AndAll() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("AndAll() returned nil, want non-nil")
			}

			rendered := Render(got)
			if rendered != tt.want {
				t.Errorf("Render(AndAll()) = %q, want %q", rendered, tt.want)
			}
		})
	}
}

func TestRenderStmt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		stmt ast.Stmt
		want string
	}{
		{
			name: "ShortVarDecl with function call",
			stmt: ShortVarDecl("ip", Call("net", "ParseIP", Field("t", "IP"))),
			want: "ip := net.ParseIP(t.IP)",
		},
		{
			name: "ShortVarDecl with identifier",
			stmt: ShortVarDecl("x", IntLit("42")),
			want: "x := 42",
		},
		{
			name: "ShortVarDecl with nil",
			stmt: ShortVarDecl("err", Nil()),
			want: "err := nil",
		},
		{
			name: "ShortVarDecl with method call",
			stmt: ShortVarDecl("v", MethodCall(Ident("s"), "String")),
			want: "v := s.String()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := RenderStmt(tt.stmt)
			if got != tt.want {
				t.Errorf("RenderStmt() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "simple field access",
			input: "t.Name",
			want:  "t.Name",
		},
		{
			name:  "comparison expression",
			input: "t.Name > 5",
			want:  "t.Name > 5",
		},
		{
			name:  "function call",
			input: "len(t.Items)",
			want:  "len(t.Items)",
		},
		{
			name:  "logical expression",
			input: "a && b",
			want:  "a && b",
		},
		{
			name:  "complex expression",
			input: `t.Age > 0 && t.Age <= 150`,
			want:  "t.Age > 0 && t.Age <= 150",
		},
		{
			name:  "equality with string",
			input: `t.Status == "active"`,
			want:  `t.Status == "active"`,
		},
		{
			name:  "nil comparison",
			input: "t.Ptr == nil",
			want:  "t.Ptr == nil",
		},
		{
			name:    "invalid expression returns error",
			input:   "if x {",
			wantErr: true,
		},
		{
			name:    "empty string returns error",
			input:   "",
			wantErr: true,
		},
		{
			name:    "unbalanced parentheses returns error",
			input:   "(a + b",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Parse(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Error("Parse() error = nil, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Parse() unexpected error: %v", err)
			}

			rendered := Render(got)
			if rendered != tt.want {
				t.Errorf("Render(Parse(%q)) = %q, want %q", tt.input, rendered, tt.want)
			}
		})
	}
}

func TestRenderNilAndInvalid(t *testing.T) {
	t.Parallel()

	t.Run("Render with nil expression returns empty string", func(t *testing.T) {
		t.Parallel()

		// format.Node returns an error when given nil, so Render returns "".
		got := Render(nil)
		if got != "" {
			t.Errorf("Render(nil) = %q, want %q", got, "")
		}
	})

	t.Run("RenderStmt with nil statement returns empty string", func(t *testing.T) {
		t.Parallel()

		got := RenderStmt(nil)
		if got != "" {
			t.Errorf("RenderStmt(nil) = %q, want %q", got, "")
		}
	})
}

func TestFieldProperties(t *testing.T) {
	t.Parallel()

	t.Run("Field returns SelectorExpr with correct identifiers", func(t *testing.T) {
		t.Parallel()

		sel := Field("t", "Name")
		ident, ok := sel.X.(*ast.Ident)
		if !ok {
			t.Fatal("Field().X is not *ast.Ident")
		}
		if ident.Name != "t" {
			t.Errorf("Field().X.Name = %q, want %q", ident.Name, "t")
		}
		if sel.Sel.Name != "Name" {
			t.Errorf("Field().Sel.Name = %q, want %q", sel.Sel.Name, "Name")
		}
	})
}

func TestIntLitKind(t *testing.T) {
	t.Parallel()

	lit := IntLit("42")
	if lit.Kind != token.INT {
		t.Errorf("IntLit().Kind = %v, want %v", lit.Kind, token.INT)
	}
	if lit.Value != "42" {
		t.Errorf("IntLit().Value = %q, want %q", lit.Value, "42")
	}
}

func TestStringLitKind(t *testing.T) {
	t.Parallel()

	lit := StringLit("hello")
	if lit.Kind != token.STRING {
		t.Errorf("StringLit().Kind = %v, want %v", lit.Kind, token.STRING)
	}
	if lit.Value != `"hello"` {
		t.Errorf("StringLit().Value = %q, want %q", lit.Value, `"hello"`)
	}
}

func TestShortVarDeclTokenType(t *testing.T) {
	t.Parallel()

	stmt := ShortVarDecl("x", IntLit("1"))
	if stmt.Tok != token.DEFINE {
		t.Errorf("ShortVarDecl().Tok = %v, want %v", stmt.Tok, token.DEFINE)
	}
	if len(stmt.Lhs) != 1 {
		t.Fatalf("ShortVarDecl().Lhs length = %d, want 1", len(stmt.Lhs))
	}
	lhsIdent, ok := stmt.Lhs[0].(*ast.Ident)
	if !ok {
		t.Fatal("ShortVarDecl().Lhs[0] is not *ast.Ident")
	}
	if lhsIdent.Name != "x" {
		t.Errorf("ShortVarDecl().Lhs[0].Name = %q, want %q", lhsIdent.Name, "x")
	}
}

func TestCallArgCount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		call    *ast.CallExpr
		wantLen int
	}{
		{
			name:    "zero args",
			call:    Call("", "foo"),
			wantLen: 0,
		},
		{
			name:    "one arg",
			call:    Call("", "len", Ident("s")),
			wantLen: 1,
		},
		{
			name:    "two args",
			call:    Call("", "append", Ident("s"), Ident("v")),
			wantLen: 2,
		},
		{
			name:    "three args",
			call:    Call("fmt", "Fprintf", Ident("w"), StringLit("%s"), Ident("v")),
			wantLen: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if len(tt.call.Args) != tt.wantLen {
				t.Errorf("Call().Args length = %d, want %d", len(tt.call.Args), tt.wantLen)
			}
		})
	}
}

func TestMethodCallStructure(t *testing.T) {
	t.Parallel()

	call := MethodCall(Ident("ip"), "To4")
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		t.Fatal("MethodCall().Fun is not *ast.SelectorExpr")
	}
	recvIdent, ok := sel.X.(*ast.Ident)
	if !ok {
		t.Fatal("MethodCall().Fun.X is not *ast.Ident")
	}
	if recvIdent.Name != "ip" {
		t.Errorf("MethodCall receiver = %q, want %q", recvIdent.Name, "ip")
	}
	if sel.Sel.Name != "To4" {
		t.Errorf("MethodCall method = %q, want %q", sel.Sel.Name, "To4")
	}
	if len(call.Args) != 0 {
		t.Errorf("MethodCall().Args length = %d, want 0", len(call.Args))
	}
}
