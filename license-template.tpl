LICENSE SUMMARY
===============

License terms can be found at the bottom of this file.

{{ range . }}
{{ .Name }}:{{ .LicenseName }}
{{ end }}

THIRD-PARTY SOFTWARE USED
=========================



{{ range . }}

* Name: {{ .Name }}
* License: [{{ .LicenseName }}]({{ .LicenseURL }})

LICENSE REQUIREMENTS & SPECIFICATIONS
======================================
```
{{ .LicenseText }}
```
{{ end }}


SOLACE CORPORATION LICENSE AGREEMENT
Home page: https://solace.com/license-software

Solace
------

Version: 10/25/2018

SOLACE CORPORATION


LICENCE AGREEMENT FOR SOLACE SOFTWARE
======================================

```
THIS LICENCE AGREEMENT and any documents expressly referred to in this
agreement (the "Agreement") between SOLACE CORPORATION, a company
incorporated under the laws of the Province of Ontario ("SOLACE") and
licensee, the party identified in the Order (as defined below) or that
otherwise accepts this Agreement (the "Licensee") (together the "Parties",
and each a "Party"), is made on the Effective Date (as defined below).

BY ACCEPTING THE TERMS OF THIS AGREEMENT, EITHER BY: A) ACCEPTING THE
AGREEMENT ONLINE, B) SIGNING THE ORDER (AS DEFINED BELOW) WHICH REFERENCES
THIS AGREEMENT, OR C) INSTALLING OR USING THE SOFTWARE AFTER BEING MADE AWARE
OF THIS AGREEMENT, THE LICENSEE ACKNOWLEDGES THAT IT HAS READ AND UNDERSTOOD
ALL OF THE PROVISIONS, AND HAS THE AUTHORITY TO AGREE TO, AND IS CONFIRMING
THAT IT IS AGREEING TO, COMPLY WITH AND BE BOUND BY, ALL OF THE TERMS AND
CONDITIONS CONTAINED HEREIN, TOGETHER WITH THE TERMS SET FORTH IN ANY ORDER.
IF, AFTER READING THIS AGREEMENT, THE LICENSEE DOES NOT ACCEPT OR AGREE TO
THE TERMS AND CONDITIONS CONTAINED HEREIN, THE LICENSEE SHALL NOT INSTALL OR
USE THE SOFTWARE.

IF YOU ARE AN AGENT OR EMPLOYEE OF ANOTHER ENTITY THEN YOU HEREBY REPRESENT
AND WARRANT THAT: (I) THE INDIVIDUAL ACCEPTING THIS AGREEMENT IS DULY
AUTHORIZED TO ACCEPT THIS AGREEMENT ON SUCH ENTITY'S BEHALF AND TO BIND SUCH
ENTITY, AND (II) SUCH ENTITY HAS FULL POWER, CORPORATE OR OTHERWISE, TO ENTER
INTO THIS AGREEMENT AND PERFORM ITS OBLIGATIONS HEREUNDER.

1	INTERPRETATION

1.1	Definitions. In this Agreement the following terms shall have the
following meanings:

"Core" means (i) a single physical processor core or hyper-thread when Solace
PubSub+ software is deployed on either a bare-metal server or a cloud or
virtualization environment that presents physical cores to the software, and
(ii) a single virtual core when deployed in a cloud or virtualization
environment that presents virtual cores to the VMR.

"Documentation" means the documentation  made accessible by SOLACE via a URL
provided to Licensee.

"Order" means (i) an electronic form provided by SOLACE on its website for
ordering Software Subscriptions, Professional Services, and/or Support and
Maintenance Services, or (ii) a written document, including a Licensee
purchase order, executed by SOLACE and Licensee pursuant to which Licensee
purchases of Software Subscriptions, Professional Services, and/or Support
and Maintenance Services from SOLACE.

"Products" means the Software, Documentation, Support and Maintenance
Services, Professional Services and other products and services that are
ordered by Licensee from SOLACE.

"Software" means the SOLACE software product(s) described in an Order.

"SOLACE Quotation" means SOLACE's sales quotation document provided by SOLACE
to a prospective customer which sets out the fees for SOLACE's Products.

"Subscription" means the right granted by SOLACE to Licensee to install and
use the Software in accordance with the terms of this Agreement and the
applicable Order, for the Subscription Term specified in the applicable
Order.

"Subscription Fee" means the fee payable by Licensee for a Subscription in
accordance with the terms hereof and the applicable Order.

"Subscription Term" means the period of time that Licensee is authorized by
SOLACE to install and use the Software (including the Documentation).

"Support and Maintenance Services" means the support services provided by
SOLACE for the Software in accordance with the Support and Maintenance Terms.

"Support and Maintenance Terms" means SOLACE'S policies, terms and conditions
for the provision of Support and Maintenance Services to its customers, a
copy of which is available on the SOLACE website at
https://solace.com/support.

"Statement of Work" or "SOW" shall mean a statement of work in the form
attached hereto as Schedule B pursuant to which the parties agree upon the
Professional Services to be provided by SOLACE to Licensee, the fees to be
charged, milestones, deliverables and such other terms and conditions as the
parties may agree upon.

1.2	Currency. Unless otherwise specified, all dollar amounts in this
Agreement, including the symbol "$", refer to United States currency.

2	LICENSE GRANT

2.1	General License to Software.

(a)	Provided Licensee complies with this Agreement, SOLACE hereby grants
to Licensee a non-exclusive, non-sublicensable (except as permitted in
accordance with Section 2.6 below), non-transferable, license, during the
term of this Agreement, to install and use the Software in object code form
during the applicable Subscription Term for the number of Cores specified in
the Order, solely for the Licensee's internal business purposes and in
accordance with the terms of this Agreement.

(b)	If Licensee requires a license from SOLACE to enable Licensee to
bundle or otherwise make available a Product with Licensee's own software,
such bundling will be pursuant to separate terms to be agreed.

2.2	Documentation. Provided Licensee complies with this Agreement,
Licensee may reproduce the Documentation, for use on an internal basis only,
and solely in support of the Licensee's licensed use of the Software.
Distribution of the Documentation outside of Licensee is prohibited without
the express written permission of SOLACE. Licensee must reproduce all
copyright and other proprietary notices that are on the original copy of the
Documentation.

2.3	Back-up Copy.  In addition to the number of copies of the Software
installed and used pursuant to Section 2.1 and paid for in accordance with
Section 5, Licensee may make one copy of each licensed Product per
Subscription solely for back-up purposes, provided that Licensee reproduces
all copyright and other proprietary notices that are on the original copy of
the Software and such back-up copy is not installed or used other than for
back-up and recovery purposes.	Back-up copies that are used as part of a
live or 'hot' back-up will be subject to additional fees.

2.4	Use Restrictions.  Licensee will not: (a) reverse engineer,
disassemble, decompile, or translate the Software (other than Sample
Applications), or otherwise attempt to derive the source code version of the
Software, except if and only to the extent expressly permitted by applicable
law, and provided that Licensee first approaches SOLACE and seeks permission
in writing; (b) except as expressly permitted in this Agreement, use the
Products or any portion thereof to create or develop any developer tools
(including plug-ins and middleware) or any software; (c) except as expressly
permitted in this Agreement, rent, lease, loan or otherwise in any manner
provide, transfer or distribute the Products or any part thereof to any third
party; (d) use the Software in violation of applicable laws; (e) circumvent
any user limits or other license timing or use restrictions that are built
into the Software; (f) except as expressly permitted in this Agreement,
reproduce, distribute, publicly perform, publicly display or create
adaptations or derivative works of or based on the Products.

2.5	Publicly Available Software. Portions of the Software include
software programs that are distributed by SOLACE pursuant to the terms and
conditions of a license granted by the copyright owner of such software
programs and which governs Customer's use of such software programs
("Publicly Available Software").  The Licensee's use of Publicly Available
Software in conjunction with the Software in a manner consistent with the
terms of this Agreement is permitted, however, the Licensee may have broader
rights under the applicable license for Publicly Available Software and
nothing contained herein is intended to impose restrictions or limitations on
the Licensee's use of the Publicly Available Software. The warranty,
indemnity and limitation of liability provisions in this Agreement will apply
to all of the Software, including Publicly Available Software included in the
Software. Copies of such Publicly Available Software license agreements are
available by contacting Licensor at support@solace.com. The source code for
certain portions of the Publicly Available Software included in the Software
(as specified in the copyright notices) is available by contacting SOLACE at
support@solcae.com within a three (3) year period from the original date of
receipt of the applicable Software or Adapter and for a fee that shall not
exceed Licensor' costs associated with the shipping of such software source
code.

2.6	Sub-licensing.	Any sub-licensing of the Software under this
Agreement must be expressly authorized by SOLACE pursuant to an Order or
otherwise in writing. Any attempt by Licensee to sub-license or otherwise
transfer the Products to a third party in breach of this restriction will be
void.  Any sub-licensing that may be permitted under this Agreement by SOLACE
will be subject to such sub-licensee agreeing to substantially similar
restrictions and obligations set out in this Agreement.  Licensee will be
fully liable for any breach by a sub-licensee of any restriction or
obligation, and SOLACE may bring a Claim against Licensee if SOLACE suffers
any Losses arising from such breach.

2.7	Evaluation Licenses.

(a)	If the Software provided to Licensee under this Agreement is
designated by SOLACE in an Order or otherwise as an evaluation release
(indicated by terms such as "pre-commercial", "alpha," "beta," "trial,"
"draft," "early access," "EA" or "evaluation") (each an "Evaluation Software
Release"), Licensee will have the limited right under this Agreement to
download and install the Software on the number of Cores identified in the
Order or, if not identified, one Core, for the Licensee's internal and
non-commercial evaluation of the Software.

(b)	Licensee acknowledges that the Evaluation Software Release may not
meet performance and compatibility standards of a production version. The
Evaluation Software Release may not operate correctly, may be substantially
modified by SOLACE prior to first commercial shipment, and may be withdrawn
completely and never issued for commercial use.

(c)	If Licensee desires other rights for the Evaluation Software Release,
Licensee must request from SOLACE a commercial release of the Software.

(d)	The limited use license granted in subsection (a) will automatically
expire on the earlier of: (i) the date when the Software is made available to
Licensee as a commercially available product, and (ii) the date specified in
the Order or, if no such date is identified in the Order, the date that is 30
days after the date of delivery or provision of the Evaluation Software
Release to Licensee. Following license expiry Licensee will permanently
delete or otherwise purge such Evaluation Software Release from Licensee's
systems and, if requested by SOLACE, certify the same.

2.8	License of APIs. Provided Licensee complies with this Agreement and
any terms that SOLACE provides, SOLACE grants to Licensee a non-exclusive,
royalty free license, during the term of this Agreement, to download, install
and use, the applicable application programming interfaces that may be made
available by SOLACE with the Software ("APIs") solely to create interfaces
between the Software and the Licensee's software or third party software on
Licensee's systems.

2.9	License to Sample Applications.

(a)	SOLACE may, in its sole discretion, provide certain sample Software
in source code or object code form for the purposes of demonstrating certain
features enabled by the Software, including demonstrating to Licensees how to
build applications using APIs, and for use by Licensees with such APIs (each,
a "Sample Application").

(b)	Whether provided separately or together with other Software, if
SOLACE provides such Sample Application to Licensee, then SOLACE hereby
grants to Licensee a non-sublicensable, non-transferable, non-exclusive,
revocable license, to install such Sample Application for Licensee's
evaluation for the same duration as the Software with which the Sample
Application is associated or such other duration as specified by SOLACE upon
delivery of the Sample Application.

3	OPTIONAL SERVICES AND SUPPORT

3.1	Optional Services.  Licensee acknowledge that certain optional
services, such as training, integration and development services may be
provided by SOLACE in association with the Products, and access to such
services will be provided only pursuant to a Statement of Work executed by
SOLACE and Licensee and may include separate and additional fees.

3.2	Support.

(a)	Provided Licensee complies with this Agreement, SOLACE will provide
Support and Maintenance Services the Software in accordance with SOLACE's
then standard Support and Maintenance Terms.  The level of support will be
dependent on whether Licensee has procured either the 'Premium Support Plan'
or 'Standard Support Plan' defined in SOLACE's Support and Maintenance Terms
and as specified in the applicable Order.

(b)	SOLACE may enhance such standard Support and Maintenance Services
from time to time in its discretion.

(c)	For greater clarity, SOLACE's then standard Support and Maintenance
Terms do not apply to Evaluation Software Releases, Sample Applications or
any free versions of the Software that may be made available.  SOLACE may
make available support related information on a free basis for such Software
on its publicly accessible website or otherwise, and such support related
information will, for greater clarity, be subject to the limitations and
exclusions in this Agreement.

4	PROPRIETARY RIGHTS

4.1	Intellectual Property Rights.  In this Agreement "Intellectual
Property Rights" means: (a) any and all proprietary rights anywhere in the
world provided under: (i) patent law; (ii) copyright law (including moral
rights); (iii) trademark law; (iv) design patent or industrial design law; or
(v) any other statutory provision or common law principle applicable to this
Agreement, including trade secret law, that may provide a right in either
hardware or information generally or the expression or use of such hardware
or information; (b) any and all applications, registrations, licenses,
sub-licenses, franchises, agreements or any other evidence of a right in any
of the foregoing.  Except for the licenses expressly granted herein, othing
in this Agreement or the provision of the Products conveys or otherwise
provides to Licensee title, interest or any Intellectual Property Rights in
or to: (a) the Products, or (b) know-how, ideas, or any other subject matter
protectable under laws applicable to Intellectual Property Rights of any
jurisdiction. As between Licensee and SOLACE, SOLACE and its affiliates and
licensors are the sole and exclusive owners of the Products, including
Intellectual Property Rights therein.

4.2	Feedback.  Licensee is encouraged to provide to SOLACE suggestions,
comments and feedback related to the Products (including reporting bugs) (the
"Feedback"). Licensee hereby grants to SOLACE a license to use, copy,
distribute, modify or otherwise adapt, incorporate into any software and
documentation, including the Products, and sublicense, without attribution or
compensation to Licensee, all Feedback which SOLACE receives or otherwise
obtains from Licensee, in any form, to improve, enhance or modify the
Products or otherwise.	Licensee waives or will cause all moral rights to be
waived in any Feedback.

4.3	Third Party Licenses.  The Software may contain or require third
party software that is licensed under third party terms.  SOLACE may direct
Licensee to such third party terms, and in some instances the Software cannot
be used or further distributed without Licensee's acceptance of such terms.
Any failure of Licensee to agree to the terms applicable to such third party
software may undermine certain functionality of or prevent Licensee from
using the Software.

4.4	Open Source Software.

(a)	Licensee will not represent to third parties, or use any third party
software or code in conjunction with: (i) the Software; or (ii) any software,
products, documentation, content or other materials developed using the
Software, in such a way that: (A) creates, purports to create or has the
potential to create, obligations for SOLACE with respect to the Software; or
(B) grants, purports to grant, or has the potential to grant to any third
party any rights to or immunities under any Intellectual Property Rights of
SOLACE, as such rights exist in or relate to the Products.

(b)	Licensee will not use any Software in any manner, including through
incorporation, linking, distribution or otherwise, that will cause any
Products and any Intellectual Property Rights therein to become subject to
any encumbrance or terms and conditions of any third party or open source
license, including any open source license listed on
http://www.opensource.org/licenses/alphabetical (each an "Open Source
License").

(c)	The restrictions, limitations, exclusions and conditions referred to
under subsection (b) will apply even if SOLACE becomes aware of or fails to
act in a manner to address any violation or failure to comply therewith.  No
act by SOLACE that is undertaken under this Agreement in respect to any
Products will be construed as intending to cause any Intellectual Property
Rights that are owned or controlled by SOLACE or any of its affiliates (or
for which SOLACE or any of its affiliates has received license rights) to
become subject to any encumbrance or terms and conditions of any Open Source
License.

4.5	Use of Name and Logo.  Licensee will not display or make any use of
SOLACE's or its affiliates' names, marks or logos without the prior written
approval of SOLACE.

5	FEES AND TAXES

5.1	Fees. Licensee shall pay the applicable Subscription Fees and support
fees specified in the applicable Order. Except as otherwise specified herein
or in an Order,  Subscription Fees are based on Subscriptions purchased and
not actual usage. Payment obligations are non-cancellable, Subscription Fees
paid are non-refundable, and the number of Subscriptions purchased cannot be
decreased during the relevant Subscription Term.

5.2	Invoices and Payment.  Subscription Fees will be invoiced in advance
and otherwise in accordance with the relevant Order.  All invoices issued by
SOLACE are due and payable within 30 days of the invoice date unless
otherwise agreed in an Order.  Licensee will be responsible for any and all
sales, use, excise, import, value-added, services, consumption, and other
taxes assessed on the receipt of the Products, and any related services as a
whole.

5.3	Overdue Charges. Any payment not received from Customer by the due
date may accrue (except with respect to charges then subject to a reasonable
and good faith dispute), at Licensor' discretion, late charges at the rate of
1.5% of the outstanding balance per month (19.57% per annum), or the maximum
rate permitted by law, whichever is lower, from the date such payment was due
until the date paid.

6	CONFIDENTIALITY

6.1	Definition of Confidential Information.

In this Agreement "Confidential Information" of a Party means any information
of a Party (including in respect to SOLACE any of its affiliates, licensors,
customers, employees or subcontractors) (the "Disclosing Party"), whether
oral, written or in electronic form, which has or will come into the
possession or knowledge of the other Party (the "Receiving Party") in
connection with or as a result of entering into this Agreement that can
reasonably be considered to be confidential in the circumstances of
disclosure or which is designated as confidential. The Products, any
performance information, service levels, support terms, and results of
testing of the Software, and the terms of this Agreement are Confidential
Information of SOLACE. Notwithstanding the foregoing, "Confidential
Information" does not include information that is:

(a)	publicly available when it is received by or becomes known to the
Receiving Party or that subsequently becomes publicly available other than
through a direct or indirect act or omission of the Receiving Party (but only
after it becomes publicly available);

(b)	established by evidence to have been already known to the Receiving
Party at the time of its disclosure to the Receiving Party and is not known
by the Receiving Party to be the subject of an obligation of confidence of
any kind;

(c)	independently developed by the Receiving Party without any use of or
reference to the Confidential Information of the Disclosing Party as
established by evidence that would be acceptable to a court of competent
jurisdiction;

(d)	received by the Receiving Party in good faith without an obligation
of confidence of any kind from a third party who the Receiving Party had no
reason to believe was not lawfully in possession of such information free of
any obligation of confidence of any kind, but only until the Receiving Party
subsequently comes to have reason to believe that such information was
subject to an obligation of confidence of any kind when originally received;
or

(e)	Feedback provided by Licensee or a representative of Licensee.

6.2	Confidentiality Obligations.

(a)	Each Party will, in its capacity as a Receiving Party: (i) not use or
reproduce Confidential Information of the Disclosing Party for any purpose,
other than as may be reasonably necessary for the exercise of its rights or
the performance of its obligations set out in this Agreement; and (ii) not
disclose, provide access to, transfer or otherwise make available any
Confidential Information of the Disclosing Party to any third party except as
expressly permitted in this Agreement.

(b)	Each Party may, in its capacity as a Receiving Party, disclose
Confidential Information of the Disclosing Party: (i) if and to the extent
required by a governmental authority or otherwise as required by applicable
law, provided that the Receiving Party must first give the Disclosing Party
notice of such compelled disclosure (except where prohibited by applicable
law from doing so) and must use commercially reasonable efforts to provide
the Disclosing Party with an opportunity to take such steps as it desires to
challenge or contest such disclosure or seek a protective order.  Thereafter,
the Receiving Party may disclose the Confidential Information of the
Disclosing Party, but only to the extent required by applicable law and
subject to any protective order that applies to such disclosure; and (ii) to:
(A) its accountants, internal and external auditors and other professional
advisors if and to the extent that such persons need to know such
Confidential Information in order to provide the applicable professional
advisory services relating to the Receiving Party; and (B) employees of the
Receiving Party and its subcontractors if and to the extent that such persons
need to know such Confidential Information to perform their respective
obligations under this Agreement;

provided that any such person is aware of the provisions of this Section 6.2
and has entered into a written agreement with the Receiving Party that
includes confidentiality obligations in respect of such Confidential
Information of the Disclosing Party that are no less stringent than those
contained in this Section 6.2.

6.3	Consent to Injunctive Relief.  Any unauthorized use or disclosure of
the Confidential Information of SOLACE, its affiliates or licensors may cause
irreparable harm and significant injury to SOLACE that would be difficult to
ascertain or quantify; accordingly Licensee agrees that SOLACE will have the
right to seek and obtain injunctive or other equitable relief to enforce the
terms of this Agreement and without limiting any other rights or remedies
that SOLACE may have.

7	WARRANTY AND DISCLAIMER OF WARRANTIES.

7.1	Warranty.  SOLACE warrants that the Software will materially comply
with the Documentation during the Subscription Term.  If the Software does
not materially conform with the warranty in the prior sentence, provided that
Licensee is in compliance with the terms of this Agreement, and all
Subscription Fees are fully-paid up, SOLACE will provide the support to
Licensee in respect to the applicable Software to the extent set out in
SOLACE's then current Support and Maintenance Terms, and the provision of
support to correct the non-compliance with the warranty in this Section will
be Licensee's sole and exclusive remedy in the event of non-compliance with
the warranty in this Section by SOLACE.  All other support will be dependent
on the plan procured by Licensee, as defined in the Support and Maintenance
Terms.

7.2	Disclaimers.

a)	EXCEPT AS SET OUT IN SECTION 7.1, THE PRODUCTS AND  SUPPORT THAT MAY
BE PROVIDED BY SOLACE UNDER THIS AGREEMENT, IS PROVIDED 'AS-IS' AND 'AS
AVAILABLE'.

(b)	Except as set out in Section 7.1, the Products and  support are
without any additional warranties of any kind, whether express, implied,
collateral, statutory or otherwise. SOLACE does not warrant or make any
representations regarding the use, or the results of the use, of the Products
in terms of its correctness, accuracy, reliability, or otherwise.

(c)	SOLACE does not represent or warrant that the functionality of the
Products will meet Licensee requirements, or that the operation of the
Products will be uninterrupted or error-free, or that the Products  or any
service enabled by the use of the Software will always be available, or that
defects in the Products  will be corrected.

(d)	TO THE MAXIMUM EXTENT PERMITTED UNDER APPLICABLE LAW, SOLACE ON ITS
OWN BEHALF AND ON BEHALF OF ITS AFFILIATES AND LICENSOR(S) EXPRESSLY
DISCLAIMS ALL WARRANTIES AND CONDITIONS, WHETHER EXPRESS, IMPLIED, STATUTORY
OR OTHERWISE, INCLUDING ANY IMPLIED WARRANTIES, AND CONDITIONS OF
MERCHANTABLE QUALITY, MERCHANTABILITY, QUALITY, FITNESS FOR A PARTICULAR
PURPOSE, AND NON-INFRINGEMENT.

(e)	Some jurisdictions do not allow the exclusion of implied warranties,
so exclusions in this Article 7 will apply only to the extent permitted by
applicable law.

8	LICENSEE INDEMNITY AND EXCLUSION.

8.1	Licensee Indemnity.

(a)	Without limiting SOLACE's rights and remedies under this Agreement,
Licensee will indemnify, defend and hold SOLACE, its licensors, affiliates or
any of their respective directors, officers, employees or agents (together,
the "Solace Indemnitees") harmless from and against any and all third party
Claims and Losses incurred or otherwise suffered by each SOLACE Indemnitee
arising out of, resulting from or related to:
(i)	any use, reproduction or distribution of the Products
(notwithstanding the restrictions and obligations in this Agreement), as
modified or integrated by Licensee in Licensee application, which causes an
infringement or misappropriation of any Intellectual Property Right,
publicity or privacy right of any third parties arising in any jurisdiction
anywhere in the world, except and solely to the extent such infringement is
caused by the unmodified Software, or portions thereof, as supplied to
Licensee by SOLACE under this Agreement; or
(ii)	any use, downloading, distribution, installation, storage, execution,
or transfer of the Products in breach of this Agreement.

(b)	SOLACE may enforce the indemnity under this Article 8 on behalf of
any or all of the SOLACE Indemnitees.  Licensee may only bring a Claim
against SOLACE and not any SOLACE Indemnitees under this Agreement.

8.2	SOLACE Indemnity.

(a)	SOLACE will defend Licensee from and against any and all Claims by a
third party incurred or otherwise suffered by Licensee arising out of,
resulting from or related to a Claim that the Products licensed pursuant to
Section 2.1 infringe or misappropriate third party copyright or patent rights
in Canada or the United States of America, and indemnity Licensee from any
damages awarded by a court of final determination.

(b)	Without limitation, Section 8.2 will not be applicable and SOLACE
will not be liable to defend a Claim to the extent that such Claim is based
on: (i) Licensee's use of the Products after SOLACE notifies Licensee to
discontinue using them; (ii) Licensee combining the Products with non-SOLACE
services, products, programs or data; or (iii) Licensee altering or modifying
the Products.

(c)	If SOLACE receives information concerning an infringement or
misappropriation Claim related to the Products, SOLACE may, at its expense
and without obligation to do so, either: (i) procure the Intellectual
Property Rights or other right(s) to continue to use the Product; or (ii)
replace or modify the Product to make it non-infringing; or (iii) immediately
terminate this Agreement on written notice to Licensee, in which case SOLACE
will refund to Licensee, on a pro-rata basis, any pre-paid fees in respect to
such Product from the date of such termination to the end of the then current
Subscription Term for such Product; and this Section 8.2(c) states the sole
and exclusive remedy of Licensee and the entire liability of SOLACE for third
party infringement claims and actions.

8.3	Indemnification Procedures.  Each Party's obligations under this
Article 8 are contingent on all of the following: (i) the Party seeking the
indemnity (the "Indemnified Party") must notify the other Party (the
"Indemnifying Party"), in a timely manner and in writing of the Claim; (ii)
the Indemnified Party must give the Indemnifying Party sole control over
defense and settlement of the Claim; (iii) the Indemnified Party must provide
the Indemnifying Party with reasonable information and assistance, at the
Indemnifying Party's request, as needed in defending the Claim (the
Indemnifying Party will reimburse the Indemnified Party for reasonable
expenses that the Indemnified Party incurs in providing that assistance).
The Indemnified Party may choose to have its counsel, monitor or participate
in the defense of such a Claim provided that the Indemnified Party will be
responsible for the cost of its own counsel and the Indemnifying Party's
obligations in this Article 8 do not extend to the Indemnified Party's legal
costs should it wish to exercise such right. The Indemnifying Party will not
be responsible for any settlement made by the Indemnified Party without its
prior written consent.	The Indemnifying Party may not settle or publicize
any Claim without the Indemnified Party's prior written consent.

9	LIMITATIONS OF LIABILITY.

9.1	Definition and Limitations of Liability.

(a)	In this Agreement: "Claim" means any actual, threatened or potential
civil, criminal, administrative, regulatory, arbitral or investigative
demand, allegation, action, suit, investigation or proceeding or any other
claim or demand; and "Losses" means any and all damages, fines, penalties,
deficiencies, losses, liabilities (including settlements and judgments),
costs and expenses (including interest, court costs, reasonable fees and
expenses of lawyers, accountants and other experts and professionals or other
reasonable fees and expenses of litigation or other proceedings or of any
Claim, default or assessment).

(b)	SUBJECT TO SECTION 9.1(d), TO THE FULLEST EXTENT PERMITTED BY
APPLICABLE LAW, UNDER NO CIRCUMSTANCES WILL SOLACE INDEMNITEES BE LIABLE FOR
(A) ANY INDIRECT, INCIDENTAL, SPECIAL, PUNITIVE, EXEMPLARY OR CONSEQUENTIAL
DAMAGES; OR (B) ANY DAMAGES FOR LOSS OF BUSINESS PROFITS, BUSINESS
INTERRUPTION, OR LOSS OF BUSINESS INFORMATION, IN EACH CASE, ARISING OUT OF
OR IN CONNECTION WITH THIS AGREEMENT, INCLUDING ANY DOWNLOAD, INSTALLATION OR
USE OF, OR INABILITY TO USE, THE PRODUCTS; EVEN IF SUCH DAMAGES WERE
FORESEEABLE, AND REGARDLESS OF WHETHER THE SOLACE INDEMNITIEES HAVE BEEN
ADVISED OF THE POSSIBILITY OF SUCH DAMAGES.

(c)	SUBJECT TO SECTION 9.1(d), TO THE FULLEST EXTENT PERMITTED BY
APPLICABLE LAW, IN NO EVENT WILL SOLACE INDEMNITEES' TOTAL AGGREGATE
LIABILITY IN RESPECT OF THIS AGREEMENT, INCLUDING THE PRODUCTS AND ANY
SERVICES THAT MAY BE PROVIDED HEREUNDER, FOR ANY AND ALL LOSSES AND CLAIMS
EXCEED THE AMOUNTS PAID TO SOLACE IN THE 12 MONTHS IMMEDIATELY PRECEDING THE
EVENT GIVING RISE TO THE CLAIM.

(d)	Certain Damages Not Excluded or Limited.  NOTWITHSTANDING THE
FOREGOING, SECTIONS 9.1 (b) AND (c) DO NOT APPLY TO (I) DAMAGES ARISING FROM
A PARTY'S BREACH OF ITS CONFIDENTIALITY OBLIGATIONS HEREUNDER, (II)
INDEMNIFICATION CLAIMS, (III) DAMAGES ARISING FROM INFRINGEMENT OF A PARTY'S
INTELLECTUAL PROPERTY RIGHTS; (IV) ANY CLAIMS FOR NON-PAYMENT, (V) FRAUD OR
WILLFUL MISCONDUCT, OR (VI) BODILY INJURY OR DEATH.

(e)	This Article 9 will apply irrespective of the nature of the cause of
action, demand or Claim, including, breach of contract (including fundamental
breach), negligence (including gross negligence), tort or any other legal
theory, and will survive a fundamental breach or breaches of this Agreement
or of any remedy contained herein.

10	TERM AND TERMINATION.

10.1	Term and Renewal. This Agreement will be effective from the Effective
Date and will continue until the expiry of the Subscription Term set out in
the Order or the Agreement terminates in accordance with its terms. Subject
to payment of the applicable Software Fees, Software Subscriptions shall
automatically renew for additional periods equal to the expiring Subscription
Term or one (1) year (whichever is shorter), unless either party gives the
other notice of non-renewal at least thirty (30) days prior to the end of the
then-current Subscription Term. The Subscription Fees during any automatic
renewal term will be as set forth in the applicable Order.

10.2	Termination for Cause. A party may terminate this Agreement for cause
(i) upon 30 days' written notice to the other party of a material breach if
such breach remains uncured at the expiration of such period, or (ii) if the
other party becomes the subject of a petition in bankruptcy or any other
proceeding relating to insolvency, receivership, liquidation or assignment
for the benefit of creditors.

10.3	Termination by SOLACE. SOLACE may terminate this Agreement for cause
with immediate effect on written notice if Licensee commits a breach of
Articles 4 or 5 by Licensee.

10.4	Termination of Sample Application and Evaluation Software Release
Licenses for Convenience by SOLACE. SOLACE may terminate the licenses in
respect to the Sample Applications, Evaluation Software Releases, and any
other Products that may be licensed by SOLACE on a trial basis, at any time
for convenience, upon written notice to Licensee.

10.5	Termination of Licenses of Trial Software. Subject to Section 10.4,
if any Software is licensed for use by a Licensee on a trial basis, the
license to use such Software during a trial period will continue for such
duration set out in an Order.

10.6	Effects of Termination. Upon termination or expiry of this Agreement
or specific licenses granted hereunder for any reason, and without limiting
SOLACE's other rights or remedies under this Agreement: (a) Licensee must
permanently delete or destroy, or otherwise purge, all copies (electronic or
otherwise) of the applicable Products from Licensee's systems, and any other
Confidential Information of SOLACE, in Licensee's possession or control, and,
if requested by SOLACE, certify the same, and the license and other rights
granted to Licensee in this Agreement will terminate; (b) termination or
expiration of this Agreement or an individual Subscription will result in
termination of any applicable Support and Maintenance Services; and (c)
Licensee will not receive a return of any pre-paid fees in respect to the
applicable Products, on a pro-rata basis or otherwise, except where expressly
stated in this Agreement.

10.7	Survival. Neither the expiration nor the earlier termination of this
Agreement will release either of the Parties from any obligation or liability
that accrued prior to such expiration or termination.  The provisions of this
Agreement requiring performance or fulfilment after the expiration or earlier
termination of this Agreement, including Articles 4, 5, 5, 7, 9, 8, 9, 10,
11, 12, and 13, and such other provisions as are necessary for the
interpretation thereof and any other provisions hereof, the nature and intent
of which is to survive termination or expiration of this Agreement, will
survive the expiration or earlier termination of this Agreement.

11	AUDIT AND REMEDIATION

11.1	Audit. During the term of this Agreement and for two years
thereafter, SOLACE or any internal or external audit representative acting on
behalf of SOLACE (the "SOLACE Audit Representatives") will have the right,
and Licensee will provide access to SOLACE Audit Representatives during
regular business hours and upon reasonable prior written notice to Licensee,
to audit and inspect on a mutually agreed upon date and location any system
or facility or part of a system or facility to which Licensee has downloaded
the Software or is receiving any services (or both) in order to verify the
performance by Licensee of its obligations under this Agreement, including
the Licensee's usage of the Products in accordance with the restrictions and
terms in this Agreement.

11.2	Remediation. Without limiting SOLACE's rights and remedies under this
Agreement, if an audit conducted pursuant to this Agreement reveals any
error, deficiency or other failure to perform on the part of Licensee
including use of the Software contrary to the licenses in this Agreement or
installed on systems, computers or processors for which the Licensee has not
paid applicable Subscription Fees: (a) Licensee will immediately pay to
SOLACE any fees due and payable for Software used in breach of the
restrictions in this Agreement, plus interest at the lesser of: (i) the rate
of 1.5 percent per month compounded monthly (19.562 percent per annum); or
(ii) the maximum rate allowed by applicable law, in each case, on the amount
outstanding from the date when payment is due until the date payment in full
is received by SOLACE; and (b) pursue any other right or remedy SOLACE may
have under this Agreement.

12	EXPORT COMPLIANCE ASSURANCES

(a)	All Products obtained from SOLACE are subject to the export control
and economic sanctions laws and regulations of Canada, including the Exports
and Import Permits Act,  R.S.C. 1985, c. E-19, Area Control List, Export
Control List, and the United States, including the Export Administration
Regulations ("EAR", 15 CFR 730 et seq., http://www.bis.doc.gov/) administered
by the Department of Commerce, Bureau of Industry and Security, and the
Foreign Asset Control Regulations (31 CFR 500 et seq.,
http://www.treas.gov/offices/enforcement/ofac/) administered by the
Department of Treasury, Office of Foreign Assets Control ("OFAC"), each as
may be amended and updated from time to time.

(b)	Licensee will not, and will ensure that Licensee will not directly or
indirectly export, re-export, transfer or release (collectively, "export")
any Products to any destination, person, entity or end use prohibited or
restricted under Canadian or US law, or the laws of the jurisdiction in which
Licensee is resident or in which Licensee uses the Products, without prior
government or regulatory authorization to the extent required by applicable
laws and regulations.

(c)	The US government maintains embargoes and sanctions against the
countries listed in Country Groups E:1/2 of the EAR (Supplement 1 to part
740), including, as at the Effective Date, Cuba, Iran, North Korea, Sudan and
Syria, as amended from time to time.  Licensee will not directly or
indirectly employ any Product received from SOLACE in missile technology,
sensitive nuclear or chemical biological weapons activities, or in any manner
knowingly transfer any Product to any party for any such end use. Licensee
will not export Products listed in Supplement 2 to part 744 of the EAR for
military end-uses, as defined in part 744.21, to the People's Republic of
China. Licensee will not transfer any Product to any party listed on any of
the denied parties lists or specially designated nationals lists maintained
under said regulations without appropriate US government authorization to the
extent required by regulation. Licensee acknowledge that other countries may
have trade laws pertaining to import, use, export or distribution of
Products, and that compliance with same is Licensee responsibility.

(d)	Licensee may not use the Products if Licensee is barred from
receiving the Products under the laws of Canada, the United States or any
other country including the country in which Licensee are resident or in
which Licensee use the Products.

13	GENERAL

13.1	U.S. Government Users.	If Licensee are acting on behalf of an agency
or instrumentality of the U.S. federal government, the Product, as
applicable, are "commercial computer software" and "commercial computer
software documentation" developed exclusively at private expense by SOLACE.
Pursuant to FAR 12.212 or DFARS 227 7202 and their successors, as applicable,
use, reproduction and disclosure of the Products is governed by the terms of
this Agreement.

13.2	Entire Agreement. This Agreement, and the agreements and other
documents required to be delivered pursuant to this Agreement, constitute the
entire and exclusive agreement between SOLACE and Licensee, and sets out all
the covenants, promises, warranties, representations, conditions and
agreements between the Parties in connection with the subject matter of this
Agreement, and supersedes all prior agreements (whether written or oral,
pre-contractual or otherwise) and other communications between SOLACE and
Licensee. There are no covenants, promises, warranties, representations,
conditions or other agreements, whether oral or written, pre-contractual or
otherwise, express, implied or collateral, whether statutory or otherwise,
between the Parties in connection with the subject matter of this Agreement
except as specifically set forth in this Agreement and any document required
to be delivered pursuant to this Agreement.

13.3	Amendments. This Agreement may be modified only by a written
amendment agreed to by both Licensee and SOLACE, except that SOLACE may
modify the Documentation from time to time, provided that SOLACE does not
materially lessen the description of the functionality of the Products as a
result of such modification.

13.4	English Language. This Agreement is entered into solely in the
English language, and if for any reason any other language version is
prepared by any Party, it will be solely for convenience and the English
version will govern and control in all respects.  If Licensee are located in
the province of Quebec, Canada, the following applies:	The Parties hereby
confirm they have requested this Agreement and all related documents be
prepared in English.  Les parties ont exigé que le présent contrat et tous
les documents connexes soient rédigés en anglais.

13.5	Waiver. To be effective, any waiver by a Party of any of its rights
or any other Party's obligations under this Agreement must be made in a
writing signed by the Party to be charged with the waiver.  No failure or
forbearance by any Party to insist upon or enforce performance by any other
Party of any of the provisions of this Agreement or to exercise any rights or
remedies under this Agreement or otherwise at law or in equity will be
construed as a waiver or relinquishment to any extent of such Party's right
to assert or rely upon any such provision, right, or remedy in that or any
other instance; rather, the same will be and remain in full force and effect.
 A Party's waiver of a breach of any term will not be a waiver of any
subsequent breach of the same or another term.

13.6	Cumulative Rights. The rights of each Party hereunder are cumulative
and no exercise or enforcement by a Party of any right or remedy hereunder
will preclude the exercise or enforcement by such Party of any other right or
remedy hereunder or which such Party is otherwise entitled by law to enforce.

13.7	Severability.  If, in any jurisdiction, any provision of this
Agreement or its application to any Party or circumstance is restricted,
prohibited or unenforceable, the provision will, as to that jurisdiction, be
ineffective only to the extent of the restriction, prohibition or
unenforceability without invalidating the remaining provisions of this
Agreement and without affecting the validity or enforceability of such
provision in any other jurisdiction, or without affecting its application to
other Parties or circumstances.

13.8	Assignment.  SOLACE may assign this Agreement or any of the benefits,
rights or obligations under this Agreement without the prior written consent
of the Licensee.  Licensee may not assign this Agreement or any of the
benefits, rights or obligations under this Agreement without the prior
written consent of SOLACE. Any attempt by Licensee to so assign or transfer
is null and void. If SOLACE does consent to an assignment of this Agreement,
the transferee/assignee must be acceptable to SOLACE and agree to the terms
and conditions of this Agreement.

13.9	Further Assurances. The Parties will, with reasonable diligence, do
all things and provide all such reasonable assurances as may be required to
consummate the transactions contemplated by this Agreement, and each Party
will provide such further documents or instruments required by any other
Party as may be reasonably necessary or desirable to effect the purpose of
this Agreement and carry out its provisions.

13.10	Governing Law and Jurisdiction.  This Agreement is governed and
interpreted in accordance with the laws of the Province of Ontario and the
laws of Canada applicable therein, without giving effect to its conflict of
laws provisions. Any Claim arising out of or related to this Agreement must
be brought exclusively in a federal or provincial court located in Ottawa,
Canada, and Licensee hereby consents to the jurisdiction and venue of such
courts.  Each of the Parties irrevocably waives, to the fullest extent it may
effectively do so, the defence of an inconvenient forum to the maintenance of
such action, application or proceeding.  The Parties will not raise any
objection to the venue of any action, application, reference or other
proceeding arising out of or related to this Agreement in the federal or
provincial courts sitting in Ottawa, including the objection that the
proceedings have been brought in an inconvenient forum.  A final judgment in
any such action, application or proceeding is conclusive and may be enforced
in other jurisdictions by suit on the judgment or in any other manner
specified by law.  The United Nations Convention on Contracts for the
International Sale of Goods is expressly disclaimed and will not apply.
```

