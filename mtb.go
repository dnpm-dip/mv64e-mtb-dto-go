package mtb

import "encoding/json"

func UnmarshalMtb(data []byte) (Mtb, error) {
	var r Mtb
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Mtb) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Mtb struct {
	CarePlans                  []MTBCarePlan                     `json:"carePlans,omitempty"`
	ClaimResponses             []ClaimResponse                   `json:"claimResponses,omitempty"`
	Claims                     []ClaimElement                    `json:"claims,omitempty"`
	Consent                    Consent                           `json:"consent"`
	Diagnoses                  []MTBDiagnosis                    `json:"diagnoses,omitempty"`
	EcogStatus                 []PerformanceStatus               `json:"ecogStatus,omitempty"`
	Episode                    *MTBEpisode                       `json:"episode,omitempty"`
	GeneticCounsellingRequests []GeneticCounselingRecommendation `json:"geneticCounsellingRequests,omitempty"`
	HistologyReports           []HistologyReport                 `json:"histologyReports,omitempty"`
	LastGuidelineTherapies     []LastGuidelineTherapyElement     `json:"lastGuidelineTherapies,omitempty"`
	MolecularTherapies         []MolecularTherapy                `json:"molecularTherapies,omitempty"`
	NgsReports                 []SomaticNGSReport                `json:"ngsReports,omitempty"`
	Patient                    MtbPatient                        `json:"patient"`
	PreviousGuidelineTherapies []LastGuidelineTherapyElement     `json:"previousGuidelineTherapies,omitempty"`
	Recommendations            []MTBMedicationRecommendation     `json:"recommendations,omitempty"`
	Responses                  []Response                        `json:"responses,omitempty"`
	Specimens                  []SpecimenElement                 `json:"specimens,omitempty"`
	StudyInclusionRequests     []StudyEnrollmentRecommendation   `json:"studyInclusionRequests,omitempty"`
}

type MTBCarePlan struct {
	Description               *string                `json:"description,omitempty"`
	Diagnosis                 *string                `json:"diagnosis,omitempty"`
	GeneticCounsellingRequest *string                `json:"geneticCounsellingRequest,omitempty"`
	ID                        string                 `json:"id"`
	IssuedOn                  *string                `json:"issuedOn,omitempty"`
	NoTargetFinding           *NoTargetFinding       `json:"noTargetFinding,omitempty"`
	Patient                   NoTargetFindingPatient `json:"patient"`
	Recommendations           []string               `json:"recommendations,omitempty"`
	StudyInclusionRequests    []string               `json:"studyInclusionRequests,omitempty"`
}

type NoTargetFinding struct {
	Diagnosis string                 `json:"diagnosis"`
	IssuedOn  *string                `json:"issuedOn,omitempty"`
	Patient   NoTargetFindingPatient `json:"patient"`
}

type NoTargetFindingPatient struct {
	ID   string      `json:"id"`
	Type PatientType `json:"type"`
}

type ClaimResponse struct {
	Claim    ClaimResponseClaim         `json:"claim"`
	ID       string                     `json:"id"`
	IssuedOn string                     `json:"issuedOn"`
	Patient  NoTargetFindingPatient     `json:"patient"`
	Reason   *ClaimResponseStatusReason `json:"reason,omitempty"`
	Status   CodingClaimResponseStatus  `json:"status"`
}

type ClaimResponseClaim struct {
	ID   *string                 `json:"id,omitempty"`
	Type *ClaimResponseClaimType `json:"type,omitempty"`
}

type CodingClaimResponseStatus struct {
	Code    ClaimResponseStatus `json:"code"`
	Display *string             `json:"display,omitempty"`
	System  *string             `json:"system,omitempty"`
}

type ClaimElement struct {
	ID       string                 `json:"id"`
	IssuedOn string                 `json:"issuedOn"`
	Patient  NoTargetFindingPatient `json:"patient"`
	Therapy  *string                `json:"therapy,omitempty"`
}

