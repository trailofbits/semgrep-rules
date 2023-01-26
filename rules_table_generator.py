#!/usr/bin/env python

from pathlib import Path
from urllib.parse import quote
import yaml
import sys


LANGUAGES = ['go', 'python', 'rs']
IMPACT_MAP = {
    'LOW': "🟩",
    'MEDIUM': "🟧",
    'HIGH': "🟥",
    None: "🌫️",
}
CONFIDENCE_MAP = {
    'LOW': "🌕",
    'MEDIUM': "🌗",
    'HIGH': "🌘",
    None: "",
}

def main():
    for language in LANGUAGES:

        rules_for_lang = []
        for rule in Path(language).rglob('*.yaml'):
            try:
                rule_data = yaml.safe_load(rule.open())
            except yaml.YAMLError as err:
                print(f"Error reading {rule} - {err}", file=sys.stderr)
                continue

            if rule_data is None or 'rules' not in rule_data:
                print(f"Error for {rule} - missing rules", file=sys.stderr)
                continue
            rule_data = rule_data['rules']
            if len(rule_data) > 1:
                print(f"Error for {rule} - only one rule per file is supported", file=sys.stderr)
                continue
            if len(rule_data) == 0:
                print(f"Error for {rule} - missing any rule", file=sys.stderr)
                continue
            rule_data = rule_data[0]

            rule_link = '.'.join(rule.parts[1:-1] + (rule.stem,))
            rule_link = '.'.join(("r/trailofbits", language, rule_link, rule_data['id']))
            rules_for_lang.append((rule, rule_data, rule_link))

        if len(rules_for_lang) > 0:
            print(f"### {language}")
            print("")
            print("| ID | Playground | Impact | Confidence | Description |")
            print("| -- | :--------: | :----: | :--------: | ----------- |")

            for rule, rule_data, rule_link in sorted(rules_for_lang):
                rule_meta = rule_data.get('metadata', {})
                print(f"| [{rule_data['id']}]({rule}) | [🛝🔗](https://semgrep.dev/playground/{quote(rule_link)}) | {IMPACT_MAP[rule_meta.get('impact')]} | {CONFIDENCE_MAP[rule_meta.get('confidence')]} | {rule_meta.get('description', '')} |")

            print("\n")


if __name__ == "__main__":
    main()
