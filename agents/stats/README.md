# Statistics Research & Validation System

A comprehensive system of Claude Code agents for finding, validating, and verifying statistics from web sources with strict anti-hallucination safeguards. Designed for academic papers, professional presentations, and research requiring highly accurate, properly cited statistics.

---

## 📦 Detailed Installation Guide

### Understanding Claude Code Agent Locations

Claude Code looks for agent definitions in:
- `~/claude/agents/` - User agents (recommended)
- Project-specific `.claude/agents/` - Project-scoped agents
- System agents (Claude Code built-ins)

### Step-by-Step: File System Installation

#### Step 1: Create Directory Structure

```bash
# Create agents directory
mkdir -p ~/claude/agents

# Create resources directory for checklist
mkdir -p ~/claude/resources

# Optional: Create project-specific location
mkdir -p ~/your-project/.claude/agents
```

#### Step 2: Download or Copy Agent Files

**Option A: Download from repository**
```bash
cd ~/claude/agents

# Download orchestrator
curl -O https://your-repo/Statistics_Research_Orchestrator_Agent.md
mv Statistics_Research_Orchestrator_Agent.md stats_orchestrator.md

# Optional: Download individual agents
curl -O https://your-repo/Statistics_Research_Validation_Agent.md
mv Statistics_Research_Validation_Agent.md stats_researcher.md

curl -O https://your-repo/Statistics_Verification_Agent.md
mv Statistics_Verification_Agent.md stats_verifier.md
```

**Option B: Manual file creation**
```bash
# Create the orchestrator agent
cat > ~/claude/agents/stats_orchestrator.md << 'EOF'
# Statistics Research Orchestrator Agent

## Purpose
[Paste the full content from Statistics_Research_Orchestrator_Agent.md]
EOF
```

**Option C: Copy from local files**
```bash
# If you have the files locally
cp Statistics_Research_Orchestrator_Agent.md ~/claude/agents/stats_orchestrator.md
cp Statistics_Research_Validation_Agent.md ~/claude/agents/stats_researcher.md
cp Statistics_Verification_Agent.md ~/claude/agents/stats_verifier.md
```

#### Step 3: Install Verification Checklist

```bash
# Save to shared resources
cp Statistics_Verification_Checklist.md ~/claude/resources/verification_checklist.md

# Or save to your project
cp Statistics_Verification_Checklist.md ~/your-project/docs/verification_checklist.md
```

#### Step 4: Verify Installation

```bash
# List installed agents
ls -la ~/claude/agents/

# Should show:
# stats_orchestrator.md
# stats_researcher.md (optional)
# stats_verifier.md (optional)

# Restart Claude Code (if needed)
# Then test:
/agent stats_orchestrator --help
```

### File Naming Conventions

**Important:** File names determine agent invocation:

| File Name | Invocation Command | Notes |
|-----------|-------------------|-------|
| `stats_orchestrator.md` | `/agent stats_orchestrator` | Recommended |
| `stats-orchestrator.md` | `/agent stats-orchestrator` | Also works |
| `statsOrchestrator.md` | `/agent statsOrchestrator` | camelCase works |

**We recommend:** `snake_case` format (`stats_orchestrator.md`)

### Project-Specific Installation

For project-scoped agents (only available in that project):

```bash
cd ~/your-project

# Create project agents directory
mkdir -p .claude/agents

# Copy agent files
cp ~/downloads/Statistics_Research_Orchestrator_Agent.md \
   .claude/agents/stats_orchestrator.md

# Add to .gitignore if you don't want to commit
echo ".claude/agents/" >> .gitignore

# Or commit if you want team access
git add .claude/agents/
git commit -m "Add statistics research agents"
```

### Verification Checklist Integration

The checklist should be a **standalone reference document**, not embedded in agents. Here's why:

**❌ Don't embed in agent:**
- Makes agent file too large
- Harder to maintain
- Checklist is for humans, not agents
- Can't print or share easily

