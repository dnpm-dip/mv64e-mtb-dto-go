package mtb

import (
	"bytes"
	"encoding/json"
)

func UnmarshalMtb(data []byte) (Mtb, error) {
	var r Mtb
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	err := dec.Decode(&r)
	return r, err
}

func (r *Mtb) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Mtb struct {
	CarePlans              []MTBCarePlan              `json:"carePlans,omitempty"`
	ClaimResponses         []ClaimResponse            `json:"claimResponses,omitempty"`
	Claims                 []Claim                    `json:"claims,omitempty"`
	Diagnoses              []MTBDiagnosis             `json:"diagnoses"`
	EpisodesOfCare         []MTBEpisodeOfCare         `json:"episodesOfCare"`
	FollowUPS              []FollowUp                 `json:"followUps,omitempty"`
	GuidelineProcedures    []OncoProcedure            `json:"guidelineProcedures,omitempty"`
	GuidelineTherapies     []MTBSystemicTherapy       `json:"guidelineTherapies,omitempty"`
	HistologyReports       []HistologyReport          `json:"histologyReports,omitempty"`
	IhcReports             []IHCReport                `json:"ihcReports,omitempty"`
	NgsReports             []SomaticNGSReportMetadata `json:"ngsReports,omitempty"`
	Patient                Patient                    `json:"patient"`
	PerformanceStatus      []PerformanceStatus        `json:"performanceStatus,omitempty"`
	PriorDiagnosticReports []PriorDiagnosticReport    `json:"priorDiagnosticReports,omitempty"`
	Responses              []Response                 `json:"responses,omitempty"`
	Specimens              []TumorSpecimen            `json:"specimens,omitempty"`
	SystemicTherapies      []SystemicTherapy          `json:"systemicTherapies,omitempty"`
}

type MTBCarePlan struct {
	GeneticCounselingRecommendation *GeneticCounselingRecommendation   `json:"geneticCounselingRecommendation,omitempty"`
	HistologyReevaluationRequests   []HistologyReevaluationRequest     `json:"histologyReevaluationRequests,omitempty"`
	ID                              string                             `json:"id"`
	IssuedOn                        string                             `json:"issuedOn"`
	MedicationRecommendations       []MTBMedicationRecommendation      `json:"medicationRecommendations,omitempty"`
	Notes                           []string                           `json:"notes,omitempty"`
	Patient                         Reference                          `json:"patient"`
	ProcedureRecommendations        []ProcedureRecommendation          `json:"procedureRecommendations,omitempty"`
	Reason                          *Reference                         `json:"reason,omitempty"`
	RebiopsyRequests                []RebiopsyRequest                  `json:"rebiopsyRequests,omitempty"`
	StatusReason                    *CodingMTBCarePlanStatusReason     `json:"statusReason,omitempty"`
	StudyEnrollmentRecommendations  []MTBStudyEnrollmentRecommendation `json:"studyEnrollmentRecommendations,omitempty"`
}

type GeneticCounselingRecommendation struct {
	ID       string                                      `json:"id"`
	IssuedOn string                                      `json:"issuedOn"`
	Patient  Reference                                   `json:"patient"`
	Reason   CodingGeneticCounselingRecommendationReason `json:"reason"`
}

type Reference struct {
	Display *string `json:"display,omitempty"`
	ID      string  `json:"id"`
	System  *string `json:"system,omitempty"`
	Type    *string `json:"type,omitempty"`
}

type CodingGeneticCounselingRecommendationReason struct {
	Code    ReasonCode `json:"code"`
	Display *string    `json:"display,omitempty"`
	System  *string    `json:"system,omitempty"`
	Version *string    `json:"version,omitempty"`
}

type HistologyReevaluationRequest struct {
	ID       string    `json:"id"`
	IssuedOn string    `json:"issuedOn"`
	Patient  Reference `json:"patient"`
	Specimen Reference `json:"specimen"`
}

type MTBMedicationRecommendation struct {
	Category           *CodingMTBMedicationRecommendationCategory `json:"category,omitempty"`
	ID                 string                                     `json:"id"`
	IssuedOn           string                                     `json:"issuedOn"`
	LevelOfEvidence    *LevelOfEvidence                           `json:"levelOfEvidence,omitempty"`
	Medication         []CodingATCUnregisteredMedication          `json:"medication"`
	Patient            Reference                                  `json:"patient"`
	Priority           CodingRecommendationPriority               `json:"priority"`
	Reason             *Reference                                 `json:"reason,omitempty"`
	SupportingVariants []GeneAlterationReference                  `json:"supportingVariants,omitempty"`
	UseType            *CodingMTBMedicationRecommendationUseType  `json:"useType,omitempty"`
}

type CodingMTBMedicationRecommendationCategory struct {
	Code    CodingMTBMedicationRecommendationCategoryCode `json:"code"`
	Display *string                                       `json:"display,omitempty"`
	System  *string                                       `json:"system,omitempty"`
	Version *string                                       `json:"version,omitempty"`
}

type LevelOfEvidence struct {
	Addendums    []CodingLevelOfEvidenceAddendum `json:"addendums,omitempty"`
	Grading      CodingLevelOfEvidenceGrading    `json:"grading"`
	Publications []PublicationReference          `json:"publications,omitempty"`
}

type CodingLevelOfEvidenceAddendum struct {
	Code    AddendumCode `json:"code"`
	Display *string      `json:"display,omitempty"`
	System  *string      `json:"system,omitempty"`
	Version *string      `json:"version,omitempty"`
}

type CodingLevelOfEvidenceGrading struct {
	Code    LevelOfEvidenceCode `json:"code"`
	Display *string             `json:"display,omitempty"`
	System  *string             `json:"system,omitempty"`
	Version *string             `json:"version,omitempty"`
}

type PublicationReference struct {
	Display *string           `json:"display,omitempty"`
	ID      string            `json:"id"`
	System  PublicationSystem `json:"system"`
	Type    *string           `json:"type,omitempty"`
}

type CodingATCUnregisteredMedication struct {
	Code    string                    `json:"code"`
	Display *string                   `json:"display,omitempty"`
	System  RequestedMedicationSystem `json:"system"`
	Version *string                   `json:"version,omitempty"`
}

type CodingRecommendationPriority struct {
	Code    PriorityCode `json:"code"`
	Display *string      `json:"display,omitempty"`
	System  *string      `json:"system,omitempty"`
	Version *string      `json:"version,omitempty"`
}

type GeneAlterationReference struct {
	Display *string   `json:"display,omitempty"`
	Gene    *Coding   `json:"gene,omitempty"`
	Variant Reference `json:"variant"`
}

type Coding struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	System  *string `json:"system,omitempty"`
	Version *string `json:"version,omitempty"`
}