type Consent struct {
	ID      *string                 `json:"id,omitempty"`
	Patient *NoTargetFindingPatient `json:"patient,omitempty"`
	Status  *ConsentStatus          `json:"status,omitempty"`
}

type MTBDiagnosis struct {
	Code                     Coding                 `json:"code"`
	GuidelineTreatmentStatus *Coding                `json:"guidelineTreatmentStatus,omitempty"`
	HistologyResults         []string               `json:"histologyResults,omitempty"`
	IcdO3T                   *Coding                `json:"icdO3T,omitempty"`
	ID                       string                 `json:"id"`
	Patient                  NoTargetFindingPatient `json:"patient"`
	RecordedOn               *string                `json:"recordedOn,omitempty"`
	StatusHistory            []StatusHistory        `json:"statusHistory,omitempty"`
	WhoGrade                 *Coding                `json:"whoGrade,omitempty"`
}

type Coding struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	System  *string `json:"system,omitempty"`
	Version *string `json:"version,omitempty"`
}

type StatusHistory struct {
	Date   string                  `json:"date"`
	Status MTBDiagnosisTumorSpread `json:"status"`
}

type PerformanceStatus struct {
	EffectiveDate string                 `json:"effectiveDate"`
	ID            string                 `json:"id"`
	Patient       NoTargetFindingPatient `json:"patient"`
	Value         CodingECOG             `json:"value"`
}

type CodingECOG struct {
	Code    PurpleCode `json:"code"`
	Display *string    `json:"display,omitempty"`
	System  *string    `json:"system,omitempty"`
	Version *string    `json:"version,omitempty"`
}

type MTBEpisode struct {
	ID      string                 `json:"id"`
	Patient NoTargetFindingPatient `json:"patient"`
	Period  PeriodLocalDate        `json:"period"`
}

type PeriodLocalDate struct {
	End   *string `json:"end,omitempty"`
	Start string  `json:"start"`
}

type GeneticCounselingRecommendation struct {
	ID       string                 `json:"id"`
	IssuedOn *string                `json:"issuedOn,omitempty"`
	Patient  NoTargetFindingPatient `json:"patient"`
	Reason   string                 `json:"reason"`
}

type HistologyReport struct {
	ID               string                  `json:"id"`
	IssuedOn         string                  `json:"issuedOn"`
	Patient          NoTargetFindingPatient  `json:"patient"`
	Specimen         HistologyReportSpecimen `json:"specimen"`
	TumorCellContent *TumorCellContent       `json:"tumorCellContent,omitempty"`
	TumorMorphology  *TumorMorphology        `json:"tumorMorphology,omitempty"`
}

type HistologyReportSpecimen struct {
	ID   string       `json:"id"`
	Type SpecimenType `json:"type"`
}

type TumorCellContent struct {
	ID       string                 `json:"id"`
	Method   TumorCellContentMethod `json:"method"`
	Specimen string                 `json:"specimen"`
	Value    float64                `json:"value"`
}

type TumorMorphology struct {
	ID       string                 `json:"id"`
	Note     *string                `json:"note,omitempty"`
	Patient  NoTargetFindingPatient `json:"patient"`
	Specimen string                 `json:"specimen"`
	Value    Coding                 `json:"value"`
}

type LastGuidelineTherapyElement struct {
	BasedOn       *string                    `json:"basedOn,omitempty"`
	Diagnosis     *string                    `json:"diagnosis,omitempty"`
	ID            string                     `json:"id"`
	Medication    []Coding                   `json:"medication,omitempty"`
	NotDoneReason *CodingTherapyStatusReason `json:"notDoneReason,omitempty"`
	Note          *string                    `json:"note,omitempty"`
	Patient       NoTargetFindingPatient     `json:"patient"`
	Period        *PeriodLocalDate           `json:"period,omitempty"`
	ReasonStopped *CodingTherapyStatusReason `json:"reasonStopped,omitempty"`
	RecordedOn    *string                    `json:"recordedOn,omitempty"`
	Status        *TherapyStatus             `json:"status,omitempty"`
	TherapyLine   *int64                     `json:"therapyLine,omitempty"`
}

