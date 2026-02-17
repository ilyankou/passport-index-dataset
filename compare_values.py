#!/usr/bin/env python3
"""Compare unique Requirement values between current and legacy tidy CSVs."""

import csv
import os

def get_requirements(path):
    reqs = set()
    with open(path) as f:
        reader = csv.reader(f)
        next(reader)
        for row in reader:
            if len(row) >= 3:
                reqs.add(row[2])
    return reqs

legacy_dirs = sorted(d for d in os.listdir("legacy") if os.path.isdir(f"legacy/{d}"))
if not legacy_dirs:
    print("No legacy directories found")
    exit(1)

latest = legacy_dirs[-1]
print(f"Comparing with legacy/{latest}\n")
old = get_requirements(f"legacy/{latest}/passport-index-tidy.csv")
new = get_requirements("passport-index-tidy.csv")

print("=== Only in LEGACY ===")
for v in sorted(old - new):
    print(f"  {v}")

print()
print("=== Only in NEW ===")
for v in sorted(new - old):
    print(f"  {v}")

print()
print("=== In both ===")
for v in sorted(old & new):
    print(f"  {v}")