type CodingMTBMedicationRecommendationUseType struct {
	Code    UseTypeCode `json:"code"`
	Display *string     `json:"display,omitempty"`
	System  *string     `json:"system,omitempty"`
	Version *string     `json:"version,omitempty"`
}

type ProcedureRecommendation struct {
	Code               CodingMTBProcedureRecommendationCategory `json:"code"`
	ID                 string                                   `json:"id"`
	IssuedOn           string                                   `json:"issuedOn"`
	LevelOfEvidence    *LevelOfEvidence                         `json:"levelOfEvidence,omitempty"`
	Patient            Reference                                `json:"patient"`
	Priority           CodingRecommendationPriority             `json:"priority"`
	Reason             *Reference                               `json:"reason,omitempty"`
	SupportingVariants []GeneAlterationReference                `json:"supportingVariants,omitempty"`
}

type CodingMTBProcedureRecommendationCategory struct {
	Code    CodingMTBProcedureRecommendationCategoryCode `json:"code"`
	Display *string                                      `json:"display,omitempty"`
	System  *string                                      `json:"system,omitempty"`
	Version *string                                      `json:"version,omitempty"`
}

type RebiopsyRequest struct {
	ID          string    `json:"id"`
	IssuedOn    string    `json:"issuedOn"`
	Patient     Reference `json:"patient"`
	TumorEntity Reference `json:"tumorEntity"`
}

type CodingMTBCarePlanStatusReason struct {
	Code    CodingMTBCarePlanStatusReasonCode `json:"code"`
	Display *string                           `json:"display,omitempty"`
	System  *string                           `json:"system,omitempty"`
	Version *string                           `json:"version,omitempty"`
}

type MTBStudyEnrollmentRecommendation struct {
	ID                 string                            `json:"id"`
	IssuedOn           string                            `json:"issuedOn"`
	LevelOfEvidence    *CodingLevelOfEvidenceGrading     `json:"levelOfEvidence,omitempty"`
	Medication         []CodingATCUnregisteredMedication `json:"medication,omitempty"`
	Patient            Reference                         `json:"patient"`
	Priority           CodingRecommendationPriority      `json:"priority"`
	Reason             Reference                         `json:"reason"`
	Study              []StudyReference                  `json:"study"`
	SupportingVariants []GeneAlterationReference         `json:"supportingVariants,omitempty"`
}

type StudyReference struct {
	Display *string     `json:"display,omitempty"`
	ID      string      `json:"id"`
	System  StudySystem `json:"system"`
	Type    *string     `json:"type,omitempty"`
}

type ClaimResponse struct {
	Claim        Reference                        `json:"claim"`
	ID           string                           `json:"id"`
	IssuedOn     string                           `json:"issuedOn"`
	Patient      Reference                        `json:"patient"`
	Status       CodingClaimResponseStatus        `json:"status"`
	StatusReason *CodingClaimResponseStatusReason `json:"statusReason,omitempty"`
}

type CodingClaimResponseStatus struct {
	Code    CodingClaimResponseStatusCode `json:"code"`
	Display *string                       `json:"display,omitempty"`
	System  *string                       `json:"system,omitempty"`
	Version *string                       `json:"version,omitempty"`
}

type CodingClaimResponseStatusReason struct {
	Code    CodingClaimResponseStatusReasonCode `json:"code"`
	Display *string                             `json:"display,omitempty"`
	System  *string                             `json:"system,omitempty"`
	Version *string                             `json:"version,omitempty"`
}

type Claim struct {
	ID                  string                            `json:"id"`
	IssuedOn            string                            `json:"issuedOn"`
	Patient             Reference                         `json:"patient"`
	Recommendation      Reference                         `json:"recommendation"`
	RequestedMedication []CodingATCUnregisteredMedication `json:"requestedMedication,omitempty"`
	Stage               *CodingClaimStage                 `json:"stage,omitempty"`
}

type CodingClaimStage struct {
	Code    StageCode `json:"code"`
	Display *string   `json:"display,omitempty"`
	System  *string   `json:"system,omitempty"`
	Version *string   `json:"version,omitempty"`
}

type MTBDiagnosis struct {
	Code                     Coding                                      `json:"code"`
	GermlineCodes            []Coding                                    `json:"germlineCodes,omitempty"`
	Grading                  *Grading                                    `json:"grading,omitempty"`
	GuidelineTreatmentStatus *CodingMTBDiagnosisGuidelineTreatmentStatus `json:"guidelineTreatmentStatus,omitempty"`
	Histology                []Reference                                 `json:"histology,omitempty"`
	ID                       string                                      `json:"id"`
	Notes                    []string                                    `json:"notes,omitempty"`
	Patient                  Reference                                   `json:"patient"`
	RecordedOn               *string                                     `json:"recordedOn,omitempty"`
	Staging                  *Staging                                    `json:"staging,omitempty"`
	Topography               *Coding                                     `json:"topography,omitempty"`
	Type                     Type                                        `json:"type"`
}

type Grading struct {
	History []TumorGrading `json:"history"`
}

type TumorGrading struct {
	Codes []Coding `json:"codes"`
	Date  string   `json:"date"`
}

type CodingMTBDiagnosisGuidelineTreatmentStatus struct {
	Code    GuidelineTreatmentStatusCode `json:"code"`
	Display *string                      `json:"display,omitempty"`
	System  *string                      `json:"system,omitempty"`
	Version *string                      `json:"version,omitempty"`
}

type Staging struct {
	History []TumorStaging `json:"history"`
}

type TumorStaging struct {
	Date                 string                   `json:"date"`
	Method               CodingTumorStagingMethod `json:"method"`
	OtherClassifications []Coding                 `json:"otherClassifications,omitempty"`
	TnmClassification    *TnmClassification       `json:"tnmClassification,omitempty"`
}

type CodingTumorStagingMethod struct {
	Code    CodingTumorStagingMethodCode `json:"code"`
	Display *string                      `json:"display,omitempty"`
	System  *string                      `json:"system,omitempty"`
	Version *string                      `json:"version,omitempty"`
}

type TnmClassification struct {
	Metastasis Coding `json:"metastasis"`
	Nodes      Coding `json:"nodes"`
	Tumor      Coding `json:"tumor"`
}

type Type struct {
	History []History `json:"history"`
}

type History struct {
	Date  string             `json:"date"`
	Value CodingMTBDiagnosis `json:"value"`
}

type CodingMTBDiagnosis struct {
	Code    CodingMTBDiagnosisCode `json:"code"`
	Display *string                `json:"display,omitempty"`
	System  *string                `json:"system,omitempty"`
	Version *string                `json:"version,omitempty"`
}

type MTBEpisodeOfCare struct {
	Diagnoses []Reference `json:"diagnoses,omitempty"`
	ID        string      `json:"id"`
	Patient   Reference   `json:"patient"`
	Period    PeriodDate  `json:"period"`
}

