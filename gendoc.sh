# Generate Myzone RFC internet draft using kdrfc v3
cat template.txt myzone-convention.md > draft-myzone-00.md
kdrfc -3 draft-myzone-00.md 