type CodingTherapyStatusReason struct {
	Code    NotDoneReasonCode `json:"code"`
	Display *string           `json:"display,omitempty"`
	System  *string           `json:"system,omitempty"`
	Version *string           `json:"version,omitempty"`
}

type MolecularTherapy struct {
	History []LastGuidelineTherapyElement `json:"history"`
}

type SomaticNGSReport struct {
	Brcaness           *float64               `json:"brcaness,omitempty"`
	CopyNumberVariants []Cnv                  `json:"copyNumberVariants,omitempty"`
	DnaFusions         []DNAFusion            `json:"dnaFusions,omitempty"`
	ID                 string                 `json:"id"`
	IssueDate          *string                `json:"issueDate,omitempty"`
	Metadata           []Metadatum            `json:"metadata"`
	MSI                *float64               `json:"msi,omitempty"`
	Patient            NoTargetFindingPatient `json:"patient"`
	RnaFusions         []RNAFusion            `json:"rnaFusions,omitempty"`
	RnaSeqs            []RNASeq               `json:"rnaSeqs,omitempty"`
	SequencingType     Coding                 `json:"sequencingType"`
	SimpleVariants     []Snv                  `json:"simpleVariants,omitempty"`
	Specimen           NgsReportSpecimen      `json:"specimen"`
	Tmb                *float64               `json:"tmb,omitempty"`
	TumorCellContent   *TumorCellContent      `json:"tumorCellContent,omitempty"`
}

type Cnv struct {
	Chromosome            Chromosome             `json:"chromosome"`
	CNA                   *float64               `json:"cnA,omitempty"`
	CNB                   *float64               `json:"cnB,omitempty"`
	CopyNumberNeutralLoH  []CopyNumberNeutralLoH `json:"copyNumberNeutralLoH,omitempty"`
	EndRange              EndRange               `json:"endRange"`
	ID                    string                 `json:"id"`
	RelativeCopyNumber    *float64               `json:"relativeCopyNumber,omitempty"`
	ReportedAffectedGenes []ReportedAffectedGene `json:"reportedAffectedGenes,omitempty"`
	ReportedFocality      *string                `json:"reportedFocality,omitempty"`
	StartRange            StartRange             `json:"startRange"`
	TotalCopyNumber       *int64                 `json:"totalCopyNumber,omitempty"`
	Type                  CNVType                `json:"type"`
}

type CopyNumberNeutralLoH struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
}

type EndRange struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type ReportedAffectedGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
}

type StartRange struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type DNAFusion struct {
	FusionPartner3Prime DnaFusionFusionPartner3Prime `json:"fusionPartner3prime"`
	FusionPartner5Prime DnaFusionFusionPartner5Prime `json:"fusionPartner5prime"`
	ID                  string                       `json:"id"`
	ReportedNumReads    int64                        `json:"reportedNumReads"`
}

type DnaFusionFusionPartner3Prime struct {
	Chromosome Chromosome `json:"chromosome"`
	Gene       PurpleGene `json:"gene"`
	Position   float64    `json:"position"`
}

type PurpleGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
}

type DnaFusionFusionPartner5Prime struct {
	Chromosome Chromosome `json:"chromosome"`
	Gene       FluffyGene `json:"gene"`
	Position   float64    `json:"position"`
}

type FluffyGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
}

type Metadatum struct {
	KitManufacturer string `json:"kitManufacturer"`
	KitType         string `json:"kitType"`
	Pipeline        string `json:"pipeline"`
	ReferenceGenome string `json:"referenceGenome"`
	Sequencer       string `json:"sequencer"`
}