type PeriodDate struct {
	End   *string `json:"end,omitempty"`
	Start string  `json:"start"`
}

type FollowUp struct {
	Date          string                       `json:"date"`
	PatientStatus *CodingFollowUpPatientStatus `json:"patientStatus,omitempty"`
}

type CodingFollowUpPatientStatus struct {
	Code    PatientStatusCode `json:"code"`
	Display *string           `json:"display,omitempty"`
	System  *string           `json:"system,omitempty"`
	Version *string           `json:"version,omitempty"`
}

type OncoProcedure struct {
	BasedOn      *Reference                    `json:"basedOn,omitempty"`
	Code         CodingOncoProcedure           `json:"code"`
	ID           string                        `json:"id"`
	Intent       *CodingMTBTherapyIntent       `json:"intent,omitempty"`
	Notes        []string                      `json:"notes,omitempty"`
	Patient      Reference                     `json:"patient"`
	Period       *PeriodDate                   `json:"period,omitempty"`
	Reason       *Reference                    `json:"reason,omitempty"`
	RecordedOn   string                        `json:"recordedOn"`
	Status       CodingTherapyStatus           `json:"status"`
	StatusReason *CodingMTBTherapyStatusReason `json:"statusReason,omitempty"`
	TherapyLine  *int64                        `json:"therapyLine,omitempty"`
}

type CodingOncoProcedure struct {
	Code    CodingOncoProcedureCode `json:"code"`
	Display *string                 `json:"display,omitempty"`
	System  *string                 `json:"system,omitempty"`
	Version *string                 `json:"version,omitempty"`
}

type CodingMTBTherapyIntent struct {
	Code    IntentCode `json:"code"`
	Display *string    `json:"display,omitempty"`
	System  *string    `json:"system,omitempty"`
	Version *string    `json:"version,omitempty"`
}

type CodingTherapyStatus struct {
	Code    CodingTherapyStatusCode `json:"code"`
	Display *string                 `json:"display,omitempty"`
	System  *string                 `json:"system,omitempty"`
	Version *string                 `json:"version,omitempty"`
}

type CodingMTBTherapyStatusReason struct {
	Code    CodingMTBTherapyStatusReasonCode `json:"code"`
	Display *string                          `json:"display,omitempty"`
	System  *string                          `json:"system,omitempty"`
	Version *string                          `json:"version,omitempty"`
}

type MTBSystemicTherapy struct {
	BasedOn                         *Reference                                               `json:"basedOn,omitempty"`
	Category                        *CodingMTBSystemicTherapyCategory                        `json:"category,omitempty"`
	ID                              string                                                   `json:"id"`
	Intent                          *CodingMTBTherapyIntent                                  `json:"intent,omitempty"`
	Medication                      []CodingATCUnregisteredMedication                        `json:"medication,omitempty"`
	Notes                           []string                                                 `json:"notes,omitempty"`
	Patient                         Reference                                                `json:"patient"`
	Period                          *PeriodDate                                              `json:"period,omitempty"`
	Reason                          *Reference                                               `json:"reason,omitempty"`
	RecommendationFulfillmentStatus *CodingMTBSystemicTherapyRecommendationFulfillmentStatus `json:"recommendationFulfillmentStatus,omitempty"`
	RecordedOn                      string                                                   `json:"recordedOn"`
	Status                          CodingTherapyStatus                                      `json:"status"`
	StatusReason                    *CodingMTBTherapyStatusReason                            `json:"statusReason,omitempty"`
	TherapyLine                     *int64                                                   `json:"therapyLine,omitempty"`
}

type CodingMTBSystemicTherapyCategory struct {
	Code    CodingMTBSystemicTherapyCategoryCode `json:"code"`
	Display *string                              `json:"display,omitempty"`
	System  *string                              `json:"system,omitempty"`
	Version *string                              `json:"version,omitempty"`
}

type CodingMTBSystemicTherapyRecommendationFulfillmentStatus struct {
	Code    RecommendationFulfillmentStatusCode `json:"code"`
	Display *string                             `json:"display,omitempty"`
	System  *string                             `json:"system,omitempty"`
	Version *string                             `json:"version,omitempty"`
}

type HistologyReport struct {
	ID       string                 `json:"id"`
	IssuedOn string                 `json:"issuedOn"`
	Patient  Reference              `json:"patient"`
	Results  HistologyReportResults `json:"results"`
	Specimen Reference              `json:"specimen"`
}

type HistologyReportResults struct {
	TumorCellContent *TumorCellContent `json:"tumorCellContent,omitempty"`
	TumorMorphology  TumorMorphology   `json:"tumorMorphology"`
}

type TumorCellContent struct {
	ID       string                       `json:"id"`
	Method   CodingTumorCellContentMethod `json:"method"`
	Patient  Reference                    `json:"patient"`
	Specimen Reference                    `json:"specimen"`
	Value    float64                      `json:"value"`
}

type CodingTumorCellContentMethod struct {
	Code    CodingTumorCellContentMethodCode `json:"code"`
	Display *string                          `json:"display,omitempty"`
	System  *string                          `json:"system,omitempty"`
	Version *string                          `json:"version,omitempty"`
}

type TumorMorphology struct {
	ID       string    `json:"id"`
	Notes    *string   `json:"notes,omitempty"`
	Patient  Reference `json:"patient"`
	Specimen Reference `json:"specimen"`
	Value    Coding    `json:"value"`
}

type IHCReport struct {
	BlockIDS  []string         `json:"blockIds"`
	ID        string           `json:"id"`
	IssuedOn  string           `json:"issuedOn"`
	JournalID string           `json:"journalId"`
	Patient   Reference        `json:"patient"`
	Results   IhcReportResults `json:"results"`
	Specimen  Reference        `json:"specimen"`
}

type IhcReportResults struct {
	MSIMmr            []MSIMmr            `json:"msiMmr"`
	ProteinExpression []ProteinExpression `json:"proteinExpression"`
}

type MSIMmr struct {
	ICScore  *CodingProteinExpressionICScore `json:"icScore,omitempty"`
	ID       string                          `json:"id"`
	Patient  Reference                       `json:"patient"`
	Protein  Coding                          `json:"protein"`
	TcScore  *CodingProteinExpressionTCScore `json:"tcScore,omitempty"`
	TpsScore *int64                          `json:"tpsScore,omitempty"`
	Value    CodingProteinExpressionResult   `json:"value"`
}

type CodingProteinExpressionICScore struct {
	Code    ICScoreCode `json:"code"`
	Display *string     `json:"display,omitempty"`
	System  *string     `json:"system,omitempty"`
	Version *string     `json:"version,omitempty"`
}

