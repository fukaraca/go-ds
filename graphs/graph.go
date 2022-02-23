package graphs

import (
	"fmt"
)

//DirectedGraph is a struct for directed graph
type directedGraph struct {
	Vertice map[int]*VertexD
	arcsLen int
	vertLen int
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
	}
}

//AddVertex adds new vertex to related directed graph
func (d *directedGraph) AddVertex(id int) error {
	if d.IsExist(id) {
		return fmt.Errorf("vertex for given id %d is already exist", id)
	}
	tempArc := make(map[*Arc]bool)
	d.Vertice[id] = &VertexD{
		Id:     id,
		Degree: 0,
		Arcs:   tempArc,
	}
	d.vertLen++
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
	neoArc := &Arc{}
	neoArc.From = d.Vertice[from]
	neoArc.To = d.Vertice[to]
	d.Vertice[from].Arcs[neoArc] = true
	d.Vertice[to].Arcs[neoArc] = false
	d.Vertice[from].Degree++
	d.arcsLen++
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
	for k, _ := range d.Vertice[from].Arcs {
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

//GetArcs returns arcs that origins from given id vertex
func (d *directedGraph) GetArcs(id int) (map[*Arc]bool, error) {
	if !d.IsExist(id) {
		return nil, fmt.Errorf("id %d is not exist", id)
	}
	return d.Vertice[id].Arcs, nil

}

/*GetAdjacents returns all adjacents of the given vertex as struct. This struct has 2 fields;

importer: arrow of the arc(edge) points outward,

exporter: arrow of the arc(edge) points inward .*/
func (d *directedGraph) GetAdjacents(id int) (*struct {
	Importers []*VertexD //Adjacent list which has a relation FROM given id vertex. To clarify vertex[id] -> vertex[importer]
	Exporters []*VertexD //Adjacent list which has relation TO given id vertex. To clarify vertex[id] <- vertex[exporter]
}, error) {
	if !d.IsExist(id) {
		return nil, fmt.Errorf("there is no such vertex with id:%d", id)
	}
	temp := &struct {
		Importers []*VertexD //Adjacent list which has a relation FROM given id vertex
		Exporters []*VertexD //Adjacent list which has relation TO given id vertex
	}{}
	for k, from := range d.Vertice[id].Arcs {
		fmt.Println(&k.From, from)
		if from {
			temp.Importers = append(temp.Importers, k.To)
		} else {
			temp.Exporters = append(temp.Exporters, k.From)
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

//NotAdjacent error occurs when there is no adjacent found
var NotAdjacent = fmt.Errorf("not adjacent")

/*IsAdjacent returns boolean and error, if any.

In case of id1->id2 returns true and nil,

inc ase of id1<-id2 returns false and nil,

if there is no adjacency between them returns false and graphs.NotAdjacent error*/
func (d *directedGraph) IsAdjacent(id1, id2 int) (bool, error) {
	if !d.IsExist(id1) || !d.IsExist(id2) {
		return false, fmt.Errorf("at least one of the vertice(%d or/and %d ) is not exist", id1, id2)
	}
	for k, _ := range d.Vertice[id1].Arcs {
		if k.From.Id == id1 && k.To.Id == id2 {
			return true, nil
		}
		if k.From.Id == id2 && k.To.Id == id1 {
			return false, nil
		}
	}
	return false, NotAdjacent
}

//UndirectedGraph is a struct for undirected graph
type UndirectedGraph struct {
}

type vertexU struct {
	id     int
	degree int // total number of edges
}
type edge struct {
	adjacents []*vertexU
}

/*
https://en.wikipedia.org/wiki/Graph_(abstract_data_type)
The basic operations provided by a graph data structure G usually include:[1]

adjacent(G, x, y): tests whether there is an edge from the vertex x to the vertex y;
neighbors(G, x): lists all vertices y such that there is an edge from the vertex x to the vertex y;
//add_vertex(G, x): adds the vertex x, if it is not there;
remove_vertex(G, x): removes the vertex x, if it is there;
add_edge(G, x, y): adds the edge from the vertex x to the vertex y, if it is not there;
remove_edge(G, x, y): removes the edge from the vertex x to the vertex y, if it is there;
get_vertex_value(G, x): returns the value associated with the vertex x;
set_vertex_value(G, x, v): sets the value associated with the vertex x to v.
Structures that associate values to the edges usually also provide:[1]

get_edge_value(G, x, y): returns the value associated with the edge (x, y);
set_edge_value(G, x, y, v): sets the value associated with the edge (x, y) to v.
*/
