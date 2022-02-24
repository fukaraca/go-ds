package graphs

import (
	"fmt"
	"sync"
)

//DirectedGraph is a struct for directed graph
type directedGraph struct {
	Vertice map[int]*VertexD
	arcsLen int
	vertLen int
	lock    *sync.Mutex
}

//VertexD represents vertex for directed graph
type VertexD struct {
	Id     int
	Degree int           // total number of arcs that origin from
	Arcs   map[*Arc]bool //true if arc is directed outside
}

type Arc struct {
	From *VertexD //importer side of the arrow
	To   *VertexD //exporter side of the arrow
}

//NewDirected return a new directed graph
func NewDirected() *directedGraph {
	temp := make(map[int]*VertexD)
	return &directedGraph{
		Vertice: temp,
		arcsLen: 0,
		vertLen: 0,
		lock:    &sync.Mutex{},
	}
}

//AddVertex adds new vertex to related directed graph
func (d *directedGraph) AddVertex(id int) error {
	if d.IsExist(id) {
		return fmt.Errorf("vertex for given id %d is already exist", id)
	}
	d.lock.Lock()
	tempArc := make(map[*Arc]bool)
	d.Vertice[id] = &VertexD{
		Id:     id,
		Degree: 0,
		Arcs:   tempArc,
	}
	d.vertLen++
	d.lock.Unlock()
	return nil
}

//AddArc method connects given vertices at given direction
func (d *directedGraph) AddArc(from, to int) error {
	if !d.IsExist(from) || !d.IsExist(to) {
		return fmt.Errorf("at least one of the vertice(%d or/and %d ) is not exist", from, to)
	}

	if ok, _ := d.IsAdjacent(from, to); ok {
		return fmt.Errorf("there is already arc from %d to %d", from, to)
	}
	d.lock.Lock()
	neoArc := &Arc{}
	neoArc.From = d.Vertice[from]
	neoArc.To = d.Vertice[to]
	d.Vertice[from].Arcs[neoArc] = true
	d.Vertice[to].Arcs[neoArc] = false
	d.Vertice[from].Degree++
	d.arcsLen++
	d.lock.Unlock()
	return nil
}

//RemoveVertex removes vertex id if exists
func (d *directedGraph) RemoveVertex(id int) error {
	if !d.IsExist(id) {
		return fmt.Errorf("given id %d is not exist", id)
	}

	d.arcsLen -= d.Vertice[id].Degree
	d.vertLen--
	//handle orphaned arcs
	for arc, ok := range d.Vertice[id].Arcs {
		if ok { //if arc points outward
			delete(d.Vertice[arc.To.Id].Arcs, arc) //delete arc at the adjacent
		} else { //else if arc points inward
			delete(d.Vertice[arc.From.Id].Arcs, arc)
			d.Vertice[arc.From.Id].Degree-- //delete arc from adjacent and decr the degree ad arc count by one
			d.arcsLen--
		}

	}
	delete(d.Vertice, id)
	return nil
}

//RemoveArc removes the given arc
func (d *directedGraph) RemoveArc(from, to int) error {
	if !d.IsExist(from) || !d.IsExist(to) {
		return fmt.Errorf("at least one of the vertice(%d or/and %d ) is not exist", from, to)
	}
	control := false
	temp := &Arc{}
	for k := range d.Vertice[from].Arcs {
		if k.From.Id == from && k.To.Id == to {
			delete(d.Vertice[from].Arcs, k)
			temp = k
			d.Vertice[from].Degree--
			d.arcsLen--
			control = true
			break
		}
	}
	if !control {
		return fmt.Errorf("there is no such arc from %d", from)
	}
	delete(d.Vertice[to].Arcs, temp)

	return nil
}

//IsExist method return true if given vertex id exists
func (d *directedGraph) IsExist(id int) bool {
	if _, ok := d.Vertice[id]; ok {
		return true
	}
	return false
}

//GetVertex returns vertex for given id
func (d *directedGraph) GetVertex(id int) (*VertexD, error) {
	if !d.IsExist(id) {
		return nil, fmt.Errorf("there is no such vertex[id] : %d", id)
	}
	return d.Vertice[id], nil

}

//GetArcs returns arcs that origins from given id vertex
func (d *directedGraph) GetArcs(id int) (map[*Arc]bool, error) {
	if !d.IsExist(id) {
		return nil, fmt.Errorf("id %d is not exist", id)
	}
	return d.Vertice[id].Arcs, nil

}

