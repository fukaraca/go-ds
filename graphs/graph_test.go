package graphs

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestNewDirected(t *testing.T) {
	qw := NewDirected()
	assert.NotEqual(t, nil, qw.RemoveArc(5, 6), "there must be no arc to be removed")
	assert.NotEqual(t, nil, qw.RemoveVertex(5), "there must be no vertex to be removed ")
	_, err := qw.GetVertex(5)
	assert.NotEqual(t, nil, err, "error must be returned since there was no actual vertex")
	_, err = qw.GetArcs(5)
	assert.NotEqual(t, nil, err, "error must be returned since there was no actual arc")
	_, err = qw.GetAdjacents(5)
	assert.NotEqual(t, nil, err, "error must be returned since there was no actual vertex")
	_, err = qw.IsAdjacent(5, 6)
	assert.NotEqual(t, nil, err, "error must be returned since there was no actual vertex to be checked for adjacency")
	ok := qw.IsExist(5)
	assert.Equal(t, false, ok, "5 id vertex shouldn't be exist at the moment")

	err = qw.AddVertex(1)
	err = qw.AddVertex(2)
	err = qw.AddVertex(3)
	err = qw.AddVertex(4)
	err = qw.AddVertex(5)
	err = qw.AddVertex(5)
	assert.NotEqual(t, nil, err, "there must be an error for trying to add already exist vertex")
	err = qw.AddVertex(6)
	assert.Equal(t, 6, qw.Len().VertLen, "there must be 6 vertex at the moment")
	err = qw.AddArc(1, 2)
	err = qw.AddArc(1, 3)
	err = qw.AddArc(1, 4)
	err = qw.AddArc(1, 5)
	err = qw.AddArc(2, 5)
	err = qw.AddArc(3, 5)
	err = qw.AddArc(3, 5)
	assert.NotEqual(t, nil, err, "there must be an error for trying to add arc that already exist")
	err = qw.AddArc(4, 5)
	err = qw.AddArc(6, 5)
	assert.Equal(t, 8, qw.Len().ArcLen, "there must be 8 arcs at the moment")
	err = qw.RemoveArc(6, 5)
	err = qw.RemoveArc(4, 5)
	err = qw.RemoveArc(3, 5)
	assert.Equal(t, 5, qw.Len().ArcLen, "there must be 5 arcs at the moment")
	err = qw.AddArc(3, 5)
	err = qw.AddArc(4, 5)
	err = qw.AddArc(6, 5)

	_, err = qw.GetArcs(5)
	assert.Equal(t, nil, err, "it expected to return nil error normally")
	_, err = qw.GetVertex(5)
	assert.Equal(t, nil, err, "it is expected to return nil error normally")
	_, err = qw.GetAdjacents(5)
	assert.Equal(t, nil, err, "it is expected to return nil error normally")
	err = qw.AddVertex(15) //isolated vertex
	_, err = qw.IsAdjacent(15, 5)
	assert.Equal(t, NotAdjacentError, err, "it is expected to return NotAdjacentError")
	err = qw.AddArc(5, 6)

	err = qw.RemoveVertex(5)
	assert.Equal(t, 3, qw.Len().ArcLen, "there must be 3 arcs left at the moment")
	assert.Equal(t, 6, qw.Len().VertLen, "there must be 6 vertex left at the moment")
	err = qw.AddArc(5, 15)
	assert.NotEqual(t, nil, err, "there must be an error since vertex 5 is not exist")
	err = qw.RemoveArc(2, 4)
	assert.NotEqual(t, nil, err, "there must be an error since there is no arc between 2 and 4")
	adj, err := qw.GetAdjacents(1)
	assert.Equal(t, 3, adj.Total, "there must be 3 arcs totally all importer ")
}

