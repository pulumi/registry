---
user-invocable: false
description: Review comment templates for different contributor types and actions
---

# Message Templates

Select template based on contributor type (from Step 1).

**Bot templates**: Professional/technical tone, no emojis

## Template Matrix

| Action | External | Internal | Bot |
|--------|----------|----------|-----|
| **Approve** | "Thank you! [≤1 sentence praise if warranted]. Welcome! 🎉" | "LGTM! [feedback/suggestions]" | **pulumi-bot**: "Metadata update reviewed. Changes verified."<br>**Other**: "Automated changes approved." |
| **Approve and merge** | "Thank you! [≤1 sentence praise]. Auto-merge enabled. 🎉" | "LGTM! Auto-merge enabled." | **pulumi-bot**: "Metadata update verified. Auto-merge enabled."<br>**Other**: "Automated changes approved. Auto-merge enabled." |
| **Make changes and approve** | "Applied minor formatting fixes. Thank you! 🙏" | "Applied formatting fixes. LGTM!" | N/A (excluded for bots) |
| **Request changes** | Opening: "Thank you!"<br>Acknowledge positives (≤1 sentence)<br>Issues with line numbers<br>Suggestion blocks<br>Close: "Mention @claude if you need help" | Professional opening<br>Issues with line numbers<br>Suggestion blocks<br>Clear action items | Technical issue description<br>Line numbers<br>What needs changing<br>No suggestion blocks<br>Close: "This automated PR may need closing and regeneration after fixing source configuration." |
| **Close PR** | "Thank you for contributing!"<br>Clear, kind closing reason<br>Acknowledge valuable aspects<br>Suggest alternatives<br>"We appreciate your interest in Pulumi!" | Clear closing reason | **pulumi-bot**: "Closing. [Technical reason: conflicts, superseded, etc.]"<br>**Other**: "Closing. [Technical reason]" |

## Tone Guidelines

### External Contributors

Warm and welcoming tone with explicit gratitude, appropriate emojis (🎉, 🙏), and community-building language.

**Example**: "Thank you for your first contribution to Pulumi Registry! This package addition looks great. Welcome! 🎉"

### Internal Contributors

Professional and efficient tone. Brief acknowledgments, technical focus, "LGTM" acceptable. Direct and clear.

**Example**: "LGTM! Nice improvements to the schema validation."

### Bot PRs (All Types)

Technical and factual tone. No emojis. State what was checked/verified.

**Example pulumi-bot**: "Package metadata update for aws v7.21.0 reviewed. Version bump and schema URL verified."

## Implementation Notes

- Always use confirmed/edited content from preview step
- Base template selection on contributor type detected in Step 1
- For pulumi-bot: Keep messages terse and technical
- Keep bot messages factual and concise
- Adjust tone but maintain professionalism across all contributor types