/*GetAdjacents returns all adjacents of the given vertex as struct. This struct has 2 fields;

Importers: arrow of the arc(edge) points outward.To clarify vertex[id] -> vertex[importer].

Exporters: arrow of the arc(edge) points inward. To clarify vertex[id] <- vertex[exporter] .*/
func (d *directedGraph) GetAdjacents(id int) (*struct {
	Importers []*VertexD //Adjacent list which has a relation FROM given id vertex. To clarify vertex[id] -> vertex[importer]
	Exporters []*VertexD //Adjacent list which has relation TO given id vertex. To clarify vertex[id] <- vertex[exporter]
	Total     int
}, error) {
	if !d.IsExist(id) {
		return nil, fmt.Errorf("there is no such vertex with id:%d", id)
	}
	temp := &struct {
		Importers []*VertexD //Adjacent list which has a relation FROM given id vertex
		Exporters []*VertexD //Adjacent list which has relation TO given id vertex
		Total     int
	}{}
	for k, from := range d.Vertice[id].Arcs {
		if from {
			temp.Importers = append(temp.Importers, k.To)
			temp.Total++
		} else {
			temp.Exporters = append(temp.Exporters, k.From)
			temp.Total++
		}
	}
	return temp, nil
}

//Len returns total number of arcs and vertices
func (d *directedGraph) Len() struct {
	VertLen int
	ArcLen  int
} {
	return struct {
		VertLen int
		ArcLen  int
	}{d.vertLen, d.arcsLen}

}

//NotAdjacentError error occurs when there is no adjacent found
var NotAdjacentError = fmt.Errorf("not adjacent")

/*IsAdjacent returns boolean and error, if any.

In case of id1->id2 returns true and nil,

inc ase of id1<-id2 returns false and nil,

if there is no adjacency between them returns false and graphs.NotAdjacentError error*/
func (d *directedGraph) IsAdjacent(id1, id2 int) (bool, error) {
	if !d.IsExist(id1) || !d.IsExist(id2) {
		return false, fmt.Errorf("at least one of the vertice(%d or/and %d ) is not exist", id1, id2)
	}
	for arc, from := range d.Vertice[id1].Arcs {

		if from {
			if arc.From.Id == id1 && arc.To.Id == id2 {
				return true, nil
			}
		}
	}
	for arc, from := range d.Vertice[id1].Arcs {
		if !from {
			if arc.From.Id == id2 && arc.To.Id == id1 {
				return false, nil
			}
		}
	}
	return false, NotAdjacentError
}

//UndirectedGraph is a struct for undirected graph
type undirectedGraph struct {
	Vertice map[int]*VertexU
	edgeLen int
	vertlen int
	lock    *sync.Mutex
}

type VertexU struct {
	Id     int
	Degree int           // total number of edges
	Edges  map[int]*Edge // map[ adjacent-vertex-id ] edge
}
type Edge struct {
	Adjacents map[int]*VertexU // map[vertex-id]vertex-pointer
}

//NewUndirected creates a new undirected graph and returns
func NewUndirected() *undirectedGraph {
	temp := make(map[int]*VertexU)
	return &undirectedGraph{
		Vertice: temp,
		edgeLen: 0, //total count of edges
		vertlen: 0, //total count of vertices
		lock:    &sync.Mutex{},
	}
}

//AddVertex creates a new vertex
func (g *undirectedGraph) AddVertex(id int) error {
	if g.IsExist(id) {
		return fmt.Errorf("vertex %d couldn't be added because already exist", id)
	}
	g.lock.Lock()
	tempEdges := make(map[int]*Edge)
	g.Vertice[id] = &VertexU{
		Id:     id,
		Degree: 0,
		Edges:  tempEdges,
	}
	g.vertlen++
	g.lock.Unlock()
	return nil
}

//AddEdge creates a new edge and connects to given vertices
func (g *undirectedGraph) AddEdge(id1, id2 int) error {
	if ok, err := g.IsAdjacent(id1, id2); ok || (!ok && err != NotAdjacentError) {
		if ok {
			return fmt.Errorf("%d and %d is already adjacent", id1, id2)
		} else {
			return err
		}
	}
	g.lock.Lock()
	tempEdge := &Edge{Adjacents: make(map[int]*VertexU)}
	tempEdge.Adjacents[id1], tempEdge.Adjacents[id2] = g.Vertice[id1], g.Vertice[id2]
	g.Vertice[id1].Edges[id2] = tempEdge //key is adjacent id value is edge itself
	g.Vertice[id1].Degree++
	g.Vertice[id2].Edges[id1] = tempEdge
	g.Vertice[id2].Degree++
	g.edgeLen++
	g.lock.Unlock()
	return nil
}

