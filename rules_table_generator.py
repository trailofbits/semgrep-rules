#!/usr/bin/env python

from pathlib import Path
from urllib.parse import quote
import yaml
import sys

LANGUAGES = ['go', 'python', 'rs', 'javascript']
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
        for rule_path in Path(language).rglob('*.yaml'):
            try:
                rules_data = yaml.safe_load(rule_path.open())
            except yaml.YAMLError as err:
                print(f"Error reading {rule_path} - {err}", file=sys.stderr)
                continue

            if rules_data is None or 'rules' not in rules_data:
                print(f"Error for {rule_path} - missing rules", file=sys.stderr)
                continue

            rules_data = rules_data['rules']
            if len(rules_data) == 0:
                print(f"Error for {rule_path} - missing any rule", file=sys.stderr)
                continue

            for rule_data in rules_data:
                rule_link = '.'.join(rule_path.parts[1:-1] + (rule_path.stem,))
                rule_link = '.'.join(("r/trailofbits", language, rule_link, rule_data['id']))
                rules_for_lang.append((rule_path, rule_data, rule_link))

        if len(rules_for_lang) > 0:
            print(f"### {language}")
            print("")
            print("| ID | Playground | Impact | Confidence | Description |")
            print("| -- | :--------: | :----: | :--------: | ----------- |")

            for rule_path, rule_data, rule_link in sorted(rules_for_lang, key=lambda x: (x[0], x[1]['id'])):
                rule_meta = rule_data.get('metadata', {})
                print(
                    f"| [{rule_data['id']}]({rule_path}) | [🛝🔗](https://semgrep.dev/playground/{quote(rule_link)}) | {IMPACT_MAP[rule_meta.get('impact')]} | {CONFIDENCE_MAP[rule_meta.get('confidence')]} | {rule_meta.get('description', '')} |"
                )

            print("\n")


if __name__ == "__main__":
    main()
