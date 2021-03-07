# This script converts all learnxinyminutes' docs into go source files
# for easy embedding 
import sys
import os

if len(sys.argv) < 2:
    print("Usage: pkg.py <path-to learnxinyminutes-docs>")
    sys.exit(1)

template = """
package pages

const {lang} = {src}
"""

ls = []

for entry in os.scandir(sys.argv[1]):
    if not entry.is_file or not entry.name.endswith(".markdown") or entry.name.startswith("README") or entry.name.startswith("CONTRIBUTING"):
        continue

    with open(sys.argv[1] + "/" + entry.name, "r") as f:
        # Skip the frontmatter
        f = f.read()
        print(f"{entry.name}: File Size: {len(f) / 1000} KB")
        if f.startswith("---"):
            f = "---".join(f.split("---")[2:]).strip()
        name = entry.name.split(".")[0]
        with open(f"pages/{name.replace('+', 'plus')}.go", "w") as w:
            # Store the whole source as a string, but be sure to escape some things to avoid breaking the string.
            src = "\""
            src += f.replace("\\", "\\\\").replace("\n", "\\n").replace("\t", "\\t").replace("\r", "\\r").replace('"', '\\"')
            src += '"'
            w.write(template.format(lang=name.title().replace("-", "_").replace("+", "Plus"), src=src))
            ls.append(name.lower())

ls.sort()

src2 = """package pages

var Pages = map[string]string{
"""

for x in ls:
    src2 += f"  \"{x}\": {x.title().replace('-', '_').replace('+', 'Plus')},\n"

src2 += "}"

with open("pages/pages.go", "w") as f:
    f.write(src2)
