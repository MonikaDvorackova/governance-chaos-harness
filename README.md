# Governance Chaos Harness

Deterministic randomized testing of the fail-closed approval invariant. Mutates event order, duplicates gates, and injects rejections across 50,000 sequences.

Run `go run .`. The invariant is explicit: deployment is permitted only if every required gate has an approved event and no rejection exists. Extend with persisted state, retries, and concurrent event delivery.