type RNAFusion struct {
	CosmicID            *string                      `json:"cosmicId,omitempty"`
	Effect              *string                      `json:"effect,omitempty"`
	FusionPartner3Prime RnaFusionFusionPartner3Prime `json:"fusionPartner3prime"`
	FusionPartner5Prime RnaFusionFusionPartner5Prime `json:"fusionPartner5prime"`
	ID                  string                       `json:"id"`
	ReportedNumReads    int64                        `json:"reportedNumReads"`
}

type RnaFusionFusionPartner3Prime struct {
	Exon         string        `json:"exon"`
	Gene         TentacledGene `json:"gene"`
	Position     float64       `json:"position"`
	Strand       StrandEnum    `json:"strand"`
	TranscriptID string        `json:"transcriptId"`
}

type TentacledGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
}

type RnaFusionFusionPartner5Prime struct {
	Exon         string     `json:"exon"`
	Gene         StickyGene `json:"gene"`
	Position     float64    `json:"position"`
	Strand       StrandEnum `json:"strand"`
	TranscriptID string     `json:"transcriptId"`
}

type StickyGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
}

type RNASeq struct {
	CohortRanking               *int64     `json:"cohortRanking,omitempty"`
	EnsemblID                   string     `json:"ensemblId"`
	EntrezID                    string     `json:"entrezId"`
	FragmentsPerKilobaseMillion float64    `json:"fragmentsPerKilobaseMillion"`
	FromNGS                     bool       `json:"fromNGS"`
	Gene                        RnaSeqGene `json:"gene"`
	ID                          string     `json:"id"`
	LibrarySize                 int64      `json:"librarySize"`
	RawCounts                   int64      `json:"rawCounts"`
	TissueCorrectedExpression   bool       `json:"tissueCorrectedExpression"`
	TranscriptID                string     `json:"transcriptId"`
}

type RnaSeqGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
}

type Snv struct {
	AllelicFrequency float64            `json:"allelicFrequency"`
	AltAllele        string             `json:"altAllele"`
	AminoAcidChange  *Coding            `json:"aminoAcidChange,omitempty"`
	Chromosome       Chromosome         `json:"chromosome"`
	CosmicID         *string            `json:"cosmicId,omitempty"`
	DBSNPID          *string            `json:"dbSNPId,omitempty"`
	DnaChange        *Coding            `json:"dnaChange,omitempty"`
	Gene             *SimpleVariantGene `json:"gene,omitempty"`
	ID               string             `json:"id"`
	Interpretation   *Coding            `json:"interpretation,omitempty"`
	ReadDepth        int64              `json:"readDepth"`
	RefAllele        string             `json:"refAllele"`
	StartEnd         StartEnd           `json:"startEnd"`
}

type SimpleVariantGene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
}

type StartEnd struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type NgsReportSpecimen struct {
	ID   string       `json:"id"`
	Type SpecimenType `json:"type"`
}

type MtbPatient struct {
	Address     *Address     `json:"address,omitempty"`
	BirthDate   string       `json:"birthDate"`
	DateOfDeath *string      `json:"dateOfDeath,omitempty"`
	Gender      CodingGender `json:"gender"`
	ID          string       `json:"id"`
	Insurance   *string      `json:"insurance,omitempty"`
}

type Address struct {
	MunicipalityCode string `json:"municipalityCode"`
}

type CodingGender struct {
	Code    Gender  `json:"code"`
	Display *string `json:"display,omitempty"`
	System  *string `json:"system,omitempty"`
	Version *string `json:"version,omitempty"`
}

type MTBMedicationRecommendation struct {
	Diagnosis          string                         `json:"diagnosis"`
	ID                 string                         `json:"id"`
	IssuedOn           *string                        `json:"issuedOn,omitempty"`
	LevelOfEvidence    *LevelOfEvidence               `json:"levelOfEvidence,omitempty"`
	Medication         []Coding                       `json:"medication,omitempty"`
	NgsReport          *string                        `json:"ngsReport,omitempty"`
	Patient            NoTargetFindingPatient         `json:"patient"`
	Priority           *TherapyRecommendationPriority `json:"priority,omitempty"`
	SupportingVariants []string                       `json:"supportingVariants,omitempty"`
}

