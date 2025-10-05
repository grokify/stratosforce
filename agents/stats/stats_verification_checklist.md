# Statistics Verification Checklist

## Purpose
This checklist guides manual verification of statistics extracted by the Statistics Research & Validation Agent before use in papers, presentations, or other professional work.

---

## Pre-Verification Setup

- [ ] Obtain the JSON output from the stats-researcher agent
- [ ] Create a verification tracking spreadsheet or document
- [ ] Set aside adequate time (estimate 5-10 minutes per statistic)
- [ ] Ensure access to all source URLs (check for paywalls/access restrictions)

---

## For Each Statistic: Verification Process

### Step 1: Source Accessibility
- [ ] Click the source URL and confirm it loads successfully
- [ ] Verify the URL matches the description (correct organization/publication)
- [ ] Check publication date matches what's claimed
- [ ] Note if source is behind paywall or requires authentication

**If source is inaccessible:**
- [ ] Try alternative access methods (library access, institutional login)
- [ ] Search for archived versions (Wayback Machine, Google Cache)
- [ ] Mark as `source_inaccessible` in verification status
- [ ] Consider excluding this statistic from your work

### Step 2: Statistic Presence Verification
- [ ] Locate the specific statistic within the source document
- [ ] Use Ctrl+F / Cmd+F to search for the exact number
- [ ] Verify the statistic appears in the claimed context
- [ ] Confirm it's not taken out of context or misrepresented

**If statistic cannot be found:**
- [ ] Check if it might be in a table, chart, or figure
- [ ] Verify you're looking at the right section/page
- [ ] Check for different number formats (e.g., 14M vs 14,000,000)
- [ ] If still not found, mark as `verified_inaccurate`

### Step 3: Number Accuracy Check
- [ ] Confirm the exact number matches (no rounding errors)
- [ ] Verify the unit of measurement is correct
- [ ] Check for any qualifiers (e.g., "approximately", "estimated")
- [ ] Confirm any percentage signs, decimal places, or significant figures

**Common errors to watch for:**
- Million vs. Billion confusion
- Percentage vs. percentage point confusion
- Annual vs. quarterly figures
- Nominal vs. real/adjusted values

### Step 4: Context Verification
- [ ] Verify the time period/date is accurate
- [ ] Confirm geographic scope (global, national, regional)
- [ ] Check methodology notes are accurate
- [ ] Verify any population or sample size claims
- [ ] Confirm margin of error if mentioned

### Step 5: Cross-Reference Check (if applicable)
- [ ] If agent claims multiple sources confirm this statistic, verify each one
- [ ] Check that cross-references truly report the same figure
- [ ] Note any variations in how the statistic is presented
- [ ] Verify the sources are truly independent (not citing each other)

### Step 6: Source Credibility Assessment
- [ ] Confirm the source type is correctly categorized
- [ ] Verify claimed credibility rating is appropriate
- [ ] Check author credentials if applicable
- [ ] For case studies, verify the case_study_details are accurate

### Step 7: Documentation
- [ ] Update the `verification` object in the JSON with:
  - Your name/ID in `verified_by`
  - Today's date in `verification_date`
  - Appropriate `verification_status`
  - Any notes in `verification_notes`
- [ ] Flag any discrepancies for investigation
- [ ] Note any additional context you discovered

---

## Overall Report Verification

### After Verifying Individual Statistics:

- [ ] Review the summary section - does it accurately reflect findings?
- [ ] Check that key_findings list contains only verified statistics
- [ ] Verify the metadata counts are accurate
- [ ] Review discrepancies section for accuracy
- [ ] Ensure limitations section includes verification limitations

### Final Quality Check:

- [ ] Calculate verification rate: (verified accurate / total statistics) %
- [ ] Assess if remaining statistics are sufficient for your needs
- [ ] Identify any critical gaps in the research
- [ ] Decide if additional research is needed
- [ ] Determine if confidence levels should be adjusted based on verification

---

## Red Flags to Watch For

🚩 **Immediate Concerns:**
- URL returns 404 or different content than described
- Statistic does not appear anywhere in the source
- Number is significantly different from what's reported
- Source is not credible or authoritative for the topic
- Multiple statistics from same source fail verification

