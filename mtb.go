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
	CarePlans              []MtbCarePlan           `json:"carePlans,omitempty"`
	ClaimResponses         []ClaimResponse         `json:"claimResponses,omitempty"`
	Claims                 []Claim                 `json:"claims,omitempty"`
	Diagnoses              []MtbDiagnosis          `json:"diagnoses"`
	EpisodesOfCare         []MtbEpisodeOfCare      `json:"episodesOfCare"`
	FamilyMemberHistories  []FamilyMemberHistory   `json:"familyMemberHistories,omitempty"`
	FollowUPS              []FollowUp              `json:"followUps,omitempty"`
	GuidelineProcedures    []OncoProcedure         `json:"guidelineProcedures,omitempty"`
	GuidelineTherapies     []MtbSystemicTherapy    `json:"guidelineTherapies,omitempty"`
	HistologyReports       []HistologyReport       `json:"histologyReports,omitempty"`
	IhcReports             []IhcReport             `json:"ihcReports,omitempty"`
	MsiFindings            []Msi                   `json:"msiFindings,omitempty"`
	Metadata               *MvhMetadata            `json:"metadata,omitempty"`
	NgsReports             []SomaticNgsReport      `json:"ngsReports,omitempty"`
	Patient                Patient                 `json:"patient"`
	PerformanceStatus      []PerformanceStatus     `json:"performanceStatus,omitempty"`
	PriorDiagnosticReports []PriorDiagnosticReport `json:"priorDiagnosticReports,omitempty"`
	Responses              []Response              `json:"responses,omitempty"`
	Specimens              []TumorSpecimen         `json:"specimens,omitempty"`
	SystemicTherapies      []SystemicTherapy       `json:"systemicTherapies,omitempty"`
}

type MtbCarePlan struct {
	GeneticCounselingRecommendation *GeneticCounselingRecommendation               `json:"geneticCounselingRecommendation,omitempty"`
	HistologyReevaluationRequests   []HistologyReevaluationRequest                 `json:"histologyReevaluationRequests,omitempty"`
	ID                              string                                         `json:"id"`
	IssuedOn                        string                                         `json:"issuedOn"`
	MedicationRecommendations       []MtbMedicationRecommendation                  `json:"medicationRecommendations,omitempty"`
	NoSequencingPerformedReason     *CarePlanNoSequencingPerformedReasonCoding     `json:"noSequencingPerformedReason,omitempty"`
	Notes                           []string                                       `json:"notes,omitempty"`
	Patient                         Reference                                      `json:"patient"`
	ProcedureRecommendations        []ProcedureRecommendation                      `json:"procedureRecommendations,omitempty"`
	Reason                          *Reference                                     `json:"reason,omitempty"`
	RebiopsyRequests                []RebiopsyRequest                              `json:"rebiopsyRequests,omitempty"`
	RecommendationsMissingReason    *MtbCarePlanRecommendationsMissingReasonCoding `json:"recommendationsMissingReason,omitempty"`
	StudyEnrollmentRecommendations  []MtbStudyEnrollmentRecommendation             `json:"studyEnrollmentRecommendations,omitempty"`
}

type GeneticCounselingRecommendation struct {
	ID       string                                      `json:"id"`
	IssuedOn string                                      `json:"issuedOn"`
	Patient  Reference                                   `json:"patient"`
	Reason   GeneticCounselingRecommendationReasonCoding `json:"reason"`
}

type Reference struct {
	Display *string `json:"display,omitempty"`
	ID      string  `json:"id"`
	System  *string `json:"system,omitempty"`
	Type    *string `json:"type,omitempty"`
}

type GeneticCounselingRecommendationReasonCoding struct {
	Code    GeneticCounselingRecommendationReasonCodingCode `json:"code"`
	Display *string                                         `json:"display,omitempty"`
	System  *string                                         `json:"system,omitempty"`
	Version *string                                         `json:"version,omitempty"`
}

type HistologyReevaluationRequest struct {
	ID       string    `json:"id"`
	IssuedOn string    `json:"issuedOn"`
	Patient  Reference `json:"patient"`
	Specimen Reference `json:"specimen"`
}

type MtbMedicationRecommendation struct {
	Category           []MtbMedicationRecommendationCategoryCoding `json:"category,omitempty"`
	ID                 string                                      `json:"id"`
	IssuedOn           string                                      `json:"issuedOn"`
	LevelOfEvidence    *LevelOfEvidence                            `json:"levelOfEvidence,omitempty"`
	Medication         []AtcUnregisteredMedicationCoding           `json:"medication"`
	Patient            Reference                                   `json:"patient"`
	Priority           RecommendationPriorityCoding                `json:"priority"`
	Reason             *Reference                                  `json:"reason,omitempty"`
	SupportingVariants []GeneAlterationReference                   `json:"supportingVariants,omitempty"`
	UseType            *MtbMedicationRecommendationUseTypeCoding   `json:"useType,omitempty"`
}

type MtbMedicationRecommendationCategoryCoding struct {
	Code    MtbMedicationRecommendationCategoryCodingCode `json:"code"`
	Display *string                                       `json:"display,omitempty"`
	System  *string                                       `json:"system,omitempty"`
	Version *string                                       `json:"version,omitempty"`
}

type LevelOfEvidence struct {
	Addendums    []LevelOfEvidenceAddendumCoding `json:"addendums,omitempty"`
	Grading      LevelOfEvidenceGradingCoding    `json:"grading"`
	Publications []PublicationReference          `json:"publications,omitempty"`
}

type LevelOfEvidenceAddendumCoding struct {
	Code    LevelOfEvidenceAddendumCodingCode `json:"code"`
	Display *string                           `json:"display,omitempty"`
	System  *string                           `json:"system,omitempty"`
	Version *string                           `json:"version,omitempty"`
}

type LevelOfEvidenceGradingCoding struct {
	Code    LevelOfEvidenceGradingCodingCode `json:"code"`
	Display *string                          `json:"display,omitempty"`
	System  *string                          `json:"system,omitempty"`
	Version *string                          `json:"version,omitempty"`
}

type PublicationReference struct {
	Display *string           `json:"display,omitempty"`
	ID      string            `json:"id"`
	System  PublicationSystem `json:"system"`
	Type    *string           `json:"type,omitempty"`
}

type AtcUnregisteredMedicationCoding struct {
	Code    string                    `json:"code"`
	Display *string                   `json:"display,omitempty"`
	System  RequestedMedicationSystem `json:"system"`
	Version *string                   `json:"version,omitempty"`
}

