# Generate Myzone RFC internet draft using kdrfc v3
# Replace embedded images ![...]  with reference [...]  
cat template.txt myzone-standard.md | sed -e 's/\!\[/\[/g' > draft-myzone-00.md

kdrfc -3 draft-myzone-00.md 

