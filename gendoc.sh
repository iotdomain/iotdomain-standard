# Generate IoT Domain RFC internet draft using kdrfc v3
# Replace embedded images ![...]  with reference [...]  
cat template.txt iotdomain-standard.md | sed -e 's/\!\[/\[/g' > iotdomain-draft-00.md

kdrfc -3 iotdomain-draft-00.md 