type CodingProteinExpressionTCScore struct {
	Code    TcScoreCode `json:"code"`
	Display *string     `json:"display,omitempty"`
	System  *string     `json:"system,omitempty"`
	Version *string     `json:"version,omitempty"`
}

type CodingProteinExpressionResult struct {
	Code    CodingProteinExpressionResultCode `json:"code"`
	Display *string                           `json:"display,omitempty"`
	System  *string                           `json:"system,omitempty"`
	Version *string                           `json:"version,omitempty"`
}

type ProteinExpression struct {
	ICScore  *CodingProteinExpressionICScore `json:"icScore,omitempty"`
	ID       string                          `json:"id"`
	Patient  Reference                       `json:"patient"`
	Protein  Coding                          `json:"protein"`
	TcScore  *CodingProteinExpressionTCScore `json:"tcScore,omitempty"`
	TpsScore *int64                          `json:"tpsScore,omitempty"`
	Value    CodingProteinExpressionResult   `json:"value"`
}

type SomaticNGSReportMetadata struct {
	ID       string           `json:"id"`
	IssuedOn string           `json:"issuedOn"`
	Metadata []Metadata       `json:"metadata"`
	Patient  Reference        `json:"patient"`
	Results  NgsReportResults `json:"results"`
	Specimen Reference        `json:"specimen"`
	Type     CodingNGSReport  `json:"type"`
}

type Metadata struct {
	KitManufacturer string `json:"kitManufacturer"`
	KitType         string `json:"kitType"`
	Pipeline        string `json:"pipeline"`
	ReferenceGenome string `json:"referenceGenome"`
	Sequencer       string `json:"sequencer"`
}

type NgsReportResults struct {
	Brcaness           *BRCAness         `json:"brcaness,omitempty"`
	CopyNumberVariants []Cnv             `json:"copyNumberVariants"`
	DnaFusions         []DNAFusion       `json:"dnaFusions"`
	HrdScore           *HRDScore         `json:"hrdScore,omitempty"`
	RnaFusions         []RNAFusion       `json:"rnaFusions"`
	RnaSeqs            []RNASeq          `json:"rnaSeqs"`
	SimpleVariants     []Snv             `json:"simpleVariants"`
	Tmb                *Tmb              `json:"tmb,omitempty"`
	TumorCellContent   *TumorCellContent `json:"tumorCellContent,omitempty"`
}

type BRCAness struct {
	ConfidenceRange ConfidenceRange `json:"confidenceRange"`
	ID              string          `json:"id"`
	Patient         Reference       `json:"patient"`
	Specimen        Reference       `json:"specimen"`
	Value           float64         `json:"value"`
}

type ConfidenceRange struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type Cnv struct {
	CNA                   *float64                        `json:"cnA,omitempty"`
	CNB                   *float64                        `json:"cnB,omitempty"`
	Chromosome            Chromosome                      `json:"chromosome"`
	CopyNumberNeutralLoH  []Coding                        `json:"copyNumberNeutralLoH,omitempty"`
	EndRange              *EndRange                       `json:"endRange,omitempty"`
	ExternalIDS           []VariantExternalID             `json:"externalIds,omitempty"`
	ID                    string                          `json:"id"`
	Localization          []CodingBaseVariantLocalization `json:"localization,omitempty"`
	Patient               Reference                       `json:"patient"`
	RelativeCopyNumber    *float64                        `json:"relativeCopyNumber,omitempty"`
	ReportedAffectedGenes []Coding                        `json:"reportedAffectedGenes,omitempty"`
	ReportedFocality      *string                         `json:"reportedFocality,omitempty"`
	StartRange            *StartRange                     `json:"startRange,omitempty"`
	TotalCopyNumber       *int64                          `json:"totalCopyNumber,omitempty"`
	Type                  CodingCNV                       `json:"type"`
}

type EndRange struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type VariantExternalID struct {
	System ExternalIDSystem `json:"system"`
	Value  string           `json:"value"`
}

type CodingBaseVariantLocalization struct {
	Code    CodingBaseVariantLocalizationCode `json:"code"`
	Display *string                           `json:"display,omitempty"`
	System  *string                           `json:"system,omitempty"`
	Version *string                           `json:"version,omitempty"`
}

type StartRange struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type CodingCNV struct {
	Code    CodingCNVCode `json:"code"`
	Display *string       `json:"display,omitempty"`
	System  *string       `json:"system,omitempty"`
	Version *string       `json:"version,omitempty"`
}

type DNAFusion struct {
	ExternalIDS         []VariantExternalID             `json:"externalIds,omitempty"`
	FusionPartner3Prime DnaFusionFusionPartner3Prime    `json:"fusionPartner3prime"`
	FusionPartner5Prime DnaFusionFusionPartner5Prime    `json:"fusionPartner5prime"`
	ID                  string                          `json:"id"`
	Localization        []CodingBaseVariantLocalization `json:"localization,omitempty"`
	Patient             Reference                       `json:"patient"`
	ReportedNumReads    int64                           `json:"reportedNumReads"`
}

type DnaFusionFusionPartner3Prime struct {
	Chromosome Chromosome `json:"chromosome"`
	Gene       Coding     `json:"gene"`
	Position   float64    `json:"position"`
}

type DnaFusionFusionPartner5Prime struct {
	Chromosome Chromosome `json:"chromosome"`
	Gene       Coding     `json:"gene"`
	Position   float64    `json:"position"`
}

type HRDScore struct {
	Components     Components                    `json:"components"`
	ID             string                        `json:"id"`
	Interpretation *CodingHRDScoreInterpretation `json:"interpretation,omitempty"`
	Patient        Reference                     `json:"patient"`
	Specimen       Reference                     `json:"specimen"`
	Value          float64                       `json:"value"`
}

type Components struct {
	Loh float64 `json:"loh"`
	Lst float64 `json:"lst"`
	Tai float64 `json:"tai"`
}

type CodingHRDScoreInterpretation struct {
	Code    InterpretationCode `json:"code"`
	Display *string            `json:"display,omitempty"`
	System  *string            `json:"system,omitempty"`
	Version *string            `json:"version,omitempty"`
}

type RNAFusion struct {
	Effect              *string                         `json:"effect,omitempty"`
	ExternalIDS         []VariantExternalID             `json:"externalIds,omitempty"`
	FusionPartner3Prime RnaFusionFusionPartner3Prime    `json:"fusionPartner3prime"`
	FusionPartner5Prime RnaFusionFusionPartner5Prime    `json:"fusionPartner5prime"`
	ID                  string                          `json:"id"`
	Localization        []CodingBaseVariantLocalization `json:"localization,omitempty"`
	Patient             Reference                       `json:"patient"`
	ReportedNumReads    int64                           `json:"reportedNumReads"`
}