//RemoveVertex deletes given vertex and associated edges
func (g *undirectedGraph) RemoveVertex(id int) error {
	if !g.IsExist(id) {
		return fmt.Errorf("vertex couldn't be removed because given id is not exist")
	}
	g.lock.Lock()
	for adj, _ := range g.Vertice[id].Edges {
		delete(g.Vertice[adj].Edges, id)
		g.Vertice[adj].Degree--
		g.edgeLen--
	}
	delete(g.Vertice, id)
	g.vertlen--
	g.lock.Unlock()
	return nil
}

//RemoveEdge deletes given edge from related vertices
func (g *undirectedGraph) RemoveEdge(id1, id2 int) error {
	/*	if !g.IsExist(id1) || !g.IsExist(id2) {
			return fmt.Errorf("at least one of id's is not exist: %d, %d", id1, id2)
		}
	*/ok, err := g.IsAdjacent(id1, id2)
	if !ok {
		return err
	}

	g.lock.Lock()
	delete(g.Vertice[id1].Edges, id2)
	g.Vertice[id1].Degree--
	delete(g.Vertice[id2].Edges, id1)
	g.Vertice[id2].Degree--
	g.edgeLen--
	g.lock.Unlock()

	return nil
}

//GetAdjacents returns adjacent vertice id' s.
func (g *undirectedGraph) GetAdjacents(id int) ([]int, error) {
	if !g.IsExist(id) {
		return nil, fmt.Errorf("list couldn't be gotten because given id is not exist")
	}
	ret := []int{}
	for adj, _ := range g.Vertice[id].Edges {
		ret = append(ret, adj)
	}
	return ret, nil
}

//GetVertex returns given vertex
func (g *undirectedGraph) GetVertex(id int) (*VertexU, error) {
	if !g.IsExist(id) {
		return nil, fmt.Errorf("vertex couldn't be gotten because given id is not exist")
	}
	return g.Vertice[id], nil
}

//GetEdge returns edge that associated given vertices
func (g *undirectedGraph) GetEdge(id1, id2 int) (*Edge, error) {
	ok, err := g.IsAdjacent(id1, id2)
	if !ok {
		return nil, err
	}
	return g.Vertice[id1].Edges[id2], nil
}

/*IsAdjacent returns true and nil if two vertex is adjacent.

if there is no adjacency between them returns false and graphs.NotAdjacentError error*/
func (g *undirectedGraph) IsAdjacent(id1, id2 int) (bool, error) {
	if !g.IsExist(id1) || !g.IsExist(id2) {
		return false, fmt.Errorf("at least one of id's is not exist: %d, %d", id1, id2)
	}
	if _, ok := g.Vertice[id1].Edges[id2]; ok {
		return true, nil
	}
	return false, NotAdjacentError
}

//IsExist method return true if given vertex id exists
func (g *undirectedGraph) IsExist(id int) bool {
	_, ok := g.Vertice[id]
	return ok
}

//Len returns total count of edges and vertices for related undirected graph. For adjacent count use GetAdjacents function.
func (g *undirectedGraph) Len() (ret struct {
	VertexLength int
	EdgeLength   int
}) {
	ret.EdgeLength, ret.VertexLength = g.edgeLen, g.vertlen
	return ret
}

/*
https://en.wikipedia.org/wiki/Graph_(abstract_data_type)
The basic operations provided by a graph data structure G usually include:[1]

//adjacent(G, x, y): tests whether there is an edge from the vertex x to the vertex y;
//neighbors(G, x): lists all vertices y such that there is an edge from the vertex x to the vertex y;
//add_vertex(G, x): adds the vertex x, if it is not there;
//remove_vertex(G, x): removes the vertex x, if it is there;
//add_edge(G, x, y): adds the edge from the vertex x to the vertex y, if it is not there;
//remove_edge(G, x, y): removes the edge from the vertex x to the vertex y, if it is there;
//get_vertex_value(G, x): returns the value associated with the vertex x;
set_vertex_value(G, x, v): sets the value associated with the vertex x to v.
Structures that associate values to the edges usually also provide:[1]

get_edge_value(G, x, y): returns the value associated with the edge (x, y);
set_edge_value(G, x, y, v): sets the value associated with the edge (x, y) to v.
*/
