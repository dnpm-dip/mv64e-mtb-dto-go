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
	Diagnoses                  []MTBDiagnosis                    `json:"diagnoses,omitempty"`
	Episode                    *MTBEpisode                       `json:"episode,omitempty"`
	EpisodesOfCare             []EpisodeOfCare                   `json:"episodesOfCare,omitempty"`
	GeneticCounsellingRequests []GeneticCounselingRecommendation `json:"geneticCounsellingRequests,omitempty"`
	GuidelineProcedures        []OncoProdecure                   `json:"guidelineProcedures,omitempty"`
	GuidelineTherapies         []GuidelineTherapyElement         `json:"guidelineTherapies,omitempty"`
	HistologyReports           []HistologyReport                 `json:"histologyReports,omitempty"`
	IhcReports                 []IHCReport                       `json:"ihcReports,omitempty"`
	MolecularTherapies         []MolecularTherapy                `json:"molecularTherapies,omitempty"`
	NgsReports                 []SomaticNGSReport                `json:"ngsReports,omitempty"`
	Patient                    MtbPatient                        `json:"patient"`
	PerformanceStatus          []PerformanceStatus               `json:"performanceStatus,omitempty"`
	Recommendations            []MTBMedicationRecommendation     `json:"recommendations,omitempty"`
	Responses                  []Response                        `json:"responses,omitempty"`
	Specimens                  []SpecimenElement                 `json:"specimens,omitempty"`
	StudyInclusionRequests     []StudyEnrollmentRecommendation   `json:"studyInclusionRequests,omitempty"`
	Therapies                  []Therapy                         `json:"therapies,omitempty"`
}

type MTBCarePlan struct {
	Diagnosis                       *string                          `json:"diagnosis,omitempty"`
	GeneticCounselingRecommendation *GeneticCounselingRecommendation `json:"geneticCounselingRecommendation,omitempty"`
	ID                              string                           `json:"id"`
	Indication                      *Reference                       `json:"indication,omitempty"`
	IssuedOn                        *string                          `json:"issuedOn,omitempty"`
	MedicationRecommendations       []MTBMedicationRecommendation    `json:"medicationRecommendations,omitempty"`
	NoTargetFinding                 *NoTargetFinding                 `json:"noTargetFinding,omitempty"`
	Notes                           *string                          `json:"notes,omitempty"`
	Patient                         GeneticCounsellingRequestPatient `json:"patient"`
	StatusReason                    *CodingCarePlanStatusReason      `json:"statusReason,omitempty"`
	StudyEnrollmentRecommendations  []StudyEnrollmentRecommendation  `json:"studyEnrollmentRecommendations,omitempty"`
}

type GeneticCounselingRecommendation struct {
	ID       string                           `json:"id"`
	IssuedOn *string                          `json:"issuedOn,omitempty"`
	Patient  GeneticCounsellingRequestPatient `json:"patient"`
	Reason   Coding                           `json:"reason"`
}

type GeneticCounsellingRequestPatient struct {
	ID   string      `json:"id"`
	Type PatientType `json:"type"`
}

type Coding struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	System  *string `json:"system,omitempty"`
	Version *string `json:"version,omitempty"`
}

type Reference struct {
	Display *string `json:"display,omitempty"`
	ID      string  `json:"id"`
	Type    *string `json:"type,omitempty"`
}

type MTBMedicationRecommendation struct {
	ID                 string                               `json:"id"`
	Indication         *Reference                           `json:"indication,omitempty"`
	IssuedOn           *string                              `json:"issuedOn,omitempty"`
	LevelOfEvidence    *LevelOfEvidence                     `json:"levelOfEvidence,omitempty"`
	Medication         []Coding                             `json:"medication,omitempty"`
	NgsReport          *string                              `json:"ngsReport,omitempty"`
	Patient            GeneticCounsellingRequestPatient     `json:"patient"`
	Priority           *CodingTherapyRecommendationPriority `json:"priority,omitempty"`
	SupportingVariants []Reference                          `json:"supportingVariants,omitempty"`
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
	System *EXTIDSystem `json:"system,omitempty"`
	Value  string       `json:"value"`
}