type LevelOfEvidence struct {
	Addendums    []CodingLevelOfEvidenceAddendum `json:"addendums,omitempty"`
	Grading      CodingLevelOfEvidenceGrading    `json:"grading"`
	Publications []ReferencePublication          `json:"publications,omitempty"`
}

type CodingLevelOfEvidenceAddendum struct {
	Code    AddendumCode `json:"code"`
	Display *string      `json:"display,omitempty"`
	System  *string      `json:"system,omitempty"`
	Version *string      `json:"version,omitempty"`
}

type CodingLevelOfEvidenceGrading struct {
	Code    GradingCode `json:"code"`
	Display *string     `json:"display,omitempty"`
	System  *string     `json:"system,omitempty"`
	Version *string     `json:"version,omitempty"`
}

type ReferencePublication struct {
	EXTID *EXTID  `json:"extId,omitempty"`
	Type  *string `json:"type,omitempty"`
	URI   *string `json:"uri,omitempty"`
}

type EXTID struct {
	System *System `json:"system,omitempty"`
	Value  string  `json:"value"`
}

type Response struct {
	EffectiveDate string                 `json:"effectiveDate"`
	ID            string                 `json:"id"`
	Patient       NoTargetFindingPatient `json:"patient"`
	Therapy       ResponseTherapy        `json:"therapy"`
	Value         CodingRECIST           `json:"value"`
}

type ResponseTherapy struct {
	ID   string              `json:"id"`
	Type ResponseTherapyType `json:"type"`
}

type CodingRECIST struct {
	Code    FluffyCode `json:"code"`
	Display *string    `json:"display,omitempty"`
	System  *string    `json:"system,omitempty"`
	Version *string    `json:"version,omitempty"`
}

type SpecimenElement struct {
	Collection *Collection              `json:"collection,omitempty"`
	Icd10      *Coding                  `json:"icd10,omitempty"`
	ID         string                   `json:"id"`
	Patient    NoTargetFindingPatient   `json:"patient"`
	Type       *CodingTumorSpecimenType `json:"type,omitempty"`
}

type Collection struct {
	Date         string                                    `json:"date"`
	Localization CodingTumorSpecimenCollectionLocalization `json:"localization"`
	Method       CodingTumorSpecimenCollectionMethod       `json:"method"`
}

type CodingTumorSpecimenCollectionLocalization struct {
	Code    TumorSpecimenCollectionLocalization `json:"code"`
	Display *string                             `json:"display,omitempty"`
	System  *string                             `json:"system,omitempty"`
}

type CodingTumorSpecimenCollectionMethod struct {
	Code    TumorSpecimenCollectionMethod `json:"code"`
	Display *string                       `json:"display,omitempty"`
	System  *string                       `json:"system,omitempty"`
}

type CodingTumorSpecimenType struct {
	Code    TumorSpecimenType `json:"code"`
	Display *string           `json:"display,omitempty"`
	System  *string           `json:"system,omitempty"`
}

type StudyEnrollmentRecommendation struct {
	ID        string                 `json:"id"`
	IssuedOn  *string                `json:"issuedOn,omitempty"`
	NctNumber string                 `json:"nctNumber"`
	Patient   NoTargetFindingPatient `json:"patient"`
	Reason    string                 `json:"reason"`
}

type PatientType string

const (
	Patient PatientType = "Patient"
)

type ClaimResponseClaimType string

const (
	Claim ClaimResponseClaimType = "Claim"
)

type ClaimResponseStatusReason string