type RnaFusionFusionPartner3Prime struct {
	ExonID       string       `json:"exonId"`
	Gene         Coding       `json:"gene"`
	Position     float64      `json:"position"`
	Strand       StrandEnum   `json:"strand"`
	TranscriptID TranscriptID `json:"transcriptId"`
}

type TranscriptID struct {
	System TranscriptIDSystem `json:"system"`
	Value  string             `json:"value"`
}

type RnaFusionFusionPartner5Prime struct {
	ExonID       string       `json:"exonId"`
	Gene         Coding       `json:"gene"`
	Position     float64      `json:"position"`
	Strand       StrandEnum   `json:"strand"`
	TranscriptID TranscriptID `json:"transcriptId"`
}

type RNASeq struct {
	CohortRanking             *int64                          `json:"cohortRanking,omitempty"`
	ExternalIDS               []VariantExternalID             `json:"externalIds,omitempty"`
	Gene                      *Coding                         `json:"gene,omitempty"`
	ID                        string                          `json:"id"`
	LibrarySize               *int64                          `json:"librarySize,omitempty"`
	Localization              []CodingBaseVariantLocalization `json:"localization,omitempty"`
	Patient                   Reference                       `json:"patient"`
	RawCounts                 int64                           `json:"rawCounts"`
	TissueCorrectedExpression *bool                           `json:"tissueCorrectedExpression,omitempty"`
	TranscriptID              *TranscriptID                   `json:"transcriptId,omitempty"`
	TranscriptsPerMillion     float64                         `json:"transcriptsPerMillion"`
	Variant                   Reference                       `json:"variant"`
}

type Snv struct {
	AllelicFrequency float64                         `json:"allelicFrequency"`
	AltAllele        string                          `json:"altAllele"`
	Chromosome       Chromosome                      `json:"chromosome"`
	DnaChange        *Coding                         `json:"dnaChange,omitempty"`
	ExonID           *string                         `json:"exonId,omitempty"`
	ExternalIDS      []VariantExternalID             `json:"externalIds,omitempty"`
	Gene             *Coding                         `json:"gene,omitempty"`
	ID               string                          `json:"id"`
	Interpretation   *CodingClinVar                  `json:"interpretation,omitempty"`
	Localization     []CodingBaseVariantLocalization `json:"localization,omitempty"`
	Patient          Reference                       `json:"patient"`
	Position         Position                        `json:"position"`
	ProteinChange    *Coding                         `json:"proteinChange,omitempty"`
	ReadDepth        int64                           `json:"readDepth"`
	RefAllele        string                          `json:"refAllele"`
	TranscriptID     TranscriptID                    `json:"transcriptId"`
}

type CodingClinVar struct {
	Code    CodingClinVarCode `json:"code"`
	Display *string           `json:"display,omitempty"`
	System  *string           `json:"system,omitempty"`
	Version *string           `json:"version,omitempty"`
}

type Position struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type Tmb struct {
	ID             string                   `json:"id"`
	Interpretation *CodingTMBInterpretation `json:"interpretation,omitempty"`
	Patient        Reference                `json:"patient"`
	Specimen       Reference                `json:"specimen"`
	Value          TMBResult                `json:"value"`
}

type CodingTMBInterpretation struct {
	Code    InterpretationCode `json:"code"`
	Display *string            `json:"display,omitempty"`
	System  *string            `json:"system,omitempty"`
	Version *string            `json:"version,omitempty"`
}

type TMBResult struct {
	Unit  *string `json:"unit,omitempty"`
	Value float64 `json:"value"`
}

type CodingNGSReport struct {
	Code    CodingNGSReportCode `json:"code"`
	Display *string             `json:"display,omitempty"`
	System  *string             `json:"system,omitempty"`
	Version *string             `json:"version,omitempty"`
}

type Patient struct {
	Address         Address            `json:"address"`
	Age             *Age               `json:"age,omitempty"`
	BirthDate       string             `json:"birthDate"`
	DateOfDeath     *string            `json:"dateOfDeath,omitempty"`
	Gender          CodingGender       `json:"gender"`
	HealthInsurance HealthInsurance    `json:"healthInsurance"`
	ID              string             `json:"id"`
	ManagingSite    *Coding            `json:"managingSite,omitempty"`
	VitalStatus     *CodingVitalStatus `json:"vitalStatus,omitempty"`
}

type Address struct {
	MunicipalityCode string `json:"municipalityCode"`
}

type Age struct {
	Unit  Unit    `json:"unit"`
	Value float64 `json:"value"`
}

type CodingGender struct {
	Code    GenderCode `json:"code"`
	Display *string    `json:"display,omitempty"`
	System  *string    `json:"system,omitempty"`
	Version *string    `json:"version,omitempty"`
}

type HealthInsurance struct {
	Reference *Reference            `json:"reference,omitempty"`
	Type      CodingHealthInsurance `json:"type"`
}

type CodingHealthInsurance struct {
	Code    CodingHealthInsuranceCode `json:"code"`
	Display *string                   `json:"display,omitempty"`
	System  *string                   `json:"system,omitempty"`
	Version *string                   `json:"version,omitempty"`
}

type CodingVitalStatus struct {
	Code    VitalStatusCode `json:"code"`
	Display *string         `json:"display,omitempty"`
	System  *string         `json:"system,omitempty"`
	Version *string         `json:"version,omitempty"`
}

type PerformanceStatus struct {
	EffectiveDate string     `json:"effectiveDate"`
	ID            string     `json:"id"`
	Patient       Reference  `json:"patient"`
	Value         CodingECOG `json:"value"`
}

type CodingECOG struct {
	Code    CodingECOGCode `json:"code"`
	Display *string        `json:"display,omitempty"`
	System  *string        `json:"system,omitempty"`
	Version *string        `json:"version,omitempty"`
}

type PriorDiagnosticReport struct {
	ID        string                          `json:"id"`
	IssuedOn  string                          `json:"issuedOn"`
	Patient   Reference                       `json:"patient"`
	Performer *Reference                      `json:"performer,omitempty"`
	Results   []string                        `json:"results,omitempty"`
	Specimen  Reference                       `json:"specimen"`
	Type      CodingMolecularDiagnosticReport `json:"type"`
}

type CodingMolecularDiagnosticReport struct {
	Code    CodingMolecularDiagnosticReportCode `json:"code"`
	Display *string                             `json:"display,omitempty"`
	System  *string                             `json:"system,omitempty"`
	Version *string                             `json:"version,omitempty"`
}

type Response struct {
	EffectiveDate string               `json:"effectiveDate"`
	ID            string               `json:"id"`
	Method        CodingResponseMethod `json:"method"`
	Patient       Reference            `json:"patient"`
	Therapy       Reference            `json:"therapy"`
	Value         CodingRECIST         `json:"value"`
}

