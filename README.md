# goimports_degrouper
Micro utility to degroup imports in Go files

Goimports has one problem, if your import section in file already separated on groups by blank line, 
goimport ignore this imports and doesn't rearrange it. 

This small utility written in Go takes filename as only one argument and remove al blank lines (degrouping) import section. 
So after that file can be secsesfully processed by goimports.