🚩 **Moderate Concerns:**
- Minor discrepancies in numbers (rounding differences)
- Context missing or incomplete
- Publication date uncertain or unclear
- Cross-references don't fully align
- Source credibility lower than claimed

🚩 **Interpretation Issues:**
- Statistic taken out of context
- Important caveats or limitations omitted
- Generalizability overstated
- Timeframe ambiguous or incorrect
- Geographic scope unclear

---

## Verification Status Guidelines

### Use `verified_accurate` when:
- Statistic found exactly as reported in source
- Context accurately represented
- No material discrepancies
- Source is accessible and credible

### Use `verified_inaccurate` when:
- Statistic not found in source
- Number is wrong or significantly different
- Context is misrepresented
- Source is misattributed

### Use `source_inaccessible` when:
- URL is broken or paywalled
- Source requires authentication you don't have
- Page has been removed or moved
- Content is not available in your region

### Use `not_verified` when:
- Haven't completed verification yet
- Deferred to later review
- Low priority statistic

---

## Best Practices

### Before Using Statistics in Your Work:
1. ✅ Verify at least 100% of statistics you plan to directly cite
2. ✅ Spot-check at least 25% of background statistics
3. ✅ Always verify the most critical/impactful claims
4. ✅ Keep verification records for audit trail
5. ✅ Note verification date (statistics can become outdated)

### When Writing Your Paper/Presentation:
1. ✅ Only cite statistics marked as `verified_accurate`
2. ✅ Include the original source in your citations, not the agent
3. ✅ Preserve all context, qualifiers, and limitations
4. ✅ Note if statistic is from a single source vs. confirmed by multiple
5. ✅ Update your own citations/bibliography with proper formatting

### If Verification Reveals Issues:
1. ✅ Remove or replace inaccurate statistics
2. ✅ Lower confidence levels if appropriate
3. ✅ Add caveats or qualifications
4. ✅ Consider re-running the agent with different parameters
5. ✅ Document what was changed and why

---

## Verification Record Template

Use this template to track your verification work:

```
Statistics Verification Record
===============================
Project: [Your project name]
Verifier: [Your name]
Date: [YYYY-MM-DD]
Agent Output File: [filename.json]

Summary:
- Total statistics in report: __
- Statistics verified: __
- Verified accurate: __
- Verified inaccurate: __
- Source inaccessible: __
- Verification rate: __%

Issues Found:
1. [Describe any issues]
2. [...]

Actions Taken:
1. [What you did to address issues]
2. [...]

Approved for use: YES / NO / CONDITIONAL
Conditions: [If conditional, what conditions must be met]

Notes:
[Any additional observations]
```

---

## Time Estimates

- Simple statistic (accessible source, clear number): 3-5 minutes
- Complex statistic (long document, table/chart data): 10-15 minutes
- Problematic statistic (broken link, hard to find): 20+ minutes

**Example:** Verifying a 20-statistic report thoroughly might take 2-4 hours.

---

## When to Re-Run the Agent

Consider having the agent search again if:
- More than 30% of statistics fail verification
- Critical statistics are inaccurate or inaccessible
- You discover better sources during verification
- The topic requires more current data
- You need statistics the agent didn't find

---

## Additional Resources

- **Citation Management**: Use tools like Zotero, Mendeley, or EndNote to manage verified sources
- **Fact-Checking**: Cross-reference with fact-checking sites for controversial claims
- **Data Repositories**: Check original data sources (government databases, research repositories)
- **Peer Review**: Have a colleague verify critical statistics independently

---

## Final Checklist Before Publication

- [ ] All cited statistics are verified accurate
- [ ] Source URLs are current and accessible
- [ ] Citations properly formatted per your style guide
- [ ] Context and limitations appropriately noted
- [ ] Verification records archived for future reference
- [ ] Any necessary permissions obtained for using data
- [ ] Co-authors/advisors have reviewed the statistics

---

**Remember:** The agent is a research assistant, not a substitute for due diligence. Always verify before you trust!