type CodingResponseMethod struct {
	Code    CodingResponseMethodCode `json:"code"`
	Display *string                  `json:"display,omitempty"`
	System  *string                  `json:"system,omitempty"`
	Version *string                  `json:"version,omitempty"`
}

type CodingRECIST struct {
	Code    CodingRECISTCode `json:"code"`
	Display *string          `json:"display,omitempty"`
	System  *string          `json:"system,omitempty"`
	Version *string          `json:"version,omitempty"`
}

type TumorSpecimen struct {
	Collection *Collection         `json:"collection,omitempty"`
	Diagnosis  Reference           `json:"diagnosis"`
	ID         string              `json:"id"`
	Patient    Reference           `json:"patient"`
	Type       CodingTumorSpecimen `json:"type"`
}

type Collection struct {
	Date         *string                                   `json:"date,omitempty"`
	Localization CodingTumorSpecimenCollectionLocalization `json:"localization"`
	Method       CodingTumorSpecimenCollectionMethod       `json:"method"`
}

type CodingTumorSpecimenCollectionLocalization struct {
	Code    CodingTumorSpecimenCollectionLocalizationCode `json:"code"`
	Display *string                                       `json:"display,omitempty"`
	System  *string                                       `json:"system,omitempty"`
	Version *string                                       `json:"version,omitempty"`
}

type CodingTumorSpecimenCollectionMethod struct {
	Code    CodingTumorSpecimenCollectionMethodCode `json:"code"`
	Display *string                                 `json:"display,omitempty"`
	System  *string                                 `json:"system,omitempty"`
	Version *string                                 `json:"version,omitempty"`
}

type CodingTumorSpecimen struct {
	Code    CodingTumorSpecimenCode `json:"code"`
	Display *string                 `json:"display,omitempty"`
	System  *string                 `json:"system,omitempty"`
	Version *string                 `json:"version,omitempty"`
}

type SystemicTherapy struct {
	History []MTBSystemicTherapy `json:"history"`
}

type ReasonCode string

const (
	FamilyAnamnesis ReasonCode = "family-anamnesis"
	PurpleOther     ReasonCode = "other"
	PurpleUnknown   ReasonCode = "unknown"
	SecondaryTumor  ReasonCode = "secondary-tumor"
	SelfAnamnesis   ReasonCode = "self-anamnesis"
)

type CodingMTBMedicationRecommendationCategoryCode string

const (
	Ch       CodingMTBMedicationRecommendationCategoryCode = "CH"
	Ho       CodingMTBMedicationRecommendationCategoryCode = "HO"
	IM       CodingMTBMedicationRecommendationCategoryCode = "IM"
	PurpleSO CodingMTBMedicationRecommendationCategoryCode = "SO"
	Sz       CodingMTBMedicationRecommendationCategoryCode = "SZ"
	Zs       CodingMTBMedicationRecommendationCategoryCode = "ZS"
)

type AddendumCode string

const (
	Is AddendumCode = "is"
	Iv AddendumCode = "iv"
	R  AddendumCode = "R"
	Z  AddendumCode = "Z"
)

type LevelOfEvidenceCode string

const (
	CodeUndefined LevelOfEvidenceCode = "undefined"
	M1A           LevelOfEvidenceCode = "m1A"
	M1B           LevelOfEvidenceCode = "m1B"
	M1C           LevelOfEvidenceCode = "m1C"
	M2A           LevelOfEvidenceCode = "m2A"
	M2B           LevelOfEvidenceCode = "m2B"
	M2C           LevelOfEvidenceCode = "m2C"
	M3            LevelOfEvidenceCode = "m3"
	M4            LevelOfEvidenceCode = "m4"
)

type PublicationSystem string

const (
	HTTPSPubmedNcbiNlmNihGov PublicationSystem = "https://pubmed.ncbi.nlm.nih.gov"
	HTTPSWWWDoiOrg           PublicationSystem = "https://www.doi.org"
)

type RequestedMedicationSystem string

const (
	HTTPFhirDeCodeSystemBfarmAtc RequestedMedicationSystem = "http://fhir.de/CodeSystem/bfarm/atc"
	SystemUndefined              RequestedMedicationSystem = "undefined"
)

type PriorityCode string

const (
	Purple1 PriorityCode = "1"
	Purple2 PriorityCode = "2"
	Purple3 PriorityCode = "3"
	Purple4 PriorityCode = "4"
)

type UseTypeCode string

const (
	Compassionate UseTypeCode = "compassionate"
	FluffyUnknown UseTypeCode = "unknown"
	InLabel       UseTypeCode = "in-label"
	OffLabel      UseTypeCode = "off-label"
	SECPreventive UseTypeCode = "sec-preventive"
)

type CodingMTBProcedureRecommendationCategoryCode string

const (
	As       CodingMTBProcedureRecommendationCategoryCode = "AS"
	FluffySO CodingMTBProcedureRecommendationCategoryCode = "SO"
	Op       CodingMTBProcedureRecommendationCategoryCode = "OP"
	St       CodingMTBProcedureRecommendationCategoryCode = "ST"
	Ws       CodingMTBProcedureRecommendationCategoryCode = "WS"
	Ww       CodingMTBProcedureRecommendationCategoryCode = "WW"
)

type CodingMTBCarePlanStatusReasonCode string

const (
	FluffyOther                    CodingMTBCarePlanStatusReasonCode = "other"
	NoTarget                       CodingMTBCarePlanStatusReasonCode = "no-target"
	NonGeneticCause                CodingMTBCarePlanStatusReasonCode = "non-genetic-cause"
	NotRareDisease                 CodingMTBCarePlanStatusReasonCode = "not-rare-disease"
	Psychosomatic                  CodingMTBCarePlanStatusReasonCode = "psychosomatic"
	TargetedDiagnosticsRecommended CodingMTBCarePlanStatusReasonCode = "targeted-diagnostics-recommended"
)

type StudySystem string

const (
	Drks    StudySystem = "DRKS"
	Eudamed StudySystem = "EUDAMED"
	EudraCT StudySystem = "Eudra-CT"
	Nct     StudySystem = "NCT"
)

type CodingClaimResponseStatusCode string

const (
	Accepted         CodingClaimResponseStatusCode = "accepted"
	Rejected         CodingClaimResponseStatusCode = "rejected"
	TentacledUnknown CodingClaimResponseStatusCode = "unknown"
)

type CodingClaimResponseStatusReasonCode string

