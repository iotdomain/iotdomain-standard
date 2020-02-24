# Generate IotConnect RFC internet draft using kdrfc v3
# Replace embedded images ![...]  with reference [...]  
cat template.txt iotconnect-standard.md | sed -e 's/\!\[/\[/g' > iotconnect-draft-00.md

kdrfc -3 iotconnect-draft-00.md 