type CodingTherapyRecommendationPriority struct {
	Code    TherapyRecommendationPriority `json:"code"`
	Display *string                       `json:"display,omitempty"`
	System  *string                       `json:"system,omitempty"`
}

type NoTargetFinding struct {
	Diagnosis string                           `json:"diagnosis"`
	IssuedOn  *string                          `json:"issuedOn,omitempty"`
	Patient   GeneticCounsellingRequestPatient `json:"patient"`
}

type CodingCarePlanStatusReason struct {
	Code    string  `json:"code"`
	Display *string `json:"display,omitempty"`
	System  *string `json:"system,omitempty"`
}

type StudyEnrollmentRecommendation struct {
	ID                 string                           `json:"id"`
	IssuedOn           *string                          `json:"issuedOn,omitempty"`
	LevelOfEvidence    *Coding                          `json:"levelOfEvidence,omitempty"`
	Patient            GeneticCounsellingRequestPatient `json:"patient"`
	Reason             Reference                        `json:"reason"`
	Studies            []Study                          `json:"studies"`
	SupportingVariants []Reference                      `json:"supportingVariants,omitempty"`
}

type Study struct {
	System string `json:"system"`
	Value  string `json:"value"`
}

type ClaimResponse struct {
	Claim        ClaimResponseClaim               `json:"claim"`
	ID           string                           `json:"id"`
	IssuedOn     string                           `json:"issuedOn"`
	Patient      GeneticCounsellingRequestPatient `json:"patient"`
	Status       CodingClaimResponseStatus        `json:"status"`
	StatusReason *CodingClaimResponseStatusReason `json:"statusReason,omitempty"`
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

type CodingClaimResponseStatusReason struct {
	Code    ClaimResponseStatusReason `json:"code"`
	Display *string                   `json:"display,omitempty"`
	System  *string                   `json:"system,omitempty"`
}

type ClaimElement struct {
	ID             string                           `json:"id"`
	IssuedOn       string                           `json:"issuedOn"`
	Patient        GeneticCounsellingRequestPatient `json:"patient"`
	Recommendation *Recommendation                  `json:"recommendation,omitempty"`
	Stage          *Coding                          `json:"stage,omitempty"`
}

type Recommendation struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type MTBDiagnosis struct {
	Code                     Coding                           `json:"code"`
	GuidelineTreatmentStatus *Coding                          `json:"guidelineTreatmentStatus,omitempty"`
	HistologyResults         []string                         `json:"histologyResults,omitempty"`
	ID                       string                           `json:"id"`
	IcdO3T                   *Coding                          `json:"icdO3T,omitempty"`
	Patient                  GeneticCounsellingRequestPatient `json:"patient"`
	RecordedOn               *string                          `json:"recordedOn,omitempty"`
	StageHistory             []StageHistory                   `json:"stageHistory,omitempty"`
	Topography               *Coding                          `json:"topography,omitempty"`
	TumorGrade               *CodingTumorGrade                `json:"tumorGrade,omitempty"`
	WhoGrade                 *Coding                          `json:"whoGrade,omitempty"`
	WhoGrading               *Coding                          `json:"whoGrading,omitempty"`
}

type StageHistory struct {
	Date  string            `json:"date"`
	Stage CodingTumorSpread `json:"stage"`
}

type CodingTumorSpread struct {
	Code    StageCode `json:"code"`
	Display *string   `json:"display,omitempty"`
	System  *string   `json:"system,omitempty"`
}

type CodingTumorGrade struct {
	Code    TumorGradeCode `json:"code"`
	Display *string        `json:"display,omitempty"`
	System  *string        `json:"system,omitempty"`
	Version *string        `json:"version,omitempty"`
}

type MTBEpisode struct {
	ID      string                           `json:"id"`
	Patient GeneticCounsellingRequestPatient `json:"patient"`
	Period  PeriodLocalDate                  `json:"period"`
}

type PeriodLocalDate struct {
	End   *string `json:"end,omitempty"`
	Start string  `json:"start"`
}

type EpisodeOfCare struct {
	Diagnoses []Reference     `json:"diagnoses,omitempty"`
	ID        string          `json:"id"`
	Patient   Reference       `json:"patient"`
	Period    PeriodLocalDate `json:"period"`
}

type OncoProdecure struct {
	BasedOn      *string                          `json:"basedOn,omitempty"`
	Code         *Coding                          `json:"code,omitempty"`
	Diagnosis    *string                          `json:"diagnosis,omitempty"`
	ID           string                           `json:"id"`
	Indication   *Reference                       `json:"indication,omitempty"`
	Notes        *string                          `json:"notes,omitempty"`
	Patient      GeneticCounsellingRequestPatient `json:"patient"`
	Period       *PeriodLocalDate                 `json:"period,omitempty"`
	RecordedOn   *string                          `json:"recordedOn,omitempty"`
	Status       *CodingTherapyStatus             `json:"status,omitempty"`
	StatusReason *CodingTherapyStatusReason       `json:"statusReason,omitempty"`
	TherapyLine  *int64                           `json:"therapyLine,omitempty"`
}

type CodingTherapyStatus struct {
	Code    TherapyStatus `json:"code"`
	Display *string       `json:"display,omitempty"`
	System  *string       `json:"system,omitempty"`
}

type CodingTherapyStatusReason struct {
	Code    StatusReasonCode `json:"code"`
	Display *string          `json:"display,omitempty"`
	System  *string          `json:"system,omitempty"`
	Version *string          `json:"version,omitempty"`
}

type GuidelineTherapyElement struct {
	BasedOn      *Reference                       `json:"basedOn,omitempty"`
	Diagnosis    *string                          `json:"diagnosis,omitempty"`
	ID           string                           `json:"id"`
	Indication   *Reference                       `json:"indication,omitempty"`
	Medication   []Coding                         `json:"medication,omitempty"`
	Notes        *string                          `json:"notes,omitempty"`
	Patient      GeneticCounsellingRequestPatient `json:"patient"`
	Period       *PeriodLocalDate                 `json:"period,omitempty"`
	RecordedOn   *string                          `json:"recordedOn,omitempty"`
	Status       *CodingTherapyStatus             `json:"status,omitempty"`
	StatusReason *CodingTherapyStatusReason       `json:"statusReason,omitempty"`
	TherapyLine  *int64                           `json:"therapyLine,omitempty"`
}

type HistologyReport struct {
	ID       string                           `json:"id"`
	IssuedOn string                           `json:"issuedOn"`
	Patient  GeneticCounsellingRequestPatient `json:"patient"`
	Results  HistologyReportResults           `json:"results"`
	Specimen HistologyReportSpecimen          `json:"specimen"`
}

type HistologyReportResults struct {
	TumorCellContent TumorCellContent `json:"tumorCellContent"`
	TumorMorphology  TumorMorphology  `json:"tumorMorphology"`
}

type TumorCellContent struct {
	ID       string                            `json:"id"`
	Method   CodingTumorCellContentMethod      `json:"method"`
	Patient  *GeneticCounsellingRequestPatient `json:"patient,omitempty"`
	Specimen TumorCellContentSpecimen          `json:"specimen"`
	Value    float64                           `json:"value"`
}

type CodingTumorCellContentMethod struct {
	Code    TumorCellContentMethod `json:"code"`
	Display *string                `json:"display,omitempty"`
	System  *string                `json:"system,omitempty"`
}

type TumorCellContentSpecimen struct {
	ID   string       `json:"id"`
	Type SpecimenType `json:"type"`
}

type TumorMorphology struct {
	ID       string                           `json:"id"`
	Notes    *string                          `json:"notes,omitempty"`
	Patient  GeneticCounsellingRequestPatient `json:"patient"`
	Specimen TumorMorphologySpecimen          `json:"specimen"`
	Value    Coding                           `json:"value"`
}

type TumorMorphologySpecimen struct {
	ID   string       `json:"id"`
	Type SpecimenType `json:"type"`
}

type HistologyReportSpecimen struct {
	ID   string       `json:"id"`
	Type SpecimenType `json:"type"`
}

type IHCReport struct {
	BlockID                  ExternalID                `json:"blockId"`
	Date                     string                    `json:"date"`
	ID                       string                    `json:"id"`
	JournalID                ExternalID                `json:"journalId"`
	MSIMmrResults            []MSIMmrResult            `json:"msiMmrResults"`
	Patient                  Reference                 `json:"patient"`
	ProteinExpressionResults []ProteinExpressionResult `json:"proteinExpressionResults"`
	Specimen                 Reference                 `json:"specimen"`
}

type ExternalID struct {
	System *string `json:"system,omitempty"`
	Value  string  `json:"value"`
}

type MSIMmrResult struct {
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
	Code    ProteinExpressionResultCode `json:"code"`
	Display *string                     `json:"display,omitempty"`
	System  *string                     `json:"system,omitempty"`
	Version *string                     `json:"version,omitempty"`
}

type ProteinExpressionResult struct {
	ICScore  *CodingProteinExpressionICScore `json:"icScore,omitempty"`
	ID       string                          `json:"id"`
	Patient  Reference                       `json:"patient"`
	Protein  Coding                          `json:"protein"`
	TcScore  *CodingProteinExpressionTCScore `json:"tcScore,omitempty"`
	TpsScore *int64                          `json:"tpsScore,omitempty"`
	Value    CodingProteinExpressionResult   `json:"value"`
}

type MolecularTherapy struct {
	History []GuidelineTherapyElement `json:"history"`
}

type SomaticNGSReport struct {
	ID             string                           `json:"id"`
	IssuedOn       *string                          `json:"issuedOn,omitempty"`
	MSI            *float64                         `json:"msi,omitempty"`
	Metadata       []Metadatum                      `json:"metadata"`
	Patient        GeneticCounsellingRequestPatient `json:"patient"`
	Results        *NgsReportResults                `json:"results,omitempty"`
	SequencingType Coding                           `json:"sequencingType"`
	Specimen       NgsReportSpecimen                `json:"specimen"`
}

type Metadatum struct {
	KitManufacturer string `json:"kitManufacturer"`
	KitType         string `json:"kitType"`
	Pipeline        string `json:"pipeline"`
	ReferenceGenome string `json:"referenceGenome"`
	Sequencer       string `json:"sequencer"`
}

type NgsReportResults struct {
	Brcaness           *BRCAness         `json:"brcaness,omitempty"`
	CopyNumberVariants []Cnv             `json:"copyNumberVariants,omitempty"`
	DnaFusions         []DNAFusion       `json:"dnaFusions,omitempty"`
	HrdScore           *HRDScore         `json:"hrdScore,omitempty"`
	RnaFusions         []RNAFusion       `json:"rnaFusions,omitempty"`
	RnaSeqs            []RNASeq          `json:"rnaSeqs,omitempty"`
	SimpleVariants     []Snv             `json:"simpleVariants,omitempty"`
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
	CNA                   *float64                          `json:"cnA,omitempty"`
	CNB                   *float64                          `json:"cnB,omitempty"`
	Chromosome            CodingChromosome                  `json:"chromosome"`
	CopyNumberNeutralLoH  []CodingGene                      `json:"copyNumberNeutralLoH,omitempty"`
	EndRange              EndRange                          `json:"endRange"`
	ID                    string                            `json:"id"`
	Indication            *Reference                        `json:"indication,omitempty"`
	Patient               *GeneticCounsellingRequestPatient `json:"patient,omitempty"`
	RelativeCopyNumber    *float64                          `json:"relativeCopyNumber,omitempty"`
	ReportedAffectedGenes []CodingGene                      `json:"reportedAffectedGenes,omitempty"`
	ReportedFocality      *string                           `json:"reportedFocality,omitempty"`
	StartRange            StartRange                        `json:"startRange"`
	TotalCopyNumber       *int64                            `json:"totalCopyNumber,omitempty"`
	Type                  CodingCNVType                     `json:"type"`
}

type CodingChromosome struct {
	Code    ChromosomeCode    `json:"code"`
	Display *string           `json:"display,omitempty"`
	System  *ChromosomeSystem `json:"system,omitempty"`
}

type CodingGene struct {
	Code    string      `json:"code"`
	Display *string     `json:"display,omitempty"`
	System  *GeneSystem `json:"system,omitempty"`
}

type EndRange struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type StartRange struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type CodingCNVType struct {
	Code    CNVType `json:"code"`
	Display *string `json:"display,omitempty"`
	System  *string `json:"system,omitempty"`
}

type DNAFusion struct {
	FusionPartner3Prime DnaFusionFusionPartner3Prime `json:"fusionPartner3prime"`
	FusionPartner5Prime DnaFusionFusionPartner5Prime `json:"fusionPartner5prime"`
	ID                  string                       `json:"id"`
	ReportedNumReads    int64                        `json:"reportedNumReads"`
}

type DnaFusionFusionPartner3Prime struct {
	Chromosome CodingChromosome `json:"chromosome"`
	Gene       CodingGene       `json:"gene"`
	Position   float64          `json:"position"`
}

type DnaFusionFusionPartner5Prime struct {
	Chromosome CodingChromosome `json:"chromosome"`
	Gene       Gene             `json:"gene"`
	Position   float64          `json:"position"`
}

type Gene struct {
	EnsemblID *string `json:"ensemblId,omitempty"`
	HgncID    *string `json:"hgncId,omitempty"`
}

type HRDScore struct {
	Components     Components `json:"components"`
	ID             string     `json:"id"`
	Interpretation *Coding    `json:"interpretation,omitempty"`
	Patient        Reference  `json:"patient"`
	Specimen       Reference  `json:"specimen"`
	Value          float64    `json:"value"`
}

type Components struct {
	Loh float64 `json:"loh"`
	Lst float64 `json:"lst"`
	Tai float64 `json:"tai"`
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
	Exon         string     `json:"exon"`
	Gene         CodingGene `json:"gene"`
	Position     float64    `json:"position"`
	Strand       StrandEnum `json:"strand"`
	TranscriptID string     `json:"transcriptId"`
}

type RnaFusionFusionPartner5Prime struct {
	Exon         string     `json:"exon"`
	Gene         CodingGene `json:"gene"`
	Position     float64    `json:"position"`
	Strand       StrandEnum `json:"strand"`
	TranscriptID string     `json:"transcriptId"`
}

type RNASeq struct {
	CohortRanking               *int64     `json:"cohortRanking,omitempty"`
	EnsemblID                   string     `json:"ensemblId"`
	EntrezID                    string     `json:"entrezId"`
	FragmentsPerKilobaseMillion float64    `json:"fragmentsPerKilobaseMillion"`
	FromNGS                     bool       `json:"fromNGS"`
	Gene                        CodingGene `json:"gene"`
	ID                          string     `json:"id"`
	LibrarySize                 int64      `json:"librarySize"`
	RawCounts                   int64      `json:"rawCounts"`
	TissueCorrectedExpression   bool       `json:"tissueCorrectedExpression"`
	TranscriptID                string     `json:"transcriptId"`
}

type Snv struct {
	AllelicFrequency float64                          `json:"allelicFrequency"`
	AltAllele        string                           `json:"altAllele"`
	AminoAcidChange  *Coding                          `json:"aminoAcidChange,omitempty"`
	Chromosome       CodingChromosome                 `json:"chromosome"`
	CosmicID         *string                          `json:"cosmicId,omitempty"`
	DBSNPID          *string                          `json:"dbSNPId,omitempty"`
	DnaChange        *Coding                          `json:"dnaChange,omitempty"`
	ExternalIDS      []ExternalID                     `json:"externalIds,omitempty"`
	Gene             *CodingGene                      `json:"gene,omitempty"`
	ID               string                           `json:"id"`
	Interpretation   *Coding                          `json:"interpretation,omitempty"`
	Patient          GeneticCounsellingRequestPatient `json:"patient"`
	Position         Position                         `json:"position"`
	ProteinChange    *Coding                          `json:"proteinChange,omitempty"`
	ReadDepth        int64                            `json:"readDepth"`
	RefAllele        string                           `json:"refAllele"`
	TranscriptID     *ExternalID                      `json:"transcriptId,omitempty"`
}

type Position struct {
	End   *float64 `json:"end,omitempty"`
	Start float64  `json:"start"`
}

type Tmb struct {
	ID             string    `json:"id"`
	Interpretation *Coding   `json:"interpretation,omitempty"`
	Patient        Reference `json:"patient"`
	Specimen       Reference `json:"specimen"`
	Value          Value     `json:"value"`
}

type Value struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type NgsReportSpecimen struct {
	ID   string       `json:"id"`
	Type SpecimenType `json:"type"`
}

type MtbPatient struct {
	Address         *Address         `json:"address,omitempty"`
	Age             *ValueWithUnit   `json:"age,omitempty"`
	BirthDate       string           `json:"birthDate"`
	DateOfDeath     *string          `json:"dateOfDeath,omitempty"`
	Gender          CodingGender     `json:"gender"`
	HealthInsurance *HealthInsurance `json:"healthInsurance,omitempty"`
	ID              string           `json:"id"`
	VitalStatus     *VitalStatus     `json:"vitalStatus,omitempty"`
}

type Address struct {
	MunicipalityCode string `json:"municipalityCode"`
}

type ValueWithUnit struct {
	Unit  Unit    `json:"unit"`
	Value float64 `json:"value"`
}

type CodingGender struct {
	Code    Gender  `json:"code"`
	Display *string `json:"display,omitempty"`
	System  *string `json:"system,omitempty"`
	Version *string `json:"version,omitempty"`
}

type HealthInsurance struct {
	Display *string    `json:"display,omitempty"`
	EXTID   ExternalID `json:"extId"`
	Type    *Type      `json:"type,omitempty"`
}

type VitalStatus struct {
	Code    VitalStatusCode `json:"code"`
	Display *string         `json:"display,omitempty"`
	System  *string         `json:"system,omitempty"`
}

type PerformanceStatus struct {
	EffectiveDate string                           `json:"effectiveDate"`
	ID            string                           `json:"id"`
	Patient       GeneticCounsellingRequestPatient `json:"patient"`
	Value         CodingECOG                       `json:"value"`
}

type CodingECOG struct {
	Code    EcogCode `json:"code"`
	Display *string  `json:"display,omitempty"`
	System  *string  `json:"system,omitempty"`
	Version *string  `json:"version,omitempty"`
}

type Response struct {
	EffectiveDate string                           `json:"effectiveDate"`
	ID            string                           `json:"id"`
	Patient       GeneticCounsellingRequestPatient `json:"patient"`
	Therapy       ResponseTherapy                  `json:"therapy"`
	Value         CodingRECIST                     `json:"value"`
}

type ResponseTherapy struct {
	ID   string              `json:"id"`
	Type ResponseTherapyType `json:"type"`
}

type CodingRECIST struct {
	Code    RecistCode `json:"code"`
	Display *string    `json:"display,omitempty"`
	System  *string    `json:"system,omitempty"`
	Version *string    `json:"version,omitempty"`
}

type SpecimenElement struct {
	Collection *Collection                      `json:"collection,omitempty"`
	Diagnosis  *Diagnosis                       `json:"diagnosis,omitempty"`
	ID         string                           `json:"id"`
	Patient    GeneticCounsellingRequestPatient `json:"patient"`
	Type       *CodingTumorSpecimenType         `json:"type,omitempty"`
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

type Diagnosis struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type CodingTumorSpecimenType struct {
	Code    TumorSpecimenType `json:"code"`
	Display *string           `json:"display,omitempty"`
	System  *string           `json:"system,omitempty"`
}

type Therapy struct {
	History []GuidelineTherapyElement `json:"history"`
}

type PatientType string

const (
	Patient PatientType = "Patient"
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

type EXTIDSystem string

const (
	HTTPSPubmedNcbiNlmNihGov EXTIDSystem = "https://pubmed.ncbi.nlm.nih.gov/"
)

type TherapyRecommendationPriority string

const (
	Prio1 TherapyRecommendationPriority = "1"
	Prio2 TherapyRecommendationPriority = "2"
	Prio3 TherapyRecommendationPriority = "3"
	Prio4 TherapyRecommendationPriority = "4"
)

type ClaimResponseClaimType string

const (
	Claim ClaimResponseClaimType = "Claim"
)

type ClaimResponseStatus string

const (
	Accepted                   ClaimResponseStatus = "accepted"
	ClaimResponseStatusUnknown ClaimResponseStatus = "unknown"
	Rejected                   ClaimResponseStatus = "rejected"
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

type StageCode string

const (
	Local         StageCode = "local"
	Metastasized  StageCode = "metastasized"
	PurpleUnknown StageCode = "unknown"
	TumorFree     StageCode = "tumor-free"
)

type TumorGradeCode string

const (
	G1 TumorGradeCode = "G1"
	G2 TumorGradeCode = "G2"
	G3 TumorGradeCode = "G3"
	G4 TumorGradeCode = "G4"
	Gx TumorGradeCode = "GX"
)

type TherapyStatus string

const (
	Completed            TherapyStatus = "completed"
	NotDone              TherapyStatus = "not-done"
	OnGoing              TherapyStatus = "on-going"
	Stopped              TherapyStatus = "stopped"
	TherapyStatusUnknown TherapyStatus = "unknown"
)

type StatusReasonCode string

const (
	ChronicRemission    StatusReasonCode = "chronic-remission"
	CodeOther           StatusReasonCode = "other"
	ContinuedExternally StatusReasonCode = "continued-externally"
	Deterioration       StatusReasonCode = "deterioration"
	FluffyUnknown       StatusReasonCode = "unknown"
	LostToFu            StatusReasonCode = "lost-to-fu"
	MedicalReason       StatusReasonCode = "medical-reason"
	NoIndication        StatusReasonCode = "no-indication"
	OtherTherapyChosen  StatusReasonCode = "other-therapy-chosen"
	PatientDeath        StatusReasonCode = "patient-death"
	PatientRefusal      StatusReasonCode = "patient-refusal"
	PatientWish         StatusReasonCode = "patient-wish"
	PaymentEnded        StatusReasonCode = "payment-ended"
	PaymentPending      StatusReasonCode = "payment-pending"
	PaymentRefused      StatusReasonCode = "payment-refused"
	Progression         StatusReasonCode = "progression"
	Toxicity            StatusReasonCode = "toxicity"
)

type TumorCellContentMethod string

const (
	Bioinformatic TumorCellContentMethod = "bioinformatic"
	Histologic    TumorCellContentMethod = "histologic"
)

type SpecimenType string

const (
	TumorSpecimen SpecimenType = "TumorSpecimen"
)

type ICScoreCode string

const (
	ICScoreCode0 ICScoreCode = "0"
	ICScoreCode1 ICScoreCode = "1"
	ICScoreCode2 ICScoreCode = "2"
	ICScoreCode3 ICScoreCode = "3"
)

type TcScoreCode string

const (
	TcScoreCode0 TcScoreCode = "0"
	TcScoreCode1 TcScoreCode = "1"
	TcScoreCode2 TcScoreCode = "2"
	TcScoreCode3 TcScoreCode = "3"
	TcScoreCode4 TcScoreCode = "4"
	TcScoreCode5 TcScoreCode = "5"
	TcScoreCode6 TcScoreCode = "6"
)

type ProteinExpressionResultCode string

const (
	Exp                              ProteinExpressionResultCode = "exp"
	NotExp                           ProteinExpressionResultCode = "not-exp"
	TentacledUnknown                 ProteinExpressionResultCode = "unknown"
	ProteinExpressionResultCode1Plus ProteinExpressionResultCode = "1+"
	ProteinExpressionResultCode2Plus ProteinExpressionResultCode = "2+"
	ProteinExpressionResultCode3Plus ProteinExpressionResultCode = "3+"
)

type ChromosomeCode string

const (
	Chr1  ChromosomeCode = "chr1"
	Chr10 ChromosomeCode = "chr10"
	Chr11 ChromosomeCode = "chr11"
	Chr12 ChromosomeCode = "chr12"
	Chr13 ChromosomeCode = "chr13"
	Chr14 ChromosomeCode = "chr14"
	Chr15 ChromosomeCode = "chr15"
	Chr16 ChromosomeCode = "chr16"
	Chr17 ChromosomeCode = "chr17"
	Chr18 ChromosomeCode = "chr18"
	Chr19 ChromosomeCode = "chr19"
	Chr2  ChromosomeCode = "chr2"
	Chr20 ChromosomeCode = "chr20"
	Chr21 ChromosomeCode = "chr21"
	Chr22 ChromosomeCode = "chr22"
	Chr3  ChromosomeCode = "chr3"
	Chr4  ChromosomeCode = "chr4"
	Chr5  ChromosomeCode = "chr5"
	Chr6  ChromosomeCode = "chr6"
	Chr7  ChromosomeCode = "chr7"
	Chr8  ChromosomeCode = "chr8"
	Chr9  ChromosomeCode = "chr9"
	ChrX  ChromosomeCode = "chrX"
	ChrY  ChromosomeCode = "chrY"
)

type ChromosomeSystem string

const (
	Chromosome ChromosomeSystem = "chromosome"
)

type GeneSystem string

const (
	HTTPSWWWGenenamesOrg GeneSystem = "https://www.genenames.org/"
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

type Unit string

const (
	Years Unit = "Years"
)

type Gender string

const (
	Female        Gender = "female"
	GenderOther   Gender = "other"
	GenderUnknown Gender = "unknown"
	Male          Gender = "male"
)

type Type string

const (
	Organization Type = "Organization"
)

type VitalStatusCode string

const (
	Alive    VitalStatusCode = "alive"
	Deceased VitalStatusCode = "deceased"
)

type EcogCode string

const (
	EcogCode0 EcogCode = "0"
	EcogCode1 EcogCode = "1"
	EcogCode2 EcogCode = "2"
	EcogCode3 EcogCode = "3"
	EcogCode4 EcogCode = "4"
)

type ResponseTherapyType string

const (
	MTBMedicationTherapy ResponseTherapyType = "MTBMedicationTherapy"
)

type RecistCode string

const (
	CR  RecistCode = "CR"
	Mr  RecistCode = "MR"
	Na  RecistCode = "NA"
	Nya RecistCode = "NYA"
	PD  RecistCode = "PD"
	PR  RecistCode = "PR"
	SD  RecistCode = "SD"
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