const (
	ApprovalRevocation               ClaimResponseStatusReason = "approval-revocation"
	ClaimResponseStatusReasonOther   ClaimResponseStatusReason = "other"
	ClaimResponseStatusReasonUnknown ClaimResponseStatusReason = "unknown"
	FormalReasons                    ClaimResponseStatusReason = "formal-reasons"
	InclusionInStudy                 ClaimResponseStatusReason = "inclusion-in-study"
	InsufficientEvidence             ClaimResponseStatusReason = "insufficient-evidence"
	OtherTherapyRecommended          ClaimResponseStatusReason = "other-therapy-recommended"
	StandardTherapyNotExhausted      ClaimResponseStatusReason = "standard-therapy-not-exhausted"
)

type ClaimResponseStatus string

const (
	Accepted                    ClaimResponseStatus = "accepted"
	ClaimResponseStatusRejected ClaimResponseStatus = "rejected"
	ClaimResponseStatusUnknown  ClaimResponseStatus = "unknown"
)

type ConsentStatus string

const (
	Active                ConsentStatus = "active"
	ConsentStatusRejected ConsentStatus = "rejected"
)

type MTBDiagnosisTumorSpread string

const (
	Local                          MTBDiagnosisTumorSpread = "local"
	MTBDiagnosisTumorSpreadUnknown MTBDiagnosisTumorSpread = "unknown"
	Metastasized                   MTBDiagnosisTumorSpread = "metastasized"
	TumorFree                      MTBDiagnosisTumorSpread = "tumor-free"
)

type PurpleCode string

const (
	Code0 PurpleCode = "0"
	Code1 PurpleCode = "1"
	Code2 PurpleCode = "2"
	Code3 PurpleCode = "3"
	Code4 PurpleCode = "4"
)

type SpecimenType string

const (
	TumorSpecimen SpecimenType = "TumorSpecimen"
)

type TumorCellContentMethod string

const (
	Bioinformatic TumorCellContentMethod = "bioinformatic"
	Histologic    TumorCellContentMethod = "histologic"
)

type NotDoneReasonCode string

const (
	ChronicRemission    NotDoneReasonCode = "chronic-remission"
	CodeOther           NotDoneReasonCode = "other"
	CodeUnknown         NotDoneReasonCode = "unknown"
	ContinuedExternally NotDoneReasonCode = "continued-externally"
	Deterioration       NotDoneReasonCode = "deterioration"
	LostToFu            NotDoneReasonCode = "lost-to-fu"
	MedicalReason       NotDoneReasonCode = "medical-reason"
	NoIndication        NotDoneReasonCode = "no-indication"
	OtherTherapyChosen  NotDoneReasonCode = "other-therapy-chosen"
	PatientDeath        NotDoneReasonCode = "patient-death"
	PatientRefusal      NotDoneReasonCode = "patient-refusal"
	PatientWish         NotDoneReasonCode = "patient-wish"
	PaymentEnded        NotDoneReasonCode = "payment-ended"
	PaymentPending      NotDoneReasonCode = "payment-pending"
	PaymentRefused      NotDoneReasonCode = "payment-refused"
	Progression         NotDoneReasonCode = "progression"
	Toxicity            NotDoneReasonCode = "toxicity"
)

type TherapyStatus string

const (
	Completed            TherapyStatus = "completed"
	NotDone              TherapyStatus = "not-done"
	OnGoing              TherapyStatus = "on-going"
	Stopped              TherapyStatus = "stopped"
	TherapyStatusUnknown TherapyStatus = "unknown"
)

type Chromosome string

