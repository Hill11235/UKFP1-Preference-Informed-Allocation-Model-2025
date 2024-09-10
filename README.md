# deanery-model

Uses a rough Monte Carlo simulation model to determine probability of being placed in a given deanery given a selection. Uses the [existing first preference competition ratios](https://foundationprogramme.nhs.uk/programmes/2-year-foundation-programme/ukfp/competition-ratios/) as well as an attempt to replicate the preference information allocation algorithm that is used.

### Overview

Uses the existing first preference competition ratios to estimate overall deanery popularity. Use this to simulate student location selections (number equal to total of available positions), then run PIA m times with a given preference selection. Use the m PIA runs to determine probability of ending up in a given deanery based on the selection.


### Preference Informed Allocation (PIA)
Based on the information and flow diagram shared [here](https://madeinheene.hee.nhs.uk/Portals/12/UKFP%202024%20Applicant%20Guide%20to%20Allocation%20-%20Preference%20Informed%20Allocation%20.pdf). Two passes are made:

- 1st pass-algorithm works through randomly generated list. If a place is available in applicant's first choice foundation school, they will be allocated. If not, they will be skipped. Works through every applicant, giving only first choices if they are available.
- 2nd pass-algorithm will again work through list. Any unplaced applicants will be allocated a place in their highest preference that still has places. (After the first pass).

See UKFP 2024 Preference Informed Allocation Webinar on youtube for more information.

### Assumptions
- Does not consider pre-allocation.
- Does not consider linked applications.
- Number of available allocations == number of students applying.
- Currently assumes relative popularity based on ratios minus those locations which are over-subscribed. This will should reflect reality as no one should be mad enough to put consecutive over subscribed deaneries at the top of their list.
- Loads of others.