func TestNewUndirected(t *testing.T) {
	qw := NewUndirected()
	assert.Equal(t, false, qw.IsExist(5), "it is expected to return, because there is no vertex yet")
	assert.Equal(t, nil, qw.AddVertex(5), "it is expected to return nil error")
	assert.NotEqual(t, nil, qw.AddVertex(5), "it shouldn't return nil error bc vertex-5 is already exist")
	assert.Equal(t, nil, qw.RemoveVertex(5), "it should remove vertex-5 with no problem")
	assert.NotEqual(t, nil, qw.RemoveVertex(5), "it should throw an error says there is no such vertex to remove")
	_, err := qw.GetVertex(5)
	assert.NotEqual(t, nil, err, "it is not expected to return a nil error because requested vertex is not exist")
	for i := 0; i < 6; i++ {
		_ = qw.AddVertex(i)
	}
	//there is 6 isolated vertice
	err = qw.AddEdge(1, 2)
	assert.Equal(t, nil, err, "it is expected to return nil")
	_ = qw.AddEdge(1, 3)
	_ = qw.AddEdge(1, 4)
	_ = qw.AddEdge(1, 5)
	_ = qw.AddEdge(1, 0)
	_ = qw.AddEdge(1, 1)
	err = qw.AddEdge(1, 2) //already exist
	assert.NotEqual(t, nil, err, "it is not expected to return nil because we are trying to add already existed edge")
	//we have 6 vertices and 6 edges. All edges are adjacents of vertex-1
	assert.Equal(t, 6, qw.Len().EdgeLength, "6 edges are expected")
	assert.Equal(t, 6, qw.Len().VertexLength, "6 vertices are expected")
	err = qw.RemoveVertex(1) //this action will orphan all edges connected to vertex-1 so edgelength must be 0
	assert.Equal(t, nil, err, "err must be nil")
	assert.Equal(t, 0, qw.Len().EdgeLength)

	err = qw.RemoveEdge(1, 7)
	assert.NotEqual(t, nil, err, "we are expecting 'not existed vertex' error")

	qw.AddEdge(2, 5)
	err = qw.AddEdge(1, 5)
	assert.NotEqual(t, nil, err, "we excpect an error because vertex-1 is not exist")
	qw.AddEdge(3, 5)
	qw.AddEdge(4, 5)
	err = qw.AddEdge(5, 5)
	assert.Equal(t, nil, err, "it is not expected to get an error here")
	//there must be 4 edges now

	assert.Equal(t, nil, qw.RemoveEdge(5, 5), "we expect no error")
	assert.Equal(t, nil, qw.RemoveEdge(5, 4), "we expect no error")
	assert.Equal(t, 2, qw.Len().EdgeLength, "we expect 2 edges totally")

	//clean-fresh start
	for i := 0; i < 6; i++ {
		qw.RemoveVertex(i)
	}
	//there is 0 vertex and 0 edges now
	assert.Equal(t, 0, qw.Len().VertexLength, "it is expected to return length of vertices 0")

	for i := 0; i < 10; i++ {
		err = qw.AddVertex(i)
		if err != nil {
			t.Error("vertex couldn't be created?")
		}
	}
	//10 vertice

	qw.AddEdge(0, 5)
	qw.AddEdge(2, 5)
	qw.AddEdge(4, 5)
	qw.AddEdge(6, 5)
	qw.AddEdge(8, 5)
	//5 is adjacent with doubles
	adj, _ := qw.GetAdjacents(5)
	sort.Ints(adj)
	assert.Equal(t, []int{0, 2, 4, 6, 8}, adj, "adjacents of vertex 5 must be 0,4,2,6,8")
	_, err = qw.GetAdjacents(15)
	assert.NotEqual(t, nil, err, "since there is no vertex-15, error must be occured")

	_, err = qw.GetVertex(5)
	assert.Equal(t, nil, err, "it is not expected to return non-nil error")
	_, err = qw.GetEdge(5, 15) //vertex-15 is not existed
	assert.NotEqual(t, nil, err, "it is not expected to return nil error")
	_, err = qw.GetEdge(5, 8)
	assert.Equal(t, nil, err, "it is expected to return nil for error")
	ok, err := qw.IsAdjacent(5, 18)
	assert.Equal(t, false, ok, "it is expected to return false error")
	assert.NotEqual(t, NotAdjacentError, err, "it is not expected to return notadjacent also")
	ok, _ = qw.IsAdjacent(5, 8)
	assert.Equal(t, true, ok, "it is expected to return true because 5 and 8 is neighbor")
	ok, err = qw.IsAdjacent(1, 9) //they are not adjacent
	assert.Equal(t, NotAdjacentError, err, "it is expected to return notadjacenterror")

}
