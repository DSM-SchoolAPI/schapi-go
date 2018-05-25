package schapi

type SchoolKind int

const (
	KinderGarten SchoolKind = 1 + iota
	ElementrySchool
	MiddleSchool
	HighSchool
)

type SchoolRegion string

const (
	Seoul     SchoolRegion = "stu.sen.go.kr"
	Incheon                = "stu.ice.go.kr"
	Busan                  = "stu.pen.go.kr"
	Gwangju                = "stu.gen.go.kr"
	Daejeon                = "stu.dje.go.kr"
	Daegu                  = "stu.dge.go.kr"
	Sejong                 = "stu.sje.go.kr"
	Ulsan                  = "stu.use.go.kr"
	Gyeonggi               = "stu.goe.go.kr"
	Kangwon                = "stu.kew.go.kr"
	Chungbuk               = "stu.cbe.go.kr"
	Chungnam               = "stu.cne.go.kr"
	Gyeongbuk              = "stu.gbe.go.kr"
	Gyeongnam              = "stu.gne.go.kr"
	Jeonbuk                = "stu.jbe.go.kr"
	Jeonnam                = "stu.jne.go.kr"
	Jeju                   = "stu.jje.go.kr"
)

type SchoolCode string

type SchoolAPI struct {
	Kind   SchoolKind   `json:"kind"`
	Region SchoolRegion `json:"region"`
	Code   SchoolCode   `json:"code"`
}

func NewSchoolAPI(kind SchoolKind, region SchoolRegion, code SchoolCode) *SchoolAPI {
	return &SchoolAPI{
		Kind: kind,
		Region: region,
		Code: code,
	}
}
