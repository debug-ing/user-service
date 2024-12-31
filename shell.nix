{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go_1_23
    pkgs.git
  ];

  shellHook = ''
    go install golang.org/x/lint/golint@latest    
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  '';
}