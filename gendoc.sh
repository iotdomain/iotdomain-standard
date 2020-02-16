# Generate Myzone RFC internet draft using kdrfc v3
cat template.txt myzone-standard.md > draft-myzone-00.md
kdrfc -3 draft-myzone-00.md 