const (
	ApprovalRevocation          CodingClaimResponseStatusReasonCode = "approval-revocation"
	FormalReasons               CodingClaimResponseStatusReasonCode = "formal-reasons"
	InclusionInStudy            CodingClaimResponseStatusReasonCode = "inclusion-in-study"
	InsufficientEvidence        CodingClaimResponseStatusReasonCode = "insufficient-evidence"
	OtherTherapyRecommended     CodingClaimResponseStatusReasonCode = "other-therapy-recommended"
	StandardTherapyNotExhausted CodingClaimResponseStatusReasonCode = "standard-therapy-not-exhausted"
	StickyUnknown               CodingClaimResponseStatusReasonCode = "unknown"
	TentacledOther              CodingClaimResponseStatusReasonCode = "other"
)

type StageCode string

const (
	FollowUpClaim StageCode = "follow-up-claim"
	IndigoUnknown StageCode = "unknown"
	InitialClaim  StageCode = "initial-claim"
	Revocation    StageCode = "revocation"
)

type GuidelineTreatmentStatusCode string

const (
	Exhausted             GuidelineTreatmentStatusCode = "exhausted"
	Impossible            GuidelineTreatmentStatusCode = "impossible"
	IndecentUnknown       GuidelineTreatmentStatusCode = "unknown"
	NoGuidelinesAvailable GuidelineTreatmentStatusCode = "no-guidelines-available"
	NonExhausted          GuidelineTreatmentStatusCode = "non-exhausted"
)

type CodingTumorStagingMethodCode string

const (
	Clinical   CodingTumorStagingMethodCode = "clinical"
	Pathologic CodingTumorStagingMethodCode = "pathologic"
)

type CodingMTBDiagnosisCode string

const (
	Main         CodingMTBDiagnosisCode = "main"
	Metachronous CodingMTBDiagnosisCode = "metachronous"
	Secondary    CodingMTBDiagnosisCode = "secondary"
)

type PatientStatusCode string

const (
	PurpleLostToFu PatientStatusCode = "lost-to-fu"
)

type CodingOncoProcedureCode string

const (
	NuclearMedicine CodingOncoProcedureCode = "nuclear-medicine"
	RadioTherapy    CodingOncoProcedureCode = "radio-therapy"
	Surgery         CodingOncoProcedureCode = "surgery"
)

type IntentCode string

const (
	K       IntentCode = "K"
	P       IntentCode = "P"
	PurpleS IntentCode = "S"
	X       IntentCode = "X"
)

type CodingTherapyStatusCode string

const (
	Completed        CodingTherapyStatusCode = "completed"
	HilariousUnknown CodingTherapyStatusCode = "unknown"
	NotDone          CodingTherapyStatusCode = "not-done"
	OnGoing          CodingTherapyStatusCode = "on-going"
	Stopped          CodingTherapyStatusCode = "stopped"
)

type CodingMTBTherapyStatusReasonCode string

const (
	BestSupportiveCare                   CodingMTBTherapyStatusReasonCode = "best-supportive-care"
	ChronicRemission                     CodingMTBTherapyStatusReasonCode = "chronic-remission"
	Deterioration                        CodingMTBTherapyStatusReasonCode = "deterioration"
	FluffyLostToFu                       CodingMTBTherapyStatusReasonCode = "lost-to-fu"
	MedicalReasons                       CodingMTBTherapyStatusReasonCode = "medical-reasons"
	NoIndication                         CodingMTBTherapyStatusReasonCode = "no-indication"
	OtherTherapyChosen                   CodingMTBTherapyStatusReasonCode = "other-therapy-chosen"
	PatientDeath                         CodingMTBTherapyStatusReasonCode = "patient-death"
	PatientRefusal                       CodingMTBTherapyStatusReasonCode = "patient-refusal"
	PatientWish                          CodingMTBTherapyStatusReasonCode = "patient-wish"
	PaymentEnded                         CodingMTBTherapyStatusReasonCode = "payment-ended"
	PaymentPending                       CodingMTBTherapyStatusReasonCode = "payment-pending"
	PaymentRefused                       CodingMTBTherapyStatusReasonCode = "payment-refused"
	Progression                          CodingMTBTherapyStatusReasonCode = "progression"
	RegularCompletion                    CodingMTBTherapyStatusReasonCode = "regular-completion"
	RegularCompletionWithDosageReduction CodingMTBTherapyStatusReasonCode = "regular-completion-with-dosage-reduction"
	RegularCompletionWithSubstanceChange CodingMTBTherapyStatusReasonCode = "regular-completion-with-substance-change"
	StickyOther                          CodingMTBTherapyStatusReasonCode = "other"
	Toxicity                             CodingMTBTherapyStatusReasonCode = "toxicity"
)

type CodingMTBSystemicTherapyCategoryCode string

const (
	A       CodingMTBSystemicTherapyCategoryCode = "A"
	FluffyS CodingMTBSystemicTherapyCategoryCode = "S"
	I       CodingMTBSystemicTherapyCategoryCode = "I"
	N       CodingMTBSystemicTherapyCategoryCode = "N"
	O       CodingMTBSystemicTherapyCategoryCode = "O"
)

type RecommendationFulfillmentStatusCode string

const (
	Complete RecommendationFulfillmentStatusCode = "complete"
	Partial  RecommendationFulfillmentStatusCode = "partial"
)

type CodingTumorCellContentMethodCode string

const (
	Bioinformatic CodingTumorCellContentMethodCode = "bioinformatic"
	Histologic    CodingTumorCellContentMethodCode = "histologic"
)

type ICScoreCode string

const (
	Fluffy1 ICScoreCode = "1"
	Fluffy2 ICScoreCode = "2"
	Fluffy3 ICScoreCode = "3"
	Purple0 ICScoreCode = "0"
)

type TcScoreCode string

const (
	Fluffy0    TcScoreCode = "0"
	Fluffy4    TcScoreCode = "4"
	Purple5    TcScoreCode = "5"
	Tentacled1 TcScoreCode = "1"
	Tentacled2 TcScoreCode = "2"
	Tentacled3 TcScoreCode = "3"
	The6       TcScoreCode = "6"
)

type CodingProteinExpressionResultCode string

