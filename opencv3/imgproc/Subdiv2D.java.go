package imgproc

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const PTLOC_INSIDE = 0
const PTLOC_VERTEX = 1
const PTLOC_ON_EDGE = 2
const NEXT_AROUND_ORG = 0x00
const NEXT_AROUND_DST = 0x22
const PREV_AROUND_ORG = 0x11
const PREV_AROUND_DST = 0x33
const NEXT_AROUND_LEFT = 0x13
const NEXT_AROUND_RIGHT = 0x31
const PREV_AROUND_LEFT = 0x20
const PREV_AROUND_RIGHT = 0x02

const PTLOC_ERROR = -2
const PTLOC_OUTSIDE_RECT = -1

type Subdiv2D struct {
	nativeObj int64
}

func NewSubdiv2D(addr int64) (rcvr *Subdiv2D) {
	rcvr = &Subdiv2D{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewSubdiv2D2(rect *Rect) (rcvr *Subdiv2D) {
	rcvr = &Subdiv2D{}
	rcvr.nativeObj = Subdiv2DNative_Subdiv2D_0(rect.X, rect.Y, rect.Width, rect.Height)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewSubdiv2D3() (rcvr *Subdiv2D) {
	rcvr = &Subdiv2D{}
	rcvr.nativeObj = Subdiv2DNative_Subdiv2D_1()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}

func (rcvr *Subdiv2D) EdgeDst(edge int, dstpt *Point) int {
	dstpt_out := make([]float64, 2)
	retVal := Subdiv2DNative_edgeDst_0(rcvr.GetNativeObjAddr(), edge, dstpt_out)
	if dstpt != nil {
		dstpt.X = dstpt_out[0]
		dstpt.Y = dstpt_out[1]
	}
	return retVal
}
func (rcvr *Subdiv2D) EdgeDst2(edge int) int {
	retVal := Subdiv2DNative_edgeDst_1(rcvr.GetNativeObjAddr(), edge)
	return retVal
}

func (rcvr *Subdiv2D) EdgeOrg(edge int, orgpt *Point) int {
	orgpt_out := make([]float64, 2)
	retVal := Subdiv2DNative_edgeOrg_0(rcvr.GetNativeObjAddr(), edge, orgpt_out)
	if orgpt != nil {
		orgpt.X = orgpt_out[0]
		orgpt.Y = orgpt_out[1]
	}
	return retVal
}
func (rcvr *Subdiv2D) EdgeOrg2(edge int) int {
	retVal := Subdiv2DNative_edgeOrg_1(rcvr.GetNativeObjAddr(), edge)
	return retVal
}

func (rcvr *Subdiv2D) finalize() {
	Subdiv2DNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *Subdiv2D) FindNearest(pt *Point, nearestPt *Point) int {
	nearestPt_out := make([]float64, 2)
	retVal := Subdiv2DNative_findNearest_0(rcvr.GetNativeObjAddr(), pt.X, pt.Y, nearestPt_out)
	if nearestPt != nil {
		nearestPt.X = nearestPt_out[0]
		nearestPt.Y = nearestPt_out[1]
	}
	return retVal
}
func (rcvr *Subdiv2D) FindNearest2(pt *Point) int {
	retVal := Subdiv2DNative_findNearest_1(rcvr.GetNativeObjAddr(), pt.X, pt.Y)
	return retVal
}

func (rcvr *Subdiv2D) GetEdge(edge int, nextEdgeType int) int {
	retVal := Subdiv2DNative_getEdge_0(rcvr.GetNativeObjAddr(), edge, nextEdgeType)
	return retVal
}
func (rcvr *Subdiv2D) GetEdgeList(edgeList *MatOfFloat4) {
	edgeList_mat := edgeList
	Subdiv2DNative_getEdgeList_0(rcvr.GetNativeObjAddr(), edgeList_mat.GetNativeObjAddr())
	return
}

func (rcvr *Subdiv2D) GetLeadingEdgeList(leadingEdgeList *MatOfInt) {
	leadingEdgeList_mat := leadingEdgeList
	Subdiv2DNative_getLeadingEdgeList_0(rcvr.GetNativeObjAddr(), leadingEdgeList_mat.GetNativeObjAddr())
	return
}

func (rcvr *Subdiv2D) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *Subdiv2D) GetTriangleList(triangleList *MatOfFloat6) {
	triangleList_mat := triangleList
	Subdiv2DNative_getTriangleList_0(rcvr.GetNativeObjAddr(), triangleList_mat.GetNativeObjAddr())
	return
}

func (rcvr *Subdiv2D) GetVertex(vertex int, firstEdge []int) *Point {
	firstEdge_out := make([]float64, 1)
	retVal := NewPoint3(Subdiv2DNative_getVertex_0(rcvr.GetNativeObjAddr(), vertex, firstEdge_out))
	if firstEdge != nil {
		firstEdge[0] = int(firstEdge_out[0])
	}
	return retVal
}
func (rcvr *Subdiv2D) GetVertex2(vertex int) *Point {
	retVal := NewPoint3(Subdiv2DNative_getVertex_1(rcvr.GetNativeObjAddr(), vertex))
	return retVal
}

func (rcvr *Subdiv2D) GetVoronoiFacetList(idx *MatOfInt, facetCenters *MatOfPoint2f) (facetList []*MatOfPoint2f) {
	idx_mat := idx
	facetList_mat := NewMat2()
	facetCenters_mat := facetCenters
	Subdiv2DNative_getVoronoiFacetList_0(rcvr.GetNativeObjAddr(), idx_mat.GetNativeObjAddr(), facetList_mat.GetNativeObjAddr(), facetCenters_mat.GetNativeObjAddr())
	facetList = ConvertersToVectorVectorPoint2f(facetList_mat)
	facetList_mat.Release()
	return
}

func (rcvr *Subdiv2D) InitDelaunay(rect *Rect) {
	Subdiv2DNative_initDelaunay_0(rcvr.GetNativeObjAddr(), rect.X, rect.Y, rect.Width, rect.Height)
	return
}

func (rcvr *Subdiv2D) Insert(pt *Point) int {
	retVal := Subdiv2DNative_insert_0(rcvr.GetNativeObjAddr(), pt.X, pt.Y)
	return retVal
}
func (rcvr *Subdiv2D) Insert2(ptvec *MatOfPoint2f) {
	ptvec_mat := ptvec
	Subdiv2DNative_insert_1(rcvr.GetNativeObjAddr(), ptvec_mat.GetNativeObjAddr())
	return
}

func (rcvr *Subdiv2D) Locate(pt *Point, edge []int, vertex []int) int {
	edge_out := make([]float64, 1)
	vertex_out := make([]float64, 1)
	retVal := Subdiv2DNative_locate_0(rcvr.GetNativeObjAddr(), pt.X, pt.Y, edge_out, vertex_out)
	if edge != nil {
		edge[0] = int(edge_out[0])
	}
	if vertex != nil {
		vertex[0] = int(vertex_out[0])
	}
	return retVal
}

func (rcvr *Subdiv2D) NextEdge(edge int) int {
	retVal := Subdiv2DNative_nextEdge_0(rcvr.GetNativeObjAddr(), edge)
	return retVal
}

func (rcvr *Subdiv2D) RotateEdge(edge int, rotate int) int {
	retVal := Subdiv2DNative_rotateEdge_0(rcvr.GetNativeObjAddr(), edge, rotate)
	return retVal
}

func (rcvr *Subdiv2D) SymEdge(edge int) int {
	retVal := Subdiv2DNative_symEdge_0(rcvr.GetNativeObjAddr(), edge)
	return retVal
}