const (
	Chr1  Chromosome = "chr1"
	Chr10 Chromosome = "chr10"
	Chr11 Chromosome = "chr11"
	Chr12 Chromosome = "chr12"
	Chr13 Chromosome = "chr13"
	Chr14 Chromosome = "chr14"
	Chr15 Chromosome = "chr15"
	Chr16 Chromosome = "chr16"
	Chr17 Chromosome = "chr17"
	Chr18 Chromosome = "chr18"
	Chr19 Chromosome = "chr19"
	Chr2  Chromosome = "chr2"
	Chr20 Chromosome = "chr20"
	Chr21 Chromosome = "chr21"
	Chr22 Chromosome = "chr22"
	Chr3  Chromosome = "chr3"
	Chr4  Chromosome = "chr4"
	Chr5  Chromosome = "chr5"
	Chr6  Chromosome = "chr6"
	Chr7  Chromosome = "chr7"
	Chr8  Chromosome = "chr8"
	Chr9  Chromosome = "chr9"
	ChrX  Chromosome = "chrX"
	ChrY  Chromosome = "chrY"
)

type CNVType string

const (
	HighLevelGain CNVType = "high-level-gain"
	Loss          CNVType = "loss"
	LowLevelGain  CNVType = "low-level-gain"
)

type StrandEnum string

const (
	Empty           StrandEnum = "+"
	RNAFusionStrand StrandEnum = "-"
)

type Gender string

const (
	Female        Gender = "female"
	GenderOther   Gender = "other"
	GenderUnknown Gender = "unknown"
	Male          Gender = "male"
)

type AddendumCode string

const (
	Is AddendumCode = "is"
	Iv AddendumCode = "iv"
	R  AddendumCode = "R"
	Z  AddendumCode = "Z"
)

type GradingCode string

const (
	M1A       GradingCode = "m1A"
	M1B       GradingCode = "m1B"
	M1C       GradingCode = "m1C"
	M2A       GradingCode = "m2A"
	M2B       GradingCode = "m2B"
	M2C       GradingCode = "m2C"
	M3        GradingCode = "m3"
	M4        GradingCode = "m4"
	Undefined GradingCode = "undefined"
)

type System string

const (
	HTTPSPubmedNcbiNlmNihGov System = "https://pubmed.ncbi.nlm.nih.gov/"
)

type TherapyRecommendationPriority string

const (
	Priority1 TherapyRecommendationPriority = "1"
	Priority2 TherapyRecommendationPriority = "2"
	Priority3 TherapyRecommendationPriority = "3"
	Priority4 TherapyRecommendationPriority = "4"
)

type ResponseTherapyType string

const (
	MTBMedicationTherapy ResponseTherapyType = "MTBMedicationTherapy"
)

type FluffyCode string

const (
	CR  FluffyCode = "CR"
	Mr  FluffyCode = "MR"
	Na  FluffyCode = "NA"
	Nya FluffyCode = "NYA"
	PD  FluffyCode = "PD"
	PR  FluffyCode = "PR"
	SD  FluffyCode = "SD"
)

type TumorSpecimenCollectionLocalization string

const (
	Metastasis                                 TumorSpecimenCollectionLocalization = "metastasis"
	PrimaryTumor                               TumorSpecimenCollectionLocalization = "primary-tumor"
	TumorSpecimenCollectionLocalizationUnknown TumorSpecimenCollectionLocalization = "unknown"
)

type TumorSpecimenCollectionMethod string

const (
	Biopsy                                    TumorSpecimenCollectionMethod = "biopsy"
	Cytology                                  TumorSpecimenCollectionMethod = "cytology"
	Resection                                 TumorSpecimenCollectionMethod = "resection"
	TumorSpecimenCollectionMethodLiquidBiopsy TumorSpecimenCollectionMethod = "liquid-biopsy"
	TumorSpecimenCollectionMethodUnknown      TumorSpecimenCollectionMethod = "unknown"
)

type TumorSpecimenType string

const (
	CryoFrozen                    TumorSpecimenType = "cryo-frozen"
	Ffpe                          TumorSpecimenType = "FFPE"
	FreshTissue                   TumorSpecimenType = "fresh-tissue"
	TumorSpecimenTypeLiquidBiopsy TumorSpecimenType = "liquid-biopsy"
	TumorSpecimenTypeUnknown      TumorSpecimenType = "unknown"
)