const (
	AmbitiousUnknown CodingProteinExpressionResultCode = "unknown"
	Exp              CodingProteinExpressionResultCode = "exp"
	NotExp           CodingProteinExpressionResultCode = "not-exp"
	The1             CodingProteinExpressionResultCode = "1+"
	The2             CodingProteinExpressionResultCode = "2+"
	The3             CodingProteinExpressionResultCode = "3+"
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

type ExternalIDSystem string

const (
	HTTPSCancerSangerACUkCosmic ExternalIDSystem = "https://cancer.sanger.ac.uk/cosmic"
	HTTPSWWWNcbiNlmNihGovEntrez ExternalIDSystem = "https://www.ncbi.nlm.nih.gov/entrez"
	HTTPSWWWNcbiNlmNihGovSnp    ExternalIDSystem = "https://www.ncbi.nlm.nih.gov/snp"
	PurpleHTTPSWWWEnsemblOrg    ExternalIDSystem = "https://www.ensembl.org"
)

type CodingBaseVariantLocalizationCode string

const (
	CodingRegion     CodingBaseVariantLocalizationCode = "coding-region"
	Intergenic       CodingBaseVariantLocalizationCode = "intergenic"
	Intronic         CodingBaseVariantLocalizationCode = "intronic"
	RegulatoryRegion CodingBaseVariantLocalizationCode = "regulatory-region"
	SplicingRegion   CodingBaseVariantLocalizationCode = "splicing-region"
)

type CodingCNVCode string

const (
	HighLevelGain CodingCNVCode = "high-level-gain"
	Loss          CodingCNVCode = "loss"
	LowLevelGain  CodingCNVCode = "low-level-gain"
)

type InterpretationCode string

const (
	High         InterpretationCode = "high"
	Intermediate InterpretationCode = "intermediate"
	Low          InterpretationCode = "low"
)

type StrandEnum string

const (
	Empty           StrandEnum = "+"
	RNAFusionStrand StrandEnum = "-"
)

type TranscriptIDSystem string

const (
	FluffyHTTPSWWWEnsemblOrg    TranscriptIDSystem = "https://www.ensembl.org"
	HTTPSWWWNcbiNlmNihGovRefseq TranscriptIDSystem = "https://www.ncbi.nlm.nih.gov/refseq"
)

type CodingClinVarCode string

const (
	Fluffy5    CodingClinVarCode = "5"
	Sticky1    CodingClinVarCode = "1"
	Sticky2    CodingClinVarCode = "2"
	Sticky3    CodingClinVarCode = "3"
	Tentacled4 CodingClinVarCode = "4"
)

type CodingNGSReportCode string

const (
	IndigoOther           CodingNGSReportCode = "other"
	PurpleArray           CodingNGSReportCode = "array"
	PurpleExome           CodingNGSReportCode = "exome"
	PurpleGenomeLongRead  CodingNGSReportCode = "genome-long-read"
	PurpleGenomeShortRead CodingNGSReportCode = "genome-short-read"
	PurpleKaryotyping     CodingNGSReportCode = "karyotyping"
	PurplePanel           CodingNGSReportCode = "panel"
	PurpleSingle          CodingNGSReportCode = "single"
)

type Unit string

const (
	Months Unit = "Months"
	Years  Unit = "Years"
)

type GenderCode string

const (
	CunningUnknown GenderCode = "unknown"
	Female         GenderCode = "female"
	IndecentOther  GenderCode = "other"
	Male           GenderCode = "male"
)

type CodingHealthInsuranceCode string

const (
	Bei CodingHealthInsuranceCode = "BEI"
	Bg  CodingHealthInsuranceCode = "BG"
	Gkv CodingHealthInsuranceCode = "GKV"
	Gpv CodingHealthInsuranceCode = "GPV"
	Pkv CodingHealthInsuranceCode = "PKV"
	Ppv CodingHealthInsuranceCode = "PPV"
	Sel CodingHealthInsuranceCode = "SEL"
	Soz CodingHealthInsuranceCode = "SOZ"
	Unk CodingHealthInsuranceCode = "UNK"
)

type VitalStatusCode string

const (
	Alive    VitalStatusCode = "alive"
	Deceased VitalStatusCode = "deceased"
)

type CodingECOGCode string

const (
	Indigo1    CodingECOGCode = "1"
	Indigo2    CodingECOGCode = "2"
	Indigo3    CodingECOGCode = "3"
	Sticky4    CodingECOGCode = "4"
	Tentacled0 CodingECOGCode = "0"
	Tentacled5 CodingECOGCode = "5"
)

type CodingMolecularDiagnosticReportCode string

const (
	Fish                  CodingMolecularDiagnosticReportCode = "FISH"
	FluffyArray           CodingMolecularDiagnosticReportCode = "array"
	FluffyExome           CodingMolecularDiagnosticReportCode = "exome"
	FluffyGenomeLongRead  CodingMolecularDiagnosticReportCode = "genome-long-read"
	FluffyGenomeShortRead CodingMolecularDiagnosticReportCode = "genome-short-read"
	FluffyKaryotyping     CodingMolecularDiagnosticReportCode = "karyotyping"
	FluffyPanel           CodingMolecularDiagnosticReportCode = "panel"
	FluffySingle          CodingMolecularDiagnosticReportCode = "single"
	FusionPanel           CodingMolecularDiagnosticReportCode = "fusion-panel"
	GenePanel             CodingMolecularDiagnosticReportCode = "gene-panel"
	HilariousOther        CodingMolecularDiagnosticReportCode = "other"
	Pcr                   CodingMolecularDiagnosticReportCode = "PCR"
)

type CodingResponseMethodCode string

const (
	Rano   CodingResponseMethodCode = "RANO"
	Recist CodingResponseMethodCode = "RECIST"
)

type CodingRECISTCode string

const (
	CR CodingRECISTCode = "CR"
	Mr CodingRECISTCode = "MR"
	Na CodingRECISTCode = "NA"
	PD CodingRECISTCode = "PD"
	PR CodingRECISTCode = "PR"
	SD CodingRECISTCode = "SD"
)

type CodingTumorSpecimenCollectionLocalizationCode string

const (
	CellfreeDna        CodingTumorSpecimenCollectionLocalizationCode = "cellfree-dna"
	LocalRecurrence    CodingTumorSpecimenCollectionLocalizationCode = "local-recurrence"
	MagentaUnknown     CodingTumorSpecimenCollectionLocalizationCode = "unknown"
	Metastasis         CodingTumorSpecimenCollectionLocalizationCode = "metastasis"
	PrimaryTumor       CodingTumorSpecimenCollectionLocalizationCode = "primary-tumor"
	RegionalLymphNodes CodingTumorSpecimenCollectionLocalizationCode = "regional-lymph-nodes"
)

type CodingTumorSpecimenCollectionMethodCode string

const (
	Biopsy             CodingTumorSpecimenCollectionMethodCode = "biopsy"
	Cytology           CodingTumorSpecimenCollectionMethodCode = "cytology"
	FriskyUnknown      CodingTumorSpecimenCollectionMethodCode = "unknown"
	PurpleLiquidBiopsy CodingTumorSpecimenCollectionMethodCode = "liquid-biopsy"
	Resection          CodingTumorSpecimenCollectionMethodCode = "resection"
)

type CodingTumorSpecimenCode string

const (
	CryoFrozen         CodingTumorSpecimenCode = "cryo-frozen"
	Ffpe               CodingTumorSpecimenCode = "FFPE"
	FluffyLiquidBiopsy CodingTumorSpecimenCode = "liquid-biopsy"
	FreshTissue        CodingTumorSpecimenCode = "fresh-tissue"
	MischievousUnknown CodingTumorSpecimenCode = "unknown"
)