**✅ Do keep separate:**
- Easy to update independently
- Can be printed for reference
- Shareable with team members
- Can be customized per project

**Recommended locations:**

```bash
# Global (for all projects)
~/claude/resources/verification_checklist.md

# Project-specific
~/your-project/docs/verification_checklist.md
~/your-project/.claude/resources/verification_checklist.md

# Team shared (in repository)
~/your-project/docs/statistics_verification_process.md
```

**Reference from agent (optional):**

If you want the agent to mention the checklist location, add this to the agent file:

```markdown
## Human Verification

For detailed manual verification steps, see the verification checklist at:
- Global: `~/claude/resources/verification_checklist.md`
- Project: `.claude/resources/verification_checklist.md`

Open this file in your editor for step-by-step verification guidance.
```

### Directory Structure Example

```
~/claude/
├── agents/
│   ├── stats_orchestrator.md          # Main orchestrator agent
│   ├── stats_researcher.md             # Research agent (optional)
│   └── stats_verifier.md               # Verification agent (optional)
└── resources/
    └── verification_checklist.md       # Human verification guide

~/your-project/
├── .claude/
│   ├── agents/                         # Project-specific agents (optional)
│   └── resources/
│       └── verification_checklist.md   # Project-specific checklist
├── research_output.json                # Generated by agents
├── verified_output.json                # Generated by agents
└── docs/
    └── statistics_sources.md           # Your research documentation
```

### Updating Agents

When new versions are released:

```bash
# Backup current version
cp ~/claude/agents/stats_orchestrator.md \
   ~/claude/agents/stats_orchestrator.md.backup

# Download new version
curl -o ~/claude/agents/stats_orchestrator.md \
  https://your-repo/Statistics_Research_Orchestrator_Agent.md

# Restart Claude Code
# Test the updated agent
/agent stats_orchestrator "test query"
```

### Troubleshooting Installation

**Agent not found:**
```bash
# Check if file exists
ls -la ~/claude/agents/stats_orchestrator.md

# Check permissions
chmod 644 ~/claude/agents/stats_orchestrator.md

# Restart Claude Code
```

**Agent file format issues:**
- Ensure file is plain text Markdown (`.md` extension)
- Check for proper UTF-8 encoding
- Verify no binary characters

**Path issues on Windows:**
- Use: `%USERPROFILE%\claude\agents\stats_orchestrator.md`
- Or: `C:\Users\YourName\claude\agents\stats_orchestrator.md`

**Path issues on macOS/Linux:**
- Use: `~/claude/agents/stats_orchestrator.md`
- Ensure `~` expands to your home directory

---

## 🎯 Quick Overview

This system provides **three ways to get verified statistics**:

### 🌟 Recommended: One-Command Orchestrator
```bash
/agent stats-orchestrator "Find statistics on global EV adoption in 2024"
```
→ **2-5 minutes** → Researched AND verified statistics, ready to use

### ⚙️ Advanced: Individual Agents  
```bash
/agent stats-researcher "Your query"
/agent stats-verifier --input research_output.json
```
→ **5-15 minutes** → Full control over each step

### 📋 Manual: Human Verification
→ **Additional 2-4 hours** → Use checklist for critical work

---

## 📦 What's Included

| Component | Purpose | When to Use |
|-----------|---------|-------------|
| **stats-orchestrator** ⭐ | One-command research + verification | Most use cases (90%) |
| **stats-researcher** | Find and extract statistics | Custom workflows, learning |
| **stats-verifier** | Verify stats against sources | Custom workflows, debugging |
| **Verification Checklist** | Human review guide | Critical work, final checks |
| **Go Package** | Programmatic data access | Integration, automation |

---

## 🚀 Quick Start

### Installation

**Method 1: Automated Installation (Easiest)**

```bash
# Download and run the installation script
curl -O https://your-repo/setup.sh
chmod +x setup.sh
./setup.sh
```

The script will:
- Create `~/claude/agents/` and `~/claude/resources/` directories
- Install the orchestrator agent
- Install the verification checklist
- Optionally install individual agents
- Backup any existing files