type RecommendationPriorityCoding struct {
	Code    RecommendationPriorityCodingCode `json:"code"`
	Display *string                          `json:"display,omitempty"`
	System  *string                          `json:"system,omitempty"`
	Version *string                          `json:"version,omitempty"`
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

type MtbMedicationRecommendationUseTypeCoding struct {
	Code    MtbMedicationRecommendationUseTypeCodingCode `json:"code"`
	Display *string                                      `json:"display,omitempty"`
	System  *string                                      `json:"system,omitempty"`
	Version *string                                      `json:"version,omitempty"`
}

type CarePlanNoSequencingPerformedReasonCoding struct {
	Code    NoSequencingPerformedReasonCode `json:"code"`
	Display *string                         `json:"display,omitempty"`
	System  *string                         `json:"system,omitempty"`
	Version *string                         `json:"version,omitempty"`
}

type ProcedureRecommendation struct {
	Code               MtbProcedureRecommendationCategoryCoding `json:"code"`
	ID                 string                                   `json:"id"`
	IssuedOn           string                                   `json:"issuedOn"`
	LevelOfEvidence    *LevelOfEvidence                         `json:"levelOfEvidence,omitempty"`
	Patient            Reference                                `json:"patient"`
	Priority           RecommendationPriorityCoding             `json:"priority"`
	Reason             *Reference                               `json:"reason,omitempty"`
	SupportingVariants []GeneAlterationReference                `json:"supportingVariants,omitempty"`
}

type MtbProcedureRecommendationCategoryCoding struct {
	Code    MtbProcedureRecommendationCategoryCodingCode `json:"code"`
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

type MtbCarePlanRecommendationsMissingReasonCoding struct {
	Code    MtbCarePlanRecommendationsMissingReasonCodingCode `json:"code"`
	Display *string                                           `json:"display,omitempty"`
	System  *string                                           `json:"system,omitempty"`
	Version *string                                           `json:"version,omitempty"`
}

type MtbStudyEnrollmentRecommendation struct {
	ID                 string                            `json:"id"`
	IssuedOn           string                            `json:"issuedOn"`
	LevelOfEvidence    *LevelOfEvidence                  `json:"levelOfEvidence,omitempty"`
	Medication         []AtcUnregisteredMedicationCoding `json:"medication,omitempty"`
	Patient            Reference                         `json:"patient"`
	Priority           RecommendationPriorityCoding      `json:"priority"`
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
	Claim        Reference                         `json:"claim"`
	ID           string                            `json:"id"`
	IssuedOn     string                            `json:"issuedOn"`
	Patient      Reference                         `json:"patient"`
	Status       *ClaimResponseStatusCoding        `json:"status,omitempty"`
	StatusReason []ClaimResponseStatusReasonCoding `json:"statusReason,omitempty"`
}

type ClaimResponseStatusCoding struct {
	Code    ClaimResponseStatusCodingCode `json:"code"`
	Display *string                       `json:"display,omitempty"`
	System  *string                       `json:"system,omitempty"`
	Version *string                       `json:"version,omitempty"`
}

type ClaimResponseStatusReasonCoding struct {
	Code    ClaimResponseStatusReasonCodingCode `json:"code"`
	Display *string                             `json:"display,omitempty"`
	System  *string                             `json:"system,omitempty"`
	Version *string                             `json:"version,omitempty"`
}

type Claim struct {
	ID                  string                            `json:"id"`
	IssuedOn            string                            `json:"issuedOn"`
	Patient             Reference                         `json:"patient"`
	Recommendation      Reference                         `json:"recommendation"`
	RequestedMedication []AtcUnregisteredMedicationCoding `json:"requestedMedication,omitempty"`
	Stage               *ClaimStageCoding                 `json:"stage,omitempty"`
}

type ClaimStageCoding struct {
	Code    ClaimStageCodingCode `json:"code"`
	Display *string              `json:"display,omitempty"`
	System  *string              `json:"system,omitempty"`
	Version *string              `json:"version,omitempty"`
}

type MtbDiagnosis struct {
	Code                     Coding                                      `json:"code"`
	GermlineCodes            []Coding                                    `json:"germlineCodes,omitempty"`
	Grading                  *Grading                                    `json:"grading,omitempty"`
	GuidelineTreatmentStatus *MtbDiagnosisGuidelineTreatmentStatusCoding `json:"guidelineTreatmentStatus,omitempty"`
	Histology                []Reference                                 `json:"histology,omitempty"`
	ID                       string                                      `json:"id"`
	Notes                    []string                                    `json:"notes,omitempty"`
	Patient                  Reference                                   `json:"patient"`
	RecordedOn               string                                      `json:"recordedOn"`
	Staging                  *Staging                                    `json:"staging,omitempty"`
	Topography               Coding                                      `json:"topography"`
	Type                     Type                                        `json:"type"`
}

type Grading struct {
	History []TumorGrading `json:"history"`
}

type TumorGrading struct {
	Codes []Coding `json:"codes"`
	Date  string   `json:"date"`
}

type MtbDiagnosisGuidelineTreatmentStatusCoding struct {
	Code    MtbDiagnosisGuidelineTreatmentStatusCodingCode `json:"code"`
	Display *string                                        `json:"display,omitempty"`
	System  *string                                        `json:"system,omitempty"`
	Version *string                                        `json:"version,omitempty"`
}

type Staging struct {
	History []TumorStaging `json:"history"`
}

type TumorStaging struct {
	Date                 string                   `json:"date"`
	Method               TumorStagingMethodCoding `json:"method"`
	OtherClassifications []Coding                 `json:"otherClassifications,omitempty"`
	TnmClassification    *TnmClassification       `json:"tnmClassification,omitempty"`
}

type TumorStagingMethodCoding struct {
	Code    TumorStagingMethodCodingCode `json:"code"`
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
	Value MtbDiagnosisCoding `json:"value"`
}

type MtbDiagnosisCoding struct {
	Code    ValueCode `json:"code"`
	Display *string   `json:"display,omitempty"`
	System  *string   `json:"system,omitempty"`
	Version *string   `json:"version,omitempty"`
}

type MtbEpisodeOfCare struct {
	Diagnoses []Reference `json:"diagnoses,omitempty"`
	ID        string      `json:"id"`
	Patient   Reference   `json:"patient"`
	Period    PeriodDate  `json:"period"`
}

type PeriodDate struct {
	End   *string `json:"end,omitempty"`
	Start string  `json:"start"`
}

type FamilyMemberHistory struct {
	ID           string                                    `json:"id"`
	Patient      Reference                                 `json:"patient"`
	Relationship FamilyMemberHistoryRelationshipTypeCoding `json:"relationship"`
}

type FamilyMemberHistoryRelationshipTypeCoding struct {
	Code    FamilyMemberHistoryRelationshipTypeCodingCode `json:"code"`
	Display *string                                       `json:"display,omitempty"`
	System  *string                                       `json:"system,omitempty"`
	Version *string                                       `json:"version,omitempty"`
}

type FollowUp struct {
	Date            string                       `json:"date"`
	LastContactDate *string                      `json:"lastContactDate,omitempty"`
	Patient         Reference                    `json:"patient"`
	PatientStatus   *FollowUpPatientStatusCoding `json:"patientStatus,omitempty"`
}

type FollowUpPatientStatusCoding struct {
	Code    FollowUpPatientStatusCodingCode `json:"code"`
	Display *string                         `json:"display,omitempty"`
	System  *string                         `json:"system,omitempty"`
	Version *string                         `json:"version,omitempty"`
}

type OncoProcedure struct {
	BasedOn      *Reference                    `json:"basedOn,omitempty"`
	Code         OncoProcedureCoding           `json:"code"`
	ID           string                        `json:"id"`
	Intent       *MtbTherapyIntentCoding       `json:"intent,omitempty"`
	Notes        []string                      `json:"notes,omitempty"`
	Patient      Reference                     `json:"patient"`
	Period       *PeriodDate                   `json:"period,omitempty"`
	Reason       *Reference                    `json:"reason,omitempty"`
	RecordedOn   string                        `json:"recordedOn"`
	Status       TherapyStatusCoding           `json:"status"`
	StatusReason *MtbTherapyStatusReasonCoding `json:"statusReason,omitempty"`
	TherapyLine  *int64                        `json:"therapyLine,omitempty"`
}

type OncoProcedureCoding struct {
	Code    OncoProcedureCodingCode `json:"code"`
	Display *string                 `json:"display,omitempty"`
	System  *string                 `json:"system,omitempty"`
	Version *string                 `json:"version,omitempty"`
}

type MtbTherapyIntentCoding struct {
	Code    MtbTherapyIntentCodingCode `json:"code"`
	Display *string                    `json:"display,omitempty"`
	System  *string                    `json:"system,omitempty"`
	Version *string                    `json:"version,omitempty"`
}

type TherapyStatusCoding struct {
	Code    TherapyStatusCodingCode `json:"code"`
	Display *string                 `json:"display,omitempty"`
	System  *string                 `json:"system,omitempty"`
	Version *string                 `json:"version,omitempty"`
}

type MtbTherapyStatusReasonCoding struct {
	Code    MtbTherapyStatusReasonCodingCode `json:"code"`
	Display *string                          `json:"display,omitempty"`
	System  *string                          `json:"system,omitempty"`
	Version *string                          `json:"version,omitempty"`
}

type MtbSystemicTherapy struct {
	BasedOn                         *Reference                                               `json:"basedOn,omitempty"`
	Category                        *MtbSystemicTherapyCategoryCoding                        `json:"category,omitempty"`
	Dosage                          *MtbSystemicTherapyDosageDensityCoding                   `json:"dosage,omitempty"`
	ID                              string                                                   `json:"id"`
	Intent                          *MtbTherapyIntentCoding                                  `json:"intent,omitempty"`
	Medication                      []AtcUnregisteredMedicationCoding                        `json:"medication,omitempty"`
	Notes                           []string                                                 `json:"notes,omitempty"`
	Patient                         Reference                                                `json:"patient"`
	Period                          *PeriodDate                                              `json:"period,omitempty"`
	Reason                          *Reference                                               `json:"reason,omitempty"`
	RecommendationFulfillmentStatus *MtbSystemicTherapyRecommendationFulfillmentStatusCoding `json:"recommendationFulfillmentStatus,omitempty"`
	RecordedOn                      string                                                   `json:"recordedOn"`
	Status                          TherapyStatusCoding                                      `json:"status"`
	StatusReason                    *MtbTherapyStatusReasonCoding                            `json:"statusReason,omitempty"`
	TherapyLine                     *int64                                                   `json:"therapyLine,omitempty"`
}

type MtbSystemicTherapyCategoryCoding struct {
	Code    MtbSystemicTherapyCategoryCodingCode `json:"code"`
	Display *string                              `json:"display,omitempty"`
	System  *string                              `json:"system,omitempty"`
	Version *string                              `json:"version,omitempty"`
}

type MtbSystemicTherapyDosageDensityCoding struct {
	Code    MtbSystemicTherapyDosageDensityCodingCode `json:"code"`
	Display *string                                   `json:"display,omitempty"`
	System  *string                                   `json:"system,omitempty"`
	Version *string                                   `json:"version,omitempty"`
}

type MtbSystemicTherapyRecommendationFulfillmentStatusCoding struct {
	Code    MtbSystemicTherapyRecommendationFulfillmentStatusCodingCode `json:"code"`
	Display *string                                                     `json:"display,omitempty"`
	System  *string                                                     `json:"system,omitempty"`
	Version *string                                                     `json:"version,omitempty"`
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
	Method   TumorCellContentMethodCoding `json:"method"`
	Patient  Reference                    `json:"patient"`
	Specimen Reference                    `json:"specimen"`
	Value    float64                      `json:"value"`
}

type TumorCellContentMethodCoding struct {
	Code    TumorCellContentMethodCodingCode `json:"code"`
	Display *string                          `json:"display,omitempty"`
	System  *string                          `json:"system,omitempty"`
	Version *string                          `json:"version,omitempty"`
}

type TumorMorphology struct {
	ID       string    `json:"id"`
	Note     *string   `json:"note,omitempty"`
	Patient  Reference `json:"patient"`
	Specimen Reference `json:"specimen"`
	Value    Coding    `json:"value"`
}

type IhcReport struct {
	ID       string           `json:"id"`
	IssuedOn string           `json:"issuedOn"`
	Patient  Reference        `json:"patient"`
	Results  IhcReportResults `json:"results"`
	Specimen Reference        `json:"specimen"`
}

type IhcReportResults struct {
	MsiMmr            []MsiMmr            `json:"msiMmr"`
	ProteinExpression []ProteinExpression `json:"proteinExpression"`
}

type MsiMmr struct {
	ICScore  *ProteinExpressionICScoreCoding `json:"icScore,omitempty"`
	ID       string                          `json:"id"`
	Patient  Reference                       `json:"patient"`
	Protein  Coding                          `json:"protein"`
	TcScore  *ProteinExpressionTcScoreCoding `json:"tcScore,omitempty"`
	TpsScore *int64                          `json:"tpsScore,omitempty"`
	Value    ProteinExpressionResultCoding   `json:"value"`
}

type ProteinExpressionICScoreCoding struct {
	Code    ProteinExpressionICScoreCodingCode `json:"code"`
	Display *string                            `json:"display,omitempty"`
	System  *string                            `json:"system,omitempty"`
	Version *string                            `json:"version,omitempty"`
}

type ProteinExpressionTcScoreCoding struct {
	Code    ProteinExpressionTcScoreCodingCode `json:"code"`
	Display *string                            `json:"display,omitempty"`
	System  *string                            `json:"system,omitempty"`
	Version *string                            `json:"version,omitempty"`
}

type ProteinExpressionResultCoding struct {
	Code    ProteinExpressionResultCodingCode `json:"code"`
	Display *string                           `json:"display,omitempty"`
	System  *string                           `json:"system,omitempty"`
	Version *string                           `json:"version,omitempty"`
}

type ProteinExpression struct {
	ICScore  *ProteinExpressionICScoreCoding `json:"icScore,omitempty"`
	ID       string                          `json:"id"`
	Patient  Reference                       `json:"patient"`
	Protein  Coding                          `json:"protein"`
	TcScore  *ProteinExpressionTcScoreCoding `json:"tcScore,omitempty"`
	TpsScore *int64                          `json:"tpsScore,omitempty"`
	Value    ProteinExpressionResultCoding   `json:"value"`
}

type Msi struct {
	ID             string                  `json:"id"`
	Interpretation MsiInterpretationCoding `json:"interpretation"`
	Method         MsiMethodCoding         `json:"method"`
	Patient        Reference               `json:"patient"`
	Specimen       Reference               `json:"specimen"`
	Value          float64                 `json:"value"`
}

type MsiInterpretationCoding struct {
	Code    MsiInterpretationCodingCode `json:"code"`
	Display *string                     `json:"display,omitempty"`
	System  *string                     `json:"system,omitempty"`
	Version *string                     `json:"version,omitempty"`
}

type MsiMethodCoding struct {
	Code    MsiMethodCodingCode `json:"code"`
	Display *string             `json:"display,omitempty"`
	System  *string             `json:"system,omitempty"`
	Version *string             `json:"version,omitempty"`
}

type MvhMetadata struct {
	ModelProjectConsent ModelProjectConsent      `json:"modelProjectConsent"`
	ResearchConsents    []map[string]interface{} `json:"researchConsents,omitempty"`
	TransferTAN         string                   `json:"transferTAN"`
	Type                MvhSubmissionType        `json:"type"`
}

type ModelProjectConsent struct {
	Date       *string     `json:"date,omitempty"`
	Provisions []Provision `json:"provisions"`
	Version    string      `json:"version"`
}

type Provision struct {
	Date    string                     `json:"date"`
	Purpose ModelProjectConsentPurpose `json:"purpose"`
	Type    ConsentProvision           `json:"type"`
}

type SomaticNgsReport struct {
	ID       string              `json:"id"`
	IssuedOn string              `json:"issuedOn"`
	Metadata []NgsReportMetadata `json:"metadata"`
	Patient  Reference           `json:"patient"`
	Results  NgsReportResults    `json:"results"`
	Specimen Reference           `json:"specimen"`
	Type     NgsReportCoding     `json:"type"`
}

type NgsReportMetadata struct {
	KitManufacturer string `json:"kitManufacturer"`
	KitType         string `json:"kitType"`
	Pipeline        string `json:"pipeline"`
	ReferenceGenome string `json:"referenceGenome"`
	Sequencer       string `json:"sequencer"`
}

type NgsReportResults struct {
	Brcaness           *Brcaness         `json:"brcaness,omitempty"`
	CopyNumberVariants []Cnv             `json:"copyNumberVariants,omitempty"`
	DnaFusions         []DnaFusion       `json:"dnaFusions,omitempty"`
	HrdScore           *HrdScore         `json:"hrdScore,omitempty"`
	RnaFusions         []RnaFusion       `json:"rnaFusions,omitempty"`
	RnaSeqs            []RnaSeq          `json:"rnaSeqs,omitempty"`
	SimpleVariants     []Snv             `json:"simpleVariants,omitempty"`
	Tmb                *Tmb              `json:"tmb,omitempty"`
	TumorCellContent   *TumorCellContent `json:"tumorCellContent,omitempty"`
}

type Brcaness struct {
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
	Localization          []BaseVariantLocalizationCoding `json:"localization,omitempty"`
	Patient               Reference                       `json:"patient"`
	RelativeCopyNumber    *float64                        `json:"relativeCopyNumber,omitempty"`
	ReportedAffectedGenes []Coding                        `json:"reportedAffectedGenes,omitempty"`
	ReportedFocality      *string                         `json:"reportedFocality,omitempty"`
	StartRange            *StartRange                     `json:"startRange,omitempty"`
	TotalCopyNumber       *int64                          `json:"totalCopyNumber,omitempty"`
	Type                  CnvCoding                       `json:"type"`
}

type EndRange struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type VariantExternalID struct {
	System ExternalIDSystem `json:"system"`
	Value  string           `json:"value"`
}

type BaseVariantLocalizationCoding struct {
	Code    BaseVariantLocalizationCodingCode `json:"code"`
	Display *string                           `json:"display,omitempty"`
	System  *string                           `json:"system,omitempty"`
	Version *string                           `json:"version,omitempty"`
}

type StartRange struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type CnvCoding struct {
	Code    CnvCodingCode `json:"code"`
	Display *string       `json:"display,omitempty"`
	System  *string       `json:"system,omitempty"`
	Version *string       `json:"version,omitempty"`
}

type DnaFusion struct {
	ExternalIDS         []VariantExternalID             `json:"externalIds,omitempty"`
	FusionPartner3Prime DnaFusionFusionPartner3Prime    `json:"fusionPartner3prime"`
	FusionPartner5Prime DnaFusionFusionPartner5Prime    `json:"fusionPartner5prime"`
	ID                  string                          `json:"id"`
	Localization        []BaseVariantLocalizationCoding `json:"localization,omitempty"`
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

type HrdScore struct {
	Components     Components                    `json:"components"`
	ID             string                        `json:"id"`
	Interpretation *HrdScoreInterpretationCoding `json:"interpretation,omitempty"`
	Patient        Reference                     `json:"patient"`
	Specimen       Reference                     `json:"specimen"`
	Value          float64                       `json:"value"`
}

type Components struct {
	Loh float64 `json:"loh"`
	Lst float64 `json:"lst"`
	Tai float64 `json:"tai"`
}

type HrdScoreInterpretationCoding struct {
	Code    InterpretationCodingCode `json:"code"`
	Display *string                  `json:"display,omitempty"`
	System  *string                  `json:"system,omitempty"`
	Version *string                  `json:"version,omitempty"`
}

type RnaFusion struct {
	Effect              *string                         `json:"effect,omitempty"`
	ExternalIDS         []VariantExternalID             `json:"externalIds,omitempty"`
	FusionPartner3Prime RnaFusionFusionPartner3Prime    `json:"fusionPartner3prime"`
	FusionPartner5Prime RnaFusionFusionPartner5Prime    `json:"fusionPartner5prime"`
	ID                  string                          `json:"id"`
	Localization        []BaseVariantLocalizationCoding `json:"localization,omitempty"`
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

type RnaSeq struct {
	CohortRanking             *int64                          `json:"cohortRanking,omitempty"`
	ExternalIDS               []VariantExternalID             `json:"externalIds,omitempty"`
	Gene                      *Coding                         `json:"gene,omitempty"`
	ID                        string                          `json:"id"`
	LibrarySize               *int64                          `json:"librarySize,omitempty"`
	Localization              []BaseVariantLocalizationCoding `json:"localization,omitempty"`
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
	DnaChange        string                          `json:"dnaChange"`
	ExonID           *string                         `json:"exonId,omitempty"`
	ExternalIDS      []VariantExternalID             `json:"externalIds,omitempty"`
	Gene             Coding                          `json:"gene"`
	ID               string                          `json:"id"`
	Interpretation   *ClinVarCoding                  `json:"interpretation,omitempty"`
	Localization     []BaseVariantLocalizationCoding `json:"localization,omitempty"`
	Patient          Reference                       `json:"patient"`
	Position         Position                        `json:"position"`
	ProteinChange    *string                         `json:"proteinChange,omitempty"`
	ReadDepth        int64                           `json:"readDepth"`
	RefAllele        string                          `json:"refAllele"`
	TranscriptID     TranscriptID                    `json:"transcriptId"`
}

type ClinVarCoding struct {
	Code    ClinVarCodingCode `json:"code"`
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
	Interpretation *TmbInterpretationCoding `json:"interpretation,omitempty"`
	Patient        Reference                `json:"patient"`
	Specimen       Reference                `json:"specimen"`
	Value          TmbResult                `json:"value"`
}

type TmbInterpretationCoding struct {
	Code    InterpretationCodingCode `json:"code"`
	Display *string                  `json:"display,omitempty"`
	System  *string                  `json:"system,omitempty"`
	Version *string                  `json:"version,omitempty"`
}

type TmbResult struct {
	Unit  *string `json:"unit,omitempty"`
	Value float64 `json:"value"`
}

type NgsReportCoding struct {
	Code    NgsReportCodingCode `json:"code"`
	Display *string             `json:"display,omitempty"`
	System  *string             `json:"system,omitempty"`
	Version *string             `json:"version,omitempty"`
}

type Patient struct {
	Address         Address            `json:"address"`
	Age             *Age               `json:"age,omitempty"`
	BirthDate       string             `json:"birthDate"`
	DateOfDeath     *string            `json:"dateOfDeath,omitempty"`
	Gender          GenderCoding       `json:"gender"`
	HealthInsurance HealthInsurance    `json:"healthInsurance"`
	ID              string             `json:"id"`
	ManagingSite    *Coding            `json:"managingSite,omitempty"`
	VitalStatus     *VitalStatusCoding `json:"vitalStatus,omitempty"`
}

type Address struct {
	MunicipalityCode string `json:"municipalityCode"`
}

type Age struct {
	Unit  Unit    `json:"unit"`
	Value float64 `json:"value"`
}

type GenderCoding struct {
	Code    GenderCodingCode `json:"code"`
	Display *string          `json:"display,omitempty"`
	System  *string          `json:"system,omitempty"`
	Version *string          `json:"version,omitempty"`
}

type HealthInsurance struct {
	Reference *Reference            `json:"reference,omitempty"`
	Type      HealthInsuranceCoding `json:"type"`
}

type HealthInsuranceCoding struct {
	Code    HealthInsuranceCodingCode `json:"code"`
	Display *string                   `json:"display,omitempty"`
	System  *string                   `json:"system,omitempty"`
	Version *string                   `json:"version,omitempty"`
}

type VitalStatusCoding struct {
	Code    VitalStatusCodingCode `json:"code"`
	Display *string               `json:"display,omitempty"`
	System  *string               `json:"system,omitempty"`
	Version *string               `json:"version,omitempty"`
}

type PerformanceStatus struct {
	EffectiveDate string     `json:"effectiveDate"`
	ID            string     `json:"id"`
	Patient       Reference  `json:"patient"`
	Value         EcogCoding `json:"value"`
}

type EcogCoding struct {
	Code    EcogCodingCode `json:"code"`
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
	Type      MolecularDiagnosticReportCoding `json:"type"`
}

type MolecularDiagnosticReportCoding struct {
	Code    MolecularDiagnosticReportCodingCode `json:"code"`
	Display *string                             `json:"display,omitempty"`
	System  *string                             `json:"system,omitempty"`
	Version *string                             `json:"version,omitempty"`
}

type Response struct {
	EffectiveDate string               `json:"effectiveDate"`
	ID            string               `json:"id"`
	Method        ResponseMethodCoding `json:"method"`
	Patient       Reference            `json:"patient"`
	Therapy       Reference            `json:"therapy"`
	Value         RecistCoding         `json:"value"`
}

type ResponseMethodCoding struct {
	Code    ResponseMethodCodingCode `json:"code"`
	Display *string                  `json:"display,omitempty"`
	System  *string                  `json:"system,omitempty"`
	Version *string                  `json:"version,omitempty"`
}

type RecistCoding struct {
	Code    RecistCodingCode `json:"code"`
	Display *string          `json:"display,omitempty"`
	System  *string          `json:"system,omitempty"`
	Version *string          `json:"version,omitempty"`
}

type TumorSpecimen struct {
	Collection *Collection         `json:"collection,omitempty"`
	Diagnosis  Reference           `json:"diagnosis"`
	ID         string              `json:"id"`
	Patient    Reference           `json:"patient"`
	Type       TumorSpecimenCoding `json:"type"`
}

type Collection struct {
	Date         *string                                   `json:"date,omitempty"`
	Localization TumorSpecimenCollectionLocalizationCoding `json:"localization"`
	Method       TumorSpecimenCollectionMethodCoding       `json:"method"`
}

type TumorSpecimenCollectionLocalizationCoding struct {
	Code    TumorSpecimenCollectionLocalizationCodingCode `json:"code"`
	Display *string                                       `json:"display,omitempty"`
	System  *string                                       `json:"system,omitempty"`
	Version *string                                       `json:"version,omitempty"`
}

type TumorSpecimenCollectionMethodCoding struct {
	Code    TumorSpecimenCollectionMethodCodingCode `json:"code"`
	Display *string                                 `json:"display,omitempty"`
	System  *string                                 `json:"system,omitempty"`
	Version *string                                 `json:"version,omitempty"`
}

type TumorSpecimenCoding struct {
	Code    TumorSpecimenCodingCode `json:"code"`
	Display *string                 `json:"display,omitempty"`
	System  *string                 `json:"system,omitempty"`
	Version *string                 `json:"version,omitempty"`
}

type SystemicTherapy struct {
	History []MtbSystemicTherapy `json:"history"`
}

type GeneticCounselingRecommendationReasonCodingCode string

const (
	FamilyAnamnesis                                        GeneticCounselingRecommendationReasonCodingCode = "family-anamnesis"
	GeneticCounselingRecommendationReasonCodingCodeOther   GeneticCounselingRecommendationReasonCodingCode = "other"
	GeneticCounselingRecommendationReasonCodingCodeUnknown GeneticCounselingRecommendationReasonCodingCode = "unknown"
	SecondaryTumor                                         GeneticCounselingRecommendationReasonCodingCode = "secondary-tumor"
	SelfAnamnesis                                          GeneticCounselingRecommendationReasonCodingCode = "self-anamnesis"
)

type MtbMedicationRecommendationCategoryCodingCode string

const (
	MtbMedicationRecommendationCategoryCodingCodeCh MtbMedicationRecommendationCategoryCodingCode = "CH"
	MtbMedicationRecommendationCategoryCodingCodeHo MtbMedicationRecommendationCategoryCodingCode = "HO"
	MtbMedicationRecommendationCategoryCodingCodeIm MtbMedicationRecommendationCategoryCodingCode = "IM"
	MtbMedicationRecommendationCategoryCodingCodeSO MtbMedicationRecommendationCategoryCodingCode = "SO"
	MtbMedicationRecommendationCategoryCodingCodeSz MtbMedicationRecommendationCategoryCodingCode = "SZ"
	MtbMedicationRecommendationCategoryCodingCodeZs MtbMedicationRecommendationCategoryCodingCode = "ZS"
)

type LevelOfEvidenceAddendumCodingCode string

const (
	Is LevelOfEvidenceAddendumCodingCode = "is"
	Iv LevelOfEvidenceAddendumCodingCode = "iv"
	R  LevelOfEvidenceAddendumCodingCode = "R"
	Z  LevelOfEvidenceAddendumCodingCode = "Z"
)

type LevelOfEvidenceGradingCodingCode string

const (
	LevelOfEvidenceGradingCodingCodeUndefined LevelOfEvidenceGradingCodingCode = "undefined"
	LevelOfEvidenceGradingCodingCodeM1A       LevelOfEvidenceGradingCodingCode = "m1A"
	LevelOfEvidenceGradingCodingCodeM1B       LevelOfEvidenceGradingCodingCode = "m1B"
	LevelOfEvidenceGradingCodingCodeM1C       LevelOfEvidenceGradingCodingCode = "m1C"
	LevelOfEvidenceGradingCodingCodeM2A       LevelOfEvidenceGradingCodingCode = "m2A"
	LevelOfEvidenceGradingCodingCodeM2B       LevelOfEvidenceGradingCodingCode = "m2B"
	LevelOfEvidenceGradingCodingCodeM2C       LevelOfEvidenceGradingCodingCode = "m2C"
	LevelOfEvidenceGradingCodingCodeM3        LevelOfEvidenceGradingCodingCode = "m3"
	LevelOfEvidenceGradingCodingCodeM4        LevelOfEvidenceGradingCodingCode = "m4"
)

type PublicationSystem string

const (
	PubmedNcbiNlmNihGov PublicationSystem = "https://pubmed.ncbi.nlm.nih.gov"
	DoiOrg              PublicationSystem = "https://www.doi.org"
)

type RequestedMedicationSystem string

const (
	FhirDeCodeSystemBfarmAtc RequestedMedicationSystem = "http://fhir.de/CodeSystem/bfarm/atc"
	SystemUndefined          RequestedMedicationSystem = "undefined"
)

type RecommendationPriorityCodingCode string

const (
	RecommendationPriorityCodingCode1 RecommendationPriorityCodingCode = "1"
	RecommendationPriorityCodingCode2 RecommendationPriorityCodingCode = "2"
	RecommendationPriorityCodingCode3 RecommendationPriorityCodingCode = "3"
	RecommendationPriorityCodingCode4 RecommendationPriorityCodingCode = "4"
)

type MtbMedicationRecommendationUseTypeCodingCode string

const (
	MtbMedicationRecommendationUseTypeCodingCodeCompassionate MtbMedicationRecommendationUseTypeCodingCode = "compassionate"
	MtbMedicationRecommendationUseTypeCodingCodeInLabel       MtbMedicationRecommendationUseTypeCodingCode = "in-label"
	MtbMedicationRecommendationUseTypeCodingCodeUnknown       MtbMedicationRecommendationUseTypeCodingCode = "unknown"
	MtbMedicationRecommendationUseTypeCodingCodeOffLabel      MtbMedicationRecommendationUseTypeCodingCode = "off-label"
	MtbMedicationRecommendationUseTypeCodingCodeSECPreventive MtbMedicationRecommendationUseTypeCodingCode = "sec-preventive"
)

type NoSequencingPerformedReasonCode string

const (
	MtbDiagnosisCodingCodeOther    NoSequencingPerformedReasonCode = "other"
	NonGeneticCause                NoSequencingPerformedReasonCode = "non-genetic-cause"
	NotRareDisease                 NoSequencingPerformedReasonCode = "not-rare-disease"
	Psychosomatic                  NoSequencingPerformedReasonCode = "psychosomatic"
	TargetedDiagnosticsRecommended NoSequencingPerformedReasonCode = "targeted-diagnostics-recommended"
)

type MtbProcedureRecommendationCategoryCodingCode string

const (
	MtbProcedureRecommendationCategoryCodingCodeAs MtbProcedureRecommendationCategoryCodingCode = "AS"
	MtbProcedureRecommendationCategoryCodingCodeSO MtbProcedureRecommendationCategoryCodingCode = "SO"
	MtbProcedureRecommendationCategoryCodingCodeOp MtbProcedureRecommendationCategoryCodingCode = "OP"
	MtbProcedureRecommendationCategoryCodingCodeSt MtbProcedureRecommendationCategoryCodingCode = "ST"
	MtbProcedureRecommendationCategoryCodingCodeWs MtbProcedureRecommendationCategoryCodingCode = "WS"
	MtbProcedureRecommendationCategoryCodingCodeWw MtbProcedureRecommendationCategoryCodingCode = "WW"
)

type MtbCarePlanRecommendationsMissingReasonCodingCode string

const (
	NoTarget MtbCarePlanRecommendationsMissingReasonCodingCode = "no-target"
)

type StudySystem string

const (
	Drks    StudySystem = "DRKS"
	Eudamed StudySystem = "EUDAMED"
	EudraCT StudySystem = "Eudra-CT"
	Nct     StudySystem = "NCT"
)

type ClaimResponseStatusCodingCode string

const (
	Accepted                             ClaimResponseStatusCodingCode = "accepted"
	ClaimResponseStatusCodingCodeUnknown ClaimResponseStatusCodingCode = "unknown"
	Rejected                             ClaimResponseStatusCodingCode = "rejected"
)

type ClaimResponseStatusReasonCodingCode string

const (
	ClaimResponseStatusReasonCodingCodeApprovalRevocation          ClaimResponseStatusReasonCodingCode = "approval-revocation"
	ClaimResponseStatusReasonCodingCodeOther                       ClaimResponseStatusReasonCodingCode = "other"
	ClaimResponseStatusReasonCodingCodeUnknown                     ClaimResponseStatusReasonCodingCode = "unknown"
	ClaimResponseStatusReasonCodingCodeFormalReasons               ClaimResponseStatusReasonCodingCode = "formal-reasons"
	ClaimResponseStatusReasonCodingCodeInclusionInStudy            ClaimResponseStatusReasonCodingCode = "inclusion-in-study"
	ClaimResponseStatusReasonCodingCodeInsufficientEvidence        ClaimResponseStatusReasonCodingCode = "insufficient-evidence"
	ClaimResponseStatusReasonCodingCodeOtherTherapyRecommended     ClaimResponseStatusReasonCodingCode = "other-therapy-recommended"
	ClaimResponseStatusReasonCodingCodeStandardTherapyNotExhausted ClaimResponseStatusReasonCodingCode = "standard-therapy-not-exhausted"
)

type ClaimStageCodingCode string

const (
	ClaimStageCodingCodeUnknown       ClaimStageCodingCode = "unknown"
	ClaimStageCodingCodeFollowUpClaim ClaimStageCodingCode = "follow-up-claim"
	ClaimStageCodingCodeInitialClaim  ClaimStageCodingCode = "initial-claim"
	ClaimStageCodingCodeRevocation    ClaimStageCodingCode = "revocation"
)

type MtbDiagnosisGuidelineTreatmentStatusCodingCode string

const (
	MtbDiagnosisGuidelineTreatmentStatusCodingCodeExhausted             MtbDiagnosisGuidelineTreatmentStatusCodingCode = "exhausted"
	MtbDiagnosisGuidelineTreatmentStatusCodingCodeImpossible            MtbDiagnosisGuidelineTreatmentStatusCodingCode = "impossible"
	MtbDiagnosisGuidelineTreatmentStatusCodingCodeUnknown               MtbDiagnosisGuidelineTreatmentStatusCodingCode = "unknown"
	MtbDiagnosisGuidelineTreatmentStatusCodingCodeNoGuidelinesAvailable MtbDiagnosisGuidelineTreatmentStatusCodingCode = "no-guidelines-available"
	MtbDiagnosisGuidelineTreatmentStatusCodingCodeNonExhausted          MtbDiagnosisGuidelineTreatmentStatusCodingCode = "non-exhausted"
)

type TumorStagingMethodCodingCode string

const (
	Clinical   TumorStagingMethodCodingCode = "clinical"
	Pathologic TumorStagingMethodCodingCode = "pathologic"
)

type ValueCode string

const (
	Main         ValueCode = "main"
	Metachronous ValueCode = "metachronous"
	Secondary    ValueCode = "secondary"
)

type FamilyMemberHistoryRelationshipTypeCodingCode string

const (
	EXT     FamilyMemberHistoryRelationshipTypeCodingCode = "EXT"
	Fammemb FamilyMemberHistoryRelationshipTypeCodingCode = "FAMMEMB"
)

type FollowUpPatientStatusCodingCode string

const (
	FollowUpPatientStatusCodingCodeLostToFu FollowUpPatientStatusCodingCode = "lost-to-fu"
)

type OncoProcedureCodingCode string

const (
	NuclearMedicine OncoProcedureCodingCode = "nuclear-medicine"
	RadioTherapy    OncoProcedureCodingCode = "radio-therapy"
	Surgery         OncoProcedureCodingCode = "surgery"
)

type MtbTherapyIntentCodingCode string

const (
	MtbTherapyIntentCodingCodeK MtbTherapyIntentCodingCode = "K"
	MtbTherapyIntentCodingCodeS MtbTherapyIntentCodingCode = "S"
	MtbTherapyIntentCodingCodeP MtbTherapyIntentCodingCode = "P"
	MtbTherapyIntentCodingCodeX MtbTherapyIntentCodingCode = "X"
)

type TherapyStatusCodingCode string

const (
	TherapyStatusCodingCodeCompleted TherapyStatusCodingCode = "completed"
	TherapyStatusCodingCodeNotDone   TherapyStatusCodingCode = "not-done"
	TherapyStatusCodingCodeOnGoing   TherapyStatusCodingCode = "on-going"
	TherapyStatusCodingCodeStopped   TherapyStatusCodingCode = "stopped"
	TherapyStatusCodingCodeUnknown   TherapyStatusCodingCode = "unknown"
)

type MtbTherapyStatusReasonCodingCode string

const (
	MtbTherapyStatusReasonCodingCodeBestSupportiveCare                   MtbTherapyStatusReasonCodingCode = "best-supportive-care"
	MtbTherapyStatusReasonCodingCodeChronicRemission                     MtbTherapyStatusReasonCodingCode = "chronic-remission"
	MtbTherapyStatusReasonCodingCodeDeterioration                        MtbTherapyStatusReasonCodingCode = "deterioration"
	MtbTherapyStatusReasonCodingCodeMedicalReasons                       MtbTherapyStatusReasonCodingCode = "medical-reasons"
	MtbTherapyStatusReasonCodingCodeLostToFu                             MtbTherapyStatusReasonCodingCode = "lost-to-fu"
	MtbTherapyStatusReasonCodingCodeOther                                MtbTherapyStatusReasonCodingCode = "other"
	MtbTherapyStatusReasonCodingCodeNoIndication                         MtbTherapyStatusReasonCodingCode = "no-indication"
	MtbTherapyStatusReasonCodingCodeOtherTherapyChosen                   MtbTherapyStatusReasonCodingCode = "other-therapy-chosen"
	MtbTherapyStatusReasonCodingCodePatientDeath                         MtbTherapyStatusReasonCodingCode = "patient-death"
	MtbTherapyStatusReasonCodingCodePatientRefusal                       MtbTherapyStatusReasonCodingCode = "patient-refusal"
	MtbTherapyStatusReasonCodingCodePatientWish                          MtbTherapyStatusReasonCodingCode = "patient-wish"
	MtbTherapyStatusReasonCodingCodePaymentEnded                         MtbTherapyStatusReasonCodingCode = "payment-ended"
	MtbTherapyStatusReasonCodingCodePaymentPending                       MtbTherapyStatusReasonCodingCode = "payment-pending"
	MtbTherapyStatusReasonCodingCodePaymentRefused                       MtbTherapyStatusReasonCodingCode = "payment-refused"
	MtbTherapyStatusReasonCodingCodeProgression                          MtbTherapyStatusReasonCodingCode = "progression"
	MtbTherapyStatusReasonCodingCodeRegularCompletion                    MtbTherapyStatusReasonCodingCode = "regular-completion"
	MtbTherapyStatusReasonCodingCodeRegularCompletionWithDosageReduction MtbTherapyStatusReasonCodingCode = "regular-completion-with-dosage-reduction"
	MtbTherapyStatusReasonCodingCodeRegularCompletionWithSubstanceChange MtbTherapyStatusReasonCodingCode = "regular-completion-with-substance-change"
	MtbTherapyStatusReasonCodingCodeToxicity                             MtbTherapyStatusReasonCodingCode = "toxicity"
)

type MtbSystemicTherapyCategoryCodingCode string

const (
	MtbSystemicTherapyCategoryCodingCodeA MtbSystemicTherapyCategoryCodingCode = "A"
	MtbSystemicTherapyCategoryCodingCodeI MtbSystemicTherapyCategoryCodingCode = "I"
	MtbSystemicTherapyCategoryCodingCodeS MtbSystemicTherapyCategoryCodingCode = "S"
	MtbSystemicTherapyCategoryCodingCodeN MtbSystemicTherapyCategoryCodingCode = "N"
	MtbSystemicTherapyCategoryCodingCodeO MtbSystemicTherapyCategoryCodingCode = "O"
)

type MtbSystemicTherapyDosageDensityCodingCode string

const (
	Over50  MtbSystemicTherapyDosageDensityCodingCode = "over-50%"
	Under50 MtbSystemicTherapyDosageDensityCodingCode = "under-50%"
)

type MtbSystemicTherapyRecommendationFulfillmentStatusCodingCode string

const (
	Complete MtbSystemicTherapyRecommendationFulfillmentStatusCodingCode = "complete"
	Partial  MtbSystemicTherapyRecommendationFulfillmentStatusCodingCode = "partial"
)

type TumorCellContentMethodCodingCode string

const (
	Bioinformatic TumorCellContentMethodCodingCode = "bioinformatic"
	Histologic    TumorCellContentMethodCodingCode = "histologic"
)

type ProteinExpressionICScoreCodingCode string

const (
	ProteinExpressionICScoreCodingCode0 ProteinExpressionICScoreCodingCode = "0"
	ProteinExpressionICScoreCodingCode1 ProteinExpressionICScoreCodingCode = "1"
	ProteinExpressionICScoreCodingCode2 ProteinExpressionICScoreCodingCode = "2"
	ProteinExpressionICScoreCodingCode3 ProteinExpressionICScoreCodingCode = "3"
)

type ProteinExpressionTcScoreCodingCode string

const (
	ProteinExpressionTcScoreCodingCode0 ProteinExpressionTcScoreCodingCode = "0"
	ProteinExpressionTcScoreCodingCode1 ProteinExpressionTcScoreCodingCode = "1"
	ProteinExpressionTcScoreCodingCode2 ProteinExpressionTcScoreCodingCode = "2"
	ProteinExpressionTcScoreCodingCode3 ProteinExpressionTcScoreCodingCode = "3"
	ProteinExpressionTcScoreCodingCode4 ProteinExpressionTcScoreCodingCode = "4"
	ProteinExpressionTcScoreCodingCode5 ProteinExpressionTcScoreCodingCode = "5"
	ProteinExpressionTcScoreCodingCode6 ProteinExpressionTcScoreCodingCode = "6"
)

type ProteinExpressionResultCodingCode string

const (
	ProteinExpressionResultCodingCodeExp       ProteinExpressionResultCodingCode = "exp"
	ProteinExpressionResultCodingCodeNotExp    ProteinExpressionResultCodingCode = "not-exp"
	ProteinExpressionResultCodingCodeUnknown   ProteinExpressionResultCodingCode = "unknown"
	ProteinExpressionResultCodingCodeCode1Plus ProteinExpressionResultCodingCode = "1+"
	ProteinExpressionResultCodingCodeCode2Plus ProteinExpressionResultCodingCode = "2+"
	ProteinExpressionResultCodingCodeCode3Plus ProteinExpressionResultCodingCode = "3+"
)

type MsiInterpretationCodingCode string

const (
	MsiInterpretationCodingCodeMsiHigh       MsiInterpretationCodingCode = "msi-high"
	MsiInterpretationCodingCodeUnknown       MsiInterpretationCodingCode = "unknown"
	MsiInterpretationCodingCodeMsiLow        MsiInterpretationCodingCode = "msi-low"
	MsiInterpretationCodingCodeMmrDeficient  MsiInterpretationCodingCode = "mmr-deficient"
	MsiInterpretationCodingCodeMmrProficient MsiInterpretationCodingCode = "mmr-proficient"
	MsiInterpretationCodingCodeStable        MsiInterpretationCodingCode = "stable"
)

type MsiMethodCodingCode string

const (
	MsiMethodCodingCodeIhc           MsiMethodCodingCode = "IHC"
	MsiMethodCodingCodeBioinformatic MsiMethodCodingCode = "bioinformatic"
	MsiMethodCodingCodePCR           MsiMethodCodingCode = "PCR"
)

type ModelProjectConsentPurpose string

const (
	CaseIdentification ModelProjectConsentPurpose = "case-identification"
	Reidentification   ModelProjectConsentPurpose = "reidentification"
	Sequencing         ModelProjectConsentPurpose = "sequencing"
)

type ConsentProvision string

const (
	Deny   ConsentProvision = "deny"
	Permit ConsentProvision = "permit"
)

type MvhSubmissionType string

const (
	Addition   MvhSubmissionType = "addition"
	Correction MvhSubmissionType = "correction"
	Followup   MvhSubmissionType = "followup"
	Initial    MvhSubmissionType = "initial"
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
	ChrMt Chromosome = "chrMT"
)

type ExternalIDSystem string

const (
	CancerSangerACUkCosmic ExternalIDSystem = "https://cancer.sanger.ac.uk/cosmic"
	NcbiNlmNihGovEntrez    ExternalIDSystem = "https://www.ncbi.nlm.nih.gov/entrez"
	NcbiNlmNihGovSnp       ExternalIDSystem = "https://www.ncbi.nlm.nih.gov/snp"
	EnsemblOrg             ExternalIDSystem = "https://www.ensembl.org"
)

type BaseVariantLocalizationCodingCode string

const (
	CodingRegion     BaseVariantLocalizationCodingCode = "coding-region"
	Intergenic       BaseVariantLocalizationCodingCode = "intergenic"
	Intronic         BaseVariantLocalizationCodingCode = "intronic"
	RegulatoryRegion BaseVariantLocalizationCodingCode = "regulatory-region"
	SplicingRegion   BaseVariantLocalizationCodingCode = "splicing-region"
)

type CnvCodingCode string

const (
	HighLevelGain CnvCodingCode = "high-level-gain"
	Loss          CnvCodingCode = "loss"
	LowLevelGain  CnvCodingCode = "low-level-gain"
)

type InterpretationCodingCode string

const (
	High         InterpretationCodingCode = "high"
	Intermediate InterpretationCodingCode = "intermediate"
	Low          InterpretationCodingCode = "low"
)

type StrandEnum string

const (
	Empty           StrandEnum = "+"
	RnaFusionStrand StrandEnum = "-"
)

type TranscriptIDSystem string

const (
	FluffyHTTPSWWWEnsemblOrg    TranscriptIDSystem = "https://www.ensembl.org"
	HTTPSWWWNcbiNlmNihGovRefseq TranscriptIDSystem = "https://www.ncbi.nlm.nih.gov/refseq"
)

type ClinVarCodingCode string

const (
	ClinVarCodingCode1 ClinVarCodingCode = "1"
	ClinVarCodingCode2 ClinVarCodingCode = "2"
	ClinVarCodingCode3 ClinVarCodingCode = "3"
	ClinVarCodingCode4 ClinVarCodingCode = "4"
	ClinVarCodingCode5 ClinVarCodingCode = "5"
)

type NgsReportCodingCode string

const (
	NgsReportCodingCodeArray           NgsReportCodingCode = "array"
	NgsReportCodingCodeExome           NgsReportCodingCode = "exome"
	NgsReportCodingCodeGenomeLongRead  NgsReportCodingCode = "genome-long-read"
	NgsReportCodingCodeGenomeShortRead NgsReportCodingCode = "genome-short-read"
	NgsReportCodingCodeKaryotyping     NgsReportCodingCode = "karyotyping"
	NgsReportCodingCodeOther           NgsReportCodingCode = "other"
	NgsReportCodingCodePanel           NgsReportCodingCode = "panel"
	NgsReportCodingCodeSingle          NgsReportCodingCode = "single"
)

type Unit string

const (
	Months Unit = "Months"
	Years  Unit = "Years"
)

type GenderCodingCode string

const (
	GenderCodingCodeFemale  GenderCodingCode = "female"
	GenderCodingCodeOther   GenderCodingCode = "other"
	GenderCodingCodeUnknown GenderCodingCode = "unknown"
	GenderCodingCodeMale    GenderCodingCode = "male"
)

type HealthInsuranceCodingCode string

const (
	Bei HealthInsuranceCodingCode = "BEI"
	Bg  HealthInsuranceCodingCode = "BG"
	Gkv HealthInsuranceCodingCode = "GKV"
	Gpv HealthInsuranceCodingCode = "GPV"
	Pkv HealthInsuranceCodingCode = "PKV"
	Ppv HealthInsuranceCodingCode = "PPV"
	Sel HealthInsuranceCodingCode = "SEL"
	Skt HealthInsuranceCodingCode = "SKT"
	Soz HealthInsuranceCodingCode = "SOZ"
	Unk HealthInsuranceCodingCode = "UNK"
)

type VitalStatusCodingCode string

const (
	Alive    VitalStatusCodingCode = "alive"
	Deceased VitalStatusCodingCode = "deceased"
)

type EcogCodingCode string

const (
	EcogCodingCode0 EcogCodingCode = "0"
	EcogCodingCode1 EcogCodingCode = "1"
	EcogCodingCode2 EcogCodingCode = "2"
	EcogCodingCode3 EcogCodingCode = "3"
	EcogCodingCode4 EcogCodingCode = "4"
	EcogCodingCode5 EcogCodingCode = "5"
)

type MolecularDiagnosticReportCodingCode string

const (
	MolecularDiagnosticReportCodingCodeFish            MolecularDiagnosticReportCodingCode = "FISH"
	MolecularDiagnosticReportCodingCodeFusionPanel     MolecularDiagnosticReportCodingCode = "fusion-panel"
	MolecularDiagnosticReportCodingCodeGenePanel       MolecularDiagnosticReportCodingCode = "gene-panel"
	MolecularDiagnosticReportCodingCodeArray           MolecularDiagnosticReportCodingCode = "array"
	MolecularDiagnosticReportCodingCodeExome           MolecularDiagnosticReportCodingCode = "exome"
	MolecularDiagnosticReportCodingCodeGenomeLongRead  MolecularDiagnosticReportCodingCode = "genome-long-read"
	MolecularDiagnosticReportCodingCodeGenomeShortRead MolecularDiagnosticReportCodingCode = "genome-short-read"
	MolecularDiagnosticReportCodingCodeKaryotyping     MolecularDiagnosticReportCodingCode = "karyotyping"
	MolecularDiagnosticReportCodingCodeOther           MolecularDiagnosticReportCodingCode = "other"
	MolecularDiagnosticReportCodingCodePanel           MolecularDiagnosticReportCodingCode = "panel"
	MolecularDiagnosticReportCodingCodeSingle          MolecularDiagnosticReportCodingCode = "single"
	MolecularDiagnosticReportCodingCodePcr             MolecularDiagnosticReportCodingCode = "PCR"
)

type ResponseMethodCodingCode string

const (
	Rano   ResponseMethodCodingCode = "RANO"
	Recist ResponseMethodCodingCode = "RECIST"
)

type RecistCodingCode string

const (
	CR RecistCodingCode = "CR"
	Mr RecistCodingCode = "MR"
	Na RecistCodingCode = "NA"
	PD RecistCodingCode = "PD"
	PR RecistCodingCode = "PR"
	SD RecistCodingCode = "SD"
)

type TumorSpecimenCollectionLocalizationCodingCode string

const (
	TumorSpecimenCollectionLocalizationCodingCodeCellfreeDna        TumorSpecimenCollectionLocalizationCodingCode = "cellfree-dna"
	TumorSpecimenCollectionLocalizationCodingCodeLocalRecurrence    TumorSpecimenCollectionLocalizationCodingCode = "local-recurrence"
	TumorSpecimenCollectionLocalizationCodingCodeMetastasis         TumorSpecimenCollectionLocalizationCodingCode = "metastasis"
	TumorSpecimenCollectionLocalizationCodingCodePrimaryTumor       TumorSpecimenCollectionLocalizationCodingCode = "primary-tumor"
	TumorSpecimenCollectionLocalizationCodingCodeRegionalLymphNodes TumorSpecimenCollectionLocalizationCodingCode = "regional-lymph-nodes"
	TumorSpecimenCollectionLocalizationCodingCodeUnknown            TumorSpecimenCollectionLocalizationCodingCode = "unknown"
)

type TumorSpecimenCollectionMethodCodingCode string

const (
	TumorSpecimenCollectionMethodCodingCodeBiopsy       TumorSpecimenCollectionMethodCodingCode = "biopsy"
	TumorSpecimenCollectionMethodCodingCodeCytology     TumorSpecimenCollectionMethodCodingCode = "cytology"
	TumorSpecimenCollectionMethodCodingCodeResection    TumorSpecimenCollectionMethodCodingCode = "resection"
	TumorSpecimenCollectionMethodCodingCodeLiquidBiopsy TumorSpecimenCollectionMethodCodingCode = "liquid-biopsy"
	TumorSpecimenCollectionMethodCodingCodeUnknown      TumorSpecimenCollectionMethodCodingCode = "unknown"
)

type TumorSpecimenCodingCode string

const (
	TumorSpecimenCodingCodeCryoFrozen   TumorSpecimenCodingCode = "cryo-frozen"
	TumorSpecimenCodingCodeFfpe         TumorSpecimenCodingCode = "FFPE"
	TumorSpecimenCodingCodeFreshTissue  TumorSpecimenCodingCode = "fresh-tissue"
	TumorSpecimenCodingCodeLiquidBiopsy TumorSpecimenCodingCode = "liquid-biopsy"
	TumorSpecimenCodingCodeUnknown      TumorSpecimenCodingCode = "unknown"
)