---

**Method 2: Manual File System Installation**

1. **Create the agents directory** (if it doesn't exist):
```bash
mkdir -p ~/claude/agents
```

2. **Download and save the agent files**:

# Save the orchestrator agent

```bash
curl -o ~/claude/agents/stats_orchestrator.md https://raw.githubusercontent.com/grokify/stratosforce/refs/heads/main/agents/stats/stats_orchestrator_agent.md

# Or manually save the files to:

```
# ~/claude/agents/stats_orchestrator_agent.md
# ~/claude/agents/stats_researcher_agent.md
# ~/claude/agents/stats_verifier_agent.md
```

3. **Save the verification checklist**:

# Save to a shared location

```bash
mkdir -p ~/claude/resources
curl -o ~/claude/resources/verification_checklist.md 'https://raw.githubusercontent.com/grokify/stratosforce/refs/heads/main/agents/stats/stats_verification_checklist.md'
```

# Or save to your project directory

```
cp stats_verification_checklist.md ~/your-project/
```

4. **Restart Claude Code** (if needed) and the agents will be automatically available.

**That's it!** The agents are now available via `/agent stats_orchestrator`.

---

**Method 3: Interactive Installation**

If you prefer to use Claude Code's interactive agent creation:

```bash
/agent create stats-orchestrator
# Paste content from Statistics_Research_Orchestrator_Agent.md when prompted
```

<details>
<summary><strong>Optional: Install Individual Agents</strong></summary>

**File System Method:**
```bash
# Save individual agent files
cp Statistics_Research_Validation_Agent.md ~/claude/agents/stats_researcher.md
cp Statistics_Verification_Agent.md ~/claude/agents/stats_verifier.md
```

**Interactive Method:**
```bash
# Research agent
/agent create stats-researcher
# Paste from stats_research_agent.md

# Verification agent  
/agent create stats-verifier
# Paste from stats_verification_agent.md
```
</details>

---

### Verification Checklist Setup

The verification checklist can be used in two ways:

**Option 1: Standalone File (Recommended)**
- Best for: Human reviewers, printing, sharing with team
- Location: `~/claude/resources/verification_checklist.md` or your project folder
- Usage: Open in your editor or Markdown viewer when needed

**Option 2: Reference in Agent (Advanced)**
- Best for: Agent can reference it during verification
- The agent Markdown files can include a path reference:
```markdown
For detailed verification guidance, see:
file://~/claude/resources/verification_checklist.md
```

**We recommend Option 1** - Keep the checklist as a standalone reference document that humans use after the automated verification process.

### First Query

```bash
/agent stats-orchestrator "Find statistics on renewable energy growth in 2024"
```

**You'll get:**
- ✅ **READY FOR USE** - Statistics verified, use with confidence (70%+ verification)
- ⚠️ **NEEDS REVIEW** - Some issues found, review before using (40-70%)
- ❌ **RE-RESEARCH** - Quality too low, try different approach (<40%)

---

## 📖 System Architecture

```
┌─────────────────────────────────────────┐
│      stats-orchestrator (Front-end)     │
│  • Coordinates workflow                 │
│  • Quality assessment                   │
│  • Smart presentation                   │
└──────────────┬──────────────────────────┘
               │
      ┌────────┴────────┐
      ↓                 ↓
┌─────────────────┐ ┌──────────────────┐
│stats-researcher │→│ stats-verifier   │
│ • Web search    │ │ • Fetch sources  │
│ • Extract stats │ │ • Verify accuracy│
│ • Cross-validate│ │ • Flag issues    │
└─────────────v───┘ └──────────────────┘
      │                 │
      └────────┬────────┘
               ↓
      ┌────────────────┐
      │  JSON Output   │
      │ + Verification │
      │   Report       │
      └────────────────┘
               ↓
      ┌────────────────┐
      │ Human Review   │
      │  (if needed)   │
      └────────────────┘
```

---

## 🎯 Usage Examples

### Example 1: Academic Research

**Query:**
```bash
/agent stats-orchestrator \
  "Find statistics on climate change economic impact: GDP loss, \
   infrastructure costs, adaptation spending. Prefer peer-reviewed sources."
```

**Result in 3 minutes:**
```
Status: ✅ READY FOR USE
Verification: 12/14 statistics verified (86%)

Ready-to-Use Statistics:
• Global economic cost by 2050: $23T (World Bank, verified)
• Annual adaptation spending needed: $300B (UNEP, verified)
• GDP impact under 2.5°C warming: 7.9% (Nature Climate, verified)
[+ 9 more verified statistics]

Next Steps:
✅ Use these statistics with confidence
📝 Cite using provided URLs
```

### Example 2: Business Presentation

**Query:**
```bash
/agent stats-orchestrator \
  --url "https://www.gartner.com/market-report" \
  --mode expansion \
  "Extract cloud market statistics and find corroborating data"
```

**Result:**
```
Status: ⚠️ NEEDS REVIEW
Verification: 10/15 statistics verified (67%)

Ready-to-Use Statistics: [10 verified stats]

⚠️ Requires Attention:
• Market growth: Gartner 21.5% vs IDC 18.3% (different timeframes)
• 2 statistics in charts (manual extraction needed)

Recommendation: Use verified stats, note timeframe differences
```

### Example 3: Fact-Checking

**Query:**
```bash
/agent stats-orchestrator \
  "Verify claim: EV sales hit 18M units in 2024, 25% market share"
```

**Result:**
```
Status: ⚠️ CLAIM PARTIALLY INACCURATE

Sales Volume: ✅ Accurate (17.8M ≈ 18M)
Market Share: ❌ Inaccurate (claimed 25%, actual 18-19%)

Verdict: Sales correct, market share overstated by ~7 points
```

---

## 🔄 Workflows

### Workflow 1: Orchestrated (Recommended)

```bash
# One command
/agent stats-orchestrator "your query"
```

**Timeline:**
- 0:00 - Agent starts research
- 0:30 - Research complete (8-15 statistics found)
- 0:45 - Verification begins automatically
- 2:00 - Verification complete
- 2:05 - Results presented with quality assessment

**When to use:** 95% of cases

---

### Workflow 2: Manual Control

```bash
# Step 1: Research (30-60 seconds)
/agent stats-researcher "your query"

# Step 2: Review research_output.json
cat research_output.json

# Step 3: Verify (1-3 minutes)
/agent stats-verifier --input research_output.json

# Step 4: Review verification_report.md
cat verification_report.md
```

**When to use:** Learning, debugging, custom integration

---

### Workflow 3: High-Stakes (Orchestrator + Manual)

```bash
# Step 1: Automated research + verification
/agent stats-orchestrator "your query"

# Step 2: Review orchestrator output
# Even if ✅ READY FOR USE

# Step 3: Manual spot-checks
# Use checklist for 3-5 most critical statistics

# Step 4: Use with high confidence
```

**When to use:** Academic papers, legal filings, board presentations

---

## 📊 Understanding Results

### ✅ READY FOR USE (≥70% verified)

**What it means:**
- Most statistics verified in sources
- Quality meets publication standards
- Safe to use with normal practices

**What to do:**
1. Use the verified statistics
2. Cite sources properly
3. Optional: Spot-check 1-2 sources
4. Include context in your work

---

### ⚠️ NEEDS REVIEW (40-70% verified)

**What it means:**
- Some stats verified, others have issues
- Usable but requires attention
- May have discrepancies or inaccessible sources

**What to do:**
1. Use the verified statistics  
2. Review flagged items with checklist
3. Manually verify inaccessible sources
4. Exclude or correct inaccurate statistics

---

### ❌ RE-RESEARCH RECOMMENDED (<40% verified)

**What it means:**
- Too few reliable statistics found
- Widespread verification failures
- Not suitable for use as-is

**What to do:**
1. Read orchestrator's analysis
2. Try suggested refinements
3. Consider if data exists on topic
4. Use different search strategy

---

## 🎓 Best Practices

### Research Phase

**✅ DO:**
- Be specific: "Global EV sales 2024 by region"
- Specify sources: "Prefer government and academic sources"
- Include timeframe: "Most recent data available"
- Choose right mode: Known source? Use expansion

**❌ DON'T:**
- Be vague: "EV statistics"
- Accept old data without checking
- Ignore verification warnings
- Skip source quality assessment

### Verification Phase

**✅ DO:**
- Always run verification (orchestrator does this)
- Review flagged statistics manually
- Use checklist for critical work
- Document your verification process

**❌ DON'T:**
- Use unverified statistics
- Ignore low verification rates
- Skip manual review for high-stakes work
- Trust without verifying sources

### Publication Phase

**✅ DO:**
- Use only verified_accurate statistics
- Cite original sources (not Claude)
- Include context and limitations
- Keep audit trail (save all JSON files)

**❌ DON'T:**
- Cite "Claude" or the agents
- Omit caveats or qualifications
- Use statistics marked as inaccurate
- Delete verification records

---

## 🔧 Advanced Features

### Research Modes

**Multi-Source (Default)**
```bash
/agent stats-orchestrator "topic"
```
- Searches multiple credible sources
- Cross-validates findings
- Best for: General research

**Single-Source**
```bash
/agent stats-orchestrator \
  --url "https://source.com/report" \
  --mode single-source
```
- Extracts from one document only
- No external searches
- Best for: Analyzing a known report

**Expansion**
```bash
/agent stats-orchestrator \
  --url "https://source.com/report" \
  --mode expansion
```
- Starts with provided source
- Finds corroborating data
- Best for: Validating a primary source

### Custom Thresholds

```bash
/agent stats-orchestrator \
  --min-verification-rate 80 \
  --verification-scope high-priority \
  "your query"
```

---

## 🛡️ Anti-Hallucination Safeguards

### Research Agent Protections
- ✅ Never invents statistics
- ✅ Never fabricates URLs
- ✅ Reports exact figures only
- ✅ Only links verified sources
- ✅ Transparent about confidence

### Verification Agent Protections
- ✅ Actually fetches every source
- ✅ Searches for exact statistics
- ✅ Reports what it finds (or doesn't)
- ✅ Flags uncertainties
- ✅ Never assumes correctness

### Orchestrator Protections
- ✅ Won't present unverified stats
- ✅ Quality gates (70% threshold)
- ✅ Transparent about issues
- ✅ Guides remediation
- ✅ Prevents premature use

---

## 📚 Component Documentation

### Individual Agent Docs

- **Orchestrator Agent:** [Statistics_Research_Orchestrator_Agent.md](./Statistics_Research_Orchestrator_Agent.md)
- **Research Agent:** [Statistics_Research_Validation_Agent.md](./Statistics_Research_Validation_Agent.md)
- **Verification Agent:** [Statistics_Verification_Agent.md](./Statistics_Verification_Agent.md)
- **Verification Checklist:** [Statistics_Verification_Checklist.md](./Statistics_Verification_Checklist.md)

### Go Package

- **Package Docs:** [go-package/README.md](./go-package/README.md)
- **Source Code:** [go-package/statsresearch.go](./go-package/statsresearch.go)
- **Tests:** [go-package/statsresearch_test.go](./go-package/statsresearch_test.go)
- **Examples:** [go-package/examples.go](./go-package/examples.go)

---

## 💼 Use Cases

| Use Case | Recommended Approach | Time Estimate |
|----------|---------------------|---------------|
| Academic paper | Orchestrator + Manual review | 30-60 min |
| Business presentation | Orchestrator only | 5-10 min |
| Blog post / article | Orchestrator only | 3-5 min |
| Fact-checking | Orchestrator only | 2-5 min |
| Legal filing | Orchestrator + Full manual | 2-4 hours |
| Grant proposal | Orchestrator + Manual review | 45-90 min |
| News article | Orchestrator only | 5-10 min |
| Research report | Orchestrator + Spot checks | 20-40 min |

---

## ❓ FAQ

### Q: Can I trust results without manual verification?

**A:** If orchestrator says ✅ READY FOR USE and verification rate is ≥70%, you can use them with normal citation practices. However, we recommend manual spot-checks for:
- Academic papers (verify key statistics)
- High-stakes presentations
- Legal or regulatory work

### Q: What's the difference between validation and verification?

**A:**
- **Validation** = Cross-referencing during research (agent checks multiple sources agree)
- **Verification** = Confirming stats actually appear in sources (agent fetches and checks the actual source documents)

### Q: How long does the process take?

**A:**
- Orchestrator: 2-5 minutes for most queries
- Manual agents: 5-15 minutes
- Human review: 5-10 minutes per statistic (if needed)

### Q: What if sources are paywalled?

**A:** The verification agent will mark them as "source_inaccessible." Try:
- Institutional library access
- Author's website or ResearchGate
- Contact authors directly
- Use statistics from accessible sources only

### Q: Can I modify the agents?

**A:** Yes! All agents are in Markdown and fully customizable. You can:
- Adjust source preferences
- Modify verification thresholds
- Add custom validation rules
- Change output formats

### Q: What if I find an error after publication?

**A:**
1. Check your verification records
2. Re-verify the statistic
3. Issue correction/errata if needed
4. Update JSON files for future reference
5. Review process to prevent recurrence

### Q: How should I cite these statistics?

**A:** Cite the original source, NEVER Claude or the agents:

**Good:** "Global EV sales reached 14M units (IEA, 2024)."
**Bad:** "According to Claude, EV sales..."

Always use source URLs from JSON output.

---

## 🔧 Troubleshooting

### Low Verification Rate (<50%)

**Causes:**
- Query too broad
- Topic has limited data
- Sources behind paywalls
- Marketing content (unverifiable)

**Solutions:**
- Narrow your query
- Add source preferences
- Try expansion mode with known source
- Check if data exists on topic

### Keep Getting "Re-Research Recommended"

**Causes:**
- Query too vague
- Wrong time period
- Niche topic
- Wrong sources

**Solutions:**
- Use orchestrator's suggestions
- Try different timeframe
- Break into smaller queries
- Use expansion mode

### Statistics Don't Match Expectations

**Possible reasons:**
- Your source might be outdated
- Different methodology
- Different time periods
- Your assumptions incorrect

**What to do:**
- Review verified sources
- Check methodologies
- Look for timeframe differences
- Consider both might be valid

---

## 🎉 Benefits Summary

### Time Savings
- ⏱️ **Manual research:** 2-4 hours
- ⏱️ **With orchestrator:** 5-15 minutes
- 📊 **Time saved:** ~90%

### Quality Improvements
- ✅ Eliminates hallucinations (triple verification)
- ✅ Proper source attribution (every stat linked)
- ✅ Transparent confidence levels
- ✅ Complete audit trail

### Use Case Coverage
- 📝 Academic papers
- 💼 Business presentations
- 📰 Journalism
- 🏛️ Policy analysis
- 💰 Grant writing
- ✓ Fact-checking

---

## 📞 Support & Resources

- **Issues:** Report bugs or request features
- **Discussions:** Ask questions, share use cases
- **Documentation:** Full technical documentation
- **Examples:** Additional example workflows

---

## 🚀 Getting Started Checklist

- [ ] Install stats-orchestrator agent
- [ ] Run first test query
- [ ] Review output format
- [ ] Try different research modes
- [ ] Download verification checklist
- [ ] (Optional) Install Go package
- [ ] (Optional) Install individual agents

---

**Version:** 1.0.0  
**Last Updated:** October 2025  
**Compatibility:** Claude Code with Claude Sonnet 4.5+

**Built for researchers who demand accuracy and rigorous verification.**
